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

// AddLeafArgs is an auto generated low-level Go binding around an user-defined struct.
type AddLeafArgs struct {
	ChallengeId            [32]byte
	ClaimId                [32]byte
	Height                 *big.Int
	HistoryCommitment      [32]byte
	FirstState             [32]byte
	FirstStatehistoryProof []byte
	LastState              [32]byte
	LastStatehistoryProof  []byte
}

// ChallengeVertex is an auto generated low-level Go binding around an user-defined struct.
type ChallengeVertex struct {
	PredecessorId                   [32]byte
	SuccessionChallenge             [32]byte
	HistoryCommitment               [32]byte
	Height                          *big.Int
	ClaimId                         [32]byte
	Staker                          common.Address
	Status                          uint8
	PresumptiveSuccessorId          [32]byte
	PresumptiveSuccessorLastUpdated *big.Int
	FlushedPsTime                   *big.Int
	LowestHeightSucessorId          [32]byte
}

// IChallengeManagerMetaData contains all meta data concerning the IChallengeManager contract.
var IChallengeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"historyCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"firstState\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"firstStatehistoryProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"lastState\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"lastStatehistoryProof\",\"type\":\"bytes\"}],\"internalType\":\"structAddLeafArgs\",\"name\":\"leafData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof2\",\"type\":\"bytes\"}],\"name\":\"addLeaf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prefixHistoryCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefixProof\",\"type\":\"bytes\"}],\"name\":\"bisect\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"confirmForPsTimer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"confirmForSucessionChallengeWin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"startId\",\"type\":\"bytes32\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"child1Id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"child2Id\",\"type\":\"bytes32\"}],\"name\":\"createSubChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"getCurrentPsTimer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"getVertex\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"predecessorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"successionChallenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"historyCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"enumStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"presumptiveSuccessorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"presumptiveSuccessorLastUpdated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flushedPsTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lowestHeightSucessorId\",\"type\":\"bytes32\"}],\"internalType\":\"structChallengeVertex\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prefixHistoryCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefixProof\",\"type\":\"bytes\"}],\"name\":\"merge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"vertexExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"}],\"name\":\"winningClaim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IChallengeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IChallengeManagerMetaData.ABI instead.
var IChallengeManagerABI = IChallengeManagerMetaData.ABI

// IChallengeManager is an auto generated Go binding around an Ethereum contract.
type IChallengeManager struct {
	IChallengeManagerCaller     // Read-only binding to the contract
	IChallengeManagerTransactor // Write-only binding to the contract
	IChallengeManagerFilterer   // Log filterer for contract events
}

// IChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeManagerSession struct {
	Contract     *IChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeManagerCallerSession struct {
	Contract *IChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeManagerTransactorSession struct {
	Contract     *IChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeManagerRaw struct {
	Contract *IChallengeManager // Generic contract binding to access the raw methods on
}

// IChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeManagerCallerRaw struct {
	Contract *IChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactorRaw struct {
	Contract *IChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeManager creates a new instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManager(address common.Address, backend bind.ContractBackend) (*IChallengeManager, error) {
	contract, err := bindIChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeManager{IChallengeManagerCaller: IChallengeManagerCaller{contract: contract}, IChallengeManagerTransactor: IChallengeManagerTransactor{contract: contract}, IChallengeManagerFilterer: IChallengeManagerFilterer{contract: contract}}, nil
}

// NewIChallengeManagerCaller creates a new read-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*IChallengeManagerCaller, error) {
	contract, err := bindIChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerCaller{contract: contract}, nil
}

// NewIChallengeManagerTransactor creates a new write-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeManagerTransactor, error) {
	contract, err := bindIChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerTransactor{contract: contract}, nil
}

// NewIChallengeManagerFilterer creates a new log filterer instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeManagerFilterer, error) {
	contract, err := bindIChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerFilterer{contract: contract}, nil
}

