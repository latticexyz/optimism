package driver

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	damgr "github.com/ethereum-optimism/optimism/alt-da/mgr"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-node/rollup/sync"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

type Metrics interface {
	RecordPipelineReset()
	RecordPublishingError()
	RecordDerivationError()

	RecordReceivedUnsafePayload(payload *eth.ExecutionPayload)

	RecordL1Ref(name string, ref eth.L1BlockRef)
	RecordL2Ref(name string, ref eth.L2BlockRef)
	RecordChannelInputBytes(inputCompressedBytes int)
	RecordHeadChannelOpened()
	RecordChannelTimedOut()
	RecordFrame()

	RecordUnsafePayloadsBuffer(length uint64, memSize uint64, next eth.BlockID)

	SetDerivationIdle(idle bool)

	RecordL1ReorgDepth(d uint64)

	EngineMetrics
	L1FetcherMetrics
	SequencerMetrics
}

type L1Chain interface {
	derive.L1Fetcher
	L1BlockRefByLabel(context.Context, eth.BlockLabel) (eth.L1BlockRef, error)
}

type L2Chain interface {
	derive.Engine
	L2BlockRefByLabel(ctx context.Context, label eth.BlockLabel) (eth.L2BlockRef, error)
	L2BlockRefByHash(ctx context.Context, l2Hash common.Hash) (eth.L2BlockRef, error)
	L2BlockRefByNumber(ctx context.Context, num uint64) (eth.L2BlockRef, error)
}

type SubjectiveEngine interface {
	L2Chain
	SetSubjectiveSafeHead(ctx context.Context, safeHash common.Hash) error
}

type DerivationPipeline interface {
	Reset()
	Step(ctx context.Context) error
	AddUnsafePayload(payload *eth.ExecutionPayload)
	UnsafeL2SyncTarget() eth.L2BlockRef
	Finalize(ref eth.L1BlockRef)
	FinalizedL1() eth.L1BlockRef
	Finalized() eth.L2BlockRef
	SafeL2Head() eth.L2BlockRef
	UnsafeL2Head() eth.L2BlockRef
	PendingSafeL2Head() eth.L2BlockRef
	Origin() eth.L1BlockRef
	EngineReady() bool
	EngineSyncTarget() eth.L2BlockRef
}

type L1StateIface interface {
	HandleNewL1HeadBlock(head eth.L1BlockRef)
	HandleNewL1SafeBlock(safe eth.L1BlockRef)
	HandleNewL1FinalizedBlock(finalized eth.L1BlockRef)

	L1Head() eth.L1BlockRef
	L1Safe() eth.L1BlockRef
	L1Finalized() eth.L1BlockRef
}

type SequencerIface interface {
	StartBuildingBlock(ctx context.Context) error
	CompleteBuildingBlock(ctx context.Context) (*eth.ExecutionPayload, error)
	PlanNextSequencerAction() time.Duration
	RunNextSequencerAction(ctx context.Context) (*eth.ExecutionPayload, error)
	BuildingOnto() eth.L2BlockRef
	CancelBuildingBlock(ctx context.Context)
}

type Network interface {
	// PublishL2Payload is called by the driver whenever there is a new payload to publish, synchronously with the driver main loop.
	PublishL2Payload(ctx context.Context, payload *eth.ExecutionPayload) error
}

type AltSync interface {
	// RequestL2Range informs the sync source that the given range of L2 blocks is missing,
	// and should be retrieved from any available alternative syncing source.
	// The start and end of the range are exclusive:
	// the start is the head we already have, the end is the first thing we have queued up.
	// It's the task of the alt-sync mechanism to use this hint to fetch the right payloads.
	// Note that the end and start may not be consistent: in this case the sync method should fetch older history
	//
	// If the end value is zeroed, then the sync-method may determine the end free of choice,
	// e.g. sync till the chain head meets the wallclock time. This functionality is optional:
	// a fixed target to sync towards may be determined by picking up payloads through P2P gossip or other sources.
	//
	// The sync results should be returned back to the driver via the OnUnsafeL2Payload(ctx, payload) method.
	// The latest requested range should always take priority over previous requests.
	// There may be overlaps in requested ranges.
	// An error may be returned if the scheduling fails immediately, e.g. a context timeout.
	RequestL2Range(ctx context.Context, start, end eth.L2BlockRef) error
}

