package protocol

import (
	"context"
	"crypto/rand"
	"github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
)

func TestChallenge_EndToEndResolution(t *testing.T) {
	ctx := context.Background()
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	divergingAt := uint64(1)
	numRoots := uint64(5)
	correctStateRoots := createDivergingStateRoots(t, divergingAt, numRoots)
	wrongStateRoots := createDivergingStateRoots(t, divergingAt, numRoots)

	require.Equal(t, wrongStateRoots[0], correctStateRoots[0])
	require.NotEqual(t, wrongStateRoots[1], correctStateRoots[1])

	correctManager := statemanager.New(correctStateRoots)
	wrongManager := statemanager.New(wrongStateRoots)
	alice := common.BytesToAddress([]byte("alice"))
	bob := common.BytesToAddress([]byte("bob"))

	chain := NewAssertionChain(
		ctx,
		util.NewArtificialTimeReference(),
		time.Second,
	)

	bal := big.NewInt(0).Mul(AssertionStake, big.NewInt(100))
	chain.AddToBalance(tx, alice, bal)
	chain.AddToBalance(tx, bob, bal)
	commit := util.StateCommitment{
		StateRoot: correctStateRoots[1],
		Height:    1,
	}

	genesis := chain.LatestConfirmed(tx)
	assertion1, err := chain.CreateLeaf(
		tx,
		genesis,
		commit,
		alice,
	)
	require.NoError(t, err)
	assertion1.status = ConfirmedAssertionState
	chain.latestConfirmed = assertion1.SequenceNum

	// Creates two conflicting leaves in the assertion chain.
	aliceCommit := util.StateCommitment{
		StateRoot: correctStateRoots[3],
		Height:    3,
	}
	bobCommit := util.StateCommitment{
		StateRoot: wrongStateRoots[3],
		Height:    3,
	}
	aliceAssertion, bobAssertion := setupAssertionChainFork(
		t,
		chain,
		assertion1,
		alice,
		bob,
		aliceCommit,
		bobCommit,
	)

	// Create a BlockChallenge on genesis.
	challenge, err := assertion1.CreateChallenge(tx, ctx, alice)
	require.NoError(t, err)

	// Alice and bob will add challenge leaves to the BlockChallenge.
	aliceV := addBlockChallengeLeaf(
		t,
		correctStateRoots,
		challenge,
		aliceAssertion,
		alice,
	)
	bobV := addBlockChallengeLeaf(
		t,
		wrongStateRoots,
		challenge,
		bobAssertion,
		bob,
	)

	// Leaves have height 2, so after bisecting to 1, and a merge to 1,
	// they will be at a one-step-fork from the root vertex of the challenge.
	bobBisected := bisect(t, bobV, wrongManager, bob)
	aliceMerged := merge(t, aliceV, bobBisected, correctManager, alice)
	_ = aliceMerged

	ok, err := chain.IsAtOneStepFork(
		tx,
		challenge.Hash(),
		bobV.Commitment,
		bobV.Prev.Unwrap().Commitment,
	)
	require.NoError(t, err)
	require.True(t, ok)

	ok, err = chain.IsAtOneStepFork(
		tx,
		challenge.Hash(),
		aliceV.Commitment,
		aliceV.Prev.Unwrap().Commitment,
	)
	require.NoError(t, err)
	require.True(t, ok)

	// The non-presumptive vertex, created by Bob, should then open a
	// BigStepChallenge on vertex 1, which is the parent vertex
	// of the one-step-fork.
	require.True(t, bobBisected.IsPresumptiveSuccessor())

	subChal, err := bobBisected.CreateBigStepChallenge(tx, alice)
	require.NoError(t, err)

	// Alice and bob add leaves to the BigStepChallenge.
	aliceV = addBigStepChallengeLeaf(
		t,
		correctStateRoots,
		3,
		subChal,
		bobBisected,
		alice,
	)

	bobV = addBigStepChallengeLeaf(
		t,
		wrongStateRoots,
		3,
		subChal,
		bobBisected,
		bob,
	)

	// // A single bisection and merge will lead to a one-step-fork
	// // once more.
	// bobBisected = bisect(t, bobV, wrongManager, bob)
	// aliceMerged = merge(t, aliceV, bobBisected, correctManager, alice)
	// _ = aliceMerged

	// The non-presumptive vertex should then open a
	// SmallStepChallenge on vertex 1, which is the parent vertex
	// of the one-step-fork.

	// Alice and bob add leaves to the SmallStepChallenge.
	// A single bisection and merge will lead to a one-step-fork
	// once more.

	// After a bisection and merge, Alice and Bob reach a one-step-fork
	// which will be resolved with a one-step-proof. Alice should win.
}

