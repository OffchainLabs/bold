package validator

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/rollupgen"
	statemanager "github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("prefix", "validator")

type Opt = func(val *Validator)

// Validator defines a validator client instances in the assertion protocol, which will be
// an active participant in interacting with the on-chain contracts.
type Validator struct {
	chain                                  protocol.Protocol
	chalManagerAddr                        common.Address
	rollupAddr                             common.Address
	rollup                                 *rollupgen.RollupCore
	rollupFilterer                         *rollupgen.RollupCoreFilterer
	chalManager                            *challengeV2gen.EdgeChallengeManagerFilterer
	backend                                bind.ContractBackend
	stateManager                           statemanager.Manager
	address                                common.Address
	name                                   string
	createdAssertions                      map[common.Hash]protocol.Assertion
	assertionsLock                         sync.RWMutex
	sequenceNumbersByParentStateCommitment map[common.Hash][]protocol.AssertionSequenceNumber
	assertions                             map[protocol.AssertionSequenceNumber]protocol.Assertion
	timeRef                                util.TimeReference
	edgeTrackerWakeInterval                time.Duration
	newAssertionCheckInterval              time.Duration
}

// WithName is a human-readable identifier for this validator client for logging purposes.
func WithName(name string) Opt {
	return func(val *Validator) {
		val.name = name
	}
}

// WithAddress gives a staker address to the validator.
func WithAddress(addr common.Address) Opt {
	return func(val *Validator) {

		val.address = addr
	}
}

// WithTimeReference adds a time reference interface to the validator.
func WithTimeReference(ref util.TimeReference) Opt {
	return func(val *Validator) {
		val.timeRef = ref
	}
}

// WithEdgeTrackerWakeInterval specifies how often each edge tracker goroutine will
// act on its responsibilities.
func WithEdgeTrackerWakeInterval(d time.Duration) Opt {
	return func(val *Validator) {
		val.edgeTrackerWakeInterval = d
	}
}

// WithNewAssertionCheckInterval specifies how often handle assertions goroutine will
// act on its responsibilities.
func WithNewAssertionCheckInterval(d time.Duration) Opt {
	return func(val *Validator) {
		val.newAssertionCheckInterval = d
	}
}

// New sets up a validator client instances provided a protocol, state manager,
// and additional options.
func New(
	ctx context.Context,
	chain protocol.Protocol,
	backend bind.ContractBackend,
	stateManager statemanager.Manager,
	rollupAddr common.Address,
	opts ...Opt,
) (*Validator, error) {
	v := &Validator{
		backend:                                backend,
		chain:                                  chain,
		stateManager:                           stateManager,
		address:                                common.Address{},
		createdAssertions:                      make(map[common.Hash]protocol.Assertion),
		sequenceNumbersByParentStateCommitment: make(map[common.Hash][]protocol.AssertionSequenceNumber),
		assertions:                             make(map[protocol.AssertionSequenceNumber]protocol.Assertion),
		timeRef:                                util.NewRealTimeReference(),
		rollupAddr:                             rollupAddr,
		edgeTrackerWakeInterval:                time.Millisecond * 100,
		newAssertionCheckInterval:              time.Second,
	}
	for _, o := range opts {
		o(v)
	}
	genesisAssertion, err := v.chain.AssertionBySequenceNum(ctx, 1)
	if err != nil {
		return nil, err
	}
	chalManager, err := v.chain.SpecChallengeManager(ctx)
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
	v.rollup = rollup
	v.rollupFilterer = rollupFilterer
	v.assertions[0] = genesisAssertion
	v.chalManagerAddr = chalManagerAddr
	v.chalManager = chalManagerFilterer
	return v, nil
}

func (v *Validator) Start(ctx context.Context) {
	go v.pollForAssertions(ctx)
	log.WithField(
		"address",
		v.address.Hex(),
	).Info("Started validator client")
}

// Processes new leaf creation events from the protocol that were not initiated by self.
func (v *Validator) onLeafCreated(
	ctx context.Context,
	assertion protocol.Assertion,
) error {
	assertionStateHash, err := assertion.StateHash()
	if err != nil {
		return err
	}
	log.WithFields(logrus.Fields{
		"name":      v.name,
		"stateHash": fmt.Sprintf("%#x", assertionStateHash),
	}).Info("New assertion appended to protocol")

	// Keep track of assertions by parent state root to more easily detect forks.
	assertionPrevSeqNum, err := assertion.PrevSeqNum()
	if err != nil {
		return err
	}
	prevAssertion, err := v.chain.AssertionBySequenceNum(ctx, assertionPrevSeqNum)
	if err != nil {
		return err
	}

	v.assertionsLock.Lock()
	key, err := prevAssertion.StateHash()
	if err != nil {
		return err
	}
	v.sequenceNumbersByParentStateCommitment[key] = append(
		v.sequenceNumbersByParentStateCommitment[key],
		assertion.SeqNum(),
	)
	hasForked := len(v.sequenceNumbersByParentStateCommitment[key]) > 1
	v.assertionsLock.Unlock()

	// If this leaf's creation has not triggered fork, we have nothing else to do.
	if !hasForked {
		log.Info("No fork detected in assertion tree upon leaf creation")
		return nil
	}

	return v.challengeAssertion(ctx, assertion)
}
