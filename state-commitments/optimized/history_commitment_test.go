package optimized

import (
	"fmt"
	"testing"

	"github.com/OffchainLabs/bold/state-commitments/history"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// func TestPrefixProofGeneration(t *testing.T) {
// 	ctx := context.Background()
// 	merkleTreeContract, _ := setupMerkleTreeContract(t)
// 	verify := func(t *testing.T, computed *prefixProofComputation) {
// 		prefixExpRaw := make([][32]byte, len(computed.prefixExpansion))
// 		for i := 0; i < len(computed.prefixExpansion); i++ {
// 			var r [32]byte
// 			copy(r[:], computed.prefixExpansion[i][:])
// 			prefixExpRaw[i] = r
// 		}
// 		proofRaw := make([][32]byte, len(computed.proof))
// 		for i := 0; i < len(computed.proof); i++ {
// 			var r [32]byte
// 			copy(r[:], computed.proof[i][:])
// 			proofRaw[i] = r
// 		}
// 		err := prefixproofs.VerifyPrefixProof(&prefixproofs.VerifyPrefixProofConfig{
// 			PreRoot:      computed.prefixRoot,
// 			PreSize:      computed.prefixTotalLeaves,
// 			PostRoot:     computed.fullRoot,
// 			PostSize:     computed.fullTreeTotalLeaves,
// 			PreExpansion: computed.prefixExpansion,
// 			PrefixProof:  computed.proof,
// 		})
// 		require.NoError(t, err)
// 		err = merkleTreeContract.VerifyPrefixProof(
// 			&bind.CallOpts{},
// 			computed.prefixRoot,
// 			new(big.Int).SetUint64(computed.prefixTotalLeaves),
// 			computed.fullRoot,
// 			new(big.Int).SetUint64(computed.fullTreeTotalLeaves),
// 			prefixExpRaw,
// 			proofRaw,
// 		)
// 		require.NoError(t, err)
// 	}
// 	tests := []struct {
// 		realLength    uint64
// 		virtualLength uint64
// 	}{
// 		{1, 4},
// 		{2, 4},
// 		{3, 4},
// 		{4, 4},
// 		{1, 8},
// 		{2, 8},
// 		{3, 8},
// 		{4, 8},
// 		{5, 8},
// 		{6, 8},
// 		{7, 8},
// 		{8, 8},
// 		{1, 16},
// 	}

// 	for _, tt := range tests {
// 		for virtual := tt.realLength; virtual < tt.virtualLength; virtual++ {
// 			limit := nextPowerOf2(virtual)
// 			for prefixIndex := uint64(0); prefixIndex < virtual-1; prefixIndex++ {
// 				t.Run(fmt.Sprintf("real length %d, virtual %d, limit %d, prefix index %d", tt.realLength, virtual, limit, prefixIndex), func(t *testing.T) {
// 					legacy := computeLegacyPrefixProof(t, ctx, virtual, prefixIndex)
// 					optimized := computeOptimizedPrefixProof(t, tt.realLength, virtual, limit, prefixIndex)
// 					verify(t, legacy)
// 					verify(t, optimized)
// 				})
// 			}
// 		}
// 	}
// }

// func BenchmarkPrefixProofGeneration_Legacy(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		prefixIndex := 13384
// 		simpleHash := crypto.Keccak256Hash([]byte("foo"))
// 		hashes := make([]common.Hash, 1<<14)
// 		for i := 0; i < len(hashes); i++ {
// 			hashes[i] = simpleHash
// 		}

// 		lowCommitmentNumLeaves := prefixIndex + 1
// 		hiCommitmentNumLeaves := (1 << 14)
// 		prefixExpansion, err := prefixproofs.ExpansionFromLeaves(hashes[:lowCommitmentNumLeaves])
// 		require.NoError(b, err)
// 		_, err = prefixproofs.GeneratePrefixProof(
// 			uint64(lowCommitmentNumLeaves),
// 			prefixExpansion,
// 			hashes[lowCommitmentNumLeaves:hiCommitmentNumLeaves],
// 			prefixproofs.RootFetcherFromExpansion,
// 		)
// 		require.NoError(b, err)
// 	}
// }

// func BenchmarkPrefixProofGeneration_Optimized(b *testing.B) {
// 	b.StopTimer()
// 	simpleHash := crypto.Keccak256Hash([]byte("foo"))
// 	hashes := []common.Hash{crypto.Keccak256Hash(simpleHash[:])}
// 	prefixIndex := uint64(13384)
// 	virtual := uint64(1 << 14)
// 	builder, err := NewBuilder().Virtual(virtual).Limit(virtual).Build()
// 	require.NoError(b, err)
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		_, _, err := builder.GeneratePrefixProof(prefixIndex, hashes)
// 		require.NoError(b, err)
// 	}
// }

// type prefixProofComputation struct {
// 	prefixRoot          common.Hash
// 	fullRoot            common.Hash
// 	prefixTotalLeaves   uint64
// 	fullTreeTotalLeaves uint64
// 	prefixExpansion     []common.Hash
// 	proof               []common.Hash
// }

// func computeOptimizedPrefixProof(t *testing.T, numRealHashes int, virtual int, limit int, prefixIndex int) *prefixProofComputation {
// 	// Computes the prefix proof and expansion.
// 	simpleHash := crypto.Keccak256Hash([]byte("foo"))
// 	hashes := make([]common.Hash, numRealHashes)
// 	for i := 0; i < len(hashes); i++ {
// 		hashes[i] = crypto.Keccak256Hash(simpleHash[:])
// 	}
// 	prefixExp, proof, err := prefixAndProof(prefixIndex, hashes, virtual)
// 	require.NoError(t, err)
// 	lastLeafFillers = nil

// 	// Computes the prefix tree.
// 	hashes = make([]common.Hash, prefixIndex+1)
// 	for i := 0; i < len(hashes); i++ {
// 		hashes[i] = crypto.Keccak256Hash(simpleHash[:])
// 	}
// 	computeVirtualSparseTree(hashes, prefixIndex+1, 0)
// 	hashes = make([]common.Hash, prefixIndex+1)
// 	for i := 0; i < len(hashes); i++ {
// 		hashes[i] = crypto.Keccak256Hash(simpleHash[:])
// 	}
// 	prefixRoot, err := computeVirtualSparseTree(hashes, prefixIndex+1, nextPowerOf2(prefixIndex+1))
// 	require.NoError(t, err)

// 	// Computes the full tree.
// 	lastLeafFillers = nil
// 	hashes = make([]common.Hash, numRealHashes)
// 	for i := 0; i < len(hashes); i++ {
// 		hashes[i] = crypto.Keccak256Hash(simpleHash[:])
// 	}
// 	computeVirtualSparseTree(hashes, virtual, 0)
// 	hashes = make([]common.Hash, numRealHashes)
// 	for i := 0; i < len(hashes); i++ {
// 		hashes[i] = crypto.Keccak256Hash(simpleHash[:])
// 	}
// 	fullTreeRoot, err := computeVirtualSparseTree(hashes, virtual, limit)
// 	require.NoError(t, err)
// 	prefixExp = trimTrailingZeroHashes(prefixExp)
// 	proof = trimZeroes(proof)
// 	return &prefixProofComputation{
// 		prefixRoot:          prefixRoot,
// 		fullRoot:            fullTreeRoot,
// 		prefixTotalLeaves:   uint64(prefixIndex) + 1,
// 		fullTreeTotalLeaves: uint64(virtual),
// 		prefixExpansion:     prefixExp,
// 		proof:               proof,
// 	}
// }

// func computeLegacyPrefixProof(t *testing.T, ctx context.Context, numHashes int, prefixIndex int) *prefixProofComputation {
// 	simpleHash := crypto.Keccak256Hash([]byte("foo"))
// 	hashes := make([]common.Hash, numHashes)
// 	for i := 0; i < len(hashes); i++ {
// 		hashes[i] = simpleHash
// 	}
// 	manager, err := statemanager.NewWithMockedStateRoots(hashes)
// 	require.NoError(t, err)

// 	wasmModuleRoot := common.Hash{}
// 	startMessageNumber := l2stateprovider.Height(0)
// 	fromMessageNumber := l2stateprovider.Height(prefixIndex)
// 	req := &l2stateprovider.HistoryCommitmentRequest{
// 		WasmModuleRoot:              wasmModuleRoot,
// 		FromBatch:                   0,
// 		ToBatch:                     10,
// 		UpperChallengeOriginHeights: []l2stateprovider.Height{},
// 		FromHeight:                  startMessageNumber,
// 		UpToHeight:                  option.Some(l2stateprovider.Height(fromMessageNumber)),
// 	}
// 	loCommit, err := manager.HistoryCommitment(ctx, req)
// 	require.NoError(t, err)

// 	req.UpToHeight = option.Some(l2stateprovider.Height(numHashes - 1))
// 	hiCommit, err := manager.HistoryCommitment(ctx, req)
// 	require.NoError(t, err)

// 	packedProof, err := manager.PrefixProof(ctx, req, fromMessageNumber)
// 	require.NoError(t, err)

// 	data, err := statemanager.ProofArgs.Unpack(packedProof)
// 	require.NoError(t, err)
// 	preExpansion, ok := data[0].([][32]byte)
// 	require.Equal(t, true, ok)
// 	proof, ok := data[1].([][32]byte)
// 	require.Equal(t, true, ok)

// 	preExpansionHashes := make([]common.Hash, len(preExpansion))
// 	for i := 0; i < len(preExpansion); i++ {
// 		preExpansionHashes[i] = preExpansion[i]
// 	}
// 	prefixProof := make([]common.Hash, len(proof))
// 	for i := 0; i < len(proof); i++ {
// 		prefixProof[i] = proof[i]
// 	}
// 	return &prefixProofComputation{
// 		prefixRoot:          loCommit.Merkle,
// 		fullRoot:            hiCommit.Merkle,
// 		prefixTotalLeaves:   uint64(prefixIndex) + 1,
// 		fullTreeTotalLeaves: uint64(numHashes),
// 		prefixExpansion:     preExpansionHashes,
// 		proof:               prefixProof,
// 	}
// }

func TestLegacyVsOptimized(t *testing.T) {
	end := uint64(1 << 9)
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	for i := uint64(1); i < end; i++ {
		limit := nextPowerOf2(i)
		for j := i; j < limit; j++ {
			hashedLeaves := make([]common.Hash, i)
			for i := range hashedLeaves {
				hashedLeaves[i] = crypto.Keccak256Hash(simpleHash[:])
			}
			builder := NewBuilder().Virtual(uint64(j)).Limit(limit)
			committer, err := builder.Build()
			require.NoError(t, err)
			computedRoot, err := committer.ComputeRoot(hashedLeaves)
			require.NoError(t, err)

			legacyInputLeaves := make([]common.Hash, j)
			for i := range legacyInputLeaves {
				legacyInputLeaves[i] = simpleHash
			}
			histCommit, err := history.New(legacyInputLeaves)
			require.NoError(t, err)
			require.Equal(t, computedRoot, histCommit.Merkle)
		}
	}
}

func TestLegacyVsOptimizedEdgeCases(t *testing.T) {
	simpleHash := crypto.Keccak256Hash([]byte("foo"))

	tests := []struct {
		realLength    int
		virtualLength int
		limit         int
	}{
		{12, 14, 16},
		{8, 10, 16},
		{6, 6, 8},
		{10, 16, 16},
		{4, 8, 8},
		{1, 5, 8},
		{3, 5, 8},
		{5, 5, 8},
		{1023, 1024, 1024},
		{(1 << 14) - 7, (1 << 14) - 7, 1 << 14},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("real length %d, virtual %d, limit %d", tt.realLength, tt.virtualLength, tt.limit), func(t *testing.T) {
			hashedLeaves := make([]common.Hash, tt.realLength)
			for i := range hashedLeaves {
				hashedLeaves[i] = crypto.Keccak256Hash(simpleHash[:])
			}
			builder := NewBuilder().Virtual(uint64(tt.virtualLength)).Limit(uint64(tt.limit))
			committer, err := builder.Build()
			require.NoError(t, err)
			computedRoot, err := committer.ComputeRoot(hashedLeaves)
			require.NoError(t, err)

			leaves := make([]common.Hash, tt.virtualLength)
			for i := range leaves {
				leaves[i] = simpleHash
			}
			histCommit, err := history.New(leaves)
			require.NoError(t, err)
			require.Equal(t, computedRoot, histCommit.Merkle)
		})
	}
}

