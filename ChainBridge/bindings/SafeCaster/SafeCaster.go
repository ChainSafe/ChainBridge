// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SafeCaster

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SafeCasterMetaData contains all meta data concerning the SafeCaster contract.
var SafeCasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"input\",\"type\":\"uint256\"}],\"name\":\"toUint200\",\"outputs\":[{\"internalType\":\"uint200\",\"name\":\"\",\"type\":\"uint200\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610105806100206000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063720ad67414602d575b600080fd5b604760048036036020811015604157600080fd5b50356063565b604080516001600160c81b039092168252519081900360200190f35b6000606c826072565b92915050565b6000600160c81b821060cb576040805162461bcd60e51b815260206004820152601e60248201527f76616c756520646f6573206e6f742066697420696e2032303020626974730000604482015290519081900360640190fd5b509056fea26469706673582212201b9720cab50db25fce506a438b8d8aebc48e87401e69675ec0d69b7eb684596a64736f6c634300060c0033",
}

// SafeCasterABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeCasterMetaData.ABI instead.
var SafeCasterABI = SafeCasterMetaData.ABI

// SafeCasterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeCasterMetaData.Bin instead.
var SafeCasterBin = SafeCasterMetaData.Bin

// DeploySafeCaster deploys a new Ethereum contract, binding an instance of SafeCaster to it.
func DeploySafeCaster(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeCaster, error) {
	parsed, err := SafeCasterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeCasterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeCaster{SafeCasterCaller: SafeCasterCaller{contract: contract}, SafeCasterTransactor: SafeCasterTransactor{contract: contract}, SafeCasterFilterer: SafeCasterFilterer{contract: contract}}, nil
}

// SafeCaster is an auto generated Go binding around an Ethereum contract.
type SafeCaster struct {
	SafeCasterCaller     // Read-only binding to the contract
	SafeCasterTransactor // Write-only binding to the contract
	SafeCasterFilterer   // Log filterer for contract events
}

// SafeCasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeCasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeCasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeCasterSession struct {
	Contract     *SafeCaster       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCasterCallerSession struct {
	Contract *SafeCasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SafeCasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeCasterTransactorSession struct {
	Contract     *SafeCasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SafeCasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeCasterRaw struct {
	Contract *SafeCaster // Generic contract binding to access the raw methods on
}

// SafeCasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCasterCallerRaw struct {
	Contract *SafeCasterCaller // Generic read-only contract binding to access the raw methods on
}

// SafeCasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeCasterTransactorRaw struct {
	Contract *SafeCasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeCaster creates a new instance of SafeCaster, bound to a specific deployed contract.
func NewSafeCaster(address common.Address, backend bind.ContractBackend) (*SafeCaster, error) {
	contract, err := bindSafeCaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeCaster{SafeCasterCaller: SafeCasterCaller{contract: contract}, SafeCasterTransactor: SafeCasterTransactor{contract: contract}, SafeCasterFilterer: SafeCasterFilterer{contract: contract}}, nil
}

// NewSafeCasterCaller creates a new read-only instance of SafeCaster, bound to a specific deployed contract.
func NewSafeCasterCaller(address common.Address, caller bind.ContractCaller) (*SafeCasterCaller, error) {
	contract, err := bindSafeCaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCasterCaller{contract: contract}, nil
}

// NewSafeCasterTransactor creates a new write-only instance of SafeCaster, bound to a specific deployed contract.
func NewSafeCasterTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeCasterTransactor, error) {
	contract, err := bindSafeCaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCasterTransactor{contract: contract}, nil
}

// NewSafeCasterFilterer creates a new log filterer instance of SafeCaster, bound to a specific deployed contract.
func NewSafeCasterFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeCasterFilterer, error) {
	contract, err := bindSafeCaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeCasterFilterer{contract: contract}, nil
}

// bindSafeCaster binds a generic wrapper to an already deployed contract.
func bindSafeCaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeCasterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCaster *SafeCasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCaster.Contract.SafeCasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCaster *SafeCasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCaster.Contract.SafeCasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCaster *SafeCasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCaster.Contract.SafeCasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCaster *SafeCasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCaster *SafeCasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCaster *SafeCasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCaster.Contract.contract.Transact(opts, method, params...)
}

// ToUint200 is a free data retrieval call binding the contract method 0x720ad674.
//
// Solidity: function toUint200(uint256 input) pure returns(uint200)
func (_SafeCaster *SafeCasterCaller) ToUint200(opts *bind.CallOpts, input *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SafeCaster.contract.Call(opts, &out, "toUint200", input)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToUint200 is a free data retrieval call binding the contract method 0x720ad674.
//
// Solidity: function toUint200(uint256 input) pure returns(uint200)
func (_SafeCaster *SafeCasterSession) ToUint200(input *big.Int) (*big.Int, error) {
	return _SafeCaster.Contract.ToUint200(&_SafeCaster.CallOpts, input)
}

// ToUint200 is a free data retrieval call binding the contract method 0x720ad674.
//
// Solidity: function toUint200(uint256 input) pure returns(uint200)
func (_SafeCaster *SafeCasterCallerSession) ToUint200(input *big.Int) (*big.Int, error) {
	return _SafeCaster.Contract.ToUint200(&_SafeCaster.CallOpts, input)
}
