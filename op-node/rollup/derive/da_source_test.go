package derive

import (
	"context"
	"io"
	"math/big"
	"math/rand"
	"testing"

	daapi "github.com/ethereum-optimism/optimism/alt-da/api"
	daclient "github.com/ethereum-optimism/optimism/alt-da/client"
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
	"github.com/stretchr/testify/require"
)

// TestDASource tests the logic in response to all DA service responses.
func TestDASource(t *testing.T) {
	logger := testlog.Logger(t, log.LvlCrit)
	ctx := context.Background()
	storage := daclient.NewMockClient(logger)
	engine := &testutils.MockEngine{}
	l1F := &testutils.MockL1Source{}

	daCfg := damgr.Config{ChallengeWindow: 4}
	da := damgr.NewAltDA(logger, daCfg, storage, engine)
	rng := rand.New(rand.NewSource(1234))

	l1Time := uint64(2)
	refA := testutils.RandomBlockRef(rng)
	refB := eth.L1BlockRef{
		Hash:       testutils.RandomHash(rng),
		Number:     refA.Number + 1,
		ParentHash: refA.Hash,
		Time:       refA.Time + l1Time,
	}
	refC := eth.L1BlockRef{
		Hash:       testutils.RandomHash(rng),
		Number:     refB.Number + 1,
		ParentHash: refB.Hash,
		Time:       refB.Time + l1Time,
	}
	refD := eth.L1BlockRef{
		Hash:       testutils.RandomHash(rng),
		Number:     refC.Number + 1,
		ParentHash: refC.Hash,
		Time:       refC.Time + l1Time,
	}
	refE := eth.L1BlockRef{
		Hash:       testutils.RandomHash(rng),
		Number:     refD.Number + 1,
		ParentHash: refD.Hash,
		Time:       refD.Time + l1Time,
	}
	refF := eth.L1BlockRef{
		Hash:       testutils.RandomHash(rng),
		Number:     refE.Number + 1,
		ParentHash: refE.Hash,
		Time:       refE.Time + l1Time,
	}
	l1Refs := []eth.L1BlockRef{refA, refB, refC, refD, refE, refF}

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
		SeqWindowSize:     2,
		BatchInboxAddress: batcherInbox,
	}
	refA1 := eth.L2BlockRef{
		Hash:           testutils.RandomHash(rng),
		Number:         refA0.Number + 1,
		ParentHash:     refA0.Hash,
		Time:           refA0.Time + cfg.BlockTime,
		L1Origin:       refA.ID(),
		SequenceNumber: 1,
	}
	refB0 := eth.L2BlockRef{
		Hash:           testutils.RandomHash(rng),
		Number:         refA1.Number + 1,
		ParentHash:     refA1.Hash,
		Time:           refA1.Time + cfg.BlockTime,
		L1Origin:       refB.ID(),
		SequenceNumber: 0,
	}
	refB1 := eth.L2BlockRef{
		Hash:           testutils.RandomHash(rng),
		Number:         refB0.Number + 1,
		ParentHash:     refB0.Hash,
		Time:           refB0.Time + cfg.BlockTime,
		L1Origin:       refB.ID(),
		SequenceNumber: 1,
	}
	refC0 := eth.L2BlockRef{
		Hash:           testutils.RandomHash(rng),
		Number:         refB1.Number + 1,
		ParentHash:     refB1.Hash,
		Time:           refB1.Time + cfg.BlockTime,
		L1Origin:       refC.ID(),
		SequenceNumber: 0,
	}
	refC1 := eth.L2BlockRef{
		Hash:           testutils.RandomHash(rng),
		Number:         refC0.Number + 1,
		ParentHash:     refC0.Hash,
		Time:           refC0.Time + cfg.BlockTime,
		L1Origin:       refC.ID(),
		SequenceNumber: 1,
	}
	refD0 := eth.L2BlockRef{
		Hash:           testutils.RandomHash(rng),
		Number:         refC1.Number + 1,
		ParentHash:     refC1.Hash,
		Time:           refC1.Time + cfg.BlockTime,
		L1Origin:       refD.ID(),
		SequenceNumber: 0,
	}
	refD1 := eth.L2BlockRef{
		Hash:           testutils.RandomHash(rng),
		Number:         refD0.Number + 1,
		ParentHash:     refD0.Hash,
		Time:           refD0.Time + cfg.BlockTime,
		L1Origin:       refD.ID(),
		SequenceNumber: 1,
	}
	refE0 := eth.L2BlockRef{
		Hash:           testutils.RandomHash(rng),
		Number:         refD1.Number + 1,
		ParentHash:     refD1.Hash,
		Time:           refD1.Time + cfg.BlockTime,
		L1Origin:       refE.ID(),
		SequenceNumber: 0,
	}
	refE1 := eth.L2BlockRef{
		Hash:           testutils.RandomHash(rng),
		Number:         refE0.Number + 1,
		ParentHash:     refE0.Hash,
		Time:           refE0.Time + cfg.BlockTime,
		L1Origin:       refE.ID(),
		SequenceNumber: 1,
	}
	refF0 := eth.L2BlockRef{
		Hash:           testutils.RandomHash(rng),
		Number:         refE1.Number + 1,
		ParentHash:     refE1.Hash,
		Time:           refE1.Time + cfg.BlockTime,
		L1Origin:       refF.ID(),
		SequenceNumber: 0,
	}
	refF1 := eth.L2BlockRef{
		Hash:           testutils.RandomHash(rng),
		Number:         refF0.Number + 1,
		ParentHash:     refF0.Hash,
		Time:           refF0.Time + cfg.BlockTime,
		L1Origin:       refF.ID(),
		SequenceNumber: 1,
	}
	l2Refs := []eth.L2BlockRef{refA0, refA1, refB0, refB1, refC0, refC1, refD0, refD1, refE0, refE1, refF0, refF1}

	for i, ref := range l2Refs {
		t.Log("block", i, "hash", ref.Hash)
	}

	inputs := make([][]byte, len(l1Refs))
	comms := make([][]byte, len(l1Refs))

	signer := cfg.L1Signer()

	for i, ref := range l1Refs {
		input := testutils.RandomData(rng, 2000)
		inputs[i] = input
		comm, _ := storage.SetPreImage(ctx, input)
		comms[i] = comm
		t.Log("block", i, "hash", comm)

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
		l1F.ExpectInfoAndTxsByHash(ref.Hash, testutils.RandomBlockInfo(rng), txs, nil)
	}

	engine.ExpectSetSafeHead(refA1.Hash, nil)

	// Set active challenges
	da.SetChallengeStatus(comms[0], refA.Number, daapi.ChallengeActive)
	da.SetChallengeStatus(comms[1], refB.Number, daapi.ChallengeActive)

	factory := NewDASourceFactory(logger, cfg, l1F, da)

	for i := uint64(0); i <= daCfg.ChallengeWindow; i++ {
		ref := l1Refs[i]
		src := factory.OpenData(ctx, ref.ID(), batcherAddr)

		data, err := src.Next(ctx)
		if uint64(i) == daCfg.ChallengeWindow {
			require.ErrorIs(t, ErrReset, err)
		} else {
			require.NoError(t, err)
			require.Equal(t, hexutil.Bytes(inputs[i]), data)
		}

		for j := i * 2; j < i*2+2; j++ {
			t.Log("setting safe head", j, l2Refs[j].Hash)
			da.SetSubjectiveSafeHead(ctx, l2Refs[j])
		}

	}
	engine.ExpectSetSafeHead(refA1.Hash, nil)
	engine.ExpectSetSafeHead(refB1.Hash, nil)

	// Rederive and move safe head forward
	for i, ref := range l1Refs {
		src := factory.OpenData(ctx, ref.ID(), batcherAddr)

		data, err := src.Next(ctx)
		if i == 0 {
			require.Equal(t, io.EOF, err)
		} else if uint64(i) == daCfg.ChallengeWindow+1 {
			require.ErrorIs(t, ErrReset, err)
		} else {
			require.NoError(t, err)
			require.Equal(t, hexutil.Bytes(inputs[i]), data)
		}

		for j := i * 2; j < i*2+2; j++ {
			t.Log("setting safe head", j, l2Refs[j].Hash)
			da.SetSubjectiveSafeHead(ctx, l2Refs[j])
		}

	}

	engine.ExpectSetSafeHead(refB1.Hash, nil)
	// Rederive again starting from refB
	for i := 1; i < len(l1Refs); i++ {
		ref := l1Refs[i]
		src := factory.OpenData(ctx, ref.ID(), batcherAddr)

		data, err := src.Next(ctx)
		if i == 1 {
			require.Equal(t, io.EOF, err)
		} else {
			require.NoError(t, err)
			require.Equal(t, hexutil.Bytes(inputs[i]), data)
		}

		for j := i * 2; j < i*2+2; j++ {
			t.Log("setting safe head", j, l2Refs[j].Hash)
			da.SetSubjectiveSafeHead(ctx, l2Refs[j])
		}

	}

	engine.AssertExpectations(t)
}
