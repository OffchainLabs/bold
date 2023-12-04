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
	Bin: "0x6101c06040523480156200001257600080fd5b50604051620067a5380380620067a5833981016040819052620000359162000285565b6001600160a01b038a166200005d5760405163641f043160e11b815260040160405180910390fd5b6001600160a01b03808b166101005288166200008c5760405163fb60b0ef60e01b815260040160405180910390fd5b6001600160a01b038816610120526001600160401b038916600003620000c557604051632283bb7360e21b815260040160405180910390fd5b6001600160401b03891660e0526001600160a01b0380851660a05260c0849052821662000105576040516301e1d91560e31b815260040160405180910390fd5b6001600160a01b03821660805262000129876200023f602090811b620018df17901c565b6200014f57604051633abfb6ff60e21b8152600481018890526024015b60405180910390fd5b8661014081815250506200016e866200023f60201b620018df1760201c565b6200019057604051633abfb6ff60e21b81526004810187905260240162000146565b856101608181525050620001af856200023f60201b620018df1760201c565b620001d157604051633abfb6ff60e21b81526004810186905260240162000146565b61018085905260ff8116600003620001fc57604051632a18f5b960e21b815260040160405180910390fd5b60fd8160ff161115620002285760405163040d23bf60e41b815260ff8216600482015260240162000146565b60ff166101a052506200037c975050505050505050565b6000816000036200025257506000919050565b60006200026160018462000354565b929092161592915050565b6001600160a01b03811681146200028257600080fd5b50565b6000806000806000806000806000806101408b8d031215620002a657600080fd5b8a51620002b3816200026c565b60208c0151909a506001600160401b0381168114620002d157600080fd5b60408c0151909950620002e4816200026c565b8098505060608b0151965060808b0151955060a08b0151945060c08b01516200030d816200026c565b60e08c01516101008d0151919550935062000328816200026c565b6101208c015190925060ff811681146200034157600080fd5b809150509295989b9194979a5092959850565b818103818111156200037657634e487b7160e01b600052601160045260246000fd5b92915050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516162e2620004c3600039600081816103b0015281816106bf01528181610a1b01528181610a7301528181610be101528181610e38015281816111dd015261151c01526000818161058e0152610d5901526000818161028c0152610d190152600081816102260152610cd90152600081816102ed015281816109f901528181610a5101526114f501526000818161032c015281816107880152818161081b015281816108a50152818161094101528181610f2a01528181610fdf0152818161108101528181611128015281816113c001526114580152600081816102c601526111bc0152600081816103e901528181610abd015261129f01526000818161035301528181610a9c015261127e0152600081816105540152610b0b01526162e26000f3fe608060405234801561001057600080fd5b50600436106101ce5760003560e01c80635d9e244411610104578063bce6f54f116100a2578063e94e051e11610071578063e94e051e1461054f578063eae0328b14610576578063f8ee77d614610589578063fda2892e146105b057600080fd5b8063bce6f54f1461046a578063c32d8c631461048a578063c8bc4e4314610507578063e5b123da1461052f57600080fd5b8063748926f3116100de578063748926f31461041e578063750e0c0f146104315780638c1b3a4014610444578063908517e91461045757600080fd5b80635d9e2444146103ab57806360c7dc47146103e457806364deed591461040b57600080fd5b806342e1aaa81161017157806348dd29241161014b57806348dd29241461032757806351ed6a301461034e57806354b64151146103755780635a48e0f41461039857600080fd5b806342e1aaa8146102ae57806346c2781a146102c157806348923bc5146102e857600080fd5b80631dce5166116101ad5780631dce5166146102215780632eaa0043146102485780633e35f5e81461025b578063416e66571461028757600080fd5b80624d8efe146101d357806305fae141146101f95780630f73bfad1461020c575b600080fd5b6101e66101e136600461545c565b6105d0565b6040519081526020015b60405180910390f35b6101e66102073660046154a6565b61066b565b61021f61021a3660046154e1565b610bd8565b005b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b61021f610256366004615503565b610c5a565b61026e610269366004615503565b610caa565b60405167ffffffffffffffff90911681526020016101f0565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b6101e66102bc366004615537565b610cbc565b61026e7f000000000000000000000000000000000000000000000000000000000000000081565b61030f7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101f0565b61030f7f000000000000000000000000000000000000000000000000000000000000000081565b61030f7f000000000000000000000000000000000000000000000000000000000000000081565b610388610383366004615503565b610dbf565b60405190151581526020016101f0565b6101e66103a6366004615503565b610dcb565b6103d27f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff90911681526020016101f0565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b61021f610419366004615676565b610dd7565b61021f61042c366004615503565b611265565b61038861043f366004615503565b611357565b61021f610452366004615726565b611381565b610388610465366004615503565b61158d565b6101e6610478366004615503565b60009081526001602052604090205490565b6101e66104983660046157ef565b6040805160f89690961b7fff00000000000000000000000000000000000000000000000000000000000000166020808801919091526021870195909552604186019390935260618501919091526081808501919091528151808503909101815260a19093019052815191012090565b61051a610515366004615831565b611599565b604080519283526020830191909152016101f0565b6101e661053d366004615503565b60009081526002602052604090205490565b61030f7f000000000000000000000000000000000000000000000000000000000000000081565b6101e6610584366004615503565b611750565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b6105c36105be366004615503565b611764565b6040516101f091906158db565b6040805160f888901b7fff000000000000000000000000000000000000000000000000000000000000001660208083019190915260218201889052604182018790526061820186905260818083018690528351808403909101815260a18301845280519082012060c183015260e18083018590528351808403909101815261010190920190925280519101206000905b979650505050505050565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905260006106e36106bd60208601866159b1565b7f0000000000000000000000000000000000000000000000000000000000000000611909565b905060006106f082610cbc565b90506106fa6153b0565b600083600281111561070e5761070e6158b1565b03610a485761072060a08701876159cc565b905060000361074257604051630c9ccac560e41b815260040160405180910390fd5b60008061075260a08901896159cc565b81019061075f9190615b6c565b80516020820151604080840151905163f9cee9df60e01b81529497509295506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016945063f9cee9df936107c49360608f0135939291600401615c5a565b60006040518083038186803b1580156107dc57600080fd5b505afa1580156107f0573d6000803e3d6000fd5b505050602080830151845191850151604080870151905163f9cee9df60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016955063f9cee9df94610854949390929091600401615c5a565b60006040518083038186803b15801561086c57600080fd5b505afa158015610880573d6000803e3d6000fd5b505050506040518060c0016040528089606001358152602001826020015181526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e531d8c78b606001356040518263ffffffff1660e01b81526004016108f591815260200190565b602060405180830381865afa158015610912573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109369190615c81565b1515815260200160007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166356bbc9e685602001516040518263ffffffff1660e01b815260040161099191815260200190565b602060405180830381865afa1580156109ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d29190615ca3565b67ffffffffffffffff161181528351602082015282516040909101529250610a3f600089857f0000000000000000000000000000000000000000000000000000000000000000887f0000000000000000000000000000000000000000000000000000000000000000611991565b95505050610a9a565b610a97600087837f0000000000000000000000000000000000000000000000000000000000000000867f0000000000000000000000000000000000000000000000000000000000000000611991565b93505b7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03821615801590610af357508015155b15610b445760008660c00151610b095730610b2b565b7f00000000000000000000000000000000000000000000000000000000000000005b9050610b426001600160a01b038416338385611a1b565b505b8560400151866020015187600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a489606001518a608001518b60a001518c60c001518d60e00151604051610bc3959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a45050925195945050505050565b610c05600083837f0000000000000000000000000000000000000000000000000000000000000000611ad2565b6000828152602081905260409020610c1c90611c59565b827fb924f3aa473645c7cf5b10262f927ae4ccf869d7fc239c17144b0c67490d1c7383604051610c4e91815260200190565b60405180910390a35050565b610c65600082611ce8565b6000818152602081905260409020610c7c90611c59565b60405182907f0d27fcaf1adc41547a5cfc99d2364f6c0dc7e81c9fc3fe8cb38abb409b48358a90600090a350565b6000610cb68183611e9e565b92915050565b600080826002811115610cd157610cd16158b1565b03610cfd57507f0000000000000000000000000000000000000000000000000000000000000000919050565b6001826002811115610d1157610d116158b1565b03610d3d57507f0000000000000000000000000000000000000000000000000000000000000000919050565b6002826002811115610d5157610d516158b1565b03610d7d57507f0000000000000000000000000000000000000000000000000000000000000000919050565b816040517f3bf2e1ec000000000000000000000000000000000000000000000000000000008152600401610db19190615cc0565b60405180910390fd5b919050565b6000610cb6818361205b565b6000610cb68183612090565b600080835111610de75783610e0f565b8260018451610df69190615ce9565b81518110610e0657610e06615cfc565b60200260200101515b90506000610e1d81836120e1565b90506000610e5c8260090160099054906101000a900460ff167f0000000000000000000000000000000000000000000000000000000000000000611909565b90506000816002811115610e7257610e726158b1565b14610eba5760098201546040517fec72dc5d000000000000000000000000000000000000000000000000000000008152600160481b90910460ff166004820152602401610db1565b610ec382612136565b610f2557610ed08261215a565b600883015460078401546040517fe58c830800000000000000000000000000000000000000000000000000000000815260048101939093526001600160a01b0390911660248301526044820152606401610db1565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633083622885600701546040518263ffffffff1660e01b8152600401610f7a91815260200190565b602060405180830381865afa158015610f97573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fbb9190615c81565b905080156111ac57600784015460405163f9cee9df60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163f9cee9df9161102191908a9060a08201359060c083013590600401615d12565b60006040518083038186803b15801561103957600080fd5b505afa15801561104d573d6000803e3d6000fd5b50506040517f1171558500000000000000000000000000000000000000000000000000000000815260a089013560048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316925063117155859150602401602060405180830381865afa1580156110d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110f69190615ca3565b6040517f56bbc9e600000000000000000000000000000000000000000000000000000000815260a088013560048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906356bbc9e690602401602060405180830381865afa158015611177573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061119b9190615ca3565b6111a59190615d90565b91506111b1565b600091505b6000611201818a8a867f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000612213565b60008a815260208190526040902090915061121b90611c59565b60405167ffffffffffffffff831681528a907f9cd2c77f6772dd0fb07e9972aee8ddddf9fd13f1a65abb03a3de88b07dc59af99060200160405180910390a3505050505050505050565b600061127181836120e1565b905061127c81612497565b7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b038216158015906112d557508015155b156112f65760088301546112f6906001600160a01b038481169116836125dd565b600084815260208190526040902061130d90611c59565b604080516001600160a01b03851681526020810184905286917fa635398959ddb5ce3b14537edfc25b2e671274c9b8cad0f4bd634752e69007b6910160405180910390a350505050565b600081815260208190526040812060080154600160a01b900467ffffffffffffffff161515610cb6565b600061138d8189612090565b6040517f04972af90000000000000000000000000000000000000000000000000000000081529091506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906304972af9906113f79084908a90600401615dc6565b60006040518083038186803b15801561140f57600080fd5b505afa158015611423573d6000803e3d6000fd5b50505050600060405180606001604052808860800160208101906114479190615e40565b67ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e78cea926040518163ffffffff1660e01b8152600401602060405180830381865afa1580156114b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114d89190615e5d565b6001600160a01b031681528835602090910152905061154060008a7f00000000000000000000000000000000000000000000000000000000000000008b858b8b8b8b7f000000000000000000000000000000000000000000000000000000000000000061262b565b600089815260208190526040902061155790611c59565b6040518a907fe11db4b27bc8c6ea5943ecbb205ae1ca8d56c42c719717aaf8a53d43d0cee7c290600090a3505050505050505050565b6000610cb681836128a1565b60008060008060006115e5898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509594939250506129819050565b81519295509093509150158061167d578260400151836020015184600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a4866060015187608001518860a001518960c001518a60e00151604051611674959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a45b8160400151826020015183600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a4856060015186608001518760a001518860c001518960e001516040516116fc959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a48151604051821515815285908c907f7340510d24b7ec9b5c100f5500d93429d80d00d46f0d18e4e85d0c4cc22b99249060200160405180910390a4505191989197509095505050505050565b6000610cb661175f82846120e1565b612d89565b604080516101c081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a08101919091526117e06000836120e1565b604080516101c0810182528254815260018084015460208301526002840154928201929092526003830154606082015260048301546080820152600583015460a0820152600683015460c0820152600783015460e082015260088301546001600160a01b038116610100830152600160a01b900467ffffffffffffffff9081166101208301526009840154908116610140830152909291610160840191600160401b900460ff1690811115611897576118976158b1565b60018111156118a8576118a86158b1565b81526009919091015460ff600160481b8204811660208401526a010000000000000000000090910416151560409091015292915050565b6000816000036118f157506000919050565b60006118fe600184615ce9565b929092161592915050565b60008260ff1660000361191e57506000610cb6565b8160ff168360ff161161193357506001610cb6565b61193e826001615e7a565b60ff168360ff160361195257506002610cb6565b6040517fae0da57800000000000000000000000000000000000000000000000000000000815260ff808516600483015283166024820152604401610db1565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905290806119e08989898988612dce565b9150915060006119f1838a8861333c565b90506000611a0083838c61349a565b9050611a0c8b8261353e565b9b9a5050505050505050505050565b6040516001600160a01b0380851660248301528316604482015260648101829052611acc9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152613914565b50505050565b600083815260208590526040902060080154600160a01b900467ffffffffffffffff16611b145760405162a7b02b60e01b815260048101849052602401610db1565b600082815260208590526040902060080154600160a01b900467ffffffffffffffff16611b565760405162a7b02b60e01b815260048101849052602401610db1565b6001600083815260208690526040902060090154600160401b900460ff166001811115611b8557611b856158b1565b14611bc45760008281526020859052604090819020600901549051633bc499ed60e21b8152610db1918491600160401b90910460ff1690600401615e93565b611bd0848484846139f9565b6000828152602085905260409020600701548314611c3857600082815260208590526040908190206007015490517fc2adc3e8000000000000000000000000000000000000000000000000000000008152610db1918591600401918252602082015260400190565b611c428484613b53565b6000838152602085905260409020611acc90613bd6565b6009810154815460028301546001840154600485015460408051600160481b90960460f81b7fff00000000000000000000000000000000000000000000000000000000000000166020808801919091526021870195909552604186019390935260618501919091526081808501919091528151808503909101815260a190930190528151910120600090610cb6565b600081815260208390526040902060080154600160a01b900467ffffffffffffffff16611d2a5760405162a7b02b60e01b815260048101829052602401610db1565b60008181526020839052604080822060050154808352912060080154600160a01b900467ffffffffffffffff16611d765760405162a7b02b60e01b815260048101829052602401610db1565b6001600082815260208590526040902060090154600160401b900460ff166001811115611da557611da56158b1565b14611de45760008181526020849052604090819020600901549051633bc499ed60e21b8152610db1918391600160401b90910460ff1690600401615e93565b60008281526020849052604080822060060154808352912060080154600160a01b900467ffffffffffffffff16611e305760405162a7b02b60e01b815260048101829052602401610db1565b6001600082815260208690526040902060090154600160401b900460ff166001811115611e5f57611e5f6158b1565b14611c385760008181526020859052604090819020600901549051633bc499ed60e21b8152610db1918391600160401b90910460ff1690600401615e93565b600081815260208390526040812060080154600160a01b900467ffffffffffffffff16611ee05760405162a7b02b60e01b815260048101839052602401610db1565b6000828152602084905260408120611ef790611c59565b6000818152600186016020526040812054919250819003611f44576040517fda10f67c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040516815539492559053115160ba1b6020820152602901604051602081830303815290604052805190602001208103611fae57600084815260208690526040902060080154611fa590600160a01b900467ffffffffffffffff1643615d90565b92505050610cb6565b600081815260208690526040902060080154600160a01b900467ffffffffffffffff16611ff05760405162a7b02b60e01b815260048101829052602401610db1565b600081815260208690526040808220600890810154878452919092209091015467ffffffffffffffff600160a01b9283900481169290910416808211156120465761203b8183615d90565b945050505050610cb6565b6000945050505050610cb6565b505092915050565b600061206783836128a1565b80156120895750600082815260208490526040902061208590612d89565b6001145b9392505050565b60008061209d84846120e1565b90505b6009810154600160481b900460ff16156120d957805460009081526001850160205260409020546120d185826120e1565b9150506120a0565b549392505050565b600081815260208390526040812060080154600160a01b900467ffffffffffffffff166121235760405162a7b02b60e01b815260048101839052602401610db1565b5060009081526020919091526040902090565b600781015460009015801590610cb6575050600801546001600160a01b0316151590565b60098101548154600283015460018401546004850154600386015460408051600160481b90970460f81b7fff00000000000000000000000000000000000000000000000000000000000000166020808901919091526021880196909652604187019490945260618601929092526081808601919091528251808603909101815260a18501835280519084012060c185015260e1808501919091528151808503909101815261010190930190528151910120600090610cb6565b600085815260208790526040812060080154600160a01b900467ffffffffffffffff166122555760405162a7b02b60e01b815260048101879052602401610db1565b8560006122628983611e9e565b905060005b87518110156123fb5760006122958b8a848151811061228857612288615cfc565b60200260200101516120e1565b905083816005015414806122ac5750838160060154145b156122f0576122c38b6122be8361215a565b611e9e565b6122cd9084615ea7565b92508882815181106122e1576122e1615cfc565b602002602001015193506123e8565b600084815260208c9052604090206007015489518a908490811061231657612316615cfc565b602002602001015103612355576123488b8a848151811061233957612339615cfc565b602002602001015186896139f9565b6122c38b6122be8361215a565b83816005015482600601548b858151811061237257612372615cfc565b60200260200101518e6000016000898152602001908152602001600020600701546040517f6ebd28c9000000000000000000000000000000000000000000000000000000008152600401610db1959493929190948552602085019390935260408401919091526060830152608082015260a00190565b50806123f381615ec8565b915050612267565b506124068682615ea7565b90508467ffffffffffffffff168167ffffffffffffffff16101561246a576040517f11a8d4d000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808316600483015286166024820152604401610db1565b6124748989613b53565b600088815260208a90526040902061248b90613bd6565b98975050505050505050565b60016009820154600160401b900460ff1660018111156124b9576124b96158b1565b146124f1576124c78161215a565b6009820154604051633bc499ed60e21b8152610db19291600160401b900460ff1690600401615e93565b6124fa81612136565b61255c576125078161215a565b600882015460078301546040517fe58c830800000000000000000000000000000000000000000000000000000000815260048101939093526001600160a01b0390911660248301526044820152606401610db1565b60098101546a0100000000000000000000900460ff1615156001036125b9576125848161215a565b6040517f307f7669000000000000000000000000000000000000000000000000000000008152600401610db191815260200190565b60090180546aff0000000000000000000019166a0100000000000000000000179055565b6040516001600160a01b0383166024820152604481018290526126269084907fa9059cbb0000000000000000000000000000000000000000000000000000000090606401611a68565b505050565b60006126378b8b6120e1565b600290810154915060008b815260208d9052604090206009015461266590600160481b900460ff1684611909565b6002811115612676576126766158b1565b146126cc5760008a815260208c90526040908190206009015490517f348aefdf000000000000000000000000000000000000000000000000000000008152600160481b90910460ff166004820152602401610db1565b60008a815260208c9052604090206126e390612d89565b6001146127365760008a815260208c90526040902061270190612d89565b6040517f6b595e50000000000000000000000000000000000000000000000000000000008152600401610db191815260200190565b6127918b60000160008c815260200190815260200160002060010154896000013583898980806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250613c7292505050565b60006001600160a01b038a1663b5112fd289848c356127b360208f018f6159cc565b6040518663ffffffff1660e01b81526004016127d3959493929190615ee2565b602060405180830381865afa1580156127f0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128149190615f44565b60008c815260208e905260409020600301549091506128729082612839856001615f5d565b888880806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250613c7292505050565b61287c8c8c613b53565b60008b815260208d90526040902061289390613bd6565b505050505050505050505050565b600081815260208390526040812060080154600160a01b900467ffffffffffffffff166128e35760405162a7b02b60e01b815260048101839052602401610db1565b60008281526020849052604081206128fa90611c59565b6000818152600186016020526040812054919250819003612947576040517fda10f67c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040516815539492559053115160ba1b602082015260290160408051601f1981840301815291905280516020909101201415949350505050565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915260008087815260208990526040902060090154600160401b900460ff166001811115612a3057612a306158b1565b14612a8857600086815260208890526040908190206009015490517f23f8405d000000000000000000000000000000000000000000000000000000008152610db1918891600160401b90910460ff1690600401615e93565b612a9287876128a1565b612acb576040517f80e07e4500000000000000000000000000000000000000000000000000000000815260048101879052602401610db1565b6000868152602088905260408120604080516101c0810182528254815260018084015460208301526002840154928201929092526003830154606082015260048301546080820152600583015460a0820152600683015460c0820152600783015460e082015260088301546001600160a01b038116610100830152600160a01b900467ffffffffffffffff9081166101208301526009840154908116610140830152909291610160840191600160401b900460ff1690811115612b9057612b906158b1565b6001811115612ba157612ba16158b1565b81526009919091015460ff600160481b8204811660208401526a01000000000000000000009091041615156040918201528101516080820151919250600091612bea9190613cff565b905060008087806020019051810190612c039190615fcb565b9092509050612c3389612c17856001615f5d565b60608701516080880151612c2c906001615f5d565b8686613dac565b50506040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526000612c968560000151866020015187604001518d888a61018001516140cf565b9050612ca1816141c9565b600081815260208e90526040902060080154909350600160a01b900467ffffffffffffffff16612cd857612cd58c8261353e565b91505b506040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101919091526000612d3b86600001518c8789606001518a608001518b61018001516140cf565b9050612d478d8261353e565b915050612d778382600001518e60000160008f81526020019081526020016000206142799092919063ffffffff16565b919b909a509098509650505050505050565b60008082600201548360040154612da09190615ce9565b905080600003610cb657612db38361215a565b60405162a7b02b60e01b8152600401610db191815260200190565b60408051606080820183526000808352602083015291810191909152600080612e03612dfd60208901896159b1565b85611909565b6002811115612e1457612e146158b1565b0361313c5760208501518551600003612e59576040517fe96bc61a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8551606088013514612ea75785516040517f5b177a3c000000000000000000000000000000000000000000000000000000008152600481019190915260608801356024820152604401610db1565b8560400151612ee2576040517fc169243600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8560600151612f1d576040517fb45d1c3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612f2a60a08801886159cc565b9050600003612f4c57604051630c9ccac560e41b815260040160405180910390fd5b6000612f5b60a08901896159cc565b810190612f689190615b6c565b50909150600090508760800151602001516002811115612f8a57612f8a6158b1565b03612fc1576040517f46365e5200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008760a00151602001516002811115612fdd57612fdd6158b1565b03613014576040517f8999857d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60808701516040517fc39619c40000000000000000000000000000000000000000000000000000000081526000916001600160a01b0389169163c39619c49161305f9160040161602f565b602060405180830381865afa15801561307c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130a09190615f44565b90506000876001600160a01b031663c39619c48a60a001516040518263ffffffff1660e01b81526004016130d4919061602f565b602060405180830381865afa1580156130f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131159190615f44565b60408051606081018252938452602084019190915282019290925293509091506133329050565b61314a87876060013561205b565b613186576040517fff6d9bd700000000000000000000000000000000000000000000000000000000815260608701356004820152602401610db1565b60608601356000908152602088905260408120906131a382611c59565b905060006009830154600160401b900460ff1660018111156131c7576131c76158b1565b146131fe576040517f12459ffd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600982015461321790600160481b900460ff16866142f9565b60ff1661322760208a018a6159b1565b60ff16146132895761323c60208901896159b1565b60098301546040517fac9e611600000000000000000000000000000000000000000000000000000000815260ff9283166004820152600160481b9091049091166024820152604401610db1565b61329660a08901896159cc565b90506000036132b857604051630c9ccac560e41b815260040160405180910390fd5b6000808080806132cb60a08e018e6159cc565b8101906132d8919061603d565b945094509450945094506132f6876001015486896002015486613c72565b61330a876003015485896004015485613c72565b6040518060600160405280868152602001858152602001828152508698509850505050505050505b9550959350505050565b6040805160008082526020820190925281906133629061335d90875161431b565b614351565b905061336d836118df565b6133a6576040517feafedbfc00000000000000000000000000000000000000000000000000000000815260048101849052602401610db1565b828460400135146133f057604080517fdfcc62bc00000000000000000000000000000000000000000000000000000000815290850135600482015260248101849052604401610db1565b61340c8460200135866020015186604001358860400151613c72565b61341960808501856159cc565b9050600003613454576040517f342a075200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008061346460808701876159cc565b81019061347191906160d9565b909250905061348f8360016020890135612c2c60408b013583615f5d565b509095945050505050565b604080516101c081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081019190915261353684846000602086018035906040880135906060890135903390613531908b6159b1565b6144f1565b949350505050565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905290613588836141c9565b600081815260208690526040902060080154909150600160a01b900467ffffffffffffffff16156135e8576040517fbcedf3de00000000000000000000000000000000000000000000000000000000815260048101829052602401610db1565b600081815260208581526040918290208551815590850151600180830191909155918501516002820155606085015160038201556080850151600482015560a0850151600582015560c0850151600682015560e0850151600782015561010085015160088201805461012088015167ffffffffffffffff908116600160a01b027fffffffff000000000000000000000000000000000000000000000000000000009092166001600160a01b0390941693909317179055610140860151600983018054919092167fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000821681178355610160880151889590939268ffffffffffffffffff191690911790600160401b908490811115613707576137076158b1565b02179055506101808281015160099290920180546101a0909401517fffffffffffffffffffffffffffffffffffffffffff0000ffffffffffffffffff909416600160481b60ff909416939093026aff000000000000000000001916929092176a01000000000000000000009315159390930292909217905583015183516040808601516020808801516080890151845160f89790971b7fff0000000000000000000000000000000000000000000000000000000000000016878401526021870195909552604186019290925260618501919091526081808501939093528151808503909301835260a190930190528051910120600090600081815260018701602052604081205491925081900361385c576040516815539492559053115160ba1b602082015260290160408051601f198184030181529181528151602092830120600085815260018a019093529120556138a4565b6040516815539492559053115160ba1b60208201526029016040516020818303038152906040528051906020012081036138a457600082815260018701602052604090208390555b6040805161010081018252848152602080820185905287518284015260e088015160608301526000868152908990529190912060808201906138e590612d89565b815261018087015160ff166020820152911515604083015260e090950151151560609091015250919392505050565b6000613969826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166146909092919063ffffffff16565b80519091501561262657808060200190518101906139879190615c81565b6126265760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610db1565b600082815260208590526040808220548583529120613a1790611c59565b14613a7e576000838152602085905260409020613a3390611c59565b600083815260208690526040908190205490517fe2e27f8700000000000000000000000000000000000000000000000000000000815260048101929092526024820152604401610db1565b600082815260208590526040808220600990810154868452919092209091015460ff600160481b92839004811692613ab8920416836142f9565b60ff1614611acc5760008381526020859052604090206009015483908390613aea90600160481b900460ff16846142f9565b600085815260208890526040908190206009015490517f7e726d150000000000000000000000000000000000000000000000000000000081526004810194909452602484019290925260ff9081166044840152600160481b909104166064820152608401610db1565b6000818152602083905260408120613b6a90611c59565b60008181526002850160205260409020549091508015613bc0576040517fdd7028f00000000000000000000000000000000000000000000000000000000081526004810184905260248101829052604401610db1565b5060009081526002909201602052604090912055565b60006009820154600160401b900460ff166001811115613bf857613bf86158b1565b14613c4957613c068161215a565b60098201546040517f23f8405d000000000000000000000000000000000000000000000000000000008152610db19291600160401b900460ff1690600401615e93565b600901805467ffffffffffffffff431668ffffffffffffffffff1990911617600160401b179055565b6000613ca7828486604051602001613c8c91815260200190565b6040516020818303038152906040528051906020012061469f565b9050808514613cf85760405162461bcd60e51b815260206004820152601760248201527f496e76616c696420696e636c7573696f6e2070726f6f660000000000000000006044820152606401610db1565b5050505050565b60006002613d0d8484615ce9565b1015613d4f576040517f902985940000000000000000000000000000000000000000000000000000000081526004810184905260248101839052604401610db1565b613d598383615ce9565b600203613d7257613d6b836001615f5d565b9050610cb6565b600083613d80600185615ce9565b1890506000613d8e8261475a565b9050600019811b80613da1600187615ce9565b169695505050505050565b60008511613dfc5760405162461bcd60e51b815260206004820152601460248201527f5072652d73697a652063616e6e6f7420626520300000000000000000000000006044820152606401610db1565b85613e0683614351565b14613e535760405162461bcd60e51b815260206004820152601b60248201527f50726520657870616e73696f6e20726f6f74206d69736d6174636800000000006044820152606401610db1565b84613e5d83614896565b14613ed05760405162461bcd60e51b815260206004820152602160248201527f5072652073697a6520646f6573206e6f74206d6174636820657870616e73696f60448201527f6e000000000000000000000000000000000000000000000000000000000000006064820152608401610db1565b828510613f1f5760405162461bcd60e51b815260206004820181905260248201527f5072652073697a65206e6f74206c657373207468616e20706f73742073697a656044820152606401610db1565b6000859050600080613f3485600087516148f1565b90505b85831015613ff7576000613f4b8488614a70565b905084518310613f9d5760405162461bcd60e51b815260206004820152601260248201527f496e646578206f7574206f662072616e676500000000000000000000000000006044820152606401610db1565b613fc18282878681518110613fb457613fb4615cfc565b6020026020010151614b5a565b91506001811b613fd18186615f5d565b945087851115613fe357613fe3616133565b83613fed81615ec8565b9450505050613f37565b8661400182614351565b146140745760405162461bcd60e51b815260206004820152602260248201527f506f737420657870616e73696f6e20726f6f74206e6f7420657175616c20706f60448201527f73740000000000000000000000000000000000000000000000000000000000006064820152608401610db1565b835182146140c45760405162461bcd60e51b815260206004820152601660248201527f496e636f6d706c6574652070726f6f66207573616765000000000000000000006044820152606401610db1565b505050505050505050565b604080516101c081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081019190915261414d87878787876150d9565b50604080516101c08101825296875260208701959095529385019290925260608401526080830152600060a0830181905260c0830181905260e08301819052610100830181905267ffffffffffffffff43166101208401526101408301819052610160830181905260ff9091166101808301526101a082015290565b610180810151815160408084015160208086015160808701516060880151855160f89890981b7fff0000000000000000000000000000000000000000000000000000000000000016888501526021880196909652604187019390935260618601526081808601929092528251808603909201825260a18501835281519181019190912060c185015260e1808501939093528151808503909301835261010190930190528051910120600090610cb6565b600583015415158061428e5750600683015415155b156142e95761429c8361215a565b600584015460068501546040517f8b0e71d0000000000000000000000000000000000000000000000000000000008152600481019390935260248301919091526044820152606401610db1565b6005830191909155600690910155565b600080614307846001615e7a565b90506143138184611909565b509392505050565b60606120898360008460405160200161433691815260200190565b60405160208183030381529060405280519060200120614b5a565b6000808251116143a35760405162461bcd60e51b815260206004820152601660248201527f456d707479206d65726b6c6520657870616e73696f6e000000000000000000006044820152606401610db1565b6040825111156143f55760405162461bcd60e51b815260206004820152601a60248201527f4d65726b6c6520657870616e73696f6e20746f6f206c617267650000000000006044820152606401610db1565b6000805b83518110156144ea57600084828151811061441657614416615cfc565b60200260200101519050826000801b0361448257801561447d57809250600185516144419190615ce9565b821461447d57604051614464908490600090602001918252602082015260400190565b6040516020818303038152906040528051906020012092505b6144d7565b80156144a1576040805160208101839052908101849052606001614464565b6040516144be908490600090602001918252602082015260400190565b6040516020818303038152906040528051906020012092505b50806144e281615ec8565b9150506143f9565b5092915050565b604080516101c081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a08101919091526001600160a01b0383166145a2576040517ff289e65700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008490036145dd576040517f6932bcfd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6145ea89898989896150d9565b604051806101c001604052808a81526020018981526020018881526020018781526020018681526020016000801b81526020016000801b8152602001858152602001846001600160a01b031681526020014367ffffffffffffffff168152602001600067ffffffffffffffff1681526020016000600181111561466f5761466f6158b1565b815260ff841660208201526000604090910152905098975050505050505050565b606061353684846000856151cd565b82516000906101008111156146eb576040517ffdac331e000000000000000000000000000000000000000000000000000000008152600481018290526101006024820152604401610db1565b8260005b8281101561475057600087828151811061470b5761470b615cfc565b60200260200101519050816001901b871660000361473757826000528060205260406000209250614747565b8060005282602052604060002092505b506001016146ef565b5095945050505050565b6000816000036147ac5760405162461bcd60e51b815260206004820152601c60248201527f5a65726f20686173206e6f207369676e69666963616e742062697473000000006044820152606401610db1565b70010000000000000000000000000000000082106147d757608091821c916147d49082615f5d565b90505b600160401b82106147f557604091821c916147f29082615f5d565b90505b640100000000821061481457602091821c916148119082615f5d565b90505b62010000821061483157601091821c9161482e9082615f5d565b90505b610100821061484d57600891821c9161484a9082615f5d565b90505b6010821061486857600491821c916148659082615f5d565b90505b6004821061488357600291821c916148809082615f5d565b90505b60028210610dba57610cb6600182615f5d565b600080805b83518110156144ea578381815181106148b6576148b6615cfc565b60200260200101516000801b146148df576148d281600261622d565b6148dc9083615f5d565b91505b806148e981615ec8565b91505061489b565b60608183106149425760405162461bcd60e51b815260206004820152601760248201527f5374617274206e6f74206c657373207468616e20656e640000000000000000006044820152606401610db1565b83518211156149b95760405162461bcd60e51b815260206004820152602160248201527f456e64206e6f74206c657373206f7220657175616c207468616e206c656e677460448201527f68000000000000000000000000000000000000000000000000000000000000006064820152608401610db1565b60006149c58484615ce9565b67ffffffffffffffff8111156149dd576149dd615554565b604051908082528060200260200182016040528015614a06578160200160208202803683370190505b509050835b83811015614a6757858181518110614a2557614a25615cfc565b6020026020010151828683614a3a9190615ce9565b81518110614a4a57614a4a615cfc565b602090810291909101015280614a5f81615ec8565b915050614a0b565b50949350505050565b6000818310614ac15760405162461bcd60e51b815260206004820152601760248201527f5374617274206e6f74206c657373207468616e20656e640000000000000000006044820152606401610db1565b6000614ace83851861475a565b905060006001614ade8382615f5d565b6001901b614aec9190615ce9565b90508481168482168115614b035761203b8261530a565b8015614b125761203b8161475a565b60405162461bcd60e51b815260206004820152601b60248201527f426f7468207920616e64207a2063616e6e6f74206265207a65726f00000000006044820152606401610db1565b606060408310614bac5760405162461bcd60e51b815260206004820152600e60248201527f4c6576656c20746f6f20686967680000000000000000000000000000000000006044820152606401610db1565b6000829003614bfd5760405162461bcd60e51b815260206004820152601b60248201527f43616e6e6f7420617070656e6420656d707479207375627472656500000000006044820152606401610db1565b604084511115614c4f5760405162461bcd60e51b815260206004820152601a60248201527f4d65726b6c6520657870616e73696f6e20746f6f206c617267650000000000006044820152606401610db1565b8351600003614cce576000614c65846001615f5d565b67ffffffffffffffff811115614c7d57614c7d615554565b604051908082528060200260200182016040528015614ca6578160200160208202803683370190505b50905082818581518110614cbc57614cbc615cfc565b60209081029190910101529050612089565b83518310614d445760405162461bcd60e51b815260206004820152603560248201527f4c6576656c2067726561746572207468616e2068696768657374206c6576656c60448201527f206f662063757272656e7420657870616e73696f6e00000000000000000000006064820152608401610db1565b816000614d5086614896565b90506000614d5f86600261622d565b614d699083615f5d565b90506000614d768361475a565b614d7f8361475a565b11614dcd57875167ffffffffffffffff811115614d9e57614d9e615554565b604051908082528060200260200182016040528015614dc7578160200160208202803683370190505b50614e1d565b8751614dda906001615f5d565b67ffffffffffffffff811115614df257614df2615554565b604051908082528060200260200182016040528015614e1b578160200160208202803683370190505b505b9050604081511115614e715760405162461bcd60e51b815260206004820152601c60248201527f417070656e642063726561746573206f76657273697a652074726565000000006044820152606401610db1565b60005b885181101561502d5787811015614f1b57888181518110614e9757614e97615cfc565b60200260200101516000801b14614f165760405162461bcd60e51b815260206004820152602260248201527f417070656e642061626f7665206c65617374207369676e69666963616e74206260448201527f69740000000000000000000000000000000000000000000000000000000000006064820152608401610db1565b61501b565b6000859003614f6157888181518110614f3657614f36615cfc565b6020026020010151828281518110614f5057614f50615cfc565b60200260200101818152505061501b565b888181518110614f7357614f73615cfc565b60200260200101516000801b03614fab5784828281518110614f9757614f97615cfc565b60209081029190910101526000945061501b565b6000801b828281518110614fc157614fc1615cfc565b602002602001018181525050888181518110614fdf57614fdf615cfc565b602002602001015185604051602001615002929190918252602082015260400190565b6040516020818303038152906040528051906020012094505b8061502581615ec8565b915050614e74565b508315615061578381600183516150449190615ce9565b8151811061505457615054615cfc565b6020026020010181815250505b80600182516150709190615ce9565b8151811061508057615080615cfc565b60200260200101516000801b036106605760405162461bcd60e51b815260206004820152600f60248201527f4c61737420656e747279207a65726f00000000000000000000000000000000006044820152606401610db1565b6000859003615114576040517f8d79dbbc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b828111615157576040517f2060faf40000000000000000000000000000000000000000000000000000000081526004810184905260248101829052604401610db1565b6000849003615192576040517f83c683e400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000829003613cf8576040517f5cb6e5bb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060824710156152455760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610db1565b6001600160a01b0385163b61529c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610db1565b600080866001600160a01b031685876040516152b8919061625d565b60006040518083038185875af1925050503d80600081146152f5576040519150601f19603f3d011682016040523d82523d6000602084013e6152fa565b606091505b5091509150610660828286615377565b600080821161535b5760405162461bcd60e51b815260206004820152601c60248201527f5a65726f20686173206e6f207369676e69666963616e742062697473000000006044820152606401610db1565b6000828061536a600182615ce9565b161890506120898161475a565b60608315615386575081612089565b8251156153965782518084602001fd5b8160405162461bcd60e51b8152600401610db19190616279565b6040805160c0810182526000808252602082018190529181018290526060810191909152608081016153e06153f2565b81526020016153ed6153f2565b905290565b6040518060400160405280615405615411565b81526020016000905290565b604051806040016040528061542461542d565b81526020016153ed5b60405180604001604052806002906020820280368337509192915050565b803560ff81168114610dba57600080fd5b60008060008060008060c0878903121561547557600080fd5b61547e8761544b565b9860208801359850604088013597606081013597506080810135965060a00135945092505050565b6000602082840312156154b857600080fd5b813567ffffffffffffffff8111156154cf57600080fd5b820160c0818503121561208957600080fd5b600080604083850312156154f457600080fd5b50508035926020909101359150565b60006020828403121561551557600080fd5b5035919050565b6003811061552957600080fd5b50565b8035610dba8161551c565b60006020828403121561554957600080fd5b81356120898161551c565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561558d5761558d615554565b60405290565b6040805190810167ffffffffffffffff8111828210171561558d5761558d615554565b604051601f8201601f1916810167ffffffffffffffff811182821017156155df576155df615554565b604052919050565b600067ffffffffffffffff82111561560157615601615554565b5060051b60200190565b600082601f83011261561c57600080fd5b8135602061563161562c836155e7565b6155b6565b82815260059290921b8401810191818101908684111561565057600080fd5b8286015b8481101561566b5780358352918301918301615654565b509695505050505050565b600080600083850361012081121561568d57600080fd5b84359350602085013567ffffffffffffffff8111156156ab57600080fd5b6156b78782880161560b565b93505060e0603f19820112156156cc57600080fd5b506040840190509250925092565b60008083601f8401126156ec57600080fd5b50813567ffffffffffffffff81111561570457600080fd5b6020830191508360208260051b850101111561571f57600080fd5b9250929050565b600080600080600080600087890361012081121561574357600080fd5b88359750602089013567ffffffffffffffff8082111561576257600080fd5b908a01906040828d03121561577657600080fd5b81985060a0603f198401121561578b57600080fd5b60408b01975060e08b01359250808311156157a557600080fd5b6157b18c848d016156da565b90975095506101008b01359250869150808311156157ce57600080fd5b50506157dc8a828b016156da565b989b979a50959850939692959293505050565b600080600080600060a0868803121561580757600080fd5b6158108661544b565b97602087013597506040870135966060810135965060800135945092505050565b6000806000806060858703121561584757600080fd5b8435935060208501359250604085013567ffffffffffffffff8082111561586d57600080fd5b818701915087601f83011261588157600080fd5b81358181111561589057600080fd5b8860208285010111156158a257600080fd5b95989497505060200194505050565b634e487b7160e01b600052602160045260246000fd5b600281106158d7576158d76158b1565b9052565b60006101c082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015260e083015160e083015261010080840151615949828501826001600160a01b03169052565b50506101208381015167ffffffffffffffff81168483015250506101408381015167ffffffffffffffff81168483015250506101608084015161598e828501826158c7565b50506101808381015160ff16908301526101a08084015180151582850152612053565b6000602082840312156159c357600080fd5b6120898261544b565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112615a0157600080fd5b83018035915067ffffffffffffffff821115615a1c57600080fd5b60200191503681900382131561571f57600080fd5b67ffffffffffffffff8116811461552957600080fd5b600082601f830112615a5857600080fd5b615a60615593565b806040840185811115615a7257600080fd5b845b8181101561348f578035615a8781615a31565b845260209384019301615a74565b600081830360e0811215615aa857600080fd5b615ab061556a565b915060a0811215615ac057600080fd5b615ac8615593565b6080821215615ad657600080fd5b615ade615593565b915084601f850112615aef57600080fd5b615af7615593565b806040860187811115615b0957600080fd5b865b81811015615b23578035845260209384019301615b0b565b50818552615b318882615a47565b6020860152505050818152615b486080850161552c565b6020820152808352505060a0820135602082015260c0820135604082015292915050565b60008060006101e08486031215615b8257600080fd5b833567ffffffffffffffff811115615b9957600080fd5b615ba58682870161560b565b935050615bb58560208601615a95565b9150615bc5856101008601615a95565b90509250925092565b60038110615529576155296158b1565b6158d781615bce565b805180518360005b6002811015615c0e578251825260209283019290910190600101615bef565b505050602090810151906040840160005b6002811015615c4657835167ffffffffffffffff1682529282019290820190600101615c1f565b505082015190506126266080840182615bde565b8481526101008101615c6f6020830186615be7565b60c082019390935260e0015292915050565b600060208284031215615c9357600080fd5b8151801515811461208957600080fd5b600060208284031215615cb557600080fd5b815161208981615a31565b60208101615ccd83615bce565b91905290565b634e487b7160e01b600052601160045260246000fd5b81810381811115610cb657610cb6615cd3565b634e487b7160e01b600052603260045260246000fd5b8481526101008101602060408682850137606083016040870160005b6002811015615d5e578135615d4281615a31565b67ffffffffffffffff1683529183019190830190600101615d2e565b505050506080850135615d708161551c565b615d7981615bce565b60a083015260c082019390935260e0015292915050565b67ffffffffffffffff8281168282160390808211156144ea576144ea615cd3565b6001600160a01b038116811461552957600080fd5b600060c08201905083825282356020830152602083013560408301526040830135615df081615db1565b6001600160a01b0381166060840152506060830135615e0e81615a31565b67ffffffffffffffff808216608085015260808501359150615e2f82615a31565b80821660a085015250509392505050565b600060208284031215615e5257600080fd5b813561208981615a31565b600060208284031215615e6f57600080fd5b815161208981615db1565b60ff8181168382160190811115610cb657610cb6615cd3565b8281526040810161208960208301846158c7565b67ffffffffffffffff8181168382160190808211156144ea576144ea615cd3565b60006000198203615edb57615edb615cd3565b5060010190565b855181526001600160a01b0360208701511660208201526040860151604082015284606082015283608082015260c060a08201528160c0820152818360e0830137600081830160e090810191909152601f909201601f19160101949350505050565b600060208284031215615f5657600080fd5b5051919050565b80820180821115610cb657610cb6615cd3565b600082601f830112615f8157600080fd5b81516020615f9161562c836155e7565b82815260059290921b84018101918181019086841115615fb057600080fd5b8286015b8481101561566b5780518352918301918301615fb4565b60008060408385031215615fde57600080fd5b825167ffffffffffffffff80821115615ff657600080fd5b61600286838701615f70565b9350602085015191508082111561601857600080fd5b5061602585828601615f70565b9150509250929050565b60a08101610cb68284615be7565b600080600080600060a0868803121561605557600080fd5b8535945060208601359350604086013567ffffffffffffffff8082111561607b57600080fd5b61608789838a0161560b565b9450606088013591508082111561609d57600080fd5b6160a989838a0161560b565b935060808801359150808211156160bf57600080fd5b506160cc8882890161560b565b9150509295509295909350565b600080604083850312156160ec57600080fd5b823567ffffffffffffffff8082111561610457600080fd5b6161108683870161560b565b9350602085013591508082111561612657600080fd5b506160258582860161560b565b634e487b7160e01b600052600160045260246000fd5b600181815b8085111561618457816000190482111561616a5761616a615cd3565b8085161561617757918102915b93841c939080029061614e565b509250929050565b60008261619b57506001610cb6565b816161a857506000610cb6565b81600181146161be57600281146161c8576161e4565b6001915050610cb6565b60ff8411156161d9576161d9615cd3565b50506001821b610cb6565b5060208310610133831016604e8410600b8410161715616207575081810a610cb6565b6162118383616149565b806000190482111561622557616225615cd3565b029392505050565b6000612089838361618c565b60005b8381101561625457818101518382015260200161623c565b50506000910152565b6000825161626f818460208701616239565b9190910192915050565b6020815260008251806020840152616298816040850160208701616239565b601f01601f1916919091016040019291505056fea264697066735822122097528bd78949481aeb4d5c9b8e621b8194cde0b12bbb8c2e656397120493558364736f6c63430008110033",
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

