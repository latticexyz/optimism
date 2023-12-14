package mgr

import (
	"bytes"
	"context"
	"errors"

	"github.com/ethereum-optimism/optimism/alt-da/api"
	"github.com/ethereum-optimism/optimism/alt-da/client"
	"github.com/ethereum-optimism/optimism/alt-da/metrics"
	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

var (
	ChallengeStatusEventName    = "ChallengeStatusChanged"
	ChallengeStatusEventABI     = "ChallengeStatusChanged(bytes32,uint256,uint8)"
	ChallengeStatusEventABIHash = crypto.Keccak256Hash([]byte(ChallengeStatusEventABI))
)

// L1Fetcher is the required interface for syncing the DA challenge contract state.
type L1Fetcher interface {
	InfoAndTxsByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, types.Transactions, error)
	FetchReceipts(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, types.Receipts, error)
	L1BlockRefByNumber(context.Context, uint64) (eth.L1BlockRef, error)
}

// PreImageFetcher is the required interface for fetching inputs from the storage layer.
type PreImageFetcher interface {
	GetPreImage(ctx context.Context, hash []byte) ([]byte, error)
}

type Config struct {
	// Required for filtering contract events
	DaChallengeContractAddress common.Address
	SequencerWindow            uint64
	// The number of l1 blocks after the input is committed during which one can challenge.
	ChallengeWindow uint64
	// The number of l1 blocks after a commitment is challenged during which one can resolve.
	ResolveWindow uint64
}

// AltDA keeps track of data availability and the challenge contract to determine
// when to update the safe head.
type AltDA struct {
	log     log.Logger
	cfg     Config
	metrics metrics.AltDAMetricer

	storageClient PreImageFetcher
	l1Client      L1Fetcher

	state *State

	// the latest l1 block we synced challenge contract events from
	challHead     eth.BlockID
	finalizedHead eth.L1BlockRef

	finalizedHeadSignalFunc eth.HeadSignalFn
}

func NewAltDA(log log.Logger, cfg Config, metrics metrics.AltDAMetricer, storage PreImageFetcher, l1F L1Fetcher) *AltDA {
	return &AltDA{
		log:           log,
		cfg:           cfg,
		metrics:       metrics,
		storageClient: storage,
		l1Client:      l1F,
		state:         NewState(log),
	}
}

// OnFinalizedHeadSignal sets the callback function to be called when the finalized head is updated.
// This will signal to the engine queue that will set the proper L2 block as finalized.
func (a *AltDA) OnFinalizedHeadSignal(f eth.HeadSignalFn) {
	a.finalizedHeadSignalFunc = f
}

// ChallengesHead returns the latest l1 block synced from the challenge contract.
func (a *AltDA) ChallengesHead() eth.BlockID {
	return a.challHead
}

// LoadNextChallenges increments the challenges head and process the new block if it exists.
// It is only used if the derivation pipeline stalls and we need to wait for a challenge to be resolved
// to get the next input.
func (a *AltDA) LoadNextChallenges(ctx context.Context) error {
	blkRef, err := a.l1Client.L1BlockRefByNumber(ctx, a.challHead.Number+1)
	if err != nil {
		a.log.Error("failed to fetch l1 head", "err", err)
		return err
	}
	return a.ProcessL1Origin(ctx, blkRef.ID())
}

func (a *AltDA) State() *State {
	return a.state
}

// ProcessL1Origin syncs any challenge events included in the l1 block, expires any active challenges
// after the new resolveWindow, computes and signals the new finalized head and sets the l1 block
// as the new head for tracking challenges. If forwards an error if any new challenge have expired to
// trigger a derivation reset.
func (a *AltDA) ProcessL1Origin(ctx context.Context, block eth.BlockID) error {
	// do not repeat for the same origin
	if block.Number <= a.challHead.Number {
		return nil
	}
	// sync challenges for the new block
	if err := a.LoadChallengeEvents(ctx, block); err != nil {
		return err
	}
	// advance challenge window
	bn, err := a.state.ExpireChallenges(block.Number)
	if err != nil {
		return err
	}

	if bn > a.finalizedHead.Number {
		ref, err := a.l1Client.L1BlockRefByNumber(ctx, bn)
		if err != nil {
			return err
		}
		a.finalizedHead = ref

		// if we get a greater finalized head, signal to the engine queue
		if a.finalizedHeadSignalFunc != nil {
			a.finalizedHeadSignalFunc(ctx, a.finalizedHead)

		}
		// prune old state
		a.state.Prune(bn)

	}
	a.challHead = block
	return nil
}

