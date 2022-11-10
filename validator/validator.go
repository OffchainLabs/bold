package validator

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	statemanager "github.com/OffchainLabs/new-rollup-exploration/state-manager"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("prefix", "validator")

type Opt = func(val *Validator)

type Validator struct {
	protocol                    protocol.OnChainProtocol
	stateManager                statemanager.Manager
	assertionEvents             chan protocol.AssertionChainEvent
	stateUpdateEvents           chan *statemanager.StateAdvancedEvent
	address                     common.Address
	name                        string
	knownValidatorNames         map[common.Address]string
	createdLeaves               map[common.Hash]*protocol.Assertion
	assertionsLock              sync.RWMutex
	assertionsByParentStateRoot map[common.Hash][]*protocol.Assertion
	leavesLock                  sync.RWMutex
	createLeafInterval          time.Duration
	maliciousProbability        float64
	chaosMonkeyProbability      float64
}

func WithMaliciousProbability(p float64) Opt {
	return func(val *Validator) {
		val.maliciousProbability = p
	}
}

func WithName(name string) Opt {
	return func(val *Validator) {
		val.name = name
	}
}

func WithAddress(addr common.Address) Opt {
	return func(val *Validator) {
		val.address = addr
	}
}

func WithKnownValidators(vals map[common.Address]string) Opt {
	return func(val *Validator) {
		val.knownValidatorNames = vals
	}
}

func WithCreateLeafEvery(d time.Duration) Opt {
	return func(val *Validator) {
		val.createLeafInterval = d
	}
}

func New(
	ctx context.Context,
	onChainProtocol protocol.OnChainProtocol,
	stateManager statemanager.Manager,
	opts ...Opt,
) (*Validator, error) {
	v := &Validator{
		protocol:                    onChainProtocol,
		stateManager:                stateManager,
		address:                     common.Address{},
		createLeafInterval:          5 * time.Second,
		assertionEvents:             make(chan protocol.AssertionChainEvent, 1),
		stateUpdateEvents:           make(chan *statemanager.StateAdvancedEvent, 1),
		createdLeaves:               make(map[common.Hash]*protocol.Assertion),
		assertionsByParentStateRoot: make(map[common.Hash][]*protocol.Assertion),
	}
	for _, o := range opts {
		o(v)
	}
	// TODO: Prefer an API where the caller provides the channel and we can subscribe to all challenge and
	// assertion chain events. Provide the ability to specify the type of the subscription.
	v.protocol.SubscribeChainEvents(ctx, v.assertionEvents)
	v.stateManager.SubscribeStateEvents(ctx, v.stateUpdateEvents)
	return v, nil
}

func (v *Validator) Start(ctx context.Context) {
	go v.listenForAssertionEvents(ctx)
	go v.prepareLeafCreationPeriodically(ctx)
}

// TODO: Simulate posting leaf events with some jitter delay, validators will have
// latency in posting created leaves to the protocol.
func (v *Validator) prepareLeafCreationPeriodically(ctx context.Context) {
	ticker := time.NewTicker(v.createLeafInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// Keep track of the leaf we created so we can confirm it as no rival in the future.
			leaf := v.submitLeafCreation(ctx)
			if leaf == nil {
				continue
			}
			v.leavesLock.Lock()
			v.createdLeaves[leaf.StateCommitment.StateRoot] = leaf
			v.leavesLock.Unlock()
			go v.confirmLeafAfterChallengePeriod(leaf)
		case <-ctx.Done():
			return
		}
	}
}

func (v *Validator) listenForAssertionEvents(ctx context.Context) {
	for {
		select {
		case genericEvent := <-v.assertionEvents:
			switch ev := genericEvent.(type) {
			case *protocol.CreateLeafEvent:
				go v.processLeafCreationEvent(ctx, ev)
			case *protocol.StartChallengeEvent:
				go v.processChallengeStart(ctx, ev)
			case *protocol.ConfirmEvent:
				log.WithField(
					"sequenceNum", ev.SeqNum,
				).Info("Leaf with sequence number confirmed on-chain")
			default:
				log.WithField("ev", fmt.Sprintf("%+v", ev)).Error("Not a recognized chain event")
			}
		case <-ctx.Done():
			return
		}
	}
}

