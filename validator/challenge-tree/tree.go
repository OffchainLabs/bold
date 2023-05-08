package challengetree

import (
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
)

type edgeId string
type claimId string
type originId string
type mutualId string
type commit string

type edge struct {
	id           edgeId
	edgeType     protocol.EdgeType
	startHeight  uint64
	startCommit  commit
	endHeight    uint64
	endCommit    commit
	originId     originId
	claimId      claimId
	lowerChildId edgeId
	upperChildId edgeId
	creationTime uint64
}

func (e *edge) computeMutualId() mutualId {
	return mutualId(e.id[:len(e.id)-1]) // Strip off the last char.
	// return mutualId(fmt.Sprintf(
	// 	"%d-%#x-%d-%#x-%d",
	// 	e.edgeType,
	// 	e.originId,
	// 	e.startHeight,
	// 	e.startCommit,
	// 	e.endHeight,
	// ))
}

// A challenge tree keeps track of edges whose history commitments the honest node agrees with.
// All edges tracked in this data structure are part of the same, top-level assertion challenge.
type challengeTree struct {
	timeRef            util.TimeReference
	edges              *threadsafe.Map[edgeId, *edge]
	mutualIds          *threadsafe.Map[mutualId, *threadsafe.Set[edgeId]]
	rivaledEdges       *threadsafe.Set[edgeId]
	computedPathTimers *threadsafe.Map[edgeId, uint64]
}
