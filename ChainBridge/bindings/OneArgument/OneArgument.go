// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OneArgument

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

// OneArgumentABI is the input ABI used to generate the binding from.
const OneArgumentABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"argumentOne\",\"type\":\"uint256\"}],\"name\":\"OneArgumentCalled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"argumentOne\",\"type\":\"uint256\"}],\"name\":\"oneArgument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OneArgumentBin is the compiled bytecode used for deploying new contracts.
var OneArgumentBin = "0x6080604052348015600f57600080fd5b5060be8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063c95cf0d814602d575b600080fd5b605660048036036020811015604157600080fd5b81019080803590602001909291905050506058565b005b807f29ab08c845830c69b55a1fba5c95718f65dc24361a471e3da14cd5ff2b37315960405160405180910390a25056fea2646970667358221220f5d9df102f413a664744d3a135d0827dd447379c2856d373efa98a988bf4442864736f6c63430006040033"

// DeployOneArgument deploys a new Ethereum contract, binding an instance of OneArgument to it.
func DeployOneArgument(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneArgument, error) {
	parsed, err := abi.JSON(strings.NewReader(OneArgumentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OneArgumentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneArgument{OneArgumentCaller: OneArgumentCaller{contract: contract}, OneArgumentTransactor: OneArgumentTransactor{contract: contract}, OneArgumentFilterer: OneArgumentFilterer{contract: contract}}, nil
}

// OneArgument is an auto generated Go binding around an Ethereum contract.
type OneArgument struct {
	OneArgumentCaller     // Read-only binding to the contract
	OneArgumentTransactor // Write-only binding to the contract
	OneArgumentFilterer   // Log filterer for contract events
}

// OneArgumentCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneArgumentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneArgumentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneArgumentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneArgumentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneArgumentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneArgumentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneArgumentSession struct {
	Contract     *OneArgument      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneArgumentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneArgumentCallerSession struct {
	Contract *OneArgumentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OneArgumentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneArgumentTransactorSession struct {
	Contract     *OneArgumentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OneArgumentRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneArgumentRaw struct {
	Contract *OneArgument // Generic contract binding to access the raw methods on
}

// OneArgumentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneArgumentCallerRaw struct {
	Contract *OneArgumentCaller // Generic read-only contract binding to access the raw methods on
}

// OneArgumentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneArgumentTransactorRaw struct {
	Contract *OneArgumentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneArgument creates a new instance of OneArgument, bound to a specific deployed contract.
func NewOneArgument(address common.Address, backend bind.ContractBackend) (*OneArgument, error) {
	contract, err := bindOneArgument(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneArgument{OneArgumentCaller: OneArgumentCaller{contract: contract}, OneArgumentTransactor: OneArgumentTransactor{contract: contract}, OneArgumentFilterer: OneArgumentFilterer{contract: contract}}, nil
}

// NewOneArgumentCaller creates a new read-only instance of OneArgument, bound to a specific deployed contract.
func NewOneArgumentCaller(address common.Address, caller bind.ContractCaller) (*OneArgumentCaller, error) {
	contract, err := bindOneArgument(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneArgumentCaller{contract: contract}, nil
}

// NewOneArgumentTransactor creates a new write-only instance of OneArgument, bound to a specific deployed contract.
func NewOneArgumentTransactor(address common.Address, transactor bind.ContractTransactor) (*OneArgumentTransactor, error) {
	contract, err := bindOneArgument(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneArgumentTransactor{contract: contract}, nil
}

// NewOneArgumentFilterer creates a new log filterer instance of OneArgument, bound to a specific deployed contract.
func NewOneArgumentFilterer(address common.Address, filterer bind.ContractFilterer) (*OneArgumentFilterer, error) {
	contract, err := bindOneArgument(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneArgumentFilterer{contract: contract}, nil
}

// bindOneArgument binds a generic wrapper to an already deployed contract.
func bindOneArgument(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneArgumentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneArgument *OneArgumentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneArgument.Contract.OneArgumentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneArgument *OneArgumentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneArgument.Contract.OneArgumentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneArgument *OneArgumentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneArgument.Contract.OneArgumentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneArgument *OneArgumentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneArgument.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneArgument *OneArgumentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneArgument.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneArgument *OneArgumentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneArgument.Contract.contract.Transact(opts, method, params...)
}

// OneArgument is a paid mutator transaction binding the contract method 0xc95cf0d8.
//
// Solidity: function oneArgument(uint256 argumentOne) returns()
func (_OneArgument *OneArgumentTransactor) OneArgument(opts *bind.TransactOpts, argumentOne *big.Int) (*types.Transaction, error) {
	return _OneArgument.contract.Transact(opts, "oneArgument", argumentOne)
}

// OneArgument is a paid mutator transaction binding the contract method 0xc95cf0d8.
//
// Solidity: function oneArgument(uint256 argumentOne) returns()
func (_OneArgument *OneArgumentSession) OneArgument(argumentOne *big.Int) (*types.Transaction, error) {
	return _OneArgument.Contract.OneArgument(&_OneArgument.TransactOpts, argumentOne)
}

// OneArgument is a paid mutator transaction binding the contract method 0xc95cf0d8.
//
// Solidity: function oneArgument(uint256 argumentOne) returns()
func (_OneArgument *OneArgumentTransactorSession) OneArgument(argumentOne *big.Int) (*types.Transaction, error) {
	return _OneArgument.Contract.OneArgument(&_OneArgument.TransactOpts, argumentOne)
}

// OneArgumentOneArgumentCalledIterator is returned from FilterOneArgumentCalled and is used to iterate over the raw logs and unpacked data for OneArgumentCalled events raised by the OneArgument contract.
type OneArgumentOneArgumentCalledIterator struct {
	Event *OneArgumentOneArgumentCalled // Event containing the contract specifics and raw log

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
func (it *OneArgumentOneArgumentCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneArgumentOneArgumentCalled)
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
		it.Event = new(OneArgumentOneArgumentCalled)
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
func (it *OneArgumentOneArgumentCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneArgumentOneArgumentCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneArgumentOneArgumentCalled represents a OneArgumentCalled event raised by the OneArgument contract.
type OneArgumentOneArgumentCalled struct {
	ArgumentOne *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOneArgumentCalled is a free log retrieval operation binding the contract event 0x29ab08c845830c69b55a1fba5c95718f65dc24361a471e3da14cd5ff2b373159.
//
// Solidity: event OneArgumentCalled(uint256 indexed argumentOne)
func (_OneArgument *OneArgumentFilterer) FilterOneArgumentCalled(opts *bind.FilterOpts, argumentOne []*big.Int) (*OneArgumentOneArgumentCalledIterator, error) {

	var argumentOneRule []interface{}
	for _, argumentOneItem := range argumentOne {
		argumentOneRule = append(argumentOneRule, argumentOneItem)
	}

	logs, sub, err := _OneArgument.contract.FilterLogs(opts, "OneArgumentCalled", argumentOneRule)
	if err != nil {
		return nil, err
	}
	return &OneArgumentOneArgumentCalledIterator{contract: _OneArgument.contract, event: "OneArgumentCalled", logs: logs, sub: sub}, nil
}

// WatchOneArgumentCalled is a free log subscription operation binding the contract event 0x29ab08c845830c69b55a1fba5c95718f65dc24361a471e3da14cd5ff2b373159.
//
// Solidity: event OneArgumentCalled(uint256 indexed argumentOne)
func (_OneArgument *OneArgumentFilterer) WatchOneArgumentCalled(opts *bind.WatchOpts, sink chan<- *OneArgumentOneArgumentCalled, argumentOne []*big.Int) (event.Subscription, error) {

	var argumentOneRule []interface{}
	for _, argumentOneItem := range argumentOne {
		argumentOneRule = append(argumentOneRule, argumentOneItem)
	}

	logs, sub, err := _OneArgument.contract.WatchLogs(opts, "OneArgumentCalled", argumentOneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneArgumentOneArgumentCalled)
				if err := _OneArgument.contract.UnpackLog(event, "OneArgumentCalled", log); err != nil {
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

// ParseOneArgumentCalled is a log parse operation binding the contract event 0x29ab08c845830c69b55a1fba5c95718f65dc24361a471e3da14cd5ff2b373159.
//
// Solidity: event OneArgumentCalled(uint256 indexed argumentOne)
func (_OneArgument *OneArgumentFilterer) ParseOneArgumentCalled(log types.Log) (*OneArgumentOneArgumentCalled, error) {
	event := new(OneArgumentOneArgumentCalled)
	if err := _OneArgument.contract.UnpackLog(event, "OneArgumentCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
