// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IRelayer

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

// IRelayerABI is the input ABI used to generate the binding from.
const IRelayerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalRelayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"},{\"internalType\":\"enumIRelayer.RelayerActionType\",\"name\":\"action\",\"type\":\"uint8\"}],\"name\":\"voteRelayerProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposedValue\",\"type\":\"uint256\"}],\"name\":\"createRelayerThresholdProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIRelayer.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"voteRelayerThresholdProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IRelayer is an auto generated Go binding around an Ethereum contract.
type IRelayer struct {
	IRelayerCaller     // Read-only binding to the contract
	IRelayerTransactor // Write-only binding to the contract
	IRelayerFilterer   // Log filterer for contract events
}

// IRelayerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRelayerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRelayerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRelayerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRelayerSession struct {
	Contract     *IRelayer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRelayerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRelayerCallerSession struct {
	Contract *IRelayerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IRelayerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRelayerTransactorSession struct {
	Contract     *IRelayerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IRelayerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRelayerRaw struct {
	Contract *IRelayer // Generic contract binding to access the raw methods on
}

// IRelayerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRelayerCallerRaw struct {
	Contract *IRelayerCaller // Generic read-only contract binding to access the raw methods on
}

// IRelayerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRelayerTransactorRaw struct {
	Contract *IRelayerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRelayer creates a new instance of IRelayer, bound to a specific deployed contract.
func NewIRelayer(address common.Address, backend bind.ContractBackend) (*IRelayer, error) {
	contract, err := bindIRelayer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRelayer{IRelayerCaller: IRelayerCaller{contract: contract}, IRelayerTransactor: IRelayerTransactor{contract: contract}, IRelayerFilterer: IRelayerFilterer{contract: contract}}, nil
}

// NewIRelayerCaller creates a new read-only instance of IRelayer, bound to a specific deployed contract.
func NewIRelayerCaller(address common.Address, caller bind.ContractCaller) (*IRelayerCaller, error) {
	contract, err := bindIRelayer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayerCaller{contract: contract}, nil
}

// NewIRelayerTransactor creates a new write-only instance of IRelayer, bound to a specific deployed contract.
func NewIRelayerTransactor(address common.Address, transactor bind.ContractTransactor) (*IRelayerTransactor, error) {
	contract, err := bindIRelayer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayerTransactor{contract: contract}, nil
}

// NewIRelayerFilterer creates a new log filterer instance of IRelayer, bound to a specific deployed contract.
func NewIRelayerFilterer(address common.Address, filterer bind.ContractFilterer) (*IRelayerFilterer, error) {
	contract, err := bindIRelayer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRelayerFilterer{contract: contract}, nil
}

// bindIRelayer binds a generic wrapper to an already deployed contract.
func bindIRelayer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRelayerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayer *IRelayerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IRelayer.Contract.IRelayerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayer *IRelayerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayer.Contract.IRelayerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayer *IRelayerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayer.Contract.IRelayerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayer *IRelayerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IRelayer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayer *IRelayerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayer *IRelayerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayer.Contract.contract.Transact(opts, method, params...)
}

// CreateRelayerThresholdProposal is a paid mutator transaction binding the contract method 0xdf269060.
//
// Solidity: function createRelayerThresholdProposal(uint256 proposedValue) returns()
func (_IRelayer *IRelayerTransactor) CreateRelayerThresholdProposal(opts *bind.TransactOpts, proposedValue *big.Int) (*types.Transaction, error) {
	return _IRelayer.contract.Transact(opts, "createRelayerThresholdProposal", proposedValue)
}

// CreateRelayerThresholdProposal is a paid mutator transaction binding the contract method 0xdf269060.
//
// Solidity: function createRelayerThresholdProposal(uint256 proposedValue) returns()
func (_IRelayer *IRelayerSession) CreateRelayerThresholdProposal(proposedValue *big.Int) (*types.Transaction, error) {
	return _IRelayer.Contract.CreateRelayerThresholdProposal(&_IRelayer.TransactOpts, proposedValue)
}

// CreateRelayerThresholdProposal is a paid mutator transaction binding the contract method 0xdf269060.
//
// Solidity: function createRelayerThresholdProposal(uint256 proposedValue) returns()
func (_IRelayer *IRelayerTransactorSession) CreateRelayerThresholdProposal(proposedValue *big.Int) (*types.Transaction, error) {
	return _IRelayer.Contract.CreateRelayerThresholdProposal(&_IRelayer.TransactOpts, proposedValue)
}

