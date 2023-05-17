package challengetree

import (
	"context"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
)

func (ht *HonestChallengeTree) ancestryPath(
	accum []protocol.EdgeId,
	curr protocol.EdgeSnapshot,
	queryingFor protocol.EdgeId,
) ([]protocol.EdgeId, bool) {
	edge, err := ht.metadataReader.GetEdge(context.Background(), queryingFor)
	if err != nil {
		return nil, false
	}
	// if curr.Id() == queryingFor {
	// 	return accum, true
	// }
	// if !hasChildren(curr) {
	// 	return accum, false
	// }
	// accum = append(accum, curr.Id())

	// // If the edge id we are querying for is a direct child of the current edge, we append
	// // the current edge to the ancestors list and return true.
	// if isDirectChild(curr, queryingFor) {
	// 	return accum, true
	// }
	// var lowerAncestors []protocol.EdgeId
	// var foundInLowerChildren bool
	// if !curr.LowerChildSnapshot().IsNone() {
	// 	lowerChildId := curr.LowerChildSnapshot().Unwrap()
	// 	lowerChild := ht.edges.Get(lowerChildId)
	// 	lowerAncestors, foundInLowerChildren = ht.ancestorQuery(
	// 		accum, lowerChild, queryingFor,
	// 	)
	// }
	// var upperAncestors []protocol.EdgeId
	// var foundInUpperChildren bool
	// if !curr.UpperChildSnapshot().IsNone() {
	// 	upperChildId := curr.UpperChildSnapshot().Unwrap()
	// 	upperChild := ht.edges.Get(upperChildId)
	// 	upperAncestors, foundInUpperChildren = ht.ancestorQuery(
	// 		accum, upperChild, queryingFor,
	// 	)
	// }
	// // If the edge we are querying for is found in the lower children,
	// // we return the ancestry along such path.
	// if foundInLowerChildren {
	// 	return lowerAncestors, true
	// }
	// // If the edge we are querying for is found in the upper children,
	// // we return the ancestry along such path.
	// if foundInUpperChildren {
	// 	return upperAncestors, true
	// }
	// return accum, false
}
