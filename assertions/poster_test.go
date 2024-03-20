// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package assertions

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/containers/option"
	"github.com/OffchainLabs/bold/containers/threadsafe"
	"github.com/OffchainLabs/bold/solgen/go/rollupgen"
	"github.com/OffchainLabs/bold/testing/mocks"
	"github.com/OffchainLabs/bold/testing/setup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestPostAssertion(t *testing.T) {
	t.Run("new stake", func(t *testing.T) {
		ctx := context.Background()
		poster, chain, stateManager := setupPoster(t)
		_, creationInfo := setupAssertions(ctx, chain, stateManager, 10, func(int) bool { return false })
		info := creationInfo[len(creationInfo)-1]
		latestConfirmed, err := chain.LatestConfirmed(ctx)
		require.NoError(t, err)
		poster.assertionChainData.latestAgreedAssertion = latestConfirmed.Id()
		latestConfirmedInfo, err := chain.ReadAssertionCreationInfo(ctx, latestConfirmed.Id())
		require.NoError(t, err)
		poster.assertionChainData.canonicalAssertions[latestConfirmed.Id()] = latestConfirmedInfo

		execState := protocol.GoExecutionStateFromSolidity(info.AfterState)
		stateManager.On("AgreesWithExecutionState", ctx, execState).Return(nil)
		assertion := &mocks.MockAssertion{}

		latestValid := info

		chain.On(
			"ReadAssertionCreationInfo",
			ctx,
			latestValid,
		).Return(info, nil)
		chain.On("IsStaked", ctx).Return(false, nil)
		stateManager.On("ExecutionStateAfterBatchCount", ctx, info.InboxMaxCount.Uint64()).Return(execState, nil)

		chain.On("NewStakeOnNewAssertion", ctx, info, execState).Return(assertion, nil)
		posted, err := poster.PostAssertion(ctx)
		require.NoError(t, err)
		require.Equal(t, assertion, posted.Unwrap())
	})
	t.Run("existing stake", func(t *testing.T) {
		ctx := context.Background()
		poster, chain, stateManager := setupPoster(t)
		_, creationInfo := setupAssertions(ctx, chain, stateManager, 10, func(int) bool { return false })
		info := creationInfo[len(creationInfo)-1]
		latestConfirmed, err := chain.LatestConfirmed(ctx)
		require.NoError(t, err)
		poster.assertionChainData.latestAgreedAssertion = latestConfirmed.Id()
		latestConfirmedInfo, err := chain.ReadAssertionCreationInfo(ctx, latestConfirmed.Id())
		require.NoError(t, err)
		poster.assertionChainData.canonicalAssertions[latestConfirmed.Id()] = latestConfirmedInfo

		execState := protocol.GoExecutionStateFromSolidity(info.AfterState)
		stateManager.On("AgreesWithExecutionState", ctx, execState).Return(nil)
		assertion := &mocks.MockAssertion{}

		latestValid := creationInfo[len(creationInfo)-1]

		chain.On(
			"ReadAssertionCreationInfo",
			ctx,
			latestValid,
		).Return(info, nil)
		chain.On("IsStaked", ctx).Return(true, nil)

		stateManager.On("ExecutionStateAfterBatchCount", ctx, info.InboxMaxCount.Uint64()).Return(execState, nil)

		chain.On("StakeOnNewAssertion", ctx, info, execState).Return(assertion, nil)
		posted, err := poster.PostAssertion(ctx)
		require.NoError(t, err)
		require.Equal(t, assertion, posted.Unwrap())
	})
}

