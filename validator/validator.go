package validator

import (
	"context"
	"errors"
	"fmt"
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
	assertionEvents        <-chan protocol.AssertionChainEvent
	stateUpdateEvents      <-chan *statemanager.StateAdvancedEvent
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
	}
	for _, o := range opts {
		o(v)
	}
	// TODO: Prefer an API where the caller provides the channel and we can subscribe to all challenge and
	// assertion chain events. Provide the ability to specify the type of the subscription.
	v.assertionEvents = v.protocol.Subscribe(ctx)
	v.stateUpdateEvents = v.stateManager.SubscribeStateEvents(ctx)
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
			v.submitLeafCreation(ctx)
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
				localCommitment, err := v.stateManager.HistoryCommitmentAtHeight(ctx, ev.Commitment.Height)
				if err != nil {
					log.WithError(err).Error("Could not get history commitment")
					continue
				}
				if v.isCorrectLeaf(localCommitment, ev) {
					v.defendLeaf(ev)
				} else {
					v.challengeLeaf(localCommitment, ev)
				}
			case *protocol.StartChallengeEvent:
				v.processChallengeStart(ctx, ev)
			default:
				panic("not a recognized assertion chain event")
			}
		case <-ctx.Done():
			return
		}
	}
}

func (v *Validator) submitLeafCreation(ctx context.Context) {
	prevAssertion := v.protocol.LatestConfirmed()
	currentCommit := v.stateManager.LatestHistoryCommitment(ctx)
	commit := protocol.StateCommitment{
		Height: currentCommit.Height,
		State:  currentCommit.Merkle,
	}
	logFields := logrus.Fields{
		"name":                  v.name,
		"latestConfirmedHeight": fmt.Sprintf("%+v", prevAssertion.SequenceNum),
		"leafHeight":            commit.Height,
		"leafCommitmentMerkle":  util.FormatHash(commit.State),
	}
	_, err := v.protocol.CreateLeaf(prevAssertion, commit, v.address)
	switch {
	case errors.Is(err, protocol.ErrVertexAlreadyExists):
		log.WithFields(logFields).Debug("Vertex already exists, unable to create new leaf")
		return
	case errors.Is(err, protocol.ErrInvalid):
		log.WithFields(logFields).Debug("Tried to create a leaf with an older commitment")
		return
	case err != nil:
		log.WithError(err).Error("Could not create leaf")
		return
	}
	log.WithFields(logFields).Info("Submitted leaf creation")
	return
}

func (v *Validator) isFromSelf(ev *protocol.CreateLeafEvent) bool {
	return v.address == ev.Staker
}

func (v *Validator) isCorrectLeaf(localCommitment *util.HistoryCommitment, ev *protocol.CreateLeafEvent) bool {
	return localCommitment.Hash() == ev.Commitment.Hash()
}

func (v *Validator) defendLeaf(ev *protocol.CreateLeafEvent) {
	logFields := logrus.Fields{}
	if name, ok := v.knownValidatorNames[ev.Staker]; ok {
		logFields["createdBy"] = name
	}
	logFields["height"] = ev.Commitment.Height
	logFields["commitmentMerkle"] = util.FormatHash(ev.Commitment.State)
	log.WithFields(logFields).Info("New leaf matches local state")
}

func (v *Validator) challengeLeaf(localCommitment *util.HistoryCommitment, ev *protocol.CreateLeafEvent) {
	logFields := logrus.Fields{}
	if name, ok := v.knownValidatorNames[ev.Staker]; ok {
		logFields["disagreesWith"] = name
	}
	logFields["correctCommitmentHeight"] = localCommitment.Height
	logFields["badCommitmentHeight"] = ev.Commitment.Height
	logFields["correctCommitmentMerkle"] = util.FormatHash(localCommitment.Merkle)
	logFields["badCommitmentMerkle"] = util.FormatHash(ev.Commitment.State)
	log.WithFields(logFields).Warn("Disagreed with created leaf")
}

func (v *Validator) processChallengeStart(ctx context.Context, ev *protocol.StartChallengeEvent) {

}