// LoadChallengeEvents fetches the l1 block receipts and updates the challenge status
func (a *AltDA) LoadChallengeEvents(ctx context.Context, block eth.BlockID) error {
	//cached with deposits events call so not expensive
	_, receipts, err := a.l1Client.FetchReceipts(ctx, block.Hash)
	if err != nil {
		return err
	}
	a.log.Info("updating challenges", "epoch", block.Number, "numReceipts", len(receipts))
	for i, rec := range receipts {
		if rec.Status != types.ReceiptStatusSuccessful {
			continue
		}
		for j, log := range rec.Logs {
			if log.Address == a.cfg.DaChallengeContractAddress && len(log.Topics) > 0 && log.Topics[0] == ChallengeStatusEventABIHash {
				event, err := DecodeChallengeStatusEvent(log)
				if err != nil {
					a.log.Error("failed to decode challenge event", "block", block.Number, "tx", i, "log", j, "err", err)
					continue
				}
				a.log.Info("found challenge event", "block", block.Number, "tx", i, "log", j, "event", event)
				hash := event.ChallengedHash[:]
				bn := event.ChallengedBlockNumber.Uint64()
				// if we are not tracking the commitment from processing the l1 origin in derivation, this challenge is invalid.
				if !a.state.IsTracking(hash, bn) {
					a.log.Warn("skipping invalid challenge", "block", bn)
					continue
				}
				status := api.ChallengeStatusCode(event.Status)
				if status == api.ChallengeResolved {
					// cached with input resolution call so not expensive
					_, txs, err := a.l1Client.InfoAndTxsByHash(ctx, block.Hash)
					if err != nil {
						a.log.Error("failed to fetch l1 block", "block", block.Number, "err", err)
						continue
					}
					tx := txs[i]
					// txs and receipts should be in the same order
					if tx.Hash() != rec.TxHash {
						a.log.Error("tx hash mismatch", "block", block.Number, "tx", i, "log", j, "txHash", tx.Hash(), "receiptTxHash", rec.TxHash)
						continue
					}
					input, err := DecodeResolvedInput(tx.Data())
					if err != nil {
						a.log.Error("failed to decode resolved input", "block", block.Number, "tx", i, "err", err)
						continue
					}
					if input != nil && bytes.Equal(crypto.Keccak256(input), event.ChallengedHash[:]) {
						a.log.Debug("resolved input", "block", block.Number, "tx", i)
						a.state.SetResolvedChallenge(hash, input, log.BlockNumber)
					}
				}
				if status == api.ChallengeActive {
					a.state.SetActiveChallenge(hash, log.BlockNumber, a.cfg.ResolveWindow)
				}
			}
		}

	}
	return nil
}

// GetPreImage combines results from the storage client and the challenge contract to inform the derivation pipeline on what to do.
func (a *AltDA) GetPreImage(ctx context.Context, key []byte, blockNumber uint64) (*api.Response, error) {
	// If the challenge head is ahead in the case of a pipeline reset or stall, we might have synced a
	// challenge event for this commitment. Otherwise we mark the commitment as part of the cannonical
	// chain so potential future challenge events can be selected.
	ch := a.state.GetOrTrackChallenge(key, blockNumber, a.cfg.ChallengeWindow)
	val, err := a.storageClient.GetPreImage(ctx, key)
	res := api.NewResponse(api.Available, val)
	// data is not found in storage or may be available but challenge was
	// still expired so it needs to be skipped.
	notFound := errors.Is(client.ErrNotFound, err)
	switch ch.challengeStatus {
	case api.ChallengeActive:
		if a.isExpired(ch.expiresAt) {
			res.SetExpiredChallenge()
		} else if notFound {
			res.SetPendingChallenge()
		}
	case api.ChallengeExpired:
		res.SetExpiredChallenge()
	case api.ChallengeResolved:
		resolvedInput, err := a.state.GetResolvedInput(key)
		if err == nil {
			res.SetResolved(resolvedInput)
		}
	default:
		if notFound {
			// If we can still challenge the data, set as pending a keep retrying
			// while the challenge head steps forward.
			if a.isExpired(ch.expiresAt) {
				res.SetFatal()
			} else {
				res.SetPendingChallenge()
			}
		}
	}

	if notFound {
		a.log.Error("preimage not available", "status", res.Status)
	}

	if err != nil && !notFound {
		a.log.Error("failed to get preimage", "err", err)
		// the storage client request failed for some other reason
		// in which case derivation pipeline should be retried
		return nil, err
	}

	return res, nil
}

// isExpired returns whether the expiration block is lower or equal to the current head
func (a *AltDA) isExpired(bn uint64) bool {
	return a.challHead.Number >= bn
}

// DecodeChallengeStatusEvent decodes the challenge status event from the log data and the indexed challenged
// hash and block number from the topics.
func DecodeChallengeStatusEvent(log *types.Log) (*bindings.DataAvailabilityChallengeChallengeStatusChanged, error) {
	// abi lazy loaded
	dacAbi, _ := bindings.DataAvailabilityChallengeMetaData.GetAbi()
	var event bindings.DataAvailabilityChallengeChallengeStatusChanged
	err := dacAbi.UnpackIntoInterface(&event, ChallengeStatusEventName, log.Data)
	if err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range dacAbi.Events[ChallengeStatusEventName].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(&event, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	return &event, nil
}

// DecodeResolvedInput decodes the preimage bytes from the tx input data.
func DecodeResolvedInput(data []byte) ([]byte, error) {
	dacAbi, _ := bindings.DataAvailabilityChallengeMetaData.GetAbi()

	args := make(map[string]interface{})
	err := dacAbi.Methods["resolve"].Inputs.UnpackIntoMap(args, data[4:])
	if err != nil {
		return nil, err
	}
	return args["preImage"].([]byte), nil
}