// EdgeChallengeManagerFactoryMetaData contains all meta data concerning the EdgeChallengeManagerFactory contract.
var EdgeChallengeManagerFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAssertionChain\",\"name\":\"_assertionChain\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_challengePeriodBlocks\",\"type\":\"uint64\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"_oneStepProofEntry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"layerZeroBlockEdgeHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"layerZeroBigStepEdgeHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"layerZeroSmallStepEdgeHeight\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_excessStakeReceiver\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_numBigStepLevel\",\"type\":\"uint8\"}],\"name\":\"createChallengeManager\",\"outputs\":[{\"internalType\":\"contractEdgeChallengeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615f54806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063c9fade8e14610030575b600080fd5b61004361003e36600461012b565b61005f565b6040516001600160a01b03909116815260200160405180910390f35b60008a8a8a8a8a8a8a8a8a8a60405161007790610106565b6001600160a01b039a8b16815267ffffffffffffffff90991660208a015296891660408901526060880195909552608087019390935260a0860191909152851660c085015260e08401529290921661010082015260ff90911661012082015261014001604051809103906000f0801580156100f6573d6000803e3d6000fd5b509b9a5050505050505050505050565b615d34806101eb83390190565b6001600160a01b038116811461012857600080fd5b50565b6000806000806000806000806000806101408b8d03121561014b57600080fd5b8a3561015681610113565b995060208b013567ffffffffffffffff8116811461017357600080fd5b985060408b013561018381610113565b975060608b0135965060808b0135955060a08b0135945060c08b01356101a881610113565b935060e08b013592506101008b01356101c081610113565b91506101208b013560ff811681146101d757600080fd5b809150509295989b9194979a509295985056fe6101c06040523480156200001257600080fd5b5060405162005d3438038062005d34833981016040819052620000359162000285565b6001600160a01b038a166200005d5760405163641f043160e11b815260040160405180910390fd5b6001600160a01b03808b166101005288166200008c5760405163fb60b0ef60e01b815260040160405180910390fd5b6001600160a01b038816610120526001600160401b038916600003620000c557604051632283bb7360e21b815260040160405180910390fd5b6001600160401b03891660e0526001600160a01b0380851660a05260c0849052821662000105576040516301e1d91560e31b815260040160405180910390fd5b6001600160a01b03821660805262000129876200023f602090811b620016cc17901c565b6200014f57604051633abfb6ff60e21b8152600481018890526024015b60405180910390fd5b8661014081815250506200016e866200023f60201b620016cc1760201c565b6200019057604051633abfb6ff60e21b81526004810187905260240162000146565b856101608181525050620001af856200023f60201b620016cc1760201c565b620001d157604051633abfb6ff60e21b81526004810186905260240162000146565b61018085905260ff8116600003620001fc57604051632a18f5b960e21b815260040160405180910390fd5b60fd8160ff161115620002285760405163040d23bf60e41b815260ff8216600482015260240162000146565b60ff166101a052506200037c975050505050505050565b6000816000036200025257506000919050565b60006200026160018462000354565b929092161592915050565b6001600160a01b03811681146200028257600080fd5b50565b6000806000806000806000806000806101408b8d031215620002a657600080fd5b8a51620002b3816200026c565b60208c0151909a506001600160401b0381168114620002d157600080fd5b60408c0151909950620002e4816200026c565b8098505060608b0151965060808b0151955060a08b0151945060c08b01516200030d816200026c565b60e08c01516101008d0151919550935062000328816200026c565b6101208c015190925060ff811681146200034157600080fd5b809150509295989b9194979a5092959850565b818103818111156200037657634e487b7160e01b600052601160045260246000fd5b92915050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a051615871620004c3600039600081816103af0152818161059e015281816108f90152818161095101528181610abf01528181610cfd0152818161103e01526113610152600081816105230152610c3701526000818161028b0152610bf70152600081816102260152610bb70152600081816102ec015281816108d70152818161092f015261133a01526000818161032b01528181610667015281816106fa015281816107840152818161082001528181610dbd01528181610e7201528181610efb01528181610f8901528181611206015261129d0152600081816102c5015261101d0152600081816103e80152818161099b01526110ff0152600081816103520152818161097a01526110de0152600081816104e901526109e901526158716000f3fe608060405234801561001057600080fd5b50600436106101ce5760003560e01c80635d9e244411610104578063bce6f54f116100a2578063e94e051e11610071578063e94e051e146104e4578063eae0328b1461050b578063f8ee77d61461051e578063fda2892e1461054557600080fd5b8063bce6f54f14610469578063c32d8c6314610489578063c8bc4e431461049c578063e5b123da146104c457600080fd5b8063748926f3116100de578063748926f31461041d578063750e0c0f146104305780638c1b3a4014610443578063908517e91461045657600080fd5b80635d9e2444146103aa57806360c7dc47146103e357806364deed591461040a57600080fd5b806342e1aaa81161017157806348dd29241161014b57806348dd29241461032657806351ed6a301461034d57806354b64151146103745780635a48e0f41461039757600080fd5b806342e1aaa8146102ad57806346c2781a146102c057806348923bc5146102e757600080fd5b80631dce5166116101ad5780631dce5166146102215780632eaa0043146102485780633e35f5e81461025b578063416e66571461028657600080fd5b80624d8efe146101d357806305fae141146101f95780630f73bfad1461020c575b600080fd5b6101e66101e1366004614a20565b610565565b6040519081526020015b60405180910390f35b6101e6610207366004614a6a565b610580565b61021f61021a366004614aa4565b610ab6565b005b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b61021f610256366004614ac6565b610b38565b61026e610269366004614ac6565b610b88565b6040516001600160401b0390911681526020016101f0565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b6101e66102bb366004614afa565b610b9a565b61026e7f000000000000000000000000000000000000000000000000000000000000000081565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101f0565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b610387610382366004614ac6565b610c84565b60405190151581526020016101f0565b6101e66103a5366004614ac6565b610c90565b6103d17f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff90911681526020016101f0565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b61021f610418366004614c35565b610c9c565b61021f61042b366004614ac6565b6110c5565b61038761043e366004614ac6565b6111b7565b61021f610451366004614ce3565b6111e0565b610387610464366004614ac6565b6113d2565b6101e6610477366004614ac6565b60009081526001602052604090205490565b6101e6610497366004614dab565b6113de565b6104af6104aa366004614ded565b6113f7565b604080519283526020830191909152016101f0565b6101e66104d2366004614ac6565b60009081526002602052604090205490565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b6101e6610519366004614ac6565b6115ae565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b610558610553366004614ac6565b6115c2565b6040516101f09190614e96565b60006105758787878787876116f6565b979650505050505050565b600061058a6148bd565b60006105c261059c6020860186614f6a565b7f000000000000000000000000000000000000000000000000000000000000000061173b565b905060006105cf82610b9a565b90506105d9614901565b60008360028111156105ed576105ed614e6c565b03610926576105ff60a0870187614f85565b905060000361062157604051630c9ccac560e41b815260040160405180910390fd5b60008061063160a0890189614f85565b81019061063e9190615105565b80516020820151604080840151905163f9cee9df60e01b81529497509295506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016945063f9cee9df936106a39360608f01359392916004016151f1565b60006040518083038186803b1580156106bb57600080fd5b505afa1580156106cf573d6000803e3d6000fd5b505050602080830151845191850151604080870151905163f9cee9df60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016955063f9cee9df946107339493909290916004016151f1565b60006040518083038186803b15801561074b57600080fd5b505afa15801561075f573d6000803e3d6000fd5b505050506040518060c0016040528089606001358152602001826020015181526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e531d8c78b606001356040518263ffffffff1660e01b81526004016107d491815260200190565b602060405180830381865afa1580156107f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108159190615218565b1515815260200160007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166356bbc9e685602001516040518263ffffffff1660e01b815260040161087091815260200190565b602060405180830381865afa15801561088d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b1919061523a565b6001600160401b0316118152835160208201528251604090910152925061091d600089857f0000000000000000000000000000000000000000000000000000000000000000887f00000000000000000000000000000000000000000000000000000000000000006117aa565b95505050610978565b610975600087837f0000000000000000000000000000000000000000000000000000000000000000867f00000000000000000000000000000000000000000000000000000000000000006117aa565b93505b7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b038216158015906109d157508015155b15610a225760008660c001516109e75730610a09565b7f00000000000000000000000000000000000000000000000000000000000000005b9050610a206001600160a01b0384163383856117fd565b505b8560400151866020015187600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a489606001518a608001518b60a001518c60c001518d60e00151604051610aa1959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a45050925195945050505050565b610ae3600083837f0000000000000000000000000000000000000000000000000000000000000000611883565b6000828152602081905260409020610afa906119ef565b827fb924f3aa473645c7cf5b10262f927ae4ccf869d7fc239c17144b0c67490d1c7383604051610b2c91815260200190565b60405180910390a35050565b610b43600082611a1f565b6000818152602081905260409020610b5a906119ef565b60405182907f0d27fcaf1adc41547a5cfc99d2364f6c0dc7e81c9fc3fe8cb38abb409b48358a90600090a350565b6000610b948183611bd2565b92915050565b600080826002811115610baf57610baf614e6c565b03610bdb57507f0000000000000000000000000000000000000000000000000000000000000000919050565b6001826002811115610bef57610bef614e6c565b03610c1b57507f0000000000000000000000000000000000000000000000000000000000000000919050565b6002826002811115610c2f57610c2f614e6c565b03610c5b57507f0000000000000000000000000000000000000000000000000000000000000000919050565b81604051630efcb87b60e21b8152600401610c769190615257565b60405180910390fd5b919050565b6000610b948183611d72565b6000610b948183611da7565b600080835111610cac5783610cd4565b8260018451610cbb9190615280565b81518110610ccb57610ccb615293565b60200260200101515b90506000610ce28183611df8565b90506000610d218260090160099054906101000a900460ff167f000000000000000000000000000000000000000000000000000000000000000061173b565b90506000816002811115610d3757610d37614e6c565b14610d6657600982015460405163ec72dc5d60e01b8152600160481b90910460ff166004820152602401610c76565b610d6f82611e4c565b610db857610d7c82611e70565b60088301546007840154604051631cb1906160e31b815260048101939093526001600160a01b0390911660248301526044820152606401610c76565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633083622885600701546040518263ffffffff1660e01b8152600401610e0d91815260200190565b602060405180830381865afa158015610e2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e4e9190615218565b9050801561100d57600784015460405163f9cee9df60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163f9cee9df91610eb491908a9060a08201359060c0830135906004016152a9565b60006040518083038186803b158015610ecc57600080fd5b505afa158015610ee0573d6000803e3d6000fd5b5050604051631171558560e01b815260a089013560048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316925063117155859150602401602060405180830381865afa158015610f4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f70919061523a565b604051632b5de4f360e11b815260a088013560048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906356bbc9e690602401602060405180830381865afa158015610fd8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ffc919061523a565b6110069190615326565b9150611012565b600091505b6000611062818a8a867f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611ea5565b60008a815260208190526040902090915061107c906119ef565b6040516001600160401b03831681528a907f9cd2c77f6772dd0fb07e9972aee8ddddf9fd13f1a65abb03a3de88b07dc59af99060200160405180910390a3505050505050505050565b60006110d18183611df8565b90506110dc816120f3565b7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0382161580159061113557508015155b15611156576008830154611156906001600160a01b038481169116836121f2565b600084815260208190526040902061116d906119ef565b604080516001600160a01b03851681526020810184905286917fa635398959ddb5ce3b14537edfc25b2e671274c9b8cad0f4bd634752e69007b6910160405180910390a350505050565b600081815260208190526040812060080154600160a01b90046001600160401b03161515610b94565b60006111ec8189611da7565b6040516304972af960e01b81529091506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906304972af99061123d9084908a9060040161535b565b60006040518083038186803b15801561125557600080fd5b505afa158015611269573d6000803e3d6000fd5b505050506000604051806060016040528088608001602081019061128d91906153d4565b6001600160401b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e78cea926040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061131d91906153f1565b6001600160a01b031681528835602090910152905061138560008a7f00000000000000000000000000000000000000000000000000000000000000008b858b8b8b8b7f0000000000000000000000000000000000000000000000000000000000000000612227565b600089815260208190526040902061139c906119ef565b6040518a907fe11db4b27bc8c6ea5943ecbb205ae1ca8d56c42c719717aaf8a53d43d0cee7c290600090a3505050505050505050565b6000610b94818361246b565b60006113ed8686868686612531565b9695505050505050565b6000806000806000611443898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509594939250506125a39050565b8151929550909350915015806114db578260400151836020015184600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a4866060015187608001518860a001518960c001518a60e001516040516114d2959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a45b8160400151826020015183600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a4856060015186608001518760a001518860c001518960e0015160405161155a959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a48151604051821515815285908c907f7340510d24b7ec9b5c100f5500d93429d80d00d46f0d18e4e85d0c4cc22b99249060200160405180910390a4505191989197509095505050505050565b6000610b946115bd8284611df8565b612892565b6115ca614943565b6115d5600083611df8565b604080516101c0810182528254815260018084015460208301526002840154928201929092526003830154606082015260048301546080820152600583015460a0820152600683015460c0820152600783015460e082015260088301546001600160a01b038116610100830152600160a01b90046001600160401b039081166101208301526009840154908116610140830152909291610160840191600160401b900460ff169081111561168b5761168b614e6c565b600181111561169c5761169c614e6c565b81526009919091015460ff600160481b820481166020840152600160501b90910416151560409091015292915050565b6000816000036116de57506000919050565b60006116eb600184615280565b929092161592915050565b60006117058787878787612531565b60408051602081019290925281018390526060016040516020818303038152906040528051906020012090509695505050505050565b60008260ff1660000361175057506000610b94565b8160ff168360ff161161176557506001610b94565b61177082600161540e565b60ff168360ff160361178457506002610b94565b6040516315c1b4af60e31b815260ff808516600483015283166024820152604401610c76565b6117b26148bd565b6000806117c289898989886128d7565b9150915060006117d3838a88612d64565b905060006117e283838c612e77565b90506117ee8b82612eb2565b9b9a5050505050505050505050565b6040516001600160a01b038085166024830152831660448201526064810182905261187d9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166001600160e01b031990931692909217909152613189565b50505050565b600083815260208590526040902060080154600160a01b90046001600160401b03166118c45760405162a7b02b60e01b815260048101849052602401610c76565b600082815260208590526040902060080154600160a01b90046001600160401b03166119055760405162a7b02b60e01b815260048101849052602401610c76565b6001600083815260208690526040902060090154600160401b900460ff16600181111561193457611934614e6c565b146119735760008281526020859052604090819020600901549051633bc499ed60e21b8152610c76918491600160401b90910460ff1690600401615427565b61197f8484848461325b565b60008281526020859052604090206007015483146119ce5760008281526020859052604090819020600701549051631855b87d60e31b8152610c76918591600401918252602082015260400190565b6119d88484613383565b600083815260208590526040902061187d906133ed565b6000610b948260090160099054906101000a900460ff168360000154846002015485600101548660040154612531565b600081815260208390526040902060080154600160a01b90046001600160401b0316611a605760405162a7b02b60e01b815260048101829052602401610c76565b60008181526020839052604080822060050154808352912060080154600160a01b90046001600160401b0316611aab5760405162a7b02b60e01b815260048101829052602401610c76565b6001600082815260208590526040902060090154600160401b900460ff166001811115611ada57611ada614e6c565b14611b195760008181526020849052604090819020600901549051633bc499ed60e21b8152610c76918391600160401b90910460ff1690600401615427565b60008281526020849052604080822060060154808352912060080154600160a01b90046001600160401b0316611b645760405162a7b02b60e01b815260048101829052602401610c76565b6001600082815260208690526040902060090154600160401b900460ff166001811115611b9357611b93614e6c565b146119ce5760008181526020859052604090819020600901549051633bc499ed60e21b8152610c76918391600160401b90910460ff1690600401615427565b600081815260208390526040812060080154600160a01b90046001600160401b0316611c135760405162a7b02b60e01b815260048101839052602401610c76565b6000828152602084905260408120611c2a906119ef565b6000818152600186016020526040812054919250819003611c5e576040516336843d9f60e21b815260040160405180910390fd5b6040516815539492559053115160ba1b6020820152602901604051602081830303815290604052805190602001208103611cc757600084815260208690526040902060080154611cbe90600160a01b90046001600160401b031643615326565b92505050610b94565b600081815260208690526040902060080154600160a01b90046001600160401b0316611d085760405162a7b02b60e01b815260048101829052602401610c76565b60008181526020869052604080822060089081015487845291909220909101546001600160401b03600160a01b928390048116929091041680821115611d5d57611d528183615326565b945050505050610b94565b6000945050505050610b94565b505092915050565b6000611d7e838361246b565b8015611da057506000828152602084905260409020611d9c90612892565b6001145b9392505050565b600080611db48484611df8565b90505b6009810154600160481b900460ff1615611df05780546000908152600185016020526040902054611de88582611df8565b915050611db7565b549392505050565b600081815260208390526040812060080154600160a01b90046001600160401b0316611e395760405162a7b02b60e01b815260048101839052602401610c76565b5060009081526020919091526040902090565b600781015460009015801590610b94575050600801546001600160a01b0316151590565b6000610b948260090160099054906101000a900460ff16836000015484600201548560010154866004015487600301546116f6565b600085815260208790526040812060080154600160a01b90046001600160401b0316611ee65760405162a7b02b60e01b815260048101879052602401610c76565b856000611ef38983611bd2565b905060005b8751811015612073576000611f268b8a8481518110611f1957611f19615293565b6020026020010151611df8565b90508381600501541480611f3d5750838160060154145b15611f8157611f548b611f4f83611e70565b611bd2565b611f5e908461543b565b9250888281518110611f7257611f72615293565b60200260200101519350612060565b600084815260208c9052604090206007015489518a9084908110611fa757611fa7615293565b602002602001015103611fe657611fd98b8a8481518110611fca57611fca615293565b6020026020010151868961325b565b611f548b611f4f83611e70565b83816005015482600601548b858151811061200357612003615293565b60200260200101518e600001600089815260200190815260200160002060070154604051636ebd28c960e01b8152600401610c76959493929190948552602085019390935260408401919091526060830152608082015260a00190565b508061206b8161545b565b915050611ef8565b5061207e868261543b565b9050846001600160401b0316816001600160401b031610156120c65760405163011a8d4d60e41b81526001600160401b03808316600483015286166024820152604401610c76565b6120d08989613383565b600088815260208a9052604090206120e7906133ed565b98975050505050505050565b60016009820154600160401b900460ff16600181111561211557612115614e6c565b1461214d5761212381611e70565b6009820154604051633bc499ed60e21b8152610c769291600160401b900460ff1690600401615427565b61215681611e4c565b61219f5761216381611e70565b60088201546007830154604051631cb1906160e31b815260048101939093526001600160a01b0390911660248301526044820152606401610c76565b6009810154600160501b900460ff1615156001036121dc576121c081611e70565b60405163307f766960e01b8152600401610c7691815260200190565b600901805460ff60501b1916600160501b179055565b6040516001600160a01b03831660248201526044810182905261222290849063a9059cbb60e01b90606401611831565b505050565b60006122338b8b611df8565b600290810154915060008b815260208d9052604090206009015461226190600160481b900460ff168461173b565b600281111561227257612272614e6c565b146122af5760008a815260208c905260409081902060090154905163348aefdf60e01b8152600160481b90910460ff166004820152602401610c76565b60008a815260208c9052604090206122c690612892565b6001146123005760008a815260208c9052604090206122e490612892565b6040516306b595e560e41b8152600401610c7691815260200190565b61235b8b60000160008c81526020019081526020016000206001015489600001358389898080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061346f92505050565b60006001600160a01b038a1663b5112fd289848c3561237d60208f018f614f85565b6040518663ffffffff1660e01b815260040161239d959493929190615474565b602060405180830381865afa1580156123ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123de91906154d6565b60008c815260208e9052604090206003015490915061243c90826124038560016154ef565b88888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061346f92505050565b6124468c8c613383565b60008b815260208d90526040902061245d906133ed565b505050505050505050505050565b600081815260208390526040812060080154600160a01b90046001600160401b03166124ac5760405162a7b02b60e01b815260048101839052602401610c76565b60008281526020849052604081206124c3906119ef565b60008181526001860160205260408120549192508190036124f7576040516336843d9f60e21b815260040160405180910390fd5b6040516815539492559053115160ba1b602082015260290160408051601f1981840301815291905280516020909101201415949350505050565b6040517fff0000000000000000000000000000000000000000000000000000000000000060f887901b1660208201526021810185905260418101849052606181018390526081810182905260009060a10160405160208183030381529060405280519060200120905095945050505050565b60006125ad6148bd565b6125b56148bd565b60008087815260208990526040902060090154600160401b900460ff1660018111156125e3576125e3614e6c565b1461262257600086815260208890526040908190206009015490516323f8405d60e01b8152610c76918891600160401b90910460ff1690600401615427565b61262c878761246b565b61264c576040516380e07e4560e01b815260048101879052602401610c76565b6000868152602088905260408120604080516101c0810182528254815260018084015460208301526002840154928201929092526003830154606082015260048301546080820152600583015460a0820152600683015460c0820152600783015460e082015260088301546001600160a01b038116610100830152600160a01b90046001600160401b039081166101208301526009840154908116610140830152909291610160840191600160401b900460ff169081111561271057612710614e6c565b600181111561272157612721614e6c565b81526009919091015460ff600160481b820481166020840152600160501b909104161515604091820152810151608082015191925060009161276391906134fc565b90506000808780602001905181019061277c919061555d565b90925090506127ac896127908560016154ef565b606087015160808801516127a59060016154ef565b8686613590565b505060006127b86148bd565b60006127d98560000151866020015187604001518d888a610180015161387c565b90506127e48161390c565b600081815260208e90526040902060080154909350600160a01b90046001600160401b031661281a576128178c82612eb2565b91505b506128236148bd565b600061284486600001518c8789606001518a608001518b610180015161387c565b90506128508d82612eb2565b9150506128808382600001518e60000160008f81526020019081526020016000206139359092919063ffffffff16565b919b909a509098509650505050505050565b600080826002015483600401546128a99190615280565b905080600003610b94576128bc83611e70565b60405162a7b02b60e01b8152600401610c7691815260200190565b6040805160608082018352600080835260208301529181019190915260008061290c6129066020890189614f6a565b8561173b565b600281111561291d5761291d614e6c565b03612b965760208501518551600003612949576040516374b5e30d60e11b815260040160405180910390fd5b855160608801351461297e5785516040516316c5de8f60e21b8152600481019190915260608801356024820152604401610c76565b85604001516129a0576040516360b4921b60e11b815260040160405180910390fd5b85606001516129c257604051635a2e8e1d60e11b815260040160405180910390fd5b6129cf60a0880188614f85565b90506000036129f157604051630c9ccac560e41b815260040160405180910390fd5b6000612a0060a0890189614f85565b810190612a0d9190615105565b50909150600090508760800151602001516002811115612a2f57612a2f614e6c565b03612a4d5760405163231b2f2960e11b815260040160405180910390fd5b60008760a00151602001516002811115612a6957612a69614e6c565b03612a8757604051638999857d60e01b815260040160405180910390fd5b60808701516040516330e5867160e21b81526000916001600160a01b0389169163c39619c491612ab9916004016155c0565b602060405180830381865afa158015612ad6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612afa91906154d6565b90506000876001600160a01b031663c39619c48a60a001516040518263ffffffff1660e01b8152600401612b2e91906155c0565b602060405180830381865afa158015612b4b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b6f91906154d6565b6040805160608101825293845260208401919091528201929092529350909150612d5a9050565b612ba4878760600135611d72565b612be0576040517fff6d9bd700000000000000000000000000000000000000000000000000000000815260608701356004820152602401610c76565b6060860135600090815260208890526040812090612bfd826119ef565b905060006009830154600160401b900460ff166001811115612c2157612c21614e6c565b14612c3f576040516312459ffd60e01b815260040160405180910390fd5b6009820154612c5890600160481b900460ff168661399c565b60ff16612c6860208a018a614f6a565b60ff1614612cb157612c7d6020890189614f6a565b600983015460405163564f308b60e11b815260ff9283166004820152600160481b9091049091166024820152604401610c76565b612cbe60a0890189614f85565b9050600003612ce057604051630c9ccac560e41b815260040160405180910390fd5b600080808080612cf360a08e018e614f85565b810190612d0091906155ce565b94509450945094509450612d1e87600101548689600201548661346f565b612d3287600301548589600401548561346f565b6040518060600160405280868152602001858152602001828152508698509850505050505050505b9550959350505050565b604080516000808252602082019092528190612d8a90612d859087516139be565b6139f4565b9050612d95836116cc565b612db557604051633abfb6ff60e21b815260048101849052602401610c76565b82846040013514612de657604080516337f318af60e21b815290850135600482015260248101849052604401610c76565b612e02846020013586602001518660400135886040015161346f565b612e0f6080850185614f85565b9050600003612e3157604051631a1503a960e11b815260040160405180910390fd5b600080612e416080870187614f85565b810190612e4e9190615669565b9092509050612e6c83600160208901356127a560408b0135836154ef565b509095945050505050565b612e7f614943565b612eaa84846000602086018035906040880135906060890135903390612ea5908b614f6a565b613b94565b949350505050565b612eba6148bd565b6000612ec58361390c565b600081815260208690526040902060080154909150600160a01b90046001600160401b031615612f0b57604051635e76f9ef60e11b815260048101829052602401610c76565b600081815260208581526040918290208551815590850151600180830191909155918501516002820155606085015160038201556080850151600482015560a0850151600582015560c0850151600682015560e085015160078201556101008501516008820180546101208801516001600160401b03908116600160a01b026001600160e01b03199092166001600160a01b03909416939093171790556101408601516009830180549190921667ffffffffffffffff19821681178355610160880151889590939268ffffffffffffffffff191690911790600160401b908490811115612ffa57612ffa614e6c565b021790555061018082810151600990920180546101a0909401511515600160501b0260ff60501b1960ff909416600160481b02939093166affff000000000000000000199094169390931791909117909155830151835160408501516020860151608087015160009461307294909390929091612531565b60008181526001870160205260408120549192508190036130d1576040516815539492559053115160ba1b602082015260290160408051601f198184030181529181528151602092830120600085815260018a01909352912055613119565b6040516815539492559053115160ba1b602082015260290160405160208183030381529060405280519060200120810361311957600082815260018701602052604090208390555b6040805161010081018252848152602080820185905287518284015260e0880151606083015260008681529089905291909120608082019061315a90612892565b815261018087015160ff166020820152911515604083015260e090950151151560609091015250919392505050565b60006131de826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613c969092919063ffffffff16565b80519091501561222257808060200190518101906131fc9190615218565b6122225760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610c76565b600082815260208590526040808220548583529120613279906119ef565b146132c7576000838152602085905260409020613295906119ef565b6000838152602086905260409081902054905163e2e27f8760e01b815260048101929092526024820152604401610c76565b600082815260208590526040808220600990810154868452919092209091015460ff600160481b928390048116926133019204168361399c565b60ff161461187d576000838152602085905260409020600901548390839061333390600160481b900460ff168461399c565b60008581526020889052604090819020600901549051637e726d1560e01b81526004810194909452602484019290925260ff9081166044840152600160481b909104166064820152608401610c76565b600081815260208390526040812061339a906119ef565b600081815260028501602052604090205490915080156133d757604051630dd7028f60e41b81526004810184905260248101829052604401610c76565b5060009081526002909201602052604090912055565b60006009820154600160401b900460ff16600181111561340f5761340f614e6c565b146134475761341d81611e70565b60098201546040516323f8405d60e01b8152610c769291600160401b900460ff1690600401615427565b60090180546001600160401b03431668ffffffffffffffffff1990911617600160401b179055565b60006134a482848660405160200161348991815260200190565b60405160208183030381529060405280519060200120613ca5565b90508085146134f55760405162461bcd60e51b815260206004820152601760248201527f496e76616c696420696e636c7573696f6e2070726f6f660000000000000000006044820152606401610c76565b5050505050565b6000600261350a8484615280565b10156135335760405163240a616560e21b81526004810184905260248101839052604401610c76565b61353d8383615280565b6002036135565761354f8360016154ef565b9050610b94565b600083613564600185615280565b189050600061357282613d47565b9050600019811b80613585600187615280565b169695505050505050565b600085116135e05760405162461bcd60e51b815260206004820152601460248201527f5072652d73697a652063616e6e6f7420626520300000000000000000000000006044820152606401610c76565b856135ea836139f4565b146136375760405162461bcd60e51b815260206004820152601b60248201527f50726520657870616e73696f6e20726f6f74206d69736d6174636800000000006044820152606401610c76565b8461364183613e76565b146136985760405162461bcd60e51b815260206004820152602160248201527f5072652073697a6520646f6573206e6f74206d6174636820657870616e73696f6044820152603760f91b6064820152608401610c76565b8285106136e75760405162461bcd60e51b815260206004820181905260248201527f5072652073697a65206e6f74206c657373207468616e20706f73742073697a656044820152606401610c76565b60008590506000806136fc8560008751613ed1565b90505b858310156137bf5760006137138488614033565b9050845183106137655760405162461bcd60e51b815260206004820152601260248201527f496e646578206f7574206f662072616e676500000000000000000000000000006044820152606401610c76565b613789828287868151811061377c5761377c615293565b602002602001015161411d565b91506001811b61379981866154ef565b9450878511156137ab576137ab6156c2565b836137b58161545b565b94505050506136ff565b866137c9826139f4565b146138215760405162461bcd60e51b815260206004820152602260248201527f506f737420657870616e73696f6e20726f6f74206e6f7420657175616c20706f6044820152611cdd60f21b6064820152608401610c76565b835182146138715760405162461bcd60e51b815260206004820152601660248201527f496e636f6d706c6574652070726f6f66207573616765000000000000000000006044820152606401610c76565b505050505050505050565b613884614943565b6138918787878787614661565b50604080516101c08101825296875260208701959095529385019290925260608401526080830152600060a0830181905260c0830181905260e0830181905261010083018190526001600160401b0343166101208401526101408301819052610160830181905260ff9091166101808301526101a082015290565b6000610b94826101800151836000015184604001518560200151866080015187606001516116f6565b600583015415158061394a5750600683015415155b1561398c5761395883611e70565b600584015460068501546040516308b0e71d60e41b8152600481019390935260248301919091526044820152606401610c76565b6005830191909155600690910155565b6000806139aa84600161540e565b90506139b6818461173b565b509392505050565b6060611da0836000846040516020016139d991815260200190565b6040516020818303038152906040528051906020012061411d565b600080825111613a465760405162461bcd60e51b815260206004820152601660248201527f456d707479206d65726b6c6520657870616e73696f6e000000000000000000006044820152606401610c76565b604082511115613a985760405162461bcd60e51b815260206004820152601a60248201527f4d65726b6c6520657870616e73696f6e20746f6f206c617267650000000000006044820152606401610c76565b6000805b8351811015613b8d576000848281518110613ab957613ab9615293565b60200260200101519050826000801b03613b25578015613b205780925060018551613ae49190615280565b8214613b2057604051613b07908490600090602001918252602082015260400190565b6040516020818303038152906040528051906020012092505b613b7a565b8015613b44576040805160208101839052908101849052606001613b07565b604051613b61908490600090602001918252602082015260400190565b6040516020818303038152906040528051906020012092505b5080613b858161545b565b915050613a9c565b5092915050565b613b9c614943565b6001600160a01b038316613bc35760405163f289e65760e01b815260040160405180910390fd5b6000849003613be557604051636932bcfd60e01b815260040160405180910390fd5b613bf28989898989614661565b604051806101c001604052808a81526020018981526020018881526020018781526020018681526020016000801b81526020016000801b8152602001858152602001846001600160a01b03168152602001436001600160401b0316815260200160006001600160401b0316815260200160006001811115613c7557613c75614e6c565b815260ff841660208201526000604090910152905098975050505050505050565b6060612eaa84846000856146f1565b8251600090610100811115613cd857604051637ed6198f60e11b8152600481018290526101006024820152604401610c76565b8260005b82811015613d3d576000878281518110613cf857613cf8615293565b60200260200101519050816001901b8716600003613d2457826000528060205260406000209250613d34565b8060005282602052604060002092505b50600101613cdc565b5095945050505050565b600081600003613d995760405162461bcd60e51b815260206004820152601c60248201527f5a65726f20686173206e6f207369676e69666963616e742062697473000000006044820152606401610c76565b600160801b8210613db757608091821c91613db490826154ef565b90505b600160401b8210613dd557604091821c91613dd290826154ef565b90505b6401000000008210613df457602091821c91613df190826154ef565b90505b620100008210613e1157601091821c91613e0e90826154ef565b90505b6101008210613e2d57600891821c91613e2a90826154ef565b90505b60108210613e4857600491821c91613e4590826154ef565b90505b60048210613e6357600291821c91613e6090826154ef565b90505b60028210610c7f57610b946001826154ef565b600080805b8351811015613b8d57838181518110613e9657613e96615293565b60200260200101516000801b14613ebf57613eb28160026157bc565b613ebc90836154ef565b91505b80613ec98161545b565b915050613e7b565b6060818310613f225760405162461bcd60e51b815260206004820152601760248201527f5374617274206e6f74206c657373207468616e20656e640000000000000000006044820152606401610c76565b8351821115613f7d5760405162461bcd60e51b815260206004820152602160248201527f456e64206e6f74206c657373206f7220657175616c207468616e206c656e67746044820152600d60fb1b6064820152608401610c76565b6000613f898484615280565b6001600160401b03811115613fa057613fa0614b17565b604051908082528060200260200182016040528015613fc9578160200160208202803683370190505b509050835b8381101561402a57858181518110613fe857613fe8615293565b6020026020010151828683613ffd9190615280565b8151811061400d5761400d615293565b6020908102919091010152806140228161545b565b915050613fce565b50949350505050565b60008183106140845760405162461bcd60e51b815260206004820152601760248201527f5374617274206e6f74206c657373207468616e20656e640000000000000000006044820152606401610c76565b6000614091838518613d47565b9050600060016140a183826154ef565b6001901b6140af9190615280565b905084811684821681156140c657611d5282614817565b80156140d557611d5281613d47565b60405162461bcd60e51b815260206004820152601b60248201527f426f7468207920616e64207a2063616e6e6f74206265207a65726f00000000006044820152606401610c76565b6060604083106141605760405162461bcd60e51b815260206004820152600e60248201526d098caeccad840e8dede40d0d2ced60931b6044820152606401610c76565b60008290036141b15760405162461bcd60e51b815260206004820152601b60248201527f43616e6e6f7420617070656e6420656d707479207375627472656500000000006044820152606401610c76565b6040845111156142035760405162461bcd60e51b815260206004820152601a60248201527f4d65726b6c6520657870616e73696f6e20746f6f206c617267650000000000006044820152606401610c76565b83516000036142815760006142198460016154ef565b6001600160401b0381111561423057614230614b17565b604051908082528060200260200182016040528015614259578160200160208202803683370190505b5090508281858151811061426f5761426f615293565b60209081029190910101529050611da0565b835183106142f75760405162461bcd60e51b815260206004820152603560248201527f4c6576656c2067726561746572207468616e2068696768657374206c6576656c60448201527f206f662063757272656e7420657870616e73696f6e00000000000000000000006064820152608401610c76565b81600061430386613e76565b905060006143128660026157bc565b61431c90836154ef565b9050600061432983613d47565b61433283613d47565b1161437f5787516001600160401b0381111561435057614350614b17565b604051908082528060200260200182016040528015614379578160200160208202803683370190505b506143ce565b875161438c9060016154ef565b6001600160401b038111156143a3576143a3614b17565b6040519080825280602002602001820160405280156143cc578160200160208202803683370190505b505b90506040815111156144225760405162461bcd60e51b815260206004820152601c60248201527f417070656e642063726561746573206f76657273697a652074726565000000006044820152606401610c76565b60005b88518110156145c357878110156144b15788818151811061444857614448615293565b60200260200101516000801b146144ac5760405162461bcd60e51b815260206004820152602260248201527f417070656e642061626f7665206c65617374207369676e69666963616e7420626044820152611a5d60f21b6064820152608401610c76565b6145b1565b60008590036144f7578881815181106144cc576144cc615293565b60200260200101518282815181106144e6576144e6615293565b6020026020010181815250506145b1565b88818151811061450957614509615293565b60200260200101516000801b03614541578482828151811061452d5761452d615293565b6020908102919091010152600094506145b1565b6000801b82828151811061455757614557615293565b60200260200101818152505088818151811061457557614575615293565b602002602001015185604051602001614598929190918252602082015260400190565b6040516020818303038152906040528051906020012094505b806145bb8161545b565b915050614425565b5083156145f7578381600183516145da9190615280565b815181106145ea576145ea615293565b6020026020010181815250505b80600182516146069190615280565b8151811061461657614616615293565b60200260200101516000801b036105755760405162461bcd60e51b815260206004820152600f60248201526e4c61737420656e747279207a65726f60881b6044820152606401610c76565b60008590036146835760405163235e76ef60e21b815260040160405180910390fd5b8281116146ad576040516308183ebd60e21b81526004810184905260248101829052604401610c76565b60008490036146cf576040516320f1a0f960e21b815260040160405180910390fd5b60008290036134f557604051635cb6e5bb60e01b815260040160405180910390fd5b6060824710156147525760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610c76565b6001600160a01b0385163b6147a95760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610c76565b600080866001600160a01b031685876040516147c591906157ec565b60006040518083038185875af1925050503d8060008114614802576040519150601f19603f3d011682016040523d82523d6000602084013e614807565b606091505b5091509150610575828286614884565b60008082116148685760405162461bcd60e51b815260206004820152601c60248201527f5a65726f20686173206e6f207369676e69666963616e742062697473000000006044820152606401610c76565b60008280614877600182615280565b16189050611da081613d47565b60608315614893575081611da0565b8251156148a35782518084602001fd5b8160405162461bcd60e51b8152600401610c769190615808565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915290565b6040805160c0810182526000808252602082018190529181018290526060810191909152608081016149316149b6565b815260200161493e6149b6565b905290565b604080516101c081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905290610160820190815260006020820181905260409091015290565b60405180604001604052806149c96149d5565b81526020016000905290565b60405180604001604052806149e86149f1565b815260200161493e5b60405180604001604052806002906020820280368337509192915050565b803560ff81168114610c7f57600080fd5b60008060008060008060c08789031215614a3957600080fd5b614a4287614a0f565b9860208801359850604088013597606081013597506080810135965060a00135945092505050565b600060208284031215614a7c57600080fd5b81356001600160401b03811115614a9257600080fd5b820160c08185031215611da057600080fd5b60008060408385031215614ab757600080fd5b50508035926020909101359150565b600060208284031215614ad857600080fd5b5035919050565b60038110614aec57600080fd5b50565b8035610c7f81614adf565b600060208284031215614b0c57600080fd5b8135611da081614adf565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715614b4f57614b4f614b17565b60405290565b604080519081016001600160401b0381118282101715614b4f57614b4f614b17565b604051601f8201601f191681016001600160401b0381118282101715614b9f57614b9f614b17565b604052919050565b60006001600160401b03821115614bc057614bc0614b17565b5060051b60200190565b600082601f830112614bdb57600080fd5b81356020614bf0614beb83614ba7565b614b77565b82815260059290921b84018101918181019086841115614c0f57600080fd5b8286015b84811015614c2a5780358352918301918301614c13565b509695505050505050565b6000806000838503610120811215614c4c57600080fd5b8435935060208501356001600160401b03811115614c6957600080fd5b614c7587828801614bca565b93505060e0603f1982011215614c8a57600080fd5b506040840190509250925092565b60008083601f840112614caa57600080fd5b5081356001600160401b03811115614cc157600080fd5b6020830191508360208260051b8501011115614cdc57600080fd5b9250929050565b6000806000806000806000878903610120811215614d0057600080fd5b8835975060208901356001600160401b0380821115614d1e57600080fd5b908a01906040828d031215614d3257600080fd5b81985060a0603f1984011215614d4757600080fd5b60408b01975060e08b0135925080831115614d6157600080fd5b614d6d8c848d01614c98565b90975095506101008b0135925086915080831115614d8a57600080fd5b5050614d988a828b01614c98565b989b979a50959850939692959293505050565b600080600080600060a08688031215614dc357600080fd5b614dcc86614a0f565b97602087013597506040870135966060810135965060800135945092505050565b60008060008060608587031215614e0357600080fd5b843593506020850135925060408501356001600160401b0380821115614e2857600080fd5b818701915087601f830112614e3c57600080fd5b813581811115614e4b57600080fd5b886020828501011115614e5d57600080fd5b95989497505060200194505050565b634e487b7160e01b600052602160045260246000fd5b60028110614e9257614e92614e6c565b9052565b60006101c082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015260e083015160e083015261010080840151614f04828501826001600160a01b03169052565b5050610120838101516001600160401b038116848301525050610140838101516001600160401b03811684830152505061016080840151614f4782850182614e82565b50506101808381015160ff16908301526101a08084015180151582850152611d6a565b600060208284031215614f7c57600080fd5b611da082614a0f565b6000808335601e19843603018112614f9c57600080fd5b8301803591506001600160401b03821115614fb657600080fd5b602001915036819003821315614cdc57600080fd5b6001600160401b0381168114614aec57600080fd5b600082601f830112614ff157600080fd5b614ff9614b55565b80604084018581111561500b57600080fd5b845b81811015612e6c57803561502081614fcb565b84526020938401930161500d565b600081830360e081121561504157600080fd5b615049614b2d565b915060a081121561505957600080fd5b615061614b55565b608082121561506f57600080fd5b615077614b55565b915084601f85011261508857600080fd5b615090614b55565b8060408601878111156150a257600080fd5b865b818110156150bc5780358452602093840193016150a4565b508185526150ca8882614fe0565b60208601525050508181526150e160808501614aef565b6020820152808352505060a0820135602082015260c0820135604082015292915050565b60008060006101e0848603121561511b57600080fd5b83356001600160401b0381111561513157600080fd5b61513d86828701614bca565b93505061514d856020860161502e565b915061515d85610100860161502e565b90509250925092565b60038110614aec57614aec614e6c565b614e9281615166565b805180518360005b60028110156151a6578251825260209283019290910190600101615187565b505050602090810151906040840160005b60028110156151dd5783516001600160401b0316825292820192908201906001016151b7565b505082015190506122226080840182615176565b8481526101008101615206602083018661517f565b60c082019390935260e0015292915050565b60006020828403121561522a57600080fd5b81518015158114611da057600080fd5b60006020828403121561524c57600080fd5b8151611da081614fcb565b6020810161526483615166565b91905290565b634e487b7160e01b600052601160045260246000fd5b81810381811115610b9457610b9461526a565b634e487b7160e01b600052603260045260246000fd5b8481526101008101602060408682850137606083016040870160005b60028110156152f45781356152d981614fcb565b6001600160401b0316835291830191908301906001016152c5565b50505050608085013561530681614adf565b61530f81615166565b60a083015260c082019390935260e0015292915050565b6001600160401b03828116828216039080821115613b8d57613b8d61526a565b6001600160a01b0381168114614aec57600080fd5b600060c0820190508382528235602083015260208301356040830152604083013561538581615346565b6001600160a01b03811660608401525060608301356153a381614fcb565b6001600160401b038082166080850152608085013591506153c382614fcb565b80821660a085015250509392505050565b6000602082840312156153e657600080fd5b8135611da081614fcb565b60006020828403121561540357600080fd5b8151611da081615346565b60ff8181168382160190811115610b9457610b9461526a565b82815260408101611da06020830184614e82565b6001600160401b03818116838216019080821115613b8d57613b8d61526a565b60006001820161546d5761546d61526a565b5060010190565b855181526001600160a01b0360208701511660208201526040860151604082015284606082015283608082015260c060a08201528160c0820152818360e0830137600081830160e090810191909152601f909201601f19160101949350505050565b6000602082840312156154e857600080fd5b5051919050565b80820180821115610b9457610b9461526a565b600082601f83011261551357600080fd5b81516020615523614beb83614ba7565b82815260059290921b8401810191818101908684111561554257600080fd5b8286015b84811015614c2a5780518352918301918301615546565b6000806040838503121561557057600080fd5b82516001600160401b038082111561558757600080fd5b61559386838701615502565b935060208501519150808211156155a957600080fd5b506155b685828601615502565b9150509250929050565b60a08101610b94828461517f565b600080600080600060a086880312156155e657600080fd5b853594506020860135935060408601356001600160401b038082111561560b57600080fd5b61561789838a01614bca565b9450606088013591508082111561562d57600080fd5b61563989838a01614bca565b9350608088013591508082111561564f57600080fd5b5061565c88828901614bca565b9150509295509295909350565b6000806040838503121561567c57600080fd5b82356001600160401b038082111561569357600080fd5b61569f86838701614bca565b935060208501359150808211156156b557600080fd5b506155b685828601614bca565b634e487b7160e01b600052600160045260246000fd5b600181815b808511156157135781600019048211156156f9576156f961526a565b8085161561570657918102915b93841c93908002906156dd565b509250929050565b60008261572a57506001610b94565b8161573757506000610b94565b816001811461574d576002811461575757615773565b6001915050610b94565b60ff8411156157685761576861526a565b50506001821b610b94565b5060208310610133831016604e8410600b8410161715615796575081810a610b94565b6157a083836156d8565b80600019048211156157b4576157b461526a565b029392505050565b6000611da0838361571b565b60005b838110156157e35781810151838201526020016157cb565b50506000910152565b600082516157fe8184602087016157c8565b9190910192915050565b60208152600082518060208401526158278160408501602087016157c8565b601f01601f1916919091016040019291505056fea2646970667358221220c6166c323a2b08d2097a71af0d5cc4fc85b36353dc7c8c73fa542edb7a1cf89d64736f6c63430008110033a264697066735822122089b21db83106aba9406bea16b491e20b11192e258f382f4fb83dd1a1b2f7400564736f6c63430008110033",
}

// EdgeChallengeManagerFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use EdgeChallengeManagerFactoryMetaData.ABI instead.
var EdgeChallengeManagerFactoryABI = EdgeChallengeManagerFactoryMetaData.ABI

// EdgeChallengeManagerFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EdgeChallengeManagerFactoryMetaData.Bin instead.
var EdgeChallengeManagerFactoryBin = EdgeChallengeManagerFactoryMetaData.Bin

// DeployEdgeChallengeManagerFactory deploys a new Ethereum contract, binding an instance of EdgeChallengeManagerFactory to it.
func DeployEdgeChallengeManagerFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EdgeChallengeManagerFactory, error) {
	parsed, err := EdgeChallengeManagerFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EdgeChallengeManagerFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EdgeChallengeManagerFactory{EdgeChallengeManagerFactoryCaller: EdgeChallengeManagerFactoryCaller{contract: contract}, EdgeChallengeManagerFactoryTransactor: EdgeChallengeManagerFactoryTransactor{contract: contract}, EdgeChallengeManagerFactoryFilterer: EdgeChallengeManagerFactoryFilterer{contract: contract}}, nil
}

// EdgeChallengeManagerFactory is an auto generated Go binding around an Ethereum contract.
type EdgeChallengeManagerFactory struct {
	EdgeChallengeManagerFactoryCaller     // Read-only binding to the contract
	EdgeChallengeManagerFactoryTransactor // Write-only binding to the contract
	EdgeChallengeManagerFactoryFilterer   // Log filterer for contract events
}

// EdgeChallengeManagerFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EdgeChallengeManagerFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeChallengeManagerFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EdgeChallengeManagerFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeChallengeManagerFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EdgeChallengeManagerFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeChallengeManagerFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EdgeChallengeManagerFactorySession struct {
	Contract     *EdgeChallengeManagerFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// EdgeChallengeManagerFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EdgeChallengeManagerFactoryCallerSession struct {
	Contract *EdgeChallengeManagerFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// EdgeChallengeManagerFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EdgeChallengeManagerFactoryTransactorSession struct {
	Contract     *EdgeChallengeManagerFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// EdgeChallengeManagerFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EdgeChallengeManagerFactoryRaw struct {
	Contract *EdgeChallengeManagerFactory // Generic contract binding to access the raw methods on
}

// EdgeChallengeManagerFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EdgeChallengeManagerFactoryCallerRaw struct {
	Contract *EdgeChallengeManagerFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// EdgeChallengeManagerFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EdgeChallengeManagerFactoryTransactorRaw struct {
	Contract *EdgeChallengeManagerFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEdgeChallengeManagerFactory creates a new instance of EdgeChallengeManagerFactory, bound to a specific deployed contract.
func NewEdgeChallengeManagerFactory(address common.Address, backend bind.ContractBackend) (*EdgeChallengeManagerFactory, error) {
	contract, err := bindEdgeChallengeManagerFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManagerFactory{EdgeChallengeManagerFactoryCaller: EdgeChallengeManagerFactoryCaller{contract: contract}, EdgeChallengeManagerFactoryTransactor: EdgeChallengeManagerFactoryTransactor{contract: contract}, EdgeChallengeManagerFactoryFilterer: EdgeChallengeManagerFactoryFilterer{contract: contract}}, nil
}

