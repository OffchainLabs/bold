package validator

import (
	"context"
	"sync"
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

// The challenge watcher implements a singleton service in the validator runtime
// that is in charge of scanning through all edge creation events via a polling
// mechanism
type challengeWatcher struct {
	stateManager       statemanager.Manager
	chain              protocol.AssertionChain
	pollEventsInterval time.Duration
	lock               sync.RWMutex
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
		challenges:         threadsafe.NewMap[protocol.AssertionId, *challenge](),
		backend:            backend,
		stateManager:       manager,
		validatorName:      validatorName,
	}
}

// Checks if a confirmed, level zero edge exists that claims a particular
// claim id for a given challenge namespace (for a top-level assertion).
func (w *challengeWatcher) confirmedEdgeWithClaimExists(
	topLevelParentAssertionId protocol.AssertionId,
	claimId protocol.ClaimId,
) (bool, error) {
	w.lock.RLock()
	defer w.lock.RUnlock()
	challenge, ok := w.challenges.TryGet(topLevelParentAssertionId)
	if !ok {
		return false, errors.New("assertion does not have an associated challenge")
	}
	return challenge.confirmedLevelZeroEdgeClaimIds.Has(claimId), nil
}

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
	// Start from the latest confirmed assertion's creation block.
	latestConfirmed, err := w.chain.LatestConfirmed(ctx)
	if err != nil {
		panic(err)
	}
	firstBlock := latestConfirmed.CreatedAtBlock()
	fromBlock := firstBlock

	latestBlock, err := w.backend.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Error(err)
		return
	}
	toBlock := latestBlock.Number.Uint64()
	//log.Infof("&&&&&&&&&SCANNING %d to %d", fromBlock, toBlock)

	challengeManager, err := w.chain.SpecChallengeManager(ctx)
	if err != nil {
		log.Error(err)
		return
	}
	filterer, err := challengeV2gen.NewEdgeChallengeManagerFilterer(challengeManager.Address(), w.backend)
	if err != nil {
		log.Error(err)
		return
	}
	filterOpts := &bind.FilterOpts{
		Start:   fromBlock,
		End:     &toBlock,
		Context: ctx,
	}
	if err = w.checkForEdgeAdded(filterOpts, filterer); err != nil {
		log.Error(err)
		return
	}
	// if err = w.checkForEdgeConfirmedByOneStepProof(filterOpts, challengeManager, filterer); err != nil {
	// 	return err
	// }
	// Watcher needs access to the challenge manager. If it sees an edge it agrees with (honest),
	// it will then persist that in the honest ancestors branch. It needs to keep track of ancestors
	// in a special order.
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
			//log.Infof("&&&&&&&&&SCANNING %d to %d", fromBlock, toBlock)

			if fromBlock == toBlock {
				continue
			}

			challengeManager, err := w.chain.SpecChallengeManager(ctx)
			if err != nil {
				log.Error(err)
				continue
			}
			filterer, err := challengeV2gen.NewEdgeChallengeManagerFilterer(challengeManager.Address(), w.backend)
			if err != nil {
				log.Error(err)
				continue
			}
			filterOpts := &bind.FilterOpts{
				Start:   fromBlock,
				End:     &toBlock,
				Context: ctx,
			}
			if err = w.checkForEdgeAdded(filterOpts, filterer); err != nil {
				log.Error(err)
				continue
			}
			// if err = w.checkForEdgeConfirmedByOneStepProof(filterOpts, challengeManager, filterer); err != nil {
			// 	return err
			// }
			// Watcher needs access to the challenge manager. If it sees an edge it agrees with (honest),
			// it will then persist that in the honest ancestors branch. It needs to keep track of ancestors
			// in a special order.
			fromBlock = toBlock
		case <-ctx.Done():
			return
		}
	}
}

func (w *challengeWatcher) checkForEdgeAdded(
	filterOpts *bind.FilterOpts,
	filterer *challengeV2gen.EdgeChallengeManagerFilterer,
) error {
	it, err := filterer.FilterEdgeAdded(filterOpts, nil, nil, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err = it.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	ctx := context.TODO()
	challengeManager, err := w.chain.SpecChallengeManager(ctx)
	if err != nil {
		return err
	}
	for it.Next() {
		if it.Error() != nil {
			log.WithError(err).Error("***************WEIRDOOOOOO 1st")
			return err // TODO: Handle better.
		}
		edgeAdded := it.Event
		//log.Infof("GOT EVENT CREATION %#x", edgeAdded.EdgeId)
		edgeOpt, err := challengeManager.GetEdge(ctx, protocol.EdgeId(edgeAdded.EdgeId))
		if err != nil {
			return err
		}
		if edgeOpt.IsNone() {
			return fmt.Errorf("no edge found with id %#x", edgeAdded.EdgeId)
		}
		edge := edgeOpt.Unwrap()

		assertionId, err := edge.PrevAssertionId(ctx)
		if err != nil {
			return err
		}
		chal, ok := w.challenges.TryGet(assertionId)
		if !ok {
			tree := challengetree.New(
				protocol.AssertionId(edgeAdded.OriginId),
				w.chain,
				w.stateManager,
				w.validatorName,
			)
			chal = &challenge{
				honestEdgeTree:                 tree,
				confirmedLevelZeroEdgeClaimIds: threadsafe.NewSet[protocol.ClaimId](),
			}
			w.challenges.Put(assertionId, chal)
		}
		if err := chal.honestEdgeTree.AddEdge(ctx, edge); err != nil {
			return err
		}
	}
	return nil
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
		log.Warn("Watcher: Edge was confirmed by one step proof")
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
	w.lock.Lock()
	defer w.lock.Unlock()
	log.Warn("Watcher: Level zero edge was confirmed")
	// claimId := edge.ClaimId().Unwrap()
	// chal := w.challenges[protocol.AssertionId{}]
	// chal.confirmedLevelZeroEdgeClaimIds.Insert(claimId)
	return nil
}
