package validator

import (
	"context"
	"sync"
	"time"

	"fmt"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	statemanager "github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
)

type challengeWatcher struct {
	// Will keep track of ancestor histories for honest
	// branches per challenge.
	// Will scan for all previous events, and poll for new ones.
	// Will scan for level zero edges being confirmed and track
	// their claim id in this struct.
	stateManager       statemanager.Manager
	chain              protocol.AssertionChain
	pollEventsInterval time.Duration
	lock               sync.RWMutex
	challenges         map[protocol.AssertionId]*challenge
	backend            bind.ContractBackend
}

func NewWatcher(
	chain protocol.AssertionChain,
	manager statemanager.Manager,
	backend bind.ContractBackend,
	interval time.Duration,
) *challengeWatcher {
	return &challengeWatcher{
		chain:              chain,
		pollEventsInterval: interval,
		challenges:         make(map[protocol.AssertionId]*challenge),
		backend:            backend,
	}
}

// Checks if a confirmed, level zero edge exists that claims a particular
// claim id for a given challenge namespace (for a top-level assertion).
func (w *challengeWatcher) ConfirmedEdgeWithClaimExists(
	topLevelParentAssertionId protocol.AssertionId,
	claimId protocol.ClaimId,
) (bool, error) {
	w.lock.RLock()
	defer w.lock.RUnlock()
	challenge, ok := w.challenges[topLevelParentAssertionId]
	if !ok {
		return false, errors.New("assertion does not have an associated challenge")
	}
	return challenge.confirmedLevelZeroEdgeClaimIds.has(claimId), nil
}

// TODO: Implement using ancestor branches.
func (w *challengeWatcher) BranchTotalUnrivaledBlocks(
	topLevelParentAssertionId protocol.AssertionId,
	edgeId protocol.EdgeId,
) (uint64, error) {
	return 0, nil
}

func (w *challengeWatcher) Ancestors(
	topLevelParentAssertionId protocol.AssertionId,
	edgeId protocol.EdgeId,
) ([]protocol.EdgeId, error) {
	return nil, nil
}

type challenge struct {
	honestAncestorsBranch          *ancestorsBranch
	confirmedLevelZeroEdgeClaimIds *set[protocol.ClaimId]
}

type ancestorsBranch struct {
	ordered              []protocol.EdgeId
	allAncestors         []protocol.EdgeId // Perhaps linked list instead?
	rivaled              *set[protocol.EdgeId]
	totalBlocksUnrivaled uint64
	// Maybe need to keep time unrivaled up to a certain point...? Perhaps in a slice.
	// Cumulative unrivaled time list, for example.
}

func (a *ancestorsBranch) updateTotalBlocksUnrivaled(
	ctx context.Context,
	challengeManager protocol.SpecChallengeManager,
	backend bind.ContractBackend,
) {
	caller, err := challengeV2gen.NewEdgeChallengeManagerCaller(challengeManager.Address(), backend)
	if err != nil {
		panic(err)
	}
	var total uint64
	for _, id := range a.allAncestors {
		if a.rivaled.has(id) {
			continue
		}
		blocksUnrivaled, err := caller.TimeUnrivaled(
			&bind.CallOpts{Context: ctx}, id,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		total += blocksUnrivaled.Uint64()
	}
}

func (w *challengeWatcher) Watch(ctx context.Context) error {
	// Start from the latest confirmed assertion's creation block.
	latestConfirmed, err := w.chain.LatestConfirmed(ctx)
	if err != nil {
		return err
	}
	firstBlock, err := latestConfirmed.CreatedAtBlock()
	if err != nil {
		return err
	}
	fromBlock := firstBlock

	// Some kind of backoff so as to not spam the node with requests.
	ticker := time.NewTicker(w.pollEventsInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			latestBlock, err := w.backend.HeaderByNumber(ctx, nil)
			if err != nil {
				return err
			}
			toBlock := latestBlock.Number.Uint64()

			if fromBlock == toBlock {
				continue
			}

			challengeManager, err := w.chain.SpecChallengeManager(ctx)
			if err != nil {
				return err
			}
			filterer, err := challengeV2gen.NewEdgeChallengeManagerFilterer(challengeManager.Address(), w.backend)
			if err != nil {
				return err
			}
			filterOpts := &bind.FilterOpts{
				Start:   fromBlock,
				End:     &toBlock,
				Context: ctx,
			}
			if err = w.checkForEdgeAdded(filterOpts, filterer); err != nil {
				return err
			}
			if err = w.checkForEdgeConfirmedByOneStepProof(filterOpts, challengeManager, filterer); err != nil {
				return err
			}
			// Watcher needs access to the challenge manager. If it sees an edge it agrees with (honest),
			// it will then persist that in the honest ancestors branch. It needs to keep track of ancestors
			// in a special order.
			fromBlock = toBlock
		case <-ctx.Done():
			return nil
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
	for it.Next() {
		if it.Error() != nil {
			return err // TODO: Handle better.
		}
		edgeAdded := it.Event

		// ClaimID: Entire namespace for a challenge's 3 levels.
		// OriginID: Namespace per subchallenge level.
		if protocol.EdgeType(edgeAdded.EType) == protocol.BlockChallengeEdge && edgeAdded.ClaimId != [32]byte{} {
			w.lock.Lock()
			if _, ok := w.challenges[edgeAdded.ClaimId]; !ok {
				w.challenges[edgeAdded.ClaimId] = &challenge{
					honestAncestorsBranch:          &ancestorsBranch{},
					confirmedLevelZeroEdgeClaimIds: newSet[protocol.ClaimId](),
				}
			}
			w.lock.Unlock()
		}

		// We see a new edge: what do?
		// - For each branch in a challenge, we need to figure out where to place it on the branch
		//   - We need to figure out its parent/child relationship to other edges being tracked here
		// - We need to figure out what top-level assertion it corresponds to.
		//   - However, we only have this information for level zero edges, using a claim id
		//   - for all others, we have an origin id that can be used for namespacing information

		// TODO: Optimization: figure out if the edge being added has a history commitment we agree with
		// locally, and only track the edges this watcher agrees with.
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
	claimId := edge.ClaimId().Unwrap()
	chal := w.challenges[protocol.AssertionId{}]
	chal.confirmedLevelZeroEdgeClaimIds.insert(claimId)
	return nil
}
