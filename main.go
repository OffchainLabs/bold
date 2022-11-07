package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	statemanager "github.com/OffchainLabs/new-rollup-exploration/state-manager"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/OffchainLabs/new-rollup-exploration/validator"
)

const challengePeriod = 100 * time.Second

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	timeRef := util.NewRealTimeReference()
	chain := protocol.NewAssertionChain(ctx, timeRef, challengePeriod)
	manager := statemanager.NewSimulatedManager(
		statemanager.WithL2BlockTimes(5 * time.Second),
	)

	// We start our simulation with a single, honest validator.
	val, err := validator.New(
		ctx,
		chain,
		manager,
		validator.WithMaliciousProbability(0), // Not a malicious validator for now...
	)
	if err != nil {
		panic(err)
	}

	// Begin the validator process in the background.
	// The validator will not only be responsible for listening to new leaf and challenge creation events,
	// but will also participate in defending challenges it agrees with and challenge assertions it
	// disagrees with. Honest validators will also be responsible for creating new leaves in the assertion
	// tree based on the local state advancing, controlled by the state manager.
	//
	// All honest validators will have the same Merkle root as the state transition advances the chain,
	// and therefore all leaves posted by honest validators should have the same commitment. We give all honest
	// validators the same state reader. State reader is also advancing its state in the background:
	//
	// TODO: For the purposes of simulation, we plan to create several validators that can either be malicious, honest,
	// or chaos monkeys with some probability and we want to observe the behavior of the system.
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
	go val.Validate(ctx)

	// Simulate advancing an L2 state via state transitions using our state manager. This can be configured
	// to advance at different rates, and validators that use this state manager will be notified
	// of when there is a new state created. This will trigger honest validators to submit leaf creation events
	// to the on-chain protocol.
	go manager.AdvanceL2State(ctx)

	// Await a shutdown signal, which will trigger context cancellation across the program.
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
	cancel()
}