// bindIChallengeManager binds a generic wrapper to an already deployed contract.
func bindIChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.IChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentPsTimer is a free data retrieval call binding the contract method 0x78425139.
//
// Solidity: function getCurrentPsTimer(bytes32 challengeId, bytes32 vId) view returns(uint256)
func (_IChallengeManager *IChallengeManagerCaller) GetCurrentPsTimer(opts *bind.CallOpts, challengeId [32]byte, vId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "getCurrentPsTimer", challengeId, vId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPsTimer is a free data retrieval call binding the contract method 0x78425139.
//
// Solidity: function getCurrentPsTimer(bytes32 challengeId, bytes32 vId) view returns(uint256)
func (_IChallengeManager *IChallengeManagerSession) GetCurrentPsTimer(challengeId [32]byte, vId [32]byte) (*big.Int, error) {
	return _IChallengeManager.Contract.GetCurrentPsTimer(&_IChallengeManager.CallOpts, challengeId, vId)
}

// GetCurrentPsTimer is a free data retrieval call binding the contract method 0x78425139.
//
// Solidity: function getCurrentPsTimer(bytes32 challengeId, bytes32 vId) view returns(uint256)
func (_IChallengeManager *IChallengeManagerCallerSession) GetCurrentPsTimer(challengeId [32]byte, vId [32]byte) (*big.Int, error) {
	return _IChallengeManager.Contract.GetCurrentPsTimer(&_IChallengeManager.CallOpts, challengeId, vId)
}

// GetVertex is a free data retrieval call binding the contract method 0x950fd09d.
//
// Solidity: function getVertex(bytes32 challengeId, bytes32 vId) view returns((bytes32,bytes32,bytes32,uint256,bytes32,address,uint8,bytes32,uint256,uint256,bytes32))
func (_IChallengeManager *IChallengeManagerCaller) GetVertex(opts *bind.CallOpts, challengeId [32]byte, vId [32]byte) (ChallengeVertex, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "getVertex", challengeId, vId)

	if err != nil {
		return *new(ChallengeVertex), err
	}

	out0 := *abi.ConvertType(out[0], new(ChallengeVertex)).(*ChallengeVertex)

	return out0, err

}

// GetVertex is a free data retrieval call binding the contract method 0x950fd09d.
//
// Solidity: function getVertex(bytes32 challengeId, bytes32 vId) view returns((bytes32,bytes32,bytes32,uint256,bytes32,address,uint8,bytes32,uint256,uint256,bytes32))
func (_IChallengeManager *IChallengeManagerSession) GetVertex(challengeId [32]byte, vId [32]byte) (ChallengeVertex, error) {
	return _IChallengeManager.Contract.GetVertex(&_IChallengeManager.CallOpts, challengeId, vId)
}

// GetVertex is a free data retrieval call binding the contract method 0x950fd09d.
//
// Solidity: function getVertex(bytes32 challengeId, bytes32 vId) view returns((bytes32,bytes32,bytes32,uint256,bytes32,address,uint8,bytes32,uint256,uint256,bytes32))
func (_IChallengeManager *IChallengeManagerCallerSession) GetVertex(challengeId [32]byte, vId [32]byte) (ChallengeVertex, error) {
	return _IChallengeManager.Contract.GetVertex(&_IChallengeManager.CallOpts, challengeId, vId)
}

// VertexExists is a free data retrieval call binding the contract method 0xf0ebc431.
//
// Solidity: function vertexExists(bytes32 challengeId, bytes32 vId) view returns(bool)
func (_IChallengeManager *IChallengeManagerCaller) VertexExists(opts *bind.CallOpts, challengeId [32]byte, vId [32]byte) (bool, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "vertexExists", challengeId, vId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VertexExists is a free data retrieval call binding the contract method 0xf0ebc431.
//
// Solidity: function vertexExists(bytes32 challengeId, bytes32 vId) view returns(bool)
func (_IChallengeManager *IChallengeManagerSession) VertexExists(challengeId [32]byte, vId [32]byte) (bool, error) {
	return _IChallengeManager.Contract.VertexExists(&_IChallengeManager.CallOpts, challengeId, vId)
}

// VertexExists is a free data retrieval call binding the contract method 0xf0ebc431.
//
// Solidity: function vertexExists(bytes32 challengeId, bytes32 vId) view returns(bool)
func (_IChallengeManager *IChallengeManagerCallerSession) VertexExists(challengeId [32]byte, vId [32]byte) (bool, error) {
	return _IChallengeManager.Contract.VertexExists(&_IChallengeManager.CallOpts, challengeId, vId)
}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_IChallengeManager *IChallengeManagerCaller) WinningClaim(opts *bind.CallOpts, challengeId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "winningClaim", challengeId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_IChallengeManager *IChallengeManagerSession) WinningClaim(challengeId [32]byte) ([32]byte, error) {
	return _IChallengeManager.Contract.WinningClaim(&_IChallengeManager.CallOpts, challengeId)
}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_IChallengeManager *IChallengeManagerCallerSession) WinningClaim(challengeId [32]byte) ([32]byte, error) {
	return _IChallengeManager.Contract.WinningClaim(&_IChallengeManager.CallOpts, challengeId)
}

