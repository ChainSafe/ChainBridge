// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IERC20Handler

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

// IERC20HandlerABI is the input ABI used to generate the binding from.
const IERC20HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20Handler is an auto generated Go binding around an Ethereum contract.
type IERC20Handler struct {
	IERC20HandlerCaller     // Read-only binding to the contract
	IERC20HandlerTransactor // Write-only binding to the contract
	IERC20HandlerFilterer   // Log filterer for contract events
}

// IERC20HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20HandlerSession struct {
	Contract     *IERC20Handler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20HandlerCallerSession struct {
	Contract *IERC20HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IERC20HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20HandlerTransactorSession struct {
	Contract     *IERC20HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IERC20HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20HandlerRaw struct {
	Contract *IERC20Handler // Generic contract binding to access the raw methods on
}

// IERC20HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20HandlerCallerRaw struct {
	Contract *IERC20HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20HandlerTransactorRaw struct {
	Contract *IERC20HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Handler creates a new instance of IERC20Handler, bound to a specific deployed contract.
func NewIERC20Handler(address common.Address, backend bind.ContractBackend) (*IERC20Handler, error) {
	contract, err := bindIERC20Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Handler{IERC20HandlerCaller: IERC20HandlerCaller{contract: contract}, IERC20HandlerTransactor: IERC20HandlerTransactor{contract: contract}, IERC20HandlerFilterer: IERC20HandlerFilterer{contract: contract}}, nil
}

// NewIERC20HandlerCaller creates a new read-only instance of IERC20Handler, bound to a specific deployed contract.
func NewIERC20HandlerCaller(address common.Address, caller bind.ContractCaller) (*IERC20HandlerCaller, error) {
	contract, err := bindIERC20Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20HandlerCaller{contract: contract}, nil
}

// NewIERC20HandlerTransactor creates a new write-only instance of IERC20Handler, bound to a specific deployed contract.
func NewIERC20HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20HandlerTransactor, error) {
	contract, err := bindIERC20Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20HandlerTransactor{contract: contract}, nil
}

// NewIERC20HandlerFilterer creates a new log filterer instance of IERC20Handler, bound to a specific deployed contract.
func NewIERC20HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20HandlerFilterer, error) {
	contract, err := bindIERC20Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20HandlerFilterer{contract: contract}, nil
}

// bindIERC20Handler binds a generic wrapper to an already deployed contract.
func bindIERC20Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Handler *IERC20HandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20Handler.Contract.IERC20HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Handler *IERC20HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Handler.Contract.IERC20HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Handler *IERC20HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Handler.Contract.IERC20HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Handler *IERC20HandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Handler *IERC20HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Handler *IERC20HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Handler.Contract.contract.Transact(opts, method, params...)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x1cad5a40.
//
// Solidity: function depositERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_IERC20Handler *IERC20HandlerTransactor) DepositERC20(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Handler.contract.Transact(opts, "depositERC20", tokenAddress, owner, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x1cad5a40.
//
// Solidity: function depositERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_IERC20Handler *IERC20HandlerSession) DepositERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Handler.Contract.DepositERC20(&_IERC20Handler.TransactOpts, tokenAddress, owner, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x1cad5a40.
//
// Solidity: function depositERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_IERC20Handler *IERC20HandlerTransactorSession) DepositERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Handler.Contract.DepositERC20(&_IERC20Handler.TransactOpts, tokenAddress, owner, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address tokenAddress, address recipient, uint256 amount) returns()
func (_IERC20Handler *IERC20HandlerTransactor) WithdrawERC20(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Handler.contract.Transact(opts, "withdrawERC20", tokenAddress, recipient, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address tokenAddress, address recipient, uint256 amount) returns()
func (_IERC20Handler *IERC20HandlerSession) WithdrawERC20(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Handler.Contract.WithdrawERC20(&_IERC20Handler.TransactOpts, tokenAddress, recipient, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address tokenAddress, address recipient, uint256 amount) returns()
func (_IERC20Handler *IERC20HandlerTransactorSession) WithdrawERC20(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Handler.Contract.WithdrawERC20(&_IERC20Handler.TransactOpts, tokenAddress, recipient, amount)
}
var RuntimeBytecode = "0x"
