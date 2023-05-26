package validator

import (
	"context"
	"time"

	"fmt"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	statemanager "github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	challengetree "github.com/OffchainLabs/challenge-protocol-v2/validator/challenge-tree"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
)

// Represents a set of honest edges being tracked in a top-level challenge and all the
// associated subchallenge honest edges along with some more metadata used for
// computing information needed for confirmations. Each time an edge is created onchain,
// the challenge watcher service will add it to its respective "trackedChallenge"
// namespaced under the top-level assertion id the edge belongs to.
type trackedChallenge struct {
	honestEdgeTree                 *challengetree.HonestChallengeTree
	confirmedLevelZeroEdgeClaimIds *threadsafe.Set[protocol.ClaimId]
}

// The challenge watcher implements a service in the validator runtime
// that is in charge of scanning through all edge creation events via a polling
// mechanism. It will keep track of edges the validator's state provider agrees with
// within trackedChallenge instances. The challenge watcher provides two useful
// methods: (a) the ability to compute the honest path timer of an edge, and
// (b) the ability to check if an edge with a certain claim id has been confirmed. Both
// are used during the confirmation process in edge tracker goroutines.
type challengeWatcher struct {
	stateManager       statemanager.Manager
	chain              protocol.AssertionChain
	pollEventsInterval time.Duration
	challenges         *threadsafe.Map[protocol.AssertionId, *trackedChallenge]
	backend            bind.ContractBackend
	validatorName      string
}

func newChallengeWatcher(
	chain protocol.AssertionChain,
	manager statemanager.Manager,
	backend bind.ContractBackend,
	interval time.Duration,
	validatorName string,
) *challengeWatcher {
	return &challengeWatcher{
		chain:              chain,
		pollEventsInterval: interval,
		challenges:         threadsafe.NewMap[protocol.AssertionId, *trackedChallenge](),
		backend:            backend,
		stateManager:       manager,
		validatorName:      validatorName,
	}
}

// Checks if a confirmed, level zero edge exists that claims a particular
// edge id for a tracked challenge. This is used during the confirmation process of edges
// within edge tracker goroutines.
func (w *challengeWatcher) confirmedEdgeWithClaimExists(
	topLevelParentAssertionId protocol.AssertionId,
	claimId protocol.ClaimId,
) (bool, error) {
	challenge, ok := w.challenges.TryGet(topLevelParentAssertionId)
	if !ok {
		return false, errors.New("assertion does not have an associated challenge")
	}
	return challenge.confirmedLevelZeroEdgeClaimIds.Has(claimId), nil
}

// Computes the honest path timer for an edge id within an assertion id challenge
// namespace. This is used during the confirmation process for edges in
// edge tracker goroutine logic.
func (w *challengeWatcher) computeHonestPathTimer(
	ctx context.Context,
	topLevelParentAssertionId protocol.AssertionId,
	edgeId protocol.EdgeId,
) (challengetree.PathTimer, challengetree.HonestAncestors, error) {
	header, err := w.backend.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, nil, err
	}
	if !header.Number.IsUint64() {
		return 0, nil, errors.New("latest block header number is not a uint64")
	}
	blockNumber := header.Number.Uint64()
	chal, ok := w.challenges.TryGet(topLevelParentAssertionId)
	if !ok {
		return 0, nil, fmt.Errorf(
			"could not get challenge for top level assertion %#x",
			topLevelParentAssertionId,
		)
	}
	return chal.honestEdgeTree.HonestPathTimer(ctx, edgeId, blockNumber)
}

func (w *challengeWatcher) watch(ctx context.Context) {
	scanRange, err := retryUntilSucceeds(ctx, func() (filterRange, error) {
		return w.getStartEndBlockNum(ctx)
	})
	if err != nil {
		log.Error(err)
		return
	}
	fromBlock := scanRange.startBlockNum
	toBlock := scanRange.endBlockNum

	_, err = retryUntilSucceeds(ctx, func() (bool, error) {
		return true, w.checkForEdgeAdded(ctx, fromBlock, toBlock)
	})
	if err != nil {
		log.Error(err)
		return
	}
	fromBlock = toBlock

	// Some kind of backoff so as to not spam the node with requests.
	ticker := time.NewTicker(w.pollEventsInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			latestBlock, err := w.backend.HeaderByNumber(ctx, nil)
			if err != nil {
				log.Error(err)
				continue
			}
			toBlock := latestBlock.Number.Uint64()
			if fromBlock == toBlock {
				continue
			}
			if err = w.checkForEdgeAdded(ctx, fromBlock, toBlock); err != nil {
				log.Error(err)
				continue
			}
			fromBlock = toBlock
		case <-ctx.Done():
			return
		}
	}
}

