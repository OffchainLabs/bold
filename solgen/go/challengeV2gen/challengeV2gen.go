// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challengeV2gen

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ChallengeEdge is an auto generated low-level Go binding around an user-defined struct.
type ChallengeEdge struct {
	OriginId         [32]byte
	StartHistoryRoot [32]byte
	StartHeight      *big.Int
	EndHistoryRoot   [32]byte
	EndHeight        *big.Int
	LowerChildId     [32]byte
	UpperChildId     [32]byte
	ClaimId          [32]byte
	Staker           common.Address
	CreatedAtBlock   uint64
	ConfirmedAtBlock uint64
	Status           uint8
	Level            uint8
	Refunded         bool
}

// ConfigData is an auto generated low-level Go binding around an user-defined struct.
type ConfigData struct {
	WasmModuleRoot      [32]byte
	RequiredStake       *big.Int
	ChallengeManager    common.Address
	ConfirmPeriodBlocks uint64
	NextInboxPosition   uint64
}

// CreateEdgeArgs is an auto generated low-level Go binding around an user-defined struct.
type CreateEdgeArgs struct {
	Level          uint8
	EndHistoryRoot [32]byte
	EndHeight      *big.Int
	ClaimId        [32]byte
	PrefixProof    []byte
	Proof          []byte
}

// ExecutionState is an auto generated low-level Go binding around an user-defined struct.
type ExecutionState struct {
	GlobalState   GlobalState
	MachineStatus uint8
}

// ExecutionStateData is an auto generated low-level Go binding around an user-defined struct.
type ExecutionStateData struct {
	ExecutionState    ExecutionState
	PrevAssertionHash [32]byte
	InboxAcc          [32]byte
}

// GlobalState is an auto generated low-level Go binding around an user-defined struct.
type GlobalState struct {
	Bytes32Vals [2][32]byte
	U64Vals     [2]uint64
}

// OneStepData is an auto generated low-level Go binding around an user-defined struct.
type OneStepData struct {
	BeforeHash [32]byte
	Proof      []byte
}

