// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package SimpleEmitter

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SimpleEmitterABI is the input ABI used to generate the binding from.
const SimpleEmitterABI = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"_addr","type":"address"},{"indexed":false,"internalType":"bytes32","name":"_hash","type":"bytes32"}],"name":"DepositAsset","type":"event"},{"payable":false,"stateMutability":"nonpayable","type":"fallback"}]`

// SimpleEmitterBin is the compiled bytecode used for deploying new contracts.
const SimpleEmitterBin = `60806040523480156100115760006000fd5b50610017565b6101a8806100266000396000f3fe60806040526004361061000d575b34801561001a5760006000fd5b505b600060023343604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b6020831015156100ad57805182525b602082019150602081019050602083039250610087565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa1580156100f0573d600060003e3d6000fd5b5050506040513d60208110156101065760006000fd5b810190808051906020019092919050505090503373ffffffffffffffffffffffffffffffffffffffff167f073221459ee71dc19e3af5573dea23a57451d68b87132e600a1f4f4d948419508260405180826000191660001916815260200191505060405180910390a2505b00fea265627a7a72315820d32a199b2a10503d801478bd70dc699134119da41b738e4c9ca0b8f7637a2cf264736f6c634300050c0032`

// DeploySimpleEmitter deploys a new Ethereum contract, binding an instance of SimpleEmitter to it.
func DeploySimpleEmitter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleEmitter, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleEmitterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleEmitterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleEmitter{SimpleEmitterCaller: SimpleEmitterCaller{contract: contract}, SimpleEmitterTransactor: SimpleEmitterTransactor{contract: contract}}, nil
}

// SimpleEmitter is an auto generated Go binding around an Ethereum contract.
type SimpleEmitter struct {
	SimpleEmitterCaller     // Read-only binding to the contract
	SimpleEmitterTransactor // Write-only binding to the contract
}

// SimpleEmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleEmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleEmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleEmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleEmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleEmitterSession struct {
	Contract     *SimpleEmitter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleEmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleEmitterCallerSession struct {
	Contract *SimpleEmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleEmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleEmitterTransactorSession struct {
	Contract     *SimpleEmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleEmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleEmitterRaw struct {
	Contract *SimpleEmitter // Generic contract binding to access the raw methods on
}

// SimpleEmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleEmitterCallerRaw struct {
	Contract *SimpleEmitterCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleEmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleEmitterTransactorRaw struct {
	Contract *SimpleEmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleEmitter creates a new instance of SimpleEmitter, bound to a specific deployed contract.
func NewSimpleEmitter(address common.Address, backend bind.ContractBackend) (*SimpleEmitter, error) {
	contract, err := bindSimpleEmitter(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &SimpleEmitter{SimpleEmitterCaller: SimpleEmitterCaller{contract: contract}, SimpleEmitterTransactor: SimpleEmitterTransactor{contract: contract}}, nil
}

// NewSimpleEmitterCaller creates a new read-only instance of SimpleEmitter, bound to a specific deployed contract.
func NewSimpleEmitterCaller(address common.Address, caller bind.ContractCaller) (*SimpleEmitterCaller, error) {
	contract, err := bindSimpleEmitter(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleEmitterCaller{contract: contract}, nil
}

// NewSimpleEmitterTransactor creates a new write-only instance of SimpleEmitter, bound to a specific deployed contract.
func NewSimpleEmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleEmitterTransactor, error) {
	contract, err := bindSimpleEmitter(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SimpleEmitterTransactor{contract: contract}, nil
}

// bindSimpleEmitter binds a generic wrapper to an already deployed contract.
func bindSimpleEmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleEmitterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleEmitter *SimpleEmitterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleEmitter.Contract.SimpleEmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleEmitter *SimpleEmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleEmitter.Contract.SimpleEmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleEmitter *SimpleEmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleEmitter.Contract.SimpleEmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleEmitter *SimpleEmitterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleEmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleEmitter *SimpleEmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleEmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleEmitter *SimpleEmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleEmitter.Contract.contract.Transact(opts, method, params...)
}