func (v *Validator) submitLeafCreation(ctx context.Context) *protocol.Assertion {
	randDuration := rand.Int31n(2000) // 2000 ms simulating latency in submitting leaf creation.
	time.Sleep(time.Millisecond * time.Duration(randDuration))

	// Ensure that we only build on a valid parent from this validator's perspective.
	// TODO: Instead of iterating over all assertions, validator should load up all created assertions since
	// the latest confirmed one in prod and update that list as faster way of looking up valid parents.
	// the validator should also have ready access to historical commitments to make sure it can select
	// the valid parent based on its commitment state root.
	// TODO: Turn into its own method and test thoroughly.
	parentAssertion, err := v.findLatestValidAssertion(ctx)
	if err != nil {
		log.WithError(err).Error("Could not find valid parent assertion to build leaf upon")
		return nil
	}

	// TODO: Fix! State commit and history commit are not the same thing.
	currentCommit := v.stateManager.LatestHistoryCommitment(ctx)
	stateCommit := protocol.StateCommitment{
		Height:    currentCommit.Height,
		StateRoot: currentCommit.Merkle,
	}
	logFields := logrus.Fields{
		"name":                       v.name,
		"latestValidParentHeight":    fmt.Sprintf("%+v", parentAssertion.StateCommitment.Height),
		"latestValidParentStateRoot": util.FormatHash(parentAssertion.StateCommitment.StateRoot),
		"leafHeight":                 currentCommit.Height,
		"leafCommitmentMerkle":       util.FormatHash(currentCommit.Merkle),
	}
	leaf, err := v.protocol.CreateLeaf(parentAssertion, stateCommit, v.address)
	switch {
	case errors.Is(err, protocol.ErrVertexAlreadyExists):
		log.WithFields(logFields).Debug("Vertex already exists, unable to create new leaf")
		return nil
	case errors.Is(err, protocol.ErrInvalid):
		log.WithFields(logFields).Debug("Tried to create a leaf with an older commitment")
		return nil
	case err != nil:
		log.WithError(err).Error("Could not create leaf")
		return nil
	}
	log.WithFields(logFields).Info("Submitted leaf creation")
	return leaf
}

// Finds the latest valid assertion a validator should build their new leaves upon. This starts from
// the latest confirmed assertion and makes it down the tree to the latest assertion that has a state
// root matching in the validator's database.
func (v *Validator) findLatestValidAssertion(ctx context.Context) (*protocol.Assertion, error) {
	latestValidParent := v.protocol.LatestConfirmed()
	for s := latestValidParent.SequenceNum; s < v.protocol.NumAssertions(); s++ {
		a, err := v.protocol.AssertionBySequenceNumber(s)
		if err != nil {
			return nil, err
		}
		if v.stateManager.HasStateRoot(ctx, a.StateCommitment.StateRoot) {
			latestValidParent = a
		}
	}
	return latestValidParent, nil
}

// For a leaf created by a validator, we confirm the leaf has no rival after the challenge deadline has passed.
// This function is meant to be ran as a goroutine for each leaf created by the validator.
func (v *Validator) confirmLeafAfterChallengePeriod(leaf *protocol.Assertion) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(v.protocol.ChallengePeriodLength()))
	defer cancel()
	<-ctx.Done()
	logFields := logrus.Fields{
		"height":      leaf.StateCommitment.Height,
		"sequenceNum": leaf.SequenceNum,
	}
	if err := leaf.ConfirmNoRival(); err != nil {
		log.WithError(err).WithFields(logFields).Warn("Could not confirm that created leaf had no rival")
		return
	}
	log.WithFields(logFields).Info("Confirmed leaf passed challenge period successfully on-chain")
}

func (v *Validator) processLeafCreationEvent(ctx context.Context, ev *protocol.CreateLeafEvent) {
	// TODO: Ignore all events from self, not just CreateLeafEvent.
	if v.isFromSelf(ev) {
		return
	}

	// Detect if there is a fork, then decide if we want to challenge.
	// We check if the parent assertion has > 1 child.
	assertion, err := v.protocol.AssertionBySequenceNumber(ev.SeqNum)
	if err != nil {
		log.WithError(err).Error("Could ont get assertion")
	}
	v.assertionsLock.Lock()
	defer v.assertionsLock.Unlock()
	if assertion.Prev().IsEmpty() {
		v.assertionsByParentStateRoot[common.Hash{}] = append(
			v.assertionsByParentStateRoot[common.Hash{}],
			assertion,
		)
	} else {
		parentAssertion := assertion.Prev().OpenKnownFull()
		v.assertionsByParentStateRoot[parentAssertion.StateCommitment.StateRoot] = append(
			v.assertionsByParentStateRoot[parentAssertion.StateCommitment.StateRoot],
			assertion,
		)
	}
	var hasForked bool

	localCommitment, err := v.stateManager.HistoryCommitmentAtHeight(ctx, ev.StateCommitment.Height)
	if err != nil {
		log.WithError(err).Error("Could not get history commitment")
		return
	}
	commit := protocol.StateCommitment{
		Height:    localCommitment.Height,
		StateRoot: localCommitment.Merkle,
	}
	_ = commit
}

