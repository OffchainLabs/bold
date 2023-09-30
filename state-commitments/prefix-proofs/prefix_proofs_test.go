// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package prefixproofs_test

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/OffchainLabs/bold/solgen/go/mocksgen"
	prefixproofs "github.com/OffchainLabs/bold/state-commitments/prefix-proofs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestAppendCompleteSubTree(t *testing.T) {
	// Test case: Level >= MAX_LEVEL
	_, err := prefixproofs.AppendCompleteSubTree([]common.Hash{{1}}, prefixproofs.MAX_LEVEL, common.Hash{2})
	require.ErrorContains(t, err, "level too high")

	// Test case: Empty Subtree Root
	_, err = prefixproofs.AppendCompleteSubTree([]common.Hash{{1}}, 1, common.Hash{})
	require.ErrorContains(t, err, "cannot append empty")

	// Test case: Expansion Too Large
	_, err = prefixproofs.AppendCompleteSubTree(make([]common.Hash, prefixproofs.MAX_LEVEL+1), 1, common.Hash{2})
	require.ErrorContains(t, err, "merkle expansion to large")

	// Test case: Empty 'me' Array
	_, err = prefixproofs.AppendCompleteSubTree([]common.Hash{}, 1, common.Hash{2})
	require.NoError(t, err)

	// Test case: Level >= len(me)
	_, err = prefixproofs.AppendCompleteSubTree([]common.Hash{{1}}, 2, common.Hash{2})
	require.ErrorContains(t, err, "failing before for loop: level too high")
}

func TestGeneratePrefixProof(t *testing.T) {
	defaultLeaves := []common.Hash{{1}, {2}}

	// Test case: Zero PrefixHeight
	_, err := prefixproofs.GeneratePrefixProof(0, nil, defaultLeaves, nil)
	require.ErrorContains(t, err, "prefixHeight was 0")

	// Test case: Zero Length of Leaves
	_, err = prefixproofs.GeneratePrefixProof(1, nil, []common.Hash{}, nil)
	require.ErrorContains(t, err, "length of leaves was 0")
}

func TestRoot(t *testing.T) {
	t.Run("tree with exactly size MAX_LEVEL should pass validation", func(t *testing.T) {
		tree := make([]common.Hash, prefixproofs.MAX_LEVEL)
		_, err := prefixproofs.Root(tree)
		require.NotEqual(t, prefixproofs.ErrLevelTooHigh, err)
	})
	t.Run("tree too large", func(t *testing.T) {
		tree := make([]common.Hash, prefixproofs.MAX_LEVEL+1)
		_, err := prefixproofs.Root(tree)
		require.Equal(t, prefixproofs.ErrExpansionTooLarge, err)
	})
	t.Run("empty tree", func(t *testing.T) {
		tree := make([]common.Hash, 0)
		_, err := prefixproofs.Root(tree)
		require.Equal(t, prefixproofs.ErrRootForEmpty, err)
	})
	t.Run("single element returns itself", func(t *testing.T) {
		tree := make([]common.Hash, 1)
		tree[0] = common.HexToHash("0x1234")
		root, err := prefixproofs.Root(tree)
		require.NoError(t, err)
		require.Equal(t, tree[0], root)
	})
}

func TestLeastSignificantBit_GoSolidityEquivalence(t *testing.T) {
	merkleTreeContract, _ := setupMerkleTreeContract(t)
	runBitEquivalenceTest(t, merkleTreeContract.LeastSignificantBit, prefixproofs.LeastSignificantBit)
}

func TestMostSignificantBit_GoSolidityEquivalence(t *testing.T) {
	merkleTreeContract, _ := setupMerkleTreeContract(t)
	runBitEquivalenceTest(t, merkleTreeContract.MostSignificantBit, prefixproofs.MostSignificantBit)
}

func FuzzPrefixProof_MaximumAppendBetween_GoSolidityEquivalence(f *testing.F) {
	type prePost struct {
		pre  uint64
		post uint64
	}
	testcases := []prePost{
		{4, 8},
		{10, 0},
		{0, 0},
		{0, 1},
		{3, 3},
		{3, 4},
		{0, 15},
		{128, 512},
		{128, 200},
		{128, 1 << 20},
		{1 << 20, 1<<20 + 1},
	}
	for _, tc := range testcases {
		f.Add(tc.pre, tc.post)
	}
	merkleTreeContract, _ := setupMerkleTreeContract(f)
	opts := &bind.CallOpts{}
	f.Fuzz(func(t *testing.T, pre, post uint64) {
		gotGo, err1 := prefixproofs.MaximumAppendBetween(pre, post)
		gotSol, err2 := merkleTreeContract.MaximumAppendBetween(opts, big.NewInt(int64(pre)), big.NewInt(int64(post)))
		if err1 == nil && err2 == nil {
			if !gotSol.IsUint64() {
				t.Fatal("sol result was not a uint64")
			}
			if gotSol.Uint64() != gotGo {
				t.Errorf("sol %d != go %d", gotSol.Uint64(), gotGo)
			}
		}
	})
}

