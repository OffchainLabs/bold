// Package inclusionproofs defines a series of utilities for generating and verifying
// traditional Merkle proofs of data.
//
// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE
package inclusionproofs

import (
	"github.com/OffchainLabs/bold/mmap"
	prefixproofs "github.com/OffchainLabs/bold/state-commitments/prefix-proofs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"runtime"
	"sync"
)

var (
	ErrProofTooLong  = errors.New("merkle proof too long")
	ErrInvalidLeaves = errors.New("invalid number of leaves for merkle tree")
)

// FullTree generates a Merkle tree from a list of leaves.
func FullTree(leavesMmap mmap.Mmap) ([]mmap.Mmap, error) {
	msb, err := prefixproofs.MostSignificantBit(uint64(leavesMmap.Length()))
	if err != nil {
		return nil, err
	}
	lsb, err := prefixproofs.LeastSignificantBit(uint64(leavesMmap.Length()))
	if err != nil {
		return nil, err
	}
	maxLevel := msb + 1
	if msb == lsb {
		maxLevel = msb
	}

	layers := make([]mmap.Mmap, maxLevel+1)
	layers[0] = leavesMmap
	l := uint64(1)

	prevLayer := leavesMmap
	for prevLayer.Length() > 1 {
		nextLayer, err := mmap.NewMmap((prevLayer.Length() + 1) / 2)
		if err != nil {
			return nil, err
		}
		for i := 0; i < nextLayer.Length(); i++ {
			if 2*i+1 < prevLayer.Length() {
				nextLayer.Set(i, crypto.Keccak256Hash(prevLayer.Get(2*i).Bytes(), prevLayer.Get(2*i+1).Bytes()))
			} else {
				nextLayer.Set(i, crypto.Keccak256Hash(prevLayer.Get(2*i).Bytes(), (common.Hash{}).Bytes()))
			}
		}
		layers[l] = nextLayer
		prevLayer = nextLayer
		l++
	}
	return layers, nil
}

// GenerateInclusionProof from a list of Merkle leaves at a specified index.
func GenerateInclusionProof(leavesMmap mmap.Mmap, idx uint64) ([]common.Hash, error) {
	numLeaves := leavesMmap.Length()
	if numLeaves == 0 {
		return nil, ErrInvalidLeaves
	}
	if idx >= uint64(numLeaves) {
		return nil, ErrInvalidLeaves
	}
	if numLeaves == 1 {
		return make([]common.Hash, 0), nil
	}
	rehashed, err := mmap.NewMmap(numLeaves)
	if err != nil {
		return nil, err
	}
	var waitGroup sync.WaitGroup
	gomaxprocs := runtime.GOMAXPROCS(-1)
	waitGroup.Add(gomaxprocs)
	batchSize := numLeaves / gomaxprocs
	batchRemainder := numLeaves % gomaxprocs
	for i := 0; i < gomaxprocs-1; i++ {
		start := i * batchSize
		go func() {
			defer waitGroup.Done()
			for j := start; j < start+batchSize; j++ {
				rehashed.Set(j, crypto.Keccak256Hash(leavesMmap.Get(j).Bytes()))
			}
		}()
	}
	start := (gomaxprocs - 1) * batchSize
	go func() {
		defer waitGroup.Done()
		for j := start; j < start+batchSize+batchRemainder; j++ {
			rehashed.Set(j, crypto.Keccak256Hash(leavesMmap.Get(j).Bytes()))
		}
	}()
	waitGroup.Wait()

	fullT, err := FullTree(rehashed)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(fullT); i++ {
		defer fullT[i].Free()
	}
	maxLevel, err := prefixproofs.MostSignificantBit(uint64(rehashed.Length()) - 1)
	if err != nil {
		return nil, err
	}
	proof := make([]common.Hash, maxLevel+1)

	for level := uint64(0); level <= maxLevel; level++ {
		levelIndex := idx >> level
		counterpartIndex := levelIndex ^ 1
		layer := fullT[level]
		counterpart := common.Hash{}
		if counterpartIndex <= uint64(layer.Length())-1 {
			counterpart = layer.Get(int(counterpartIndex))
		}
		proof[level] = counterpart
	}

	return proof, nil
}

// CalculateRootFromProof calculates a Merkle root from a Merkle proof, index, and leaf.
func CalculateRootFromProof(proof []common.Hash, index uint64, leaf common.Hash) (common.Hash, error) {
	if len(proof) > 256 {
		return common.Hash{}, ErrProofTooLong
	}
	h := crypto.Keccak256Hash(leaf[:])
	for i := 0; i < len(proof); i++ {
		node := proof[i]
		if index&(1<<i) == 0 {
			h = crypto.Keccak256Hash(h[:], node[:])
		} else {
			h = crypto.Keccak256Hash(node[:], h[:])
		}
	}
	return h, nil
}
