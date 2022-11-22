package validator

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// TODO: A challenge ID in our challenge manager is a unique identifier for that challenge.
// Currently, it is using the parent assertion's state commitment hash as an ID,
// but this value may not work if we want to manage subchallenges in the same way.
type challengeID common.Hash

// Each challenge has a lifecycle we need to manage. A single challenge's entire lifecycle should
// be managed in a goroutine specific to that challenge. A challenge goroutine will exit if
//
// (a) A winner has been found (meaning all subchallenges are resolved), or
// (b) The chess clock for the challenge ends
//
// The validator has a challenge manager which is able to dispatch events from the global feed
// to specific challenge goroutines. A challenge goroutine is spawned upon receiving
// a ChallengeStarted event.
type challengeManager struct {
	lock                   sync.RWMutex
	chain                  protocol.ChainReadWriter
	challenges             map[challengeID]*challengeWorker
	challengeEventsBufSize int
}

type challengeWorker struct {
	challenge *protocol.Challenge
	leaves    []*protocol.ChallengeVertex
	events    chan protocol.ChallengeEvent
}

func newChallengeManager(chain protocol.ChainReadWriter) *challengeManager {
	return &challengeManager{
		chain:                  chain,
		challenges:             make(map[challengeID]*challengeWorker),
		challengeEventsBufSize: 100, // TODO: Make configurable.
	}
}

func (c *challengeManager) numChallenges() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.challenges)
}

// Dispatches an incoming generic challenge event to the respective challenge's worker.
func (c *challengeManager) dispatch(ev protocol.ChallengeEvent) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	ch, ok := c.challenges[challengeID(ev.ParentStateCommitmentHash())]
	if !ok {
		return
	}
	ch.events <- ev
}

func (c *challengeManager) spawnChallenge(ctx context.Context, challenge *protocol.Challenge) {
	c.lock.Lock()
	ch := make(chan protocol.ChallengeEvent, c.challengeEventsBufSize)
	id := challenge.ParentStateCommitment().Hash()
	worker := &challengeWorker{
		challenge: challenge,
		events:    ch,
	}
	c.challenges[challengeID(id)] = worker
	c.lock.Unlock()
	go worker.runChallengeLifecycle(ctx, c, ch)
}

func (w *challengeWorker) runChallengeLifecycle(
	ctx context.Context,
	manager *challengeManager,
	evs chan protocol.ChallengeEvent,
) {
	// Manage chess clock moves for the validator.
	// Listen for challenge completion, win
	// Cleanup the challenge goroutine once done.
	defer close(evs)
	ownChessClock := time.Now()
	_ = ownChessClock
	// Figure out if we are at a one-step fork, and then depending on who's turn it is,
	// spawn a subchallenge (BigStepChallenge).
	manager.chain.Tx(func(at *protocol.ActiveTx, ocp protocol.OnChainProtocol) error {
		return nil
	})
	vertex, _ := w.challenge.AddLeaf(nil, nil, util.HistoryCommitment{})
	_ = vertex
	for {
		select {
		case genericEvent := <-evs:
			switch ev := genericEvent.(type) {
			case *protocol.ChallengeLeafEvent:
				//go func() {
				//if err := c.onChallengeLeafAdded(ctx, ev); err != nil {
				//log.WithError(err).Error("Could not process leaf creation event")
				//}
				//}()
			case *protocol.ChallengeBisectEvent:
				//go func() {
				//if err := c.onBisectionEvent(ctx, ev); err != nil {
				//log.WithError(err).Error("Could not process challenge start event")
				//}
				//}()
			case *protocol.ChallengeMergeEvent:
				//go func() {
				//if err := c.onMergeEvent(ctx, ev); err != nil {
				//log.WithError(err).Error("Could not process challenge start event")
				//}
				//}()
			default:
				log.WithField("ev", fmt.Sprintf("%+v", ev)).Error("Not a recognized chain event")
			}
		case <-ctx.Done():
			return
		}
	}
}