func addBlockChallengeLeaf(
	t *testing.T,
	stateRoots []common.Hash,
	chal *Challenge,
	assertion *Assertion,
	staker common.Address,
) *ChallengeVertex {
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	height := assertion.StateCommitment.Height - assertion.Prev.Unwrap().StateCommitment.Height
	commit, err := util.NewHistoryCommitment(
		height,
		stateRoots[:height],
		util.WithLastElementProof(stateRoots[:height+2]),
	)
	require.NoError(t, err)
	leaf, err := chal.AddLeaf(
		tx,
		assertion,
		commit,
		staker,
	)
	require.NoError(t, err)
	return leaf
}

func bisect(
	t *testing.T,
	v *ChallengeVertex,
	manager statemanager.Manager,
	staker common.Address,
) *ChallengeVertex {
	ctx := context.Background()
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	prevHeight := v.Prev.Unwrap().Commitment.Height
	vHeight := v.Commitment.Height
	bisectionHeight, err := util.BisectionPoint(prevHeight, vHeight)
	require.NoError(t, err)

	proof, err := manager.PrefixProof(ctx, bisectionHeight, vHeight)
	require.NoError(t, err)

	history, err := manager.HistoryCommitmentUpTo(ctx, bisectionHeight)
	require.NoError(t, err)

	bisectedTo, err := v.Bisect(
		tx,
		history,
		proof,
		staker,
	)
	require.NoError(t, err)
	return bisectedTo
}

func merge(
	t *testing.T,
	mergingFrom,
	mergingTo *ChallengeVertex,
	manager statemanager.Manager,
	staker common.Address,
) *ChallengeVertex {
	ctx := context.Background()
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	mergingToHeight := mergingTo.Commitment.Height
	mergingFromHeight := mergingFrom.Commitment.Height

	proof, err := manager.PrefixProof(ctx, mergingToHeight, mergingFromHeight)
	require.NoError(t, err)

	err = mergingFrom.Merge(
		tx,
		mergingTo,
		proof,
		staker,
	)
	require.NoError(t, err)
	return mergingTo
}

func addBigStepChallengeLeaf(
	t *testing.T,
	stateRoots []common.Hash,
	height uint64,
	chal *Challenge,
	vertex *ChallengeVertex,
	staker common.Address,
) *ChallengeVertex {
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	commit, err := util.NewHistoryCommitment(
		height,
		stateRoots[:height],
		util.WithLastElementProof(stateRoots[:height+2]),
	)
	require.NoError(t, err)
	leaf, err := chal.AddSubchallengeLeaf(
		tx,
		vertex,
		BigStepChallenge,
		commit,
		staker,
	)
	require.NoError(t, err)
	return leaf
}

func setupAssertionChainFork(
	t *testing.T,
	chain *AssertionChain,
	latest *Assertion,
	staker1,
	staker2 common.Address,
	commit1,
	commit2 util.StateCommitment,
) (*Assertion, *Assertion) {
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	bal := big.NewInt(0).Mul(AssertionStake, big.NewInt(100))
	chain.AddToBalance(tx, staker1, bal)
	chain.AddToBalance(tx, staker2, bal)

	assertion1, err := chain.CreateLeaf(
		tx,
		latest,
		commit1,
		staker1,
	)
	require.NoError(t, err)
	assertion2, err := chain.CreateLeaf(
		tx,
		latest,
		commit2,
		staker2,
	)
	require.NoError(t, err)
	return assertion1, assertion2
}

func createDivergingStateRoots(t *testing.T, divergingAt, numRoots uint64) []common.Hash {
	stateRoots := make([]common.Hash, numRoots)
	for i := uint64(0); i < numRoots; i++ {
		if divergingAt == 0 || i < divergingAt {
			stateRoots[i] = util.HashForUint(i)
		} else {
			divergingRoot := make([]byte, 32)
			_, err := rand.Read(divergingRoot)
			require.NoError(t, err)
			stateRoots[i] = common.BytesToHash(divergingRoot)
		}
	}
	return stateRoots
}
