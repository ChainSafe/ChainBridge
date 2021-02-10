// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IDepositExecute

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

// IDepositExecuteABI is the input ABI used to generate the binding from.
const IDepositExecuteABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IDepositExecute is an auto generated Go binding around an Ethereum contract.
type IDepositExecute struct {
	IDepositExecuteCaller     // Read-only binding to the contract
	IDepositExecuteTransactor // Write-only binding to the contract
	IDepositExecuteFilterer   // Log filterer for contract events
}

// IDepositExecuteCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDepositExecuteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositExecuteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDepositExecuteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositExecuteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDepositExecuteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositExecuteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDepositExecuteSession struct {
	Contract     *IDepositExecute  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDepositExecuteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDepositExecuteCallerSession struct {
	Contract *IDepositExecuteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IDepositExecuteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDepositExecuteTransactorSession struct {
	Contract     *IDepositExecuteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IDepositExecuteRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDepositExecuteRaw struct {
	Contract *IDepositExecute // Generic contract binding to access the raw methods on
}

// IDepositExecuteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDepositExecuteCallerRaw struct {
	Contract *IDepositExecuteCaller // Generic read-only contract binding to access the raw methods on
}

// IDepositExecuteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDepositExecuteTransactorRaw struct {
	Contract *IDepositExecuteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDepositExecute creates a new instance of IDepositExecute, bound to a specific deployed contract.
func NewIDepositExecute(address common.Address, backend bind.ContractBackend) (*IDepositExecute, error) {
	contract, err := bindIDepositExecute(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDepositExecute{IDepositExecuteCaller: IDepositExecuteCaller{contract: contract}, IDepositExecuteTransactor: IDepositExecuteTransactor{contract: contract}, IDepositExecuteFilterer: IDepositExecuteFilterer{contract: contract}}, nil
}

// NewIDepositExecuteCaller creates a new read-only instance of IDepositExecute, bound to a specific deployed contract.
func NewIDepositExecuteCaller(address common.Address, caller bind.ContractCaller) (*IDepositExecuteCaller, error) {
	contract, err := bindIDepositExecute(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDepositExecuteCaller{contract: contract}, nil
}

// NewIDepositExecuteTransactor creates a new write-only instance of IDepositExecute, bound to a specific deployed contract.
func NewIDepositExecuteTransactor(address common.Address, transactor bind.ContractTransactor) (*IDepositExecuteTransactor, error) {
	contract, err := bindIDepositExecute(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDepositExecuteTransactor{contract: contract}, nil
}

// NewIDepositExecuteFilterer creates a new log filterer instance of IDepositExecute, bound to a specific deployed contract.
func NewIDepositExecuteFilterer(address common.Address, filterer bind.ContractFilterer) (*IDepositExecuteFilterer, error) {
	contract, err := bindIDepositExecute(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDepositExecuteFilterer{contract: contract}, nil
}

// bindIDepositExecute binds a generic wrapper to an already deployed contract.
func bindIDepositExecute(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDepositExecuteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDepositExecute *IDepositExecuteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDepositExecute.Contract.IDepositExecuteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDepositExecute *IDepositExecuteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDepositExecute.Contract.IDepositExecuteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDepositExecute *IDepositExecuteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDepositExecute.Contract.IDepositExecuteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDepositExecute *IDepositExecuteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDepositExecute.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDepositExecute *IDepositExecuteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDepositExecute.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDepositExecute *IDepositExecuteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDepositExecute.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_IDepositExecute *IDepositExecuteTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _IDepositExecute.contract.Transact(opts, "deposit", resourceID, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_IDepositExecute *IDepositExecuteSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _IDepositExecute.Contract.Deposit(&_IDepositExecute.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_IDepositExecute *IDepositExecuteTransactorSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _IDepositExecute.Contract.Deposit(&_IDepositExecute.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_IDepositExecute *IDepositExecuteTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _IDepositExecute.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_IDepositExecute *IDepositExecuteSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _IDepositExecute.Contract.ExecuteProposal(&_IDepositExecute.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_IDepositExecute *IDepositExecuteTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _IDepositExecute.Contract.ExecuteProposal(&_IDepositExecute.TransactOpts, resourceID, data)
}