// If a leaf has been added, we then check if we should add a competing leaf, bisect, or merge
// and then perform the corresponding action.
func (c *challengeManager) onChallengeLeafAdded(ctx context.Context, ev *protocol.ChallengeLeafEvent) error {
	return nil
}

// If a bisection has occurred, we need to determine if we should merge or bisect.
func (c *challengeManager) onBisectionEvent(ctx context.Context, ev *protocol.ChallengeBisectEvent) error {
	return nil
}

// If we have seen a merge event, what happens...??
func (c *challengeManager) onMergeEvent(ctx context.Context, ev *protocol.ChallengeMergeEvent) error {
	return nil
}

// Process new challenge creation events from the protocol that were not initiated by self.
func (v *Validator) onChallengeStarted(ctx context.Context, ev *protocol.StartChallengeEvent) error {
	if ev == nil {
		return nil
	}
	// Ignore challenges initiated by self.
	if v.isFromSelf(ev.Challenger) {
		return nil
	}
	// Checks if the challenge has to do with a vertex we created.
	v.leavesLock.RLock()
	leaf, ok := v.createdLeaves[ev.ParentStateCommitment.StateRoot]
	if !ok {
		v.leavesLock.RUnlock()
		// TODO: Act on the honest vertices even if this challenge does not have to do with us by
		// keeping track of associated challenge vertices' clocks and acting if the associated
		// staker we agree with is not performing their responsibilities on time. As an honest
		// validator, we should participate in confirming valid assertions.
		return nil
	}
	v.leavesLock.RUnlock()
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
		"challengingStateRoot": fmt.Sprintf("%#x", leaf.StateCommitment.StateRoot),
		"challengingHeight":    leaf.StateCommitment.Height,
	}).Warn("Received challenge for a created leaf")
	return nil
}

// Initiates a challenge on a created leaf event.
func (v *Validator) challengeLeaf(ctx context.Context, ev *protocol.CreateLeafEvent) error {
	// Retrieves the parent assertion to begin the challenge on.
	var parentAssertion *protocol.Assertion
	var err error
	if err = v.chain.Call(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		parentAssertion, err = p.AssertionBySequenceNum(tx, ev.PrevSeqNum)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	logFields := logrus.Fields{}
	logFields["name"] = v.name
	logFields["height"] = ev.StateCommitment.Height
	logFields["stateRoot"] = fmt.Sprintf("%#x", ev.StateCommitment.StateRoot)
	log.WithFields(logFields).Info("Initiating challenge on leaf validator disagrees with")

	var challenge *protocol.Challenge
	var challengeVertex *protocol.ChallengeVertex
	if err = v.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		challenge, err = parentAssertion.CreateChallenge(tx, ctx)
		if err != nil {
			return err
		}
		challengeVertex, err = challenge.AddLeaf(tx, parentAssertion, util.HistoryCommitment{})
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return errors.Wrapf(err, "could not create challenge on leaf with sequence number: %d", ev.PrevSeqNum)
	}
	_ = challenge
	return nil
}

// Prepares to defend a leaf that matches our local history and is part of a fork
// in the assertions tree. This leaf may be challenged and the local validator should
// be ready to perform proper challenge moves on the assertion if no one else is making them.
func (v *Validator) defendLeaf(ctx context.Context, ev *protocol.CreateLeafEvent) error {
	logFields := logrus.Fields{}
	if name, ok := v.knownValidatorNames[ev.Staker]; ok {
		logFields["createdBy"] = name
	}
	logFields["name"] = v.name
	logFields["height"] = ev.StateCommitment.Height
	logFields["stateRoot"] = fmt.Sprintf("%#x", ev.StateCommitment.StateRoot)
	log.WithFields(logFields).Info(
		"New leaf created by another validator matching local state has " +
			"forked the protocol, preparing to defend",
	)
	return nil
}
