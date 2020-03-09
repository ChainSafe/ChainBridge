// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IERC721Handler

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

// IERC721HandlerABI is the input ABI used to generate the binding from.
const IERC721HandlerABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"depositERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721Handler is an auto generated Go binding around an Ethereum contract.
type IERC721Handler struct {
	IERC721HandlerCaller     // Read-only binding to the contract
	IERC721HandlerTransactor // Write-only binding to the contract
	IERC721HandlerFilterer   // Log filterer for contract events
}

// IERC721HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721HandlerSession struct {
	Contract     *IERC721Handler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721HandlerCallerSession struct {
	Contract *IERC721HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC721HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721HandlerTransactorSession struct {
	Contract     *IERC721HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC721HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721HandlerRaw struct {
	Contract *IERC721Handler // Generic contract binding to access the raw methods on
}

// IERC721HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721HandlerCallerRaw struct {
	Contract *IERC721HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721HandlerTransactorRaw struct {
	Contract *IERC721HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Handler creates a new instance of IERC721Handler, bound to a specific deployed contract.
func NewIERC721Handler(address common.Address, backend bind.ContractBackend) (*IERC721Handler, error) {
	contract, err := bindIERC721Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Handler{IERC721HandlerCaller: IERC721HandlerCaller{contract: contract}, IERC721HandlerTransactor: IERC721HandlerTransactor{contract: contract}, IERC721HandlerFilterer: IERC721HandlerFilterer{contract: contract}}, nil
}

// NewIERC721HandlerCaller creates a new read-only instance of IERC721Handler, bound to a specific deployed contract.
func NewIERC721HandlerCaller(address common.Address, caller bind.ContractCaller) (*IERC721HandlerCaller, error) {
	contract, err := bindIERC721Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721HandlerCaller{contract: contract}, nil
}

// NewIERC721HandlerTransactor creates a new write-only instance of IERC721Handler, bound to a specific deployed contract.
func NewIERC721HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721HandlerTransactor, error) {
	contract, err := bindIERC721Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721HandlerTransactor{contract: contract}, nil
}

// NewIERC721HandlerFilterer creates a new log filterer instance of IERC721Handler, bound to a specific deployed contract.
func NewIERC721HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721HandlerFilterer, error) {
	contract, err := bindIERC721Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721HandlerFilterer{contract: contract}, nil
}

// bindIERC721Handler binds a generic wrapper to an already deployed contract.
func bindIERC721Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Handler *IERC721HandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721Handler.Contract.IERC721HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Handler *IERC721HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Handler.Contract.IERC721HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Handler *IERC721HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Handler.Contract.IERC721HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Handler *IERC721HandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Handler *IERC721HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Handler *IERC721HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Handler.Contract.contract.Transact(opts, method, params...)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x331ded1a.
//
// Solidity: function depositERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_IERC721Handler *IERC721HandlerTransactor) DepositERC721(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _IERC721Handler.contract.Transact(opts, "depositERC721", tokenAddress, owner, tokenID)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x331ded1a.
//
// Solidity: function depositERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_IERC721Handler *IERC721HandlerSession) DepositERC721(tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _IERC721Handler.Contract.DepositERC721(&_IERC721Handler.TransactOpts, tokenAddress, owner, tokenID)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x331ded1a.
//
// Solidity: function depositERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_IERC721Handler *IERC721HandlerTransactorSession) DepositERC721(tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _IERC721Handler.Contract.DepositERC721(&_IERC721Handler.TransactOpts, tokenAddress, owner, tokenID)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x4025feb2.
//
// Solidity: function withdrawERC721(address tokenAddress, address recipient, uint256 tokenID) returns()
func (_IERC721Handler *IERC721HandlerTransactor) WithdrawERC721(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _IERC721Handler.contract.Transact(opts, "withdrawERC721", tokenAddress, recipient, tokenID)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x4025feb2.
//
// Solidity: function withdrawERC721(address tokenAddress, address recipient, uint256 tokenID) returns()
func (_IERC721Handler *IERC721HandlerSession) WithdrawERC721(tokenAddress common.Address, recipient common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _IERC721Handler.Contract.WithdrawERC721(&_IERC721Handler.TransactOpts, tokenAddress, recipient, tokenID)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x4025feb2.
//
// Solidity: function withdrawERC721(address tokenAddress, address recipient, uint256 tokenID) returns()
func (_IERC721Handler *IERC721HandlerTransactorSession) WithdrawERC721(tokenAddress common.Address, recipient common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _IERC721Handler.Contract.WithdrawERC721(&_IERC721Handler.TransactOpts, tokenAddress, recipient, tokenID)
}
