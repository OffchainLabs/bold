package protocol

import (
	"context"
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
	correctStateRoots := correctBlockHashesForTest(10)
	wrongStateRoots := wrongBlockHashesForTest(10)
	correctManager := statemanager.New(correctStateRoots)
	wrongManager := statemanager.New(wrongStateRoots)
	_ = correctManager
	_ = wrongManager
	chain := NewAssertionChain(
		ctx,
		util.NewArtificialTimeReference(),
		time.Second,
	)

	// Creates two conflicting leaves in the assertion chain.
	alice := common.BytesToAddress([]byte("alice"))
	bob := common.BytesToAddress([]byte("bob"))

	commit1 := util.StateCommitment{
		StateRoot: correctStateRoots[2],
		Height:    2,
	}
	commit2 := util.StateCommitment{
		StateRoot: wrongStateRoots[2],
		Height:    2,
	}
	assertion1, assertion2 := setupAssertionChainFork(
		t,
		ctx,
		chain,
		alice,
		bob,
		commit1,
		commit2,
	)
	_ = assertion1
	_ = assertion2

	genesis := chain.LatestConfirmed(tx)

	// Create a challenge on genesis.
	challenge, err := genesis.CreateChallenge(tx, ctx, alice)
	require.NoError(t, err)

	_ = challenge

	// Alice and bob will add challenge leaves.

	// Leaves have height 2, so a bisection from Alice to 1
	// and a merge from Bob to 1 will lead to a one-step-fork

	// The non-presumptive vertex should then open a
	// BigStepChallenge on vertex 1, which is the parent vertex
	// of the one-step-fork.

	// Alice and bob add leaves to the BigStepChallenge.
	// A single bisection and merge will lead to a one-step-fork
	// once more.

	// The non-presumptive vertex should then open a
	// SmallStepChallenge on vertex 1, which is the parent vertex
	// of the one-step-fork.

	// Alice and bob add leaves to the SmallStepChallenge.
	// A single bisection and merge will lead to a one-step-fork
	// once more.

	// After a bisection and merge, Alice and Bob reach a one-step-fork
	// which will be resolved with a one-step-proof. Alice should win.
	// expectedBisectionHeight := uint64(4)
	// lo := expectedBisectionHeight

	// hi := uint64(6)
	// loExp := util.ExpansionFromLeaves(wrongBlockHashes[:lo])
	// badCommit, err := util.NewHistoryCommitment(
	// 	hi,
	// 	wrongBlockHashes[:hi],
	// 	util.WithLastElementProof(wrongBlockHashes[:hi+1]),
	// )
	// require.NoError(t, err)

	// badLeaf, err := challenge.AddLeaf(
	// 	tx,
	// 	wrongBranch,
	// 	badCommit,
	// 	staker1,
	// )
	// require.NoError(t, err)

	// goodCommit, err := util.NewHistoryCommitment(
	// 	hi,
	// 	correctBlockHashes[:hi],
	// 	util.WithLastElementProof(correctBlockHashes[:hi+1]),
	// )
	// require.NoError(t, err)
	// goodLeaf, err := challenge.AddLeaf(
	// 	tx,
	// 	correctBranch,
	// 	goodCommit,
	// 	staker2,
	// )
	// require.NoError(t, err)
}

func setupAssertionChainFork(
	t *testing.T,
	ctx context.Context,
	chain *AssertionChain,
	staker1,
	staker2 common.Address,
	commit1,
	commit2 util.StateCommitment,
) (*Assertion, *Assertion) {
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	bal := big.NewInt(0).Mul(AssertionStake, big.NewInt(100))
	chain.AddToBalance(tx, staker1, bal)
	chain.AddToBalance(tx, staker2, bal)

	genesis := chain.LatestConfirmed(tx)
	assertion1, err := chain.CreateLeaf(
		tx,
		genesis,
		commit1,
		staker1,
	)
	require.NoError(t, err)
	assertion2, err := chain.CreateLeaf(
		tx,
		genesis,
		commit2,
		staker2,
	)
	require.NoError(t, err)
	return assertion1, assertion2
}
