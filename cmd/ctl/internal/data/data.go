package data

import (
	_ "embed"
	"encoding/json"

	"github.com/OffchainLabs/challenge-protocol-v2/api"
)

//go:embed edges.json
var edgesJSON []byte

func LoadEdgesFromDisk() ([]*api.Edge, error) {
	var edges []*api.Edge
	if err := json.Unmarshal(edgesJSON, &edges); err != nil {
		return nil, err
	}

	return edges, nil
}

func CurrentBlockNumber() uint {
	return 103
}
