package mgr

import (
	"container/heap"
	"errors"

	"github.com/ethereum-optimism/optimism/alt-da/api"
	"github.com/ethereum-optimism/optimism/alt-da/metrics"
	"github.com/ethereum/go-ethereum/log"
)

var ErrChallengeExpired = errors.New("challenge expired")

// Commitment keeps track of the onchain state of an input commitment.
type Commitment struct {
	hash            []byte                  // the keccak256 hash of the input
	input           []byte                  // the input itself if it was resolved onchain
	expiresAt       uint64                  // represents the block number after which the commitment can no longer be challenged or if challenged no longer be resolved.
	blockNumber     uint64                  // block where the commitment is included as calldata to the batcher inbox
	challengeStatus api.ChallengeStatusCode // latest known challenge status
}

// CommQueue is a queue of commitments ordered by block number.
type CommQueue []*Commitment

var _ heap.Interface = (*CommQueue)(nil)

func (c CommQueue) Len() int { return len(c) }

func (c CommQueue) Less(i, j int) bool {
	return c[i].blockNumber < c[j].blockNumber
}

func (c CommQueue) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *CommQueue) Push(x any) {
	*c = append(*c, x.(*Commitment))
}

func (c *CommQueue) Pop() any {
	old := *c
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*c = old[0 : n-1]
	return item
}

// State tracks the commitment and their challenges in order of l1 inclusion.
type State struct {
	comms       CommQueue
	commsByHash map[string]*Commitment
	log         log.Logger
	metrics     metrics.AltDAMetricer
}

func NewState(log log.Logger, m metrics.AltDAMetricer) *State {
	return &State{
		comms:       make(CommQueue, 0),
		commsByHash: make(map[string]*Commitment),
		log:         log,
		metrics:     m,
	}
}

// IsTracking returns whether we currently have a commitment for the given hash.
func (s *State) IsTracking(comm []byte, bn uint64) bool {
	if c, ok := s.commsByHash[string(comm)]; ok {
		return c.blockNumber == bn
	}
	return false
}

// SetActiveChallenge switches the state of a given commitment to active challenge. Noop if
// the commitment is not tracked as we don't want to track challenges for invalid commitments.
func (s *State) SetActiveChallenge(comm []byte, challengedAt uint64, resolveWindow uint64) {
	if c, ok := s.commsByHash[string(comm)]; ok {
		c.expiresAt = challengedAt + resolveWindow
		c.challengeStatus = api.ChallengeActive
		s.metrics.RecordActiveChallenge(c.blockNumber, challengedAt, comm)
	}
}

// SetResolvedChallenge switches the state of a given commitment to resolved. Noop if
// the commitment is not tracked as we don't want to track challenges for invalid commitments.
// The input posted onchain is stored in the state for later retrieval.
func (s *State) SetResolvedChallenge(comm []byte, input []byte, resolvedAt uint64) {
	if c, ok := s.commsByHash[string(comm)]; ok {
		c.challengeStatus = api.ChallengeResolved
		c.expiresAt = resolvedAt
		c.input = input
		s.metrics.RecordResolvedChallenge(comm)
	}
}

// SetInputCommitment initializes a new commitment and adds it to the state.
func (s *State) SetInputCommitment(comm []byte, committedAt uint64, challengeWindow uint64) *Commitment {
	c := &Commitment{
		hash:        comm,
		expiresAt:   committedAt + challengeWindow,
		blockNumber: committedAt,
	}
	s.log.Debug("append commitment", "expiresAt", c.expiresAt, "blockNumber", c.blockNumber)
	heap.Push(&s.comms, c)
	s.commsByHash[string(comm)] = c

	return c
}

// GetOrTrackChallenge returns the commitment for the given hash if it is already tracked, or
// initializes a new commitment and adds it to the state.
func (s *State) GetOrTrackChallenge(comm []byte, bn uint64, challengeWindow uint64) *Commitment {
	if c, ok := s.commsByHash[string(comm)]; ok {
		return c
	}
	return s.SetInputCommitment(comm, bn, challengeWindow)
}

// GetResolvedInput returns the input bytes if the commitment was resolved onchain.
func (s *State) GetResolvedInput(comm []byte) ([]byte, error) {
	if c, ok := s.commsByHash[string(comm)]; ok {
		return c.input, nil
	}
	return nil, errors.New("commitment not found")
}

// ExpireChallenges walks back from the oldest commitment to find the latest l1 origin
// for which input data can no longer be challenged. It also marks any active challenges
// as expired based on the new latest l1 origin. If any active challenges are expired
// it returns an error to signal that a derivation pipeline reset is required.
func (s *State) ExpireChallenges(bn uint64) (uint64, error) {
	latest := uint64(0)
	var err error
	for i := 0; i < len(s.comms); i++ {
		c := s.comms[i]
		if c.expiresAt <= bn && c.blockNumber > latest {
			latest = c.blockNumber

			if c.challengeStatus == api.ChallengeActive {
				c.challengeStatus = api.ChallengeExpired
				s.metrics.RecordExpiredChallenge(c.hash)
				err = ErrChallengeExpired
			}
		} else {
			break
		}
	}
	return latest, err
}

// Prune removes commitments once they can no longer be challenged or resolved.
func (s *State) Prune(bn uint64) {
	for i := 0; i < len(s.comms); i++ {
		c := s.comms[i]
		if c.blockNumber < bn {
			s.log.Debug("prune commitment", "expiresAt", c.expiresAt, "blockNumber", c.blockNumber)
			delete(s.commsByHash, string(c.hash))
		} else {
			s.comms = s.comms[i:]
			break
		}
	}
}
