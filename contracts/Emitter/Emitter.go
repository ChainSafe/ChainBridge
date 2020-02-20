// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Emitter

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

// EmitterABI is the input ABI used to generate the binding from.
const EmitterABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_destChain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"ERCTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_destChain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"GenericTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_destChain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"NFTTransfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_destChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_destChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"depositGenericErc\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_destChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"name\":\"depositNFT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"releaseErc\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"releaseNFT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EmitterBin is the compiled bytecode used for deploying new contracts.
const EmitterBin = `60806040523480156100115760006000fd5b505b305b80600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b505b61005e565b6110af8061006d6000396000f3fe60806040523480156100115760006000fd5b50600436106100825760003560e01c8063344557331161005c57806334455733146101c95780638da5cb5b14610238578063faa9bce914610282578063fcf861651461036f57610082565b80631d4241dd1461008857806327e235e3146101015780632b4380c51461015a57610082565b60006000fd5b6100ff6004803603608081101561009f5760006000fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610486565b005b610144600480360360208110156101185760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610589565b6040518082815260200191505060405180910390f35b6101c7600480360360608110156101715760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506105a4565b005b610236600480360360608110156101e05760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106e3565b005b610240610976565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61036d600480360360608110156102995760006000fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156102e15760006000fd5b8201836020820111156102f45760006000fd5b803590602001918460018302840111640100000000831117156103175760006000fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509090919290909192905050506109a5565b005b610484600480360360a08110156103865760006000fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156103f85760006000fd5b82018360208201111561040b5760006000fd5b8035906020019184600183028401116401000000008311171561042e5760006000fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050909091929090919290505050610ac7565b005b6002600050600085815260200190815260200160002060008181505480929190600101919050909055506000600260005060008681526020019081526020016000206000505490506104e082853033610c3963ffffffff16565b80857f12a20368a5b6a9eb35068693542bef5613d39588e50c62fca6312f1b7d9242fe858786604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001935050505060405180910390a3505b50505050565b60016000506020528060005260406000206000915090505481565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156106015760006000fd5b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd3085856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b1580156106c25760006000fd5b505af11580156106d7573d600060003e3d6000fd5b50505050505b5b505050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156107405760006000fd5b81600160005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505410151515610800576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f5769746864726177616c20616d6f756e7420697320746f6f206869676821000081526020015060200191505060405180910390fd5b61085882600160005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005054610de390919063ffffffff16565b600160005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005081909090555060008390508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156109305760006000fd5b505af1158015610945573d600060003e3d6000fd5b505050506040513d602081101561095c5760006000fd5b810190808051906020019092919050505050505b5b505050565b6000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506109a2565b90565b60026000506000848152602001908152602001600020600081815054809291906001019190509090555060006002600050600085815260200190815260200160002060005054905080847f6cc01912ea1b73308485cfdf6b52b3a8c1e8f33825352b5d5fc0b5ddec3a2a2b8585604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610a855780820151818401525b602081019050610a69565b50505050905090810190601f168015610ab25780820380516001836020036101000a031916815260200191505b50935050505060405180910390a3505b505050565b600260005060008681526020019081526020016000206000818150548092919060010191905090905550600060026000506000878152602001908152602001600020600050549050610b2184303386610e3b63ffffffff16565b80867f0e6815b49056b27f8b1a0f27453bdd5caedf0eb14fa1e1c1f4f11be9cba05f0587878787604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610bf35780820151818401525b602081019050610bd7565b50505050905090810190601f168015610c205780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a3505b5050505050565b610c9183600160005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005054610f1d90919063ffffffff16565b600160005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005081909090555060008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8385876040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015610d9d5760006000fd5b505af1158015610db2573d600060003e3d6000fd5b505050506040513d6020811015610dc95760006000fd5b810190808051906020019092919050505050505b50505050565b6000610e2e83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815260200150610fb163ffffffff16565b9050610e35565b92915050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8486856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015610efc5760006000fd5b505af1158015610f11573d600060003e3d6000fd5b50505050505b50505050565b600060008284019050838110151515610fa1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081526020015060200191505060405180910390fd5b80915050610fab56505b92915050565b60008383111582901515611061576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156110265780820151818401525b60208101905061100a565b50505050905090810190601f1680156110535780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083850390508091505061107356505b939250505056fea265627a7a72315820965ab29e0fa65dd28497f313fb95430a53217986dc3d000299defe3ed99629d864736f6c634300050c0032`

