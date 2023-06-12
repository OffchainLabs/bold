package validator

import (
	"context"
	"fmt"
	"testing"

	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	"github.com/OffchainLabs/challenge-protocol-v2/containers/option"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/rollupgen"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/mocks"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/setup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var _ = ChallengeCreator(&Manager{})

func mockId(x uint64) protocol.AssertionId {
	return protocol.AssertionId(common.BytesToHash([]byte(fmt.Sprintf("%d", x))))
}

func setupAssertions(ctx context.Context, p *mocks.MockProtocol, s *mocks.MockStateManager, num int, validity func(int) bool) []protocol.Assertion {
	if num == 0 {
		return make([]protocol.Assertion, 0)
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
	for i := 1; i <= num; i++ {
		mockHash := common.BytesToHash([]byte(fmt.Sprintf("%d", i)))
		assertion := protocol.Assertion(&mocks.MockAssertion{
			MockId:        mockId(uint64(i)),
			MockPrevId:    mockId(uint64(i - 1)),
			MockHeight:    uint64(i),
			MockStateHash: mockHash,
			Prev:          option.Some(assertions[i-1].(*mocks.MockAssertion)),
		})
		assertions = append(assertions, assertion)
		p.On(
			"GetAssertion",
			ctx,
			mockId(uint64(i)),
		).Return(assertion, nil)
		mockState := rollupgen.ExecutionState{
			MachineStatus: uint8(protocol.MachineStatusFinished),
			GlobalState: rollupgen.GlobalState(protocol.GoGlobalState{
				BlockHash: mockHash,
			}.AsSolidityStruct()),
		}
		mockAssertionCreationInfo := &protocol.AssertionCreatedInfo{
			AfterState: mockState,
		}
		p.On(
			"ReadAssertionCreationInfo",
			ctx,
			mockId(uint64(i)),
		).Return(mockAssertionCreationInfo, nil)
		valid := validity(i)
		s.On("ExecutionStateBlockHeight", ctx, protocol.GoExecutionStateFromSolidity(mockState)).Return(uint64(i), valid)

		if i == 1 {
			var firstValid protocol.Assertion = genesis
			if valid {
				firstValid = assertion
			}
			p.On("LatestConfirmed", ctx).Return(firstValid, nil)
		}
	}
	p.On("LatestConfirmed", ctx).Return(assertions[0], nil)
	p.On("LatestCreatedAssertion", ctx).Return(assertions[len(assertions)-1], nil)
	return assertions
}

func setupValidator(t *testing.T) (*Manager, *mocks.MockProtocol, *mocks.MockStateManager) {
	t.Helper()
	p := &mocks.MockProtocol{}
	ctx := context.Background()
	p.On("CurrentChallengeManager", ctx).Return(&mocks.MockChallengeManager{}, nil)
	p.On("SpecChallengeManager", ctx).Return(&mocks.MockSpecChallengeManager{}, nil)
	s := &mocks.MockStateManager{}
	cfg, err := setup.ChainsWithEdgeChallengeManager()
	require.NoError(t, err)
	v, err := New(context.Background(), p, cfg.Backend, s, cfg.Addrs.Rollup)
	require.NoError(t, err)
	return v, p, s
}
