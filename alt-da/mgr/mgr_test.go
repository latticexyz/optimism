package mgr

import (
	"math/rand"
	"testing"

	"github.com/ethereum-optimism/optimism/alt-da/api"
	"github.com/ethereum-optimism/optimism/alt-da/metrics"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

func TestDAChallengeState(t *testing.T) {
	logger := testlog.Logger(t, log.LvlDebug)

	rng := rand.New(rand.NewSource(1234))
	state := NewState(logger, &metrics.NoopAltDAMetrics{})

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
	// i+4 is the block at which it was challenged
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

func TestExpireChallenges(t *testing.T) {
	logger := testlog.Logger(t, log.LvlDebug)

	rng := rand.New(rand.NewSource(1234))
	state := NewState(logger, &metrics.NoopAltDAMetrics{})

	comms := make(map[uint64][]byte)

	i := uint64(3713854)

	var finalized uint64

	challengeWindow := uint64(90)
	resolveWindow := uint64(90)

	for ; i < 3713948; i += 6 {
		comm := testutils.RandomData(rng, 32)
		comms[i] = comm
		logger.Info("set commitment", "block", i)
		cm := state.GetOrTrackChallenge(comm, i, challengeWindow)
		require.NotNil(t, cm)

		bn, err := state.ExpireChallenges(i)
		logger.Info("expire challenges", "expired", bn, "err", err)

		if bn > finalized {
			finalized = bn
			state.Prune(bn)
		}
	}

	state.SetActiveChallenge(comms[3713926], 3713948, resolveWindow)

	state.SetActiveChallenge(comms[3713932], 3713950, resolveWindow)

	for ; i < 3714038; i += 6 {
		comm := testutils.RandomData(rng, 32)
		comms[i] = comm
		logger.Info("set commitment", "block", i)
		cm := state.GetOrTrackChallenge(comm, i, challengeWindow)
		require.NotNil(t, cm)

		bn, err := state.ExpireChallenges(i)
		logger.Info("expire challenges", "expired", bn, "err", err)

		if bn > finalized {
			finalized = bn
			state.Prune(bn)
		}

	}

	bn, err := state.ExpireChallenges(3714034)
	require.NoError(t, err)
	require.Equal(t, uint64(3713920), bn)

	bn, err = state.ExpireChallenges(3714035)
	require.NoError(t, err)
	require.Equal(t, uint64(3713920), bn)

	bn, err = state.ExpireChallenges(3714036)
	require.NoError(t, err)
	require.Equal(t, uint64(3713920), bn)

	bn, err = state.ExpireChallenges(3714037)
	require.NoError(t, err)
	require.Equal(t, uint64(3713920), bn)

	bn, err = state.ExpireChallenges(3714038)
	require.ErrorIs(t, err, ErrChallengeExpired)

	for i := uint64(3713854); i < 3714044; i += 6 {
		cm := state.GetOrTrackChallenge(comms[i], i, challengeWindow)
		require.NotNil(t, cm)

		if i == 3713926 {
			require.Equal(t, api.ChallengeExpired, cm.challengeStatus)
		}
	}

	bn, err = state.ExpireChallenges(3714038)
	require.NoError(t, err)

	require.Equal(t, uint64(3713926), bn)
}