// GetTotalRelayers is a paid mutator transaction binding the contract method 0x933b4667.
//
// Solidity: function getTotalRelayers() returns(uint256)
func (_IRelayer *IRelayerTransactor) GetTotalRelayers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayer.contract.Transact(opts, "getTotalRelayers")
}

// GetTotalRelayers is a paid mutator transaction binding the contract method 0x933b4667.
//
// Solidity: function getTotalRelayers() returns(uint256)
func (_IRelayer *IRelayerSession) GetTotalRelayers() (*types.Transaction, error) {
	return _IRelayer.Contract.GetTotalRelayers(&_IRelayer.TransactOpts)
}

// GetTotalRelayers is a paid mutator transaction binding the contract method 0x933b4667.
//
// Solidity: function getTotalRelayers() returns(uint256)
func (_IRelayer *IRelayerTransactorSession) GetTotalRelayers() (*types.Transaction, error) {
	return _IRelayer.Contract.GetTotalRelayers(&_IRelayer.TransactOpts)
}

// IsRelayer is a paid mutator transaction binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayerAddress) returns(bool)
func (_IRelayer *IRelayerTransactor) IsRelayer(opts *bind.TransactOpts, relayerAddress common.Address) (*types.Transaction, error) {
	return _IRelayer.contract.Transact(opts, "isRelayer", relayerAddress)
}

// IsRelayer is a paid mutator transaction binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayerAddress) returns(bool)
func (_IRelayer *IRelayerSession) IsRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _IRelayer.Contract.IsRelayer(&_IRelayer.TransactOpts, relayerAddress)
}

// IsRelayer is a paid mutator transaction binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayerAddress) returns(bool)
func (_IRelayer *IRelayerTransactorSession) IsRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _IRelayer.Contract.IsRelayer(&_IRelayer.TransactOpts, relayerAddress)
}

// VoteRelayerProposal is a paid mutator transaction binding the contract method 0x4d4ebd7a.
//
// Solidity: function voteRelayerProposal(address proposedAddress, uint8 action) returns()
func (_IRelayer *IRelayerTransactor) VoteRelayerProposal(opts *bind.TransactOpts, proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _IRelayer.contract.Transact(opts, "voteRelayerProposal", proposedAddress, action)
}

// VoteRelayerProposal is a paid mutator transaction binding the contract method 0x4d4ebd7a.
//
// Solidity: function voteRelayerProposal(address proposedAddress, uint8 action) returns()
func (_IRelayer *IRelayerSession) VoteRelayerProposal(proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _IRelayer.Contract.VoteRelayerProposal(&_IRelayer.TransactOpts, proposedAddress, action)
}

// VoteRelayerProposal is a paid mutator transaction binding the contract method 0x4d4ebd7a.
//
// Solidity: function voteRelayerProposal(address proposedAddress, uint8 action) returns()
func (_IRelayer *IRelayerTransactorSession) VoteRelayerProposal(proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _IRelayer.Contract.VoteRelayerProposal(&_IRelayer.TransactOpts, proposedAddress, action)
}

// VoteRelayerThresholdProposal is a paid mutator transaction binding the contract method 0xe9cdaead.
//
// Solidity: function voteRelayerThresholdProposal(uint8 vote) returns()
func (_IRelayer *IRelayerTransactor) VoteRelayerThresholdProposal(opts *bind.TransactOpts, vote uint8) (*types.Transaction, error) {
	return _IRelayer.contract.Transact(opts, "voteRelayerThresholdProposal", vote)
}

// VoteRelayerThresholdProposal is a paid mutator transaction binding the contract method 0xe9cdaead.
//
// Solidity: function voteRelayerThresholdProposal(uint8 vote) returns()
func (_IRelayer *IRelayerSession) VoteRelayerThresholdProposal(vote uint8) (*types.Transaction, error) {
	return _IRelayer.Contract.VoteRelayerThresholdProposal(&_IRelayer.TransactOpts, vote)
}

// VoteRelayerThresholdProposal is a paid mutator transaction binding the contract method 0xe9cdaead.
//
// Solidity: function voteRelayerThresholdProposal(uint8 vote) returns()
func (_IRelayer *IRelayerTransactorSession) VoteRelayerThresholdProposal(vote uint8) (*types.Transaction, error) {
	return _IRelayer.Contract.VoteRelayerThresholdProposal(&_IRelayer.TransactOpts, vote)
}

var RuntimeBytecode = "0x"
