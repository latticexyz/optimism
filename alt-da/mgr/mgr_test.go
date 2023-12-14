package mgr

import (
	"math/rand"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

func TestDAChallengeState(t *testing.T) {
	logger := testlog.Logger(t, log.LvlDebug)

	rng := rand.New(rand.NewSource(1234))
	state := NewState(logger)

	i := uint64(1)

	challengeWindow := uint64(6)
	resolveWindow := uint64(6)

	// track commitments in the first 10 blocks
	for ; i < 10; i++ {
		state.SetInputCommitment(testutils.RandomData(rng, 32), i, challengeWindow)
	}

	// blocks are safe after the challenge window expires
	bn, err := state.ExpireChallenges(i)
	require.NoError(t, err)
	require.Equal(t, uint64(4), bn)

	// track the next commitment and mark it as challenged
	c := testutils.RandomData(rng, 32)
	state.SetInputCommitment(c, i, challengeWindow)
	state.SetActiveChallenge(c, i+4, resolveWindow)

	for j := i + 1; j < i+8; j++ {
		state.SetInputCommitment(testutils.RandomData(rng, 32), j, challengeWindow)
	}

	// safe l1 origin should not extend past the resolve window
	bn, err = state.ExpireChallenges(i + 8)
	require.NoError(t, err)
	require.Equal(t, i-1, bn)

	for j := i + 8; j < i+12; j++ {
		state.SetInputCommitment(testutils.RandomData(rng, 32), j, challengeWindow)
	}

	// no more active challenges, the head can catch up to the challenge window
	bn, err = state.ExpireChallenges(i + 12)
	require.ErrorIs(t, err, ErrChallengeExpired)
	require.Equal(t, i+6, bn)

	// cleanup state we don't need anymore
	state.Prune(i + 12)
	bn, err = state.ExpireChallenges(i + 12)
	require.NoError(t, err)
	require.Equal(t, i+6, bn)

	i = i + 12
	// add one more commitment and challenge it
	c = testutils.RandomData(rng, 32)
	state.SetInputCommitment(c, i, challengeWindow)
	// challenge 3 blocks after
	state.SetActiveChallenge(c, i+3, resolveWindow)

	// exceed the challenge window with more commitments
	for j := i + 1; j < i+8; j++ {
		state.SetInputCommitment(testutils.RandomData(rng, 32), j, challengeWindow)
	}

	// finalized head should not extend past the resolve window
	bn, err = state.ExpireChallenges(i + 8)
	require.NoError(t, err)
	require.Equal(t, i-1, bn)

	input := testutils.RandomData(rng, 100)
	// resolve the challenge
	state.SetResolvedChallenge(c, input, i+8)

	// finalized head catches up
	bn, err = state.ExpireChallenges(i + 9)
	require.NoError(t, err)
	require.Equal(t, i+3, bn)

	storedInput, err := state.GetResolvedInput(c)
	require.NoError(t, err)
	require.Equal(t, input, storedInput)
}
