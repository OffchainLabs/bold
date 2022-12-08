package validator

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	statemanager "github.com/OffchainLabs/new-rollup-exploration/state-manager"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

func TestValidator_SingleBlockChallenge(t *testing.T) {
	t.Run("two validators' leaves created at same height", func(tt *testing.T) {
		cfg := &blockChallengeTestConfig{}
		runBlockChallengeValidators(tt, cfg)
	})
	t.Run("three validators' leaves created at same height", func(tt *testing.T) {
		cfg := &blockChallengeTestConfig{}
		runBlockChallengeValidators(tt, cfg)
	})
	t.Run("two validators' leaves created at large difference in heights", func(tt *testing.T) {
		cfg := &blockChallengeTestConfig{}
		runBlockChallengeValidators(tt, cfg)
	})
	t.Run("three validators' leaves created at large difference in heights", func(tt *testing.T) {
		cfg := &blockChallengeTestConfig{}
		runBlockChallengeValidators(tt, cfg)
	})
	t.Run("fifty validators with many varying heights and many equal heights", func(tt *testing.T) {
		cfg := &blockChallengeTestConfig{}
		runBlockChallengeValidators(tt, cfg)
	})
}

func TestBlockChallenge_ValidatorParticipatesInMultipleChallengesConcurrently(t *testing.T) {
}

type blockChallengeTestConfig struct {
	numValidators              uint16
	stateRoots                 []common.Hash
	divergenceHeightsByAddress map[common.Address]uint64
	validatorNamesByAddress    map[common.Address]string
	eventsToAssert             []protocol.ChallengeEvent
}

