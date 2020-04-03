// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ICentrifugeAssetHandler

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

// ICentrifugeAssetHandlerABI is the input ABI used to generate the binding from.
const ICentrifugeAssetHandlerABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"metaDataHash\",\"type\":\"bytes32\"}],\"name\":\"depositAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ICentrifugeAssetHandler is an auto generated Go binding around an Ethereum contract.
type ICentrifugeAssetHandler struct {
	ICentrifugeAssetHandlerCaller     // Read-only binding to the contract
	ICentrifugeAssetHandlerTransactor // Write-only binding to the contract
	ICentrifugeAssetHandlerFilterer   // Log filterer for contract events
}

// ICentrifugeAssetHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICentrifugeAssetHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICentrifugeAssetHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICentrifugeAssetHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICentrifugeAssetHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICentrifugeAssetHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICentrifugeAssetHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICentrifugeAssetHandlerSession struct {
	Contract     *ICentrifugeAssetHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ICentrifugeAssetHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICentrifugeAssetHandlerCallerSession struct {
	Contract *ICentrifugeAssetHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ICentrifugeAssetHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICentrifugeAssetHandlerTransactorSession struct {
	Contract     *ICentrifugeAssetHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ICentrifugeAssetHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICentrifugeAssetHandlerRaw struct {
	Contract *ICentrifugeAssetHandler // Generic contract binding to access the raw methods on
}

// ICentrifugeAssetHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICentrifugeAssetHandlerCallerRaw struct {
	Contract *ICentrifugeAssetHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ICentrifugeAssetHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICentrifugeAssetHandlerTransactorRaw struct {
	Contract *ICentrifugeAssetHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICentrifugeAssetHandler creates a new instance of ICentrifugeAssetHandler, bound to a specific deployed contract.
func NewICentrifugeAssetHandler(address common.Address, backend bind.ContractBackend) (*ICentrifugeAssetHandler, error) {
	contract, err := bindICentrifugeAssetHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICentrifugeAssetHandler{ICentrifugeAssetHandlerCaller: ICentrifugeAssetHandlerCaller{contract: contract}, ICentrifugeAssetHandlerTransactor: ICentrifugeAssetHandlerTransactor{contract: contract}, ICentrifugeAssetHandlerFilterer: ICentrifugeAssetHandlerFilterer{contract: contract}}, nil
}

// NewICentrifugeAssetHandlerCaller creates a new read-only instance of ICentrifugeAssetHandler, bound to a specific deployed contract.
func NewICentrifugeAssetHandlerCaller(address common.Address, caller bind.ContractCaller) (*ICentrifugeAssetHandlerCaller, error) {
	contract, err := bindICentrifugeAssetHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICentrifugeAssetHandlerCaller{contract: contract}, nil
}

// NewICentrifugeAssetHandlerTransactor creates a new write-only instance of ICentrifugeAssetHandler, bound to a specific deployed contract.
func NewICentrifugeAssetHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ICentrifugeAssetHandlerTransactor, error) {
	contract, err := bindICentrifugeAssetHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICentrifugeAssetHandlerTransactor{contract: contract}, nil
}

// NewICentrifugeAssetHandlerFilterer creates a new log filterer instance of ICentrifugeAssetHandler, bound to a specific deployed contract.
func NewICentrifugeAssetHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ICentrifugeAssetHandlerFilterer, error) {
	contract, err := bindICentrifugeAssetHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICentrifugeAssetHandlerFilterer{contract: contract}, nil
}

// bindICentrifugeAssetHandler binds a generic wrapper to an already deployed contract.
func bindICentrifugeAssetHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICentrifugeAssetHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICentrifugeAssetHandler *ICentrifugeAssetHandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICentrifugeAssetHandler.Contract.ICentrifugeAssetHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICentrifugeAssetHandler *ICentrifugeAssetHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICentrifugeAssetHandler.Contract.ICentrifugeAssetHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICentrifugeAssetHandler *ICentrifugeAssetHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICentrifugeAssetHandler.Contract.ICentrifugeAssetHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICentrifugeAssetHandler *ICentrifugeAssetHandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICentrifugeAssetHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICentrifugeAssetHandler *ICentrifugeAssetHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICentrifugeAssetHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICentrifugeAssetHandler *ICentrifugeAssetHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICentrifugeAssetHandler.Contract.contract.Transact(opts, method, params...)
}

// DepositAsset is a paid mutator transaction binding the contract method 0x48ed85ce.
//
// Solidity: function depositAsset(bytes32 metaDataHash) returns()
func (_ICentrifugeAssetHandler *ICentrifugeAssetHandlerTransactor) DepositAsset(opts *bind.TransactOpts, metaDataHash [32]byte) (*types.Transaction, error) {
	return _ICentrifugeAssetHandler.contract.Transact(opts, "depositAsset", metaDataHash)
}

// DepositAsset is a paid mutator transaction binding the contract method 0x48ed85ce.
//
// Solidity: function depositAsset(bytes32 metaDataHash) returns()
func (_ICentrifugeAssetHandler *ICentrifugeAssetHandlerSession) DepositAsset(metaDataHash [32]byte) (*types.Transaction, error) {
	return _ICentrifugeAssetHandler.Contract.DepositAsset(&_ICentrifugeAssetHandler.TransactOpts, metaDataHash)
}

// DepositAsset is a paid mutator transaction binding the contract method 0x48ed85ce.
//
// Solidity: function depositAsset(bytes32 metaDataHash) returns()
func (_ICentrifugeAssetHandler *ICentrifugeAssetHandlerTransactorSession) DepositAsset(metaDataHash [32]byte) (*types.Transaction, error) {
	return _ICentrifugeAssetHandler.Contract.DepositAsset(&_ICentrifugeAssetHandler.TransactOpts, metaDataHash)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_ICentrifugeAssetHandler *ICentrifugeAssetHandlerTransactor) ExecuteDeposit(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ICentrifugeAssetHandler.contract.Transact(opts, "executeDeposit", data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_ICentrifugeAssetHandler *ICentrifugeAssetHandlerSession) ExecuteDeposit(data []byte) (*types.Transaction, error) {
	return _ICentrifugeAssetHandler.Contract.ExecuteDeposit(&_ICentrifugeAssetHandler.TransactOpts, data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_ICentrifugeAssetHandler *ICentrifugeAssetHandlerTransactorSession) ExecuteDeposit(data []byte) (*types.Transaction, error) {
	return _ICentrifugeAssetHandler.Contract.ExecuteDeposit(&_ICentrifugeAssetHandler.TransactOpts, data)
}
var RuntimeBytecode = "0x"
