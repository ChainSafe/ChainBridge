// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WithDepositer

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

// WithDepositerMetaData contains all meta data concerning the WithDepositer contract.
var WithDepositerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"argumentOne\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"argumentTwo\",\"type\":\"uint256\"}],\"name\":\"WithDepositerCalled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"argumentOne\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"argumentTwo\",\"type\":\"uint256\"}],\"name\":\"withDepositer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5060d68061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80630a357c4f14602d575b600080fd5b605660048036036040811015604157600080fd5b506001600160a01b0381351690602001356058565b005b604080516001600160a01b03841681526020810183905281517f0c190bb7953fe744192f379d4b36b025584bb5f386755d164648a99b8ff9e5bb929181900390910190a1505056fea2646970667358221220ad208175e56fe0394f2bc26a74fdfca7de5aecba84b64d1f7d3a97ff526208d764736f6c634300060c0033",
}

// WithDepositerABI is the input ABI used to generate the binding from.
// Deprecated: Use WithDepositerMetaData.ABI instead.
var WithDepositerABI = WithDepositerMetaData.ABI

// WithDepositerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WithDepositerMetaData.Bin instead.
var WithDepositerBin = WithDepositerMetaData.Bin

// DeployWithDepositer deploys a new Ethereum contract, binding an instance of WithDepositer to it.
func DeployWithDepositer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WithDepositer, error) {
	parsed, err := WithDepositerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WithDepositerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WithDepositer{WithDepositerCaller: WithDepositerCaller{contract: contract}, WithDepositerTransactor: WithDepositerTransactor{contract: contract}, WithDepositerFilterer: WithDepositerFilterer{contract: contract}}, nil
}

// WithDepositer is an auto generated Go binding around an Ethereum contract.
type WithDepositer struct {
	WithDepositerCaller     // Read-only binding to the contract
	WithDepositerTransactor // Write-only binding to the contract
	WithDepositerFilterer   // Log filterer for contract events
}

// WithDepositerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithDepositerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithDepositerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithDepositerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithDepositerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithDepositerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithDepositerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithDepositerSession struct {
	Contract     *WithDepositer    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WithDepositerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithDepositerCallerSession struct {
	Contract *WithDepositerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WithDepositerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithDepositerTransactorSession struct {
	Contract     *WithDepositerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WithDepositerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithDepositerRaw struct {
	Contract *WithDepositer // Generic contract binding to access the raw methods on
}

// WithDepositerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithDepositerCallerRaw struct {
	Contract *WithDepositerCaller // Generic read-only contract binding to access the raw methods on
}

// WithDepositerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithDepositerTransactorRaw struct {
	Contract *WithDepositerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithDepositer creates a new instance of WithDepositer, bound to a specific deployed contract.
func NewWithDepositer(address common.Address, backend bind.ContractBackend) (*WithDepositer, error) {
	contract, err := bindWithDepositer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WithDepositer{WithDepositerCaller: WithDepositerCaller{contract: contract}, WithDepositerTransactor: WithDepositerTransactor{contract: contract}, WithDepositerFilterer: WithDepositerFilterer{contract: contract}}, nil
}

// NewWithDepositerCaller creates a new read-only instance of WithDepositer, bound to a specific deployed contract.
func NewWithDepositerCaller(address common.Address, caller bind.ContractCaller) (*WithDepositerCaller, error) {
	contract, err := bindWithDepositer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithDepositerCaller{contract: contract}, nil
}

// NewWithDepositerTransactor creates a new write-only instance of WithDepositer, bound to a specific deployed contract.
func NewWithDepositerTransactor(address common.Address, transactor bind.ContractTransactor) (*WithDepositerTransactor, error) {
	contract, err := bindWithDepositer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithDepositerTransactor{contract: contract}, nil
}

// NewWithDepositerFilterer creates a new log filterer instance of WithDepositer, bound to a specific deployed contract.
func NewWithDepositerFilterer(address common.Address, filterer bind.ContractFilterer) (*WithDepositerFilterer, error) {
	contract, err := bindWithDepositer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithDepositerFilterer{contract: contract}, nil
}

// bindWithDepositer binds a generic wrapper to an already deployed contract.
func bindWithDepositer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WithDepositerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithDepositer *WithDepositerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithDepositer.Contract.WithDepositerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithDepositer *WithDepositerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithDepositer.Contract.WithDepositerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithDepositer *WithDepositerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithDepositer.Contract.WithDepositerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithDepositer *WithDepositerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithDepositer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithDepositer *WithDepositerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithDepositer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithDepositer *WithDepositerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithDepositer.Contract.contract.Transact(opts, method, params...)
}