func TestVirtualSparse(t *testing.T) {
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	t.Run("real length 1, virtual length 3, limit 4", func(t *testing.T) {
		builder := NewBuilder().Virtual(3).Limit(4)
		committer, err := builder.Build()
		require.NoError(t, err)
		computedRoot, err := committer.ComputeRoot([]common.Hash{crypto.Keccak256Hash(simpleHash[:])})
		require.NoError(t, err)

		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := history.New(leaves)
		require.NoError(t, err)
		require.Equal(t, histCommit.Merkle, computedRoot)
	})
	t.Run("real length 2, virtual length 3, limit 4", func(t *testing.T) {
		hashedLeaves := []common.Hash{
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
		}
		builder := NewBuilder().Virtual(3).Limit(4)
		committer, err := builder.Build()
		require.NoError(t, err)
		computedRoot, err := committer.ComputeRoot(hashedLeaves)
		require.NoError(t, err)
		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := history.New(leaves)
		require.NoError(t, err)
		require.Equal(t, histCommit.Merkle, computedRoot)
	})
	t.Run("real length 3, virtual length 3, limit 4", func(t *testing.T) {
		hashedLeaves := []common.Hash{
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
		}
		builder := NewBuilder().Virtual(3).Limit(4)
		committer, err := builder.Build()
		require.NoError(t, err)
		computedRoot, err := committer.ComputeRoot(hashedLeaves)
		require.NoError(t, err)
		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := history.New(leaves)
		require.NoError(t, err)
		require.Equal(t, histCommit.Merkle, computedRoot)
	})
	t.Run("real length 4, virtual length 4, limit 4", func(t *testing.T) {
		hashedLeaves := []common.Hash{
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
		}
		builder := NewBuilder().Virtual(4).Limit(4)
		committer, err := builder.Build()
		require.NoError(t, err)
		computedRoot, err := committer.ComputeRoot(hashedLeaves)
		require.NoError(t, err)
		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := history.New(leaves)
		require.NoError(t, err)
		require.Equal(t, histCommit.Merkle, computedRoot)
	})
	t.Run("real length 1, virtual length 5, limit 8", func(t *testing.T) {
		hashedLeaves := []common.Hash{
			crypto.Keccak256Hash(simpleHash[:]),
		}
		builder := NewBuilder().Virtual(5).Limit(8)
		committer, err := builder.Build()
		require.NoError(t, err)
		computedRoot, err := committer.ComputeRoot(hashedLeaves)
		require.NoError(t, err)

		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := history.New(leaves)
		require.NoError(t, err)
		require.Equal(t, computedRoot, histCommit.Merkle)
	})
	t.Run("real length 12, virtual length 14, limit 16", func(t *testing.T) {
		hashedLeaves := make([]common.Hash, 12)
		for i := range hashedLeaves {
			hashedLeaves[i] = crypto.Keccak256Hash(simpleHash[:])
		}
		builder := NewBuilder().Virtual(14).Limit(16)
		committer, err := builder.Build()
		require.NoError(t, err)
		computedRoot, err := committer.ComputeRoot(hashedLeaves)
		require.NoError(t, err)

		leaves := make([]common.Hash, 14)
		for i := range leaves {
			leaves[i] = simpleHash
		}
		histCommit, err := history.New(leaves)
		require.NoError(t, err)
		require.Equal(t, computedRoot, histCommit.Merkle)
	})
}

