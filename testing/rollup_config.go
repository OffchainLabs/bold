// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/challenge-protocol-v2/blob/main/LICENSE

package challenge_testing

import (
	"math/big"

	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/rollupgen"
	"github.com/ethereum/go-ethereum/common"
)

const (
	LevelZeroBlockEdgeHeight     = 1 << 5
	LevelZeroBigStepEdgeHeight   = 1 << 5
	LevelZeroSmallStepEdgeHeight = 1 << 5
)

type LevelZeroHeights struct {
	BlockChallengeHeight     uint64
	BigStepChallengeHeight   uint64
	SmallStepChallengeHeight uint64
}

type Opt func(c *rollupgen.Config)

func WithLevelZeroHeights(heights *LevelZeroHeights) Opt {
	return func(c *rollupgen.Config) {
		c.LayerZeroBlockEdgeHeight = new(big.Int).SetUint64(heights.BlockChallengeHeight)
		c.LayerZeroBigStepEdgeHeight = new(big.Int).SetUint64(heights.BigStepChallengeHeight)
		c.LayerZeroSmallStepEdgeHeight = new(big.Int).SetUint64(heights.SmallStepChallengeHeight)
	}
}

func GenerateRollupConfig(
	prod bool,
	wasmModuleRoot common.Hash,
	rollupOwner common.Address,
	chainId *big.Int,
	loserStakeEscrow common.Address,
	miniStakeValue *big.Int,
	opts ...Opt,
) rollupgen.Config {
	var confirmPeriod uint64
	if prod {
		confirmPeriod = 45818
	} else {
		confirmPeriod = 25
	}

	cfg := rollupgen.Config{
		MiniStakeValue:      miniStakeValue,
		ConfirmPeriodBlocks: confirmPeriod,
		StakeToken:          common.Address{},
		BaseStake:           big.NewInt(100),
		WasmModuleRoot:      wasmModuleRoot,
		Owner:               rollupOwner,
		LoserStakeEscrow:    loserStakeEscrow,
		ChainId:             chainId,
		SequencerInboxMaxTimeVariation: rollupgen.ISequencerInboxMaxTimeVariation{
			DelayBlocks:   big.NewInt(60 * 60 * 24 / 15),
			FutureBlocks:  big.NewInt(12),
			DelaySeconds:  big.NewInt(60 * 60 * 24),
			FutureSeconds: big.NewInt(60 * 60),
		},
		LayerZeroBlockEdgeHeight:     big.NewInt(LevelZeroBlockEdgeHeight),
		LayerZeroBigStepEdgeHeight:   big.NewInt(LevelZeroBigStepEdgeHeight),
		LayerZeroSmallStepEdgeHeight: big.NewInt(LevelZeroSmallStepEdgeHeight),
	}
	for _, o := range opts {
		o(&cfg)
	}
	return cfg
}
