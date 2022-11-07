package main

import (
	"context"
	"time"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	statemanager "github.com/OffchainLabs/new-rollup-exploration/state-manager"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/OffchainLabs/new-rollup-exploration/validator"
)

const challengePeriod = 100 * time.Second

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeRef := util.NewRealTimeReference()
	chain := protocol.NewAssertionChain(ctx, timeRef, challengePeriod)
	_ = chain
	manager := statemanager.NewSimulatedManager()

	val, err := validator.New(
		ctx,
		chain,
		manager,
		validator.WithMaliciousProbability(0),
	)
	if err != nil {
		panic(err)
	}
	_ = val

	// go chain.Start(ctx)
	// create N validators, for each, run in background goroutines
	// for v := range validator { go v.Validate(ctx) }
	//
	// In the background, make validators create leaves for a chain that is "advancing". State transitions
	// are happening and there are new Merkle commitments that must be posted on-chain and await a challenge game.
	// Observe behavior using real time-references and simulate different probabilities of malicious / honest
	// validators. All honest validators will have the same Merkle root as the state transition advances the chain,
	// and therefore all leaves posted by honest validators should have the same commitment. We give all honest
	// validators the same state reader. State reader is also advancing its state in the background:
	//
	// Advances a chain in the background, simulating state transitions running and block heights.
	// go stateManager.Advance(ctx)
	//
	// TODO: Create either a metrics collector that will gather information about the challenge games being
	// played and create an API that can extract a graphviz of the current assertion chain to visualize
	// the actual tree of assertions. Bonus: add detail such as presumtive status, chess clocks, etc.
	//
	// Observe:
	//  1. All honest validators
	//  2. Malicious validators issuing challenges on honestly-created leaves
	//  3. Honest validators issuing challenges on maliciously-created leaves
	//  4. Chaos monkey validators operating alongside honest ones.
	//
	// Await a graceful shutdown signal...
}
