package server

import (
	"math/big"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/ethereum/go-ethereum/common"
)

type JsonChallenge struct {
	AssertionHash common.Hash              `json:"assertionHash"`
	CreationBlock uint64                   `json:"creationBlock"`
	Status        protocol.AssertionStatus `json:"status"`
	Config        *JsonChallengeConfig     `json:"config"`
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
	Hash                common.Hash              `json:"hash"`
	ConfirmPeriodBlocks uint64                   `json:"confirmPeriodBlocks"`
	RequiredStake       string                   `json:"requiredStake"`
	ParentAssertionHash common.Hash              `json:"parentAssertionHash"`
	InboxMaxCount       string                   `json:"inboxMaxCount"`
	AfterInboxBatchAcc  common.Hash              `json:"afterInboxBatchAcc"`
	WasmModuleRoot      common.Hash              `json:"wasmModuleRoot"`
	ChallengeManager    common.Address           `json:"challengeManager"`
	CreationBlock       uint64                   `json:"creationBlock"`
	TransactionHash     common.Hash              `json:"transactionHash"`
	BeforeState         *protocol.ExecutionState `json:"arbitrumBeforeState"`
	AfterState          *protocol.ExecutionState `json:"arbitrumAfterState"`
	FirstChildBlock     uint64                   `json:"firstChildBlock"`
	SecondChildBlock    uint64                   `json:"secondChildBlock"`
	IsFirstChild        bool                     `json:"isFirstChild"`
	Status              protocol.AssertionStatus `json:"status"`
	ConfigHash          common.Hash              `json:"configHash"`
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

type JsonCommitment struct {
	Height uint64      `json:"height"`
	Hash   common.Hash `json:"hash"`
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
