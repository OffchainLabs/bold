package validator

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
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
	protocol               protocol.OnChainProtocol
	stateManager           statemanager.Manager
	assertionEvents        chan protocol.AssertionChainEvent
	stateUpdateEvents      chan *statemanager.StateAdvancedEvent
	address                common.Address
	name                   string
	knownValidatorNames    map[common.Address]string
	createLeafInterval     time.Duration
	maliciousProbability   float64
	chaosMonkeyProbability float64
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
		protocol:           onChainProtocol,
		stateManager:       stateManager,
		address:            common.Address{},
		createLeafInterval: 5 * time.Second,
		assertionEvents:    make(chan protocol.AssertionChainEvent, 1),
		stateUpdateEvents:  make(chan *statemanager.StateAdvancedEvent, 1),
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
				// TODO: Ignore all events from self, not just CreateLeafEvent.
				if v.isFromSelf(ev) {
					continue
				}
				localCommitment, err := v.stateManager.HistoryCommitmentAtHeight(ctx, ev.StateCommitment.Height)
				if err != nil {
					log.WithError(err).Error("Could not get history commitment")
					continue
				}
				commit := protocol.StateCommitment{
					Height:    localCommitment.Height,
					StateRoot: localCommitment.Merkle,
				}
				if v.isCorrectLeaf(commit, ev) {
					v.defendLeaf(ev)
				} else {
					v.challengeLeaf(commit, ev)
				}
			case *protocol.StartChallengeEvent:
				v.processChallengeStart(ctx, ev)
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
	prevAssertion := v.protocol.LatestConfirmed()
	var latestValidParent *protocol.Assertion
	if v.stateManager.HasStateRoot(ctx, prevAssertion.StateCommitment.StateRoot) {
		latestValidParent = prevAssertion
	}
	for s := prevAssertion.SequenceNum; s < v.protocol.NumAssertions(); s++ {
		a, err := v.protocol.AssertionBySequenceNumber(s)
		if err != nil {
			log.WithError(err).Error("Could not retrieve assertion by sequence number")
			return nil
		}
		if v.stateManager.HasStateRoot(ctx, a.StateCommitment.StateRoot) {
			latestValidParent = prevAssertion
		}
	}
	if latestValidParent == nil {
		latestValidParent = v.protocol.Genesis()
	}

	// TODO: Fix! State commit and history commit are not the same thing.
	currentCommit := v.stateManager.LatestHistoryCommitment(ctx)
	stateCommit := protocol.StateCommitment{
		Height:    currentCommit.Height,
		StateRoot: currentCommit.Merkle,
	}
	logFields := logrus.Fields{
		"name":                  v.name,
		"latestConfirmedHeight": fmt.Sprintf("%+v", prevAssertion.SequenceNum),
		"leafHeight":            currentCommit.Height,
		"leafCommitmentMerkle":  util.FormatHash(currentCommit.Merkle),
	}
	leaf, err := v.protocol.CreateLeaf(prevAssertion, stateCommit, v.address)
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
		log.WithError(err).WithFields(logFields).Error("Could not confirm that created leaf had no rival")
		return
	}
	log.WithFields(logFields).Info("Confirmed leaf passed challenge period successfully on-chain")
}

func (v *Validator) isFromSelf(ev *protocol.CreateLeafEvent) bool {
	return v.address == ev.Staker
}

func (v *Validator) isCorrectLeaf(localCommitment protocol.StateCommitment, ev *protocol.CreateLeafEvent) bool {
	return localCommitment.Hash() == ev.StateCommitment.Hash()
}

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

func (v *Validator) challengeLeaf(localCommitment protocol.StateCommitment, ev *protocol.CreateLeafEvent) {
	logFields := logrus.Fields{}
	if name, ok := v.knownValidatorNames[ev.Staker]; ok {
		logFields["disagreesWith"] = name
	}
	logFields["name"] = v.name
	logFields["correctCommitmentHeight"] = localCommitment.Height
	logFields["badCommitmentHeight"] = ev.StateCommitment.Height
	logFields["correctStateRoot"] = util.FormatHash(localCommitment.StateRoot)
	logFields["badStateRoot"] = util.FormatHash(ev.StateCommitment.StateRoot)
	log.WithFields(logFields).Warn("Disagreed with created leaf")
}

func (v *Validator) processChallengeStart(ctx context.Context, ev *protocol.StartChallengeEvent) {

}
