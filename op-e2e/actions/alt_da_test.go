package actions

import (
	"math/big"
	"math/rand"
	"testing"

	daclient "github.com/ethereum-optimism/optimism/alt-da/client"
	damgr "github.com/ethereum-optimism/optimism/alt-da/mgr"
	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-node/rollup/sync"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

// L2AltDA is a test harness for manipulating DA service state.
type L2AltDA struct {
	log        log.Logger
	storage    *daclient.DAErrFaker
	daMgr      *damgr.AltDA
	contract   *bindings.DataAvailabilityChallenge
	batcher    *L2Batcher
	sequencer  *L2Sequencer
	engine     *L2Engine
	engCl      *sources.EngineClient
	sd         *e2eutils.SetupData
	dp         *e2eutils.DeployParams
	miner      *L1Miner
	alice      *CrossLayerUser
	lastComm   []byte
	lastCommBn uint64
}

func NewL2AltDA(log log.Logger, p *e2eutils.TestParams, t Testing) *L2AltDA {
	dp := e2eutils.MakeDeployParams(t, p)
	sd := e2eutils.Setup(t, dp, defaultAlloc)

	miner := NewL1Miner(t, log, sd.L1Cfg)
	l1Client := miner.EthClient()

	jwtPath := e2eutils.WriteDefaultJWT(t)
	engine := NewL2Engine(t, log, sd.L2Cfg, sd.RollupCfg.Genesis.L1, jwtPath)
	engCl := engine.EngineClient(t, sd.RollupCfg)

	storage := &daclient.DAErrFaker{Client: daclient.NewMockClient(log)}

	daMgr := damgr.NewAltDA(log, *sd.DaCfg, storage, engCl)

	l1F, err := sources.NewL1Client(miner.RPCClient(), log, nil, sources.L1ClientDefaultConfig(sd.RollupCfg, false, sources.RPCKindBasic))
	require.NoError(t, err)

	dataSrc := derive.NewDASourceFactory(log, sd.RollupCfg, l1F, daMgr)
	modEng := NewModEngine(engCl, daMgr)
	sequencer := NewL2Sequencer(t, log, l1F, modEng, sd.RollupCfg, 0, dataSrc)
	miner.ActL1SetFeeRecipient(common.Address{'A'})
	sequencer.ActL2PipelineFull(t)

	batcher := NewL2Batcher(log, sd.RollupCfg, &BatcherCfg{
		MinL1TxSize: 0,
		MaxL1TxSize: 128_000,
		BatcherKey:  dp.Secrets.Batcher,
		AltDA:       storage,
	}, sequencer.RollupClient(), l1Client, engine.EthClient(), engCl)

	addresses := e2eutils.CollectAddresses(sd, dp)
	cl := engine.EthClient()
	l2UserEnv := &BasicUserEnv[*L2Bindings]{
		EthCl:          cl,
		Signer:         types.LatestSigner(sd.L2Cfg.Config),
		AddressCorpora: addresses,
		Bindings:       NewL2Bindings(t, cl, engine.GethClient()),
	}
	alice := NewCrossLayerUser(log, dp.Secrets.Alice, rand.New(rand.NewSource(0xa57b)))
	alice.L2.SetUserEnv(l2UserEnv)

	contract, err := bindings.NewDataAvailabilityChallenge(sd.DaCfg.DaChallengeContractAddress, l1Client)
	require.NoError(t, err)
	return &L2AltDA{
		log:       log,
		storage:   storage,
		daMgr:     daMgr,
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

func (a *L2AltDA) StorageClient() *daclient.DAErrFaker {
	return a.storage
}

func (a *L2AltDA) NewVerifier(t Testing) *L2Verifier {
	jwtPath := e2eutils.WriteDefaultJWT(t)
	engine := NewL2Engine(t, a.log, a.sd.L2Cfg, a.sd.RollupCfg.Genesis.L1, jwtPath)
	engCl := engine.EngineClient(t, a.sd.RollupCfg)
	daMgr := damgr.NewAltDA(a.log, *a.sd.DaCfg, a.storage, engCl)
	modEng := NewModEngine(engCl, daMgr)

	l1F, err := sources.NewL1Client(a.miner.RPCClient(), a.log, nil, sources.L1ClientDefaultConfig(a.sd.RollupCfg, false, sources.RPCKindBasic))
	require.NoError(t, err)

	dataSrc := derive.NewDASourceFactory(a.log, a.sd.RollupCfg, l1F, daMgr)

	return NewL2Verifier(t, a.log, l1F, modEng, a.sd.RollupCfg, &sync.Config{}, dataSrc)
}

func (a *L2AltDA) ActSequencerIncludeTx(t Testing) {
	a.alice.L2.ActResetTxOpts(t)
	a.alice.L2.ActSetTxToAddr(&a.dp.Addresses.Bob)(t)
	a.alice.L2.ActMakeTx(t)

	a.sequencer.ActL2PipelineFull(t)

	a.sequencer.ActL2StartBlock(t)
	a.engine.ActL2IncludeTx(a.alice.Address())(t)
	a.sequencer.ActL2EndBlock(t)
}

func (a *L2AltDA) ActNewL2Tx(t Testing) {
	a.ActSequencerIncludeTx(t)

	a.batcher.ActL2BatchBuffer(t)
	a.batcher.ActL2ChannelClose(t)
	a.batcher.ActL2BatchSubmit(t, func(tx *types.DynamicFeeTx) {
		a.lastComm = tx.Data
	})

	a.miner.ActL1StartBlock(3)(t)
	a.miner.ActL1IncludeTx(a.dp.Addresses.Batcher)(t)
	a.miner.ActL1EndBlock(t)

	a.lastCommBn = a.miner.l1Chain.CurrentBlock().Number.Uint64()
}

func (a *L2AltDA) ActDeleteLastInput(t Testing) {
	a.storage.Client.DeleteData(a.lastComm)
}

func (a *L2AltDA) ActChallengeLastInput(t Testing) {
	a.ActChallengeInput(t, a.lastComm, a.lastCommBn)

	a.log.Info("challenged last input", "block", a.lastCommBn)
}

func (a *L2AltDA) ActChallengeInput(t Testing, comm []byte, bn uint64) {
	bondValue, err := a.contract.BondSize(&bind.CallOpts{})
	require.NoError(t, err)

	txOpts, err := bind.NewKeyedTransactorWithChainID(a.dp.Secrets.Alice, a.sd.L1Cfg.Config.ChainID)
	require.NoError(t, err)

	txOpts.Value = bondValue
	_, err = a.contract.Deposit(txOpts)
	require.NoError(t, err)

	a.miner.ActL1StartBlock(3)(t)
	a.miner.ActL1IncludeTx(a.alice.Address())(t)
	a.miner.ActL1EndBlock(t)

	txOpts, err = bind.NewKeyedTransactorWithChainID(a.dp.Secrets.Alice, a.sd.L1Cfg.Config.ChainID)
	require.NoError(t, err)

	commArray := (*[32]byte)(comm)
	_, err = a.contract.Challenge(txOpts, big.NewInt(int64(bn)), *commArray)
	require.NoError(t, err)

	a.miner.ActL1StartBlock(3)(t)
	a.miner.ActL1IncludeTx(a.alice.Address())(t)
	a.miner.ActL1EndBlock(t)
}

func (a *L2AltDA) ActExpireLastInput(t Testing) {
	reorgWindow := a.sd.DaCfg.ChallengeWindow + a.sd.DaCfg.ResolveWindow
	for a.miner.l1Chain.CurrentBlock().Number.Uint64() <= a.lastCommBn+reorgWindow {
		a.miner.ActL1StartBlock(3)(t)
		a.miner.ActL1EndBlock(t)
	}
}

func (a *L2AltDA) ActResolveLastChallenge(t Testing) {
	input, err := a.storage.GetPreImage(t.Ctx(), a.lastComm)
	require.NoError(t, err)

	txOpts, err := bind.NewKeyedTransactorWithChainID(a.dp.Secrets.Alice, a.sd.L1Cfg.Config.ChainID)

	commArray := (*[32]byte)(a.lastComm)
	_, err = a.contract.Resolve(txOpts, big.NewInt(int64(a.lastCommBn)), *commArray, input)
	require.NoError(t, err)

	a.miner.ActL1StartBlock(3)(t)
	a.miner.ActL1IncludeTx(a.alice.Address())(t)
	a.miner.ActL1EndBlock(t)
}

func (a *L2AltDA) ActL1Blocks(t Testing, n uint64) {
	for i := uint64(0); i < n; i++ {
		a.miner.ActL1StartBlock(3)(t)
		a.miner.ActL1EndBlock(t)
	}
}

func (a *L2AltDA) GetLastTxBlock(t Testing) *types.Block {
	rcpt, err := a.engine.EthClient().TransactionReceipt(t.Ctx(), a.alice.L2.lastTxHash)
	require.NoError(t, err)
	blk, err := a.engine.EthClient().BlockByHash(t.Ctx(), rcpt.BlockHash)
	require.NoError(t, err)
	return blk
}

// - Scenario 1: Commitment is challenged but never resolved, chain reorgs when challenge window expires.
// - Scenario 2: Commitment is challenged after sequencer derived the chain but data disappears. A verifier
// derivation pipeline stalls until the challenge is resolved and then resumes with data from the contract.
// - Scenario 3: DA storage service goes offline while sequencer keeps making blocks. When storage comes back online, it should be able to catch up.
// - Scenario 4: Commitment is challenged but with a wrong block number.

func TestAltDA_ChallengeExpired(gt *testing.T) {
	t := NewDefaultTesting(gt)
	p := &e2eutils.TestParams{
		MaxSequencerDrift:   2,
		SequencerWindowSize: 4,
		ChannelTimeout:      4,
		L1BlockTime:         3,
		DaChallengeWindow:   6,
		DaResolveWindow:     6,
	}
	log := testlog.Logger(t, log.LvlDebug)
	harness := NewL2AltDA(log, p, t)

	// generate enough initial l1 blocks to have a safe head.
	harness.ActL1Blocks(t, 5)

	// Include a new l2 transaction, submitting an input commitment to the l1.
	harness.ActNewL2Tx(t)

	// Challenge the input commitment on the l1 challenge contract.
	harness.ActChallengeLastInput(t)

	blk := harness.GetLastTxBlock(t)

	// catch up the sequencer derivation pipeline with the new l1 blocks.
	harness.sequencer.ActL2PipelineFull(t)

	// create enough l1 blocks to expire the resolve window.
	harness.ActExpireLastInput(t)

	// catch up the sequencer derivation pipeline with the new l1 blocks.
	harness.sequencer.ActL2PipelineFull(t)

	// make sure that the safe head was correctly updated on the engine.
	l2Safe, err := harness.engCl.L2BlockRefByLabel(t.Ctx(), eth.Safe)
	require.NoError(t, err)
	require.Equal(t, uint64(11), l2Safe.Number)

	newBlk, err := harness.engine.EthClient().BlockByNumber(t.Ctx(), blk.Number())
	require.NoError(t, err)

	// reorg happened even though data was available
	require.NotEqual(t, blk.Hash(), newBlk.Hash())

	// now delete the data from the storage service so it is not available at all
	// to the verifier derivation pipeline.
	harness.ActDeleteLastInput(t)

	syncStatus := harness.sequencer.SyncStatus()

	// verifier is able to sync with expired missing data
	verifier := harness.NewVerifier(t)
	verifier.ActL2PipelineFull(t)

	verifSyncStatus := verifier.SyncStatus()

	require.Equal(t, syncStatus.SafeL2, verifSyncStatus.SafeL2)
}

func TestAltDA_ChallengeResolved(gt *testing.T) {
	t := NewDefaultTesting(gt)
	p := &e2eutils.TestParams{
		MaxSequencerDrift:   2,
		SequencerWindowSize: 4,
		ChannelTimeout:      4,
		L1BlockTime:         3,
		DaChallengeWindow:   6,
		DaResolveWindow:     6,
	}
	log := testlog.Logger(t, log.LvlDebug)
	harness := NewL2AltDA(log, p, t)

	// include a new l2 transaction, submitting an input commitment to the l1.
	harness.ActNewL2Tx(t)

	// generate 3 l1 blocks.
	harness.ActL1Blocks(t, 3)

	// challenge the input commitment for that l2 transaction on the l1 challenge contract.
	harness.ActChallengeLastInput(t)

	// catch up sequencer derivatio pipeline.
	// this syncs the latest event within the AltDA manager.
	harness.sequencer.ActL2PipelineFull(t)

	// resolve the challenge on the l1 challenge contract.
	harness.ActResolveLastChallenge(t)

	// catch up the sequencer derivation pipeline with the new l1 blocks.
	// this syncs the resolved status and input data within the AltDA manager.
	harness.sequencer.ActL2PipelineFull(t)

	// delete the data from the storage service so it is not available at all
	// to the verifier derivation pipeline.
	harness.ActDeleteLastInput(t)

	syncStatus := harness.sequencer.SyncStatus()

	// new verifier is able to sync and resolve the input from calldata
	verifier := harness.NewVerifier(t)
	verifier.ActL2PipelineFull(t)

	verifSyncStatus := verifier.SyncStatus()

	require.Equal(t, syncStatus.SafeL2, verifSyncStatus.SafeL2)
}

func TestAltDA_StorageError(gt *testing.T) {
	t := NewDefaultTesting(gt)
	p := &e2eutils.TestParams{
		MaxSequencerDrift:   2,
		SequencerWindowSize: 4,
		ChannelTimeout:      4,
		L1BlockTime:         3,
		DaChallengeWindow:   6,
		DaResolveWindow:     6,
	}
	log := testlog.Logger(t, log.LvlDebug)
	harness := NewL2AltDA(log, p, t)

	// include a new l2 transaction, submitting an input commitment to the l1.
	harness.ActNewL2Tx(t)

	txBlk := harness.GetLastTxBlock(t)

	// mock a storage client error when trying to get the pre-image.
	// this simulates the storage service going offline for example.
	harness.storage.ActGetPreImageFail()

	// try to derive the l2 chain from the submitted inputs commitments.
	// the storage call will fail the first time then succeed.
	harness.sequencer.ActL2PipelineFull(t)

	// sequencer derivation was able to sync to latest l1 origin
	syncStatus := harness.sequencer.SyncStatus()
	require.Equal(t, uint64(1), syncStatus.SafeL2.Number)
	require.Equal(t, txBlk.Hash(), syncStatus.SafeL2.Hash)
}

func TestAltDA_ChallengeBadBlockNumber(gt *testing.T) {
	t := NewDefaultTesting(gt)
	p := &e2eutils.TestParams{
		MaxSequencerDrift:   2,
		SequencerWindowSize: 4,
		ChannelTimeout:      4,
		L1BlockTime:         3,
		DaChallengeWindow:   6,
		DaResolveWindow:     6,
	}
	log := testlog.Logger(t, log.LvlDebug)
	harness := NewL2AltDA(log, p, t)

	// generate 3 blocks of l1 chain
	harness.ActL1Blocks(t, 3)

	// include a new transaction on l2
	harness.ActNewL2Tx(t)

	// move the l1 chain so the challenge window expires
	harness.ActExpireLastInput(t)

	// catch up derivation
	harness.sequencer.ActL2PipelineFull(t)

	// challenge the input but with a wrong block number
	// in the current challenge window
	harness.ActChallengeInput(t, harness.lastComm, 14)

	// catch up derivation
	harness.sequencer.ActL2PipelineFull(t)

	// da mgr should not have save the challenge
	_, found := harness.daMgr.GetChallengeStatus(harness.lastComm, 14)
	require.False(t, found)
}
