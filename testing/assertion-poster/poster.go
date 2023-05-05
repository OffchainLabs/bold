package assertionposter

import (
	"context"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol/sol-implementation"
	"github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("prefix", "assertion-poster")

type Poster struct {
	chain                  protocol.Protocol
	stateManager           statemanager.Manager
	postAssertionsInterval time.Duration
	validatorName          string
}

type Opt func(p *Poster)

func WithPostAssertionsInterval(d time.Duration) Opt {
	return func(p *Poster) {
		p.postAssertionsInterval = d
	}
}

func WithName(name string) Opt {
	return func(p *Poster) {
		p.validatorName = name
	}
}

func New(
	chain protocol.Protocol,
	stateManager statemanager.Manager,
	opts ...Opt,
) *Poster {
	p := &Poster{
		chain:                  chain,
		stateManager:           stateManager,
		postAssertionsInterval: time.Minute,
		validatorName:          "unknown",
	}
	for _, o := range opts {
		o(p)
	}
	return p
}

func (p *Poster) Start(ctx context.Context) {
	if _, err := p.PostLatestAssertion(ctx); err != nil {
		log.WithError(err).Error("Could not submit latest assertion to L1")
	}
	ticker := time.NewTicker(p.postAssertionsInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if _, err := p.PostLatestAssertion(ctx); err != nil {
				log.WithError(err).Error("Could not submit latest assertion to L1")
			}
		case <-ctx.Done():
			return
		}
	}
}

// Posts the latest claim of the Node's L2 state as an assertion to the L1 protocol smart contracts.
// TODO: Include leaf creation validity conditions which are more complex than this.
// For example, a validator must include messages from the inbox that were not included
// by the last validator in the last leaf's creation.
func (v *Poster) PostLatestAssertion(ctx context.Context) (protocol.Assertion, error) {
	// Ensure that we only build on a valid parent from this validator's perspective.
	// the validator should also have ready access to historical commitments to make sure it can select
	// the valid parent based on its commitment state root.
	parentAssertionSeq, err := v.findLatestValidAssertion(ctx)
	if err != nil {
		return nil, err
	}
	parentAssertion, err := v.chain.AssertionBySequenceNum(ctx, parentAssertionSeq)
	if err != nil {
		return nil, err
	}
	parentAssertionStateHash, err := parentAssertion.StateHash()
	if err != nil {
		return nil, err
	}
	parentAssertionState, err := v.stateManager.AssertionExecutionState(ctx, parentAssertionStateHash)
	if err != nil {
		return nil, err
	}
	assertionToCreate, err := v.stateManager.LatestAssertionCreationData(ctx)
	if err != nil {
		return nil, err
	}
	assertion, err := v.chain.CreateAssertion(
		ctx,
		parentAssertionState,
		assertionToCreate.State,
		assertionToCreate.InboxMaxCount,
	)
	switch {
	case errors.Is(err, solimpl.ErrAlreadyExists):
		return nil, errors.Wrap(err, "assertion already exists, was unable to post")
	case err != nil:
		return nil, err
	}
	assertionState, err := assertion.StateHash()
	if err != nil {
		return nil, err
	}
	logFields := logrus.Fields{
		"name":               v.validatorName,
		"parentStateHash":    util.Trunc(parentAssertionStateHash.Bytes()),
		"assertionStateHash": util.Trunc(assertionState.Bytes()),
	}
	log.WithFields(logFields).Info("Submitted latest L2 state claim as an assertion to L1")
	return assertion, nil
}

// Finds the latest valid assertion sequence num a validator should build their new leaves upon. This walks
// down from the number of assertions in the protocol down until it finds
// an assertion that we have a state commitment for.
func (p *Poster) findLatestValidAssertion(ctx context.Context) (protocol.AssertionSequenceNumber, error) {
	numAssertions, err := p.chain.NumAssertions(ctx)
	if err != nil {
		return 0, err
	}
	latestConfirmedFetched, err := p.chain.LatestConfirmed(ctx)
	if err != nil {
		return 0, err
	}
	latestConfirmed := latestConfirmedFetched.SeqNum()
	for s := protocol.AssertionSequenceNumber(numAssertions); s > latestConfirmed; s-- {
		a, err := p.chain.AssertionBySequenceNum(ctx, s)
		if err != nil {
			return 0, err
		}
		stateHash, err := a.StateHash()
		if err != nil {
			return 0, err
		}
		if p.stateManager.HasStateCommitment(ctx, util.StateCommitment{
			StateRoot: stateHash,
		}) {
			return a.SeqNum(), nil
		}
	}
	return latestConfirmed, nil
}
