package challengetree

import (
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	"github.com/ethereum/go-ethereum/common"
)

// TrackedEdge defines an interface for a minimal set of getters we need on
// an edge a challenge tree is tracking.
type TrackedEdge interface {
	Id() protocol.EdgeId
	EdgeType() protocol.EdgeType
	StartCommitment() (uint64, common.Hash)
	EndCommitment() (uint64, common.Hash)
	OriginId() protocol.OriginId
	LowerChild() protocol.EdgeId
	UpperChild() protocol.EdgeId
	CreatedAtBlock() uint64
	ComputeMutualId() protocol.MutualId
}

// A challenge tree keeps track of edges whose history commitments the honest node agrees with.
// All edges tracked in this data structure are part of the same, top-level assertion challenge.
type challengeTree struct {
	edges        *threadsafe.Map[protocol.EdgeId, TrackedEdge]
	mutualIds    *threadsafe.Map[protocol.MutualId, *threadsafe.Set[protocol.EdgeId]]
	rivaledEdges *threadsafe.Set[protocol.EdgeId]
}