type SequencerStateListener interface {
	SequencerStarted() error
	SequencerStopped() error
}

// ModEngine wraps an engine client and forwards safe head calls to subjective safe head.
type ModEngine struct {
	SubjectiveEngine
	da  *damgr.AltDA
	log log.Logger
}

// ForkchoiceUpdate forwards the safe head to the da service and then calls the underlying engine
// subjective call.
func (meng *ModEngine) ForkchoiceUpdate(ctx context.Context, state *eth.ForkchoiceState, attr *eth.PayloadAttributes) (*eth.ForkchoiceUpdatedResult, error) {
	if state.SafeBlockHash != (common.Hash{}) {
		safeHash := state.SafeBlockHash
		// clear the safe hash so it isn't set during the forkchoice update
		state.SafeBlockHash = common.Hash{}

		result, err := meng.SubjectiveEngine.ForkchoiceUpdate(ctx, state, attr)
		if err != nil {
			return nil, err
		}

		// then set the subjective safe head instead of the safe head
		if err := meng.SetSubjectiveSafeHead(ctx, safeHash); err != nil {
			return nil, err
		}

		// forward the l2ref to the da service
		blkRef, err := meng.L2BlockRefByHash(ctx, safeHash)
		// error isn't a show stopper here, we can still update the engine
		if err != nil || meng.da.SetSubjectiveSafeHead(ctx, blkRef) != nil {
			log.Error("failed to notify da service about new l2 head", "err", err)
		}
		return result, nil
	}

	return meng.SubjectiveEngine.ForkchoiceUpdate(ctx, state, attr)
}

func NewModEngine(eng SubjectiveEngine, da *damgr.AltDA, log log.Logger) *ModEngine {
	return &ModEngine{
		SubjectiveEngine: eng,
		da:               da,
		log:              log,
	}
}

// NewDriver composes an events handler that tracks L1 state, triggers L2 derivation, and optionally sequences new L2 blocks.
func NewDriver(driverCfg *Config, cfg *rollup.Config, l2 L2Chain, l1 L1Chain, altSync AltSync, dataSrc derive.DataAvailabilitySource, network Network, log log.Logger, snapshotLog log.Logger, metrics Metrics, sequencerStateListener SequencerStateListener, syncCfg *sync.Config) *Driver {
	l1 = NewMeteredL1Fetcher(l1, metrics)
	l1State := NewL1State(log, metrics)
	sequencerConfDepth := NewConfDepth(driverCfg.SequencerConfDepth, l1State.L1Head, l1)
	findL1Origin := NewL1OriginSelector(log, cfg, sequencerConfDepth)
	verifConfDepth := NewConfDepth(driverCfg.VerifierConfDepth, l1State.L1Head, l1)
	derivationPipeline := derive.NewDerivationPipeline(log, cfg, verifConfDepth, dataSrc, l2, metrics, syncCfg)
	attrBuilder := derive.NewFetchingAttributesBuilder(cfg, l1, l2)
	engine := derivationPipeline
	meteredEngine := NewMeteredEngine(cfg, engine, metrics, log)
	sequencer := NewSequencer(log, cfg, meteredEngine, attrBuilder, findL1Origin, metrics)

	return &Driver{
		l1State:          l1State,
		derivation:       derivationPipeline,
		stateReq:         make(chan chan struct{}),
		forceReset:       make(chan chan struct{}, 10),
		startSequencer:   make(chan hashAndErrorChannel, 10),
		stopSequencer:    make(chan chan hashAndError, 10),
		sequencerActive:  make(chan chan bool, 10),
		sequencerNotifs:  sequencerStateListener,
		config:           cfg,
		driverConfig:     driverCfg,
		done:             make(chan struct{}),
		log:              log,
		snapshotLog:      snapshotLog,
		l1:               l1,
		l2:               l2,
		sequencer:        sequencer,
		network:          network,
		metrics:          metrics,
		l1HeadSig:        make(chan eth.L1BlockRef, 10),
		l1SafeSig:        make(chan eth.L1BlockRef, 10),
		l1FinalizedSig:   make(chan eth.L1BlockRef, 10),
		unsafeL2Payloads: make(chan *eth.ExecutionPayload, 10),
		altSync:          altSync,
	}
}
