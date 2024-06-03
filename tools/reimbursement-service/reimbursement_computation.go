package main

import (
	"context"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/ethereum/go-ethereum/common"
)

type reimbursementArgs struct {
	claimedAssertion common.Hash
	claimedEdge      common.Hash
	itemTyp          claimType
	challengeLvl     protocol.ChallengeLevel
}

func (s *service) executeReimbursement(ctx context.Context, args reimbursementArgs) {
	var protocolGraph *protocolGraph
	var challengeRoot *edge
	// If assertion, build the protocol graph for the challenge.
	if args.itemTyp == assertionTyp {
		graph, err := s.buildProtocolGraphForChallenge(ctx, args.claimedAssertion)
		if err != nil {
			panic(err)
		}
		s.protocolGraphLock.Lock()
		protocolGraph = graph
		s.protocolGraphLock.Unlock()
		challengeRoot = protocolGraph.getClaimingEdgeAtLvl(protocol.NewBlockChallengeLevel(), args.claimedAssertion)
	} else {
		s.protocolGraphLock.RLock()
		graph, ok := s.protocolGraphsByClaimedAssertion[args.claimedAssertion]
		s.protocolGraphLock.RUnlock()
		if !ok {
			panic("Could not find protocol graph for claimed assertion")
		}
		protocolGraph = graph
		challengeRoot = protocolGraph.getClaimingEdgeAtLvl(args.challengeLvl, args.claimedEdge)
	}
	// Get all the terminal nodes descending from item at the current level.
	terminalNodes := protocolGraph.terminalEdgesAtLvl(protocol.ChallengeLevel(challengeRoot.Level))
	confirmedRefinements := make([]*edge, 0)
	unconfirmedRefinements := make([]*edge, 0)
	// Get all the nodes that have a claim id equal to a terminal node extracted above.
	for _, terminalNode := range terminalNodes {
		refinement := protocolGraph.getClaimingEdgeAtLvl(protocol.ChallengeLevel(challengeRoot.Level)+1, terminalNode.id)
		if refinement == nil {
			continue
		}
		// Filter those refinement nodes that are confirmed vs. unconfirmed.
		if protocol.EdgeStatus(refinement.Status) == protocol.EdgeConfirmed {
			confirmedRefinements = append(confirmedRefinements, refinement)
		} else {
			unconfirmedRefinements = append(unconfirmedRefinements, refinement)
		}
	}

	// Add the unconfirmed refinement nodes to the watchset, then remove A from the watchset.
	for _, unconfirmed := range unconfirmedRefinements {
		s.watchList.Put(unconfirmed.id, edgeTyp)
	}
	// Process service fee payments and remove items.
	if args.itemTyp == assertionTyp {
		s.watchList.Delete(args.claimedAssertion)
		s.serviceFeePaymentRequests <- &serviceFeePaymentRequest{
			claimedItemTyp: args.itemTyp,
			claimedItem:    args.claimedAssertion,
		}
	} else {
		s.watchList.Delete(args.claimedEdge)
		var eg *edge
		s.protocolGraphLock.RLock()
		graph := s.protocolGraphsByClaimedAssertion[args.claimedAssertion]
		for _, edgesByLvl := range graph.edgesByLevel {
			for hash, edg := range edgesByLvl {
				if hash == args.claimedEdge {
					eg = edg
					break
				}
			}
		}
		s.protocolGraphLock.RUnlock()
		s.serviceFeePaymentRequests <- &serviceFeePaymentRequest{
			claimedItemTyp: args.itemTyp,
			claimedItem:    args.claimedEdge,
			edgeCreationTx: eg.txHash,
		}
	}

	// For each confirmed refinement, execute the reimbursement function.
	for _, refinement := range confirmedRefinements {
		s.executeReimbursement(ctx, reimbursementArgs{
			claimedAssertion: args.claimedAssertion,
			claimedEdge:      refinement.ClaimId,
			challengeLvl:     protocol.ChallengeLevel(refinement.Level),
			itemTyp:          edgeTyp,
		})
	}

	// Process gas payments.
	// Extract all branches descending from item at the current level that have
	// big enough timers. They should have their gas refunded.
	confirmableBranchesAtLevel := protocolGraph.extractEssentialBranchesAtLvl(
		protocol.ChallengeLevel(challengeRoot.Level),
		challengeRoot,
	)
	// Submit a payment to process.
	s.gasPaymentRequests <- &gasPaymentRequest{
		confirmedAssertionInChallenge: args.claimedAssertion,
		essentialChallengeBranches:    confirmableBranchesAtLevel,
	}
}
