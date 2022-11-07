package validator

import (
	"context"
	"fmt"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
)

type Opt = func(val *Validator) error

type Validator struct {
	protocol             protocol.OnChainProtocol
	assertionEvents      <-chan protocol.AssertionChainEvent
	maliciousProbability float64
}

func WithMaliciousProbability(p float64) Opt {
	return func(val *Validator) error {
		val.maliciousProbability = p
		return nil
	}
}

func New(ctx context.Context, onChainProtocol protocol.OnChainProtocol, opts ...Opt) (*Validator, error) {
	v := &Validator{
		protocol: onChainProtocol,
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
				fmt.Println(ev)
			default:
				panic("not a recognized assertion chain event")
			}
		case <-ctx.Done():
			return
		}
	}
}
