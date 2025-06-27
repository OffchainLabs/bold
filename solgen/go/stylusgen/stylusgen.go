// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stylusgen

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

// StylusDeployerMetaData contains all meta data concerning the StylusDeployer contract.
var StylusDeployerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"}],\"name\":\"ContractDeploymentError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ContractInitializationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitValueButNotInitData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"excessValue\",\"type\":\"uint256\"}],\"name\":\"RefundExcessValueError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"deployedContract\",\"type\":\"address\"}],\"name\":\"ContractDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"initValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"initSalt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"requiresActivation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506108dd806100206000396000f3fe6080604052600436106100345760003560e01c8063835d1d4c146100395780639f40b3851461006e578063a9a8e4e91461009c575b600080fd5b34801561004557600080fd5b50610059610054366004610605565b6100d4565b60405190151581526020015b60405180910390f35b34801561007a57600080fd5b5061008e610089366004610684565b6101f0565b604051908152602001610065565b6100af6100aa3660046106d0565b610226565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610065565b6040517fd70c0ca700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82163f6004820152600090819060719063d70c0ca790602401602060405180830381865afa925050508015610161575060408051601f3d908101601f1916820190925261015e91810190610765565b60015b61016d57506000610170565b90505b607173ffffffffffffffffffffffffffffffffffffffff1663a996e0c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101e09190610765565b61ffff9182169116141592915050565b600083838360405160200161020793929190610787565b6040516020818303038152906040528051906020012090509392505050565b6000811561023c576102398286866101f0565b91505b600061027f88888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250879250610534915050565b9050600061028c826100d4565b90506000811561033c5760006102a287346107a1565b6040517f58c780c200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201529091506071906358c780c2908390602401604080518083038185885af1158015610312573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061033791906107db565b925050505b86156103ff576000808473ffffffffffffffffffffffffffffffffffffffff16888b8b60405161036d929190610807565b60006040518083038185875af1925050503d80600081146103aa576040519150601f19603f3d011682016040523d82523d6000602084013e6103af565b606091505b5091509150816103f85784816040517f88d8f57d0000000000000000000000000000000000000000000000000000000081526004016103ef92919061085d565b60405180910390fd5b5050610436565b8515610436576040517ecc797100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008661044383346107a1565b61044d91906107a1565b905080156104dc57604051600090339083908381818185875af1925050503d8060008114610497576040519150601f19603f3d011682016040523d82523d6000602084013e61049c565b606091505b50509050806104da576040517f3ea99169000000000000000000000000000000000000000000000000000000008152600481018390526024016103ef565b505b60405173ffffffffffffffffffffffffffffffffffffffff851681527f8ffcdc15a283d706d38281f500270d8b5a656918f555de0913d7455e3e6bc1bf9060200160405180910390a150919998505050505050505050565b60008251600003610571576040517f21744a5900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821561058a57828451602086016000f59050610596565b8351602085016000f090505b3d1519811516156105ad576040513d6000823e3d81fd5b73ffffffffffffffffffffffffffffffffffffffff81166105fc57836040517f794c92ce0000000000000000000000000000000000000000000000000000000081526004016103ef9190610894565b90505b92915050565b60006020828403121561061757600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146105fc57600080fd5b60008083601f84011261064d57600080fd5b50813567ffffffffffffffff81111561066557600080fd5b60208301915083602082850101111561067d57600080fd5b9250929050565b60008060006040848603121561069957600080fd5b83359250602084013567ffffffffffffffff8111156106b757600080fd5b6106c38682870161063b565b9497909650939450505050565b600080600080600080608087890312156106e957600080fd5b863567ffffffffffffffff8082111561070157600080fd5b61070d8a838b0161063b565b9098509650602089013591508082111561072657600080fd5b5061073389828a0161063b565b979a9699509760408101359660609091013595509350505050565b805161ffff8116811461076057600080fd5b919050565b60006020828403121561077757600080fd5b6107808261074e565b9392505050565b838152818360208301376000910160200190815292915050565b818103818111156105ff577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080604083850312156107ee57600080fd5b6107f78361074e565b9150602083015190509250929050565b8183823760009101908152919050565b6000815180845260005b8181101561083d57602081850181015186830182015201610821565b506000602082860101526020601f19601f83011685010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff8316815260406020820152600061088c6040830184610817565b949350505050565b602081526000610780602083018461081756fea26469706673582212202ffe15188103190347fe5bb6d173aaf7108665ba4de6a82347fb04622c2affa064736f6c63430008110033",
}

// StylusDeployerABI is the input ABI used to generate the binding from.
// Deprecated: Use StylusDeployerMetaData.ABI instead.
var StylusDeployerABI = StylusDeployerMetaData.ABI

// StylusDeployerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StylusDeployerMetaData.Bin instead.
var StylusDeployerBin = StylusDeployerMetaData.Bin

