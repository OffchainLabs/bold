package validator

import (
	"context"
	"time"

	"fmt"
	"math/big"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

var (
	topicByName                      = make(map[string]common.Hash)
	nameByTopic                      = make(map[common.Hash]string)
	edgeAddedEvent                   = "EdgeAdded"
	edgeBisectedEvent                = "EdgeBisected"
	edgeConfirmedByChildrenEvent     = "EdgeConfiredByChildren"
	edgeConfirmedByTimeEvent         = "EdgeConfirmedByTime"
	edgeConfirmedByOneStepProofEvent = "EdgeConfirmedByOneStepProof"
	edgeConfirmedByClaimEvent        = "EdgeConfirmedByClaim"
)

func init() {
	abi, err := challengeV2gen.EdgeChallengeManagerMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	topicByName = map[string]common.Hash{
		edgeAddedEvent:                   crypto.Keccak256Hash([]byte(abi.Events[edgeAddedEvent].Sig)),
		edgeBisectedEvent:                crypto.Keccak256Hash([]byte(abi.Events[edgeBisectedEvent].Sig)),
		edgeConfirmedByChildrenEvent:     crypto.Keccak256Hash([]byte(abi.Events[edgeConfirmedByChildrenEvent].Sig)),
		edgeConfirmedByTimeEvent:         crypto.Keccak256Hash([]byte(abi.Events[edgeConfirmedByTimeEvent].Sig)),
		edgeConfirmedByOneStepProofEvent: crypto.Keccak256Hash([]byte(abi.Events[edgeConfirmedByOneStepProofEvent].Sig)),
		edgeConfirmedByClaimEvent:        crypto.Keccak256Hash([]byte(abi.Events[edgeConfirmedByClaimEvent].Sig)),
	}
	for k, v := range topicByName {
		nameByTopic[v] = k
	}
}

type challengeWatcher struct {
	// Will keep track of ancestor histories for honest
	// branches per challenge.
	// Will scan for all previous events, and poll for new ones.
	// Will scan for level zero edges being confirmed and track
	// their claim id in this struct.
	chain              protocol.AssertionChain
	pollEventsInterval time.Duration
	challenges         map[protocol.AssertionId]*challenge
}

func NewWatcher(
	chain protocol.AssertionChain,
	interval time.Duration,
) *challengeWatcher {
	return &challengeWatcher{
		chain:              chain,
		pollEventsInterval: interval,
		challenges:         make(map[protocol.AssertionId]*challenge),
	}
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
) {
	caller, err := challengeV2gen.NewEdgeChallengeManagerCaller(challengeManager.Address(), nil)
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

// TODO: Panic if something occurs
func (w *challengeWatcher) Watch(ctx context.Context) error {
	// Start from the latest confirmed assertion's creation block.
	// latestConfirmed, err := w.chain.LatestConfirmed(ctx)
	// if err != nil {
	// 	return err
	// }
	// assertionNode, ok := latestConfirmed.(*Assertion)
	// if !ok {
	// 	return errors.New("not ok")
	// }
	// inner, err := assertionNode.inner()
	// if err != nil {
	// 	return err
	// }
	// firstBlock := inner.CreatedAtBlock
	firstBlock := uint64(0)
	fromBlock := firstBlock

	// Some kind of backoff so as to not spam the node with requests.
	ticker := time.NewTicker(w.pollEventsInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// latestBlock, err := w.chain.backend.HeaderByNumber(ctx, nil)
			// if err != nil {
			// 	return err
			// }
			// toBlock := latestBlock.Number.Uint64()
			toBlock := uint64(0)

			challengeManager, err := w.chain.SpecChallengeManager(ctx)
			if err != nil {
				return err
			}
			topics := make([]common.Hash, 0, len(topicByName))
			for _, v := range topicByName {
				topics = append(topics, v)
			}
			var query = ethereum.FilterQuery{
				FromBlock: new(big.Int).SetUint64(fromBlock),
				ToBlock:   new(big.Int).SetUint64(toBlock),
				Addresses: []common.Address{challengeManager.Address()},
				Topics:    [][]common.Hash{topics, {}},
			}
			logs, err := w.chain.backend.FilterLogs(ctx, query)
			if err != nil {
				return err
			}
			if len(logs) == 0 {
				continue
			}
			// For each level zero edge creation event, get the edge type.
			// Then, use the edge type to add to a challenge set in the challenge watcher.
			for _, l := range logs {
				topicName, ok := nameByTopic[l.Topics[0]]
				if !ok {
					continue
				}

				if isConfirmationTopic(topicName) {
					// If edge is being confirmed, we check if it has a non-zero claimId.
					// If so, it is a level zero edge, and we keep track of it in the watcher.
				}

				switch {
				case topicName == "EdgeAdded":
					// Switch on the log type.
					edgeAdded, err := challengeManager.filterer.ParseEdgeAdded(l)
					if err != nil {
						return err
					}
					if protocol.EdgeType(edgeAdded.EType) == protocol.BlockChallengeEdge {
						if _, ok := w.challenges[edgeAdded.ClaimId]; !ok {
							w.challenges[edgeAdded.ClaimId] = &challenge{
								honestAncestorsBranch:          &ancestorsBranch{},
								confirmedLevelZeroEdgeClaimIds: newSet[protocol.ClaimId](),
							}
						}
					}
				}

				// Watcher needs access to the challenge manager. If it sees an edge it agrees with (honest),
				// it will then persist that in the honest ancestors branch. It needs to keep track of ancestors
				// in a special order.
			}
			fromBlock = toBlock
		case <-ctx.Done():
			return nil
		}
	}
}

func isConfirmationTopic(topicName string) bool {
	return topicName == edgeConfirmedByChildrenEvent ||
		topicName == edgeConfirmedByClaimEvent ||
		topicName == edgeConfirmedByOneStepProofEvent ||
		topicName == edgeConfirmedByTimeEvent
}

func newChallengeWatcher() *challengeWatcher {
	return &challengeWatcher{}
}
