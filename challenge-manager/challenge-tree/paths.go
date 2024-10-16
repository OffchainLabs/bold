package challengetree

import (
	"container/heap"
	"container/list"
	"context"
	"fmt"

	protocol "github.com/offchainlabs/bold/chain-abstraction"
	"github.com/offchainlabs/bold/containers"
	"github.com/offchainlabs/bold/containers/option"
	"github.com/pkg/errors"
)

type ComputePathWeightArgs struct {
	Child    protocol.EdgeId
	Ancestor protocol.EdgeId
	BlockNum uint64
}

// ComputePathWeight from a child edge to a specified ancestor edge. A weight is the sum of the local timers
// of all edges along the path.
//
// Invariant: assumes ComputeAncestors returns a list of ancestors ordered from child to parent,
// not including the edge id we are querying ancestors for.
func (ht *RoyalChallengeTree) ComputePathWeight(
	ctx context.Context,
	args ComputePathWeightArgs,
) (uint64, error) {
	child, ok := ht.edges.TryGet(args.Child)
	if !ok {
		return 0, fmt.Errorf("child edge not yet tracked %#x", args.Child.Hash)
	}
	if !ht.edges.Has(args.Ancestor) {
		return 0, fmt.Errorf("ancestor not yet tracked %#x", args.Ancestor.Hash)
	}
	localTimer, err := ht.LocalTimer(child, args.BlockNum)
	if err != nil {
		return 0, err
	}
	if args.Child == args.Ancestor {
		return localTimer, nil
	}
	ancestors, err := ht.ComputeAncestors(ctx, args.Child, args.BlockNum)
	if err != nil {
		return 0, err
	}
	pathWeight := localTimer
	found := false
	for _, an := range ancestors {
		localTimer, err := ht.LocalTimer(an, args.BlockNum)
		if err != nil {
			return 0, err
		}
		pathWeight += localTimer
		if an.Id() == args.Ancestor {
			found = true
			break
		}
	}
	if !found {
		return 0, errors.New("expected ancestor not found in computed ancestors list")
	}
	return pathWeight, nil
}

type essentialPath []protocol.EdgeId

type isConfirmableArgs struct {
	essentialNode         protocol.EdgeId
	confirmationThreshold uint64
	blockNum              uint64
}

// Find all the paths down from an essential node, and
// compute the local timer of each edge along the path. This is
// a recursive computation that goes down the tree rooted at the essential
// node and ends once it finds edges that either do not have children,
// or are terminal nodes that end in children that are incorrectly constructed
// or non-essential.
//
// After the paths are computed, we then compute the path weight of each
// and insert each weight into a min-heap. If the min element of this heap
// has a weight >= the confirmation threshold, the
// essential node is then confirmable.
//
// Note: the specified argument essential node must indeed be essential, otherwise,
// this function will error.
func (ht *RoyalChallengeTree) IsConfirmableEssentialNode(
	ctx context.Context,
	args isConfirmableArgs,
) (bool, []essentialPath, uint64, error) {
	essentialNode, ok := ht.edges.TryGet(args.essentialNode)
	if !ok {
		return false, nil, 0, fmt.Errorf("essential node not found")
	}
	if essentialNode.ClaimId().IsNone() {
		return false, nil, 0, fmt.Errorf("specified input argument %#x is not essential", args.essentialNode.Hash)
	}
	essentialPaths, essentialTimers, err := ht.findEssentialPaths(
		ctx,
		essentialNode,
		args.blockNum,
	)
	if err != nil {
		return false, nil, 0, err
	}
	if len(essentialPaths) == 0 || len(essentialTimers) == 0 {
		return false, nil, 0, fmt.Errorf("no essential paths found")
	}
	// An essential node is confirmable if all of its essential paths
	// down the tree have a path weight >= the confirmation threshold.
	// To do this, we compute the path weight of each path and insert
	// into a min heap. Then, it is sufficient to check that the minimum
	// element of the heap is >= the confirmation threshold.
	pathWeights := newPathWeightMinHeap()
	for _, timers := range essentialTimers {
		pathWeight := uint64(0)
		for _, timer := range timers {
			pathWeight += timer
		}
		pathWeights.Push(pathWeight)
	}
	if pathWeights.items.Len() == 0 {
		return false, nil, 0, fmt.Errorf("no path weights computed")
	}
	minWeight := pathWeights.Pop()
	allEssentialPathsConfirmable := minWeight >= args.confirmationThreshold
	return allEssentialPathsConfirmable, essentialPaths, minWeight, nil
}

type essentialLocalTimers []uint64