func FuzzPrefixProof_BitUtils_GoSolidityEquivalence(f *testing.F) {
	testcases := []uint64{
		0,
		2,
		3,
		4,
		7,
		8,
		100,
		1 << 32,
		1<<32 - 1,
		1<<32 + 1,
		1 << 40,
	}
	for _, tc := range testcases {
		f.Add(tc)
	}
	merkleTreeContract, _ := setupMerkleTreeContract(f)
	opts := &bind.CallOpts{}
	f.Fuzz(func(t *testing.T, x uint64) {
		lsbSol, _ := merkleTreeContract.LeastSignificantBit(opts, big.NewInt(int64(x)))
		lsbGo, _ := prefixproofs.LeastSignificantBit(x)
		if lsbSol != nil {
			if !lsbSol.IsUint64() {
				t.Fatal("lsb sol not a uint64")
			}
			if lsbSol.Uint64() != lsbGo {
				t.Errorf("Mismatch lsb sol=%d, go=%d", lsbSol, lsbGo)
			}
		}
		msbSol, _ := merkleTreeContract.MostSignificantBit(opts, big.NewInt(int64(x)))
		msbGo, _ := prefixproofs.MostSignificantBit(x)
		if msbSol != nil {
			if !msbSol.IsUint64() {
				t.Fatal("msb sol not a uint64")
			}
			if msbSol.Uint64() != msbGo {
				t.Errorf("Mismatch msb sol=%d, go=%d", msbSol, msbGo)
			}
		}
	})
}

func runBitEquivalenceTest(
	t testing.TB,
	solFunc func(opts *bind.CallOpts, x *big.Int) (*big.Int, error),
	goFunc func(x uint64) (uint64, error),
) {
	opts := &bind.CallOpts{}
	for _, tt := range []struct {
		num        uint64
		wantSolErr bool
		solErr     string
		wantGoErr  bool
		goErr      error
	}{
		{
			num:        0,
			wantSolErr: true,
			solErr:     "has no significant bits",
			wantGoErr:  true,
			goErr:      prefixproofs.ErrCannotBeZero,
		},
		{num: 2},
		{num: 3},
		{num: 4},
		{num: 7},
		{num: 8},
		{num: 10},
		{num: 100},
		{num: 256},
		{num: 1 << 32},
		{num: 1<<32 + 1},
		{num: 1<<32 - 1},
		{num: 10231920391293},
	} {
		lsbSol, err := solFunc(opts, big.NewInt(int64(tt.num)))
		if tt.wantSolErr {
			require.NotNil(t, err)
			require.ErrorContains(t, err, tt.solErr)
		} else {
			require.NoError(t, err)
		}

		lsbGo, err := goFunc(tt.num)
		if tt.wantGoErr {
			require.NotNil(t, err)
			require.ErrorIs(t, err, tt.goErr)
		} else {
			require.NoError(t, err)
			require.Equal(t, lsbSol.Uint64(), lsbGo)
		}
	}
}

func setupMerkleTreeContract(t testing.TB) (*mocksgen.MerkleTreeAccess, *backends.SimulatedBackend) {
	numChains := uint64(1)
	accs, backend := setupAccounts(t, numChains)
	_, _, merkleTreeContract, err := mocksgen.DeployMerkleTreeAccess(accs[0].txOpts, backend)
	if err != nil {
		t.Fatal(err)
	}
	backend.Commit()
	return merkleTreeContract, backend
}

// Represents a test EOA account in the simulated backend,
type testAccount struct {
	accountAddr common.Address
	txOpts      *bind.TransactOpts
}

func setupAccounts(t testing.TB, numAccounts uint64) ([]*testAccount, *backends.SimulatedBackend) {
	genesis := make(core.GenesisAlloc)
	gasLimit := uint64(100000000)

	accs := make([]*testAccount, numAccounts)
	for i := uint64(0); i < numAccounts; i++ {
		privKey, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		pubKeyECDSA, ok := privKey.Public().(*ecdsa.PublicKey)
		if !ok {
			t.Fatal("not ok")
		}

		// Strip off the 0x and the first 2 characters 04 which is always the
		// EC prefix and is not required.
		publicKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)[4:]
		var pubKey = make([]byte, 48)
		copy(pubKey, publicKeyBytes)

		addr := crypto.PubkeyToAddress(privKey.PublicKey)
		chainID := big.NewInt(1337)
		txOpts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
		if err != nil {
			t.Fatal(err)
		}
		startingBalance, _ := new(big.Int).SetString(
			"100000000000000000000000000000000000000",
			10,
		)
		genesis[addr] = core.GenesisAccount{Balance: startingBalance}
		accs[i] = &testAccount{
			accountAddr: addr,
			txOpts:      txOpts,
		}
	}
	backend := backends.NewSimulatedBackend(genesis, gasLimit)
	return accs, backend
}