func (v *Validator) processChallengeStart(ctx context.Context, ev *protocol.StartChallengeEvent) {
	// Checks if the challenge has to do with a vertex we created.
	challengedAssertion, err := v.protocol.AssertionBySequenceNumber(ev.ParentSeqNum)
	if err != nil {
		log.WithError(err).Error("Could not get assertion by sequence number")
	}
	v.leavesLock.RLock()
	defer v.leavesLock.RUnlock()
	leaf, ok := v.createdLeaves[challengedAssertion.StateCommitment.StateRoot]
	if !ok {
		// TODO: Act on the honest vertices even if this challenge does not have to do with us.
		return
	}
	challengerName := "unknown-name"
	if !leaf.Staker.IsEmpty() {
		if name, ok := v.knownValidatorNames[leaf.Staker.OpenKnownFull()]; ok {
			challengerName = name
		} else {
			challengerName = leaf.Staker.OpenKnownFull().Hex()
		}
	}
	log.WithFields(logrus.Fields{
		"name":                 v.name,
		"challenger":           challengerName,
		"challengingStateRoot": util.FormatHash(leaf.StateCommitment.StateRoot),
		"challengingHeight":    leaf.StateCommitment.Height,
	}).Warn("Received challenge for a created leaf!")
}

// TODO: Defend a leaf if it is not created by us, but is a valid leaf from our perspective.
func (v *Validator) defendLeaf(ev *protocol.CreateLeafEvent) {
	logFields := logrus.Fields{}
	if name, ok := v.knownValidatorNames[ev.Staker]; ok {
		logFields["createdBy"] = name
	}
	logFields["name"] = v.name
	logFields["height"] = ev.StateCommitment.Height
	logFields["stateRoot"] = util.FormatHash(ev.StateCommitment.StateRoot)
	log.WithFields(logFields).Info("New leaf matches local state")
}

// Initiates a challenge on a created leaf.
// TODO: We can only challenge if there are two competing leaves under a parent vertex. We challenge the parent.
// We should only perform this action if there is a fork, so the validator needs to be aware if a leaf creation
// event is an actual fork or not.
func (v *Validator) challengeLeaf(
	ctx context.Context,
	localCommitment protocol.StateCommitment,
	ev *protocol.CreateLeafEvent,
) {
	assertion, err := v.protocol.AssertionBySequenceNumber(ev.SeqNum)
	if err != nil {
		log.WithError(err).Error("Could not retrieve assertion for challenge")
		return
	}
	logFields := logrus.Fields{}
	logFields["disagreesWith"] = "unknown-name"
	if name, ok := v.knownValidatorNames[ev.Staker]; ok {
		logFields["disagreesWith"] = name
	}
	logFields["name"] = v.name
	logFields["correctCommitmentHeight"] = localCommitment.Height
	logFields["badCommitmentHeight"] = ev.StateCommitment.Height
	logFields["correctStateRoot"] = util.FormatHash(localCommitment.StateRoot)
	logFields["badStateRoot"] = util.FormatHash(ev.StateCommitment.StateRoot)
	log.WithFields(logFields).Infof("Disagreed with other posted leaf, submitting challenge to protocol")

	challenge, err := assertion.CreateChallenge(ctx)
	if err != nil {
		log.WithError(err).Error("Could not issue challenge")
		return
	}
	log.WithFields(logFields).Infof("Submitted challenge: %+v", challenge)
}

func (v *Validator) isFromSelf(ev *protocol.CreateLeafEvent) bool {
	return v.address == ev.Staker
}

func (v *Validator) isCorrectLeaf(localCommitment protocol.StateCommitment, ev *protocol.CreateLeafEvent) bool {
	return localCommitment.Hash() == ev.StateCommitment.Hash()
}
