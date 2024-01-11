package api

import (
	"math/big"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/ethereum/go-ethereum/common"
)

type JsonChallenge struct {
	AssertionHash common.Hash `json:"assertionHash"`
}

type JsonChallengeConfig struct {
	ConfirmPeriodBlocks          uint64         `json:"confirmPeriodBlocks"`
	StakeToken                   common.Address `json:"stakeToken"`
	BaseStake                    *big.Int       `json:"baseStake"`
	WasmModuleRoot               common.Hash    `json:"wasmModuleRoot"`
	MiniStakeValue               *big.Int       `json:"miniStakeValue"`
	LayerZeroBlockEdgeHeight     uint64         `json:"layerZeroBlockEdgeHeight"`
	LayerZeroBigStepEdgeHeight   uint64         `json:"layerZeroBigStepEdgeHeight"`
	LayerZeroSmallStepEdgeHeight uint64         `json:"layerZeroSmallStepEdgeHeight"`
	NumBigStepLevel              uint8          `json:"numBigStepLevel"`
	ChallengeGracePeriodBlocks   uint64         `json:"challengeGracePeriodBlocks"`
}

type JsonAssertion struct {
	Hash                     common.Hash            `json:"hash" db:"Hash"`
	ConfirmPeriodBlocks      uint64                 `json:"confirmPeriodBlocks" db:"ConfirmPeriodBlocks"`
	RequiredStake            string                 `json:"requiredStake" db:"RequiredStake"`
	ParentAssertionHash      common.Hash            `json:"parentAssertionHash" db:"ParentAssertionHash"`
	InboxMaxCount            string                 `json:"inboxMaxCount" db:"InboxMaxCount"`
	AfterInboxBatchAcc       common.Hash            `json:"afterInboxBatchAcc" db:"AfterInboxBatchAcc"`
	WasmModuleRoot           common.Hash            `json:"wasmModuleRoot" db:"WasmModuleRoot"`
	ChallengeManager         common.Address         `json:"challengeManager" db:"ChallengeManager"`
	CreationBlock            uint64                 `json:"creationBlock" db:"CreationBlock"`
	TransactionHash          common.Hash            `json:"transactionHash" db:"TransactionHash"`
	BeforeStateBlockHash     common.Hash            `json:"beforeStateBlockHash" db:"BeforeStateBlockHash"`
	BeforeStateSendRoot      common.Hash            `json:"beforeStateSendRoot" db:"BeforeStateSendRoot"`
	BeforeStateBatch         uint64                 `json:"beforeStateBatch" db:"BeforeStateBatch"`
	BeforeStatePosInBatch    uint64                 `json:"beforeStatePosInBatch" db:"BeforeStatePosInBatch"`
	BeforeStateMachineStatus protocol.MachineStatus `json:"beforeStateMachineStatus" db:"BeforeStateMachineStatus"`
	AfterStateBlockHash      common.Hash            `json:"afterStateBlockHash" db:"AfterStateBlockHash"`
	AfterStateSendRoot       common.Hash            `json:"afterStateSendRoot" db:"AfterStateSendRoot"`
	AfterStateBatch          uint64                 `json:"afterStateBatch" db:"AfterStateBatch"`
	AfterStatePosInBatch     uint64                 `json:"afterStatePosInBatch" db:"AfterStatePosInBatch"`
	AfterStateMachineStatus  protocol.MachineStatus `json:"afterStateMachineStatus" db:"AfterStateMachineStatus"`
	FirstChildBlock          *uint64                `json:"firstChildBlock" db:"FirstChildBlock"`
	SecondChildBlock         *uint64                `json:"secondChildBlock" db:"SecondChildBlock"`
	IsFirstChild             bool                   `json:"isFirstChild" db:"IsFirstChild"`
	Status                   string                 `json:"status" db:"Status"`
	ConfigHash               common.Hash            `json:"configHash" db:"ConfigHash"`
}

type JsonEdge struct {
	Id                common.Hash    `json:"id" db:"Id"`
	ChallengeLevel    uint8          `json:"challengeLevel" db:"ChallengeLevel"`
	StartHistoryRoot  common.Hash    `json:"startHistoryRoot" db:"StartHistoryRoot"`
	StartHeight       uint64         `json:"startHeight" db:"StartHeight"`
	EndHistoryRoot    common.Hash    `json:"endHistoryRoot" db:"EndHistoryRoot"`
	EndHeight         uint64         `json:"endHeight" db:"EndHeight"`
	CreatedAtBlock    uint64         `json:"createdAtBlock" db:"CreatedAtBlock"`
	MutualId          common.Hash    `json:"mutualId" db:"MutualId"`
	OriginId          common.Hash    `json:"originId" db:"OriginId"`
	ClaimId           common.Hash    `json:"claimId" db:"ClaimId"`
	HasChildren       bool           `json:"hasChildren" db:"HasChildren"`
	LowerChildId      common.Hash    `json:"lowerChildId" db:"LowerChildId"`
	UpperChildId      common.Hash    `json:"upperChildId" db:"UpperChildId"`
	MiniStaker        common.Address `json:"miniStaker" db:"MiniStaker"`
	AssertionHash     common.Hash    `json:"assertionHash" db:"AssertionHash"`
	TimeUnrivaled     uint64         `json:"timeUnrivaled" db:"TimeUnrivaled"`
	HasRival          bool           `json:"hasRival" db:"HasRival"`
	Status            string         `json:"status" db:"Status"`
	HasLengthOneRival bool           `json:"hasLengthOneRival" db:"HasLengthOneRival"`
	// Honest validator's point of view
	IsHonest            bool   `json:"isHonest"`
	IsRelevant          bool   `json:"isRelevant"`
	CumulativePathTimer uint64 `json:"cumulativePathTimer"`
}

type JsonStakeInfo struct {
	StakerAddresses       []common.Address `json:"stakerAddresses"`
	NumberOfMinistakes    uint64           `json:"numberOfMiniStakes"`
	StartCommitmentHeight uint64           `json:"startCommitmentHeight"`
	EndCommitmentHeight   uint64           `json:"endCommitmentHeight"`
}

type JsonMiniStakes struct {
	AssertionHash common.Hash    `json:"assertionHash"`
	Level         string         `json:"level"`
	StakeInfo     *JsonStakeInfo `json:"stakeInfo"`
}