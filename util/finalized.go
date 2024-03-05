package util

import (
	"flag"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"
)

func GetFinalizedBlockNumber() *big.Int {
	// If we are running tests, we want to use the latest block number since
	// simulated backends only support the latest block number.
	if flag.Lookup("test.v") != nil {
		return nil
	}
	return big.NewInt(int64(rpc.FinalizedBlockNumber))
}
