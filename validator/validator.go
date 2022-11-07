package validator

import (
	"context"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
)

type Opt = func(val *Validator) error

type Validator struct {
	protocol             protocol.OnChainProtocol
	maliciousProbability float64
}

func WithOnChainProtocol(p protocol.OnChainProtocol) Opt {
	return func(val *Validator) error {
		val.protocol = p
		return nil
	}
}

func WithMaliciousProbability(p float64) Opt {
	return func(val *Validator) error {
		val.maliciousProbability = p
		return nil
	}
}

func New(ctx context.Context, opts ...Opt) (*Validator, error) {
	v := &Validator{}
	for _, o := range opts {
		if err := o(v); err != nil {
			return nil, err
		}
	}
	return v, nil
}