// EdgeChallengeManagerMetaData contains all meta data concerning the EdgeChallengeManager contract.
var EdgeChallengeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAssertionChain\",\"name\":\"_assertionChain\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_challengePeriodBlocks\",\"type\":\"uint64\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"_oneStepProofEntry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"layerZeroBlockEdgeHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"layerZeroBigStepEdgeHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"layerZeroSmallStepEdgeHeight\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_excessStakeReceiver\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_numBigStepLevel\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AssertionHashEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"h2\",\"type\":\"bytes32\"}],\"name\":\"AssertionHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssertionNoSibling\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssertionNotPending\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"levels\",\"type\":\"uint8\"}],\"name\":\"BigStepLevelsTooMany\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lowerChildId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"upperChildId\",\"type\":\"bytes32\"}],\"name\":\"ChildrenAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"argLevel\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"claimLevel\",\"type\":\"uint8\"}],\"name\":\"ClaimEdgeInvalidLevel\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"}],\"name\":\"ClaimEdgeNotLengthOneRival\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimEdgeNotPending\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"EdgeAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"EdgeAlreadyRefunded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"claimingEdgeId\",\"type\":\"bytes32\"}],\"name\":\"EdgeClaimMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"edgeId2\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"level1\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"level2\",\"type\":\"uint8\"}],\"name\":\"EdgeLevelInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lowerChildId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"upperChildId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"ancestorEdgeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"}],\"name\":\"EdgeNotAncestor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"internalType\":\"enumEdgeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"EdgeNotConfirmed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"EdgeNotExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"}],\"name\":\"EdgeNotLayerZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"EdgeNotLengthOne\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"internalType\":\"enumEdgeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"EdgeNotPending\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"}],\"name\":\"EdgeTypeNotBlock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"}],\"name\":\"EdgeTypeNotSmallStep\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"EdgeUnrivaled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAssertionChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyChallengePeriod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyClaimId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyEdgeSpecificProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyEndMachineStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyEndRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyFirstRival\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyOneStepProofEntry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyOriginId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyPrefixProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyStakeReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyStaker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyStartMachineStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyStartRoot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"h1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"h2\",\"type\":\"uint256\"}],\"name\":\"HeightDiffLtTwo\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"thresholdBlocks\",\"type\":\"uint256\"}],\"name\":\"InsufficientConfirmationBlocks\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumEdgeType\",\"name\":\"edgeType\",\"type\":\"uint8\"}],\"name\":\"InvalidEdgeType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeight\",\"type\":\"uint256\"}],\"name\":\"InvalidEndHeight\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"InvalidHeights\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"numBigStepLevels\",\"type\":\"uint8\"}],\"name\":\"LevelTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxProofLength\",\"type\":\"uint256\"}],\"name\":\"MerkleProofTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"NotPowerOfTwo\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mutualId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"originId\",\"type\":\"bytes32\"}],\"name\":\"OriginIdMutualIdMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"confirmedRivalId\",\"type\":\"bytes32\"}],\"name\":\"RivalEdgeConfirmed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroBigStepLevels\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mutualId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"hasRival\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLayerZero\",\"type\":\"bool\"}],\"name\":\"EdgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"lowerChildId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"upperChildId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"lowerChildAlreadyExists\",\"type\":\"bool\"}],\"name\":\"EdgeBisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mutualId\",\"type\":\"bytes32\"}],\"name\":\"EdgeConfirmedByChildren\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mutualId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"claimingEdgeId\",\"type\":\"bytes32\"}],\"name\":\"EdgeConfirmedByClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mutualId\",\"type\":\"bytes32\"}],\"name\":\"EdgeConfirmedByOneStepProof\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mutualId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"totalTimeUnrivaled\",\"type\":\"uint64\"}],\"name\":\"EdgeConfirmedByTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mutualId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"EdgeRefunded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LAYERZERO_BIGSTEPEDGE_HEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LAYERZERO_BLOCKEDGE_HEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LAYERZERO_SMALLSTEPEDGE_HEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NUM_BIGSTEP_LEVEL\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assertionChain\",\"outputs\":[{\"internalType\":\"contractIAssertionChain\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bisectionHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefixProof\",\"type\":\"bytes\"}],\"name\":\"bisectEdge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"originId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"startHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"endHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"endHistoryRoot\",\"type\":\"bytes32\"}],\"name\":\"calculateEdgeId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"originId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"startHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"endHeight\",\"type\":\"uint256\"}],\"name\":\"calculateMutualId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengePeriodBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"confirmEdgeByChildren\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"claimingEdgeId\",\"type\":\"bytes32\"}],\"name\":\"confirmEdgeByClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structOneStepData\",\"name\":\"oneStepData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"requiredStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"challengeManager\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"confirmPeriodBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextInboxPosition\",\"type\":\"uint64\"}],\"internalType\":\"structConfigData\",\"name\":\"prevConfig\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"beforeHistoryInclusionProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"afterHistoryInclusionProof\",\"type\":\"bytes32[]\"}],\"name\":\"confirmEdgeByOneStepProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"ancestorEdges\",\"type\":\"bytes32[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState\",\"name\":\"globalState\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus\",\"name\":\"machineStatus\",\"type\":\"uint8\"}],\"internalType\":\"structExecutionState\",\"name\":\"executionState\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"prevAssertionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inboxAcc\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionStateData\",\"name\":\"claimStateData\",\"type\":\"tuple\"}],\"name\":\"confirmEdgeByTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mutualId\",\"type\":\"bytes32\"}],\"name\":\"confirmedRival\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"endHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"endHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefixProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structCreateEdgeArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"createLayerZeroEdge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"edgeExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"edgeLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excessStakeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mutualId\",\"type\":\"bytes32\"}],\"name\":\"firstRival\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"getEdge\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"originId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"startHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"endHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"endHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lowerChildId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"upperChildId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"createdAtBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"confirmedAtBlock\",\"type\":\"uint64\"},{\"internalType\":\"enumEdgeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"refunded\",\"type\":\"bool\"}],\"internalType\":\"structChallengeEdge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumEdgeType\",\"name\":\"eType\",\"type\":\"uint8\"}],\"name\":\"getLayerZeroEndHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"getPrevAssertionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"hasLengthOneRival\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"hasRival\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oneStepProofEntry\",\"outputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"refundStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"timeUnrivaled\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101c06040523480156200001257600080fd5b5060405162005eaa38038062005eaa833981016040819052620000359162000285565b6001600160a01b038a166200005d5760405163641f043160e11b815260040160405180910390fd5b6001600160a01b03808b166101005288166200008c5760405163fb60b0ef60e01b815260040160405180910390fd5b6001600160a01b038816610120526001600160401b038916600003620000c557604051632283bb7360e21b815260040160405180910390fd5b6001600160401b03891660e0526001600160a01b0380851660a05260c0849052821662000105576040516301e1d91560e31b815260040160405180910390fd5b6001600160a01b03821660805262000129876200023f602090811b6200170217901c565b6200014f57604051633abfb6ff60e21b8152600481018890526024015b60405180910390fd5b8661014081815250506200016e866200023f60201b620017021760201c565b6200019057604051633abfb6ff60e21b81526004810187905260240162000146565b856101608181525050620001af856200023f60201b620017021760201c565b620001d157604051633abfb6ff60e21b81526004810186905260240162000146565b61018085905260ff8116600003620001fc57604051632a18f5b960e21b815260040160405180910390fd5b60fd8160ff161115620002285760405163040d23bf60e41b815260ff8216600482015260240162000146565b60ff166101a052506200037c975050505050505050565b6000816000036200025257506000919050565b60006200026160018462000354565b929092161592915050565b6001600160a01b03811681146200028257600080fd5b50565b6000806000806000806000806000806101408b8d031215620002a657600080fd5b8a51620002b3816200026c565b60208c0151909a506001600160401b0381168114620002d157600080fd5b60408c0151909950620002e4816200026c565b8098505060608b0151965060808b0151955060a08b0151945060c08b01516200030d816200026c565b60e08c01516101008d0151919550935062000328816200026c565b6101208c015190925060ff811681146200034157600080fd5b809150509295989b9194979a5092959850565b818103818111156200037657634e487b7160e01b600052601160045260246000fd5b92915050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516159e7620004c3600039600081816103af015281816105d40152818161092f0152818161098701528181610af501528181610d330152818161107401526113970152600081816105230152610c6d01526000818161028b0152610c2d0152600081816102260152610bed0152600081816102ec0152818161090d01528181610965015261137001526000818161032b0152818161069d01528181610730015281816107ba0152818161085601528181610df301528181610ea801528181610f3101528181610fbf0152818161123c01526112d30152600081816102c501526110530152600081816103e8015281816109d10152611135015260008181610352015281816109b001526111140152600081816104e90152610a1f01526159e76000f3fe608060405234801561001057600080fd5b50600436106101ce5760003560e01c80635d9e244411610104578063bce6f54f116100a2578063e94e051e11610071578063e94e051e146104e4578063eae0328b1461050b578063f8ee77d61461051e578063fda2892e1461054557600080fd5b8063bce6f54f14610469578063c32d8c6314610489578063c8bc4e431461049c578063e5b123da146104c457600080fd5b8063748926f3116100de578063748926f31461041d578063750e0c0f146104305780638c1b3a4014610443578063908517e91461045657600080fd5b80635d9e2444146103aa57806360c7dc47146103e357806364deed591461040a57600080fd5b806342e1aaa81161017157806348dd29241161014b57806348dd29241461032657806351ed6a301461034d57806354b64151146103745780635a48e0f41461039757600080fd5b806342e1aaa8146102ad57806346c2781a146102c057806348923bc5146102e757600080fd5b80631dce5166116101ad5780631dce5166146102215780632eaa0043146102485780633e35f5e81461025b578063416e66571461028657600080fd5b80624d8efe146101d357806305fae141146101f95780630f73bfad1461020c575b600080fd5b6101e66101e1366004614b96565b610565565b6040519081526020015b60405180910390f35b6101e6610207366004614be0565b610580565b61021f61021a366004614c1a565b610aec565b005b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b61021f610256366004614c3c565b610b6e565b61026e610269366004614c3c565b610bbe565b6040516001600160401b0390911681526020016101f0565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b6101e66102bb366004614c70565b610bd0565b61026e7f000000000000000000000000000000000000000000000000000000000000000081565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101f0565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b610387610382366004614c3c565b610cba565b60405190151581526020016101f0565b6101e66103a5366004614c3c565b610cc6565b6103d17f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff90911681526020016101f0565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b61021f610418366004614dab565b610cd2565b61021f61042b366004614c3c565b6110fb565b61038761043e366004614c3c565b6111ed565b61021f610451366004614e59565b611216565b610387610464366004614c3c565b611408565b6101e6610477366004614c3c565b60009081526001602052604090205490565b6101e6610497366004614f21565b611414565b6104af6104aa366004614f63565b61142d565b604080519283526020830191909152016101f0565b6101e66104d2366004614c3c565b60009081526002602052604090205490565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b6101e6610519366004614c3c565b6115e4565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b610558610553366004614c3c565b6115f8565b6040516101f0919061500c565b600061057587878787878761172c565b979650505050505050565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905260006105f86105d260208601866150e0565b7f0000000000000000000000000000000000000000000000000000000000000000611771565b9050600061060582610bd0565b905061060f614a77565b600083600281111561062357610623614fe2565b0361095c5761063560a08701876150fb565b905060000361065757604051630c9ccac560e41b815260040160405180910390fd5b60008061066760a08901896150fb565b810190610674919061527b565b80516020820151604080840151905163f9cee9df60e01b81529497509295506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016945063f9cee9df936106d99360608f0135939291600401615367565b60006040518083038186803b1580156106f157600080fd5b505afa158015610705573d6000803e3d6000fd5b505050602080830151845191850151604080870151905163f9cee9df60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016955063f9cee9df94610769949390929091600401615367565b60006040518083038186803b15801561078157600080fd5b505afa158015610795573d6000803e3d6000fd5b505050506040518060c0016040528089606001358152602001826020015181526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e531d8c78b606001356040518263ffffffff1660e01b815260040161080a91815260200190565b602060405180830381865afa158015610827573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061084b919061538e565b1515815260200160007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166356bbc9e685602001516040518263ffffffff1660e01b81526004016108a691815260200190565b602060405180830381865afa1580156108c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e791906153b0565b6001600160401b03161181528351602082015282516040909101529250610953600089857f0000000000000000000000000000000000000000000000000000000000000000887f00000000000000000000000000000000000000000000000000000000000000006117e0565b955050506109ae565b6109ab600087837f0000000000000000000000000000000000000000000000000000000000000000867f00000000000000000000000000000000000000000000000000000000000000006117e0565b93505b7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03821615801590610a0757508015155b15610a585760008660c00151610a1d5730610a3f565b7f00000000000000000000000000000000000000000000000000000000000000005b9050610a566001600160a01b03841633838561186a565b505b8560400151866020015187600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a489606001518a608001518b60a001518c60c001518d60e00151604051610ad7959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a45050925195945050505050565b610b19600083837f00000000000000000000000000000000000000000000000000000000000000006118f0565b6000828152602081905260409020610b3090611a5c565b827fb924f3aa473645c7cf5b10262f927ae4ccf869d7fc239c17144b0c67490d1c7383604051610b6291815260200190565b60405180910390a35050565b610b79600082611a8c565b6000818152602081905260409020610b9090611a5c565b60405182907f0d27fcaf1adc41547a5cfc99d2364f6c0dc7e81c9fc3fe8cb38abb409b48358a90600090a350565b6000610bca8183611c3f565b92915050565b600080826002811115610be557610be5614fe2565b03610c1157507f0000000000000000000000000000000000000000000000000000000000000000919050565b6001826002811115610c2557610c25614fe2565b03610c5157507f0000000000000000000000000000000000000000000000000000000000000000919050565b6002826002811115610c6557610c65614fe2565b03610c9157507f0000000000000000000000000000000000000000000000000000000000000000919050565b81604051630efcb87b60e21b8152600401610cac91906153cd565b60405180910390fd5b919050565b6000610bca8183611ddf565b6000610bca8183611e14565b600080835111610ce25783610d0a565b8260018451610cf191906153f6565b81518110610d0157610d01615409565b60200260200101515b90506000610d188183611e65565b90506000610d578260090160099054906101000a900460ff167f0000000000000000000000000000000000000000000000000000000000000000611771565b90506000816002811115610d6d57610d6d614fe2565b14610d9c57600982015460405163ec72dc5d60e01b8152600160481b90910460ff166004820152602401610cac565b610da582611eb9565b610dee57610db282611edd565b60088301546007840154604051631cb1906160e31b815260048101939093526001600160a01b0390911660248301526044820152606401610cac565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633083622885600701546040518263ffffffff1660e01b8152600401610e4391815260200190565b602060405180830381865afa158015610e60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e84919061538e565b9050801561104357600784015460405163f9cee9df60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163f9cee9df91610eea91908a9060a08201359060c08301359060040161541f565b60006040518083038186803b158015610f0257600080fd5b505afa158015610f16573d6000803e3d6000fd5b5050604051631171558560e01b815260a089013560048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316925063117155859150602401602060405180830381865afa158015610f82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa691906153b0565b604051632b5de4f360e11b815260a088013560048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906356bbc9e690602401602060405180830381865afa15801561100e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103291906153b0565b61103c919061549c565b9150611048565b600091505b6000611098818a8a867f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611f12565b60008a81526020819052604090209091506110b290611a5c565b6040516001600160401b03831681528a907f9cd2c77f6772dd0fb07e9972aee8ddddf9fd13f1a65abb03a3de88b07dc59af99060200160405180910390a3505050505050505050565b60006111078183611e65565b905061111281612160565b7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0382161580159061116b57508015155b1561118c57600883015461118c906001600160a01b03848116911683612266565b60008481526020819052604090206111a390611a5c565b604080516001600160a01b03851681526020810184905286917fa635398959ddb5ce3b14537edfc25b2e671274c9b8cad0f4bd634752e69007b6910160405180910390a350505050565b600081815260208190526040812060080154600160a01b90046001600160401b03161515610bca565b60006112228189611e14565b6040516304972af960e01b81529091506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906304972af9906112739084908a906004016154d1565b60006040518083038186803b15801561128b57600080fd5b505afa15801561129f573d6000803e3d6000fd5b50505050600060405180606001604052808860800160208101906112c3919061554a565b6001600160401b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e78cea926040518163ffffffff1660e01b8152600401602060405180830381865afa15801561132f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113539190615567565b6001600160a01b03168152883560209091015290506113bb60008a7f00000000000000000000000000000000000000000000000000000000000000008b858b8b8b8b7f000000000000000000000000000000000000000000000000000000000000000061229b565b60008981526020819052604090206113d290611a5c565b6040518a907fe11db4b27bc8c6ea5943ecbb205ae1ca8d56c42c719717aaf8a53d43d0cee7c290600090a3505050505050505050565b6000610bca81836124df565b600061142386868686866125a5565b9695505050505050565b6000806000806000611479898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509594939250506126179050565b815192955090935091501580611511578260400151836020015184600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a4866060015187608001518860a001518960c001518a60e00151604051611508959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a45b8160400151826020015183600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a4856060015186608001518760a001518860c001518960e00151604051611590959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a48151604051821515815285908c907f7340510d24b7ec9b5c100f5500d93429d80d00d46f0d18e4e85d0c4cc22b99249060200160405180910390a4505191989197509095505050505050565b6000610bca6115f38284611e65565b6129e4565b611600614ab9565b61160b600083611e65565b604080516101c0810182528254815260018084015460208301526002840154928201929092526003830154606082015260048301546080820152600583015460a0820152600683015460c0820152600783015460e082015260088301546001600160a01b038116610100830152600160a01b90046001600160401b039081166101208301526009840154908116610140830152909291610160840191600160401b900460ff16908111156116c1576116c1614fe2565b60018111156116d2576116d2614fe2565b81526009919091015460ff600160481b820481166020840152600160501b90910416151560409091015292915050565b60008160000361171457506000919050565b60006117216001846153f6565b929092161592915050565b600061173b87878787876125a5565b60408051602081019290925281018390526060016040516020818303038152906040528051906020012090509695505050505050565b60008260ff1660000361178657506000610bca565b8160ff168360ff161161179b57506001610bca565b6117a6826001615584565b60ff168360ff16036117ba57506002610bca565b6040516315c1b4af60e31b815260ff808516600483015283166024820152604401610cac565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052908061182f8989898988612a29565b915091506000611840838a88612eb6565b9050600061184f83838c612fc9565b905061185b8b82613004565b9b9a5050505050505050505050565b6040516001600160a01b03808516602483015283166044820152606481018290526118ea9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166001600160e01b031990931692909217909152613319565b50505050565b600083815260208590526040902060080154600160a01b90046001600160401b03166119315760405162a7b02b60e01b815260048101849052602401610cac565b600082815260208590526040902060080154600160a01b90046001600160401b03166119725760405162a7b02b60e01b815260048101849052602401610cac565b6001600083815260208690526040902060090154600160401b900460ff1660018111156119a1576119a1614fe2565b146119e05760008281526020859052604090819020600901549051633bc499ed60e21b8152610cac918491600160401b90910460ff169060040161559d565b6119ec848484846133eb565b6000828152602085905260409020600701548314611a3b5760008281526020859052604090819020600701549051631855b87d60e31b8152610cac918591600401918252602082015260400190565b611a458484613513565b60008381526020859052604090206118ea9061357d565b6000610bca8260090160099054906101000a900460ff1683600001548460020154856001015486600401546125a5565b600081815260208390526040902060080154600160a01b90046001600160401b0316611acd5760405162a7b02b60e01b815260048101829052602401610cac565b60008181526020839052604080822060050154808352912060080154600160a01b90046001600160401b0316611b185760405162a7b02b60e01b815260048101829052602401610cac565b6001600082815260208590526040902060090154600160401b900460ff166001811115611b4757611b47614fe2565b14611b865760008181526020849052604090819020600901549051633bc499ed60e21b8152610cac918391600160401b90910460ff169060040161559d565b60008281526020849052604080822060060154808352912060080154600160a01b90046001600160401b0316611bd15760405162a7b02b60e01b815260048101829052602401610cac565b6001600082815260208690526040902060090154600160401b900460ff166001811115611c0057611c00614fe2565b14611a3b5760008181526020859052604090819020600901549051633bc499ed60e21b8152610cac918391600160401b90910460ff169060040161559d565b600081815260208390526040812060080154600160a01b90046001600160401b0316611c805760405162a7b02b60e01b815260048101839052602401610cac565b6000828152602084905260408120611c9790611a5c565b6000818152600186016020526040812054919250819003611ccb576040516336843d9f60e21b815260040160405180910390fd5b6040516815539492559053115160ba1b6020820152602901604051602081830303815290604052805190602001208103611d3457600084815260208690526040902060080154611d2b90600160a01b90046001600160401b03164361549c565b92505050610bca565b600081815260208690526040902060080154600160a01b90046001600160401b0316611d755760405162a7b02b60e01b815260048101829052602401610cac565b60008181526020869052604080822060089081015487845291909220909101546001600160401b03600160a01b928390048116929091041680821115611dca57611dbf818361549c565b945050505050610bca565b6000945050505050610bca565b505092915050565b6000611deb83836124df565b8015611e0d57506000828152602084905260409020611e09906129e4565b6001145b9392505050565b600080611e218484611e65565b90505b6009810154600160481b900460ff1615611e5d5780546000908152600185016020526040902054611e558582611e65565b915050611e24565b549392505050565b600081815260208390526040812060080154600160a01b90046001600160401b0316611ea65760405162a7b02b60e01b815260048101839052602401610cac565b5060009081526020919091526040902090565b600781015460009015801590610bca575050600801546001600160a01b0316151590565b6000610bca8260090160099054906101000a900460ff168360000154846002015485600101548660040154876003015461172c565b600085815260208790526040812060080154600160a01b90046001600160401b0316611f535760405162a7b02b60e01b815260048101879052602401610cac565b856000611f608983611c3f565b905060005b87518110156120e0576000611f938b8a8481518110611f8657611f86615409565b6020026020010151611e65565b90508381600501541480611faa5750838160060154145b15611fee57611fc18b611fbc83611edd565b611c3f565b611fcb90846155b1565b9250888281518110611fdf57611fdf615409565b602002602001015193506120cd565b600084815260208c9052604090206007015489518a908490811061201457612014615409565b602002602001015103612053576120468b8a848151811061203757612037615409565b602002602001015186896133eb565b611fc18b611fbc83611edd565b83816005015482600601548b858151811061207057612070615409565b60200260200101518e600001600089815260200190815260200160002060070154604051636ebd28c960e01b8152600401610cac959493929190948552602085019390935260408401919091526060830152608082015260a00190565b50806120d8816155d1565b915050611f65565b506120eb86826155b1565b9050846001600160401b0316816001600160401b031610156121335760405163011a8d4d60e41b81526001600160401b03808316600483015286166024820152604401610cac565b61213d8989613513565b600088815260208a9052604090206121549061357d565b98975050505050505050565b60016009820154600160401b900460ff16600181111561218257612182614fe2565b146121ba5761219081611edd565b6009820154604051633bc499ed60e21b8152610cac9291600160401b900460ff169060040161559d565b6121c381611eb9565b61220c576121d081611edd565b60088201546007830154604051631cb1906160e31b815260048101939093526001600160a01b0390911660248301526044820152606401610cac565b6009810154600160501b900460ff1615156001036122495761222d81611edd565b60405163307f766960e01b8152600401610cac91815260200190565b60090180546aff000000000000000000001916600160501b179055565b6040516001600160a01b03831660248201526044810182905261229690849063a9059cbb60e01b9060640161189e565b505050565b60006122a78b8b611e65565b600290810154915060008b815260208d905260409020600901546122d590600160481b900460ff1684611771565b60028111156122e6576122e6614fe2565b146123235760008a815260208c905260409081902060090154905163348aefdf60e01b8152600160481b90910460ff166004820152602401610cac565b60008a815260208c90526040902061233a906129e4565b6001146123745760008a815260208c905260409020612358906129e4565b6040516306b595e560e41b8152600401610cac91815260200190565b6123cf8b60000160008c8152602001908152602001600020600101548960000135838989808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506135ff92505050565b60006001600160a01b038a1663b5112fd289848c356123f160208f018f6150fb565b6040518663ffffffff1660e01b81526004016124119594939291906155ea565b602060405180830381865afa15801561242e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612452919061564c565b60008c815260208e905260409020600301549091506124b09082612477856001615665565b8888808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506135ff92505050565b6124ba8c8c613513565b60008b815260208d9052604090206124d19061357d565b505050505050505050505050565b600081815260208390526040812060080154600160a01b90046001600160401b03166125205760405162a7b02b60e01b815260048101839052602401610cac565b600082815260208490526040812061253790611a5c565b600081815260018601602052604081205491925081900361256b576040516336843d9f60e21b815260040160405180910390fd5b6040516815539492559053115160ba1b602082015260290160408051601f1981840301815291905280516020909101201415949350505050565b6040517fff0000000000000000000000000000000000000000000000000000000000000060f887901b1660208201526021810185905260418101849052606181018390526081810182905260009060a10160405160208183030381529060405280519060200120905095945050505050565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915260008087815260208990526040902060090154600160401b900460ff1660018111156126c6576126c6614fe2565b1461270557600086815260208890526040908190206009015490516323f8405d60e01b8152610cac918891600160401b90910460ff169060040161559d565b61270f87876124df565b61272f576040516380e07e4560e01b815260048101879052602401610cac565b6000868152602088905260408120604080516101c0810182528254815260018084015460208301526002840154928201929092526003830154606082015260048301546080820152600583015460a0820152600683015460c0820152600783015460e082015260088301546001600160a01b038116610100830152600160a01b90046001600160401b039081166101208301526009840154908116610140830152909291610160840191600160401b900460ff16908111156127f3576127f3614fe2565b600181111561280457612804614fe2565b81526009919091015460ff600160481b820481166020840152600160501b9091041615156040918201528101516080820151919250600091612846919061368c565b90506000808780602001905181019061285f91906156d3565b909250905061288f89612873856001615665565b60608701516080880151612888906001615665565b8686613720565b50506040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905260006128f28560000151866020015187604001518d888a6101800151613a0c565b90506128fd81613a9c565b600081815260208e90526040902060080154909350600160a01b90046001600160401b0316612933576129308c82613004565b91505b506040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810191909152600061299686600001518c8789606001518a608001518b6101800151613a0c565b90506129a28d82613004565b9150506129d28382600001518e60000160008f8152602001908152602001600020613ac59092919063ffffffff16565b919b909a509098509650505050505050565b600080826002015483600401546129fb91906153f6565b905080600003610bca57612a0e83611edd565b60405162a7b02b60e01b8152600401610cac91815260200190565b60408051606080820183526000808352602083015291810191909152600080612a5e612a5860208901896150e0565b85611771565b6002811115612a6f57612a6f614fe2565b03612ce85760208501518551600003612a9b576040516374b5e30d60e11b815260040160405180910390fd5b8551606088013514612ad05785516040516316c5de8f60e21b8152600481019190915260608801356024820152604401610cac565b8560400151612af2576040516360b4921b60e11b815260040160405180910390fd5b8560600151612b1457604051635a2e8e1d60e11b815260040160405180910390fd5b612b2160a08801886150fb565b9050600003612b4357604051630c9ccac560e41b815260040160405180910390fd5b6000612b5260a08901896150fb565b810190612b5f919061527b565b50909150600090508760800151602001516002811115612b8157612b81614fe2565b03612b9f5760405163231b2f2960e11b815260040160405180910390fd5b60008760a00151602001516002811115612bbb57612bbb614fe2565b03612bd957604051638999857d60e01b815260040160405180910390fd5b60808701516040516330e5867160e21b81526000916001600160a01b0389169163c39619c491612c0b91600401615736565b602060405180830381865afa158015612c28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c4c919061564c565b90506000876001600160a01b031663c39619c48a60a001516040518263ffffffff1660e01b8152600401612c809190615736565b602060405180830381865afa158015612c9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cc1919061564c565b6040805160608101825293845260208401919091528201929092529350909150612eac9050565b612cf6878760600135611ddf565b612d32576040517fff6d9bd700000000000000000000000000000000000000000000000000000000815260608701356004820152602401610cac565b6060860135600090815260208890526040812090612d4f82611a5c565b905060006009830154600160401b900460ff166001811115612d7357612d73614fe2565b14612d91576040516312459ffd60e01b815260040160405180910390fd5b6009820154612daa90600160481b900460ff1686613b2c565b60ff16612dba60208a018a6150e0565b60ff1614612e0357612dcf60208901896150e0565b600983015460405163564f308b60e11b815260ff9283166004820152600160481b9091049091166024820152604401610cac565b612e1060a08901896150fb565b9050600003612e3257604051630c9ccac560e41b815260040160405180910390fd5b600080808080612e4560a08e018e6150fb565b810190612e529190615744565b94509450945094509450612e708760010154868960020154866135ff565b612e848760030154858960040154856135ff565b6040518060600160405280868152602001858152602001828152508698509850505050505050505b9550959350505050565b604080516000808252602082019092528190612edc90612ed7908751613b4e565b613b84565b9050612ee783611702565b612f0757604051633abfb6ff60e21b815260048101849052602401610cac565b82846040013514612f3857604080516337f318af60e21b815290850135600482015260248101849052604401610cac565b612f5484602001358660200151866040013588604001516135ff565b612f6160808501856150fb565b9050600003612f8357604051631a1503a960e11b815260040160405180910390fd5b600080612f9360808701876150fb565b810190612fa091906157df565b9092509050612fbe836001602089013561288860408b013583615665565b509095945050505050565b612fd1614ab9565b612ffc84846000602086018035906040880135906060890135903390612ff7908b6150e0565b613d24565b949350505050565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290529061304e83613a9c565b600081815260208690526040902060080154909150600160a01b90046001600160401b03161561309457604051635e76f9ef60e11b815260048101829052602401610cac565b600081815260208581526040918290208551815590850151600180830191909155918501516002820155606085015160038201556080850151600482015560a0850151600582015560c0850151600682015560e085015160078201556101008501516008820180546101208801516001600160401b03908116600160a01b026001600160e01b03199092166001600160a01b03909416939093171790556101408601516009830180549190921667ffffffffffffffff19821681178355610160880151889590939268ffffffffffffffffff191690911790600160401b90849081111561318357613183614fe2565b021790555061018082810151600990920180546101a0909401511515600160501b026aff000000000000000000001960ff909416600160481b02939093166affff0000000000000000001990941693909317919091179091558301518351604085015160208601516080870151600094613202949093909290916125a5565b6000818152600187016020526040812054919250819003613261576040516815539492559053115160ba1b602082015260290160408051601f198184030181529181528151602092830120600085815260018a019093529120556132a9565b6040516815539492559053115160ba1b60208201526029016040516020818303038152906040528051906020012081036132a957600082815260018701602052604090208390555b6040805161010081018252848152602080820185905287518284015260e088015160608301526000868152908990529190912060808201906132ea906129e4565b815261018087015160ff166020820152911515604083015260e090950151151560609091015250919392505050565b600061336e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613e269092919063ffffffff16565b805190915015612296578080602001905181019061338c919061538e565b6122965760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610cac565b60008281526020859052604080822054858352912061340990611a5c565b1461345757600083815260208590526040902061342590611a5c565b6000838152602086905260409081902054905163e2e27f8760e01b815260048101929092526024820152604401610cac565b600082815260208590526040808220600990810154868452919092209091015460ff600160481b9283900481169261349192041683613b2c565b60ff16146118ea57600083815260208590526040902060090154839083906134c390600160481b900460ff1684613b2c565b60008581526020889052604090819020600901549051637e726d1560e01b81526004810194909452602484019290925260ff9081166044840152600160481b909104166064820152608401610cac565b600081815260208390526040812061352a90611a5c565b6000818152600285016020526040902054909150801561356757604051630dd7028f60e41b81526004810184905260248101829052604401610cac565b5060009081526002909201602052604090912055565b60006009820154600160401b900460ff16600181111561359f5761359f614fe2565b146135d7576135ad81611edd565b60098201546040516323f8405d60e01b8152610cac9291600160401b900460ff169060040161559d565b60090180546001600160401b03431668ffffffffffffffffff1990911617600160401b179055565b600061363482848660405160200161361991815260200190565b60405160208183030381529060405280519060200120613e35565b90508085146136855760405162461bcd60e51b815260206004820152601760248201527f496e76616c696420696e636c7573696f6e2070726f6f660000000000000000006044820152606401610cac565b5050505050565b6000600261369a84846153f6565b10156136c35760405163240a616560e21b81526004810184905260248101839052604401610cac565b6136cd83836153f6565b6002036136e6576136df836001615665565b9050610bca565b6000836136f46001856153f6565b189050600061370282613ed7565b9050600019811b806137156001876153f6565b169695505050505050565b600085116137705760405162461bcd60e51b815260206004820152601460248201527f5072652d73697a652063616e6e6f7420626520300000000000000000000000006044820152606401610cac565b8561377a83613b84565b146137c75760405162461bcd60e51b815260206004820152601b60248201527f50726520657870616e73696f6e20726f6f74206d69736d6174636800000000006044820152606401610cac565b846137d183614013565b146138285760405162461bcd60e51b815260206004820152602160248201527f5072652073697a6520646f6573206e6f74206d6174636820657870616e73696f6044820152603760f91b6064820152608401610cac565b8285106138775760405162461bcd60e51b815260206004820181905260248201527f5072652073697a65206e6f74206c657373207468616e20706f73742073697a656044820152606401610cac565b600085905060008061388c856000875161406e565b90505b8583101561394f5760006138a384886141d0565b9050845183106138f55760405162461bcd60e51b815260206004820152601260248201527f496e646578206f7574206f662072616e676500000000000000000000000000006044820152606401610cac565b613919828287868151811061390c5761390c615409565b60200260200101516142ba565b91506001811b6139298186615665565b94508785111561393b5761393b615838565b83613945816155d1565b945050505061388f565b8661395982613b84565b146139b15760405162461bcd60e51b815260206004820152602260248201527f506f737420657870616e73696f6e20726f6f74206e6f7420657175616c20706f6044820152611cdd60f21b6064820152608401610cac565b83518214613a015760405162461bcd60e51b815260206004820152601660248201527f496e636f6d706c6574652070726f6f66207573616765000000000000000000006044820152606401610cac565b505050505050505050565b613a14614ab9565b613a21878787878761481b565b50604080516101c08101825296875260208701959095529385019290925260608401526080830152600060a0830181905260c0830181905260e0830181905261010083018190526001600160401b0343166101208401526101408301819052610160830181905260ff9091166101808301526101a082015290565b6000610bca8261018001518360000151846040015185602001518660800151876060015161172c565b6005830154151580613ada5750600683015415155b15613b1c57613ae883611edd565b600584015460068501546040516308b0e71d60e41b8152600481019390935260248301919091526044820152606401610cac565b6005830191909155600690910155565b600080613b3a846001615584565b9050613b468184611771565b509392505050565b6060611e0d83600084604051602001613b6991815260200190565b604051602081830303815290604052805190602001206142ba565b600080825111613bd65760405162461bcd60e51b815260206004820152601660248201527f456d707479206d65726b6c6520657870616e73696f6e000000000000000000006044820152606401610cac565b604082511115613c285760405162461bcd60e51b815260206004820152601a60248201527f4d65726b6c6520657870616e73696f6e20746f6f206c617267650000000000006044820152606401610cac565b6000805b8351811015613d1d576000848281518110613c4957613c49615409565b60200260200101519050826000801b03613cb5578015613cb05780925060018551613c7491906153f6565b8214613cb057604051613c97908490600090602001918252602082015260400190565b6040516020818303038152906040528051906020012092505b613d0a565b8015613cd4576040805160208101839052908101849052606001613c97565b604051613cf1908490600090602001918252602082015260400190565b6040516020818303038152906040528051906020012092505b5080613d15816155d1565b915050613c2c565b5092915050565b613d2c614ab9565b6001600160a01b038316613d535760405163f289e65760e01b815260040160405180910390fd5b6000849003613d7557604051636932bcfd60e01b815260040160405180910390fd5b613d82898989898961481b565b604051806101c001604052808a81526020018981526020018881526020018781526020018681526020016000801b81526020016000801b8152602001858152602001846001600160a01b03168152602001436001600160401b0316815260200160006001600160401b0316815260200160006001811115613e0557613e05614fe2565b815260ff841660208201526000604090910152905098975050505050505050565b6060612ffc84846000856148ab565b8251600090610100811115613e6857604051637ed6198f60e11b8152600481018290526101006024820152604401610cac565b8260005b82811015613ecd576000878281518110613e8857613e88615409565b60200260200101519050816001901b8716600003613eb457826000528060205260406000209250613ec4565b8060005282602052604060002092505b50600101613e6c565b5095945050505050565b600081600003613f295760405162461bcd60e51b815260206004820152601c60248201527f5a65726f20686173206e6f207369676e69666963616e742062697473000000006044820152606401610cac565b7001000000000000000000000000000000008210613f5457608091821c91613f519082615665565b90505b600160401b8210613f7257604091821c91613f6f9082615665565b90505b6401000000008210613f9157602091821c91613f8e9082615665565b90505b620100008210613fae57601091821c91613fab9082615665565b90505b6101008210613fca57600891821c91613fc79082615665565b90505b60108210613fe557600491821c91613fe29082615665565b90505b6004821061400057600291821c91613ffd9082615665565b90505b60028210610cb557610bca600182615665565b600080805b8351811015613d1d5783818151811061403357614033615409565b60200260200101516000801b1461405c5761404f816002615932565b6140599083615665565b91505b80614066816155d1565b915050614018565b60608183106140bf5760405162461bcd60e51b815260206004820152601760248201527f5374617274206e6f74206c657373207468616e20656e640000000000000000006044820152606401610cac565b835182111561411a5760405162461bcd60e51b815260206004820152602160248201527f456e64206e6f74206c657373206f7220657175616c207468616e206c656e67746044820152600d60fb1b6064820152608401610cac565b600061412684846153f6565b6001600160401b0381111561413d5761413d614c8d565b604051908082528060200260200182016040528015614166578160200160208202803683370190505b509050835b838110156141c75785818151811061418557614185615409565b602002602001015182868361419a91906153f6565b815181106141aa576141aa615409565b6020908102919091010152806141bf816155d1565b91505061416b565b50949350505050565b60008183106142215760405162461bcd60e51b815260206004820152601760248201527f5374617274206e6f74206c657373207468616e20656e640000000000000000006044820152606401610cac565b600061422e838518613ed7565b90506000600161423e8382615665565b6001901b61424c91906153f6565b9050848116848216811561426357611dbf826149d1565b801561427257611dbf81613ed7565b60405162461bcd60e51b815260206004820152601b60248201527f426f7468207920616e64207a2063616e6e6f74206265207a65726f00000000006044820152606401610cac565b60606040831061430c5760405162461bcd60e51b815260206004820152600e60248201527f4c6576656c20746f6f20686967680000000000000000000000000000000000006044820152606401610cac565b600082900361435d5760405162461bcd60e51b815260206004820152601b60248201527f43616e6e6f7420617070656e6420656d707479207375627472656500000000006044820152606401610cac565b6040845111156143af5760405162461bcd60e51b815260206004820152601a60248201527f4d65726b6c6520657870616e73696f6e20746f6f206c617267650000000000006044820152606401610cac565b835160000361442d5760006143c5846001615665565b6001600160401b038111156143dc576143dc614c8d565b604051908082528060200260200182016040528015614405578160200160208202803683370190505b5090508281858151811061441b5761441b615409565b60209081029190910101529050611e0d565b835183106144a35760405162461bcd60e51b815260206004820152603560248201527f4c6576656c2067726561746572207468616e2068696768657374206c6576656c60448201527f206f662063757272656e7420657870616e73696f6e00000000000000000000006064820152608401610cac565b8160006144af86614013565b905060006144be866002615932565b6144c89083615665565b905060006144d583613ed7565b6144de83613ed7565b1161452b5787516001600160401b038111156144fc576144fc614c8d565b604051908082528060200260200182016040528015614525578160200160208202803683370190505b5061457a565b8751614538906001615665565b6001600160401b0381111561454f5761454f614c8d565b604051908082528060200260200182016040528015614578578160200160208202803683370190505b505b90506040815111156145ce5760405162461bcd60e51b815260206004820152601c60248201527f417070656e642063726561746573206f76657273697a652074726565000000006044820152606401610cac565b60005b885181101561476f578781101561465d578881815181106145f4576145f4615409565b60200260200101516000801b146146585760405162461bcd60e51b815260206004820152602260248201527f417070656e642061626f7665206c65617374207369676e69666963616e7420626044820152611a5d60f21b6064820152608401610cac565b61475d565b60008590036146a35788818151811061467857614678615409565b602002602001015182828151811061469257614692615409565b60200260200101818152505061475d565b8881815181106146b5576146b5615409565b60200260200101516000801b036146ed57848282815181106146d9576146d9615409565b60209081029190910101526000945061475d565b6000801b82828151811061470357614703615409565b60200260200101818152505088818151811061472157614721615409565b602002602001015185604051602001614744929190918252602082015260400190565b6040516020818303038152906040528051906020012094505b80614767816155d1565b9150506145d1565b5083156147a35783816001835161478691906153f6565b8151811061479657614796615409565b6020026020010181815250505b80600182516147b291906153f6565b815181106147c2576147c2615409565b60200260200101516000801b036105755760405162461bcd60e51b815260206004820152600f60248201527f4c61737420656e747279207a65726f00000000000000000000000000000000006044820152606401610cac565b600085900361483d5760405163235e76ef60e21b815260040160405180910390fd5b828111614867576040516308183ebd60e21b81526004810184905260248101829052604401610cac565b6000849003614889576040516320f1a0f960e21b815260040160405180910390fd5b600082900361368557604051635cb6e5bb60e01b815260040160405180910390fd5b60608247101561490c5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610cac565b6001600160a01b0385163b6149635760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610cac565b600080866001600160a01b0316858760405161497f9190615962565b60006040518083038185875af1925050503d80600081146149bc576040519150601f19603f3d011682016040523d82523d6000602084013e6149c1565b606091505b5091509150610575828286614a3e565b6000808211614a225760405162461bcd60e51b815260206004820152601c60248201527f5a65726f20686173206e6f207369676e69666963616e742062697473000000006044820152606401610cac565b60008280614a316001826153f6565b16189050611e0d81613ed7565b60608315614a4d575081611e0d565b825115614a5d5782518084602001fd5b8160405162461bcd60e51b8152600401610cac919061597e565b6040805160c081018252600080825260208201819052918101829052606081019190915260808101614aa7614b2c565b8152602001614ab4614b2c565b905290565b604080516101c081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905290610160820190815260006020820181905260409091015290565b6040518060400160405280614b3f614b4b565b81526020016000905290565b6040518060400160405280614b5e614b67565b8152602001614ab45b60405180604001604052806002906020820280368337509192915050565b803560ff81168114610cb557600080fd5b60008060008060008060c08789031215614baf57600080fd5b614bb887614b85565b9860208801359850604088013597606081013597506080810135965060a00135945092505050565b600060208284031215614bf257600080fd5b81356001600160401b03811115614c0857600080fd5b820160c08185031215611e0d57600080fd5b60008060408385031215614c2d57600080fd5b50508035926020909101359150565b600060208284031215614c4e57600080fd5b5035919050565b60038110614c6257600080fd5b50565b8035610cb581614c55565b600060208284031215614c8257600080fd5b8135611e0d81614c55565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715614cc557614cc5614c8d565b60405290565b604080519081016001600160401b0381118282101715614cc557614cc5614c8d565b604051601f8201601f191681016001600160401b0381118282101715614d1557614d15614c8d565b604052919050565b60006001600160401b03821115614d3657614d36614c8d565b5060051b60200190565b600082601f830112614d5157600080fd5b81356020614d66614d6183614d1d565b614ced565b82815260059290921b84018101918181019086841115614d8557600080fd5b8286015b84811015614da05780358352918301918301614d89565b509695505050505050565b6000806000838503610120811215614dc257600080fd5b8435935060208501356001600160401b03811115614ddf57600080fd5b614deb87828801614d40565b93505060e0603f1982011215614e0057600080fd5b506040840190509250925092565b60008083601f840112614e2057600080fd5b5081356001600160401b03811115614e3757600080fd5b6020830191508360208260051b8501011115614e5257600080fd5b9250929050565b6000806000806000806000878903610120811215614e7657600080fd5b8835975060208901356001600160401b0380821115614e9457600080fd5b908a01906040828d031215614ea857600080fd5b81985060a0603f1984011215614ebd57600080fd5b60408b01975060e08b0135925080831115614ed757600080fd5b614ee38c848d01614e0e565b90975095506101008b0135925086915080831115614f0057600080fd5b5050614f0e8a828b01614e0e565b989b979a50959850939692959293505050565b600080600080600060a08688031215614f3957600080fd5b614f4286614b85565b97602087013597506040870135966060810135965060800135945092505050565b60008060008060608587031215614f7957600080fd5b843593506020850135925060408501356001600160401b0380821115614f9e57600080fd5b818701915087601f830112614fb257600080fd5b813581811115614fc157600080fd5b886020828501011115614fd357600080fd5b95989497505060200194505050565b634e487b7160e01b600052602160045260246000fd5b6002811061500857615008614fe2565b9052565b60006101c082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015260e083015160e08301526101008084015161507a828501826001600160a01b03169052565b5050610120838101516001600160401b038116848301525050610140838101516001600160401b038116848301525050610160808401516150bd82850182614ff8565b50506101808381015160ff16908301526101a08084015180151582850152611dd7565b6000602082840312156150f257600080fd5b611e0d82614b85565b6000808335601e1984360301811261511257600080fd5b8301803591506001600160401b0382111561512c57600080fd5b602001915036819003821315614e5257600080fd5b6001600160401b0381168114614c6257600080fd5b600082601f83011261516757600080fd5b61516f614ccb565b80604084018581111561518157600080fd5b845b81811015612fbe57803561519681615141565b845260209384019301615183565b600081830360e08112156151b757600080fd5b6151bf614ca3565b915060a08112156151cf57600080fd5b6151d7614ccb565b60808212156151e557600080fd5b6151ed614ccb565b915084601f8501126151fe57600080fd5b615206614ccb565b80604086018781111561521857600080fd5b865b8181101561523257803584526020938401930161521a565b508185526152408882615156565b602086015250505081815261525760808501614c65565b6020820152808352505060a0820135602082015260c0820135604082015292915050565b60008060006101e0848603121561529157600080fd5b83356001600160401b038111156152a757600080fd5b6152b386828701614d40565b9350506152c385602086016151a4565b91506152d38561010086016151a4565b90509250925092565b60038110614c6257614c62614fe2565b615008816152dc565b805180518360005b600281101561531c5782518252602092830192909101906001016152fd565b505050602090810151906040840160005b60028110156153535783516001600160401b03168252928201929082019060010161532d565b5050820151905061229660808401826152ec565b848152610100810161537c60208301866152f5565b60c082019390935260e0015292915050565b6000602082840312156153a057600080fd5b81518015158114611e0d57600080fd5b6000602082840312156153c257600080fd5b8151611e0d81615141565b602081016153da836152dc565b91905290565b634e487b7160e01b600052601160045260246000fd5b81810381811115610bca57610bca6153e0565b634e487b7160e01b600052603260045260246000fd5b8481526101008101602060408682850137606083016040870160005b600281101561546a57813561544f81615141565b6001600160401b03168352918301919083019060010161543b565b50505050608085013561547c81614c55565b615485816152dc565b60a083015260c082019390935260e0015292915050565b6001600160401b03828116828216039080821115613d1d57613d1d6153e0565b6001600160a01b0381168114614c6257600080fd5b600060c082019050838252823560208301526020830135604083015260408301356154fb816154bc565b6001600160a01b038116606084015250606083013561551981615141565b6001600160401b0380821660808501526080850135915061553982615141565b80821660a085015250509392505050565b60006020828403121561555c57600080fd5b8135611e0d81615141565b60006020828403121561557957600080fd5b8151611e0d816154bc565b60ff8181168382160190811115610bca57610bca6153e0565b82815260408101611e0d6020830184614ff8565b6001600160401b03818116838216019080821115613d1d57613d1d6153e0565b6000600182016155e3576155e36153e0565b5060010190565b855181526001600160a01b0360208701511660208201526040860151604082015284606082015283608082015260c060a08201528160c0820152818360e0830137600081830160e090810191909152601f909201601f19160101949350505050565b60006020828403121561565e57600080fd5b5051919050565b80820180821115610bca57610bca6153e0565b600082601f83011261568957600080fd5b81516020615699614d6183614d1d565b82815260059290921b840181019181810190868411156156b857600080fd5b8286015b84811015614da057805183529183019183016156bc565b600080604083850312156156e657600080fd5b82516001600160401b03808211156156fd57600080fd5b61570986838701615678565b9350602085015191508082111561571f57600080fd5b5061572c85828601615678565b9150509250929050565b60a08101610bca82846152f5565b600080600080600060a0868803121561575c57600080fd5b853594506020860135935060408601356001600160401b038082111561578157600080fd5b61578d89838a01614d40565b945060608801359150808211156157a357600080fd5b6157af89838a01614d40565b935060808801359150808211156157c557600080fd5b506157d288828901614d40565b9150509295509295909350565b600080604083850312156157f257600080fd5b82356001600160401b038082111561580957600080fd5b61581586838701614d40565b9350602085013591508082111561582b57600080fd5b5061572c85828601614d40565b634e487b7160e01b600052600160045260246000fd5b600181815b8085111561588957816000190482111561586f5761586f6153e0565b8085161561587c57918102915b93841c9390800290615853565b509250929050565b6000826158a057506001610bca565b816158ad57506000610bca565b81600181146158c357600281146158cd576158e9565b6001915050610bca565b60ff8411156158de576158de6153e0565b50506001821b610bca565b5060208310610133831016604e8410600b841016171561590c575081810a610bca565b615916838361584e565b806000190482111561592a5761592a6153e0565b029392505050565b6000611e0d8383615891565b60005b83811015615959578181015183820152602001615941565b50506000910152565b6000825161597481846020870161593e565b9190910192915050565b602081526000825180602084015261599d81604085016020870161593e565b601f01601f1916919091016040019291505056fea2646970667358221220f02817538950bbce0611d3b3d11a69230681b7ecc3bbbd0feb00a4ef64163bd164736f6c63430008110033",
}

// EdgeChallengeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use EdgeChallengeManagerMetaData.ABI instead.
var EdgeChallengeManagerABI = EdgeChallengeManagerMetaData.ABI

// EdgeChallengeManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EdgeChallengeManagerMetaData.Bin instead.
var EdgeChallengeManagerBin = EdgeChallengeManagerMetaData.Bin

// DeployEdgeChallengeManager deploys a new Ethereum contract, binding an instance of EdgeChallengeManager to it.
func DeployEdgeChallengeManager(auth *bind.TransactOpts, backend bind.ContractBackend, _assertionChain common.Address, _challengePeriodBlocks uint64, _oneStepProofEntry common.Address, layerZeroBlockEdgeHeight *big.Int, layerZeroBigStepEdgeHeight *big.Int, layerZeroSmallStepEdgeHeight *big.Int, _stakeToken common.Address, _stakeAmount *big.Int, _excessStakeReceiver common.Address, _numBigStepLevel uint8) (common.Address, *types.Transaction, *EdgeChallengeManager, error) {
	parsed, err := EdgeChallengeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EdgeChallengeManagerBin), backend, _assertionChain, _challengePeriodBlocks, _oneStepProofEntry, layerZeroBlockEdgeHeight, layerZeroBigStepEdgeHeight, layerZeroSmallStepEdgeHeight, _stakeToken, _stakeAmount, _excessStakeReceiver, _numBigStepLevel)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EdgeChallengeManager{EdgeChallengeManagerCaller: EdgeChallengeManagerCaller{contract: contract}, EdgeChallengeManagerTransactor: EdgeChallengeManagerTransactor{contract: contract}, EdgeChallengeManagerFilterer: EdgeChallengeManagerFilterer{contract: contract}}, nil
}

// EdgeChallengeManager is an auto generated Go binding around an Ethereum contract.
type EdgeChallengeManager struct {
	EdgeChallengeManagerCaller     // Read-only binding to the contract
	EdgeChallengeManagerTransactor // Write-only binding to the contract
	EdgeChallengeManagerFilterer   // Log filterer for contract events
}

// EdgeChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EdgeChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EdgeChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EdgeChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EdgeChallengeManagerSession struct {
	Contract     *EdgeChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EdgeChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EdgeChallengeManagerCallerSession struct {
	Contract *EdgeChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// EdgeChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EdgeChallengeManagerTransactorSession struct {
	Contract     *EdgeChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// EdgeChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EdgeChallengeManagerRaw struct {
	Contract *EdgeChallengeManager // Generic contract binding to access the raw methods on
}

// EdgeChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EdgeChallengeManagerCallerRaw struct {
	Contract *EdgeChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// EdgeChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EdgeChallengeManagerTransactorRaw struct {
	Contract *EdgeChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEdgeChallengeManager creates a new instance of EdgeChallengeManager, bound to a specific deployed contract.
func NewEdgeChallengeManager(address common.Address, backend bind.ContractBackend) (*EdgeChallengeManager, error) {
	contract, err := bindEdgeChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManager{EdgeChallengeManagerCaller: EdgeChallengeManagerCaller{contract: contract}, EdgeChallengeManagerTransactor: EdgeChallengeManagerTransactor{contract: contract}, EdgeChallengeManagerFilterer: EdgeChallengeManagerFilterer{contract: contract}}, nil
}

// NewEdgeChallengeManagerCaller creates a new read-only instance of EdgeChallengeManager, bound to a specific deployed contract.
func NewEdgeChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*EdgeChallengeManagerCaller, error) {
	contract, err := bindEdgeChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManagerCaller{contract: contract}, nil
}

// NewEdgeChallengeManagerTransactor creates a new write-only instance of EdgeChallengeManager, bound to a specific deployed contract.
func NewEdgeChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*EdgeChallengeManagerTransactor, error) {
	contract, err := bindEdgeChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManagerTransactor{contract: contract}, nil
}

// NewEdgeChallengeManagerFilterer creates a new log filterer instance of EdgeChallengeManager, bound to a specific deployed contract.
func NewEdgeChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*EdgeChallengeManagerFilterer, error) {
	contract, err := bindEdgeChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManagerFilterer{contract: contract}, nil
}

// bindEdgeChallengeManager binds a generic wrapper to an already deployed contract.
func bindEdgeChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EdgeChallengeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EdgeChallengeManager *EdgeChallengeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EdgeChallengeManager.Contract.EdgeChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EdgeChallengeManager *EdgeChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.EdgeChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EdgeChallengeManager *EdgeChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.EdgeChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EdgeChallengeManager *EdgeChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EdgeChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EdgeChallengeManager *EdgeChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EdgeChallengeManager *EdgeChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// LAYERZEROBIGSTEPEDGEHEIGHT is a free data retrieval call binding the contract method 0x416e6657.
//
// Solidity: function LAYERZERO_BIGSTEPEDGE_HEIGHT() view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) LAYERZEROBIGSTEPEDGEHEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "LAYERZERO_BIGSTEPEDGE_HEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LAYERZEROBIGSTEPEDGEHEIGHT is a free data retrieval call binding the contract method 0x416e6657.
//
// Solidity: function LAYERZERO_BIGSTEPEDGE_HEIGHT() view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) LAYERZEROBIGSTEPEDGEHEIGHT() (*big.Int, error) {
	return _EdgeChallengeManager.Contract.LAYERZEROBIGSTEPEDGEHEIGHT(&_EdgeChallengeManager.CallOpts)
}

// LAYERZEROBIGSTEPEDGEHEIGHT is a free data retrieval call binding the contract method 0x416e6657.
//
// Solidity: function LAYERZERO_BIGSTEPEDGE_HEIGHT() view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) LAYERZEROBIGSTEPEDGEHEIGHT() (*big.Int, error) {
	return _EdgeChallengeManager.Contract.LAYERZEROBIGSTEPEDGEHEIGHT(&_EdgeChallengeManager.CallOpts)
}

// LAYERZEROBLOCKEDGEHEIGHT is a free data retrieval call binding the contract method 0x1dce5166.
//
// Solidity: function LAYERZERO_BLOCKEDGE_HEIGHT() view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) LAYERZEROBLOCKEDGEHEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "LAYERZERO_BLOCKEDGE_HEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LAYERZEROBLOCKEDGEHEIGHT is a free data retrieval call binding the contract method 0x1dce5166.
//
// Solidity: function LAYERZERO_BLOCKEDGE_HEIGHT() view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) LAYERZEROBLOCKEDGEHEIGHT() (*big.Int, error) {
	return _EdgeChallengeManager.Contract.LAYERZEROBLOCKEDGEHEIGHT(&_EdgeChallengeManager.CallOpts)
}

// LAYERZEROBLOCKEDGEHEIGHT is a free data retrieval call binding the contract method 0x1dce5166.
//
// Solidity: function LAYERZERO_BLOCKEDGE_HEIGHT() view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) LAYERZEROBLOCKEDGEHEIGHT() (*big.Int, error) {
	return _EdgeChallengeManager.Contract.LAYERZEROBLOCKEDGEHEIGHT(&_EdgeChallengeManager.CallOpts)
}

// LAYERZEROSMALLSTEPEDGEHEIGHT is a free data retrieval call binding the contract method 0xf8ee77d6.
//
// Solidity: function LAYERZERO_SMALLSTEPEDGE_HEIGHT() view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) LAYERZEROSMALLSTEPEDGEHEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "LAYERZERO_SMALLSTEPEDGE_HEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LAYERZEROSMALLSTEPEDGEHEIGHT is a free data retrieval call binding the contract method 0xf8ee77d6.
//
// Solidity: function LAYERZERO_SMALLSTEPEDGE_HEIGHT() view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) LAYERZEROSMALLSTEPEDGEHEIGHT() (*big.Int, error) {
	return _EdgeChallengeManager.Contract.LAYERZEROSMALLSTEPEDGEHEIGHT(&_EdgeChallengeManager.CallOpts)
}

