// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package challengemanager

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/bold/assertions"
	protocol "github.com/offchainlabs/bold/chain-abstraction"
	watcher "github.com/offchainlabs/bold/challenge-manager/chain-watcher"
	edgetracker "github.com/offchainlabs/bold/challenge-manager/edge-tracker"
	"github.com/offchainlabs/bold/challenge-manager/types"
	"github.com/offchainlabs/bold/containers/option"
	l2stateprovider "github.com/offchainlabs/bold/layer2-state-provider"
	"github.com/offchainlabs/bold/solgen/go/challengeV2gen"
	"github.com/offchainlabs/bold/solgen/go/rollupgen"
	"github.com/offchainlabs/bold/testing/mocks"
	"github.com/offchainlabs/bold/testing/setup"
	customTime "github.com/offchainlabs/bold/time"
	"github.com/stretchr/testify/require"
)

var _ = types.RivalHandler(&Manager{})

func TestEdgeTracker_Act(t *testing.T) {
	ctx := context.Background()
	createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{}, setup.WithMockOneStepProver())
	require.NoError(t, err)

	tkr, _ := setupEdgeTrackersForBisection(t, ctx, createdData, option.None[uint64]())
	err = tkr.Act(ctx)
	require.NoError(t, err)
	require.Equal(t, edgetracker.EdgeBisecting, tkr.CurrentState())

	err = tkr.Act(ctx)
	require.NoError(t, err)
	require.Equal(t, edgetracker.EdgeAwaitingChallengeCompletion, tkr.CurrentState())

	err = tkr.Act(ctx)
	require.NoError(t, err)
	require.Equal(t, edgetracker.EdgeAwaitingChallengeCompletion, tkr.CurrentState())
}

func TestEdgeTracker_Act_ConfirmedByTime(t *testing.T) {
	ctx := context.Background()
	createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{}, setup.WithMockOneStepProver())
	require.NoError(t, err)

	chalManager := createdData.Chains[0].SpecChallengeManager()
	chalPeriodBlocks, err := chalManager.ChallengePeriodBlocks(ctx)
	require.NoError(t, err)

	// Delay the evil root edge creation by a challenge period.
	delayEvilRootEdgeCreation := option.Some(chalPeriodBlocks)
	honestTracker, evilTracker := setupEdgeTrackersForBisection(t, ctx, createdData, delayEvilRootEdgeCreation)

	honestEdgeOpt, err := chalManager.GetEdge(ctx, honestTracker.EdgeId())
	require.NoError(t, err)
	require.Equal(t, false, honestEdgeOpt.IsNone())

	evilEdgeOpt, err := chalManager.GetEdge(ctx, evilTracker.EdgeId())
	require.NoError(t, err)
	require.Equal(t, false, evilEdgeOpt.IsNone())

	// Expect our edge to be confirmed right away.
	err = honestTracker.Act(ctx)
	require.NoError(t, err)
	require.Equal(t, edgetracker.EdgeAwaitingChallengeCompletion, honestTracker.CurrentState())
	require.Equal(t, true, honestTracker.ShouldDespawn(ctx))
}

type verifiedHonestMock struct {
	*mocks.MockSpecEdge
}

func (verifiedHonestMock) Honest() {}

