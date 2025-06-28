// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocksgen

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

// BufferConfig is an auto generated low-level Go binding around an user-defined struct.
type BufferConfig struct {
	Threshold            uint64
	Max                  uint64
	ReplenishRateInBasis uint64
}

// DelayProof is an auto generated low-level Go binding around an user-defined struct.
type DelayProof struct {
	BeforeDelayedAcc [32]byte
	DelayedMessage   MessagesMessage
}

// ExecutionContext is an auto generated low-level Go binding around an user-defined struct.
type ExecutionContext struct {
	MaxInboxMessagesRead  *big.Int
	Bridge                common.Address
	InitialWasmModuleRoot [32]byte
}

// ExecutionState is an auto generated low-level Go binding around an user-defined struct.
type ExecutionState struct {
	GlobalState   GlobalState
	MachineStatus uint8
}

// GlobalState is an auto generated low-level Go binding around an user-defined struct.
type GlobalState struct {
	Bytes32Vals [2][32]byte
	U64Vals     [2]uint64
}

// IBridgeTimeBounds is an auto generated low-level Go binding around an user-defined struct.
type IBridgeTimeBounds struct {
	MinTimestamp   uint64
	MaxTimestamp   uint64
	MinBlockNumber uint64
	MaxBlockNumber uint64
}

// ISequencerInboxMaxTimeVariation is an auto generated low-level Go binding around an user-defined struct.
type ISequencerInboxMaxTimeVariation struct {
	DelayBlocks   *big.Int
	FutureBlocks  *big.Int
	DelaySeconds  *big.Int
	FutureSeconds *big.Int
}

// MessagesMessage is an auto generated low-level Go binding around an user-defined struct.
type MessagesMessage struct {
	Kind            uint8
	Sender          common.Address
	BlockNumber     uint64
	Timestamp       uint64
	InboxSeqNum     *big.Int
	BaseFeeL1       *big.Int
	MessageDataHash [32]byte
}

// BridgeStubMetaData contains all meta data concerning the BridgeStub contract.
var BridgeStubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"BadSequencerMessageNumber\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BridgeCallTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"InboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"OutboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"RollupUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSequencerInbox\",\"type\":\"address\"}],\"name\":\"SequencerInboxUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptFundsFromOldBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeOutbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedDelayedInboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"allowedDelayedInboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedOutboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedOutboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delayedInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"enqueueDelayedMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"enqueueSequencerMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seqMessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"acc\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerReportedSubMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setDelayedInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencerInbox\",\"type\":\"address\"}],\"name\":\"setSequencerInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"batchPoster\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"submitBatchSpendingReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"updateRollupAddress\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610f1e806100206000396000f3fe60806040526004361061018a5760003560e01c8063ab5d8943116100d6578063d5719dc21161007f578063e77145f411610059578063e77145f414610223578063eca067ad14610487578063ee35f3271461049c57600080fd5b8063d5719dc214610427578063e1758bd814610447578063e76f5c8d1461046757600080fd5b8063c4d66de8116100b0578063c4d66de8146102d1578063cb23bcb5146103f7578063cee3d7281461040c57600080fd5b8063ab5d894314610357578063ad48cb5e14610377578063ae60bd13146103bb57600080fd5b80637a88b10711610138578063919cc70611610112578063919cc706146102d1578063945e1147146102f15780639e5d4c491461032957600080fd5b80637a88b1071461025b57806386598a561461027e5780638db5993b146102be57600080fd5b806347fb24c51161016957806347fb24c5146102035780634f61f850146102255780635fca4a161461024557600080fd5b806284120c1461018f57806316bf5579146101b3578063413b35bd146101d3575b600080fd5b34801561019b57600080fd5b506005545b6040519081526020015b60405180910390f35b3480156101bf57600080fd5b506101a06101ce366004610c28565b6104bc565b3480156101df57600080fd5b506101f36101ee366004610c59565b6104dd565b60405190151581526020016101aa565b34801561020f57600080fd5b5061022361021e366004610c7d565b61052d565b005b34801561023157600080fd5b50610223610240366004610c59565b61077b565b34801561025157600080fd5b506101a060075481565b34801561026757600080fd5b506101a0610276366004610cbb565b600092915050565b34801561028a57600080fd5b5061029e610299366004610ce7565b6107dc565b6040805194855260208501939093529183015260608201526080016101aa565b6101a06102cc366004610d19565b61092c565b3480156102dd57600080fd5b506102236102ec366004610c59565b6109a4565b3480156102fd57600080fd5b5061031161030c366004610c28565b6109ec565b6040516001600160a01b0390911681526020016101aa565b34801561033557600080fd5b50610349610344366004610d60565b610a16565b6040516101aa929190610e0d565b34801561036357600080fd5b50600354610311906001600160a01b031681565b34801561038357600080fd5b506008546103a99074010000000000000000000000000000000000000000900460ff1681565b60405160ff90911681526020016101aa565b3480156103c757600080fd5b506101f36103d6366004610c59565b6001600160a01b031660009081526020819052604090206001015460ff1690565b34801561040357600080fd5b506103116104dd565b34801561041857600080fd5b506102236102ec366004610c7d565b34801561043357600080fd5b506101a0610442366004610c28565b610ab2565b34801561045357600080fd5b50600854610311906001600160a01b031681565b34801561047357600080fd5b50610311610482366004610c28565b610ac2565b34801561049357600080fd5b506004546101a0565b3480156104a857600080fd5b50600654610311906001600160a01b031681565b600581815481106104cc57600080fd5b600091825260209091200154905081565b60405162461bcd60e51b815260206004820152600f60248201527f4e4f545f494d504c454d454e544544000000000000000000000000000000000060448201526000906064015b60405180910390fd5b6001600160a01b03821660008181526020818152604091829020600181015492518515158152909360ff90931692917f6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521910160405180910390a2821515811515036105985750505050565b82156106315760408051808201825260018054825260208083018281526001600160a01b0389166000818152928390529482209351845551928201805460ff1916931515939093179092558054808201825591527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601805473ffffffffffffffffffffffffffffffffffffffff19169091179055610775565b60018054610640908290610e49565b8154811061065057610650610e6a565b6000918252602090912001548254600180546001600160a01b0390931692909190811061067f5761067f610e6a565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550816000015460008060018560000154815481106106cc576106cc610e6a565b60009182526020808320909101546001600160a01b03168352820192909252604001902055600180548061070257610702610e80565b6000828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b038616825281905260408120908155600101805460ff191690555b50505050565b6006805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a9060200160405180910390a150565b60008060008085600754141580156107f357508515155b8015610800575060075415155b15610845576007546040517fe2051feb000000000000000000000000000000000000000000000000000000008152600481019190915260248101879052604401610524565b600785905560055493508315610883576005805461086590600190610e49565b8154811061087557610875610e6a565b906000526020600020015492505b86156108b4576004610896600189610e49565b815481106108a6576108a6610e6a565b906000526020600020015491505b60408051602081018590529081018990526060810183905260800160408051601f198184030181529190528051602090910120600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0018190559398929750909550919350915050565b3360009081526020819052604081206001015460ff1661098e5760405162461bcd60e51b815260206004820152600e60248201527f4e4f545f46524f4d5f494e424f580000000000000000000000000000000000006044820152606401610524565b61099c848443424887610ad2565b949350505050565b60405162461bcd60e51b815260206004820152600f60248201527f4e4f545f494d504c454d454e54454400000000000000000000000000000000006044820152606401610524565b600281815481106109fc57600080fd5b6000918252602090912001546001600160a01b0316905081565b60006060610a5b868686868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610b8d92505050565b60405191935091506001600160a01b0387169033907f2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d46690610aa190899089908990610e96565b60405180910390a394509492505050565b600481815481106104cc57600080fd5b600181815481106109fc57600080fd5b60045460408051600060208083018290526021830182905260358301829052603d8301829052604583018290526065830182905260858084018790528451808503909101815260a59093019093528151919092012090919060008215610b5d576004610b3f600185610e49565b81548110610b4f57610b4f610e6a565b906000526020600020015490505b6004610b698284610bf9565b81546001810183556000928352602090922090910155509098975050505050505050565b60006060846001600160a01b03168484604051610baa9190610ecc565b60006040518083038185875af1925050503d8060008114610be7576040519150601f19603f3d011682016040523d82523d6000602084013e610bec565b606091505b5090969095509350505050565b604080516020808201859052818301849052825180830384018152606090920190925280519101205b92915050565b600060208284031215610c3a57600080fd5b5035919050565b6001600160a01b0381168114610c5657600080fd5b50565b600060208284031215610c6b57600080fd5b8135610c7681610c41565b9392505050565b60008060408385031215610c9057600080fd5b8235610c9b81610c41565b915060208301358015158114610cb057600080fd5b809150509250929050565b60008060408385031215610cce57600080fd5b8235610cd981610c41565b946020939093013593505050565b60008060008060808587031215610cfd57600080fd5b5050823594602084013594506040840135936060013592509050565b600080600060608486031215610d2e57600080fd5b833560ff81168114610d3f57600080fd5b92506020840135610d4f81610c41565b929592945050506040919091013590565b60008060008060608587031215610d7657600080fd5b8435610d8181610c41565b935060208501359250604085013567ffffffffffffffff80821115610da557600080fd5b818701915087601f830112610db957600080fd5b813581811115610dc857600080fd5b886020828501011115610dda57600080fd5b95989497505060200194505050565b60005b83811015610e04578181015183820152602001610dec565b50506000910152565b82151581526040602082015260008251806040840152610e34816060850160208701610de9565b601f01601f1916919091016060019392505050565b81810381811115610c2257634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b83815260406020820152816040820152818360608301376000818301606090810191909152601f909201601f1916010192915050565b60008251610ede818460208701610de9565b919091019291505056fea264697066735822122088bf79f4990f518aaa542f439b1d9471f0bb42050304a87194a98a98e474b62c64736f6c63430008110033",
}

// BridgeStubABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeStubMetaData.ABI instead.
var BridgeStubABI = BridgeStubMetaData.ABI

// BridgeStubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeStubMetaData.Bin instead.
var BridgeStubBin = BridgeStubMetaData.Bin

// DeployBridgeStub deploys a new Ethereum contract, binding an instance of BridgeStub to it.
func DeployBridgeStub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeStub, error) {
	parsed, err := BridgeStubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeStubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeStub{BridgeStubCaller: BridgeStubCaller{contract: contract}, BridgeStubTransactor: BridgeStubTransactor{contract: contract}, BridgeStubFilterer: BridgeStubFilterer{contract: contract}}, nil
}

// BridgeStub is an auto generated Go binding around an Ethereum contract.
type BridgeStub struct {
	BridgeStubCaller     // Read-only binding to the contract
	BridgeStubTransactor // Write-only binding to the contract
	BridgeStubFilterer   // Log filterer for contract events
}

// BridgeStubCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeStubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeStubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeStubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeStubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeStubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeStubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeStubSession struct {
	Contract     *BridgeStub       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeStubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeStubCallerSession struct {
	Contract *BridgeStubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BridgeStubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeStubTransactorSession struct {
	Contract     *BridgeStubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BridgeStubRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeStubRaw struct {
	Contract *BridgeStub // Generic contract binding to access the raw methods on
}

// BridgeStubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeStubCallerRaw struct {
	Contract *BridgeStubCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeStubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeStubTransactorRaw struct {
	Contract *BridgeStubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeStub creates a new instance of BridgeStub, bound to a specific deployed contract.
func NewBridgeStub(address common.Address, backend bind.ContractBackend) (*BridgeStub, error) {
	contract, err := bindBridgeStub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeStub{BridgeStubCaller: BridgeStubCaller{contract: contract}, BridgeStubTransactor: BridgeStubTransactor{contract: contract}, BridgeStubFilterer: BridgeStubFilterer{contract: contract}}, nil
}

// NewBridgeStubCaller creates a new read-only instance of BridgeStub, bound to a specific deployed contract.
func NewBridgeStubCaller(address common.Address, caller bind.ContractCaller) (*BridgeStubCaller, error) {
	contract, err := bindBridgeStub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeStubCaller{contract: contract}, nil
}

// NewBridgeStubTransactor creates a new write-only instance of BridgeStub, bound to a specific deployed contract.
func NewBridgeStubTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeStubTransactor, error) {
	contract, err := bindBridgeStub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeStubTransactor{contract: contract}, nil
}

// NewBridgeStubFilterer creates a new log filterer instance of BridgeStub, bound to a specific deployed contract.
func NewBridgeStubFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeStubFilterer, error) {
	contract, err := bindBridgeStub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeStubFilterer{contract: contract}, nil
}

// bindBridgeStub binds a generic wrapper to an already deployed contract.
func bindBridgeStub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeStubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeStub *BridgeStubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeStub.Contract.BridgeStubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeStub *BridgeStubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeStub.Contract.BridgeStubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeStub *BridgeStubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeStub.Contract.BridgeStubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeStub *BridgeStubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeStub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeStub *BridgeStubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeStub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeStub *BridgeStubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeStub.Contract.contract.Transact(opts, method, params...)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeStub *BridgeStubCaller) ActiveOutbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "activeOutbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeStub *BridgeStubSession) ActiveOutbox() (common.Address, error) {
	return _BridgeStub.Contract.ActiveOutbox(&_BridgeStub.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeStub *BridgeStubCallerSession) ActiveOutbox() (common.Address, error) {
	return _BridgeStub.Contract.ActiveOutbox(&_BridgeStub.CallOpts)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubCaller) AllowedDelayedInboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "allowedDelayedInboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeStub.Contract.AllowedDelayedInboxList(&_BridgeStub.CallOpts, arg0)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubCallerSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeStub.Contract.AllowedDelayedInboxList(&_BridgeStub.CallOpts, arg0)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeStub *BridgeStubCaller) AllowedDelayedInboxes(opts *bind.CallOpts, inbox common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "allowedDelayedInboxes", inbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeStub *BridgeStubSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeStub.Contract.AllowedDelayedInboxes(&_BridgeStub.CallOpts, inbox)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeStub *BridgeStubCallerSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeStub.Contract.AllowedDelayedInboxes(&_BridgeStub.CallOpts, inbox)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubCaller) AllowedOutboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "allowedOutboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeStub.Contract.AllowedOutboxList(&_BridgeStub.CallOpts, arg0)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubCallerSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeStub.Contract.AllowedOutboxList(&_BridgeStub.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address ) pure returns(bool)
func (_BridgeStub *BridgeStubCaller) AllowedOutboxes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "allowedOutboxes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address ) pure returns(bool)
func (_BridgeStub *BridgeStubSession) AllowedOutboxes(arg0 common.Address) (bool, error) {
	return _BridgeStub.Contract.AllowedOutboxes(&_BridgeStub.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address ) pure returns(bool)
func (_BridgeStub *BridgeStubCallerSession) AllowedOutboxes(arg0 common.Address) (bool, error) {
	return _BridgeStub.Contract.AllowedOutboxes(&_BridgeStub.CallOpts, arg0)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubCaller) DelayedInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "delayedInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeStub.Contract.DelayedInboxAccs(&_BridgeStub.CallOpts, arg0)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubCallerSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeStub.Contract.DelayedInboxAccs(&_BridgeStub.CallOpts, arg0)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCaller) DelayedMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "delayedMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.DelayedMessageCount(&_BridgeStub.CallOpts)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCallerSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.DelayedMessageCount(&_BridgeStub.CallOpts)
}

// Initialize is a free data retrieval call binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ) pure returns()
func (_BridgeStub *BridgeStubCaller) Initialize(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "initialize", arg0)

	if err != nil {
		return err
	}

	return err

}

// Initialize is a free data retrieval call binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ) pure returns()
func (_BridgeStub *BridgeStubSession) Initialize(arg0 common.Address) error {
	return _BridgeStub.Contract.Initialize(&_BridgeStub.CallOpts, arg0)
}

// Initialize is a free data retrieval call binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ) pure returns()
func (_BridgeStub *BridgeStubCallerSession) Initialize(arg0 common.Address) error {
	return _BridgeStub.Contract.Initialize(&_BridgeStub.CallOpts, arg0)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_BridgeStub *BridgeStubCaller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "nativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_BridgeStub *BridgeStubSession) NativeToken() (common.Address, error) {
	return _BridgeStub.Contract.NativeToken(&_BridgeStub.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_BridgeStub *BridgeStubCallerSession) NativeToken() (common.Address, error) {
	return _BridgeStub.Contract.NativeToken(&_BridgeStub.CallOpts)
}

// NativeTokenDecimals is a free data retrieval call binding the contract method 0xad48cb5e.
//
// Solidity: function nativeTokenDecimals() view returns(uint8)
func (_BridgeStub *BridgeStubCaller) NativeTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "nativeTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NativeTokenDecimals is a free data retrieval call binding the contract method 0xad48cb5e.
//
// Solidity: function nativeTokenDecimals() view returns(uint8)
func (_BridgeStub *BridgeStubSession) NativeTokenDecimals() (uint8, error) {
	return _BridgeStub.Contract.NativeTokenDecimals(&_BridgeStub.CallOpts)
}

// NativeTokenDecimals is a free data retrieval call binding the contract method 0xad48cb5e.
//
// Solidity: function nativeTokenDecimals() view returns(uint8)
func (_BridgeStub *BridgeStubCallerSession) NativeTokenDecimals() (uint8, error) {
	return _BridgeStub.Contract.NativeTokenDecimals(&_BridgeStub.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() pure returns(address)
func (_BridgeStub *BridgeStubCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() pure returns(address)
func (_BridgeStub *BridgeStubSession) Rollup() (common.Address, error) {
	return _BridgeStub.Contract.Rollup(&_BridgeStub.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() pure returns(address)
func (_BridgeStub *BridgeStubCallerSession) Rollup() (common.Address, error) {
	return _BridgeStub.Contract.Rollup(&_BridgeStub.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeStub *BridgeStubCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeStub *BridgeStubSession) SequencerInbox() (common.Address, error) {
	return _BridgeStub.Contract.SequencerInbox(&_BridgeStub.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeStub *BridgeStubCallerSession) SequencerInbox() (common.Address, error) {
	return _BridgeStub.Contract.SequencerInbox(&_BridgeStub.CallOpts)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubCaller) SequencerInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "sequencerInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeStub.Contract.SequencerInboxAccs(&_BridgeStub.CallOpts, arg0)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubCallerSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeStub.Contract.SequencerInboxAccs(&_BridgeStub.CallOpts, arg0)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCaller) SequencerMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "sequencerMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.SequencerMessageCount(&_BridgeStub.CallOpts)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCallerSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.SequencerMessageCount(&_BridgeStub.CallOpts)
}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCaller) SequencerReportedSubMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "sequencerReportedSubMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubSession) SequencerReportedSubMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.SequencerReportedSubMessageCount(&_BridgeStub.CallOpts)
}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCallerSession) SequencerReportedSubMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.SequencerReportedSubMessageCount(&_BridgeStub.CallOpts)
}

// SetOutbox is a free data retrieval call binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address , bool ) pure returns()
func (_BridgeStub *BridgeStubCaller) SetOutbox(opts *bind.CallOpts, arg0 common.Address, arg1 bool) error {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "setOutbox", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// SetOutbox is a free data retrieval call binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address , bool ) pure returns()
func (_BridgeStub *BridgeStubSession) SetOutbox(arg0 common.Address, arg1 bool) error {
	return _BridgeStub.Contract.SetOutbox(&_BridgeStub.CallOpts, arg0, arg1)
}

// SetOutbox is a free data retrieval call binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address , bool ) pure returns()
func (_BridgeStub *BridgeStubCallerSession) SetOutbox(arg0 common.Address, arg1 bool) error {
	return _BridgeStub.Contract.SetOutbox(&_BridgeStub.CallOpts, arg0, arg1)
}

// UpdateRollupAddress is a free data retrieval call binding the contract method 0x919cc706.
//
// Solidity: function updateRollupAddress(address ) pure returns()
func (_BridgeStub *BridgeStubCaller) UpdateRollupAddress(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "updateRollupAddress", arg0)

	if err != nil {
		return err
	}

	return err

}

// UpdateRollupAddress is a free data retrieval call binding the contract method 0x919cc706.
//
// Solidity: function updateRollupAddress(address ) pure returns()
func (_BridgeStub *BridgeStubSession) UpdateRollupAddress(arg0 common.Address) error {
	return _BridgeStub.Contract.UpdateRollupAddress(&_BridgeStub.CallOpts, arg0)
}

// UpdateRollupAddress is a free data retrieval call binding the contract method 0x919cc706.
//
// Solidity: function updateRollupAddress(address ) pure returns()
func (_BridgeStub *BridgeStubCallerSession) UpdateRollupAddress(arg0 common.Address) error {
	return _BridgeStub.Contract.UpdateRollupAddress(&_BridgeStub.CallOpts, arg0)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeStub *BridgeStubTransactor) AcceptFundsFromOldBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "acceptFundsFromOldBridge")
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeStub *BridgeStubSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeStub.Contract.AcceptFundsFromOldBridge(&_BridgeStub.TransactOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeStub *BridgeStubTransactorSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeStub.Contract.AcceptFundsFromOldBridge(&_BridgeStub.TransactOpts)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeStub *BridgeStubTransactor) EnqueueDelayedMessage(opts *bind.TransactOpts, kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "enqueueDelayedMessage", kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeStub *BridgeStubSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.Contract.EnqueueDelayedMessage(&_BridgeStub.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeStub *BridgeStubTransactorSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.Contract.EnqueueDelayedMessage(&_BridgeStub.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeStub *BridgeStubTransactor) EnqueueSequencerMessage(opts *bind.TransactOpts, dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "enqueueSequencerMessage", dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeStub *BridgeStubSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeStub.Contract.EnqueueSequencerMessage(&_BridgeStub.TransactOpts, dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeStub *BridgeStubTransactorSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeStub.Contract.EnqueueSequencerMessage(&_BridgeStub.TransactOpts, dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeStub *BridgeStubTransactor) ExecuteCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "executeCall", to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeStub *BridgeStubSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeStub.Contract.ExecuteCall(&_BridgeStub.TransactOpts, to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeStub *BridgeStubTransactorSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeStub.Contract.ExecuteCall(&_BridgeStub.TransactOpts, to, value, data)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeStub *BridgeStubTransactor) SetDelayedInbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "setDelayedInbox", inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeStub *BridgeStubSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeStub.Contract.SetDelayedInbox(&_BridgeStub.TransactOpts, inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeStub *BridgeStubTransactorSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeStub.Contract.SetDelayedInbox(&_BridgeStub.TransactOpts, inbox, enabled)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeStub *BridgeStubTransactor) SetSequencerInbox(opts *bind.TransactOpts, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "setSequencerInbox", _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeStub *BridgeStubSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeStub.Contract.SetSequencerInbox(&_BridgeStub.TransactOpts, _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeStub *BridgeStubTransactorSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeStub.Contract.SetSequencerInbox(&_BridgeStub.TransactOpts, _sequencerInbox)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeStub *BridgeStubTransactor) SubmitBatchSpendingReport(opts *bind.TransactOpts, batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "submitBatchSpendingReport", batchPoster, dataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeStub *BridgeStubSession) SubmitBatchSpendingReport(batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.Contract.SubmitBatchSpendingReport(&_BridgeStub.TransactOpts, batchPoster, dataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeStub *BridgeStubTransactorSession) SubmitBatchSpendingReport(batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.Contract.SubmitBatchSpendingReport(&_BridgeStub.TransactOpts, batchPoster, dataHash)
}

// BridgeStubBridgeCallTriggeredIterator is returned from FilterBridgeCallTriggered and is used to iterate over the raw logs and unpacked data for BridgeCallTriggered events raised by the BridgeStub contract.
type BridgeStubBridgeCallTriggeredIterator struct {
	Event *BridgeStubBridgeCallTriggered // Event containing the contract specifics and raw log

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
func (it *BridgeStubBridgeCallTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubBridgeCallTriggered)
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
		it.Event = new(BridgeStubBridgeCallTriggered)
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
func (it *BridgeStubBridgeCallTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubBridgeCallTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubBridgeCallTriggered represents a BridgeCallTriggered event raised by the BridgeStub contract.
type BridgeStubBridgeCallTriggered struct {
	Outbox common.Address
	To     common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeCallTriggered is a free log retrieval operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeStub *BridgeStubFilterer) FilterBridgeCallTriggered(opts *bind.FilterOpts, outbox []common.Address, to []common.Address) (*BridgeStubBridgeCallTriggeredIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeStubBridgeCallTriggeredIterator{contract: _BridgeStub.contract, event: "BridgeCallTriggered", logs: logs, sub: sub}, nil
}

// WatchBridgeCallTriggered is a free log subscription operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeStub *BridgeStubFilterer) WatchBridgeCallTriggered(opts *bind.WatchOpts, sink chan<- *BridgeStubBridgeCallTriggered, outbox []common.Address, to []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubBridgeCallTriggered)
				if err := _BridgeStub.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
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

// ParseBridgeCallTriggered is a log parse operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeStub *BridgeStubFilterer) ParseBridgeCallTriggered(log types.Log) (*BridgeStubBridgeCallTriggered, error) {
	event := new(BridgeStubBridgeCallTriggered)
	if err := _BridgeStub.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeStubInboxToggleIterator is returned from FilterInboxToggle and is used to iterate over the raw logs and unpacked data for InboxToggle events raised by the BridgeStub contract.
type BridgeStubInboxToggleIterator struct {
	Event *BridgeStubInboxToggle // Event containing the contract specifics and raw log

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
func (it *BridgeStubInboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubInboxToggle)
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
		it.Event = new(BridgeStubInboxToggle)
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
func (it *BridgeStubInboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubInboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubInboxToggle represents a InboxToggle event raised by the BridgeStub contract.
type BridgeStubInboxToggle struct {
	Inbox   common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInboxToggle is a free log retrieval operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) FilterInboxToggle(opts *bind.FilterOpts, inbox []common.Address) (*BridgeStubInboxToggleIterator, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeStubInboxToggleIterator{contract: _BridgeStub.contract, event: "InboxToggle", logs: logs, sub: sub}, nil
}

// WatchInboxToggle is a free log subscription operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) WatchInboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeStubInboxToggle, inbox []common.Address) (event.Subscription, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubInboxToggle)
				if err := _BridgeStub.contract.UnpackLog(event, "InboxToggle", log); err != nil {
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

// ParseInboxToggle is a log parse operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) ParseInboxToggle(log types.Log) (*BridgeStubInboxToggle, error) {
	event := new(BridgeStubInboxToggle)
	if err := _BridgeStub.contract.UnpackLog(event, "InboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeStubMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the BridgeStub contract.
type BridgeStubMessageDeliveredIterator struct {
	Event *BridgeStubMessageDelivered // Event containing the contract specifics and raw log

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
func (it *BridgeStubMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubMessageDelivered)
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
		it.Event = new(BridgeStubMessageDelivered)
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
func (it *BridgeStubMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubMessageDelivered represents a MessageDelivered event raised by the BridgeStub contract.
type BridgeStubMessageDelivered struct {
	MessageIndex    *big.Int
	BeforeInboxAcc  [32]byte
	Inbox           common.Address
	Kind            uint8
	Sender          common.Address
	MessageDataHash [32]byte
	BaseFeeL1       *big.Int
	Timestamp       uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeStub *BridgeStubFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (*BridgeStubMessageDeliveredIterator, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &BridgeStubMessageDeliveredIterator{contract: _BridgeStub.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeStub *BridgeStubFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *BridgeStubMessageDelivered, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubMessageDelivered)
				if err := _BridgeStub.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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

// ParseMessageDelivered is a log parse operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeStub *BridgeStubFilterer) ParseMessageDelivered(log types.Log) (*BridgeStubMessageDelivered, error) {
	event := new(BridgeStubMessageDelivered)
	if err := _BridgeStub.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeStubOutboxToggleIterator is returned from FilterOutboxToggle and is used to iterate over the raw logs and unpacked data for OutboxToggle events raised by the BridgeStub contract.
type BridgeStubOutboxToggleIterator struct {
	Event *BridgeStubOutboxToggle // Event containing the contract specifics and raw log

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
func (it *BridgeStubOutboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubOutboxToggle)
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
		it.Event = new(BridgeStubOutboxToggle)
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
func (it *BridgeStubOutboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubOutboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubOutboxToggle represents a OutboxToggle event raised by the BridgeStub contract.
type BridgeStubOutboxToggle struct {
	Outbox  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOutboxToggle is a free log retrieval operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) FilterOutboxToggle(opts *bind.FilterOpts, outbox []common.Address) (*BridgeStubOutboxToggleIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeStubOutboxToggleIterator{contract: _BridgeStub.contract, event: "OutboxToggle", logs: logs, sub: sub}, nil
}

// WatchOutboxToggle is a free log subscription operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) WatchOutboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeStubOutboxToggle, outbox []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubOutboxToggle)
				if err := _BridgeStub.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
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

// ParseOutboxToggle is a log parse operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) ParseOutboxToggle(log types.Log) (*BridgeStubOutboxToggle, error) {
	event := new(BridgeStubOutboxToggle)
	if err := _BridgeStub.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeStubRollupUpdatedIterator is returned from FilterRollupUpdated and is used to iterate over the raw logs and unpacked data for RollupUpdated events raised by the BridgeStub contract.
type BridgeStubRollupUpdatedIterator struct {
	Event *BridgeStubRollupUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeStubRollupUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubRollupUpdated)
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
		it.Event = new(BridgeStubRollupUpdated)
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
func (it *BridgeStubRollupUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubRollupUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubRollupUpdated represents a RollupUpdated event raised by the BridgeStub contract.
type BridgeStubRollupUpdated struct {
	Rollup common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRollupUpdated is a free log retrieval operation binding the contract event 0xae1f5aa15f6ff844896347ceca2a3c24c8d3a27785efdeacd581a0a95172784a.
//
// Solidity: event RollupUpdated(address rollup)
func (_BridgeStub *BridgeStubFilterer) FilterRollupUpdated(opts *bind.FilterOpts) (*BridgeStubRollupUpdatedIterator, error) {

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "RollupUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeStubRollupUpdatedIterator{contract: _BridgeStub.contract, event: "RollupUpdated", logs: logs, sub: sub}, nil
}

// WatchRollupUpdated is a free log subscription operation binding the contract event 0xae1f5aa15f6ff844896347ceca2a3c24c8d3a27785efdeacd581a0a95172784a.
//
// Solidity: event RollupUpdated(address rollup)
func (_BridgeStub *BridgeStubFilterer) WatchRollupUpdated(opts *bind.WatchOpts, sink chan<- *BridgeStubRollupUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "RollupUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubRollupUpdated)
				if err := _BridgeStub.contract.UnpackLog(event, "RollupUpdated", log); err != nil {
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

// ParseRollupUpdated is a log parse operation binding the contract event 0xae1f5aa15f6ff844896347ceca2a3c24c8d3a27785efdeacd581a0a95172784a.
//
// Solidity: event RollupUpdated(address rollup)
func (_BridgeStub *BridgeStubFilterer) ParseRollupUpdated(log types.Log) (*BridgeStubRollupUpdated, error) {
	event := new(BridgeStubRollupUpdated)
	if err := _BridgeStub.contract.UnpackLog(event, "RollupUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeStubSequencerInboxUpdatedIterator is returned from FilterSequencerInboxUpdated and is used to iterate over the raw logs and unpacked data for SequencerInboxUpdated events raised by the BridgeStub contract.
type BridgeStubSequencerInboxUpdatedIterator struct {
	Event *BridgeStubSequencerInboxUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeStubSequencerInboxUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubSequencerInboxUpdated)
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
		it.Event = new(BridgeStubSequencerInboxUpdated)
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
func (it *BridgeStubSequencerInboxUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubSequencerInboxUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubSequencerInboxUpdated represents a SequencerInboxUpdated event raised by the BridgeStub contract.
type BridgeStubSequencerInboxUpdated struct {
	NewSequencerInbox common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSequencerInboxUpdated is a free log retrieval operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeStub *BridgeStubFilterer) FilterSequencerInboxUpdated(opts *bind.FilterOpts) (*BridgeStubSequencerInboxUpdatedIterator, error) {

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeStubSequencerInboxUpdatedIterator{contract: _BridgeStub.contract, event: "SequencerInboxUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerInboxUpdated is a free log subscription operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeStub *BridgeStubFilterer) WatchSequencerInboxUpdated(opts *bind.WatchOpts, sink chan<- *BridgeStubSequencerInboxUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubSequencerInboxUpdated)
				if err := _BridgeStub.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
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

// ParseSequencerInboxUpdated is a log parse operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeStub *BridgeStubFilterer) ParseSequencerInboxUpdated(log types.Log) (*BridgeStubSequencerInboxUpdated, error) {
	event := new(BridgeStubSequencerInboxUpdated)
	if err := _BridgeStub.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnproxiedMetaData contains all meta data concerning the BridgeUnproxied contract.
var BridgeUnproxiedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"BadSequencerMessageNumber\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"}],\"name\":\"InvalidOutboxSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotDelayedInbox\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotOutbox\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotRollupOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotSequencerInbox\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BridgeCallTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"InboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"OutboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"RollupUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSequencerInbox\",\"type\":\"address\"}],\"name\":\"SequencerInboxUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DUMMYVAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptFundsFromOldBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeOutbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedDelayedInboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"allowedDelayedInboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedOutboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"}],\"name\":\"allowedOutboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delayedInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"enqueueDelayedMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"enqueueSequencerMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seqMessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"acc\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"rollup_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerReportedSubMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setDelayedInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencerInbox\",\"type\":\"address\"}],\"name\":\"setSequencerInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMsgCount\",\"type\":\"uint256\"}],\"name\":\"setSequencerReportedSubMessageCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"submitBatchSpendingReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"_rollup\",\"type\":\"address\"}],\"name\":\"updateRollupAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405230608052600060335534801561001957600080fd5b50600580546001600160a01b03199081166001600160a01b03179091556008805490911633179055608051611a8f61005c6000396000610f480152611a8f6000f3fe60806040526004361061018a5760003560e01c80639e5d4c49116100d6578063d5719dc21161007f578063eca067ad11610059578063eca067ad14610478578063ee35f3271461048d578063f81ff3b3146104ad57600080fd5b8063d5719dc214610438578063e76f5c8d14610458578063e77145f41461023f57600080fd5b8063c4d66de8116100b0578063c4d66de8146103d8578063cb23bcb5146103f8578063cee3d7281461041857600080fd5b80639e5d4c4914610358578063ab5d894314610386578063ae60bd131461039b57600080fd5b80637a88b10711610138578063919cc70611610112578063919cc706146102ea578063927dcfab1461030a578063945e11471461032057600080fd5b80637a88b1071461027757806386598a56146102975780638db5993b146102d757600080fd5b806347fb24c51161016957806347fb24c51461021f5780634f61f850146102415780635fca4a161461026157600080fd5b806284120c1461018f57806316bf5579146101b3578063413b35bd146101d3575b600080fd5b34801561019b57600080fd5b506007545b6040519081526020015b60405180910390f35b3480156101bf57600080fd5b506101a06101ce36600461177f565b6104cd565b3480156101df57600080fd5b5061020f6101ee3660046117ad565b6001600160a01b031660009081526002602052604090206001015460ff1690565b60405190151581526020016101aa565b34801561022b57600080fd5b5061023f61023a3660046117d1565b6104ee565b005b34801561024d57600080fd5b5061023f61025c3660046117ad565b6107f4565b34801561026d57600080fd5b506101a0600a5481565b34801561028357600080fd5b506101a061029236600461180f565b610920565b3480156102a357600080fd5b506102b76102b236600461183b565b610981565b6040805194855260208501939093529183015260608201526080016101aa565b6101a06102e536600461186d565b610b17565b3480156102f657600080fd5b5061023f6103053660046117ad565b610b2d565b34801561031657600080fd5b506101a060335481565b34801561032c57600080fd5b5061034061033b36600461177f565b610c52565b6040516001600160a01b0390911681526020016101aa565b34801561036457600080fd5b506103786103733660046118b4565b610c7c565b6040516101aa929190611961565b34801561039257600080fd5b50610340610e12565b3480156103a757600080fd5b5061020f6103b63660046117ad565b6001600160a01b03166000908152600160208190526040909120015460ff1690565b3480156103e457600080fd5b5061023f6103f33660046117ad565b610e55565b34801561040457600080fd5b50600854610340906001600160a01b031681565b34801561042457600080fd5b5061023f6104333660046117d1565b611079565b34801561044457600080fd5b506101a061045336600461177f565b6113e7565b34801561046457600080fd5b5061034061047336600461177f565b6113f7565b34801561048457600080fd5b506006546101a0565b34801561049957600080fd5b50600954610340906001600160a01b031681565b3480156104b957600080fd5b5061023f6104c836600461177f565b611407565b600781815481106104dd57600080fd5b600091825260209091200154905081565b6008546001600160a01b031633146105bd5760085460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa15801561054a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056e919061199d565b9050336001600160a01b038216146105bb57600854604051630739600760e01b81523360048201526001600160a01b03918216602482015290821660448201526064015b60405180910390fd5b505b6001600160a01b0382166000818152600160208181526040928390209182015492518515158152919360ff90931692917f6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521910160405180910390a2821515811515036106295750505050565b82156106c457604080518082018252600380548252600160208084018281526001600160a01b038a166000818152928490529582209451855551938201805460ff1916941515949094179093558154908101825591527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01805473ffffffffffffffffffffffffffffffffffffffff191690911790556107ed565b600380546106d4906001906119ba565b815481106106e4576106e46119db565b6000918252602090912001548254600380546001600160a01b03909316929091908110610713576107136119db565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160000154600160006003856000015481548110610761576107616119db565b60009182526020808320909101546001600160a01b031683528201929092526040019020556003805480610797576107976119f1565b600082815260208082208301600019908101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b03861682526001908190526040822091825501805460ff191690555b50505b5050565b6008546001600160a01b031633146108be5760085460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610850573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610874919061199d565b9050336001600160a01b038216146108bc57600854604051630739600760e01b81523360048201526001600160a01b03918216602482015290821660448201526064016105b2565b505b6009805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a906020015b60405180910390a150565b6009546000906001600160a01b03163314610969576040517f88f84f040000000000000000000000000000000000000000000000000000000081523360048201526024016105b2565b610978600d84434248876114d6565b90505b92915050565b6009546000908190819081906001600160a01b031633146109d0576040517f88f84f040000000000000000000000000000000000000000000000000000000081523360048201526024016105b2565b85600a54141580156109e157508515155b80156109ee5750600a5415155b15610a3357600a546040517fe2051feb0000000000000000000000000000000000000000000000000000000081526004810191909152602481018790526044016105b2565b600a85905560075493508315610a6e576007610a506001866119ba565b81548110610a6057610a606119db565b906000526020600020015492505b8615610a9f576006610a816001896119ba565b81548110610a9157610a916119db565b906000526020600020015491505b60408051602081018590529081018990526060810183905260800160408051601f198184030181529190528051602090910120600780546001810182556000919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688018190559398929750909550919350915050565b6000610b25848484346116a8565b949350505050565b6008546001600160a01b03163314610bf75760085460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610b89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bad919061199d565b9050336001600160a01b03821614610bf557600854604051630739600760e01b81523360048201526001600160a01b03918216602482015290821660448201526064016105b2565b505b6008805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527fae1f5aa15f6ff844896347ceca2a3c24c8d3a27785efdeacd581a0a95172784a90602001610915565b60048181548110610c6257600080fd5b6000918252602090912001546001600160a01b0316905081565b3360009081526002602052604081206001015460609060ff16610ccd576040517f32ea82ab0000000000000000000000000000000000000000000000000000000081523360048201526024016105b2565b8215801590610ce457506001600160a01b0386163b155b15610d26576040517fb5cf5b8f0000000000000000000000000000000000000000000000000000000081526001600160a01b03871660048201526024016105b2565b6005805473ffffffffffffffffffffffffffffffffffffffff1981163317909155604080516020601f87018190048102820181019092528581526001600160a01b0390921691610d949189918991899089908190840183828082843760009201919091525061171092505050565b6005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038581169190911790915560405192955090935088169033907f2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d46690610e00908a908a908a90611a07565b60405180910390a35094509492505050565b6005546000906001600160a01b03167fffffffffffffffffffffffff00000000000000000000000000000000000000018101610e5057600091505090565b919050565b600054610100900460ff1615808015610e755750600054600160ff909116105b80610e8f5750303b158015610e8f575060005460ff166001145b610f1b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016105b2565b6000805460ff191660011790558015610f3e576000805461ff0019166101001790555b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003610ff6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c000000000000000000000000000000000000000060648201526084016105b2565b600580546001600160a01b0373ffffffffffffffffffffffffffffffffffffffff1991821681179092556008805490911691841691909117905580156107f0576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6008546001600160a01b031633146111435760085460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa1580156110d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110f9919061199d565b9050336001600160a01b0382161461114157600854604051630739600760e01b81523360048201526001600160a01b03918216602482015290821660448201526064016105b2565b505b7fffffffffffffffffffffffff00000000000000000000000000000000000000016001600160a01b038316016111b0576040517f77abed100000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016105b2565b6001600160a01b038216600081815260026020908152604091829020600181015492518515158152909360ff90931692917f49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa910160405180910390a28215158115150361121d5750505050565b82156112b957604080518082018252600480548252600160208084018281526001600160a01b038a16600081815260029093529582209451855551938201805460ff1916941515949094179093558154908101825591527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01805473ffffffffffffffffffffffffffffffffffffffff191690911790556107ed565b600480546112c9906001906119ba565b815481106112d9576112d96119db565b6000918252602090912001548254600480546001600160a01b03909316929091908110611308576113086119db565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160000154600260006004856000015481548110611356576113566119db565b60009182526020808320909101546001600160a01b03168352820192909252604001902055600480548061138c5761138c6119f1565b600082815260208082208301600019908101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b03861682526002905260408120908155600101805460ff1916905550505050565b600681815481106104dd57600080fd5b60038181548110610c6257600080fd5b6008546001600160a01b031633146114d15760085460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015611463573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611487919061199d565b9050336001600160a01b038216146114cf57600854604051630739600760e01b81523360048201526001600160a01b03918216602482015290821660448201526064016105b2565b505b600a55565b600654604080517fff0000000000000000000000000000000000000000000000000000000000000060f88a901b166020808301919091527fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060608a901b1660218301527fffffffffffffffff00000000000000000000000000000000000000000000000060c089811b8216603585015288901b16603d830152604582018490526065820186905260858083018690528351808403909101815260a5909201909252805191012060009190600082156115d35760066115b56001856119ba565b815481106115c5576115c56119db565b906000526020600020015490505b6040805160208082018490528183018590528251808303840181526060830180855281519190920120600680546001810182556000919091527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f015533905260ff8c1660808201526001600160a01b038b1660a082015260c0810187905260e0810188905267ffffffffffffffff89166101008201529051829185917f5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1918190036101200190a3509098975050505050505050565b3360009081526001602081905260408220015460ff166116f6576040517fb6c60ea30000000000000000000000000000000000000000000000000000000081523360048201526024016105b2565b60006117068686434248896114d6565b9695505050505050565b60006060846001600160a01b0316848460405161172d9190611a3d565b60006040518083038185875af1925050503d806000811461176a576040519150601f19603f3d011682016040523d82523d6000602084013e61176f565b606091505b5090969095509350505050565b50565b60006020828403121561179157600080fd5b5035919050565b6001600160a01b038116811461177c57600080fd5b6000602082840312156117bf57600080fd5b81356117ca81611798565b9392505050565b600080604083850312156117e457600080fd5b82356117ef81611798565b91506020830135801515811461180457600080fd5b809150509250929050565b6000806040838503121561182257600080fd5b823561182d81611798565b946020939093013593505050565b6000806000806080858703121561185157600080fd5b5050823594602084013594506040840135936060013592509050565b60008060006060848603121561188257600080fd5b833560ff8116811461189357600080fd5b925060208401356118a381611798565b929592945050506040919091013590565b600080600080606085870312156118ca57600080fd5b84356118d581611798565b935060208501359250604085013567ffffffffffffffff808211156118f957600080fd5b818701915087601f83011261190d57600080fd5b81358181111561191c57600080fd5b88602082850101111561192e57600080fd5b95989497505060200194505050565b60005b83811015611958578181015183820152602001611940565b50506000910152565b8215158152604060208201526000825180604084015261198881606085016020870161193d565b601f01601f1916919091016060019392505050565b6000602082840312156119af57600080fd5b81516117ca81611798565b8181038181111561097b57634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b83815260406020820152816040820152818360608301376000818301606090810191909152601f909201601f1916010192915050565b60008251611a4f81846020870161193d565b919091019291505056fea2646970667358221220ac97507e47410f72055ecd86b5d120c3799da11edac6acfc5cc8a086e4e7de5564736f6c63430008110033",
}

// BridgeUnproxiedABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeUnproxiedMetaData.ABI instead.
var BridgeUnproxiedABI = BridgeUnproxiedMetaData.ABI

// BridgeUnproxiedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeUnproxiedMetaData.Bin instead.
var BridgeUnproxiedBin = BridgeUnproxiedMetaData.Bin

// DeployBridgeUnproxied deploys a new Ethereum contract, binding an instance of BridgeUnproxied to it.
func DeployBridgeUnproxied(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeUnproxied, error) {
	parsed, err := BridgeUnproxiedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeUnproxiedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeUnproxied{BridgeUnproxiedCaller: BridgeUnproxiedCaller{contract: contract}, BridgeUnproxiedTransactor: BridgeUnproxiedTransactor{contract: contract}, BridgeUnproxiedFilterer: BridgeUnproxiedFilterer{contract: contract}}, nil
}

// BridgeUnproxied is an auto generated Go binding around an Ethereum contract.
type BridgeUnproxied struct {
	BridgeUnproxiedCaller     // Read-only binding to the contract
	BridgeUnproxiedTransactor // Write-only binding to the contract
	BridgeUnproxiedFilterer   // Log filterer for contract events
}

// BridgeUnproxiedCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeUnproxiedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeUnproxiedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeUnproxiedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeUnproxiedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeUnproxiedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeUnproxiedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeUnproxiedSession struct {
	Contract     *BridgeUnproxied  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeUnproxiedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeUnproxiedCallerSession struct {
	Contract *BridgeUnproxiedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BridgeUnproxiedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeUnproxiedTransactorSession struct {
	Contract     *BridgeUnproxiedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BridgeUnproxiedRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeUnproxiedRaw struct {
	Contract *BridgeUnproxied // Generic contract binding to access the raw methods on
}

// BridgeUnproxiedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeUnproxiedCallerRaw struct {
	Contract *BridgeUnproxiedCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeUnproxiedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeUnproxiedTransactorRaw struct {
	Contract *BridgeUnproxiedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeUnproxied creates a new instance of BridgeUnproxied, bound to a specific deployed contract.
func NewBridgeUnproxied(address common.Address, backend bind.ContractBackend) (*BridgeUnproxied, error) {
	contract, err := bindBridgeUnproxied(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxied{BridgeUnproxiedCaller: BridgeUnproxiedCaller{contract: contract}, BridgeUnproxiedTransactor: BridgeUnproxiedTransactor{contract: contract}, BridgeUnproxiedFilterer: BridgeUnproxiedFilterer{contract: contract}}, nil
}

// NewBridgeUnproxiedCaller creates a new read-only instance of BridgeUnproxied, bound to a specific deployed contract.
func NewBridgeUnproxiedCaller(address common.Address, caller bind.ContractCaller) (*BridgeUnproxiedCaller, error) {
	contract, err := bindBridgeUnproxied(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedCaller{contract: contract}, nil
}

// NewBridgeUnproxiedTransactor creates a new write-only instance of BridgeUnproxied, bound to a specific deployed contract.
func NewBridgeUnproxiedTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeUnproxiedTransactor, error) {
	contract, err := bindBridgeUnproxied(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedTransactor{contract: contract}, nil
}

// NewBridgeUnproxiedFilterer creates a new log filterer instance of BridgeUnproxied, bound to a specific deployed contract.
func NewBridgeUnproxiedFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeUnproxiedFilterer, error) {
	contract, err := bindBridgeUnproxied(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedFilterer{contract: contract}, nil
}

// bindBridgeUnproxied binds a generic wrapper to an already deployed contract.
func bindBridgeUnproxied(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeUnproxiedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeUnproxied *BridgeUnproxiedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeUnproxied.Contract.BridgeUnproxiedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeUnproxied *BridgeUnproxiedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.BridgeUnproxiedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeUnproxied *BridgeUnproxiedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.BridgeUnproxiedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeUnproxied *BridgeUnproxiedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeUnproxied.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeUnproxied *BridgeUnproxiedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeUnproxied *BridgeUnproxiedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.contract.Transact(opts, method, params...)
}

// DUMMYVAR is a free data retrieval call binding the contract method 0x927dcfab.
//
// Solidity: function DUMMYVAR() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedCaller) DUMMYVAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "DUMMYVAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DUMMYVAR is a free data retrieval call binding the contract method 0x927dcfab.
//
// Solidity: function DUMMYVAR() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedSession) DUMMYVAR() (*big.Int, error) {
	return _BridgeUnproxied.Contract.DUMMYVAR(&_BridgeUnproxied.CallOpts)
}

// DUMMYVAR is a free data retrieval call binding the contract method 0x927dcfab.
//
// Solidity: function DUMMYVAR() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) DUMMYVAR() (*big.Int, error) {
	return _BridgeUnproxied.Contract.DUMMYVAR(&_BridgeUnproxied.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCaller) ActiveOutbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "activeOutbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedSession) ActiveOutbox() (common.Address, error) {
	return _BridgeUnproxied.Contract.ActiveOutbox(&_BridgeUnproxied.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) ActiveOutbox() (common.Address, error) {
	return _BridgeUnproxied.Contract.ActiveOutbox(&_BridgeUnproxied.CallOpts)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCaller) AllowedDelayedInboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "allowedDelayedInboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeUnproxied.Contract.AllowedDelayedInboxList(&_BridgeUnproxied.CallOpts, arg0)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeUnproxied.Contract.AllowedDelayedInboxList(&_BridgeUnproxied.CallOpts, arg0)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeUnproxied *BridgeUnproxiedCaller) AllowedDelayedInboxes(opts *bind.CallOpts, inbox common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "allowedDelayedInboxes", inbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeUnproxied *BridgeUnproxiedSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeUnproxied.Contract.AllowedDelayedInboxes(&_BridgeUnproxied.CallOpts, inbox)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeUnproxied.Contract.AllowedDelayedInboxes(&_BridgeUnproxied.CallOpts, inbox)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCaller) AllowedOutboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "allowedOutboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeUnproxied.Contract.AllowedOutboxList(&_BridgeUnproxied.CallOpts, arg0)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeUnproxied.Contract.AllowedOutboxList(&_BridgeUnproxied.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeUnproxied *BridgeUnproxiedCaller) AllowedOutboxes(opts *bind.CallOpts, outbox common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "allowedOutboxes", outbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeUnproxied *BridgeUnproxiedSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _BridgeUnproxied.Contract.AllowedOutboxes(&_BridgeUnproxied.CallOpts, outbox)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _BridgeUnproxied.Contract.AllowedOutboxes(&_BridgeUnproxied.CallOpts, outbox)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeUnproxied *BridgeUnproxiedCaller) DelayedInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "delayedInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeUnproxied *BridgeUnproxiedSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeUnproxied.Contract.DelayedInboxAccs(&_BridgeUnproxied.CallOpts, arg0)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeUnproxied.Contract.DelayedInboxAccs(&_BridgeUnproxied.CallOpts, arg0)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedCaller) DelayedMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "delayedMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeUnproxied.Contract.DelayedMessageCount(&_BridgeUnproxied.CallOpts)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeUnproxied.Contract.DelayedMessageCount(&_BridgeUnproxied.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedSession) Rollup() (common.Address, error) {
	return _BridgeUnproxied.Contract.Rollup(&_BridgeUnproxied.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) Rollup() (common.Address, error) {
	return _BridgeUnproxied.Contract.Rollup(&_BridgeUnproxied.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedSession) SequencerInbox() (common.Address, error) {
	return _BridgeUnproxied.Contract.SequencerInbox(&_BridgeUnproxied.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) SequencerInbox() (common.Address, error) {
	return _BridgeUnproxied.Contract.SequencerInbox(&_BridgeUnproxied.CallOpts)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeUnproxied *BridgeUnproxiedCaller) SequencerInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "sequencerInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeUnproxied *BridgeUnproxiedSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeUnproxied.Contract.SequencerInboxAccs(&_BridgeUnproxied.CallOpts, arg0)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeUnproxied.Contract.SequencerInboxAccs(&_BridgeUnproxied.CallOpts, arg0)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedCaller) SequencerMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "sequencerMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeUnproxied.Contract.SequencerMessageCount(&_BridgeUnproxied.CallOpts)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeUnproxied.Contract.SequencerMessageCount(&_BridgeUnproxied.CallOpts)
}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedCaller) SequencerReportedSubMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "sequencerReportedSubMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedSession) SequencerReportedSubMessageCount() (*big.Int, error) {
	return _BridgeUnproxied.Contract.SequencerReportedSubMessageCount(&_BridgeUnproxied.CallOpts)
}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) SequencerReportedSubMessageCount() (*big.Int, error) {
	return _BridgeUnproxied.Contract.SequencerReportedSubMessageCount(&_BridgeUnproxied.CallOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactor) AcceptFundsFromOldBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "acceptFundsFromOldBridge")
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeUnproxied *BridgeUnproxiedSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.AcceptFundsFromOldBridge(&_BridgeUnproxied.TransactOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.AcceptFundsFromOldBridge(&_BridgeUnproxied.TransactOpts)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedTransactor) EnqueueDelayedMessage(opts *bind.TransactOpts, kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "enqueueDelayedMessage", kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.EnqueueDelayedMessage(&_BridgeUnproxied.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.EnqueueDelayedMessage(&_BridgeUnproxied.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeUnproxied *BridgeUnproxiedTransactor) EnqueueSequencerMessage(opts *bind.TransactOpts, dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "enqueueSequencerMessage", dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeUnproxied *BridgeUnproxiedSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.EnqueueSequencerMessage(&_BridgeUnproxied.TransactOpts, dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.EnqueueSequencerMessage(&_BridgeUnproxied.TransactOpts, dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeUnproxied *BridgeUnproxiedTransactor) ExecuteCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "executeCall", to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeUnproxied *BridgeUnproxiedSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.ExecuteCall(&_BridgeUnproxied.TransactOpts, to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.ExecuteCall(&_BridgeUnproxied.TransactOpts, to, value, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactor) Initialize(opts *bind.TransactOpts, rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "initialize", rollup_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeUnproxied *BridgeUnproxiedSession) Initialize(rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.Initialize(&_BridgeUnproxied.TransactOpts, rollup_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) Initialize(rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.Initialize(&_BridgeUnproxied.TransactOpts, rollup_)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactor) SetDelayedInbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "setDelayedInbox", inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeUnproxied *BridgeUnproxiedSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetDelayedInbox(&_BridgeUnproxied.TransactOpts, inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetDelayedInbox(&_BridgeUnproxied.TransactOpts, inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactor) SetOutbox(opts *bind.TransactOpts, outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "setOutbox", outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeUnproxied *BridgeUnproxiedSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetOutbox(&_BridgeUnproxied.TransactOpts, outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetOutbox(&_BridgeUnproxied.TransactOpts, outbox, enabled)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactor) SetSequencerInbox(opts *bind.TransactOpts, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "setSequencerInbox", _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeUnproxied *BridgeUnproxiedSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetSequencerInbox(&_BridgeUnproxied.TransactOpts, _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetSequencerInbox(&_BridgeUnproxied.TransactOpts, _sequencerInbox)
}

// SetSequencerReportedSubMessageCount is a paid mutator transaction binding the contract method 0xf81ff3b3.
//
// Solidity: function setSequencerReportedSubMessageCount(uint256 newMsgCount) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactor) SetSequencerReportedSubMessageCount(opts *bind.TransactOpts, newMsgCount *big.Int) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "setSequencerReportedSubMessageCount", newMsgCount)
}

// SetSequencerReportedSubMessageCount is a paid mutator transaction binding the contract method 0xf81ff3b3.
//
// Solidity: function setSequencerReportedSubMessageCount(uint256 newMsgCount) returns()
func (_BridgeUnproxied *BridgeUnproxiedSession) SetSequencerReportedSubMessageCount(newMsgCount *big.Int) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetSequencerReportedSubMessageCount(&_BridgeUnproxied.TransactOpts, newMsgCount)
}

// SetSequencerReportedSubMessageCount is a paid mutator transaction binding the contract method 0xf81ff3b3.
//
// Solidity: function setSequencerReportedSubMessageCount(uint256 newMsgCount) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) SetSequencerReportedSubMessageCount(newMsgCount *big.Int) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetSequencerReportedSubMessageCount(&_BridgeUnproxied.TransactOpts, newMsgCount)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address sender, bytes32 messageDataHash) returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedTransactor) SubmitBatchSpendingReport(opts *bind.TransactOpts, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "submitBatchSpendingReport", sender, messageDataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address sender, bytes32 messageDataHash) returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedSession) SubmitBatchSpendingReport(sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SubmitBatchSpendingReport(&_BridgeUnproxied.TransactOpts, sender, messageDataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address sender, bytes32 messageDataHash) returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) SubmitBatchSpendingReport(sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SubmitBatchSpendingReport(&_BridgeUnproxied.TransactOpts, sender, messageDataHash)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x919cc706.
//
// Solidity: function updateRollupAddress(address _rollup) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactor) UpdateRollupAddress(opts *bind.TransactOpts, _rollup common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "updateRollupAddress", _rollup)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x919cc706.
//
// Solidity: function updateRollupAddress(address _rollup) returns()
func (_BridgeUnproxied *BridgeUnproxiedSession) UpdateRollupAddress(_rollup common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.UpdateRollupAddress(&_BridgeUnproxied.TransactOpts, _rollup)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x919cc706.
//
// Solidity: function updateRollupAddress(address _rollup) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) UpdateRollupAddress(_rollup common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.UpdateRollupAddress(&_BridgeUnproxied.TransactOpts, _rollup)
}

// BridgeUnproxiedBridgeCallTriggeredIterator is returned from FilterBridgeCallTriggered and is used to iterate over the raw logs and unpacked data for BridgeCallTriggered events raised by the BridgeUnproxied contract.
type BridgeUnproxiedBridgeCallTriggeredIterator struct {
	Event *BridgeUnproxiedBridgeCallTriggered // Event containing the contract specifics and raw log

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
func (it *BridgeUnproxiedBridgeCallTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnproxiedBridgeCallTriggered)
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
		it.Event = new(BridgeUnproxiedBridgeCallTriggered)
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
func (it *BridgeUnproxiedBridgeCallTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnproxiedBridgeCallTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnproxiedBridgeCallTriggered represents a BridgeCallTriggered event raised by the BridgeUnproxied contract.
type BridgeUnproxiedBridgeCallTriggered struct {
	Outbox common.Address
	To     common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeCallTriggered is a free log retrieval operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) FilterBridgeCallTriggered(opts *bind.FilterOpts, outbox []common.Address, to []common.Address) (*BridgeUnproxiedBridgeCallTriggeredIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.FilterLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedBridgeCallTriggeredIterator{contract: _BridgeUnproxied.contract, event: "BridgeCallTriggered", logs: logs, sub: sub}, nil
}

// WatchBridgeCallTriggered is a free log subscription operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) WatchBridgeCallTriggered(opts *bind.WatchOpts, sink chan<- *BridgeUnproxiedBridgeCallTriggered, outbox []common.Address, to []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.WatchLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnproxiedBridgeCallTriggered)
				if err := _BridgeUnproxied.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
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

// ParseBridgeCallTriggered is a log parse operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) ParseBridgeCallTriggered(log types.Log) (*BridgeUnproxiedBridgeCallTriggered, error) {
	event := new(BridgeUnproxiedBridgeCallTriggered)
	if err := _BridgeUnproxied.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnproxiedInboxToggleIterator is returned from FilterInboxToggle and is used to iterate over the raw logs and unpacked data for InboxToggle events raised by the BridgeUnproxied contract.
type BridgeUnproxiedInboxToggleIterator struct {
	Event *BridgeUnproxiedInboxToggle // Event containing the contract specifics and raw log

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
func (it *BridgeUnproxiedInboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnproxiedInboxToggle)
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
		it.Event = new(BridgeUnproxiedInboxToggle)
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
func (it *BridgeUnproxiedInboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnproxiedInboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnproxiedInboxToggle represents a InboxToggle event raised by the BridgeUnproxied contract.
type BridgeUnproxiedInboxToggle struct {
	Inbox   common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInboxToggle is a free log retrieval operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) FilterInboxToggle(opts *bind.FilterOpts, inbox []common.Address) (*BridgeUnproxiedInboxToggleIterator, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.FilterLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedInboxToggleIterator{contract: _BridgeUnproxied.contract, event: "InboxToggle", logs: logs, sub: sub}, nil
}

// WatchInboxToggle is a free log subscription operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) WatchInboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeUnproxiedInboxToggle, inbox []common.Address) (event.Subscription, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.WatchLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnproxiedInboxToggle)
				if err := _BridgeUnproxied.contract.UnpackLog(event, "InboxToggle", log); err != nil {
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

// ParseInboxToggle is a log parse operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) ParseInboxToggle(log types.Log) (*BridgeUnproxiedInboxToggle, error) {
	event := new(BridgeUnproxiedInboxToggle)
	if err := _BridgeUnproxied.contract.UnpackLog(event, "InboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnproxiedInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeUnproxied contract.
type BridgeUnproxiedInitializedIterator struct {
	Event *BridgeUnproxiedInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeUnproxiedInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnproxiedInitialized)
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
		it.Event = new(BridgeUnproxiedInitialized)
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
func (it *BridgeUnproxiedInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnproxiedInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnproxiedInitialized represents a Initialized event raised by the BridgeUnproxied contract.
type BridgeUnproxiedInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeUnproxiedInitializedIterator, error) {

	logs, sub, err := _BridgeUnproxied.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedInitializedIterator{contract: _BridgeUnproxied.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeUnproxiedInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeUnproxied.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnproxiedInitialized)
				if err := _BridgeUnproxied.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) ParseInitialized(log types.Log) (*BridgeUnproxiedInitialized, error) {
	event := new(BridgeUnproxiedInitialized)
	if err := _BridgeUnproxied.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnproxiedMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the BridgeUnproxied contract.
type BridgeUnproxiedMessageDeliveredIterator struct {
	Event *BridgeUnproxiedMessageDelivered // Event containing the contract specifics and raw log

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
func (it *BridgeUnproxiedMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnproxiedMessageDelivered)
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
		it.Event = new(BridgeUnproxiedMessageDelivered)
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
func (it *BridgeUnproxiedMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnproxiedMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnproxiedMessageDelivered represents a MessageDelivered event raised by the BridgeUnproxied contract.
type BridgeUnproxiedMessageDelivered struct {
	MessageIndex    *big.Int
	BeforeInboxAcc  [32]byte
	Inbox           common.Address
	Kind            uint8
	Sender          common.Address
	MessageDataHash [32]byte
	BaseFeeL1       *big.Int
	Timestamp       uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (*BridgeUnproxiedMessageDeliveredIterator, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.FilterLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedMessageDeliveredIterator{contract: _BridgeUnproxied.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *BridgeUnproxiedMessageDelivered, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.WatchLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnproxiedMessageDelivered)
				if err := _BridgeUnproxied.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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

// ParseMessageDelivered is a log parse operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) ParseMessageDelivered(log types.Log) (*BridgeUnproxiedMessageDelivered, error) {
	event := new(BridgeUnproxiedMessageDelivered)
	if err := _BridgeUnproxied.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnproxiedOutboxToggleIterator is returned from FilterOutboxToggle and is used to iterate over the raw logs and unpacked data for OutboxToggle events raised by the BridgeUnproxied contract.
type BridgeUnproxiedOutboxToggleIterator struct {
	Event *BridgeUnproxiedOutboxToggle // Event containing the contract specifics and raw log

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
func (it *BridgeUnproxiedOutboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnproxiedOutboxToggle)
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
		it.Event = new(BridgeUnproxiedOutboxToggle)
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
func (it *BridgeUnproxiedOutboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnproxiedOutboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnproxiedOutboxToggle represents a OutboxToggle event raised by the BridgeUnproxied contract.
type BridgeUnproxiedOutboxToggle struct {
	Outbox  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOutboxToggle is a free log retrieval operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) FilterOutboxToggle(opts *bind.FilterOpts, outbox []common.Address) (*BridgeUnproxiedOutboxToggleIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.FilterLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedOutboxToggleIterator{contract: _BridgeUnproxied.contract, event: "OutboxToggle", logs: logs, sub: sub}, nil
}

// WatchOutboxToggle is a free log subscription operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) WatchOutboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeUnproxiedOutboxToggle, outbox []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.WatchLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnproxiedOutboxToggle)
				if err := _BridgeUnproxied.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
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

// ParseOutboxToggle is a log parse operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) ParseOutboxToggle(log types.Log) (*BridgeUnproxiedOutboxToggle, error) {
	event := new(BridgeUnproxiedOutboxToggle)
	if err := _BridgeUnproxied.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnproxiedRollupUpdatedIterator is returned from FilterRollupUpdated and is used to iterate over the raw logs and unpacked data for RollupUpdated events raised by the BridgeUnproxied contract.
type BridgeUnproxiedRollupUpdatedIterator struct {
	Event *BridgeUnproxiedRollupUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeUnproxiedRollupUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnproxiedRollupUpdated)
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
		it.Event = new(BridgeUnproxiedRollupUpdated)
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
func (it *BridgeUnproxiedRollupUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnproxiedRollupUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnproxiedRollupUpdated represents a RollupUpdated event raised by the BridgeUnproxied contract.
type BridgeUnproxiedRollupUpdated struct {
	Rollup common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRollupUpdated is a free log retrieval operation binding the contract event 0xae1f5aa15f6ff844896347ceca2a3c24c8d3a27785efdeacd581a0a95172784a.
//
// Solidity: event RollupUpdated(address rollup)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) FilterRollupUpdated(opts *bind.FilterOpts) (*BridgeUnproxiedRollupUpdatedIterator, error) {

	logs, sub, err := _BridgeUnproxied.contract.FilterLogs(opts, "RollupUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedRollupUpdatedIterator{contract: _BridgeUnproxied.contract, event: "RollupUpdated", logs: logs, sub: sub}, nil
}

// WatchRollupUpdated is a free log subscription operation binding the contract event 0xae1f5aa15f6ff844896347ceca2a3c24c8d3a27785efdeacd581a0a95172784a.
//
// Solidity: event RollupUpdated(address rollup)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) WatchRollupUpdated(opts *bind.WatchOpts, sink chan<- *BridgeUnproxiedRollupUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeUnproxied.contract.WatchLogs(opts, "RollupUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnproxiedRollupUpdated)
				if err := _BridgeUnproxied.contract.UnpackLog(event, "RollupUpdated", log); err != nil {
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

// ParseRollupUpdated is a log parse operation binding the contract event 0xae1f5aa15f6ff844896347ceca2a3c24c8d3a27785efdeacd581a0a95172784a.
//
// Solidity: event RollupUpdated(address rollup)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) ParseRollupUpdated(log types.Log) (*BridgeUnproxiedRollupUpdated, error) {
	event := new(BridgeUnproxiedRollupUpdated)
	if err := _BridgeUnproxied.contract.UnpackLog(event, "RollupUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnproxiedSequencerInboxUpdatedIterator is returned from FilterSequencerInboxUpdated and is used to iterate over the raw logs and unpacked data for SequencerInboxUpdated events raised by the BridgeUnproxied contract.
type BridgeUnproxiedSequencerInboxUpdatedIterator struct {
	Event *BridgeUnproxiedSequencerInboxUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeUnproxiedSequencerInboxUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnproxiedSequencerInboxUpdated)
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
		it.Event = new(BridgeUnproxiedSequencerInboxUpdated)
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
func (it *BridgeUnproxiedSequencerInboxUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnproxiedSequencerInboxUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnproxiedSequencerInboxUpdated represents a SequencerInboxUpdated event raised by the BridgeUnproxied contract.
type BridgeUnproxiedSequencerInboxUpdated struct {
	NewSequencerInbox common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSequencerInboxUpdated is a free log retrieval operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) FilterSequencerInboxUpdated(opts *bind.FilterOpts) (*BridgeUnproxiedSequencerInboxUpdatedIterator, error) {

	logs, sub, err := _BridgeUnproxied.contract.FilterLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedSequencerInboxUpdatedIterator{contract: _BridgeUnproxied.contract, event: "SequencerInboxUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerInboxUpdated is a free log subscription operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) WatchSequencerInboxUpdated(opts *bind.WatchOpts, sink chan<- *BridgeUnproxiedSequencerInboxUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeUnproxied.contract.WatchLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnproxiedSequencerInboxUpdated)
				if err := _BridgeUnproxied.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
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

// ParseSequencerInboxUpdated is a log parse operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) ParseSequencerInboxUpdated(log types.Log) (*BridgeUnproxiedSequencerInboxUpdated, error) {
	event := new(BridgeUnproxiedSequencerInboxUpdated)
	if err := _BridgeUnproxied.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWETH9MetaData contains all meta data concerning the IWETH9 contract.
var IWETH9MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IWETH9ABI is the input ABI used to generate the binding from.
// Deprecated: Use IWETH9MetaData.ABI instead.
var IWETH9ABI = IWETH9MetaData.ABI

// IWETH9 is an auto generated Go binding around an Ethereum contract.
type IWETH9 struct {
	IWETH9Caller     // Read-only binding to the contract
	IWETH9Transactor // Write-only binding to the contract
	IWETH9Filterer   // Log filterer for contract events
}

// IWETH9Caller is an auto generated read-only Go binding around an Ethereum contract.
type IWETH9Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IWETH9Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWETH9Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWETH9Session struct {
	Contract     *IWETH9           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETH9CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWETH9CallerSession struct {
	Contract *IWETH9Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWETH9TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWETH9TransactorSession struct {
	Contract     *IWETH9Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETH9Raw is an auto generated low-level Go binding around an Ethereum contract.
type IWETH9Raw struct {
	Contract *IWETH9 // Generic contract binding to access the raw methods on
}

// IWETH9CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWETH9CallerRaw struct {
	Contract *IWETH9Caller // Generic read-only contract binding to access the raw methods on
}

// IWETH9TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWETH9TransactorRaw struct {
	Contract *IWETH9Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIWETH9 creates a new instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9(address common.Address, backend bind.ContractBackend) (*IWETH9, error) {
	contract, err := bindIWETH9(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWETH9{IWETH9Caller: IWETH9Caller{contract: contract}, IWETH9Transactor: IWETH9Transactor{contract: contract}, IWETH9Filterer: IWETH9Filterer{contract: contract}}, nil
}

// NewIWETH9Caller creates a new read-only instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Caller(address common.Address, caller bind.ContractCaller) (*IWETH9Caller, error) {
	contract, err := bindIWETH9(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH9Caller{contract: contract}, nil
}

// NewIWETH9Transactor creates a new write-only instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Transactor(address common.Address, transactor bind.ContractTransactor) (*IWETH9Transactor, error) {
	contract, err := bindIWETH9(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH9Transactor{contract: contract}, nil
}

// NewIWETH9Filterer creates a new log filterer instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Filterer(address common.Address, filterer bind.ContractFilterer) (*IWETH9Filterer, error) {
	contract, err := bindIWETH9(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWETH9Filterer{contract: contract}, nil
}

// bindIWETH9 binds a generic wrapper to an already deployed contract.
func bindIWETH9(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWETH9MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH9 *IWETH9Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH9.Contract.IWETH9Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH9 *IWETH9Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.Contract.IWETH9Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH9 *IWETH9Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH9.Contract.IWETH9Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH9 *IWETH9CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH9.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH9 *IWETH9TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH9 *IWETH9TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH9.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9Session) Deposit() (*types.Transaction, error) {
	return _IWETH9.Contract.Deposit(&_IWETH9.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9TransactorSession) Deposit() (*types.Transaction, error) {
	return _IWETH9.Contract.Deposit(&_IWETH9.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_IWETH9 *IWETH9Transactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_IWETH9 *IWETH9Session) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Withdraw(&_IWETH9.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_IWETH9 *IWETH9TransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Withdraw(&_IWETH9.TransactOpts, _amount)
}

// InboxStubMetaData contains all meta data concerning the InboxStub contract.
var InboxStubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"calculateRetryableSubmissionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"createRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDataSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedUnsignedTransactionToFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendUnsignedTransactionToFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sendWithdrawEthToFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"name\":\"setAllowList\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setAllowListEnabled\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"unsafeCreateRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506201cccc608052608051610dee61003360003960006103fb0152610dee6000f3fe6080604052600436106101a05760003560e01c80638456cb59116100e1578063c474d2c51161008a578063e78cea9211610064578063e78cea92146103bc578063e8eb1dc3146103e9578063ee35f3271461041d578063efeadb6d1461044a57600080fd5b8063c474d2c51461037e578063e3de72a51461039c578063e6bd12cf146102aa57600080fd5b8063a66b327d116100bb578063a66b327d14610328578063b75436bb14610343578063babcc5391461036357600080fd5b80638456cb591461021d5780638a631aa6146102d35780638b3240a0146102ee57600080fd5b80635075788b1161014e578063679b6ded11610128578063679b6ded1461029c57806367ef3ab8146102aa5780636e6e8a6a1461029c57806370665f14146102b857600080fd5b80635075788b146101a55780635c975abb1461025c5780635e9167581461028e57600080fd5b80633f4ba83a1161017f5780633f4ba83a1461021d578063439370b114610234578063485cc9551461023c57600080fd5b8062f72382146101a55780631fe927cf146101d857806322bd5c1c146101f8575b600080fd5b3480156101b157600080fd5b506101c56101c0366004610816565b610465565b6040519081526020015b60405180910390f35b3480156101e457600080fd5b506101c56101f3366004610893565b6104b5565b34801561020457600080fd5b5061020d610465565b60405190151581526020016101cf565b34801561022957600080fd5b50610232610560565b005b6101c5610465565b34801561024857600080fd5b506102326102573660046108d5565b6105a8565b34801561026857600080fd5b5060015461020d9074010000000000000000000000000000000000000000900460ff1681565b6101c56101c036600461090e565b6101c56101c0366004610978565b6101c56101c0366004610a1d565b3480156102c457600080fd5b506101c56101c0366004610a90565b3480156102df57600080fd5b506101c56101c0366004610add565b3480156102fa57600080fd5b50610303610465565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101cf565b34801561033457600080fd5b506101c56101c0366004610b32565b34801561034f57600080fd5b506101c561035e366004610893565b610656565b34801561036f57600080fd5b5061020d6101c0366004610b54565b34801561038a57600080fd5b50610232610399366004610b54565b50565b3480156103a857600080fd5b506102326103b7366004610c83565b6106b2565b3480156103c857600080fd5b506000546103039073ffffffffffffffffffffffffffffffffffffffff1681565b3480156103f557600080fd5b506101c57f000000000000000000000000000000000000000000000000000000000000000081565b34801561042957600080fd5b506001546103039073ffffffffffffffffffffffffffffffffffffffff1681565b34801561045657600080fd5b506102326103b7366004610d45565b60405162461bcd60e51b815260206004820152600f60248201527f4e4f545f494d504c454d454e544544000000000000000000000000000000000060448201526000906064015b60405180910390fd5b60003332146105065760405162461bcd60e51b815260206004820152600b60248201527f6f726967696e206f6e6c7900000000000000000000000000000000000000000060448201526064016104ac565b600061052b600333868660405161051e929190610d60565b60405180910390206106fa565b60405190915081907fab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c90600090a29392505050565b60405162461bcd60e51b815260206004820152600f60248201527f4e4f5420494d504c454d454e544544000000000000000000000000000000000060448201526064016104ac565b60005473ffffffffffffffffffffffffffffffffffffffff161561060e5760405162461bcd60e51b815260206004820152600c60248201527f414c52454144595f494e4954000000000000000000000000000000000000000060448201526064016104ac565b50600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60008061066f600333868660405161051e929190610d60565b9050807fff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b85856040516106a3929190610d70565b60405180910390a29392505050565b60405162461bcd60e51b815260206004820152600f60248201527f4e4f545f494d504c454d454e544544000000000000000000000000000000000060448201526064016104ac565b600080546040517f8db5993b00000000000000000000000000000000000000000000000000000000815260ff8616600482015273ffffffffffffffffffffffffffffffffffffffff85811660248301526044820185905290911690638db5993b90349060640160206040518083038185885af115801561077e573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906107a39190610d9f565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461039957600080fd5b60008083601f8401126107df57600080fd5b50813567ffffffffffffffff8111156107f757600080fd5b60208301915083602082850101111561080f57600080fd5b9250929050565b600080600080600080600060c0888a03121561083157600080fd5b8735965060208801359550604088013594506060880135610851816107ab565b93506080880135925060a088013567ffffffffffffffff81111561087457600080fd5b6108808a828b016107cd565b989b979a50959850939692959293505050565b600080602083850312156108a657600080fd5b823567ffffffffffffffff8111156108bd57600080fd5b6108c9858286016107cd565b90969095509350505050565b600080604083850312156108e857600080fd5b82356108f3816107ab565b91506020830135610903816107ab565b809150509250929050565b60008060008060006080868803121561092657600080fd5b8535945060208601359350604086013561093f816107ab565b9250606086013567ffffffffffffffff81111561095b57600080fd5b610967888289016107cd565b969995985093965092949392505050565b60008060008060008060008060006101008a8c03121561099757600080fd5b89356109a2816107ab565b985060208a0135975060408a0135965060608a01356109c0816107ab565b955060808a01356109d0816107ab565b945060a08a0135935060c08a0135925060e08a013567ffffffffffffffff8111156109fa57600080fd5b610a068c828d016107cd565b915080935050809150509295985092959850929598565b60008060008060008060a08789031215610a3657600080fd5b8635955060208701359450604087013593506060870135610a56816107ab565b9250608087013567ffffffffffffffff811115610a7257600080fd5b610a7e89828a016107cd565b979a9699509497509295939492505050565b600080600080600060a08688031215610aa857600080fd5b853594506020860135935060408601359250606086013591506080860135610acf816107ab565b809150509295509295909350565b60008060008060008060a08789031215610af657600080fd5b86359550602087013594506040870135610b0f816107ab565b935060608701359250608087013567ffffffffffffffff811115610a7257600080fd5b60008060408385031215610b4557600080fd5b50508035926020909101359150565b600060208284031215610b6657600080fd5b8135610b71816107ab565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610bd057610bd0610b78565b604052919050565b600067ffffffffffffffff821115610bf257610bf2610b78565b5060051b60200190565b80358015158114610c0c57600080fd5b919050565b600082601f830112610c2257600080fd5b81356020610c37610c3283610bd8565b610ba7565b82815260059290921b84018101918181019086841115610c5657600080fd5b8286015b84811015610c7857610c6b81610bfc565b8352918301918301610c5a565b509695505050505050565b60008060408385031215610c9657600080fd5b823567ffffffffffffffff80821115610cae57600080fd5b818501915085601f830112610cc257600080fd5b81356020610cd2610c3283610bd8565b82815260059290921b84018101918181019089841115610cf157600080fd5b948201945b83861015610d18578535610d09816107ab565b82529482019490820190610cf6565b96505086013592505080821115610d2e57600080fd5b50610d3b85828601610c11565b9150509250929050565b600060208284031215610d5757600080fd5b610b7182610bfc565b8183823760009101908152919050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b600060208284031215610db157600080fd5b505191905056fea26469706673582212201750ce18bd8947cb314cf031be672ec17e4236e2177e6d92060c9fc1240f446564736f6c63430008110033",
}

// InboxStubABI is the input ABI used to generate the binding from.
// Deprecated: Use InboxStubMetaData.ABI instead.
var InboxStubABI = InboxStubMetaData.ABI

// InboxStubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InboxStubMetaData.Bin instead.
var InboxStubBin = InboxStubMetaData.Bin

// DeployInboxStub deploys a new Ethereum contract, binding an instance of InboxStub to it.
func DeployInboxStub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InboxStub, error) {
	parsed, err := InboxStubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InboxStubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InboxStub{InboxStubCaller: InboxStubCaller{contract: contract}, InboxStubTransactor: InboxStubTransactor{contract: contract}, InboxStubFilterer: InboxStubFilterer{contract: contract}}, nil
}

// InboxStub is an auto generated Go binding around an Ethereum contract.
type InboxStub struct {
	InboxStubCaller     // Read-only binding to the contract
	InboxStubTransactor // Write-only binding to the contract
	InboxStubFilterer   // Log filterer for contract events
}

// InboxStubCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxStubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxStubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxStubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxStubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxStubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxStubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxStubSession struct {
	Contract     *InboxStub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxStubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxStubCallerSession struct {
	Contract *InboxStubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// InboxStubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxStubTransactorSession struct {
	Contract     *InboxStubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// InboxStubRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxStubRaw struct {
	Contract *InboxStub // Generic contract binding to access the raw methods on
}

// InboxStubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxStubCallerRaw struct {
	Contract *InboxStubCaller // Generic read-only contract binding to access the raw methods on
}

// InboxStubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxStubTransactorRaw struct {
	Contract *InboxStubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInboxStub creates a new instance of InboxStub, bound to a specific deployed contract.
func NewInboxStub(address common.Address, backend bind.ContractBackend) (*InboxStub, error) {
	contract, err := bindInboxStub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InboxStub{InboxStubCaller: InboxStubCaller{contract: contract}, InboxStubTransactor: InboxStubTransactor{contract: contract}, InboxStubFilterer: InboxStubFilterer{contract: contract}}, nil
}

// NewInboxStubCaller creates a new read-only instance of InboxStub, bound to a specific deployed contract.
func NewInboxStubCaller(address common.Address, caller bind.ContractCaller) (*InboxStubCaller, error) {
	contract, err := bindInboxStub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxStubCaller{contract: contract}, nil
}

// NewInboxStubTransactor creates a new write-only instance of InboxStub, bound to a specific deployed contract.
func NewInboxStubTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxStubTransactor, error) {
	contract, err := bindInboxStub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxStubTransactor{contract: contract}, nil
}

// NewInboxStubFilterer creates a new log filterer instance of InboxStub, bound to a specific deployed contract.
func NewInboxStubFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxStubFilterer, error) {
	contract, err := bindInboxStub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxStubFilterer{contract: contract}, nil
}

// bindInboxStub binds a generic wrapper to an already deployed contract.
func bindInboxStub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InboxStubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InboxStub *InboxStubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InboxStub.Contract.InboxStubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InboxStub *InboxStubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxStub.Contract.InboxStubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InboxStub *InboxStubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InboxStub.Contract.InboxStubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InboxStub *InboxStubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InboxStub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InboxStub *InboxStubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxStub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InboxStub *InboxStubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InboxStub.Contract.contract.Transact(opts, method, params...)
}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() pure returns(bool)
func (_InboxStub *InboxStubCaller) AllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "allowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() pure returns(bool)
func (_InboxStub *InboxStubSession) AllowListEnabled() (bool, error) {
	return _InboxStub.Contract.AllowListEnabled(&_InboxStub.CallOpts)
}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() pure returns(bool)
func (_InboxStub *InboxStubCallerSession) AllowListEnabled() (bool, error) {
	return _InboxStub.Contract.AllowListEnabled(&_InboxStub.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_InboxStub *InboxStubCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_InboxStub *InboxStubSession) Bridge() (common.Address, error) {
	return _InboxStub.Contract.Bridge(&_InboxStub.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_InboxStub *InboxStubCallerSession) Bridge() (common.Address, error) {
	return _InboxStub.Contract.Bridge(&_InboxStub.CallOpts)
}

// CalculateRetryableSubmissionFee is a free data retrieval call binding the contract method 0xa66b327d.
//
// Solidity: function calculateRetryableSubmissionFee(uint256 , uint256 ) pure returns(uint256)
func (_InboxStub *InboxStubCaller) CalculateRetryableSubmissionFee(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "calculateRetryableSubmissionFee", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRetryableSubmissionFee is a free data retrieval call binding the contract method 0xa66b327d.
//
// Solidity: function calculateRetryableSubmissionFee(uint256 , uint256 ) pure returns(uint256)
func (_InboxStub *InboxStubSession) CalculateRetryableSubmissionFee(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _InboxStub.Contract.CalculateRetryableSubmissionFee(&_InboxStub.CallOpts, arg0, arg1)
}

// CalculateRetryableSubmissionFee is a free data retrieval call binding the contract method 0xa66b327d.
//
// Solidity: function calculateRetryableSubmissionFee(uint256 , uint256 ) pure returns(uint256)
func (_InboxStub *InboxStubCallerSession) CalculateRetryableSubmissionFee(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _InboxStub.Contract.CalculateRetryableSubmissionFee(&_InboxStub.CallOpts, arg0, arg1)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0x8b3240a0.
//
// Solidity: function getProxyAdmin() pure returns(address)
func (_InboxStub *InboxStubCaller) GetProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "getProxyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAdmin is a free data retrieval call binding the contract method 0x8b3240a0.
//
// Solidity: function getProxyAdmin() pure returns(address)
func (_InboxStub *InboxStubSession) GetProxyAdmin() (common.Address, error) {
	return _InboxStub.Contract.GetProxyAdmin(&_InboxStub.CallOpts)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0x8b3240a0.
//
// Solidity: function getProxyAdmin() pure returns(address)
func (_InboxStub *InboxStubCallerSession) GetProxyAdmin() (common.Address, error) {
	return _InboxStub.Contract.GetProxyAdmin(&_InboxStub.CallOpts)
}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address ) pure returns(bool)
func (_InboxStub *InboxStubCaller) IsAllowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "isAllowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address ) pure returns(bool)
func (_InboxStub *InboxStubSession) IsAllowed(arg0 common.Address) (bool, error) {
	return _InboxStub.Contract.IsAllowed(&_InboxStub.CallOpts, arg0)
}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address ) pure returns(bool)
func (_InboxStub *InboxStubCallerSession) IsAllowed(arg0 common.Address) (bool, error) {
	return _InboxStub.Contract.IsAllowed(&_InboxStub.CallOpts, arg0)
}

// MaxDataSize is a free data retrieval call binding the contract method 0xe8eb1dc3.
//
// Solidity: function maxDataSize() view returns(uint256)
func (_InboxStub *InboxStubCaller) MaxDataSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "maxDataSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDataSize is a free data retrieval call binding the contract method 0xe8eb1dc3.
//
// Solidity: function maxDataSize() view returns(uint256)
func (_InboxStub *InboxStubSession) MaxDataSize() (*big.Int, error) {
	return _InboxStub.Contract.MaxDataSize(&_InboxStub.CallOpts)
}

// MaxDataSize is a free data retrieval call binding the contract method 0xe8eb1dc3.
//
// Solidity: function maxDataSize() view returns(uint256)
func (_InboxStub *InboxStubCallerSession) MaxDataSize() (*big.Int, error) {
	return _InboxStub.Contract.MaxDataSize(&_InboxStub.CallOpts)
}

// Pause is a free data retrieval call binding the contract method 0x8456cb59.
//
// Solidity: function pause() pure returns()
func (_InboxStub *InboxStubCaller) Pause(opts *bind.CallOpts) error {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "pause")

	if err != nil {
		return err
	}

	return err

}

// Pause is a free data retrieval call binding the contract method 0x8456cb59.
//
// Solidity: function pause() pure returns()
func (_InboxStub *InboxStubSession) Pause() error {
	return _InboxStub.Contract.Pause(&_InboxStub.CallOpts)
}

// Pause is a free data retrieval call binding the contract method 0x8456cb59.
//
// Solidity: function pause() pure returns()
func (_InboxStub *InboxStubCallerSession) Pause() error {
	return _InboxStub.Contract.Pause(&_InboxStub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InboxStub *InboxStubCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InboxStub *InboxStubSession) Paused() (bool, error) {
	return _InboxStub.Contract.Paused(&_InboxStub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InboxStub *InboxStubCallerSession) Paused() (bool, error) {
	return _InboxStub.Contract.Paused(&_InboxStub.CallOpts)
}

// SendContractTransaction is a free data retrieval call binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCaller) SendContractTransaction(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*big.Int, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "sendContractTransaction", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SendContractTransaction is a free data retrieval call binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubSession) SendContractTransaction(arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendContractTransaction(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendContractTransaction is a free data retrieval call binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCallerSession) SendContractTransaction(arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendContractTransaction(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendUnsignedTransaction is a free data retrieval call binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCaller) SendUnsignedTransaction(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "sendUnsignedTransaction", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SendUnsignedTransaction is a free data retrieval call binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubSession) SendUnsignedTransaction(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendUnsignedTransaction(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// SendUnsignedTransaction is a free data retrieval call binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCallerSession) SendUnsignedTransaction(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendUnsignedTransaction(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// SendUnsignedTransactionToFork is a free data retrieval call binding the contract method 0x00f72382.
//
// Solidity: function sendUnsignedTransactionToFork(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCaller) SendUnsignedTransactionToFork(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "sendUnsignedTransactionToFork", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SendUnsignedTransactionToFork is a free data retrieval call binding the contract method 0x00f72382.
//
// Solidity: function sendUnsignedTransactionToFork(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubSession) SendUnsignedTransactionToFork(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendUnsignedTransactionToFork(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// SendUnsignedTransactionToFork is a free data retrieval call binding the contract method 0x00f72382.
//
// Solidity: function sendUnsignedTransactionToFork(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCallerSession) SendUnsignedTransactionToFork(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendUnsignedTransactionToFork(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// SendWithdrawEthToFork is a free data retrieval call binding the contract method 0x70665f14.
//
// Solidity: function sendWithdrawEthToFork(uint256 , uint256 , uint256 , uint256 , address ) pure returns(uint256)
func (_InboxStub *InboxStubCaller) SendWithdrawEthToFork(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "sendWithdrawEthToFork", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SendWithdrawEthToFork is a free data retrieval call binding the contract method 0x70665f14.
//
// Solidity: function sendWithdrawEthToFork(uint256 , uint256 , uint256 , uint256 , address ) pure returns(uint256)
func (_InboxStub *InboxStubSession) SendWithdrawEthToFork(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address) (*big.Int, error) {
	return _InboxStub.Contract.SendWithdrawEthToFork(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendWithdrawEthToFork is a free data retrieval call binding the contract method 0x70665f14.
//
// Solidity: function sendWithdrawEthToFork(uint256 , uint256 , uint256 , uint256 , address ) pure returns(uint256)
func (_InboxStub *InboxStubCallerSession) SendWithdrawEthToFork(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address) (*big.Int, error) {
	return _InboxStub.Contract.SendWithdrawEthToFork(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_InboxStub *InboxStubCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_InboxStub *InboxStubSession) SequencerInbox() (common.Address, error) {
	return _InboxStub.Contract.SequencerInbox(&_InboxStub.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_InboxStub *InboxStubCallerSession) SequencerInbox() (common.Address, error) {
	return _InboxStub.Contract.SequencerInbox(&_InboxStub.CallOpts)
}

// SetAllowList is a free data retrieval call binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] , bool[] ) pure returns()
func (_InboxStub *InboxStubCaller) SetAllowList(opts *bind.CallOpts, arg0 []common.Address, arg1 []bool) error {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "setAllowList", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// SetAllowList is a free data retrieval call binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] , bool[] ) pure returns()
func (_InboxStub *InboxStubSession) SetAllowList(arg0 []common.Address, arg1 []bool) error {
	return _InboxStub.Contract.SetAllowList(&_InboxStub.CallOpts, arg0, arg1)
}

// SetAllowList is a free data retrieval call binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] , bool[] ) pure returns()
func (_InboxStub *InboxStubCallerSession) SetAllowList(arg0 []common.Address, arg1 []bool) error {
	return _InboxStub.Contract.SetAllowList(&_InboxStub.CallOpts, arg0, arg1)
}

// SetAllowListEnabled is a free data retrieval call binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool ) pure returns()
func (_InboxStub *InboxStubCaller) SetAllowListEnabled(opts *bind.CallOpts, arg0 bool) error {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "setAllowListEnabled", arg0)

	if err != nil {
		return err
	}

	return err

}

// SetAllowListEnabled is a free data retrieval call binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool ) pure returns()
func (_InboxStub *InboxStubSession) SetAllowListEnabled(arg0 bool) error {
	return _InboxStub.Contract.SetAllowListEnabled(&_InboxStub.CallOpts, arg0)
}

// SetAllowListEnabled is a free data retrieval call binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool ) pure returns()
func (_InboxStub *InboxStubCallerSession) SetAllowListEnabled(arg0 bool) error {
	return _InboxStub.Contract.SetAllowListEnabled(&_InboxStub.CallOpts, arg0)
}

// Unpause is a free data retrieval call binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() pure returns()
func (_InboxStub *InboxStubCaller) Unpause(opts *bind.CallOpts) error {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "unpause")

	if err != nil {
		return err
	}

	return err

}

// Unpause is a free data retrieval call binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() pure returns()
func (_InboxStub *InboxStubSession) Unpause() error {
	return _InboxStub.Contract.Unpause(&_InboxStub.CallOpts)
}

// Unpause is a free data retrieval call binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() pure returns()
func (_InboxStub *InboxStubCallerSession) Unpause() error {
	return _InboxStub.Contract.Unpause(&_InboxStub.CallOpts)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) CreateRetryableTicket(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "createRetryableTicket", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubSession) CreateRetryableTicket(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.CreateRetryableTicket(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) CreateRetryableTicket(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.CreateRetryableTicket(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_InboxStub *InboxStubTransactor) DepositEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "depositEth")
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_InboxStub *InboxStubSession) DepositEth() (*types.Transaction, error) {
	return _InboxStub.Contract.DepositEth(&_InboxStub.TransactOpts)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) DepositEth() (*types.Transaction, error) {
	return _InboxStub.Contract.DepositEth(&_InboxStub.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address ) returns()
func (_InboxStub *InboxStubTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "initialize", _bridge, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address ) returns()
func (_InboxStub *InboxStubSession) Initialize(_bridge common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _InboxStub.Contract.Initialize(&_InboxStub.TransactOpts, _bridge, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address ) returns()
func (_InboxStub *InboxStubTransactorSession) Initialize(_bridge common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _InboxStub.Contract.Initialize(&_InboxStub.TransactOpts, _bridge, arg1)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_InboxStub *InboxStubTransactor) PostUpgradeInit(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "postUpgradeInit", _bridge)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_InboxStub *InboxStubSession) PostUpgradeInit(_bridge common.Address) (*types.Transaction, error) {
	return _InboxStub.Contract.PostUpgradeInit(&_InboxStub.TransactOpts, _bridge)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_InboxStub *InboxStubTransactorSession) PostUpgradeInit(_bridge common.Address) (*types.Transaction, error) {
	return _InboxStub.Contract.PostUpgradeInit(&_InboxStub.TransactOpts, _bridge)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) SendL1FundedContractTransaction(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "sendL1FundedContractTransaction", arg0, arg1, arg2, arg3)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubSession) SendL1FundedContractTransaction(arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedContractTransaction(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) SendL1FundedContractTransaction(arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedContractTransaction(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) SendL1FundedUnsignedTransaction(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "sendL1FundedUnsignedTransaction", arg0, arg1, arg2, arg3, arg4)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubSession) SendL1FundedUnsignedTransaction(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedUnsignedTransaction(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) SendL1FundedUnsignedTransaction(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedUnsignedTransaction(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendL1FundedUnsignedTransactionToFork is a paid mutator transaction binding the contract method 0xe6bd12cf.
//
// Solidity: function sendL1FundedUnsignedTransactionToFork(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) SendL1FundedUnsignedTransactionToFork(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "sendL1FundedUnsignedTransactionToFork", arg0, arg1, arg2, arg3, arg4)
}

// SendL1FundedUnsignedTransactionToFork is a paid mutator transaction binding the contract method 0xe6bd12cf.
//
// Solidity: function sendL1FundedUnsignedTransactionToFork(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubSession) SendL1FundedUnsignedTransactionToFork(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedUnsignedTransactionToFork(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendL1FundedUnsignedTransactionToFork is a paid mutator transaction binding the contract method 0xe6bd12cf.
//
// Solidity: function sendL1FundedUnsignedTransactionToFork(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) SendL1FundedUnsignedTransactionToFork(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedUnsignedTransactionToFork(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL2Message(&_InboxStub.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL2Message(&_InboxStub.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubTransactor) SendL2MessageFromOrigin(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "sendL2MessageFromOrigin", messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL2MessageFromOrigin(&_InboxStub.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubTransactorSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL2MessageFromOrigin(&_InboxStub.TransactOpts, messageData)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) UnsafeCreateRetryableTicket(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "unsafeCreateRetryableTicket", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubSession) UnsafeCreateRetryableTicket(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.UnsafeCreateRetryableTicket(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) UnsafeCreateRetryableTicket(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.UnsafeCreateRetryableTicket(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// InboxStubInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the InboxStub contract.
type InboxStubInboxMessageDeliveredIterator struct {
	Event *InboxStubInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *InboxStubInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxStubInboxMessageDelivered)
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
		it.Event = new(InboxStubInboxMessageDelivered)
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
func (it *InboxStubInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxStubInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxStubInboxMessageDelivered represents a InboxMessageDelivered event raised by the InboxStub contract.
type InboxStubInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_InboxStub *InboxStubFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxStubInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _InboxStub.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxStubInboxMessageDeliveredIterator{contract: _InboxStub.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_InboxStub *InboxStubFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *InboxStubInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _InboxStub.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxStubInboxMessageDelivered)
				if err := _InboxStub.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_InboxStub *InboxStubFilterer) ParseInboxMessageDelivered(log types.Log) (*InboxStubInboxMessageDelivered, error) {
	event := new(InboxStubInboxMessageDelivered)
	if err := _InboxStub.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxStubInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the InboxStub contract.
type InboxStubInboxMessageDeliveredFromOriginIterator struct {
	Event *InboxStubInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *InboxStubInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxStubInboxMessageDeliveredFromOrigin)
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
		it.Event = new(InboxStubInboxMessageDeliveredFromOrigin)
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
func (it *InboxStubInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxStubInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxStubInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the InboxStub contract.
type InboxStubInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_InboxStub *InboxStubFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxStubInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _InboxStub.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxStubInboxMessageDeliveredFromOriginIterator{contract: _InboxStub.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_InboxStub *InboxStubFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *InboxStubInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _InboxStub.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxStubInboxMessageDeliveredFromOrigin)
				if err := _InboxStub.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_InboxStub *InboxStubFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*InboxStubInboxMessageDeliveredFromOrigin, error) {
	event := new(InboxStubInboxMessageDeliveredFromOrigin)
	if err := _InboxStub.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleTreeAccessMetaData contains all meta data concerning the MerkleTreeAccess contract.
var MerkleTreeAccessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxProofLength\",\"type\":\"uint256\"}],\"name\":\"MerkleProofTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"me\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subtreeRoot\",\"type\":\"bytes32\"}],\"name\":\"appendCompleteSubTree\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"me\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"name\":\"appendLeaf\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"leastSignificantBit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endSize\",\"type\":\"uint256\"}],\"name\":\"maximumAppendBetween\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"mostSignificantBit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"me\",\"type\":\"bytes32[]\"}],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"verifyInclusionProof\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"preRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"postRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"postSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"preExpansion\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"verifyPrefixProof\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506116b8806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063c22c47a41161005b578063c22c47a4146100ff578063ca11325314610112578063d230d23f14610125578063e6bcbc651461013857600080fd5b80635fb9c3d41461008d57806367905a7e146100a25780636bd58993146100cb578063bc2f0640146100de575b600080fd5b6100a061009b3660046112de565b61014b565b005b6100b56100b0366004611367565b610161565b6040516100c291906113b5565b60405180910390f35b6100a06100d93660046113f9565b610178565b6100f16100ec366004611453565b61018a565b6040519081526020016100c2565b6100b561010d366004611475565b61019f565b6100f16101203660046114ba565b6101ab565b6100f16101333660046114f7565b6101b6565b6100f16101463660046114f7565b6101c1565b6101598686868686866101cc565b505050505050565b606061016e8484846104f4565b90505b9392505050565b61018484848484610a7e565b50505050565b60006101968383610b0b565b90505b92915050565b60606101968383610c00565b600061019982610c36565b600061019982610dd6565b600061019982610e3f565b600085116102215760405162461bcd60e51b815260206004820152601460248201527f5072652d73697a652063616e6e6f74206265203000000000000000000000000060448201526064015b60405180910390fd5b8561022b83610c36565b146102785760405162461bcd60e51b815260206004820152601b60248201527f50726520657870616e73696f6e20726f6f74206d69736d6174636800000000006044820152606401610218565b8461028283610f85565b146102f55760405162461bcd60e51b815260206004820152602160248201527f5072652073697a6520646f6573206e6f74206d6174636820657870616e73696f60448201527f6e000000000000000000000000000000000000000000000000000000000000006064820152608401610218565b8285106103445760405162461bcd60e51b815260206004820181905260248201527f5072652073697a65206e6f74206c657373207468616e20706f73742073697a656044820152606401610218565b60008590506000806103598560008751610fe0565b90505b8583101561041c5760006103708488610b0b565b9050845183106103c25760405162461bcd60e51b815260206004820152601260248201527f496e646578206f7574206f662072616e676500000000000000000000000000006044820152606401610218565b6103e682828786815181106103d9576103d9611510565b60200260200101516104f4565b91506001811b6103f6818661153c565b9450878511156104085761040861154f565b8361041281611565565b945050505061035c565b8661042682610c36565b146104995760405162461bcd60e51b815260206004820152602260248201527f506f737420657870616e73696f6e20726f6f74206e6f7420657175616c20706f60448201527f73740000000000000000000000000000000000000000000000000000000000006064820152608401610218565b835182146104e95760405162461bcd60e51b815260206004820152601660248201527f496e636f6d706c6574652070726f6f66207573616765000000000000000000006044820152606401610218565b505050505050505050565b6060604083106105465760405162461bcd60e51b815260206004820152600e60248201527f4c6576656c20746f6f20686967680000000000000000000000000000000000006044820152606401610218565b60008290036105975760405162461bcd60e51b815260206004820152601b60248201527f43616e6e6f7420617070656e6420656d707479207375627472656500000000006044820152606401610218565b6040845111156105e95760405162461bcd60e51b815260206004820152601a60248201527f4d65726b6c6520657870616e73696f6e20746f6f206c617267650000000000006044820152606401610218565b83516000036106685760006105ff84600161153c565b67ffffffffffffffff8111156106175761061761121a565b604051908082528060200260200182016040528015610640578160200160208202803683370190505b5090508281858151811061065657610656611510565b60209081029190910101529050610171565b835183106106de5760405162461bcd60e51b815260206004820152603560248201527f4c6576656c2067726561746572207468616e2068696768657374206c6576656c60448201527f206f662063757272656e7420657870616e73696f6e00000000000000000000006064820152608401610218565b8160006106ea86610f85565b905060006106f9866002611663565b610703908361153c565b9050600061071083610e3f565b61071983610e3f565b1161076757875167ffffffffffffffff8111156107385761073861121a565b604051908082528060200260200182016040528015610761578160200160208202803683370190505b506107b7565b875161077490600161153c565b67ffffffffffffffff81111561078c5761078c61121a565b6040519080825280602002602001820160405280156107b5578160200160208202803683370190505b505b905060408151111561080b5760405162461bcd60e51b815260206004820152601c60248201527f417070656e642063726561746573206f76657273697a652074726565000000006044820152606401610218565b60005b88518110156109c757878110156108b55788818151811061083157610831611510565b60200260200101516000801b146108b05760405162461bcd60e51b815260206004820152602260248201527f417070656e642061626f7665206c65617374207369676e69666963616e74206260448201527f69740000000000000000000000000000000000000000000000000000000000006064820152608401610218565b6109b5565b60008590036108fb578881815181106108d0576108d0611510565b60200260200101518282815181106108ea576108ea611510565b6020026020010181815250506109b5565b88818151811061090d5761090d611510565b60200260200101516000801b03610945578482828151811061093157610931611510565b6020908102919091010152600094506109b5565b6000801b82828151811061095b5761095b611510565b60200260200101818152505088818151811061097957610979611510565b60200260200101518560405160200161099c929190918252602082015260400190565b6040516020818303038152906040528051906020012094505b806109bf81611565565b91505061080e565b5083156109fb578381600183516109de919061166f565b815181106109ee576109ee611510565b6020026020010181815250505b8060018251610a0a919061166f565b81518110610a1a57610a1a611510565b60200260200101516000801b03610a735760405162461bcd60e51b815260206004820152600f60248201527f4c61737420656e747279207a65726f00000000000000000000000000000000006044820152606401610218565b979650505050505050565b6000610ab3828486604051602001610a9891815260200190565b6040516020818303038152906040528051906020012061115f565b9050808514610b045760405162461bcd60e51b815260206004820152601760248201527f496e76616c696420696e636c7573696f6e2070726f6f660000000000000000006044820152606401610218565b5050505050565b6000818310610b5c5760405162461bcd60e51b815260206004820152601760248201527f5374617274206e6f74206c657373207468616e20656e640000000000000000006044820152606401610218565b6000610b69838518610e3f565b905060006001610b79838261153c565b6001901b610b87919061166f565b90508481168482168115610ba957610b9e82610dd6565b945050505050610199565b8015610bb857610b9e81610e3f565b60405162461bcd60e51b815260206004820152601b60248201527f426f7468207920616e64207a2063616e6e6f74206265207a65726f00000000006044820152606401610218565b606061019683600084604051602001610c1b91815260200190565b604051602081830303815290604052805190602001206104f4565b600080825111610c885760405162461bcd60e51b815260206004820152601660248201527f456d707479206d65726b6c6520657870616e73696f6e000000000000000000006044820152606401610218565b604082511115610cda5760405162461bcd60e51b815260206004820152601a60248201527f4d65726b6c6520657870616e73696f6e20746f6f206c617267650000000000006044820152606401610218565b6000805b8351811015610dcf576000848281518110610cfb57610cfb611510565b60200260200101519050826000801b03610d67578015610d625780925060018551610d26919061166f565b8214610d6257604051610d49908490600090602001918252602082015260400190565b6040516020818303038152906040528051906020012092505b610dbc565b8015610d86576040805160208101839052908101849052606001610d49565b604051610da3908490600090602001918252602082015260400190565b6040516020818303038152906040528051906020012092505b5080610dc781611565565b915050610cde565b5092915050565b6000808211610e275760405162461bcd60e51b815260206004820152601c60248201527f5a65726f20686173206e6f207369676e69666963616e742062697473000000006044820152606401610218565b60008280610e3660018261166f565b16189050610171815b600081600003610e915760405162461bcd60e51b815260206004820152601c60248201527f5a65726f20686173206e6f207369676e69666963616e742062697473000000006044820152606401610218565b7001000000000000000000000000000000008210610ebc57608091821c91610eb9908261153c565b90505b680100000000000000008210610edf57604091821c91610edc908261153c565b90505b6401000000008210610efe57602091821c91610efb908261153c565b90505b620100008210610f1b57601091821c91610f18908261153c565b90505b6101008210610f3757600891821c91610f34908261153c565b90505b60108210610f5257600491821c91610f4f908261153c565b90505b60048210610f6d57600291821c91610f6a908261153c565b90505b60028210610f805761019960018261153c565b919050565b600080805b8351811015610dcf57838181518110610fa557610fa5611510565b60200260200101516000801b14610fce57610fc1816002611663565b610fcb908361153c565b91505b80610fd881611565565b915050610f8a565b60608183106110315760405162461bcd60e51b815260206004820152601760248201527f5374617274206e6f74206c657373207468616e20656e640000000000000000006044820152606401610218565b83518211156110a85760405162461bcd60e51b815260206004820152602160248201527f456e64206e6f74206c657373206f7220657175616c207468616e206c656e677460448201527f68000000000000000000000000000000000000000000000000000000000000006064820152608401610218565b60006110b4848461166f565b67ffffffffffffffff8111156110cc576110cc61121a565b6040519080825280602002602001820160405280156110f5578160200160208202803683370190505b509050835b838110156111565785818151811061111457611114611510565b6020026020010151828683611129919061166f565b8151811061113957611139611510565b60209081029190910101528061114e81611565565b9150506110fa565b50949350505050565b82516000906101008111156111ab576040517ffdac331e000000000000000000000000000000000000000000000000000000008152600481018290526101006024820152604401610218565b8260005b828110156112105760008782815181106111cb576111cb611510565b60200260200101519050816001901b87166000036111f757826000528060205260406000209250611207565b8060005282602052604060002092505b506001016111af565b5095945050505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261124157600080fd5b8135602067ffffffffffffffff8083111561125e5761125e61121a565b8260051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f830116810181811084821117156112a1576112a161121a565b6040529384528581018301938381019250878511156112bf57600080fd5b83870191505b84821015610a73578135835291830191908301906112c5565b60008060008060008060c087890312156112f757600080fd5b86359550602087013594506040870135935060608701359250608087013567ffffffffffffffff8082111561132b57600080fd5b6113378a838b01611230565b935060a089013591508082111561134d57600080fd5b5061135a89828a01611230565b9150509295509295509295565b60008060006060848603121561137c57600080fd5b833567ffffffffffffffff81111561139357600080fd5b61139f86828701611230565b9660208601359650604090950135949350505050565b6020808252825182820181905260009190848201906040850190845b818110156113ed578351835292840192918401916001016113d1565b50909695505050505050565b6000806000806080858703121561140f57600080fd5b843593506020850135925060408501359150606085013567ffffffffffffffff81111561143b57600080fd5b61144787828801611230565b91505092959194509250565b6000806040838503121561146657600080fd5b50508035926020909101359150565b6000806040838503121561148857600080fd5b823567ffffffffffffffff81111561149f57600080fd5b6114ab85828601611230565b95602094909401359450505050565b6000602082840312156114cc57600080fd5b813567ffffffffffffffff8111156114e357600080fd5b6114ef84828501611230565b949350505050565b60006020828403121561150957600080fd5b5035919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b8082018082111561019957610199611526565b634e487b7160e01b600052600160045260246000fd5b6000600019820361157857611578611526565b5060010190565b600181815b808511156115ba5781600019048211156115a0576115a0611526565b808516156115ad57918102915b93841c9390800290611584565b509250929050565b6000826115d157506001610199565b816115de57506000610199565b81600181146115f457600281146115fe5761161a565b6001915050610199565b60ff84111561160f5761160f611526565b50506001821b610199565b5060208310610133831016604e8410600b841016171561163d575081810a610199565b611647838361157f565b806000190482111561165b5761165b611526565b029392505050565b600061019683836115c2565b818103818111156101995761019961152656fea26469706673582212202b0e9f2f912577079fd7d6f7b774ab5bdd0efa3fde13a34fb2b5e98fcfaf2c4c64736f6c63430008110033",
}

// MerkleTreeAccessABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleTreeAccessMetaData.ABI instead.
var MerkleTreeAccessABI = MerkleTreeAccessMetaData.ABI

// MerkleTreeAccessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleTreeAccessMetaData.Bin instead.
var MerkleTreeAccessBin = MerkleTreeAccessMetaData.Bin

// DeployMerkleTreeAccess deploys a new Ethereum contract, binding an instance of MerkleTreeAccess to it.
func DeployMerkleTreeAccess(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleTreeAccess, error) {
	parsed, err := MerkleTreeAccessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleTreeAccessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleTreeAccess{MerkleTreeAccessCaller: MerkleTreeAccessCaller{contract: contract}, MerkleTreeAccessTransactor: MerkleTreeAccessTransactor{contract: contract}, MerkleTreeAccessFilterer: MerkleTreeAccessFilterer{contract: contract}}, nil
}

// MerkleTreeAccess is an auto generated Go binding around an Ethereum contract.
type MerkleTreeAccess struct {
	MerkleTreeAccessCaller     // Read-only binding to the contract
	MerkleTreeAccessTransactor // Write-only binding to the contract
	MerkleTreeAccessFilterer   // Log filterer for contract events
}

// MerkleTreeAccessCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleTreeAccessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeAccessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTreeAccessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeAccessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleTreeAccessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeAccessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleTreeAccessSession struct {
	Contract     *MerkleTreeAccess // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleTreeAccessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleTreeAccessCallerSession struct {
	Contract *MerkleTreeAccessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MerkleTreeAccessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTreeAccessTransactorSession struct {
	Contract     *MerkleTreeAccessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MerkleTreeAccessRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleTreeAccessRaw struct {
	Contract *MerkleTreeAccess // Generic contract binding to access the raw methods on
}

// MerkleTreeAccessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleTreeAccessCallerRaw struct {
	Contract *MerkleTreeAccessCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleTreeAccessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTreeAccessTransactorRaw struct {
	Contract *MerkleTreeAccessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleTreeAccess creates a new instance of MerkleTreeAccess, bound to a specific deployed contract.
func NewMerkleTreeAccess(address common.Address, backend bind.ContractBackend) (*MerkleTreeAccess, error) {
	contract, err := bindMerkleTreeAccess(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeAccess{MerkleTreeAccessCaller: MerkleTreeAccessCaller{contract: contract}, MerkleTreeAccessTransactor: MerkleTreeAccessTransactor{contract: contract}, MerkleTreeAccessFilterer: MerkleTreeAccessFilterer{contract: contract}}, nil
}

// NewMerkleTreeAccessCaller creates a new read-only instance of MerkleTreeAccess, bound to a specific deployed contract.
func NewMerkleTreeAccessCaller(address common.Address, caller bind.ContractCaller) (*MerkleTreeAccessCaller, error) {
	contract, err := bindMerkleTreeAccess(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeAccessCaller{contract: contract}, nil
}

// NewMerkleTreeAccessTransactor creates a new write-only instance of MerkleTreeAccess, bound to a specific deployed contract.
func NewMerkleTreeAccessTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTreeAccessTransactor, error) {
	contract, err := bindMerkleTreeAccess(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeAccessTransactor{contract: contract}, nil
}

// NewMerkleTreeAccessFilterer creates a new log filterer instance of MerkleTreeAccess, bound to a specific deployed contract.
func NewMerkleTreeAccessFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleTreeAccessFilterer, error) {
	contract, err := bindMerkleTreeAccess(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeAccessFilterer{contract: contract}, nil
}

// bindMerkleTreeAccess binds a generic wrapper to an already deployed contract.
func bindMerkleTreeAccess(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MerkleTreeAccessMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTreeAccess *MerkleTreeAccessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTreeAccess.Contract.MerkleTreeAccessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTreeAccess *MerkleTreeAccessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTreeAccess.Contract.MerkleTreeAccessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTreeAccess *MerkleTreeAccessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTreeAccess.Contract.MerkleTreeAccessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTreeAccess *MerkleTreeAccessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTreeAccess.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTreeAccess *MerkleTreeAccessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTreeAccess.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTreeAccess *MerkleTreeAccessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTreeAccess.Contract.contract.Transact(opts, method, params...)
}

// AppendCompleteSubTree is a free data retrieval call binding the contract method 0x67905a7e.
//
// Solidity: function appendCompleteSubTree(bytes32[] me, uint256 level, bytes32 subtreeRoot) pure returns(bytes32[])
func (_MerkleTreeAccess *MerkleTreeAccessCaller) AppendCompleteSubTree(opts *bind.CallOpts, me [][32]byte, level *big.Int, subtreeRoot [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _MerkleTreeAccess.contract.Call(opts, &out, "appendCompleteSubTree", me, level, subtreeRoot)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// AppendCompleteSubTree is a free data retrieval call binding the contract method 0x67905a7e.
//
// Solidity: function appendCompleteSubTree(bytes32[] me, uint256 level, bytes32 subtreeRoot) pure returns(bytes32[])
func (_MerkleTreeAccess *MerkleTreeAccessSession) AppendCompleteSubTree(me [][32]byte, level *big.Int, subtreeRoot [32]byte) ([][32]byte, error) {
	return _MerkleTreeAccess.Contract.AppendCompleteSubTree(&_MerkleTreeAccess.CallOpts, me, level, subtreeRoot)
}

// AppendCompleteSubTree is a free data retrieval call binding the contract method 0x67905a7e.
//
// Solidity: function appendCompleteSubTree(bytes32[] me, uint256 level, bytes32 subtreeRoot) pure returns(bytes32[])
func (_MerkleTreeAccess *MerkleTreeAccessCallerSession) AppendCompleteSubTree(me [][32]byte, level *big.Int, subtreeRoot [32]byte) ([][32]byte, error) {
	return _MerkleTreeAccess.Contract.AppendCompleteSubTree(&_MerkleTreeAccess.CallOpts, me, level, subtreeRoot)
}

// AppendLeaf is a free data retrieval call binding the contract method 0xc22c47a4.
//
// Solidity: function appendLeaf(bytes32[] me, bytes32 leaf) pure returns(bytes32[])
func (_MerkleTreeAccess *MerkleTreeAccessCaller) AppendLeaf(opts *bind.CallOpts, me [][32]byte, leaf [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _MerkleTreeAccess.contract.Call(opts, &out, "appendLeaf", me, leaf)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// AppendLeaf is a free data retrieval call binding the contract method 0xc22c47a4.
//
// Solidity: function appendLeaf(bytes32[] me, bytes32 leaf) pure returns(bytes32[])
func (_MerkleTreeAccess *MerkleTreeAccessSession) AppendLeaf(me [][32]byte, leaf [32]byte) ([][32]byte, error) {
	return _MerkleTreeAccess.Contract.AppendLeaf(&_MerkleTreeAccess.CallOpts, me, leaf)
}

// AppendLeaf is a free data retrieval call binding the contract method 0xc22c47a4.
//
// Solidity: function appendLeaf(bytes32[] me, bytes32 leaf) pure returns(bytes32[])
func (_MerkleTreeAccess *MerkleTreeAccessCallerSession) AppendLeaf(me [][32]byte, leaf [32]byte) ([][32]byte, error) {
	return _MerkleTreeAccess.Contract.AppendLeaf(&_MerkleTreeAccess.CallOpts, me, leaf)
}

// LeastSignificantBit is a free data retrieval call binding the contract method 0xd230d23f.
//
// Solidity: function leastSignificantBit(uint256 x) pure returns(uint256)
func (_MerkleTreeAccess *MerkleTreeAccessCaller) LeastSignificantBit(opts *bind.CallOpts, x *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTreeAccess.contract.Call(opts, &out, "leastSignificantBit", x)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LeastSignificantBit is a free data retrieval call binding the contract method 0xd230d23f.
//
// Solidity: function leastSignificantBit(uint256 x) pure returns(uint256)
func (_MerkleTreeAccess *MerkleTreeAccessSession) LeastSignificantBit(x *big.Int) (*big.Int, error) {
	return _MerkleTreeAccess.Contract.LeastSignificantBit(&_MerkleTreeAccess.CallOpts, x)
}

// LeastSignificantBit is a free data retrieval call binding the contract method 0xd230d23f.
//
// Solidity: function leastSignificantBit(uint256 x) pure returns(uint256)
func (_MerkleTreeAccess *MerkleTreeAccessCallerSession) LeastSignificantBit(x *big.Int) (*big.Int, error) {
	return _MerkleTreeAccess.Contract.LeastSignificantBit(&_MerkleTreeAccess.CallOpts, x)
}

// MaximumAppendBetween is a free data retrieval call binding the contract method 0xbc2f0640.
//
// Solidity: function maximumAppendBetween(uint256 startSize, uint256 endSize) pure returns(uint256)
func (_MerkleTreeAccess *MerkleTreeAccessCaller) MaximumAppendBetween(opts *bind.CallOpts, startSize *big.Int, endSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTreeAccess.contract.Call(opts, &out, "maximumAppendBetween", startSize, endSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumAppendBetween is a free data retrieval call binding the contract method 0xbc2f0640.
//
// Solidity: function maximumAppendBetween(uint256 startSize, uint256 endSize) pure returns(uint256)
func (_MerkleTreeAccess *MerkleTreeAccessSession) MaximumAppendBetween(startSize *big.Int, endSize *big.Int) (*big.Int, error) {
	return _MerkleTreeAccess.Contract.MaximumAppendBetween(&_MerkleTreeAccess.CallOpts, startSize, endSize)
}

// MaximumAppendBetween is a free data retrieval call binding the contract method 0xbc2f0640.
//
// Solidity: function maximumAppendBetween(uint256 startSize, uint256 endSize) pure returns(uint256)
func (_MerkleTreeAccess *MerkleTreeAccessCallerSession) MaximumAppendBetween(startSize *big.Int, endSize *big.Int) (*big.Int, error) {
	return _MerkleTreeAccess.Contract.MaximumAppendBetween(&_MerkleTreeAccess.CallOpts, startSize, endSize)
}

// MostSignificantBit is a free data retrieval call binding the contract method 0xe6bcbc65.
//
// Solidity: function mostSignificantBit(uint256 x) pure returns(uint256)
func (_MerkleTreeAccess *MerkleTreeAccessCaller) MostSignificantBit(opts *bind.CallOpts, x *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTreeAccess.contract.Call(opts, &out, "mostSignificantBit", x)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MostSignificantBit is a free data retrieval call binding the contract method 0xe6bcbc65.
//
// Solidity: function mostSignificantBit(uint256 x) pure returns(uint256)
func (_MerkleTreeAccess *MerkleTreeAccessSession) MostSignificantBit(x *big.Int) (*big.Int, error) {
	return _MerkleTreeAccess.Contract.MostSignificantBit(&_MerkleTreeAccess.CallOpts, x)
}

// MostSignificantBit is a free data retrieval call binding the contract method 0xe6bcbc65.
//
// Solidity: function mostSignificantBit(uint256 x) pure returns(uint256)
func (_MerkleTreeAccess *MerkleTreeAccessCallerSession) MostSignificantBit(x *big.Int) (*big.Int, error) {
	return _MerkleTreeAccess.Contract.MostSignificantBit(&_MerkleTreeAccess.CallOpts, x)
}

// Root is a free data retrieval call binding the contract method 0xca113253.
//
// Solidity: function root(bytes32[] me) pure returns(bytes32)
func (_MerkleTreeAccess *MerkleTreeAccessCaller) Root(opts *bind.CallOpts, me [][32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MerkleTreeAccess.contract.Call(opts, &out, "root", me)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xca113253.
//
// Solidity: function root(bytes32[] me) pure returns(bytes32)
func (_MerkleTreeAccess *MerkleTreeAccessSession) Root(me [][32]byte) ([32]byte, error) {
	return _MerkleTreeAccess.Contract.Root(&_MerkleTreeAccess.CallOpts, me)
}

// Root is a free data retrieval call binding the contract method 0xca113253.
//
// Solidity: function root(bytes32[] me) pure returns(bytes32)
func (_MerkleTreeAccess *MerkleTreeAccessCallerSession) Root(me [][32]byte) ([32]byte, error) {
	return _MerkleTreeAccess.Contract.Root(&_MerkleTreeAccess.CallOpts, me)
}

// VerifyInclusionProof is a free data retrieval call binding the contract method 0x6bd58993.
//
// Solidity: function verifyInclusionProof(bytes32 rootHash, bytes32 leaf, uint256 index, bytes32[] proof) pure returns()
func (_MerkleTreeAccess *MerkleTreeAccessCaller) VerifyInclusionProof(opts *bind.CallOpts, rootHash [32]byte, leaf [32]byte, index *big.Int, proof [][32]byte) error {
	var out []interface{}
	err := _MerkleTreeAccess.contract.Call(opts, &out, "verifyInclusionProof", rootHash, leaf, index, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyInclusionProof is a free data retrieval call binding the contract method 0x6bd58993.
//
// Solidity: function verifyInclusionProof(bytes32 rootHash, bytes32 leaf, uint256 index, bytes32[] proof) pure returns()
func (_MerkleTreeAccess *MerkleTreeAccessSession) VerifyInclusionProof(rootHash [32]byte, leaf [32]byte, index *big.Int, proof [][32]byte) error {
	return _MerkleTreeAccess.Contract.VerifyInclusionProof(&_MerkleTreeAccess.CallOpts, rootHash, leaf, index, proof)
}

// VerifyInclusionProof is a free data retrieval call binding the contract method 0x6bd58993.
//
// Solidity: function verifyInclusionProof(bytes32 rootHash, bytes32 leaf, uint256 index, bytes32[] proof) pure returns()
func (_MerkleTreeAccess *MerkleTreeAccessCallerSession) VerifyInclusionProof(rootHash [32]byte, leaf [32]byte, index *big.Int, proof [][32]byte) error {
	return _MerkleTreeAccess.Contract.VerifyInclusionProof(&_MerkleTreeAccess.CallOpts, rootHash, leaf, index, proof)
}

// VerifyPrefixProof is a free data retrieval call binding the contract method 0x5fb9c3d4.
//
// Solidity: function verifyPrefixProof(bytes32 preRoot, uint256 preSize, bytes32 postRoot, uint256 postSize, bytes32[] preExpansion, bytes32[] proof) pure returns()
func (_MerkleTreeAccess *MerkleTreeAccessCaller) VerifyPrefixProof(opts *bind.CallOpts, preRoot [32]byte, preSize *big.Int, postRoot [32]byte, postSize *big.Int, preExpansion [][32]byte, proof [][32]byte) error {
	var out []interface{}
	err := _MerkleTreeAccess.contract.Call(opts, &out, "verifyPrefixProof", preRoot, preSize, postRoot, postSize, preExpansion, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyPrefixProof is a free data retrieval call binding the contract method 0x5fb9c3d4.
//
// Solidity: function verifyPrefixProof(bytes32 preRoot, uint256 preSize, bytes32 postRoot, uint256 postSize, bytes32[] preExpansion, bytes32[] proof) pure returns()
func (_MerkleTreeAccess *MerkleTreeAccessSession) VerifyPrefixProof(preRoot [32]byte, preSize *big.Int, postRoot [32]byte, postSize *big.Int, preExpansion [][32]byte, proof [][32]byte) error {
	return _MerkleTreeAccess.Contract.VerifyPrefixProof(&_MerkleTreeAccess.CallOpts, preRoot, preSize, postRoot, postSize, preExpansion, proof)
}

// VerifyPrefixProof is a free data retrieval call binding the contract method 0x5fb9c3d4.
//
// Solidity: function verifyPrefixProof(bytes32 preRoot, uint256 preSize, bytes32 postRoot, uint256 postSize, bytes32[] preExpansion, bytes32[] proof) pure returns()
func (_MerkleTreeAccess *MerkleTreeAccessCallerSession) VerifyPrefixProof(preRoot [32]byte, preSize *big.Int, postRoot [32]byte, postSize *big.Int, preExpansion [][32]byte, proof [][32]byte) error {
	return _MerkleTreeAccess.Contract.VerifyPrefixProof(&_MerkleTreeAccess.CallOpts, preRoot, preSize, postRoot, postSize, preExpansion, proof)
}

// MockRollupEventInboxMetaData contains all meta data concerning the MockRollupEventInbox contract.
var MockRollupEventInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HadZeroInit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"chainConfig\",\"type\":\"string\"}],\"name\":\"rollupInitialized\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRollupAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b506080516109ae6100366000396000818160e801526102a701526109ae6000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063cb23bcb511610050578063cb23bcb514610089578063cf8d56d6146100b8578063e78cea92146100cb57600080fd5b80636ae71f121461006c578063c4d66de814610076575b600080fd5b6100746100de565b005b6100746100843660046107a2565b61029d565b60015461009c906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b6100746100c63660046107c6565b610491565b60005461009c906001600160a01b031681565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036101815760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c000000000000000000000000000000000000000060648201526084015b60405180910390fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61038054336001600160a01b038216146101f7576040517f23295f0e0000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b0382166024820152604401610178565b60008054906101000a90046001600160a01b03166001600160a01b031663cb23bcb56040518163ffffffff1660e01b8152600401602060405180830381865afa158015610248573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061026c9190610842565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03929092169190911790555050565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361033b5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c00000000000000000000000000000000000000006064820152608401610178565b6000546001600160a01b03161561037e576040517fef34ca5c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381166103be576040517f1ad0f74300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038316908117909155604080517fcb23bcb5000000000000000000000000000000000000000000000000000000008152905163cb23bcb5916004808201926020929091908290030181865afa15801561043d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104619190610842565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039290921691909117905550565b6001546001600160a01b031633146104eb5760405162461bcd60e51b815260206004820152600b60248201527f4f4e4c595f524f4c4c55500000000000000000000000000000000000000000006044820152606401610178565b806105385760405162461bcd60e51b815260206004820152601260248201527f454d5054595f434841494e5f434f4e46494700000000000000000000000000006044820152606401610178565b6001806105436106c4565b156105b857606c6001600160a01b031663f5d6ded76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610587573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ab919061085f565b6105b59082610878565b90505b600085838387876040516020016105d39594939291906108b8565b60408051808303601f190181529082905260008054825160208401207f8db5993b000000000000000000000000000000000000000000000000000000008552600b6004860152602485018390526044850152919350916001600160a01b0390911690638db5993b906064016020604051808303816000875af115801561065d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610681919061085f565b9050807fff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b836040516106b39190610929565b60405180910390a250505050505050565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f051038f200000000000000000000000000000000000000000000000000000000179052905160009182918291606491610730919061095c565b600060405180830381855afa9150503d806000811461076b576040519150601f19603f3d011682016040523d82523d6000602084013e610770565b606091505b5091509150818015610783575080516020145b9250505090565b6001600160a01b038116811461079f57600080fd5b50565b6000602082840312156107b457600080fd5b81356107bf8161078a565b9392505050565b6000806000604084860312156107db57600080fd5b83359250602084013567ffffffffffffffff808211156107fa57600080fd5b818601915086601f83011261080e57600080fd5b81358181111561081d57600080fd5b87602082850101111561082f57600080fd5b6020830194508093505050509250925092565b60006020828403121561085457600080fd5b81516107bf8161078a565b60006020828403121561087157600080fd5b5051919050565b808201808211156108b2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b8581527fff000000000000000000000000000000000000000000000000000000000000008560f81b1660208201528360218201528183604183013760009101604101908152949350505050565b60005b83811015610920578181015183820152602001610908565b50506000910152565b6020815260008251806020840152610948816040850160208701610905565b601f01601f19169190910160400192915050565b6000825161096e818460208701610905565b919091019291505056fea264697066735822122064f04a22bc2a3eedc58c40c6037652216e97a8fe5c9b9067c2ba8dbf49eaa56d64736f6c63430008110033",
}

// MockRollupEventInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use MockRollupEventInboxMetaData.ABI instead.
var MockRollupEventInboxABI = MockRollupEventInboxMetaData.ABI

// MockRollupEventInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockRollupEventInboxMetaData.Bin instead.
var MockRollupEventInboxBin = MockRollupEventInboxMetaData.Bin

// DeployMockRollupEventInbox deploys a new Ethereum contract, binding an instance of MockRollupEventInbox to it.
func DeployMockRollupEventInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockRollupEventInbox, error) {
	parsed, err := MockRollupEventInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockRollupEventInboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockRollupEventInbox{MockRollupEventInboxCaller: MockRollupEventInboxCaller{contract: contract}, MockRollupEventInboxTransactor: MockRollupEventInboxTransactor{contract: contract}, MockRollupEventInboxFilterer: MockRollupEventInboxFilterer{contract: contract}}, nil
}

// MockRollupEventInbox is an auto generated Go binding around an Ethereum contract.
type MockRollupEventInbox struct {
	MockRollupEventInboxCaller     // Read-only binding to the contract
	MockRollupEventInboxTransactor // Write-only binding to the contract
	MockRollupEventInboxFilterer   // Log filterer for contract events
}

// MockRollupEventInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockRollupEventInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockRollupEventInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockRollupEventInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockRollupEventInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockRollupEventInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockRollupEventInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockRollupEventInboxSession struct {
	Contract     *MockRollupEventInbox // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MockRollupEventInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockRollupEventInboxCallerSession struct {
	Contract *MockRollupEventInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// MockRollupEventInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockRollupEventInboxTransactorSession struct {
	Contract     *MockRollupEventInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// MockRollupEventInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockRollupEventInboxRaw struct {
	Contract *MockRollupEventInbox // Generic contract binding to access the raw methods on
}

// MockRollupEventInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockRollupEventInboxCallerRaw struct {
	Contract *MockRollupEventInboxCaller // Generic read-only contract binding to access the raw methods on
}

// MockRollupEventInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockRollupEventInboxTransactorRaw struct {
	Contract *MockRollupEventInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockRollupEventInbox creates a new instance of MockRollupEventInbox, bound to a specific deployed contract.
func NewMockRollupEventInbox(address common.Address, backend bind.ContractBackend) (*MockRollupEventInbox, error) {
	contract, err := bindMockRollupEventInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockRollupEventInbox{MockRollupEventInboxCaller: MockRollupEventInboxCaller{contract: contract}, MockRollupEventInboxTransactor: MockRollupEventInboxTransactor{contract: contract}, MockRollupEventInboxFilterer: MockRollupEventInboxFilterer{contract: contract}}, nil
}

// NewMockRollupEventInboxCaller creates a new read-only instance of MockRollupEventInbox, bound to a specific deployed contract.
func NewMockRollupEventInboxCaller(address common.Address, caller bind.ContractCaller) (*MockRollupEventInboxCaller, error) {
	contract, err := bindMockRollupEventInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockRollupEventInboxCaller{contract: contract}, nil
}

// NewMockRollupEventInboxTransactor creates a new write-only instance of MockRollupEventInbox, bound to a specific deployed contract.
func NewMockRollupEventInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*MockRollupEventInboxTransactor, error) {
	contract, err := bindMockRollupEventInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockRollupEventInboxTransactor{contract: contract}, nil
}

// NewMockRollupEventInboxFilterer creates a new log filterer instance of MockRollupEventInbox, bound to a specific deployed contract.
func NewMockRollupEventInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*MockRollupEventInboxFilterer, error) {
	contract, err := bindMockRollupEventInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockRollupEventInboxFilterer{contract: contract}, nil
}

// bindMockRollupEventInbox binds a generic wrapper to an already deployed contract.
func bindMockRollupEventInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockRollupEventInboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockRollupEventInbox *MockRollupEventInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockRollupEventInbox.Contract.MockRollupEventInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockRollupEventInbox *MockRollupEventInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockRollupEventInbox.Contract.MockRollupEventInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockRollupEventInbox *MockRollupEventInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockRollupEventInbox.Contract.MockRollupEventInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockRollupEventInbox *MockRollupEventInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockRollupEventInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockRollupEventInbox *MockRollupEventInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockRollupEventInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockRollupEventInbox *MockRollupEventInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockRollupEventInbox.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_MockRollupEventInbox *MockRollupEventInboxCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockRollupEventInbox.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_MockRollupEventInbox *MockRollupEventInboxSession) Bridge() (common.Address, error) {
	return _MockRollupEventInbox.Contract.Bridge(&_MockRollupEventInbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_MockRollupEventInbox *MockRollupEventInboxCallerSession) Bridge() (common.Address, error) {
	return _MockRollupEventInbox.Contract.Bridge(&_MockRollupEventInbox.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_MockRollupEventInbox *MockRollupEventInboxCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockRollupEventInbox.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_MockRollupEventInbox *MockRollupEventInboxSession) Rollup() (common.Address, error) {
	return _MockRollupEventInbox.Contract.Rollup(&_MockRollupEventInbox.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_MockRollupEventInbox *MockRollupEventInboxCallerSession) Rollup() (common.Address, error) {
	return _MockRollupEventInbox.Contract.Rollup(&_MockRollupEventInbox.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_MockRollupEventInbox *MockRollupEventInboxTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _MockRollupEventInbox.contract.Transact(opts, "initialize", _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_MockRollupEventInbox *MockRollupEventInboxSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _MockRollupEventInbox.Contract.Initialize(&_MockRollupEventInbox.TransactOpts, _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_MockRollupEventInbox *MockRollupEventInboxTransactorSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _MockRollupEventInbox.Contract.Initialize(&_MockRollupEventInbox.TransactOpts, _bridge)
}

// RollupInitialized is a paid mutator transaction binding the contract method 0xcf8d56d6.
//
// Solidity: function rollupInitialized(uint256 chainId, string chainConfig) returns()
func (_MockRollupEventInbox *MockRollupEventInboxTransactor) RollupInitialized(opts *bind.TransactOpts, chainId *big.Int, chainConfig string) (*types.Transaction, error) {
	return _MockRollupEventInbox.contract.Transact(opts, "rollupInitialized", chainId, chainConfig)
}

// RollupInitialized is a paid mutator transaction binding the contract method 0xcf8d56d6.
//
// Solidity: function rollupInitialized(uint256 chainId, string chainConfig) returns()
func (_MockRollupEventInbox *MockRollupEventInboxSession) RollupInitialized(chainId *big.Int, chainConfig string) (*types.Transaction, error) {
	return _MockRollupEventInbox.Contract.RollupInitialized(&_MockRollupEventInbox.TransactOpts, chainId, chainConfig)
}

// RollupInitialized is a paid mutator transaction binding the contract method 0xcf8d56d6.
//
// Solidity: function rollupInitialized(uint256 chainId, string chainConfig) returns()
func (_MockRollupEventInbox *MockRollupEventInboxTransactorSession) RollupInitialized(chainId *big.Int, chainConfig string) (*types.Transaction, error) {
	return _MockRollupEventInbox.Contract.RollupInitialized(&_MockRollupEventInbox.TransactOpts, chainId, chainConfig)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_MockRollupEventInbox *MockRollupEventInboxTransactor) UpdateRollupAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockRollupEventInbox.contract.Transact(opts, "updateRollupAddress")
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_MockRollupEventInbox *MockRollupEventInboxSession) UpdateRollupAddress() (*types.Transaction, error) {
	return _MockRollupEventInbox.Contract.UpdateRollupAddress(&_MockRollupEventInbox.TransactOpts)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_MockRollupEventInbox *MockRollupEventInboxTransactorSession) UpdateRollupAddress() (*types.Transaction, error) {
	return _MockRollupEventInbox.Contract.UpdateRollupAddress(&_MockRollupEventInbox.TransactOpts)
}

// MockRollupEventInboxInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the MockRollupEventInbox contract.
type MockRollupEventInboxInboxMessageDeliveredIterator struct {
	Event *MockRollupEventInboxInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *MockRollupEventInboxInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockRollupEventInboxInboxMessageDelivered)
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
		it.Event = new(MockRollupEventInboxInboxMessageDelivered)
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
func (it *MockRollupEventInboxInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockRollupEventInboxInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockRollupEventInboxInboxMessageDelivered represents a InboxMessageDelivered event raised by the MockRollupEventInbox contract.
type MockRollupEventInboxInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_MockRollupEventInbox *MockRollupEventInboxFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*MockRollupEventInboxInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _MockRollupEventInbox.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &MockRollupEventInboxInboxMessageDeliveredIterator{contract: _MockRollupEventInbox.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_MockRollupEventInbox *MockRollupEventInboxFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *MockRollupEventInboxInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _MockRollupEventInbox.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockRollupEventInboxInboxMessageDelivered)
				if err := _MockRollupEventInbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_MockRollupEventInbox *MockRollupEventInboxFilterer) ParseInboxMessageDelivered(log types.Log) (*MockRollupEventInboxInboxMessageDelivered, error) {
	event := new(MockRollupEventInboxInboxMessageDelivered)
	if err := _MockRollupEventInbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockRollupEventInboxInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the MockRollupEventInbox contract.
type MockRollupEventInboxInboxMessageDeliveredFromOriginIterator struct {
	Event *MockRollupEventInboxInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *MockRollupEventInboxInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockRollupEventInboxInboxMessageDeliveredFromOrigin)
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
		it.Event = new(MockRollupEventInboxInboxMessageDeliveredFromOrigin)
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
func (it *MockRollupEventInboxInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockRollupEventInboxInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockRollupEventInboxInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the MockRollupEventInbox contract.
type MockRollupEventInboxInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_MockRollupEventInbox *MockRollupEventInboxFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*MockRollupEventInboxInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _MockRollupEventInbox.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &MockRollupEventInboxInboxMessageDeliveredFromOriginIterator{contract: _MockRollupEventInbox.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_MockRollupEventInbox *MockRollupEventInboxFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *MockRollupEventInboxInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _MockRollupEventInbox.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockRollupEventInboxInboxMessageDeliveredFromOrigin)
				if err := _MockRollupEventInbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_MockRollupEventInbox *MockRollupEventInboxFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*MockRollupEventInboxInboxMessageDeliveredFromOrigin, error) {
	event := new(MockRollupEventInboxInboxMessageDeliveredFromOrigin)
	if err := _MockRollupEventInbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxBlobMockMetaData contains all meta data concerning the SequencerInboxBlobMock contract.
var SequencerInboxBlobMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxDataSize_\",\"type\":\"uint256\"},{\"internalType\":\"contractIReader4844\",\"name\":\"reader_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isUsingFeeToken_\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDelayBufferable_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"AlreadyValidDASKeyset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadBufferConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadMaxTimeVariation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"BadSequencerNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSetFeeTokenPricer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataBlobsNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDataLength\",\"type\":\"uint256\"}],\"name\":\"DataTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayProofRequired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayedBackwards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayedTooFar\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Deprecated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtraGasNotUint64\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceIncludeBlockTooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HadZeroInit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectMessagePreimage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"InitParamZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDelayedAccPreimage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"name\":\"InvalidHeaderFlag\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeysetTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingDataHashes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"NoSuchKeyset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBatchPoster\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NotBatchPosterManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCodelessOrigin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotDelayBufferable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotForked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupNotChanged\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBatchPosterManager\",\"type\":\"address\"}],\"name\":\"BatchPosterManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"batchPoster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBatchPoster\",\"type\":\"bool\"}],\"name\":\"BatchPosterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"replenishRateInBasis\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBufferConfig\",\"name\":\"bufferConfig\",\"type\":\"tuple\"}],\"name\":\"BufferConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeTokenPricer\",\"type\":\"address\"}],\"name\":\"FeeTokenPricerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keysetHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidateKeyset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation\",\"type\":\"tuple\"}],\"name\":\"MaxTimeVariationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"OwnerFunctionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchSequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SequencerBatchData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchSequenceNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxBlockNumber\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIBridge.TimeBounds\",\"name\":\"timeBounds\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"enumIBridge.BatchDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"}],\"name\":\"SequencerBatchDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSequencer\",\"type\":\"bool\"}],\"name\":\"SequencerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keysetHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"keysetBytes\",\"type\":\"bytes\"}],\"name\":\"SetValidKeyset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BROTLI_MESSAGE_HEADER_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CUSTOM_DA_MESSAGE_HEADER_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAS_MESSAGE_HEADER_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DATA_AUTHENTICATED_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DATA_BLOB_HEADER_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HEADER_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TREE_DAS_MESSAGE_HEADER_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_HEAVY_MESSAGE_HEADER_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"addSequencerL2Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"beforeDelayedAcc\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMessages.Message\",\"name\":\"delayedMessage\",\"type\":\"tuple\"}],\"internalType\":\"structDelayProof\",\"name\":\"delayProof\",\"type\":\"tuple\"}],\"name\":\"addSequencerL2BatchDelayProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"addSequencerL2BatchFromBlobs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"beforeDelayedAcc\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMessages.Message\",\"name\":\"delayedMessage\",\"type\":\"tuple\"}],\"internalType\":\"structDelayProof\",\"name\":\"delayProof\",\"type\":\"tuple\"}],\"name\":\"addSequencerL2BatchFromBlobsDelayProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addSequencerL2BatchFromOrigin\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"addSequencerL2BatchFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"beforeDelayedAcc\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMessages.Message\",\"name\":\"delayedMessage\",\"type\":\"tuple\"}],\"internalType\":\"structDelayProof\",\"name\":\"delayProof\",\"type\":\"tuple\"}],\"name\":\"addSequencerL2BatchFromOriginDelayProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchPosterManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buffer\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"bufferBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"prevBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"replenishRateInBasis\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"prevSequencedBlockNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dasKeySetInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidKeyset\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"creationBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTokenPricer\",\"outputs\":[{\"internalType\":\"contractIFeeTokenPricer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint64[2]\",\"name\":\"l1BlockAndTime\",\"type\":\"uint64[2]\"},{\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"forceInclusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"}],\"name\":\"forceInclusionDeadline\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"getKeysetCreationBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"inboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"replenishRateInBasis\",\"type\":\"uint64\"}],\"internalType\":\"structBufferConfig\",\"name\":\"bufferConfig_\",\"type\":\"tuple\"},{\"internalType\":\"contractIFeeTokenPricer\",\"name\":\"feeTokenPricer_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"invalidateKeysetHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBatchPoster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDelayBufferable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUsingFeeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"isValidKeysetHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDataSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTimeVariation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"replenishRateInBasis\",\"type\":\"uint64\"}],\"internalType\":\"structBufferConfig\",\"name\":\"bufferConfig_\",\"type\":\"tuple\"}],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reader4844\",\"outputs\":[{\"internalType\":\"contractIReader4844\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeDelayAfterFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBatchPosterManager\",\"type\":\"address\"}],\"name\":\"setBatchPosterManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"replenishRateInBasis\",\"type\":\"uint64\"}],\"internalType\":\"structBufferConfig\",\"name\":\"bufferConfig_\",\"type\":\"tuple\"}],\"name\":\"setBufferConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFeeTokenPricer\",\"name\":\"feeTokenPricer_\",\"type\":\"address\"}],\"name\":\"setFeeTokenPricer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isBatchPoster_\",\"type\":\"bool\"}],\"name\":\"setIsBatchPoster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSequencer_\",\"type\":\"bool\"}],\"name\":\"setIsSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation_\",\"type\":\"tuple\"}],\"name\":\"setMaxTimeVariation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"keysetBytes\",\"type\":\"bytes\"}],\"name\":\"setValidKeyset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDelayedMessagesRead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRollupAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x610180604052306080526202000060a05246610100526200002b62000115602090811b6200318617901c565b1515610120523480156200003e57600080fd5b50604051620057f5380380620057f58339810160408190526200006191620001c8565b838383838360e081815250506101205115620000a6576001600160a01b03831615620000a0576040516386657a5360e01b815260040160405180910390fd5b620000ef565b6001600160a01b038316620000ef576040516380fc2c0360e01b815260206004820152600a60248201526914995859195c8d0e0d0d60b21b604482015260640160405180910390fd5b6001600160a01b0390921660c052151561014052151561016052506200025b9350505050565b60408051600481526024810182526020810180516001600160e01b03166302881c7960e11b1790529051600091829182916064916200015591906200022a565b600060405180830381855afa9150503d806000811462000192576040519150601f19603f3d011682016040523d82523d6000602084013e62000197565b606091505b5091509150818015620001ab575080516020145b9250505090565b80518015158114620001c357600080fd5b919050565b60008060008060808587031215620001df57600080fd5b845160208601519094506001600160a01b0381168114620001ff57600080fd5b92506200020f60408601620001b2565b91506200021f60608601620001b2565b905092959194509250565b6000825160005b818110156200024d576020818601810151858301520162000231565b506000920191825250919050565b60805160a05160c05160e051610100516101205161014051610160516154856200037060003960008181610483015281816112fc015281816117f001528181611ecc01528181612333015281816126b901528181612cd101528181612e660152818161324e0152613490015260008181610628015281816109b501528181612578015281816126e801528181613f57015261407e015260008181612ab20152818161359c0152613f9701526000818161218e0152613a0901526000818161077c01528181614345015261439a0152600081816105db01528181610f9a0152611e7501526000818161117f015281816114bd01528181611d6a015261208401526000818161220d01526123cb01526154856000f3fe608060405234801561001057600080fd5b50600436106103155760003560e01c80637fa3a40e116101a7578063d1ce8da8116100ee578063e78cea9211610097578063edaafe2011610071578063edaafe20146107c6578063f19815781461084f578063f60a50911461086257600080fd5b8063e78cea9214610764578063e8eb1dc314610777578063ebea461d1461079e57600080fd5b8063dd44e6e0116100c8578063dd44e6e0146106fe578063e0bc97291461072a578063e5a358c81461073d57600080fd5b8063d1ce8da8146106b1578063d9dd67ab146106c4578063dab341a4146106d757600080fd5b806396cc5c7811610150578063b31761f81161012a578063b31761f814610678578063cb23bcb51461068b578063cc2a1a0c1461069e57600080fd5b806396cc5c781461064a578063a655d93714610652578063a84840b71461066557600080fd5b80638f111f3c116101815780638f111f3c146105fd578063917cf8ac1461061057806392d9f7821461062357600080fd5b80637fa3a40e146105ba57806384420860146105c35780638d910dde146105d657600080fd5b80632f3985a71161026b5780636d46e987116102145780636f12b0c9116101ee5780636f12b0c914610530578063715ea34b1461054357806371c3e6fe1461059757600080fd5b80636d46e987146104e75780636e6200551461050a5780636e7df3e71461051d57600080fd5b806369cacded1161024557806369cacded146104a55780636ae71f12146104b85780636c890450146104c057600080fd5b80632f3985a7146104585780633e5aa0821461046b5780634b678a661461047e57600080fd5b80631f956632116102cd578063258f0495116102a7578063258f04951461041657806327957a49146104295780632cbf74e51461043157600080fd5b80631f956632146103c55780631ff64790146103d857806322291e8d146103eb57600080fd5b806306f13056116102fe57806306f13056146103745780631637be481461038a57806316af91a7146103bd57600080fd5b806302c992751461031a578063036f7ed31461035f575b600080fd5b6103417f200000000000000000000000000000000000000000000000000000000000000081565b6040516001600160f81b031990911681526020015b60405180910390f35b61037261036d3660046148af565b61086d565b005b61037c610a97565b604051908152602001610356565b6103ad6103983660046148cc565b60009081526008602052604090205460ff1690565b6040519015158152602001610356565b610341600081565b6103726103d33660046148f3565b610b21565b6103726103e63660046148af565b610c8c565b600e546103fe906001600160a01b031681565b6040516001600160a01b039091168152602001610356565b61037c6104243660046148cc565b610df1565b61037c602881565b6103417f500000000000000000000000000000000000000000000000000000000000000081565b610372610466366004614a4b565b610e5e565b610372610479366004614a67565b610f97565b6103ad7f000000000000000000000000000000000000000000000000000000000000000081565b6103726104b3366004614b12565b611285565b6103726115c5565b6103417f080000000000000000000000000000000000000000000000000000000000000081565b6103ad6104f53660046148af565b60096020526000908152604090205460ff1681565b610372610518366004614b12565b61179d565b61037261052b3660046148f3565b61184f565b61037261053e366004614ba0565b6119ba565b6105776105513660046148cc565b60086020526000908152604090205460ff811690610100900467ffffffffffffffff1682565b60408051921515835267ffffffffffffffff909116602083015201610356565b6103ad6105a53660046148af565b60036020526000908152604090205460ff1681565b61037c60005481565b6103726105d13660046148cc565b6119ec565b6103fe7f000000000000000000000000000000000000000000000000000000000000000081565b61037261060b366004614c0b565b611b61565b61037261061e366004614c88565b611e72565b6103ad7f000000000000000000000000000000000000000000000000000000000000000081565b61037261218b565b610372610660366004614a4b565b612203565b610372610673366004614ce4565b6123c1565b610372610686366004614d4b565b612786565b6002546103fe906001600160a01b031681565b600b546103fe906001600160a01b031681565b6103726106bf366004614db1565b6128e5565b61037c6106d23660046148cc565b612c32565b6103417f010000000000000000000000000000000000000000000000000000000000000081565b61071161070c366004614df3565b612cbf565b60405167ffffffffffffffff9091168152602001610356565b610372610738366004614c0b565b612d22565b6103417f400000000000000000000000000000000000000000000000000000000000000081565b6001546103fe906001600160a01b031681565b61037c7f000000000000000000000000000000000000000000000000000000000000000081565b6107a6612daa565b604080519485526020850193909352918301526060820152608001610356565b600c54600d5461080c9167ffffffffffffffff8082169268010000000000000000808404831693600160801b8104841693600160c01b9091048116928082169290041686565b6040805167ffffffffffffffff978816815295871660208701529386169385019390935290841660608401528316608083015290911660a082015260c001610356565b61037261085d366004614e1f565b612de3565b610341600160ff1b81565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e49190614e87565b6001600160a01b0316336001600160a01b0316146109b35760025460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610945573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109699190614e87565b6040517f23295f0e0000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152911660248201526044015b60405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000610a0a576040517fe13123b400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527fe83d6153add50e41b8ee6c1115c4178687349bb12bc3902a50b1f6ad78a0c5419060200160405180910390a16040516006907fea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e90600090a250565b600154604080517e84120c00000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916284120c9160048083019260209291908290030181865afa158015610af8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b1c9190614ea4565b905090565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b989190614e87565b6001600160a01b0316336001600160a01b031614158015610bc45750600b546001600160a01b03163314155b15610bfd576040517f660b3b420000000000000000000000000000000000000000000000000000000081523360048201526024016109aa565b6001600160a01b038216600081815260096020908152604091829020805460ff19168515159081179091558251938452908301527feb12a9a53eec138c91b27b4f912a257bd690c18fc8bde744be92a0365eb9b87e910160405180910390a16040516004907fea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e90600090a25050565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cdf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d039190614e87565b6001600160a01b0316336001600160a01b031614610d645760025460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610945573d6000803e3d6000fd5b600b805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f3cd6c184800297a0f2b00926a683cbe76890bb7fd01480ac0a10ed6c8f7f66599060200160405180910390a16040516005907fea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e90600090a250565b600081815260086020908152604080832081518083019092525460ff811615158252610100900467ffffffffffffffff16918101829052908203610e4a5760405162f20c5d60e01b8152600481018490526024016109aa565b6020015167ffffffffffffffff1692915050565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610eb1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed59190614e87565b6001600160a01b0316336001600160a01b031614610f365760025460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610945573d6000803e3d6000fd5b610f3f8161324c565b60408051825167ffffffffffffffff908116825260208085015182169083015283830151168183015290517faa7a2d8175dee3b637814ad6346005dfcc357165396fb8327f649effe8abcf859181900360600190a150565b827f000000000000000000000000000000000000000000000000000000000000000060005a3360009081526003602052604090205490915060ff16610fef57604051632dd9fc9760e01b815260040160405180910390fd5b610ff88761348c565b1561101657604051630e5da8fb60e01b815260040160405180910390fd5b611022888887876134d4565b6001600160a01b0383161561127b57366000602061104183601f614ed3565b61104b9190614ee6565b905061020061105b600283614fec565b6110659190614ee6565b611070826006614ffb565b61107a9190614ed3565b6110849084614ed3565b925061108e61360b565b61109b57600091506111ce565b6001600160a01b038416156111ce57836001600160a01b031663e83a2d826040518163ffffffff1660e01b8152600401600060405180830381865afa92505050801561110957506040513d6000823e601f3d908101601f191682016040526111069190810190615012565b60015b156111ce578051156111cc576000856001600160a01b0316631f6d6ef76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611155573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111799190614ea4565b905048817f000000000000000000000000000000000000000000000000000000000000000084516111aa9190614ffb565b6111b49190614ffb565b6111be9190614ee6565b6111c89086614ed3565b9450505b505b846001600160a01b031663e3db8a49335a6111e990876150b8565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b0390921660048301526024820152604481018590526064016020604051808303816000875af1158015611253573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061127791906150cb565b5050505b5050505050505050565b836000805a905061129461360b565b6112ca576040517fc8958ead00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526003602052604090205460ff166112fa57604051632dd9fc9760e01b815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000061133857604051631192b39960e31b815260040160405180910390fd5b6113508861134b368790038701876150e8565b61361e565b6113608b8b8b8b8a8a600161372b565b6001600160a01b0383161561127757366000602061137f83601f614ed3565b6113899190614ee6565b9050610200611399600283614fec565b6113a39190614ee6565b6113ae826006614ffb565b6113b89190614ed3565b6113c29084614ed3565b92506113cc61360b565b6113d9576000915061150c565b6001600160a01b0384161561150c57836001600160a01b031663e83a2d826040518163ffffffff1660e01b8152600401600060405180830381865afa92505050801561144757506040513d6000823e601f3d908101601f191682016040526114449190810190615012565b60015b1561150c5780511561150a576000856001600160a01b0316631f6d6ef76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611493573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114b79190614ea4565b905048817f000000000000000000000000000000000000000000000000000000000000000084516114e89190614ffb565b6114f29190614ffb565b6114fc9190614ee6565b6115069086614ed3565b9450505b505b846001600160a01b031663e3db8a49335a61152790876150b8565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b0390921660048301526024820152604481018590526064016020604051808303816000875af1158015611591573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115b591906150cb565b5050505050505050505050505050565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611618573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061163c9190614e87565b6001600160a01b0316336001600160a01b03161461169d5760025460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610945573d6000803e3d6000fd5b600154604080517fcb23bcb500000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163cb23bcb59160048083019260209291908290030181865afa158015611700573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117249190614e87565b6002549091506001600160a01b0380831691160361176e576040517fd054909f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b836000805a3360009081526003602052604090205490915060ff161580156117d057506002546001600160a01b03163314155b156117ee57604051632dd9fc9760e01b815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000061182c57604051631192b39960e31b815260040160405180910390fd5b61183f8861134b368790038701876150e8565b6113608b8b8b8b8a8a600061372b565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118c69190614e87565b6001600160a01b0316336001600160a01b0316141580156118f25750600b546001600160a01b03163314155b1561192b576040517f660b3b420000000000000000000000000000000000000000000000000000000081523360048201526024016109aa565b6001600160a01b038216600081815260036020908152604091829020805460ff19168515159081179091558251938452908301527f28bcc5626d357efe966b4b0876aa1ee8ab99e26da4f131f6a2623f1800701c21910160405180910390a16040516001907fea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e90600090a25050565b6040517fc73b9d7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a639190614e87565b6001600160a01b0316336001600160a01b031614611ac45760025460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610945573d6000803e3d6000fd5b60008181526008602052604090205460ff16611af55760405162f20c5d60e01b8152600481018290526024016109aa565b600081815260086020526040808220805460ff191690555182917f5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a91a26040516003907fea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e90600090a250565b826000805a9050611b7061360b565b611ba6576040517fc8958ead00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526003602052604090205460ff16611bd657604051632dd9fc9760e01b815260040160405180910390fd5b611bdf8761348c565b15611bfd57604051630e5da8fb60e01b815260040160405180910390fd5b611c0d8a8a8a8a8989600161372b565b6001600160a01b03831615611e66573660006020611c2c83601f614ed3565b611c369190614ee6565b9050610200611c46600283614fec565b611c509190614ee6565b611c5b826006614ffb565b611c659190614ed3565b611c6f9084614ed3565b9250611c7961360b565b611c865760009150611db9565b6001600160a01b03841615611db957836001600160a01b031663e83a2d826040518163ffffffff1660e01b8152600401600060405180830381865afa925050508015611cf457506040513d6000823e601f3d908101601f19168201604052611cf19190810190615012565b60015b15611db957805115611db7576000856001600160a01b0316631f6d6ef76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d649190614ea4565b905048817f00000000000000000000000000000000000000000000000000000000000000008451611d959190614ffb565b611d9f9190614ffb565b611da99190614ee6565b611db39086614ed3565b9450505b505b846001600160a01b031663e3db8a49335a611dd490876150b8565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b0390921660048301526024820152604481018590526064016020604051808303816000875af1158015611e3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e6291906150cb565b5050505b50505050505050505050565b837f000000000000000000000000000000000000000000000000000000000000000060005a3360009081526003602052604090205490915060ff16611eca57604051632dd9fc9760e01b815260040160405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000611f0857604051631192b39960e31b815260040160405180910390fd5b611f1b8861134b368790038701876150e8565b611f27898988886134d4565b6001600160a01b03831615612180573660006020611f4683601f614ed3565b611f509190614ee6565b9050610200611f60600283614fec565b611f6a9190614ee6565b611f75826006614ffb565b611f7f9190614ed3565b611f899084614ed3565b9250611f9361360b565b611fa057600091506120d3565b6001600160a01b038416156120d357836001600160a01b031663e83a2d826040518163ffffffff1660e01b8152600401600060405180830381865afa92505050801561200e57506040513d6000823e601f3d908101601f1916820160405261200b9190810190615012565b60015b156120d3578051156120d1576000856001600160a01b0316631f6d6ef76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561205a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061207e9190614ea4565b905048817f000000000000000000000000000000000000000000000000000000000000000084516120af9190614ffb565b6120b99190614ffb565b6120c39190614ee6565b6120cd9086614ed3565b9450505b505b846001600160a01b031663e3db8a49335a6120ee90876150b8565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b0390921660048301526024820152604481018590526064016020604051808303816000875af1158015612158573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061217c91906150cb565b5050505b505050505050505050565b467f0000000000000000000000000000000000000000000000000000000000000000036121e4576040517fa301bb0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7801000000000000000100000000000000010000000000000001600a55565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036122bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c000000000000000000000000000000000000000060648201526084016109aa565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61038054336001600160a01b03821614612331576040517f23295f0e0000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b03821660248201526044016109aa565b7f000000000000000000000000000000000000000000000000000000000000000061236f57604051631192b39960e31b815260040160405180910390fd5b600c5467ffffffffffffffff16156123b3576040517fef34ca5c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6123bc8361324c565b505050565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003612479576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c000000000000000000000000000000000000000060648201526084016109aa565b6001546001600160a01b0316156124bc576040517fef34ca5c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0384166124fc576040517f1ad0f74300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000846001600160a01b031663e1758bd86040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612558575060408051601f3d908101601f1916820190925261255591810190614e87565b60015b15612573576001600160a01b0381161561257157600191505b505b8015157f00000000000000000000000000000000000000000000000000000000000000001515146125d0576040517fc3e31f8d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038716908117909155604080517fcb23bcb5000000000000000000000000000000000000000000000000000000008152905163cb23bcb5916004808201926020929091908290030181865afa15801561264f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126739190614e87565b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03929092169190911790556126b76126b236869003860186614d4b565b613856565b7f0000000000000000000000000000000000000000000000000000000000000000156126e6576126e68361324c565b7f000000000000000000000000000000000000000000000000000000000000000015801561271c57506001600160a01b03821615155b15612753576040517fe13123b400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50600e805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055505050565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156127d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127fd9190614e87565b6001600160a01b0316336001600160a01b03161461285e5760025460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610945573d6000803e3d6000fd5b61286781613856565b60408051825181526020808401519082015282820151818301526060808401519082015290517faa6a58dad31128ff7ecc2b80987ee6e003df80bc50cd8d0b0d1af0e07da6d19d9181900360800190a16040516000907fea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e908290a250565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612938573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061295c9190614e87565b6001600160a01b0316336001600160a01b0316146129bd5760025460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610945573d6000803e3d6000fd5b600082826040516129cf929190615196565b6040519081900381207ffe000000000000000000000000000000000000000000000000000000000000006020830152602182015260410160408051601f1981840301815291905280516020909101209050600160ff1b8118620100008310612a63576040517fb3d1f41200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008181526008602052604090205460ff1615612aaf576040517ffa2fddda000000000000000000000000000000000000000000000000000000008152600481018290526024016109aa565b437f000000000000000000000000000000000000000000000000000000000000000015612b3c5760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612b15573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b399190614ea4565b90505b6040805180820182526001815267ffffffffffffffff8381166020808401918252600087815260089091528490209251835491517fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000009092169015157fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff161761010091909216021790555182907fabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef65572290612bf790889088906151a6565b60405180910390a26040516002907fea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e90600090a25050505050565b6001546040517f16bf5579000000000000000000000000000000000000000000000000000000008152600481018390526000916001600160a01b0316906316bf557990602401602060405180830381865afa158015612c95573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cb99190614ea4565b92915050565b600a5460009067ffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000015612d11576000612d02600c85613982565b9050612d0d816139d1565b9150505b612d1b81846151d5565b9392505050565b826000805a3360009081526003602052604090205490915060ff16158015612d5557506002546001600160a01b03163314155b15612d7357604051632dd9fc9760e01b815260040160405180910390fd5b612d7c8761348c565b15612d9a57604051630e5da8fb60e01b815260040160405180910390fd5b611c0d8a8a8a8a8989600061372b565b600080600080600080600080612dbe613a01565b67ffffffffffffffff9384169b50918316995082169750169450505050505b90919293565b6000548611612e1e576040517f7d73e6fa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000612e548684612e326020890189614df3565b612e4260408a0160208b01614df3565b612e4d60018d6150b8565b8988613a78565b600a5490915067ffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000015612ec557612ea2612e9a6020880188614df3565b600c90613b1d565b600c54612eb89067ffffffffffffffff166139d1565b67ffffffffffffffff1690505b4381612ed46020890189614df3565b67ffffffffffffffff16612ee89190614ed3565b10612f1f576040517fad3515d900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006001891115612fa8576001546001600160a01b031663d5719dc2612f4660028c6150b8565b6040518263ffffffff1660e01b8152600401612f6491815260200190565b602060405180830381865afa158015612f81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fa59190614ea4565b90505b60408051602080820184905281830186905282518083038401815260609092019092528051910120600180546001600160a01b03169063d5719dc290612fee908d6150b8565b6040518263ffffffff1660e01b815260040161300c91815260200190565b602060405180830381865afa158015613029573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061304d9190614ea4565b14613084576040517f13947fd700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806130908b613ba3565b9150915060008b90506000600160009054906101000a90046001600160a01b03166001600160a01b0316635fca4a166040518163ffffffff1660e01b8152600401602060405180830381865afa1580156130ee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131129190614ea4565b90508060008080806131278988838880613be8565b93509350935093508083857f7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7856000548d600260405161316a9493929190615213565b60405180910390a4505050505050505050505050505050505050565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f051038f2000000000000000000000000000000000000000000000000000000001790529051600091829182916064916131f291906152ac565b600060405180830381855afa9150503d806000811461322d576040519150601f19603f3d011682016040523d82523d6000602084013e613232565b606091505b5091509150818015613245575080516020145b9250505090565b7f000000000000000000000000000000000000000000000000000000000000000061328a57604051631192b39960e31b815260040160405180910390fd5b61329381613da5565b6132c9576040517fda1c8eb600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c5467ffffffffffffffff1615806132f557506020810151600c5467ffffffffffffffff9182169116115b15613321576020810151600c805467ffffffffffffffff191667ffffffffffffffff9092169190911790555b8051600c5467ffffffffffffffff9182169116101561335e578051600c805467ffffffffffffffff191667ffffffffffffffff9092169190911790555b602081810151600c805484517fffffffffffffffff00000000000000000000000000000000ffffffffffffffff9091166801000000000000000067ffffffffffffffff948516027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1617600160801b91841691909102179055604080840151600d805467ffffffffffffffff1916919093161790915560005460015482517feca067ad000000000000000000000000000000000000000000000000000000008152925191936001600160a01b039091169263eca067ad92600480830193928290030181865afa158015613455573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134799190614ea4565b0361348957613489600c43613b1d565b50565b60007f000000000000000000000000000000000000000000000000000000000000000080156134bc575060005482115b8015612cb957506134cd600c613e0d565b1592915050565b60008060006134e286613e40565b9250925092506000806000806134fc878b60008c8c613be8565b93509350935093508a841415801561351657506000198b14155b15613557576040517fac7411c900000000000000000000000000000000000000000000000000000000815260048101859052602481018c90526044016109aa565b80838c7f7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7856000548b60036040516135929493929190615213565b60405180910390a47f0000000000000000000000000000000000000000000000000000000000000000156135f2576040517f86657a5300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6135fa61360b565b156112775761127787854888613f49565b60003332148015610b1c575050333b1590565b60005482111561372757613632600c61429b565b1561372757600154600080546040517fd5719dc200000000000000000000000000000000000000000000000000000000815291926001600160a01b03169163d5719dc2916136869160040190815260200190565b602060405180830381865afa1580156136a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136c79190614ea4565b90506136dc81836000015184602001516142cc565b613712576040517fc334542d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020820151604001516123bc90600c90613b1d565b5050565b600080613739888888614311565b9150915060008060008061375d868b89613754576000613756565b8d5b8c8c613be8565b93509350935093508c841415801561377757506000198d14155b156137b8576040517fac7411c900000000000000000000000000000000000000000000000000000000815260048101859052602481018e90526044016109aa565b8083857f7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7856000548a8d6137ed5760016137f0565b60005b6040516138009493929190615213565b60405180910390a486611e6257837ffe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c208d8d60405161383f9291906151a6565b60405180910390a250505050505050505050505050565b805167ffffffffffffffff10806138785750602081015167ffffffffffffffff105b8061388e5750604081015167ffffffffffffffff105b806138a45750606081015167ffffffffffffffff105b156138db576040517f09cfba7500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600a80546020840151604085015160609095015167ffffffffffffffff908116600160c01b0277ffffffffffffffffffffffffffffffffffffffffffffffff968216600160801b02969096166fffffffffffffffffffffffffffffffff92821668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009094169190951617919091171691909117919091179055565b81546001830154600091612d1b9167ffffffffffffffff600160c01b8304811692868216928282169268010000000000000000808304821693600160801b81048316939190048216911661451f565b600a5460009067ffffffffffffffff908116908316106139fd57600a5467ffffffffffffffff16612cb9565b5090565b6000808080467f000000000000000000000000000000000000000000000000000000000000000014613a3e57506001925082915081905080612ddd565b5050600a5467ffffffffffffffff808216935068010000000000000000820481169250600160801b8204811691600160c01b900416612ddd565b6040516001600160f81b031960f889901b1660208201526bffffffffffffffffffffffff19606088901b1660218201527fffffffffffffffff00000000000000000000000000000000000000000000000060c087811b8216603584015286901b16603d82015260458101849052606581018390526085810182905260009060a5016040516020818303038152906040528051906020012090505b979650505050505050565b613b278282613982565b825467ffffffffffffffff928316600160c01b0277ffffffffffffffffffffffffffffffff000000000000000090911691831691909117178255600190910180544390921668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff909216919091179055565b6040805160808101825260008082526020820181905291810182905260608101829052600080613bd2856145e6565b8151602090920191909120969095509350505050565b600080600080600054881015613c2a576040517f7d73e6fa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160009054906101000a90046001600160a01b03166001600160a01b031663eca067ad6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613c7d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ca19190614ea4565b881115613cda576040517f925f8bd300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040517f86598a56000000000000000000000000000000000000000000000000000000008152600481018b9052602481018a905260448101889052606481018790526001600160a01b03909116906386598a56906084016080604051808303816000875af1158015613d53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d7791906152c8565b60008c9055929650909450925090508615613d9957613d998985486000613f49565b95509550955095915050565b805160009067ffffffffffffffff1615801590613dcf5750602082015167ffffffffffffffff1615155b8015613deb5750612710826040015167ffffffffffffffff1611155b8015612cb95750506020810151905167ffffffffffffffff9182169116111590565b805460009067ffffffffffffffff600160801b8204811691613e3891600160c01b90910416436150b8565b111592915050565b604080516080810182526000808252602082018190529181018290526060810182905260408051606081018252600080825260208201819052918101829052600080613e8b876145e6565b9092509050633b9aca0060006003613ea66202000084614ffb565b613eb09190614ffb565b60405190915084907f500000000000000000000000000000000000000000000000000000000000000090613ee89088906020016152fe565b60408051601f1981840301815290829052613f07939291602001615332565b604051602081830303815290604052805190602001208360004811613f2d576000613f37565b613f374884614ee6565b97509750975050505050509193909250565b600e546001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000008015613f8957506001600160a01b038116155b15613f945750614295565b327f00000000000000000000000000000000000000000000000000000000000000001561403a576000606c6001600160a01b031663c6f7de0e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613ffc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906140209190614ea4565b905061402c4882614ee6565b6140369085614ed3565b9350505b67ffffffffffffffff83111561407c576040517f04d5501200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000080156140b157506001600160a01b03821615155b1561413f576000826001600160a01b031663e6aa216c6040518163ffffffff1660e01b81526004016020604051808303816000875af11580156140f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061411c9190614ea4565b9050670de0b6b3a76400006141318287614ffb565b61413b9190614ee6565b9450505b604080514260208201526bffffffffffffffffffffffff19606084901b16918101919091526054810187905260748101869052609481018590527fffffffffffffffff00000000000000000000000000000000000000000000000060c085901b1660b482015260009060bc0160408051808303601f1901815290829052600154815160208301207f7a88b1070000000000000000000000000000000000000000000000000000000084526001600160a01b0386811660048601526024850191909152919350600092911690637a88b107906044016020604051808303816000875af1158015614232573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906142569190614ea4565b9050807fff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b836040516142889190615375565b60405180910390a2505050505b50505050565b60006142a682613e0d565b1580612cb95750505467ffffffffffffffff680100000000000000008204811691161090565b6000614307836142db846146be565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b9093149392505050565b60408051608081018252600080825260208201819052918101829052606081018290526000614341856028614ed3565b90507f00000000000000000000000000000000000000000000000000000000000000008111156143c6576040517f4634691b000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000060248201526044016109aa565b6000806143d2866145e6565b909250905086156144e557614402888860008181106143f3576143f36151fd565b9050013560f81c60f81b6146eb565b61445a5787876000818110614419576144196151fd565b6040517f6b3333560000000000000000000000000000000000000000000000000000000081529201356001600160f81b0319166004830152506024016109aa565b600160ff1b8888600081614470576144706151fd565b6001600160f81b031992013592909216161580159150614491575060218710155b156144e55760006144a6602160018a8c6153a8565b6144af916153d2565b60008181526008602052604090205490915060ff166144e35760405162f20c5d60e01b8152600481018290526024016109aa565b505b8188886040516020016144fa939291906153f0565b60408051601f1981840301815291905280516020909101209890975095505050505050565b60008088881161453057600061453a565b61453a89896150b8565b9050600089871161454c576000614556565b6145568a886150b8565b90506127106145658584614ffb565b61456f9190614ee6565b6145799089614ed3565b9750600086821161458b576000614595565b61459587836150b8565b9050828111156145a25750815b808911156145d7576145b4818a6150b8565b9850868911156145d7578589116145cb57886145cd565b855b9350505050613b12565b50949998505050505050505050565b604080516080810182526000808252602082018190529181018290526060808201839052916146136147b2565b9050600081600001518260200151836040015184606001518860405160200161469395949392919060c095861b7fffffffffffffffff000000000000000000000000000000000000000000000000908116825294861b8516600882015292851b8416601084015290841b8316601883015290921b16602082015260280190565b604051602081830303815290604052905060288151146146b5576146b5615418565b94909350915050565b6000612cb9826000015183602001518460400151856060015186608001518760a001518860c00151613a78565b60006001600160f81b03198216158061471157506001600160f81b03198216600160ff1b145b8061474557506001600160f81b031982167f8800000000000000000000000000000000000000000000000000000000000000145b8061477957506001600160f81b031982167f2000000000000000000000000000000000000000000000000000000000000000145b80612cb957506001600160f81b031982167f01000000000000000000000000000000000000000000000000000000000000001492915050565b604080516080810182526000808252602082018190529181018290526060810191909152604080516080810182526000808252602082018190529181018290526060810191909152600080600080614808613a01565b93509350935093508167ffffffffffffffff164211156148395761482c824261542e565b67ffffffffffffffff1685525b61484381426151d5565b67ffffffffffffffff9081166020870152841643111561487757614867844361542e565b67ffffffffffffffff1660408601525b61488183436151d5565b67ffffffffffffffff1660608601525092949350505050565b6001600160a01b038116811461348957600080fd5b6000602082840312156148c157600080fd5b8135612d1b8161489a565b6000602082840312156148de57600080fd5b5035919050565b801515811461348957600080fd5b6000806040838503121561490657600080fd5b82356149118161489a565b91506020830135614921816148e5565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156149655761496561492c565b60405290565b60405160e0810167ffffffffffffffff811182821017156149655761496561492c565b604051601f8201601f1916810167ffffffffffffffff811182821017156149b7576149b761492c565b604052919050565b803567ffffffffffffffff811681146149d757600080fd5b919050565b6000606082840312156149ee57600080fd5b6040516060810181811067ffffffffffffffff82111715614a1157614a1161492c565b604052905080614a20836149bf565b8152614a2e602084016149bf565b6020820152614a3f604084016149bf565b60408201525092915050565b600060608284031215614a5d57600080fd5b612d1b83836149dc565b600080600080600060a08688031215614a7f57600080fd5b85359450602086013593506040860135614a988161489a565b94979396509394606081013594506080013592915050565b60008083601f840112614ac257600080fd5b50813567ffffffffffffffff811115614ada57600080fd5b602083019150836020828501011115614af257600080fd5b9250929050565b60006101008284031215614b0c57600080fd5b50919050565b6000806000806000806000806101c0898b031215614b2f57600080fd5b88359750602089013567ffffffffffffffff811115614b4d57600080fd5b614b598b828c01614ab0565b909850965050604089013594506060890135614b748161489a565b93506080890135925060a08901359150614b918a60c08b01614af9565b90509295985092959890939650565b600080600080600060808688031215614bb857600080fd5b85359450602086013567ffffffffffffffff811115614bd657600080fd5b614be288828901614ab0565b909550935050604086013591506060860135614bfd8161489a565b809150509295509295909350565b600080600080600080600060c0888a031215614c2657600080fd5b87359650602088013567ffffffffffffffff811115614c4457600080fd5b614c508a828b01614ab0565b909750955050604088013593506060880135614c6b8161489a565b969995985093969295946080840135945060a09093013592915050565b6000806000806000806101a08789031215614ca257600080fd5b86359550602087013594506040870135614cbb8161489a565b93506060870135925060808701359150614cd88860a08901614af9565b90509295509295509295565b600080600080848603610120811215614cfc57600080fd5b8535614d078161489a565b94506080601f1982011215614d1b57600080fd5b50602085019250614d2f8660a087016149dc565b9150610100850135614d408161489a565b939692955090935050565b600060808284031215614d5d57600080fd5b6040516080810181811067ffffffffffffffff82111715614d8057614d8061492c565b8060405250823581526020830135602082015260408301356040820152606083013560608201528091505092915050565b60008060208385031215614dc457600080fd5b823567ffffffffffffffff811115614ddb57600080fd5b614de785828601614ab0565b90969095509350505050565b600060208284031215614e0557600080fd5b612d1b826149bf565b803560ff811681146149d757600080fd5b60008060008060008060e08789031215614e3857600080fd5b86359550614e4860208801614e0e565b94506080870188811115614e5b57600080fd5b60408801945035925060a0870135614e728161489a565b8092505060c087013590509295509295509295565b600060208284031215614e9957600080fd5b8151612d1b8161489a565b600060208284031215614eb657600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b80820180821115612cb957612cb9614ebd565b600082614f0357634e487b7160e01b600052601260045260246000fd5b500490565b600181815b80851115614f43578160001904821115614f2957614f29614ebd565b80851615614f3657918102915b93841c9390800290614f0d565b509250929050565b600082614f5a57506001612cb9565b81614f6757506000612cb9565b8160018114614f7d5760028114614f8757614fa3565b6001915050612cb9565b60ff841115614f9857614f98614ebd565b50506001821b612cb9565b5060208310610133831016604e8410600b8410161715614fc6575081810a612cb9565b614fd08383614f08565b8060001904821115614fe457614fe4614ebd565b029392505050565b6000612d1b60ff841683614f4b565b8082028115828204841417612cb957612cb9614ebd565b6000602080838503121561502557600080fd5b825167ffffffffffffffff8082111561503d57600080fd5b818501915085601f83011261505157600080fd5b8151818111156150635761506361492c565b8060051b915061507484830161498e565b818152918301840191848101908884111561508e57600080fd5b938501935b838510156150ac57845182529385019390850190615093565b98975050505050505050565b81810381811115612cb957612cb9614ebd565b6000602082840312156150dd57600080fd5b8151612d1b816148e5565b60008183036101008112156150fc57600080fd5b615104614942565b8335815260e0601f198301121561511a57600080fd5b61512261496b565b915061513060208501614e0e565b825260408401356151408161489a565b6020830152615151606085016149bf565b6040830152615162608085016149bf565b606083015260a0840135608083015260c084013560a083015260e084013560c0830152816020820152809250505092915050565b8183823760009101908152919050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b67ffffffffffffffff8181168382160190808211156151f6576151f6614ebd565b5092915050565b634e487b7160e01b600052603260045260246000fd5b600060e08201905085825284602083015267ffffffffffffffff8085511660408401528060208601511660608401528060408601511660808401528060608601511660a0840152506004831061527957634e487b7160e01b600052602160045260246000fd5b8260c083015295945050505050565b60005b838110156152a357818101518382015260200161528b565b50506000910152565b600082516152be818460208701615288565b9190910192915050565b600080600080608085870312156152de57600080fd5b505082516020840151604085015160609095015191969095509092509050565b60008183825b6003811015615323578151835260209283019290910190600101615304565b50505060608201905092915050565b60008451615344818460208901615288565b6001600160f81b031985169083019081528351615368816001840160208801615288565b0160010195945050505050565b6020815260008251806020840152615394816040850160208701615288565b601f01601f19169190910160400192915050565b600080858511156153b857600080fd5b838611156153c557600080fd5b5050820193919092039150565b80356020831015612cb957600019602084900360031b1b1692915050565b60008451615402818460208901615288565b8201838582376000930192835250909392505050565b634e487b7160e01b600052600160045260246000fd5b67ffffffffffffffff8281168282160390808211156151f6576151f6614ebd56fea2646970667358221220a8f5f3072bac3ddf42ee781020ceda6db756fa895529e710389eb0ef7372edf564736f6c63430008110033",
}

// SequencerInboxBlobMockABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerInboxBlobMockMetaData.ABI instead.
var SequencerInboxBlobMockABI = SequencerInboxBlobMockMetaData.ABI

// SequencerInboxBlobMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SequencerInboxBlobMockMetaData.Bin instead.
var SequencerInboxBlobMockBin = SequencerInboxBlobMockMetaData.Bin

// DeploySequencerInboxBlobMock deploys a new Ethereum contract, binding an instance of SequencerInboxBlobMock to it.
func DeploySequencerInboxBlobMock(auth *bind.TransactOpts, backend bind.ContractBackend, maxDataSize_ *big.Int, reader_ common.Address, isUsingFeeToken_ bool, isDelayBufferable_ bool) (common.Address, *types.Transaction, *SequencerInboxBlobMock, error) {
	parsed, err := SequencerInboxBlobMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SequencerInboxBlobMockBin), backend, maxDataSize_, reader_, isUsingFeeToken_, isDelayBufferable_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SequencerInboxBlobMock{SequencerInboxBlobMockCaller: SequencerInboxBlobMockCaller{contract: contract}, SequencerInboxBlobMockTransactor: SequencerInboxBlobMockTransactor{contract: contract}, SequencerInboxBlobMockFilterer: SequencerInboxBlobMockFilterer{contract: contract}}, nil
}

// SequencerInboxBlobMock is an auto generated Go binding around an Ethereum contract.
type SequencerInboxBlobMock struct {
	SequencerInboxBlobMockCaller     // Read-only binding to the contract
	SequencerInboxBlobMockTransactor // Write-only binding to the contract
	SequencerInboxBlobMockFilterer   // Log filterer for contract events
}

// SequencerInboxBlobMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerInboxBlobMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxBlobMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerInboxBlobMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxBlobMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerInboxBlobMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxBlobMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerInboxBlobMockSession struct {
	Contract     *SequencerInboxBlobMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SequencerInboxBlobMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerInboxBlobMockCallerSession struct {
	Contract *SequencerInboxBlobMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// SequencerInboxBlobMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerInboxBlobMockTransactorSession struct {
	Contract     *SequencerInboxBlobMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// SequencerInboxBlobMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerInboxBlobMockRaw struct {
	Contract *SequencerInboxBlobMock // Generic contract binding to access the raw methods on
}

// SequencerInboxBlobMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerInboxBlobMockCallerRaw struct {
	Contract *SequencerInboxBlobMockCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerInboxBlobMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerInboxBlobMockTransactorRaw struct {
	Contract *SequencerInboxBlobMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencerInboxBlobMock creates a new instance of SequencerInboxBlobMock, bound to a specific deployed contract.
func NewSequencerInboxBlobMock(address common.Address, backend bind.ContractBackend) (*SequencerInboxBlobMock, error) {
	contract, err := bindSequencerInboxBlobMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMock{SequencerInboxBlobMockCaller: SequencerInboxBlobMockCaller{contract: contract}, SequencerInboxBlobMockTransactor: SequencerInboxBlobMockTransactor{contract: contract}, SequencerInboxBlobMockFilterer: SequencerInboxBlobMockFilterer{contract: contract}}, nil
}

// NewSequencerInboxBlobMockCaller creates a new read-only instance of SequencerInboxBlobMock, bound to a specific deployed contract.
func NewSequencerInboxBlobMockCaller(address common.Address, caller bind.ContractCaller) (*SequencerInboxBlobMockCaller, error) {
	contract, err := bindSequencerInboxBlobMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockCaller{contract: contract}, nil
}

// NewSequencerInboxBlobMockTransactor creates a new write-only instance of SequencerInboxBlobMock, bound to a specific deployed contract.
func NewSequencerInboxBlobMockTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerInboxBlobMockTransactor, error) {
	contract, err := bindSequencerInboxBlobMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockTransactor{contract: contract}, nil
}

// NewSequencerInboxBlobMockFilterer creates a new log filterer instance of SequencerInboxBlobMock, bound to a specific deployed contract.
func NewSequencerInboxBlobMockFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerInboxBlobMockFilterer, error) {
	contract, err := bindSequencerInboxBlobMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockFilterer{contract: contract}, nil
}

// bindSequencerInboxBlobMock binds a generic wrapper to an already deployed contract.
func bindSequencerInboxBlobMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SequencerInboxBlobMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInboxBlobMock *SequencerInboxBlobMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInboxBlobMock.Contract.SequencerInboxBlobMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInboxBlobMock *SequencerInboxBlobMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SequencerInboxBlobMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInboxBlobMock *SequencerInboxBlobMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SequencerInboxBlobMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInboxBlobMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.contract.Transact(opts, method, params...)
}

// BROTLIMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x16af91a7.
//
// Solidity: function BROTLI_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) BROTLIMESSAGEHEADERFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "BROTLI_MESSAGE_HEADER_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// BROTLIMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x16af91a7.
//
// Solidity: function BROTLI_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) BROTLIMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxBlobMock.Contract.BROTLIMESSAGEHEADERFLAG(&_SequencerInboxBlobMock.CallOpts)
}

// BROTLIMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x16af91a7.
//
// Solidity: function BROTLI_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) BROTLIMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxBlobMock.Contract.BROTLIMESSAGEHEADERFLAG(&_SequencerInboxBlobMock.CallOpts)
}

// CUSTOMDAMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0xdab341a4.
//
// Solidity: function CUSTOM_DA_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) CUSTOMDAMESSAGEHEADERFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "CUSTOM_DA_MESSAGE_HEADER_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// CUSTOMDAMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0xdab341a4.
//
// Solidity: function CUSTOM_DA_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) CUSTOMDAMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxBlobMock.Contract.CUSTOMDAMESSAGEHEADERFLAG(&_SequencerInboxBlobMock.CallOpts)
}

// CUSTOMDAMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0xdab341a4.
//
// Solidity: function CUSTOM_DA_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) CUSTOMDAMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxBlobMock.Contract.CUSTOMDAMESSAGEHEADERFLAG(&_SequencerInboxBlobMock.CallOpts)
}

// DASMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0xf60a5091.
//
// Solidity: function DAS_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) DASMESSAGEHEADERFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "DAS_MESSAGE_HEADER_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// DASMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0xf60a5091.
//
// Solidity: function DAS_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) DASMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxBlobMock.Contract.DASMESSAGEHEADERFLAG(&_SequencerInboxBlobMock.CallOpts)
}

// DASMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0xf60a5091.
//
// Solidity: function DAS_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) DASMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxBlobMock.Contract.DASMESSAGEHEADERFLAG(&_SequencerInboxBlobMock.CallOpts)
}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) DATAAUTHENTICATEDFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "DATA_AUTHENTICATED_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) DATAAUTHENTICATEDFLAG() ([1]byte, error) {
	return _SequencerInboxBlobMock.Contract.DATAAUTHENTICATEDFLAG(&_SequencerInboxBlobMock.CallOpts)
}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) DATAAUTHENTICATEDFLAG() ([1]byte, error) {
	return _SequencerInboxBlobMock.Contract.DATAAUTHENTICATEDFLAG(&_SequencerInboxBlobMock.CallOpts)
}

// DATABLOBHEADERFLAG is a free data retrieval call binding the contract method 0x2cbf74e5.
//
// Solidity: function DATA_BLOB_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) DATABLOBHEADERFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "DATA_BLOB_HEADER_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// DATABLOBHEADERFLAG is a free data retrieval call binding the contract method 0x2cbf74e5.
//
// Solidity: function DATA_BLOB_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) DATABLOBHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxBlobMock.Contract.DATABLOBHEADERFLAG(&_SequencerInboxBlobMock.CallOpts)
}

// DATABLOBHEADERFLAG is a free data retrieval call binding the contract method 0x2cbf74e5.
//
// Solidity: function DATA_BLOB_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) DATABLOBHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxBlobMock.Contract.DATABLOBHEADERFLAG(&_SequencerInboxBlobMock.CallOpts)
}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) HEADERLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "HEADER_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) HEADERLENGTH() (*big.Int, error) {
	return _SequencerInboxBlobMock.Contract.HEADERLENGTH(&_SequencerInboxBlobMock.CallOpts)
}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) HEADERLENGTH() (*big.Int, error) {
	return _SequencerInboxBlobMock.Contract.HEADERLENGTH(&_SequencerInboxBlobMock.CallOpts)
}

// TREEDASMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x6c890450.
//
// Solidity: function TREE_DAS_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) TREEDASMESSAGEHEADERFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "TREE_DAS_MESSAGE_HEADER_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// TREEDASMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x6c890450.
//
// Solidity: function TREE_DAS_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) TREEDASMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxBlobMock.Contract.TREEDASMESSAGEHEADERFLAG(&_SequencerInboxBlobMock.CallOpts)
}

// TREEDASMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x6c890450.
//
// Solidity: function TREE_DAS_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) TREEDASMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxBlobMock.Contract.TREEDASMESSAGEHEADERFLAG(&_SequencerInboxBlobMock.CallOpts)
}

// ZEROHEAVYMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x02c99275.
//
// Solidity: function ZERO_HEAVY_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) ZEROHEAVYMESSAGEHEADERFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "ZERO_HEAVY_MESSAGE_HEADER_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// ZEROHEAVYMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x02c99275.
//
// Solidity: function ZERO_HEAVY_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) ZEROHEAVYMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxBlobMock.Contract.ZEROHEAVYMESSAGEHEADERFLAG(&_SequencerInboxBlobMock.CallOpts)
}

// ZEROHEAVYMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x02c99275.
//
// Solidity: function ZERO_HEAVY_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) ZEROHEAVYMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxBlobMock.Contract.ZEROHEAVYMESSAGEHEADERFLAG(&_SequencerInboxBlobMock.CallOpts)
}

// AddSequencerL2BatchFromOrigin6f12b0c9 is a free data retrieval call binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 , bytes , uint256 , address ) pure returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) AddSequencerL2BatchFromOrigin6f12b0c9(opts *bind.CallOpts, arg0 *big.Int, arg1 []byte, arg2 *big.Int, arg3 common.Address) error {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "addSequencerL2BatchFromOrigin", arg0, arg1, arg2, arg3)

	if err != nil {
		return err
	}

	return err

}

// AddSequencerL2BatchFromOrigin6f12b0c9 is a free data retrieval call binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 , bytes , uint256 , address ) pure returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) AddSequencerL2BatchFromOrigin6f12b0c9(arg0 *big.Int, arg1 []byte, arg2 *big.Int, arg3 common.Address) error {
	return _SequencerInboxBlobMock.Contract.AddSequencerL2BatchFromOrigin6f12b0c9(&_SequencerInboxBlobMock.CallOpts, arg0, arg1, arg2, arg3)
}

// AddSequencerL2BatchFromOrigin6f12b0c9 is a free data retrieval call binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 , bytes , uint256 , address ) pure returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) AddSequencerL2BatchFromOrigin6f12b0c9(arg0 *big.Int, arg1 []byte, arg2 *big.Int, arg3 common.Address) error {
	return _SequencerInboxBlobMock.Contract.AddSequencerL2BatchFromOrigin6f12b0c9(&_SequencerInboxBlobMock.CallOpts, arg0, arg1, arg2, arg3)
}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) BatchCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "batchCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) BatchCount() (*big.Int, error) {
	return _SequencerInboxBlobMock.Contract.BatchCount(&_SequencerInboxBlobMock.CallOpts)
}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) BatchCount() (*big.Int, error) {
	return _SequencerInboxBlobMock.Contract.BatchCount(&_SequencerInboxBlobMock.CallOpts)
}

// BatchPosterManager is a free data retrieval call binding the contract method 0xcc2a1a0c.
//
// Solidity: function batchPosterManager() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) BatchPosterManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "batchPosterManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BatchPosterManager is a free data retrieval call binding the contract method 0xcc2a1a0c.
//
// Solidity: function batchPosterManager() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) BatchPosterManager() (common.Address, error) {
	return _SequencerInboxBlobMock.Contract.BatchPosterManager(&_SequencerInboxBlobMock.CallOpts)
}

// BatchPosterManager is a free data retrieval call binding the contract method 0xcc2a1a0c.
//
// Solidity: function batchPosterManager() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) BatchPosterManager() (common.Address, error) {
	return _SequencerInboxBlobMock.Contract.BatchPosterManager(&_SequencerInboxBlobMock.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) Bridge() (common.Address, error) {
	return _SequencerInboxBlobMock.Contract.Bridge(&_SequencerInboxBlobMock.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) Bridge() (common.Address, error) {
	return _SequencerInboxBlobMock.Contract.Bridge(&_SequencerInboxBlobMock.CallOpts)
}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint64 bufferBlocks, uint64 max, uint64 threshold, uint64 prevBlockNumber, uint64 replenishRateInBasis, uint64 prevSequencedBlockNumber)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) Buffer(opts *bind.CallOpts) (struct {
	BufferBlocks             uint64
	Max                      uint64
	Threshold                uint64
	PrevBlockNumber          uint64
	ReplenishRateInBasis     uint64
	PrevSequencedBlockNumber uint64
}, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "buffer")

	outstruct := new(struct {
		BufferBlocks             uint64
		Max                      uint64
		Threshold                uint64
		PrevBlockNumber          uint64
		ReplenishRateInBasis     uint64
		PrevSequencedBlockNumber uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BufferBlocks = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Max = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Threshold = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.PrevBlockNumber = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.ReplenishRateInBasis = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.PrevSequencedBlockNumber = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint64 bufferBlocks, uint64 max, uint64 threshold, uint64 prevBlockNumber, uint64 replenishRateInBasis, uint64 prevSequencedBlockNumber)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) Buffer() (struct {
	BufferBlocks             uint64
	Max                      uint64
	Threshold                uint64
	PrevBlockNumber          uint64
	ReplenishRateInBasis     uint64
	PrevSequencedBlockNumber uint64
}, error) {
	return _SequencerInboxBlobMock.Contract.Buffer(&_SequencerInboxBlobMock.CallOpts)
}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint64 bufferBlocks, uint64 max, uint64 threshold, uint64 prevBlockNumber, uint64 replenishRateInBasis, uint64 prevSequencedBlockNumber)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) Buffer() (struct {
	BufferBlocks             uint64
	Max                      uint64
	Threshold                uint64
	PrevBlockNumber          uint64
	ReplenishRateInBasis     uint64
	PrevSequencedBlockNumber uint64
}, error) {
	return _SequencerInboxBlobMock.Contract.Buffer(&_SequencerInboxBlobMock.CallOpts)
}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(bool isValidKeyset, uint64 creationBlock)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) DasKeySetInfo(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "dasKeySetInfo", arg0)

	outstruct := new(struct {
		IsValidKeyset bool
		CreationBlock uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsValidKeyset = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CreationBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(bool isValidKeyset, uint64 creationBlock)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) DasKeySetInfo(arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	return _SequencerInboxBlobMock.Contract.DasKeySetInfo(&_SequencerInboxBlobMock.CallOpts, arg0)
}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(bool isValidKeyset, uint64 creationBlock)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) DasKeySetInfo(arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	return _SequencerInboxBlobMock.Contract.DasKeySetInfo(&_SequencerInboxBlobMock.CallOpts, arg0)
}

// FeeTokenPricer is a free data retrieval call binding the contract method 0x22291e8d.
//
// Solidity: function feeTokenPricer() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) FeeTokenPricer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "feeTokenPricer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenPricer is a free data retrieval call binding the contract method 0x22291e8d.
//
// Solidity: function feeTokenPricer() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) FeeTokenPricer() (common.Address, error) {
	return _SequencerInboxBlobMock.Contract.FeeTokenPricer(&_SequencerInboxBlobMock.CallOpts)
}

// FeeTokenPricer is a free data retrieval call binding the contract method 0x22291e8d.
//
// Solidity: function feeTokenPricer() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) FeeTokenPricer() (common.Address, error) {
	return _SequencerInboxBlobMock.Contract.FeeTokenPricer(&_SequencerInboxBlobMock.CallOpts)
}

// ForceInclusionDeadline is a free data retrieval call binding the contract method 0xdd44e6e0.
//
// Solidity: function forceInclusionDeadline(uint64 blockNumber) view returns(uint64)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) ForceInclusionDeadline(opts *bind.CallOpts, blockNumber uint64) (uint64, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "forceInclusionDeadline", blockNumber)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceInclusionDeadline is a free data retrieval call binding the contract method 0xdd44e6e0.
//
// Solidity: function forceInclusionDeadline(uint64 blockNumber) view returns(uint64)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) ForceInclusionDeadline(blockNumber uint64) (uint64, error) {
	return _SequencerInboxBlobMock.Contract.ForceInclusionDeadline(&_SequencerInboxBlobMock.CallOpts, blockNumber)
}

// ForceInclusionDeadline is a free data retrieval call binding the contract method 0xdd44e6e0.
//
// Solidity: function forceInclusionDeadline(uint64 blockNumber) view returns(uint64)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) ForceInclusionDeadline(blockNumber uint64) (uint64, error) {
	return _SequencerInboxBlobMock.Contract.ForceInclusionDeadline(&_SequencerInboxBlobMock.CallOpts, blockNumber)
}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) GetKeysetCreationBlock(opts *bind.CallOpts, ksHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "getKeysetCreationBlock", ksHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) GetKeysetCreationBlock(ksHash [32]byte) (*big.Int, error) {
	return _SequencerInboxBlobMock.Contract.GetKeysetCreationBlock(&_SequencerInboxBlobMock.CallOpts, ksHash)
}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) GetKeysetCreationBlock(ksHash [32]byte) (*big.Int, error) {
	return _SequencerInboxBlobMock.Contract.GetKeysetCreationBlock(&_SequencerInboxBlobMock.CallOpts, ksHash)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) InboxAccs(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "inboxAccs", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) InboxAccs(index *big.Int) ([32]byte, error) {
	return _SequencerInboxBlobMock.Contract.InboxAccs(&_SequencerInboxBlobMock.CallOpts, index)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) InboxAccs(index *big.Int) ([32]byte, error) {
	return _SequencerInboxBlobMock.Contract.InboxAccs(&_SequencerInboxBlobMock.CallOpts, index)
}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) IsBatchPoster(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "isBatchPoster", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) IsBatchPoster(arg0 common.Address) (bool, error) {
	return _SequencerInboxBlobMock.Contract.IsBatchPoster(&_SequencerInboxBlobMock.CallOpts, arg0)
}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) IsBatchPoster(arg0 common.Address) (bool, error) {
	return _SequencerInboxBlobMock.Contract.IsBatchPoster(&_SequencerInboxBlobMock.CallOpts, arg0)
}

// IsDelayBufferable is a free data retrieval call binding the contract method 0x4b678a66.
//
// Solidity: function isDelayBufferable() view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) IsDelayBufferable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "isDelayBufferable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelayBufferable is a free data retrieval call binding the contract method 0x4b678a66.
//
// Solidity: function isDelayBufferable() view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) IsDelayBufferable() (bool, error) {
	return _SequencerInboxBlobMock.Contract.IsDelayBufferable(&_SequencerInboxBlobMock.CallOpts)
}

// IsDelayBufferable is a free data retrieval call binding the contract method 0x4b678a66.
//
// Solidity: function isDelayBufferable() view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) IsDelayBufferable() (bool, error) {
	return _SequencerInboxBlobMock.Contract.IsDelayBufferable(&_SequencerInboxBlobMock.CallOpts)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) IsSequencer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "isSequencer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) IsSequencer(arg0 common.Address) (bool, error) {
	return _SequencerInboxBlobMock.Contract.IsSequencer(&_SequencerInboxBlobMock.CallOpts, arg0)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) IsSequencer(arg0 common.Address) (bool, error) {
	return _SequencerInboxBlobMock.Contract.IsSequencer(&_SequencerInboxBlobMock.CallOpts, arg0)
}

// IsUsingFeeToken is a free data retrieval call binding the contract method 0x92d9f782.
//
// Solidity: function isUsingFeeToken() view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) IsUsingFeeToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "isUsingFeeToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUsingFeeToken is a free data retrieval call binding the contract method 0x92d9f782.
//
// Solidity: function isUsingFeeToken() view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) IsUsingFeeToken() (bool, error) {
	return _SequencerInboxBlobMock.Contract.IsUsingFeeToken(&_SequencerInboxBlobMock.CallOpts)
}

// IsUsingFeeToken is a free data retrieval call binding the contract method 0x92d9f782.
//
// Solidity: function isUsingFeeToken() view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) IsUsingFeeToken() (bool, error) {
	return _SequencerInboxBlobMock.Contract.IsUsingFeeToken(&_SequencerInboxBlobMock.CallOpts)
}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) IsValidKeysetHash(opts *bind.CallOpts, ksHash [32]byte) (bool, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "isValidKeysetHash", ksHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) IsValidKeysetHash(ksHash [32]byte) (bool, error) {
	return _SequencerInboxBlobMock.Contract.IsValidKeysetHash(&_SequencerInboxBlobMock.CallOpts, ksHash)
}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) IsValidKeysetHash(ksHash [32]byte) (bool, error) {
	return _SequencerInboxBlobMock.Contract.IsValidKeysetHash(&_SequencerInboxBlobMock.CallOpts, ksHash)
}

// MaxDataSize is a free data retrieval call binding the contract method 0xe8eb1dc3.
//
// Solidity: function maxDataSize() view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) MaxDataSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "maxDataSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDataSize is a free data retrieval call binding the contract method 0xe8eb1dc3.
//
// Solidity: function maxDataSize() view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) MaxDataSize() (*big.Int, error) {
	return _SequencerInboxBlobMock.Contract.MaxDataSize(&_SequencerInboxBlobMock.CallOpts)
}

// MaxDataSize is a free data retrieval call binding the contract method 0xe8eb1dc3.
//
// Solidity: function maxDataSize() view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) MaxDataSize() (*big.Int, error) {
	return _SequencerInboxBlobMock.Contract.MaxDataSize(&_SequencerInboxBlobMock.CallOpts)
}

// MaxTimeVariation is a free data retrieval call binding the contract method 0xebea461d.
//
// Solidity: function maxTimeVariation() view returns(uint256, uint256, uint256, uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) MaxTimeVariation(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "maxTimeVariation")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// MaxTimeVariation is a free data retrieval call binding the contract method 0xebea461d.
//
// Solidity: function maxTimeVariation() view returns(uint256, uint256, uint256, uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) MaxTimeVariation() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _SequencerInboxBlobMock.Contract.MaxTimeVariation(&_SequencerInboxBlobMock.CallOpts)
}

// MaxTimeVariation is a free data retrieval call binding the contract method 0xebea461d.
//
// Solidity: function maxTimeVariation() view returns(uint256, uint256, uint256, uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) MaxTimeVariation() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _SequencerInboxBlobMock.Contract.MaxTimeVariation(&_SequencerInboxBlobMock.CallOpts)
}

// Reader4844 is a free data retrieval call binding the contract method 0x8d910dde.
//
// Solidity: function reader4844() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) Reader4844(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "reader4844")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Reader4844 is a free data retrieval call binding the contract method 0x8d910dde.
//
// Solidity: function reader4844() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) Reader4844() (common.Address, error) {
	return _SequencerInboxBlobMock.Contract.Reader4844(&_SequencerInboxBlobMock.CallOpts)
}

// Reader4844 is a free data retrieval call binding the contract method 0x8d910dde.
//
// Solidity: function reader4844() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) Reader4844() (common.Address, error) {
	return _SequencerInboxBlobMock.Contract.Reader4844(&_SequencerInboxBlobMock.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) Rollup() (common.Address, error) {
	return _SequencerInboxBlobMock.Contract.Rollup(&_SequencerInboxBlobMock.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) Rollup() (common.Address, error) {
	return _SequencerInboxBlobMock.Contract.Rollup(&_SequencerInboxBlobMock.CallOpts)
}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCaller) TotalDelayedMessagesRead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxBlobMock.contract.Call(opts, &out, "totalDelayedMessagesRead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) TotalDelayedMessagesRead() (*big.Int, error) {
	return _SequencerInboxBlobMock.Contract.TotalDelayedMessagesRead(&_SequencerInboxBlobMock.CallOpts)
}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockCallerSession) TotalDelayedMessagesRead() (*big.Int, error) {
	return _SequencerInboxBlobMock.Contract.TotalDelayedMessagesRead(&_SequencerInboxBlobMock.CallOpts)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0xe0bc9729.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) AddSequencerL2Batch(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "addSequencerL2Batch", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0xe0bc9729.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) AddSequencerL2Batch(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.AddSequencerL2Batch(&_SequencerInboxBlobMock.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0xe0bc9729.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) AddSequencerL2Batch(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.AddSequencerL2Batch(&_SequencerInboxBlobMock.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchDelayProof is a paid mutator transaction binding the contract method 0x6e620055.
//
// Solidity: function addSequencerL2BatchDelayProof(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) AddSequencerL2BatchDelayProof(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "addSequencerL2BatchDelayProof", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchDelayProof is a paid mutator transaction binding the contract method 0x6e620055.
//
// Solidity: function addSequencerL2BatchDelayProof(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) AddSequencerL2BatchDelayProof(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.AddSequencerL2BatchDelayProof(&_SequencerInboxBlobMock.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchDelayProof is a paid mutator transaction binding the contract method 0x6e620055.
//
// Solidity: function addSequencerL2BatchDelayProof(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) AddSequencerL2BatchDelayProof(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.AddSequencerL2BatchDelayProof(&_SequencerInboxBlobMock.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchFromBlobs is a paid mutator transaction binding the contract method 0x3e5aa082.
//
// Solidity: function addSequencerL2BatchFromBlobs(uint256 sequenceNumber, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) AddSequencerL2BatchFromBlobs(opts *bind.TransactOpts, sequenceNumber *big.Int, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "addSequencerL2BatchFromBlobs", sequenceNumber, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromBlobs is a paid mutator transaction binding the contract method 0x3e5aa082.
//
// Solidity: function addSequencerL2BatchFromBlobs(uint256 sequenceNumber, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) AddSequencerL2BatchFromBlobs(sequenceNumber *big.Int, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.AddSequencerL2BatchFromBlobs(&_SequencerInboxBlobMock.TransactOpts, sequenceNumber, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromBlobs is a paid mutator transaction binding the contract method 0x3e5aa082.
//
// Solidity: function addSequencerL2BatchFromBlobs(uint256 sequenceNumber, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) AddSequencerL2BatchFromBlobs(sequenceNumber *big.Int, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.AddSequencerL2BatchFromBlobs(&_SequencerInboxBlobMock.TransactOpts, sequenceNumber, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromBlobsDelayProof is a paid mutator transaction binding the contract method 0x917cf8ac.
//
// Solidity: function addSequencerL2BatchFromBlobsDelayProof(uint256 sequenceNumber, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) AddSequencerL2BatchFromBlobsDelayProof(opts *bind.TransactOpts, sequenceNumber *big.Int, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "addSequencerL2BatchFromBlobsDelayProof", sequenceNumber, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchFromBlobsDelayProof is a paid mutator transaction binding the contract method 0x917cf8ac.
//
// Solidity: function addSequencerL2BatchFromBlobsDelayProof(uint256 sequenceNumber, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) AddSequencerL2BatchFromBlobsDelayProof(sequenceNumber *big.Int, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.AddSequencerL2BatchFromBlobsDelayProof(&_SequencerInboxBlobMock.TransactOpts, sequenceNumber, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchFromBlobsDelayProof is a paid mutator transaction binding the contract method 0x917cf8ac.
//
// Solidity: function addSequencerL2BatchFromBlobsDelayProof(uint256 sequenceNumber, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) AddSequencerL2BatchFromBlobsDelayProof(sequenceNumber *big.Int, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.AddSequencerL2BatchFromBlobsDelayProof(&_SequencerInboxBlobMock.TransactOpts, sequenceNumber, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchFromOrigin8f111f3c is a paid mutator transaction binding the contract method 0x8f111f3c.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) AddSequencerL2BatchFromOrigin8f111f3c(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "addSequencerL2BatchFromOrigin0", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromOrigin8f111f3c is a paid mutator transaction binding the contract method 0x8f111f3c.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) AddSequencerL2BatchFromOrigin8f111f3c(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.AddSequencerL2BatchFromOrigin8f111f3c(&_SequencerInboxBlobMock.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromOrigin8f111f3c is a paid mutator transaction binding the contract method 0x8f111f3c.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) AddSequencerL2BatchFromOrigin8f111f3c(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.AddSequencerL2BatchFromOrigin8f111f3c(&_SequencerInboxBlobMock.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromOriginDelayProof is a paid mutator transaction binding the contract method 0x69cacded.
//
// Solidity: function addSequencerL2BatchFromOriginDelayProof(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) AddSequencerL2BatchFromOriginDelayProof(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "addSequencerL2BatchFromOriginDelayProof", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchFromOriginDelayProof is a paid mutator transaction binding the contract method 0x69cacded.
//
// Solidity: function addSequencerL2BatchFromOriginDelayProof(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) AddSequencerL2BatchFromOriginDelayProof(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.AddSequencerL2BatchFromOriginDelayProof(&_SequencerInboxBlobMock.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchFromOriginDelayProof is a paid mutator transaction binding the contract method 0x69cacded.
//
// Solidity: function addSequencerL2BatchFromOriginDelayProof(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) AddSequencerL2BatchFromOriginDelayProof(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.AddSequencerL2BatchFromOriginDelayProof(&_SequencerInboxBlobMock.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) ForceInclusion(opts *bind.TransactOpts, _totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "forceInclusion", _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.ForceInclusion(&_SequencerInboxBlobMock.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.ForceInclusion(&_SequencerInboxBlobMock.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xa84840b7.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_, (uint64,uint64,uint64) bufferConfig_, address feeTokenPricer_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) Initialize(opts *bind.TransactOpts, bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation, bufferConfig_ BufferConfig, feeTokenPricer_ common.Address) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "initialize", bridge_, maxTimeVariation_, bufferConfig_, feeTokenPricer_)
}

// Initialize is a paid mutator transaction binding the contract method 0xa84840b7.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_, (uint64,uint64,uint64) bufferConfig_, address feeTokenPricer_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) Initialize(bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation, bufferConfig_ BufferConfig, feeTokenPricer_ common.Address) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.Initialize(&_SequencerInboxBlobMock.TransactOpts, bridge_, maxTimeVariation_, bufferConfig_, feeTokenPricer_)
}

// Initialize is a paid mutator transaction binding the contract method 0xa84840b7.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_, (uint64,uint64,uint64) bufferConfig_, address feeTokenPricer_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) Initialize(bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation, bufferConfig_ BufferConfig, feeTokenPricer_ common.Address) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.Initialize(&_SequencerInboxBlobMock.TransactOpts, bridge_, maxTimeVariation_, bufferConfig_, feeTokenPricer_)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) InvalidateKeysetHash(opts *bind.TransactOpts, ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "invalidateKeysetHash", ksHash)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) InvalidateKeysetHash(ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.InvalidateKeysetHash(&_SequencerInboxBlobMock.TransactOpts, ksHash)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) InvalidateKeysetHash(ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.InvalidateKeysetHash(&_SequencerInboxBlobMock.TransactOpts, ksHash)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xa655d937.
//
// Solidity: function postUpgradeInit((uint64,uint64,uint64) bufferConfig_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) PostUpgradeInit(opts *bind.TransactOpts, bufferConfig_ BufferConfig) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "postUpgradeInit", bufferConfig_)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xa655d937.
//
// Solidity: function postUpgradeInit((uint64,uint64,uint64) bufferConfig_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) PostUpgradeInit(bufferConfig_ BufferConfig) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.PostUpgradeInit(&_SequencerInboxBlobMock.TransactOpts, bufferConfig_)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xa655d937.
//
// Solidity: function postUpgradeInit((uint64,uint64,uint64) bufferConfig_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) PostUpgradeInit(bufferConfig_ BufferConfig) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.PostUpgradeInit(&_SequencerInboxBlobMock.TransactOpts, bufferConfig_)
}

// RemoveDelayAfterFork is a paid mutator transaction binding the contract method 0x96cc5c78.
//
// Solidity: function removeDelayAfterFork() returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) RemoveDelayAfterFork(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "removeDelayAfterFork")
}

// RemoveDelayAfterFork is a paid mutator transaction binding the contract method 0x96cc5c78.
//
// Solidity: function removeDelayAfterFork() returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) RemoveDelayAfterFork() (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.RemoveDelayAfterFork(&_SequencerInboxBlobMock.TransactOpts)
}

// RemoveDelayAfterFork is a paid mutator transaction binding the contract method 0x96cc5c78.
//
// Solidity: function removeDelayAfterFork() returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) RemoveDelayAfterFork() (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.RemoveDelayAfterFork(&_SequencerInboxBlobMock.TransactOpts)
}

// SetBatchPosterManager is a paid mutator transaction binding the contract method 0x1ff64790.
//
// Solidity: function setBatchPosterManager(address newBatchPosterManager) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) SetBatchPosterManager(opts *bind.TransactOpts, newBatchPosterManager common.Address) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "setBatchPosterManager", newBatchPosterManager)
}

// SetBatchPosterManager is a paid mutator transaction binding the contract method 0x1ff64790.
//
// Solidity: function setBatchPosterManager(address newBatchPosterManager) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) SetBatchPosterManager(newBatchPosterManager common.Address) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SetBatchPosterManager(&_SequencerInboxBlobMock.TransactOpts, newBatchPosterManager)
}

// SetBatchPosterManager is a paid mutator transaction binding the contract method 0x1ff64790.
//
// Solidity: function setBatchPosterManager(address newBatchPosterManager) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) SetBatchPosterManager(newBatchPosterManager common.Address) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SetBatchPosterManager(&_SequencerInboxBlobMock.TransactOpts, newBatchPosterManager)
}

// SetBufferConfig is a paid mutator transaction binding the contract method 0x2f3985a7.
//
// Solidity: function setBufferConfig((uint64,uint64,uint64) bufferConfig_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) SetBufferConfig(opts *bind.TransactOpts, bufferConfig_ BufferConfig) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "setBufferConfig", bufferConfig_)
}

// SetBufferConfig is a paid mutator transaction binding the contract method 0x2f3985a7.
//
// Solidity: function setBufferConfig((uint64,uint64,uint64) bufferConfig_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) SetBufferConfig(bufferConfig_ BufferConfig) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SetBufferConfig(&_SequencerInboxBlobMock.TransactOpts, bufferConfig_)
}

// SetBufferConfig is a paid mutator transaction binding the contract method 0x2f3985a7.
//
// Solidity: function setBufferConfig((uint64,uint64,uint64) bufferConfig_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) SetBufferConfig(bufferConfig_ BufferConfig) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SetBufferConfig(&_SequencerInboxBlobMock.TransactOpts, bufferConfig_)
}

// SetFeeTokenPricer is a paid mutator transaction binding the contract method 0x036f7ed3.
//
// Solidity: function setFeeTokenPricer(address feeTokenPricer_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) SetFeeTokenPricer(opts *bind.TransactOpts, feeTokenPricer_ common.Address) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "setFeeTokenPricer", feeTokenPricer_)
}

// SetFeeTokenPricer is a paid mutator transaction binding the contract method 0x036f7ed3.
//
// Solidity: function setFeeTokenPricer(address feeTokenPricer_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) SetFeeTokenPricer(feeTokenPricer_ common.Address) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SetFeeTokenPricer(&_SequencerInboxBlobMock.TransactOpts, feeTokenPricer_)
}

// SetFeeTokenPricer is a paid mutator transaction binding the contract method 0x036f7ed3.
//
// Solidity: function setFeeTokenPricer(address feeTokenPricer_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) SetFeeTokenPricer(feeTokenPricer_ common.Address) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SetFeeTokenPricer(&_SequencerInboxBlobMock.TransactOpts, feeTokenPricer_)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) SetIsBatchPoster(opts *bind.TransactOpts, addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "setIsBatchPoster", addr, isBatchPoster_)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) SetIsBatchPoster(addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SetIsBatchPoster(&_SequencerInboxBlobMock.TransactOpts, addr, isBatchPoster_)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) SetIsBatchPoster(addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SetIsBatchPoster(&_SequencerInboxBlobMock.TransactOpts, addr, isBatchPoster_)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool isSequencer_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) SetIsSequencer(opts *bind.TransactOpts, addr common.Address, isSequencer_ bool) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "setIsSequencer", addr, isSequencer_)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool isSequencer_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) SetIsSequencer(addr common.Address, isSequencer_ bool) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SetIsSequencer(&_SequencerInboxBlobMock.TransactOpts, addr, isSequencer_)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool isSequencer_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) SetIsSequencer(addr common.Address, isSequencer_ bool) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SetIsSequencer(&_SequencerInboxBlobMock.TransactOpts, addr, isSequencer_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) SetMaxTimeVariation(opts *bind.TransactOpts, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "setMaxTimeVariation", maxTimeVariation_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) SetMaxTimeVariation(maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SetMaxTimeVariation(&_SequencerInboxBlobMock.TransactOpts, maxTimeVariation_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) SetMaxTimeVariation(maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SetMaxTimeVariation(&_SequencerInboxBlobMock.TransactOpts, maxTimeVariation_)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) SetValidKeyset(opts *bind.TransactOpts, keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "setValidKeyset", keysetBytes)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) SetValidKeyset(keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SetValidKeyset(&_SequencerInboxBlobMock.TransactOpts, keysetBytes)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) SetValidKeyset(keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.SetValidKeyset(&_SequencerInboxBlobMock.TransactOpts, keysetBytes)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactor) UpdateRollupAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInboxBlobMock.contract.Transact(opts, "updateRollupAddress")
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockSession) UpdateRollupAddress() (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.UpdateRollupAddress(&_SequencerInboxBlobMock.TransactOpts)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_SequencerInboxBlobMock *SequencerInboxBlobMockTransactorSession) UpdateRollupAddress() (*types.Transaction, error) {
	return _SequencerInboxBlobMock.Contract.UpdateRollupAddress(&_SequencerInboxBlobMock.TransactOpts)
}

// SequencerInboxBlobMockBatchPosterManagerSetIterator is returned from FilterBatchPosterManagerSet and is used to iterate over the raw logs and unpacked data for BatchPosterManagerSet events raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockBatchPosterManagerSetIterator struct {
	Event *SequencerInboxBlobMockBatchPosterManagerSet // Event containing the contract specifics and raw log

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
func (it *SequencerInboxBlobMockBatchPosterManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxBlobMockBatchPosterManagerSet)
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
		it.Event = new(SequencerInboxBlobMockBatchPosterManagerSet)
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
func (it *SequencerInboxBlobMockBatchPosterManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxBlobMockBatchPosterManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxBlobMockBatchPosterManagerSet represents a BatchPosterManagerSet event raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockBatchPosterManagerSet struct {
	NewBatchPosterManager common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterBatchPosterManagerSet is a free log retrieval operation binding the contract event 0x3cd6c184800297a0f2b00926a683cbe76890bb7fd01480ac0a10ed6c8f7f6659.
//
// Solidity: event BatchPosterManagerSet(address newBatchPosterManager)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) FilterBatchPosterManagerSet(opts *bind.FilterOpts) (*SequencerInboxBlobMockBatchPosterManagerSetIterator, error) {

	logs, sub, err := _SequencerInboxBlobMock.contract.FilterLogs(opts, "BatchPosterManagerSet")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockBatchPosterManagerSetIterator{contract: _SequencerInboxBlobMock.contract, event: "BatchPosterManagerSet", logs: logs, sub: sub}, nil
}

// WatchBatchPosterManagerSet is a free log subscription operation binding the contract event 0x3cd6c184800297a0f2b00926a683cbe76890bb7fd01480ac0a10ed6c8f7f6659.
//
// Solidity: event BatchPosterManagerSet(address newBatchPosterManager)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) WatchBatchPosterManagerSet(opts *bind.WatchOpts, sink chan<- *SequencerInboxBlobMockBatchPosterManagerSet) (event.Subscription, error) {

	logs, sub, err := _SequencerInboxBlobMock.contract.WatchLogs(opts, "BatchPosterManagerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxBlobMockBatchPosterManagerSet)
				if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "BatchPosterManagerSet", log); err != nil {
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

// ParseBatchPosterManagerSet is a log parse operation binding the contract event 0x3cd6c184800297a0f2b00926a683cbe76890bb7fd01480ac0a10ed6c8f7f6659.
//
// Solidity: event BatchPosterManagerSet(address newBatchPosterManager)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) ParseBatchPosterManagerSet(log types.Log) (*SequencerInboxBlobMockBatchPosterManagerSet, error) {
	event := new(SequencerInboxBlobMockBatchPosterManagerSet)
	if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "BatchPosterManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxBlobMockBatchPosterSetIterator is returned from FilterBatchPosterSet and is used to iterate over the raw logs and unpacked data for BatchPosterSet events raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockBatchPosterSetIterator struct {
	Event *SequencerInboxBlobMockBatchPosterSet // Event containing the contract specifics and raw log

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
func (it *SequencerInboxBlobMockBatchPosterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxBlobMockBatchPosterSet)
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
		it.Event = new(SequencerInboxBlobMockBatchPosterSet)
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
func (it *SequencerInboxBlobMockBatchPosterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxBlobMockBatchPosterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxBlobMockBatchPosterSet represents a BatchPosterSet event raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockBatchPosterSet struct {
	BatchPoster   common.Address
	IsBatchPoster bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBatchPosterSet is a free log retrieval operation binding the contract event 0x28bcc5626d357efe966b4b0876aa1ee8ab99e26da4f131f6a2623f1800701c21.
//
// Solidity: event BatchPosterSet(address batchPoster, bool isBatchPoster)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) FilterBatchPosterSet(opts *bind.FilterOpts) (*SequencerInboxBlobMockBatchPosterSetIterator, error) {

	logs, sub, err := _SequencerInboxBlobMock.contract.FilterLogs(opts, "BatchPosterSet")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockBatchPosterSetIterator{contract: _SequencerInboxBlobMock.contract, event: "BatchPosterSet", logs: logs, sub: sub}, nil
}

// WatchBatchPosterSet is a free log subscription operation binding the contract event 0x28bcc5626d357efe966b4b0876aa1ee8ab99e26da4f131f6a2623f1800701c21.
//
// Solidity: event BatchPosterSet(address batchPoster, bool isBatchPoster)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) WatchBatchPosterSet(opts *bind.WatchOpts, sink chan<- *SequencerInboxBlobMockBatchPosterSet) (event.Subscription, error) {

	logs, sub, err := _SequencerInboxBlobMock.contract.WatchLogs(opts, "BatchPosterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxBlobMockBatchPosterSet)
				if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "BatchPosterSet", log); err != nil {
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

// ParseBatchPosterSet is a log parse operation binding the contract event 0x28bcc5626d357efe966b4b0876aa1ee8ab99e26da4f131f6a2623f1800701c21.
//
// Solidity: event BatchPosterSet(address batchPoster, bool isBatchPoster)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) ParseBatchPosterSet(log types.Log) (*SequencerInboxBlobMockBatchPosterSet, error) {
	event := new(SequencerInboxBlobMockBatchPosterSet)
	if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "BatchPosterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxBlobMockBufferConfigSetIterator is returned from FilterBufferConfigSet and is used to iterate over the raw logs and unpacked data for BufferConfigSet events raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockBufferConfigSetIterator struct {
	Event *SequencerInboxBlobMockBufferConfigSet // Event containing the contract specifics and raw log

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
func (it *SequencerInboxBlobMockBufferConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxBlobMockBufferConfigSet)
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
		it.Event = new(SequencerInboxBlobMockBufferConfigSet)
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
func (it *SequencerInboxBlobMockBufferConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxBlobMockBufferConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxBlobMockBufferConfigSet represents a BufferConfigSet event raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockBufferConfigSet struct {
	BufferConfig BufferConfig
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBufferConfigSet is a free log retrieval operation binding the contract event 0xaa7a2d8175dee3b637814ad6346005dfcc357165396fb8327f649effe8abcf85.
//
// Solidity: event BufferConfigSet((uint64,uint64,uint64) bufferConfig)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) FilterBufferConfigSet(opts *bind.FilterOpts) (*SequencerInboxBlobMockBufferConfigSetIterator, error) {

	logs, sub, err := _SequencerInboxBlobMock.contract.FilterLogs(opts, "BufferConfigSet")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockBufferConfigSetIterator{contract: _SequencerInboxBlobMock.contract, event: "BufferConfigSet", logs: logs, sub: sub}, nil
}

// WatchBufferConfigSet is a free log subscription operation binding the contract event 0xaa7a2d8175dee3b637814ad6346005dfcc357165396fb8327f649effe8abcf85.
//
// Solidity: event BufferConfigSet((uint64,uint64,uint64) bufferConfig)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) WatchBufferConfigSet(opts *bind.WatchOpts, sink chan<- *SequencerInboxBlobMockBufferConfigSet) (event.Subscription, error) {

	logs, sub, err := _SequencerInboxBlobMock.contract.WatchLogs(opts, "BufferConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxBlobMockBufferConfigSet)
				if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "BufferConfigSet", log); err != nil {
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

// ParseBufferConfigSet is a log parse operation binding the contract event 0xaa7a2d8175dee3b637814ad6346005dfcc357165396fb8327f649effe8abcf85.
//
// Solidity: event BufferConfigSet((uint64,uint64,uint64) bufferConfig)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) ParseBufferConfigSet(log types.Log) (*SequencerInboxBlobMockBufferConfigSet, error) {
	event := new(SequencerInboxBlobMockBufferConfigSet)
	if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "BufferConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxBlobMockFeeTokenPricerSetIterator is returned from FilterFeeTokenPricerSet and is used to iterate over the raw logs and unpacked data for FeeTokenPricerSet events raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockFeeTokenPricerSetIterator struct {
	Event *SequencerInboxBlobMockFeeTokenPricerSet // Event containing the contract specifics and raw log

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
func (it *SequencerInboxBlobMockFeeTokenPricerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxBlobMockFeeTokenPricerSet)
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
		it.Event = new(SequencerInboxBlobMockFeeTokenPricerSet)
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
func (it *SequencerInboxBlobMockFeeTokenPricerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxBlobMockFeeTokenPricerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxBlobMockFeeTokenPricerSet represents a FeeTokenPricerSet event raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockFeeTokenPricerSet struct {
	FeeTokenPricer common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFeeTokenPricerSet is a free log retrieval operation binding the contract event 0xe83d6153add50e41b8ee6c1115c4178687349bb12bc3902a50b1f6ad78a0c541.
//
// Solidity: event FeeTokenPricerSet(address feeTokenPricer)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) FilterFeeTokenPricerSet(opts *bind.FilterOpts) (*SequencerInboxBlobMockFeeTokenPricerSetIterator, error) {

	logs, sub, err := _SequencerInboxBlobMock.contract.FilterLogs(opts, "FeeTokenPricerSet")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockFeeTokenPricerSetIterator{contract: _SequencerInboxBlobMock.contract, event: "FeeTokenPricerSet", logs: logs, sub: sub}, nil
}

// WatchFeeTokenPricerSet is a free log subscription operation binding the contract event 0xe83d6153add50e41b8ee6c1115c4178687349bb12bc3902a50b1f6ad78a0c541.
//
// Solidity: event FeeTokenPricerSet(address feeTokenPricer)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) WatchFeeTokenPricerSet(opts *bind.WatchOpts, sink chan<- *SequencerInboxBlobMockFeeTokenPricerSet) (event.Subscription, error) {

	logs, sub, err := _SequencerInboxBlobMock.contract.WatchLogs(opts, "FeeTokenPricerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxBlobMockFeeTokenPricerSet)
				if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "FeeTokenPricerSet", log); err != nil {
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

// ParseFeeTokenPricerSet is a log parse operation binding the contract event 0xe83d6153add50e41b8ee6c1115c4178687349bb12bc3902a50b1f6ad78a0c541.
//
// Solidity: event FeeTokenPricerSet(address feeTokenPricer)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) ParseFeeTokenPricerSet(log types.Log) (*SequencerInboxBlobMockFeeTokenPricerSet, error) {
	event := new(SequencerInboxBlobMockFeeTokenPricerSet)
	if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "FeeTokenPricerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxBlobMockInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockInboxMessageDeliveredIterator struct {
	Event *SequencerInboxBlobMockInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *SequencerInboxBlobMockInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxBlobMockInboxMessageDelivered)
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
		it.Event = new(SequencerInboxBlobMockInboxMessageDelivered)
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
func (it *SequencerInboxBlobMockInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxBlobMockInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxBlobMockInboxMessageDelivered represents a InboxMessageDelivered event raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*SequencerInboxBlobMockInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxBlobMock.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockInboxMessageDeliveredIterator{contract: _SequencerInboxBlobMock.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *SequencerInboxBlobMockInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxBlobMock.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxBlobMockInboxMessageDelivered)
				if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) ParseInboxMessageDelivered(log types.Log) (*SequencerInboxBlobMockInboxMessageDelivered, error) {
	event := new(SequencerInboxBlobMockInboxMessageDelivered)
	if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxBlobMockInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockInboxMessageDeliveredFromOriginIterator struct {
	Event *SequencerInboxBlobMockInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *SequencerInboxBlobMockInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxBlobMockInboxMessageDeliveredFromOrigin)
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
		it.Event = new(SequencerInboxBlobMockInboxMessageDeliveredFromOrigin)
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
func (it *SequencerInboxBlobMockInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxBlobMockInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxBlobMockInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*SequencerInboxBlobMockInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxBlobMock.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockInboxMessageDeliveredFromOriginIterator{contract: _SequencerInboxBlobMock.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *SequencerInboxBlobMockInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxBlobMock.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxBlobMockInboxMessageDeliveredFromOrigin)
				if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*SequencerInboxBlobMockInboxMessageDeliveredFromOrigin, error) {
	event := new(SequencerInboxBlobMockInboxMessageDeliveredFromOrigin)
	if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxBlobMockInvalidateKeysetIterator is returned from FilterInvalidateKeyset and is used to iterate over the raw logs and unpacked data for InvalidateKeyset events raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockInvalidateKeysetIterator struct {
	Event *SequencerInboxBlobMockInvalidateKeyset // Event containing the contract specifics and raw log

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
func (it *SequencerInboxBlobMockInvalidateKeysetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxBlobMockInvalidateKeyset)
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
		it.Event = new(SequencerInboxBlobMockInvalidateKeyset)
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
func (it *SequencerInboxBlobMockInvalidateKeysetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxBlobMockInvalidateKeysetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxBlobMockInvalidateKeyset represents a InvalidateKeyset event raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockInvalidateKeyset struct {
	KeysetHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvalidateKeyset is a free log retrieval operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) FilterInvalidateKeyset(opts *bind.FilterOpts, keysetHash [][32]byte) (*SequencerInboxBlobMockInvalidateKeysetIterator, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxBlobMock.contract.FilterLogs(opts, "InvalidateKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockInvalidateKeysetIterator{contract: _SequencerInboxBlobMock.contract, event: "InvalidateKeyset", logs: logs, sub: sub}, nil
}

// WatchInvalidateKeyset is a free log subscription operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) WatchInvalidateKeyset(opts *bind.WatchOpts, sink chan<- *SequencerInboxBlobMockInvalidateKeyset, keysetHash [][32]byte) (event.Subscription, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxBlobMock.contract.WatchLogs(opts, "InvalidateKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxBlobMockInvalidateKeyset)
				if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "InvalidateKeyset", log); err != nil {
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

// ParseInvalidateKeyset is a log parse operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) ParseInvalidateKeyset(log types.Log) (*SequencerInboxBlobMockInvalidateKeyset, error) {
	event := new(SequencerInboxBlobMockInvalidateKeyset)
	if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "InvalidateKeyset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxBlobMockMaxTimeVariationSetIterator is returned from FilterMaxTimeVariationSet and is used to iterate over the raw logs and unpacked data for MaxTimeVariationSet events raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockMaxTimeVariationSetIterator struct {
	Event *SequencerInboxBlobMockMaxTimeVariationSet // Event containing the contract specifics and raw log

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
func (it *SequencerInboxBlobMockMaxTimeVariationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxBlobMockMaxTimeVariationSet)
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
		it.Event = new(SequencerInboxBlobMockMaxTimeVariationSet)
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
func (it *SequencerInboxBlobMockMaxTimeVariationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxBlobMockMaxTimeVariationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxBlobMockMaxTimeVariationSet represents a MaxTimeVariationSet event raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockMaxTimeVariationSet struct {
	MaxTimeVariation ISequencerInboxMaxTimeVariation
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMaxTimeVariationSet is a free log retrieval operation binding the contract event 0xaa6a58dad31128ff7ecc2b80987ee6e003df80bc50cd8d0b0d1af0e07da6d19d.
//
// Solidity: event MaxTimeVariationSet((uint256,uint256,uint256,uint256) maxTimeVariation)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) FilterMaxTimeVariationSet(opts *bind.FilterOpts) (*SequencerInboxBlobMockMaxTimeVariationSetIterator, error) {

	logs, sub, err := _SequencerInboxBlobMock.contract.FilterLogs(opts, "MaxTimeVariationSet")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockMaxTimeVariationSetIterator{contract: _SequencerInboxBlobMock.contract, event: "MaxTimeVariationSet", logs: logs, sub: sub}, nil
}

// WatchMaxTimeVariationSet is a free log subscription operation binding the contract event 0xaa6a58dad31128ff7ecc2b80987ee6e003df80bc50cd8d0b0d1af0e07da6d19d.
//
// Solidity: event MaxTimeVariationSet((uint256,uint256,uint256,uint256) maxTimeVariation)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) WatchMaxTimeVariationSet(opts *bind.WatchOpts, sink chan<- *SequencerInboxBlobMockMaxTimeVariationSet) (event.Subscription, error) {

	logs, sub, err := _SequencerInboxBlobMock.contract.WatchLogs(opts, "MaxTimeVariationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxBlobMockMaxTimeVariationSet)
				if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "MaxTimeVariationSet", log); err != nil {
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

// ParseMaxTimeVariationSet is a log parse operation binding the contract event 0xaa6a58dad31128ff7ecc2b80987ee6e003df80bc50cd8d0b0d1af0e07da6d19d.
//
// Solidity: event MaxTimeVariationSet((uint256,uint256,uint256,uint256) maxTimeVariation)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) ParseMaxTimeVariationSet(log types.Log) (*SequencerInboxBlobMockMaxTimeVariationSet, error) {
	event := new(SequencerInboxBlobMockMaxTimeVariationSet)
	if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "MaxTimeVariationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxBlobMockOwnerFunctionCalledIterator is returned from FilterOwnerFunctionCalled and is used to iterate over the raw logs and unpacked data for OwnerFunctionCalled events raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockOwnerFunctionCalledIterator struct {
	Event *SequencerInboxBlobMockOwnerFunctionCalled // Event containing the contract specifics and raw log

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
func (it *SequencerInboxBlobMockOwnerFunctionCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxBlobMockOwnerFunctionCalled)
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
		it.Event = new(SequencerInboxBlobMockOwnerFunctionCalled)
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
func (it *SequencerInboxBlobMockOwnerFunctionCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxBlobMockOwnerFunctionCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxBlobMockOwnerFunctionCalled represents a OwnerFunctionCalled event raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockOwnerFunctionCalled struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOwnerFunctionCalled is a free log retrieval operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) FilterOwnerFunctionCalled(opts *bind.FilterOpts, id []*big.Int) (*SequencerInboxBlobMockOwnerFunctionCalledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SequencerInboxBlobMock.contract.FilterLogs(opts, "OwnerFunctionCalled", idRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockOwnerFunctionCalledIterator{contract: _SequencerInboxBlobMock.contract, event: "OwnerFunctionCalled", logs: logs, sub: sub}, nil
}

// WatchOwnerFunctionCalled is a free log subscription operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) WatchOwnerFunctionCalled(opts *bind.WatchOpts, sink chan<- *SequencerInboxBlobMockOwnerFunctionCalled, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SequencerInboxBlobMock.contract.WatchLogs(opts, "OwnerFunctionCalled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxBlobMockOwnerFunctionCalled)
				if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
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

// ParseOwnerFunctionCalled is a log parse operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) ParseOwnerFunctionCalled(log types.Log) (*SequencerInboxBlobMockOwnerFunctionCalled, error) {
	event := new(SequencerInboxBlobMockOwnerFunctionCalled)
	if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxBlobMockSequencerBatchDataIterator is returned from FilterSequencerBatchData and is used to iterate over the raw logs and unpacked data for SequencerBatchData events raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockSequencerBatchDataIterator struct {
	Event *SequencerInboxBlobMockSequencerBatchData // Event containing the contract specifics and raw log

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
func (it *SequencerInboxBlobMockSequencerBatchDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxBlobMockSequencerBatchData)
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
		it.Event = new(SequencerInboxBlobMockSequencerBatchData)
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
func (it *SequencerInboxBlobMockSequencerBatchDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxBlobMockSequencerBatchDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxBlobMockSequencerBatchData represents a SequencerBatchData event raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockSequencerBatchData struct {
	BatchSequenceNumber *big.Int
	Data                []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchData is a free log retrieval operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) FilterSequencerBatchData(opts *bind.FilterOpts, batchSequenceNumber []*big.Int) (*SequencerInboxBlobMockSequencerBatchDataIterator, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}

	logs, sub, err := _SequencerInboxBlobMock.contract.FilterLogs(opts, "SequencerBatchData", batchSequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockSequencerBatchDataIterator{contract: _SequencerInboxBlobMock.contract, event: "SequencerBatchData", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchData is a free log subscription operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) WatchSequencerBatchData(opts *bind.WatchOpts, sink chan<- *SequencerInboxBlobMockSequencerBatchData, batchSequenceNumber []*big.Int) (event.Subscription, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}

	logs, sub, err := _SequencerInboxBlobMock.contract.WatchLogs(opts, "SequencerBatchData", batchSequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxBlobMockSequencerBatchData)
				if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "SequencerBatchData", log); err != nil {
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

// ParseSequencerBatchData is a log parse operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) ParseSequencerBatchData(log types.Log) (*SequencerInboxBlobMockSequencerBatchData, error) {
	event := new(SequencerInboxBlobMockSequencerBatchData)
	if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "SequencerBatchData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxBlobMockSequencerBatchDeliveredIterator is returned from FilterSequencerBatchDelivered and is used to iterate over the raw logs and unpacked data for SequencerBatchDelivered events raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockSequencerBatchDeliveredIterator struct {
	Event *SequencerInboxBlobMockSequencerBatchDelivered // Event containing the contract specifics and raw log

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
func (it *SequencerInboxBlobMockSequencerBatchDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxBlobMockSequencerBatchDelivered)
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
		it.Event = new(SequencerInboxBlobMockSequencerBatchDelivered)
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
func (it *SequencerInboxBlobMockSequencerBatchDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxBlobMockSequencerBatchDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxBlobMockSequencerBatchDelivered represents a SequencerBatchDelivered event raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockSequencerBatchDelivered struct {
	BatchSequenceNumber      *big.Int
	BeforeAcc                [32]byte
	AfterAcc                 [32]byte
	DelayedAcc               [32]byte
	AfterDelayedMessagesRead *big.Int
	TimeBounds               IBridgeTimeBounds
	DataLocation             uint8
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchDelivered is a free log retrieval operation binding the contract event 0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed batchSequenceNumber, bytes32 indexed beforeAcc, bytes32 indexed afterAcc, bytes32 delayedAcc, uint256 afterDelayedMessagesRead, (uint64,uint64,uint64,uint64) timeBounds, uint8 dataLocation)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) FilterSequencerBatchDelivered(opts *bind.FilterOpts, batchSequenceNumber []*big.Int, beforeAcc [][32]byte, afterAcc [][32]byte) (*SequencerInboxBlobMockSequencerBatchDeliveredIterator, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}
	var afterAccRule []interface{}
	for _, afterAccItem := range afterAcc {
		afterAccRule = append(afterAccRule, afterAccItem)
	}

	logs, sub, err := _SequencerInboxBlobMock.contract.FilterLogs(opts, "SequencerBatchDelivered", batchSequenceNumberRule, beforeAccRule, afterAccRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockSequencerBatchDeliveredIterator{contract: _SequencerInboxBlobMock.contract, event: "SequencerBatchDelivered", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchDelivered is a free log subscription operation binding the contract event 0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed batchSequenceNumber, bytes32 indexed beforeAcc, bytes32 indexed afterAcc, bytes32 delayedAcc, uint256 afterDelayedMessagesRead, (uint64,uint64,uint64,uint64) timeBounds, uint8 dataLocation)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) WatchSequencerBatchDelivered(opts *bind.WatchOpts, sink chan<- *SequencerInboxBlobMockSequencerBatchDelivered, batchSequenceNumber []*big.Int, beforeAcc [][32]byte, afterAcc [][32]byte) (event.Subscription, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}
	var afterAccRule []interface{}
	for _, afterAccItem := range afterAcc {
		afterAccRule = append(afterAccRule, afterAccItem)
	}

	logs, sub, err := _SequencerInboxBlobMock.contract.WatchLogs(opts, "SequencerBatchDelivered", batchSequenceNumberRule, beforeAccRule, afterAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxBlobMockSequencerBatchDelivered)
				if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
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

// ParseSequencerBatchDelivered is a log parse operation binding the contract event 0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed batchSequenceNumber, bytes32 indexed beforeAcc, bytes32 indexed afterAcc, bytes32 delayedAcc, uint256 afterDelayedMessagesRead, (uint64,uint64,uint64,uint64) timeBounds, uint8 dataLocation)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) ParseSequencerBatchDelivered(log types.Log) (*SequencerInboxBlobMockSequencerBatchDelivered, error) {
	event := new(SequencerInboxBlobMockSequencerBatchDelivered)
	if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxBlobMockSequencerSetIterator is returned from FilterSequencerSet and is used to iterate over the raw logs and unpacked data for SequencerSet events raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockSequencerSetIterator struct {
	Event *SequencerInboxBlobMockSequencerSet // Event containing the contract specifics and raw log

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
func (it *SequencerInboxBlobMockSequencerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxBlobMockSequencerSet)
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
		it.Event = new(SequencerInboxBlobMockSequencerSet)
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
func (it *SequencerInboxBlobMockSequencerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxBlobMockSequencerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxBlobMockSequencerSet represents a SequencerSet event raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockSequencerSet struct {
	Addr        common.Address
	IsSequencer bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSequencerSet is a free log retrieval operation binding the contract event 0xeb12a9a53eec138c91b27b4f912a257bd690c18fc8bde744be92a0365eb9b87e.
//
// Solidity: event SequencerSet(address addr, bool isSequencer)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) FilterSequencerSet(opts *bind.FilterOpts) (*SequencerInboxBlobMockSequencerSetIterator, error) {

	logs, sub, err := _SequencerInboxBlobMock.contract.FilterLogs(opts, "SequencerSet")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockSequencerSetIterator{contract: _SequencerInboxBlobMock.contract, event: "SequencerSet", logs: logs, sub: sub}, nil
}

// WatchSequencerSet is a free log subscription operation binding the contract event 0xeb12a9a53eec138c91b27b4f912a257bd690c18fc8bde744be92a0365eb9b87e.
//
// Solidity: event SequencerSet(address addr, bool isSequencer)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) WatchSequencerSet(opts *bind.WatchOpts, sink chan<- *SequencerInboxBlobMockSequencerSet) (event.Subscription, error) {

	logs, sub, err := _SequencerInboxBlobMock.contract.WatchLogs(opts, "SequencerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxBlobMockSequencerSet)
				if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "SequencerSet", log); err != nil {
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

// ParseSequencerSet is a log parse operation binding the contract event 0xeb12a9a53eec138c91b27b4f912a257bd690c18fc8bde744be92a0365eb9b87e.
//
// Solidity: event SequencerSet(address addr, bool isSequencer)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) ParseSequencerSet(log types.Log) (*SequencerInboxBlobMockSequencerSet, error) {
	event := new(SequencerInboxBlobMockSequencerSet)
	if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "SequencerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxBlobMockSetValidKeysetIterator is returned from FilterSetValidKeyset and is used to iterate over the raw logs and unpacked data for SetValidKeyset events raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockSetValidKeysetIterator struct {
	Event *SequencerInboxBlobMockSetValidKeyset // Event containing the contract specifics and raw log

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
func (it *SequencerInboxBlobMockSetValidKeysetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxBlobMockSetValidKeyset)
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
		it.Event = new(SequencerInboxBlobMockSetValidKeyset)
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
func (it *SequencerInboxBlobMockSetValidKeysetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxBlobMockSetValidKeysetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxBlobMockSetValidKeyset represents a SetValidKeyset event raised by the SequencerInboxBlobMock contract.
type SequencerInboxBlobMockSetValidKeyset struct {
	KeysetHash  [32]byte
	KeysetBytes []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetValidKeyset is a free log retrieval operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) FilterSetValidKeyset(opts *bind.FilterOpts, keysetHash [][32]byte) (*SequencerInboxBlobMockSetValidKeysetIterator, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxBlobMock.contract.FilterLogs(opts, "SetValidKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxBlobMockSetValidKeysetIterator{contract: _SequencerInboxBlobMock.contract, event: "SetValidKeyset", logs: logs, sub: sub}, nil
}

// WatchSetValidKeyset is a free log subscription operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) WatchSetValidKeyset(opts *bind.WatchOpts, sink chan<- *SequencerInboxBlobMockSetValidKeyset, keysetHash [][32]byte) (event.Subscription, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxBlobMock.contract.WatchLogs(opts, "SetValidKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxBlobMockSetValidKeyset)
				if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "SetValidKeyset", log); err != nil {
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

// ParseSetValidKeyset is a log parse operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_SequencerInboxBlobMock *SequencerInboxBlobMockFilterer) ParseSetValidKeyset(log types.Log) (*SequencerInboxBlobMockSetValidKeyset, error) {
	event := new(SequencerInboxBlobMockSetValidKeyset)
	if err := _SequencerInboxBlobMock.contract.UnpackLog(event, "SetValidKeyset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubMetaData contains all meta data concerning the SequencerInboxStub contract.
var SequencerInboxStubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation_\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"maxDataSize_\",\"type\":\"uint256\"},{\"internalType\":\"contractIReader4844\",\"name\":\"reader4844_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isUsingFeeToken_\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDelayBufferable_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"AlreadyValidDASKeyset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadBufferConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadMaxTimeVariation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"BadSequencerNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSetFeeTokenPricer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataBlobsNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDataLength\",\"type\":\"uint256\"}],\"name\":\"DataTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayProofRequired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayedBackwards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayedTooFar\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Deprecated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtraGasNotUint64\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceIncludeBlockTooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HadZeroInit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectMessagePreimage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"InitParamZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDelayedAccPreimage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"name\":\"InvalidHeaderFlag\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeysetTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingDataHashes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"NoSuchKeyset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBatchPoster\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NotBatchPosterManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCodelessOrigin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotDelayBufferable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotForked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupNotChanged\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBatchPosterManager\",\"type\":\"address\"}],\"name\":\"BatchPosterManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"batchPoster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBatchPoster\",\"type\":\"bool\"}],\"name\":\"BatchPosterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"replenishRateInBasis\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBufferConfig\",\"name\":\"bufferConfig\",\"type\":\"tuple\"}],\"name\":\"BufferConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeTokenPricer\",\"type\":\"address\"}],\"name\":\"FeeTokenPricerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keysetHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidateKeyset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation\",\"type\":\"tuple\"}],\"name\":\"MaxTimeVariationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"OwnerFunctionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchSequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SequencerBatchData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchSequenceNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxBlockNumber\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIBridge.TimeBounds\",\"name\":\"timeBounds\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"enumIBridge.BatchDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"}],\"name\":\"SequencerBatchDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSequencer\",\"type\":\"bool\"}],\"name\":\"SequencerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keysetHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"keysetBytes\",\"type\":\"bytes\"}],\"name\":\"SetValidKeyset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BROTLI_MESSAGE_HEADER_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CUSTOM_DA_MESSAGE_HEADER_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAS_MESSAGE_HEADER_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DATA_AUTHENTICATED_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DATA_BLOB_HEADER_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HEADER_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TREE_DAS_MESSAGE_HEADER_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_HEAVY_MESSAGE_HEADER_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"addInitMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"addSequencerL2Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"beforeDelayedAcc\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMessages.Message\",\"name\":\"delayedMessage\",\"type\":\"tuple\"}],\"internalType\":\"structDelayProof\",\"name\":\"delayProof\",\"type\":\"tuple\"}],\"name\":\"addSequencerL2BatchDelayProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"addSequencerL2BatchFromBlobs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"beforeDelayedAcc\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMessages.Message\",\"name\":\"delayedMessage\",\"type\":\"tuple\"}],\"internalType\":\"structDelayProof\",\"name\":\"delayProof\",\"type\":\"tuple\"}],\"name\":\"addSequencerL2BatchFromBlobsDelayProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addSequencerL2BatchFromOrigin\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"addSequencerL2BatchFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"beforeDelayedAcc\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMessages.Message\",\"name\":\"delayedMessage\",\"type\":\"tuple\"}],\"internalType\":\"structDelayProof\",\"name\":\"delayProof\",\"type\":\"tuple\"}],\"name\":\"addSequencerL2BatchFromOriginDelayProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchPosterManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buffer\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"bufferBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"prevBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"replenishRateInBasis\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"prevSequencedBlockNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dasKeySetInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidKeyset\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"creationBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTokenPricer\",\"outputs\":[{\"internalType\":\"contractIFeeTokenPricer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint64[2]\",\"name\":\"l1BlockAndTime\",\"type\":\"uint64[2]\"},{\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"forceInclusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"}],\"name\":\"forceInclusionDeadline\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"getKeysetCreationBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"inboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"replenishRateInBasis\",\"type\":\"uint64\"}],\"internalType\":\"structBufferConfig\",\"name\":\"bufferConfig_\",\"type\":\"tuple\"},{\"internalType\":\"contractIFeeTokenPricer\",\"name\":\"feeTokenPricer_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"invalidateKeysetHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBatchPoster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDelayBufferable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUsingFeeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"isValidKeysetHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDataSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTimeVariation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"replenishRateInBasis\",\"type\":\"uint64\"}],\"internalType\":\"structBufferConfig\",\"name\":\"bufferConfig_\",\"type\":\"tuple\"}],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reader4844\",\"outputs\":[{\"internalType\":\"contractIReader4844\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeDelayAfterFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBatchPosterManager\",\"type\":\"address\"}],\"name\":\"setBatchPosterManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"replenishRateInBasis\",\"type\":\"uint64\"}],\"internalType\":\"structBufferConfig\",\"name\":\"bufferConfig_\",\"type\":\"tuple\"}],\"name\":\"setBufferConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFeeTokenPricer\",\"name\":\"feeTokenPricer_\",\"type\":\"address\"}],\"name\":\"setFeeTokenPricer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isBatchPoster_\",\"type\":\"bool\"}],\"name\":\"setIsBatchPoster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSequencer_\",\"type\":\"bool\"}],\"name\":\"setIsSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation_\",\"type\":\"tuple\"}],\"name\":\"setMaxTimeVariation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"keysetBytes\",\"type\":\"bytes\"}],\"name\":\"setValidKeyset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDelayedMessagesRead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRollupAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x610180604052306080526202000060a05246610100526200002b620001c8602090811b6200338d17901c565b1515610120523480156200003e57600080fd5b5060405162005b7b38038062005b7b8339810160408190526200006191620002a1565b838383838360e081815250506101205115620000a6576001600160a01b03831615620000a0576040516386657a5360e01b815260040160405180910390fd5b620000ef565b6001600160a01b038316620000ef576040516380fc2c0360e01b815260206004820152600a60248201526914995859195c8d0e0d0d60b21b604482015260640160405180910390fd5b6001600160a01b0392831660c05290151561014052151561016052600180549982166001600160a01b03199a8b1617815560028054909a1633179099558651600a80546020808b01516040808d01516060909d01516001600160401b03908116600160c01b026001600160c01b039e8216600160801b029e909e166001600160801b0393821668010000000000000000026001600160801b0319909616919097161793909317169390931799909917905597166000908152600390975250505091909220805460ff191690931790925550620003ca9050565b60408051600481526024810182526020810180516001600160e01b03166302881c7960e11b17905290516000918291829160649162000208919062000399565b600060405180830381855afa9150503d806000811462000245576040519150601f19603f3d011682016040523d82523d6000602084013e6200024a565b606091505b50915091508180156200025e575080516020145b9250505090565b6001600160a01b03811681146200027b57600080fd5b50565b80516200028b8162000265565b919050565b805180151581146200028b57600080fd5b6000806000806000806000878903610140811215620002bf57600080fd5b8851620002cc8162000265565b60208a0151909850620002df8162000265565b96506080603f1982011215620002f457600080fd5b50604051608081016001600160401b03811182821017156200032657634e487b7160e01b600052604160045260246000fd5b806040525060408901518152606089015160208201526080890151604082015260a089015160608201528095505060c088015193506200036960e089016200027e565b92506200037a610100890162000290565b91506200038b610120890162000290565b905092959891949750929550565b6000825160005b81811015620003bc5760208186018101518583015201620003a0565b506000920191825250919050565b60805160a05160c05160e0516101005161012051610140516101605161568e620004ed6000396000818161048e0152818161153701528181611a2b0152818161210701528181612554015281816128c001528181612ed80152818161306d015281816134550152613697015260008181610646015281816109d30152818161277f015281816128ef0152818161428401526143ab015260008181612cb9015281816137a301526142c40152600081816123c90152613e1201526000818161079a0152818161472d01526147820152600081816105f901528181610fb8015281816120b00152818161406f015261414a01526000818161119d015281816116f801528181611fa501526122bf01526000818161244801526125ec015261568e6000f3fe608060405234801561001057600080fd5b50600436106103205760003560e01c80637fa3a40e116101a7578063d1ce8da8116100ee578063e78cea9211610097578063edaafe2011610071578063edaafe20146107e4578063f19815781461086d578063f60a50911461088057600080fd5b8063e78cea9214610782578063e8eb1dc314610795578063ebea461d146107bc57600080fd5b8063dd44e6e0116100c8578063dd44e6e01461071c578063e0bc972914610748578063e5a358c81461075b57600080fd5b8063d1ce8da8146106cf578063d9dd67ab146106e2578063dab341a4146106f557600080fd5b806396cc5c7811610150578063b31761f81161012a578063b31761f814610696578063cb23bcb5146106a9578063cc2a1a0c146106bc57600080fd5b806396cc5c7814610668578063a655d93714610670578063a84840b71461068357600080fd5b80638f111f3c116101815780638f111f3c1461061b578063917cf8ac1461062e57806392d9f7821461064157600080fd5b80637fa3a40e146105d857806384420860146105e15780638d910dde146105f457600080fd5b80633e5aa0821161026b5780636d46e987116102145780636f12b0c9116101ee5780636f12b0c91461054e578063715ea34b1461056157806371c3e6fe146105b557600080fd5b80636d46e987146105055780636e620055146105285780636e7df3e71461053b57600080fd5b806369cacded1161024557806369cacded146104c35780636ae71f12146104d65780636c890450146104de57600080fd5b80633e5aa082146104765780634b678a66146104895780636633ae85146104b057600080fd5b80631ff64790116102cd57806327957a49116102a757806327957a49146104345780632cbf74e51461043c5780632f3985a71461046357600080fd5b80631ff64790146103e357806322291e8d146103f6578063258f04951461042157600080fd5b80631637be48116102fe5780631637be481461039557806316af91a7146103c85780631f956632146103d057600080fd5b806302c9927514610325578063036f7ed31461036a57806306f130561461037f575b600080fd5b61034c7f200000000000000000000000000000000000000000000000000000000000000081565b6040516001600160f81b031990911681526020015b60405180910390f35b61037d610378366004614ad7565b61088b565b005b610387610ab5565b604051908152602001610361565b6103b86103a3366004614af4565b60009081526008602052604090205460ff1690565b6040519015158152602001610361565b61034c600081565b61037d6103de366004614b1b565b610b3f565b61037d6103f1366004614ad7565b610caa565b600e54610409906001600160a01b031681565b6040516001600160a01b039091168152602001610361565b61038761042f366004614af4565b610e0f565b610387602881565b61034c7f500000000000000000000000000000000000000000000000000000000000000081565b61037d610471366004614c73565b610e7c565b61037d610484366004614c8f565b610fb5565b6103b87f000000000000000000000000000000000000000000000000000000000000000081565b61037d6104be366004614af4565b6112a3565b61037d6104d1366004614d3a565b6114c0565b61037d611800565b61034c7f080000000000000000000000000000000000000000000000000000000000000081565b6103b8610513366004614ad7565b60096020526000908152604090205460ff1681565b61037d610536366004614d3a565b6119d8565b61037d610549366004614b1b565b611a8a565b61037d61055c366004614dc8565b611bf5565b61059561056f366004614af4565b60086020526000908152604090205460ff811690610100900467ffffffffffffffff1682565b60408051921515835267ffffffffffffffff909116602083015201610361565b6103b86105c3366004614ad7565b60036020526000908152604090205460ff1681565b61038760005481565b61037d6105ef366004614af4565b611c27565b6104097f000000000000000000000000000000000000000000000000000000000000000081565b61037d610629366004614e33565b611d9c565b61037d61063c366004614eb0565b6120ad565b6103b87f000000000000000000000000000000000000000000000000000000000000000081565b61037d6123c6565b61037d61067e366004614c73565b61243e565b61037d610691366004614f0c565b6125e2565b61037d6106a4366004614f73565b61298d565b600254610409906001600160a01b031681565b600b54610409906001600160a01b031681565b61037d6106dd366004614fd9565b612aec565b6103876106f0366004614af4565b612e39565b61034c7f010000000000000000000000000000000000000000000000000000000000000081565b61072f61072a36600461501b565b612ec6565b60405167ffffffffffffffff9091168152602001610361565b61037d610756366004614e33565b612f29565b61034c7f400000000000000000000000000000000000000000000000000000000000000081565b600154610409906001600160a01b031681565b6103877f000000000000000000000000000000000000000000000000000000000000000081565b6107c4612fb1565b604080519485526020850193909352918301526060820152608001610361565b600c54600d5461082a9167ffffffffffffffff8082169268010000000000000000808404831693600160801b8104841693600160c01b9091048116928082169290041686565b6040805167ffffffffffffffff978816815295871660208701529386169385019390935290841660608401528316608083015290911660a082015260c001610361565b61037d61087b366004615047565b612fea565b61034c600160ff1b81565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108de573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061090291906150af565b6001600160a01b0316336001600160a01b0316146109d15760025460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610963573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061098791906150af565b6040517f23295f0e0000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152911660248201526044015b60405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000610a28576040517fe13123b400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527fe83d6153add50e41b8ee6c1115c4178687349bb12bc3902a50b1f6ad78a0c5419060200160405180910390a16040516006907fea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e90600090a250565b600154604080517e84120c00000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916284120c9160048083019260209291908290030181865afa158015610b16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b3a91906150cc565b905090565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bb691906150af565b6001600160a01b0316336001600160a01b031614158015610be25750600b546001600160a01b03163314155b15610c1b576040517f660b3b420000000000000000000000000000000000000000000000000000000081523360048201526024016109c8565b6001600160a01b038216600081815260096020908152604091829020805460ff19168515159081179091558251938452908301527feb12a9a53eec138c91b27b4f912a257bd690c18fc8bde744be92a0365eb9b87e910160405180910390a16040516004907fea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e90600090a25050565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cfd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2191906150af565b6001600160a01b0316336001600160a01b031614610d825760025460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610963573d6000803e3d6000fd5b600b805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f3cd6c184800297a0f2b00926a683cbe76890bb7fd01480ac0a10ed6c8f7f66599060200160405180910390a16040516005907fea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e90600090a250565b600081815260086020908152604080832081518083019092525460ff811615158252610100900467ffffffffffffffff16918101829052908203610e685760405162f20c5d60e01b8152600481018490526024016109c8565b6020015167ffffffffffffffff1692915050565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ecf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef391906150af565b6001600160a01b0316336001600160a01b031614610f545760025460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610963573d6000803e3d6000fd5b610f5d81613453565b60408051825167ffffffffffffffff908116825260208085015182169083015283830151168183015290517faa7a2d8175dee3b637814ad6346005dfcc357165396fb8327f649effe8abcf859181900360600190a150565b827f000000000000000000000000000000000000000000000000000000000000000060005a3360009081526003602052604090205490915060ff1661100d57604051632dd9fc9760e01b815260040160405180910390fd5b61101687613693565b1561103457604051630e5da8fb60e01b815260040160405180910390fd5b611040888887876136db565b6001600160a01b0383161561129957366000602061105f83601f6150fb565b611069919061510e565b9050610200611079600283615214565b611083919061510e565b61108e826006615223565b61109891906150fb565b6110a290846150fb565b92506110ac613812565b6110b957600091506111ec565b6001600160a01b038416156111ec57836001600160a01b031663e83a2d826040518163ffffffff1660e01b8152600401600060405180830381865afa92505050801561112757506040513d6000823e601f3d908101601f19168201604052611124919081019061523a565b60015b156111ec578051156111ea576000856001600160a01b0316631f6d6ef76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611173573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061119791906150cc565b905048817f000000000000000000000000000000000000000000000000000000000000000084516111c89190615223565b6111d29190615223565b6111dc919061510e565b6111e690866150fb565b9450505b505b846001600160a01b031663e3db8a49335a61120790876152e0565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b0390921660048301526024820152604481018590526064016020604051808303816000875af1158015611271573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061129591906152f3565b5050505b5050505050505050565b6000816040516020016112b891815260200190565b60408051808303601f1901815290829052600154815160208301207f8db5993b000000000000000000000000000000000000000000000000000000008452600b6004850152600060248501819052604485019190915291935090916001600160a01b0390911690638db5993b906064016020604051808303816000875af1158015611347573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061136b91906150cc565b905080156113bb5760405162461bcd60e51b815260206004820152601460248201527f414c52454144595f44454c415945445f494e495400000000000000000000000060448201526064016109c8565b807fff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b836040516113eb9190615334565b60405180910390a26000806114006001613825565b9150915060008060008061141a866001600080600161386a565b9350935093509350836000146114725760405162461bcd60e51b815260206004820152601060248201527f414c52454144595f5345515f494e49540000000000000000000000000000000060448201526064016109c8565b8083857f7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7856000548a60026040516114ad9493929190615367565b60405180910390a4505050505050505050565b836000805a90506114cf613812565b611505576040517fc8958ead00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526003602052604090205460ff1661153557604051632dd9fc9760e01b815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000061157357604051631192b39960e31b815260040160405180910390fd5b61158b88611586368790038701876153dc565b613a27565b61159b8b8b8b8b8a8a6001613b34565b6001600160a01b038316156112955736600060206115ba83601f6150fb565b6115c4919061510e565b90506102006115d4600283615214565b6115de919061510e565b6115e9826006615223565b6115f391906150fb565b6115fd90846150fb565b9250611607613812565b6116145760009150611747565b6001600160a01b0384161561174757836001600160a01b031663e83a2d826040518163ffffffff1660e01b8152600401600060405180830381865afa92505050801561168257506040513d6000823e601f3d908101601f1916820160405261167f919081019061523a565b60015b1561174757805115611745576000856001600160a01b0316631f6d6ef76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f291906150cc565b905048817f000000000000000000000000000000000000000000000000000000000000000084516117239190615223565b61172d9190615223565b611737919061510e565b61174190866150fb565b9450505b505b846001600160a01b031663e3db8a49335a61176290876152e0565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b0390921660048301526024820152604481018590526064016020604051808303816000875af11580156117cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117f091906152f3565b5050505050505050505050505050565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611853573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061187791906150af565b6001600160a01b0316336001600160a01b0316146118d85760025460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610963573d6000803e3d6000fd5b600154604080517fcb23bcb500000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163cb23bcb59160048083019260209291908290030181865afa15801561193b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061195f91906150af565b6002549091506001600160a01b038083169116036119a9576040517fd054909f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b836000805a3360009081526003602052604090205490915060ff16158015611a0b57506002546001600160a01b03163314155b15611a2957604051632dd9fc9760e01b815260040160405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000611a6757604051631192b39960e31b815260040160405180910390fd5b611a7a88611586368790038701876153dc565b61159b8b8b8b8b8a8a6000613b34565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611add573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b0191906150af565b6001600160a01b0316336001600160a01b031614158015611b2d5750600b546001600160a01b03163314155b15611b66576040517f660b3b420000000000000000000000000000000000000000000000000000000081523360048201526024016109c8565b6001600160a01b038216600081815260036020908152604091829020805460ff19168515159081179091558251938452908301527f28bcc5626d357efe966b4b0876aa1ee8ab99e26da4f131f6a2623f1800701c21910160405180910390a16040516001907fea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e90600090a25050565b6040517fc73b9d7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c7a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c9e91906150af565b6001600160a01b0316336001600160a01b031614611cff5760025460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610963573d6000803e3d6000fd5b60008181526008602052604090205460ff16611d305760405162f20c5d60e01b8152600481018290526024016109c8565b600081815260086020526040808220805460ff191690555182917f5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a91a26040516003907fea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e90600090a250565b826000805a9050611dab613812565b611de1576040517fc8958ead00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526003602052604090205460ff16611e1157604051632dd9fc9760e01b815260040160405180910390fd5b611e1a87613693565b15611e3857604051630e5da8fb60e01b815260040160405180910390fd5b611e488a8a8a8a89896001613b34565b6001600160a01b038316156120a1573660006020611e6783601f6150fb565b611e71919061510e565b9050610200611e81600283615214565b611e8b919061510e565b611e96826006615223565b611ea091906150fb565b611eaa90846150fb565b9250611eb4613812565b611ec15760009150611ff4565b6001600160a01b03841615611ff457836001600160a01b031663e83a2d826040518163ffffffff1660e01b8152600401600060405180830381865afa925050508015611f2f57506040513d6000823e601f3d908101601f19168201604052611f2c919081019061523a565b60015b15611ff457805115611ff2576000856001600160a01b0316631f6d6ef76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611f7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f9f91906150cc565b905048817f00000000000000000000000000000000000000000000000000000000000000008451611fd09190615223565b611fda9190615223565b611fe4919061510e565b611fee90866150fb565b9450505b505b846001600160a01b031663e3db8a49335a61200f90876152e0565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b0390921660048301526024820152604481018590526064016020604051808303816000875af1158015612079573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061209d91906152f3565b5050505b50505050505050505050565b837f000000000000000000000000000000000000000000000000000000000000000060005a3360009081526003602052604090205490915060ff1661210557604051632dd9fc9760e01b815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000061214357604051631192b39960e31b815260040160405180910390fd5b61215688611586368790038701876153dc565b612162898988886136db565b6001600160a01b038316156123bb57366000602061218183601f6150fb565b61218b919061510e565b905061020061219b600283615214565b6121a5919061510e565b6121b0826006615223565b6121ba91906150fb565b6121c490846150fb565b92506121ce613812565b6121db576000915061230e565b6001600160a01b0384161561230e57836001600160a01b031663e83a2d826040518163ffffffff1660e01b8152600401600060405180830381865afa92505050801561224957506040513d6000823e601f3d908101601f19168201604052612246919081019061523a565b60015b1561230e5780511561230c576000856001600160a01b0316631f6d6ef76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612295573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122b991906150cc565b905048817f000000000000000000000000000000000000000000000000000000000000000084516122ea9190615223565b6122f49190615223565b6122fe919061510e565b61230890866150fb565b9450505b505b846001600160a01b031663e3db8a49335a61232990876152e0565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b0390921660048301526024820152604481018590526064016020604051808303816000875af1158015612393573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123b791906152f3565b5050505b505050505050505050565b467f00000000000000000000000000000000000000000000000000000000000000000361241f576040517fa301bb0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7801000000000000000100000000000000010000000000000001600a55565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036124dc5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c000000000000000000000000000000000000000060648201526084016109c8565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61038054336001600160a01b03821614612552576040517f23295f0e0000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b03821660248201526044016109c8565b7f000000000000000000000000000000000000000000000000000000000000000061259057604051631192b39960e31b815260040160405180910390fd5b600c5467ffffffffffffffff16156125d4576040517fef34ca5c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6125dd83613453565b505050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036126805760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c000000000000000000000000000000000000000060648201526084016109c8565b6001546001600160a01b0316156126c3576040517fef34ca5c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038416612703576040517f1ad0f74300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000846001600160a01b031663e1758bd86040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561275f575060408051601f3d908101601f1916820190925261275c918101906150af565b60015b1561277a576001600160a01b0381161561277857600191505b505b8015157f00000000000000000000000000000000000000000000000000000000000000001515146127d7576040517fc3e31f8d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038716908117909155604080517fcb23bcb5000000000000000000000000000000000000000000000000000000008152905163cb23bcb5916004808201926020929091908290030181865afa158015612856573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061287a91906150af565b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03929092169190911790556128be6128b936869003860186614f73565b613c5f565b7f0000000000000000000000000000000000000000000000000000000000000000156128ed576128ed83613453565b7f000000000000000000000000000000000000000000000000000000000000000015801561292357506001600160a01b03821615155b1561295a576040517fe13123b400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50600e805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055505050565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156129e0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a0491906150af565b6001600160a01b0316336001600160a01b031614612a655760025460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610963573d6000803e3d6000fd5b612a6e81613c5f565b60408051825181526020808401519082015282820151818301526060808401519082015290517faa6a58dad31128ff7ecc2b80987ee6e003df80bc50cd8d0b0d1af0e07da6d19d9181900360800190a16040516000907fea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e908290a250565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612b3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b6391906150af565b6001600160a01b0316336001600160a01b031614612bc45760025460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610963573d6000803e3d6000fd5b60008282604051612bd692919061548a565b6040519081900381207ffe000000000000000000000000000000000000000000000000000000000000006020830152602182015260410160408051601f1981840301815291905280516020909101209050600160ff1b8118620100008310612c6a576040517fb3d1f41200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008181526008602052604090205460ff1615612cb6576040517ffa2fddda000000000000000000000000000000000000000000000000000000008152600481018290526024016109c8565b437f000000000000000000000000000000000000000000000000000000000000000015612d435760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612d1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d4091906150cc565b90505b6040805180820182526001815267ffffffffffffffff8381166020808401918252600087815260089091528490209251835491517fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000009092169015157fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff161761010091909216021790555182907fabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef65572290612dfe908890889061549a565b60405180910390a26040516002907fea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e90600090a25050505050565b6001546040517f16bf5579000000000000000000000000000000000000000000000000000000008152600481018390526000916001600160a01b0316906316bf557990602401602060405180830381865afa158015612e9c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ec091906150cc565b92915050565b600a5460009067ffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000015612f18576000612f09600c85613d8b565b9050612f1481613dda565b9150505b612f2281846154c9565b9392505050565b826000805a3360009081526003602052604090205490915060ff16158015612f5c57506002546001600160a01b03163314155b15612f7a57604051632dd9fc9760e01b815260040160405180910390fd5b612f8387613693565b15612fa157604051630e5da8fb60e01b815260040160405180910390fd5b611e488a8a8a8a89896000613b34565b600080600080600080600080612fc5613e0a565b67ffffffffffffffff9384169b50918316995082169750169450505050505b90919293565b6000548611613025576040517f7d73e6fa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061305b8684613039602089018961501b565b61304960408a0160208b0161501b565b61305460018d6152e0565b8988613e81565b600a5490915067ffffffffffffffff167f0000000000000000000000000000000000000000000000000000000000000000156130cc576130a96130a1602088018861501b565b600c90613f26565b600c546130bf9067ffffffffffffffff16613dda565b67ffffffffffffffff1690505b43816130db602089018961501b565b67ffffffffffffffff166130ef91906150fb565b10613126576040517fad3515d900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060018911156131af576001546001600160a01b031663d5719dc261314d60028c6152e0565b6040518263ffffffff1660e01b815260040161316b91815260200190565b602060405180830381865afa158015613188573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131ac91906150cc565b90505b60408051602080820184905281830186905282518083038401815260609092019092528051910120600180546001600160a01b03169063d5719dc2906131f5908d6152e0565b6040518263ffffffff1660e01b815260040161321391815260200190565b602060405180830381865afa158015613230573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061325491906150cc565b1461328b576040517f13947fd700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806132978b613825565b9150915060008b90506000600160009054906101000a90046001600160a01b03166001600160a01b0316635fca4a166040518163ffffffff1660e01b8152600401602060405180830381865afa1580156132f5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061331991906150cc565b905080600080808061332e898883888061386a565b93509350935093508083857f7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7856000548d60026040516133719493929190615367565b60405180910390a4505050505050505050505050505050505050565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f051038f2000000000000000000000000000000000000000000000000000000001790529051600091829182916064916133f99190615507565b600060405180830381855afa9150503d8060008114613434576040519150601f19603f3d011682016040523d82523d6000602084013e613439565b606091505b509150915081801561344c575080516020145b9250505090565b7f000000000000000000000000000000000000000000000000000000000000000061349157604051631192b39960e31b815260040160405180910390fd5b61349a81613fac565b6134d0576040517fda1c8eb600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c5467ffffffffffffffff1615806134fc57506020810151600c5467ffffffffffffffff9182169116115b15613528576020810151600c805467ffffffffffffffff191667ffffffffffffffff9092169190911790555b8051600c5467ffffffffffffffff91821691161015613565578051600c805467ffffffffffffffff191667ffffffffffffffff9092169190911790555b602081810151600c805484517fffffffffffffffff00000000000000000000000000000000ffffffffffffffff9091166801000000000000000067ffffffffffffffff948516027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1617600160801b91841691909102179055604080840151600d805467ffffffffffffffff1916919093161790915560005460015482517feca067ad000000000000000000000000000000000000000000000000000000008152925191936001600160a01b039091169263eca067ad92600480830193928290030181865afa15801561365c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061368091906150cc565b0361369057613690600c43613f26565b50565b60007f000000000000000000000000000000000000000000000000000000000000000080156136c3575060005482115b8015612ec057506136d4600c614014565b1592915050565b60008060006136e986614047565b925092509250600080600080613703878b60008c8c61386a565b93509350935093508a841415801561371d57506000198b14155b1561375e576040517fac7411c900000000000000000000000000000000000000000000000000000000815260048101859052602481018c90526044016109c8565b80838c7f7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7856000548b60036040516137999493929190615367565b60405180910390a47f0000000000000000000000000000000000000000000000000000000000000000156137f9576040517f86657a5300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613801613812565b156112955761129587854888614276565b60003332148015610b3a575050333b1590565b6040805160808101825260008082526020820181905291810182905260608101829052600080613854856145c8565b8151602090920191909120969095509350505050565b6000806000806000548810156138ac576040517f7d73e6fa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160009054906101000a90046001600160a01b03166001600160a01b031663eca067ad6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156138ff573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061392391906150cc565b88111561395c576040517f925f8bd300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040517f86598a56000000000000000000000000000000000000000000000000000000008152600481018b9052602481018a905260448101889052606481018790526001600160a01b03909116906386598a56906084016080604051808303816000875af11580156139d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139f99190615523565b60008c9055929650909450925090508615613a1b57613a1b8985486000614276565b95509550955095915050565b600054821115613b3057613a3b600c614683565b15613b3057600154600080546040517fd5719dc200000000000000000000000000000000000000000000000000000000815291926001600160a01b03169163d5719dc291613a8f9160040190815260200190565b602060405180830381865afa158015613aac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ad091906150cc565b9050613ae581836000015184602001516146b4565b613b1b576040517fc334542d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020820151604001516125dd90600c90613f26565b5050565b600080613b428888886146f9565b91509150600080600080613b66868b89613b5d576000613b5f565b8d5b8c8c61386a565b93509350935093508c8414158015613b8057506000198d14155b15613bc1576040517fac7411c900000000000000000000000000000000000000000000000000000000815260048101859052602481018e90526044016109c8565b8083857f7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7856000548a8d613bf6576001613bf9565b60005b604051613c099493929190615367565b60405180910390a48661209d57837ffe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c208d8d604051613c4892919061549a565b60405180910390a250505050505050505050505050565b805167ffffffffffffffff1080613c815750602081015167ffffffffffffffff105b80613c975750604081015167ffffffffffffffff105b80613cad5750606081015167ffffffffffffffff105b15613ce4576040517f09cfba7500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600a80546020840151604085015160609095015167ffffffffffffffff908116600160c01b0277ffffffffffffffffffffffffffffffffffffffffffffffff968216600160801b02969096166fffffffffffffffffffffffffffffffff92821668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009094169190951617919091171691909117919091179055565b81546001830154600091612f229167ffffffffffffffff600160c01b8304811692868216928282169268010000000000000000808304821693600160801b810483169391900482169116614907565b600a5460009067ffffffffffffffff90811690831610613e0657600a5467ffffffffffffffff16612ec0565b5090565b6000808080467f000000000000000000000000000000000000000000000000000000000000000014613e4757506001925082915081905080612fe4565b5050600a5467ffffffffffffffff808216935068010000000000000000820481169250600160801b8204811691600160c01b900416612fe4565b6040516001600160f81b031960f889901b1660208201526bffffffffffffffffffffffff19606088901b1660218201527fffffffffffffffff00000000000000000000000000000000000000000000000060c087811b8216603584015286901b16603d82015260458101849052606581018390526085810182905260009060a5016040516020818303038152906040528051906020012090505b979650505050505050565b613f308282613d8b565b825467ffffffffffffffff928316600160c01b0277ffffffffffffffffffffffffffffffff000000000000000090911691831691909117178255600190910180544390921668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff909216919091179055565b805160009067ffffffffffffffff1615801590613fd65750602082015167ffffffffffffffff1615155b8015613ff25750612710826040015167ffffffffffffffff1611155b8015612ec05750506020810151905167ffffffffffffffff9182169116111590565b805460009067ffffffffffffffff600160801b820481169161403f91600160c01b90910416436152e0565b111592915050565b60408051608081018252600080825260208201819052918101829052606081018290526000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e83a2d826040518163ffffffff1660e01b8152600401600060405180830381865afa1580156140cb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526140f3919081019061523a565b90508051600003614130576040517f3cd27eb600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008061413c876145c8565b9150915060008351620200007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316631f6d6ef76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156141a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141ca91906150cc565b6141d49190615223565b6141de9190615223565b60405190915083907f500000000000000000000000000000000000000000000000000000000000000090614216908790602001615559565b60408051601f198184030181529082905261423593929160200161558f565b60405160208183030381529060405280519060200120826000481161425b576000614265565b614265488461510e565b965096509650505050509193909250565b600e546001600160a01b03167f000000000000000000000000000000000000000000000000000000000000000080156142b657506001600160a01b038116155b156142c157506145c2565b327f000000000000000000000000000000000000000000000000000000000000000015614367576000606c6001600160a01b031663c6f7de0e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614329573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061434d91906150cc565b9050614359488261510e565b61436390856150fb565b9350505b67ffffffffffffffff8311156143a9576040517f04d5501200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000080156143de57506001600160a01b03821615155b1561446c576000826001600160a01b031663e6aa216c6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015614425573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061444991906150cc565b9050670de0b6b3a764000061445e8287615223565b614468919061510e565b9450505b604080514260208201526bffffffffffffffffffffffff19606084901b16918101919091526054810187905260748101869052609481018590527fffffffffffffffff00000000000000000000000000000000000000000000000060c085901b1660b482015260009060bc0160408051808303601f1901815290829052600154815160208301207f7a88b1070000000000000000000000000000000000000000000000000000000084526001600160a01b0386811660048601526024850191909152919350600092911690637a88b107906044016020604051808303816000875af115801561455f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061458391906150cc565b9050807fff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b836040516145b59190615334565b60405180910390a2505050505b50505050565b604080516080808201835260008083526020808401829052838501829052606080850183905285519384018652828452838201839052838601839052838101839052855191820183905260288201839052603082018390526038820183905260c087901b7fffffffffffffffff00000000000000000000000000000000000000000000000016958201959095526048016040516020818303038152906040529050602881511461467a5761467a6155d2565b94909350915050565b600061468e82614014565b1580612ec05750505467ffffffffffffffff680100000000000000008204811691161090565b60006146ef836146c3846149ce565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b9093149392505050565b604080516080810182526000808252602082018190529181018290526060810182905260006147298560286150fb565b90507f00000000000000000000000000000000000000000000000000000000000000008111156147ae576040517f4634691b000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000060248201526044016109c8565b6000806147ba866145c8565b909250905086156148cd576147ea888860008181106147db576147db6154f1565b9050013560f81c60f81b6149fb565b6148425787876000818110614801576148016154f1565b6040517f6b3333560000000000000000000000000000000000000000000000000000000081529201356001600160f81b0319166004830152506024016109c8565b600160ff1b8888600081614858576148586154f1565b6001600160f81b031992013592909216161580159150614879575060218710155b156148cd57600061488e602160018a8c6155e8565b61489791615612565b60008181526008602052604090205490915060ff166148cb5760405162f20c5d60e01b8152600481018290526024016109c8565b505b8188886040516020016148e293929190615630565b60408051601f1981840301815291905280516020909101209890975095505050505050565b600080888811614918576000614922565b61492289896152e0565b9050600089871161493457600061493e565b61493e8a886152e0565b905061271061494d8584615223565b614957919061510e565b61496190896150fb565b9750600086821161497357600061497d565b61497d87836152e0565b90508281111561498a5750815b808911156149bf5761499c818a6152e0565b9850868911156149bf578589116149b357886149b5565b855b9350505050613f1b565b50949998505050505050505050565b6000612ec0826000015183602001518460400151856060015186608001518760a001518860c00151613e81565b60006001600160f81b031982161580614a2157506001600160f81b03198216600160ff1b145b80614a5557506001600160f81b031982167f8800000000000000000000000000000000000000000000000000000000000000145b80614a8957506001600160f81b031982167f2000000000000000000000000000000000000000000000000000000000000000145b80612ec057506001600160f81b031982167f01000000000000000000000000000000000000000000000000000000000000001492915050565b6001600160a01b038116811461369057600080fd5b600060208284031215614ae957600080fd5b8135612f2281614ac2565b600060208284031215614b0657600080fd5b5035919050565b801515811461369057600080fd5b60008060408385031215614b2e57600080fd5b8235614b3981614ac2565b91506020830135614b4981614b0d565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715614b8d57614b8d614b54565b60405290565b60405160e0810167ffffffffffffffff81118282101715614b8d57614b8d614b54565b604051601f8201601f1916810167ffffffffffffffff81118282101715614bdf57614bdf614b54565b604052919050565b803567ffffffffffffffff81168114614bff57600080fd5b919050565b600060608284031215614c1657600080fd5b6040516060810181811067ffffffffffffffff82111715614c3957614c39614b54565b604052905080614c4883614be7565b8152614c5660208401614be7565b6020820152614c6760408401614be7565b60408201525092915050565b600060608284031215614c8557600080fd5b612f228383614c04565b600080600080600060a08688031215614ca757600080fd5b85359450602086013593506040860135614cc081614ac2565b94979396509394606081013594506080013592915050565b60008083601f840112614cea57600080fd5b50813567ffffffffffffffff811115614d0257600080fd5b602083019150836020828501011115614d1a57600080fd5b9250929050565b60006101008284031215614d3457600080fd5b50919050565b6000806000806000806000806101c0898b031215614d5757600080fd5b88359750602089013567ffffffffffffffff811115614d7557600080fd5b614d818b828c01614cd8565b909850965050604089013594506060890135614d9c81614ac2565b93506080890135925060a08901359150614db98a60c08b01614d21565b90509295985092959890939650565b600080600080600060808688031215614de057600080fd5b85359450602086013567ffffffffffffffff811115614dfe57600080fd5b614e0a88828901614cd8565b909550935050604086013591506060860135614e2581614ac2565b809150509295509295909350565b600080600080600080600060c0888a031215614e4e57600080fd5b87359650602088013567ffffffffffffffff811115614e6c57600080fd5b614e788a828b01614cd8565b909750955050604088013593506060880135614e9381614ac2565b969995985093969295946080840135945060a09093013592915050565b6000806000806000806101a08789031215614eca57600080fd5b86359550602087013594506040870135614ee381614ac2565b93506060870135925060808701359150614f008860a08901614d21565b90509295509295509295565b600080600080848603610120811215614f2457600080fd5b8535614f2f81614ac2565b94506080601f1982011215614f4357600080fd5b50602085019250614f578660a08701614c04565b9150610100850135614f6881614ac2565b939692955090935050565b600060808284031215614f8557600080fd5b6040516080810181811067ffffffffffffffff82111715614fa857614fa8614b54565b8060405250823581526020830135602082015260408301356040820152606083013560608201528091505092915050565b60008060208385031215614fec57600080fd5b823567ffffffffffffffff81111561500357600080fd5b61500f85828601614cd8565b90969095509350505050565b60006020828403121561502d57600080fd5b612f2282614be7565b803560ff81168114614bff57600080fd5b60008060008060008060e0878903121561506057600080fd5b8635955061507060208801615036565b9450608087018881111561508357600080fd5b60408801945035925060a087013561509a81614ac2565b8092505060c087013590509295509295509295565b6000602082840312156150c157600080fd5b8151612f2281614ac2565b6000602082840312156150de57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b80820180821115612ec057612ec06150e5565b60008261512b57634e487b7160e01b600052601260045260246000fd5b500490565b600181815b8085111561516b578160001904821115615151576151516150e5565b8085161561515e57918102915b93841c9390800290615135565b509250929050565b60008261518257506001612ec0565b8161518f57506000612ec0565b81600181146151a557600281146151af576151cb565b6001915050612ec0565b60ff8411156151c0576151c06150e5565b50506001821b612ec0565b5060208310610133831016604e8410600b84101617156151ee575081810a612ec0565b6151f88383615130565b806000190482111561520c5761520c6150e5565b029392505050565b6000612f2260ff841683615173565b8082028115828204841417612ec057612ec06150e5565b6000602080838503121561524d57600080fd5b825167ffffffffffffffff8082111561526557600080fd5b818501915085601f83011261527957600080fd5b81518181111561528b5761528b614b54565b8060051b915061529c848301614bb6565b81815291830184019184810190888411156152b657600080fd5b938501935b838510156152d4578451825293850193908501906152bb565b98975050505050505050565b81810381811115612ec057612ec06150e5565b60006020828403121561530557600080fd5b8151612f2281614b0d565b60005b8381101561532b578181015183820152602001615313565b50506000910152565b6020815260008251806020840152615353816040850160208701615310565b601f01601f19169190910160400192915050565b600060e08201905085825284602083015267ffffffffffffffff8085511660408401528060208601511660608401528060408601511660808401528060608601511660a084015250600483106153cd57634e487b7160e01b600052602160045260246000fd5b8260c083015295945050505050565b60008183036101008112156153f057600080fd5b6153f8614b6a565b8335815260e0601f198301121561540e57600080fd5b615416614b93565b915061542460208501615036565b8252604084013561543481614ac2565b602083015261544560608501614be7565b604083015261545660808501614be7565b606083015260a0840135608083015260c084013560a083015260e084013560c0830152816020820152809250505092915050565b8183823760009101908152919050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b67ffffffffffffffff8181168382160190808211156154ea576154ea6150e5565b5092915050565b634e487b7160e01b600052603260045260246000fd5b60008251615519818460208701615310565b9190910192915050565b6000806000806080858703121561553957600080fd5b505082516020840151604085015160609095015191969095509092509050565b815160009082906020808601845b8381101561558357815185529382019390820190600101615567565b50929695505050505050565b600084516155a1818460208901615310565b6001600160f81b0319851690830190815283516155c5816001840160208801615310565b0160010195945050505050565b634e487b7160e01b600052600160045260246000fd5b600080858511156155f857600080fd5b8386111561560557600080fd5b5050820193919092039150565b80356020831015612ec057600019602084900360031b1b1692915050565b60008451615642818460208901615310565b820183858237600093019283525090939250505056fea26469706673582212201a78abc6e7b94dc9944739abe351fa2b85d352a1097881e059b203cb320f3af464736f6c63430008110033",
}

// SequencerInboxStubABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerInboxStubMetaData.ABI instead.
var SequencerInboxStubABI = SequencerInboxStubMetaData.ABI

// SequencerInboxStubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SequencerInboxStubMetaData.Bin instead.
var SequencerInboxStubBin = SequencerInboxStubMetaData.Bin

// DeploySequencerInboxStub deploys a new Ethereum contract, binding an instance of SequencerInboxStub to it.
func DeploySequencerInboxStub(auth *bind.TransactOpts, backend bind.ContractBackend, bridge_ common.Address, sequencer_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation, maxDataSize_ *big.Int, reader4844_ common.Address, isUsingFeeToken_ bool, isDelayBufferable_ bool) (common.Address, *types.Transaction, *SequencerInboxStub, error) {
	parsed, err := SequencerInboxStubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SequencerInboxStubBin), backend, bridge_, sequencer_, maxTimeVariation_, maxDataSize_, reader4844_, isUsingFeeToken_, isDelayBufferable_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SequencerInboxStub{SequencerInboxStubCaller: SequencerInboxStubCaller{contract: contract}, SequencerInboxStubTransactor: SequencerInboxStubTransactor{contract: contract}, SequencerInboxStubFilterer: SequencerInboxStubFilterer{contract: contract}}, nil
}

// SequencerInboxStub is an auto generated Go binding around an Ethereum contract.
type SequencerInboxStub struct {
	SequencerInboxStubCaller     // Read-only binding to the contract
	SequencerInboxStubTransactor // Write-only binding to the contract
	SequencerInboxStubFilterer   // Log filterer for contract events
}

// SequencerInboxStubCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerInboxStubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxStubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerInboxStubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxStubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerInboxStubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxStubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerInboxStubSession struct {
	Contract     *SequencerInboxStub // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SequencerInboxStubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerInboxStubCallerSession struct {
	Contract *SequencerInboxStubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SequencerInboxStubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerInboxStubTransactorSession struct {
	Contract     *SequencerInboxStubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SequencerInboxStubRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerInboxStubRaw struct {
	Contract *SequencerInboxStub // Generic contract binding to access the raw methods on
}

// SequencerInboxStubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerInboxStubCallerRaw struct {
	Contract *SequencerInboxStubCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerInboxStubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerInboxStubTransactorRaw struct {
	Contract *SequencerInboxStubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencerInboxStub creates a new instance of SequencerInboxStub, bound to a specific deployed contract.
func NewSequencerInboxStub(address common.Address, backend bind.ContractBackend) (*SequencerInboxStub, error) {
	contract, err := bindSequencerInboxStub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStub{SequencerInboxStubCaller: SequencerInboxStubCaller{contract: contract}, SequencerInboxStubTransactor: SequencerInboxStubTransactor{contract: contract}, SequencerInboxStubFilterer: SequencerInboxStubFilterer{contract: contract}}, nil
}

// NewSequencerInboxStubCaller creates a new read-only instance of SequencerInboxStub, bound to a specific deployed contract.
func NewSequencerInboxStubCaller(address common.Address, caller bind.ContractCaller) (*SequencerInboxStubCaller, error) {
	contract, err := bindSequencerInboxStub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubCaller{contract: contract}, nil
}

// NewSequencerInboxStubTransactor creates a new write-only instance of SequencerInboxStub, bound to a specific deployed contract.
func NewSequencerInboxStubTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerInboxStubTransactor, error) {
	contract, err := bindSequencerInboxStub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubTransactor{contract: contract}, nil
}

// NewSequencerInboxStubFilterer creates a new log filterer instance of SequencerInboxStub, bound to a specific deployed contract.
func NewSequencerInboxStubFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerInboxStubFilterer, error) {
	contract, err := bindSequencerInboxStub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubFilterer{contract: contract}, nil
}

// bindSequencerInboxStub binds a generic wrapper to an already deployed contract.
func bindSequencerInboxStub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SequencerInboxStubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInboxStub *SequencerInboxStubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInboxStub.Contract.SequencerInboxStubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInboxStub *SequencerInboxStubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SequencerInboxStubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInboxStub *SequencerInboxStubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SequencerInboxStubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInboxStub *SequencerInboxStubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInboxStub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInboxStub *SequencerInboxStubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInboxStub *SequencerInboxStubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.contract.Transact(opts, method, params...)
}

// BROTLIMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x16af91a7.
//
// Solidity: function BROTLI_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCaller) BROTLIMESSAGEHEADERFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "BROTLI_MESSAGE_HEADER_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// BROTLIMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x16af91a7.
//
// Solidity: function BROTLI_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubSession) BROTLIMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.BROTLIMESSAGEHEADERFLAG(&_SequencerInboxStub.CallOpts)
}

// BROTLIMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x16af91a7.
//
// Solidity: function BROTLI_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) BROTLIMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.BROTLIMESSAGEHEADERFLAG(&_SequencerInboxStub.CallOpts)
}

// CUSTOMDAMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0xdab341a4.
//
// Solidity: function CUSTOM_DA_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCaller) CUSTOMDAMESSAGEHEADERFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "CUSTOM_DA_MESSAGE_HEADER_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// CUSTOMDAMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0xdab341a4.
//
// Solidity: function CUSTOM_DA_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubSession) CUSTOMDAMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.CUSTOMDAMESSAGEHEADERFLAG(&_SequencerInboxStub.CallOpts)
}

// CUSTOMDAMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0xdab341a4.
//
// Solidity: function CUSTOM_DA_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) CUSTOMDAMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.CUSTOMDAMESSAGEHEADERFLAG(&_SequencerInboxStub.CallOpts)
}

// DASMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0xf60a5091.
//
// Solidity: function DAS_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCaller) DASMESSAGEHEADERFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "DAS_MESSAGE_HEADER_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// DASMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0xf60a5091.
//
// Solidity: function DAS_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubSession) DASMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.DASMESSAGEHEADERFLAG(&_SequencerInboxStub.CallOpts)
}

// DASMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0xf60a5091.
//
// Solidity: function DAS_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) DASMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.DASMESSAGEHEADERFLAG(&_SequencerInboxStub.CallOpts)
}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCaller) DATAAUTHENTICATEDFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "DATA_AUTHENTICATED_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubSession) DATAAUTHENTICATEDFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.DATAAUTHENTICATEDFLAG(&_SequencerInboxStub.CallOpts)
}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) DATAAUTHENTICATEDFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.DATAAUTHENTICATEDFLAG(&_SequencerInboxStub.CallOpts)
}

// DATABLOBHEADERFLAG is a free data retrieval call binding the contract method 0x2cbf74e5.
//
// Solidity: function DATA_BLOB_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCaller) DATABLOBHEADERFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "DATA_BLOB_HEADER_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// DATABLOBHEADERFLAG is a free data retrieval call binding the contract method 0x2cbf74e5.
//
// Solidity: function DATA_BLOB_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubSession) DATABLOBHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.DATABLOBHEADERFLAG(&_SequencerInboxStub.CallOpts)
}

// DATABLOBHEADERFLAG is a free data retrieval call binding the contract method 0x2cbf74e5.
//
// Solidity: function DATA_BLOB_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) DATABLOBHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.DATABLOBHEADERFLAG(&_SequencerInboxStub.CallOpts)
}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCaller) HEADERLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "HEADER_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubSession) HEADERLENGTH() (*big.Int, error) {
	return _SequencerInboxStub.Contract.HEADERLENGTH(&_SequencerInboxStub.CallOpts)
}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) HEADERLENGTH() (*big.Int, error) {
	return _SequencerInboxStub.Contract.HEADERLENGTH(&_SequencerInboxStub.CallOpts)
}

// TREEDASMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x6c890450.
//
// Solidity: function TREE_DAS_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCaller) TREEDASMESSAGEHEADERFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "TREE_DAS_MESSAGE_HEADER_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// TREEDASMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x6c890450.
//
// Solidity: function TREE_DAS_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubSession) TREEDASMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.TREEDASMESSAGEHEADERFLAG(&_SequencerInboxStub.CallOpts)
}

// TREEDASMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x6c890450.
//
// Solidity: function TREE_DAS_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) TREEDASMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.TREEDASMESSAGEHEADERFLAG(&_SequencerInboxStub.CallOpts)
}

// ZEROHEAVYMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x02c99275.
//
// Solidity: function ZERO_HEAVY_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCaller) ZEROHEAVYMESSAGEHEADERFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "ZERO_HEAVY_MESSAGE_HEADER_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// ZEROHEAVYMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x02c99275.
//
// Solidity: function ZERO_HEAVY_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubSession) ZEROHEAVYMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.ZEROHEAVYMESSAGEHEADERFLAG(&_SequencerInboxStub.CallOpts)
}

// ZEROHEAVYMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x02c99275.
//
// Solidity: function ZERO_HEAVY_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) ZEROHEAVYMESSAGEHEADERFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.ZEROHEAVYMESSAGEHEADERFLAG(&_SequencerInboxStub.CallOpts)
}

// AddSequencerL2BatchFromOrigin6f12b0c9 is a free data retrieval call binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 , bytes , uint256 , address ) pure returns()
func (_SequencerInboxStub *SequencerInboxStubCaller) AddSequencerL2BatchFromOrigin6f12b0c9(opts *bind.CallOpts, arg0 *big.Int, arg1 []byte, arg2 *big.Int, arg3 common.Address) error {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "addSequencerL2BatchFromOrigin", arg0, arg1, arg2, arg3)

	if err != nil {
		return err
	}

	return err

}

// AddSequencerL2BatchFromOrigin6f12b0c9 is a free data retrieval call binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 , bytes , uint256 , address ) pure returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddSequencerL2BatchFromOrigin6f12b0c9(arg0 *big.Int, arg1 []byte, arg2 *big.Int, arg3 common.Address) error {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromOrigin6f12b0c9(&_SequencerInboxStub.CallOpts, arg0, arg1, arg2, arg3)
}

// AddSequencerL2BatchFromOrigin6f12b0c9 is a free data retrieval call binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 , bytes , uint256 , address ) pure returns()
func (_SequencerInboxStub *SequencerInboxStubCallerSession) AddSequencerL2BatchFromOrigin6f12b0c9(arg0 *big.Int, arg1 []byte, arg2 *big.Int, arg3 common.Address) error {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromOrigin6f12b0c9(&_SequencerInboxStub.CallOpts, arg0, arg1, arg2, arg3)
}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCaller) BatchCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "batchCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubSession) BatchCount() (*big.Int, error) {
	return _SequencerInboxStub.Contract.BatchCount(&_SequencerInboxStub.CallOpts)
}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) BatchCount() (*big.Int, error) {
	return _SequencerInboxStub.Contract.BatchCount(&_SequencerInboxStub.CallOpts)
}

// BatchPosterManager is a free data retrieval call binding the contract method 0xcc2a1a0c.
//
// Solidity: function batchPosterManager() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCaller) BatchPosterManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "batchPosterManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BatchPosterManager is a free data retrieval call binding the contract method 0xcc2a1a0c.
//
// Solidity: function batchPosterManager() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubSession) BatchPosterManager() (common.Address, error) {
	return _SequencerInboxStub.Contract.BatchPosterManager(&_SequencerInboxStub.CallOpts)
}

// BatchPosterManager is a free data retrieval call binding the contract method 0xcc2a1a0c.
//
// Solidity: function batchPosterManager() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) BatchPosterManager() (common.Address, error) {
	return _SequencerInboxStub.Contract.BatchPosterManager(&_SequencerInboxStub.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubSession) Bridge() (common.Address, error) {
	return _SequencerInboxStub.Contract.Bridge(&_SequencerInboxStub.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) Bridge() (common.Address, error) {
	return _SequencerInboxStub.Contract.Bridge(&_SequencerInboxStub.CallOpts)
}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint64 bufferBlocks, uint64 max, uint64 threshold, uint64 prevBlockNumber, uint64 replenishRateInBasis, uint64 prevSequencedBlockNumber)
func (_SequencerInboxStub *SequencerInboxStubCaller) Buffer(opts *bind.CallOpts) (struct {
	BufferBlocks             uint64
	Max                      uint64
	Threshold                uint64
	PrevBlockNumber          uint64
	ReplenishRateInBasis     uint64
	PrevSequencedBlockNumber uint64
}, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "buffer")

	outstruct := new(struct {
		BufferBlocks             uint64
		Max                      uint64
		Threshold                uint64
		PrevBlockNumber          uint64
		ReplenishRateInBasis     uint64
		PrevSequencedBlockNumber uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BufferBlocks = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Max = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Threshold = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.PrevBlockNumber = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.ReplenishRateInBasis = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.PrevSequencedBlockNumber = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint64 bufferBlocks, uint64 max, uint64 threshold, uint64 prevBlockNumber, uint64 replenishRateInBasis, uint64 prevSequencedBlockNumber)
func (_SequencerInboxStub *SequencerInboxStubSession) Buffer() (struct {
	BufferBlocks             uint64
	Max                      uint64
	Threshold                uint64
	PrevBlockNumber          uint64
	ReplenishRateInBasis     uint64
	PrevSequencedBlockNumber uint64
}, error) {
	return _SequencerInboxStub.Contract.Buffer(&_SequencerInboxStub.CallOpts)
}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint64 bufferBlocks, uint64 max, uint64 threshold, uint64 prevBlockNumber, uint64 replenishRateInBasis, uint64 prevSequencedBlockNumber)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) Buffer() (struct {
	BufferBlocks             uint64
	Max                      uint64
	Threshold                uint64
	PrevBlockNumber          uint64
	ReplenishRateInBasis     uint64
	PrevSequencedBlockNumber uint64
}, error) {
	return _SequencerInboxStub.Contract.Buffer(&_SequencerInboxStub.CallOpts)
}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(bool isValidKeyset, uint64 creationBlock)
func (_SequencerInboxStub *SequencerInboxStubCaller) DasKeySetInfo(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "dasKeySetInfo", arg0)

	outstruct := new(struct {
		IsValidKeyset bool
		CreationBlock uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsValidKeyset = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CreationBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(bool isValidKeyset, uint64 creationBlock)
func (_SequencerInboxStub *SequencerInboxStubSession) DasKeySetInfo(arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	return _SequencerInboxStub.Contract.DasKeySetInfo(&_SequencerInboxStub.CallOpts, arg0)
}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(bool isValidKeyset, uint64 creationBlock)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) DasKeySetInfo(arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	return _SequencerInboxStub.Contract.DasKeySetInfo(&_SequencerInboxStub.CallOpts, arg0)
}

// FeeTokenPricer is a free data retrieval call binding the contract method 0x22291e8d.
//
// Solidity: function feeTokenPricer() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCaller) FeeTokenPricer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "feeTokenPricer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenPricer is a free data retrieval call binding the contract method 0x22291e8d.
//
// Solidity: function feeTokenPricer() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubSession) FeeTokenPricer() (common.Address, error) {
	return _SequencerInboxStub.Contract.FeeTokenPricer(&_SequencerInboxStub.CallOpts)
}

// FeeTokenPricer is a free data retrieval call binding the contract method 0x22291e8d.
//
// Solidity: function feeTokenPricer() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) FeeTokenPricer() (common.Address, error) {
	return _SequencerInboxStub.Contract.FeeTokenPricer(&_SequencerInboxStub.CallOpts)
}

// ForceInclusionDeadline is a free data retrieval call binding the contract method 0xdd44e6e0.
//
// Solidity: function forceInclusionDeadline(uint64 blockNumber) view returns(uint64)
func (_SequencerInboxStub *SequencerInboxStubCaller) ForceInclusionDeadline(opts *bind.CallOpts, blockNumber uint64) (uint64, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "forceInclusionDeadline", blockNumber)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceInclusionDeadline is a free data retrieval call binding the contract method 0xdd44e6e0.
//
// Solidity: function forceInclusionDeadline(uint64 blockNumber) view returns(uint64)
func (_SequencerInboxStub *SequencerInboxStubSession) ForceInclusionDeadline(blockNumber uint64) (uint64, error) {
	return _SequencerInboxStub.Contract.ForceInclusionDeadline(&_SequencerInboxStub.CallOpts, blockNumber)
}

// ForceInclusionDeadline is a free data retrieval call binding the contract method 0xdd44e6e0.
//
// Solidity: function forceInclusionDeadline(uint64 blockNumber) view returns(uint64)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) ForceInclusionDeadline(blockNumber uint64) (uint64, error) {
	return _SequencerInboxStub.Contract.ForceInclusionDeadline(&_SequencerInboxStub.CallOpts, blockNumber)
}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCaller) GetKeysetCreationBlock(opts *bind.CallOpts, ksHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "getKeysetCreationBlock", ksHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubSession) GetKeysetCreationBlock(ksHash [32]byte) (*big.Int, error) {
	return _SequencerInboxStub.Contract.GetKeysetCreationBlock(&_SequencerInboxStub.CallOpts, ksHash)
}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) GetKeysetCreationBlock(ksHash [32]byte) (*big.Int, error) {
	return _SequencerInboxStub.Contract.GetKeysetCreationBlock(&_SequencerInboxStub.CallOpts, ksHash)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInboxStub *SequencerInboxStubCaller) InboxAccs(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "inboxAccs", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInboxStub *SequencerInboxStubSession) InboxAccs(index *big.Int) ([32]byte, error) {
	return _SequencerInboxStub.Contract.InboxAccs(&_SequencerInboxStub.CallOpts, index)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) InboxAccs(index *big.Int) ([32]byte, error) {
	return _SequencerInboxStub.Contract.InboxAccs(&_SequencerInboxStub.CallOpts, index)
}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCaller) IsBatchPoster(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "isBatchPoster", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubSession) IsBatchPoster(arg0 common.Address) (bool, error) {
	return _SequencerInboxStub.Contract.IsBatchPoster(&_SequencerInboxStub.CallOpts, arg0)
}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) IsBatchPoster(arg0 common.Address) (bool, error) {
	return _SequencerInboxStub.Contract.IsBatchPoster(&_SequencerInboxStub.CallOpts, arg0)
}

// IsDelayBufferable is a free data retrieval call binding the contract method 0x4b678a66.
//
// Solidity: function isDelayBufferable() view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCaller) IsDelayBufferable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "isDelayBufferable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelayBufferable is a free data retrieval call binding the contract method 0x4b678a66.
//
// Solidity: function isDelayBufferable() view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubSession) IsDelayBufferable() (bool, error) {
	return _SequencerInboxStub.Contract.IsDelayBufferable(&_SequencerInboxStub.CallOpts)
}

// IsDelayBufferable is a free data retrieval call binding the contract method 0x4b678a66.
//
// Solidity: function isDelayBufferable() view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) IsDelayBufferable() (bool, error) {
	return _SequencerInboxStub.Contract.IsDelayBufferable(&_SequencerInboxStub.CallOpts)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCaller) IsSequencer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "isSequencer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubSession) IsSequencer(arg0 common.Address) (bool, error) {
	return _SequencerInboxStub.Contract.IsSequencer(&_SequencerInboxStub.CallOpts, arg0)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) IsSequencer(arg0 common.Address) (bool, error) {
	return _SequencerInboxStub.Contract.IsSequencer(&_SequencerInboxStub.CallOpts, arg0)
}

// IsUsingFeeToken is a free data retrieval call binding the contract method 0x92d9f782.
//
// Solidity: function isUsingFeeToken() view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCaller) IsUsingFeeToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "isUsingFeeToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUsingFeeToken is a free data retrieval call binding the contract method 0x92d9f782.
//
// Solidity: function isUsingFeeToken() view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubSession) IsUsingFeeToken() (bool, error) {
	return _SequencerInboxStub.Contract.IsUsingFeeToken(&_SequencerInboxStub.CallOpts)
}

// IsUsingFeeToken is a free data retrieval call binding the contract method 0x92d9f782.
//
// Solidity: function isUsingFeeToken() view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) IsUsingFeeToken() (bool, error) {
	return _SequencerInboxStub.Contract.IsUsingFeeToken(&_SequencerInboxStub.CallOpts)
}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCaller) IsValidKeysetHash(opts *bind.CallOpts, ksHash [32]byte) (bool, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "isValidKeysetHash", ksHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubSession) IsValidKeysetHash(ksHash [32]byte) (bool, error) {
	return _SequencerInboxStub.Contract.IsValidKeysetHash(&_SequencerInboxStub.CallOpts, ksHash)
}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) IsValidKeysetHash(ksHash [32]byte) (bool, error) {
	return _SequencerInboxStub.Contract.IsValidKeysetHash(&_SequencerInboxStub.CallOpts, ksHash)
}

// MaxDataSize is a free data retrieval call binding the contract method 0xe8eb1dc3.
//
// Solidity: function maxDataSize() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCaller) MaxDataSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "maxDataSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDataSize is a free data retrieval call binding the contract method 0xe8eb1dc3.
//
// Solidity: function maxDataSize() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubSession) MaxDataSize() (*big.Int, error) {
	return _SequencerInboxStub.Contract.MaxDataSize(&_SequencerInboxStub.CallOpts)
}

// MaxDataSize is a free data retrieval call binding the contract method 0xe8eb1dc3.
//
// Solidity: function maxDataSize() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) MaxDataSize() (*big.Int, error) {
	return _SequencerInboxStub.Contract.MaxDataSize(&_SequencerInboxStub.CallOpts)
}

// MaxTimeVariation is a free data retrieval call binding the contract method 0xebea461d.
//
// Solidity: function maxTimeVariation() view returns(uint256, uint256, uint256, uint256)
func (_SequencerInboxStub *SequencerInboxStubCaller) MaxTimeVariation(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "maxTimeVariation")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// MaxTimeVariation is a free data retrieval call binding the contract method 0xebea461d.
//
// Solidity: function maxTimeVariation() view returns(uint256, uint256, uint256, uint256)
func (_SequencerInboxStub *SequencerInboxStubSession) MaxTimeVariation() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _SequencerInboxStub.Contract.MaxTimeVariation(&_SequencerInboxStub.CallOpts)
}

// MaxTimeVariation is a free data retrieval call binding the contract method 0xebea461d.
//
// Solidity: function maxTimeVariation() view returns(uint256, uint256, uint256, uint256)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) MaxTimeVariation() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _SequencerInboxStub.Contract.MaxTimeVariation(&_SequencerInboxStub.CallOpts)
}

// Reader4844 is a free data retrieval call binding the contract method 0x8d910dde.
//
// Solidity: function reader4844() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCaller) Reader4844(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "reader4844")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Reader4844 is a free data retrieval call binding the contract method 0x8d910dde.
//
// Solidity: function reader4844() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubSession) Reader4844() (common.Address, error) {
	return _SequencerInboxStub.Contract.Reader4844(&_SequencerInboxStub.CallOpts)
}

// Reader4844 is a free data retrieval call binding the contract method 0x8d910dde.
//
// Solidity: function reader4844() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) Reader4844() (common.Address, error) {
	return _SequencerInboxStub.Contract.Reader4844(&_SequencerInboxStub.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubSession) Rollup() (common.Address, error) {
	return _SequencerInboxStub.Contract.Rollup(&_SequencerInboxStub.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) Rollup() (common.Address, error) {
	return _SequencerInboxStub.Contract.Rollup(&_SequencerInboxStub.CallOpts)
}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCaller) TotalDelayedMessagesRead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "totalDelayedMessagesRead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubSession) TotalDelayedMessagesRead() (*big.Int, error) {
	return _SequencerInboxStub.Contract.TotalDelayedMessagesRead(&_SequencerInboxStub.CallOpts)
}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) TotalDelayedMessagesRead() (*big.Int, error) {
	return _SequencerInboxStub.Contract.TotalDelayedMessagesRead(&_SequencerInboxStub.CallOpts)
}

// AddInitMessage is a paid mutator transaction binding the contract method 0x6633ae85.
//
// Solidity: function addInitMessage(uint256 chainId) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) AddInitMessage(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "addInitMessage", chainId)
}

// AddInitMessage is a paid mutator transaction binding the contract method 0x6633ae85.
//
// Solidity: function addInitMessage(uint256 chainId) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddInitMessage(chainId *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddInitMessage(&_SequencerInboxStub.TransactOpts, chainId)
}

// AddInitMessage is a paid mutator transaction binding the contract method 0x6633ae85.
//
// Solidity: function addInitMessage(uint256 chainId) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) AddInitMessage(chainId *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddInitMessage(&_SequencerInboxStub.TransactOpts, chainId)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0xe0bc9729.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) AddSequencerL2Batch(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "addSequencerL2Batch", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0xe0bc9729.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddSequencerL2Batch(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2Batch(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0xe0bc9729.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) AddSequencerL2Batch(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2Batch(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchDelayProof is a paid mutator transaction binding the contract method 0x6e620055.
//
// Solidity: function addSequencerL2BatchDelayProof(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) AddSequencerL2BatchDelayProof(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "addSequencerL2BatchDelayProof", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchDelayProof is a paid mutator transaction binding the contract method 0x6e620055.
//
// Solidity: function addSequencerL2BatchDelayProof(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddSequencerL2BatchDelayProof(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchDelayProof(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchDelayProof is a paid mutator transaction binding the contract method 0x6e620055.
//
// Solidity: function addSequencerL2BatchDelayProof(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) AddSequencerL2BatchDelayProof(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchDelayProof(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchFromBlobs is a paid mutator transaction binding the contract method 0x3e5aa082.
//
// Solidity: function addSequencerL2BatchFromBlobs(uint256 sequenceNumber, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) AddSequencerL2BatchFromBlobs(opts *bind.TransactOpts, sequenceNumber *big.Int, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "addSequencerL2BatchFromBlobs", sequenceNumber, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromBlobs is a paid mutator transaction binding the contract method 0x3e5aa082.
//
// Solidity: function addSequencerL2BatchFromBlobs(uint256 sequenceNumber, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddSequencerL2BatchFromBlobs(sequenceNumber *big.Int, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromBlobs(&_SequencerInboxStub.TransactOpts, sequenceNumber, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromBlobs is a paid mutator transaction binding the contract method 0x3e5aa082.
//
// Solidity: function addSequencerL2BatchFromBlobs(uint256 sequenceNumber, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) AddSequencerL2BatchFromBlobs(sequenceNumber *big.Int, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromBlobs(&_SequencerInboxStub.TransactOpts, sequenceNumber, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromBlobsDelayProof is a paid mutator transaction binding the contract method 0x917cf8ac.
//
// Solidity: function addSequencerL2BatchFromBlobsDelayProof(uint256 sequenceNumber, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) AddSequencerL2BatchFromBlobsDelayProof(opts *bind.TransactOpts, sequenceNumber *big.Int, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "addSequencerL2BatchFromBlobsDelayProof", sequenceNumber, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchFromBlobsDelayProof is a paid mutator transaction binding the contract method 0x917cf8ac.
//
// Solidity: function addSequencerL2BatchFromBlobsDelayProof(uint256 sequenceNumber, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddSequencerL2BatchFromBlobsDelayProof(sequenceNumber *big.Int, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromBlobsDelayProof(&_SequencerInboxStub.TransactOpts, sequenceNumber, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchFromBlobsDelayProof is a paid mutator transaction binding the contract method 0x917cf8ac.
//
// Solidity: function addSequencerL2BatchFromBlobsDelayProof(uint256 sequenceNumber, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) AddSequencerL2BatchFromBlobsDelayProof(sequenceNumber *big.Int, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromBlobsDelayProof(&_SequencerInboxStub.TransactOpts, sequenceNumber, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchFromOrigin8f111f3c is a paid mutator transaction binding the contract method 0x8f111f3c.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) AddSequencerL2BatchFromOrigin8f111f3c(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "addSequencerL2BatchFromOrigin0", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromOrigin8f111f3c is a paid mutator transaction binding the contract method 0x8f111f3c.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddSequencerL2BatchFromOrigin8f111f3c(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromOrigin8f111f3c(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromOrigin8f111f3c is a paid mutator transaction binding the contract method 0x8f111f3c.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) AddSequencerL2BatchFromOrigin8f111f3c(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromOrigin8f111f3c(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromOriginDelayProof is a paid mutator transaction binding the contract method 0x69cacded.
//
// Solidity: function addSequencerL2BatchFromOriginDelayProof(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) AddSequencerL2BatchFromOriginDelayProof(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "addSequencerL2BatchFromOriginDelayProof", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchFromOriginDelayProof is a paid mutator transaction binding the contract method 0x69cacded.
//
// Solidity: function addSequencerL2BatchFromOriginDelayProof(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddSequencerL2BatchFromOriginDelayProof(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromOriginDelayProof(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// AddSequencerL2BatchFromOriginDelayProof is a paid mutator transaction binding the contract method 0x69cacded.
//
// Solidity: function addSequencerL2BatchFromOriginDelayProof(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount, (bytes32,(uint8,address,uint64,uint64,uint256,uint256,bytes32)) delayProof) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) AddSequencerL2BatchFromOriginDelayProof(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int, delayProof DelayProof) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromOriginDelayProof(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount, delayProof)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) ForceInclusion(opts *bind.TransactOpts, _totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "forceInclusion", _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.ForceInclusion(&_SequencerInboxStub.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.ForceInclusion(&_SequencerInboxStub.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xa84840b7.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_, (uint64,uint64,uint64) bufferConfig_, address feeTokenPricer_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) Initialize(opts *bind.TransactOpts, bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation, bufferConfig_ BufferConfig, feeTokenPricer_ common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "initialize", bridge_, maxTimeVariation_, bufferConfig_, feeTokenPricer_)
}

// Initialize is a paid mutator transaction binding the contract method 0xa84840b7.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_, (uint64,uint64,uint64) bufferConfig_, address feeTokenPricer_) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) Initialize(bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation, bufferConfig_ BufferConfig, feeTokenPricer_ common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.Initialize(&_SequencerInboxStub.TransactOpts, bridge_, maxTimeVariation_, bufferConfig_, feeTokenPricer_)
}

// Initialize is a paid mutator transaction binding the contract method 0xa84840b7.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_, (uint64,uint64,uint64) bufferConfig_, address feeTokenPricer_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) Initialize(bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation, bufferConfig_ BufferConfig, feeTokenPricer_ common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.Initialize(&_SequencerInboxStub.TransactOpts, bridge_, maxTimeVariation_, bufferConfig_, feeTokenPricer_)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) InvalidateKeysetHash(opts *bind.TransactOpts, ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "invalidateKeysetHash", ksHash)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) InvalidateKeysetHash(ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.InvalidateKeysetHash(&_SequencerInboxStub.TransactOpts, ksHash)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) InvalidateKeysetHash(ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.InvalidateKeysetHash(&_SequencerInboxStub.TransactOpts, ksHash)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xa655d937.
//
// Solidity: function postUpgradeInit((uint64,uint64,uint64) bufferConfig_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) PostUpgradeInit(opts *bind.TransactOpts, bufferConfig_ BufferConfig) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "postUpgradeInit", bufferConfig_)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xa655d937.
//
// Solidity: function postUpgradeInit((uint64,uint64,uint64) bufferConfig_) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) PostUpgradeInit(bufferConfig_ BufferConfig) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.PostUpgradeInit(&_SequencerInboxStub.TransactOpts, bufferConfig_)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xa655d937.
//
// Solidity: function postUpgradeInit((uint64,uint64,uint64) bufferConfig_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) PostUpgradeInit(bufferConfig_ BufferConfig) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.PostUpgradeInit(&_SequencerInboxStub.TransactOpts, bufferConfig_)
}

// RemoveDelayAfterFork is a paid mutator transaction binding the contract method 0x96cc5c78.
//
// Solidity: function removeDelayAfterFork() returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) RemoveDelayAfterFork(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "removeDelayAfterFork")
}

// RemoveDelayAfterFork is a paid mutator transaction binding the contract method 0x96cc5c78.
//
// Solidity: function removeDelayAfterFork() returns()
func (_SequencerInboxStub *SequencerInboxStubSession) RemoveDelayAfterFork() (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.RemoveDelayAfterFork(&_SequencerInboxStub.TransactOpts)
}

// RemoveDelayAfterFork is a paid mutator transaction binding the contract method 0x96cc5c78.
//
// Solidity: function removeDelayAfterFork() returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) RemoveDelayAfterFork() (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.RemoveDelayAfterFork(&_SequencerInboxStub.TransactOpts)
}

// SetBatchPosterManager is a paid mutator transaction binding the contract method 0x1ff64790.
//
// Solidity: function setBatchPosterManager(address newBatchPosterManager) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) SetBatchPosterManager(opts *bind.TransactOpts, newBatchPosterManager common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "setBatchPosterManager", newBatchPosterManager)
}

// SetBatchPosterManager is a paid mutator transaction binding the contract method 0x1ff64790.
//
// Solidity: function setBatchPosterManager(address newBatchPosterManager) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) SetBatchPosterManager(newBatchPosterManager common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetBatchPosterManager(&_SequencerInboxStub.TransactOpts, newBatchPosterManager)
}

// SetBatchPosterManager is a paid mutator transaction binding the contract method 0x1ff64790.
//
// Solidity: function setBatchPosterManager(address newBatchPosterManager) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) SetBatchPosterManager(newBatchPosterManager common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetBatchPosterManager(&_SequencerInboxStub.TransactOpts, newBatchPosterManager)
}

// SetBufferConfig is a paid mutator transaction binding the contract method 0x2f3985a7.
//
// Solidity: function setBufferConfig((uint64,uint64,uint64) bufferConfig_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) SetBufferConfig(opts *bind.TransactOpts, bufferConfig_ BufferConfig) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "setBufferConfig", bufferConfig_)
}

// SetBufferConfig is a paid mutator transaction binding the contract method 0x2f3985a7.
//
// Solidity: function setBufferConfig((uint64,uint64,uint64) bufferConfig_) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) SetBufferConfig(bufferConfig_ BufferConfig) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetBufferConfig(&_SequencerInboxStub.TransactOpts, bufferConfig_)
}

// SetBufferConfig is a paid mutator transaction binding the contract method 0x2f3985a7.
//
// Solidity: function setBufferConfig((uint64,uint64,uint64) bufferConfig_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) SetBufferConfig(bufferConfig_ BufferConfig) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetBufferConfig(&_SequencerInboxStub.TransactOpts, bufferConfig_)
}

// SetFeeTokenPricer is a paid mutator transaction binding the contract method 0x036f7ed3.
//
// Solidity: function setFeeTokenPricer(address feeTokenPricer_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) SetFeeTokenPricer(opts *bind.TransactOpts, feeTokenPricer_ common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "setFeeTokenPricer", feeTokenPricer_)
}

// SetFeeTokenPricer is a paid mutator transaction binding the contract method 0x036f7ed3.
//
// Solidity: function setFeeTokenPricer(address feeTokenPricer_) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) SetFeeTokenPricer(feeTokenPricer_ common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetFeeTokenPricer(&_SequencerInboxStub.TransactOpts, feeTokenPricer_)
}

// SetFeeTokenPricer is a paid mutator transaction binding the contract method 0x036f7ed3.
//
// Solidity: function setFeeTokenPricer(address feeTokenPricer_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) SetFeeTokenPricer(feeTokenPricer_ common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetFeeTokenPricer(&_SequencerInboxStub.TransactOpts, feeTokenPricer_)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) SetIsBatchPoster(opts *bind.TransactOpts, addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "setIsBatchPoster", addr, isBatchPoster_)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) SetIsBatchPoster(addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetIsBatchPoster(&_SequencerInboxStub.TransactOpts, addr, isBatchPoster_)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) SetIsBatchPoster(addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetIsBatchPoster(&_SequencerInboxStub.TransactOpts, addr, isBatchPoster_)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool isSequencer_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) SetIsSequencer(opts *bind.TransactOpts, addr common.Address, isSequencer_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "setIsSequencer", addr, isSequencer_)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool isSequencer_) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) SetIsSequencer(addr common.Address, isSequencer_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetIsSequencer(&_SequencerInboxStub.TransactOpts, addr, isSequencer_)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool isSequencer_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) SetIsSequencer(addr common.Address, isSequencer_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetIsSequencer(&_SequencerInboxStub.TransactOpts, addr, isSequencer_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) SetMaxTimeVariation(opts *bind.TransactOpts, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "setMaxTimeVariation", maxTimeVariation_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) SetMaxTimeVariation(maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetMaxTimeVariation(&_SequencerInboxStub.TransactOpts, maxTimeVariation_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) SetMaxTimeVariation(maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetMaxTimeVariation(&_SequencerInboxStub.TransactOpts, maxTimeVariation_)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) SetValidKeyset(opts *bind.TransactOpts, keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "setValidKeyset", keysetBytes)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) SetValidKeyset(keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetValidKeyset(&_SequencerInboxStub.TransactOpts, keysetBytes)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) SetValidKeyset(keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetValidKeyset(&_SequencerInboxStub.TransactOpts, keysetBytes)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) UpdateRollupAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "updateRollupAddress")
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_SequencerInboxStub *SequencerInboxStubSession) UpdateRollupAddress() (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.UpdateRollupAddress(&_SequencerInboxStub.TransactOpts)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) UpdateRollupAddress() (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.UpdateRollupAddress(&_SequencerInboxStub.TransactOpts)
}

// SequencerInboxStubBatchPosterManagerSetIterator is returned from FilterBatchPosterManagerSet and is used to iterate over the raw logs and unpacked data for BatchPosterManagerSet events raised by the SequencerInboxStub contract.
type SequencerInboxStubBatchPosterManagerSetIterator struct {
	Event *SequencerInboxStubBatchPosterManagerSet // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubBatchPosterManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubBatchPosterManagerSet)
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
		it.Event = new(SequencerInboxStubBatchPosterManagerSet)
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
func (it *SequencerInboxStubBatchPosterManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubBatchPosterManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubBatchPosterManagerSet represents a BatchPosterManagerSet event raised by the SequencerInboxStub contract.
type SequencerInboxStubBatchPosterManagerSet struct {
	NewBatchPosterManager common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterBatchPosterManagerSet is a free log retrieval operation binding the contract event 0x3cd6c184800297a0f2b00926a683cbe76890bb7fd01480ac0a10ed6c8f7f6659.
//
// Solidity: event BatchPosterManagerSet(address newBatchPosterManager)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterBatchPosterManagerSet(opts *bind.FilterOpts) (*SequencerInboxStubBatchPosterManagerSetIterator, error) {

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "BatchPosterManagerSet")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubBatchPosterManagerSetIterator{contract: _SequencerInboxStub.contract, event: "BatchPosterManagerSet", logs: logs, sub: sub}, nil
}

// WatchBatchPosterManagerSet is a free log subscription operation binding the contract event 0x3cd6c184800297a0f2b00926a683cbe76890bb7fd01480ac0a10ed6c8f7f6659.
//
// Solidity: event BatchPosterManagerSet(address newBatchPosterManager)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchBatchPosterManagerSet(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubBatchPosterManagerSet) (event.Subscription, error) {

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "BatchPosterManagerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubBatchPosterManagerSet)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "BatchPosterManagerSet", log); err != nil {
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

// ParseBatchPosterManagerSet is a log parse operation binding the contract event 0x3cd6c184800297a0f2b00926a683cbe76890bb7fd01480ac0a10ed6c8f7f6659.
//
// Solidity: event BatchPosterManagerSet(address newBatchPosterManager)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseBatchPosterManagerSet(log types.Log) (*SequencerInboxStubBatchPosterManagerSet, error) {
	event := new(SequencerInboxStubBatchPosterManagerSet)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "BatchPosterManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubBatchPosterSetIterator is returned from FilterBatchPosterSet and is used to iterate over the raw logs and unpacked data for BatchPosterSet events raised by the SequencerInboxStub contract.
type SequencerInboxStubBatchPosterSetIterator struct {
	Event *SequencerInboxStubBatchPosterSet // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubBatchPosterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubBatchPosterSet)
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
		it.Event = new(SequencerInboxStubBatchPosterSet)
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
func (it *SequencerInboxStubBatchPosterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubBatchPosterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubBatchPosterSet represents a BatchPosterSet event raised by the SequencerInboxStub contract.
type SequencerInboxStubBatchPosterSet struct {
	BatchPoster   common.Address
	IsBatchPoster bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBatchPosterSet is a free log retrieval operation binding the contract event 0x28bcc5626d357efe966b4b0876aa1ee8ab99e26da4f131f6a2623f1800701c21.
//
// Solidity: event BatchPosterSet(address batchPoster, bool isBatchPoster)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterBatchPosterSet(opts *bind.FilterOpts) (*SequencerInboxStubBatchPosterSetIterator, error) {

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "BatchPosterSet")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubBatchPosterSetIterator{contract: _SequencerInboxStub.contract, event: "BatchPosterSet", logs: logs, sub: sub}, nil
}

// WatchBatchPosterSet is a free log subscription operation binding the contract event 0x28bcc5626d357efe966b4b0876aa1ee8ab99e26da4f131f6a2623f1800701c21.
//
// Solidity: event BatchPosterSet(address batchPoster, bool isBatchPoster)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchBatchPosterSet(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubBatchPosterSet) (event.Subscription, error) {

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "BatchPosterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubBatchPosterSet)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "BatchPosterSet", log); err != nil {
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

// ParseBatchPosterSet is a log parse operation binding the contract event 0x28bcc5626d357efe966b4b0876aa1ee8ab99e26da4f131f6a2623f1800701c21.
//
// Solidity: event BatchPosterSet(address batchPoster, bool isBatchPoster)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseBatchPosterSet(log types.Log) (*SequencerInboxStubBatchPosterSet, error) {
	event := new(SequencerInboxStubBatchPosterSet)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "BatchPosterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubBufferConfigSetIterator is returned from FilterBufferConfigSet and is used to iterate over the raw logs and unpacked data for BufferConfigSet events raised by the SequencerInboxStub contract.
type SequencerInboxStubBufferConfigSetIterator struct {
	Event *SequencerInboxStubBufferConfigSet // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubBufferConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubBufferConfigSet)
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
		it.Event = new(SequencerInboxStubBufferConfigSet)
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
func (it *SequencerInboxStubBufferConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubBufferConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubBufferConfigSet represents a BufferConfigSet event raised by the SequencerInboxStub contract.
type SequencerInboxStubBufferConfigSet struct {
	BufferConfig BufferConfig
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBufferConfigSet is a free log retrieval operation binding the contract event 0xaa7a2d8175dee3b637814ad6346005dfcc357165396fb8327f649effe8abcf85.
//
// Solidity: event BufferConfigSet((uint64,uint64,uint64) bufferConfig)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterBufferConfigSet(opts *bind.FilterOpts) (*SequencerInboxStubBufferConfigSetIterator, error) {

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "BufferConfigSet")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubBufferConfigSetIterator{contract: _SequencerInboxStub.contract, event: "BufferConfigSet", logs: logs, sub: sub}, nil
}

// WatchBufferConfigSet is a free log subscription operation binding the contract event 0xaa7a2d8175dee3b637814ad6346005dfcc357165396fb8327f649effe8abcf85.
//
// Solidity: event BufferConfigSet((uint64,uint64,uint64) bufferConfig)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchBufferConfigSet(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubBufferConfigSet) (event.Subscription, error) {

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "BufferConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubBufferConfigSet)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "BufferConfigSet", log); err != nil {
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

// ParseBufferConfigSet is a log parse operation binding the contract event 0xaa7a2d8175dee3b637814ad6346005dfcc357165396fb8327f649effe8abcf85.
//
// Solidity: event BufferConfigSet((uint64,uint64,uint64) bufferConfig)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseBufferConfigSet(log types.Log) (*SequencerInboxStubBufferConfigSet, error) {
	event := new(SequencerInboxStubBufferConfigSet)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "BufferConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubFeeTokenPricerSetIterator is returned from FilterFeeTokenPricerSet and is used to iterate over the raw logs and unpacked data for FeeTokenPricerSet events raised by the SequencerInboxStub contract.
type SequencerInboxStubFeeTokenPricerSetIterator struct {
	Event *SequencerInboxStubFeeTokenPricerSet // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubFeeTokenPricerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubFeeTokenPricerSet)
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
		it.Event = new(SequencerInboxStubFeeTokenPricerSet)
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
func (it *SequencerInboxStubFeeTokenPricerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubFeeTokenPricerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubFeeTokenPricerSet represents a FeeTokenPricerSet event raised by the SequencerInboxStub contract.
type SequencerInboxStubFeeTokenPricerSet struct {
	FeeTokenPricer common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFeeTokenPricerSet is a free log retrieval operation binding the contract event 0xe83d6153add50e41b8ee6c1115c4178687349bb12bc3902a50b1f6ad78a0c541.
//
// Solidity: event FeeTokenPricerSet(address feeTokenPricer)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterFeeTokenPricerSet(opts *bind.FilterOpts) (*SequencerInboxStubFeeTokenPricerSetIterator, error) {

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "FeeTokenPricerSet")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubFeeTokenPricerSetIterator{contract: _SequencerInboxStub.contract, event: "FeeTokenPricerSet", logs: logs, sub: sub}, nil
}

// WatchFeeTokenPricerSet is a free log subscription operation binding the contract event 0xe83d6153add50e41b8ee6c1115c4178687349bb12bc3902a50b1f6ad78a0c541.
//
// Solidity: event FeeTokenPricerSet(address feeTokenPricer)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchFeeTokenPricerSet(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubFeeTokenPricerSet) (event.Subscription, error) {

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "FeeTokenPricerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubFeeTokenPricerSet)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "FeeTokenPricerSet", log); err != nil {
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

// ParseFeeTokenPricerSet is a log parse operation binding the contract event 0xe83d6153add50e41b8ee6c1115c4178687349bb12bc3902a50b1f6ad78a0c541.
//
// Solidity: event FeeTokenPricerSet(address feeTokenPricer)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseFeeTokenPricerSet(log types.Log) (*SequencerInboxStubFeeTokenPricerSet, error) {
	event := new(SequencerInboxStubFeeTokenPricerSet)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "FeeTokenPricerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the SequencerInboxStub contract.
type SequencerInboxStubInboxMessageDeliveredIterator struct {
	Event *SequencerInboxStubInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubInboxMessageDelivered)
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
		it.Event = new(SequencerInboxStubInboxMessageDelivered)
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
func (it *SequencerInboxStubInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubInboxMessageDelivered represents a InboxMessageDelivered event raised by the SequencerInboxStub contract.
type SequencerInboxStubInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*SequencerInboxStubInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubInboxMessageDeliveredIterator{contract: _SequencerInboxStub.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubInboxMessageDelivered)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseInboxMessageDelivered(log types.Log) (*SequencerInboxStubInboxMessageDelivered, error) {
	event := new(SequencerInboxStubInboxMessageDelivered)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the SequencerInboxStub contract.
type SequencerInboxStubInboxMessageDeliveredFromOriginIterator struct {
	Event *SequencerInboxStubInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubInboxMessageDeliveredFromOrigin)
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
		it.Event = new(SequencerInboxStubInboxMessageDeliveredFromOrigin)
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
func (it *SequencerInboxStubInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the SequencerInboxStub contract.
type SequencerInboxStubInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*SequencerInboxStubInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubInboxMessageDeliveredFromOriginIterator{contract: _SequencerInboxStub.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubInboxMessageDeliveredFromOrigin)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*SequencerInboxStubInboxMessageDeliveredFromOrigin, error) {
	event := new(SequencerInboxStubInboxMessageDeliveredFromOrigin)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubInvalidateKeysetIterator is returned from FilterInvalidateKeyset and is used to iterate over the raw logs and unpacked data for InvalidateKeyset events raised by the SequencerInboxStub contract.
type SequencerInboxStubInvalidateKeysetIterator struct {
	Event *SequencerInboxStubInvalidateKeyset // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubInvalidateKeysetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubInvalidateKeyset)
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
		it.Event = new(SequencerInboxStubInvalidateKeyset)
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
func (it *SequencerInboxStubInvalidateKeysetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubInvalidateKeysetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubInvalidateKeyset represents a InvalidateKeyset event raised by the SequencerInboxStub contract.
type SequencerInboxStubInvalidateKeyset struct {
	KeysetHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvalidateKeyset is a free log retrieval operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterInvalidateKeyset(opts *bind.FilterOpts, keysetHash [][32]byte) (*SequencerInboxStubInvalidateKeysetIterator, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "InvalidateKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubInvalidateKeysetIterator{contract: _SequencerInboxStub.contract, event: "InvalidateKeyset", logs: logs, sub: sub}, nil
}

// WatchInvalidateKeyset is a free log subscription operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchInvalidateKeyset(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubInvalidateKeyset, keysetHash [][32]byte) (event.Subscription, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "InvalidateKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubInvalidateKeyset)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "InvalidateKeyset", log); err != nil {
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

// ParseInvalidateKeyset is a log parse operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseInvalidateKeyset(log types.Log) (*SequencerInboxStubInvalidateKeyset, error) {
	event := new(SequencerInboxStubInvalidateKeyset)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "InvalidateKeyset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubMaxTimeVariationSetIterator is returned from FilterMaxTimeVariationSet and is used to iterate over the raw logs and unpacked data for MaxTimeVariationSet events raised by the SequencerInboxStub contract.
type SequencerInboxStubMaxTimeVariationSetIterator struct {
	Event *SequencerInboxStubMaxTimeVariationSet // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubMaxTimeVariationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubMaxTimeVariationSet)
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
		it.Event = new(SequencerInboxStubMaxTimeVariationSet)
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
func (it *SequencerInboxStubMaxTimeVariationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubMaxTimeVariationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubMaxTimeVariationSet represents a MaxTimeVariationSet event raised by the SequencerInboxStub contract.
type SequencerInboxStubMaxTimeVariationSet struct {
	MaxTimeVariation ISequencerInboxMaxTimeVariation
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMaxTimeVariationSet is a free log retrieval operation binding the contract event 0xaa6a58dad31128ff7ecc2b80987ee6e003df80bc50cd8d0b0d1af0e07da6d19d.
//
// Solidity: event MaxTimeVariationSet((uint256,uint256,uint256,uint256) maxTimeVariation)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterMaxTimeVariationSet(opts *bind.FilterOpts) (*SequencerInboxStubMaxTimeVariationSetIterator, error) {

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "MaxTimeVariationSet")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubMaxTimeVariationSetIterator{contract: _SequencerInboxStub.contract, event: "MaxTimeVariationSet", logs: logs, sub: sub}, nil
}

// WatchMaxTimeVariationSet is a free log subscription operation binding the contract event 0xaa6a58dad31128ff7ecc2b80987ee6e003df80bc50cd8d0b0d1af0e07da6d19d.
//
// Solidity: event MaxTimeVariationSet((uint256,uint256,uint256,uint256) maxTimeVariation)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchMaxTimeVariationSet(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubMaxTimeVariationSet) (event.Subscription, error) {

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "MaxTimeVariationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubMaxTimeVariationSet)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "MaxTimeVariationSet", log); err != nil {
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

// ParseMaxTimeVariationSet is a log parse operation binding the contract event 0xaa6a58dad31128ff7ecc2b80987ee6e003df80bc50cd8d0b0d1af0e07da6d19d.
//
// Solidity: event MaxTimeVariationSet((uint256,uint256,uint256,uint256) maxTimeVariation)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseMaxTimeVariationSet(log types.Log) (*SequencerInboxStubMaxTimeVariationSet, error) {
	event := new(SequencerInboxStubMaxTimeVariationSet)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "MaxTimeVariationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubOwnerFunctionCalledIterator is returned from FilterOwnerFunctionCalled and is used to iterate over the raw logs and unpacked data for OwnerFunctionCalled events raised by the SequencerInboxStub contract.
type SequencerInboxStubOwnerFunctionCalledIterator struct {
	Event *SequencerInboxStubOwnerFunctionCalled // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubOwnerFunctionCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubOwnerFunctionCalled)
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
		it.Event = new(SequencerInboxStubOwnerFunctionCalled)
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
func (it *SequencerInboxStubOwnerFunctionCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubOwnerFunctionCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubOwnerFunctionCalled represents a OwnerFunctionCalled event raised by the SequencerInboxStub contract.
type SequencerInboxStubOwnerFunctionCalled struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOwnerFunctionCalled is a free log retrieval operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterOwnerFunctionCalled(opts *bind.FilterOpts, id []*big.Int) (*SequencerInboxStubOwnerFunctionCalledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "OwnerFunctionCalled", idRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubOwnerFunctionCalledIterator{contract: _SequencerInboxStub.contract, event: "OwnerFunctionCalled", logs: logs, sub: sub}, nil
}

// WatchOwnerFunctionCalled is a free log subscription operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchOwnerFunctionCalled(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubOwnerFunctionCalled, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "OwnerFunctionCalled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubOwnerFunctionCalled)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
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

// ParseOwnerFunctionCalled is a log parse operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseOwnerFunctionCalled(log types.Log) (*SequencerInboxStubOwnerFunctionCalled, error) {
	event := new(SequencerInboxStubOwnerFunctionCalled)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubSequencerBatchDataIterator is returned from FilterSequencerBatchData and is used to iterate over the raw logs and unpacked data for SequencerBatchData events raised by the SequencerInboxStub contract.
type SequencerInboxStubSequencerBatchDataIterator struct {
	Event *SequencerInboxStubSequencerBatchData // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubSequencerBatchDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubSequencerBatchData)
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
		it.Event = new(SequencerInboxStubSequencerBatchData)
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
func (it *SequencerInboxStubSequencerBatchDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubSequencerBatchDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubSequencerBatchData represents a SequencerBatchData event raised by the SequencerInboxStub contract.
type SequencerInboxStubSequencerBatchData struct {
	BatchSequenceNumber *big.Int
	Data                []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchData is a free log retrieval operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterSequencerBatchData(opts *bind.FilterOpts, batchSequenceNumber []*big.Int) (*SequencerInboxStubSequencerBatchDataIterator, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "SequencerBatchData", batchSequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubSequencerBatchDataIterator{contract: _SequencerInboxStub.contract, event: "SequencerBatchData", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchData is a free log subscription operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchSequencerBatchData(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubSequencerBatchData, batchSequenceNumber []*big.Int) (event.Subscription, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "SequencerBatchData", batchSequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubSequencerBatchData)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "SequencerBatchData", log); err != nil {
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

// ParseSequencerBatchData is a log parse operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseSequencerBatchData(log types.Log) (*SequencerInboxStubSequencerBatchData, error) {
	event := new(SequencerInboxStubSequencerBatchData)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "SequencerBatchData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubSequencerBatchDeliveredIterator is returned from FilterSequencerBatchDelivered and is used to iterate over the raw logs and unpacked data for SequencerBatchDelivered events raised by the SequencerInboxStub contract.
type SequencerInboxStubSequencerBatchDeliveredIterator struct {
	Event *SequencerInboxStubSequencerBatchDelivered // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubSequencerBatchDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubSequencerBatchDelivered)
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
		it.Event = new(SequencerInboxStubSequencerBatchDelivered)
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
func (it *SequencerInboxStubSequencerBatchDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubSequencerBatchDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubSequencerBatchDelivered represents a SequencerBatchDelivered event raised by the SequencerInboxStub contract.
type SequencerInboxStubSequencerBatchDelivered struct {
	BatchSequenceNumber      *big.Int
	BeforeAcc                [32]byte
	AfterAcc                 [32]byte
	DelayedAcc               [32]byte
	AfterDelayedMessagesRead *big.Int
	TimeBounds               IBridgeTimeBounds
	DataLocation             uint8
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchDelivered is a free log retrieval operation binding the contract event 0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed batchSequenceNumber, bytes32 indexed beforeAcc, bytes32 indexed afterAcc, bytes32 delayedAcc, uint256 afterDelayedMessagesRead, (uint64,uint64,uint64,uint64) timeBounds, uint8 dataLocation)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterSequencerBatchDelivered(opts *bind.FilterOpts, batchSequenceNumber []*big.Int, beforeAcc [][32]byte, afterAcc [][32]byte) (*SequencerInboxStubSequencerBatchDeliveredIterator, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}
	var afterAccRule []interface{}
	for _, afterAccItem := range afterAcc {
		afterAccRule = append(afterAccRule, afterAccItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "SequencerBatchDelivered", batchSequenceNumberRule, beforeAccRule, afterAccRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubSequencerBatchDeliveredIterator{contract: _SequencerInboxStub.contract, event: "SequencerBatchDelivered", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchDelivered is a free log subscription operation binding the contract event 0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed batchSequenceNumber, bytes32 indexed beforeAcc, bytes32 indexed afterAcc, bytes32 delayedAcc, uint256 afterDelayedMessagesRead, (uint64,uint64,uint64,uint64) timeBounds, uint8 dataLocation)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchSequencerBatchDelivered(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubSequencerBatchDelivered, batchSequenceNumber []*big.Int, beforeAcc [][32]byte, afterAcc [][32]byte) (event.Subscription, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}
	var afterAccRule []interface{}
	for _, afterAccItem := range afterAcc {
		afterAccRule = append(afterAccRule, afterAccItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "SequencerBatchDelivered", batchSequenceNumberRule, beforeAccRule, afterAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubSequencerBatchDelivered)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
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

// ParseSequencerBatchDelivered is a log parse operation binding the contract event 0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed batchSequenceNumber, bytes32 indexed beforeAcc, bytes32 indexed afterAcc, bytes32 delayedAcc, uint256 afterDelayedMessagesRead, (uint64,uint64,uint64,uint64) timeBounds, uint8 dataLocation)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseSequencerBatchDelivered(log types.Log) (*SequencerInboxStubSequencerBatchDelivered, error) {
	event := new(SequencerInboxStubSequencerBatchDelivered)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubSequencerSetIterator is returned from FilterSequencerSet and is used to iterate over the raw logs and unpacked data for SequencerSet events raised by the SequencerInboxStub contract.
type SequencerInboxStubSequencerSetIterator struct {
	Event *SequencerInboxStubSequencerSet // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubSequencerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubSequencerSet)
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
		it.Event = new(SequencerInboxStubSequencerSet)
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
func (it *SequencerInboxStubSequencerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubSequencerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubSequencerSet represents a SequencerSet event raised by the SequencerInboxStub contract.
type SequencerInboxStubSequencerSet struct {
	Addr        common.Address
	IsSequencer bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSequencerSet is a free log retrieval operation binding the contract event 0xeb12a9a53eec138c91b27b4f912a257bd690c18fc8bde744be92a0365eb9b87e.
//
// Solidity: event SequencerSet(address addr, bool isSequencer)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterSequencerSet(opts *bind.FilterOpts) (*SequencerInboxStubSequencerSetIterator, error) {

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "SequencerSet")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubSequencerSetIterator{contract: _SequencerInboxStub.contract, event: "SequencerSet", logs: logs, sub: sub}, nil
}

// WatchSequencerSet is a free log subscription operation binding the contract event 0xeb12a9a53eec138c91b27b4f912a257bd690c18fc8bde744be92a0365eb9b87e.
//
// Solidity: event SequencerSet(address addr, bool isSequencer)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchSequencerSet(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubSequencerSet) (event.Subscription, error) {

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "SequencerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubSequencerSet)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "SequencerSet", log); err != nil {
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

// ParseSequencerSet is a log parse operation binding the contract event 0xeb12a9a53eec138c91b27b4f912a257bd690c18fc8bde744be92a0365eb9b87e.
//
// Solidity: event SequencerSet(address addr, bool isSequencer)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseSequencerSet(log types.Log) (*SequencerInboxStubSequencerSet, error) {
	event := new(SequencerInboxStubSequencerSet)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "SequencerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubSetValidKeysetIterator is returned from FilterSetValidKeyset and is used to iterate over the raw logs and unpacked data for SetValidKeyset events raised by the SequencerInboxStub contract.
type SequencerInboxStubSetValidKeysetIterator struct {
	Event *SequencerInboxStubSetValidKeyset // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubSetValidKeysetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubSetValidKeyset)
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
		it.Event = new(SequencerInboxStubSetValidKeyset)
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
func (it *SequencerInboxStubSetValidKeysetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubSetValidKeysetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubSetValidKeyset represents a SetValidKeyset event raised by the SequencerInboxStub contract.
type SequencerInboxStubSetValidKeyset struct {
	KeysetHash  [32]byte
	KeysetBytes []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetValidKeyset is a free log retrieval operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterSetValidKeyset(opts *bind.FilterOpts, keysetHash [][32]byte) (*SequencerInboxStubSetValidKeysetIterator, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "SetValidKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubSetValidKeysetIterator{contract: _SequencerInboxStub.contract, event: "SetValidKeyset", logs: logs, sub: sub}, nil
}

// WatchSetValidKeyset is a free log subscription operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchSetValidKeyset(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubSetValidKeyset, keysetHash [][32]byte) (event.Subscription, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "SetValidKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubSetValidKeyset)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "SetValidKeyset", log); err != nil {
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

// ParseSetValidKeyset is a log parse operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseSetValidKeyset(log types.Log) (*SequencerInboxStubSetValidKeyset, error) {
	event := new(SequencerInboxStubSetValidKeyset)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "SetValidKeyset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleOneStepProofEntryMetaData contains all meta data concerning the SimpleOneStepProofEntry contract.
var SimpleOneStepProofEntryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"STEPS_PER_BATCH\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState\",\"name\":\"globalState\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus\",\"name\":\"machineStatus\",\"type\":\"uint8\"}],\"internalType\":\"structExecutionState\",\"name\":\"execState\",\"type\":\"tuple\"}],\"name\":\"getMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"}],\"name\":\"getStartMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"step\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveOneStep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610990806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806304997be4146100515780639c2009cd146100cd578063b5112fd2146100ef578063c39619c414610102575b600080fd5b6100ba61005f36600461062e565b6040517f4d616368696e653a0000000000000000000000000000000000000000000000006020820152602881018390526048810182905260009060680160405160208183030381529060405280519060200120905092915050565b6040519081526020015b60405180910390f35b6100d66107d081565b60405167ffffffffffffffff90911681526020016100c4565b6100ba6100fd366004610650565b610115565b6100ba6101103660046106ec565b610439565b600081810361016b5760405162461bcd60e51b815260206004820152600b60248201527f454d5054595f50524f4f4600000000000000000000000000000000000000000060448201526064015b60405180910390fd5b6101736105eb565b600061018085858361055f565b602084015167ffffffffffffffff90921690915290506101a185858361055f565b60208481015167ffffffffffffffff9093169201919091529050861580159061020a57508560001a60f81b7fff0000000000000000000000000000000000000000000000000000000000000016158061020a57506101fe826105c7565b67ffffffffffffffff16155b15610219578592505050610430565b8735610224836105dd565b67ffffffffffffffff161061023d578592505050610430565b8151805160209182015182850151805190840151604080517f476c6f62616c2073746174653a0000000000000000000000000000000000000081880152602d810195909552604d8501939093527fffffffffffffffff00000000000000000000000000000000000000000000000060c092831b8116606d860152911b1660758301528051808303605d018152607d9092019052805191012086146103235760405162461bcd60e51b815260206004820152600960248201527f4241445f50524f4f4600000000000000000000000000000000000000000000006044820152606401610162565b6020828101510180519061033682610730565b67ffffffffffffffff169052506020828101510151610358906107d090610757565b67ffffffffffffffff1660000361039357602082015180519061037a82610730565b67ffffffffffffffff1690525060208281015160009101525b8151805160209182015182850151805190840151604080517f476c6f62616c2073746174653a0000000000000000000000000000000000000081880152602d810195909552604d8501939093527fffffffffffffffff00000000000000000000000000000000000000000000000060c092831b8116606d860152911b1660758301528051808303605d018152607d90920190528051910120925050505b95945050505050565b6000600161044d60a08401608085016107a2565b600281111561045e5761045e61078c565b146104ab5760405162461bcd60e51b815260206004820152601260248201527f4241445f4d414348494e455f53544154555300000000000000000000000000006044820152606401610162565b6105596104bd36849003840184610889565b8051805160209182015192820151805190830151604080517f476c6f62616c2073746174653a0000000000000000000000000000000000000081870152602d810194909452604d8401959095527fffffffffffffffff00000000000000000000000000000000000000000000000060c092831b8116606d850152911b1660758201528251808203605d018152607d909101909252815191012090565b92915050565b600081815b60088110156105be5760088367ffffffffffffffff16901b925085858381811061059057610590610704565b919091013560f81c939093179250816105a881610922565b92505080806105b690610922565b915050610564565b50935093915050565b602081015160009060015b602002015192915050565b6020810151600090816105d2565b60405180604001604052806105fe610610565b815260200161060b610610565b905290565b60405180604001604052806002906020820280368337509192915050565b6000806040838503121561064157600080fd5b50508035926020909101359150565b600080600080600085870360c081121561066957600080fd5b606081121561067757600080fd5b50859450606086013593506080860135925060a086013567ffffffffffffffff808211156106a457600080fd5b818801915088601f8301126106b857600080fd5b8135818111156106c757600080fd5b8960208285010111156106d957600080fd5b9699959850939650602001949392505050565b600060a082840312156106fe57600080fd5b50919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600067ffffffffffffffff80831681810361074d5761074d61071a565b6001019392505050565b600067ffffffffffffffff8084168061078057634e487b7160e01b600052601260045260246000fd5b92169190910692915050565b634e487b7160e01b600052602160045260246000fd5b6000602082840312156107b457600080fd5b8135600381106107c357600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610803576108036107ca565b60405290565b600082601f83011261081a57600080fd5b6040516040810167ffffffffffffffff828210818311171561083e5761083e6107ca565b6040918252829185018681111561085457600080fd5b855b8181101561087d578035838116811461086f5760008081fd5b845260209384019301610856565b50929695505050505050565b60006080828403121561089b57600080fd5b6040516040810181811067ffffffffffffffff821117156108be576108be6107ca565b604052601f830184136108d057600080fd5b6108d86107e0565b8060408501868111156108ea57600080fd5b855b818110156109045780358452602093840193016108ec565b508184526109128782610809565b6020850152509195945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036109535761095361071a565b506001019056fea26469706673582212201a2f420f23788e34a9ef955580ce85abc37eaf9845add9e5fed8b919e2e0533064736f6c63430008110033",
}

// SimpleOneStepProofEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleOneStepProofEntryMetaData.ABI instead.
var SimpleOneStepProofEntryABI = SimpleOneStepProofEntryMetaData.ABI

// SimpleOneStepProofEntryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleOneStepProofEntryMetaData.Bin instead.
var SimpleOneStepProofEntryBin = SimpleOneStepProofEntryMetaData.Bin

// DeploySimpleOneStepProofEntry deploys a new Ethereum contract, binding an instance of SimpleOneStepProofEntry to it.
func DeploySimpleOneStepProofEntry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleOneStepProofEntry, error) {
	parsed, err := SimpleOneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleOneStepProofEntryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleOneStepProofEntry{SimpleOneStepProofEntryCaller: SimpleOneStepProofEntryCaller{contract: contract}, SimpleOneStepProofEntryTransactor: SimpleOneStepProofEntryTransactor{contract: contract}, SimpleOneStepProofEntryFilterer: SimpleOneStepProofEntryFilterer{contract: contract}}, nil
}

// SimpleOneStepProofEntry is an auto generated Go binding around an Ethereum contract.
type SimpleOneStepProofEntry struct {
	SimpleOneStepProofEntryCaller     // Read-only binding to the contract
	SimpleOneStepProofEntryTransactor // Write-only binding to the contract
	SimpleOneStepProofEntryFilterer   // Log filterer for contract events
}

// SimpleOneStepProofEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleOneStepProofEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleOneStepProofEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleOneStepProofEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleOneStepProofEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleOneStepProofEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleOneStepProofEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleOneStepProofEntrySession struct {
	Contract     *SimpleOneStepProofEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleOneStepProofEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleOneStepProofEntryCallerSession struct {
	Contract *SimpleOneStepProofEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// SimpleOneStepProofEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleOneStepProofEntryTransactorSession struct {
	Contract     *SimpleOneStepProofEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// SimpleOneStepProofEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleOneStepProofEntryRaw struct {
	Contract *SimpleOneStepProofEntry // Generic contract binding to access the raw methods on
}

// SimpleOneStepProofEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleOneStepProofEntryCallerRaw struct {
	Contract *SimpleOneStepProofEntryCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleOneStepProofEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleOneStepProofEntryTransactorRaw struct {
	Contract *SimpleOneStepProofEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleOneStepProofEntry creates a new instance of SimpleOneStepProofEntry, bound to a specific deployed contract.
func NewSimpleOneStepProofEntry(address common.Address, backend bind.ContractBackend) (*SimpleOneStepProofEntry, error) {
	contract, err := bindSimpleOneStepProofEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleOneStepProofEntry{SimpleOneStepProofEntryCaller: SimpleOneStepProofEntryCaller{contract: contract}, SimpleOneStepProofEntryTransactor: SimpleOneStepProofEntryTransactor{contract: contract}, SimpleOneStepProofEntryFilterer: SimpleOneStepProofEntryFilterer{contract: contract}}, nil
}

// NewSimpleOneStepProofEntryCaller creates a new read-only instance of SimpleOneStepProofEntry, bound to a specific deployed contract.
func NewSimpleOneStepProofEntryCaller(address common.Address, caller bind.ContractCaller) (*SimpleOneStepProofEntryCaller, error) {
	contract, err := bindSimpleOneStepProofEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleOneStepProofEntryCaller{contract: contract}, nil
}

// NewSimpleOneStepProofEntryTransactor creates a new write-only instance of SimpleOneStepProofEntry, bound to a specific deployed contract.
func NewSimpleOneStepProofEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleOneStepProofEntryTransactor, error) {
	contract, err := bindSimpleOneStepProofEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleOneStepProofEntryTransactor{contract: contract}, nil
}

// NewSimpleOneStepProofEntryFilterer creates a new log filterer instance of SimpleOneStepProofEntry, bound to a specific deployed contract.
func NewSimpleOneStepProofEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleOneStepProofEntryFilterer, error) {
	contract, err := bindSimpleOneStepProofEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleOneStepProofEntryFilterer{contract: contract}, nil
}

// bindSimpleOneStepProofEntry binds a generic wrapper to an already deployed contract.
func bindSimpleOneStepProofEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleOneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleOneStepProofEntry.Contract.SimpleOneStepProofEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleOneStepProofEntry.Contract.SimpleOneStepProofEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleOneStepProofEntry.Contract.SimpleOneStepProofEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleOneStepProofEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleOneStepProofEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleOneStepProofEntry.Contract.contract.Transact(opts, method, params...)
}

// STEPSPERBATCH is a free data retrieval call binding the contract method 0x9c2009cd.
//
// Solidity: function STEPS_PER_BATCH() view returns(uint64)
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntryCaller) STEPSPERBATCH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SimpleOneStepProofEntry.contract.Call(opts, &out, "STEPS_PER_BATCH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// STEPSPERBATCH is a free data retrieval call binding the contract method 0x9c2009cd.
//
// Solidity: function STEPS_PER_BATCH() view returns(uint64)
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntrySession) STEPSPERBATCH() (uint64, error) {
	return _SimpleOneStepProofEntry.Contract.STEPSPERBATCH(&_SimpleOneStepProofEntry.CallOpts)
}

// STEPSPERBATCH is a free data retrieval call binding the contract method 0x9c2009cd.
//
// Solidity: function STEPS_PER_BATCH() view returns(uint64)
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntryCallerSession) STEPSPERBATCH() (uint64, error) {
	return _SimpleOneStepProofEntry.Contract.STEPSPERBATCH(&_SimpleOneStepProofEntry.CallOpts)
}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntryCaller) GetMachineHash(opts *bind.CallOpts, execState ExecutionState) ([32]byte, error) {
	var out []interface{}
	err := _SimpleOneStepProofEntry.contract.Call(opts, &out, "getMachineHash", execState)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntrySession) GetMachineHash(execState ExecutionState) ([32]byte, error) {
	return _SimpleOneStepProofEntry.Contract.GetMachineHash(&_SimpleOneStepProofEntry.CallOpts, execState)
}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntryCallerSession) GetMachineHash(execState ExecutionState) ([32]byte, error) {
	return _SimpleOneStepProofEntry.Contract.GetMachineHash(&_SimpleOneStepProofEntry.CallOpts, execState)
}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntryCaller) GetStartMachineHash(opts *bind.CallOpts, globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SimpleOneStepProofEntry.contract.Call(opts, &out, "getStartMachineHash", globalStateHash, wasmModuleRoot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntrySession) GetStartMachineHash(globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	return _SimpleOneStepProofEntry.Contract.GetStartMachineHash(&_SimpleOneStepProofEntry.CallOpts, globalStateHash, wasmModuleRoot)
}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntryCallerSession) GetStartMachineHash(globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	return _SimpleOneStepProofEntry.Contract.GetStartMachineHash(&_SimpleOneStepProofEntry.CallOpts, globalStateHash, wasmModuleRoot)
}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 step, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntryCaller) ProveOneStep(opts *bind.CallOpts, execCtx ExecutionContext, step *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	var out []interface{}
	err := _SimpleOneStepProofEntry.contract.Call(opts, &out, "proveOneStep", execCtx, step, beforeHash, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 step, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntrySession) ProveOneStep(execCtx ExecutionContext, step *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _SimpleOneStepProofEntry.Contract.ProveOneStep(&_SimpleOneStepProofEntry.CallOpts, execCtx, step, beforeHash, proof)
}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 step, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_SimpleOneStepProofEntry *SimpleOneStepProofEntryCallerSession) ProveOneStep(execCtx ExecutionContext, step *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _SimpleOneStepProofEntry.Contract.ProveOneStep(&_SimpleOneStepProofEntry.CallOpts, execCtx, step, beforeHash, proof)
}

// TestWETH9MetaData contains all meta data concerning the TestWETH9 contract.
var TestWETH9MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200107838038062001078833981016040819052620000349162000123565b818160036200004483826200021c565b5060046200005382826200021c565b5050505050620002e8565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200008657600080fd5b81516001600160401b0380821115620000a357620000a36200005e565b604051601f8301601f19908116603f01168101908282118183101715620000ce57620000ce6200005e565b81604052838152602092508683858801011115620000eb57600080fd5b600091505b838210156200010f5785820183015181830184015290820190620000f0565b600093810190920192909252949350505050565b600080604083850312156200013757600080fd5b82516001600160401b03808211156200014f57600080fd5b6200015d8683870162000074565b935060208501519150808211156200017457600080fd5b50620001838582860162000074565b9150509250929050565b600181811c90821680620001a257607f821691505b602082108103620001c357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200021757600081815260208120601f850160051c81016020861015620001f25750805b601f850160051c820191505b818110156200021357828155600101620001fe565b5050505b505050565b81516001600160401b038111156200023857620002386200005e565b62000250816200024984546200018d565b84620001c9565b602080601f8311600181146200028857600084156200026f5750858301515b600019600386901b1c1916600185901b17855562000213565b600085815260208120601f198616915b82811015620002b95788860151825594840194600190910190840162000298565b5085821015620002d85787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610d8080620002f86000396000f3fe6080604052600436106100d25760003560e01c8063395093511161007f578063a457c2d711610059578063a457c2d71461021a578063a9059cbb1461023a578063d0e30db01461025a578063dd62ed3e1461026257600080fd5b806339509351146101af57806370a08231146101cf57806395d89b411461020557600080fd5b806323b872dd116100b057806323b872dd146101515780632e1a7d4d14610171578063313ce5671461019357600080fd5b806306fdde03146100d7578063095ea7b31461010257806318160ddd14610132575b600080fd5b3480156100e357600080fd5b506100ec6102a8565b6040516100f99190610b46565b60405180910390f35b34801561010e57600080fd5b5061012261011d366004610bce565b61033a565b60405190151581526020016100f9565b34801561013e57600080fd5b506002545b6040519081526020016100f9565b34801561015d57600080fd5b5061012261016c366004610bf8565b610354565b34801561017d57600080fd5b5061019161018c366004610c34565b610378565b005b34801561019f57600080fd5b50604051601281526020016100f9565b3480156101bb57600080fd5b506101226101ca366004610bce565b6103b3565b3480156101db57600080fd5b506101436101ea366004610c4d565b6001600160a01b031660009081526020819052604090205490565b34801561021157600080fd5b506100ec6103f2565b34801561022657600080fd5b50610122610235366004610bce565b610401565b34801561024657600080fd5b50610122610255366004610bce565b6104b0565b6101916104be565b34801561026e57600080fd5b5061014361027d366004610c6f565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6060600380546102b790610ca2565b80601f01602080910402602001604051908101604052809291908181526020018280546102e390610ca2565b80156103305780601f1061030557610100808354040283529160200191610330565b820191906000526020600020905b81548152906001019060200180831161031357829003601f168201915b5050505050905090565b6000336103488185856104ca565b60019150505b92915050565b600033610362858285610623565b61036d8585856106d3565b506001949350505050565b61038233826108ea565b604051339082156108fc029083906000818181858888f193505050501580156103af573d6000803e3d6000fd5b5050565b3360008181526001602090815260408083206001600160a01b038716845290915281205490919061034890829086906103ed908790610d24565b6104ca565b6060600480546102b790610ca2565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909190838110156104a35760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b61036d82868684036104ca565b6000336103488185856106d3565b6104c83334610a67565b565b6001600160a01b0383166105455760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f7265737300000000000000000000000000000000000000000000000000000000606482015260840161049a565b6001600160a01b0382166105c15760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f7373000000000000000000000000000000000000000000000000000000000000606482015260840161049a565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600160a01b038381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146106cd57818110156106c05760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161049a565b6106cd84848484036104ca565b50505050565b6001600160a01b03831661074f5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015260840161049a565b6001600160a01b0382166107cb5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015260840161049a565b6001600160a01b0383166000908152602081905260409020548181101561085a5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e63650000000000000000000000000000000000000000000000000000606482015260840161049a565b6001600160a01b03808516600090815260208190526040808220858503905591851681529081208054849290610891908490610d24565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516108dd91815260200190565b60405180910390a36106cd565b6001600160a01b0382166109665760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161049a565b6001600160a01b038216600090815260208190526040902054818110156109f55760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f6365000000000000000000000000000000000000000000000000000000000000606482015260840161049a565b6001600160a01b0383166000908152602081905260408120838303905560028054849290610a24908490610d37565b90915550506040518281526000906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610616565b6001600160a01b038216610abd5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161049a565b8060026000828254610acf9190610d24565b90915550506001600160a01b03821660009081526020819052604081208054839290610afc908490610d24565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b600060208083528351808285015260005b81811015610b7357858101830151858201604001528201610b57565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b80356001600160a01b0381168114610bc957600080fd5b919050565b60008060408385031215610be157600080fd5b610bea83610bb2565b946020939093013593505050565b600080600060608486031215610c0d57600080fd5b610c1684610bb2565b9250610c2460208501610bb2565b9150604084013590509250925092565b600060208284031215610c4657600080fd5b5035919050565b600060208284031215610c5f57600080fd5b610c6882610bb2565b9392505050565b60008060408385031215610c8257600080fd5b610c8b83610bb2565b9150610c9960208401610bb2565b90509250929050565b600181811c90821680610cb657607f821691505b602082108103610cef577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561034e5761034e610cf5565b8181038181111561034e5761034e610cf556fea2646970667358221220d8023966023a2784e1f1ca33411e5b774cae5e19a008cb86925c21a2a2fa3b0164736f6c63430008110033",
}

// TestWETH9ABI is the input ABI used to generate the binding from.
// Deprecated: Use TestWETH9MetaData.ABI instead.
var TestWETH9ABI = TestWETH9MetaData.ABI

// TestWETH9Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestWETH9MetaData.Bin instead.
var TestWETH9Bin = TestWETH9MetaData.Bin

// DeployTestWETH9 deploys a new Ethereum contract, binding an instance of TestWETH9 to it.
func DeployTestWETH9(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *TestWETH9, error) {
	parsed, err := TestWETH9MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestWETH9Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestWETH9{TestWETH9Caller: TestWETH9Caller{contract: contract}, TestWETH9Transactor: TestWETH9Transactor{contract: contract}, TestWETH9Filterer: TestWETH9Filterer{contract: contract}}, nil
}

// TestWETH9 is an auto generated Go binding around an Ethereum contract.
type TestWETH9 struct {
	TestWETH9Caller     // Read-only binding to the contract
	TestWETH9Transactor // Write-only binding to the contract
	TestWETH9Filterer   // Log filterer for contract events
}

// TestWETH9Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestWETH9Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestWETH9Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestWETH9Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestWETH9Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestWETH9Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestWETH9Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestWETH9Session struct {
	Contract     *TestWETH9        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestWETH9CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestWETH9CallerSession struct {
	Contract *TestWETH9Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TestWETH9TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestWETH9TransactorSession struct {
	Contract     *TestWETH9Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TestWETH9Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestWETH9Raw struct {
	Contract *TestWETH9 // Generic contract binding to access the raw methods on
}

// TestWETH9CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestWETH9CallerRaw struct {
	Contract *TestWETH9Caller // Generic read-only contract binding to access the raw methods on
}

// TestWETH9TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestWETH9TransactorRaw struct {
	Contract *TestWETH9Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestWETH9 creates a new instance of TestWETH9, bound to a specific deployed contract.
func NewTestWETH9(address common.Address, backend bind.ContractBackend) (*TestWETH9, error) {
	contract, err := bindTestWETH9(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestWETH9{TestWETH9Caller: TestWETH9Caller{contract: contract}, TestWETH9Transactor: TestWETH9Transactor{contract: contract}, TestWETH9Filterer: TestWETH9Filterer{contract: contract}}, nil
}

// NewTestWETH9Caller creates a new read-only instance of TestWETH9, bound to a specific deployed contract.
func NewTestWETH9Caller(address common.Address, caller bind.ContractCaller) (*TestWETH9Caller, error) {
	contract, err := bindTestWETH9(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestWETH9Caller{contract: contract}, nil
}

// NewTestWETH9Transactor creates a new write-only instance of TestWETH9, bound to a specific deployed contract.
func NewTestWETH9Transactor(address common.Address, transactor bind.ContractTransactor) (*TestWETH9Transactor, error) {
	contract, err := bindTestWETH9(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestWETH9Transactor{contract: contract}, nil
}

// NewTestWETH9Filterer creates a new log filterer instance of TestWETH9, bound to a specific deployed contract.
func NewTestWETH9Filterer(address common.Address, filterer bind.ContractFilterer) (*TestWETH9Filterer, error) {
	contract, err := bindTestWETH9(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestWETH9Filterer{contract: contract}, nil
}

// bindTestWETH9 binds a generic wrapper to an already deployed contract.
func bindTestWETH9(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestWETH9MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestWETH9 *TestWETH9Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestWETH9.Contract.TestWETH9Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestWETH9 *TestWETH9Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestWETH9.Contract.TestWETH9Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestWETH9 *TestWETH9Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestWETH9.Contract.TestWETH9Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestWETH9 *TestWETH9CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestWETH9.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestWETH9 *TestWETH9TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestWETH9.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestWETH9 *TestWETH9TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestWETH9.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestWETH9 *TestWETH9Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestWETH9.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestWETH9 *TestWETH9Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestWETH9.Contract.Allowance(&_TestWETH9.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestWETH9 *TestWETH9CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestWETH9.Contract.Allowance(&_TestWETH9.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestWETH9 *TestWETH9Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestWETH9.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestWETH9 *TestWETH9Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestWETH9.Contract.BalanceOf(&_TestWETH9.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestWETH9 *TestWETH9CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestWETH9.Contract.BalanceOf(&_TestWETH9.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestWETH9 *TestWETH9Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestWETH9.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestWETH9 *TestWETH9Session) Decimals() (uint8, error) {
	return _TestWETH9.Contract.Decimals(&_TestWETH9.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestWETH9 *TestWETH9CallerSession) Decimals() (uint8, error) {
	return _TestWETH9.Contract.Decimals(&_TestWETH9.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestWETH9 *TestWETH9Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestWETH9.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestWETH9 *TestWETH9Session) Name() (string, error) {
	return _TestWETH9.Contract.Name(&_TestWETH9.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestWETH9 *TestWETH9CallerSession) Name() (string, error) {
	return _TestWETH9.Contract.Name(&_TestWETH9.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestWETH9 *TestWETH9Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestWETH9.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestWETH9 *TestWETH9Session) Symbol() (string, error) {
	return _TestWETH9.Contract.Symbol(&_TestWETH9.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestWETH9 *TestWETH9CallerSession) Symbol() (string, error) {
	return _TestWETH9.Contract.Symbol(&_TestWETH9.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestWETH9 *TestWETH9Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestWETH9.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestWETH9 *TestWETH9Session) TotalSupply() (*big.Int, error) {
	return _TestWETH9.Contract.TotalSupply(&_TestWETH9.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestWETH9 *TestWETH9CallerSession) TotalSupply() (*big.Int, error) {
	return _TestWETH9.Contract.TotalSupply(&_TestWETH9.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestWETH9 *TestWETH9Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestWETH9.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestWETH9 *TestWETH9Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestWETH9.Contract.Approve(&_TestWETH9.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestWETH9 *TestWETH9TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestWETH9.Contract.Approve(&_TestWETH9.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestWETH9 *TestWETH9Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestWETH9.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestWETH9 *TestWETH9Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestWETH9.Contract.DecreaseAllowance(&_TestWETH9.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestWETH9 *TestWETH9TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestWETH9.Contract.DecreaseAllowance(&_TestWETH9.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestWETH9 *TestWETH9Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestWETH9.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestWETH9 *TestWETH9Session) Deposit() (*types.Transaction, error) {
	return _TestWETH9.Contract.Deposit(&_TestWETH9.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_TestWETH9 *TestWETH9TransactorSession) Deposit() (*types.Transaction, error) {
	return _TestWETH9.Contract.Deposit(&_TestWETH9.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestWETH9 *TestWETH9Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestWETH9.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestWETH9 *TestWETH9Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestWETH9.Contract.IncreaseAllowance(&_TestWETH9.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestWETH9 *TestWETH9TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestWETH9.Contract.IncreaseAllowance(&_TestWETH9.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestWETH9 *TestWETH9Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestWETH9.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestWETH9 *TestWETH9Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestWETH9.Contract.Transfer(&_TestWETH9.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestWETH9 *TestWETH9TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestWETH9.Contract.Transfer(&_TestWETH9.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestWETH9 *TestWETH9Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestWETH9.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestWETH9 *TestWETH9Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestWETH9.Contract.TransferFrom(&_TestWETH9.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestWETH9 *TestWETH9TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestWETH9.Contract.TransferFrom(&_TestWETH9.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_TestWETH9 *TestWETH9Transactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _TestWETH9.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_TestWETH9 *TestWETH9Session) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _TestWETH9.Contract.Withdraw(&_TestWETH9.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_TestWETH9 *TestWETH9TransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _TestWETH9.Contract.Withdraw(&_TestWETH9.TransactOpts, _amount)
}

// TestWETH9ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TestWETH9 contract.
type TestWETH9ApprovalIterator struct {
	Event *TestWETH9Approval // Event containing the contract specifics and raw log

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
func (it *TestWETH9ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestWETH9Approval)
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
		it.Event = new(TestWETH9Approval)
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
func (it *TestWETH9ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestWETH9ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestWETH9Approval represents a Approval event raised by the TestWETH9 contract.
type TestWETH9Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestWETH9 *TestWETH9Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TestWETH9ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestWETH9.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TestWETH9ApprovalIterator{contract: _TestWETH9.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestWETH9 *TestWETH9Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestWETH9Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestWETH9.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestWETH9Approval)
				if err := _TestWETH9.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestWETH9 *TestWETH9Filterer) ParseApproval(log types.Log) (*TestWETH9Approval, error) {
	event := new(TestWETH9Approval)
	if err := _TestWETH9.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestWETH9TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TestWETH9 contract.
type TestWETH9TransferIterator struct {
	Event *TestWETH9Transfer // Event containing the contract specifics and raw log

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
func (it *TestWETH9TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestWETH9Transfer)
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
		it.Event = new(TestWETH9Transfer)
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
func (it *TestWETH9TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestWETH9TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestWETH9Transfer represents a Transfer event raised by the TestWETH9 contract.
type TestWETH9Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestWETH9 *TestWETH9Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestWETH9TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestWETH9.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestWETH9TransferIterator{contract: _TestWETH9.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestWETH9 *TestWETH9Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestWETH9Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestWETH9.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestWETH9Transfer)
				if err := _TestWETH9.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestWETH9 *TestWETH9Filterer) ParseTransfer(log types.Log) (*TestWETH9Transfer, error) {
	event := new(TestWETH9Transfer)
	if err := _TestWETH9.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorMockMetaData contains all meta data concerning the UpgradeExecutorMock contract.
var UpgradeExecutorMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TargetCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"upgrade\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"UpgradeExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXECUTOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"upgrade\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"upgradeCallData\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetCallData\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"executors\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506001609755600054610100900460ff16158080156100365750600054600160ff909116105b80610061575061004f3061013760201b6108881760201c565b158015610061575060005460ff166001145b6100c85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b6000805460ff1916600117905580156100eb576000805461ff0019166101001790555b8015610131576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50610146565b6001600160a01b03163b151590565b61145680620001566000396000f3fe6080604052600436106100c75760003560e01c806375b238fc11610074578063a217fddf1161004e578063a217fddf14610262578063bca8c7b514610277578063d547741f1461028a57600080fd5b806375b238fc146101c857806391d14854146101fc578063946d92041461024257600080fd5b8063248a9ca3116100a5578063248a9ca3146101585780632f2ff15d1461018857806336568abe146101a857600080fd5b806301ffc9a7146100cc57806307bd0265146101015780631cff79cd14610143575b600080fd5b3480156100d857600080fd5b506100ec6100e7366004610fbd565b6102aa565b60405190151581526020015b60405180910390f35b34801561010d57600080fd5b506101357fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6381565b6040519081526020016100f8565b610156610151366004611062565b610343565b005b34801561016457600080fd5b50610135610173366004611108565b60009081526065602052604090206001015490565b34801561019457600080fd5b506101566101a3366004611121565b610448565b3480156101b457600080fd5b506101566101c3366004611121565b610472565b3480156101d457600080fd5b506101357fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b34801561020857600080fd5b506100ec610217366004611121565b60009182526065602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561024e57600080fd5b5061015661025d36600461114d565b6104fe565b34801561026e57600080fd5b50610135600081565b610156610285366004611062565b610773565b34801561029657600080fd5b506101566102a5366004611121565b610863565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061033d57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b7fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6361036d81610897565b6002609754036103c45760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064015b60405180910390fd5b60026097819055506103fa826040518060600160405280603a81526020016113e7603a91396001600160a01b03861691906108a4565b50826001600160a01b03167f49f6851d1cd01a518db5bdea5cffbbe90276baa2595f74250b7472b96806302e348460405161043692919061125d565b60405180910390a25050600160975550565b60008281526065602052604090206001015461046381610897565b61046d838361099a565b505050565b6001600160a01b03811633146104f05760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201527f20726f6c657320666f722073656c66000000000000000000000000000000000060648201526084016103bb565b6104fa8282610a3c565b5050565b600054610100900460ff161580801561051e5750600054600160ff909116105b806105385750303b158015610538575060005460ff166001145b6105aa5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016103bb565b6000805460ff1916600117905580156105cd576000805461ff0019166101001790555b6001600160a01b0383166106235760405162461bcd60e51b815260206004820152601b60248201527f557067726164654578656375746f723a207a65726f2061646d696e000000000060448201526064016103bb565b61062b610abf565b6106557fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177580610b3e565b61069f7fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e637fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775610b3e565b6106c97fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177584610b89565b60005b8251811015610728576107187fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6384838151811061070b5761070b61127e565b6020026020010151610b89565b610721816112aa565b90506106cc565b50801561046d576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b7fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6361079d81610897565b6002609754036107ef5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016103bb565b600260978190555061082782346040518060600160405280603181526020016113b6603191396001600160a01b038716929190610b93565b50826001600160a01b03167f4d7dbdcc249630ec373f584267f10abf44938de920c32562f5aee93959c25258348460405161043692919061125d565b60008281526065602052604090206001015461087e81610897565b61046d8383610a3c565b6001600160a01b03163b151590565b6108a18133610cdb565b50565b60606001600160a01b0384163b6109235760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60448201527f6e7472616374000000000000000000000000000000000000000000000000000060648201526084016103bb565b600080856001600160a01b03168560405161093e91906112c4565b600060405180830381855af49150503d8060008114610979576040519150601f19603f3d011682016040523d82523d6000602084013e61097e565b606091505b509150915061098e828286610d5b565b925050505b9392505050565b60008281526065602090815260408083206001600160a01b038516845290915290205460ff166104fa5760008281526065602090815260408083206001600160a01b03851684529091529020805460ff191660011790556109f83390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60008281526065602090815260408083206001600160a01b038516845290915290205460ff16156104fa5760008281526065602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b600054610100900460ff16610b3c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103bb565b565b600082815260656020526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b6104fa828261099a565b606082471015610c0b5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016103bb565b6001600160a01b0385163b610c625760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103bb565b600080866001600160a01b03168587604051610c7e91906112c4565b60006040518083038185875af1925050503d8060008114610cbb576040519150601f19603f3d011682016040523d82523d6000602084013e610cc0565b606091505b5091509150610cd0828286610d5b565b979650505050505050565b60008281526065602090815260408083206001600160a01b038516845290915290205460ff166104fa57610d19816001600160a01b03166014610d94565b610d24836020610d94565b604051602001610d359291906112e0565b60408051601f198184030181529082905262461bcd60e51b82526103bb91600401611361565b60608315610d6a575081610993565b825115610d7a5782518084602001fd5b8160405162461bcd60e51b81526004016103bb9190611361565b60606000610da3836002611374565b610dae90600261138b565b67ffffffffffffffff811115610dc657610dc661101b565b6040519080825280601f01601f191660200182016040528015610df0576020820181803683370190505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110610e2757610e2761127e565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610e8a57610e8a61127e565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506000610ec6846002611374565b610ed190600161138b565b90505b6001811115610f6e577f303132333435363738396162636465660000000000000000000000000000000085600f1660108110610f1257610f1261127e565b1a60f81b828281518110610f2857610f2861127e565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c93610f678161139e565b9050610ed4565b5083156109935760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016103bb565b600060208284031215610fcf57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461099357600080fd5b80356001600160a01b038116811461101657600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561105a5761105a61101b565b604052919050565b6000806040838503121561107557600080fd5b61107e83610fff565b915060208084013567ffffffffffffffff8082111561109c57600080fd5b818601915086601f8301126110b057600080fd5b8135818111156110c2576110c261101b565b6110d484601f19601f84011601611031565b915080825287848285010111156110ea57600080fd5b80848401858401376000848284010152508093505050509250929050565b60006020828403121561111a57600080fd5b5035919050565b6000806040838503121561113457600080fd5b8235915061114460208401610fff565b90509250929050565b6000806040838503121561116057600080fd5b61116983610fff565b915060208084013567ffffffffffffffff8082111561118757600080fd5b818601915086601f83011261119b57600080fd5b8135818111156111ad576111ad61101b565b8060051b91506111be848301611031565b81815291830184019184810190898411156111d857600080fd5b938501935b838510156111fd576111ee85610fff565b825293850193908501906111dd565b8096505050505050509250929050565b60005b83811015611228578181015183820152602001611210565b50506000910152565b6000815180845261124981602086016020860161120d565b601f01601f19169290920160200192915050565b8281526040602082015260006112766040830184611231565b949350505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060001982036112bd576112bd611294565b5060010190565b600082516112d681846020870161120d565b9190910192915050565b7f416363657373436f6e74726f6c3a206163636f756e742000000000000000000081526000835161131881601785016020880161120d565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000601791840191820152835161135581602884016020880161120d565b01602801949350505050565b6020815260006109936020830184611231565b808202811582820484141761033d5761033d611294565b8082018082111561033d5761033d611294565b6000816113ad576113ad611294565b50600019019056fe557067726164654578656375746f723a20696e6e65722063616c6c206661696c656420776974686f757420726561736f6e557067726164654578656375746f723a20696e6e65722064656c65676174652063616c6c206661696c656420776974686f757420726561736f6ea2646970667358221220f44ae3705db75af49201a3b4f4ee0652cd1d9d55dbae6e20df4e452f611744e464736f6c63430008110033",
}

// UpgradeExecutorMockABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradeExecutorMockMetaData.ABI instead.
var UpgradeExecutorMockABI = UpgradeExecutorMockMetaData.ABI

// UpgradeExecutorMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UpgradeExecutorMockMetaData.Bin instead.
var UpgradeExecutorMockBin = UpgradeExecutorMockMetaData.Bin

// DeployUpgradeExecutorMock deploys a new Ethereum contract, binding an instance of UpgradeExecutorMock to it.
func DeployUpgradeExecutorMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UpgradeExecutorMock, error) {
	parsed, err := UpgradeExecutorMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UpgradeExecutorMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UpgradeExecutorMock{UpgradeExecutorMockCaller: UpgradeExecutorMockCaller{contract: contract}, UpgradeExecutorMockTransactor: UpgradeExecutorMockTransactor{contract: contract}, UpgradeExecutorMockFilterer: UpgradeExecutorMockFilterer{contract: contract}}, nil
}

// UpgradeExecutorMock is an auto generated Go binding around an Ethereum contract.
type UpgradeExecutorMock struct {
	UpgradeExecutorMockCaller     // Read-only binding to the contract
	UpgradeExecutorMockTransactor // Write-only binding to the contract
	UpgradeExecutorMockFilterer   // Log filterer for contract events
}

// UpgradeExecutorMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeExecutorMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeExecutorMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeExecutorMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeExecutorMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeExecutorMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeExecutorMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeExecutorMockSession struct {
	Contract     *UpgradeExecutorMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UpgradeExecutorMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeExecutorMockCallerSession struct {
	Contract *UpgradeExecutorMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// UpgradeExecutorMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeExecutorMockTransactorSession struct {
	Contract     *UpgradeExecutorMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// UpgradeExecutorMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeExecutorMockRaw struct {
	Contract *UpgradeExecutorMock // Generic contract binding to access the raw methods on
}

// UpgradeExecutorMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeExecutorMockCallerRaw struct {
	Contract *UpgradeExecutorMockCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeExecutorMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeExecutorMockTransactorRaw struct {
	Contract *UpgradeExecutorMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeExecutorMock creates a new instance of UpgradeExecutorMock, bound to a specific deployed contract.
func NewUpgradeExecutorMock(address common.Address, backend bind.ContractBackend) (*UpgradeExecutorMock, error) {
	contract, err := bindUpgradeExecutorMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorMock{UpgradeExecutorMockCaller: UpgradeExecutorMockCaller{contract: contract}, UpgradeExecutorMockTransactor: UpgradeExecutorMockTransactor{contract: contract}, UpgradeExecutorMockFilterer: UpgradeExecutorMockFilterer{contract: contract}}, nil
}

// NewUpgradeExecutorMockCaller creates a new read-only instance of UpgradeExecutorMock, bound to a specific deployed contract.
func NewUpgradeExecutorMockCaller(address common.Address, caller bind.ContractCaller) (*UpgradeExecutorMockCaller, error) {
	contract, err := bindUpgradeExecutorMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorMockCaller{contract: contract}, nil
}

// NewUpgradeExecutorMockTransactor creates a new write-only instance of UpgradeExecutorMock, bound to a specific deployed contract.
func NewUpgradeExecutorMockTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeExecutorMockTransactor, error) {
	contract, err := bindUpgradeExecutorMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorMockTransactor{contract: contract}, nil
}

// NewUpgradeExecutorMockFilterer creates a new log filterer instance of UpgradeExecutorMock, bound to a specific deployed contract.
func NewUpgradeExecutorMockFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeExecutorMockFilterer, error) {
	contract, err := bindUpgradeExecutorMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorMockFilterer{contract: contract}, nil
}

// bindUpgradeExecutorMock binds a generic wrapper to an already deployed contract.
func bindUpgradeExecutorMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpgradeExecutorMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeExecutorMock *UpgradeExecutorMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeExecutorMock.Contract.UpgradeExecutorMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeExecutorMock *UpgradeExecutorMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.UpgradeExecutorMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeExecutorMock *UpgradeExecutorMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.UpgradeExecutorMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeExecutorMock *UpgradeExecutorMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeExecutorMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeExecutorMock *UpgradeExecutorMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeExecutorMock *UpgradeExecutorMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutorMock *UpgradeExecutorMockCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeExecutorMock.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutorMock *UpgradeExecutorMockSession) ADMINROLE() ([32]byte, error) {
	return _UpgradeExecutorMock.Contract.ADMINROLE(&_UpgradeExecutorMock.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutorMock *UpgradeExecutorMockCallerSession) ADMINROLE() ([32]byte, error) {
	return _UpgradeExecutorMock.Contract.ADMINROLE(&_UpgradeExecutorMock.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutorMock *UpgradeExecutorMockCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeExecutorMock.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutorMock *UpgradeExecutorMockSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UpgradeExecutorMock.Contract.DEFAULTADMINROLE(&_UpgradeExecutorMock.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutorMock *UpgradeExecutorMockCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UpgradeExecutorMock.Contract.DEFAULTADMINROLE(&_UpgradeExecutorMock.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_UpgradeExecutorMock *UpgradeExecutorMockCaller) EXECUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeExecutorMock.contract.Call(opts, &out, "EXECUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_UpgradeExecutorMock *UpgradeExecutorMockSession) EXECUTORROLE() ([32]byte, error) {
	return _UpgradeExecutorMock.Contract.EXECUTORROLE(&_UpgradeExecutorMock.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_UpgradeExecutorMock *UpgradeExecutorMockCallerSession) EXECUTORROLE() ([32]byte, error) {
	return _UpgradeExecutorMock.Contract.EXECUTORROLE(&_UpgradeExecutorMock.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UpgradeExecutorMock *UpgradeExecutorMockCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeExecutorMock.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UpgradeExecutorMock *UpgradeExecutorMockSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UpgradeExecutorMock.Contract.GetRoleAdmin(&_UpgradeExecutorMock.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UpgradeExecutorMock *UpgradeExecutorMockCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UpgradeExecutorMock.Contract.GetRoleAdmin(&_UpgradeExecutorMock.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UpgradeExecutorMock *UpgradeExecutorMockCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _UpgradeExecutorMock.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UpgradeExecutorMock *UpgradeExecutorMockSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UpgradeExecutorMock.Contract.HasRole(&_UpgradeExecutorMock.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UpgradeExecutorMock *UpgradeExecutorMockCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UpgradeExecutorMock.Contract.HasRole(&_UpgradeExecutorMock.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UpgradeExecutorMock *UpgradeExecutorMockCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _UpgradeExecutorMock.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UpgradeExecutorMock *UpgradeExecutorMockSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UpgradeExecutorMock.Contract.SupportsInterface(&_UpgradeExecutorMock.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UpgradeExecutorMock *UpgradeExecutorMockCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UpgradeExecutorMock.Contract.SupportsInterface(&_UpgradeExecutorMock.CallOpts, interfaceId)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address upgrade, bytes upgradeCallData) payable returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockTransactor) Execute(opts *bind.TransactOpts, upgrade common.Address, upgradeCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutorMock.contract.Transact(opts, "execute", upgrade, upgradeCallData)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address upgrade, bytes upgradeCallData) payable returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockSession) Execute(upgrade common.Address, upgradeCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.Execute(&_UpgradeExecutorMock.TransactOpts, upgrade, upgradeCallData)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address upgrade, bytes upgradeCallData) payable returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockTransactorSession) Execute(upgrade common.Address, upgradeCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.Execute(&_UpgradeExecutorMock.TransactOpts, upgrade, upgradeCallData)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0xbca8c7b5.
//
// Solidity: function executeCall(address target, bytes targetCallData) payable returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockTransactor) ExecuteCall(opts *bind.TransactOpts, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutorMock.contract.Transact(opts, "executeCall", target, targetCallData)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0xbca8c7b5.
//
// Solidity: function executeCall(address target, bytes targetCallData) payable returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockSession) ExecuteCall(target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.ExecuteCall(&_UpgradeExecutorMock.TransactOpts, target, targetCallData)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0xbca8c7b5.
//
// Solidity: function executeCall(address target, bytes targetCallData) payable returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockTransactorSession) ExecuteCall(target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.ExecuteCall(&_UpgradeExecutorMock.TransactOpts, target, targetCallData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutorMock.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.GrantRole(&_UpgradeExecutorMock.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.GrantRole(&_UpgradeExecutorMock.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address admin, address[] executors) returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, executors []common.Address) (*types.Transaction, error) {
	return _UpgradeExecutorMock.contract.Transact(opts, "initialize", admin, executors)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address admin, address[] executors) returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockSession) Initialize(admin common.Address, executors []common.Address) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.Initialize(&_UpgradeExecutorMock.TransactOpts, admin, executors)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address admin, address[] executors) returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockTransactorSession) Initialize(admin common.Address, executors []common.Address) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.Initialize(&_UpgradeExecutorMock.TransactOpts, admin, executors)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutorMock.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.RenounceRole(&_UpgradeExecutorMock.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.RenounceRole(&_UpgradeExecutorMock.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutorMock.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.RevokeRole(&_UpgradeExecutorMock.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UpgradeExecutorMock *UpgradeExecutorMockTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutorMock.Contract.RevokeRole(&_UpgradeExecutorMock.TransactOpts, role, account)
}

// UpgradeExecutorMockInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UpgradeExecutorMock contract.
type UpgradeExecutorMockInitializedIterator struct {
	Event *UpgradeExecutorMockInitialized // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorMockInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorMockInitialized)
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
		it.Event = new(UpgradeExecutorMockInitialized)
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
func (it *UpgradeExecutorMockInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorMockInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorMockInitialized represents a Initialized event raised by the UpgradeExecutorMock contract.
type UpgradeExecutorMockInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) FilterInitialized(opts *bind.FilterOpts) (*UpgradeExecutorMockInitializedIterator, error) {

	logs, sub, err := _UpgradeExecutorMock.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorMockInitializedIterator{contract: _UpgradeExecutorMock.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorMockInitialized) (event.Subscription, error) {

	logs, sub, err := _UpgradeExecutorMock.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorMockInitialized)
				if err := _UpgradeExecutorMock.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) ParseInitialized(log types.Log) (*UpgradeExecutorMockInitialized, error) {
	event := new(UpgradeExecutorMockInitialized)
	if err := _UpgradeExecutorMock.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorMockRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the UpgradeExecutorMock contract.
type UpgradeExecutorMockRoleAdminChangedIterator struct {
	Event *UpgradeExecutorMockRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorMockRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorMockRoleAdminChanged)
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
		it.Event = new(UpgradeExecutorMockRoleAdminChanged)
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
func (it *UpgradeExecutorMockRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorMockRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorMockRoleAdminChanged represents a RoleAdminChanged event raised by the UpgradeExecutorMock contract.
type UpgradeExecutorMockRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*UpgradeExecutorMockRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _UpgradeExecutorMock.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorMockRoleAdminChangedIterator{contract: _UpgradeExecutorMock.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorMockRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _UpgradeExecutorMock.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorMockRoleAdminChanged)
				if err := _UpgradeExecutorMock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) ParseRoleAdminChanged(log types.Log) (*UpgradeExecutorMockRoleAdminChanged, error) {
	event := new(UpgradeExecutorMockRoleAdminChanged)
	if err := _UpgradeExecutorMock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorMockRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the UpgradeExecutorMock contract.
type UpgradeExecutorMockRoleGrantedIterator struct {
	Event *UpgradeExecutorMockRoleGranted // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorMockRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorMockRoleGranted)
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
		it.Event = new(UpgradeExecutorMockRoleGranted)
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
func (it *UpgradeExecutorMockRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorMockRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorMockRoleGranted represents a RoleGranted event raised by the UpgradeExecutorMock contract.
type UpgradeExecutorMockRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UpgradeExecutorMockRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _UpgradeExecutorMock.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorMockRoleGrantedIterator{contract: _UpgradeExecutorMock.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorMockRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _UpgradeExecutorMock.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorMockRoleGranted)
				if err := _UpgradeExecutorMock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) ParseRoleGranted(log types.Log) (*UpgradeExecutorMockRoleGranted, error) {
	event := new(UpgradeExecutorMockRoleGranted)
	if err := _UpgradeExecutorMock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorMockRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the UpgradeExecutorMock contract.
type UpgradeExecutorMockRoleRevokedIterator struct {
	Event *UpgradeExecutorMockRoleRevoked // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorMockRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorMockRoleRevoked)
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
		it.Event = new(UpgradeExecutorMockRoleRevoked)
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
func (it *UpgradeExecutorMockRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorMockRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorMockRoleRevoked represents a RoleRevoked event raised by the UpgradeExecutorMock contract.
type UpgradeExecutorMockRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UpgradeExecutorMockRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _UpgradeExecutorMock.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorMockRoleRevokedIterator{contract: _UpgradeExecutorMock.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorMockRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _UpgradeExecutorMock.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorMockRoleRevoked)
				if err := _UpgradeExecutorMock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) ParseRoleRevoked(log types.Log) (*UpgradeExecutorMockRoleRevoked, error) {
	event := new(UpgradeExecutorMockRoleRevoked)
	if err := _UpgradeExecutorMock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorMockTargetCallExecutedIterator is returned from FilterTargetCallExecuted and is used to iterate over the raw logs and unpacked data for TargetCallExecuted events raised by the UpgradeExecutorMock contract.
type UpgradeExecutorMockTargetCallExecutedIterator struct {
	Event *UpgradeExecutorMockTargetCallExecuted // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorMockTargetCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorMockTargetCallExecuted)
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
		it.Event = new(UpgradeExecutorMockTargetCallExecuted)
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
func (it *UpgradeExecutorMockTargetCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorMockTargetCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorMockTargetCallExecuted represents a TargetCallExecuted event raised by the UpgradeExecutorMock contract.
type UpgradeExecutorMockTargetCallExecuted struct {
	Target common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTargetCallExecuted is a free log retrieval operation binding the contract event 0x4d7dbdcc249630ec373f584267f10abf44938de920c32562f5aee93959c25258.
//
// Solidity: event TargetCallExecuted(address indexed target, uint256 value, bytes data)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) FilterTargetCallExecuted(opts *bind.FilterOpts, target []common.Address) (*UpgradeExecutorMockTargetCallExecutedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _UpgradeExecutorMock.contract.FilterLogs(opts, "TargetCallExecuted", targetRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorMockTargetCallExecutedIterator{contract: _UpgradeExecutorMock.contract, event: "TargetCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTargetCallExecuted is a free log subscription operation binding the contract event 0x4d7dbdcc249630ec373f584267f10abf44938de920c32562f5aee93959c25258.
//
// Solidity: event TargetCallExecuted(address indexed target, uint256 value, bytes data)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) WatchTargetCallExecuted(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorMockTargetCallExecuted, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _UpgradeExecutorMock.contract.WatchLogs(opts, "TargetCallExecuted", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorMockTargetCallExecuted)
				if err := _UpgradeExecutorMock.contract.UnpackLog(event, "TargetCallExecuted", log); err != nil {
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

// ParseTargetCallExecuted is a log parse operation binding the contract event 0x4d7dbdcc249630ec373f584267f10abf44938de920c32562f5aee93959c25258.
//
// Solidity: event TargetCallExecuted(address indexed target, uint256 value, bytes data)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) ParseTargetCallExecuted(log types.Log) (*UpgradeExecutorMockTargetCallExecuted, error) {
	event := new(UpgradeExecutorMockTargetCallExecuted)
	if err := _UpgradeExecutorMock.contract.UnpackLog(event, "TargetCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorMockUpgradeExecutedIterator is returned from FilterUpgradeExecuted and is used to iterate over the raw logs and unpacked data for UpgradeExecuted events raised by the UpgradeExecutorMock contract.
type UpgradeExecutorMockUpgradeExecutedIterator struct {
	Event *UpgradeExecutorMockUpgradeExecuted // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorMockUpgradeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorMockUpgradeExecuted)
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
		it.Event = new(UpgradeExecutorMockUpgradeExecuted)
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
func (it *UpgradeExecutorMockUpgradeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorMockUpgradeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorMockUpgradeExecuted represents a UpgradeExecuted event raised by the UpgradeExecutorMock contract.
type UpgradeExecutorMockUpgradeExecuted struct {
	Upgrade common.Address
	Value   *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpgradeExecuted is a free log retrieval operation binding the contract event 0x49f6851d1cd01a518db5bdea5cffbbe90276baa2595f74250b7472b96806302e.
//
// Solidity: event UpgradeExecuted(address indexed upgrade, uint256 value, bytes data)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) FilterUpgradeExecuted(opts *bind.FilterOpts, upgrade []common.Address) (*UpgradeExecutorMockUpgradeExecutedIterator, error) {

	var upgradeRule []interface{}
	for _, upgradeItem := range upgrade {
		upgradeRule = append(upgradeRule, upgradeItem)
	}

	logs, sub, err := _UpgradeExecutorMock.contract.FilterLogs(opts, "UpgradeExecuted", upgradeRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorMockUpgradeExecutedIterator{contract: _UpgradeExecutorMock.contract, event: "UpgradeExecuted", logs: logs, sub: sub}, nil
}

// WatchUpgradeExecuted is a free log subscription operation binding the contract event 0x49f6851d1cd01a518db5bdea5cffbbe90276baa2595f74250b7472b96806302e.
//
// Solidity: event UpgradeExecuted(address indexed upgrade, uint256 value, bytes data)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) WatchUpgradeExecuted(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorMockUpgradeExecuted, upgrade []common.Address) (event.Subscription, error) {

	var upgradeRule []interface{}
	for _, upgradeItem := range upgrade {
		upgradeRule = append(upgradeRule, upgradeItem)
	}

	logs, sub, err := _UpgradeExecutorMock.contract.WatchLogs(opts, "UpgradeExecuted", upgradeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorMockUpgradeExecuted)
				if err := _UpgradeExecutorMock.contract.UnpackLog(event, "UpgradeExecuted", log); err != nil {
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

// ParseUpgradeExecuted is a log parse operation binding the contract event 0x49f6851d1cd01a518db5bdea5cffbbe90276baa2595f74250b7472b96806302e.
//
// Solidity: event UpgradeExecuted(address indexed upgrade, uint256 value, bytes data)
func (_UpgradeExecutorMock *UpgradeExecutorMockFilterer) ParseUpgradeExecuted(log types.Log) (*UpgradeExecutorMockUpgradeExecuted, error) {
	event := new(UpgradeExecutorMockUpgradeExecuted)
	if err := _UpgradeExecutorMock.contract.UnpackLog(event, "UpgradeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
