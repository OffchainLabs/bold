package mmap

import (
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sys/unix"
)

type Mmap []byte

func NewMmap(size int) (Mmap, error) {
	return unix.Mmap(-1, 0, common.HashLength*size, unix.PROT_READ|unix.PROT_WRITE, unix.MAP_ANON|unix.MAP_PRIVATE)
}

func ConvertSliceToMmap(leaves []common.Hash) (Mmap, error) {
	data, err := NewMmap(len(leaves))
	if err != nil {
		return nil, err
	}
	for i, r := range leaves {
		copy(data[i*common.HashLength:], r.Bytes())
	}
	return data, nil
}

func (l Mmap) Free() {
	err := unix.Munmap(l)
	if err != nil {
		panic(err)
	}
}

func (l Mmap) Length() int {
	return len(l) / common.HashLength
}

func (l Mmap) SubMmap(start, end int) Mmap {
	return l[start*common.HashLength : end*common.HashLength]
}

func (l Mmap) Get(i int) common.Hash {
	return common.BytesToHash(l[i*common.HashLength : (i+1)*common.HashLength])
}

func (l Mmap) Set(i int, val common.Hash) {
	copy(l[i*common.HashLength:], val.Bytes())
}
