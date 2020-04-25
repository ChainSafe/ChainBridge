// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IMinterBurner

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

// IMinterBurnerABI is the input ABI used to generate the binding from.
const IMinterBurnerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IMinterBurner is an auto generated Go binding around an Ethereum contract.
type IMinterBurner struct {
	IMinterBurnerCaller     // Read-only binding to the contract
	IMinterBurnerTransactor // Write-only binding to the contract
	IMinterBurnerFilterer   // Log filterer for contract events
}

// IMinterBurnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMinterBurnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMinterBurnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMinterBurnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMinterBurnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMinterBurnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMinterBurnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMinterBurnerSession struct {
	Contract     *IMinterBurner    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMinterBurnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMinterBurnerCallerSession struct {
	Contract *IMinterBurnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IMinterBurnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMinterBurnerTransactorSession struct {
	Contract     *IMinterBurnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IMinterBurnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMinterBurnerRaw struct {
	Contract *IMinterBurner // Generic contract binding to access the raw methods on
}

// IMinterBurnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMinterBurnerCallerRaw struct {
	Contract *IMinterBurnerCaller // Generic read-only contract binding to access the raw methods on
}

// IMinterBurnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMinterBurnerTransactorRaw struct {
	Contract *IMinterBurnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMinterBurner creates a new instance of IMinterBurner, bound to a specific deployed contract.
func NewIMinterBurner(address common.Address, backend bind.ContractBackend) (*IMinterBurner, error) {
	contract, err := bindIMinterBurner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMinterBurner{IMinterBurnerCaller: IMinterBurnerCaller{contract: contract}, IMinterBurnerTransactor: IMinterBurnerTransactor{contract: contract}, IMinterBurnerFilterer: IMinterBurnerFilterer{contract: contract}}, nil
}

// NewIMinterBurnerCaller creates a new read-only instance of IMinterBurner, bound to a specific deployed contract.
func NewIMinterBurnerCaller(address common.Address, caller bind.ContractCaller) (*IMinterBurnerCaller, error) {
	contract, err := bindIMinterBurner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMinterBurnerCaller{contract: contract}, nil
}

// NewIMinterBurnerTransactor creates a new write-only instance of IMinterBurner, bound to a specific deployed contract.
func NewIMinterBurnerTransactor(address common.Address, transactor bind.ContractTransactor) (*IMinterBurnerTransactor, error) {
	contract, err := bindIMinterBurner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMinterBurnerTransactor{contract: contract}, nil
}

// NewIMinterBurnerFilterer creates a new log filterer instance of IMinterBurner, bound to a specific deployed contract.
func NewIMinterBurnerFilterer(address common.Address, filterer bind.ContractFilterer) (*IMinterBurnerFilterer, error) {
	contract, err := bindIMinterBurner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMinterBurnerFilterer{contract: contract}, nil
}

// bindIMinterBurner binds a generic wrapper to an already deployed contract.
func bindIMinterBurner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMinterBurnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMinterBurner *IMinterBurnerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IMinterBurner.Contract.IMinterBurnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMinterBurner *IMinterBurnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMinterBurner.Contract.IMinterBurnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMinterBurner *IMinterBurnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMinterBurner.Contract.IMinterBurnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMinterBurner *IMinterBurnerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IMinterBurner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMinterBurner *IMinterBurnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMinterBurner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMinterBurner *IMinterBurnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMinterBurner.Contract.contract.Transact(opts, method, params...)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address ) returns()
func (_IMinterBurner *IMinterBurnerTransactor) SetBurnable(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IMinterBurner.contract.Transact(opts, "setBurnable", arg0)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address ) returns()
func (_IMinterBurner *IMinterBurnerSession) SetBurnable(arg0 common.Address) (*types.Transaction, error) {
	return _IMinterBurner.Contract.SetBurnable(&_IMinterBurner.TransactOpts, arg0)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address ) returns()
func (_IMinterBurner *IMinterBurnerTransactorSession) SetBurnable(arg0 common.Address) (*types.Transaction, error) {
	return _IMinterBurner.Contract.SetBurnable(&_IMinterBurner.TransactOpts, arg0)
}

var RuntimeBytecode = "0x"
