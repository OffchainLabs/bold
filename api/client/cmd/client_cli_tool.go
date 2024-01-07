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
	case "healthz":
		err = client.Healthz()
	case "forcedbupdate":
		err = client.ForceDBUpdate()
	case "ishonestpartyactive":
		err = client.IsHonestPartyActive()
	case "honestpartyhasadvantage":
		if len(args) < 4 {
			panic("Usage: client_cli_tool <url> HonestPartyHasAdvantage <assertion_hash>")
		}
		_, err = client.HonestPartyHasAdvantage(common.HexToHash(args[3]))
	case "ishonestpartyplaying":
		if len(args) < 4 {
			panic("Usage: client_cli_tool <url> IsHonestPartyPlayingSubchallenges <assertion_hash>")
		}
		_, err = client.IsHonestPartyPlayingSubchallenges(common.HexToHash(args[3]))
	case "anyhonestedgeconfirmable":
		if len(args) < 4 {
			panic("Usage: client_cli_tool <url> AnyHonestEdgeConfirmable <assertion_hash>")
		}
		_, err = client.AnyHonestEdgeConfirmable(common.HexToHash(args[3]))
	case "sybilactivityhappening":
		if len(args) < 4 {
			panic("Usage: client_cli_tool <url> SybilActivityHappening <assertion_hash>")
		}
		_, err = client.SybilActivityHappening(common.HexToHash(args[3]))
	case "evilpartyinsights":
		err = client.EvilPartyInsights()
	case "anyeviledgeconfirmed":
		if len(args) < 4 {
			panic("Usage: client_cli_tool <url> EvilPartyInsights <assertion_hash>")
		}
		_, err = client.AnyEvilEdgeConfirmed(common.HexToHash(args[3]))
	case "assertionchainhealth":
		err = client.AssertionChainHealth()
	case "listassertions":
		_, err = client.ListAssertions()
	case "allchallengeedges":
		if len(args) < 4 {
			panic("Usage: client_cli_tool <url> AllChallengeEdges <assertion_hash>")
		}
		_, err = client.AllChallengeEdges(common.HexToHash(args[3]))
	case "challengebyassertionhash":
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
