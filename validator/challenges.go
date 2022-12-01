package validator

import (
	"context"
	"fmt"
	"sync"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	statemanager "github.com/OffchainLabs/new-rollup-exploration/state-manager"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

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
	validatorName          string
	challenges             map[protocol.ChallengeID]*challengeWorker
	challengeEventsBufSize int
}

type challengeWorker struct {
	challenge          *protocol.Challenge
	validatorAddress   common.Address
	reachedOneStepFork chan struct{}
	validatorName      string
	createdVertices    *util.ThreadSafeSlice[*protocol.ChallengeVertex]
	events             chan protocol.ChallengeEvent
}

func newChallengeManager(
	chain protocol.ChainReadWriter,
	stateManager statemanager.Manager,
	validatorAddress common.Address,
	validatorName string,
) *challengeManager {
	return &challengeManager{
		chain:                  chain,
		stateManager:           stateManager,
		challenges:             make(map[protocol.ChallengeID]*challengeWorker),
		validatorAddress:       validatorAddress,
		validatorName:          validatorName,
		challengeEventsBufSize: 100, // TODO: Make configurable.
	}
}

func (c *challengeManager) spawnChallenge(
	ctx context.Context,
	challenge *protocol.Challenge,
	vertex *protocol.ChallengeVertex,
) {
	c.lock.Lock()
	ch := make(chan protocol.ChallengeEvent, c.challengeEventsBufSize)
	c.chain.SubscribeChallengeEvents(ctx, ch)
	id := challenge.ParentStateCommitment().Hash()
	if _, ok := c.challenges[protocol.ChallengeID(id)]; ok {
		c.lock.Unlock()
		log.WithFields(logrus.Fields{
			"challengeID": fmt.Sprintf("%#x", id),
			"name":        c.validatorName,
		}).Error("Attempted to spawn challenge that is already in progress")
		return
	}
	vertices := util.NewThreadSafeSlice[*protocol.ChallengeVertex]()
	vertices.Append(vertex)
	worker := &challengeWorker{
		challenge:          challenge,
		createdVertices:    vertices,
		validatorAddress:   c.validatorAddress,
		reachedOneStepFork: make(chan struct{}),
		validatorName:      c.validatorName,
		events:             ch,
	}
	c.challenges[protocol.ChallengeID(id)] = worker
	c.lock.Unlock()
	log.WithFields(logrus.Fields{
		"challengeID": fmt.Sprintf("%#x", id),
		"name":        c.validatorName,
	}).Info("Spawning challenge lifecycle manager")
	go worker.runChallengeLifecycle(ctx, c, ch)
}

func (w *challengeWorker) runChallengeLifecycle(
	ctx context.Context,
	manager *challengeManager,
	evs chan protocol.ChallengeEvent,
) {
	for {
		select {
		case genericEvent := <-evs:
			switch ev := genericEvent.(type) {
			case *protocol.ChallengeLeafEvent:
				go func() {
					if err := w.act(ctx, manager, ev.Actor, ev.History, ev.SequenceNum); err != nil {
						log.WithError(err).Error("Could not process challenge leaf added event")
					}
				}()
			case *protocol.ChallengeBisectEvent:
				go func() {
					if err := w.act(ctx, manager, ev.Actor, ev.History, ev.SequenceNum); err != nil {
						log.WithError(err).Error("Could not process bisection event")
					}
				}()
			case *protocol.ChallengeMergeEvent:
				go func() {
					if err := w.act(ctx, manager, ev.Actor, ev.History, ev.ShallowerSequenceNum); err != nil {
						log.WithError(err).Error("Could not process merge event")
					}
				}()
			default:
				log.WithField("ev", fmt.Sprintf("%+v", ev)).Error("Not a recognized challenge event")
			}
		case <-w.reachedOneStepFork:
			log.WithField(
				"name", w.validatorName,
			).Infof("Reached a one-step-fork in the challenge, now awaiting subchallenge resolution")
			// TODO: Trigger subchallenge!
			return
		case <-ctx.Done():
			return
		}
	}
}

