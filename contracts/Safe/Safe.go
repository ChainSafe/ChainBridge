// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Safe

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

// SafeABI is the input ABI used to generate the binding from.
const SafeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"releaseErc\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"releaseNFT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SafeBin is the compiled bytecode used for deploying new contracts.
const SafeBin = `60806040523480156100115760006000fd5b506040516107e83803806107e8833981810160405260208110156100355760006000fd5b81019080805190602001909291905050505b80600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5061008f565b61074a8061009e6000396000f3fe60806040523480156100115760006000fd5b50600436106100515760003560e01c806327e235e3146100575780632b4380c5146100b0578063344557331461011f5780638da5cb5b1461018e57610051565b60006000fd5b61009a6004803603602081101561006e5760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506101d8565b6040518082815260200191505060405180910390f35b61011d600480360360608110156100c75760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506101f3565b005b61018c600480360360608110156101365760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610332565b005b6101966105c5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60016000506020528060005260406000206000915090505481565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156102505760006000fd5b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd3085856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b1580156103115760006000fd5b505af1158015610326573d600060003e3d6000fd5b50505050505b5b505050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561038f5760006000fd5b81600160005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050541015151561044f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f5769746864726177616c20616d6f756e7420697320746f6f206869676821000081526020015060200191505060405180910390fd5b6104a782600160005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546105f490919063ffffffff16565b600160005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005081909090555060008390508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561057f5760006000fd5b505af1158015610594573d600060003e3d6000fd5b505050506040513d60208110156105ab5760006000fd5b810190808051906020019092919050505050505b5b505050565b6000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506105f1565b90565b600061063f83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081526020015061064c63ffffffff16565b9050610646565b92915050565b600083831115829015156106fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156106c15780820151818401525b6020810190506106a5565b50505050905090810190601f1680156106ee5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083850390508091505061070e56505b939250505056fea265627a7a7231582075b65a19c7a201fced2b5d4de4067fd420909e24a5523ddd7cbf1d79eca6a8bf64736f6c634300050c0032`

// DeploySafe deploys a new Ethereum contract, binding an instance of Safe to it.
func DeploySafe(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *Safe, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Safe{SafeCaller: SafeCaller{contract: contract}, SafeTransactor: SafeTransactor{contract: contract}, SafeFilterer: SafeFilterer{contract: contract}}, nil
}

// Safe is an auto generated Go binding around an Ethereum contract.
type Safe struct {
	SafeCaller     // Read-only binding to the contract
	SafeTransactor // Write-only binding to the contract
	SafeFilterer   // Log filterer for contract events
}

// SafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeSession struct {
	Contract     *Safe             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCallerSession struct {
	Contract *SafeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeTransactorSession struct {
	Contract     *SafeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeRaw struct {
	Contract *Safe // Generic contract binding to access the raw methods on
}

// SafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCallerRaw struct {
	Contract *SafeCaller // Generic read-only contract binding to access the raw methods on
}

// SafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeTransactorRaw struct {
	Contract *SafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafe creates a new instance of Safe, bound to a specific deployed contract.
func NewSafe(address common.Address, backend bind.ContractBackend) (*Safe, error) {
	contract, err := bindSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Safe{SafeCaller: SafeCaller{contract: contract}, SafeTransactor: SafeTransactor{contract: contract}, SafeFilterer: SafeFilterer{contract: contract}}, nil
}

// NewSafeCaller creates a new read-only instance of Safe, bound to a specific deployed contract.
func NewSafeCaller(address common.Address, caller bind.ContractCaller) (*SafeCaller, error) {
	contract, err := bindSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCaller{contract: contract}, nil
}

// NewSafeTransactor creates a new write-only instance of Safe, bound to a specific deployed contract.
func NewSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeTransactor, error) {
	contract, err := bindSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeTransactor{contract: contract}, nil
}

// NewSafeFilterer creates a new log filterer instance of Safe, bound to a specific deployed contract.
func NewSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeFilterer, error) {
	contract, err := bindSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeFilterer{contract: contract}, nil
}

// bindSafe binds a generic wrapper to an already deployed contract.
func bindSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safe *SafeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Safe.Contract.SafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safe *SafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safe.Contract.SafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safe *SafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safe.Contract.SafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safe *SafeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Safe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safe *SafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safe *SafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safe.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_Safe *SafeCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_Safe *SafeSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Safe.Contract.Balances(&_Safe.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_Safe *SafeCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Safe.Contract.Balances(&_Safe.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Safe *SafeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Safe.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Safe *SafeSession) Owner() (common.Address, error) {
	return _Safe.Contract.Owner(&_Safe.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Safe *SafeCallerSession) Owner() (common.Address, error) {
	return _Safe.Contract.Owner(&_Safe.CallOpts)
}

// ReleaseErc is a paid mutator transaction binding the contract method 0x34455733.
//
// Solidity: function releaseErc(_tokenAddress address, _value uint256, _to address) returns()
func (_Safe *SafeTransactor) ReleaseErc(opts *bind.TransactOpts, _tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "releaseErc", _tokenAddress, _value, _to)
}

// ReleaseErc is a paid mutator transaction binding the contract method 0x34455733.
//
// Solidity: function releaseErc(_tokenAddress address, _value uint256, _to address) returns()
func (_Safe *SafeSession) ReleaseErc(_tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Safe.Contract.ReleaseErc(&_Safe.TransactOpts, _tokenAddress, _value, _to)
}

// ReleaseErc is a paid mutator transaction binding the contract method 0x34455733.
//
// Solidity: function releaseErc(_tokenAddress address, _value uint256, _to address) returns()
func (_Safe *SafeTransactorSession) ReleaseErc(_tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Safe.Contract.ReleaseErc(&_Safe.TransactOpts, _tokenAddress, _value, _to)
}

// ReleaseNFT is a paid mutator transaction binding the contract method 0x2b4380c5.
//
// Solidity: function releaseNFT(_tokenAddress address, _to address, _tokenId uint256) returns()
func (_Safe *SafeTransactor) ReleaseNFT(opts *bind.TransactOpts, _tokenAddress common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "releaseNFT", _tokenAddress, _to, _tokenId)
}

// ReleaseNFT is a paid mutator transaction binding the contract method 0x2b4380c5.
//
// Solidity: function releaseNFT(_tokenAddress address, _to address, _tokenId uint256) returns()
func (_Safe *SafeSession) ReleaseNFT(_tokenAddress common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Safe.Contract.ReleaseNFT(&_Safe.TransactOpts, _tokenAddress, _to, _tokenId)
}

// ReleaseNFT is a paid mutator transaction binding the contract method 0x2b4380c5.
//
// Solidity: function releaseNFT(_tokenAddress address, _to address, _tokenId uint256) returns()
func (_Safe *SafeTransactorSession) ReleaseNFT(_tokenAddress common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Safe.Contract.ReleaseNFT(&_Safe.TransactOpts, _tokenAddress, _to, _tokenId)
}
