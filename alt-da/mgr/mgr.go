package mgr

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/ethereum-optimism/optimism/alt-da/api"
	"github.com/ethereum-optimism/optimism/alt-da/client"
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

// challenge object stores the status and if active the block where it was activated.
type challenge struct {
	status     api.ChallengeStatusCode
	input      []byte
	startBlock uint64
}

// l1Origin keeps track of all input commitments included in the l1 block, their challenge states
// and a ref to the last l2 block derived from that epoch.
type l1Origin struct {
	challenges map[string]challenge
	head       *eth.L2BlockRef
}

func newL1Origin() *l1Origin {
	return &l1Origin{
		challenges: make(map[string]challenge),
	}
}

// Keeps track of the l2 blocks for each L1 origin in the reorg window.
type L2Blocks struct {
	byL1Origin map[uint64]*l1Origin
	windowSize uint64
}

func NewL2Blocks(windowSize uint64) *L2Blocks {
	return &L2Blocks{
		byL1Origin: make(map[uint64]*l1Origin),
		windowSize: windowSize,
	}
}

func (b *L2Blocks) getOrNewOrigin(number uint64) *l1Origin {
	_, ok := b.byL1Origin[number]
	if !ok {
		b.byL1Origin[number] = newL1Origin()
	}
	return b.byL1Origin[number]
}

// AddBlock adds a new l2 block ref only if greater than the current head.
func (b *L2Blocks) AddBlock(block eth.L2BlockRef) {
	o := b.getOrNewOrigin(block.L1Origin.Number)
	if o.head == nil || block.Number >= o.head.Number {
		o.head = &block
	}
}

// As the challenge window advances, we pop the oldest l1 block and return the last l2 block in the
// sequencing window.
func (b *L2Blocks) AdvanceWindow(l1Blk eth.BlockID) (*eth.L2BlockRef, bool, error) {
	b.getOrNewOrigin(l1Blk.Number)
	if len(b.byL1Origin) > int(b.windowSize) {
		expNum := l1Blk.Number - b.windowSize
		e, ok := b.byL1Origin[expNum]
		if !ok {
			return nil, false, fmt.Errorf("epoch not found")
		}
		expired := 0
		for k, v := range e.challenges {
			if v.status == api.ChallengeActive {
				expired++
				b.SetExpiredChallenge([]byte(k), expNum)
			}
		}
		// cleanup previous one, as past new safe head
		if expNum > 0 {
			delete(b.byL1Origin, expNum-1)
		}
		return e.head, expired > 0, nil
	}
	return nil, false, fmt.Errorf("window not full")
}

// SetExpiredChallenge sets a challenge status as expired.
func (b *L2Blocks) SetExpiredChallenge(key []byte, blockNumber uint64) {
	o := b.getOrNewOrigin(blockNumber)
	o.challenges[string(key)] = challenge{status: api.ChallengeExpired}
}

// SetActiveChallenge sets the challenge status to active and the block where it was activated.
func (b *L2Blocks) SetActiveChallenge(key []byte, blockNumber uint64, start uint64) {
	o := b.getOrNewOrigin(blockNumber)
	chall := o.challenges[string(key)]
	// do not override an expired challenge
	if chall.status == api.ChallengeResolved || chall.status == api.ChallengeExpired {
		return
	}
	o.challenges[string(key)] = challenge{status: api.ChallengeActive, startBlock: start}
}

// SetResolvedChallenge sets the challenge status and the input data that resolved it.
func (b *L2Blocks) SetResolvedChallenge(key []byte, input []byte, blockNumber uint64) {
	o := b.getOrNewOrigin(blockNumber)
	o.challenges[string(key)] = challenge{status: api.ChallengeResolved, input: input}
}

// GetChallenge return the challenge state and whether it already exists.
func (b *L2Blocks) GetChallenge(key []byte, blockNumber uint64) (challenge, bool) {
	o := b.getOrNewOrigin(blockNumber)
	chall, ok := o.challenges[string(key)]
	return chall, ok
}

// GetOrTrackChallenge returns the challenge state or sets the key with a default value to track it.
func (b *L2Blocks) GetOrTrackChallenge(key []byte, blockNumber uint64) challenge {
	o := b.getOrNewOrigin(blockNumber)
	_, ok := o.challenges[string(key)]
	if !ok {
		o.challenges[string(key)] = challenge{}
	}
	return o.challenges[string(key)]
}

// GetResolvedInput returns the input data that resolved the challenge.
func (b *L2Blocks) GetResolvedInput(key []byte, blockNumber uint64) []byte {
	if _, ok := b.byL1Origin[blockNumber]; !ok {
		return nil
	}
	chall, _ := b.byL1Origin[blockNumber].challenges[string(key)]
	return chall.input
}

// L1Fetcher is the required interface for syncing the DA challenge contract state.
type L1Fetcher interface {
	InfoAndTxsByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, types.Transactions, error)
	FetchReceipts(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, types.Receipts, error)
	L1BlockRefByNumber(context.Context, uint64) (eth.L1BlockRef, error)
}

// L2Engine is the required interface for updating the L2 engine safe head.
type L2Engine interface {
	SetSafeHead(ctx context.Context, safeHash common.Hash) error
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
	log log.Logger
	cfg Config

	storageClient PreImageFetcher
	engineClient  L2Engine

	blocks *L2Blocks

	// the latest l1 block we synced challenge contract events from
	challHead eth.BlockID
}

