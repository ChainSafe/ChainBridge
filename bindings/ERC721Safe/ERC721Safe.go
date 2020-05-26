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
const ERC721SafeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"fundERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721SafeBin is the compiled bytecode used for deploying new contracts.
var ERC721SafeBin = "0x608060405234801561001057600080fd5b506105bb806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806345274738146100465780636ebcf6071461009e57806373542980146100f6575b600080fd5b6100886004803603602081101561005c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610164565b6040518082815260200191505060405180910390f35b6100e0600480360360208110156100b457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061017c565b6040518082815260200191505060405180910390f35b6101626004803603606081101561010c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610194565b005b60016020528060005260406000206000915090505481565b60006020528060005260406000206000915090505481565b60008390506101a58184308561023f565b6101f760016000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546104d890919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b600060608573ffffffffffffffffffffffffffffffffffffffff166323b872dd60e01b868686604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040518082805190602001908083835b6020831061036b5780518252602082019150602081019050602083039250610348565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146103cd576040519150601f19603f3d011682016040523d82523d6000602084013e6103d2565b606091505b50915091508161044a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4552433732313a207472616e736665722066726f6d206661696c65640000000081525060200191505060405180910390fd5b6000815111156104d05780806020019051602081101561046957600080fd5b81019080805190602001909291905050506104cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806105616025913960400191505060405180910390fd5b5b505050505050565b600080828401905083811015610556576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b809150509291505056fe4552433732313a207472616e736665722066726f6d20646964206e6f742073756363656564a26469706673582212204491835fa4ed7b0c112a7bd0fbbdbe930d03a1f06665554e3d3769217bc70c6864736f6c63430006040033"

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
// Solidity: function _balances(address ) view returns(uint256)
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
// Solidity: function _balances(address ) view returns(uint256)
func (_ERC721Safe *ERC721SafeSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC721Safe.Contract.Balances(&_ERC721Safe.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) view returns(uint256)
func (_ERC721Safe *ERC721SafeCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC721Safe.Contract.Balances(&_ERC721Safe.CallOpts, arg0)
}

// BurnedTokens is a free data retrieval call binding the contract method 0x45274738.
//
// Solidity: function _burnedTokens(address ) view returns(uint256)
func (_ERC721Safe *ERC721SafeCaller) BurnedTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Safe.contract.Call(opts, out, "_burnedTokens", arg0)
	return *ret0, err
}

// BurnedTokens is a free data retrieval call binding the contract method 0x45274738.
//
// Solidity: function _burnedTokens(address ) view returns(uint256)
func (_ERC721Safe *ERC721SafeSession) BurnedTokens(arg0 common.Address) (*big.Int, error) {
	return _ERC721Safe.Contract.BurnedTokens(&_ERC721Safe.CallOpts, arg0)
}

// BurnedTokens is a free data retrieval call binding the contract method 0x45274738.
//
// Solidity: function _burnedTokens(address ) view returns(uint256)
func (_ERC721Safe *ERC721SafeCallerSession) BurnedTokens(arg0 common.Address) (*big.Int, error) {
	return _ERC721Safe.Contract.BurnedTokens(&_ERC721Safe.CallOpts, arg0)
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
