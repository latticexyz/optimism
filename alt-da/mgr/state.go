package mgr

import (
	"errors"

	"github.com/ethereum-optimism/optimism/alt-da/api"
	"github.com/ethereum/go-ethereum/log"
)

var ErrChallengeExpired = errors.New("challenge expired")

type Commitment struct {
	hash            []byte
	input           []byte
	expiresAt       uint64
	blockNumber     uint64
	challengeStatus api.ChallengeStatusCode
}

type CommQueue []*Commitment

type State struct {
	comms       CommQueue
	commsByHash map[string]*Commitment
	log         log.Logger
}

func NewState(log log.Logger) *State {
	return &State{
		comms:       make(CommQueue, 0),
		commsByHash: make(map[string]*Commitment),
		log:         log,
	}
}

func (s *State) IsTracking(comm []byte, bn uint64) bool {
	if c, ok := s.commsByHash[string(comm)]; ok {
		return c.blockNumber == bn
	}
	return false
}

func (s *State) SetActiveChallenge(comm []byte, challengedAt uint64, resolveWindow uint64) {
	if c, ok := s.commsByHash[string(comm)]; ok {
		c.expiresAt = challengedAt + resolveWindow
		c.challengeStatus = api.ChallengeActive
	}
}

func (s *State) SetResolvedChallenge(comm []byte, input []byte, resolvedAt uint64) {
	if c, ok := s.commsByHash[string(comm)]; ok {
		c.challengeStatus = api.ChallengeResolved
		c.expiresAt = resolvedAt
		c.input = input
	}
}

func (s *State) SetInputCommitment(comm []byte, committedAt uint64, challengeWindow uint64) *Commitment {
	c := &Commitment{
		hash:        comm,
		expiresAt:   committedAt + challengeWindow,
		blockNumber: committedAt,
	}
	s.log.Debug("append commitment", "expiresAt", c.expiresAt, "blockNumber", c.blockNumber)
	s.comms = append(s.comms, c)
	s.commsByHash[string(comm)] = c

	return c
}

func (s *State) GetOrTrackChallenge(comm []byte, bn uint64, challengeWindow uint64) *Commitment {
	if c, ok := s.commsByHash[string(comm)]; ok {
		return c
	}
	return s.SetInputCommitment(comm, bn, challengeWindow)
}

func (s *State) GetResolvedInput(comm []byte) ([]byte, error) {
	if c, ok := s.commsByHash[string(comm)]; ok {
		return c.input, nil
	}
	return nil, errors.New("commitment not found")
}

func (s *State) ExpireChallenges(bn uint64) (uint64, error) {
	latest := uint64(0)
	var err error
	for i := 0; i < len(s.comms); i++ {
		c := s.comms[i]
		if c.expiresAt <= bn && c.blockNumber > latest {
			latest = c.blockNumber

			if c.challengeStatus == api.ChallengeActive {
				c.challengeStatus = api.ChallengeExpired
				err = ErrChallengeExpired
			}
		} else {
			break
		}
	}
	return latest, err
}

func (s *State) Prune(bn uint64) {
	for i := 0; i < len(s.comms); i++ {
		c := s.comms[i]
		if c.expiresAt < bn {
			s.log.Debug("prune commitment", "expiresAt", c.expiresAt, "blockNumber", c.blockNumber)
			delete(s.commsByHash, string(c.hash))
		} else {
			s.comms = s.comms[i:]
			break
		}
	}
}
