// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package assertions_test

import (
	"context"
	"testing"
	"time"

	"github.com/offchainlabs/bold/assertions"
	protocol "github.com/offchainlabs/bold/chain-abstraction"
	challengemanager "github.com/offchainlabs/bold/challenge-manager"
	"github.com/offchainlabs/bold/challenge-manager/types"
	"github.com/offchainlabs/bold/solgen/go/mocksgen"
	challenge_testing "github.com/offchainlabs/bold/testing"
	statemanager "github.com/offchainlabs/bold/testing/mocks/state-provider"
	"github.com/offchainlabs/bold/testing/setup"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
)

func TestPostAssertion(t *testing.T) {
	ctx := context.Background()
	setup, err := setup.ChainsWithEdgeChallengeManager(
		// setup.WithMockBridge(),
		setup.WithMockOneStepProver(),
		setup.WithChallengeTestingOpts(
			challenge_testing.WithLayerZeroHeights(&protocol.LayerZeroHeights{
				BlockChallengeHeight:     64,
				BigStepChallengeHeight:   32,
				SmallStepChallengeHeight: 32,
			}),
		),
	)
	require.NoError(t, err)

	bridgeBindings, err := mocksgen.NewBridgeStub(setup.Addrs.Bridge, setup.Backend)
	require.NoError(t, err)

	msgCount, err := bridgeBindings.SequencerMessageCount(setup.Chains[0].GetCallOptsWithDesiredRpcHeadBlockNumber(&bind.CallOpts{}))
	require.NoError(t, err)
	require.Equal(t, uint64(1), msgCount.Uint64())

	aliceChain := setup.Chains[0]

	stateManagerOpts := setup.StateManagerOpts
	stateManagerOpts = append(
		stateManagerOpts,
		statemanager.WithNumBatchesRead(5),
	)
	stateManager, err := statemanager.NewForSimpleMachine(stateManagerOpts...)
	require.NoError(t, err)

	chalManager, err := challengemanager.New(
		ctx,
		aliceChain,
		stateManager,
		setup.Addrs.Rollup,
		challengemanager.WithMode(types.DefensiveMode),
	)
	require.NoError(t, err)
	chalManager.Start(ctx)

	preState, err := stateManager.ExecutionStateAfterPreviousState(ctx, 0, nil, 1<<26)
	require.NoError(t, err)
	postState, err := stateManager.ExecutionStateAfterPreviousState(ctx, 1, &preState.GlobalState, 1<<26)
	require.NoError(t, err)

	assertionManager, err := assertions.NewManager(
		aliceChain,
		stateManager,
		setup.Backend,
		chalManager,
		aliceChain.RollupAddress(),
		chalManager.ChallengeManagerAddress(),
		"alice",
		time.Millisecond*200, // poll interval for assertions
		time.Hour,            // confirmation attempt interval
		stateManager,
		time.Millisecond*100, // poll interval
		time.Second*1,
		nil,
		assertions.WithDangerousReadyToPost(),
		assertions.WithPostingDisabled(),
	)
	require.NoError(t, err)

	go assertionManager.Start(ctx)

	time.Sleep(time.Second)

	posted, err := assertionManager.PostAssertion(ctx)
	require.NoError(t, err)
	require.Equal(t, true, posted.IsSome())
	require.Equal(t, postState, protocol.GoExecutionStateFromSolidity(posted.Unwrap().AfterState))
}
