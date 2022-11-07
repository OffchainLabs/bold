package validator

import (
	"context"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	statemanager "github.com/OffchainLabs/new-rollup-exploration/state-manager"
)

type Opt = func(val *Validator) error

type Validator struct {
	protocol             protocol.OnChainProtocol
	stateManager         statemanager.Manager
	assertionEvents      <-chan protocol.AssertionChainEvent
	maliciousProbability float64
}

func WithMaliciousProbability(p float64) Opt {
	return func(val *Validator) error {
		val.maliciousProbability = p
		return nil
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
	}
	for _, o := range opts {
		if err := o(v); err != nil {
			return nil, err
		}
	}
	// TODO: Prefer an API where the caller provides the channel and we can subscribe to all challenge and
	// assertion chain events. Provide the ability to specify the type of the subscription.
	v.assertionEvents = v.protocol.Subscribe(ctx)
	return v, nil
}

func (v *Validator) Validate(ctx context.Context) {
	for {
		select {
		case genericEvent := <-v.assertionEvents:
			switch ev := genericEvent.(type) {
			case *protocol.CreateLeafEvent:
				if v.isCorrectLeaf(ev) {
					v.defendLeaf(ev)
				} else {
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

func (v *Validator) isCorrectLeaf(ev *protocol.CreateLeafEvent) bool {
	return true
}

func (v *Validator) defendLeaf(ev *protocol.CreateLeafEvent) {
}

func (v *Validator) challengeLeaf(ev *protocol.CreateLeafEvent) {
}

func (v *Validator) processChallengeStart(ctx context.Context, ev *protocol.StartChallengeEvent) {

}