// LAYERZEROSMALLSTEPEDGEHEIGHT is a free data retrieval call binding the contract method 0xf8ee77d6.
//
// Solidity: function LAYERZERO_SMALLSTEPEDGE_HEIGHT() view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) LAYERZEROSMALLSTEPEDGEHEIGHT() (*big.Int, error) {
	return _EdgeChallengeManager.Contract.LAYERZEROSMALLSTEPEDGEHEIGHT(&_EdgeChallengeManager.CallOpts)
}

// NUMBIGSTEPLEVEL is a free data retrieval call binding the contract method 0x5d9e2444.
//
// Solidity: function NUM_BIGSTEP_LEVEL() view returns(uint8)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) NUMBIGSTEPLEVEL(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "NUM_BIGSTEP_LEVEL")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NUMBIGSTEPLEVEL is a free data retrieval call binding the contract method 0x5d9e2444.
//
// Solidity: function NUM_BIGSTEP_LEVEL() view returns(uint8)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) NUMBIGSTEPLEVEL() (uint8, error) {
	return _EdgeChallengeManager.Contract.NUMBIGSTEPLEVEL(&_EdgeChallengeManager.CallOpts)
}

// NUMBIGSTEPLEVEL is a free data retrieval call binding the contract method 0x5d9e2444.
//
// Solidity: function NUM_BIGSTEP_LEVEL() view returns(uint8)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) NUMBIGSTEPLEVEL() (uint8, error) {
	return _EdgeChallengeManager.Contract.NUMBIGSTEPLEVEL(&_EdgeChallengeManager.CallOpts)
}

// AssertionChain is a free data retrieval call binding the contract method 0x48dd2924.
//
// Solidity: function assertionChain() view returns(address)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) AssertionChain(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "assertionChain")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssertionChain is a free data retrieval call binding the contract method 0x48dd2924.
//
// Solidity: function assertionChain() view returns(address)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) AssertionChain() (common.Address, error) {
	return _EdgeChallengeManager.Contract.AssertionChain(&_EdgeChallengeManager.CallOpts)
}

// AssertionChain is a free data retrieval call binding the contract method 0x48dd2924.
//
// Solidity: function assertionChain() view returns(address)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) AssertionChain() (common.Address, error) {
	return _EdgeChallengeManager.Contract.AssertionChain(&_EdgeChallengeManager.CallOpts)
}

// CalculateEdgeId is a free data retrieval call binding the contract method 0x004d8efe.
//
// Solidity: function calculateEdgeId(uint8 level, bytes32 originId, uint256 startHeight, bytes32 startHistoryRoot, uint256 endHeight, bytes32 endHistoryRoot) pure returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) CalculateEdgeId(opts *bind.CallOpts, level uint8, originId [32]byte, startHeight *big.Int, startHistoryRoot [32]byte, endHeight *big.Int, endHistoryRoot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "calculateEdgeId", level, originId, startHeight, startHistoryRoot, endHeight, endHistoryRoot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateEdgeId is a free data retrieval call binding the contract method 0x004d8efe.
//
// Solidity: function calculateEdgeId(uint8 level, bytes32 originId, uint256 startHeight, bytes32 startHistoryRoot, uint256 endHeight, bytes32 endHistoryRoot) pure returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) CalculateEdgeId(level uint8, originId [32]byte, startHeight *big.Int, startHistoryRoot [32]byte, endHeight *big.Int, endHistoryRoot [32]byte) ([32]byte, error) {
	return _EdgeChallengeManager.Contract.CalculateEdgeId(&_EdgeChallengeManager.CallOpts, level, originId, startHeight, startHistoryRoot, endHeight, endHistoryRoot)
}

// CalculateEdgeId is a free data retrieval call binding the contract method 0x004d8efe.
//
// Solidity: function calculateEdgeId(uint8 level, bytes32 originId, uint256 startHeight, bytes32 startHistoryRoot, uint256 endHeight, bytes32 endHistoryRoot) pure returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) CalculateEdgeId(level uint8, originId [32]byte, startHeight *big.Int, startHistoryRoot [32]byte, endHeight *big.Int, endHistoryRoot [32]byte) ([32]byte, error) {
	return _EdgeChallengeManager.Contract.CalculateEdgeId(&_EdgeChallengeManager.CallOpts, level, originId, startHeight, startHistoryRoot, endHeight, endHistoryRoot)
}

// CalculateMutualId is a free data retrieval call binding the contract method 0xc32d8c63.
//
// Solidity: function calculateMutualId(uint8 level, bytes32 originId, uint256 startHeight, bytes32 startHistoryRoot, uint256 endHeight) pure returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) CalculateMutualId(opts *bind.CallOpts, level uint8, originId [32]byte, startHeight *big.Int, startHistoryRoot [32]byte, endHeight *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "calculateMutualId", level, originId, startHeight, startHistoryRoot, endHeight)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMutualId is a free data retrieval call binding the contract method 0xc32d8c63.
//
// Solidity: function calculateMutualId(uint8 level, bytes32 originId, uint256 startHeight, bytes32 startHistoryRoot, uint256 endHeight) pure returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) CalculateMutualId(level uint8, originId [32]byte, startHeight *big.Int, startHistoryRoot [32]byte, endHeight *big.Int) ([32]byte, error) {
	return _EdgeChallengeManager.Contract.CalculateMutualId(&_EdgeChallengeManager.CallOpts, level, originId, startHeight, startHistoryRoot, endHeight)
}

// CalculateMutualId is a free data retrieval call binding the contract method 0xc32d8c63.
//
// Solidity: function calculateMutualId(uint8 level, bytes32 originId, uint256 startHeight, bytes32 startHistoryRoot, uint256 endHeight) pure returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) CalculateMutualId(level uint8, originId [32]byte, startHeight *big.Int, startHistoryRoot [32]byte, endHeight *big.Int) ([32]byte, error) {
	return _EdgeChallengeManager.Contract.CalculateMutualId(&_EdgeChallengeManager.CallOpts, level, originId, startHeight, startHistoryRoot, endHeight)
}

// ChallengePeriodBlocks is a free data retrieval call binding the contract method 0x46c2781a.
//
// Solidity: function challengePeriodBlocks() view returns(uint64)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) ChallengePeriodBlocks(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "challengePeriodBlocks")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChallengePeriodBlocks is a free data retrieval call binding the contract method 0x46c2781a.
//
// Solidity: function challengePeriodBlocks() view returns(uint64)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) ChallengePeriodBlocks() (uint64, error) {
	return _EdgeChallengeManager.Contract.ChallengePeriodBlocks(&_EdgeChallengeManager.CallOpts)
}

// ChallengePeriodBlocks is a free data retrieval call binding the contract method 0x46c2781a.
//
// Solidity: function challengePeriodBlocks() view returns(uint64)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) ChallengePeriodBlocks() (uint64, error) {
	return _EdgeChallengeManager.Contract.ChallengePeriodBlocks(&_EdgeChallengeManager.CallOpts)
}

// ConfirmedRival is a free data retrieval call binding the contract method 0xe5b123da.
//
// Solidity: function confirmedRival(bytes32 mutualId) view returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) ConfirmedRival(opts *bind.CallOpts, mutualId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "confirmedRival", mutualId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConfirmedRival is a free data retrieval call binding the contract method 0xe5b123da.
//
// Solidity: function confirmedRival(bytes32 mutualId) view returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) ConfirmedRival(mutualId [32]byte) ([32]byte, error) {
	return _EdgeChallengeManager.Contract.ConfirmedRival(&_EdgeChallengeManager.CallOpts, mutualId)
}

// ConfirmedRival is a free data retrieval call binding the contract method 0xe5b123da.
//
// Solidity: function confirmedRival(bytes32 mutualId) view returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) ConfirmedRival(mutualId [32]byte) ([32]byte, error) {
	return _EdgeChallengeManager.Contract.ConfirmedRival(&_EdgeChallengeManager.CallOpts, mutualId)
}

// EdgeExists is a free data retrieval call binding the contract method 0x750e0c0f.
//
// Solidity: function edgeExists(bytes32 edgeId) view returns(bool)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) EdgeExists(opts *bind.CallOpts, edgeId [32]byte) (bool, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "edgeExists", edgeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EdgeExists is a free data retrieval call binding the contract method 0x750e0c0f.
//
// Solidity: function edgeExists(bytes32 edgeId) view returns(bool)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) EdgeExists(edgeId [32]byte) (bool, error) {
	return _EdgeChallengeManager.Contract.EdgeExists(&_EdgeChallengeManager.CallOpts, edgeId)
}

// EdgeExists is a free data retrieval call binding the contract method 0x750e0c0f.
//
// Solidity: function edgeExists(bytes32 edgeId) view returns(bool)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) EdgeExists(edgeId [32]byte) (bool, error) {
	return _EdgeChallengeManager.Contract.EdgeExists(&_EdgeChallengeManager.CallOpts, edgeId)
}

// EdgeLength is a free data retrieval call binding the contract method 0xeae0328b.
//
// Solidity: function edgeLength(bytes32 edgeId) view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) EdgeLength(opts *bind.CallOpts, edgeId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "edgeLength", edgeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EdgeLength is a free data retrieval call binding the contract method 0xeae0328b.
//
// Solidity: function edgeLength(bytes32 edgeId) view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) EdgeLength(edgeId [32]byte) (*big.Int, error) {
	return _EdgeChallengeManager.Contract.EdgeLength(&_EdgeChallengeManager.CallOpts, edgeId)
}

// EdgeLength is a free data retrieval call binding the contract method 0xeae0328b.
//
// Solidity: function edgeLength(bytes32 edgeId) view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) EdgeLength(edgeId [32]byte) (*big.Int, error) {
	return _EdgeChallengeManager.Contract.EdgeLength(&_EdgeChallengeManager.CallOpts, edgeId)
}

// ExcessStakeReceiver is a free data retrieval call binding the contract method 0xe94e051e.
//
// Solidity: function excessStakeReceiver() view returns(address)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) ExcessStakeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "excessStakeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExcessStakeReceiver is a free data retrieval call binding the contract method 0xe94e051e.
//
// Solidity: function excessStakeReceiver() view returns(address)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) ExcessStakeReceiver() (common.Address, error) {
	return _EdgeChallengeManager.Contract.ExcessStakeReceiver(&_EdgeChallengeManager.CallOpts)
}

// ExcessStakeReceiver is a free data retrieval call binding the contract method 0xe94e051e.
//
// Solidity: function excessStakeReceiver() view returns(address)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) ExcessStakeReceiver() (common.Address, error) {
	return _EdgeChallengeManager.Contract.ExcessStakeReceiver(&_EdgeChallengeManager.CallOpts)
}

// FirstRival is a free data retrieval call binding the contract method 0xbce6f54f.
//
// Solidity: function firstRival(bytes32 mutualId) view returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) FirstRival(opts *bind.CallOpts, mutualId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "firstRival", mutualId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FirstRival is a free data retrieval call binding the contract method 0xbce6f54f.
//
// Solidity: function firstRival(bytes32 mutualId) view returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) FirstRival(mutualId [32]byte) ([32]byte, error) {
	return _EdgeChallengeManager.Contract.FirstRival(&_EdgeChallengeManager.CallOpts, mutualId)
}

// FirstRival is a free data retrieval call binding the contract method 0xbce6f54f.
//
// Solidity: function firstRival(bytes32 mutualId) view returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) FirstRival(mutualId [32]byte) ([32]byte, error) {
	return _EdgeChallengeManager.Contract.FirstRival(&_EdgeChallengeManager.CallOpts, mutualId)
}

// GetEdge is a free data retrieval call binding the contract method 0xfda2892e.
//
// Solidity: function getEdge(bytes32 edgeId) view returns((bytes32,bytes32,uint256,bytes32,uint256,bytes32,bytes32,bytes32,address,uint64,uint64,uint8,uint8,bool))
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) GetEdge(opts *bind.CallOpts, edgeId [32]byte) (ChallengeEdge, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "getEdge", edgeId)

	if err != nil {
		return *new(ChallengeEdge), err
	}

	out0 := *abi.ConvertType(out[0], new(ChallengeEdge)).(*ChallengeEdge)

	return out0, err

}

// GetEdge is a free data retrieval call binding the contract method 0xfda2892e.
//
// Solidity: function getEdge(bytes32 edgeId) view returns((bytes32,bytes32,uint256,bytes32,uint256,bytes32,bytes32,bytes32,address,uint64,uint64,uint8,uint8,bool))
func (_EdgeChallengeManager *EdgeChallengeManagerSession) GetEdge(edgeId [32]byte) (ChallengeEdge, error) {
	return _EdgeChallengeManager.Contract.GetEdge(&_EdgeChallengeManager.CallOpts, edgeId)
}

// GetEdge is a free data retrieval call binding the contract method 0xfda2892e.
//
// Solidity: function getEdge(bytes32 edgeId) view returns((bytes32,bytes32,uint256,bytes32,uint256,bytes32,bytes32,bytes32,address,uint64,uint64,uint8,uint8,bool))
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) GetEdge(edgeId [32]byte) (ChallengeEdge, error) {
	return _EdgeChallengeManager.Contract.GetEdge(&_EdgeChallengeManager.CallOpts, edgeId)
}

// GetLayerZeroEndHeight is a free data retrieval call binding the contract method 0x42e1aaa8.
//
// Solidity: function getLayerZeroEndHeight(uint8 eType) view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) GetLayerZeroEndHeight(opts *bind.CallOpts, eType uint8) (*big.Int, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "getLayerZeroEndHeight", eType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLayerZeroEndHeight is a free data retrieval call binding the contract method 0x42e1aaa8.
//
// Solidity: function getLayerZeroEndHeight(uint8 eType) view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) GetLayerZeroEndHeight(eType uint8) (*big.Int, error) {
	return _EdgeChallengeManager.Contract.GetLayerZeroEndHeight(&_EdgeChallengeManager.CallOpts, eType)
}

// GetLayerZeroEndHeight is a free data retrieval call binding the contract method 0x42e1aaa8.
//
// Solidity: function getLayerZeroEndHeight(uint8 eType) view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) GetLayerZeroEndHeight(eType uint8) (*big.Int, error) {
	return _EdgeChallengeManager.Contract.GetLayerZeroEndHeight(&_EdgeChallengeManager.CallOpts, eType)
}

// GetPrevAssertionHash is a free data retrieval call binding the contract method 0x5a48e0f4.
//
// Solidity: function getPrevAssertionHash(bytes32 edgeId) view returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) GetPrevAssertionHash(opts *bind.CallOpts, edgeId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "getPrevAssertionHash", edgeId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPrevAssertionHash is a free data retrieval call binding the contract method 0x5a48e0f4.
//
// Solidity: function getPrevAssertionHash(bytes32 edgeId) view returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) GetPrevAssertionHash(edgeId [32]byte) ([32]byte, error) {
	return _EdgeChallengeManager.Contract.GetPrevAssertionHash(&_EdgeChallengeManager.CallOpts, edgeId)
}

// GetPrevAssertionHash is a free data retrieval call binding the contract method 0x5a48e0f4.
//
// Solidity: function getPrevAssertionHash(bytes32 edgeId) view returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) GetPrevAssertionHash(edgeId [32]byte) ([32]byte, error) {
	return _EdgeChallengeManager.Contract.GetPrevAssertionHash(&_EdgeChallengeManager.CallOpts, edgeId)
}

// HasLengthOneRival is a free data retrieval call binding the contract method 0x54b64151.
//
// Solidity: function hasLengthOneRival(bytes32 edgeId) view returns(bool)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) HasLengthOneRival(opts *bind.CallOpts, edgeId [32]byte) (bool, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "hasLengthOneRival", edgeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasLengthOneRival is a free data retrieval call binding the contract method 0x54b64151.
//
// Solidity: function hasLengthOneRival(bytes32 edgeId) view returns(bool)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) HasLengthOneRival(edgeId [32]byte) (bool, error) {
	return _EdgeChallengeManager.Contract.HasLengthOneRival(&_EdgeChallengeManager.CallOpts, edgeId)
}

// HasLengthOneRival is a free data retrieval call binding the contract method 0x54b64151.
//
// Solidity: function hasLengthOneRival(bytes32 edgeId) view returns(bool)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) HasLengthOneRival(edgeId [32]byte) (bool, error) {
	return _EdgeChallengeManager.Contract.HasLengthOneRival(&_EdgeChallengeManager.CallOpts, edgeId)
}

// HasRival is a free data retrieval call binding the contract method 0x908517e9.
//
// Solidity: function hasRival(bytes32 edgeId) view returns(bool)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) HasRival(opts *bind.CallOpts, edgeId [32]byte) (bool, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "hasRival", edgeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRival is a free data retrieval call binding the contract method 0x908517e9.
//
// Solidity: function hasRival(bytes32 edgeId) view returns(bool)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) HasRival(edgeId [32]byte) (bool, error) {
	return _EdgeChallengeManager.Contract.HasRival(&_EdgeChallengeManager.CallOpts, edgeId)
}

// HasRival is a free data retrieval call binding the contract method 0x908517e9.
//
// Solidity: function hasRival(bytes32 edgeId) view returns(bool)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) HasRival(edgeId [32]byte) (bool, error) {
	return _EdgeChallengeManager.Contract.HasRival(&_EdgeChallengeManager.CallOpts, edgeId)
}

// OneStepProofEntry is a free data retrieval call binding the contract method 0x48923bc5.
//
// Solidity: function oneStepProofEntry() view returns(address)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) OneStepProofEntry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "oneStepProofEntry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OneStepProofEntry is a free data retrieval call binding the contract method 0x48923bc5.
//
// Solidity: function oneStepProofEntry() view returns(address)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) OneStepProofEntry() (common.Address, error) {
	return _EdgeChallengeManager.Contract.OneStepProofEntry(&_EdgeChallengeManager.CallOpts)
}

// OneStepProofEntry is a free data retrieval call binding the contract method 0x48923bc5.
//
// Solidity: function oneStepProofEntry() view returns(address)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) OneStepProofEntry() (common.Address, error) {
	return _EdgeChallengeManager.Contract.OneStepProofEntry(&_EdgeChallengeManager.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) StakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "stakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) StakeAmount() (*big.Int, error) {
	return _EdgeChallengeManager.Contract.StakeAmount(&_EdgeChallengeManager.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) StakeAmount() (*big.Int, error) {
	return _EdgeChallengeManager.Contract.StakeAmount(&_EdgeChallengeManager.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) StakeToken() (common.Address, error) {
	return _EdgeChallengeManager.Contract.StakeToken(&_EdgeChallengeManager.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) StakeToken() (common.Address, error) {
	return _EdgeChallengeManager.Contract.StakeToken(&_EdgeChallengeManager.CallOpts)
}

// TimeUnrivaled is a free data retrieval call binding the contract method 0x3e35f5e8.
//
// Solidity: function timeUnrivaled(bytes32 edgeId) view returns(uint64)
func (_EdgeChallengeManager *EdgeChallengeManagerCaller) TimeUnrivaled(opts *bind.CallOpts, edgeId [32]byte) (uint64, error) {
	var out []interface{}
	err := _EdgeChallengeManager.contract.Call(opts, &out, "timeUnrivaled", edgeId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TimeUnrivaled is a free data retrieval call binding the contract method 0x3e35f5e8.
//
// Solidity: function timeUnrivaled(bytes32 edgeId) view returns(uint64)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) TimeUnrivaled(edgeId [32]byte) (uint64, error) {
	return _EdgeChallengeManager.Contract.TimeUnrivaled(&_EdgeChallengeManager.CallOpts, edgeId)
}

// TimeUnrivaled is a free data retrieval call binding the contract method 0x3e35f5e8.
//
// Solidity: function timeUnrivaled(bytes32 edgeId) view returns(uint64)
func (_EdgeChallengeManager *EdgeChallengeManagerCallerSession) TimeUnrivaled(edgeId [32]byte) (uint64, error) {
	return _EdgeChallengeManager.Contract.TimeUnrivaled(&_EdgeChallengeManager.CallOpts, edgeId)
}

// BisectEdge is a paid mutator transaction binding the contract method 0xc8bc4e43.
//
// Solidity: function bisectEdge(bytes32 edgeId, bytes32 bisectionHistoryRoot, bytes prefixProof) returns(bytes32, bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerTransactor) BisectEdge(opts *bind.TransactOpts, edgeId [32]byte, bisectionHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.contract.Transact(opts, "bisectEdge", edgeId, bisectionHistoryRoot, prefixProof)
}

// BisectEdge is a paid mutator transaction binding the contract method 0xc8bc4e43.
//
// Solidity: function bisectEdge(bytes32 edgeId, bytes32 bisectionHistoryRoot, bytes prefixProof) returns(bytes32, bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) BisectEdge(edgeId [32]byte, bisectionHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.BisectEdge(&_EdgeChallengeManager.TransactOpts, edgeId, bisectionHistoryRoot, prefixProof)
}

// BisectEdge is a paid mutator transaction binding the contract method 0xc8bc4e43.
//
// Solidity: function bisectEdge(bytes32 edgeId, bytes32 bisectionHistoryRoot, bytes prefixProof) returns(bytes32, bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerTransactorSession) BisectEdge(edgeId [32]byte, bisectionHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.BisectEdge(&_EdgeChallengeManager.TransactOpts, edgeId, bisectionHistoryRoot, prefixProof)
}

