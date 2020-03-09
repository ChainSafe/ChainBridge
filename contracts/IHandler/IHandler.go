// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IHandler

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

// IHandlerABI is the input ABI used to generate the binding from.
const IHandlerABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_originChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"executeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IHandler is an auto generated Go binding around an Ethereum contract.
type IHandler struct {
	IHandlerCaller     // Read-only binding to the contract
	IHandlerTransactor // Write-only binding to the contract
	IHandlerFilterer   // Log filterer for contract events
}

// IHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IHandlerSession struct {
	Contract     *IHandler         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IHandlerCallerSession struct {
	Contract *IHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IHandlerTransactorSession struct {
	Contract     *IHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IHandlerRaw struct {
	Contract *IHandler // Generic contract binding to access the raw methods on
}

// IHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IHandlerCallerRaw struct {
	Contract *IHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IHandlerTransactorRaw struct {
	Contract *IHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIHandler creates a new instance of IHandler, bound to a specific deployed contract.
func NewIHandler(address common.Address, backend bind.ContractBackend) (*IHandler, error) {
	contract, err := bindIHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IHandler{IHandlerCaller: IHandlerCaller{contract: contract}, IHandlerTransactor: IHandlerTransactor{contract: contract}, IHandlerFilterer: IHandlerFilterer{contract: contract}}, nil
}

// NewIHandlerCaller creates a new read-only instance of IHandler, bound to a specific deployed contract.
func NewIHandlerCaller(address common.Address, caller bind.ContractCaller) (*IHandlerCaller, error) {
	contract, err := bindIHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IHandlerCaller{contract: contract}, nil
}

// NewIHandlerTransactor creates a new write-only instance of IHandler, bound to a specific deployed contract.
func NewIHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IHandlerTransactor, error) {
	contract, err := bindIHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IHandlerTransactor{contract: contract}, nil
}

// NewIHandlerFilterer creates a new log filterer instance of IHandler, bound to a specific deployed contract.
func NewIHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IHandlerFilterer, error) {
	contract, err := bindIHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IHandlerFilterer{contract: contract}, nil
}

// bindIHandler binds a generic wrapper to an already deployed contract.
func bindIHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHandler *IHandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IHandler.Contract.IHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHandler *IHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHandler.Contract.IHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHandler *IHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHandler.Contract.IHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHandler *IHandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHandler *IHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHandler *IHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHandler.Contract.contract.Transact(opts, method, params...)
}

// ExecuteTransfer is a paid mutator transaction binding the contract method 0x5dcd6536.
//
// Solidity: function executeTransfer(uint256 _originChain, bytes _data) returns()
func (_IHandler *IHandlerTransactor) ExecuteTransfer(opts *bind.TransactOpts, _originChain *big.Int, _data []byte) (*types.Transaction, error) {
	return _IHandler.contract.Transact(opts, "executeTransfer", _originChain, _data)
}

// ExecuteTransfer is a paid mutator transaction binding the contract method 0x5dcd6536.
//
// Solidity: function executeTransfer(uint256 _originChain, bytes _data) returns()
func (_IHandler *IHandlerSession) ExecuteTransfer(_originChain *big.Int, _data []byte) (*types.Transaction, error) {
	return _IHandler.Contract.ExecuteTransfer(&_IHandler.TransactOpts, _originChain, _data)
}

// ExecuteTransfer is a paid mutator transaction binding the contract method 0x5dcd6536.
//
// Solidity: function executeTransfer(uint256 _originChain, bytes _data) returns()
func (_IHandler *IHandlerTransactorSession) ExecuteTransfer(_originChain *big.Int, _data []byte) (*types.Transaction, error) {
	return _IHandler.Contract.ExecuteTransfer(&_IHandler.TransactOpts, _originChain, _data)
}
