package validator

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	statemanager "github.com/OffchainLabs/new-rollup-exploration/state-manager"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

func TestChallenges_ValidatorsReachOneStepFork(t *testing.T) {
	// Tests that validators are able to reach a one step fork correctly
	// by playing the challenge game on their own upon observing leaves
	// they disagree with. Here's the example with Alice and Bob.
	//
	// 1. Alice and Bob create their own leaf at height 6 they disagree with
	// 2. Alice bisects to 4, Bob disagrees with the bisection and bisects to his own 4
	// 3. Alice bisects to 2, Bob agrees and merges to 2
	// 4. Alice bisects from 4 to 3, Bob agrees and merges to 3
	// 5. There is now a one-step fork from height 3 to 4
	//
	//
	//
	//                   [4]-[6]-Bob
	//                  /
	// [genesis]-[2]-[3]
	//                  \[4]-[6]-Alice
	//
	//
	//
	//

	// Alice and bob agree up to height 3. From there, their local states diverge.
	hook := test.NewGlobal()
	stateRootsInCommon := make([]common.Hash, 0)
	for i := uint64(0); i <= 3; i++ {
		stateRootsInCommon = append(stateRootsInCommon, util.HashForUint(i))
	}
	aliceRoots := append(
		stateRootsInCommon,
		common.BytesToHash([]byte("a4")),
		common.BytesToHash([]byte("a5")),
		common.BytesToHash([]byte("a6")),
	)
	bobRoots := append(
		stateRootsInCommon,
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

	// TODO: Disable leaf creation for validators, do it manually.
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

	go alice.Start(ctx)
	go bob.Start(ctx)

	leaf, err := alice.submitLeafCreation(ctx)
	require.NoError(t, err)
	_ = leaf
	leaf, err = bob.submitLeafCreation(ctx)
	require.NoError(t, err)
	_ = leaf

	AssertLogsContain(t, hook, "Submitted leaf creation")
	AssertLogsContain(t, hook, "Submitted leaf creation")

	<-ctx.Done()

	//eventsToAssert := []*protocol.ChallengeEvent{}

	//harnessObserver := make(chan protocol.ChallengeEvent)
	//defer close(harnessObserver)
	//chain.SubscribeChallengeEvents(harnessObserver)

	//for {
	//select {
	//case ev := <-harnessObserver:
	//require.Equal(t, true, expectedEvent(eventsToAssert, ev))
	//case <-ctx.Done():
	//break
	//}
	//}
}
