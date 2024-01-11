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
	"github.com/OffchainLabs/bold/state-commitments/history"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gorilla/mux"
)

var contentType = "application/json"

// Healthz checks if the API server is ready to serve queries. Returns 200 if it is ready.
//
// method:
// - GET
// - /api/v1/db/healthz
func (s *Server) Healthz(r *http.Request, w http.ResponseWriter) {
	// TODO: Respond with a 503 if the client the BOLD validator is
	// connected to is syncing.
	w.WriteHeader(http.StatusOK)
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
//   - challenged: fetch only assertions that have been challenged
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
	if _, ok := query["challenged"]; ok {
		opts = append(opts, db.WithChallenge())
	}
	assertions, err := s.backend.GetAssertions(r.Context(), opts...)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not get assertions from backend: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", contentType)
	if err = json.NewEncoder(w).Encode(assertions); err != nil {
		http.Error(w, fmt.Sprintf("Could not write response: %v", err), http.StatusInternalServerError)
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
			http.Error(w, fmt.Sprintf("Could not get latest confirmed assertion: %v", err), http.StatusInternalServerError)
			return
		}
		assertion = a
	} else {
		// Otherwise, get the assertion by hash.
		hash, err := hexutil.Decode(identifier)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not parse assertion hash: %v", err), http.StatusBadRequest)
			return
		}
		opts = append(opts, db.WithAssertionHash(protocol.AssertionHash{Hash: common.BytesToHash(hash)}))
		assertions, err := s.backend.GetAssertions(r.Context(), opts...)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not get assertions from backend: %v", err), http.StatusInternalServerError)
			return
		}
		if len(assertions) != 1 {
			http.Error(
				w,
				fmt.Sprintf("Got more than 1 matching assertion: got %d", len(assertions)),
				http.StatusInternalServerError,
			)
			return
		}
		assertion = assertions[0]
	}
	w.Header().Set("Content-Type", contentType)
	if err := json.NewEncoder(w).Encode(assertion); err != nil {
		http.Error(w, fmt.Sprintf("Could not write response: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
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
// - from_block_number: items that were created since a specific block number.
// - to_block_number: caps the response to edges up to a block number
// - origin_id: edges that have a 0x-prefixed origin id
// - mutual_id: edges that have a 0x-prefixed mutual id
// - claim_id: edges that have a 0x-prefixed claim id
// - start_commitment: edges with a start history commitment of format "height:hash", such as 32:0xdeadbeef
// - end_commitment: edges with an end history commitment of format "height:hash", such as 32:0xdeadbeef
// - challenge_level: edges in a specific challenge level. level 0 is the block challenge level
// response:
// - []*JsonEdge
func (s *Server) AllChallengeEdges(r *http.Request, w http.ResponseWriter) {
	opts := make([]db.EdgeOption, 0)
	query := r.URL.Query()
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
	if val, ok := query["status"]; ok && len(val) > 0 {
		status := parseEdgeStatus(strings.Join(val, ""))
		opts = append(opts, db.WithEdgeStatus(status))
	}
	if _, ok := query["honest"]; ok {
		opts = append(opts, db.WithHonestEdges())
	}
	if _, ok := query["root_edges"]; ok {
		opts = append(opts, db.WithRootEdges())
	}
	if val, ok := query["from_block_number"]; ok && len(val) > 0 {
		if v, err := strconv.ParseUint(val[0], 10, 64); err == nil {
			opts = append(opts, db.FromEdgeCreationBlock(v))
		}
	}
	if val, ok := query["to_block_number"]; ok && len(val) > 0 {
		if v, err := strconv.ParseUint(val[0], 10, 64); err == nil {
			opts = append(opts, db.ToEdgeCreationBlock(v))
		}
	}
	if val, ok := query["origin_id"]; ok && len(val) > 0 {
		hash, err := hexutil.Decode(strings.Join(val, ""))
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not parse origin_id: %v", err), http.StatusBadRequest)
			return
		}
		opts = append(opts, db.WithOriginId(protocol.OriginId(common.BytesToHash(hash))))
	}
	if val, ok := query["mutual_id"]; ok && len(val) > 0 {
		hash, err := hexutil.Decode(strings.Join(val, ""))
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not parse mutual_id: %v", err), http.StatusBadRequest)
			return
		}
		opts = append(opts, db.WithMutualId(protocol.MutualId(common.BytesToHash(hash))))
	}
	if val, ok := query["claim_id"]; ok && len(val) > 0 {
		hash, err := hexutil.Decode(strings.Join(val, ""))
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not parse claim_id: %v", err), http.StatusBadRequest)
			return
		}
		opts = append(opts, db.WithClaimId(protocol.ClaimId(common.BytesToHash(hash))))
	}
	if val, ok := query["start_commitment"]; ok && len(val) > 0 {
		commitStr := strings.Join(val, "")
		commitParts := strings.Split(commitStr, ":")
		if len(commitParts) != 2 {
			http.Error(w, "Wrong start history commitment format, wanted height:hash", http.StatusBadRequest)
			return
		}
		startHeight, err := strconv.ParseUint(commitParts[0], 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not parse start commit height: %v", err), http.StatusBadRequest)
			return
		}
		startHash, err := hexutil.Decode(commitParts[1])
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not parse start commit hash: %v", err), http.StatusBadRequest)
			return
		}
		opts = append(opts, db.WithStartHistoryCommitment(history.History{
			Height: startHeight,
			Merkle: common.BytesToHash(startHash),
		}))
	}
	if val, ok := query["end_commitment"]; ok && len(val) > 0 {
		commitStr := strings.Join(val, "")
		commitParts := strings.Split(commitStr, ":")
		if len(commitParts) != 2 {
			http.Error(w, "Wrong start history commitment format, wanted height:hash", http.StatusBadRequest)
			return
		}
		endHeight, err := strconv.ParseUint(commitParts[0], 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not parse end commit height: %v", err), http.StatusBadRequest)
			return
		}
		endHash, err := hexutil.Decode(commitParts[1])
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not parse end commit hash: %v", err), http.StatusBadRequest)
			return
		}
		opts = append(opts, db.WithEndHistoryCommitment(history.History{
			Height: endHeight,
			Merkle: common.BytesToHash(endHash),
		}))
	}
	if val, ok := query["challenge_level"]; ok && len(val) > 0 {
		if v, err := strconv.ParseUint(val[0], 10, 8); err == nil {
			opts = append(opts, db.WithChallengeLevel(uint8(v)))
		}
	}
	edges, err := s.backend.GetEdges(r.Context(), opts...)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not get edges from backend: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", contentType)
	if err := json.NewEncoder(w).Encode(edges); err != nil {
		http.Error(w, fmt.Sprintf("Could not write response: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func parseEdgeStatus(str string) protocol.EdgeStatus {
	return protocol.EdgeConfirmed
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
		http.Error(w, fmt.Sprintf("Could not parse assertion hash: %v", err), http.StatusBadRequest)
		return
	}
	id, err := hexutil.Decode(edgeIdStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not parse edge id: %v", err), http.StatusBadRequest)
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
		http.Error(w, fmt.Sprintf("Could not get edges from backend: %v", err), http.StatusInternalServerError)
		return
	}
	if len(edges) != 1 {
		http.Error(w, fmt.Sprintf("Got more edges than expected: %d", len(edges)), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", contentType)
	if err := json.NewEncoder(w).Encode(edges[0]); err != nil {
		http.Error(w, fmt.Sprintf("Could not write response: %v", err), http.StatusInternalServerError)
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
	assertionHashStr := vars["assertion-hash"]
	hash, err := hexutil.Decode(assertionHashStr)
	if err != nil {
		return
	}
	assertionHash := protocol.AssertionHash{Hash: common.BytesToHash(hash)}
	historyCommitment := vars["history-commitment"]
	parts := strings.Split(historyCommitment, ":")
	if len(parts) != 4 {
		return
	}
	startHeight, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return
	}
	startHash, err := hexutil.Decode(parts[1])
	if err != nil {
		return
	}
	endHeight, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return
	}
	endHash, err := hexutil.Decode(parts[3])
	if err != nil {
		return
	}
	edges, err := s.backend.GetEdges(
		r.Context(),
		db.WithEdgeAssertionHash(assertionHash),
		db.WithStartHistoryCommitment(history.History{
			Height: startHeight,
			Merkle: common.BytesToHash(startHash),
		}),
		db.WithEndHistoryCommitment(history.History{
			Height: endHeight,
			Merkle: common.BytesToHash(endHash),
		}),
		db.WithLimit(1),
	)
	if err != nil {
		return
	}
	if len(edges) != 1 {
		return
	}
	w.Header().Set("Content-Type", contentType)
	if err := json.NewEncoder(w).Encode(edges[0]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Could not write edge response: %v", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
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
	w.Header().Set("Content-Type", contentType)
	if err := json.NewEncoder(w).Encode(miniStakes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Could not write ministakes response: %v", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
}
