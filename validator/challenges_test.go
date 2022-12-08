package validator

import (
	"context"
	"crypto/rand"
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

func TestBlockChallenge(t *testing.T) {
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
	t.Run("two validators opening leaves at same height", func(t *testing.T) {
		aliceAddr := common.BytesToAddress([]byte("a"))
		bobAddr := common.BytesToAddress([]byte("b"))
		cfg := &blockChallengeTestConfig{
			numValidators:  2,
			numStateRoots:  6,
			validatorAddrs: []common.Address{aliceAddr, bobAddr},
			validatorNamesByAddress: map[common.Address]string{
				aliceAddr: "alice",
				bobAddr:   "bob",
			},
			// The heights at which the validators diverge in histories. In this test,
			// alice and bob agree up to and including height 3.
			divergenceHeightsByAddress: map[common.Address]uint64{
				aliceAddr: 3,
				bobAddr:   3,
			},
		}
		cfg.eventsToAssert = []protocol.ChallengeEvent{
			// Alice adds a challenge leaf 6, is presumptive.
			&protocol.ChallengeLeafEvent{
				Validator: aliceAddr,
			},
			// Bob adds leaf 6.
			&protocol.ChallengeLeafEvent{
				Validator: bobAddr,
			},
			// Bob bisects to 4, is presumptive.
			&protocol.ChallengeBisectEvent{
				Validator: bobAddr,
			},
			// Alice bisects to 4.
			&protocol.ChallengeBisectEvent{
				Validator: aliceAddr,
			},
			// Alice bisects to 2, is presumptive.
			&protocol.ChallengeBisectEvent{
				Validator: aliceAddr,
			},
			// Bob merges to 2.
			&protocol.ChallengeMergeEvent{
				Validator: bobAddr,
			},
			// Bob bisects from 4 to 3, is presumptive.
			&protocol.ChallengeBisectEvent{
				Validator: bobAddr,
			},
			// Alice merges to 3.
			&protocol.ChallengeMergeEvent{
				Validator: aliceAddr,
			},
			// Both challengers are now at a one-step fork, we now await subchallenge resolution.
		}
		runBlockChallengeTest(t, cfg)
	})
	t.Run("three validators opening leaves at same height", func(t *testing.T) {
		aliceAddr := common.BytesToAddress([]byte("a"))
		bobAddr := common.BytesToAddress([]byte("b"))
		charlieAddr := common.BytesToAddress([]byte("c"))
		cfg := &blockChallengeTestConfig{
			numValidators:  3,
			numStateRoots:  6,
			validatorAddrs: []common.Address{aliceAddr, bobAddr, charlieAddr},
			validatorNamesByAddress: map[common.Address]string{
				aliceAddr:   "alice",
				bobAddr:     "bob",
				charlieAddr: "charlie",
			},
			// The heights at which the validators diverge in histories. In this test,
			// alice and bob agree up to and including height 3.
			divergenceHeightsByAddress: map[common.Address]uint64{
				aliceAddr:   3,
				bobAddr:     3,
				charlieAddr: 3,
			},
		}
		// NOTE: THIS SHOULD FAIL!
		cfg.eventsToAssert = []protocol.ChallengeEvent{
			// Alice adds a challenge leaf 6, is presumptive.
			&protocol.ChallengeLeafEvent{
				Validator: aliceAddr,
			},
			// Bob adds leaf 6.
			&protocol.ChallengeLeafEvent{
				Validator: bobAddr,
			},
			// Bob bisects to 4, is presumptive.
			&protocol.ChallengeBisectEvent{
				Validator: bobAddr,
			},
			// Alice bisects to 4.
			&protocol.ChallengeBisectEvent{
				Validator: aliceAddr,
			},
			// Alice bisects to 2, is presumptive.
			&protocol.ChallengeBisectEvent{
				Validator: aliceAddr,
			},
			// Bob merges to 2.
			&protocol.ChallengeMergeEvent{
				Validator: bobAddr,
			},
			// Bob bisects from 4 to 3, is presumptive.
			&protocol.ChallengeBisectEvent{
				Validator: bobAddr,
			},
			// Alice merges to 3.
			&protocol.ChallengeMergeEvent{
				Validator: aliceAddr,
			},
			// Both challengers are now at a one-step fork, we now await subchallenge resolution.
		}
		runBlockChallengeTest(t, cfg)
	})
	t.Run("two validators opening leaves at very different heights", func(t *testing.T) {
		// TODO: Needs to change the state roots each validator has. Bob might have up to height N
		// while Alice has up to height M, which will cause them to post leaves at different heights.
	})
}

type blockChallengeTestConfig struct {
	// Number of validators we want to enter a block challenge with.
	numValidators uint16
	// Total number of state roots in the chain.
	numStateRoots uint16
	// The heights at which each validator by address diverges histories.
	divergenceHeightsByAddress map[common.Address]uint64
	// Validator human-readable names by address.
	validatorNamesByAddress map[common.Address]string
	// List of validator addresses to initialize in order.
	validatorAddrs []common.Address
	// Events we want to assert are fired from the protocol.
	eventsToAssert []protocol.ChallengeEvent
}

func runBlockChallengeTest(t testing.TB, cfg *blockChallengeTestConfig) {
	hook := test.NewGlobal()
	stateRoots := make([]common.Hash, 0)
	for i := uint64(0); i < uint64(cfg.numStateRoots); i++ {
		stateRoots = append(stateRoots, util.HashForUint(i))
	}

	ctx := context.Background()
	chain := protocol.NewAssertionChain(ctx, util.NewArtificialTimeReference(), time.Second)

	// Increase the balance for each validator in the test.
	bal := big.NewInt(0).Mul(protocol.Gwei, big.NewInt(100))
	err := chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		for addr := range cfg.validatorNamesByAddress {
			chain.AddToBalance(tx, addr, bal)
		}
		return nil
	})
	require.NoError(t, err)

	// Initialize each validator associated state roots which diverge
	// at specified points in the test config.
	validatorStateRoots := make([][]common.Hash, cfg.numValidators)
	for i := uint16(0); i < cfg.numValidators; i++ {
		validatorStateRoots[i] = make([]common.Hash, len(stateRoots))
		for r, rt := range stateRoots {
			var newRoot common.Hash
			copy(newRoot[:], rt[:])
			validatorStateRoots[i][r] = newRoot
		}

		addr := cfg.validatorAddrs[i]
		divergenceHeight := cfg.divergenceHeightsByAddress[addr]
		for h := divergenceHeight; h < uint64(len(validatorStateRoots[i])); h++ {
			divergingRoot := make([]byte, 32)
			_, err = rand.Read(divergingRoot)
			require.NoError(t, err)
			validatorStateRoots[i][h] = common.BytesToHash(divergingRoot)
		}
	}

	// Initialize each validator.
	validators := make([]*Validator, cfg.numValidators)
	for i := 0; i < len(validators); i++ {
		manager := statemanager.New(validatorStateRoots[i])
		addr := cfg.validatorAddrs[i]
		v, err := New(
			ctx,
			chain,
			manager,
			WithName(cfg.validatorNamesByAddress[addr]),
			WithAddress(addr),
			WithDisableLeafCreation(),
		)
		require.NoError(t, err)
		validators[i] = v
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	harnessObserver := make(chan protocol.ChallengeEvent, 1)
	chain.SubscribeChallengeEvents(ctx, harnessObserver)

	// Submit leaves for each validator.
	for _, val := range validators {
		_, err = val.submitLeafCreation(ctx)
		require.NoError(t, err)
		AssertLogsContain(t, hook, "Submitted leaf creation")
	}

	// We fire off each validator's background routines in a specific order.
	for _, val := range validators {
		go val.Start(ctx)
		time.Sleep(time.Millisecond * 100)
	}
	expectedEventIndex := 0
	for ev := range harnessObserver {
		if ctx.Err() != nil {
			t.Fatal("Timed out - validators were unable to reach one-step-fork in time")
		}
		if expectedEventIndex >= len(cfg.eventsToAssert) {
			t.Fatal("Received more events than were expected")
		}
		t.Logf("%+T", ev)
		wantedEv := cfg.eventsToAssert[expectedEventIndex]
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
	}
	require.Equal(t, len(cfg.eventsToAssert), expectedEventIndex)
	t.Log("Finished asserting events")
	for i := 0; i < int(cfg.numValidators); i++ {
		AssertLogsContain(t, hook, "Reached a one-step-fork")
	}
}
