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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC721SafeABI is the input ABI used to generate the binding from.
const ERC721SafeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"fundERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721SafeBin is the compiled bytecode used for deploying new contracts.
var ERC721SafeBin = "0x608060405234801561001057600080fd5b506101b2806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80637354298014610030575b600080fd5b61009c6004803603606081101561004657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061009e565b005b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561015e57600080fd5b505af1158015610172573d6000803e3d6000fd5b505050505050505056fea2646970667358221220c351cbada1955bbf33a873f3a4583be8f2e6876cf6570d1b77e8da47941f3f0a64736f6c63430006040033"

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
func (_ERC721Safe *ERC721SafeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_ERC721Safe *ERC721SafeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// FundERC721 is a paid mutator transaction binding the contract method 0x73542980.
//
// Solidity: function fundERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_ERC721Safe *ERC721SafeTransactor) FundERC721(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Safe.contract.Transact(opts, "fundERC721", tokenAddress, owner, tokenID)
}

// FundERC721 is a paid mutator transaction binding the contract method 0x73542980.
//
// Solidity: function fundERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_ERC721Safe *ERC721SafeSession) FundERC721(tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Safe.Contract.FundERC721(&_ERC721Safe.TransactOpts, tokenAddress, owner, tokenID)
}

// FundERC721 is a paid mutator transaction binding the contract method 0x73542980.
//
// Solidity: function fundERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_ERC721Safe *ERC721SafeTransactorSession) FundERC721(tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Safe.Contract.FundERC721(&_ERC721Safe.TransactOpts, tokenAddress, owner, tokenID)
}
