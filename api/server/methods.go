package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/OffchainLabs/bold/api"
	"github.com/OffchainLabs/bold/api/db"
	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gorilla/mux"
)

// Healthz checks if the API server is ready to serve queries. Returns 200 if it is ready.
//
// method:
// - GET
// - /api/v1/db/healthz
func (s *Server) Healthz() {

}

// ListAssertions up to chain head
//
// method:
// - GET
// - /api/v1/assertions
//
// request query params:
//   - limit: the max number of items in the response
//   - offset: the offset index in the DB
//   - inbox_max_count: assertions that have a specified value for InboxMaxCount
//   - from_block_number: items that were created since a specific block number. Defaults to latest confirmed assertion
//   - to_block_number: caps the response to assertions up to and including a block number
//
// response:
// - []*JsonAssertion
func (s *Server) ListAssertions(r *http.Request, w http.ResponseWriter) {
	opts := make([]db.AssertionOption, 0)
	query := r.URL.Query()
	if val, ok := query["limit"]; ok && len(val) > 0 {
		if v, err := strconv.Atoi(val[0]); err == nil {
			opts = append(opts, db.WithAssertionLimit(v))
		}
	}
	if val, ok := query["offset"]; ok && len(val) > 0 {
		if v, err := strconv.Atoi(val[0]); err == nil {
			opts = append(opts, db.WithAssertionOffset(v))
		}
	}
	if val, ok := query["inbox_max_count"]; ok && len(val) > 0 {
		opts = append(opts, db.WithInboxMaxCount(strings.Join(val, "")))
	}
	if val, ok := query["from_block_number"]; ok && len(val) > 0 {
		if v, err := strconv.ParseUint(val[0], 10, 64); err == nil {
			opts = append(opts, db.FromAssertionCreationBlock(v))
		}
	}
	if val, ok := query["to_block_number"]; ok && len(val) > 0 {
		if v, err := strconv.ParseUint(val[0], 10, 64); err == nil {
			opts = append(opts, db.ToAssertionCreationBlock(v))
		}
	}
	assertions, err := s.backend.GetAssertions(r.Context(), opts...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Could not get assertions: %v", err)))
		return
	}
	if err = json.NewEncoder(w).Encode(assertions); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Could not write assertions response: %v", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
}

