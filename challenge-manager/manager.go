// Package challengemanager includes the main entrypoint for setting up a BOLD
// challenge manager instance and challenging assertions onchain.
//
// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE
package challengemanager

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/rpc"
	apibackend "github.com/offchainlabs/bold/api/backend"
	"github.com/offchainlabs/bold/api/db"
	"github.com/offchainlabs/bold/api/server"
	protocol "github.com/offchainlabs/bold/chain-abstraction"
	watcher "github.com/offchainlabs/bold/challenge-manager/chain-watcher"
	edgetracker "github.com/offchainlabs/bold/challenge-manager/edge-tracker"
	"github.com/offchainlabs/bold/challenge-manager/types"
	"github.com/offchainlabs/bold/containers/events"
	"github.com/offchainlabs/bold/containers/option"
	"github.com/offchainlabs/bold/containers/threadsafe"
	l2stateprovider "github.com/offchainlabs/bold/layer2-state-provider"
	retry "github.com/offchainlabs/bold/runtime"
	"github.com/offchainlabs/bold/solgen/go/challengeV2gen"
	"github.com/offchainlabs/bold/solgen/go/rollupgen"
	utilTime "github.com/offchainlabs/bold/time"
	"github.com/offchainlabs/bold/util/stopwaiter"
	"github.com/pkg/errors"
)

var (
	challengeSubmittedCounter = metrics.NewRegisteredCounter("arb/validator/scanner/challenge_submitted", nil)
)

type Opt = func(val *Manager)

type assertionManager interface {
	Start(context.Context)
	StopAndWait()
	LatestAgreedAssertion() protocol.AssertionHash
	LayerZeroHeights(context.Context) (*protocol.LayerZeroHeights, error)
	SetRivalHandler(types.RivalHandler)
}

// Manager defines an offchain, challenge manager, which will be
// an active participant in interacting with the on-chain contracts.
type Manager struct {
	stopwaiter.StopWaiter
	chain                        protocol.Protocol
	chalManagerAddr              common.Address
	rollup                       *rollupgen.RollupCore
	rollupFilterer               *rollupgen.RollupCoreFilterer
	chalManager                  *challengeV2gen.EdgeChallengeManagerFilterer
	backend                      bind.ContractBackend
	assertionManager             assertionManager
	watcher                      *watcher.Watcher
	client                       *rpc.Client
	stateManager                 l2stateprovider.Provider
	address                      common.Address
	name                         string
	timeRef                      utilTime.Reference
	trackedEdgeIds               *threadsafe.Map[protocol.EdgeId, *edgetracker.Tracker]
	batchIndexForAssertionCache  *threadsafe.LruMap[protocol.AssertionHash, l2stateprovider.AssociatedAssertionMetadata]
	newBlockNotifier             *events.Producer[*gethtypes.Header]
	notifyOnNumberOfBlocks       uint64
	assertionConfirmingInterval  time.Duration
	averageTimeForBlockCreation  time.Duration
	mode                         types.Mode
	maxDelaySeconds              int
	claimedAssertionsInChallenge *threadsafe.LruSet[protocol.AssertionHash]
	headBlockSubscriptions       bool
	// API
	apiAddr string
	api     *server.Server
	apiDB   db.Database
}

// WithName is a human-readable identifier for this challenge manager for logging purposes.
func WithName(name string) Opt {
	return func(val *Manager) {
		val.name = name
	}
}

// WithAddress gives a staker address to the validator.
func WithAddress(addr common.Address) Opt {
	return func(val *Manager) {
		val.address = addr
	}
}

func WithAssertionConfirmingInterval(d time.Duration) Opt {
	return func(val *Manager) {
		val.assertionConfirmingInterval = d
	}
}

func WithAvgBlockCreationTime(d time.Duration) Opt {
	return func(val *Manager) {
		val.averageTimeForBlockCreation = d
	}
}

// Edges tick on every block received from the parent chain of the rollup, by default. Alternatively,
// they can be configured to tick every N blocks.
func WithTickEdgesOnNumberOfBlocks(n uint64) Opt {
	return func(val *Manager) {
		val.notifyOnNumberOfBlocks = n
	}
}

// WithMode specifies the mode of the challenge manager.
func WithMode(m types.Mode) Opt {
	return func(val *Manager) {
		val.mode = m
	}
}

// WithAPIEnabled specifies whether or not to enable the API and the address to listen on.
func WithAPIEnabled(addr string, db db.Database) Opt {
	return func(val *Manager) {
		val.apiAddr = addr
		val.apiDB = db
	}
}

func WithRPCClient(client *rpc.Client) Opt {
	return func(val *Manager) {
		val.client = client
	}
}

