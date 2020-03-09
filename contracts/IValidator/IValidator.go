// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IValidator

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

// IValidatorABI is the input ABI used to generate the binding from.
const IValidatorABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getTotalValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"},{\"internalType\":\"enumIValidator.ValidatorActionType\",\"name\":\"action\",\"type\":\"uint8\"}],\"name\":\"createValidatorProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"},{\"internalType\":\"enumIValidator.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"voteValidatorProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposedValue\",\"type\":\"uint256\"}],\"name\":\"createValidatorThresholdProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"enumIValidator.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"voteValidatorThresholdProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IValidator is an auto generated Go binding around an Ethereum contract.
type IValidator struct {
	IValidatorCaller     // Read-only binding to the contract
	IValidatorTransactor // Write-only binding to the contract
	IValidatorFilterer   // Log filterer for contract events
}

// IValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IValidatorSession struct {
	Contract     *IValidator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IValidatorCallerSession struct {
	Contract *IValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IValidatorTransactorSession struct {
	Contract     *IValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IValidatorRaw struct {
	Contract *IValidator // Generic contract binding to access the raw methods on
}

// IValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IValidatorCallerRaw struct {
	Contract *IValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// IValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IValidatorTransactorRaw struct {
	Contract *IValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIValidator creates a new instance of IValidator, bound to a specific deployed contract.
func NewIValidator(address common.Address, backend bind.ContractBackend) (*IValidator, error) {
	contract, err := bindIValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IValidator{IValidatorCaller: IValidatorCaller{contract: contract}, IValidatorTransactor: IValidatorTransactor{contract: contract}, IValidatorFilterer: IValidatorFilterer{contract: contract}}, nil
}

// NewIValidatorCaller creates a new read-only instance of IValidator, bound to a specific deployed contract.
func NewIValidatorCaller(address common.Address, caller bind.ContractCaller) (*IValidatorCaller, error) {
	contract, err := bindIValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IValidatorCaller{contract: contract}, nil
}

// NewIValidatorTransactor creates a new write-only instance of IValidator, bound to a specific deployed contract.
func NewIValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IValidatorTransactor, error) {
	contract, err := bindIValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IValidatorTransactor{contract: contract}, nil
}

// NewIValidatorFilterer creates a new log filterer instance of IValidator, bound to a specific deployed contract.
func NewIValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IValidatorFilterer, error) {
	contract, err := bindIValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IValidatorFilterer{contract: contract}, nil
}

// bindIValidator binds a generic wrapper to an already deployed contract.
func bindIValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IValidator *IValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IValidator.Contract.IValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IValidator *IValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValidator.Contract.IValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IValidator *IValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IValidator.Contract.IValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IValidator *IValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IValidator *IValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IValidator *IValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IValidator.Contract.contract.Transact(opts, method, params...)
}

// CreateValidatorProposal is a paid mutator transaction binding the contract method 0x331ba913.
//
// Solidity: function createValidatorProposal(address proposedAddress, uint8 action) returns()
func (_IValidator *IValidatorTransactor) CreateValidatorProposal(opts *bind.TransactOpts, proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _IValidator.contract.Transact(opts, "createValidatorProposal", proposedAddress, action)
}

// CreateValidatorProposal is a paid mutator transaction binding the contract method 0x331ba913.
//
// Solidity: function createValidatorProposal(address proposedAddress, uint8 action) returns()
func (_IValidator *IValidatorSession) CreateValidatorProposal(proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _IValidator.Contract.CreateValidatorProposal(&_IValidator.TransactOpts, proposedAddress, action)
}

// CreateValidatorProposal is a paid mutator transaction binding the contract method 0x331ba913.
//
// Solidity: function createValidatorProposal(address proposedAddress, uint8 action) returns()
func (_IValidator *IValidatorTransactorSession) CreateValidatorProposal(proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _IValidator.Contract.CreateValidatorProposal(&_IValidator.TransactOpts, proposedAddress, action)
}

// CreateValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x7fa9aaaf.
//
// Solidity: function createValidatorThresholdProposal(uint256 proposedValue) returns()
func (_IValidator *IValidatorTransactor) CreateValidatorThresholdProposal(opts *bind.TransactOpts, proposedValue *big.Int) (*types.Transaction, error) {
	return _IValidator.contract.Transact(opts, "createValidatorThresholdProposal", proposedValue)
}

// CreateValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x7fa9aaaf.
//
// Solidity: function createValidatorThresholdProposal(uint256 proposedValue) returns()
func (_IValidator *IValidatorSession) CreateValidatorThresholdProposal(proposedValue *big.Int) (*types.Transaction, error) {
	return _IValidator.Contract.CreateValidatorThresholdProposal(&_IValidator.TransactOpts, proposedValue)
}

// CreateValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x7fa9aaaf.
//
// Solidity: function createValidatorThresholdProposal(uint256 proposedValue) returns()
func (_IValidator *IValidatorTransactorSession) CreateValidatorThresholdProposal(proposedValue *big.Int) (*types.Transaction, error) {
	return _IValidator.Contract.CreateValidatorThresholdProposal(&_IValidator.TransactOpts, proposedValue)
}

// GetTotalValidators is a paid mutator transaction binding the contract method 0x295912df.
//
// Solidity: function getTotalValidators() returns(uint256)
func (_IValidator *IValidatorTransactor) GetTotalValidators(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValidator.contract.Transact(opts, "getTotalValidators")
}

// GetTotalValidators is a paid mutator transaction binding the contract method 0x295912df.
//
// Solidity: function getTotalValidators() returns(uint256)
func (_IValidator *IValidatorSession) GetTotalValidators() (*types.Transaction, error) {
	return _IValidator.Contract.GetTotalValidators(&_IValidator.TransactOpts)
}

// GetTotalValidators is a paid mutator transaction binding the contract method 0x295912df.
//
// Solidity: function getTotalValidators() returns(uint256)
func (_IValidator *IValidatorTransactorSession) GetTotalValidators() (*types.Transaction, error) {
	return _IValidator.Contract.GetTotalValidators(&_IValidator.TransactOpts)
}

// IsValidator is a paid mutator transaction binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validatorAddress) returns(bool)
func (_IValidator *IValidatorTransactor) IsValidator(opts *bind.TransactOpts, validatorAddress common.Address) (*types.Transaction, error) {
	return _IValidator.contract.Transact(opts, "isValidator", validatorAddress)
}

// IsValidator is a paid mutator transaction binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validatorAddress) returns(bool)
func (_IValidator *IValidatorSession) IsValidator(validatorAddress common.Address) (*types.Transaction, error) {
	return _IValidator.Contract.IsValidator(&_IValidator.TransactOpts, validatorAddress)
}

// IsValidator is a paid mutator transaction binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validatorAddress) returns(bool)
func (_IValidator *IValidatorTransactorSession) IsValidator(validatorAddress common.Address) (*types.Transaction, error) {
	return _IValidator.Contract.IsValidator(&_IValidator.TransactOpts, validatorAddress)
}

// VoteValidatorProposal is a paid mutator transaction binding the contract method 0x62e73263.
//
// Solidity: function voteValidatorProposal(address proposedAddress, uint8 vote) returns()
func (_IValidator *IValidatorTransactor) VoteValidatorProposal(opts *bind.TransactOpts, proposedAddress common.Address, vote uint8) (*types.Transaction, error) {
	return _IValidator.contract.Transact(opts, "voteValidatorProposal", proposedAddress, vote)
}

// VoteValidatorProposal is a paid mutator transaction binding the contract method 0x62e73263.
//
// Solidity: function voteValidatorProposal(address proposedAddress, uint8 vote) returns()
func (_IValidator *IValidatorSession) VoteValidatorProposal(proposedAddress common.Address, vote uint8) (*types.Transaction, error) {
	return _IValidator.Contract.VoteValidatorProposal(&_IValidator.TransactOpts, proposedAddress, vote)
}

// VoteValidatorProposal is a paid mutator transaction binding the contract method 0x62e73263.
//
// Solidity: function voteValidatorProposal(address proposedAddress, uint8 vote) returns()
func (_IValidator *IValidatorTransactorSession) VoteValidatorProposal(proposedAddress common.Address, vote uint8) (*types.Transaction, error) {
	return _IValidator.Contract.VoteValidatorProposal(&_IValidator.TransactOpts, proposedAddress, vote)
}

// VoteValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x9bd61e6a.
//
// Solidity: function voteValidatorThresholdProposal(uint8 vote) returns()
func (_IValidator *IValidatorTransactor) VoteValidatorThresholdProposal(opts *bind.TransactOpts, vote uint8) (*types.Transaction, error) {
	return _IValidator.contract.Transact(opts, "voteValidatorThresholdProposal", vote)
}

// VoteValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x9bd61e6a.
//
// Solidity: function voteValidatorThresholdProposal(uint8 vote) returns()
func (_IValidator *IValidatorSession) VoteValidatorThresholdProposal(vote uint8) (*types.Transaction, error) {
	return _IValidator.Contract.VoteValidatorThresholdProposal(&_IValidator.TransactOpts, vote)
}

// VoteValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x9bd61e6a.
//
// Solidity: function voteValidatorThresholdProposal(uint8 vote) returns()
func (_IValidator *IValidatorTransactorSession) VoteValidatorThresholdProposal(vote uint8) (*types.Transaction, error) {
	return _IValidator.Contract.VoteValidatorThresholdProposal(&_IValidator.TransactOpts, vote)
}
