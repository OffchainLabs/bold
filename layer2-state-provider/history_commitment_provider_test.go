package l2stateprovider

import (
	"testing"

	"github.com/OffchainLabs/bold/containers/option"
	"github.com/stretchr/testify/require"
)

func Test_computeRequiredNumberOfHashes(t *testing.T) {
	provider := &HistoryCommitmentProvider{
		challengeLeafHeights: []Height{
			4,
			8,
			16,
		},
	}

	_, err := provider.computeRequiredNumberOfHashes(
		[]Height{5},
		option.None[Height](),
	)
	require.ErrorContains(t, err, "invalid range: end 4 was < start 5")

	_, err = provider.computeRequiredNumberOfHashes(
		[]Height{0},
		option.Some(Height(5)),
	)
	require.ErrorContains(t, err, "end 5 was greater than max height for level 4")

	_, err = provider.computeRequiredNumberOfHashes(
		[]Height{0},
		option.Some(Height(5)),
	)
	require.ErrorContains(t, err, "end 5 was greater than max height for level 4")

	_, err = provider.computeRequiredNumberOfHashes(
		[]Height{0},
		option.Some(Height(5)),
	)
	require.ErrorContains(t, err, "end 5 was greater than max height for level 4")

	got, err := provider.computeRequiredNumberOfHashes(
		[]Height{0},
		option.Some(Height(4)),
	)
	require.NoError(t, err)
	require.Equal(t, uint64(5), got)

	got, err = provider.computeRequiredNumberOfHashes(
		[]Height{0, 0},
		option.Some(Height(4)),
	)
	require.NoError(t, err)
	require.Equal(t, uint64(5), got)

	got, err = provider.computeRequiredNumberOfHashes(
		[]Height{0, 0},
		option.None[Height](),
	)
	require.NoError(t, err)
	require.Equal(t, uint64(9), got)

	got, err = provider.computeRequiredNumberOfHashes(
		[]Height{0, 0, 0},
		option.None[Height](),
	)
	require.NoError(t, err)
	require.Equal(t, uint64(17), got)
}

func Test_computeMachineStartIndex(t *testing.T) {
	t.Run("block challenge level", func(t *testing.T) {
		provider := &HistoryCommitmentProvider{
			challengeLeafHeights: []Height{
				32,
				1 << 10,
				1 << 10,
			},
		}
		machineStartIdx := provider.computeMachineStartIndex(validatedStartHeights{1})
		require.Equal(t, OpcodeIndex(0), machineStartIdx)
	})
	t.Run("three subchallenge levels", func(t *testing.T) {
		provider := &HistoryCommitmentProvider{
			challengeLeafHeights: []Height{
				32, // block challenge level.
				32, // giant step challenge level.
				32, // big step challenge level.
				32, // small step challenge level.
			},
		}
		heights := []Height{
			0,
			3,
			4,
			5,
		}
		//	  3 * (32 * 32)
		//	+ 4 * (32)
		//	+ 5
		//  = 3205
		got := provider.computeMachineStartIndex(validatedStartHeights(heights))
		t.Log(got)
		require.Equal(t, OpcodeIndex(3205), got)
	})
	t.Run("four challenge levels", func(t *testing.T) {
		// take the following constants:
		//
		//	colossal_steps_per_giant_step = 16
		//	giant_steps_per_big_step      = 8
		//	big_steps_per_giant_step      = 4
		//	small_steps_per_big_step      = 2
		//
		// Let's say we want to figure out the machine start opcode index for the following inputs:
		//
		//	colossal_step=4, giant_step=5, big_step=1, small_step=0
		//
		// We can compute the opcode index using the following algorithm for the example above.
		//
		//	  4 * (8 * 4 * 2)
		//	+ 5 * (4 * 2)
		//	+ 1 * (2)
		//	+ 0
		//	= 298
		//
		provider := &HistoryCommitmentProvider{
			challengeLeafHeights: []Height{
				32, // Block challenge level.
				16,
				8,
				4,
				2,
			},
		}
		heights := []Height{
			0,
			4,
			5,
			1,
			0,
		}
		got := provider.computeMachineStartIndex(validatedStartHeights(heights))
		require.Equal(t, OpcodeIndex(298), got)
	})
}

func Test_computeStepSize(t *testing.T) {
	provider := &HistoryCommitmentProvider{
		challengeLeafHeights: []Height{
			1,
			2,
			4,
			8,
		},
	}
	t.Run("small step size", func(t *testing.T) {
		stepSize, err := provider.computeStepSize(validatedStartHeights{1, 2, 3, 4})
		require.NoError(t, err)
		// The step size for the last challenge level is always 1 opcode at a time.
		require.Equal(t, StepSize(1), stepSize)
	})
	t.Run("product of height constants for next challenge levels", func(t *testing.T) {
		stepSize, err := provider.computeStepSize(validatedStartHeights{1})
		require.NoError(t, err)
		// Product of height constants for challenge levels 1, 2, 3.
		require.Equal(t, StepSize(2*4*8), stepSize)
	})

}