func WithHeadBlockSubscriptions() Opt {
	return func(val *Manager) {
		val.headBlockSubscriptions = true
	}
}

// New sets up a challenge manager instance provided a protocol, state manager, and additional options.
func New(
	ctx context.Context,
	chain protocol.Protocol,
	stateManager l2stateprovider.Provider,
	watcher *watcher.Watcher,
	assertionManager assertionManager,
	opts ...Opt,
) (*Manager, error) {
	m := &Manager{
		backend:                      chain.Backend(),
		chain:                        chain,
		stateManager:                 stateManager,
		assertionManager:             assertionManager,
		watcher:                      watcher,
		address:                      common.Address{},
		timeRef:                      utilTime.NewRealTimeReference(),
		trackedEdgeIds:               threadsafe.NewMap(threadsafe.MapWithMetric[protocol.EdgeId, *edgetracker.Tracker]("trackedEdgeIds")),
		batchIndexForAssertionCache:  threadsafe.NewLruMap(1000, threadsafe.LruMapWithMetric[protocol.AssertionHash, l2stateprovider.AssociatedAssertionMetadata]("batchIndexForAssertionCache")),
		notifyOnNumberOfBlocks:       1,
		newBlockNotifier:             events.NewProducer[*gethtypes.Header](),
		assertionConfirmingInterval:  time.Second * 10,
		averageTimeForBlockCreation:  time.Second * 12,
		headBlockSubscriptions:       false,
		claimedAssertionsInChallenge: threadsafe.NewLruSet(1000, threadsafe.LruSetWithMetric[protocol.AssertionHash]("claimedAssertionsInChallenge")),
	}
	for _, o := range opts {
		o(m)
	}
	chalManager := m.chain.SpecChallengeManager()
	chalManagerAddr := chalManager.Address()
	rollupAddr := m.chain.RollupAddress()

	rollup, err := rollupgen.NewRollupCore(rollupAddr, m.backend)
	if err != nil {
		return nil, err
	}
	rollupFilterer, err := rollupgen.NewRollupCoreFilterer(rollupAddr, m.backend)
	if err != nil {
		return nil, err
	}
	chalManagerFilterer, err := challengeV2gen.NewEdgeChallengeManagerFilterer(chalManagerAddr, m.backend)
	if err != nil {
		return nil, err
	}
	m.rollup = rollup
	m.rollupFilterer = rollupFilterer
	m.chalManagerAddr = chalManagerAddr
	m.chalManager = chalManagerFilterer
	m.watcher.SetEdgeManager(m)

	if m.apiAddr != "" {
		bknd := apibackend.NewBackend(m.apiDB, m.chain, m.watcher, m)
		srv, err2 := server.New(m.apiAddr, bknd)
		if err2 != nil {
			return nil, err2
		}
		m.api = srv
	}
	m.assertionManager.SetRivalHandler(m)
	log.Info("Setting up challenge manager", "name", m.name, "addreess", m.address, "rollup", rollupAddr)
	return m, nil
}

func (m *Manager) GetEdgeTracker(edgeId protocol.EdgeId) option.Option[*edgetracker.Tracker] {
	if m.IsTrackingEdge(edgeId) {
		return option.Some(m.trackedEdgeIds.Get(edgeId))
	}
	return option.None[*edgetracker.Tracker]()
}

// IsTrackingEdge returns true if we are currently tracking a specified edge id as an edge tracker goroutine.
func (m *Manager) IsTrackingEdge(edgeId protocol.EdgeId) bool {
	return m.trackedEdgeIds.Has(edgeId)
}

func (m *Manager) Database() db.Database {
	return m.apiDB
}

func (m *Manager) ChallengeManagerAddress() common.Address {
	return m.chalManagerAddr
}

func (m *Manager) BlockTimes() time.Duration {
	return m.averageTimeForBlockCreation
}

// MarkTrackedEdge marks an edge id as being tracked by our challenge manager.
func (m *Manager) MarkTrackedEdge(edgeId protocol.EdgeId, tracker *edgetracker.Tracker) {
	m.trackedEdgeIds.Put(edgeId, tracker)
}

func (m *Manager) RemovedTrackedEdge(edgeId protocol.EdgeId) {
	m.trackedEdgeIds.Delete(edgeId)
}

// Mode returns the mode of the challenge manager.
func (m *Manager) Mode() types.Mode {
	return m.mode
}

// IsChallengedAssertion checks if an assertion with a given hash has a challenge.
func (m *Manager) IsClaimedByChallenge(assertionHash protocol.AssertionHash) bool {
	return m.claimedAssertionsInChallenge.Has(assertionHash)
}