// NewEdgeChallengeManagerFactoryCaller creates a new read-only instance of EdgeChallengeManagerFactory, bound to a specific deployed contract.
func NewEdgeChallengeManagerFactoryCaller(address common.Address, caller bind.ContractCaller) (*EdgeChallengeManagerFactoryCaller, error) {
	contract, err := bindEdgeChallengeManagerFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManagerFactoryCaller{contract: contract}, nil
}

// NewEdgeChallengeManagerFactoryTransactor creates a new write-only instance of EdgeChallengeManagerFactory, bound to a specific deployed contract.
func NewEdgeChallengeManagerFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*EdgeChallengeManagerFactoryTransactor, error) {
	contract, err := bindEdgeChallengeManagerFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManagerFactoryTransactor{contract: contract}, nil
}

// NewEdgeChallengeManagerFactoryFilterer creates a new log filterer instance of EdgeChallengeManagerFactory, bound to a specific deployed contract.
func NewEdgeChallengeManagerFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*EdgeChallengeManagerFactoryFilterer, error) {
	contract, err := bindEdgeChallengeManagerFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EdgeChallengeManagerFactoryFilterer{contract: contract}, nil
}

// bindEdgeChallengeManagerFactory binds a generic wrapper to an already deployed contract.
func bindEdgeChallengeManagerFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EdgeChallengeManagerFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EdgeChallengeManagerFactory *EdgeChallengeManagerFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EdgeChallengeManagerFactory.Contract.EdgeChallengeManagerFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EdgeChallengeManagerFactory *EdgeChallengeManagerFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EdgeChallengeManagerFactory.Contract.EdgeChallengeManagerFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EdgeChallengeManagerFactory *EdgeChallengeManagerFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EdgeChallengeManagerFactory.Contract.EdgeChallengeManagerFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EdgeChallengeManagerFactory *EdgeChallengeManagerFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EdgeChallengeManagerFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EdgeChallengeManagerFactory *EdgeChallengeManagerFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EdgeChallengeManagerFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EdgeChallengeManagerFactory *EdgeChallengeManagerFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EdgeChallengeManagerFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateChallengeManager is a paid mutator transaction binding the contract method 0xc9fade8e.
//
// Solidity: function createChallengeManager(address _assertionChain, uint64 _challengePeriodBlocks, address _oneStepProofEntry, uint256 layerZeroBlockEdgeHeight, uint256 layerZeroBigStepEdgeHeight, uint256 layerZeroSmallStepEdgeHeight, address _stakeToken, uint256 _stakeAmount, address _excessStakeReceiver, uint8 _numBigStepLevel) returns(address)
func (_EdgeChallengeManagerFactory *EdgeChallengeManagerFactoryTransactor) CreateChallengeManager(opts *bind.TransactOpts, _assertionChain common.Address, _challengePeriodBlocks uint64, _oneStepProofEntry common.Address, layerZeroBlockEdgeHeight *big.Int, layerZeroBigStepEdgeHeight *big.Int, layerZeroSmallStepEdgeHeight *big.Int, _stakeToken common.Address, _stakeAmount *big.Int, _excessStakeReceiver common.Address, _numBigStepLevel uint8) (*types.Transaction, error) {
	return _EdgeChallengeManagerFactory.contract.Transact(opts, "createChallengeManager", _assertionChain, _challengePeriodBlocks, _oneStepProofEntry, layerZeroBlockEdgeHeight, layerZeroBigStepEdgeHeight, layerZeroSmallStepEdgeHeight, _stakeToken, _stakeAmount, _excessStakeReceiver, _numBigStepLevel)
}