// AssertionByIdentifier since the latest confirmed assertion.
//
// method:
// - GET
// - /api/v1/assertion/<identifier>
//
// identifier options:
// - an assertion hash (0x-prefixed): gets the assertion by hash
// - "latest-confirmed": gets the latest confirmed assertion
//
// response:
// - *JsonAssertion
func (s *Server) AssertionByIdentifier(r *http.Request, w http.ResponseWriter) {
	vars := mux.Vars(r)
	identifier := vars["identifier"]

	var assertion *api.JsonAssertion
	opts := []db.AssertionOption{
		db.WithAssertionLimit(1),
	}
	if identifier == "latest-confirmed" {
		a, err := s.backend.LatestConfirmedAssertion(r.Context())
		if err != nil {
			return
		}
		assertion = a
	} else {
		// Otherwise, get the assertion by hash.
		hash, err := hexutil.Decode(identifier)
		if err != nil {
			return
		}
		opts = append(opts, db.WithAssertionHash(protocol.AssertionHash{Hash: common.BytesToHash(hash)}))
		assertions, err := s.backend.GetAssertions(r.Context(), opts...)
		if err != nil {
			return
		}
		if len(assertions) != 1 {
			return
		}
		assertion = assertions[0]
	}
	if err := json.NewEncoder(w).Encode(assertion); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Could not write assertion response: %v", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ChallengeByAssertionHash fetches information about a challenge on a specific assertion hash
// method:
// - GET
// - /api/v1/challenge/<assertion-hash>
//
// identifier options:
// - 0x-prefixed assertion hash
//
// response:
// - *JsonChallenge
func (s *Server) ChallengeByAssertionHash(r *http.Request, w http.ResponseWriter) {
	vars := mux.Vars(r)
	identifier := vars["assertion-hash"]
	_ = identifier
}

// AllChallengeEdges fetches all the edges corresponding to a challenged
// assertion with a specific hash. This assertion hash must be the "parent assertion"
// of two child assertions that originated a challenge.
//
// method:
// - GET
// - /api/v1/challenge/<assertion-hash>/edges
//
// identifier options:
// - 0x-prefixed assertion hash
//
// request query params:
// - limit: the max number of items in the response
// - offset: the offset index in the DB
// - status: filter edges that have status "confirmed", "confirmable", or "pending"
// - honest: boolean true or false to filter out honest vs. evil edges. If not set, fetches all edges in the challenge.
// - root_edges: boolean true or false to filter out only root edges (those that have a claim id)
// - from_block_number: items that were created since a specific block number. Defaults to challenge creation block
// - to_block_number: caps the response to edges up to and including a block number
// - origin_id: edges that have a 0x-prefixed origin id
// - mutual_id: edges that have a 0x-prefixed mutual id
// - claim_id: edges that have a 0x-prefixed claim id
// - start_commitment: edges with a start history commitment of format "height:hash", such as 32:0xdeadbeef
// - end_commitment: edges with an end history commitment of format "height:hash", such as 32:0xdeadbeef
// - challenge_level: edges in a specific challenge level. level 0 is the block challenge level
// - to_block_number: caps the response to edges up to and including a block number
// response:
// - []*JsonEdge
func (s *Server) AllChallengeEdges() {

}

// EdgeByIdentifier fetches an edge by its specific id in a challenge.
//
// method:
// - GET
// - /api/v1/challenge/<assertion-hash>/edges/<edge-id>
//
// identifier options:
// - 0x-prefixed assertion hash
// - 0x-prefixed edge id
//
// response:
// - *JsonEdge
func (s *Server) EdgeByIdentifier(r *http.Request, w http.ResponseWriter) {
	vars := mux.Vars(r)
	assertionHashStr := vars["assertion-hash"]
	edgeIdStr := vars["edge-id"]
	hash, err := hexutil.Decode(assertionHashStr)
	if err != nil {
		return
	}
	id, err := hexutil.Decode(edgeIdStr)
	if err != nil {
		return
	}
	assertionHash := protocol.AssertionHash{Hash: common.BytesToHash(hash)}
	edgeId := protocol.EdgeId{Hash: common.BytesToHash(id)}
	edges, err := s.backend.GetEdges(
		r.Context(),
		db.WithLimit(1),
		db.WithEdgeAssertionHash(assertionHash),
		db.WithId(edgeId),
	)
	if err != nil {
		return
	}
	if len(edges) != 1 {
		return
	}
	if err := json.NewEncoder(w).Encode(edges[0]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Could not write edge response: %v", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
}

// EdgeByHistoryCommitment fetches an edge by its specific history commitment in a challenge.
//
// method:
// - GET
// - /api/v1/challenge/<assertion-hash>/edges/<history-commitment>
//
// identifier options:
//   - 0x-prefixed assertion hash
//   - history commitment with the format startheight:starthash:endheight:endhash, such as
//     0:0xdeadbeef:32:0xdeadbeef
//
// response:
// - *JsonEdge
func (s *Server) EdgeByHistoryCommitment(r *http.Request, w http.ResponseWriter) {
	vars := mux.Vars(r)
	historyCommitment := vars["history-commitment"]
	parts := strings.Split(historyCommitment, ":")
	if len(parts) != 4 {
		return
	}

	// Extract the parts
	startHeightStr := parts[0]
	startHashStr := parts[1]
	endHeightStr := parts[2]
	endHashStr := parts[3]
}

// MiniStakes fetches all the mini-stakes present in a single challenged assertion.
//
// method:
// - GET
// - /api/v1/challenge/<assertion-hash>/ministakes
//
// identifier options:
//   - 0x-prefixed assertion hash
//
// request query params:
// - limit: the max number of items in the response
// - offset: the offset index in the DB
// - challenge_level: items in a specific challenge level. level 0 is the block challenge level
// response:
// - []*MiniStake
func (s *Server) MiniStakes(r *http.Request, w http.ResponseWriter) {
	vars := mux.Vars(r)
	assertionHashStr := vars["assertion-hash"]
	hash, err := hexutil.Decode(assertionHashStr)
	if err != nil {
		return
	}
	assertionHash := protocol.AssertionHash{Hash: common.BytesToHash(hash)}
	query := r.URL.Query()
	opts := make([]db.EdgeOption, 0)
	if val, ok := query["limit"]; ok && len(val) > 0 {
		if v, err := strconv.Atoi(val[0]); err == nil {
			opts = append(opts, db.WithLimit(v))
		}
	}
	if val, ok := query["offset"]; ok && len(val) > 0 {
		if v, err := strconv.Atoi(val[0]); err == nil {
			opts = append(opts, db.WithOffset(v))
		}
	}
	if val, ok := query["challenge_level"]; ok && len(val) > 0 {
		if v, err := strconv.ParseUint(val[0], 10, 8); err == nil {
			opts = append(opts, db.WithChallengeLevel(uint8(v)))
		}
	}
	miniStakes, err := s.backend.GetMiniStakes(r.Context(), assertionHash, opts...)
	if err != nil {
		return
	}
	if err := json.NewEncoder(w).Encode(miniStakes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Could not write ministakes response: %v", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
}
