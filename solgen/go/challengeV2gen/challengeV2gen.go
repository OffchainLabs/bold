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
	Bin: "0x6101c06040523480156200001257600080fd5b5060405162005d9638038062005d96833981016040819052620000359162000285565b6001600160a01b038a166200005d5760405163641f043160e11b815260040160405180910390fd5b6001600160a01b03808b166101005288166200008c5760405163fb60b0ef60e01b815260040160405180910390fd5b6001600160a01b038816610120526001600160401b038916600003620000c557604051632283bb7360e21b815260040160405180910390fd5b6001600160401b03891660e0526001600160a01b0380851660a05260c0849052821662000105576040516301e1d91560e31b815260040160405180910390fd5b6001600160a01b03821660805262000129876200023f602090811b620016fc17901c565b6200014f57604051633abfb6ff60e21b8152600481018890526024015b60405180910390fd5b8661014081815250506200016e866200023f60201b620016fc1760201c565b6200019057604051633abfb6ff60e21b81526004810187905260240162000146565b856101608181525050620001af856200023f60201b620016fc1760201c565b620001d157604051633abfb6ff60e21b81526004810186905260240162000146565b61018085905260ff8116600003620001fc57604051632a18f5b960e21b815260040160405180910390fd5b60fd8160ff161115620002285760405163040d23bf60e41b815260ff8216600482015260240162000146565b60ff166101a052506200037c975050505050505050565b6000816000036200025257506000919050565b60006200026160018462000354565b929092161592915050565b6001600160a01b03811681146200028257600080fd5b50565b6000806000806000806000806000806101408b8d031215620002a657600080fd5b8a51620002b3816200026c565b60208c0151909a506001600160401b0381168114620002d157600080fd5b60408c0151909950620002e4816200026c565b8098505060608b0151965060808b0151955060a08b0151945060c08b01516200030d816200026c565b60e08c01516101008d0151919550935062000328816200026c565b6101208c015190925060ff811681146200034157600080fd5b809150509295989b9194979a5092959850565b818103818111156200037657634e487b7160e01b600052601160045260246000fd5b92915050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516158c5620004d1600039600081816103af0152818161059e015281816108f90152818161095101528181610abf01528181610cfd0152818161103e015261134f01526000818161052301528181610c37015261139101526000818161028b01528181610bf701526113700152600081816102260152610bb70152600081816102ec015281816108d70152818161092f015261132801526000818161032b01528181610667015281816106fa015281816107840152818161082001528181610dbd01528181610e7201528181610efb01528181610f89015281816111f4015261128b0152600081816102c5015261101d0152600081816103e80152818161099b01526110ff0152600081816103520152818161097a01526110de0152600081816104e901526109e901526158c56000f3fe608060405234801561001057600080fd5b50600436106101ce5760003560e01c80635d9e244411610104578063bce6f54f116100a2578063e94e051e11610071578063e94e051e146104e4578063eae0328b1461050b578063f8ee77d61461051e578063fda2892e1461054557600080fd5b8063bce6f54f14610469578063c32d8c6314610489578063c8bc4e431461049c578063e5b123da146104c457600080fd5b8063748926f3116100de578063748926f31461041d578063750e0c0f146104305780638c1b3a4014610443578063908517e91461045657600080fd5b80635d9e2444146103aa57806360c7dc47146103e357806364deed591461040a57600080fd5b806342e1aaa81161017157806348dd29241161014b57806348dd29241461032657806351ed6a301461034d57806354b64151146103745780635a48e0f41461039757600080fd5b806342e1aaa8146102ad57806346c2781a146102c057806348923bc5146102e757600080fd5b80631dce5166116101ad5780631dce5166146102215780632eaa0043146102485780633e35f5e81461025b578063416e66571461028657600080fd5b80624d8efe146101d357806305fae141146101f95780630f73bfad1461020c575b600080fd5b6101e66101e1366004614a5e565b610565565b6040519081526020015b60405180910390f35b6101e6610207366004614aa8565b610580565b61021f61021a366004614ae2565b610ab6565b005b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b61021f610256366004614b04565b610b38565b61026e610269366004614b04565b610b88565b6040516001600160401b0390911681526020016101f0565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b6101e66102bb366004614b38565b610b9a565b61026e7f000000000000000000000000000000000000000000000000000000000000000081565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101f0565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b610387610382366004614b04565b610c84565b60405190151581526020016101f0565b6101e66103a5366004614b04565b610c90565b6103d17f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff90911681526020016101f0565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b61021f610418366004614c73565b610c9c565b61021f61042b366004614b04565b6110c5565b61038761043e366004614b04565b6111b7565b61021f610451366004614d21565b6111ce565b610387610464366004614b04565b611402565b6101e6610477366004614b04565b60009081526001602052604090205490565b6101e6610497366004614de9565b61140e565b6104af6104aa366004614e2b565b611427565b604080519283526020830191909152016101f0565b6101e66104d2366004614b04565b60009081526002602052604090205490565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b6101e6610519366004614b04565b6115de565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b610558610553366004614b04565b6115f2565b6040516101f09190614ed4565b6000610575878787878787611726565b979650505050505050565b600061058a6148fb565b60006105c261059c6020860186614fa8565b7f000000000000000000000000000000000000000000000000000000000000000061176b565b905060006105cf82610b9a565b90506105d961493f565b60008360028111156105ed576105ed614eaa565b03610926576105ff60a0870187614fc3565b905060000361062157604051630c9ccac560e41b815260040160405180910390fd5b60008061063160a0890189614fc3565b81019061063e9190615143565b80516020820151604080840151905163f9cee9df60e01b81529497509295506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016945063f9cee9df936106a39360608f013593929160040161522f565b60006040518083038186803b1580156106bb57600080fd5b505afa1580156106cf573d6000803e3d6000fd5b505050602080830151845191850151604080870151905163f9cee9df60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016955063f9cee9df9461073394939092909160040161522f565b60006040518083038186803b15801561074b57600080fd5b505afa15801561075f573d6000803e3d6000fd5b505050506040518060c0016040528089606001358152602001826020015181526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e531d8c78b606001356040518263ffffffff1660e01b81526004016107d491815260200190565b602060405180830381865afa1580156107f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108159190615256565b1515815260200160007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166356bbc9e685602001516040518263ffffffff1660e01b815260040161087091815260200190565b602060405180830381865afa15801561088d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b19190615278565b6001600160401b0316118152835160208201528251604090910152925061091d600089857f0000000000000000000000000000000000000000000000000000000000000000887f00000000000000000000000000000000000000000000000000000000000000006117da565b95505050610978565b610975600087837f0000000000000000000000000000000000000000000000000000000000000000867f00000000000000000000000000000000000000000000000000000000000000006117da565b93505b7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b038216158015906109d157508015155b15610a225760008660c001516109e75730610a09565b7f00000000000000000000000000000000000000000000000000000000000000005b9050610a206001600160a01b03841633838561182d565b505b8560400151866020015187600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a489606001518a608001518b60a001518c60c001518d60e00151604051610aa1959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a45050925195945050505050565b610ae3600083837f000000000000000000000000000000000000000000000000000000000000000061189e565b6000828152602081905260409020610afa906119f4565b827fb924f3aa473645c7cf5b10262f927ae4ccf869d7fc239c17144b0c67490d1c7383604051610b2c91815260200190565b60405180910390a35050565b610b43600082611a24565b6000818152602081905260409020610b5a906119f4565b60405182907f0d27fcaf1adc41547a5cfc99d2364f6c0dc7e81c9fc3fe8cb38abb409b48358a90600090a350565b6000610b948183611bb6565b92915050565b600080826002811115610baf57610baf614eaa565b03610bdb57507f0000000000000000000000000000000000000000000000000000000000000000919050565b6001826002811115610bef57610bef614eaa565b03610c1b57507f0000000000000000000000000000000000000000000000000000000000000000919050565b6002826002811115610c2f57610c2f614eaa565b03610c5b57507f0000000000000000000000000000000000000000000000000000000000000000919050565b81604051630efcb87b60e21b8152600401610c769190615295565b60405180910390fd5b919050565b6000610b948183611d40565b6000610b948183611d75565b600080835111610cac5783610cd4565b8260018451610cbb91906152be565b81518110610ccb57610ccb6152d1565b60200260200101515b90506000610ce28183611dc6565b90506000610d218260090160099054906101000a900460ff167f000000000000000000000000000000000000000000000000000000000000000061176b565b90506000816002811115610d3757610d37614eaa565b14610d6657600982015460405163ec72dc5d60e01b8152600160481b90910460ff166004820152602401610c76565b610d6f82611e0f565b610db857610d7c82611e33565b60088301546007840154604051631cb1906160e31b815260048101939093526001600160a01b0390911660248301526044820152606401610c76565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633083622885600701546040518263ffffffff1660e01b8152600401610e0d91815260200190565b602060405180830381865afa158015610e2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e4e9190615256565b9050801561100d57600784015460405163f9cee9df60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163f9cee9df91610eb491908a9060a08201359060c0830135906004016152e7565b60006040518083038186803b158015610ecc57600080fd5b505afa158015610ee0573d6000803e3d6000fd5b5050604051631171558560e01b815260a089013560048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316925063117155859150602401602060405180830381865afa158015610f4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f709190615278565b604051632b5de4f360e11b815260a088013560048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906356bbc9e690602401602060405180830381865afa158015610fd8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ffc9190615278565b6110069190615364565b9150611012565b600091505b6000611062818a8a867f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611e68565b60008a815260208190526040902090915061107c906119f4565b6040516001600160401b03831681528a907f9cd2c77f6772dd0fb07e9972aee8ddddf9fd13f1a65abb03a3de88b07dc59af99060200160405180910390a3505050505050505050565b60006110d18183611dc6565b90506110dc816120ab565b7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0382161580159061113557508015155b15611156576008830154611156906001600160a01b038481169116836121aa565b600084815260208190526040902061116d906119f4565b604080516001600160a01b03851681526020810184905286917fa635398959ddb5ce3b14537edfc25b2e671274c9b8cad0f4bd634752e69007b6910160405180910390a350505050565b6000818152602081905260408120610b94906121df565b60006111da8189611d75565b6040516304972af960e01b81529091506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906304972af99061122b9084908a90600401615399565b60006040518083038186803b15801561124357600080fd5b505afa158015611257573d6000803e3d6000fd5b505050506000604051806060016040528088608001602081019061127b9190615412565b6001600160401b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e78cea926040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061130b919061542f565b6001600160a01b03168152883560209091015290506113b560008a7f00000000000000000000000000000000000000000000000000000000000000008b858b8b8b8b7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006121f8565b60008981526020819052604090206113cc906119f4565b6040518a907fe11db4b27bc8c6ea5943ecbb205ae1ca8d56c42c719717aaf8a53d43d0cee7c290600090a3505050505050505050565b6000610b94818361252e565b600061141d86868686866125e9565b9695505050505050565b6000806000806000611473898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509594939250506126439050565b81519295509093509150158061150b578260400151836020015184600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a4866060015187608001518860a001518960c001518a60e00151604051611502959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a45b8160400151826020015183600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a4856060015186608001518760a001518860c001518960e0015160405161158a959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a48151604051821515815285908c907f7340510d24b7ec9b5c100f5500d93429d80d00d46f0d18e4e85d0c4cc22b99249060200160405180910390a4505191989197509095505050505050565b6000610b946115ed8284611dc6565b612927565b6115fa614981565b611605600083611dc6565b604080516101c0810182528254815260018084015460208301526002840154928201929092526003830154606082015260048301546080820152600583015460a0820152600683015460c0820152600783015460e082015260088301546001600160a01b038116610100830152600160a01b90046001600160401b039081166101208301526009840154908116610140830152909291610160840191600160401b900460ff16908111156116bb576116bb614eaa565b60018111156116cc576116cc614eaa565b81526009919091015460ff600160481b820481166020840152600160501b90910416151560409091015292915050565b60008160000361170e57506000919050565b600061171b6001846152be565b929092161592915050565b600061173587878787876125e9565b60408051602081019290925281018390526060016040516020818303038152906040528051906020012090509695505050505050565b60008260ff1660000361178057506000610b94565b8160ff168360ff161161179557506001610b94565b6117a082600161544c565b60ff168360ff16036117b457506002610b94565b6040516315c1b4af60e31b815260ff808516600483015283166024820152604401610c76565b6117e26148fb565b6000806117f2898989898861296c565b915091506000611803838a88612de3565b9050600061181283838c612ef6565b905061181e8b82612f31565b9b9a5050505050505050505050565b6040516001600160a01b03808516602483015283166044820152606481018290526118989085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526131fd565b50505050565b60008381526020859052604090206118b5906121df565b6118d45760405162a7b02b60e01b815260048101849052602401610c76565b60008281526020859052604090206118eb906121df565b61190a5760405162a7b02b60e01b815260048101849052602401610c76565b6001600083815260208690526040902060090154600160401b900460ff16600181111561193957611939614eaa565b146119785760008281526020859052604090819020600901549051633bc499ed60e21b8152610c76918491600160401b90910460ff1690600401615465565b611984848484846132cf565b60008281526020859052604090206007015483146119d35760008281526020859052604090819020600701549051631855b87d60e31b8152610c76918591600401918252602082015260400190565b6119dd84846133f7565b600083815260208590526040902061189890613461565b6000610b948260090160099054906101000a900460ff1683600001548460020154856001015486600401546125e9565b6000818152602083905260409020611a3b906121df565b611a5a5760405162a7b02b60e01b815260048101829052602401610c76565b600081815260208390526040808220600501548083529120611a7b906121df565b611a9a5760405162a7b02b60e01b815260048101829052602401610c76565b6001600082815260208590526040902060090154600160401b900460ff166001811115611ac957611ac9614eaa565b14611b085760008181526020849052604090819020600901549051633bc499ed60e21b8152610c76918391600160401b90910460ff1690600401615465565b600082815260208490526040808220600601548083529120611b29906121df565b611b485760405162a7b02b60e01b815260048101829052602401610c76565b6001600082815260208690526040902060090154600160401b900460ff166001811115611b7757611b77614eaa565b146119d35760008181526020859052604090819020600901549051633bc499ed60e21b8152610c76918391600160401b90910460ff1690600401615465565b6000818152602083905260408120611bcd906121df565b611bec5760405162a7b02b60e01b815260048101839052602401610c76565b6000828152602084905260408120611c03906119f4565b6000818152600186016020526040812054919250819003611c37576040516336843d9f60e21b815260040160405180910390fd5b6040516815539492559053115160ba1b6020820152602901604051602081830303815290604052805190602001208103611ca057600084815260208690526040902060080154611c9790600160a01b90046001600160401b031643615364565b92505050610b94565b6000818152602086905260409020611cb7906121df565b611cd65760405162a7b02b60e01b815260048101829052602401610c76565b60008181526020869052604080822060089081015487845291909220909101546001600160401b03600160a01b928390048116929091041680821115611d2b57611d208183615364565b945050505050610b94565b6000945050505050610b94565b505092915050565b6000611d4c838361252e565b8015611d6e57506000828152602084905260409020611d6a90612927565b6001145b9392505050565b600080611d828484611dc6565b90505b6009810154600160481b900460ff1615611dbe5780546000908152600185016020526040902054611db68582611dc6565b915050611d85565b549392505050565b6000818152602083905260408120611ddd906121df565b611dfc5760405162a7b02b60e01b815260048101839052602401610c76565b5060009081526020919091526040902090565b600781015460009015801590610b94575050600801546001600160a01b0316151590565b6000610b948260090160099054906101000a900460ff1683600001548460020154856001015486600401548760030154611726565b6000858152602087905260408120611e7f906121df565b611e9e5760405162a7b02b60e01b815260048101879052602401610c76565b856000611eab8983611bb6565b905060005b875181101561202b576000611ede8b8a8481518110611ed157611ed16152d1565b6020026020010151611dc6565b90508381600501541480611ef55750838160060154145b15611f3957611f0c8b611f0783611e33565b611bb6565b611f169084615479565b9250888281518110611f2a57611f2a6152d1565b60200260200101519350612018565b600084815260208c9052604090206007015489518a9084908110611f5f57611f5f6152d1565b602002602001015103611f9e57611f918b8a8481518110611f8257611f826152d1565b602002602001015186896132cf565b611f0c8b611f0783611e33565b83816005015482600601548b8581518110611fbb57611fbb6152d1565b60200260200101518e600001600089815260200190815260200160002060070154604051636ebd28c960e01b8152600401610c76959493929190948552602085019390935260408401919091526060830152608082015260a00190565b508061202381615499565b915050611eb0565b506120368682615479565b9050846001600160401b0316816001600160401b0316101561207e5760405163011a8d4d60e41b81526001600160401b03808316600483015286166024820152604401610c76565b61208889896133f7565b600088815260208a90526040902061209f90613461565b98975050505050505050565b60016009820154600160401b900460ff1660018111156120cd576120cd614eaa565b14612105576120db81611e33565b6009820154604051633bc499ed60e21b8152610c769291600160401b900460ff1690600401615465565b61210e81611e0f565b6121575761211b81611e33565b60088201546007830154604051631cb1906160e31b815260048101939093526001600160a01b0390911660248301526044820152606401610c76565b6009810154600160501b900460ff1615156001036121945761217881611e33565b60405163307f766960e01b8152600401610c7691815260200190565b600901805460ff60501b1916600160501b179055565b6040516001600160a01b0383166024820152604481018290526121da90849063a9059cbb60e01b90606401611861565b505050565b60080154600160a01b90046001600160401b0316151590565b60008b815260208d90526040902061220f906121df565b61222e5760405162a7b02b60e01b8152600481018c9052602401610c76565b600260008c815260208e9052604090206009015461225690600160481b900460ff168561176b565b600281111561226757612267614eaa565b146122a45760008b815260208d905260409081902060090154905163348aefdf60e01b8152600160481b90910460ff166004820152602401610c76565b60008b815260208d9052604090206122bb90612927565b6001146122f55760008b815260208d9052604090206122d990612927565b6040516306b595e560e41b8152600401610c7691815260200190565b60008b815260208d905260409020600201548b825b60018f600001600084815260200190815260200160002060090160099054906101000a900460ff1660ff1611156123b25760008f60000160008481526020019081526020016000206000015490508f60010160008281526020019081526020016000205492508f6000016000848152602001908152602001600020600201548261239491906154b2565b61239e90856154c9565b93506123aa86836154b2565b91505061230a565b505061240f8d60000160008e8152602001908152602001600020600101548b60000135838b8b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506134e392505050565b60008b6001600160a01b031663b5112fd28b848e600001358f80602001906124379190614fc3565b6040518663ffffffff1660e01b81526004016124579594939291906154dc565b602060405180830381865afa158015612474573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612498919061553d565b90506124fd8e60000160008f815260200190815260200160002060030154828460016124c491906154c9565b8a8a808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506134e392505050565b6125078e8e6133f7565b60008d815260208f90526040902061251e90613461565b5050505050505050505050505050565b6000818152602083905260408120612545906121df565b6125645760405162a7b02b60e01b815260048101839052602401610c76565b600082815260208490526040812061257b906119f4565b60008181526001860160205260408120549192508190036125af576040516336843d9f60e21b815260040160405180910390fd5b6040516815539492559053115160ba1b602082015260290160408051601f1981840301815291905280516020909101201415949350505050565b6040516001600160f81b031960f887901b1660208201526021810185905260418101849052606181018390526081810182905260009060a10160405160208183030381529060405280519060200120905095945050505050565b600061264d6148fb565b6126556148fb565b60008087815260208990526040902060090154600160401b900460ff16600181111561268357612683614eaa565b146126c257600086815260208890526040908190206009015490516323f8405d60e01b8152610c76918891600160401b90910460ff1690600401615465565b6126cc878761252e565b6126ec576040516380e07e4560e01b815260048101879052602401610c76565b6000868152602088905260408120604080516101c0810182528254815260018084015460208301526002840154928201929092526003830154606082015260048301546080820152600583015460a0820152600683015460c0820152600783015460e082015260088301546001600160a01b038116610100830152600160a01b90046001600160401b039081166101208301526009840154908116610140830152909291610160840191600160401b900460ff16908111156127b0576127b0614eaa565b60018111156127c1576127c1614eaa565b81526009919091015460ff600160481b820481166020840152600160501b90910416151560409182015281015160808201519192506000916128039190613570565b90506000808780602001905181019061281c91906155b1565b909250905061284c896128308560016154c9565b606087015160808801516128459060016154c9565b8686613604565b505060006128586148fb565b60006128798560000151866020015187604001518d888a61018001516138d5565b905061288481613965565b600081815260208e90526040902090935061289e906121df565b6128af576128ac8c82612f31565b91505b506128b86148fb565b60006128d986600001518c8789606001518a608001518b61018001516138d5565b90506128e58d82612f31565b9150506129158382600001518e60000160008f815260200190815260200160002061398e9092919063ffffffff16565b919b909a509098509650505050505050565b6000808260020154836004015461293e91906152be565b905080600003610b945761295183611e33565b60405162a7b02b60e01b8152600401610c7691815260200190565b604080516060808201835260008083526020830152918101919091526000806129a161299b6020890189614fa8565b8561176b565b60028111156129b2576129b2614eaa565b03612c2b57602085015185516000036129de576040516374b5e30d60e11b815260040160405180910390fd5b8551606088013514612a135785516040516316c5de8f60e21b8152600481019190915260608801356024820152604401610c76565b8560400151612a35576040516360b4921b60e11b815260040160405180910390fd5b8560600151612a5757604051635a2e8e1d60e11b815260040160405180910390fd5b612a6460a0880188614fc3565b9050600003612a8657604051630c9ccac560e41b815260040160405180910390fd5b6000612a9560a0890189614fc3565b810190612aa29190615143565b50909150600090508760800151602001516002811115612ac457612ac4614eaa565b03612ae25760405163231b2f2960e11b815260040160405180910390fd5b60008760a00151602001516002811115612afe57612afe614eaa565b03612b1c57604051638999857d60e01b815260040160405180910390fd5b60808701516040516330e5867160e21b81526000916001600160a01b0389169163c39619c491612b4e91600401615614565b602060405180830381865afa158015612b6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b8f919061553d565b90506000876001600160a01b031663c39619c48a60a001516040518263ffffffff1660e01b8152600401612bc39190615614565b602060405180830381865afa158015612be0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c04919061553d565b6040805160608101825293845260208401919091528201929092529350909150612dd99050565b612c39878760600135611d40565b612c5f5760405160016292642960e01b0319815260608701356004820152602401610c76565b6060860135600090815260208890526040812090612c7c826119f4565b905060006009830154600160401b900460ff166001811115612ca057612ca0614eaa565b14612cbe576040516312459ffd60e01b815260040160405180910390fd5b6009820154612cd790600160481b900460ff16866139f5565b60ff16612ce760208a018a614fa8565b60ff1614612d3057612cfc6020890189614fa8565b600983015460405163564f308b60e11b815260ff9283166004820152600160481b9091049091166024820152604401610c76565b612d3d60a0890189614fc3565b9050600003612d5f57604051630c9ccac560e41b815260040160405180910390fd5b600080808080612d7260a08e018e614fc3565b810190612d7f9190615622565b94509450945094509450612d9d8760010154868960020154866134e3565b612db18760030154858960040154856134e3565b6040518060600160405280868152602001858152602001828152508698509850505050505050505b9550959350505050565b604080516000808252602082019092528190612e0990612e04908751613a17565b613a4d565b9050612e14836116fc565b612e3457604051633abfb6ff60e21b815260048101849052602401610c76565b82846040013514612e6557604080516337f318af60e21b815290850135600482015260248101849052604401610c76565b612e8184602001358660200151866040013588604001516134e3565b612e8e6080850185614fc3565b9050600003612eb057604051631a1503a960e11b815260040160405180910390fd5b600080612ec06080870187614fc3565b810190612ecd91906156bd565b9092509050612eeb836001602089013561284560408b0135836154c9565b509095945050505050565b612efe614981565b612f2984846000602086018035906040880135906060890135903390612f24908b614fa8565b613be6565b949350505050565b612f396148fb565b6000612f4483613965565b6000818152602086905260409020909150612f5e906121df565b15612f7f57604051635e76f9ef60e11b815260048101829052602401610c76565b600081815260208581526040918290208551815590850151600180830191909155918501516002820155606085015160038201556080850151600482015560a0850151600582015560c0850151600682015560e085015160078201556101008501516008820180546101208801516001600160401b03908116600160a01b026001600160e01b03199092166001600160a01b03909416939093171790556101408601516009830180549190921667ffffffffffffffff19821681178355610160880151889590939268ffffffffffffffffff191690911790600160401b90849081111561306e5761306e614eaa565b021790555061018082810151600990920180546101a0909401511515600160501b0260ff60501b1960ff909416600160481b02939093166affff00000000000000000019909416939093179190911790915583015183516040850151602086015160808701516000946130e6949093909290916125e9565b6000818152600187016020526040812054919250819003613145576040516815539492559053115160ba1b602082015260290160408051601f198184030181529181528151602092830120600085815260018a0190935291205561318d565b6040516815539492559053115160ba1b602082015260290160405160208183030381529060405280519060200120810361318d57600082815260018701602052604090208390555b6040805161010081018252848152602080820185905287518284015260e088015160608301526000868152908990529190912060808201906131ce90612927565b815261018087015160ff166020820152911515604083015260e090950151151560609091015250919392505050565b6000613252826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613ce89092919063ffffffff16565b8051909150156121da57808060200190518101906132709190615256565b6121da5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610c76565b6000828152602085905260408082205485835291206132ed906119f4565b1461333b576000838152602085905260409020613309906119f4565b6000838152602086905260409081902054905163e2e27f8760e01b815260048101929092526024820152604401610c76565b600082815260208590526040808220600990810154868452919092209091015460ff600160481b92839004811692613375920416836139f5565b60ff161461189857600083815260208590526040902060090154839083906133a790600160481b900460ff16846139f5565b60008581526020889052604090819020600901549051637e726d1560e01b81526004810194909452602484019290925260ff9081166044840152600160481b909104166064820152608401610c76565b600081815260208390526040812061340e906119f4565b6000818152600285016020526040902054909150801561344b57604051630dd7028f60e41b81526004810184905260248101829052604401610c76565b5060009081526002909201602052604090912055565b60006009820154600160401b900460ff16600181111561348357613483614eaa565b146134bb5761349181611e33565b60098201546040516323f8405d60e01b8152610c769291600160401b900460ff1690600401615465565b60090180546001600160401b03431668ffffffffffffffffff1990911617600160401b179055565b60006135188284866040516020016134fd91815260200190565b60405160208183030381529060405280519060200120613cf7565b90508085146135695760405162461bcd60e51b815260206004820152601760248201527f496e76616c696420696e636c7573696f6e2070726f6f660000000000000000006044820152606401610c76565b5050505050565b6000600261357e84846152be565b10156135a75760405163240a616560e21b81526004810184905260248101839052604401610c76565b6135b183836152be565b6002036135ca576135c38360016154c9565b9050610b94565b6000836135d86001856152be565b18905060006135e682613d99565b9050600019811b806135f96001876152be565b169695505050505050565b6000851161364b5760405162461bcd60e51b815260206004820152601460248201527305072652d73697a652063616e6e6f7420626520360641b6044820152606401610c76565b8561365583613a4d565b146136a25760405162461bcd60e51b815260206004820152601b60248201527f50726520657870616e73696f6e20726f6f74206d69736d6174636800000000006044820152606401610c76565b846136ac83613ec8565b146137035760405162461bcd60e51b815260206004820152602160248201527f5072652073697a6520646f6573206e6f74206d6174636820657870616e73696f6044820152603760f91b6064820152608401610c76565b8285106137525760405162461bcd60e51b815260206004820181905260248201527f5072652073697a65206e6f74206c657373207468616e20706f73742073697a656044820152606401610c76565b60008590506000806137678560008751613f23565b90505b8583101561381f57600061377e848861407f565b9050845183106137c55760405162461bcd60e51b8152602060048201526012602482015271496e646578206f7574206f662072616e676560701b6044820152606401610c76565b6137e982828786815181106137dc576137dc6152d1565b6020026020010151614163565b91506001811b6137f981866154c9565b94508785111561380b5761380b615716565b8361381581615499565b945050505061376a565b8661382982613a4d565b146138815760405162461bcd60e51b815260206004820152602260248201527f506f737420657870616e73696f6e20726f6f74206e6f7420657175616c20706f6044820152611cdd60f21b6064820152608401610c76565b835182146138ca5760405162461bcd60e51b8152602060048201526016602482015275496e636f6d706c6574652070726f6f6620757361676560501b6044820152606401610c76565b505050505050505050565b6138dd614981565b6138ea878787878761469f565b50604080516101c08101825296875260208701959095529385019290925260608401526080830152600060a0830181905260c0830181905260e0830181905261010083018190526001600160401b0343166101208401526101408301819052610160830181905260ff9091166101808301526101a082015290565b6000610b9482610180015183600001518460400151856020015186608001518760600151611726565b60058301541515806139a35750600683015415155b156139e5576139b183611e33565b600584015460068501546040516308b0e71d60e41b8152600481019390935260248301919091526044820152606401610c76565b6005830191909155600690910155565b600080613a0384600161544c565b9050613a0f818461176b565b509392505050565b6060611d6e83600084604051602001613a3291815260200190565b60405160208183030381529060405280519060200120614163565b600080825111613a985760405162461bcd60e51b815260206004820152601660248201527522b6b83a3c9036b2b935b6329032bc3830b739b4b7b760511b6044820152606401610c76565b604082511115613aea5760405162461bcd60e51b815260206004820152601a60248201527f4d65726b6c6520657870616e73696f6e20746f6f206c617267650000000000006044820152606401610c76565b6000805b8351811015613bdf576000848281518110613b0b57613b0b6152d1565b60200260200101519050826000801b03613b77578015613b725780925060018551613b3691906152be565b8214613b7257604051613b59908490600090602001918252602082015260400190565b6040516020818303038152906040528051906020012092505b613bcc565b8015613b96576040805160208101839052908101849052606001613b59565b604051613bb3908490600090602001918252602082015260400190565b6040516020818303038152906040528051906020012092505b5080613bd781615499565b915050613aee565b5092915050565b613bee614981565b6001600160a01b038316613c155760405163f289e65760e01b815260040160405180910390fd5b6000849003613c3757604051636932bcfd60e01b815260040160405180910390fd5b613c44898989898961469f565b604051806101c001604052808a81526020018981526020018881526020018781526020018681526020016000801b81526020016000801b8152602001858152602001846001600160a01b03168152602001436001600160401b0316815260200160006001600160401b0316815260200160006001811115613cc757613cc7614eaa565b815260ff841660208201526000604090910152905098975050505050505050565b6060612f29848460008561472f565b8251600090610100811115613d2a57604051637ed6198f60e11b8152600481018290526101006024820152604401610c76565b8260005b82811015613d8f576000878281518110613d4a57613d4a6152d1565b60200260200101519050816001901b8716600003613d7657826000528060205260406000209250613d86565b8060005282602052604060002092505b50600101613d2e565b5095945050505050565b600081600003613deb5760405162461bcd60e51b815260206004820152601c60248201527f5a65726f20686173206e6f207369676e69666963616e742062697473000000006044820152606401610c76565b600160801b8210613e0957608091821c91613e0690826154c9565b90505b600160401b8210613e2757604091821c91613e2490826154c9565b90505b6401000000008210613e4657602091821c91613e4390826154c9565b90505b620100008210613e6357601091821c91613e6090826154c9565b90505b6101008210613e7f57600891821c91613e7c90826154c9565b90505b60108210613e9a57600491821c91613e9790826154c9565b90505b60048210613eb557600291821c91613eb290826154c9565b90505b60028210610c7f57610b946001826154c9565b600080805b8351811015613bdf57838181518110613ee857613ee86152d1565b60200260200101516000801b14613f1157613f04816002615810565b613f0e90836154c9565b91505b80613f1b81615499565b915050613ecd565b6060818310613f6e5760405162461bcd60e51b815260206004820152601760248201527614dd185c9d081b9bdd081b195cdcc81d1a185b88195b99604a1b6044820152606401610c76565b8351821115613fc95760405162461bcd60e51b815260206004820152602160248201527f456e64206e6f74206c657373206f7220657175616c207468616e206c656e67746044820152600d60fb1b6064820152608401610c76565b6000613fd584846152be565b6001600160401b03811115613fec57613fec614b55565b604051908082528060200260200182016040528015614015578160200160208202803683370190505b509050835b8381101561407657858181518110614034576140346152d1565b602002602001015182868361404991906152be565b81518110614059576140596152d1565b60209081029190910101528061406e81615499565b91505061401a565b50949350505050565b60008183106140ca5760405162461bcd60e51b815260206004820152601760248201527614dd185c9d081b9bdd081b195cdcc81d1a185b88195b99604a1b6044820152606401610c76565b60006140d7838518613d99565b9050600060016140e783826154c9565b6001901b6140f591906152be565b9050848116848216811561410c57611d2082614855565b801561411b57611d2081613d99565b60405162461bcd60e51b815260206004820152601b60248201527f426f7468207920616e64207a2063616e6e6f74206265207a65726f00000000006044820152606401610c76565b6060604083106141a65760405162461bcd60e51b815260206004820152600e60248201526d098caeccad840e8dede40d0d2ced60931b6044820152606401610c76565b60008290036141f75760405162461bcd60e51b815260206004820152601b60248201527f43616e6e6f7420617070656e6420656d707479207375627472656500000000006044820152606401610c76565b6040845111156142495760405162461bcd60e51b815260206004820152601a60248201527f4d65726b6c6520657870616e73696f6e20746f6f206c617267650000000000006044820152606401610c76565b83516000036142c757600061425f8460016154c9565b6001600160401b0381111561427657614276614b55565b60405190808252806020026020018201604052801561429f578160200160208202803683370190505b509050828185815181106142b5576142b56152d1565b60209081029190910101529050611d6e565b835183106143355760405162461bcd60e51b815260206004820152603560248201527f4c6576656c2067726561746572207468616e2068696768657374206c6576656c6044820152741037b31031bab93932b73a1032bc3830b739b4b7b760591b6064820152608401610c76565b81600061434186613ec8565b90506000614350866002615810565b61435a90836154c9565b9050600061436783613d99565b61437083613d99565b116143bd5787516001600160401b0381111561438e5761438e614b55565b6040519080825280602002602001820160405280156143b7578160200160208202803683370190505b5061440c565b87516143ca9060016154c9565b6001600160401b038111156143e1576143e1614b55565b60405190808252806020026020018201604052801561440a578160200160208202803683370190505b505b90506040815111156144605760405162461bcd60e51b815260206004820152601c60248201527f417070656e642063726561746573206f76657273697a652074726565000000006044820152606401610c76565b60005b885181101561460157878110156144ef57888181518110614486576144866152d1565b60200260200101516000801b146144ea5760405162461bcd60e51b815260206004820152602260248201527f417070656e642061626f7665206c65617374207369676e69666963616e7420626044820152611a5d60f21b6064820152608401610c76565b6145ef565b60008590036145355788818151811061450a5761450a6152d1565b6020026020010151828281518110614524576145246152d1565b6020026020010181815250506145ef565b888181518110614547576145476152d1565b60200260200101516000801b0361457f578482828151811061456b5761456b6152d1565b6020908102919091010152600094506145ef565b6000801b828281518110614595576145956152d1565b6020026020010181815250508881815181106145b3576145b36152d1565b6020026020010151856040516020016145d6929190918252602082015260400190565b6040516020818303038152906040528051906020012094505b806145f981615499565b915050614463565b5083156146355783816001835161461891906152be565b81518110614628576146286152d1565b6020026020010181815250505b806001825161464491906152be565b81518110614654576146546152d1565b60200260200101516000801b036105755760405162461bcd60e51b815260206004820152600f60248201526e4c61737420656e747279207a65726f60881b6044820152606401610c76565b60008590036146c15760405163235e76ef60e21b815260040160405180910390fd5b8281116146eb576040516308183ebd60e21b81526004810184905260248101829052604401610c76565b600084900361470d576040516320f1a0f960e21b815260040160405180910390fd5b600082900361356957604051635cb6e5bb60e01b815260040160405180910390fd5b6060824710156147905760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610c76565b6001600160a01b0385163b6147e75760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610c76565b600080866001600160a01b031685876040516148039190615840565b60006040518083038185875af1925050503d8060008114614840576040519150601f19603f3d011682016040523d82523d6000602084013e614845565b606091505b50915091506105758282866148c2565b60008082116148a65760405162461bcd60e51b815260206004820152601c60248201527f5a65726f20686173206e6f207369676e69666963616e742062697473000000006044820152606401610c76565b600082806148b56001826152be565b16189050611d6e81613d99565b606083156148d1575081611d6e565b8251156148e15782518084602001fd5b8160405162461bcd60e51b8152600401610c76919061585c565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915290565b6040805160c08101825260008082526020820181905291810182905260608101919091526080810161496f6149f4565b815260200161497c6149f4565b905290565b604080516101c081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905290610160820190815260006020820181905260409091015290565b6040518060400160405280614a07614a13565b81526020016000905290565b6040518060400160405280614a26614a2f565b815260200161497c5b60405180604001604052806002906020820280368337509192915050565b803560ff81168114610c7f57600080fd5b60008060008060008060c08789031215614a7757600080fd5b614a8087614a4d565b9860208801359850604088013597606081013597506080810135965060a00135945092505050565b600060208284031215614aba57600080fd5b81356001600160401b03811115614ad057600080fd5b820160c08185031215611d6e57600080fd5b60008060408385031215614af557600080fd5b50508035926020909101359150565b600060208284031215614b1657600080fd5b5035919050565b60038110614b2a57600080fd5b50565b8035610c7f81614b1d565b600060208284031215614b4a57600080fd5b8135611d6e81614b1d565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715614b8d57614b8d614b55565b60405290565b604080519081016001600160401b0381118282101715614b8d57614b8d614b55565b604051601f8201601f191681016001600160401b0381118282101715614bdd57614bdd614b55565b604052919050565b60006001600160401b03821115614bfe57614bfe614b55565b5060051b60200190565b600082601f830112614c1957600080fd5b81356020614c2e614c2983614be5565b614bb5565b82815260059290921b84018101918181019086841115614c4d57600080fd5b8286015b84811015614c685780358352918301918301614c51565b509695505050505050565b6000806000838503610120811215614c8a57600080fd5b8435935060208501356001600160401b03811115614ca757600080fd5b614cb387828801614c08565b93505060e0603f1982011215614cc857600080fd5b506040840190509250925092565b60008083601f840112614ce857600080fd5b5081356001600160401b03811115614cff57600080fd5b6020830191508360208260051b8501011115614d1a57600080fd5b9250929050565b6000806000806000806000878903610120811215614d3e57600080fd5b8835975060208901356001600160401b0380821115614d5c57600080fd5b908a01906040828d031215614d7057600080fd5b81985060a0603f1984011215614d8557600080fd5b60408b01975060e08b0135925080831115614d9f57600080fd5b614dab8c848d01614cd6565b90975095506101008b0135925086915080831115614dc857600080fd5b5050614dd68a828b01614cd6565b989b979a50959850939692959293505050565b600080600080600060a08688031215614e0157600080fd5b614e0a86614a4d565b97602087013597506040870135966060810135965060800135945092505050565b60008060008060608587031215614e4157600080fd5b843593506020850135925060408501356001600160401b0380821115614e6657600080fd5b818701915087601f830112614e7a57600080fd5b813581811115614e8957600080fd5b886020828501011115614e9b57600080fd5b95989497505060200194505050565b634e487b7160e01b600052602160045260246000fd5b60028110614ed057614ed0614eaa565b9052565b60006101c082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015260e083015160e083015261010080840151614f42828501826001600160a01b03169052565b5050610120838101516001600160401b038116848301525050610140838101516001600160401b03811684830152505061016080840151614f8582850182614ec0565b50506101808381015160ff16908301526101a08084015180151582850152611d38565b600060208284031215614fba57600080fd5b611d6e82614a4d565b6000808335601e19843603018112614fda57600080fd5b8301803591506001600160401b03821115614ff457600080fd5b602001915036819003821315614d1a57600080fd5b6001600160401b0381168114614b2a57600080fd5b600082601f83011261502f57600080fd5b615037614b93565b80604084018581111561504957600080fd5b845b81811015612eeb57803561505e81615009565b84526020938401930161504b565b600081830360e081121561507f57600080fd5b615087614b6b565b915060a081121561509757600080fd5b61509f614b93565b60808212156150ad57600080fd5b6150b5614b93565b915084601f8501126150c657600080fd5b6150ce614b93565b8060408601878111156150e057600080fd5b865b818110156150fa5780358452602093840193016150e2565b50818552615108888261501e565b602086015250505081815261511f60808501614b2d565b6020820152808352505060a0820135602082015260c0820135604082015292915050565b60008060006101e0848603121561515957600080fd5b83356001600160401b0381111561516f57600080fd5b61517b86828701614c08565b93505061518b856020860161506c565b915061519b85610100860161506c565b90509250925092565b60038110614b2a57614b2a614eaa565b614ed0816151a4565b805180518360005b60028110156151e45782518252602092830192909101906001016151c5565b505050602090810151906040840160005b600281101561521b5783516001600160401b0316825292820192908201906001016151f5565b505082015190506121da60808401826151b4565b848152610100810161524460208301866151bd565b60c082019390935260e0015292915050565b60006020828403121561526857600080fd5b81518015158114611d6e57600080fd5b60006020828403121561528a57600080fd5b8151611d6e81615009565b602081016152a2836151a4565b91905290565b634e487b7160e01b600052601160045260246000fd5b81810381811115610b9457610b946152a8565b634e487b7160e01b600052603260045260246000fd5b8481526101008101602060408682850137606083016040870160005b600281101561533257813561531781615009565b6001600160401b031683529183019190830190600101615303565b50505050608085013561534481614b1d565b61534d816151a4565b60a083015260c082019390935260e0015292915050565b6001600160401b03828116828216039080821115613bdf57613bdf6152a8565b6001600160a01b0381168114614b2a57600080fd5b600060c082019050838252823560208301526020830135604083015260408301356153c381615384565b6001600160a01b03166060838101919091528301356153e181615009565b6001600160401b0380821660808501526080850135915061540182615009565b80821660a085015250509392505050565b60006020828403121561542457600080fd5b8135611d6e81615009565b60006020828403121561544157600080fd5b8151611d6e81615384565b60ff8181168382160190811115610b9457610b946152a8565b82815260408101611d6e6020830184614ec0565b6001600160401b03818116838216019080821115613bdf57613bdf6152a8565b6000600182016154ab576154ab6152a8565b5060010190565b8082028115828204841417610b9457610b946152a8565b80820180821115610b9457610b946152a8565b8551815260018060a01b0360208701511660208201526040860151604082015284606082015283608082015260c060a08201528160c0820152818360e0830137600081830160e090810191909152601f909201601f19160101949350505050565b60006020828403121561554f57600080fd5b5051919050565b600082601f83011261556757600080fd5b81516020615577614c2983614be5565b82815260059290921b8401810191818101908684111561559657600080fd5b8286015b84811015614c68578051835291830191830161559a565b600080604083850312156155c457600080fd5b82516001600160401b03808211156155db57600080fd5b6155e786838701615556565b935060208501519150808211156155fd57600080fd5b5061560a85828601615556565b9150509250929050565b60a08101610b9482846151bd565b600080600080600060a0868803121561563a57600080fd5b853594506020860135935060408601356001600160401b038082111561565f57600080fd5b61566b89838a01614c08565b9450606088013591508082111561568157600080fd5b61568d89838a01614c08565b935060808801359150808211156156a357600080fd5b506156b088828901614c08565b9150509295509295909350565b600080604083850312156156d057600080fd5b82356001600160401b03808211156156e757600080fd5b6156f386838701614c08565b9350602085013591508082111561570957600080fd5b5061560a85828601614c08565b634e487b7160e01b600052600160045260246000fd5b600181815b8085111561576757816000190482111561574d5761574d6152a8565b8085161561575a57918102915b93841c9390800290615731565b509250929050565b60008261577e57506001610b94565b8161578b57506000610b94565b81600181146157a157600281146157ab576157c7565b6001915050610b94565b60ff8411156157bc576157bc6152a8565b50506001821b610b94565b5060208310610133831016604e8410600b84101617156157ea575081810a610b94565b6157f4838361572c565b8060001904821115615808576158086152a8565b029392505050565b6000611d6e838361576f565b60005b8381101561583757818101518382015260200161581f565b50506000910152565b6000825161585281846020870161581c565b9190910192915050565b602081526000825180602084015261587b81604085016020870161581c565b601f01601f1916919091016040019291505056fea264697066735822122001730c7c06f3f72068613e12e8a6cf6c3da56311c36ab4ea2d7ae4d5aa4baa9a64736f6c63430008110033",
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
	Bin: "0x608060405234801561001057600080fd5b50615fb6806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063c9fade8e14610030575b600080fd5b61004361003e36600461012b565b61005f565b6040516001600160a01b03909116815260200160405180910390f35b60008a8a8a8a8a8a8a8a8a8a60405161007790610106565b6001600160a01b039a8b16815267ffffffffffffffff90991660208a015296891660408901526060880195909552608087019390935260a0860191909152851660c085015260e08401529290921661010082015260ff90911661012082015261014001604051809103906000f0801580156100f6573d6000803e3d6000fd5b509b9a5050505050505050505050565b615d96806101eb83390190565b6001600160a01b038116811461012857600080fd5b50565b6000806000806000806000806000806101408b8d03121561014b57600080fd5b8a3561015681610113565b995060208b013567ffffffffffffffff8116811461017357600080fd5b985060408b013561018381610113565b975060608b0135965060808b0135955060a08b0135945060c08b01356101a881610113565b935060e08b013592506101008b01356101c081610113565b91506101208b013560ff811681146101d757600080fd5b809150509295989b9194979a509295985056fe6101c06040523480156200001257600080fd5b5060405162005d9638038062005d96833981016040819052620000359162000285565b6001600160a01b038a166200005d5760405163641f043160e11b815260040160405180910390fd5b6001600160a01b03808b166101005288166200008c5760405163fb60b0ef60e01b815260040160405180910390fd5b6001600160a01b038816610120526001600160401b038916600003620000c557604051632283bb7360e21b815260040160405180910390fd5b6001600160401b03891660e0526001600160a01b0380851660a05260c0849052821662000105576040516301e1d91560e31b815260040160405180910390fd5b6001600160a01b03821660805262000129876200023f602090811b620016fc17901c565b6200014f57604051633abfb6ff60e21b8152600481018890526024015b60405180910390fd5b8661014081815250506200016e866200023f60201b620016fc1760201c565b6200019057604051633abfb6ff60e21b81526004810187905260240162000146565b856101608181525050620001af856200023f60201b620016fc1760201c565b620001d157604051633abfb6ff60e21b81526004810186905260240162000146565b61018085905260ff8116600003620001fc57604051632a18f5b960e21b815260040160405180910390fd5b60fd8160ff161115620002285760405163040d23bf60e41b815260ff8216600482015260240162000146565b60ff166101a052506200037c975050505050505050565b6000816000036200025257506000919050565b60006200026160018462000354565b929092161592915050565b6001600160a01b03811681146200028257600080fd5b50565b6000806000806000806000806000806101408b8d031215620002a657600080fd5b8a51620002b3816200026c565b60208c0151909a506001600160401b0381168114620002d157600080fd5b60408c0151909950620002e4816200026c565b8098505060608b0151965060808b0151955060a08b0151945060c08b01516200030d816200026c565b60e08c01516101008d0151919550935062000328816200026c565b6101208c015190925060ff811681146200034157600080fd5b809150509295989b9194979a5092959850565b818103818111156200037657634e487b7160e01b600052601160045260246000fd5b92915050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516158c5620004d1600039600081816103af0152818161059e015281816108f90152818161095101528181610abf01528181610cfd0152818161103e015261134f01526000818161052301528181610c37015261139101526000818161028b01528181610bf701526113700152600081816102260152610bb70152600081816102ec015281816108d70152818161092f015261132801526000818161032b01528181610667015281816106fa015281816107840152818161082001528181610dbd01528181610e7201528181610efb01528181610f89015281816111f4015261128b0152600081816102c5015261101d0152600081816103e80152818161099b01526110ff0152600081816103520152818161097a01526110de0152600081816104e901526109e901526158c56000f3fe608060405234801561001057600080fd5b50600436106101ce5760003560e01c80635d9e244411610104578063bce6f54f116100a2578063e94e051e11610071578063e94e051e146104e4578063eae0328b1461050b578063f8ee77d61461051e578063fda2892e1461054557600080fd5b8063bce6f54f14610469578063c32d8c6314610489578063c8bc4e431461049c578063e5b123da146104c457600080fd5b8063748926f3116100de578063748926f31461041d578063750e0c0f146104305780638c1b3a4014610443578063908517e91461045657600080fd5b80635d9e2444146103aa57806360c7dc47146103e357806364deed591461040a57600080fd5b806342e1aaa81161017157806348dd29241161014b57806348dd29241461032657806351ed6a301461034d57806354b64151146103745780635a48e0f41461039757600080fd5b806342e1aaa8146102ad57806346c2781a146102c057806348923bc5146102e757600080fd5b80631dce5166116101ad5780631dce5166146102215780632eaa0043146102485780633e35f5e81461025b578063416e66571461028657600080fd5b80624d8efe146101d357806305fae141146101f95780630f73bfad1461020c575b600080fd5b6101e66101e1366004614a5e565b610565565b6040519081526020015b60405180910390f35b6101e6610207366004614aa8565b610580565b61021f61021a366004614ae2565b610ab6565b005b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b61021f610256366004614b04565b610b38565b61026e610269366004614b04565b610b88565b6040516001600160401b0390911681526020016101f0565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b6101e66102bb366004614b38565b610b9a565b61026e7f000000000000000000000000000000000000000000000000000000000000000081565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101f0565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b610387610382366004614b04565b610c84565b60405190151581526020016101f0565b6101e66103a5366004614b04565b610c90565b6103d17f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff90911681526020016101f0565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b61021f610418366004614c73565b610c9c565b61021f61042b366004614b04565b6110c5565b61038761043e366004614b04565b6111b7565b61021f610451366004614d21565b6111ce565b610387610464366004614b04565b611402565b6101e6610477366004614b04565b60009081526001602052604090205490565b6101e6610497366004614de9565b61140e565b6104af6104aa366004614e2b565b611427565b604080519283526020830191909152016101f0565b6101e66104d2366004614b04565b60009081526002602052604090205490565b61030e7f000000000000000000000000000000000000000000000000000000000000000081565b6101e6610519366004614b04565b6115de565b6101e67f000000000000000000000000000000000000000000000000000000000000000081565b610558610553366004614b04565b6115f2565b6040516101f09190614ed4565b6000610575878787878787611726565b979650505050505050565b600061058a6148fb565b60006105c261059c6020860186614fa8565b7f000000000000000000000000000000000000000000000000000000000000000061176b565b905060006105cf82610b9a565b90506105d961493f565b60008360028111156105ed576105ed614eaa565b03610926576105ff60a0870187614fc3565b905060000361062157604051630c9ccac560e41b815260040160405180910390fd5b60008061063160a0890189614fc3565b81019061063e9190615143565b80516020820151604080840151905163f9cee9df60e01b81529497509295506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016945063f9cee9df936106a39360608f013593929160040161522f565b60006040518083038186803b1580156106bb57600080fd5b505afa1580156106cf573d6000803e3d6000fd5b505050602080830151845191850151604080870151905163f9cee9df60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016955063f9cee9df9461073394939092909160040161522f565b60006040518083038186803b15801561074b57600080fd5b505afa15801561075f573d6000803e3d6000fd5b505050506040518060c0016040528089606001358152602001826020015181526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e531d8c78b606001356040518263ffffffff1660e01b81526004016107d491815260200190565b602060405180830381865afa1580156107f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108159190615256565b1515815260200160007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166356bbc9e685602001516040518263ffffffff1660e01b815260040161087091815260200190565b602060405180830381865afa15801561088d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b19190615278565b6001600160401b0316118152835160208201528251604090910152925061091d600089857f0000000000000000000000000000000000000000000000000000000000000000887f00000000000000000000000000000000000000000000000000000000000000006117da565b95505050610978565b610975600087837f0000000000000000000000000000000000000000000000000000000000000000867f00000000000000000000000000000000000000000000000000000000000000006117da565b93505b7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b038216158015906109d157508015155b15610a225760008660c001516109e75730610a09565b7f00000000000000000000000000000000000000000000000000000000000000005b9050610a206001600160a01b03841633838561182d565b505b8560400151866020015187600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a489606001518a608001518b60a001518c60c001518d60e00151604051610aa1959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a45050925195945050505050565b610ae3600083837f000000000000000000000000000000000000000000000000000000000000000061189e565b6000828152602081905260409020610afa906119f4565b827fb924f3aa473645c7cf5b10262f927ae4ccf869d7fc239c17144b0c67490d1c7383604051610b2c91815260200190565b60405180910390a35050565b610b43600082611a24565b6000818152602081905260409020610b5a906119f4565b60405182907f0d27fcaf1adc41547a5cfc99d2364f6c0dc7e81c9fc3fe8cb38abb409b48358a90600090a350565b6000610b948183611bb6565b92915050565b600080826002811115610baf57610baf614eaa565b03610bdb57507f0000000000000000000000000000000000000000000000000000000000000000919050565b6001826002811115610bef57610bef614eaa565b03610c1b57507f0000000000000000000000000000000000000000000000000000000000000000919050565b6002826002811115610c2f57610c2f614eaa565b03610c5b57507f0000000000000000000000000000000000000000000000000000000000000000919050565b81604051630efcb87b60e21b8152600401610c769190615295565b60405180910390fd5b919050565b6000610b948183611d40565b6000610b948183611d75565b600080835111610cac5783610cd4565b8260018451610cbb91906152be565b81518110610ccb57610ccb6152d1565b60200260200101515b90506000610ce28183611dc6565b90506000610d218260090160099054906101000a900460ff167f000000000000000000000000000000000000000000000000000000000000000061176b565b90506000816002811115610d3757610d37614eaa565b14610d6657600982015460405163ec72dc5d60e01b8152600160481b90910460ff166004820152602401610c76565b610d6f82611e0f565b610db857610d7c82611e33565b60088301546007840154604051631cb1906160e31b815260048101939093526001600160a01b0390911660248301526044820152606401610c76565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633083622885600701546040518263ffffffff1660e01b8152600401610e0d91815260200190565b602060405180830381865afa158015610e2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e4e9190615256565b9050801561100d57600784015460405163f9cee9df60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163f9cee9df91610eb491908a9060a08201359060c0830135906004016152e7565b60006040518083038186803b158015610ecc57600080fd5b505afa158015610ee0573d6000803e3d6000fd5b5050604051631171558560e01b815260a089013560048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316925063117155859150602401602060405180830381865afa158015610f4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f709190615278565b604051632b5de4f360e11b815260a088013560048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906356bbc9e690602401602060405180830381865afa158015610fd8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ffc9190615278565b6110069190615364565b9150611012565b600091505b6000611062818a8a867f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611e68565b60008a815260208190526040902090915061107c906119f4565b6040516001600160401b03831681528a907f9cd2c77f6772dd0fb07e9972aee8ddddf9fd13f1a65abb03a3de88b07dc59af99060200160405180910390a3505050505050505050565b60006110d18183611dc6565b90506110dc816120ab565b7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0382161580159061113557508015155b15611156576008830154611156906001600160a01b038481169116836121aa565b600084815260208190526040902061116d906119f4565b604080516001600160a01b03851681526020810184905286917fa635398959ddb5ce3b14537edfc25b2e671274c9b8cad0f4bd634752e69007b6910160405180910390a350505050565b6000818152602081905260408120610b94906121df565b60006111da8189611d75565b6040516304972af960e01b81529091506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906304972af99061122b9084908a90600401615399565b60006040518083038186803b15801561124357600080fd5b505afa158015611257573d6000803e3d6000fd5b505050506000604051806060016040528088608001602081019061127b9190615412565b6001600160401b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e78cea926040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061130b919061542f565b6001600160a01b03168152883560209091015290506113b560008a7f00000000000000000000000000000000000000000000000000000000000000008b858b8b8b8b7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006121f8565b60008981526020819052604090206113cc906119f4565b6040518a907fe11db4b27bc8c6ea5943ecbb205ae1ca8d56c42c719717aaf8a53d43d0cee7c290600090a3505050505050505050565b6000610b94818361252e565b600061141d86868686866125e9565b9695505050505050565b6000806000806000611473898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509594939250506126439050565b81519295509093509150158061150b578260400151836020015184600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a4866060015187608001518860a001518960c001518a60e00151604051611502959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a45b8160400151826020015183600001517faa4b66b1ce938c06e2a3f8466bae10ef62e747630e3859889f4719fc6427b5a4856060015186608001518760a001518860c001518960e0015160405161158a959493929190948552602085019390935260ff919091166040840152151560608301521515608082015260a00190565b60405180910390a48151604051821515815285908c907f7340510d24b7ec9b5c100f5500d93429d80d00d46f0d18e4e85d0c4cc22b99249060200160405180910390a4505191989197509095505050505050565b6000610b946115ed8284611dc6565b612927565b6115fa614981565b611605600083611dc6565b604080516101c0810182528254815260018084015460208301526002840154928201929092526003830154606082015260048301546080820152600583015460a0820152600683015460c0820152600783015460e082015260088301546001600160a01b038116610100830152600160a01b90046001600160401b039081166101208301526009840154908116610140830152909291610160840191600160401b900460ff16908111156116bb576116bb614eaa565b60018111156116cc576116cc614eaa565b81526009919091015460ff600160481b820481166020840152600160501b90910416151560409091015292915050565b60008160000361170e57506000919050565b600061171b6001846152be565b929092161592915050565b600061173587878787876125e9565b60408051602081019290925281018390526060016040516020818303038152906040528051906020012090509695505050505050565b60008260ff1660000361178057506000610b94565b8160ff168360ff161161179557506001610b94565b6117a082600161544c565b60ff168360ff16036117b457506002610b94565b6040516315c1b4af60e31b815260ff808516600483015283166024820152604401610c76565b6117e26148fb565b6000806117f2898989898861296c565b915091506000611803838a88612de3565b9050600061181283838c612ef6565b905061181e8b82612f31565b9b9a5050505050505050505050565b6040516001600160a01b03808516602483015283166044820152606481018290526118989085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526131fd565b50505050565b60008381526020859052604090206118b5906121df565b6118d45760405162a7b02b60e01b815260048101849052602401610c76565b60008281526020859052604090206118eb906121df565b61190a5760405162a7b02b60e01b815260048101849052602401610c76565b6001600083815260208690526040902060090154600160401b900460ff16600181111561193957611939614eaa565b146119785760008281526020859052604090819020600901549051633bc499ed60e21b8152610c76918491600160401b90910460ff1690600401615465565b611984848484846132cf565b60008281526020859052604090206007015483146119d35760008281526020859052604090819020600701549051631855b87d60e31b8152610c76918591600401918252602082015260400190565b6119dd84846133f7565b600083815260208590526040902061189890613461565b6000610b948260090160099054906101000a900460ff1683600001548460020154856001015486600401546125e9565b6000818152602083905260409020611a3b906121df565b611a5a5760405162a7b02b60e01b815260048101829052602401610c76565b600081815260208390526040808220600501548083529120611a7b906121df565b611a9a5760405162a7b02b60e01b815260048101829052602401610c76565b6001600082815260208590526040902060090154600160401b900460ff166001811115611ac957611ac9614eaa565b14611b085760008181526020849052604090819020600901549051633bc499ed60e21b8152610c76918391600160401b90910460ff1690600401615465565b600082815260208490526040808220600601548083529120611b29906121df565b611b485760405162a7b02b60e01b815260048101829052602401610c76565b6001600082815260208690526040902060090154600160401b900460ff166001811115611b7757611b77614eaa565b146119d35760008181526020859052604090819020600901549051633bc499ed60e21b8152610c76918391600160401b90910460ff1690600401615465565b6000818152602083905260408120611bcd906121df565b611bec5760405162a7b02b60e01b815260048101839052602401610c76565b6000828152602084905260408120611c03906119f4565b6000818152600186016020526040812054919250819003611c37576040516336843d9f60e21b815260040160405180910390fd5b6040516815539492559053115160ba1b6020820152602901604051602081830303815290604052805190602001208103611ca057600084815260208690526040902060080154611c9790600160a01b90046001600160401b031643615364565b92505050610b94565b6000818152602086905260409020611cb7906121df565b611cd65760405162a7b02b60e01b815260048101829052602401610c76565b60008181526020869052604080822060089081015487845291909220909101546001600160401b03600160a01b928390048116929091041680821115611d2b57611d208183615364565b945050505050610b94565b6000945050505050610b94565b505092915050565b6000611d4c838361252e565b8015611d6e57506000828152602084905260409020611d6a90612927565b6001145b9392505050565b600080611d828484611dc6565b90505b6009810154600160481b900460ff1615611dbe5780546000908152600185016020526040902054611db68582611dc6565b915050611d85565b549392505050565b6000818152602083905260408120611ddd906121df565b611dfc5760405162a7b02b60e01b815260048101839052602401610c76565b5060009081526020919091526040902090565b600781015460009015801590610b94575050600801546001600160a01b0316151590565b6000610b948260090160099054906101000a900460ff1683600001548460020154856001015486600401548760030154611726565b6000858152602087905260408120611e7f906121df565b611e9e5760405162a7b02b60e01b815260048101879052602401610c76565b856000611eab8983611bb6565b905060005b875181101561202b576000611ede8b8a8481518110611ed157611ed16152d1565b6020026020010151611dc6565b90508381600501541480611ef55750838160060154145b15611f3957611f0c8b611f0783611e33565b611bb6565b611f169084615479565b9250888281518110611f2a57611f2a6152d1565b60200260200101519350612018565b600084815260208c9052604090206007015489518a9084908110611f5f57611f5f6152d1565b602002602001015103611f9e57611f918b8a8481518110611f8257611f826152d1565b602002602001015186896132cf565b611f0c8b611f0783611e33565b83816005015482600601548b8581518110611fbb57611fbb6152d1565b60200260200101518e600001600089815260200190815260200160002060070154604051636ebd28c960e01b8152600401610c76959493929190948552602085019390935260408401919091526060830152608082015260a00190565b508061202381615499565b915050611eb0565b506120368682615479565b9050846001600160401b0316816001600160401b0316101561207e5760405163011a8d4d60e41b81526001600160401b03808316600483015286166024820152604401610c76565b61208889896133f7565b600088815260208a90526040902061209f90613461565b98975050505050505050565b60016009820154600160401b900460ff1660018111156120cd576120cd614eaa565b14612105576120db81611e33565b6009820154604051633bc499ed60e21b8152610c769291600160401b900460ff1690600401615465565b61210e81611e0f565b6121575761211b81611e33565b60088201546007830154604051631cb1906160e31b815260048101939093526001600160a01b0390911660248301526044820152606401610c76565b6009810154600160501b900460ff1615156001036121945761217881611e33565b60405163307f766960e01b8152600401610c7691815260200190565b600901805460ff60501b1916600160501b179055565b6040516001600160a01b0383166024820152604481018290526121da90849063a9059cbb60e01b90606401611861565b505050565b60080154600160a01b90046001600160401b0316151590565b60008b815260208d90526040902061220f906121df565b61222e5760405162a7b02b60e01b8152600481018c9052602401610c76565b600260008c815260208e9052604090206009015461225690600160481b900460ff168561176b565b600281111561226757612267614eaa565b146122a45760008b815260208d905260409081902060090154905163348aefdf60e01b8152600160481b90910460ff166004820152602401610c76565b60008b815260208d9052604090206122bb90612927565b6001146122f55760008b815260208d9052604090206122d990612927565b6040516306b595e560e41b8152600401610c7691815260200190565b60008b815260208d905260409020600201548b825b60018f600001600084815260200190815260200160002060090160099054906101000a900460ff1660ff1611156123b25760008f60000160008481526020019081526020016000206000015490508f60010160008281526020019081526020016000205492508f6000016000848152602001908152602001600020600201548261239491906154b2565b61239e90856154c9565b93506123aa86836154b2565b91505061230a565b505061240f8d60000160008e8152602001908152602001600020600101548b60000135838b8b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506134e392505050565b60008b6001600160a01b031663b5112fd28b848e600001358f80602001906124379190614fc3565b6040518663ffffffff1660e01b81526004016124579594939291906154dc565b602060405180830381865afa158015612474573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612498919061553d565b90506124fd8e60000160008f815260200190815260200160002060030154828460016124c491906154c9565b8a8a808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506134e392505050565b6125078e8e6133f7565b60008d815260208f90526040902061251e90613461565b5050505050505050505050505050565b6000818152602083905260408120612545906121df565b6125645760405162a7b02b60e01b815260048101839052602401610c76565b600082815260208490526040812061257b906119f4565b60008181526001860160205260408120549192508190036125af576040516336843d9f60e21b815260040160405180910390fd5b6040516815539492559053115160ba1b602082015260290160408051601f1981840301815291905280516020909101201415949350505050565b6040516001600160f81b031960f887901b1660208201526021810185905260418101849052606181018390526081810182905260009060a10160405160208183030381529060405280519060200120905095945050505050565b600061264d6148fb565b6126556148fb565b60008087815260208990526040902060090154600160401b900460ff16600181111561268357612683614eaa565b146126c257600086815260208890526040908190206009015490516323f8405d60e01b8152610c76918891600160401b90910460ff1690600401615465565b6126cc878761252e565b6126ec576040516380e07e4560e01b815260048101879052602401610c76565b6000868152602088905260408120604080516101c0810182528254815260018084015460208301526002840154928201929092526003830154606082015260048301546080820152600583015460a0820152600683015460c0820152600783015460e082015260088301546001600160a01b038116610100830152600160a01b90046001600160401b039081166101208301526009840154908116610140830152909291610160840191600160401b900460ff16908111156127b0576127b0614eaa565b60018111156127c1576127c1614eaa565b81526009919091015460ff600160481b820481166020840152600160501b90910416151560409182015281015160808201519192506000916128039190613570565b90506000808780602001905181019061281c91906155b1565b909250905061284c896128308560016154c9565b606087015160808801516128459060016154c9565b8686613604565b505060006128586148fb565b60006128798560000151866020015187604001518d888a61018001516138d5565b905061288481613965565b600081815260208e90526040902090935061289e906121df565b6128af576128ac8c82612f31565b91505b506128b86148fb565b60006128d986600001518c8789606001518a608001518b61018001516138d5565b90506128e58d82612f31565b9150506129158382600001518e60000160008f815260200190815260200160002061398e9092919063ffffffff16565b919b909a509098509650505050505050565b6000808260020154836004015461293e91906152be565b905080600003610b945761295183611e33565b60405162a7b02b60e01b8152600401610c7691815260200190565b604080516060808201835260008083526020830152918101919091526000806129a161299b6020890189614fa8565b8561176b565b60028111156129b2576129b2614eaa565b03612c2b57602085015185516000036129de576040516374b5e30d60e11b815260040160405180910390fd5b8551606088013514612a135785516040516316c5de8f60e21b8152600481019190915260608801356024820152604401610c76565b8560400151612a35576040516360b4921b60e11b815260040160405180910390fd5b8560600151612a5757604051635a2e8e1d60e11b815260040160405180910390fd5b612a6460a0880188614fc3565b9050600003612a8657604051630c9ccac560e41b815260040160405180910390fd5b6000612a9560a0890189614fc3565b810190612aa29190615143565b50909150600090508760800151602001516002811115612ac457612ac4614eaa565b03612ae25760405163231b2f2960e11b815260040160405180910390fd5b60008760a00151602001516002811115612afe57612afe614eaa565b03612b1c57604051638999857d60e01b815260040160405180910390fd5b60808701516040516330e5867160e21b81526000916001600160a01b0389169163c39619c491612b4e91600401615614565b602060405180830381865afa158015612b6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b8f919061553d565b90506000876001600160a01b031663c39619c48a60a001516040518263ffffffff1660e01b8152600401612bc39190615614565b602060405180830381865afa158015612be0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c04919061553d565b6040805160608101825293845260208401919091528201929092529350909150612dd99050565b612c39878760600135611d40565b612c5f5760405160016292642960e01b0319815260608701356004820152602401610c76565b6060860135600090815260208890526040812090612c7c826119f4565b905060006009830154600160401b900460ff166001811115612ca057612ca0614eaa565b14612cbe576040516312459ffd60e01b815260040160405180910390fd5b6009820154612cd790600160481b900460ff16866139f5565b60ff16612ce760208a018a614fa8565b60ff1614612d3057612cfc6020890189614fa8565b600983015460405163564f308b60e11b815260ff9283166004820152600160481b9091049091166024820152604401610c76565b612d3d60a0890189614fc3565b9050600003612d5f57604051630c9ccac560e41b815260040160405180910390fd5b600080808080612d7260a08e018e614fc3565b810190612d7f9190615622565b94509450945094509450612d9d8760010154868960020154866134e3565b612db18760030154858960040154856134e3565b6040518060600160405280868152602001858152602001828152508698509850505050505050505b9550959350505050565b604080516000808252602082019092528190612e0990612e04908751613a17565b613a4d565b9050612e14836116fc565b612e3457604051633abfb6ff60e21b815260048101849052602401610c76565b82846040013514612e6557604080516337f318af60e21b815290850135600482015260248101849052604401610c76565b612e8184602001358660200151866040013588604001516134e3565b612e8e6080850185614fc3565b9050600003612eb057604051631a1503a960e11b815260040160405180910390fd5b600080612ec06080870187614fc3565b810190612ecd91906156bd565b9092509050612eeb836001602089013561284560408b0135836154c9565b509095945050505050565b612efe614981565b612f2984846000602086018035906040880135906060890135903390612f24908b614fa8565b613be6565b949350505050565b612f396148fb565b6000612f4483613965565b6000818152602086905260409020909150612f5e906121df565b15612f7f57604051635e76f9ef60e11b815260048101829052602401610c76565b600081815260208581526040918290208551815590850151600180830191909155918501516002820155606085015160038201556080850151600482015560a0850151600582015560c0850151600682015560e085015160078201556101008501516008820180546101208801516001600160401b03908116600160a01b026001600160e01b03199092166001600160a01b03909416939093171790556101408601516009830180549190921667ffffffffffffffff19821681178355610160880151889590939268ffffffffffffffffff191690911790600160401b90849081111561306e5761306e614eaa565b021790555061018082810151600990920180546101a0909401511515600160501b0260ff60501b1960ff909416600160481b02939093166affff00000000000000000019909416939093179190911790915583015183516040850151602086015160808701516000946130e6949093909290916125e9565b6000818152600187016020526040812054919250819003613145576040516815539492559053115160ba1b602082015260290160408051601f198184030181529181528151602092830120600085815260018a0190935291205561318d565b6040516815539492559053115160ba1b602082015260290160405160208183030381529060405280519060200120810361318d57600082815260018701602052604090208390555b6040805161010081018252848152602080820185905287518284015260e088015160608301526000868152908990529190912060808201906131ce90612927565b815261018087015160ff166020820152911515604083015260e090950151151560609091015250919392505050565b6000613252826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613ce89092919063ffffffff16565b8051909150156121da57808060200190518101906132709190615256565b6121da5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610c76565b6000828152602085905260408082205485835291206132ed906119f4565b1461333b576000838152602085905260409020613309906119f4565b6000838152602086905260409081902054905163e2e27f8760e01b815260048101929092526024820152604401610c76565b600082815260208590526040808220600990810154868452919092209091015460ff600160481b92839004811692613375920416836139f5565b60ff161461189857600083815260208590526040902060090154839083906133a790600160481b900460ff16846139f5565b60008581526020889052604090819020600901549051637e726d1560e01b81526004810194909452602484019290925260ff9081166044840152600160481b909104166064820152608401610c76565b600081815260208390526040812061340e906119f4565b6000818152600285016020526040902054909150801561344b57604051630dd7028f60e41b81526004810184905260248101829052604401610c76565b5060009081526002909201602052604090912055565b60006009820154600160401b900460ff16600181111561348357613483614eaa565b146134bb5761349181611e33565b60098201546040516323f8405d60e01b8152610c769291600160401b900460ff1690600401615465565b60090180546001600160401b03431668ffffffffffffffffff1990911617600160401b179055565b60006135188284866040516020016134fd91815260200190565b60405160208183030381529060405280519060200120613cf7565b90508085146135695760405162461bcd60e51b815260206004820152601760248201527f496e76616c696420696e636c7573696f6e2070726f6f660000000000000000006044820152606401610c76565b5050505050565b6000600261357e84846152be565b10156135a75760405163240a616560e21b81526004810184905260248101839052604401610c76565b6135b183836152be565b6002036135ca576135c38360016154c9565b9050610b94565b6000836135d86001856152be565b18905060006135e682613d99565b9050600019811b806135f96001876152be565b169695505050505050565b6000851161364b5760405162461bcd60e51b815260206004820152601460248201527305072652d73697a652063616e6e6f7420626520360641b6044820152606401610c76565b8561365583613a4d565b146136a25760405162461bcd60e51b815260206004820152601b60248201527f50726520657870616e73696f6e20726f6f74206d69736d6174636800000000006044820152606401610c76565b846136ac83613ec8565b146137035760405162461bcd60e51b815260206004820152602160248201527f5072652073697a6520646f6573206e6f74206d6174636820657870616e73696f6044820152603760f91b6064820152608401610c76565b8285106137525760405162461bcd60e51b815260206004820181905260248201527f5072652073697a65206e6f74206c657373207468616e20706f73742073697a656044820152606401610c76565b60008590506000806137678560008751613f23565b90505b8583101561381f57600061377e848861407f565b9050845183106137c55760405162461bcd60e51b8152602060048201526012602482015271496e646578206f7574206f662072616e676560701b6044820152606401610c76565b6137e982828786815181106137dc576137dc6152d1565b6020026020010151614163565b91506001811b6137f981866154c9565b94508785111561380b5761380b615716565b8361381581615499565b945050505061376a565b8661382982613a4d565b146138815760405162461bcd60e51b815260206004820152602260248201527f506f737420657870616e73696f6e20726f6f74206e6f7420657175616c20706f6044820152611cdd60f21b6064820152608401610c76565b835182146138ca5760405162461bcd60e51b8152602060048201526016602482015275496e636f6d706c6574652070726f6f6620757361676560501b6044820152606401610c76565b505050505050505050565b6138dd614981565b6138ea878787878761469f565b50604080516101c08101825296875260208701959095529385019290925260608401526080830152600060a0830181905260c0830181905260e0830181905261010083018190526001600160401b0343166101208401526101408301819052610160830181905260ff9091166101808301526101a082015290565b6000610b9482610180015183600001518460400151856020015186608001518760600151611726565b60058301541515806139a35750600683015415155b156139e5576139b183611e33565b600584015460068501546040516308b0e71d60e41b8152600481019390935260248301919091526044820152606401610c76565b6005830191909155600690910155565b600080613a0384600161544c565b9050613a0f818461176b565b509392505050565b6060611d6e83600084604051602001613a3291815260200190565b60405160208183030381529060405280519060200120614163565b600080825111613a985760405162461bcd60e51b815260206004820152601660248201527522b6b83a3c9036b2b935b6329032bc3830b739b4b7b760511b6044820152606401610c76565b604082511115613aea5760405162461bcd60e51b815260206004820152601a60248201527f4d65726b6c6520657870616e73696f6e20746f6f206c617267650000000000006044820152606401610c76565b6000805b8351811015613bdf576000848281518110613b0b57613b0b6152d1565b60200260200101519050826000801b03613b77578015613b725780925060018551613b3691906152be565b8214613b7257604051613b59908490600090602001918252602082015260400190565b6040516020818303038152906040528051906020012092505b613bcc565b8015613b96576040805160208101839052908101849052606001613b59565b604051613bb3908490600090602001918252602082015260400190565b6040516020818303038152906040528051906020012092505b5080613bd781615499565b915050613aee565b5092915050565b613bee614981565b6001600160a01b038316613c155760405163f289e65760e01b815260040160405180910390fd5b6000849003613c3757604051636932bcfd60e01b815260040160405180910390fd5b613c44898989898961469f565b604051806101c001604052808a81526020018981526020018881526020018781526020018681526020016000801b81526020016000801b8152602001858152602001846001600160a01b03168152602001436001600160401b0316815260200160006001600160401b0316815260200160006001811115613cc757613cc7614eaa565b815260ff841660208201526000604090910152905098975050505050505050565b6060612f29848460008561472f565b8251600090610100811115613d2a57604051637ed6198f60e11b8152600481018290526101006024820152604401610c76565b8260005b82811015613d8f576000878281518110613d4a57613d4a6152d1565b60200260200101519050816001901b8716600003613d7657826000528060205260406000209250613d86565b8060005282602052604060002092505b50600101613d2e565b5095945050505050565b600081600003613deb5760405162461bcd60e51b815260206004820152601c60248201527f5a65726f20686173206e6f207369676e69666963616e742062697473000000006044820152606401610c76565b600160801b8210613e0957608091821c91613e0690826154c9565b90505b600160401b8210613e2757604091821c91613e2490826154c9565b90505b6401000000008210613e4657602091821c91613e4390826154c9565b90505b620100008210613e6357601091821c91613e6090826154c9565b90505b6101008210613e7f57600891821c91613e7c90826154c9565b90505b60108210613e9a57600491821c91613e9790826154c9565b90505b60048210613eb557600291821c91613eb290826154c9565b90505b60028210610c7f57610b946001826154c9565b600080805b8351811015613bdf57838181518110613ee857613ee86152d1565b60200260200101516000801b14613f1157613f04816002615810565b613f0e90836154c9565b91505b80613f1b81615499565b915050613ecd565b6060818310613f6e5760405162461bcd60e51b815260206004820152601760248201527614dd185c9d081b9bdd081b195cdcc81d1a185b88195b99604a1b6044820152606401610c76565b8351821115613fc95760405162461bcd60e51b815260206004820152602160248201527f456e64206e6f74206c657373206f7220657175616c207468616e206c656e67746044820152600d60fb1b6064820152608401610c76565b6000613fd584846152be565b6001600160401b03811115613fec57613fec614b55565b604051908082528060200260200182016040528015614015578160200160208202803683370190505b509050835b8381101561407657858181518110614034576140346152d1565b602002602001015182868361404991906152be565b81518110614059576140596152d1565b60209081029190910101528061406e81615499565b91505061401a565b50949350505050565b60008183106140ca5760405162461bcd60e51b815260206004820152601760248201527614dd185c9d081b9bdd081b195cdcc81d1a185b88195b99604a1b6044820152606401610c76565b60006140d7838518613d99565b9050600060016140e783826154c9565b6001901b6140f591906152be565b9050848116848216811561410c57611d2082614855565b801561411b57611d2081613d99565b60405162461bcd60e51b815260206004820152601b60248201527f426f7468207920616e64207a2063616e6e6f74206265207a65726f00000000006044820152606401610c76565b6060604083106141a65760405162461bcd60e51b815260206004820152600e60248201526d098caeccad840e8dede40d0d2ced60931b6044820152606401610c76565b60008290036141f75760405162461bcd60e51b815260206004820152601b60248201527f43616e6e6f7420617070656e6420656d707479207375627472656500000000006044820152606401610c76565b6040845111156142495760405162461bcd60e51b815260206004820152601a60248201527f4d65726b6c6520657870616e73696f6e20746f6f206c617267650000000000006044820152606401610c76565b83516000036142c757600061425f8460016154c9565b6001600160401b0381111561427657614276614b55565b60405190808252806020026020018201604052801561429f578160200160208202803683370190505b509050828185815181106142b5576142b56152d1565b60209081029190910101529050611d6e565b835183106143355760405162461bcd60e51b815260206004820152603560248201527f4c6576656c2067726561746572207468616e2068696768657374206c6576656c6044820152741037b31031bab93932b73a1032bc3830b739b4b7b760591b6064820152608401610c76565b81600061434186613ec8565b90506000614350866002615810565b61435a90836154c9565b9050600061436783613d99565b61437083613d99565b116143bd5787516001600160401b0381111561438e5761438e614b55565b6040519080825280602002602001820160405280156143b7578160200160208202803683370190505b5061440c565b87516143ca9060016154c9565b6001600160401b038111156143e1576143e1614b55565b60405190808252806020026020018201604052801561440a578160200160208202803683370190505b505b90506040815111156144605760405162461bcd60e51b815260206004820152601c60248201527f417070656e642063726561746573206f76657273697a652074726565000000006044820152606401610c76565b60005b885181101561460157878110156144ef57888181518110614486576144866152d1565b60200260200101516000801b146144ea5760405162461bcd60e51b815260206004820152602260248201527f417070656e642061626f7665206c65617374207369676e69666963616e7420626044820152611a5d60f21b6064820152608401610c76565b6145ef565b60008590036145355788818151811061450a5761450a6152d1565b6020026020010151828281518110614524576145246152d1565b6020026020010181815250506145ef565b888181518110614547576145476152d1565b60200260200101516000801b0361457f578482828151811061456b5761456b6152d1565b6020908102919091010152600094506145ef565b6000801b828281518110614595576145956152d1565b6020026020010181815250508881815181106145b3576145b36152d1565b6020026020010151856040516020016145d6929190918252602082015260400190565b6040516020818303038152906040528051906020012094505b806145f981615499565b915050614463565b5083156146355783816001835161461891906152be565b81518110614628576146286152d1565b6020026020010181815250505b806001825161464491906152be565b81518110614654576146546152d1565b60200260200101516000801b036105755760405162461bcd60e51b815260206004820152600f60248201526e4c61737420656e747279207a65726f60881b6044820152606401610c76565b60008590036146c15760405163235e76ef60e21b815260040160405180910390fd5b8281116146eb576040516308183ebd60e21b81526004810184905260248101829052604401610c76565b600084900361470d576040516320f1a0f960e21b815260040160405180910390fd5b600082900361356957604051635cb6e5bb60e01b815260040160405180910390fd5b6060824710156147905760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610c76565b6001600160a01b0385163b6147e75760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610c76565b600080866001600160a01b031685876040516148039190615840565b60006040518083038185875af1925050503d8060008114614840576040519150601f19603f3d011682016040523d82523d6000602084013e614845565b606091505b50915091506105758282866148c2565b60008082116148a65760405162461bcd60e51b815260206004820152601c60248201527f5a65726f20686173206e6f207369676e69666963616e742062697473000000006044820152606401610c76565b600082806148b56001826152be565b16189050611d6e81613d99565b606083156148d1575081611d6e565b8251156148e15782518084602001fd5b8160405162461bcd60e51b8152600401610c76919061585c565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915290565b6040805160c08101825260008082526020820181905291810182905260608101919091526080810161496f6149f4565b815260200161497c6149f4565b905290565b604080516101c081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905290610160820190815260006020820181905260409091015290565b6040518060400160405280614a07614a13565b81526020016000905290565b6040518060400160405280614a26614a2f565b815260200161497c5b60405180604001604052806002906020820280368337509192915050565b803560ff81168114610c7f57600080fd5b60008060008060008060c08789031215614a7757600080fd5b614a8087614a4d565b9860208801359850604088013597606081013597506080810135965060a00135945092505050565b600060208284031215614aba57600080fd5b81356001600160401b03811115614ad057600080fd5b820160c08185031215611d6e57600080fd5b60008060408385031215614af557600080fd5b50508035926020909101359150565b600060208284031215614b1657600080fd5b5035919050565b60038110614b2a57600080fd5b50565b8035610c7f81614b1d565b600060208284031215614b4a57600080fd5b8135611d6e81614b1d565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715614b8d57614b8d614b55565b60405290565b604080519081016001600160401b0381118282101715614b8d57614b8d614b55565b604051601f8201601f191681016001600160401b0381118282101715614bdd57614bdd614b55565b604052919050565b60006001600160401b03821115614bfe57614bfe614b55565b5060051b60200190565b600082601f830112614c1957600080fd5b81356020614c2e614c2983614be5565b614bb5565b82815260059290921b84018101918181019086841115614c4d57600080fd5b8286015b84811015614c685780358352918301918301614c51565b509695505050505050565b6000806000838503610120811215614c8a57600080fd5b8435935060208501356001600160401b03811115614ca757600080fd5b614cb387828801614c08565b93505060e0603f1982011215614cc857600080fd5b506040840190509250925092565b60008083601f840112614ce857600080fd5b5081356001600160401b03811115614cff57600080fd5b6020830191508360208260051b8501011115614d1a57600080fd5b9250929050565b6000806000806000806000878903610120811215614d3e57600080fd5b8835975060208901356001600160401b0380821115614d5c57600080fd5b908a01906040828d031215614d7057600080fd5b81985060a0603f1984011215614d8557600080fd5b60408b01975060e08b0135925080831115614d9f57600080fd5b614dab8c848d01614cd6565b90975095506101008b0135925086915080831115614dc857600080fd5b5050614dd68a828b01614cd6565b989b979a50959850939692959293505050565b600080600080600060a08688031215614e0157600080fd5b614e0a86614a4d565b97602087013597506040870135966060810135965060800135945092505050565b60008060008060608587031215614e4157600080fd5b843593506020850135925060408501356001600160401b0380821115614e6657600080fd5b818701915087601f830112614e7a57600080fd5b813581811115614e8957600080fd5b886020828501011115614e9b57600080fd5b95989497505060200194505050565b634e487b7160e01b600052602160045260246000fd5b60028110614ed057614ed0614eaa565b9052565b60006101c082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015260e083015160e083015261010080840151614f42828501826001600160a01b03169052565b5050610120838101516001600160401b038116848301525050610140838101516001600160401b03811684830152505061016080840151614f8582850182614ec0565b50506101808381015160ff16908301526101a08084015180151582850152611d38565b600060208284031215614fba57600080fd5b611d6e82614a4d565b6000808335601e19843603018112614fda57600080fd5b8301803591506001600160401b03821115614ff457600080fd5b602001915036819003821315614d1a57600080fd5b6001600160401b0381168114614b2a57600080fd5b600082601f83011261502f57600080fd5b615037614b93565b80604084018581111561504957600080fd5b845b81811015612eeb57803561505e81615009565b84526020938401930161504b565b600081830360e081121561507f57600080fd5b615087614b6b565b915060a081121561509757600080fd5b61509f614b93565b60808212156150ad57600080fd5b6150b5614b93565b915084601f8501126150c657600080fd5b6150ce614b93565b8060408601878111156150e057600080fd5b865b818110156150fa5780358452602093840193016150e2565b50818552615108888261501e565b602086015250505081815261511f60808501614b2d565b6020820152808352505060a0820135602082015260c0820135604082015292915050565b60008060006101e0848603121561515957600080fd5b83356001600160401b0381111561516f57600080fd5b61517b86828701614c08565b93505061518b856020860161506c565b915061519b85610100860161506c565b90509250925092565b60038110614b2a57614b2a614eaa565b614ed0816151a4565b805180518360005b60028110156151e45782518252602092830192909101906001016151c5565b505050602090810151906040840160005b600281101561521b5783516001600160401b0316825292820192908201906001016151f5565b505082015190506121da60808401826151b4565b848152610100810161524460208301866151bd565b60c082019390935260e0015292915050565b60006020828403121561526857600080fd5b81518015158114611d6e57600080fd5b60006020828403121561528a57600080fd5b8151611d6e81615009565b602081016152a2836151a4565b91905290565b634e487b7160e01b600052601160045260246000fd5b81810381811115610b9457610b946152a8565b634e487b7160e01b600052603260045260246000fd5b8481526101008101602060408682850137606083016040870160005b600281101561533257813561531781615009565b6001600160401b031683529183019190830190600101615303565b50505050608085013561534481614b1d565b61534d816151a4565b60a083015260c082019390935260e0015292915050565b6001600160401b03828116828216039080821115613bdf57613bdf6152a8565b6001600160a01b0381168114614b2a57600080fd5b600060c082019050838252823560208301526020830135604083015260408301356153c381615384565b6001600160a01b03166060838101919091528301356153e181615009565b6001600160401b0380821660808501526080850135915061540182615009565b80821660a085015250509392505050565b60006020828403121561542457600080fd5b8135611d6e81615009565b60006020828403121561544157600080fd5b8151611d6e81615384565b60ff8181168382160190811115610b9457610b946152a8565b82815260408101611d6e6020830184614ec0565b6001600160401b03818116838216019080821115613bdf57613bdf6152a8565b6000600182016154ab576154ab6152a8565b5060010190565b8082028115828204841417610b9457610b946152a8565b80820180821115610b9457610b946152a8565b8551815260018060a01b0360208701511660208201526040860151604082015284606082015283608082015260c060a08201528160c0820152818360e0830137600081830160e090810191909152601f909201601f19160101949350505050565b60006020828403121561554f57600080fd5b5051919050565b600082601f83011261556757600080fd5b81516020615577614c2983614be5565b82815260059290921b8401810191818101908684111561559657600080fd5b8286015b84811015614c68578051835291830191830161559a565b600080604083850312156155c457600080fd5b82516001600160401b03808211156155db57600080fd5b6155e786838701615556565b935060208501519150808211156155fd57600080fd5b5061560a85828601615556565b9150509250929050565b60a08101610b9482846151bd565b600080600080600060a0868803121561563a57600080fd5b853594506020860135935060408601356001600160401b038082111561565f57600080fd5b61566b89838a01614c08565b9450606088013591508082111561568157600080fd5b61568d89838a01614c08565b935060808801359150808211156156a357600080fd5b506156b088828901614c08565b9150509295509295909350565b600080604083850312156156d057600080fd5b82356001600160401b03808211156156e757600080fd5b6156f386838701614c08565b9350602085013591508082111561570957600080fd5b5061560a85828601614c08565b634e487b7160e01b600052600160045260246000fd5b600181815b8085111561576757816000190482111561574d5761574d6152a8565b8085161561575a57918102915b93841c9390800290615731565b509250929050565b60008261577e57506001610b94565b8161578b57506000610b94565b81600181146157a157600281146157ab576157c7565b6001915050610b94565b60ff8411156157bc576157bc6152a8565b50506001821b610b94565b5060208310610133831016604e8410600b84101617156157ea575081810a610b94565b6157f4838361572c565b8060001904821115615808576158086152a8565b029392505050565b6000611d6e838361576f565b60005b8381101561583757818101518382015260200161581f565b50506000910152565b6000825161585281846020870161581c565b9190910192915050565b602081526000825180602084015261587b81604085016020870161581c565b601f01601f1916919091016040019291505056fea264697066735822122001730c7c06f3f72068613e12e8a6cf6c3da56311c36ab4ea2d7ae4d5aa4baa9a64736f6c63430008110033a2646970667358221220235edc3d7122a9c6b0f057eefd96b3ce45137f85d4212b8521eb126b4c1aa82c64736f6c63430008110033",
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
