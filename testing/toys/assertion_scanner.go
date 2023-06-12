package toys

import (
	"context"

	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	retry "github.com/OffchainLabs/challenge-protocol-v2/runtime"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type AssertionScanner struct{}

func (ac *AssertionScanner) Scan(ctx context.Context) {
	it, err := filterer.FilterEdgeAdded(filterOpts, nil, nil, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err = it.Close(); err != nil {
			log.WithError(err).Error("Could not close filter iterator")
		}
	}()
	for it.Next() {
		if it.Error() != nil {
			return errors.Wrapf(
				err,
				"got iterator error when scanning edge creations from block %d to %d",
				filterOpts.Start,
				*filterOpts.End,
			)
		}
		_, processErr := retry.UntilSucceeds(ctx, func() (bool, error) {
			return true, w.processEdgeAddedEvent(ctx, it.Event)
		})
		if processErr != nil {
			return processErr
		}
	}
	return nil
}

// Processes new leaf creation events from the protocol that were not initiated by self.
func (v *Manager) onLeafCreated(
	ctx context.Context,
	assertion protocol.Assertion,
) error {
	log.WithFields(logrus.Fields{
		"name": v.name,
	}).Info("New assertion appended to protocol")
	isFirstChild, err := assertion.IsFirstChild()
	if err != nil {
		return err
	}
	// If this leaf is the first child, we have nothing else to do.
	if isFirstChild {
		log.Info("No fork detected in assertion tree upon leaf creation")
		return nil
	}
	return v.challengeAssertion(ctx, assertion.Id())
}
