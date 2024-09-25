package math

import "math/bits"

func Log2(u uint64) int {
	if u <= 0 {
		panic("log2 undefined for non-positive values")
	}
	return bits.Len64(u) - 1
}
