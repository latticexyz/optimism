package mgr

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/ethereum-optimism/optimism/alt-da/api"
	"github.com/ethereum-optimism/optimism/alt-da/client"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

// Makes sure the mock client behaves as expected
func TestMockClient(t *testing.T) {
	log := testlog.Logger(t, log.LvlDebug)
	storage := client.NewMockClient(log)
	engine := &testutils.MockEngine{}

	da := NewAltDA(log, Config{}, storage, engine)

	data := []byte("hello world")
	key, err := storage.SetPreImage(context.TODO(), data)
	require.NoError(t, err)

	resp, err := da.GetPreImage(context.TODO(), key, 1)
	require.NoError(t, err)

	require.Equal(t, resp.Status, api.Available)

	da.SetChallengeStatus(key, 1, api.ChallengeExpired)
	storage.DeleteData(key)

	resp, err = da.GetPreImage(context.TODO(), key, 1)
	require.NoError(t, err)
	require.Equal(t, resp.Status, api.MissingChallengeExpired)

	da.SetChallengeStatus(key, 1, api.ChallengeActive)
	resp, err = da.GetPreImage(context.TODO(), key, 1)
	require.NoError(t, err)
	require.Equal(t, resp.Status, api.MissingPendingChallenge)

	delete(da.blocks.epochs[1].challenges, string(key))
	resp, err = da.GetPreImage(context.TODO(), key, 1)
	require.NoError(t, err)
	require.Equal(t, resp.Status, api.MissingNotChallenged)
}

func RandomHash(rng *rand.Rand) (out common.Hash) {
	rng.Read(out[:])
	return
}

// Tests the challenge window logic implemented in the mock client.
func TestChallengeWindow(t *testing.T) {
	challengeWindow := uint64(12)
	log := testlog.Logger(t, log.LvlDebug)
	storage := client.NewMockClient(log)
	engine := &testutils.MockEngine{}
	l1 := &testutils.MockL1Source{}

	da := NewAltDA(log, Config{ChallengeWindow: challengeWindow}, storage, engine)

	rng := rand.New(rand.NewSource(1234))

	l1Time := uint64(2)

	l1Block := eth.L1BlockRef{
		Hash:       RandomHash(rng),
		Number:     100000,
		ParentHash: RandomHash(rng),
		Time:       rng.Uint64(),
	}
	da.UpdateChallenges(context.TODO(), l1Block.ID(), l1)

	l2Block := eth.L2BlockRef{
		Hash:           RandomHash(rng),
		Number:         0,
		ParentHash:     common.Hash{},
		Time:           l1Block.Time,
		L1Origin:       l1Block.ID(),
		SequenceNumber: 0,
	}
	da.SetSubjectiveSafeHead(context.TODO(), l2Block)

	refs := make([]eth.L2BlockRef, (challengeWindow+6)*6)
	refs[0] = l2Block

	for i := uint64(1); i < challengeWindow+6; i++ {
		// each epoch contains 6 blocks
		num := l2Block.Number
		for j := num + 1; j < num+7; j++ {
			l2Block = eth.L2BlockRef{
				Hash:           RandomHash(rng),
				Number:         j,
				ParentHash:     l2Block.Hash,
				Time:           l2Block.Time + 1,
				L1Origin:       l1Block.ID(),
				SequenceNumber: j - (num + 1),
			}
			refs[j] = l2Block
			da.SetSubjectiveSafeHead(context.TODO(), l2Block)
		}
		l1Block = eth.L1BlockRef{
			Hash:       RandomHash(rng),
			Number:     l1Block.Number + 1,
			ParentHash: l1Block.Hash,
			Time:       l1Block.Time + l1Time,
		}
		fmt.Println("i", i, "num", num)

		if i >= challengeWindow {
			num := (i - challengeWindow + 1) * 6
			expectedSafeHead := refs[num]
			engine.ExpectSetSafeHead(expectedSafeHead.Hash, nil)
		}

		_, err := da.UpdateChallenges(context.TODO(), l1Block.ID(), l1)
		require.NoError(t, err)

	}
}
