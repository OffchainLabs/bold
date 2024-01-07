package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/OffchainLabs/bold/api/client"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		clientType := reflect.TypeOf(&client.Client{})
		var methods string
		for i := 0; i < clientType.NumMethod(); i++ {
			method := clientType.Method(i)
			methods += fmt.Sprintf("%s\n", method.Name)
		}
		panic("Usage: client_cli_tool <url> <method_name>\n" + "Available methods:\n" + methods)
	}
	client := client.NewClient(args[1])
	var err error
	switch strings.ToLower(args[2]) {
	case "Healthz":
		err = client.Healthz()
	case "ForceDBUpdate":
		err = client.ForceDBUpdate()
	case "IsHonestPartyActive":
		err = client.IsHonestPartyActive()
	case "HonestPartyHasAdvantage":
		if len(args) < 4 {
			panic("Usage: client_cli_tool <url> HonestPartyHasAdvantage <assertion_hash>")
		}
		_, err = client.HonestPartyHasAdvantage(common.HexToHash(args[3]))
	case "IsHonestPartyPlayingSubchallenges":
		if len(args) < 4 {
			panic("Usage: client_cli_tool <url> IsHonestPartyPlayingSubchallenges <assertion_hash>")
		}
		_, err = client.IsHonestPartyPlayingSubchallenges(common.HexToHash(args[3]))
	case "AnyHonestEdgeConfirmable":
		if len(args) < 4 {
			panic("Usage: client_cli_tool <url> AnyHonestEdgeConfirmable <assertion_hash>")
		}
		_, err = client.AnyHonestEdgeConfirmable(common.HexToHash(args[3]))
	case "SybilActivityHappening":
		if len(args) < 4 {
			panic("Usage: client_cli_tool <url> SybilActivityHappening <assertion_hash>")
		}
		_, err = client.SybilActivityHappening(common.HexToHash(args[3]))
	case "EvilPartyInsights":
		err = client.EvilPartyInsights()
	case "AnyEvilEdgeConfirmed":
		if len(args) < 4 {
			panic("Usage: client_cli_tool <url> EvilPartyInsights <assertion_hash>")
		}
		_, err = client.AnyEvilEdgeConfirmed(common.HexToHash(args[3]))
	case "AssertionChainHealth":
		err = client.AssertionChainHealth()
	case "ListAssertions":
		_, err = client.ListAssertions()
	case "AllChallengeEdges":
		if len(args) < 4 {
			panic("Usage: client_cli_tool <url> AllChallengeEdges <assertion_hash>")
		}
		_, err = client.AllChallengeEdges(common.HexToHash(args[3]))
	case "ChallengeByAssertionHash":
		if len(args) < 4 {
			panic("Usage: client_cli_tool <url> ChallengeByAssertionHash <assertion_hash>")
		}
		_, err = client.ChallengeByAssertionHash(common.HexToHash(args[3]))

	default:
		panic(fmt.Sprintf("Unknown method: %s", args[2]))
	}
	if err != nil {
		panic(err)
	}
}