// ConfirmEdgeByChildren is a paid mutator transaction binding the contract method 0x2eaa0043.
//
// Solidity: function confirmEdgeByChildren(bytes32 edgeId) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerTransactor) ConfirmEdgeByChildren(opts *bind.TransactOpts, edgeId [32]byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.contract.Transact(opts, "confirmEdgeByChildren", edgeId)
}

// ConfirmEdgeByChildren is a paid mutator transaction binding the contract method 0x2eaa0043.
//
// Solidity: function confirmEdgeByChildren(bytes32 edgeId) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerSession) ConfirmEdgeByChildren(edgeId [32]byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.ConfirmEdgeByChildren(&_EdgeChallengeManager.TransactOpts, edgeId)
}

// ConfirmEdgeByChildren is a paid mutator transaction binding the contract method 0x2eaa0043.
//
// Solidity: function confirmEdgeByChildren(bytes32 edgeId) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerTransactorSession) ConfirmEdgeByChildren(edgeId [32]byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.ConfirmEdgeByChildren(&_EdgeChallengeManager.TransactOpts, edgeId)
}

// ConfirmEdgeByClaim is a paid mutator transaction binding the contract method 0x0f73bfad.
//
// Solidity: function confirmEdgeByClaim(bytes32 edgeId, bytes32 claimingEdgeId) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerTransactor) ConfirmEdgeByClaim(opts *bind.TransactOpts, edgeId [32]byte, claimingEdgeId [32]byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.contract.Transact(opts, "confirmEdgeByClaim", edgeId, claimingEdgeId)
}

// ConfirmEdgeByClaim is a paid mutator transaction binding the contract method 0x0f73bfad.
//
// Solidity: function confirmEdgeByClaim(bytes32 edgeId, bytes32 claimingEdgeId) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerSession) ConfirmEdgeByClaim(edgeId [32]byte, claimingEdgeId [32]byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.ConfirmEdgeByClaim(&_EdgeChallengeManager.TransactOpts, edgeId, claimingEdgeId)
}

// ConfirmEdgeByClaim is a paid mutator transaction binding the contract method 0x0f73bfad.
//
// Solidity: function confirmEdgeByClaim(bytes32 edgeId, bytes32 claimingEdgeId) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerTransactorSession) ConfirmEdgeByClaim(edgeId [32]byte, claimingEdgeId [32]byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.ConfirmEdgeByClaim(&_EdgeChallengeManager.TransactOpts, edgeId, claimingEdgeId)
}

// ConfirmEdgeByOneStepProof is a paid mutator transaction binding the contract method 0x8c1b3a40.
//
// Solidity: function confirmEdgeByOneStepProof(bytes32 edgeId, (bytes32,bytes) oneStepData, (bytes32,uint256,address,uint64,uint64) prevConfig, bytes32[] beforeHistoryInclusionProof, bytes32[] afterHistoryInclusionProof) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerTransactor) ConfirmEdgeByOneStepProof(opts *bind.TransactOpts, edgeId [32]byte, oneStepData OneStepData, prevConfig ConfigData, beforeHistoryInclusionProof [][32]byte, afterHistoryInclusionProof [][32]byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.contract.Transact(opts, "confirmEdgeByOneStepProof", edgeId, oneStepData, prevConfig, beforeHistoryInclusionProof, afterHistoryInclusionProof)
}

// ConfirmEdgeByOneStepProof is a paid mutator transaction binding the contract method 0x8c1b3a40.
//
// Solidity: function confirmEdgeByOneStepProof(bytes32 edgeId, (bytes32,bytes) oneStepData, (bytes32,uint256,address,uint64,uint64) prevConfig, bytes32[] beforeHistoryInclusionProof, bytes32[] afterHistoryInclusionProof) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerSession) ConfirmEdgeByOneStepProof(edgeId [32]byte, oneStepData OneStepData, prevConfig ConfigData, beforeHistoryInclusionProof [][32]byte, afterHistoryInclusionProof [][32]byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.ConfirmEdgeByOneStepProof(&_EdgeChallengeManager.TransactOpts, edgeId, oneStepData, prevConfig, beforeHistoryInclusionProof, afterHistoryInclusionProof)
}

// ConfirmEdgeByOneStepProof is a paid mutator transaction binding the contract method 0x8c1b3a40.
//
// Solidity: function confirmEdgeByOneStepProof(bytes32 edgeId, (bytes32,bytes) oneStepData, (bytes32,uint256,address,uint64,uint64) prevConfig, bytes32[] beforeHistoryInclusionProof, bytes32[] afterHistoryInclusionProof) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerTransactorSession) ConfirmEdgeByOneStepProof(edgeId [32]byte, oneStepData OneStepData, prevConfig ConfigData, beforeHistoryInclusionProof [][32]byte, afterHistoryInclusionProof [][32]byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.ConfirmEdgeByOneStepProof(&_EdgeChallengeManager.TransactOpts, edgeId, oneStepData, prevConfig, beforeHistoryInclusionProof, afterHistoryInclusionProof)
}

// ConfirmEdgeByTime is a paid mutator transaction binding the contract method 0x64deed59.
//
// Solidity: function confirmEdgeByTime(bytes32 edgeId, bytes32[] ancestorEdges, (((bytes32[2],uint64[2]),uint8),bytes32,bytes32) claimStateData) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerTransactor) ConfirmEdgeByTime(opts *bind.TransactOpts, edgeId [32]byte, ancestorEdges [][32]byte, claimStateData ExecutionStateData) (*types.Transaction, error) {
	return _EdgeChallengeManager.contract.Transact(opts, "confirmEdgeByTime", edgeId, ancestorEdges, claimStateData)
}

// ConfirmEdgeByTime is a paid mutator transaction binding the contract method 0x64deed59.
//
// Solidity: function confirmEdgeByTime(bytes32 edgeId, bytes32[] ancestorEdges, (((bytes32[2],uint64[2]),uint8),bytes32,bytes32) claimStateData) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerSession) ConfirmEdgeByTime(edgeId [32]byte, ancestorEdges [][32]byte, claimStateData ExecutionStateData) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.ConfirmEdgeByTime(&_EdgeChallengeManager.TransactOpts, edgeId, ancestorEdges, claimStateData)
}

// ConfirmEdgeByTime is a paid mutator transaction binding the contract method 0x64deed59.
//
// Solidity: function confirmEdgeByTime(bytes32 edgeId, bytes32[] ancestorEdges, (((bytes32[2],uint64[2]),uint8),bytes32,bytes32) claimStateData) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerTransactorSession) ConfirmEdgeByTime(edgeId [32]byte, ancestorEdges [][32]byte, claimStateData ExecutionStateData) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.ConfirmEdgeByTime(&_EdgeChallengeManager.TransactOpts, edgeId, ancestorEdges, claimStateData)
}

// CreateLayerZeroEdge is a paid mutator transaction binding the contract method 0x05fae141.
//
// Solidity: function createLayerZeroEdge((uint8,bytes32,uint256,bytes32,bytes,bytes) args) returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerTransactor) CreateLayerZeroEdge(opts *bind.TransactOpts, args CreateEdgeArgs) (*types.Transaction, error) {
	return _EdgeChallengeManager.contract.Transact(opts, "createLayerZeroEdge", args)
}

// CreateLayerZeroEdge is a paid mutator transaction binding the contract method 0x05fae141.
//
// Solidity: function createLayerZeroEdge((uint8,bytes32,uint256,bytes32,bytes,bytes) args) returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerSession) CreateLayerZeroEdge(args CreateEdgeArgs) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.CreateLayerZeroEdge(&_EdgeChallengeManager.TransactOpts, args)
}

// CreateLayerZeroEdge is a paid mutator transaction binding the contract method 0x05fae141.
//
// Solidity: function createLayerZeroEdge((uint8,bytes32,uint256,bytes32,bytes,bytes) args) returns(bytes32)
func (_EdgeChallengeManager *EdgeChallengeManagerTransactorSession) CreateLayerZeroEdge(args CreateEdgeArgs) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.CreateLayerZeroEdge(&_EdgeChallengeManager.TransactOpts, args)
}

// RefundStake is a paid mutator transaction binding the contract method 0x748926f3.
//
// Solidity: function refundStake(bytes32 edgeId) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerTransactor) RefundStake(opts *bind.TransactOpts, edgeId [32]byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.contract.Transact(opts, "refundStake", edgeId)
}

// RefundStake is a paid mutator transaction binding the contract method 0x748926f3.
//
// Solidity: function refundStake(bytes32 edgeId) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerSession) RefundStake(edgeId [32]byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.RefundStake(&_EdgeChallengeManager.TransactOpts, edgeId)
}

// RefundStake is a paid mutator transaction binding the contract method 0x748926f3.
//
// Solidity: function refundStake(bytes32 edgeId) returns()
func (_EdgeChallengeManager *EdgeChallengeManagerTransactorSession) RefundStake(edgeId [32]byte) (*types.Transaction, error) {
	return _EdgeChallengeManager.Contract.RefundStake(&_EdgeChallengeManager.TransactOpts, edgeId)
}

// EdgeChallengeManagerEdgeAddedIterator is returned from FilterEdgeAdded and is used to iterate over the raw logs and unpacked data for EdgeAdded events raised by the EdgeChallengeManager contract.
type EdgeChallengeManagerEdgeAddedIterator struct {
	Event *EdgeChallengeManagerEdgeAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EdgeChallengeManagerEdgeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeChallengeManagerEdgeAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EdgeChallengeManagerEdgeAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EdgeChallengeManagerEdgeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeChallengeManagerEdgeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeChallengeManagerEdgeAdded represents a EdgeAdded event raised by the EdgeChallengeManager contract.
type EdgeChallengeManagerEdgeAdded struct {
	EdgeId      [32]byte
	MutualId    [32]byte
	OriginId    [32]byte
	ClaimId     [32]byte
	Length      *big.Int
	Level       uint8
	HasRival    bool
	IsLayerZero bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEdgeAdded is a free log retrieval operation binding the contract event 0xaa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a4.
//
// Solidity: event EdgeAdded(bytes32 indexed edgeId, bytes32 indexed mutualId, bytes32 indexed originId, bytes32 claimId, uint256 length, uint8 level, bool hasRival, bool isLayerZero)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) FilterEdgeAdded(opts *bind.FilterOpts, edgeId [][32]byte, mutualId [][32]byte, originId [][32]byte) (*EdgeChallengeManagerEdgeAddedIterator, error) {

	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}
	var mutualIdRule []interface{}
	for _, mutualIdItem := range mutualId {
		mutualIdRule = append(mutualIdRule, mutualIdItem)
	}
	var originIdRule []interface{}
	for _, originIdItem := range originId {
		originIdRule = append(originIdRule, originIdItem)
	}

	logs, sub, err := _EdgeChallengeManager.contract.FilterLogs(opts, "EdgeAdded", edgeIdRule, mutualIdRule, originIdRule)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManagerEdgeAddedIterator{contract: _EdgeChallengeManager.contract, event: "EdgeAdded", logs: logs, sub: sub}, nil
}

