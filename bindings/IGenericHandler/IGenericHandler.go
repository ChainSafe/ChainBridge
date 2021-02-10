// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IGenericHandler

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IGenericHandlerABI is the input ABI used to generate the binding from.
const IGenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IGenericHandler is an auto generated Go binding around an Ethereum contract.
type IGenericHandler struct {
	IGenericHandlerCaller     // Read-only binding to the contract
	IGenericHandlerTransactor // Write-only binding to the contract
	IGenericHandlerFilterer   // Log filterer for contract events
}

// IGenericHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGenericHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGenericHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGenericHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGenericHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGenericHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGenericHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGenericHandlerSession struct {
	Contract     *IGenericHandler  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGenericHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGenericHandlerCallerSession struct {
	Contract *IGenericHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IGenericHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGenericHandlerTransactorSession struct {
	Contract     *IGenericHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IGenericHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGenericHandlerRaw struct {
	Contract *IGenericHandler // Generic contract binding to access the raw methods on
}

// IGenericHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGenericHandlerCallerRaw struct {
	Contract *IGenericHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IGenericHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGenericHandlerTransactorRaw struct {
	Contract *IGenericHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGenericHandler creates a new instance of IGenericHandler, bound to a specific deployed contract.
func NewIGenericHandler(address common.Address, backend bind.ContractBackend) (*IGenericHandler, error) {
	contract, err := bindIGenericHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGenericHandler{IGenericHandlerCaller: IGenericHandlerCaller{contract: contract}, IGenericHandlerTransactor: IGenericHandlerTransactor{contract: contract}, IGenericHandlerFilterer: IGenericHandlerFilterer{contract: contract}}, nil
}

// NewIGenericHandlerCaller creates a new read-only instance of IGenericHandler, bound to a specific deployed contract.
func NewIGenericHandlerCaller(address common.Address, caller bind.ContractCaller) (*IGenericHandlerCaller, error) {
	contract, err := bindIGenericHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGenericHandlerCaller{contract: contract}, nil
}

// NewIGenericHandlerTransactor creates a new write-only instance of IGenericHandler, bound to a specific deployed contract.
func NewIGenericHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IGenericHandlerTransactor, error) {
	contract, err := bindIGenericHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGenericHandlerTransactor{contract: contract}, nil
}

// NewIGenericHandlerFilterer creates a new log filterer instance of IGenericHandler, bound to a specific deployed contract.
func NewIGenericHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IGenericHandlerFilterer, error) {
	contract, err := bindIGenericHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGenericHandlerFilterer{contract: contract}, nil
}

// bindIGenericHandler binds a generic wrapper to an already deployed contract.
func bindIGenericHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGenericHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGenericHandler *IGenericHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGenericHandler.Contract.IGenericHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGenericHandler *IGenericHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGenericHandler.Contract.IGenericHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGenericHandler *IGenericHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGenericHandler.Contract.IGenericHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGenericHandler *IGenericHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGenericHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGenericHandler *IGenericHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGenericHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGenericHandler *IGenericHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGenericHandler.Contract.contract.Transact(opts, method, params...)
}

// SetResource is a paid mutator transaction binding the contract method 0xbba8185a.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_IGenericHandler *IGenericHandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _IGenericHandler.contract.Transact(opts, "setResource", resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// SetResource is a paid mutator transaction binding the contract method 0xbba8185a.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_IGenericHandler *IGenericHandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _IGenericHandler.Contract.SetResource(&_IGenericHandler.TransactOpts, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// SetResource is a paid mutator transaction binding the contract method 0xbba8185a.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_IGenericHandler *IGenericHandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _IGenericHandler.Contract.SetResource(&_IGenericHandler.TransactOpts, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}
