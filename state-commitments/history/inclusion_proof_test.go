package history

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/go-cmp/cmp"
)

func TestPaddingStartIndexAtLevel(t *testing.T) {
	tests := []struct {
		name   string
		leaves uint64
		level  uint64
		want   uint64
	}{
		{
			name:   "leaves 5, level 0",
			leaves: 5,
			level:  0,
			want:   5,
		},
		{
			name:   "leaves 5, level 1",
			leaves: 5,
			level:  1,
			want:   2,
		},
		{
			name:   "leaves 5, level 2",
			leaves: 5,
			level:  2,
			want:   1,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := paddingStartIndexAtLevel(tc.leaves, tc.level)
			if tc.want != got {
				t.Errorf("paddingStartIndexAtLevel(%d, %d): want %d, got %d", tc.leaves, tc.level, tc.want, got)
			}
		})
	}
}

func TestPaddingEndIndexAtLevel(t *testing.T) {
	tests := []struct {
		name    string
		virtual uint64
		level   uint64
		want    uint64
	}{
		{
			name:    "virtual 9, level 0",
			virtual: 9,
			level:   0,
			want:    8,
		},
		{
			name:    "virtual 9, level 1",
			virtual: 9,
			level:   1,
			want:    3,
		},
		{
			name:    "virtual 9, level 2",
			virtual: 9,
			level:   2,
			want:    1,
		},
		{
			name:    "virtual 5, level 0",
			virtual: 5,
			level:   0,
			want:    4,
		},
		{
			name:    "virtual 5, level 1",
			virtual: 5,
			level:   1,
			want:    1,
		},
		{
			name:    "virtual 5, level 2",
			virtual: 5,
			level:   2,
			want:    0,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := paddingEndIndexAtLevel(tc.virtual, tc.level)
			if tc.want != got {
				t.Errorf("paddingEndIndexAtLevel(%d, %d): want %d, got %d", tc.virtual, tc.level, tc.want, got)
			}
		})
	}
}

func TestInclusionProofs(t *testing.T) {
	aLeaf := common.HexToHash("0xA")
	bLeaf := common.HexToHash("0xB")
	aHash := crypto.Keccak256Hash(aLeaf[:])
	bHash := crypto.Keccak256Hash(bLeaf[:])
	zHash := emptyHash
	abHash := crypto.Keccak256Hash(append(aHash[:], bHash[:]...))
	bzHash := crypto.Keccak256Hash(append(bHash[:], emptyHash[:]...))
	bbHash := crypto.Keccak256Hash(append(bHash[:], bHash[:]...))
	tests := []struct {
		name string
		idx  uint64
		lvs  []common.Hash
		virt uint64
		want []common.Hash
	}{
		{
			name: "empty leaves",
			idx:  0,
			lvs:  []common.Hash{},
			virt: 0,
			want: nil,
		},
		{
			name: "single leaf",
			idx:  0,
			lvs:  []common.Hash{aLeaf},
			virt: 1,
			want: nil,
		},
		{
			name: "two leaves - idx 0",
			idx:  0,
			lvs:  []common.Hash{aLeaf, bLeaf},
			virt: 2,
			want: []common.Hash{bHash},
		},
		{
			name: "two leaves - idx 1",
			idx:  1,
			lvs:  []common.Hash{aLeaf, bLeaf},
			virt: 2,
			want: []common.Hash{aHash},
		},
		{
			name: "two leaves - idx 1, virtual 3",
			idx:  1,
			lvs:  []common.Hash{aLeaf, bLeaf},
			virt: 3,
			want: []common.Hash{aHash, bzHash},
		},
		{
			name: "two leaves - idx 2, virtual 3",
			idx:  2,
			lvs:  []common.Hash{aLeaf, bLeaf},
			virt: 3,
			want: []common.Hash{zHash, abHash},
		},
		{
			name: "two leaves - idx 1, virtual 4",
			idx:  1,
			lvs:  []common.Hash{aLeaf, bLeaf},
			virt: 4,
			want: []common.Hash{aHash, bbHash},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hc := NewCommitter()
			got, err := hc.computeMerkleProof(tc.idx, tc.lvs, tc.virt)
			if err != nil {
				t.Errorf("computeMerkelProof(%d, %v, %d): err %v", tc.idx, tc.lvs, tc.virt, err)
			}
			if diff := cmp.Diff(tc.want, got, cmp.Transformer("toHex", func(c common.Hash) string { return c.Hex() })); diff != "" {
				t.Errorf("unexpected proof (-want, +got):\n%s", diff)
			}
		})
	}
}
