// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20Safe

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

// ERC20SafeABI is the input ABI used to generate the binding from.
const ERC20SafeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC20SafeBin is the compiled bytecode used for deploying new contracts.
var ERC20SafeBin = "0x608060405234801561001057600080fd5b5060d08061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80636ebcf60714602d575b600080fd5b606c60048036036020811015604157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506082565b6040518082815260200191505060405180910390f35b6000602052806000526040600020600091509050548156fea264697066735822122092588d61635b9221c0383224c3632aa6e05c27ce1c2dde88884312fb44ca361164736f6c63430006040033"

// DeployERC20Safe deploys a new Ethereum contract, binding an instance of ERC20Safe to it.
func DeployERC20Safe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Safe, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20SafeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20SafeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Safe{ERC20SafeCaller: ERC20SafeCaller{contract: contract}, ERC20SafeTransactor: ERC20SafeTransactor{contract: contract}, ERC20SafeFilterer: ERC20SafeFilterer{contract: contract}}, nil
}

// ERC20Safe is an auto generated Go binding around an Ethereum contract.
type ERC20Safe struct {
	ERC20SafeCaller     // Read-only binding to the contract
	ERC20SafeTransactor // Write-only binding to the contract
	ERC20SafeFilterer   // Log filterer for contract events
}

// ERC20SafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20SafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20SafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20SafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20SafeSession struct {
	Contract     *ERC20Safe        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20SafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20SafeCallerSession struct {
	Contract *ERC20SafeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ERC20SafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20SafeTransactorSession struct {
	Contract     *ERC20SafeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERC20SafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20SafeRaw struct {
	Contract *ERC20Safe // Generic contract binding to access the raw methods on
}

// ERC20SafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20SafeCallerRaw struct {
	Contract *ERC20SafeCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20SafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20SafeTransactorRaw struct {
	Contract *ERC20SafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Safe creates a new instance of ERC20Safe, bound to a specific deployed contract.
func NewERC20Safe(address common.Address, backend bind.ContractBackend) (*ERC20Safe, error) {
	contract, err := bindERC20Safe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Safe{ERC20SafeCaller: ERC20SafeCaller{contract: contract}, ERC20SafeTransactor: ERC20SafeTransactor{contract: contract}, ERC20SafeFilterer: ERC20SafeFilterer{contract: contract}}, nil
}

// NewERC20SafeCaller creates a new read-only instance of ERC20Safe, bound to a specific deployed contract.
func NewERC20SafeCaller(address common.Address, caller bind.ContractCaller) (*ERC20SafeCaller, error) {
	contract, err := bindERC20Safe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20SafeCaller{contract: contract}, nil
}

// NewERC20SafeTransactor creates a new write-only instance of ERC20Safe, bound to a specific deployed contract.
func NewERC20SafeTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20SafeTransactor, error) {
	contract, err := bindERC20Safe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20SafeTransactor{contract: contract}, nil
}

// NewERC20SafeFilterer creates a new log filterer instance of ERC20Safe, bound to a specific deployed contract.
func NewERC20SafeFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20SafeFilterer, error) {
	contract, err := bindERC20Safe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20SafeFilterer{contract: contract}, nil
}

// bindERC20Safe binds a generic wrapper to an already deployed contract.
func bindERC20Safe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20SafeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Safe *ERC20SafeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Safe.Contract.ERC20SafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Safe *ERC20SafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Safe.Contract.ERC20SafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Safe *ERC20SafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Safe.Contract.ERC20SafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Safe *ERC20SafeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Safe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Safe *ERC20SafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Safe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Safe *ERC20SafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Safe.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC20Safe *ERC20SafeCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Safe.contract.Call(opts, out, "_balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC20Safe *ERC20SafeSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC20Safe.Contract.Balances(&_ERC20Safe.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC20Safe *ERC20SafeCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC20Safe.Contract.Balances(&_ERC20Safe.CallOpts, arg0)
}
var RuntimeBytecode = "0x6080604052348015600f57600080fd5b506004361060285760003560e01c80636ebcf60714602d575b600080fd5b606c60048036036020811015604157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506082565b6040518082815260200191505060405180910390f35b6000602052806000526040600020600091509050548156fea264697066735822122092588d61635b9221c0383224c3632aa6e05c27ce1c2dde88884312fb44ca361164736f6c63430006040033"
