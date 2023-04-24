package validator

import (
	"context"
	"sync"
	"time"

	"fmt"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
)

type challengeWatcher struct {
	// Will keep track of ancestor histories for honest
	// branches per challenge.
	// Will scan for all previous events, and poll for new ones.
	// Will scan for level zero edges being confirmed and track
	// their claim id in this struct.
	chain              protocol.AssertionChain
	pollEventsInterval time.Duration
	lock               sync.RWMutex
	challenges         map[protocol.AssertionId]*challenge
	backend            bind.ContractBackend
}

func NewWatcher(
	chain protocol.AssertionChain,
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

type set[T comparable] struct {
	items map[T]bool
}

func newSet[T comparable]() *set[T] {
	return &set[T]{
		items: make(map[T]bool),
	}
}

func (s *set[T]) insert(t T) {
	s.items[t] = true
}

func (s *set[T]) has(t T) bool {
	return s.items[t]
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
		fmt.Println("ONE STEP PROOF CONFIRMATION")
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
	claimId := edge.ClaimId().Unwrap()
	chal := w.challenges[protocol.AssertionId{}]
	chal.confirmedLevelZeroEdgeClaimIds.insert(claimId)
	return nil
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
		fmt.Println("EDGE ADDED WHOA")
		edgeAdded := it.Event
		if protocol.EdgeType(edgeAdded.EType) == protocol.BlockChallengeEdge {
			w.lock.Lock()
			if _, ok := w.challenges[edgeAdded.ClaimId]; !ok {
				w.challenges[edgeAdded.ClaimId] = &challenge{
					honestAncestorsBranch:          &ancestorsBranch{},
					confirmedLevelZeroEdgeClaimIds: newSet[protocol.ClaimId](),
				}
			}
			w.lock.Unlock()
		}
	}
	return nil
}
