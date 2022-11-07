package validator

import (
	"context"
	"fmt"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	statemanager "github.com/OffchainLabs/new-rollup-exploration/state-manager"
	"github.com/ethereum/go-ethereum/common"
)

type Opt = func(val *Validator)

type Validator struct {
	protocol             protocol.OnChainProtocol
	stateManager         statemanager.Manager
	assertionEvents      <-chan protocol.AssertionChainEvent
	stateUpdateEvents    <-chan *statemanager.StateAdvancedEvent
	address              common.Address
	maliciousProbability float64
}

func WithMaliciousProbability(p float64) Opt {
	return func(val *Validator) {
		val.maliciousProbability = p
	}
}

func WithAddress(addr common.Address) Opt {
	return func(val *Validator) {
		val.address = addr
	}
}

func New(
	ctx context.Context,
	onChainProtocol protocol.OnChainProtocol,
	stateManager statemanager.Manager,
	opts ...Opt,
) (*Validator, error) {
	v := &Validator{
		protocol:     onChainProtocol,
		stateManager: stateManager,
		address:      common.Address{},
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
	go v.listenForStateUpdates(ctx)
}

func (v *Validator) listenForStateUpdates(ctx context.Context) {
	for {
		select {
		case stateUpdated := <-v.stateUpdateEvents:
			fmt.Printf(
				"Received a state update event from state manager: height %d, %#x\n",
				stateUpdated.HistoryCommitment.Height,
				stateUpdated.HistoryCommitment.Hash(),
			)
			fmt.Println("Submitting leaf creation event to chain")
			stateCommit := protocol.StateCommitment{
				Height: stateUpdated.HistoryCommitment.Height,
				State:  stateUpdated.HistoryCommitment.Hash(),
			}
			prevAssertion := v.protocol.LatestConfirmed()

			fmt.Printf("Latest confirmed is %d and %#x\n", prevAssertion.SequenceNum, prevAssertion.StateCommitment.Hash())

			// TODO: Simulate posting leaf events with some jitter delay, validators will have latency
			// in posting created leaves to the protocol.
			assertion, err := v.protocol.CreateLeaf(prevAssertion, stateCommit, v.address)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Created assertion %+v\n", assertion)
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
				if v.isCorrectLeaf(ctx, ev) {
					fmt.Printf("Got a correct leaf at height %d, %#x\n", ev.Commitment.Height, ev.Commitment.Hash())
					v.defendLeaf(ev)
				} else {
					fmt.Printf("WRONG leaf at height %d, %#x\n", ev.Commitment.Height, ev.Commitment.Hash())
					v.challengeLeaf(ev)
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

func (v *Validator) isCorrectLeaf(ctx context.Context, ev *protocol.CreateLeafEvent) bool {
	localCommitment, err := v.stateManager.HistoryCommitmentAtHeight(ctx, ev.Commitment.Height)
	if err != nil {
		panic(err)
	}
	return localCommitment != ev.Commitment.Hash()
}

func (v *Validator) defendLeaf(ev *protocol.CreateLeafEvent) {
}

func (v *Validator) challengeLeaf(ev *protocol.CreateLeafEvent) {
}

func (v *Validator) processChallengeStart(ctx context.Context, ev *protocol.StartChallengeEvent) {

}
