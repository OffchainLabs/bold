package validator

import (
	"testing"
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

	// TODO: Setup a history for both validators such that they disagree at the
	// transition from height 3 to 4. They should agree at the history up to and including
	// the state at height 3, then they will diverge.

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//chain := protocol.NewAssertionChain(ctx, util.NewArtificialTimeReference(), time.Second)
	//manager := &mocks.MockStateManager{}
	//staker1 := common.BytesToAddress([]byte("foo"))
	//staker2 := common.BytesToAddress([]byte("bar"))

	//// TODO: Disable leaf creation for validators, do it manually.
	//v1, err := New(ctx, chain, manager, WithAddress(staker1))
	//require.NoError(t, err)
	//v2, err := New(ctx, chain, manager, WithAddress(staker2))
	//require.NoError(t, err)

	//go v1.Start(ctx)
	//go v2.Start(ctx)

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
