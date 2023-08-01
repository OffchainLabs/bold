package api

import "github.com/ethereum/go-ethereum/common"

type Assertion struct {
	ID                  common.Hash    `json:"id"` // AKA: AssertionHash (?)
	ParentAssertionHash common.Hash    `json:"parentAssertionHash"`
	RequiredStake       string         `json:"requiredStake"` // Unit: wei
	ExecutionHash       common.Hash    `json:"executionHash"`
	InboxMaxCount       string         `json:"inboxMaxCount"`
	AfterInboxBatchAcc  common.Hash    `json:"afterInboxBatchAcc"`
	WasmModuleRoot      common.Hash    `json:"wasmModuleRoot"`
	ChallengeManager    common.Address `json:"challengeManager"`
}
