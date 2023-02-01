// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
)

// IAssertionChainMetaData contains all meta data concerning the IAssertionChain contract.
var IAssertionChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getFirstChildCreationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getInboxMsgCountSeen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getPredecessorId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getStateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getSuccessionChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"isFirstChild\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	parsed, err := abi.JSON(strings.NewReader(IAssertionChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// GetFirstChildCreationTime is a free data retrieval call binding the contract method 0x43ed6ad9.
//
// Solidity: function getFirstChildCreationTime(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainCaller) GetFirstChildCreationTime(opts *bind.CallOpts, assertionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "getFirstChildCreationTime", assertionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFirstChildCreationTime is a free data retrieval call binding the contract method 0x43ed6ad9.
//
// Solidity: function getFirstChildCreationTime(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainSession) GetFirstChildCreationTime(assertionId [32]byte) (*big.Int, error) {
	return _IAssertionChain.Contract.GetFirstChildCreationTime(&_IAssertionChain.CallOpts, assertionId)
}

// GetFirstChildCreationTime is a free data retrieval call binding the contract method 0x43ed6ad9.
//
// Solidity: function getFirstChildCreationTime(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainCallerSession) GetFirstChildCreationTime(assertionId [32]byte) (*big.Int, error) {
	return _IAssertionChain.Contract.GetFirstChildCreationTime(&_IAssertionChain.CallOpts, assertionId)
}

// GetHeight is a free data retrieval call binding the contract method 0x896efbf2.
//
// Solidity: function getHeight(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainCaller) GetHeight(opts *bind.CallOpts, assertionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "getHeight", assertionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHeight is a free data retrieval call binding the contract method 0x896efbf2.
//
// Solidity: function getHeight(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainSession) GetHeight(assertionId [32]byte) (*big.Int, error) {
	return _IAssertionChain.Contract.GetHeight(&_IAssertionChain.CallOpts, assertionId)
}

// GetHeight is a free data retrieval call binding the contract method 0x896efbf2.
//
// Solidity: function getHeight(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainCallerSession) GetHeight(assertionId [32]byte) (*big.Int, error) {
	return _IAssertionChain.Contract.GetHeight(&_IAssertionChain.CallOpts, assertionId)
}

// GetInboxMsgCountSeen is a free data retrieval call binding the contract method 0x7cfd5ab9.
//
// Solidity: function getInboxMsgCountSeen(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainCaller) GetInboxMsgCountSeen(opts *bind.CallOpts, assertionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "getInboxMsgCountSeen", assertionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInboxMsgCountSeen is a free data retrieval call binding the contract method 0x7cfd5ab9.
//
// Solidity: function getInboxMsgCountSeen(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainSession) GetInboxMsgCountSeen(assertionId [32]byte) (*big.Int, error) {
	return _IAssertionChain.Contract.GetInboxMsgCountSeen(&_IAssertionChain.CallOpts, assertionId)
}

// GetInboxMsgCountSeen is a free data retrieval call binding the contract method 0x7cfd5ab9.
//
// Solidity: function getInboxMsgCountSeen(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainCallerSession) GetInboxMsgCountSeen(assertionId [32]byte) (*big.Int, error) {
	return _IAssertionChain.Contract.GetInboxMsgCountSeen(&_IAssertionChain.CallOpts, assertionId)
}

// GetPredecessorId is a free data retrieval call binding the contract method 0x9ca565d4.
//
// Solidity: function getPredecessorId(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainCaller) GetPredecessorId(opts *bind.CallOpts, assertionId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "getPredecessorId", assertionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPredecessorId is a free data retrieval call binding the contract method 0x9ca565d4.
//
// Solidity: function getPredecessorId(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainSession) GetPredecessorId(assertionId [32]byte) ([32]byte, error) {
	return _IAssertionChain.Contract.GetPredecessorId(&_IAssertionChain.CallOpts, assertionId)
}

// GetPredecessorId is a free data retrieval call binding the contract method 0x9ca565d4.
//
// Solidity: function getPredecessorId(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainCallerSession) GetPredecessorId(assertionId [32]byte) ([32]byte, error) {
	return _IAssertionChain.Contract.GetPredecessorId(&_IAssertionChain.CallOpts, assertionId)
}

// GetStateHash is a free data retrieval call binding the contract method 0xf9bce634.
//
// Solidity: function getStateHash(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainCaller) GetStateHash(opts *bind.CallOpts, assertionId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "getStateHash", assertionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStateHash is a free data retrieval call binding the contract method 0xf9bce634.
//
// Solidity: function getStateHash(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainSession) GetStateHash(assertionId [32]byte) ([32]byte, error) {
	return _IAssertionChain.Contract.GetStateHash(&_IAssertionChain.CallOpts, assertionId)
}

// GetStateHash is a free data retrieval call binding the contract method 0xf9bce634.
//
// Solidity: function getStateHash(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainCallerSession) GetStateHash(assertionId [32]byte) ([32]byte, error) {
	return _IAssertionChain.Contract.GetStateHash(&_IAssertionChain.CallOpts, assertionId)
}

// GetSuccessionChallenge is a free data retrieval call binding the contract method 0x2fefa18a.
//
// Solidity: function getSuccessionChallenge(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainCaller) GetSuccessionChallenge(opts *bind.CallOpts, assertionId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "getSuccessionChallenge", assertionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSuccessionChallenge is a free data retrieval call binding the contract method 0x2fefa18a.
//
// Solidity: function getSuccessionChallenge(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainSession) GetSuccessionChallenge(assertionId [32]byte) ([32]byte, error) {
	return _IAssertionChain.Contract.GetSuccessionChallenge(&_IAssertionChain.CallOpts, assertionId)
}

// GetSuccessionChallenge is a free data retrieval call binding the contract method 0x2fefa18a.
//
// Solidity: function getSuccessionChallenge(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainCallerSession) GetSuccessionChallenge(assertionId [32]byte) ([32]byte, error) {
	return _IAssertionChain.Contract.GetSuccessionChallenge(&_IAssertionChain.CallOpts, assertionId)
}

// IsFirstChild is a free data retrieval call binding the contract method 0x30836228.
//
// Solidity: function isFirstChild(bytes32 assertionId) view returns(bool)
func (_IAssertionChain *IAssertionChainCaller) IsFirstChild(opts *bind.CallOpts, assertionId [32]byte) (bool, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "isFirstChild", assertionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFirstChild is a free data retrieval call binding the contract method 0x30836228.
//
// Solidity: function isFirstChild(bytes32 assertionId) view returns(bool)
func (_IAssertionChain *IAssertionChainSession) IsFirstChild(assertionId [32]byte) (bool, error) {
	return _IAssertionChain.Contract.IsFirstChild(&_IAssertionChain.CallOpts, assertionId)
}

// IsFirstChild is a free data retrieval call binding the contract method 0x30836228.
//
// Solidity: function isFirstChild(bytes32 assertionId) view returns(bool)
func (_IAssertionChain *IAssertionChainCallerSession) IsFirstChild(assertionId [32]byte) (bool, error) {
	return _IAssertionChain.Contract.IsFirstChild(&_IAssertionChain.CallOpts, assertionId)
}
