// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package MinterRole

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MinterRoleABI is the input ABI used to generate the binding from.
const MinterRoleABI = `[{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"}],"name":"MinterAdded","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"}],"name":"MinterRemoved","type":"event"},{"constant":true,"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"isMinter","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"addMinter","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"renounceMinter","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`

// MinterRole is an auto generated Go binding around an Ethereum contract.
type MinterRole struct {
	MinterRoleCaller     // Read-only binding to the contract
	MinterRoleTransactor // Write-only binding to the contract
}

// MinterRoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinterRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinterRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinterRoleSession struct {
	Contract     *MinterRole       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinterRoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinterRoleCallerSession struct {
	Contract *MinterRoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MinterRoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinterRoleTransactorSession struct {
	Contract     *MinterRoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MinterRoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinterRoleRaw struct {
	Contract *MinterRole // Generic contract binding to access the raw methods on
}

// MinterRoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinterRoleCallerRaw struct {
	Contract *MinterRoleCaller // Generic read-only contract binding to access the raw methods on
}

// MinterRoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinterRoleTransactorRaw struct {
	Contract *MinterRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinterRole creates a new instance of MinterRole, bound to a specific deployed contract.
func NewMinterRole(address common.Address, backend bind.ContractBackend) (*MinterRole, error) {
	contract, err := bindMinterRole(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &MinterRole{MinterRoleCaller: MinterRoleCaller{contract: contract}, MinterRoleTransactor: MinterRoleTransactor{contract: contract}}, nil
}

// NewMinterRoleCaller creates a new read-only instance of MinterRole, bound to a specific deployed contract.
func NewMinterRoleCaller(address common.Address, caller bind.ContractCaller) (*MinterRoleCaller, error) {
	contract, err := bindMinterRole(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MinterRoleCaller{contract: contract}, nil
}

// NewMinterRoleTransactor creates a new write-only instance of MinterRole, bound to a specific deployed contract.
func NewMinterRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*MinterRoleTransactor, error) {
	contract, err := bindMinterRole(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MinterRoleTransactor{contract: contract}, nil
}

// bindMinterRole binds a generic wrapper to an already deployed contract.
func bindMinterRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinterRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterRole *MinterRoleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinterRole.Contract.MinterRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterRole *MinterRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRole.Contract.MinterRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterRole *MinterRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterRole.Contract.MinterRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterRole *MinterRoleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinterRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterRole *MinterRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterRole *MinterRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterRole.Contract.contract.Transact(opts, method, params...)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(account address) constant returns(bool)
func (_MinterRole *MinterRoleCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MinterRole.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(account address) constant returns(bool)
func (_MinterRole *MinterRoleSession) IsMinter(account common.Address) (bool, error) {
	return _MinterRole.Contract.IsMinter(&_MinterRole.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(account address) constant returns(bool)
func (_MinterRole *MinterRoleCallerSession) IsMinter(account common.Address) (bool, error) {
	return _MinterRole.Contract.IsMinter(&_MinterRole.CallOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(account address) returns()
func (_MinterRole *MinterRoleTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MinterRole.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(account address) returns()
func (_MinterRole *MinterRoleSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _MinterRole.Contract.AddMinter(&_MinterRole.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(account address) returns()
func (_MinterRole *MinterRoleTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _MinterRole.Contract.AddMinter(&_MinterRole.TransactOpts, account)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRole.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleSession) RenounceMinter() (*types.Transaction, error) {
	return _MinterRole.Contract.RenounceMinter(&_MinterRole.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _MinterRole.Contract.RenounceMinter(&_MinterRole.TransactOpts)
}