// CreateChallengeManager is a paid mutator transaction binding the contract method 0xc9fade8e.
//
// Solidity: function createChallengeManager(address _assertionChain, uint64 _challengePeriodBlocks, address _oneStepProofEntry, uint256 layerZeroBlockEdgeHeight, uint256 layerZeroBigStepEdgeHeight, uint256 layerZeroSmallStepEdgeHeight, address _stakeToken, uint256 _stakeAmount, address _excessStakeReceiver, uint8 _numBigStepLevel) returns(address)
func (_EdgeChallengeManagerFactory *EdgeChallengeManagerFactorySession) CreateChallengeManager(_assertionChain common.Address, _challengePeriodBlocks uint64, _oneStepProofEntry common.Address, layerZeroBlockEdgeHeight *big.Int, layerZeroBigStepEdgeHeight *big.Int, layerZeroSmallStepEdgeHeight *big.Int, _stakeToken common.Address, _stakeAmount *big.Int, _excessStakeReceiver common.Address, _numBigStepLevel uint8) (*types.Transaction, error) {
	return _EdgeChallengeManagerFactory.Contract.CreateChallengeManager(&_EdgeChallengeManagerFactory.TransactOpts, _assertionChain, _challengePeriodBlocks, _oneStepProofEntry, layerZeroBlockEdgeHeight, layerZeroBigStepEdgeHeight, layerZeroSmallStepEdgeHeight, _stakeToken, _stakeAmount, _excessStakeReceiver, _numBigStepLevel)
}

