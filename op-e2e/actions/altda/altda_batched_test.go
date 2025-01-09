package altda

import (
	"math/rand"
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-e2e/config"

	"github.com/ethereum-optimism/optimism/op-e2e/actions/helpers"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	altda "github.com/ethereum-optimism/optimism/op-alt-da"
	"github.com/ethereum-optimism/optimism/op-alt-da/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
)

// L2AltDA is a test harness for manipulating AltDA state.

type AltDAParamBatched func(p *e2eutils.TestParams)

func NewL2AltDABatched(t helpers.Testing, params ...AltDAParamBatched) *L2AltDA {
	p := &e2eutils.TestParams{
		MaxSequencerDrift:   40,
		SequencerWindowSize: 12,
		ChannelTimeout:      12,
		L1BlockTime:         12,
		UseAltDA:            true,
		AllocType:           config.AllocTypeAltDA,
	}
	for _, apply := range params {
		apply(p)
	}
	log := testlog.Logger(t, log.LvlDebug)

	dp := e2eutils.MakeDeployParams(t, p)
	sd := e2eutils.Setup(t, dp, helpers.DefaultAlloc)

	require.True(t, sd.RollupCfg.AltDAEnabled())

	miner := helpers.NewL1Miner(t, log, sd.L1Cfg)
	l1Client := miner.EthClient()

	jwtPath := e2eutils.WriteDefaultJWT(t)
	engine := helpers.NewL2Engine(t, log, sd.L2Cfg, jwtPath)
	engCl := engine.EngineClient(t, sd.RollupCfg)

	storage := &altda.DAErrFaker{Client: altda.NewMockDAClient(log)}

	l1F, err := sources.NewL1Client(miner.RPCClient(), log, nil, sources.L1ClientDefaultConfig(sd.RollupCfg, false, sources.RPCKindBasic))
	require.NoError(t, err)

	altDACfg, err := sd.RollupCfg.GetOPAltDAConfig()
	require.NoError(t, err)

	daMgr := altda.NewAltDAWithStorage(log, altDACfg, storage, &altda.NoopMetrics{})

	sequencer := helpers.NewL2Sequencer(t, log, l1F, miner.BlobStore(), daMgr, engCl, sd.RollupCfg, 0)
	miner.ActL1SetFeeRecipient(common.Address{'A'})
	sequencer.ActL2PipelineFull(t)

	batcher := helpers.NewL2Batcher(log, sd.RollupCfg, helpers.BatchedCommsBatcherCfg(dp, storage), sequencer.RollupClient(), l1Client, engine.EthClient(), engCl)

	addresses := e2eutils.CollectAddresses(sd, dp)
	cl := engine.EthClient()
	l2UserEnv := &helpers.BasicUserEnv[*helpers.L2Bindings]{
		EthCl:          cl,
		Signer:         types.LatestSigner(sd.L2Cfg.Config),
		AddressCorpora: addresses,
		Bindings:       helpers.NewL2Bindings(t, cl, engine.GethClient()),
	}
	alice := helpers.NewCrossLayerUser(log, dp.Secrets.Alice, rand.New(rand.NewSource(0xa57b)), p.AllocType)
	alice.L2.SetUserEnv(l2UserEnv)

	contract, err := bindings.NewDataAvailabilityChallenge(sd.RollupCfg.AltDAConfig.DAChallengeAddress, l1Client)
	require.NoError(t, err)

	challengeWindow, err := contract.ChallengeWindow(nil)
	require.NoError(t, err)
	require.Equal(t, altDACfg.ChallengeWindow, challengeWindow.Uint64())

	resolveWindow, err := contract.ResolveWindow(nil)
	require.NoError(t, err)
	require.Equal(t, altDACfg.ResolveWindow, resolveWindow.Uint64())

	return &L2AltDA{
		log:       log,
		storage:   storage,
		daMgr:     daMgr,
		altDACfg:  altDACfg,
		contract:  contract,
		batcher:   batcher,
		sequencer: sequencer,
		engine:    engine,
		engCl:     engCl,
		sd:        sd,
		dp:        dp,
		miner:     miner,
		alice:     alice,
	}
}

func TestAltDABatched_Derivation(gt *testing.T) {
	t := helpers.NewDefaultTesting(gt)
	harness := NewL2AltDABatched(t)

	harness.ActL1Blocks(t, 5)

	cl := harness.engine.EthClient()

	rng := rand.New(rand.NewSource(555))

	// build 2 L2 blocks filled with large txs of random data
	for i := 0; i < 20; i++ {
		aliceNonce, err := cl.PendingNonceAt(t.Ctx(), harness.dp.Addresses.Alice)
		status := harness.sequencer.SyncStatus()
		// build empty L1 blocks as necessary, so the L2 sequencer can continue to include txs while not drifting too far out
		if status.UnsafeL2.Time >= status.HeadL1.Time+12 {
			harness.miner.ActEmptyBlock(t)
		}
		harness.sequencer.ActL1HeadSignal(t)
		harness.sequencer.ActL2StartBlock(t)
		baseFee := harness.engine.L2Chain().CurrentBlock().BaseFee
		// fill the block with large L2 txs from alice
		for n := aliceNonce; ; n++ {
			require.NoError(t, err)
			signer := types.LatestSigner(harness.sd.L2Cfg.Config)
			data := make([]byte, 120_000) // very large L2 txs, as large as the tx-pool will accept
			_, err := rng.Read(data[:])   // fill with random bytes, to make compression ineffective
			require.NoError(t, err)
			gas, err := core.IntrinsicGas(data, nil, false, true, true, false)
			require.NoError(t, err)
			if gas > harness.engine.EngineApi.RemainingBlockGas() {
				break
			}
			tx := types.MustSignNewTx(harness.dp.Secrets.Alice, signer, &types.DynamicFeeTx{
				ChainID:   harness.sd.L2Cfg.Config.ChainID,
				Nonce:     n,
				GasTipCap: big.NewInt(2 * params.GWei),
				GasFeeCap: new(big.Int).Add(new(big.Int).Mul(baseFee, big.NewInt(2)), big.NewInt(2*params.GWei)),
				Gas:       gas,
				To:        &harness.dp.Addresses.Bob,
				Value:     big.NewInt(0),
				Data:      data,
			})
			require.NoError(t, cl.SendTransaction(t.Ctx(), tx))
			harness.engine.ActL2IncludeTx(harness.dp.Addresses.Alice)(t)
		}
		harness.sequencer.ActL2EndBlock(t)
	}

	harness.batcher.ActL2SubmitBatchedCommitments(t, 1, func(tx *types.DynamicFeeTx) {
		// skip txdata version byte
		harness.lastComm = tx.Data[1:]
	})

	harness.miner.ActL1StartBlock(12)(t)
	harness.miner.ActL1IncludeTx(harness.dp.Addresses.Batcher)(t)
	harness.miner.ActL1EndBlock(t)

	harness.lastCommBn = harness.miner.L1Chain().CurrentBlock().Number.Uint64()
}