// DeployEmitter deploys a new Ethereum contract, binding an instance of Emitter to it.
func DeployEmitter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Emitter, error) {
	parsed, err := abi.JSON(strings.NewReader(EmitterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EmitterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Emitter{EmitterCaller: EmitterCaller{contract: contract}, EmitterTransactor: EmitterTransactor{contract: contract}, EmitterFilterer: EmitterFilterer{contract: contract}}, nil
}

// Emitter is an auto generated Go binding around an Ethereum contract.
type Emitter struct {
	EmitterCaller     // Read-only binding to the contract
	EmitterTransactor // Write-only binding to the contract
	EmitterFilterer   // Log filterer for contract events
}

// EmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type EmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EmitterSession struct {
	Contract     *Emitter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EmitterCallerSession struct {
	Contract *EmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EmitterTransactorSession struct {
	Contract     *EmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type EmitterRaw struct {
	Contract *Emitter // Generic contract binding to access the raw methods on
}

// EmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EmitterCallerRaw struct {
	Contract *EmitterCaller // Generic read-only contract binding to access the raw methods on
}

// EmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EmitterTransactorRaw struct {
	Contract *EmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEmitter creates a new instance of Emitter, bound to a specific deployed contract.
func NewEmitter(address common.Address, backend bind.ContractBackend) (*Emitter, error) {
	contract, err := bindEmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Emitter{EmitterCaller: EmitterCaller{contract: contract}, EmitterTransactor: EmitterTransactor{contract: contract}, EmitterFilterer: EmitterFilterer{contract: contract}}, nil
}

// NewEmitterCaller creates a new read-only instance of Emitter, bound to a specific deployed contract.
func NewEmitterCaller(address common.Address, caller bind.ContractCaller) (*EmitterCaller, error) {
	contract, err := bindEmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EmitterCaller{contract: contract}, nil
}

// NewEmitterTransactor creates a new write-only instance of Emitter, bound to a specific deployed contract.
func NewEmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*EmitterTransactor, error) {
	contract, err := bindEmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EmitterTransactor{contract: contract}, nil
}

// NewEmitterFilterer creates a new log filterer instance of Emitter, bound to a specific deployed contract.
func NewEmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*EmitterFilterer, error) {
	contract, err := bindEmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EmitterFilterer{contract: contract}, nil
}

// bindEmitter binds a generic wrapper to an already deployed contract.
func bindEmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EmitterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Emitter *EmitterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Emitter.Contract.EmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Emitter *EmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emitter.Contract.EmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Emitter *EmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Emitter.Contract.EmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Emitter *EmitterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Emitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Emitter *EmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Emitter *EmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Emitter.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_Emitter *EmitterCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Emitter.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_Emitter *EmitterSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Emitter.Contract.Balances(&_Emitter.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_Emitter *EmitterCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Emitter.Contract.Balances(&_Emitter.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Emitter *EmitterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Emitter.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Emitter *EmitterSession) Owner() (common.Address, error) {
	return _Emitter.Contract.Owner(&_Emitter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Emitter *EmitterCallerSession) Owner() (common.Address, error) {
	return _Emitter.Contract.Owner(&_Emitter.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xfaa9bce9.
//
// Solidity: function deposit(_destChain uint256, _destAddress address, _data bytes) returns()
func (_Emitter *EmitterTransactor) Deposit(opts *bind.TransactOpts, _destChain *big.Int, _destAddress common.Address, _data []byte) (*types.Transaction, error) {
	return _Emitter.contract.Transact(opts, "deposit", _destChain, _destAddress, _data)
}

// Deposit is a paid mutator transaction binding the contract method 0xfaa9bce9.
//
// Solidity: function deposit(_destChain uint256, _destAddress address, _data bytes) returns()
func (_Emitter *EmitterSession) Deposit(_destChain *big.Int, _destAddress common.Address, _data []byte) (*types.Transaction, error) {
	return _Emitter.Contract.Deposit(&_Emitter.TransactOpts, _destChain, _destAddress, _data)
}

// Deposit is a paid mutator transaction binding the contract method 0xfaa9bce9.
//
// Solidity: function deposit(_destChain uint256, _destAddress address, _data bytes) returns()
func (_Emitter *EmitterTransactorSession) Deposit(_destChain *big.Int, _destAddress common.Address, _data []byte) (*types.Transaction, error) {
	return _Emitter.Contract.Deposit(&_Emitter.TransactOpts, _destChain, _destAddress, _data)
}

// DepositGenericErc is a paid mutator transaction binding the contract method 0x1d4241dd.
//
// Solidity: function depositGenericErc(_destChain uint256, _value uint256, _to address, _tokenAddress address) returns()
func (_Emitter *EmitterTransactor) DepositGenericErc(opts *bind.TransactOpts, _destChain *big.Int, _value *big.Int, _to common.Address, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Emitter.contract.Transact(opts, "depositGenericErc", _destChain, _value, _to, _tokenAddress)
}

// DepositGenericErc is a paid mutator transaction binding the contract method 0x1d4241dd.
//
// Solidity: function depositGenericErc(_destChain uint256, _value uint256, _to address, _tokenAddress address) returns()
func (_Emitter *EmitterSession) DepositGenericErc(_destChain *big.Int, _value *big.Int, _to common.Address, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Emitter.Contract.DepositGenericErc(&_Emitter.TransactOpts, _destChain, _value, _to, _tokenAddress)
}

// DepositGenericErc is a paid mutator transaction binding the contract method 0x1d4241dd.
//
// Solidity: function depositGenericErc(_destChain uint256, _value uint256, _to address, _tokenAddress address) returns()
func (_Emitter *EmitterTransactorSession) DepositGenericErc(_destChain *big.Int, _value *big.Int, _to common.Address, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Emitter.Contract.DepositGenericErc(&_Emitter.TransactOpts, _destChain, _value, _to, _tokenAddress)
}

// DepositNFT is a paid mutator transaction binding the contract method 0xfcf86165.
//
// Solidity: function depositNFT(_destChain uint256, _to address, _tokenAddress address, _tokenId uint256, _metaData bytes) returns()
func (_Emitter *EmitterTransactor) DepositNFT(opts *bind.TransactOpts, _destChain *big.Int, _to common.Address, _tokenAddress common.Address, _tokenId *big.Int, _metaData []byte) (*types.Transaction, error) {
	return _Emitter.contract.Transact(opts, "depositNFT", _destChain, _to, _tokenAddress, _tokenId, _metaData)
}

// DepositNFT is a paid mutator transaction binding the contract method 0xfcf86165.
//
// Solidity: function depositNFT(_destChain uint256, _to address, _tokenAddress address, _tokenId uint256, _metaData bytes) returns()
func (_Emitter *EmitterSession) DepositNFT(_destChain *big.Int, _to common.Address, _tokenAddress common.Address, _tokenId *big.Int, _metaData []byte) (*types.Transaction, error) {
	return _Emitter.Contract.DepositNFT(&_Emitter.TransactOpts, _destChain, _to, _tokenAddress, _tokenId, _metaData)
}

// DepositNFT is a paid mutator transaction binding the contract method 0xfcf86165.
//
// Solidity: function depositNFT(_destChain uint256, _to address, _tokenAddress address, _tokenId uint256, _metaData bytes) returns()
func (_Emitter *EmitterTransactorSession) DepositNFT(_destChain *big.Int, _to common.Address, _tokenAddress common.Address, _tokenId *big.Int, _metaData []byte) (*types.Transaction, error) {
	return _Emitter.Contract.DepositNFT(&_Emitter.TransactOpts, _destChain, _to, _tokenAddress, _tokenId, _metaData)
}

// ReleaseErc is a paid mutator transaction binding the contract method 0x34455733.
//
// Solidity: function releaseErc(_tokenAddress address, _value uint256, _to address) returns()
func (_Emitter *EmitterTransactor) ReleaseErc(opts *bind.TransactOpts, _tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Emitter.contract.Transact(opts, "releaseErc", _tokenAddress, _value, _to)
}

// ReleaseErc is a paid mutator transaction binding the contract method 0x34455733.
//
// Solidity: function releaseErc(_tokenAddress address, _value uint256, _to address) returns()
func (_Emitter *EmitterSession) ReleaseErc(_tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Emitter.Contract.ReleaseErc(&_Emitter.TransactOpts, _tokenAddress, _value, _to)
}

// ReleaseErc is a paid mutator transaction binding the contract method 0x34455733.
//
// Solidity: function releaseErc(_tokenAddress address, _value uint256, _to address) returns()
func (_Emitter *EmitterTransactorSession) ReleaseErc(_tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Emitter.Contract.ReleaseErc(&_Emitter.TransactOpts, _tokenAddress, _value, _to)
}

// ReleaseNFT is a paid mutator transaction binding the contract method 0x2b4380c5.
//
// Solidity: function releaseNFT(_tokenAddress address, _to address, _tokenId uint256) returns()
func (_Emitter *EmitterTransactor) ReleaseNFT(opts *bind.TransactOpts, _tokenAddress common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Emitter.contract.Transact(opts, "releaseNFT", _tokenAddress, _to, _tokenId)
}

// ReleaseNFT is a paid mutator transaction binding the contract method 0x2b4380c5.
//
// Solidity: function releaseNFT(_tokenAddress address, _to address, _tokenId uint256) returns()
func (_Emitter *EmitterSession) ReleaseNFT(_tokenAddress common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Emitter.Contract.ReleaseNFT(&_Emitter.TransactOpts, _tokenAddress, _to, _tokenId)
}

// ReleaseNFT is a paid mutator transaction binding the contract method 0x2b4380c5.
//
// Solidity: function releaseNFT(_tokenAddress address, _to address, _tokenId uint256) returns()
func (_Emitter *EmitterTransactorSession) ReleaseNFT(_tokenAddress common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Emitter.Contract.ReleaseNFT(&_Emitter.TransactOpts, _tokenAddress, _to, _tokenId)
}

// EmitterERCTransferIterator is returned from FilterERCTransfer and is used to iterate over the raw logs and unpacked data for ERCTransfer events raised by the Emitter contract.
type EmitterERCTransferIterator struct {
	Event *EmitterERCTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EmitterERCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmitterERCTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EmitterERCTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EmitterERCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmitterERCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmitterERCTransfer represents a ERCTransfer event raised by the Emitter contract.
type EmitterERCTransfer struct {
	DestChain    *big.Int
	DepositId    *big.Int
	To           common.Address
	Amount       *big.Int
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterERCTransfer is a free log retrieval operation binding the contract event 0x12a20368a5b6a9eb35068693542bef5613d39588e50c62fca6312f1b7d9242fe.
//
// Solidity: e ERCTransfer(_destChain indexed uint256, _depositId indexed uint256, _to address, _amount uint256, _tokenAddress address)
func (_Emitter *EmitterFilterer) FilterERCTransfer(opts *bind.FilterOpts, _destChain []*big.Int, _depositId []*big.Int) (*EmitterERCTransferIterator, error) {

	var _destChainRule []interface{}
	for _, _destChainItem := range _destChain {
		_destChainRule = append(_destChainRule, _destChainItem)
	}
	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}

	logs, sub, err := _Emitter.contract.FilterLogs(opts, "ERCTransfer", _destChainRule, _depositIdRule)
	if err != nil {
		return nil, err
	}
	return &EmitterERCTransferIterator{contract: _Emitter.contract, event: "ERCTransfer", logs: logs, sub: sub}, nil
}

// WatchERCTransfer is a free log subscription operation binding the contract event 0x12a20368a5b6a9eb35068693542bef5613d39588e50c62fca6312f1b7d9242fe.
//
// Solidity: e ERCTransfer(_destChain indexed uint256, _depositId indexed uint256, _to address, _amount uint256, _tokenAddress address)
func (_Emitter *EmitterFilterer) WatchERCTransfer(opts *bind.WatchOpts, sink chan<- *EmitterERCTransfer, _destChain []*big.Int, _depositId []*big.Int) (event.Subscription, error) {

	var _destChainRule []interface{}
	for _, _destChainItem := range _destChain {
		_destChainRule = append(_destChainRule, _destChainItem)
	}
	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}

	logs, sub, err := _Emitter.contract.WatchLogs(opts, "ERCTransfer", _destChainRule, _depositIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmitterERCTransfer)
				if err := _Emitter.contract.UnpackLog(event, "ERCTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EmitterGenericTransferIterator is returned from FilterGenericTransfer and is used to iterate over the raw logs and unpacked data for GenericTransfer events raised by the Emitter contract.
type EmitterGenericTransferIterator struct {
	Event *EmitterGenericTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EmitterGenericTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmitterGenericTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EmitterGenericTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EmitterGenericTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmitterGenericTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmitterGenericTransfer represents a GenericTransfer event raised by the Emitter contract.
type EmitterGenericTransfer struct {
	DestChain   *big.Int
	DepositId   *big.Int
	DestAddress common.Address
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGenericTransfer is a free log retrieval operation binding the contract event 0x6cc01912ea1b73308485cfdf6b52b3a8c1e8f33825352b5d5fc0b5ddec3a2a2b.
//
// Solidity: e GenericTransfer(_destChain indexed uint256, _depositId indexed uint256, _destAddress address, _data bytes)
func (_Emitter *EmitterFilterer) FilterGenericTransfer(opts *bind.FilterOpts, _destChain []*big.Int, _depositId []*big.Int) (*EmitterGenericTransferIterator, error) {

	var _destChainRule []interface{}
	for _, _destChainItem := range _destChain {
		_destChainRule = append(_destChainRule, _destChainItem)
	}
	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}

	logs, sub, err := _Emitter.contract.FilterLogs(opts, "GenericTransfer", _destChainRule, _depositIdRule)
	if err != nil {
		return nil, err
	}
	return &EmitterGenericTransferIterator{contract: _Emitter.contract, event: "GenericTransfer", logs: logs, sub: sub}, nil
}

// WatchGenericTransfer is a free log subscription operation binding the contract event 0x6cc01912ea1b73308485cfdf6b52b3a8c1e8f33825352b5d5fc0b5ddec3a2a2b.
//
// Solidity: e GenericTransfer(_destChain indexed uint256, _depositId indexed uint256, _destAddress address, _data bytes)
func (_Emitter *EmitterFilterer) WatchGenericTransfer(opts *bind.WatchOpts, sink chan<- *EmitterGenericTransfer, _destChain []*big.Int, _depositId []*big.Int) (event.Subscription, error) {

	var _destChainRule []interface{}
	for _, _destChainItem := range _destChain {
		_destChainRule = append(_destChainRule, _destChainItem)
	}
	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}

	logs, sub, err := _Emitter.contract.WatchLogs(opts, "GenericTransfer", _destChainRule, _depositIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmitterGenericTransfer)
				if err := _Emitter.contract.UnpackLog(event, "GenericTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EmitterNFTTransferIterator is returned from FilterNFTTransfer and is used to iterate over the raw logs and unpacked data for NFTTransfer events raised by the Emitter contract.
type EmitterNFTTransferIterator struct {
	Event *EmitterNFTTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EmitterNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmitterNFTTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EmitterNFTTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EmitterNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmitterNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmitterNFTTransfer represents a NFTTransfer event raised by the Emitter contract.
type EmitterNFTTransfer struct {
	DestChain    *big.Int
	DepositId    *big.Int
	To           common.Address
	TokenAddress common.Address
	TokenId      *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNFTTransfer is a free log retrieval operation binding the contract event 0x0e6815b49056b27f8b1a0f27453bdd5caedf0eb14fa1e1c1f4f11be9cba05f05.
//
// Solidity: e NFTTransfer(_destChain indexed uint256, _depositId indexed uint256, _to address, _tokenAddress address, _tokenId uint256, _data bytes)
func (_Emitter *EmitterFilterer) FilterNFTTransfer(opts *bind.FilterOpts, _destChain []*big.Int, _depositId []*big.Int) (*EmitterNFTTransferIterator, error) {

	var _destChainRule []interface{}
	for _, _destChainItem := range _destChain {
		_destChainRule = append(_destChainRule, _destChainItem)
	}
	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}

	logs, sub, err := _Emitter.contract.FilterLogs(opts, "NFTTransfer", _destChainRule, _depositIdRule)
	if err != nil {
		return nil, err
	}
	return &EmitterNFTTransferIterator{contract: _Emitter.contract, event: "NFTTransfer", logs: logs, sub: sub}, nil
}

// WatchNFTTransfer is a free log subscription operation binding the contract event 0x0e6815b49056b27f8b1a0f27453bdd5caedf0eb14fa1e1c1f4f11be9cba05f05.
//
// Solidity: e NFTTransfer(_destChain indexed uint256, _depositId indexed uint256, _to address, _tokenAddress address, _tokenId uint256, _data bytes)
func (_Emitter *EmitterFilterer) WatchNFTTransfer(opts *bind.WatchOpts, sink chan<- *EmitterNFTTransfer, _destChain []*big.Int, _depositId []*big.Int) (event.Subscription, error) {

	var _destChainRule []interface{}
	for _, _destChainItem := range _destChain {
		_destChainRule = append(_destChainRule, _destChainItem)
	}
	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}

	logs, sub, err := _Emitter.contract.WatchLogs(opts, "NFTTransfer", _destChainRule, _depositIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmitterNFTTransfer)
				if err := _Emitter.contract.UnpackLog(event, "NFTTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