// CreateChallengeManager is a paid mutator transaction binding the contract method 0xc9fade8e.
//
// Solidity: function createChallengeManager(address _assertionChain, uint64 _challengePeriodBlocks, address _oneStepProofEntry, uint256 layerZeroBlockEdgeHeight, uint256 layerZeroBigStepEdgeHeight, uint256 layerZeroSmallStepEdgeHeight, address _stakeToken, uint256 _stakeAmount, address _excessStakeReceiver, uint8 _numBigStepLevel) returns(address)
func (_EdgeChallengeManagerFactory *EdgeChallengeManagerFactoryTransactorSession) CreateChallengeManager(_assertionChain common.Address, _challengePeriodBlocks uint64, _oneStepProofEntry common.Address, layerZeroBlockEdgeHeight *big.Int, layerZeroBigStepEdgeHeight *big.Int, layerZeroSmallStepEdgeHeight *big.Int, _stakeToken common.Address, _stakeAmount *big.Int, _excessStakeReceiver common.Address, _numBigStepLevel uint8) (*types.Transaction, error) {
	return _EdgeChallengeManagerFactory.Contract.CreateChallengeManager(&_EdgeChallengeManagerFactory.TransactOpts, _assertionChain, _challengePeriodBlocks, _oneStepProofEntry, layerZeroBlockEdgeHeight, layerZeroBigStepEdgeHeight, layerZeroSmallStepEdgeHeight, _stakeToken, _stakeAmount, _excessStakeReceiver, _numBigStepLevel)
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
