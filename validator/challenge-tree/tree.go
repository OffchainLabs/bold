package challengetree

import (
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
)

type metadataReader interface {
	TopLevelAssertion(edgeId protocol.EdgeId) (protocol.AssertionId, error)
}

// A challenge tree keeps track of edges whose history commitments the honest node agrees with.
// All edges tracked in this data structure are part of the same, top-level assertion challenge.
type challengeTree struct {
	edges                        *threadsafe.Map[protocol.EdgeId, protocol.EdgeSnapshot]
	mutualIds                    *threadsafe.Map[protocol.MutualId, *threadsafe.Set[protocol.EdgeId]]
	rivaledEdges                 *threadsafe.Set[protocol.EdgeId]
	topLevelAssertionId          protocol.AssertionId
	honestBlockChalLevelZeroEdge util.Option[protocol.SpecEdge]
	metadataReader               metadataReader
}
