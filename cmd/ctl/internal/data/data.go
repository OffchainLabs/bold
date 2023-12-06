package data

import (
	_ "embed"
	"encoding/json"

	"github.com/OffchainLabs/bold/api"
)

//go:embed edges.json
var edgesJSON []byte

//go:embed assertions.json
var assertionsJSON []byte

func LoadEdgesFromDisk() ([]*api.Edge, error) {
	var edges []*api.Edge
	if err := json.Unmarshal(edgesJSON, &edges); err != nil {
		return nil, err
	}

	return edges, nil
}

func LoadAssertionsFromDisk() ([]*api.Assertion, error) {
	var assertions []*api.Assertion
	if err := json.Unmarshal(assertionsJSON, &assertions); err != nil {
		return nil, err
	}
	return assertions, nil
}

func CurrentBlockNumber() uint {
	return 103
}