func TestMaximumDepthHistoryCommitment(t *testing.T) {
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	hashedLeaves := []common.Hash{
		crypto.Keccak256Hash(simpleHash[:]),
	}
	builder := NewBuilder().Virtual(1 << 26).Limit(1 << 26)
	committer, err := builder.Build()
	require.NoError(t, err)
	_, err = committer.ComputeRoot(hashedLeaves)
	require.NoError(t, err)
}

func BenchmarkMaximumDepthHistoryCommitment(b *testing.B) {
	b.StopTimer()
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	hashedLeaves := []common.Hash{
		crypto.Keccak256Hash(simpleHash[:]),
	}
	builder := NewBuilder().Virtual(1 << 26).Limit(1 << 26)
	committer, err := builder.Build()
	require.NoError(b, err)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err = committer.ComputeRoot(hashedLeaves)
		require.NoError(b, err)
	}
}

// func setupMerkleTreeContract(t testing.TB) (*mocksgen.MerkleTreeAccess, *simulated.Backend) {
// 	numChains := uint64(1)
// 	accs, backend := setupAccounts(t, numChains)
// 	_, _, merkleTreeContract, err := mocksgen.DeployMerkleTreeAccess(accs[0].txOpts, backend.Client())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	backend.Commit()
// 	return merkleTreeContract, backend
// }

