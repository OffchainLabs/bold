package main

import (
	"context"
	"flag"
	"math/big"
	"strings"

	"github.com/OffchainLabs/bold/solgen/go/mocksgen"
	"github.com/OffchainLabs/bold/solgen/go/rollupgen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	valPrivKeys       = flag.String("validator-priv-keys", "", "comma-separated, validator private keys to fund and approve mock ERC20 stake token")
	l1ChainIdStr      = flag.String("l1-chain-id", "11155111", "l1 chain id")
	l1EndpointUrl     = flag.String("l1-endpoint", "ws://localhost:8546", "l1 endpoint")
	rollupAddrStr     = flag.String("rollup-address", "", "rollup address")
	stakeTokenAddrStr = flag.String("stake-token-address", "", "rollup address")
	gweiToDeposit     = flag.Uint64("gwei-to-deposit", 10_000, "tokens to deposit")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	endpoint, err := rpc.DialWebsocket(ctx, *l1EndpointUrl, "*")
	if err != nil {
		panic(err)
	}
	client := ethclient.NewClient(endpoint)
	l1ChainId, ok := new(big.Int).SetString(*l1ChainIdStr, 10)
	if !ok {
		panic("not big int")
	}
	if *valPrivKeys == "" {
		panic("no validator private keys set")
	}
	privKeyStrings := strings.Split(*valPrivKeys, ",")
	for _, privKeyStr := range privKeyStrings {
		validatorPrivateKey, err := crypto.HexToECDSA(privKeyStr)
		if err != nil {
			panic(err)
		}
		txOpts, err := bind.NewKeyedTransactorWithChainID(validatorPrivateKey, l1ChainId)
		if err != nil {
			panic(err)
		}

		rollupAddr := common.HexToAddress(*rollupAddrStr)
		rollupBindings, err := rollupgen.NewRollupUserLogicCaller(rollupAddr, client)
		if err != nil {
			panic(err)
		}
		chalManagerAddr, err := rollupBindings.ChallengeManager(&bind.CallOpts{})
		if err != nil {
			panic(err)
		}

		stakeTokenAddr := common.HexToAddress(*stakeTokenAddrStr)
		tokenBindings, err := mocksgen.NewTestWETH9(stakeTokenAddr, client)
		if err != nil {
			panic(err)
		}
		depositAmount := new(big.Int).SetUint64(*gweiToDeposit * params.GWei)
		txOpts.Value = depositAmount
		if _, err = tokenBindings.Deposit(txOpts); err != nil {
			panic(err)
		}
		txOpts.Value = big.NewInt(0)
		maxUint256 := new(big.Int)
		maxUint256.Exp(big.NewInt(2), big.NewInt(256), nil).Sub(maxUint256, big.NewInt(1))
		_, err = tokenBindings.Approve(txOpts, rollupAddr, maxUint256)
		if err != nil {
			panic(err)
		}
		_, err = tokenBindings.Approve(txOpts, chalManagerAddr, maxUint256)
		if err != nil {
			panic(err)
		}

	}
}
