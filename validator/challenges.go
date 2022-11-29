package validator

import (
	"context"
	"fmt"
	"sync"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	statemanager "github.com/OffchainLabs/new-rollup-exploration/state-manager"
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
	stateManager           statemanager.Manager
	validatorAddress       common.Address
	challenges             map[challengeID]*challengeWorker
	challengeEventsBufSize int
}

type challengeWorker struct {
	lock              sync.RWMutex
	challenge         *protocol.Challenge
	validatorAddress  common.Address
	leavesByParentSeq map[uint64][]*protocol.ChallengeVertex
	events            chan protocol.ChallengeEvent
}

func newChallengeManager(chain protocol.ChainReadWriter, validatorAddress common.Address) *challengeManager {
	return &challengeManager{
		chain:                  chain,
		challenges:             make(map[challengeID]*challengeWorker),
		validatorAddress:       validatorAddress,
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

func (c *challengeManager) spawnChallenge(
	ctx context.Context,
	challenge *protocol.Challenge,
	vertex *protocol.ChallengeVertex,
) {
	c.lock.Lock()
	ch := make(chan protocol.ChallengeEvent, c.challengeEventsBufSize)
	id := challenge.ParentStateCommitment().Hash()
	if _, ok := c.challenges[challengeID(id)]; ok {
		c.lock.Unlock()
		log.WithField(
			"challengeID", fmt.Sprintf("%#x", id),
		).Error("Attempted to spawn challenge that is already in progress")
		return
	}
	parentSeqNum := vertex.Prev.SequenceNum
	worker := &challengeWorker{
		challenge: challenge,
		leavesByParentSeq: map[uint64][]*protocol.ChallengeVertex{
			parentSeqNum: {vertex},
		},
		validatorAddress: c.validatorAddress,
		events:           ch,
	}
	c.challenges[challengeID(id)] = worker
	c.lock.Unlock()
	log.WithField("challengeID", fmt.Sprintf("%#x", id)).Info("Spawning challenge lifecycle manager")
	go worker.runChallengeLifecycle(ctx, c, ch)
}

func (w *challengeWorker) runChallengeLifecycle(
	ctx context.Context,
	manager *challengeManager,
	evs chan protocol.ChallengeEvent,
) {
	// TODO:
	// Manage chess clock moves for the validator.
	// Listen for challenge completion, win
	// Cleanup the challenge goroutine once done.
	// TODO: Figure out if we are at a one-step fork, and then depending on who's turn it is,
	// spawn a subchallenge (BigStepChallenge).
	defer close(evs)
	for {
		select {
		case genericEvent := <-evs:
			switch ev := genericEvent.(type) {
			case *protocol.ChallengeLeafEvent:
				go func() {
					if err := w.onChallengeLeafAdded(ctx, manager, ev); err != nil {
						log.WithError(err).Error("Could not process challenge leaf added event")
					}
				}()
			case *protocol.ChallengeBisectEvent:
				go func() {
					if err := w.onBisectionEvent(ctx, ev); err != nil {
						log.WithError(err).Error("Could not process bisection event")
					}
				}()
			case *protocol.ChallengeMergeEvent:
				go func() {
					if err := w.onMergeEvent(ctx, ev); err != nil {
						log.WithError(err).Error("Could not process merge event")
					}
				}()
			default:
				log.WithField("ev", fmt.Sprintf("%+v", ev)).Error("Not a recognized challenge event")
			}
		case <-ctx.Done():
			return
		}
	}
}

// TODO: The methods below need to be able to produce historical commitments for leaves.
//
// If a leaf has been added, we then check if we should add a competing leaf, bisect, or merge
// and then perform the corresponding action.
func (w *challengeWorker) onChallengeLeafAdded(
	ctx context.Context, manager *challengeManager, ev *protocol.ChallengeLeafEvent,
) error {
	// Ignore challenges initiated by self.
	if isFromSelf(w.validatorAddress, ev.Challenger) {
		return nil
	}
	w.lock.Lock()
	rivals := w.leavesByParentSeq[ev.ParentSeqNum]
	if len(rivals) == 0 {
		return nil
	}
	// TODO: Implement this conditional.
	// If we have the history commitment the new leaf claims to have, we do not need to act.
	// TODO: Check if this is the correct assumption.
	if manager.stateManager.HasHistoryCommitment(ctx, ev.History) {
		return nil
	}
	// Otherwise, we must bisect to our own historical commitment and produce
	// a proof of the vertex we want to bisect to.
	manager.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		// TODO: we need to bisect a rival leaf we have created.
		bisectedVertex, err := rivals[0].Bisect(tx, ev.History, nil, w.validatorAddress)
		if err != nil {
			return err
		}
		_ = bisectedVertex
		return nil
	})
	return nil
}

