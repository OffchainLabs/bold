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
			CreationBlock: 1,
		},
		&mock.Edge{
			ID:            "bar",
			CreationBlock: 1,
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
	expectedBody := []byte(`[{"id":"0x0000000000000000000000000000000000000000000000000000000000666f6f","type":"block_challenge_edge","startCommitment":{"height":0,"hash":"0x0000000000000000000000000000000000000000000000000000000000000000"},"endCommitment":{"height":0,"hash":"0x0000000000000000000000000000000000000000000000000000000000000000"},"createdAtBlock":1,"mutualId":"0x00000000000000000000000000000000000000000000000000302d2d302d2d30","originId":"0x0000000000000000000000000000000000000000000000000000000000000000","claimId":"0x0000000000000000000000000000000000000000000000000000000000000000","hasChildren":false,"lowerChildId":"0x0000000000000000000000000000000000000000000000000000000000000000","upperChildId":"0x0000000000000000000000000000000000000000000000000000000000000000","miniStaker":"0x0000000000000000000000000000000000000000","assertionHash":"0x0000000000000000000000000000000000000000000000000000000000000000","timeUnrivaled":0,"hasRival":false,"status":"pending","hasLengthOneRival":false,"topLevelClaimHeight":{"blockChallengeOriginHeight":0,"bigStepChallengeOriginHeight":0}},{"id":"0x0000000000000000000000000000000000000000000000000000000000626172","type":"block_challenge_edge","startCommitment":{"height":0,"hash":"0x0000000000000000000000000000000000000000000000000000000000000000"},"endCommitment":{"height":0,"hash":"0x0000000000000000000000000000000000000000000000000000000000000000"},"createdAtBlock":1,"mutualId":"0x00000000000000000000000000000000000000000000000000302d2d302d2d30","originId":"0x0000000000000000000000000000000000000000000000000000000000000000","claimId":"0x0000000000000000000000000000000000000000000000000000000000000000","hasChildren":false,"lowerChildId":"0x0000000000000000000000000000000000000000000000000000000000000000","upperChildId":"0x0000000000000000000000000000000000000000000000000000000000000000","miniStaker":"0x0000000000000000000000000000000000000000","assertionHash":"0x0000000000000000000000000000000000000000000000000000000000000000","timeUnrivaled":0,"hasRival":false,"status":"pending","hasLengthOneRival":false,"topLevelClaimHeight":{"blockChallengeOriginHeight":0,"bigStepChallengeOriginHeight":0}}]`)

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
