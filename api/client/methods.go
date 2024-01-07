package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/OffchainLabs/bold/api/server"
	protocol "github.com/OffchainLabs/bold/chain-abstraction"

	"github.com/ethereum/go-ethereum/common"
)

// Healthz Checks if the server is healthy.
func (s *Client) Healthz() error {
	_, err := s.httpGet(server.HealthzPath)
	if err != nil {
		fmt.Printf("Server is not healthy: %s\n", err)
	} else {
		fmt.Printf("Server is healthy\n")
	}
	return err
}

// ForceDBUpdate Forces the server to update its database.
func (s *Client) ForceDBUpdate() error {
	res, err := http.Post(s.url+server.ForceDBUpdatePath, "application/json", nil)
	if err != nil {
		fmt.Printf("Error while forcing DB update: %s\n", err)
		return err
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Unable to force DB update: %s\n", http.StatusText(res.StatusCode))
		return fmt.Errorf("HTTP error with status %d returned by server: %s", res.StatusCode, http.StatusText(res.StatusCode))
	}

	fmt.Printf("DB update forced\n")
	return nil
}

func (s *Client) IsHonestPartyActive() error {
	// TODO: implement
	return nil
}

// HonestPartyHasAdvantage Checks if honest party has an advantage over evil party.
// 1.) If there are no unrivaled evil edges.
// 2.) And every evil edge is rivaled by an honest edge with a higher cumulative path timer.
func (s *Client) HonestPartyHasAdvantage(assertionHash common.Hash) (bool, error) {
	edges, err := s.getEdges(assertionHash)
	if err != nil {
		return false, err
	}
	honestHaveAdvantage :=
		noUnrivaledEvilEdge(edges) &&
			evilEdgesHaveLowerCumulativePathTimer(edges)
	if honestHaveAdvantage {
		fmt.Printf("Honests have advantage\n")
	}
	return honestHaveAdvantage, nil
}

// IsHonestPartyPlayingSubchallenges Checks if honest party is playing subchallenges (i.e., if there is at least one honest edge).
func (s *Client) IsHonestPartyPlayingSubchallenges(assertionHash common.Hash) (bool, error) {
	edges, err := s.getEdges(assertionHash)
	if err != nil {
		return false, err
	}
	numHonestEdges := 0
	for _, edge := range edges {
		if edge.IsHonest {
			numHonestEdges++
		}
	}
	fmt.Printf("Number of edges: %d\n", len(edges))
	fmt.Printf("Number of honest edges: %d\n", numHonestEdges)
	return numHonestEdges > 0, nil
}

// AnyHonestEdgeConfirmable Checks if any honest edge is confirmable via children, claim, timer, or OSP.
func (s *Client) AnyHonestEdgeConfirmable(assertionHash common.Hash) (bool, error) {
	edges, err := s.getEdges(assertionHash)
	if err != nil {
		return false, err
	}

	anyHonestEdgeConfirmable := false
	for _, edge := range edges {
		if edge.IsHonest {
			if edge.IsConfirmable {
				fmt.Printf("Honest edge %s is %s\n", edge.Id, edge.ConfirmableBy)
				anyHonestEdgeConfirmable = true
			}
		}
	}
	if !anyHonestEdgeConfirmable {
		fmt.Printf("No Honest edge is confirmable\n")
	}
	return anyHonestEdgeConfirmable, nil
}

// SybilActivityHappening Checks if sybil activity is happening, i.e., if there are more than 2 ministakes per challenge level.
func (s *Client) SybilActivityHappening(assertionHash common.Hash) (bool, error) {
	body, err := s.httpGet(strings.Replace(server.MiniStakesPath, server.AssertionHash, assertionHash.String(), 1))
	if err != nil {
		fmt.Printf("Error while fetching ministakes: %s\n", err)
		return false, err
	}
	var ministakesList []*server.JsonMiniStakes
	err = json.Unmarshal(body, &ministakesList)
	if err != nil {
		fmt.Printf("Error while parsing ministakes: %s\n", err)
		return false, err
	}

	sybilActivityHappening := false
	for _, ministakes := range ministakesList {
		if ministakes.StakeInfo.NumberOfMinistakes > 2 {
			fmt.Printf("Sybil activity happening at level %s\n", ministakes.Level)
			fmt.Printf("Number of ministakes: %d\n", ministakes.StakeInfo.NumberOfMinistakes)
			sybilActivityHappening = true
		}
	}
	return sybilActivityHappening, nil
}