func NewAltDA(log log.Logger, cfg Config, storage PreImageFetcher, engine L2Engine) *AltDA {
	reorgWindow := cfg.ChallengeWindow + cfg.ResolveWindow
	return &AltDA{
		log:           log,
		cfg:           cfg,
		storageClient: storage,
		engineClient:  engine,
		blocks:        NewL2Blocks(reorgWindow),
	}
}

// StepChallenges continues syncing the challenges by fetching the next l1 block.
func (a *AltDA) StepChallenges(ctx context.Context, l1 L1Fetcher) error {
	blkRef, err := l1.L1BlockRefByNumber(ctx, a.challHead.Number+1)
	if err != nil {
		a.log.Error("failed to fetch l1 head", "err", err)
		return err
	}
	blkId := blkRef.ID()
	return a.UpdateChallenges(ctx, blkId, l1)
}

// AdvanceWindow advances the challenge + reorg window and returns whether we need to reorg.
func (a *AltDA) AdvanceWindow(ctx context.Context, l1Head eth.BlockID, l1 L1Fetcher) (bool, error) {
	safeHead, reset, _ := a.blocks.AdvanceWindow(l1Head)
	if safeHead != nil {
		a.log.Info("setting safe head", "safeHead", safeHead, "origin", safeHead.L1Origin.Number)
		if err := a.engineClient.SetSafeHead(ctx, safeHead.Hash); err != nil {
			return false, err
		}
	}
	return reset, nil
}

// UpdateChallenges fetches the l1 block receipts and updates the challenge status
func (a *AltDA) UpdateChallenges(ctx context.Context, block eth.BlockID, l1 L1Fetcher) error {
	// noop if we already processed this block
	if block.Number <= a.challHead.Number {
		return nil
	}
	//cached with deposits events call so not expensive
	_, receipts, err := l1.FetchReceipts(ctx, block.Hash)
	if err != nil {
		return err
	}
	a.log.Info("updating challenges", "epoch", block.Number, "reorgWindow", a.blocks.windowSize, "numReceipts", len(receipts))
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
				_, has := a.blocks.GetChallenge(hash, bn)
				// if we are not tracking the commitment from processing the l1 origin in derivation, this challenge is invalid.
				if !has {
					a.log.Warn("skipping invalid challenge", "block", bn)
					continue
				}
				status := api.ChallengeStatusCode(event.Status)
				if status == api.ChallengeResolved {
					// cached with input resolution call so not expensive
					_, txs, err := l1.InfoAndTxsByHash(ctx, block.Hash)
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
						a.blocks.SetResolvedChallenge(hash, input, bn)
					}
				}
				if status == api.ChallengeActive {
					a.blocks.SetActiveChallenge(hash, bn, log.BlockNumber)
				}
			}
		}

	}
	a.challHead = block
	return nil
}

// GetPreImage combines results from the storage client and the challenge contract to inform the derivation pipeline on what to do.
func (a *AltDA) GetPreImage(ctx context.Context, key []byte, blockNumber uint64) (*api.Response, error) {
	ch := a.blocks.GetOrTrackChallenge(key, blockNumber)
	val, err := a.storageClient.GetPreImage(ctx, key)
	res := api.NewResponse(api.Available, val)
	// data is not found in storage or may be available but challenge was
	// still expired so it needs to be skipped.
	if errors.Is(client.ErrNotFound, err) || ch.status == api.ChallengeExpired {
		switch ch.status {
		case api.ChallengeActive:
			if a.isInResolveWindow(blockNumber) {
				res.SetPendingChallenge()
			} else {
				res.SetExpiredChallenge()
			}
		case api.ChallengeExpired:
			res.SetExpiredChallenge()
		case api.ChallengeResolved:
			resolvedInput := a.blocks.GetResolvedInput(key, blockNumber)
			if resolvedInput != nil {
				res.SetResolved(resolvedInput)
			}
		default:
			// If we can still challenge the data, set as pending a keep retrying
			// while the challenge head steps forward.
			if a.isInChallengeWindow(blockNumber) {
				res.SetPendingChallenge()
			} else {
				res.SetFatal()
			}
		}

		a.log.Error("preimage not available", "status", res.Status)
		return res, nil
	}

	if err != nil {
		a.log.Error("failed to get preimage", "err", err)
		// the storage client request failed for some other reason
		// in which case derivation pipeline should be retried
		return nil, err
	}

	return res, nil
}

// the DA service must keep track of every new subjective safe head
// so as to mark the new safe head when the challenge window expires
func (a *AltDA) SetSubjectiveSafeHead(ctx context.Context, ref eth.L2BlockRef) error {
	a.log.Debug("setting subjective safe head", "ref", ref, "origin", ref.L1Origin)
	a.blocks.AddBlock(ref)
	return nil
}

// GetChallengeStatus returns the status of the challenge for the given key and block number and if it is
// already tracked (i.e. the input data was previously requested during derivation).
func (a *AltDA) GetChallengeStatus(key []byte, blockNumber uint64) (api.ChallengeStatusCode, bool) {
	ch, ok := a.blocks.GetChallenge(key, blockNumber)
	return ch.status, ok
}

// checks if the given block number is in the challenge window based on the latest challenge head.
func (a *AltDA) isInChallengeWindow(blockNumber uint64) bool {
	return a.challHead.Number >= blockNumber && a.challHead.Number <= blockNumber+a.cfg.ChallengeWindow
}

// checks if the given block is in the resolve window based on the latest challenge head.
func (a *AltDA) isInResolveWindow(blockNumber uint64) bool {
	return a.challHead.Number >= blockNumber && a.challHead.Number <= blockNumber+a.cfg.ChallengeWindow+a.cfg.ResolveWindow
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
