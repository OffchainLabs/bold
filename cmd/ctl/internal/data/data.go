package data

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/OffchainLabs/challenge-protocol-v2/api"
)

func LoadEdgesFromDisk() ([]*api.Edge, error) {
	f, err := os.Open("/home/preston/data/edges.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var edges []*api.Edge
	if err := json.Unmarshal(b, &edges); err != nil {
		return nil, err
	}

	return edges, nil
}