func (s *Client) EvilPartyInsights() error {
	// TODO: implement
	return nil
}

// AnyEvilEdgeConfirmed Checks if any evil edges have been confirmed.
func (s *Client) AnyEvilEdgeConfirmed(assertionHash common.Hash) (bool, error) {
	edges, err := s.getEdges(assertionHash)
	if err != nil {
		return false, err
	}

	anyEvilEdgeConfirmed := false
	for _, edge := range edges {
		if !edge.IsHonest {
			if edge.Status == protocol.EdgeConfirmed.String() {
				fmt.Printf("Evil edge %s is confirmed\n", edge.Id)
				anyEvilEdgeConfirmed = true
			}
		}
	}
	if !anyEvilEdgeConfirmed {
		fmt.Printf("No evil edge is confirmed\n")
	}
	return anyEvilEdgeConfirmed, nil
}

func (s *Client) AssertionChainHealth() error {
	// TODO: implement
	return nil
}

// ListAssertions fetches all assertions from the server.
func (s *Client) ListAssertions() ([]*server.JsonAssertion, error) {
	body, err := s.httpGet(server.ListAssertionsPath)
	if err != nil {
		fmt.Printf("Error while fetching assertions: %s\n", err)
		return nil, err
	}
	var assertions []*server.JsonAssertion
	err = json.Unmarshal(body, &assertions)
	if err != nil {
		fmt.Printf("Error while parsing assertions: %s\n", err)
		return nil, err
	}
	fmt.Printf("List of assertions: %s\n", string(body))
	return assertions, nil
}

// AllChallengeEdges fetches all the edges corresponding to a challenge
func (s *Client) AllChallengeEdges(assertionHash common.Hash) ([]*server.JsonEdge, error) {
	body, err := s.httpGet(strings.Replace(server.AllChallengeEdgesPath, server.AssertionHash, assertionHash.String(), 1))
	if err != nil {
		fmt.Printf("Error while fetching challenge edges: %s\n", err)
		return nil, err
	}
	var edges []*server.JsonEdge
	err = json.Unmarshal(body, &edges)
	if err != nil {
		fmt.Printf("Error while parsing challenge edges: %s\n", err)
		return nil, err
	}
	fmt.Printf("All challenge edges: %s\n", string(body))
	return edges, nil
}

// ChallengeByAssertionHash fetches information about a challenge on a specific assertion hash
func (s *Client) ChallengeByAssertionHash(assertionHash common.Hash) (*server.JsonChallenge, error) {
	body, err := s.httpGet(strings.Replace(server.ChallengeByAssertionHashPath, server.AssertionHash, assertionHash.String(), 1))
	if err != nil {
		fmt.Printf("Error while fetching challenge: %s\n", err)
		return nil, err
	}
	var challenge server.JsonChallenge
	err = json.Unmarshal(body, &challenge)
	if err != nil {
		fmt.Printf("Error while parsing challenge: %s\n", err)
		return nil, err
	}
	fmt.Printf("Challenge: %s\n", string(body))
	return &challenge, nil
}

