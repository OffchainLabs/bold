package main

import (
	"container/list"
	"context"
	"errors"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/containers/option"
	"github.com/OffchainLabs/bold/solgen/go/challengeV2gen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type edge struct {
	id     common.Hash
	txHash common.Hash
	*challengeV2gen.ChallengeEdge
}

type protocolGraph struct {
	edgesByLevel          map[protocol.ChallengeLevel]map[common.Hash]*edge
	chalManager           *challengeV2gen.EdgeChallengeManager
	challengePeriodBlocks uint64
}

func (s *service) getEdgeClaimingAssertion(
	ctx context.Context,
	challengeCreationBlock uint64,
	claimingAssertion common.Hash,
) *edge {
	it, err := s.chalManager.FilterEdgeAdded(&bind.FilterOpts{
		Start:   challengeCreationBlock,
		End:     nil,
		Context: ctx,
	}, nil, nil, nil)
	defer func() {
		if err = it.Close(); err != nil {
			log.Error("Could not close filter iterator", "err", err)
		}
	}()
	for it.Next() {
		if it.Error() != nil {
			panic(err)
		}
		if it.Event.ClaimId == claimingAssertion {
			eg, err := s.chalManager.GetEdge(&bind.CallOpts{}, it.Event.EdgeId)
			if err != nil {
				panic(err)
			}
			return &edge{
				id:            it.Event.EdgeId,
				txHash:        it.Event.Raw.TxHash,
				ChallengeEdge: &eg,
			}
		}
	}
	return nil
}

func (s *service) buildProtocolGraphForChallenge(
	ctx context.Context, claimedAssertion common.Hash,
) (*protocolGraph, error) {
	// Scan for edge creations from parent's second child creation block up to and including
	// the block at which the item was confirmed.
	parentAssertion, err := s.rollupBindings.GetAssertion(&bind.CallOpts{}, claimedAssertion)
	if err != nil {
		return nil, err
	}

	eg := s.getEdgeClaimingAssertion(ctx, parentAssertion.SecondChildBlock, claimedAssertion)
	if protocol.EdgeStatus(eg.Status) != protocol.EdgeConfirmed {
		return nil, errors.New("edge is not confirmed")
	}
	upToConfirmedEdge := eg.id

	// Get the first edge that claims the assertion in the challenge, then get all edges
	// from that edge up to the block that first edge was confirmed.
	confirmedEdge, err := s.chalManager.GetEdge(&bind.CallOpts{}, upToConfirmedEdge)
	if err != nil {
		return nil, err
	}
	startBlock := parentAssertion.SecondChildBlock
	endBlock := confirmedEdge.ConfirmedAtBlock
	if endBlock <= startBlock {
		return nil, errors.New("end block is less than or equal to start block")
	}
	chalPeriodBlocks, err := s.chalManager.ChallengePeriodBlocks(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	graph := &protocolGraph{
		edgesByLevel:          make(map[protocol.ChallengeLevel]map[common.Hash]*edge),
		chalManager:           s.chalManager,
		challengePeriodBlocks: chalPeriodBlocks,
	}
	it, err := s.chalManager.FilterEdgeAdded(&bind.FilterOpts{
		Start:   startBlock,
		End:     &endBlock,
		Context: ctx,
	}, nil, nil, nil)
	defer func() {
		if err = it.Close(); err != nil {
			log.Error("Could not close filter iterator", "err", err)
		}
	}()
	for it.Next() {
		if it.Error() != nil {
			panic(err)
		}
		lvl := protocol.ChallengeLevel(it.Event.Level)
		eg, err := s.chalManager.GetEdge(&bind.CallOpts{}, it.Event.EdgeId)
		if err != nil {
			panic(err)
		}
		if _, ok := graph.edgesByLevel[lvl]; !ok {
			graph.edgesByLevel[lvl] = make(map[common.Hash]*edge)
		}
		graph.edgesByLevel[lvl][it.Event.EdgeId] = &edge{
			id:            it.Event.EdgeId,
			txHash:        it.Event.Raw.TxHash,
			ChallengeEdge: &eg,
		}
	}
	return graph, nil
}

func (pg *protocolGraph) terminalEdgesAtLvl(level protocol.ChallengeLevel) []*edge {
	terminalEdges := make([]*edge, 0)
	edgesAtLvl, ok := pg.edgesByLevel[level]
	if !ok {
		return terminalEdges
	}
	for _, eg := range edgesAtLvl {
		start := eg.StartHeight.Uint64()
		end := eg.EndHeight.Uint64()
		if start > end {
			continue
		}
		if end-start == 1 {
			terminalEdges = append(terminalEdges, eg)
		}
	}
	return terminalEdges
}

func (pg *protocolGraph) getClaimingEdgeAtLvl(level protocol.ChallengeLevel, claimedEdge common.Hash) *edge {
	edgesAtLvl, ok := pg.edgesByLevel[level]
	if !ok {
		return nil
	}
	for _, eg := range edgesAtLvl {
		if eg.ClaimId == claimedEdge {
			return eg
		}
	}
	return nil
}

func (pg *protocolGraph) extractEssentialBranchesAtLvl(level protocol.ChallengeLevel, root *edge) [][]*edge {
	allPaths := make([][]*edge, 0)
	if pg.onchainTimer(root) < pg.challengePeriodBlocks {
		return allPaths
	}
	edgesAtLvl := pg.edgesByLevel[level]
	type visited struct {
		essentialEdge *edge
		path          []*edge
	}
	stack := newStack[*visited]()
	stack.push(&visited{
		essentialEdge: root,
		path:          []*edge{root},
	})
	for stack.len() > 0 {
		curr := stack.pop().Unwrap()
		currentNode, path := curr.essentialEdge, curr.path

		// Node not viable for refund.
		if pg.onchainTimer(currentNode) < pg.challengePeriodBlocks {
			continue
		}
		hasChildren := currentNode.LowerChildId != common.Hash{} ||
			currentNode.UpperChildId != common.Hash{}
		if hasChildren {
			lowerChild := edgesAtLvl[currentNode.LowerChildId]
			upperChild := edgesAtLvl[currentNode.UpperChildId]
			stack.push(&visited{
				essentialEdge: lowerChild,
				path:          append(path, lowerChild),
			})
			stack.push(&visited{
				essentialEdge: upperChild,
				path:          append(path, upperChild),
			})
		} else {
			allPaths = append(allPaths, path)
		}
	}
	return allPaths
}

func (pg *protocolGraph) onchainTimer(eg *edge) uint64 {
	onchainEdge, err := pg.chalManager.GetEdge(&bind.CallOpts{}, eg.id)
	if err != nil {
		panic(err)
	}
	return onchainEdge.TotalTimeUnrivaledCache
}

type stack[T any] struct {
	dll *list.List
}

func newStack[T any]() *stack[T] {
	return &stack[T]{dll: list.New()}
}

func (s *stack[T]) len() int {
	return s.dll.Len()
}

func (s *stack[T]) push(x T) {
	s.dll.PushBack(x)
}

func (s *stack[T]) pop() option.Option[T] {
	if s.dll.Len() == 0 {
		return option.None[T]()
	}
	tail := s.dll.Back()
	val := tail.Value
	s.dll.Remove(tail)
	return option.Some(val.(T))
}
