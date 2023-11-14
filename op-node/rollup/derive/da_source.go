package derive

import (
	"context"
	"fmt"

	da "github.com/ethereum-optimism/optimism/alt-da/api"
	damgr "github.com/ethereum-optimism/optimism/alt-da/mgr"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type DAFetcher interface {
	GetPreImage(ctx context.Context, comm []byte, blockNumber uint64) (*da.Response, error)
	AdvanceWindow(ctx context.Context, epoch eth.BlockID, l1 damgr.L1Fetcher) (bool, error)
	StepChallenges(ctx context.Context, l1 damgr.L1Fetcher) error
	UpdateChallenges(ctx context.Context, epoch eth.BlockID, l1 damgr.L1Fetcher) error
}

// DASourceFactory connects to the DA service to fetch input data and keep track of
// expired challenges.
type DASourceFactory struct {
	log log.Logger
	da  DAFetcher
	cfg *rollup.Config
	l1  damgr.L1Fetcher
}

func NewDASourceFactory(log log.Logger, cfg *rollup.Config, l1 damgr.L1Fetcher, da DAFetcher) *DASourceFactory {
	return &DASourceFactory{
		log: log,
		da:  da,
		cfg: cfg,
		l1:  l1,
	}
}

// OpenData refreshes the expired challenge state once for each new l1 block and determines whether
// the pipeline should be reset by comparing the new state to the previous state.
func (f *DASourceFactory) OpenData(ctx context.Context, id eth.BlockID, batcherAddr common.Address) DataIter {
	// the challenge state will only change after a new l1 block is processed so we fetch it here
	// instead of for every batch.
	err := f.da.UpdateChallenges(ctx, id, f.l1)
	reset, err := f.da.AdvanceWindow(ctx, id, f.l1)
	if reset {
		log.Info("challenge number updated, need reset...")
		// the error is forwarded to the iterator so it will be returned first thing.
		err = NewResetError(fmt.Errorf("new expired challenge"))
	}

	return NewDASource(f.log, f.l1, NewDataSource(ctx, f.log, f.cfg, f.l1, id, batcherAddr), f.da, id, err)
}

type DASource struct {
	log log.Logger
	l1  damgr.L1Fetcher
	src DataIter
	da  DAFetcher
	id  eth.BlockID
	// keep track of a pending commitment so we can keep trying to fetch the pre-image
	comm []byte
	err  error
}

func NewDASource(log log.Logger, l1 damgr.L1Fetcher, src DataIter, da DAFetcher, id eth.BlockID, err error) *DASource {
	return &DASource{
		log: log,
		l1:  l1,
		src: src,
		da:  da,
		id:  id,
		err: err,
	}
}

func (s *DASource) Next(ctx context.Context) (eth.Data, error) {
	// if we need to reset the pipeline, return the error first.
	if s.err != nil {
		err := s.err
		s.err = nil
		return nil, err
	}

	if s.comm == nil {
		var err error
		// the l1 source returns the input commitment for the batch.
		s.comm, err = s.src.Next(ctx)
		if err != nil {
			return nil, err
		}
	}
	// wrap in eth.Data for logging as hex string.
	comm := eth.Data(s.comm)
	// use the commitment to fetch the pre-image from the DA service.
	resp, err := s.da.GetPreImage(ctx, s.comm, s.id.Number)
	if err != nil {
		tempErr := NewTemporaryError(fmt.Errorf("failed to fetch input data with comm %v from da service: %w", comm, err))
		return nil, tempErr
	}
	// if the challenge is expired, we always skip to the next batch.
	if resp.IsExpired() {
		s.log.Warn("challenge expired, skipping batch", "comm", comm)
		s.comm = nil
		return s.Next(ctx)
	}

	// if the data is available regardless of the challenge status we
	// return and continue to the next batch.
	if resp.IsAvailable() {
		s.comm = nil
		// safe as IsAvailable already checks for nil
		return *resp.Data, nil
	}

	// if the data is not available but a challenge is open we loop forever
	// until the challenge is resolved or expired.
	if resp.IsPending() {
		tempErr := NewTemporaryError(fmt.Errorf("data for comm %v not available, waiting for challenge to resolve or expire", comm))
		// keep advancing the l1 head until the challenge is resolved or expired.
		s.da.StepChallenges(ctx, s.l1)
		return nil, tempErr
	}

	// If the data is not available the derivation will stall forever.
	return nil, NewCriticalError(fmt.Errorf("data for comm %v not available", comm))
}