// AddLeaf is a paid mutator transaction binding the contract method 0xb241493b.
//
// Solidity: function addLeaf((bytes32,bytes32,uint256,bytes32,bytes32,bytes,bytes32,bytes) leafData, bytes proof1, bytes proof2) returns()
func (_IChallengeManager *IChallengeManagerTransactor) AddLeaf(opts *bind.TransactOpts, leafData AddLeafArgs, proof1 []byte, proof2 []byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "addLeaf", leafData, proof1, proof2)
}

// AddLeaf is a paid mutator transaction binding the contract method 0xb241493b.
//
// Solidity: function addLeaf((bytes32,bytes32,uint256,bytes32,bytes32,bytes,bytes32,bytes) leafData, bytes proof1, bytes proof2) returns()
func (_IChallengeManager *IChallengeManagerSession) AddLeaf(leafData AddLeafArgs, proof1 []byte, proof2 []byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.AddLeaf(&_IChallengeManager.TransactOpts, leafData, proof1, proof2)
}

// AddLeaf is a paid mutator transaction binding the contract method 0xb241493b.
//
// Solidity: function addLeaf((bytes32,bytes32,uint256,bytes32,bytes32,bytes,bytes32,bytes) leafData, bytes proof1, bytes proof2) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) AddLeaf(leafData AddLeafArgs, proof1 []byte, proof2 []byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.AddLeaf(&_IChallengeManager.TransactOpts, leafData, proof1, proof2)
}

// Bisect is a paid mutator transaction binding the contract method 0x85e57f3d.
//
// Solidity: function bisect(bytes32 challengeId, bytes32 vId, bytes32 prefixHistoryCommitment, bytes prefixProof) returns()
func (_IChallengeManager *IChallengeManagerTransactor) Bisect(opts *bind.TransactOpts, challengeId [32]byte, vId [32]byte, prefixHistoryCommitment [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "bisect", challengeId, vId, prefixHistoryCommitment, prefixProof)
}

// Bisect is a paid mutator transaction binding the contract method 0x85e57f3d.
//
// Solidity: function bisect(bytes32 challengeId, bytes32 vId, bytes32 prefixHistoryCommitment, bytes prefixProof) returns()
func (_IChallengeManager *IChallengeManagerSession) Bisect(challengeId [32]byte, vId [32]byte, prefixHistoryCommitment [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Bisect(&_IChallengeManager.TransactOpts, challengeId, vId, prefixHistoryCommitment, prefixProof)
}

// Bisect is a paid mutator transaction binding the contract method 0x85e57f3d.
//
// Solidity: function bisect(bytes32 challengeId, bytes32 vId, bytes32 prefixHistoryCommitment, bytes prefixProof) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) Bisect(challengeId [32]byte, vId [32]byte, prefixHistoryCommitment [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Bisect(&_IChallengeManager.TransactOpts, challengeId, vId, prefixHistoryCommitment, prefixProof)
}

// ConfirmForPsTimer is a paid mutator transaction binding the contract method 0x8995f555.
//
// Solidity: function confirmForPsTimer(bytes32 challengeId, bytes32 vId) returns()
func (_IChallengeManager *IChallengeManagerTransactor) ConfirmForPsTimer(opts *bind.TransactOpts, challengeId [32]byte, vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "confirmForPsTimer", challengeId, vId)
}

// ConfirmForPsTimer is a paid mutator transaction binding the contract method 0x8995f555.
//
// Solidity: function confirmForPsTimer(bytes32 challengeId, bytes32 vId) returns()
func (_IChallengeManager *IChallengeManagerSession) ConfirmForPsTimer(challengeId [32]byte, vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.ConfirmForPsTimer(&_IChallengeManager.TransactOpts, challengeId, vId)
}

// ConfirmForPsTimer is a paid mutator transaction binding the contract method 0x8995f555.
//
// Solidity: function confirmForPsTimer(bytes32 challengeId, bytes32 vId) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) ConfirmForPsTimer(challengeId [32]byte, vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.ConfirmForPsTimer(&_IChallengeManager.TransactOpts, challengeId, vId)
}

// ConfirmForSucessionChallengeWin is a paid mutator transaction binding the contract method 0x4690f553.
//
// Solidity: function confirmForSucessionChallengeWin(bytes32 challengeId, bytes32 vId) returns()
func (_IChallengeManager *IChallengeManagerTransactor) ConfirmForSucessionChallengeWin(opts *bind.TransactOpts, challengeId [32]byte, vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "confirmForSucessionChallengeWin", challengeId, vId)
}