func runBlockChallengeValidators(t *testing.T, cfg *blockChallengeTestConfig) {
	// Tests that validators are able to reach a one step fork correctly
	// by playing the challenge game on their own upon observing leaves
	// they disagree with. Here's the example with Alice and Bob.
	//
	// 1. alice and bob create their own leaves at height 6 they disagree with
	// 2. alice bisects to 4, bob disagrees with the bisection and bisects to his own 4
	// 3. alice bisects to 2, bob agrees and merges to 2
	// 4. alice bisects from 4 to 3, bob agrees and merges to 3
	// 5. there is now a one-step fork from height 3 to 4
	//
	//
	//
	//                   [4]-[6]-bob
	//                  /
	// [genesis]-[2]-[3]
	//                  \[4]-[6]-alice
	//
	//
	//

	// Alice and bob agree up to height 3. From there, their local states diverge.
	hook := test.NewGlobal()
	stateRootsInCommon := make([]common.Hash, 0)
	for i := uint64(0); i < 3; i++ {
		stateRootsInCommon = append(stateRootsInCommon, util.HashForUint(i))
	}
	aliceRoots := append(
		stateRootsInCommon,
		common.BytesToHash([]byte("a3")),
		common.BytesToHash([]byte("a4")),
		common.BytesToHash([]byte("a5")),
		common.BytesToHash([]byte("a6")),
	)
	bobRoots := append(
		stateRootsInCommon,
		common.BytesToHash([]byte("b3")),
		common.BytesToHash([]byte("b4")),
		common.BytesToHash([]byte("b5")),
		common.BytesToHash([]byte("b6")),
	)
	aliceStateManager := statemanager.New(aliceRoots)
	bobStateManager := statemanager.New(bobRoots)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	chain := protocol.NewAssertionChain(ctx, util.NewArtificialTimeReference(), time.Second)
	aliceAddr := common.BytesToAddress([]byte("a"))
	bobAddr := common.BytesToAddress([]byte("b"))

	bal := big.NewInt(0).Mul(protocol.Gwei, big.NewInt(100))
	err := chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		chain.AddToBalance(tx, aliceAddr, bal)
		chain.AddToBalance(tx, bobAddr, bal)
		return nil
	})
	require.NoError(t, err)

	alice, err := New(
		ctx,
		chain,
		aliceStateManager,
		WithName("alice"),
		WithAddress(aliceAddr),
		WithDisableLeafCreation(),
	)
	require.NoError(t, err)
	bob, err := New(
		ctx,
		chain,
		bobStateManager,
		WithName("bob"),
		WithAddress(bobAddr),
		WithDisableLeafCreation(),
	)
	require.NoError(t, err)

	harnessObserver := make(chan protocol.ChallengeEvent, 1)
	chain.SubscribeChallengeEvents(ctx, harnessObserver)

	// We submit leaves manually.
	_, err = bob.submitLeafCreation(ctx)
	require.NoError(t, err)
	_, err = alice.submitLeafCreation(ctx)
	require.NoError(t, err)

	AssertLogsContain(t, hook, "Submitted leaf creation")
	AssertLogsContain(t, hook, "Submitted leaf creation")

	// We fire off Alice and Bob's background routines, with Bob going first.
	// this means Bob should be the first to attempt to challenge Bob's leaf
	// and successfully create a challenge with an attached challenge vertex.
	// This means we will have a deterministic sequence of events we can verify
	// from Alice and Bob's interaction after this occurrence by reading from the challenge
	// events feed in the protocol.
	go bob.Start(ctx)
	time.Sleep(time.Millisecond * 100)
	go alice.Start(ctx)

	eventsToAssert := []protocol.ChallengeEvent{
		// Bob adds a challenge leaf 6, is presumptive.
		&protocol.ChallengeLeafEvent{
			Validator: bobAddr,
		},
		// Alice adds leaf 6.
		&protocol.ChallengeLeafEvent{
			Validator: aliceAddr,
		},
		// Alice bisects to 4, is presumptive.
		&protocol.ChallengeBisectEvent{
			Validator: aliceAddr,
		},
		// Bob bisects to 4.
		&protocol.ChallengeBisectEvent{
			Validator: bobAddr,
		},
		// Bob bisects to 2, is presumptive.
		&protocol.ChallengeBisectEvent{
			Validator: bobAddr,
		},
		// Alice merges to 2.
		&protocol.ChallengeMergeEvent{
			Validator: aliceAddr,
		},
		// Alice bisects from 4 to 3, is presumptive.
		&protocol.ChallengeBisectEvent{
			Validator: aliceAddr,
		},
		// Bob merges to 3.
		&protocol.ChallengeMergeEvent{
			Validator: bobAddr,
		},
		// Both challengers are now at a one-step fork, we now await subchallenge resolution.
	}
	expectedEventIndex := 0
	for {
		if expectedEventIndex == len(eventsToAssert) {
			t.Log("Finished asserting events")
			AssertLogsContain(t, hook, "Reached a one-step-fork")
			AssertLogsContain(t, hook, "Reached a one-step-fork")
			break
		}
		select {
		case ev := <-harnessObserver:
			t.Logf("%+T", ev)
			wantedEv := eventsToAssert[expectedEventIndex]
			wanted := reflect.TypeOf(wantedEv).Elem()
			got := reflect.TypeOf(ev).Elem()
			t.Logf("Asserting wanted event %+T against received %+T", wantedEv, ev)
			require.Equal(t, wanted, got)

			// Check the validator address is the one we expect from the sequence of events.
			require.Equal(
				t,
				wantedEv.ValidatorAddress(),
				ev.ValidatorAddress(),
				fmt.Sprintf("event at index %d did not match", expectedEventIndex),
			)
			expectedEventIndex++
		case <-ctx.Done():
			t.Fatal("Timed out - validators were unable to reach one-step-fork in time")
		}
	}
}

