package util

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func FormatHash(h common.Hash) string {
	return fmt.Sprintf("%#x", h[:8])
}

func FormatAddr(h common.Address) string {
	return fmt.Sprintf("%#x", h[len(h)-4:])
}
