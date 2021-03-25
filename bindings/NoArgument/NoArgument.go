// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package NoArgument

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NoArgumentABI is the input ABI used to generate the binding from.
const NoArgumentABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"NoArgumentCalled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"noArgument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NoArgumentBin is the compiled bytecode used for deploying new contracts.
var NoArgumentBin = "0x6080604052348015600f57600080fd5b5060998061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063568959ca14602d575b600080fd5b60336035565b005b7fc582abe1670c5a7f7cad8f171e4af03c793dd9f59fee6714179f56b6e9aea26f60405160405180910390a156fea2646970667358221220016fd768f6d334a67219af52840dab7b38ffaf47ec0931ffcdc800ec411c523264736f6c63430006040033"

// DeployNoArgument deploys a new Ethereum contract, binding an instance of NoArgument to it.
func DeployNoArgument(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NoArgument, error) {
	parsed, err := abi.JSON(strings.NewReader(NoArgumentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NoArgumentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NoArgument{NoArgumentCaller: NoArgumentCaller{contract: contract}, NoArgumentTransactor: NoArgumentTransactor{contract: contract}, NoArgumentFilterer: NoArgumentFilterer{contract: contract}}, nil
}

// NoArgument is an auto generated Go binding around an Ethereum contract.
type NoArgument struct {
	NoArgumentCaller     // Read-only binding to the contract
	NoArgumentTransactor // Write-only binding to the contract
	NoArgumentFilterer   // Log filterer for contract events
}

// NoArgumentCaller is an auto generated read-only Go binding around an Ethereum contract.
type NoArgumentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoArgumentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NoArgumentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoArgumentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NoArgumentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoArgumentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NoArgumentSession struct {
	Contract     *NoArgument       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NoArgumentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NoArgumentCallerSession struct {
	Contract *NoArgumentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NoArgumentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NoArgumentTransactorSession struct {
	Contract     *NoArgumentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NoArgumentRaw is an auto generated low-level Go binding around an Ethereum contract.
type NoArgumentRaw struct {
	Contract *NoArgument // Generic contract binding to access the raw methods on
}

// NoArgumentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NoArgumentCallerRaw struct {
	Contract *NoArgumentCaller // Generic read-only contract binding to access the raw methods on
}

// NoArgumentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NoArgumentTransactorRaw struct {
	Contract *NoArgumentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNoArgument creates a new instance of NoArgument, bound to a specific deployed contract.
func NewNoArgument(address common.Address, backend bind.ContractBackend) (*NoArgument, error) {
	contract, err := bindNoArgument(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NoArgument{NoArgumentCaller: NoArgumentCaller{contract: contract}, NoArgumentTransactor: NoArgumentTransactor{contract: contract}, NoArgumentFilterer: NoArgumentFilterer{contract: contract}}, nil
}

// NewNoArgumentCaller creates a new read-only instance of NoArgument, bound to a specific deployed contract.
func NewNoArgumentCaller(address common.Address, caller bind.ContractCaller) (*NoArgumentCaller, error) {
	contract, err := bindNoArgument(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NoArgumentCaller{contract: contract}, nil
}

// NewNoArgumentTransactor creates a new write-only instance of NoArgument, bound to a specific deployed contract.
func NewNoArgumentTransactor(address common.Address, transactor bind.ContractTransactor) (*NoArgumentTransactor, error) {
	contract, err := bindNoArgument(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NoArgumentTransactor{contract: contract}, nil
}

// NewNoArgumentFilterer creates a new log filterer instance of NoArgument, bound to a specific deployed contract.
func NewNoArgumentFilterer(address common.Address, filterer bind.ContractFilterer) (*NoArgumentFilterer, error) {
	contract, err := bindNoArgument(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NoArgumentFilterer{contract: contract}, nil
}

// bindNoArgument binds a generic wrapper to an already deployed contract.
func bindNoArgument(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NoArgumentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoArgument *NoArgumentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoArgument.Contract.NoArgumentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoArgument *NoArgumentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoArgument.Contract.NoArgumentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoArgument *NoArgumentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoArgument.Contract.NoArgumentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoArgument *NoArgumentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoArgument.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoArgument *NoArgumentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoArgument.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoArgument *NoArgumentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoArgument.Contract.contract.Transact(opts, method, params...)
}

// NoArgument is a paid mutator transaction binding the contract method 0x568959ca.
//
// Solidity: function noArgument() returns()
func (_NoArgument *NoArgumentTransactor) NoArgument(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoArgument.contract.Transact(opts, "noArgument")
}

// NoArgument is a paid mutator transaction binding the contract method 0x568959ca.
//
// Solidity: function noArgument() returns()
func (_NoArgument *NoArgumentSession) NoArgument() (*types.Transaction, error) {
	return _NoArgument.Contract.NoArgument(&_NoArgument.TransactOpts)
}

// NoArgument is a paid mutator transaction binding the contract method 0x568959ca.
//
// Solidity: function noArgument() returns()
func (_NoArgument *NoArgumentTransactorSession) NoArgument() (*types.Transaction, error) {
	return _NoArgument.Contract.NoArgument(&_NoArgument.TransactOpts)
}

// NoArgumentNoArgumentCalledIterator is returned from FilterNoArgumentCalled and is used to iterate over the raw logs and unpacked data for NoArgumentCalled events raised by the NoArgument contract.
type NoArgumentNoArgumentCalledIterator struct {
	Event *NoArgumentNoArgumentCalled // Event containing the contract specifics and raw log

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
func (it *NoArgumentNoArgumentCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoArgumentNoArgumentCalled)
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
		it.Event = new(NoArgumentNoArgumentCalled)
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
func (it *NoArgumentNoArgumentCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoArgumentNoArgumentCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoArgumentNoArgumentCalled represents a NoArgumentCalled event raised by the NoArgument contract.
type NoArgumentNoArgumentCalled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNoArgumentCalled is a free log retrieval operation binding the contract event 0xc582abe1670c5a7f7cad8f171e4af03c793dd9f59fee6714179f56b6e9aea26f.
//
// Solidity: event NoArgumentCalled()
func (_NoArgument *NoArgumentFilterer) FilterNoArgumentCalled(opts *bind.FilterOpts) (*NoArgumentNoArgumentCalledIterator, error) {

	logs, sub, err := _NoArgument.contract.FilterLogs(opts, "NoArgumentCalled")
	if err != nil {
		return nil, err
	}
	return &NoArgumentNoArgumentCalledIterator{contract: _NoArgument.contract, event: "NoArgumentCalled", logs: logs, sub: sub}, nil
}

// WatchNoArgumentCalled is a free log subscription operation binding the contract event 0xc582abe1670c5a7f7cad8f171e4af03c793dd9f59fee6714179f56b6e9aea26f.
//
// Solidity: event NoArgumentCalled()
func (_NoArgument *NoArgumentFilterer) WatchNoArgumentCalled(opts *bind.WatchOpts, sink chan<- *NoArgumentNoArgumentCalled) (event.Subscription, error) {

	logs, sub, err := _NoArgument.contract.WatchLogs(opts, "NoArgumentCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoArgumentNoArgumentCalled)
				if err := _NoArgument.contract.UnpackLog(event, "NoArgumentCalled", log); err != nil {
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

// ParseNoArgumentCalled is a log parse operation binding the contract event 0xc582abe1670c5a7f7cad8f171e4af03c793dd9f59fee6714179f56b6e9aea26f.
//
// Solidity: event NoArgumentCalled()
func (_NoArgument *NoArgumentFilterer) ParseNoArgumentCalled(log types.Log) (*NoArgumentNoArgumentCalled, error) {
	event := new(NoArgumentNoArgumentCalled)
	if err := _NoArgument.contract.UnpackLog(event, "NoArgumentCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