func (w *challengeWorker) act(
	ctx context.Context,
	manager *challengeManager,
	eventActor common.Address,
	eventHistoryCommit util.HistoryCommitment,
	eventSequenceNum uint64,
) error {
	if isFromSelf(w.validatorAddress, eventActor) {
		return nil
	}
	if w.createdVertices.Empty() {
		return nil
	}
	// Go down the tree to find the first vertex we created that has a commitment height >
	// the vertex seen from the merge event.
	vertexToActUpon := w.createdVertices.Last().Unwrap()
	numVertices := w.createdVertices.Len()
	for i := numVertices - 1; i > 0; i-- {
		vertex := w.createdVertices.Get(i).Unwrap()
		if vertex.Commitment.Height > eventHistoryCommit.Height {
			vertexToActUpon = vertex
			break
		}
	}
	log.WithFields(logrus.Fields{
		"name":   w.validatorName,
		"height": eventHistoryCommit.Height,
	}).Infof("Received a merge event")

	mergedToOurs := eventHistoryCommit.Height == vertexToActUpon.Commitment.Height
	if mergedToOurs {
		log.WithFields(logrus.Fields{
			"name":         w.validatorName,
			"mergedHeight": eventHistoryCommit.Height,
		}).Info("Other validator merged to our vertex")
	}

	// Make a merge move.
	if manager.stateManager.HasHistoryCommitment(ctx, eventHistoryCommit) {
		log.WithField(
			"name", w.validatorName,
		).Info("Agreed with bisection event from other challenger, starting a merge move")
		if err := w.merge(ctx, manager, vertexToBisect, eventSequenceNum); err != nil {
			return errors.Wrap(err, "failed to merge to a bisected vertex")
		}
	}

	// Bisect.
	hasPresumptiveSuccessor := vertexToActUpon.IsPresumptiveSuccessor()
	currentVertex := vertexToActUpon

	for !hasPresumptiveSuccessor {
		if currentVertex.Commitment.Height == currentVertex.Prev.Commitment.Height+1 {
			w.reachedOneStepFork <- struct{}{}
			break
		}
		bisectedVertex, err := w.bisect(ctx, currentVertex.Prev.Commitment.Height, currentVertex, manager)
		if err != nil {
			return err
		}
		currentVertex = bisectedVertex
		w.createdVertices.Append(currentVertex)
	}
	return nil
}

func (w *challengeWorker) bisect(
	ctx context.Context,
	parentHeight uint64,
	validatorChallengeVertex *protocol.ChallengeVertex,
	manager *challengeManager,
) (*protocol.ChallengeVertex, error) {
	toHeight := validatorChallengeVertex.Commitment.Height
	bisectTo, err := util.BisectionPoint(parentHeight, toHeight)
	if err != nil {
		return nil, err
	}
	historyCommit, err := manager.stateManager.HistoryCommitmentUpTo(ctx, bisectTo)
	if err != nil {
		return nil, err
	}
	proof, err := manager.stateManager.PrefixProof(ctx, bisectTo, toHeight)
	if err != nil {
		return nil, err
	}
	if err := util.VerifyPrefixProof(historyCommit, validatorChallengeVertex.Commitment, proof); err != nil {
		return nil, err
	}
	// Otherwise, we must bisect to our own historical commitment and produce
	// a proof of the vertex we want to bisect to.
	var bisectedVertex *protocol.ChallengeVertex
	err = manager.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		bisectedVertex, err = validatorChallengeVertex.Bisect(tx, historyCommit, proof, w.validatorAddress)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"could not bisect vertex with sequence %d and challenger %#x to height %d with history %d and %#x",
			validatorChallengeVertex.SequenceNum,
			validatorChallengeVertex.Challenger,
			bisectTo,
			historyCommit.Height,
			historyCommit.Merkle,
		)
	}
	log.WithFields(logrus.Fields{
		"name":                   w.validatorName,
		"IsPresumptiveSuccessor": bisectedVertex.IsPresumptiveSuccessor(),
	}).Infof(
		"Successfully bisected to vertex with height %d and commit %#x",
		bisectedVertex.Commitment.Height,
		bisectedVertex.Commitment.Merkle,
	)
	return bisectedVertex, nil
}