// MiniStakes fetches all ministakes for a specific assertion hash
func (s *Client) MiniStakes(assertionHash common.Hash) ([]*server.JsonMiniStakes, error) {
	body, err := s.httpGet(strings.Replace(server.MiniStakesPath, server.AssertionHash, assertionHash.String(), 1))
	if err != nil {
		fmt.Printf("Error while fetching ministakes: %s\n", err)
		return nil, err
	}
	var ministakesList []*server.JsonMiniStakes
	err = json.Unmarshal(body, &ministakesList)
	if err != nil {
		fmt.Printf("Error while parsing ministakes: %s\n", err)
		return nil, err
	}
	fmt.Printf("Ministakes: %s\n", string(body))
	return ministakesList, nil
}

func (s *Client) httpGet(path string) ([]byte, error) {
	res, err := http.Get(s.url + path)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error with status %d returned by server: %s", res.StatusCode, http.StatusText(res.StatusCode))
	}

	return io.ReadAll(res.Body)
}

func (s *Client) getEdges(assertionHash common.Hash) ([]*server.JsonEdge, error) {
	body, err := s.httpGet(strings.Replace(server.AllChallengeEdgesPath, server.AssertionHash, assertionHash.String(), 1))
	if err != nil {
		fmt.Printf("Error while fetching challenge edges: %s\n", err)
		return nil, err
	}
	var edges []*server.JsonEdge
	err = json.Unmarshal(body, &edges)
	if err != nil {
		fmt.Printf("Error while parsing challenge edges: %s\n", err)
		return nil, err
	}
	return edges, nil
}

func noUnrivaledEvilEdge(edges []*server.JsonEdge) bool {
	noUnrivaledEvilEdge := true
	for _, edge := range edges {
		if !edge.HasRival {
			if !edge.IsHonest {
				fmt.Printf("Evil edge %s is not rivaled\n", edge.Id)
				noUnrivaledEvilEdge = false
			}
		}
	}
	return noUnrivaledEvilEdge
}

func evilEdgesHaveLowerCumulativePathTimer(edges []*server.JsonEdge) bool {
	evilEdgesHaveLowerCumulativePathTimer := true
	honestEdgesMap, evilEdgesMap := getHonestEvilEdgeMap(edges)

	for _, evilEdges := range evilEdgesMap {
		for _, evilEdge := range evilEdges {
			if honestEdgesMap[evilEdge.MutualId] == nil {
				fmt.Printf("Evil edge %s is rivaled by no honest edge\n", evilEdge.Id)
				evilEdgesHaveLowerCumulativePathTimer = false
			} else {
				if evilEdge.CumulativePathTimer > honestEdgesMap[evilEdge.MutualId][0].CumulativePathTimer {
					fmt.Printf(
						"Evil edge %s"+
							" with cumulative path timer %d"+
							" is rivaled by honest edge %s"+
							" with cumulative path timer %d\n",
						evilEdge.Id,
						evilEdge.CumulativePathTimer,
						honestEdgesMap[evilEdge.MutualId][0].Id,
						honestEdgesMap[evilEdge.MutualId][0].CumulativePathTimer,
					)
					evilEdgesHaveLowerCumulativePathTimer = false
				}
			}
		}
	}
	return evilEdgesHaveLowerCumulativePathTimer
}

func getHonestEvilEdgeMap(edges []*server.JsonEdge) (map[common.Hash][]*server.JsonEdge, map[common.Hash][]*server.JsonEdge) {
	honestEdgesMap := make(map[common.Hash][]*server.JsonEdge)
	evilEdgesMap := make(map[common.Hash][]*server.JsonEdge)
	for _, edge := range edges {
		if edge.IsHonest {
			if honestEdgesMap[edge.MutualId] == nil {
				honestEdgesMap[edge.MutualId] = []*server.JsonEdge{}
			}
			honestEdgesMap[edge.MutualId] = append(honestEdgesMap[edge.MutualId], edge)

		} else {
			if evilEdgesMap[edge.MutualId] == nil {
				evilEdgesMap[edge.MutualId] = []*server.JsonEdge{}
			}
			evilEdgesMap[edge.MutualId] = append(evilEdgesMap[edge.MutualId], edge)
		}
	}
	return honestEdgesMap, evilEdgesMap
}