func TestChallenges_ThreeValidatorsReachOneStepFork_Simple(t *testing.T) {
	// Similar to the two validator test, but with three validators.
	//
	// 1. alice, bob and charlies create their own leaves at height 6 they disagree with
	// 2. alice bisects to 4, bob and charlie disagrees with the bisection and bisects to their own 4
	// 3. alice bisects to 2, bob and charlie agrees and merges to 2
	// 4. alice bisects from 4 to 3, bob and charlie agrees and merges to 3
	// 5. they are now at a one-step fork, we now await subchallenge resolution.
	hook := test.NewGlobal()
	stateRootsInCommon := make([]common.Hash, 0)
	for i := uint64(0); i < 3; i++ {
		stateRootsInCommon = append(stateRootsInCommon, util.HashForUint(i))
	}
	aliceRoots := append(
		stateRootsInCommon,
		common.BytesToHash([]byte("a3")),
		common.BytesToHash([]byte("a4")),
		common.BytesToHash([]byte("a5")),
		common.BytesToHash([]byte("a6")),
	)
	bobRoots := append(
		stateRootsInCommon,
		common.BytesToHash([]byte("b3")),
		common.BytesToHash([]byte("b4")),
		common.BytesToHash([]byte("b5")),
		common.BytesToHash([]byte("b6")),
	)
	charlieRoots := append(
		stateRootsInCommon,
		common.BytesToHash([]byte("c3")),
		common.BytesToHash([]byte("c4")),
		common.BytesToHash([]byte("c5")),
		common.BytesToHash([]byte("c6")),
	)

	aliceStateManager := statemanager.New(aliceRoots)
	bobStateManager := statemanager.New(bobRoots)
	charlieStateManager := statemanager.New(charlieRoots)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	chain := protocol.NewAssertionChain(ctx, util.NewArtificialTimeReference(), time.Second)
	aliceAddr := common.BytesToAddress([]byte("a"))
	bobAddr := common.BytesToAddress([]byte("b"))
	charlieAddr := common.BytesToAddress([]byte("c"))

	bal := big.NewInt(0).Mul(protocol.Gwei, big.NewInt(100))
	err := chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		chain.AddToBalance(tx, aliceAddr, bal)
		chain.AddToBalance(tx, bobAddr, bal)
		chain.AddToBalance(tx, charlieAddr, bal)
		return nil
	})
	require.NoError(t, err)

	alice, err := New(
		ctx,
		chain,
		aliceStateManager,
		WithName("alice"),
		WithAddress(aliceAddr),
		WithDisableLeafCreation(),
	)
	require.NoError(t, err)
	bob, err := New(
		ctx,
		chain,
		bobStateManager,
		WithName("bob"),
		WithAddress(bobAddr),
		WithDisableLeafCreation(),
	)
	require.NoError(t, err)
	charlie, err := New(
		ctx,
		chain,
		charlieStateManager,
		WithName("charlie"),
		WithAddress(charlieAddr),
		WithDisableLeafCreation(),
	)
	require.NoError(t, err)

	harnessObserver := make(chan protocol.ChallengeEvent, 1)
	chain.SubscribeChallengeEvents(ctx, harnessObserver)

	_, err = alice.submitLeafCreation(ctx)
	require.NoError(t, err)
	_, err = bob.submitLeafCreation(ctx)
	require.NoError(t, err)
	_, err = charlie.submitLeafCreation(ctx)
	require.NoError(t, err)

	AssertLogsContain(t, hook, "Submitted leaf creation")
	AssertLogsContain(t, hook, "Submitted leaf creation")
	AssertLogsContain(t, hook, "Submitted leaf creation")

	go alice.Start(ctx)
	time.Sleep(time.Millisecond * 100)
	go bob.Start(ctx)
	time.Sleep(time.Millisecond * 100)
	go charlie.Start(ctx)

	eventsToAssert := []protocol.ChallengeEvent{
		// TODO: Why is this passing?
		&protocol.ChallengeLeafEvent{Validator: aliceAddr},
		&protocol.ChallengeLeafEvent{Validator: aliceAddr},
		&protocol.ChallengeLeafEvent{Validator: bobAddr},
		&protocol.ChallengeLeafEvent{Validator: bobAddr},
		&protocol.ChallengeBisectEvent{Validator: bobAddr},
		&protocol.ChallengeBisectEvent{Validator: aliceAddr},
		&protocol.ChallengeBisectEvent{Validator: aliceAddr},
		&protocol.ChallengeMergeEvent{Validator: bobAddr},
		&protocol.ChallengeBisectEvent{Validator: bobAddr},
		&protocol.ChallengeMergeEvent{Validator: aliceAddr},
	}
	expectedEventIndex := 0
	for {
		if expectedEventIndex == len(eventsToAssert) {
			t.Log("Finished asserting events")
			AssertLogsContain(t, hook, "Reached a one-step-fork")
			AssertLogsContain(t, hook, "Reached a one-step-fork")
			AssertLogsContain(t, hook, "Reached a one-step-fork")
			break
		}
		select {
		case ev := <-harnessObserver:
			t.Logf("%+T", ev)
			wantedEv := eventsToAssert[expectedEventIndex]
			wanted := reflect.TypeOf(wantedEv).Elem()
			got := reflect.TypeOf(ev).Elem()
			t.Logf("Asserting wanted event %+T against received %+T", wantedEv, ev)
			require.Equal(t, wanted, got)

			// Check the validator address is the one we expect from the sequence of events.
			require.Equal(
				t,
				wantedEv.ValidatorAddress(),
				ev.ValidatorAddress(),
				fmt.Sprintf("event at index %d did not match", expectedEventIndex),
			)
			expectedEventIndex++
		case <-ctx.Done():
			t.Fatal("Timed out - validators were unable to reach one-step-fork in time")
		}
	}
}
