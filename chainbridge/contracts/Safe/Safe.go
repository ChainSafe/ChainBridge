// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Safe

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SafeABI is the input ABI used to generate the binding from.
const SafeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"lockErc\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"releaseErc\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Safe is an auto generated Go binding around an Ethereum contract.
type Safe struct {
	SafeCaller     // Read-only binding to the contract
	SafeTransactor // Write-only binding to the contract
	SafeFilterer   // Log filterer for contract events
}

// SafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeSession struct {
	Contract     *Safe             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCallerSession struct {
	Contract *SafeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeTransactorSession struct {
	Contract     *SafeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeRaw struct {
	Contract *Safe // Generic contract binding to access the raw methods on
}

// SafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCallerRaw struct {
	Contract *SafeCaller // Generic read-only contract binding to access the raw methods on
}

// SafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeTransactorRaw struct {
	Contract *SafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafe creates a new instance of Safe, bound to a specific deployed contract.
func NewSafe(address common.Address, backend bind.ContractBackend) (*Safe, error) {
	contract, err := bindSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Safe{SafeCaller: SafeCaller{contract: contract}, SafeTransactor: SafeTransactor{contract: contract}, SafeFilterer: SafeFilterer{contract: contract}}, nil
}

// NewSafeCaller creates a new read-only instance of Safe, bound to a specific deployed contract.
func NewSafeCaller(address common.Address, caller bind.ContractCaller) (*SafeCaller, error) {
	contract, err := bindSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCaller{contract: contract}, nil
}

// NewSafeTransactor creates a new write-only instance of Safe, bound to a specific deployed contract.
func NewSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeTransactor, error) {
	contract, err := bindSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeTransactor{contract: contract}, nil
}

// NewSafeFilterer creates a new log filterer instance of Safe, bound to a specific deployed contract.
func NewSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeFilterer, error) {
	contract, err := bindSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeFilterer{contract: contract}, nil
}

// bindSafe binds a generic wrapper to an already deployed contract.
func bindSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safe *SafeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Safe.Contract.SafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safe *SafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safe.Contract.SafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safe *SafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safe.Contract.SafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safe *SafeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Safe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safe *SafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safe *SafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safe.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Safe *SafeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Safe *SafeSession) Owner() (common.Address, error) {
	return _Safe.Contract.Owner(&_Safe.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Safe *SafeCallerSession) Owner() (common.Address, error) {
	return _Safe.Contract.Owner(&_Safe.CallOpts)
}

// Lock is a paid mutator transaction binding the contract method 0x3fea56b8.
//
// Solidity: function lock(_tokenAddress address, _value uint256, _to address, _from address) returns()
func (_Safe *SafeTransactor) Lock(opts *bind.TransactOpts, _tokenAddress common.Address, _value *big.Int, _to common.Address, _from common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "lock", _tokenAddress, _value, _to, _from)
}

// Lock is a paid mutator transaction binding the contract method 0x3fea56b8.
//
// Solidity: function lock(_tokenAddress address, _value uint256, _to address, _from address) returns()
func (_Safe *SafeSession) Lock(_tokenAddress common.Address, _value *big.Int, _to common.Address, _from common.Address) (*types.Transaction, error) {
	return _Safe.Contract.Lock(&_Safe.TransactOpts, _tokenAddress, _value, _to, _from)
}

// Lock is a paid mutator transaction binding the contract method 0x3fea56b8.
//
// Solidity: function lock(_tokenAddress address, _value uint256, _to address, _from address) returns()
func (_Safe *SafeTransactorSession) Lock(_tokenAddress common.Address, _value *big.Int, _to common.Address, _from common.Address) (*types.Transaction, error) {
	return _Safe.Contract.Lock(&_Safe.TransactOpts, _tokenAddress, _value, _to, _from)
}

// LockErc is a paid mutator transaction binding the contract method 0x6d3cdd97.
//
// Solidity: function lockErc(_tokenAddress address, _value uint256, _to address, _from address) returns()
func (_Safe *SafeTransactor) LockErc(opts *bind.TransactOpts, _tokenAddress common.Address, _value *big.Int, _to common.Address, _from common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "lockErc", _tokenAddress, _value, _to, _from)
}

// LockErc is a paid mutator transaction binding the contract method 0x6d3cdd97.
//
// Solidity: function lockErc(_tokenAddress address, _value uint256, _to address, _from address) returns()
func (_Safe *SafeSession) LockErc(_tokenAddress common.Address, _value *big.Int, _to common.Address, _from common.Address) (*types.Transaction, error) {
	return _Safe.Contract.LockErc(&_Safe.TransactOpts, _tokenAddress, _value, _to, _from)
}

// LockErc is a paid mutator transaction binding the contract method 0x6d3cdd97.
//
// Solidity: function lockErc(_tokenAddress address, _value uint256, _to address, _from address) returns()
func (_Safe *SafeTransactorSession) LockErc(_tokenAddress common.Address, _value *big.Int, _to common.Address, _from common.Address) (*types.Transaction, error) {
	return _Safe.Contract.LockErc(&_Safe.TransactOpts, _tokenAddress, _value, _to, _from)
}

// Release is a paid mutator transaction binding the contract method 0x15f41501.
//
// Solidity: function release(_tokenAddress address, _value uint256, _to address) returns()
func (_Safe *SafeTransactor) Release(opts *bind.TransactOpts, _tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "release", _tokenAddress, _value, _to)
}

// Release is a paid mutator transaction binding the contract method 0x15f41501.
//
// Solidity: function release(_tokenAddress address, _value uint256, _to address) returns()
func (_Safe *SafeSession) Release(_tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Safe.Contract.Release(&_Safe.TransactOpts, _tokenAddress, _value, _to)
}

// Release is a paid mutator transaction binding the contract method 0x15f41501.
//
// Solidity: function release(_tokenAddress address, _value uint256, _to address) returns()
func (_Safe *SafeTransactorSession) Release(_tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Safe.Contract.Release(&_Safe.TransactOpts, _tokenAddress, _value, _to)
}

// ReleaseErc is a paid mutator transaction binding the contract method 0x34455733.
//
// Solidity: function releaseErc(_tokenAddress address, _value uint256, _to address) returns()
func (_Safe *SafeTransactor) ReleaseErc(opts *bind.TransactOpts, _tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "releaseErc", _tokenAddress, _value, _to)
}

// ReleaseErc is a paid mutator transaction binding the contract method 0x34455733.
//
// Solidity: function releaseErc(_tokenAddress address, _value uint256, _to address) returns()
func (_Safe *SafeSession) ReleaseErc(_tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Safe.Contract.ReleaseErc(&_Safe.TransactOpts, _tokenAddress, _value, _to)
}

// ReleaseErc is a paid mutator transaction binding the contract method 0x34455733.
//
// Solidity: function releaseErc(_tokenAddress address, _value uint256, _to address) returns()
func (_Safe *SafeTransactorSession) ReleaseErc(_tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Safe.Contract.ReleaseErc(&_Safe.TransactOpts, _tokenAddress, _value, _to)
}