func (w *challengeWorker) merge(
	ctx context.Context,
	manager *challengeManager,
	validatorChallengeVertex *protocol.ChallengeVertex,
	newPrevSeqNum uint64,
) error {
	var bisectedVertex *protocol.ChallengeVertex
	var err error
	err = manager.chain.Call(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		id := protocol.ChallengeID(w.challenge.ParentStateCommitment().Hash())
		bisectedVertex, err = p.ChallengeVertexBySequenceNum(tx, id, newPrevSeqNum)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "could not read challenge vertex from protocol")
	}
	bisectionHeight := bisectedVertex.Commitment.Height
	historyCommit, err := manager.stateManager.HistoryCommitmentUpTo(ctx, bisectionHeight)
	if err != nil {
		return err
	}
	currentCommit := validatorChallengeVertex.Commitment
	proof, err := manager.stateManager.PrefixProof(ctx, bisectionHeight, currentCommit.Height)
	if err != nil {
		return err
	}
	if err := util.VerifyPrefixProof(historyCommit, currentCommit, proof); err != nil {
		return err
	}
	if err := manager.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		return validatorChallengeVertex.Merge(tx, bisectedVertex, proof, w.validatorAddress)
	}); err != nil {
		return errors.Wrapf(
			err,
			"could not merge vertex with height %d and commit %#x to height %x and commit %#x",
			currentCommit.Height,
			currentCommit.Merkle,
			bisectionHeight,
			bisectedVertex.Commitment.Merkle,
		)
	}
	log.WithFields(logrus.Fields{
		"name": w.validatorName,
	}).Infof(
		"Successfully merged to vertex with height %d and commit %#x",
		bisectedVertex.Commitment.Height,
		bisectedVertex.Commitment.Merkle,
	)
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
	if !leaf.Staker.IsNone() {
		if name, ok := v.knownValidatorNames[leaf.Staker.Unwrap()]; ok {
			challengerName = name
		} else {
			challengerName = leaf.Staker.Unwrap().Hex()
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
		challenge, err = p.ChallengeByID(tx, protocol.ChallengeID(parentAssertion.StateCommitment.Hash()))
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
	logFields["disagreedHeight"] = ev.StateCommitment.Height
	logFields["disagreedLeaf"] = fmt.Sprintf("%#x", ev.StateCommitment.StateRoot)
	log.WithFields(logFields).Info("Initiating challenge on parent of leaf validator disagrees with")

	var challenge *protocol.Challenge
	var challengeVertex *protocol.ChallengeVertex
	err = v.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		challenge, err = parentAssertion.CreateChallenge(tx, ctx, v.address)
		if err != nil {
			return errors.Wrap(err, "cannot make challenge")
		}
		return nil
	})
	switch {
	case errors.Is(err, protocol.ErrChallengeAlreadyExists):
		log.Info("Challenge on leaf already exists, reading existing challenge from protocol")
		if err = v.chain.Call(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
			challenge, err = p.ChallengeByID(tx, protocol.ChallengeID(parentAssertion.StateCommitment.Hash()))
			if err != nil {
				return errors.Wrap(err, "cannot make challenge")
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "could not get challenge by ID")
		}
	case err != nil:
		return errors.Wrapf(
			err,
			"could not initiate challenge on assertion with seq num %d",
			parentAssertion.SequenceNum,
		)
	default:
	}

	if challenge == nil {
		return errors.New("got nil challenge from protocol")
	}

	// We produce a historical commiment to add a leaf to the initiated challenge
	// by retrieving it from our local state manager.
	// TODO: Do we produce a historial commitment at the height == our latest height?
	historyCommit, err := v.stateManager.LatestHistoryCommitment(ctx)
	if err != nil {
		return err
	}

	if err = v.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		challengeVertex, err = challenge.AddLeaf(tx, currentAssertion, historyCommit, v.address)
		if err != nil {
			return errors.Wrap(err, "cannot add leaf")
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "could not add leaf to challenge")
	}

	// Start tracking the challenge and created vertex using the challenge manager.
	v.challengeManager.spawnChallenge(ctx, challenge, challengeVertex)

	logFields = logrus.Fields{}
	logFields["name"] = v.name
	logFields["parentAssertionSeqNum"] = parentAssertion.SequenceNum
	logFields["parentAssertionStateRoot"] = fmt.Sprintf("%#x", parentAssertion.StateCommitment.StateRoot)
	logFields["challengeID"] = fmt.Sprintf("%#x", parentAssertion.StateCommitment.Hash())
	log.WithFields(logFields).Info("Successfully created challenge and added leaf, now tracking events")

	return nil
}
