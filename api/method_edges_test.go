package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	"github.com/OffchainLabs/challenge-protocol-v2/challenge-manager/challenge-tree/mock"
)

func TestListEdges(t *testing.T) {
	s, d := NewTestServer(t)

	d.Edges = []protocol.SpecEdge{
		&mock.Edge{
			ID:            "foo",
			EdgeType:      protocol.BlockChallengeEdge,
			StartHeight:   100,
			StartCommit:   "foo_start_commit",
			EndHeight:     150,
			EndCommit:     "foo_end_commit",
			OriginID:      "foo_origin_id",
			ClaimID:       "foo_claim_id",
			LowerChildID:  "foo_lower_child_id",
			UpperChildID:  "foo_upper_child_id",
			CreationBlock: 1,
		},
		&mock.Edge{
			ID:            "bar",
			EdgeType:      protocol.BigStepChallengeEdge,
			StartHeight:   110,
			StartCommit:   "bar_start_commit",
			EndHeight:     160,
			EndCommit:     "bar_end_commit",
			OriginID:      "bar_origin_id",
			ClaimID:       "bar_claim_id",
			LowerChildID:  "bar_lower_child_id",
			UpperChildID:  "bar_upper_child_id",
			CreationBlock: 2,
		},
		&mock.Edge{
			ID:            "baz",
			EdgeType:      protocol.SmallStepChallengeEdge,
			StartHeight:   111,
			StartCommit:   "baz_start_commit",
			EndHeight:     161,
			EndCommit:     "baz_end_commit",
			OriginID:      "baz_origin_id",
			ClaimID:       "baz_claim_id",
			LowerChildID:  "baz_lower_child_id",
			UpperChildID:  "baz_upper_child_id",
			CreationBlock: 5,
		},
	}

	req, err := http.NewRequest("GET", "/edges", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Serve the request with the http recorder.
	s.Router().ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expectedBody := []byte(`[{"id":"0x0000000000000000000000000000000000000000000000000000000000666f6f","type":"block_challenge_edge","startCommitment":{"height":100,"hash":"0x00000000000000000000000000000000666f6f5f73746172745f636f6d6d6974"},"endCommitment":{"height":150,"hash":"0x000000000000000000000000000000000000666f6f5f656e645f636f6d6d6974"},"createdAtBlock":1,"mutualId":"0x6967696e5f69642d3130302d666f6f5f73746172745f636f6d6d69742d313530","originId":"0x00000000000000000000000000000000000000666f6f5f6f726967696e5f6964","claimId":"0x0000000000000000000000000000000000000000666f6f5f636c61696d5f6964","hasChildren":true,"lowerChildId":"0x0000000000000000000000000000666f6f5f6c6f7765725f6368696c645f6964","upperChildId":"0x0000000000000000000000000000666f6f5f75707065725f6368696c645f6964","miniStaker":"0x0000000000000000000000000000000000000000","assertionHash":"0x0000000000000000000000000000000000000000000000000000000000000000","timeUnrivaled":0,"hasRival":false,"status":"pending","hasLengthOneRival":false,"topLevelClaimHeight":{"blockChallengeOriginHeight":0,"bigStepChallengeOriginHeight":0}},{"id":"0x0000000000000000000000000000000000000000000000000000000000626172","type":"big_step_challenge_edge","startCommitment":{"height":110,"hash":"0x000000000000000000000000000000006261725f73746172745f636f6d6d6974"},"endCommitment":{"height":160,"hash":"0x0000000000000000000000000000000000006261725f656e645f636f6d6d6974"},"createdAtBlock":2,"mutualId":"0x6967696e5f69642d3131302d6261725f73746172745f636f6d6d69742d313630","originId":"0x000000000000000000000000000000000000006261725f6f726967696e5f6964","claimId":"0x00000000000000000000000000000000000000006261725f636c61696d5f6964","hasChildren":true,"lowerChildId":"0x00000000000000000000000000006261725f6c6f7765725f6368696c645f6964","upperChildId":"0x00000000000000000000000000006261725f75707065725f6368696c645f6964","miniStaker":"0x0000000000000000000000000000000000000000","assertionHash":"0x0000000000000000000000000000000000000000000000000000000000000000","timeUnrivaled":0,"hasRival":false,"status":"pending","hasLengthOneRival":false,"topLevelClaimHeight":{"blockChallengeOriginHeight":0,"bigStepChallengeOriginHeight":0}},{"id":"0x000000000000000000000000000000000000000000000000000000000062617a","type":"small_step_challenge_edge","startCommitment":{"height":111,"hash":"0x0000000000000000000000000000000062617a5f73746172745f636f6d6d6974"},"endCommitment":{"height":161,"hash":"0x00000000000000000000000000000000000062617a5f656e645f636f6d6d6974"},"createdAtBlock":5,"mutualId":"0x6967696e5f69642d3131312d62617a5f73746172745f636f6d6d69742d313631","originId":"0x0000000000000000000000000000000000000062617a5f6f726967696e5f6964","claimId":"0x000000000000000000000000000000000000000062617a5f636c61696d5f6964","hasChildren":true,"lowerChildId":"0x000000000000000000000000000062617a5f6c6f7765725f6368696c645f6964","upperChildId":"0x000000000000000000000000000062617a5f75707065725f6368696c645f6964","miniStaker":"0x0000000000000000000000000000000000000000","assertionHash":"0x0000000000000000000000000000000000000000000000000000000000000000","timeUnrivaled":0,"hasRival":false,"status":"pending","hasLengthOneRival":false,"topLevelClaimHeight":{"blockChallengeOriginHeight":0,"bigStepChallengeOriginHeight":0}}]`)

	if rr.Body.String() != string(expectedBody) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expectedBody))
	}

}

func TestGetEdge(t *testing.T) {
	s, _ := NewTestServer(t)

	req, err := http.NewRequest("GET", "/edges/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Serve the request with the http recorder.
	s.Router().ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotImplemented)
	}
}