// WithDepositer is a paid mutator transaction binding the contract method 0x0a357c4f.
//
// Solidity: function withDepositer(address argumentOne, uint256 argumentTwo) returns()
func (_WithDepositer *WithDepositerTransactor) WithDepositer(opts *bind.TransactOpts, argumentOne common.Address, argumentTwo *big.Int) (*types.Transaction, error) {
	return _WithDepositer.contract.Transact(opts, "withDepositer", argumentOne, argumentTwo)
}

// WithDepositer is a paid mutator transaction binding the contract method 0x0a357c4f.
//
// Solidity: function withDepositer(address argumentOne, uint256 argumentTwo) returns()
func (_WithDepositer *WithDepositerSession) WithDepositer(argumentOne common.Address, argumentTwo *big.Int) (*types.Transaction, error) {
	return _WithDepositer.Contract.WithDepositer(&_WithDepositer.TransactOpts, argumentOne, argumentTwo)
}

// WithDepositer is a paid mutator transaction binding the contract method 0x0a357c4f.
//
// Solidity: function withDepositer(address argumentOne, uint256 argumentTwo) returns()
func (_WithDepositer *WithDepositerTransactorSession) WithDepositer(argumentOne common.Address, argumentTwo *big.Int) (*types.Transaction, error) {
	return _WithDepositer.Contract.WithDepositer(&_WithDepositer.TransactOpts, argumentOne, argumentTwo)
}

// WithDepositerWithDepositerCalledIterator is returned from FilterWithDepositerCalled and is used to iterate over the raw logs and unpacked data for WithDepositerCalled events raised by the WithDepositer contract.
type WithDepositerWithDepositerCalledIterator struct {
	Event *WithDepositerWithDepositerCalled // Event containing the contract specifics and raw log

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
func (it *WithDepositerWithDepositerCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithDepositerWithDepositerCalled)
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
		it.Event = new(WithDepositerWithDepositerCalled)
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
func (it *WithDepositerWithDepositerCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithDepositerWithDepositerCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithDepositerWithDepositerCalled represents a WithDepositerCalled event raised by the WithDepositer contract.
type WithDepositerWithDepositerCalled struct {
	ArgumentOne common.Address
	ArgumentTwo *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithDepositerCalled is a free log retrieval operation binding the contract event 0x0c190bb7953fe744192f379d4b36b025584bb5f386755d164648a99b8ff9e5bb.
//
// Solidity: event WithDepositerCalled(address argumentOne, uint256 argumentTwo)
func (_WithDepositer *WithDepositerFilterer) FilterWithDepositerCalled(opts *bind.FilterOpts) (*WithDepositerWithDepositerCalledIterator, error) {

	logs, sub, err := _WithDepositer.contract.FilterLogs(opts, "WithDepositerCalled")
	if err != nil {
		return nil, err
	}
	return &WithDepositerWithDepositerCalledIterator{contract: _WithDepositer.contract, event: "WithDepositerCalled", logs: logs, sub: sub}, nil
}

// WatchWithDepositerCalled is a free log subscription operation binding the contract event 0x0c190bb7953fe744192f379d4b36b025584bb5f386755d164648a99b8ff9e5bb.
//
// Solidity: event WithDepositerCalled(address argumentOne, uint256 argumentTwo)
func (_WithDepositer *WithDepositerFilterer) WatchWithDepositerCalled(opts *bind.WatchOpts, sink chan<- *WithDepositerWithDepositerCalled) (event.Subscription, error) {

	logs, sub, err := _WithDepositer.contract.WatchLogs(opts, "WithDepositerCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithDepositerWithDepositerCalled)
				if err := _WithDepositer.contract.UnpackLog(event, "WithDepositerCalled", log); err != nil {
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

// ParseWithDepositerCalled is a log parse operation binding the contract event 0x0c190bb7953fe744192f379d4b36b025584bb5f386755d164648a99b8ff9e5bb.
//
// Solidity: event WithDepositerCalled(address argumentOne, uint256 argumentTwo)
func (_WithDepositer *WithDepositerFilterer) ParseWithDepositerCalled(log types.Log) (*WithDepositerWithDepositerCalled, error) {
	event := new(WithDepositerWithDepositerCalled)
	if err := _WithDepositer.contract.UnpackLog(event, "WithDepositerCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