func Test_getEdgeTrackers(t *testing.T) {
	ctx := context.Background()

	v, m, s := setupValidator(t)
	edge := &mocks.MockSpecEdge{}
	edge.On("Id").Return(protocol.EdgeId{Hash: common.BytesToHash([]byte("foo"))})
	edge.On("GetReversedChallengeLevel").Return(protocol.ChallengeLevel(2))
	edge.On("MutualId").Return(protocol.MutualId{})
	edge.On("OriginId").Return(protocol.OriginId{})
	edge.On("CreatedAtBlock").Return(uint64(1), nil)
	parentAssertionHash := protocol.AssertionHash{Hash: common.BytesToHash([]byte("par"))}
	assertionHash := protocol.AssertionHash{Hash: common.BytesToHash([]byte("bar"))}
	edge.On("ClaimId").Return(option.Some(protocol.ClaimId(assertionHash.Hash)))
	edge.On("AssertionHash", ctx).Return(assertionHash, nil)
	edge.On("StartCommitment").Return(protocol.Height(0), common.Hash{})
	edge.On("EndCommitment").Return(protocol.Height(0), common.Hash{})
	edge.On("GetChallengeLevel").Return(protocol.ChallengeLevel(0))
	m.On("ReadAssertionCreationInfo", ctx, assertionHash).Return(&protocol.AssertionCreatedInfo{
		BeforeState: rollupgen.AssertionState{
			GlobalState: rollupgen.GlobalState{
				U64Vals: [2]uint64{1, 0},
			},
		},
		AfterState: rollupgen.AssertionState{
			GlobalState: rollupgen.GlobalState{
				U64Vals: [2]uint64{100, 0},
			},
		},
		ParentAssertionHash: parentAssertionHash.Hash,
	}, nil)
	m.On("ReadAssertionCreationInfo", ctx, parentAssertionHash).Return(&protocol.AssertionCreatedInfo{
		InboxMaxCount: big.NewInt(100),
	}, nil)
	s.On("ExecutionStateMsgCount", ctx, &protocol.ExecutionState{}).Return(uint64(1), nil)

	require.NoError(t, v.watcher.AddVerifiedHonestEdge(ctx, verifiedHonestMock{edge}))

	trk, err := v.getTrackerForEdge(ctx, protocol.SpecEdge(edge))
	require.NoError(t, err)

	require.Equal(t, l2stateprovider.Batch(1), l2stateprovider.Batch(trk.AssertionInfo().FromState.Batch))
	require.Equal(t, l2stateprovider.Batch(100), trk.AssertionInfo().BatchLimit)
}

type assertionManagerConfig struct {
	c          protocol.AssertionChain
	ep         l2stateprovider.ExecutionProvider
	b          *setup.SimulatedBackendWrapper
	rollupAddr common.Address
	name       string
	mode       types.Mode
}

