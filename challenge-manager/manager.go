package challengemanager

import (
	"context"
	"time"

	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	watcher "github.com/OffchainLabs/challenge-protocol-v2/challenge-manager/chain-watcher"
	"github.com/OffchainLabs/challenge-protocol-v2/containers/threadsafe"
	l2stateprovider "github.com/OffchainLabs/challenge-protocol-v2/layer2-state-provider"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/rollupgen"
	utilTime "github.com/OffchainLabs/challenge-protocol-v2/time"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("prefix", "challenge-manager")

type Opt = func(val *Manager)

// Manager defines an offchain, challenge manager, which will be
// an active participant in interacting with the on-chain contracts.
type Manager struct {
	chain                   protocol.Protocol
	chalManagerAddr         common.Address
	rollupAddr              common.Address
	rollup                  *rollupgen.RollupCore
	rollupFilterer          *rollupgen.RollupCoreFilterer
	chalManager             *challengeV2gen.EdgeChallengeManagerFilterer
	backend                 bind.ContractBackend
	stateManager            l2stateprovider.Provider
	address                 common.Address
	name                    string
	timeRef                 utilTime.Reference
	edgeTrackerWakeInterval time.Duration
	initialSyncCompleted    chan struct{}
	chainWatcherInterval    time.Duration
	watcher                 *watcher.Watcher
	trackedEdgeIds          *threadsafe.Set[protocol.EdgeId]
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

// WithEdgeTrackerWakeInterval specifies how often each edge tracker goroutine will
// act on its responsibilities.
func WithEdgeTrackerWakeInterval(d time.Duration) Opt {
	return func(val *Manager) {
		val.edgeTrackerWakeInterval = d
	}
}

// New sets up a challenge manager instance provided a protocol, state manager, and additional options.
func New(
	ctx context.Context,
	chain protocol.Protocol,
	backend bind.ContractBackend,
	stateManager l2stateprovider.Provider,
	rollupAddr common.Address,
	opts ...Opt,
) (*Manager, error) {
	m := &Manager{
		backend:                 backend,
		chain:                   chain,
		stateManager:            stateManager,
		address:                 common.Address{},
		timeRef:                 utilTime.NewRealTimeReference(),
		rollupAddr:              rollupAddr,
		edgeTrackerWakeInterval: time.Millisecond * 100,
		chainWatcherInterval:    time.Millisecond * 500,
		initialSyncCompleted:    make(chan struct{}),
		trackedEdgeIds:          threadsafe.NewSet[protocol.EdgeId](),
	}
	for _, o := range opts {
		o(m)
	}
	chalManager, err := m.chain.SpecChallengeManager(ctx)
	if err != nil {
		return nil, err
	}
	chalManagerAddr := chalManager.Address()

	rollup, err := rollupgen.NewRollupCore(rollupAddr, backend)
	if err != nil {
		return nil, err
	}
	rollupFilterer, err := rollupgen.NewRollupCoreFilterer(rollupAddr, backend)
	if err != nil {
		return nil, err
	}
	chalManagerFilterer, err := challengeV2gen.NewEdgeChallengeManagerFilterer(chalManagerAddr, backend)
	if err != nil {
		return nil, err
	}
	m.rollup = rollup
	m.rollupFilterer = rollupFilterer
	m.chalManagerAddr = chalManagerAddr
	m.chalManager = chalManagerFilterer
	m.watcher = watcher.New(m.chain, m.stateManager, backend, m.chainWatcherInterval, m.name)
	return m, nil
}

func (m *Manager) IsTrackingEdge(edgeId protocol.EdgeId) bool {
	return m.trackedEdgeIds.Has(edgeId)
}

func (m *Manager) MarkTrackedEdge(edgeId protocol.EdgeId) {
	m.trackedEdgeIds.Insert(edgeId)
}

func (m *Manager) Start(ctx context.Context) {
	log.WithField(
		"address",
		m.address.Hex(),
	).Info("Started challenge manager")

	// Start watching for ongoing chain events in the background.
	go m.watcher.Watch(ctx, m.initialSyncCompleted)

	// Then, wait until the chain event watcher has synced up with
	// all edges from the chain since the latest confirmed assertion up to the latest block number.
	if err := m.syncEdges(ctx); err != nil {
		log.WithError(err).Errorf("Could sync with edges")
	}
}

// waitForSync waits for a notificataion that initial sync of onchain edges is complete.
func (m *Manager) waitForSync(ctx context.Context) error {
	select {
	case <-m.initialSyncCompleted:
		return nil
	case <-ctx.Done():
		return errors.New("context closed, exiting goroutine")
	}
}