// DeployStylusDeployer deploys a new Ethereum contract, binding an instance of StylusDeployer to it.
func DeployStylusDeployer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StylusDeployer, error) {
	parsed, err := StylusDeployerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StylusDeployerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StylusDeployer{StylusDeployerCaller: StylusDeployerCaller{contract: contract}, StylusDeployerTransactor: StylusDeployerTransactor{contract: contract}, StylusDeployerFilterer: StylusDeployerFilterer{contract: contract}}, nil
}

// StylusDeployer is an auto generated Go binding around an Ethereum contract.
type StylusDeployer struct {
	StylusDeployerCaller     // Read-only binding to the contract
	StylusDeployerTransactor // Write-only binding to the contract
	StylusDeployerFilterer   // Log filterer for contract events
}

// StylusDeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StylusDeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StylusDeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StylusDeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StylusDeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StylusDeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StylusDeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StylusDeployerSession struct {
	Contract     *StylusDeployer   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StylusDeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StylusDeployerCallerSession struct {
	Contract *StylusDeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StylusDeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StylusDeployerTransactorSession struct {
	Contract     *StylusDeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StylusDeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StylusDeployerRaw struct {
	Contract *StylusDeployer // Generic contract binding to access the raw methods on
}

// StylusDeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StylusDeployerCallerRaw struct {
	Contract *StylusDeployerCaller // Generic read-only contract binding to access the raw methods on
}

// StylusDeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StylusDeployerTransactorRaw struct {
	Contract *StylusDeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStylusDeployer creates a new instance of StylusDeployer, bound to a specific deployed contract.
func NewStylusDeployer(address common.Address, backend bind.ContractBackend) (*StylusDeployer, error) {
	contract, err := bindStylusDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StylusDeployer{StylusDeployerCaller: StylusDeployerCaller{contract: contract}, StylusDeployerTransactor: StylusDeployerTransactor{contract: contract}, StylusDeployerFilterer: StylusDeployerFilterer{contract: contract}}, nil
}

// NewStylusDeployerCaller creates a new read-only instance of StylusDeployer, bound to a specific deployed contract.
func NewStylusDeployerCaller(address common.Address, caller bind.ContractCaller) (*StylusDeployerCaller, error) {
	contract, err := bindStylusDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StylusDeployerCaller{contract: contract}, nil
}

// NewStylusDeployerTransactor creates a new write-only instance of StylusDeployer, bound to a specific deployed contract.
func NewStylusDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*StylusDeployerTransactor, error) {
	contract, err := bindStylusDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StylusDeployerTransactor{contract: contract}, nil
}

// NewStylusDeployerFilterer creates a new log filterer instance of StylusDeployer, bound to a specific deployed contract.
func NewStylusDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*StylusDeployerFilterer, error) {
	contract, err := bindStylusDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StylusDeployerFilterer{contract: contract}, nil
}

// bindStylusDeployer binds a generic wrapper to an already deployed contract.
func bindStylusDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StylusDeployerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StylusDeployer *StylusDeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StylusDeployer.Contract.StylusDeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StylusDeployer *StylusDeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StylusDeployer.Contract.StylusDeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StylusDeployer *StylusDeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StylusDeployer.Contract.StylusDeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StylusDeployer *StylusDeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StylusDeployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StylusDeployer *StylusDeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StylusDeployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StylusDeployer *StylusDeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StylusDeployer.Contract.contract.Transact(opts, method, params...)
}

