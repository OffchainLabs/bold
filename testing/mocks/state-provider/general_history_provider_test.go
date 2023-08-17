package stateprovider

import (
	"fmt"
	"testing"

	"github.com/OffchainLabs/bold/containers/option"
	l2stateprovider "github.com/OffchainLabs/bold/layer2-state-provider"
	"github.com/stretchr/testify/require"
)

func TestHistoryCommitment(t *testing.T) {

}

func Test_computeMachineStartIndex(t *testing.T) {
	// Figure out the actual opcode index we should move the machine to
	// when we compute the history commitment. As there are different levels of challenge
	// granularity, we have to do some math to figure out the correct index.
	// Take, for example:
	//
	// 	challengeLeafHeights = [4, 8]
	//
	// This means there are 4 big steps per block challenge and 8 opcodes per big step.
	//
	// With the following inputs:
	//
	//  bigStepsPerBlockChal := 4
	//  opcodesPerBigStep := 8
	// 	bigStepRange := {From: 2, To: 3}
	// 	smallStepRange := {From: 6, To: 7}
	//
	// We want to move our machine to opcode (2 * 4 * 8) + 6 = 70
	//
	provider := &L2StateBackend{
		challengeLeafHeights: []uint64{
			4,
			8,
		},
	}
	got, err := provider.computeMachineStartIndex([]l2stateprovider.ClaimHeight{
		{From: 2, To: option.Some(uint64(3))},
		{From: 6, To: option.Some(uint64(7))},
	})
	require.NoError(t, err)
	t.Log(got)
	require.Equal(t, uint64(70), got)
}

func Test_computeStepIncrement(t *testing.T) {
	provider := &L2StateBackend{
		challengeLeafHeights: []uint64{
			1 << 11,
			1 << 12,
			1 << 13,
			1 << 14,
		},
	}
	requestedChallengeLevel := uint64(len(provider.challengeLeafHeights))
	_, err := provider.computeStepIncrement(requestedChallengeLevel)
	require.ErrorContains(t, err, fmt.Sprintf("requested challenge level %d >=", requestedChallengeLevel))

	requestedChallengeLevel = uint64(len(provider.challengeLeafHeights) - 1)
	got, err := provider.computeStepIncrement(requestedChallengeLevel)
	require.NoError(t, err)
	require.Equal(t, uint64(1<<14), got)

	requestedChallengeLevel = uint64(0)
	got, err = provider.computeStepIncrement(requestedChallengeLevel)
	require.NoError(t, err)
	require.Equal(t, uint64(1<<11), got)

	requestedChallengeLevel = uint64(1)
	got, err = provider.computeStepIncrement(requestedChallengeLevel)
	require.NoError(t, err)
	require.Equal(t, uint64(1<<12), got)
}
