// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IDepositHandler

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

// IDepositHandlerABI is the input ABI used to generate the binding from.
const IDepositHandlerABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IDepositHandler is an auto generated Go binding around an Ethereum contract.
type IDepositHandler struct {
	IDepositHandlerCaller     // Read-only binding to the contract
	IDepositHandlerTransactor // Write-only binding to the contract
	IDepositHandlerFilterer   // Log filterer for contract events
}

// IDepositHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDepositHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDepositHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDepositHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDepositHandlerSession struct {
	Contract     *IDepositHandler  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDepositHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDepositHandlerCallerSession struct {
	Contract *IDepositHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IDepositHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDepositHandlerTransactorSession struct {
	Contract     *IDepositHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IDepositHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDepositHandlerRaw struct {
	Contract *IDepositHandler // Generic contract binding to access the raw methods on
}

// IDepositHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDepositHandlerCallerRaw struct {
	Contract *IDepositHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IDepositHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDepositHandlerTransactorRaw struct {
	Contract *IDepositHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDepositHandler creates a new instance of IDepositHandler, bound to a specific deployed contract.
func NewIDepositHandler(address common.Address, backend bind.ContractBackend) (*IDepositHandler, error) {
	contract, err := bindIDepositHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDepositHandler{IDepositHandlerCaller: IDepositHandlerCaller{contract: contract}, IDepositHandlerTransactor: IDepositHandlerTransactor{contract: contract}, IDepositHandlerFilterer: IDepositHandlerFilterer{contract: contract}}, nil
}

// NewIDepositHandlerCaller creates a new read-only instance of IDepositHandler, bound to a specific deployed contract.
func NewIDepositHandlerCaller(address common.Address, caller bind.ContractCaller) (*IDepositHandlerCaller, error) {
	contract, err := bindIDepositHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDepositHandlerCaller{contract: contract}, nil
}

// NewIDepositHandlerTransactor creates a new write-only instance of IDepositHandler, bound to a specific deployed contract.
func NewIDepositHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IDepositHandlerTransactor, error) {
	contract, err := bindIDepositHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDepositHandlerTransactor{contract: contract}, nil
}

// NewIDepositHandlerFilterer creates a new log filterer instance of IDepositHandler, bound to a specific deployed contract.
func NewIDepositHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IDepositHandlerFilterer, error) {
	contract, err := bindIDepositHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDepositHandlerFilterer{contract: contract}, nil
}

// bindIDepositHandler binds a generic wrapper to an already deployed contract.
func bindIDepositHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDepositHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDepositHandler *IDepositHandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IDepositHandler.Contract.IDepositHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDepositHandler *IDepositHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDepositHandler.Contract.IDepositHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDepositHandler *IDepositHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDepositHandler.Contract.IDepositHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDepositHandler *IDepositHandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IDepositHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDepositHandler *IDepositHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDepositHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDepositHandler *IDepositHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDepositHandler.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_IDepositHandler *IDepositHandlerTransactor) Deposit(opts *bind.TransactOpts, destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _IDepositHandler.contract.Transact(opts, "deposit", destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_IDepositHandler *IDepositHandlerSession) Deposit(destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _IDepositHandler.Contract.Deposit(&_IDepositHandler.TransactOpts, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_IDepositHandler *IDepositHandlerTransactorSession) Deposit(destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _IDepositHandler.Contract.Deposit(&_IDepositHandler.TransactOpts, destinationChainID, depositNonce, depositer, data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_IDepositHandler *IDepositHandlerTransactor) ExecuteDeposit(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IDepositHandler.contract.Transact(opts, "executeDeposit", data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_IDepositHandler *IDepositHandlerSession) ExecuteDeposit(data []byte) (*types.Transaction, error) {
	return _IDepositHandler.Contract.ExecuteDeposit(&_IDepositHandler.TransactOpts, data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_IDepositHandler *IDepositHandlerTransactorSession) ExecuteDeposit(data []byte) (*types.Transaction, error) {
	return _IDepositHandler.Contract.ExecuteDeposit(&_IDepositHandler.TransactOpts, data)
}
var RuntimeBytecode = "0x"
