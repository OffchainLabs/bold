package state_hashes

import "github.com/ethereum/go-ethereum/common"

// StateHashes is a wrapper around a slice of state hashes, which
// provides a Length method to return the number of hashes in the slice,
// and an At method to return the hash at a given index.
// If the requested index is greater than the length of the state hashes,
// the last hash in the slice is returned.
type StateHashes struct {
	Hashes []common.Hash `json:"hashes"`
	Len    uint64        `json:"len"`
}

func NewStateHashes(hashes []common.Hash, length uint64) *StateHashes {
	return &StateHashes{
		Hashes: hashes,
		Len:    length,
	}
}

func (s *StateHashes) Length() uint64 {
	return s.Len
}

func (s *StateHashes) At(i uint64) common.Hash {
	if i >= s.Len {
		panic("index out of range")
	}
	if len(s.Hashes) == 0 {
		panic("empty state hashes")
	}
	if uint64(len(s.Hashes)) > i {
		return s.Hashes[i]
	}
	// If the requested index is greater than the length of the state hashes, return the last hash
	return s.Hashes[len(s.Hashes)-1]
}

func (s *StateHashes) SubSlice(start, end uint64) *StateHashes {
	if start > s.Len {
		panic("index out of range")
	}
	if end > s.Len {
		panic("index out of range")
	}
	if start > end {
		panic("invalid slice indices")
	}
	// If start and end are within the length of the state hashes, return the sub-slice
	if end < uint64(len(s.Hashes)) {
		return &StateHashes{
			Hashes: s.Hashes[start:end],
			Len:    end - start,
		}
	}
	// If start is within the length of the state hashes, but end is greater than the length of the state hashes,
	// return the sub-slice from start to the end of the state hashes
	if start < uint64(len(s.Hashes)) {
		return &StateHashes{
			Hashes: s.Hashes[start:],
			Len:    end - start,
		}
	}
	// If both start and end are greater than the length of the state hashes, return the last hash in the slice
	return &StateHashes{
		Hashes: s.Hashes[len(s.Hashes)-1:],
		Len:    end - start,
	}
}
