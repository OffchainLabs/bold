package api

import (
	"math/big"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/ethereum/go-ethereum/common"
)

// type AssertionCreatedInfo struct {
// 	ConfirmPeriodBlocks uint64
// 	RequiredStake       *big.Int
// 	ParentAssertionHash common.Hash
// 	BeforeState         rollupgen.ExecutionState
// 	AfterState          rollupgen.ExecutionState
// 	InboxMaxCount       *big.Int
// 	AfterInboxBatchAcc  common.Hash
// 	AssertionHash       common.Hash
// 	WasmModuleRoot      common.Hash
// 	ChallengeManager    common.Address
// }

type Assertion struct {
	ConfirmPeriodBlocks uint64         `json:"confirmPeriodBlocks"`
	RequiredStake       string         `json:"requiredStake"`
	ParentAssertionHash common.Hash    `json:"parentAssertionHash"`
	InboxMaxCount       string         `json:"inboxMaxCount"`
	AfterInboxBatchAcc  common.Hash    `json:"afterInboxBatchAcc"`
	AssertionHash       common.Hash    `json:"assertionHash"`
	WasmModuleRoot      common.Hash    `json:"wasmModuleRoot"`
	ChallengeManager    common.Address `json:"challengeManager"`

	// TODO: Before / After states?
}

func AssertionCreatedInfoToAssertion(aci *protocol.AssertionCreatedInfo) *Assertion {
	if aci == nil {
		return nil
	}

	return &Assertion{
		ConfirmPeriodBlocks: aci.ConfirmPeriodBlocks,
		RequiredStake:       big.NewInt(0).Set(aci.RequiredStake).String(),
		ParentAssertionHash: aci.ParentAssertionHash,
		InboxMaxCount:       big.NewInt(0).Set(aci.InboxMaxCount).String(),
		AfterInboxBatchAcc:  aci.AfterInboxBatchAcc,
		AssertionHash:       aci.AssertionHash,
		WasmModuleRoot:      aci.WasmModuleRoot,
		ChallengeManager:    aci.ChallengeManager,
	}
}
