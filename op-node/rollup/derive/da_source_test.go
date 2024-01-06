package derive

import (
	"context"
	"io"
	"math/big"
	"math/rand"
	"testing"

	daclient "github.com/ethereum-optimism/optimism/alt-da/client"
	"github.com/ethereum-optimism/optimism/alt-da/metrics"
	damgr "github.com/ethereum-optimism/optimism/alt-da/mgr"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockFinalitySignal struct {
	mock.Mock
}

func (m *MockFinalitySignal) OnNewL1Finalized(ctx context.Context, blockRef eth.L1BlockRef) {
	m.MethodCalled("OnNewL1Finalized", blockRef)
}

func (m *MockFinalitySignal) ExpectL1Finalized(blockRef eth.L1BlockRef) {
	m.On("OnNewL1Finalized", blockRef).Once()
}

// TestDASource tests the logic in response to all DA service responses.
func TestDASource(t *testing.T) {
	logger := testlog.Logger(t, log.LvlDebug)
	ctx := context.Background()
	storage := daclient.NewMockClient(logger)
	l1F := &testutils.MockL1Source{}

	daCfg := damgr.Config{ChallengeWindow: 90, ResolveWindow: 90}
	da := damgr.NewAltDA(logger, daCfg, &metrics.NoopAltDAMetrics{}, storage, l1F)
	rng := rand.New(rand.NewSource(1234))

	finalitySignal := &MockFinalitySignal{}
	da.OnFinalizedHeadSignal(finalitySignal.OnNewL1Finalized)

	l1Time := uint64(2)
	refA := testutils.RandomBlockRef(rng)
	refA.Number = 1
	l1Refs := []eth.L1BlockRef{refA}

	refA0 := eth.L2BlockRef{
		Hash:           testutils.RandomHash(rng),
		Number:         0,
		ParentHash:     common.Hash{},
		Time:           refA.Time,
		L1Origin:       refA.ID(),
		SequenceNumber: 0,
	}
	batcherPriv := testutils.RandomKey()
	batcherAddr := crypto.PubkeyToAddress(batcherPriv.PublicKey)
	batcherInbox := common.Address{42}
	cfg := &rollup.Config{
		Genesis: rollup.Genesis{
			L1:     refA.ID(),
			L2:     refA0.ID(),
			L2Time: refA0.Time,
		},
		BlockTime:         1,
		SeqWindowSize:     20,
		BatchInboxAddress: batcherInbox,
	}

	var inputs [][]byte
	var comms [][]byte

	signer := cfg.L1Signer()
	factory := NewDASourceFactory(logger, cfg, l1F, da)

	c := 0

	for i := uint64(0); i <= daCfg.ChallengeWindow+daCfg.ResolveWindow; i++ {
		parent := l1Refs[len(l1Refs)-1]
		// create a new mock l1 ref
		ref := eth.L1BlockRef{
			Hash:       testutils.RandomHash(rng),
			Number:     parent.Number + 1,
			ParentHash: parent.Hash,
			Time:       parent.Time + l1Time,
		}
		l1Refs = append(l1Refs, ref)
		logger.Info("new l1 block", "ref", ref)
		// called for each l1 block to sync challenges
		l1F.ExpectL1BlockRefByNumber(ref.Number, ref, nil)
		l1F.ExpectFetchReceipts(ref.Hash, nil, types.Receipts{}, nil)

		// add an input commitment every 6 l1 blocks
		hasComm := i%6 == 0
		if hasComm {
			input := testutils.RandomData(rng, 2000)
			comm, _ := storage.SetPreImage(ctx, input)
			inputs = append(inputs, input)
			comms = append(comms, comm)

			logger.Info("submitting batch", "comm", comm, "block", ref.Number)
			tx, err := types.SignNewTx(batcherPriv, signer, &types.DynamicFeeTx{
				ChainID:   signer.ChainID(),
				Nonce:     0,
				GasTipCap: big.NewInt(2 * params.GWei),
				GasFeeCap: big.NewInt(30 * params.GWei),
				Gas:       100_000,
				To:        &batcherInbox,
				Value:     big.NewInt(int64(0)),
				Data:      comm,
			})
			require.NoError(t, err)

			txs := []*types.Transaction{tx}
			// called once per derivation
			l1F.ExpectInfoAndTxsByHash(ref.Hash, testutils.RandomBlockInfo(rng), txs, nil)
			l1F.ExpectInfoAndTxsByHash(ref.Hash, testutils.RandomBlockInfo(rng), txs, nil)

			finalitySignal.ExpectL1Finalized(ref)
		} else {
			l1F.ExpectInfoAndTxsByHash(ref.Hash, testutils.RandomBlockInfo(rng), []*types.Transaction{}, nil)
			l1F.ExpectInfoAndTxsByHash(ref.Hash, testutils.RandomBlockInfo(rng), []*types.Transaction{}, nil)
			// challenge the first 4 commitments as soon as we have collected them all
			if len(comms) >= 4 && c < 7 {
				// skip a block between each challenge transaction
				if c%2 == 0 {
					da.State().SetActiveChallenge(comms[c/2], ref.Number, daCfg.ResolveWindow)
					logger.Info("setting active challenge", "comm", comms[c/2])
				}
				c++
			}
		}

		src := factory.OpenData(ctx, ref.ID(), batcherAddr)
		data, err := src.Next(ctx)
		// our first challenge expires
		if i == daCfg.ResolveWindow+19 {
			require.ErrorIs(t, err, ErrReset)
			break
		}
		if !hasComm {
			require.ErrorIs(t, err, io.EOF)
		} else {
			// check that each commitment is resolved
			require.NoError(t, err)
			require.Equal(t, hexutil.Bytes(inputs[len(inputs)-1]), data)
		}
	}

	logger.Info("pipeline reset ..................................")

	c = 1

	// restart derivation from the last finalized head
	for i := 1; i < len(l1Refs)+2; i++ {

		var ref eth.L1BlockRef
		// first we run through all the existing l1 blocks
		if i < len(l1Refs) {
			ref = l1Refs[i]
			logger.Info("re deriving block", "ref", ref, "i", i)

			if i == len(l1Refs)-1 {
				l1F.ExpectFetchReceipts(ref.Hash, nil, types.Receipts{}, nil)
			}
			// once past the l1 head, continue generating new l1 refs
		} else {
			parent := l1Refs[len(l1Refs)-1]
			ref = eth.L1BlockRef{
				Hash:       testutils.RandomHash(rng),
				Number:     parent.Number + 1,
				ParentHash: parent.Hash,
				Time:       parent.Time + l1Time,
			}
			l1Refs = append(l1Refs, ref)
			logger.Info("new l1 block", "ref", ref)
			// called for each l1 block to sync challenges
			l1F.ExpectL1BlockRefByNumber(ref.Number, ref, nil)
			l1F.ExpectFetchReceipts(ref.Hash, nil, types.Receipts{}, nil)

			l1F.ExpectInfoAndTxsByHash(ref.Hash, testutils.RandomBlockInfo(rng), []*types.Transaction{}, nil)
			l1F.ExpectInfoAndTxsByHash(ref.Hash, testutils.RandomBlockInfo(rng), []*types.Transaction{}, nil)
		}

		src := factory.OpenData(ctx, ref.ID(), batcherAddr)
		data, err := src.Next(ctx)

		// the next challenge expires
		if uint64(i) == daCfg.ResolveWindow+22 {
			require.ErrorIs(t, err, ErrReset)
			break
		}

		hasComm := (i-1)%6 == 0 && i != 1
		if !hasComm {
			require.ErrorIs(t, err, io.EOF)
		} else {
			require.NoError(t, err)
			require.Equal(t, hexutil.Bytes(inputs[c]), data)
			c++
		}
	}
}
