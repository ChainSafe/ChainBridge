// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package ISafe

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ISafeABI is the input ABI used to generate the binding from.
const ISafeABI = `[{"constant":false,"inputs":[{"internalType":"address","name":"_tokenAddress","type":"address"},{"internalType":"uint256","name":"_value","type":"uint256"},{"internalType":"address","name":"_to","type":"address"}],"name":"releaseErc","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"_tokenAddress","type":"address"},{"internalType":"address","name":"_to","type":"address"},{"internalType":"uint256","name":"_tokenId","type":"uint256"}],"name":"releaseNFT","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`

// ISafe is an auto generated Go binding around an Ethereum contract.
type ISafe struct {
	ISafeCaller     // Read-only binding to the contract
	ISafeTransactor // Write-only binding to the contract
}

// ISafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISafeSession struct {
	Contract     *ISafe            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISafeCallerSession struct {
	Contract *ISafeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ISafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISafeTransactorSession struct {
	Contract     *ISafeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISafeRaw struct {
	Contract *ISafe // Generic contract binding to access the raw methods on
}

// ISafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISafeCallerRaw struct {
	Contract *ISafeCaller // Generic read-only contract binding to access the raw methods on
}

// ISafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISafeTransactorRaw struct {
	Contract *ISafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISafe creates a new instance of ISafe, bound to a specific deployed contract.
func NewISafe(address common.Address, backend bind.ContractBackend) (*ISafe, error) {
	contract, err := bindISafe(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &ISafe{ISafeCaller: ISafeCaller{contract: contract}, ISafeTransactor: ISafeTransactor{contract: contract}}, nil
}

// NewISafeCaller creates a new read-only instance of ISafe, bound to a specific deployed contract.
func NewISafeCaller(address common.Address, caller bind.ContractCaller) (*ISafeCaller, error) {
	contract, err := bindISafe(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ISafeCaller{contract: contract}, nil
}

// NewISafeTransactor creates a new write-only instance of ISafe, bound to a specific deployed contract.
func NewISafeTransactor(address common.Address, transactor bind.ContractTransactor) (*ISafeTransactor, error) {
	contract, err := bindISafe(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ISafeTransactor{contract: contract}, nil
}

// bindISafe binds a generic wrapper to an already deployed contract.
func bindISafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISafeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISafe *ISafeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISafe.Contract.ISafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISafe *ISafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISafe.Contract.ISafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISafe *ISafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISafe.Contract.ISafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISafe *ISafeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISafe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISafe *ISafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISafe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISafe *ISafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISafe.Contract.contract.Transact(opts, method, params...)
}

// ReleaseErc is a paid mutator transaction binding the contract method 0x34455733.
//
// Solidity: function releaseErc(_tokenAddress address, _value uint256, _to address) returns()
func (_ISafe *ISafeTransactor) ReleaseErc(opts *bind.TransactOpts, _tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ISafe.contract.Transact(opts, "releaseErc", _tokenAddress, _value, _to)
}

// ReleaseErc is a paid mutator transaction binding the contract method 0x34455733.
//
// Solidity: function releaseErc(_tokenAddress address, _value uint256, _to address) returns()
func (_ISafe *ISafeSession) ReleaseErc(_tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ISafe.Contract.ReleaseErc(&_ISafe.TransactOpts, _tokenAddress, _value, _to)
}

// ReleaseErc is a paid mutator transaction binding the contract method 0x34455733.
//
// Solidity: function releaseErc(_tokenAddress address, _value uint256, _to address) returns()
func (_ISafe *ISafeTransactorSession) ReleaseErc(_tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ISafe.Contract.ReleaseErc(&_ISafe.TransactOpts, _tokenAddress, _value, _to)
}

// ReleaseNFT is a paid mutator transaction binding the contract method 0x2b4380c5.
//
// Solidity: function releaseNFT(_tokenAddress address, _to address, _tokenId uint256) returns()
func (_ISafe *ISafeTransactor) ReleaseNFT(opts *bind.TransactOpts, _tokenAddress common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ISafe.contract.Transact(opts, "releaseNFT", _tokenAddress, _to, _tokenId)
}

// ReleaseNFT is a paid mutator transaction binding the contract method 0x2b4380c5.
//
// Solidity: function releaseNFT(_tokenAddress address, _to address, _tokenId uint256) returns()
func (_ISafe *ISafeSession) ReleaseNFT(_tokenAddress common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ISafe.Contract.ReleaseNFT(&_ISafe.TransactOpts, _tokenAddress, _to, _tokenId)
}

// ReleaseNFT is a paid mutator transaction binding the contract method 0x2b4380c5.
//
// Solidity: function releaseNFT(_tokenAddress address, _to address, _tokenId uint256) returns()
func (_ISafe *ISafeTransactorSession) ReleaseNFT(_tokenAddress common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ISafe.Contract.ReleaseNFT(&_ISafe.TransactOpts, _tokenAddress, _to, _tokenId)
}
