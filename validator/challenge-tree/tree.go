package challengetree

import (
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
)

// A challenge tree keeps track of edges whose history commitments the honest node agrees with.
// All edges tracked in this data structure are part of the same, top-level assertion challenge.
type challengeTree struct {
	edges        *threadsafe.Map[protocol.EdgeId, protocol.EdgeGetter]
	mutualIds    *threadsafe.Map[protocol.MutualId, *threadsafe.Set[protocol.EdgeId]]
	rivaledEdges *threadsafe.Set[protocol.EdgeId]
}