// WatchEdgeAdded is a free log subscription operation binding the contract event 0xaa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a4.
//
// Solidity: event EdgeAdded(bytes32 indexed edgeId, bytes32 indexed mutualId, bytes32 indexed originId, bytes32 claimId, uint256 length, uint8 level, bool hasRival, bool isLayerZero)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) WatchEdgeAdded(opts *bind.WatchOpts, sink chan<- *EdgeChallengeManagerEdgeAdded, edgeId [][32]byte, mutualId [][32]byte, originId [][32]byte) (event.Subscription, error) {

	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}
	var mutualIdRule []interface{}
	for _, mutualIdItem := range mutualId {
		mutualIdRule = append(mutualIdRule, mutualIdItem)
	}
	var originIdRule []interface{}
	for _, originIdItem := range originId {
		originIdRule = append(originIdRule, originIdItem)
	}

	logs, sub, err := _EdgeChallengeManager.contract.WatchLogs(opts, "EdgeAdded", edgeIdRule, mutualIdRule, originIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeChallengeManagerEdgeAdded)
				if err := _EdgeChallengeManager.contract.UnpackLog(event, "EdgeAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEdgeAdded is a log parse operation binding the contract event 0xaa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a4.
//
// Solidity: event EdgeAdded(bytes32 indexed edgeId, bytes32 indexed mutualId, bytes32 indexed originId, bytes32 claimId, uint256 length, uint8 level, bool hasRival, bool isLayerZero)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) ParseEdgeAdded(log types.Log) (*EdgeChallengeManagerEdgeAdded, error) {
	event := new(EdgeChallengeManagerEdgeAdded)
	if err := _EdgeChallengeManager.contract.UnpackLog(event, "EdgeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeChallengeManagerEdgeBisectedIterator is returned from FilterEdgeBisected and is used to iterate over the raw logs and unpacked data for EdgeBisected events raised by the EdgeChallengeManager contract.
type EdgeChallengeManagerEdgeBisectedIterator struct {
	Event *EdgeChallengeManagerEdgeBisected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EdgeChallengeManagerEdgeBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeChallengeManagerEdgeBisected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EdgeChallengeManagerEdgeBisected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EdgeChallengeManagerEdgeBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeChallengeManagerEdgeBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeChallengeManagerEdgeBisected represents a EdgeBisected event raised by the EdgeChallengeManager contract.
type EdgeChallengeManagerEdgeBisected struct {
	EdgeId                  [32]byte
	LowerChildId            [32]byte
	UpperChildId            [32]byte
	LowerChildAlreadyExists bool
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterEdgeBisected is a free log retrieval operation binding the contract event 0x7340510d24b7ec9b5c100f5500d93429d80d00d46f0d18e4e85d0c4cc22b9924.
//
// Solidity: event EdgeBisected(bytes32 indexed edgeId, bytes32 indexed lowerChildId, bytes32 indexed upperChildId, bool lowerChildAlreadyExists)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) FilterEdgeBisected(opts *bind.FilterOpts, edgeId [][32]byte, lowerChildId [][32]byte, upperChildId [][32]byte) (*EdgeChallengeManagerEdgeBisectedIterator, error) {

	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}
	var lowerChildIdRule []interface{}
	for _, lowerChildIdItem := range lowerChildId {
		lowerChildIdRule = append(lowerChildIdRule, lowerChildIdItem)
	}
	var upperChildIdRule []interface{}
	for _, upperChildIdItem := range upperChildId {
		upperChildIdRule = append(upperChildIdRule, upperChildIdItem)
	}

	logs, sub, err := _EdgeChallengeManager.contract.FilterLogs(opts, "EdgeBisected", edgeIdRule, lowerChildIdRule, upperChildIdRule)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManagerEdgeBisectedIterator{contract: _EdgeChallengeManager.contract, event: "EdgeBisected", logs: logs, sub: sub}, nil
}

// WatchEdgeBisected is a free log subscription operation binding the contract event 0x7340510d24b7ec9b5c100f5500d93429d80d00d46f0d18e4e85d0c4cc22b9924.
//
// Solidity: event EdgeBisected(bytes32 indexed edgeId, bytes32 indexed lowerChildId, bytes32 indexed upperChildId, bool lowerChildAlreadyExists)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) WatchEdgeBisected(opts *bind.WatchOpts, sink chan<- *EdgeChallengeManagerEdgeBisected, edgeId [][32]byte, lowerChildId [][32]byte, upperChildId [][32]byte) (event.Subscription, error) {

	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}
	var lowerChildIdRule []interface{}
	for _, lowerChildIdItem := range lowerChildId {
		lowerChildIdRule = append(lowerChildIdRule, lowerChildIdItem)
	}
	var upperChildIdRule []interface{}
	for _, upperChildIdItem := range upperChildId {
		upperChildIdRule = append(upperChildIdRule, upperChildIdItem)
	}

	logs, sub, err := _EdgeChallengeManager.contract.WatchLogs(opts, "EdgeBisected", edgeIdRule, lowerChildIdRule, upperChildIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeChallengeManagerEdgeBisected)
				if err := _EdgeChallengeManager.contract.UnpackLog(event, "EdgeBisected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEdgeBisected is a log parse operation binding the contract event 0x7340510d24b7ec9b5c100f5500d93429d80d00d46f0d18e4e85d0c4cc22b9924.
//
// Solidity: event EdgeBisected(bytes32 indexed edgeId, bytes32 indexed lowerChildId, bytes32 indexed upperChildId, bool lowerChildAlreadyExists)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) ParseEdgeBisected(log types.Log) (*EdgeChallengeManagerEdgeBisected, error) {
	event := new(EdgeChallengeManagerEdgeBisected)
	if err := _EdgeChallengeManager.contract.UnpackLog(event, "EdgeBisected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeChallengeManagerEdgeConfirmedByChildrenIterator is returned from FilterEdgeConfirmedByChildren and is used to iterate over the raw logs and unpacked data for EdgeConfirmedByChildren events raised by the EdgeChallengeManager contract.
type EdgeChallengeManagerEdgeConfirmedByChildrenIterator struct {
	Event *EdgeChallengeManagerEdgeConfirmedByChildren // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EdgeChallengeManagerEdgeConfirmedByChildrenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeChallengeManagerEdgeConfirmedByChildren)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EdgeChallengeManagerEdgeConfirmedByChildren)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EdgeChallengeManagerEdgeConfirmedByChildrenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeChallengeManagerEdgeConfirmedByChildrenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeChallengeManagerEdgeConfirmedByChildren represents a EdgeConfirmedByChildren event raised by the EdgeChallengeManager contract.
type EdgeChallengeManagerEdgeConfirmedByChildren struct {
	EdgeId   [32]byte
	MutualId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEdgeConfirmedByChildren is a free log retrieval operation binding the contract event 0x0d27fcaf1adc41547a5cfc99d2364f6c0dc7e81c9fc3fe8cb38abb409b48358a.
//
// Solidity: event EdgeConfirmedByChildren(bytes32 indexed edgeId, bytes32 indexed mutualId)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) FilterEdgeConfirmedByChildren(opts *bind.FilterOpts, edgeId [][32]byte, mutualId [][32]byte) (*EdgeChallengeManagerEdgeConfirmedByChildrenIterator, error) {

	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}
	var mutualIdRule []interface{}
	for _, mutualIdItem := range mutualId {
		mutualIdRule = append(mutualIdRule, mutualIdItem)
	}

	logs, sub, err := _EdgeChallengeManager.contract.FilterLogs(opts, "EdgeConfirmedByChildren", edgeIdRule, mutualIdRule)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManagerEdgeConfirmedByChildrenIterator{contract: _EdgeChallengeManager.contract, event: "EdgeConfirmedByChildren", logs: logs, sub: sub}, nil
}

// WatchEdgeConfirmedByChildren is a free log subscription operation binding the contract event 0x0d27fcaf1adc41547a5cfc99d2364f6c0dc7e81c9fc3fe8cb38abb409b48358a.
//
// Solidity: event EdgeConfirmedByChildren(bytes32 indexed edgeId, bytes32 indexed mutualId)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) WatchEdgeConfirmedByChildren(opts *bind.WatchOpts, sink chan<- *EdgeChallengeManagerEdgeConfirmedByChildren, edgeId [][32]byte, mutualId [][32]byte) (event.Subscription, error) {

	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}
	var mutualIdRule []interface{}
	for _, mutualIdItem := range mutualId {
		mutualIdRule = append(mutualIdRule, mutualIdItem)
	}

	logs, sub, err := _EdgeChallengeManager.contract.WatchLogs(opts, "EdgeConfirmedByChildren", edgeIdRule, mutualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeChallengeManagerEdgeConfirmedByChildren)
				if err := _EdgeChallengeManager.contract.UnpackLog(event, "EdgeConfirmedByChildren", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEdgeConfirmedByChildren is a log parse operation binding the contract event 0x0d27fcaf1adc41547a5cfc99d2364f6c0dc7e81c9fc3fe8cb38abb409b48358a.
//
// Solidity: event EdgeConfirmedByChildren(bytes32 indexed edgeId, bytes32 indexed mutualId)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) ParseEdgeConfirmedByChildren(log types.Log) (*EdgeChallengeManagerEdgeConfirmedByChildren, error) {
	event := new(EdgeChallengeManagerEdgeConfirmedByChildren)
	if err := _EdgeChallengeManager.contract.UnpackLog(event, "EdgeConfirmedByChildren", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeChallengeManagerEdgeConfirmedByClaimIterator is returned from FilterEdgeConfirmedByClaim and is used to iterate over the raw logs and unpacked data for EdgeConfirmedByClaim events raised by the EdgeChallengeManager contract.
type EdgeChallengeManagerEdgeConfirmedByClaimIterator struct {
	Event *EdgeChallengeManagerEdgeConfirmedByClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EdgeChallengeManagerEdgeConfirmedByClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeChallengeManagerEdgeConfirmedByClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EdgeChallengeManagerEdgeConfirmedByClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EdgeChallengeManagerEdgeConfirmedByClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeChallengeManagerEdgeConfirmedByClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeChallengeManagerEdgeConfirmedByClaim represents a EdgeConfirmedByClaim event raised by the EdgeChallengeManager contract.
type EdgeChallengeManagerEdgeConfirmedByClaim struct {
	EdgeId         [32]byte
	MutualId       [32]byte
	ClaimingEdgeId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEdgeConfirmedByClaim is a free log retrieval operation binding the contract event 0xb924f3aa473645c7cf5b10262f927ae4ccf869d7fc239c17144b0c67490d1c73.
//
// Solidity: event EdgeConfirmedByClaim(bytes32 indexed edgeId, bytes32 indexed mutualId, bytes32 claimingEdgeId)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) FilterEdgeConfirmedByClaim(opts *bind.FilterOpts, edgeId [][32]byte, mutualId [][32]byte) (*EdgeChallengeManagerEdgeConfirmedByClaimIterator, error) {

	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}
	var mutualIdRule []interface{}
	for _, mutualIdItem := range mutualId {
		mutualIdRule = append(mutualIdRule, mutualIdItem)
	}

	logs, sub, err := _EdgeChallengeManager.contract.FilterLogs(opts, "EdgeConfirmedByClaim", edgeIdRule, mutualIdRule)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManagerEdgeConfirmedByClaimIterator{contract: _EdgeChallengeManager.contract, event: "EdgeConfirmedByClaim", logs: logs, sub: sub}, nil
}

// WatchEdgeConfirmedByClaim is a free log subscription operation binding the contract event 0xb924f3aa473645c7cf5b10262f927ae4ccf869d7fc239c17144b0c67490d1c73.
//
// Solidity: event EdgeConfirmedByClaim(bytes32 indexed edgeId, bytes32 indexed mutualId, bytes32 claimingEdgeId)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) WatchEdgeConfirmedByClaim(opts *bind.WatchOpts, sink chan<- *EdgeChallengeManagerEdgeConfirmedByClaim, edgeId [][32]byte, mutualId [][32]byte) (event.Subscription, error) {

	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}
	var mutualIdRule []interface{}
	for _, mutualIdItem := range mutualId {
		mutualIdRule = append(mutualIdRule, mutualIdItem)
	}

	logs, sub, err := _EdgeChallengeManager.contract.WatchLogs(opts, "EdgeConfirmedByClaim", edgeIdRule, mutualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeChallengeManagerEdgeConfirmedByClaim)
				if err := _EdgeChallengeManager.contract.UnpackLog(event, "EdgeConfirmedByClaim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEdgeConfirmedByClaim is a log parse operation binding the contract event 0xb924f3aa473645c7cf5b10262f927ae4ccf869d7fc239c17144b0c67490d1c73.
//
// Solidity: event EdgeConfirmedByClaim(bytes32 indexed edgeId, bytes32 indexed mutualId, bytes32 claimingEdgeId)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) ParseEdgeConfirmedByClaim(log types.Log) (*EdgeChallengeManagerEdgeConfirmedByClaim, error) {
	event := new(EdgeChallengeManagerEdgeConfirmedByClaim)
	if err := _EdgeChallengeManager.contract.UnpackLog(event, "EdgeConfirmedByClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeChallengeManagerEdgeConfirmedByOneStepProofIterator is returned from FilterEdgeConfirmedByOneStepProof and is used to iterate over the raw logs and unpacked data for EdgeConfirmedByOneStepProof events raised by the EdgeChallengeManager contract.
type EdgeChallengeManagerEdgeConfirmedByOneStepProofIterator struct {
	Event *EdgeChallengeManagerEdgeConfirmedByOneStepProof // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EdgeChallengeManagerEdgeConfirmedByOneStepProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeChallengeManagerEdgeConfirmedByOneStepProof)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EdgeChallengeManagerEdgeConfirmedByOneStepProof)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EdgeChallengeManagerEdgeConfirmedByOneStepProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeChallengeManagerEdgeConfirmedByOneStepProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeChallengeManagerEdgeConfirmedByOneStepProof represents a EdgeConfirmedByOneStepProof event raised by the EdgeChallengeManager contract.
type EdgeChallengeManagerEdgeConfirmedByOneStepProof struct {
	EdgeId   [32]byte
	MutualId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEdgeConfirmedByOneStepProof is a free log retrieval operation binding the contract event 0xe11db4b27bc8c6ea5943ecbb205ae1ca8d56c42c719717aaf8a53d43d0cee7c2.
//
// Solidity: event EdgeConfirmedByOneStepProof(bytes32 indexed edgeId, bytes32 indexed mutualId)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) FilterEdgeConfirmedByOneStepProof(opts *bind.FilterOpts, edgeId [][32]byte, mutualId [][32]byte) (*EdgeChallengeManagerEdgeConfirmedByOneStepProofIterator, error) {

	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}
	var mutualIdRule []interface{}
	for _, mutualIdItem := range mutualId {
		mutualIdRule = append(mutualIdRule, mutualIdItem)
	}

	logs, sub, err := _EdgeChallengeManager.contract.FilterLogs(opts, "EdgeConfirmedByOneStepProof", edgeIdRule, mutualIdRule)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManagerEdgeConfirmedByOneStepProofIterator{contract: _EdgeChallengeManager.contract, event: "EdgeConfirmedByOneStepProof", logs: logs, sub: sub}, nil
}

// WatchEdgeConfirmedByOneStepProof is a free log subscription operation binding the contract event 0xe11db4b27bc8c6ea5943ecbb205ae1ca8d56c42c719717aaf8a53d43d0cee7c2.
//
// Solidity: event EdgeConfirmedByOneStepProof(bytes32 indexed edgeId, bytes32 indexed mutualId)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) WatchEdgeConfirmedByOneStepProof(opts *bind.WatchOpts, sink chan<- *EdgeChallengeManagerEdgeConfirmedByOneStepProof, edgeId [][32]byte, mutualId [][32]byte) (event.Subscription, error) {

	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}
	var mutualIdRule []interface{}
	for _, mutualIdItem := range mutualId {
		mutualIdRule = append(mutualIdRule, mutualIdItem)
	}

	logs, sub, err := _EdgeChallengeManager.contract.WatchLogs(opts, "EdgeConfirmedByOneStepProof", edgeIdRule, mutualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeChallengeManagerEdgeConfirmedByOneStepProof)
				if err := _EdgeChallengeManager.contract.UnpackLog(event, "EdgeConfirmedByOneStepProof", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEdgeConfirmedByOneStepProof is a log parse operation binding the contract event 0xe11db4b27bc8c6ea5943ecbb205ae1ca8d56c42c719717aaf8a53d43d0cee7c2.
//
// Solidity: event EdgeConfirmedByOneStepProof(bytes32 indexed edgeId, bytes32 indexed mutualId)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) ParseEdgeConfirmedByOneStepProof(log types.Log) (*EdgeChallengeManagerEdgeConfirmedByOneStepProof, error) {
	event := new(EdgeChallengeManagerEdgeConfirmedByOneStepProof)
	if err := _EdgeChallengeManager.contract.UnpackLog(event, "EdgeConfirmedByOneStepProof", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeChallengeManagerEdgeConfirmedByTimeIterator is returned from FilterEdgeConfirmedByTime and is used to iterate over the raw logs and unpacked data for EdgeConfirmedByTime events raised by the EdgeChallengeManager contract.
type EdgeChallengeManagerEdgeConfirmedByTimeIterator struct {
	Event *EdgeChallengeManagerEdgeConfirmedByTime // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EdgeChallengeManagerEdgeConfirmedByTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeChallengeManagerEdgeConfirmedByTime)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EdgeChallengeManagerEdgeConfirmedByTime)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EdgeChallengeManagerEdgeConfirmedByTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeChallengeManagerEdgeConfirmedByTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeChallengeManagerEdgeConfirmedByTime represents a EdgeConfirmedByTime event raised by the EdgeChallengeManager contract.
type EdgeChallengeManagerEdgeConfirmedByTime struct {
	EdgeId             [32]byte
	MutualId           [32]byte
	TotalTimeUnrivaled uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEdgeConfirmedByTime is a free log retrieval operation binding the contract event 0x9cd2c77f6772dd0fb07e9972aee8ddddf9fd13f1a65abb03a3de88b07dc59af9.
//
// Solidity: event EdgeConfirmedByTime(bytes32 indexed edgeId, bytes32 indexed mutualId, uint64 totalTimeUnrivaled)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) FilterEdgeConfirmedByTime(opts *bind.FilterOpts, edgeId [][32]byte, mutualId [][32]byte) (*EdgeChallengeManagerEdgeConfirmedByTimeIterator, error) {

	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}
	var mutualIdRule []interface{}
	for _, mutualIdItem := range mutualId {
		mutualIdRule = append(mutualIdRule, mutualIdItem)
	}

	logs, sub, err := _EdgeChallengeManager.contract.FilterLogs(opts, "EdgeConfirmedByTime", edgeIdRule, mutualIdRule)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManagerEdgeConfirmedByTimeIterator{contract: _EdgeChallengeManager.contract, event: "EdgeConfirmedByTime", logs: logs, sub: sub}, nil
}

// WatchEdgeConfirmedByTime is a free log subscription operation binding the contract event 0x9cd2c77f6772dd0fb07e9972aee8ddddf9fd13f1a65abb03a3de88b07dc59af9.
//
// Solidity: event EdgeConfirmedByTime(bytes32 indexed edgeId, bytes32 indexed mutualId, uint64 totalTimeUnrivaled)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) WatchEdgeConfirmedByTime(opts *bind.WatchOpts, sink chan<- *EdgeChallengeManagerEdgeConfirmedByTime, edgeId [][32]byte, mutualId [][32]byte) (event.Subscription, error) {

	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}
	var mutualIdRule []interface{}
	for _, mutualIdItem := range mutualId {
		mutualIdRule = append(mutualIdRule, mutualIdItem)
	}

	logs, sub, err := _EdgeChallengeManager.contract.WatchLogs(opts, "EdgeConfirmedByTime", edgeIdRule, mutualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeChallengeManagerEdgeConfirmedByTime)
				if err := _EdgeChallengeManager.contract.UnpackLog(event, "EdgeConfirmedByTime", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEdgeConfirmedByTime is a log parse operation binding the contract event 0x9cd2c77f6772dd0fb07e9972aee8ddddf9fd13f1a65abb03a3de88b07dc59af9.
//
// Solidity: event EdgeConfirmedByTime(bytes32 indexed edgeId, bytes32 indexed mutualId, uint64 totalTimeUnrivaled)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) ParseEdgeConfirmedByTime(log types.Log) (*EdgeChallengeManagerEdgeConfirmedByTime, error) {
	event := new(EdgeChallengeManagerEdgeConfirmedByTime)
	if err := _EdgeChallengeManager.contract.UnpackLog(event, "EdgeConfirmedByTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeChallengeManagerEdgeRefundedIterator is returned from FilterEdgeRefunded and is used to iterate over the raw logs and unpacked data for EdgeRefunded events raised by the EdgeChallengeManager contract.
type EdgeChallengeManagerEdgeRefundedIterator struct {
	Event *EdgeChallengeManagerEdgeRefunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EdgeChallengeManagerEdgeRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeChallengeManagerEdgeRefunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EdgeChallengeManagerEdgeRefunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EdgeChallengeManagerEdgeRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeChallengeManagerEdgeRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeChallengeManagerEdgeRefunded represents a EdgeRefunded event raised by the EdgeChallengeManager contract.
type EdgeChallengeManagerEdgeRefunded struct {
	EdgeId      [32]byte
	MutualId    [32]byte
	StakeToken  common.Address
	StakeAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEdgeRefunded is a free log retrieval operation binding the contract event 0xa635398959ddb5ce3b14537edfc25b2e671274c9b8cad0f4bd634752e69007b6.
//
// Solidity: event EdgeRefunded(bytes32 indexed edgeId, bytes32 indexed mutualId, address stakeToken, uint256 stakeAmount)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) FilterEdgeRefunded(opts *bind.FilterOpts, edgeId [][32]byte, mutualId [][32]byte) (*EdgeChallengeManagerEdgeRefundedIterator, error) {

	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}
	var mutualIdRule []interface{}
	for _, mutualIdItem := range mutualId {
		mutualIdRule = append(mutualIdRule, mutualIdItem)
	}

	logs, sub, err := _EdgeChallengeManager.contract.FilterLogs(opts, "EdgeRefunded", edgeIdRule, mutualIdRule)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManagerEdgeRefundedIterator{contract: _EdgeChallengeManager.contract, event: "EdgeRefunded", logs: logs, sub: sub}, nil
}

// WatchEdgeRefunded is a free log subscription operation binding the contract event 0xa635398959ddb5ce3b14537edfc25b2e671274c9b8cad0f4bd634752e69007b6.
//
// Solidity: event EdgeRefunded(bytes32 indexed edgeId, bytes32 indexed mutualId, address stakeToken, uint256 stakeAmount)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) WatchEdgeRefunded(opts *bind.WatchOpts, sink chan<- *EdgeChallengeManagerEdgeRefunded, edgeId [][32]byte, mutualId [][32]byte) (event.Subscription, error) {

	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}
	var mutualIdRule []interface{}
	for _, mutualIdItem := range mutualId {
		mutualIdRule = append(mutualIdRule, mutualIdItem)
	}

	logs, sub, err := _EdgeChallengeManager.contract.WatchLogs(opts, "EdgeRefunded", edgeIdRule, mutualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeChallengeManagerEdgeRefunded)
				if err := _EdgeChallengeManager.contract.UnpackLog(event, "EdgeRefunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEdgeRefunded is a log parse operation binding the contract event 0xa635398959ddb5ce3b14537edfc25b2e671274c9b8cad0f4bd634752e69007b6.
//
// Solidity: event EdgeRefunded(bytes32 indexed edgeId, bytes32 indexed mutualId, address stakeToken, uint256 stakeAmount)
func (_EdgeChallengeManager *EdgeChallengeManagerFilterer) ParseEdgeRefunded(log types.Log) (*EdgeChallengeManagerEdgeRefunded, error) {
	event := new(EdgeChallengeManagerEdgeRefunded)
	if err := _EdgeChallengeManager.contract.UnpackLog(event, "EdgeRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAssertionChainMetaData contains all meta data concerning the IAssertionChain contract.
var IAssertionChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionHash\",\"type\":\"bytes32\"}],\"name\":\"getFirstChildCreationBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionHash\",\"type\":\"bytes32\"}],\"name\":\"getSecondChildCreationBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionHash\",\"type\":\"bytes32\"}],\"name\":\"isFirstChild\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionHash\",\"type\":\"bytes32\"}],\"name\":\"isPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionHash\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState\",\"name\":\"globalState\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus\",\"name\":\"machineStatus\",\"type\":\"uint8\"}],\"internalType\":\"structExecutionState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"prevAssertionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inboxAcc\",\"type\":\"bytes32\"}],\"name\":\"validateAssertionHash\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"requiredStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"challengeManager\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"confirmPeriodBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextInboxPosition\",\"type\":\"uint64\"}],\"internalType\":\"structConfigData\",\"name\":\"configData\",\"type\":\"tuple\"}],\"name\":\"validateConfig\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IAssertionChainABI is the input ABI used to generate the binding from.
// Deprecated: Use IAssertionChainMetaData.ABI instead.
var IAssertionChainABI = IAssertionChainMetaData.ABI

// IAssertionChain is an auto generated Go binding around an Ethereum contract.
type IAssertionChain struct {
	IAssertionChainCaller     // Read-only binding to the contract
	IAssertionChainTransactor // Write-only binding to the contract
	IAssertionChainFilterer   // Log filterer for contract events
}

// IAssertionChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAssertionChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssertionChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAssertionChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssertionChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAssertionChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssertionChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAssertionChainSession struct {
	Contract     *IAssertionChain  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAssertionChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAssertionChainCallerSession struct {
	Contract *IAssertionChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IAssertionChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAssertionChainTransactorSession struct {
	Contract     *IAssertionChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IAssertionChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAssertionChainRaw struct {
	Contract *IAssertionChain // Generic contract binding to access the raw methods on
}

// IAssertionChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAssertionChainCallerRaw struct {
	Contract *IAssertionChainCaller // Generic read-only contract binding to access the raw methods on
}

// IAssertionChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAssertionChainTransactorRaw struct {
	Contract *IAssertionChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAssertionChain creates a new instance of IAssertionChain, bound to a specific deployed contract.
func NewIAssertionChain(address common.Address, backend bind.ContractBackend) (*IAssertionChain, error) {
	contract, err := bindIAssertionChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAssertionChain{IAssertionChainCaller: IAssertionChainCaller{contract: contract}, IAssertionChainTransactor: IAssertionChainTransactor{contract: contract}, IAssertionChainFilterer: IAssertionChainFilterer{contract: contract}}, nil
}

// NewIAssertionChainCaller creates a new read-only instance of IAssertionChain, bound to a specific deployed contract.
func NewIAssertionChainCaller(address common.Address, caller bind.ContractCaller) (*IAssertionChainCaller, error) {
	contract, err := bindIAssertionChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAssertionChainCaller{contract: contract}, nil
}

// NewIAssertionChainTransactor creates a new write-only instance of IAssertionChain, bound to a specific deployed contract.
func NewIAssertionChainTransactor(address common.Address, transactor bind.ContractTransactor) (*IAssertionChainTransactor, error) {
	contract, err := bindIAssertionChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAssertionChainTransactor{contract: contract}, nil
}

// NewIAssertionChainFilterer creates a new log filterer instance of IAssertionChain, bound to a specific deployed contract.
func NewIAssertionChainFilterer(address common.Address, filterer bind.ContractFilterer) (*IAssertionChainFilterer, error) {
	contract, err := bindIAssertionChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAssertionChainFilterer{contract: contract}, nil
}

// bindIAssertionChain binds a generic wrapper to an already deployed contract.
func bindIAssertionChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAssertionChainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAssertionChain *IAssertionChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAssertionChain.Contract.IAssertionChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAssertionChain *IAssertionChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAssertionChain.Contract.IAssertionChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAssertionChain *IAssertionChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAssertionChain.Contract.IAssertionChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAssertionChain *IAssertionChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAssertionChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAssertionChain *IAssertionChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAssertionChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAssertionChain *IAssertionChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAssertionChain.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IAssertionChain *IAssertionChainCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IAssertionChain *IAssertionChainSession) Bridge() (common.Address, error) {
	return _IAssertionChain.Contract.Bridge(&_IAssertionChain.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IAssertionChain *IAssertionChainCallerSession) Bridge() (common.Address, error) {
	return _IAssertionChain.Contract.Bridge(&_IAssertionChain.CallOpts)
}

// GetFirstChildCreationBlock is a free data retrieval call binding the contract method 0x11715585.
//
// Solidity: function getFirstChildCreationBlock(bytes32 assertionHash) view returns(uint64)
func (_IAssertionChain *IAssertionChainCaller) GetFirstChildCreationBlock(opts *bind.CallOpts, assertionHash [32]byte) (uint64, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "getFirstChildCreationBlock", assertionHash)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetFirstChildCreationBlock is a free data retrieval call binding the contract method 0x11715585.
//
// Solidity: function getFirstChildCreationBlock(bytes32 assertionHash) view returns(uint64)
func (_IAssertionChain *IAssertionChainSession) GetFirstChildCreationBlock(assertionHash [32]byte) (uint64, error) {
	return _IAssertionChain.Contract.GetFirstChildCreationBlock(&_IAssertionChain.CallOpts, assertionHash)
}

// GetFirstChildCreationBlock is a free data retrieval call binding the contract method 0x11715585.
//
// Solidity: function getFirstChildCreationBlock(bytes32 assertionHash) view returns(uint64)
func (_IAssertionChain *IAssertionChainCallerSession) GetFirstChildCreationBlock(assertionHash [32]byte) (uint64, error) {
	return _IAssertionChain.Contract.GetFirstChildCreationBlock(&_IAssertionChain.CallOpts, assertionHash)
}

// GetSecondChildCreationBlock is a free data retrieval call binding the contract method 0x56bbc9e6.
//
// Solidity: function getSecondChildCreationBlock(bytes32 assertionHash) view returns(uint64)
func (_IAssertionChain *IAssertionChainCaller) GetSecondChildCreationBlock(opts *bind.CallOpts, assertionHash [32]byte) (uint64, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "getSecondChildCreationBlock", assertionHash)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetSecondChildCreationBlock is a free data retrieval call binding the contract method 0x56bbc9e6.
//
// Solidity: function getSecondChildCreationBlock(bytes32 assertionHash) view returns(uint64)
func (_IAssertionChain *IAssertionChainSession) GetSecondChildCreationBlock(assertionHash [32]byte) (uint64, error) {
	return _IAssertionChain.Contract.GetSecondChildCreationBlock(&_IAssertionChain.CallOpts, assertionHash)
}

// GetSecondChildCreationBlock is a free data retrieval call binding the contract method 0x56bbc9e6.
//
// Solidity: function getSecondChildCreationBlock(bytes32 assertionHash) view returns(uint64)
func (_IAssertionChain *IAssertionChainCallerSession) GetSecondChildCreationBlock(assertionHash [32]byte) (uint64, error) {
	return _IAssertionChain.Contract.GetSecondChildCreationBlock(&_IAssertionChain.CallOpts, assertionHash)
}

// IsFirstChild is a free data retrieval call binding the contract method 0x30836228.
//
// Solidity: function isFirstChild(bytes32 assertionHash) view returns(bool)
func (_IAssertionChain *IAssertionChainCaller) IsFirstChild(opts *bind.CallOpts, assertionHash [32]byte) (bool, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "isFirstChild", assertionHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFirstChild is a free data retrieval call binding the contract method 0x30836228.
//
// Solidity: function isFirstChild(bytes32 assertionHash) view returns(bool)
func (_IAssertionChain *IAssertionChainSession) IsFirstChild(assertionHash [32]byte) (bool, error) {
	return _IAssertionChain.Contract.IsFirstChild(&_IAssertionChain.CallOpts, assertionHash)
}

// IsFirstChild is a free data retrieval call binding the contract method 0x30836228.
//
// Solidity: function isFirstChild(bytes32 assertionHash) view returns(bool)
func (_IAssertionChain *IAssertionChainCallerSession) IsFirstChild(assertionHash [32]byte) (bool, error) {
	return _IAssertionChain.Contract.IsFirstChild(&_IAssertionChain.CallOpts, assertionHash)
}

// IsPending is a free data retrieval call binding the contract method 0xe531d8c7.
//
// Solidity: function isPending(bytes32 assertionHash) view returns(bool)
func (_IAssertionChain *IAssertionChainCaller) IsPending(opts *bind.CallOpts, assertionHash [32]byte) (bool, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "isPending", assertionHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPending is a free data retrieval call binding the contract method 0xe531d8c7.
//
// Solidity: function isPending(bytes32 assertionHash) view returns(bool)
func (_IAssertionChain *IAssertionChainSession) IsPending(assertionHash [32]byte) (bool, error) {
	return _IAssertionChain.Contract.IsPending(&_IAssertionChain.CallOpts, assertionHash)
}

// IsPending is a free data retrieval call binding the contract method 0xe531d8c7.
//
// Solidity: function isPending(bytes32 assertionHash) view returns(bool)
func (_IAssertionChain *IAssertionChainCallerSession) IsPending(assertionHash [32]byte) (bool, error) {
	return _IAssertionChain.Contract.IsPending(&_IAssertionChain.CallOpts, assertionHash)
}

// ValidateAssertionHash is a free data retrieval call binding the contract method 0xf9cee9df.
//
// Solidity: function validateAssertionHash(bytes32 assertionHash, ((bytes32[2],uint64[2]),uint8) state, bytes32 prevAssertionHash, bytes32 inboxAcc) view returns()
func (_IAssertionChain *IAssertionChainCaller) ValidateAssertionHash(opts *bind.CallOpts, assertionHash [32]byte, state ExecutionState, prevAssertionHash [32]byte, inboxAcc [32]byte) error {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "validateAssertionHash", assertionHash, state, prevAssertionHash, inboxAcc)

	if err != nil {
		return err
	}

	return err

}

// ValidateAssertionHash is a free data retrieval call binding the contract method 0xf9cee9df.
//
// Solidity: function validateAssertionHash(bytes32 assertionHash, ((bytes32[2],uint64[2]),uint8) state, bytes32 prevAssertionHash, bytes32 inboxAcc) view returns()
func (_IAssertionChain *IAssertionChainSession) ValidateAssertionHash(assertionHash [32]byte, state ExecutionState, prevAssertionHash [32]byte, inboxAcc [32]byte) error {
	return _IAssertionChain.Contract.ValidateAssertionHash(&_IAssertionChain.CallOpts, assertionHash, state, prevAssertionHash, inboxAcc)
}

// ValidateAssertionHash is a free data retrieval call binding the contract method 0xf9cee9df.
//
// Solidity: function validateAssertionHash(bytes32 assertionHash, ((bytes32[2],uint64[2]),uint8) state, bytes32 prevAssertionHash, bytes32 inboxAcc) view returns()
func (_IAssertionChain *IAssertionChainCallerSession) ValidateAssertionHash(assertionHash [32]byte, state ExecutionState, prevAssertionHash [32]byte, inboxAcc [32]byte) error {
	return _IAssertionChain.Contract.ValidateAssertionHash(&_IAssertionChain.CallOpts, assertionHash, state, prevAssertionHash, inboxAcc)
}

// ValidateConfig is a free data retrieval call binding the contract method 0x04972af9.
//
// Solidity: function validateConfig(bytes32 assertionHash, (bytes32,uint256,address,uint64,uint64) configData) view returns()
func (_IAssertionChain *IAssertionChainCaller) ValidateConfig(opts *bind.CallOpts, assertionHash [32]byte, configData ConfigData) error {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "validateConfig", assertionHash, configData)

	if err != nil {
		return err
	}

	return err

}

// ValidateConfig is a free data retrieval call binding the contract method 0x04972af9.
//
// Solidity: function validateConfig(bytes32 assertionHash, (bytes32,uint256,address,uint64,uint64) configData) view returns()
func (_IAssertionChain *IAssertionChainSession) ValidateConfig(assertionHash [32]byte, configData ConfigData) error {
	return _IAssertionChain.Contract.ValidateConfig(&_IAssertionChain.CallOpts, assertionHash, configData)
}

// ValidateConfig is a free data retrieval call binding the contract method 0x04972af9.
//
// Solidity: function validateConfig(bytes32 assertionHash, (bytes32,uint256,address,uint64,uint64) configData) view returns()
func (_IAssertionChain *IAssertionChainCallerSession) ValidateConfig(assertionHash [32]byte, configData ConfigData) error {
	return _IAssertionChain.Contract.ValidateConfig(&_IAssertionChain.CallOpts, assertionHash, configData)
}

// IEdgeChallengeManagerMetaData contains all meta data concerning the IEdgeChallengeManager contract.
var IEdgeChallengeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bisectionHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefixProof\",\"type\":\"bytes\"}],\"name\":\"bisectEdge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"originId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"startHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"endHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"endHistoryRoot\",\"type\":\"bytes32\"}],\"name\":\"calculateEdgeId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"originId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"startHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"endHeight\",\"type\":\"uint256\"}],\"name\":\"calculateMutualId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengePeriodBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"confirmEdgeByChildren\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"claimingEdgeId\",\"type\":\"bytes32\"}],\"name\":\"confirmEdgeByClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structOneStepData\",\"name\":\"oneStepData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"requiredStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"challengeManager\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"confirmPeriodBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextInboxPosition\",\"type\":\"uint64\"}],\"internalType\":\"structConfigData\",\"name\":\"prevConfig\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"beforeHistoryInclusionProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"afterHistoryInclusionProof\",\"type\":\"bytes32[]\"}],\"name\":\"confirmEdgeByOneStepProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"ancestorEdgeIds\",\"type\":\"bytes32[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState\",\"name\":\"globalState\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus\",\"name\":\"machineStatus\",\"type\":\"uint8\"}],\"internalType\":\"structExecutionState\",\"name\":\"executionState\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"prevAssertionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inboxAcc\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionStateData\",\"name\":\"claimStateData\",\"type\":\"tuple\"}],\"name\":\"confirmEdgeByTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mutualId\",\"type\":\"bytes32\"}],\"name\":\"confirmedRival\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"endHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"endHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefixProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structCreateEdgeArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"createLayerZeroEdge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"edgeExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"edgeLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"firstRival\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"getEdge\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"originId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"startHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"endHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"endHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lowerChildId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"upperChildId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"createdAtBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"confirmedAtBlock\",\"type\":\"uint64\"},{\"internalType\":\"enumEdgeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"refunded\",\"type\":\"bool\"}],\"internalType\":\"structChallengeEdge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumEdgeType\",\"name\":\"eType\",\"type\":\"uint8\"}],\"name\":\"getLayerZeroEndHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"getPrevAssertionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"hasLengthOneRival\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"hasRival\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oneStepProofEntry\",\"outputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"refundStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"timeUnrivaled\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IEdgeChallengeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IEdgeChallengeManagerMetaData.ABI instead.
var IEdgeChallengeManagerABI = IEdgeChallengeManagerMetaData.ABI

// IEdgeChallengeManager is an auto generated Go binding around an Ethereum contract.
type IEdgeChallengeManager struct {
	IEdgeChallengeManagerCaller     // Read-only binding to the contract
	IEdgeChallengeManagerTransactor // Write-only binding to the contract
	IEdgeChallengeManagerFilterer   // Log filterer for contract events
}

// IEdgeChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEdgeChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEdgeChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEdgeChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEdgeChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEdgeChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEdgeChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEdgeChallengeManagerSession struct {
	Contract     *IEdgeChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IEdgeChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEdgeChallengeManagerCallerSession struct {
	Contract *IEdgeChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IEdgeChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEdgeChallengeManagerTransactorSession struct {
	Contract     *IEdgeChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IEdgeChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEdgeChallengeManagerRaw struct {
	Contract *IEdgeChallengeManager // Generic contract binding to access the raw methods on
}

// IEdgeChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEdgeChallengeManagerCallerRaw struct {
	Contract *IEdgeChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IEdgeChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEdgeChallengeManagerTransactorRaw struct {
	Contract *IEdgeChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEdgeChallengeManager creates a new instance of IEdgeChallengeManager, bound to a specific deployed contract.
func NewIEdgeChallengeManager(address common.Address, backend bind.ContractBackend) (*IEdgeChallengeManager, error) {
	contract, err := bindIEdgeChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEdgeChallengeManager{IEdgeChallengeManagerCaller: IEdgeChallengeManagerCaller{contract: contract}, IEdgeChallengeManagerTransactor: IEdgeChallengeManagerTransactor{contract: contract}, IEdgeChallengeManagerFilterer: IEdgeChallengeManagerFilterer{contract: contract}}, nil
}

// NewIEdgeChallengeManagerCaller creates a new read-only instance of IEdgeChallengeManager, bound to a specific deployed contract.
func NewIEdgeChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*IEdgeChallengeManagerCaller, error) {
	contract, err := bindIEdgeChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEdgeChallengeManagerCaller{contract: contract}, nil
}

// NewIEdgeChallengeManagerTransactor creates a new write-only instance of IEdgeChallengeManager, bound to a specific deployed contract.
func NewIEdgeChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IEdgeChallengeManagerTransactor, error) {
	contract, err := bindIEdgeChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEdgeChallengeManagerTransactor{contract: contract}, nil
}

// NewIEdgeChallengeManagerFilterer creates a new log filterer instance of IEdgeChallengeManager, bound to a specific deployed contract.
func NewIEdgeChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IEdgeChallengeManagerFilterer, error) {
	contract, err := bindIEdgeChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEdgeChallengeManagerFilterer{contract: contract}, nil
}

// bindIEdgeChallengeManager binds a generic wrapper to an already deployed contract.
func bindIEdgeChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEdgeChallengeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEdgeChallengeManager *IEdgeChallengeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEdgeChallengeManager.Contract.IEdgeChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEdgeChallengeManager *IEdgeChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.IEdgeChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEdgeChallengeManager *IEdgeChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.IEdgeChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEdgeChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// CalculateEdgeId is a free data retrieval call binding the contract method 0x004d8efe.
//
// Solidity: function calculateEdgeId(uint8 level, bytes32 originId, uint256 startHeight, bytes32 startHistoryRoot, uint256 endHeight, bytes32 endHistoryRoot) pure returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCaller) CalculateEdgeId(opts *bind.CallOpts, level uint8, originId [32]byte, startHeight *big.Int, startHistoryRoot [32]byte, endHeight *big.Int, endHistoryRoot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IEdgeChallengeManager.contract.Call(opts, &out, "calculateEdgeId", level, originId, startHeight, startHistoryRoot, endHeight, endHistoryRoot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateEdgeId is a free data retrieval call binding the contract method 0x004d8efe.
//
// Solidity: function calculateEdgeId(uint8 level, bytes32 originId, uint256 startHeight, bytes32 startHistoryRoot, uint256 endHeight, bytes32 endHistoryRoot) pure returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) CalculateEdgeId(level uint8, originId [32]byte, startHeight *big.Int, startHistoryRoot [32]byte, endHeight *big.Int, endHistoryRoot [32]byte) ([32]byte, error) {
	return _IEdgeChallengeManager.Contract.CalculateEdgeId(&_IEdgeChallengeManager.CallOpts, level, originId, startHeight, startHistoryRoot, endHeight, endHistoryRoot)
}

// CalculateEdgeId is a free data retrieval call binding the contract method 0x004d8efe.
//
// Solidity: function calculateEdgeId(uint8 level, bytes32 originId, uint256 startHeight, bytes32 startHistoryRoot, uint256 endHeight, bytes32 endHistoryRoot) pure returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerSession) CalculateEdgeId(level uint8, originId [32]byte, startHeight *big.Int, startHistoryRoot [32]byte, endHeight *big.Int, endHistoryRoot [32]byte) ([32]byte, error) {
	return _IEdgeChallengeManager.Contract.CalculateEdgeId(&_IEdgeChallengeManager.CallOpts, level, originId, startHeight, startHistoryRoot, endHeight, endHistoryRoot)
}

// CalculateMutualId is a free data retrieval call binding the contract method 0xc32d8c63.
//
// Solidity: function calculateMutualId(uint8 level, bytes32 originId, uint256 startHeight, bytes32 startHistoryRoot, uint256 endHeight) pure returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCaller) CalculateMutualId(opts *bind.CallOpts, level uint8, originId [32]byte, startHeight *big.Int, startHistoryRoot [32]byte, endHeight *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IEdgeChallengeManager.contract.Call(opts, &out, "calculateMutualId", level, originId, startHeight, startHistoryRoot, endHeight)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMutualId is a free data retrieval call binding the contract method 0xc32d8c63.
//
// Solidity: function calculateMutualId(uint8 level, bytes32 originId, uint256 startHeight, bytes32 startHistoryRoot, uint256 endHeight) pure returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) CalculateMutualId(level uint8, originId [32]byte, startHeight *big.Int, startHistoryRoot [32]byte, endHeight *big.Int) ([32]byte, error) {
	return _IEdgeChallengeManager.Contract.CalculateMutualId(&_IEdgeChallengeManager.CallOpts, level, originId, startHeight, startHistoryRoot, endHeight)
}

// CalculateMutualId is a free data retrieval call binding the contract method 0xc32d8c63.
//
// Solidity: function calculateMutualId(uint8 level, bytes32 originId, uint256 startHeight, bytes32 startHistoryRoot, uint256 endHeight) pure returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerSession) CalculateMutualId(level uint8, originId [32]byte, startHeight *big.Int, startHistoryRoot [32]byte, endHeight *big.Int) ([32]byte, error) {
	return _IEdgeChallengeManager.Contract.CalculateMutualId(&_IEdgeChallengeManager.CallOpts, level, originId, startHeight, startHistoryRoot, endHeight)
}

// ChallengePeriodBlocks is a free data retrieval call binding the contract method 0x46c2781a.
//
// Solidity: function challengePeriodBlocks() view returns(uint64)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCaller) ChallengePeriodBlocks(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IEdgeChallengeManager.contract.Call(opts, &out, "challengePeriodBlocks")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChallengePeriodBlocks is a free data retrieval call binding the contract method 0x46c2781a.
//
// Solidity: function challengePeriodBlocks() view returns(uint64)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) ChallengePeriodBlocks() (uint64, error) {
	return _IEdgeChallengeManager.Contract.ChallengePeriodBlocks(&_IEdgeChallengeManager.CallOpts)
}

// ChallengePeriodBlocks is a free data retrieval call binding the contract method 0x46c2781a.
//
// Solidity: function challengePeriodBlocks() view returns(uint64)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerSession) ChallengePeriodBlocks() (uint64, error) {
	return _IEdgeChallengeManager.Contract.ChallengePeriodBlocks(&_IEdgeChallengeManager.CallOpts)
}

// ConfirmedRival is a free data retrieval call binding the contract method 0xe5b123da.
//
// Solidity: function confirmedRival(bytes32 mutualId) view returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCaller) ConfirmedRival(opts *bind.CallOpts, mutualId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IEdgeChallengeManager.contract.Call(opts, &out, "confirmedRival", mutualId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConfirmedRival is a free data retrieval call binding the contract method 0xe5b123da.
//
// Solidity: function confirmedRival(bytes32 mutualId) view returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) ConfirmedRival(mutualId [32]byte) ([32]byte, error) {
	return _IEdgeChallengeManager.Contract.ConfirmedRival(&_IEdgeChallengeManager.CallOpts, mutualId)
}

// ConfirmedRival is a free data retrieval call binding the contract method 0xe5b123da.
//
// Solidity: function confirmedRival(bytes32 mutualId) view returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerSession) ConfirmedRival(mutualId [32]byte) ([32]byte, error) {
	return _IEdgeChallengeManager.Contract.ConfirmedRival(&_IEdgeChallengeManager.CallOpts, mutualId)
}

// EdgeExists is a free data retrieval call binding the contract method 0x750e0c0f.
//
// Solidity: function edgeExists(bytes32 edgeId) view returns(bool)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCaller) EdgeExists(opts *bind.CallOpts, edgeId [32]byte) (bool, error) {
	var out []interface{}
	err := _IEdgeChallengeManager.contract.Call(opts, &out, "edgeExists", edgeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EdgeExists is a free data retrieval call binding the contract method 0x750e0c0f.
//
// Solidity: function edgeExists(bytes32 edgeId) view returns(bool)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) EdgeExists(edgeId [32]byte) (bool, error) {
	return _IEdgeChallengeManager.Contract.EdgeExists(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// EdgeExists is a free data retrieval call binding the contract method 0x750e0c0f.
//
// Solidity: function edgeExists(bytes32 edgeId) view returns(bool)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerSession) EdgeExists(edgeId [32]byte) (bool, error) {
	return _IEdgeChallengeManager.Contract.EdgeExists(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// EdgeLength is a free data retrieval call binding the contract method 0xeae0328b.
//
// Solidity: function edgeLength(bytes32 edgeId) view returns(uint256)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCaller) EdgeLength(opts *bind.CallOpts, edgeId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IEdgeChallengeManager.contract.Call(opts, &out, "edgeLength", edgeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EdgeLength is a free data retrieval call binding the contract method 0xeae0328b.
//
// Solidity: function edgeLength(bytes32 edgeId) view returns(uint256)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) EdgeLength(edgeId [32]byte) (*big.Int, error) {
	return _IEdgeChallengeManager.Contract.EdgeLength(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// EdgeLength is a free data retrieval call binding the contract method 0xeae0328b.
//
// Solidity: function edgeLength(bytes32 edgeId) view returns(uint256)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerSession) EdgeLength(edgeId [32]byte) (*big.Int, error) {
	return _IEdgeChallengeManager.Contract.EdgeLength(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// FirstRival is a free data retrieval call binding the contract method 0xbce6f54f.
//
// Solidity: function firstRival(bytes32 edgeId) view returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCaller) FirstRival(opts *bind.CallOpts, edgeId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IEdgeChallengeManager.contract.Call(opts, &out, "firstRival", edgeId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FirstRival is a free data retrieval call binding the contract method 0xbce6f54f.
//
// Solidity: function firstRival(bytes32 edgeId) view returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) FirstRival(edgeId [32]byte) ([32]byte, error) {
	return _IEdgeChallengeManager.Contract.FirstRival(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// FirstRival is a free data retrieval call binding the contract method 0xbce6f54f.
//
// Solidity: function firstRival(bytes32 edgeId) view returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerSession) FirstRival(edgeId [32]byte) ([32]byte, error) {
	return _IEdgeChallengeManager.Contract.FirstRival(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// GetEdge is a free data retrieval call binding the contract method 0xfda2892e.
//
// Solidity: function getEdge(bytes32 edgeId) view returns((bytes32,bytes32,uint256,bytes32,uint256,bytes32,bytes32,bytes32,address,uint64,uint64,uint8,uint8,bool))
func (_IEdgeChallengeManager *IEdgeChallengeManagerCaller) GetEdge(opts *bind.CallOpts, edgeId [32]byte) (ChallengeEdge, error) {
	var out []interface{}
	err := _IEdgeChallengeManager.contract.Call(opts, &out, "getEdge", edgeId)

	if err != nil {
		return *new(ChallengeEdge), err
	}

	out0 := *abi.ConvertType(out[0], new(ChallengeEdge)).(*ChallengeEdge)

	return out0, err

}

// GetEdge is a free data retrieval call binding the contract method 0xfda2892e.
//
// Solidity: function getEdge(bytes32 edgeId) view returns((bytes32,bytes32,uint256,bytes32,uint256,bytes32,bytes32,bytes32,address,uint64,uint64,uint8,uint8,bool))
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) GetEdge(edgeId [32]byte) (ChallengeEdge, error) {
	return _IEdgeChallengeManager.Contract.GetEdge(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// GetEdge is a free data retrieval call binding the contract method 0xfda2892e.
//
// Solidity: function getEdge(bytes32 edgeId) view returns((bytes32,bytes32,uint256,bytes32,uint256,bytes32,bytes32,bytes32,address,uint64,uint64,uint8,uint8,bool))
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerSession) GetEdge(edgeId [32]byte) (ChallengeEdge, error) {
	return _IEdgeChallengeManager.Contract.GetEdge(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// GetLayerZeroEndHeight is a free data retrieval call binding the contract method 0x42e1aaa8.
//
// Solidity: function getLayerZeroEndHeight(uint8 eType) view returns(uint256)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCaller) GetLayerZeroEndHeight(opts *bind.CallOpts, eType uint8) (*big.Int, error) {
	var out []interface{}
	err := _IEdgeChallengeManager.contract.Call(opts, &out, "getLayerZeroEndHeight", eType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLayerZeroEndHeight is a free data retrieval call binding the contract method 0x42e1aaa8.
//
// Solidity: function getLayerZeroEndHeight(uint8 eType) view returns(uint256)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) GetLayerZeroEndHeight(eType uint8) (*big.Int, error) {
	return _IEdgeChallengeManager.Contract.GetLayerZeroEndHeight(&_IEdgeChallengeManager.CallOpts, eType)
}

// GetLayerZeroEndHeight is a free data retrieval call binding the contract method 0x42e1aaa8.
//
// Solidity: function getLayerZeroEndHeight(uint8 eType) view returns(uint256)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerSession) GetLayerZeroEndHeight(eType uint8) (*big.Int, error) {
	return _IEdgeChallengeManager.Contract.GetLayerZeroEndHeight(&_IEdgeChallengeManager.CallOpts, eType)
}

// GetPrevAssertionHash is a free data retrieval call binding the contract method 0x5a48e0f4.
//
// Solidity: function getPrevAssertionHash(bytes32 edgeId) view returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCaller) GetPrevAssertionHash(opts *bind.CallOpts, edgeId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IEdgeChallengeManager.contract.Call(opts, &out, "getPrevAssertionHash", edgeId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPrevAssertionHash is a free data retrieval call binding the contract method 0x5a48e0f4.
//
// Solidity: function getPrevAssertionHash(bytes32 edgeId) view returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) GetPrevAssertionHash(edgeId [32]byte) ([32]byte, error) {
	return _IEdgeChallengeManager.Contract.GetPrevAssertionHash(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// GetPrevAssertionHash is a free data retrieval call binding the contract method 0x5a48e0f4.
//
// Solidity: function getPrevAssertionHash(bytes32 edgeId) view returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerSession) GetPrevAssertionHash(edgeId [32]byte) ([32]byte, error) {
	return _IEdgeChallengeManager.Contract.GetPrevAssertionHash(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// HasLengthOneRival is a free data retrieval call binding the contract method 0x54b64151.
//
// Solidity: function hasLengthOneRival(bytes32 edgeId) view returns(bool)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCaller) HasLengthOneRival(opts *bind.CallOpts, edgeId [32]byte) (bool, error) {
	var out []interface{}
	err := _IEdgeChallengeManager.contract.Call(opts, &out, "hasLengthOneRival", edgeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasLengthOneRival is a free data retrieval call binding the contract method 0x54b64151.
//
// Solidity: function hasLengthOneRival(bytes32 edgeId) view returns(bool)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) HasLengthOneRival(edgeId [32]byte) (bool, error) {
	return _IEdgeChallengeManager.Contract.HasLengthOneRival(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// HasLengthOneRival is a free data retrieval call binding the contract method 0x54b64151.
//
// Solidity: function hasLengthOneRival(bytes32 edgeId) view returns(bool)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerSession) HasLengthOneRival(edgeId [32]byte) (bool, error) {
	return _IEdgeChallengeManager.Contract.HasLengthOneRival(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// HasRival is a free data retrieval call binding the contract method 0x908517e9.
//
// Solidity: function hasRival(bytes32 edgeId) view returns(bool)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCaller) HasRival(opts *bind.CallOpts, edgeId [32]byte) (bool, error) {
	var out []interface{}
	err := _IEdgeChallengeManager.contract.Call(opts, &out, "hasRival", edgeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRival is a free data retrieval call binding the contract method 0x908517e9.
//
// Solidity: function hasRival(bytes32 edgeId) view returns(bool)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) HasRival(edgeId [32]byte) (bool, error) {
	return _IEdgeChallengeManager.Contract.HasRival(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// HasRival is a free data retrieval call binding the contract method 0x908517e9.
//
// Solidity: function hasRival(bytes32 edgeId) view returns(bool)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerSession) HasRival(edgeId [32]byte) (bool, error) {
	return _IEdgeChallengeManager.Contract.HasRival(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// OneStepProofEntry is a free data retrieval call binding the contract method 0x48923bc5.
//
// Solidity: function oneStepProofEntry() view returns(address)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCaller) OneStepProofEntry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEdgeChallengeManager.contract.Call(opts, &out, "oneStepProofEntry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OneStepProofEntry is a free data retrieval call binding the contract method 0x48923bc5.
//
// Solidity: function oneStepProofEntry() view returns(address)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) OneStepProofEntry() (common.Address, error) {
	return _IEdgeChallengeManager.Contract.OneStepProofEntry(&_IEdgeChallengeManager.CallOpts)
}

// OneStepProofEntry is a free data retrieval call binding the contract method 0x48923bc5.
//
// Solidity: function oneStepProofEntry() view returns(address)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerSession) OneStepProofEntry() (common.Address, error) {
	return _IEdgeChallengeManager.Contract.OneStepProofEntry(&_IEdgeChallengeManager.CallOpts)
}

// TimeUnrivaled is a free data retrieval call binding the contract method 0x3e35f5e8.
//
// Solidity: function timeUnrivaled(bytes32 edgeId) view returns(uint64)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCaller) TimeUnrivaled(opts *bind.CallOpts, edgeId [32]byte) (uint64, error) {
	var out []interface{}
	err := _IEdgeChallengeManager.contract.Call(opts, &out, "timeUnrivaled", edgeId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TimeUnrivaled is a free data retrieval call binding the contract method 0x3e35f5e8.
//
// Solidity: function timeUnrivaled(bytes32 edgeId) view returns(uint64)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) TimeUnrivaled(edgeId [32]byte) (uint64, error) {
	return _IEdgeChallengeManager.Contract.TimeUnrivaled(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// TimeUnrivaled is a free data retrieval call binding the contract method 0x3e35f5e8.
//
// Solidity: function timeUnrivaled(bytes32 edgeId) view returns(uint64)
func (_IEdgeChallengeManager *IEdgeChallengeManagerCallerSession) TimeUnrivaled(edgeId [32]byte) (uint64, error) {
	return _IEdgeChallengeManager.Contract.TimeUnrivaled(&_IEdgeChallengeManager.CallOpts, edgeId)
}

// BisectEdge is a paid mutator transaction binding the contract method 0xc8bc4e43.
//
// Solidity: function bisectEdge(bytes32 edgeId, bytes32 bisectionHistoryRoot, bytes prefixProof) returns(bytes32, bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactor) BisectEdge(opts *bind.TransactOpts, edgeId [32]byte, bisectionHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.contract.Transact(opts, "bisectEdge", edgeId, bisectionHistoryRoot, prefixProof)
}

// BisectEdge is a paid mutator transaction binding the contract method 0xc8bc4e43.
//
// Solidity: function bisectEdge(bytes32 edgeId, bytes32 bisectionHistoryRoot, bytes prefixProof) returns(bytes32, bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) BisectEdge(edgeId [32]byte, bisectionHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.BisectEdge(&_IEdgeChallengeManager.TransactOpts, edgeId, bisectionHistoryRoot, prefixProof)
}

// BisectEdge is a paid mutator transaction binding the contract method 0xc8bc4e43.
//
// Solidity: function bisectEdge(bytes32 edgeId, bytes32 bisectionHistoryRoot, bytes prefixProof) returns(bytes32, bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactorSession) BisectEdge(edgeId [32]byte, bisectionHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.BisectEdge(&_IEdgeChallengeManager.TransactOpts, edgeId, bisectionHistoryRoot, prefixProof)
}

// ConfirmEdgeByChildren is a paid mutator transaction binding the contract method 0x2eaa0043.
//
// Solidity: function confirmEdgeByChildren(bytes32 edgeId) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactor) ConfirmEdgeByChildren(opts *bind.TransactOpts, edgeId [32]byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.contract.Transact(opts, "confirmEdgeByChildren", edgeId)
}

// ConfirmEdgeByChildren is a paid mutator transaction binding the contract method 0x2eaa0043.
//
// Solidity: function confirmEdgeByChildren(bytes32 edgeId) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) ConfirmEdgeByChildren(edgeId [32]byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.ConfirmEdgeByChildren(&_IEdgeChallengeManager.TransactOpts, edgeId)
}

// ConfirmEdgeByChildren is a paid mutator transaction binding the contract method 0x2eaa0043.
//
// Solidity: function confirmEdgeByChildren(bytes32 edgeId) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactorSession) ConfirmEdgeByChildren(edgeId [32]byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.ConfirmEdgeByChildren(&_IEdgeChallengeManager.TransactOpts, edgeId)
}

// ConfirmEdgeByClaim is a paid mutator transaction binding the contract method 0x0f73bfad.
//
// Solidity: function confirmEdgeByClaim(bytes32 edgeId, bytes32 claimingEdgeId) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactor) ConfirmEdgeByClaim(opts *bind.TransactOpts, edgeId [32]byte, claimingEdgeId [32]byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.contract.Transact(opts, "confirmEdgeByClaim", edgeId, claimingEdgeId)
}

// ConfirmEdgeByClaim is a paid mutator transaction binding the contract method 0x0f73bfad.
//
// Solidity: function confirmEdgeByClaim(bytes32 edgeId, bytes32 claimingEdgeId) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) ConfirmEdgeByClaim(edgeId [32]byte, claimingEdgeId [32]byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.ConfirmEdgeByClaim(&_IEdgeChallengeManager.TransactOpts, edgeId, claimingEdgeId)
}

// ConfirmEdgeByClaim is a paid mutator transaction binding the contract method 0x0f73bfad.
//
// Solidity: function confirmEdgeByClaim(bytes32 edgeId, bytes32 claimingEdgeId) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactorSession) ConfirmEdgeByClaim(edgeId [32]byte, claimingEdgeId [32]byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.ConfirmEdgeByClaim(&_IEdgeChallengeManager.TransactOpts, edgeId, claimingEdgeId)
}

// ConfirmEdgeByOneStepProof is a paid mutator transaction binding the contract method 0x8c1b3a40.
//
// Solidity: function confirmEdgeByOneStepProof(bytes32 edgeId, (bytes32,bytes) oneStepData, (bytes32,uint256,address,uint64,uint64) prevConfig, bytes32[] beforeHistoryInclusionProof, bytes32[] afterHistoryInclusionProof) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactor) ConfirmEdgeByOneStepProof(opts *bind.TransactOpts, edgeId [32]byte, oneStepData OneStepData, prevConfig ConfigData, beforeHistoryInclusionProof [][32]byte, afterHistoryInclusionProof [][32]byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.contract.Transact(opts, "confirmEdgeByOneStepProof", edgeId, oneStepData, prevConfig, beforeHistoryInclusionProof, afterHistoryInclusionProof)
}

// ConfirmEdgeByOneStepProof is a paid mutator transaction binding the contract method 0x8c1b3a40.
//
// Solidity: function confirmEdgeByOneStepProof(bytes32 edgeId, (bytes32,bytes) oneStepData, (bytes32,uint256,address,uint64,uint64) prevConfig, bytes32[] beforeHistoryInclusionProof, bytes32[] afterHistoryInclusionProof) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) ConfirmEdgeByOneStepProof(edgeId [32]byte, oneStepData OneStepData, prevConfig ConfigData, beforeHistoryInclusionProof [][32]byte, afterHistoryInclusionProof [][32]byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.ConfirmEdgeByOneStepProof(&_IEdgeChallengeManager.TransactOpts, edgeId, oneStepData, prevConfig, beforeHistoryInclusionProof, afterHistoryInclusionProof)
}

// ConfirmEdgeByOneStepProof is a paid mutator transaction binding the contract method 0x8c1b3a40.
//
// Solidity: function confirmEdgeByOneStepProof(bytes32 edgeId, (bytes32,bytes) oneStepData, (bytes32,uint256,address,uint64,uint64) prevConfig, bytes32[] beforeHistoryInclusionProof, bytes32[] afterHistoryInclusionProof) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactorSession) ConfirmEdgeByOneStepProof(edgeId [32]byte, oneStepData OneStepData, prevConfig ConfigData, beforeHistoryInclusionProof [][32]byte, afterHistoryInclusionProof [][32]byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.ConfirmEdgeByOneStepProof(&_IEdgeChallengeManager.TransactOpts, edgeId, oneStepData, prevConfig, beforeHistoryInclusionProof, afterHistoryInclusionProof)
}

// ConfirmEdgeByTime is a paid mutator transaction binding the contract method 0x64deed59.
//
// Solidity: function confirmEdgeByTime(bytes32 edgeId, bytes32[] ancestorEdgeIds, (((bytes32[2],uint64[2]),uint8),bytes32,bytes32) claimStateData) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactor) ConfirmEdgeByTime(opts *bind.TransactOpts, edgeId [32]byte, ancestorEdgeIds [][32]byte, claimStateData ExecutionStateData) (*types.Transaction, error) {
	return _IEdgeChallengeManager.contract.Transact(opts, "confirmEdgeByTime", edgeId, ancestorEdgeIds, claimStateData)
}

// ConfirmEdgeByTime is a paid mutator transaction binding the contract method 0x64deed59.
//
// Solidity: function confirmEdgeByTime(bytes32 edgeId, bytes32[] ancestorEdgeIds, (((bytes32[2],uint64[2]),uint8),bytes32,bytes32) claimStateData) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) ConfirmEdgeByTime(edgeId [32]byte, ancestorEdgeIds [][32]byte, claimStateData ExecutionStateData) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.ConfirmEdgeByTime(&_IEdgeChallengeManager.TransactOpts, edgeId, ancestorEdgeIds, claimStateData)
}

// ConfirmEdgeByTime is a paid mutator transaction binding the contract method 0x64deed59.
//
// Solidity: function confirmEdgeByTime(bytes32 edgeId, bytes32[] ancestorEdgeIds, (((bytes32[2],uint64[2]),uint8),bytes32,bytes32) claimStateData) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactorSession) ConfirmEdgeByTime(edgeId [32]byte, ancestorEdgeIds [][32]byte, claimStateData ExecutionStateData) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.ConfirmEdgeByTime(&_IEdgeChallengeManager.TransactOpts, edgeId, ancestorEdgeIds, claimStateData)
}

// CreateLayerZeroEdge is a paid mutator transaction binding the contract method 0x05fae141.
//
// Solidity: function createLayerZeroEdge((uint8,bytes32,uint256,bytes32,bytes,bytes) args) returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactor) CreateLayerZeroEdge(opts *bind.TransactOpts, args CreateEdgeArgs) (*types.Transaction, error) {
	return _IEdgeChallengeManager.contract.Transact(opts, "createLayerZeroEdge", args)
}

// CreateLayerZeroEdge is a paid mutator transaction binding the contract method 0x05fae141.
//
// Solidity: function createLayerZeroEdge((uint8,bytes32,uint256,bytes32,bytes,bytes) args) returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) CreateLayerZeroEdge(args CreateEdgeArgs) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.CreateLayerZeroEdge(&_IEdgeChallengeManager.TransactOpts, args)
}

// CreateLayerZeroEdge is a paid mutator transaction binding the contract method 0x05fae141.
//
// Solidity: function createLayerZeroEdge((uint8,bytes32,uint256,bytes32,bytes,bytes) args) returns(bytes32)
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactorSession) CreateLayerZeroEdge(args CreateEdgeArgs) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.CreateLayerZeroEdge(&_IEdgeChallengeManager.TransactOpts, args)
}

// RefundStake is a paid mutator transaction binding the contract method 0x748926f3.
//
// Solidity: function refundStake(bytes32 edgeId) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactor) RefundStake(opts *bind.TransactOpts, edgeId [32]byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.contract.Transact(opts, "refundStake", edgeId)
}

// RefundStake is a paid mutator transaction binding the contract method 0x748926f3.
//
// Solidity: function refundStake(bytes32 edgeId) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerSession) RefundStake(edgeId [32]byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.RefundStake(&_IEdgeChallengeManager.TransactOpts, edgeId)
}

// RefundStake is a paid mutator transaction binding the contract method 0x748926f3.
//
// Solidity: function refundStake(bytes32 edgeId) returns()
func (_IEdgeChallengeManager *IEdgeChallengeManagerTransactorSession) RefundStake(edgeId [32]byte) (*types.Transaction, error) {
	return _IEdgeChallengeManager.Contract.RefundStake(&_IEdgeChallengeManager.TransactOpts, edgeId)
}
