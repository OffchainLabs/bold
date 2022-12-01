package validator

import (
	"context"
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

func TestChallenges_ValidatorsReachOneStepFork_Simple(t *testing.T) {
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
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

	go alice.Start(ctx)
	go bob.Start(ctx)

	_, err = alice.submitLeafCreation(ctx)
	require.NoError(t, err)
	_, err = bob.submitLeafCreation(ctx)
	require.NoError(t, err)

	AssertLogsContain(t, hook, "Submitted leaf creation")
	AssertLogsContain(t, hook, "Submitted leaf creation")

	// TODO: There is unpredictability in who creates the first leaf. We should find
	// a way to make it deterministic so we can assert details of each event in addition
	// to the type of event being fired.
	eventsToAssert := []protocol.ChallengeEvent{
		// Alice adds leaf 6, is presumptive.
		&protocol.ChallengeLeafEvent{},
		// Bob adds leaf 6.
		&protocol.ChallengeLeafEvent{},
		// Alice bisects to 4, is presumptive.
		&protocol.ChallengeBisectEvent{},
		// Bob bisects to 4.
		&protocol.ChallengeBisectEvent{},
		// Bob bisects to 2, is presumptive.
		&protocol.ChallengeBisectEvent{},
		// Alice merges to 2.
		&protocol.ChallengeMergeEvent{},
		// Alice bisects from 4 to 3, is presumptive.
		&protocol.ChallengeBisectEvent{},
		// Bob merges to 3.
		&protocol.ChallengeMergeEvent{},
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
			expectedEventIndex++
		case <-ctx.Done():
			t.Fatal("Timed out - validators were unable to reach one-step-fork in time")
		}
	}
}