// Use a depth-first-search approach (DFS) to gather the
// essential branches of the protocol graph. We manage our own
// visitor stack to avoid recursion.
//
// Invariant: the input node must be essential.
func (ht *RoyalChallengeTree) findEssentialPaths(
	ctx context.Context,
	essentialNode protocol.ReadOnlyEdge,
	blockNum uint64,
) ([]essentialPath, []essentialLocalTimers, error) {
	allPaths := make([]essentialPath, 0)
	allTimers := make([]essentialLocalTimers, 0)

	type visited struct {
		essentialNode protocol.ReadOnlyEdge
		path          essentialPath
		localTimers   essentialLocalTimers
	}
	stack := newStack[*visited]()

	localTimer, err := ht.LocalTimer(essentialNode, blockNum)
	if err != nil {
		return nil, nil, err
	}

	stack.push(&visited{
		essentialNode: essentialNode,
		path:          essentialPath{essentialNode.Id()},
		localTimers:   essentialLocalTimers{localTimer},
	})

	for stack.len() > 0 {
		curr := stack.pop().Unwrap()
		currentNode, currentTimers, path := curr.essentialNode, curr.localTimers, curr.path
		isClaimedEdge, claimingEdge := ht.isClaimedEdge(ctx, currentNode)

		hasChildren, err := currentNode.HasChildren(ctx)
		if err != nil {
			return nil, nil, err
		}
		if hasChildren {
			lowerChildIdOpt, err := currentNode.LowerChild(ctx)
			if err != nil {
				return nil, nil, err
			}
			upperChildIdOpt, err := currentNode.UpperChild(ctx)
			if err != nil {
				return nil, nil, err
			}
			lowerChildId, upperChildId := lowerChildIdOpt.Unwrap(), upperChildIdOpt.Unwrap()
			lowerChild, ok := ht.edges.TryGet(lowerChildId)
			if !ok {
				return nil, nil, fmt.Errorf("lower child not yet tracked")
			}
			upperChild, ok := ht.edges.TryGet(upperChildId)
			if !ok {
				return nil, nil, fmt.Errorf("upper child not yet tracked")
			}
			lowerTimer, err := ht.LocalTimer(lowerChild, blockNum)
			if err != nil {
				return nil, nil, err
			}
			upperTimer, err := ht.LocalTimer(upperChild, blockNum)
			if err != nil {
				return nil, nil, err
			}
			lowerPath := append(path, lowerChildId)
			upperPath := append(path, upperChildId)
			lowerTimers := append(currentTimers, lowerTimer)
			upperTimers := append(currentTimers, upperTimer)
			stack.push(&visited{
				essentialNode: lowerChild,
				path:          lowerPath,
				localTimers:   lowerTimers,
			})
			stack.push(&visited{
				essentialNode: upperChild,
				path:          upperPath,
				localTimers:   upperTimers,
			})
			continue
		} else if isClaimedEdge {
			// Figure out if the node is a terminal node that has a refinement, in which
			// case we need to continue the search down the next challenge level,
			claimingEdgeTimer, err := ht.LocalTimer(claimingEdge, blockNum)
			if err != nil {
				return nil, nil, err
			}
			claimingPath := append(path, claimingEdge.Id())
			claimingTimers := append(currentTimers, claimingEdgeTimer)
			stack.push(&visited{
				essentialNode: claimingEdge,
				path:          claimingPath,
				localTimers:   claimingTimers,
			})
			continue
		}

		// Otherwise, the node is a qualified leaf and we can push to the list of paths
		// and all the timers of the path.
		// Onchain actions expect ordered paths from leaf to root, so we
		// preserve that ordering to make it easier for callers to use this data.
		containers.Reverse(path)
		containers.Reverse(currentTimers)
		allPaths = append(allPaths, path)
		allTimers = append(allTimers, currentTimers)
	}
	return allPaths, allTimers, nil
}

func (ht *RoyalChallengeTree) isClaimedEdge(ctx context.Context, edge protocol.ReadOnlyEdge) (bool, protocol.ReadOnlyEdge) {
	if isProofNode(ctx, edge) {
		return false, nil
	}
	if !hasLengthOne(edge) {
		return false, nil
	}
	// Note: the specification requires that the claiming edge is correctly constructed.
	// This is not checked here, because the honest validator only tracks
	// essential edges as an invariant.
	claimingEdge, ok := ht.findClaimingEdge(edge.Id())
	if !ok {
		return false, nil
	}
	return true, claimingEdge
}

// Proof nodes are nodes that have length one at the lowest challenge level.
func isProofNode(ctx context.Context, edge protocol.ReadOnlyEdge) bool {
	isSmallStep := edge.GetChallengeLevel() == protocol.ChallengeLevel(edge.GetTotalChallengeLevels(ctx)-1)
	return isSmallStep && hasLengthOne(edge)
}

type uint64Heap []uint64

func (h uint64Heap) Len() int           { return len(h) }
func (h uint64Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h uint64Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h uint64Heap) Peek() uint64 {
	return h[0]
}

func (h *uint64Heap) Push(x any) {
	*h = append(*h, x.(uint64))
}

func (h *uint64Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type pathWeightMinHeap struct {
	items *uint64Heap
}

func newPathWeightMinHeap() *pathWeightMinHeap {
	items := &uint64Heap{}
	heap.Init(items)
	return &pathWeightMinHeap{items}
}

func (h *pathWeightMinHeap) Len() int {
	return h.items.Len()
}

func (h *pathWeightMinHeap) Push(item uint64) {
	heap.Push(h.items, item)
}

func (h *pathWeightMinHeap) Pop() uint64 {
	return heap.Pop(h.items).(uint64)
}

func (h *pathWeightMinHeap) Peek() option.Option[uint64] {
	if h.items.Len() == 0 {
		return option.None[uint64]()
	}
	return option.Some(h.items.Peek())
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
