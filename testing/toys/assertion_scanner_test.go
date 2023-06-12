package toys

import (
	"context"
	"testing"

	"github.com/OffchainLabs/challenge-protocol-v2/testing/mocks"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/setup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

func Test_onLeafCreation(t *testing.T) {
	ctx := context.Background()
	t.Run("no fork detected", func(t *testing.T) {
		logsHook := test.NewGlobal()
		v, _, _ := setupValidator(t)

		prev := &mocks.MockAssertion{
			MockPrevId:       mockId(1),
			MockId:           mockId(1),
			MockStateHash:    common.Hash{},
			MockIsFirstChild: true,
		}
		ev := &mocks.MockAssertion{
			MockPrevId:       mockId(1),
			MockId:           mockId(2),
			MockStateHash:    common.BytesToHash([]byte("bar")),
			MockIsFirstChild: true,
		}

		p := &mocks.MockProtocol{}
		p.On("SpecChallengeManager", ctx).Return(&mocks.MockSpecChallengeManager{}, nil)
		p.On("GetAssertion", ctx, mockId(0)).Return(prev, nil)
		v.chain = p

		err := v.onLeafCreated(ctx, ev)
		require.NoError(t, err)
		AssertLogsContain(t, logsHook, "New assertion appended")
		AssertLogsContain(t, logsHook, "No fork detected in assertion tree")
	})
	t.Run("fork leads validator to challenge leaf", func(t *testing.T) {
		logsHook := test.NewGlobal()
		ctx := context.Background()
		createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{
			DivergeBlockHeight: 5,
		})
		require.NoError(t, err)

		validator, err := New(
			ctx,
			createdData.Chains[1],
			createdData.Backend,
			createdData.HonestStateManager,
			createdData.Addrs.Rollup,
		)
		require.NoError(t, err)

		err = validator.onLeafCreated(ctx, createdData.Leaf1)
		require.NoError(t, err)

		anotherValidator, err := New(
			ctx,
			createdData.Chains[0],
			createdData.Backend,
			createdData.EvilStateManager,
			createdData.Addrs.Rollup,
		)
		require.NoError(t, err)

		err = anotherValidator.onLeafCreated(ctx, createdData.Leaf2)
		require.NoError(t, err)

		AssertLogsContain(t, logsHook, "New assertion appended")
		AssertLogsContain(t, logsHook, "Successfully created level zero edge")

		err = anotherValidator.onLeafCreated(ctx, createdData.Leaf2)
		require.NoError(t, err)
	})
}