// MaxDelaySeconds returns the maximum number of seconds that the challenge manager will wait open a challenge.
func (m *Manager) MaxDelaySeconds() int {
	return m.maxDelaySeconds
}

// TrackEdge spawns an edge tracker for an edge if it is not currently being tracked.
func (m *Manager) TrackEdge(ctx context.Context, edge protocol.SpecEdge) error {
	if m.trackedEdgeIds.Has(edge.Id()) {
		return nil
	}
	trk, err := m.getTrackerForEdge(ctx, edge)
	if err != nil {
		return err
	}
	m.LaunchThread(trk.Spawn)
	return nil
}

// Gets an edge tracker for an edge by retrieving its associated assertion creation info.
func (m *Manager) getTrackerForEdge(ctx context.Context, edge protocol.SpecEdge) (*edgetracker.Tracker, error) {
	// Retry until you get the previous assertion Hash.
	assertionHash, err := retry.UntilSucceeds(ctx, func() (protocol.AssertionHash, error) {
		return edge.AssertionHash(ctx)
	})
	if err != nil {
		return nil, err
	}
	blockChallengeRootEdge, err := m.watcher.HonestBlockChallengeRootEdge(ctx, assertionHash)
	if err != nil {
		return nil, err
	}
	if blockChallengeRootEdge.ClaimId().IsNone() {
		return nil, fmt.Errorf(
			"block challenge root edge %#x did not have a claim id for challenged assertion %#x",
			blockChallengeRootEdge.Id(),
			assertionHash,
		)
	}
	claimedAssertionId := blockChallengeRootEdge.ClaimId().Unwrap()

	// Smart caching to avoid querying the same assertion number and creation info multiple times.
	// Edges in the same challenge should have the same creation info.
	cachedHeightAndInboxMsgCount, ok := m.batchIndexForAssertionCache.TryGet(protocol.AssertionHash{Hash: common.Hash(claimedAssertionId)})
	var edgeTrackerAssertionInfo l2stateprovider.AssociatedAssertionMetadata
	if !ok {
		assertionCreationInfo, creationErr := retry.UntilSucceeds(ctx, func() (*protocol.AssertionCreatedInfo, error) {
			return m.chain.ReadAssertionCreationInfo(ctx, protocol.AssertionHash{Hash: common.Hash(claimedAssertionId)})
		})
		if creationErr != nil {
			return nil, creationErr
		}
		prevCreationInfo, prevCreationErr := retry.UntilSucceeds(ctx, func() (*protocol.AssertionCreatedInfo, error) {
			return m.chain.ReadAssertionCreationInfo(ctx, protocol.AssertionHash{Hash: assertionCreationInfo.ParentAssertionHash})
		})
		if prevCreationErr != nil {
			return nil, prevCreationErr
		}
		if prevCreationInfo.InboxMaxCount == nil {
			return nil, errors.New("prevCreationInfo.InboxMaxCount is nil")
		}
		if !prevCreationInfo.InboxMaxCount.IsUint64() {
			return nil, fmt.Errorf("inbox max count is not a uint64: %v", prevCreationInfo.InboxMaxCount)
		}
		fromState := protocol.GoGlobalStateFromSolidity(assertionCreationInfo.BeforeState.GlobalState)
		edgeTrackerAssertionInfo = l2stateprovider.AssociatedAssertionMetadata{
			FromState:            fromState,
			BatchLimit:           l2stateprovider.Batch(prevCreationInfo.InboxMaxCount.Uint64()),
			WasmModuleRoot:       prevCreationInfo.WasmModuleRoot,
			ClaimedAssertionHash: common.Hash(claimedAssertionId),
		}
		m.batchIndexForAssertionCache.Put(protocol.AssertionHash{Hash: common.Hash(claimedAssertionId)}, edgeTrackerAssertionInfo)
	} else {
		edgeTrackerAssertionInfo = cachedHeightAndInboxMsgCount
	}
	return retry.UntilSucceeds(ctx, func() (*edgetracker.Tracker, error) {
		return edgetracker.New(
			ctx,
			edge,
			m.chain,
			m.stateManager,
			m.watcher,
			m,
			&edgeTrackerAssertionInfo,
			edgetracker.WithTimeReference(m.timeRef),
			edgetracker.WithValidatorName(m.name),
		)
	})
}

func (m *Manager) Watcher() *watcher.Watcher {
	return m.watcher
}

func (m *Manager) ChallengeManager() *challengeV2gen.EdgeChallengeManagerFilterer {
	return m.chalManager
}

func (m *Manager) NewBlockSubscriber() *events.Producer[*gethtypes.Header] {
	return m.newBlockNotifier
}