// InitSalt is a free data retrieval call binding the contract method 0x9f40b385.
//
// Solidity: function initSalt(bytes32 salt, bytes initData) pure returns(bytes32)
func (_StylusDeployer *StylusDeployerCaller) InitSalt(opts *bind.CallOpts, salt [32]byte, initData []byte) ([32]byte, error) {
	var out []interface{}
	err := _StylusDeployer.contract.Call(opts, &out, "initSalt", salt, initData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InitSalt is a free data retrieval call binding the contract method 0x9f40b385.
//
// Solidity: function initSalt(bytes32 salt, bytes initData) pure returns(bytes32)
func (_StylusDeployer *StylusDeployerSession) InitSalt(salt [32]byte, initData []byte) ([32]byte, error) {
	return _StylusDeployer.Contract.InitSalt(&_StylusDeployer.CallOpts, salt, initData)
}

// InitSalt is a free data retrieval call binding the contract method 0x9f40b385.
//
// Solidity: function initSalt(bytes32 salt, bytes initData) pure returns(bytes32)
func (_StylusDeployer *StylusDeployerCallerSession) InitSalt(salt [32]byte, initData []byte) ([32]byte, error) {
	return _StylusDeployer.Contract.InitSalt(&_StylusDeployer.CallOpts, salt, initData)
}

// RequiresActivation is a free data retrieval call binding the contract method 0x835d1d4c.
//
// Solidity: function requiresActivation(address addr) view returns(bool)
func (_StylusDeployer *StylusDeployerCaller) RequiresActivation(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _StylusDeployer.contract.Call(opts, &out, "requiresActivation", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequiresActivation is a free data retrieval call binding the contract method 0x835d1d4c.
//
// Solidity: function requiresActivation(address addr) view returns(bool)
func (_StylusDeployer *StylusDeployerSession) RequiresActivation(addr common.Address) (bool, error) {
	return _StylusDeployer.Contract.RequiresActivation(&_StylusDeployer.CallOpts, addr)
}

// RequiresActivation is a free data retrieval call binding the contract method 0x835d1d4c.
//
// Solidity: function requiresActivation(address addr) view returns(bool)
func (_StylusDeployer *StylusDeployerCallerSession) RequiresActivation(addr common.Address) (bool, error) {
	return _StylusDeployer.Contract.RequiresActivation(&_StylusDeployer.CallOpts, addr)
}

// Deploy is a paid mutator transaction binding the contract method 0xa9a8e4e9.
//
// Solidity: function deploy(bytes bytecode, bytes initData, uint256 initValue, bytes32 salt) payable returns(address)
func (_StylusDeployer *StylusDeployerTransactor) Deploy(opts *bind.TransactOpts, bytecode []byte, initData []byte, initValue *big.Int, salt [32]byte) (*types.Transaction, error) {
	return _StylusDeployer.contract.Transact(opts, "deploy", bytecode, initData, initValue, salt)
}

// Deploy is a paid mutator transaction binding the contract method 0xa9a8e4e9.
//
// Solidity: function deploy(bytes bytecode, bytes initData, uint256 initValue, bytes32 salt) payable returns(address)
func (_StylusDeployer *StylusDeployerSession) Deploy(bytecode []byte, initData []byte, initValue *big.Int, salt [32]byte) (*types.Transaction, error) {
	return _StylusDeployer.Contract.Deploy(&_StylusDeployer.TransactOpts, bytecode, initData, initValue, salt)
}

// Deploy is a paid mutator transaction binding the contract method 0xa9a8e4e9.
//
// Solidity: function deploy(bytes bytecode, bytes initData, uint256 initValue, bytes32 salt) payable returns(address)
func (_StylusDeployer *StylusDeployerTransactorSession) Deploy(bytecode []byte, initData []byte, initValue *big.Int, salt [32]byte) (*types.Transaction, error) {
	return _StylusDeployer.Contract.Deploy(&_StylusDeployer.TransactOpts, bytecode, initData, initValue, salt)
}

// StylusDeployerContractDeployedIterator is returned from FilterContractDeployed and is used to iterate over the raw logs and unpacked data for ContractDeployed events raised by the StylusDeployer contract.
type StylusDeployerContractDeployedIterator struct {
	Event *StylusDeployerContractDeployed // Event containing the contract specifics and raw log

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
func (it *StylusDeployerContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StylusDeployerContractDeployed)
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
		it.Event = new(StylusDeployerContractDeployed)
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
func (it *StylusDeployerContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StylusDeployerContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StylusDeployerContractDeployed represents a ContractDeployed event raised by the StylusDeployer contract.
type StylusDeployerContractDeployed struct {
	DeployedContract common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterContractDeployed is a free log retrieval operation binding the contract event 0x8ffcdc15a283d706d38281f500270d8b5a656918f555de0913d7455e3e6bc1bf.
//
// Solidity: event ContractDeployed(address deployedContract)
func (_StylusDeployer *StylusDeployerFilterer) FilterContractDeployed(opts *bind.FilterOpts) (*StylusDeployerContractDeployedIterator, error) {

	logs, sub, err := _StylusDeployer.contract.FilterLogs(opts, "ContractDeployed")
	if err != nil {
		return nil, err
	}
	return &StylusDeployerContractDeployedIterator{contract: _StylusDeployer.contract, event: "ContractDeployed", logs: logs, sub: sub}, nil
}

// WatchContractDeployed is a free log subscription operation binding the contract event 0x8ffcdc15a283d706d38281f500270d8b5a656918f555de0913d7455e3e6bc1bf.
//
// Solidity: event ContractDeployed(address deployedContract)
func (_StylusDeployer *StylusDeployerFilterer) WatchContractDeployed(opts *bind.WatchOpts, sink chan<- *StylusDeployerContractDeployed) (event.Subscription, error) {

	logs, sub, err := _StylusDeployer.contract.WatchLogs(opts, "ContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StylusDeployerContractDeployed)
				if err := _StylusDeployer.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
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

// ParseContractDeployed is a log parse operation binding the contract event 0x8ffcdc15a283d706d38281f500270d8b5a656918f555de0913d7455e3e6bc1bf.
//
// Solidity: event ContractDeployed(address deployedContract)
func (_StylusDeployer *StylusDeployerFilterer) ParseContractDeployed(log types.Log) (*StylusDeployerContractDeployed, error) {
	event := new(StylusDeployerContractDeployed)
	if err := _StylusDeployer.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