// If a bisection has occurred, we need to determine if we should merge or bisect.
func (w *challengeWorker) onBisectionEvent(ctx context.Context, ev *protocol.ChallengeBisectEvent) error {
	// If we agree with the history commitment, we make a merge move on the vertex.
	return nil
}

// If we have seen a merge event, what happens...??
func (w *challengeWorker) onMergeEvent(ctx context.Context, ev *protocol.ChallengeMergeEvent) error {
	return nil
}

// Process new challenge creation events from the protocol that were not initiated by self.
func (v *Validator) onChallengeStarted(ctx context.Context, ev *protocol.StartChallengeEvent) error {
	if ev == nil {
		return nil
	}
	// Ignore challenges initiated by self.
	if isFromSelf(v.address, ev.Challenger) {
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

	// TODO: Do we produce a historial commitment at the height == our latest height?
	historyCommit, err := v.stateManager.LatestHistoryCommitment(ctx)
	if err != nil {
		return err
	}

	// We then add a leaf to the challenge using a historical commitment at our latest height.
	var challenge *protocol.Challenge
	var challengeVertex *protocol.ChallengeVertex
	if err = v.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		parentAssertion, err := p.AssertionBySequenceNum(tx, ev.ParentSeqNum)
		if err != nil {
			return err
		}
		challenge, err = p.ChallengeByParentCommitmentHash(tx, parentAssertion.StateCommitment.Hash())
		if err != nil {
			return err
		}
		// TODO: What if the challenge already has a leaf we agree with?
		// TODO: Match on error ErrDuplicateLeaf to add nicer logs to the user.
		challengeVertex, err = challenge.AddLeaf(tx, parentAssertion, historyCommit, v.address)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return errors.Wrapf(err, "could not create challenge on leaf with sequence number: %d", ev.ParentSeqNum)
	}

	// Start tracking the challenge and created vertex using the challenge manager.
	v.challengeManager.spawnChallenge(ctx, challenge, challengeVertex)

	return nil
}

// Initiates a challenge on a created leaf event.
func (v *Validator) challengeLeaf(ctx context.Context, ev *protocol.CreateLeafEvent) error {
	// Retrieves the parent assertion to begin the challenge on.
	var parentAssertion *protocol.Assertion
	var currentAssertion *protocol.Assertion
	var err error
	if err = v.chain.Call(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		parentAssertion, err = p.AssertionBySequenceNum(tx, ev.PrevSeqNum)
		if err != nil {
			return err
		}
		currentAssertion, err = p.AssertionBySequenceNum(tx, ev.SeqNum)
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

	// We produce a historical commiment to add a leaf to the initiated challenge
	// by retrieving it from our local state manager.
	// TODO: Do we produce a historial commitment at the height == our latest height?
	historyCommit, err := v.stateManager.LatestHistoryCommitment(ctx)
	if err != nil {
		return err
	}

	var challenge *protocol.Challenge
	var challengeVertex *protocol.ChallengeVertex
	if err = v.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		challenge, err = parentAssertion.CreateChallenge(tx, ctx, v.address)
		if err != nil {
			return errors.Wrap(err, "cannot make challenge")
		}
		challengeVertex, err = challenge.AddLeaf(tx, currentAssertion, historyCommit, v.address)
		if err != nil {
			return errors.Wrap(err, "cannot add leaf")
		}
		return nil
	}); err != nil {
		return errors.Wrapf(err, "could not create challenge on leaf with sequence number: %d", ev.PrevSeqNum)
	}

	// Start tracking the challenge and created vertex using the challenge manager.
	v.challengeManager.spawnChallenge(ctx, challenge, challengeVertex)

	log.WithFields(logFields).Info("Successfully created challenge and added leaf, now tracking events")

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
