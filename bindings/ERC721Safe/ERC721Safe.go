// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC721Safe

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

// ERC721SafeABI is the input ABI used to generate the binding from.
const ERC721SafeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC721SafeBin is the compiled bytecode used for deploying new contracts.
var ERC721SafeBin = "0x608060405234801561001057600080fd5b5060d08061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80636ebcf60714602d575b600080fd5b606c60048036036020811015604157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506082565b6040518082815260200191505060405180910390f35b6000602052806000526040600020600091509050548156fea2646970667358221220dd637b610bffaebe18c9a9961d034f718c71dbe4a1f6795efab196f5ea93796064736f6c63430006040033"

// DeployERC721Safe deploys a new Ethereum contract, binding an instance of ERC721Safe to it.
func DeployERC721Safe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Safe, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721SafeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721SafeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Safe{ERC721SafeCaller: ERC721SafeCaller{contract: contract}, ERC721SafeTransactor: ERC721SafeTransactor{contract: contract}, ERC721SafeFilterer: ERC721SafeFilterer{contract: contract}}, nil
}

// ERC721Safe is an auto generated Go binding around an Ethereum contract.
type ERC721Safe struct {
	ERC721SafeCaller     // Read-only binding to the contract
	ERC721SafeTransactor // Write-only binding to the contract
	ERC721SafeFilterer   // Log filterer for contract events
}

// ERC721SafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721SafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721SafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721SafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721SafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721SafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721SafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721SafeSession struct {
	Contract     *ERC721Safe       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721SafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721SafeCallerSession struct {
	Contract *ERC721SafeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC721SafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721SafeTransactorSession struct {
	Contract     *ERC721SafeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC721SafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721SafeRaw struct {
	Contract *ERC721Safe // Generic contract binding to access the raw methods on
}

// ERC721SafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721SafeCallerRaw struct {
	Contract *ERC721SafeCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721SafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721SafeTransactorRaw struct {
	Contract *ERC721SafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Safe creates a new instance of ERC721Safe, bound to a specific deployed contract.
func NewERC721Safe(address common.Address, backend bind.ContractBackend) (*ERC721Safe, error) {
	contract, err := bindERC721Safe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Safe{ERC721SafeCaller: ERC721SafeCaller{contract: contract}, ERC721SafeTransactor: ERC721SafeTransactor{contract: contract}, ERC721SafeFilterer: ERC721SafeFilterer{contract: contract}}, nil
}

// NewERC721SafeCaller creates a new read-only instance of ERC721Safe, bound to a specific deployed contract.
func NewERC721SafeCaller(address common.Address, caller bind.ContractCaller) (*ERC721SafeCaller, error) {
	contract, err := bindERC721Safe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721SafeCaller{contract: contract}, nil
}

// NewERC721SafeTransactor creates a new write-only instance of ERC721Safe, bound to a specific deployed contract.
func NewERC721SafeTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721SafeTransactor, error) {
	contract, err := bindERC721Safe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721SafeTransactor{contract: contract}, nil
}

// NewERC721SafeFilterer creates a new log filterer instance of ERC721Safe, bound to a specific deployed contract.
func NewERC721SafeFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721SafeFilterer, error) {
	contract, err := bindERC721Safe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721SafeFilterer{contract: contract}, nil
}

// bindERC721Safe binds a generic wrapper to an already deployed contract.
func bindERC721Safe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721SafeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Safe *ERC721SafeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Safe.Contract.ERC721SafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Safe *ERC721SafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Safe.Contract.ERC721SafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Safe *ERC721SafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Safe.Contract.ERC721SafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Safe *ERC721SafeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Safe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Safe *ERC721SafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Safe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Safe *ERC721SafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Safe.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC721Safe *ERC721SafeCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Safe.contract.Call(opts, out, "_balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC721Safe *ERC721SafeSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC721Safe.Contract.Balances(&_ERC721Safe.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC721Safe *ERC721SafeCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC721Safe.Contract.Balances(&_ERC721Safe.CallOpts, arg0)
}

var RuntimeBytecode = "0x6080604052348015600f57600080fd5b506004361060285760003560e01c80636ebcf60714602d575b600080fd5b606c60048036036020811015604157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506082565b6040518082815260200191505060405180910390f35b6000602052806000526040600020600091509050548156fea2646970667358221220dd637b610bffaebe18c9a9961d034f718c71dbe4a1f6795efab196f5ea93796064736f6c63430006040033"