func (w *challengeWatcher) checkForEdgeAdded(
	ctx context.Context,
	startBlock,
	endBlock uint64,
) error {
	challengeManager, err := w.chain.SpecChallengeManager(ctx)
	if err != nil {
		return err
	}
	filterer, err := challengeV2gen.NewEdgeChallengeManagerFilterer(challengeManager.Address(), w.backend)
	if err != nil {
		return err
	}
	filterOpts := &bind.FilterOpts{
		Start:   startBlock,
		End:     &endBlock,
		Context: ctx,
	}
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
				startBlock,
				endBlock,
			)
		}
		_, processErr := retryUntilSucceeds(ctx, func() (bool, error) {
			return true, w.processEdgeAddedEvent(ctx, challengeManager, it.Event)
		})
		if processErr != nil {
			return processErr
		}
	}
	return nil
}

func (w *challengeWatcher) processEdgeAddedEvent(
	ctx context.Context,
	challengeManager protocol.SpecChallengeManager,
	event *challengeV2gen.EdgeChallengeManagerEdgeAdded,
) error {
	edgeOpt, err := challengeManager.GetEdge(ctx, protocol.EdgeId(event.EdgeId))
	if err != nil {
		return err
	}
	if edgeOpt.IsNone() {
		return fmt.Errorf("no edge found with id %#x", event.EdgeId)
	}
	edge := edgeOpt.Unwrap()

	assertionId, err := edge.PrevAssertionId(ctx)
	if err != nil {
		return err
	}
	chal, ok := w.challenges.TryGet(assertionId)
	if !ok {
		tree := challengetree.New(
			protocol.AssertionId(event.OriginId),
			w.chain,
			w.stateManager,
			w.validatorName,
		)
		chal = &trackedChallenge{
			honestEdgeTree:                 tree,
			confirmedLevelZeroEdgeClaimIds: threadsafe.NewSet[protocol.ClaimId](),
		}
		w.challenges.Put(assertionId, chal)
	}
	return chal.honestEdgeTree.AddEdge(ctx, edge)
}

func (w *challengeWatcher) checkForEdgeConfirmedByOneStepProof(
	filterOpts *bind.FilterOpts,
	manager protocol.SpecChallengeManager,
	filterer *challengeV2gen.EdgeChallengeManagerFilterer,
) error {
	it, err := filterer.FilterEdgeConfirmedByOneStepProof(filterOpts, nil, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err = it.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	for it.Next() {
		if it.Error() != nil {
			return err // TODO: Handle better.
		}
		if err := w.checkLevelZeroEdgeConfirmed(filterOpts.Context, manager, it.Event.EdgeId); err != nil {
			return err
		}
	}
	return nil
}

func (w *challengeWatcher) checkLevelZeroEdgeConfirmed(
	ctx context.Context,
	manager protocol.SpecChallengeManager,
	edgeId protocol.EdgeId,
) error {
	edgeOpt, err := manager.GetEdge(ctx, edgeId)
	if err != nil {
		return err
	}
	if edgeOpt.IsNone() {
		return errors.New("no edge found")
	}
	edge := edgeOpt.Unwrap()
	if edge.ClaimId().IsNone() {
		return nil
	}
	// claimId := edge.ClaimId().Unwrap()
	// chal := w.challenges[protocol.AssertionId{}]
	// chal.confirmedLevelZeroEdgeClaimIds.Insert(claimId)
	return nil
}

type filterRange struct {
	startBlockNum uint64
	endBlockNum   uint64
}

func (w *challengeWatcher) getStartEndBlockNum(ctx context.Context) (filterRange, error) {
	latestConfirmed, err := w.chain.LatestConfirmed(ctx)
	if err != nil {
		return filterRange{}, err
	}
	firstBlock, err := latestConfirmed.CreatedAtBlock()
	if err != nil {
		return filterRange{}, err
	}
	startBlock := firstBlock
	header, err := w.backend.HeaderByNumber(ctx, nil)
	if err != nil {
		return filterRange{}, err
	}
	return filterRange{
		startBlockNum: startBlock,
		endBlockNum:   header.Number.Uint64(),
	}, nil
}