// ConfirmForSucessionChallengeWin is a paid mutator transaction binding the contract method 0x4690f553.
//
// Solidity: function confirmForSucessionChallengeWin(bytes32 challengeId, bytes32 vId) returns()
func (_IChallengeManager *IChallengeManagerSession) ConfirmForSucessionChallengeWin(challengeId [32]byte, vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.ConfirmForSucessionChallengeWin(&_IChallengeManager.TransactOpts, challengeId, vId)
}

// ConfirmForSucessionChallengeWin is a paid mutator transaction binding the contract method 0x4690f553.
//
// Solidity: function confirmForSucessionChallengeWin(bytes32 challengeId, bytes32 vId) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) ConfirmForSucessionChallengeWin(challengeId [32]byte, vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.ConfirmForSucessionChallengeWin(&_IChallengeManager.TransactOpts, challengeId, vId)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xf696dc55.
//
// Solidity: function createChallenge(bytes32 startId) returns(bytes32)
func (_IChallengeManager *IChallengeManagerTransactor) CreateChallenge(opts *bind.TransactOpts, startId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "createChallenge", startId)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xf696dc55.
//
// Solidity: function createChallenge(bytes32 startId) returns(bytes32)
func (_IChallengeManager *IChallengeManagerSession) CreateChallenge(startId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.CreateChallenge(&_IChallengeManager.TransactOpts, startId)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xf696dc55.
//
// Solidity: function createChallenge(bytes32 startId) returns(bytes32)
func (_IChallengeManager *IChallengeManagerTransactorSession) CreateChallenge(startId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.CreateChallenge(&_IChallengeManager.TransactOpts, startId)
}

// CreateSubChallenge is a paid mutator transaction binding the contract method 0xe924e7e4.
//
// Solidity: function createSubChallenge(bytes32 challengeId, bytes32 child1Id, bytes32 child2Id) returns()
func (_IChallengeManager *IChallengeManagerTransactor) CreateSubChallenge(opts *bind.TransactOpts, challengeId [32]byte, child1Id [32]byte, child2Id [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "createSubChallenge", challengeId, child1Id, child2Id)
}

// CreateSubChallenge is a paid mutator transaction binding the contract method 0xe924e7e4.
//
// Solidity: function createSubChallenge(bytes32 challengeId, bytes32 child1Id, bytes32 child2Id) returns()
func (_IChallengeManager *IChallengeManagerSession) CreateSubChallenge(challengeId [32]byte, child1Id [32]byte, child2Id [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.CreateSubChallenge(&_IChallengeManager.TransactOpts, challengeId, child1Id, child2Id)
}

// CreateSubChallenge is a paid mutator transaction binding the contract method 0xe924e7e4.
//
// Solidity: function createSubChallenge(bytes32 challengeId, bytes32 child1Id, bytes32 child2Id) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) CreateSubChallenge(challengeId [32]byte, child1Id [32]byte, child2Id [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.CreateSubChallenge(&_IChallengeManager.TransactOpts, challengeId, child1Id, child2Id)
}

// Merge is a paid mutator transaction binding the contract method 0x62959d37.
//
// Solidity: function merge(bytes32 challengeId, bytes32 vId, bytes32 prefixHistoryCommitment, bytes prefixProof) returns()
func (_IChallengeManager *IChallengeManagerTransactor) Merge(opts *bind.TransactOpts, challengeId [32]byte, vId [32]byte, prefixHistoryCommitment [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "merge", challengeId, vId, prefixHistoryCommitment, prefixProof)
}

// Merge is a paid mutator transaction binding the contract method 0x62959d37.
//
// Solidity: function merge(bytes32 challengeId, bytes32 vId, bytes32 prefixHistoryCommitment, bytes prefixProof) returns()
func (_IChallengeManager *IChallengeManagerSession) Merge(challengeId [32]byte, vId [32]byte, prefixHistoryCommitment [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Merge(&_IChallengeManager.TransactOpts, challengeId, vId, prefixHistoryCommitment, prefixProof)
}

// Merge is a paid mutator transaction binding the contract method 0x62959d37.
//
// Solidity: function merge(bytes32 challengeId, bytes32 vId, bytes32 prefixHistoryCommitment, bytes prefixProof) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) Merge(challengeId [32]byte, vId [32]byte, prefixHistoryCommitment [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Merge(&_IChallengeManager.TransactOpts, challengeId, vId, prefixHistoryCommitment, prefixProof)
}