func setupAssertions(
	ctx context.Context,
	p *mocks.MockProtocol,
	s *mocks.MockStateManager,
	num int,
	validity func(int) bool,
) ([]protocol.Assertion, []*protocol.AssertionCreatedInfo) {
	if num == 0 {
		return make([]protocol.Assertion, 0), make([]*protocol.AssertionCreatedInfo, 0)
	}
	genesis := &mocks.MockAssertion{
		MockId:        mockId(0),
		MockPrevId:    mockId(0),
		MockHeight:    0,
		MockStateHash: common.Hash{},
		Prev:          option.None[*mocks.MockAssertion](),
	}
	p.On(
		"GetAssertion",
		ctx,
		mockId(uint64(0)),
	).Return(genesis, nil)
	assertions := []protocol.Assertion{genesis}
	creationInfo := make([]*protocol.AssertionCreatedInfo, 0)
	for i := 1; i <= num; i++ {
		mockHash := common.BytesToHash([]byte(fmt.Sprintf("%d", i)))
		mockAssertion := &mocks.MockAssertion{
			MockId:        mockId(uint64(i)),
			MockPrevId:    mockId(uint64(i - 1)),
			MockHeight:    uint64(i),
			MockStateHash: mockHash,
			Prev:          option.Some(assertions[i-1].(*mocks.MockAssertion)),
		}
		assertions = append(assertions, protocol.Assertion(mockAssertion))
		p.On(
			"GetAssertion",
			ctx,
			mockId(uint64(i)),
		).Return(protocol.Assertion(mockAssertion), nil)
		mockState := rollupgen.ExecutionState{
			MachineStatus: uint8(protocol.MachineStatusFinished),
			GlobalState: rollupgen.GlobalState(protocol.GoGlobalState{
				BlockHash: mockHash,
			}.AsSolidityStruct()),
		}
		mockAssertionCreationInfo := &protocol.AssertionCreatedInfo{
			AfterState:    mockState,
			InboxMaxCount: new(big.Int).SetUint64(uint64(i)),
		}
		creationInfo = append(creationInfo, mockAssertionCreationInfo)
		p.On(
			"ReadAssertionCreationInfo",
			ctx,
			mockId(uint64(i)),
		).Return(mockAssertionCreationInfo, nil)
		valid := validity(i)
		if !valid {
			s.On("AgreesWithExecutionState", ctx, protocol.GoExecutionStateFromSolidity(mockState)).Return(errors.New("invalid"))
		} else {
			s.On("AgreesWithExecutionState", ctx, protocol.GoExecutionStateFromSolidity(mockState)).Return(nil)
		}

	}
	var assertionHashes []protocol.AssertionHash
	for _, assertion := range assertions {
		assertionHashes = append(assertionHashes, assertion.Id())
	}
	p.On("LatestConfirmed", ctx).Return(genesis, nil)
	p.On("LatestCreatedAssertionHashes", ctx).Return(assertionHashes[1:], nil)
	return assertions, creationInfo
}

func setupPoster(t *testing.T) (*Manager, *mocks.MockProtocol, *mocks.MockStateManager) {
	t.Helper()
	chain := &mocks.MockProtocol{}
	ctx := context.Background()
	chain.On("CurrentChallengeManager", ctx).Return(&mocks.MockChallengeManager{}, nil)
	chain.On("SpecChallengeManager", ctx).Return(&mocks.MockSpecChallengeManager{}, nil)
	stateProvider := &mocks.MockStateManager{}
	_, err := setup.ChainsWithEdgeChallengeManager()
	require.NoError(t, err)
	p := &Manager{
		chain:               chain,
		stateManager:        stateProvider,
		submittedAssertions: threadsafe.NewLruSet[common.Hash](1000),
		isReadyToPost:       true,
		assertionChainData: &assertionChainData{
			latestAgreedAssertion: protocol.AssertionHash{},
			canonicalAssertions:   make(map[protocol.AssertionHash]*protocol.AssertionCreatedInfo),
		},
	}
	return p, chain, stateProvider
}

func mockId(x uint64) protocol.AssertionHash {
	return protocol.AssertionHash{Hash: common.BytesToHash([]byte(fmt.Sprintf("%d", x)))}
}
