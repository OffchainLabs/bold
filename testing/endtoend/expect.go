package endtoend

import (
	"context"
	"testing"
	"time"

	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	retry "github.com/OffchainLabs/challenge-protocol-v2/runtime"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/rollupgen"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/endtoend/internal/backend"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
)

// expect is a function that will be called asynchronously to verify some success criteria
// for the given scenario.
type expect func(t *testing.T, ctx context.Context, be backend.Backend) error

// Expects that an assertion is confirmed by challenge win.
func expectAssertionConfirmedByChallengeWinner(t *testing.T, ctx context.Context, be backend.Backend) error {
	t.Run("assertion confirmed", func(t *testing.T) {
		rc, err := retry.UntilSucceeds(ctx, func() (*rollupgen.RollupCore, error) {
			return rollupgen.NewRollupCore(be.ContractAddresses().Rollup, be.Client())
		})
		require.NoError(t, err)

		latestHonestStaked, err := retry.UntilSucceeds(ctx, func() ([32]byte, error) {
			return rc.LatestStakedAssertion(&bind.CallOpts{Context: ctx}, be.Alice().From)
		})
		require.NoError(t, err)

		var confirmed bool
		for ctx.Err() == nil && !confirmed {
			i, err := retry.UntilSucceeds(ctx, func() (*rollupgen.RollupCoreAssertionConfirmedIterator, error) {
				return rc.FilterAssertionConfirmed(nil, nil)
			})
			require.NoError(t, err)
			for i.Next() {
				// Only check if the assertion id belongs to the honest staker.
				if i.Event.AssertionId != latestHonestStaked {
					continue
				}
				assertionNode, err := retry.UntilSucceeds(ctx, func() (rollupgen.AssertionNode, error) {
					return rc.GetAssertion(&bind.CallOpts{Context: ctx}, i.Event.AssertionId)
				})
				require.NoError(t, err)
				if assertionNode.Status != uint8(protocol.AssertionConfirmed) {
					t.Fatal("Confirmed assertion with unfinished state")
				}
				confirmed = true

				require.NoError(t, err)
				break
			}
			time.Sleep(500 * time.Millisecond) // Don't spam the backend.
		}

		if !confirmed {
			t.Fatal("assertion was not confirmed")
		}
	})
	return nil
}

// expectAliceAndBobStaked monitors EdgeAdded events until Alice and Bob are observed adding edges
// with a stake.
func expectAliceAndBobStaked(t *testing.T, ctx context.Context, be backend.Backend) error {
	t.Run("alice and bob staked", func(t *testing.T) {
		ecm, err := retry.UntilSucceeds(ctx, func() (*challengeV2gen.EdgeChallengeManager, error) {
			return edgeManager(be)
		})
		require.NoError(t, err)

		var aliceStaked, bobStaked bool
		for ctx.Err() == nil && (!aliceStaked || !bobStaked) {
			i, err := retry.UntilSucceeds(ctx, func() (*challengeV2gen.EdgeChallengeManagerEdgeAddedIterator, error) {
				return ecm.FilterEdgeAdded(nil, nil, nil, nil)
			})
			require.NoError(t, err)
			for i.Next() {
				e, err := retry.UntilSucceeds(ctx, func() (challengeV2gen.ChallengeEdge, error) {
					return ecm.GetEdge(nil, i.Event.EdgeId)
				})
				require.NoError(t, err)

				switch e.Staker {
				case be.Alice().From:
					aliceStaked = true
				case be.Bob().From:
					bobStaked = true
				}

				time.Sleep(500 * time.Millisecond) // Don't spam the backend.
			}
		}

		if !aliceStaked {
			t.Error("alice did not stake")
		}
		if !bobStaked {
			t.Error("bob did not stake")
		}
	})

	return nil
}