func (m *Manager) Start(ctx context.Context) {
	m.StopWaiter.Start(ctx, m)
	log.Info("Started challenge manager",
		"validatorAddress", m.address.Hex(),
	)

	// Start the assertion manager.
	m.LaunchThread(m.assertionManager.Start)

	// Watcher tower and resolve modes don't monitor challenges.
	if m.mode == types.WatchTowerMode || m.mode == types.ResolveMode {
		return
	}

	// Start watching for parent chain block events in the background.
	m.LaunchThread(m.listenForBlockEvents)

	// Start watching for ongoing chain events in the background.
	m.LaunchThread(m.watcher.Start)

	if m.api != nil {
		m.LaunchThread(func(ctx context.Context) {
			if err := m.api.Start(ctx); err != nil {
				log.Error("Could not start API server",
					"address", m.apiAddr,
					"err", err,
				)
			}
		})
	}
}

func (m *Manager) StopAndWait() {
	m.StopWaiter.StopAndWait()
	m.assertionManager.StopAndWait()
	m.watcher.StopAndWait()
	if m.api != nil {
		m.api.StopAndWait()
	}
}

func (m *Manager) listenForBlockEvents(ctx context.Context) {
	// If the chain watcher has not yet scraped and caught up all BOLD
	// events up to the latest head, then we fire "block notification" events
	// every second. This will help the tracked edges act fast if we are
	// just starting up the validator or catching up to a challenge.
	m.fastTickWhileCatchingUp(ctx)

	// Then, once the watcher has reached the latest head, we
	// fire off a block notifications events normally.
	if m.headBlockSubscriptions {
		m.tickOnHeadBlockSubscriptions(ctx)
	} else {
		m.tickAtInterval(ctx)
	}
}

func (m *Manager) tickOnHeadBlockSubscriptions(ctx context.Context) {
	ch := make(chan *gethtypes.Header, 100)
	sub, err := m.chain.Backend().SubscribeNewHead(ctx, ch)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()
	numBlocksReceived := uint64(0)
	for {
		select {
		case header := <-ch:
			numBlocksReceived += 1
			// Only broadcast every N blocks received. This is important for Orbit chains
			// that have parent chains with very fast block times, such as Arbitrum One, as broadcasting
			// every 250ms would otherwise be too frequent.
			if numBlocksReceived%m.notifyOnNumberOfBlocks == 0 {
				m.newBlockNotifier.Broadcast(ctx, header)
			}
		case <-sub.Err():
		case <-ctx.Done():
			return
		}
	}
}

func (m *Manager) tickAtInterval(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 12)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
			m.newBlockNotifier.Broadcast(ctx, nil)
		}
	}
}

func (m *Manager) fastTickWhileCatchingUp(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
			m.newBlockNotifier.Broadcast(ctx, nil)
			if m.watcher.IsSynced() {
				return
			}
		}
	}
}

func (m *Manager) LatestConfirmedState(ctx context.Context) (protocol.GoGlobalState, error) {
	latestConfirmed, err := m.chain.LatestConfirmed(ctx, m.chain.GetCallOptsWithDesiredRpcHeadBlockNumber(&bind.CallOpts{Context: ctx}))
	if err != nil {
		return protocol.GoGlobalState{}, err
	}
	info, err := m.chain.ReadAssertionCreationInfo(ctx, latestConfirmed.Id())
	if err != nil {
		return protocol.GoGlobalState{}, err
	}
	return protocol.GoExecutionStateFromSolidity(info.AfterState).GlobalState, nil
}

func (m *Manager) LatestAgreedState(ctx context.Context) (protocol.GoGlobalState, error) {
	latestAgreedAssertion := m.assertionManager.LatestAgreedAssertion()
	info, err := m.chain.ReadAssertionCreationInfo(ctx, latestAgreedAssertion)
	if err != nil {
		return protocol.GoGlobalState{}, err
	}
	return protocol.GoExecutionStateFromSolidity(info.AfterState).GlobalState, nil
}

func (m *Manager) logChallengeConfigs(ctx context.Context) error {
	cm := m.chain.SpecChallengeManager()
	bigStepNum := cm.NumBigSteps()
	challengePeriodBlocks, err := cm.ChallengePeriodBlocks(ctx)
	if err != nil {
		return err
	}
	layerZeroHeights, err := m.assertionManager.LayerZeroHeights(ctx)
	if err != nil {
		return err
	}
	log.Info("Opening challenge with the following configuration",
		"address", cm.Address(),
		"bigStepNumber", bigStepNum,
		"challengePeriodBlocks", challengePeriodBlocks,
		"layerZeroHeights", layerZeroHeights,
	)
	return nil
}
