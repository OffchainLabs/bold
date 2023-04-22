package solimpl

import (
	"context"
	"time"

	"fmt"
	"math/big"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

type challengeWatcher struct {
	// Will keep track of ancestor histories for honest
	// branches per challenge.
	// Will scan for all previous events, and poll for new ones.
	// Will scan for level zero edges being confirmed and track
	// their claim id in this struct.
	chain              *AssertionChain
	pollEventsInterval time.Duration
	challenges         map[protocol.AssertionId]*challenge
}

// Checks if a confirmed, level zero edge exists that claims a particular
// claim id for a given challenge namespace (for a top-level assertion).
func (w *challengeWatcher) ConfirmedEdgeWithClaimExists(
	topLevelParentAssertionId protocol.AssertionId,
	claimId protocol.ClaimId,
) (bool, error) {
	challenge, ok := w.challenges[topLevelParentAssertionId]
	if !ok {
		return false, errors.New("assertion does not have an associated challenge")
	}
	return challenge.confirmedLevelZeroEdgeClaimIds.has(claimId), nil
}

type challenge struct {
	honestAncestorsBranch          *ancestorsBranch
	confirmedLevelZeroEdgeClaimIds set[protocol.ClaimId]
}

type ancestorsBranch struct {
	ordered              []protocol.EdgeId
	allAncestors         []protocol.EdgeId // Perhaps linked list instead?
	rivaled              set[protocol.EdgeId]
	totalBlocksUnrivaled uint64
	// Maybe need to keep time unrivaled up to a certain point...? Perhaps in a slice.
}

func (a *ancestorsBranch) updateTotalBlocksUnrivaled(
	ctx context.Context,
	challengeManager *SpecChallengeManager,
) {
	var total uint64
	for _, id := range a.allAncestors {
		if a.rivaled.has(id) {
			continue
		}
		blocksUnrivaled, err := challengeManager.caller.TimeUnrivaled(
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

// TODO: Panic if something occurs
func (w *challengeWatcher) watch(ctx context.Context) error {
	latestConfirmed, err := w.chain.LatestConfirmed(ctx)
	if err != nil {
		return err
	}
	assertionNode, ok := latestConfirmed.(*Assertion)
	if !ok {
		return errors.New("not ok")
	}
	inner, err := assertionNode.inner()
	if err != nil {
		return err
	}
	firstBlock := inner.CreatedAtBlock
	fromBlock := firstBlock

	// Some kind of backoff so as to not spam the node with requests.
	ticker := time.NewTicker(w.pollEventsInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			latestBlock, err := w.chain.backend.HeaderByNumber(ctx, nil)
			if err != nil {
				return err
			}
			toBlock := latestBlock.Number.Uint64()
			topic := common.BytesToHash(hexutil.MustDecode(
				"0xc9cc7aa3617dc3853c50ebf6703ec797191654dcc781255bed2057dce23b0e33",
			))
			var query = ethereum.FilterQuery{
				FromBlock: new(big.Int).SetUint64(fromBlock),
				ToBlock:   new(big.Int).SetUint64(toBlock),
				Addresses: []common.Address{w.chain.rollupAddr},
				Topics:    [][]common.Hash{{topic}, {}},
			}
			logs, err := w.chain.backend.FilterLogs(ctx, query)
			if err != nil {
				return err
			}
			if len(logs) == 0 {
				continue
			}
			for _, l := range logs {
				parsedLog, err := w.chain.rollup.ParseAssertionCreated(l)
				if err != nil {
					return err
				}
				fmt.Println(parsedLog)
			}
			fromBlock = toBlock
		case <-ctx.Done():
			return nil
		}
	}
}

func newChallengeWatcher() *challengeWatcher {
	return &challengeWatcher{}
}
