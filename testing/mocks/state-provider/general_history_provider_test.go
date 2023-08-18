package stateprovider

import (
	"context"
	"testing"

	"github.com/OffchainLabs/bold/containers/option"
	l2stateprovider "github.com/OffchainLabs/bold/layer2-state-provider"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestHistoryCommitment(t *testing.T) {
	ctx := context.Background()
	wasmModuleRoot := common.Hash{}
	provider := &L2StateBackend{
		challengeLeafHeights: []uint64{
			4,
			8,
			16,
		},
	}
	_, err := provider.HistoryCommitment(
		ctx,
		wasmModuleRoot,
		0,
		nil, // No start heights provided.
		option.None[l2stateprovider.Height](),
	)
	require.ErrorContains(t, err, "must specify at least one start height")
}

func Test_computeRequiredNumberOfHashes(t *testing.T) {
	provider := &L2StateBackend{
		challengeLeafHeights: []uint64{
			4,
			8,
			16,
		},
	}
	_, err := provider.computeRequiredNumberOfHashes(nil, option.None[l2stateprovider.Height]())
	require.ErrorContains(t, err, "must provide start heights")

	_, err = provider.computeRequiredNumberOfHashes(
		[]l2stateprovider.Height{1, 1, 1, 1},
		option.None[l2stateprovider.Height](),
	)
	require.ErrorContains(t, err, "challenge level 3 is out of range")

	_, err = provider.computeRequiredNumberOfHashes(
		[]l2stateprovider.Height{5},
		option.None[l2stateprovider.Height](),
	)
	require.ErrorContains(t, err, "invalid range: end 4 was < start 5")

	_, err = provider.computeRequiredNumberOfHashes(
		[]l2stateprovider.Height{0},
		option.Some(l2stateprovider.Height(5)),
	)
	require.ErrorContains(t, err, "end 5 was greater than max height for level 4")

	_, err = provider.computeRequiredNumberOfHashes(
		[]l2stateprovider.Height{0},
		option.Some(l2stateprovider.Height(5)),
	)
	require.ErrorContains(t, err, "end 5 was greater than max height for level 4")

	_, err = provider.computeRequiredNumberOfHashes(
		[]l2stateprovider.Height{0},
		option.Some(l2stateprovider.Height(5)),
	)
	require.ErrorContains(t, err, "end 5 was greater than max height for level 4")

	got, err := provider.computeRequiredNumberOfHashes(
		[]l2stateprovider.Height{0},
		option.Some(l2stateprovider.Height(4)),
	)
	require.NoError(t, err)
	require.Equal(t, uint64(5), got)

	got, err = provider.computeRequiredNumberOfHashes(
		[]l2stateprovider.Height{0, 0},
		option.Some(l2stateprovider.Height(4)),
	)
	require.NoError(t, err)
	require.Equal(t, uint64(5), got)

	got, err = provider.computeRequiredNumberOfHashes(
		[]l2stateprovider.Height{0, 0},
		option.None[l2stateprovider.Height](),
	)
	require.NoError(t, err)
	require.Equal(t, uint64(9), got)

	got, err = provider.computeRequiredNumberOfHashes(
		[]l2stateprovider.Height{0, 0, 0},
		option.None[l2stateprovider.Height](),
	)
	require.NoError(t, err)
	require.Equal(t, uint64(17), got)
}

func Test_computeMachineStartIndex(t *testing.T) {
}

func Test_computeStepIncrement(t *testing.T) {
}
