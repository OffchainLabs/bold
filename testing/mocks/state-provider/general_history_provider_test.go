package stateprovider

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHistoryCommitment(t *testing.T) {

}

func Test_computeMachineStartIndex(t *testing.T) {
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