// // Represents a test EOA account in the simulated backend,
// type testAccount struct {
// 	accountAddr common.Address
// 	txOpts      *bind.TransactOpts
// }

// func setupAccounts(t testing.TB, numAccounts uint64) ([]*testAccount, *simulated.Backend) {
// 	genesis := make(core.GenesisAlloc)
// 	gasLimit := uint64(100000000)

// 	accs := make([]*testAccount, numAccounts)
// 	for i := uint64(0); i < numAccounts; i++ {
// 		privKey, err := crypto.GenerateKey()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		pubKeyECDSA, ok := privKey.Public().(*ecdsa.PublicKey)
// 		if !ok {
// 			t.Fatal("not ok")
// 		}

// 		// Strip off the 0x and the first 2 characters 04 which is always the
// 		// EC prefix and is not required.
// 		publicKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)[4:]
// 		var pubKey = make([]byte, 48)
// 		copy(pubKey, publicKeyBytes)

// 		addr := crypto.PubkeyToAddress(privKey.PublicKey)
// 		chainID := big.NewInt(1337)
// 		txOpts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		startingBalance, _ := new(big.Int).SetString(
// 			"100000000000000000000000000000000000000",
// 			10,
// 		)
// 		genesis[addr] = core.GenesisAccount{Balance: startingBalance}
// 		accs[i] = &testAccount{
// 			accountAddr: addr,
// 			txOpts:      txOpts,
// 		}
// 	}
// 	backend := simulated.NewBackend(genesis, simulated.WithBlockGasLimit(gasLimit))
// 	return accs, backend
// }