func setupDefaultAssertionManager(conf assertionManagerConfig, t *testing.T) (assertionManager, error) {
	t.Helper()

	m, err := assertions.NewManager(
		conf.c,
		conf.ep,
		conf.name,
		conf.mode,
	)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func setupEdgeTrackersForBisection(
	t *testing.T,
	ctx context.Context,
	createdData *setup.CreatedValidatorFork,
	delayEvilRootEdgeCreationByBlocks option.Option[uint64],
) (*edgetracker.Tracker, *edgetracker.Tracker) {
	t.Helper()
	honestAsserter, err := setupDefaultAssertionManager(
		assertionManagerConfig{
			c:          createdData.Chains[0],
			ep:         createdData.HonestStateManager,
			b:          createdData.Backend,
			rollupAddr: createdData.Addrs.Rollup,
			name:       "alice",
			mode:       types.MakeMode,
		}, t)
	require.NoError(t, err)
	honestValidator, err := New(
		ctx,
		createdData.Chains[0],
		createdData.HonestStateManager,
		honestAsserter,
		createdData.Addrs.Rollup,
		WithName("alice"),
		WithMode(types.MakeMode),
	)
	require.NoError(t, err)

	evilAsserter, err := setupDefaultAssertionManager(
		assertionManagerConfig{
			c:          createdData.Chains[1],
			ep:         createdData.EvilStateManager,
			b:          createdData.Backend,
			rollupAddr: createdData.Addrs.Rollup,
			name:       "bob",
			mode:       types.MakeMode,
		}, t)
	require.NoError(t, err)
	evilValidator, err := New(
		ctx,
		createdData.Chains[1],
		createdData.EvilStateManager,
		evilAsserter,
		createdData.Addrs.Rollup,
		WithName("bob"),
		WithMode(types.MakeMode),
	)
	require.NoError(t, err)

	honestEdge, _, _, _, err := honestValidator.addBlockChallengeLevelZeroEdge(ctx, createdData.Leaf1)
	require.NoError(t, err)

	// If we specify an optional amount of blocks to delay the evil root edge creation by, do so
	// by committing blocks to the simulated backend.
	if !delayEvilRootEdgeCreationByBlocks.IsNone() {
		delay := delayEvilRootEdgeCreationByBlocks.Unwrap()
		for i := uint64(0); i < delay; i++ {
			createdData.Backend.Commit()
		}
	}

	evilEdge, _, _, _, err := evilValidator.addBlockChallengeLevelZeroEdge(ctx, createdData.Leaf2)
	require.NoError(t, err)

	// Check unrivaled statuses.
	hasRival, err := honestEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, false, !hasRival)

	chalManager := createdData.Chains[0].SpecChallengeManager()
	managerBindings, err := challengeV2gen.NewEdgeChallengeManagerCaller(chalManager.Address(), createdData.Backend)
	require.NoError(t, err)
	numBigStepLevelsRaw, err := managerBindings.NUMBIGSTEPLEVEL(createdData.Chains[0].GetCallOptsWithDesiredRpcHeadBlockNumber(&bind.CallOpts{Context: ctx}))
	require.NoError(t, err)
	numBigStepLevels := numBigStepLevelsRaw

	honestWatcher, err := watcher.New(honestValidator.chain, honestValidator, honestValidator.stateManager, createdData.Backend, time.Second, numBigStepLevels, "alice", nil, honestValidator.assertionConfirmingInterval, honestValidator.averageTimeForBlockCreation, nil)
	require.NoError(t, err)
	honestValidator.watcher = honestWatcher
	assertionInfo := &l2stateprovider.AssociatedAssertionMetadata{
		FromState:      protocol.GoGlobalState{Batch: 0, PosInBatch: 0},
		BatchLimit:     1,
		WasmModuleRoot: common.Hash{},
	}
	tracker1, err := edgetracker.New(
		ctx,
		honestEdge,
		honestValidator.chain,
		createdData.HonestStateManager,
		honestWatcher,
		honestValidator,
		assertionInfo,
		edgetracker.WithTimeReference(customTime.NewArtificialTimeReference()),
		edgetracker.WithValidatorName(honestValidator.name),
	)
	require.NoError(t, err)

	evilWatcher, err := watcher.New(evilValidator.chain, evilValidator, evilValidator.stateManager, createdData.Backend, time.Second, numBigStepLevels, "alice", nil, evilValidator.assertionConfirmingInterval, evilValidator.averageTimeForBlockCreation, nil)
	require.NoError(t, err)
	evilValidator.watcher = evilWatcher
	tracker2, err := edgetracker.New(
		ctx,
		evilEdge,
		evilValidator.chain,
		createdData.EvilStateManager,
		evilWatcher,
		evilValidator,
		assertionInfo,
		edgetracker.WithTimeReference(customTime.NewArtificialTimeReference()),
		edgetracker.WithValidatorName(evilValidator.name),
	)
	require.NoError(t, err)

	require.NoError(t, honestWatcher.AddVerifiedHonestEdge(ctx, honestEdge))
	_, err = honestWatcher.AddEdge(ctx, evilEdge)
	require.NoError(t, err)
	require.NoError(t, evilWatcher.AddVerifiedHonestEdge(ctx, evilEdge))
	_, err = evilWatcher.AddEdge(ctx, honestEdge)
	require.NoError(t, err)

	return tracker1, tracker2
}

func setupValidator(t *testing.T) (*Manager, *mocks.MockProtocol, *mocks.MockStateManager) {
	t.Helper()
	p := &mocks.MockProtocol{}
	ctx := context.Background()
	cm := &mocks.MockSpecChallengeManager{}
	p.On("CurrentChallengeManager", ctx).Return(cm, nil)
	p.On("SpecChallengeManager").Return(cm)
	cm.On("NumBigSteps", ctx).Return(uint8(1), nil)
	s := &mocks.MockStateManager{}
	cfg, err := setup.ChainsWithEdgeChallengeManager(setup.WithMockOneStepProver())
	require.NoError(t, err)
	p.On("Backend").Return(cfg.Backend, nil)
	a, err := setupDefaultAssertionManager(
		assertionManagerConfig{
			c:          cfg.Chains[0],
			ep:         s,
			b:          cfg.Backend,
			rollupAddr: cfg.Addrs.Rollup,
			name:       "alice",
			mode:       types.MakeMode,
		}, t)
	require.NoError(t, err)
	v, err := New(context.Background(), p, s, a, cfg.Addrs.Rollup, WithMode(types.MakeMode))
	require.NoError(t, err)
	return v, p, s
}
