// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ThreeArguments

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

// ThreeArgumentsABI is the input ABI used to generate the binding from.
const ThreeArgumentsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"argumentOne\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int8\",\"name\":\"argumentTwo\",\"type\":\"int8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"argumentThree\",\"type\":\"bool\"}],\"name\":\"ThreeArgumentsCalled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"argumentOne\",\"type\":\"string\"},{\"internalType\":\"int8\",\"name\":\"argumentTwo\",\"type\":\"int8\"},{\"internalType\":\"bool\",\"name\":\"argumentThree\",\"type\":\"bool\"}],\"name\":\"threeArguments\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ThreeArgumentsBin is the compiled bytecode used for deploying new contracts.
var ThreeArgumentsBin = "0x608060405234801561001057600080fd5b5061017b806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80639280b90514610030575b600080fd5b6100c06004803603606081101561004657600080fd5b810190808035906020019064010000000081111561006357600080fd5b82018360208201111561007557600080fd5b8035906020019184600183028401116401000000008311171561009757600080fd5b9091929391929390803560000b90602001909291908035151590602001909291905050506100c2565b005b7fd589183661fa75f94e2db32f4eb7ebb50f4154c160e15eb43f772a46f360a3a88484848460405180806020018460000b60000b8152602001831515151581526020018281038252868682818152602001925080828437600081840152601f19601f8201169050808301925050509550505050505060405180910390a15050505056fea26469706673582212208f4aec4d91a71d18b17c8e25a5cfcf245b92e9397024904d41da26b422c4b6a764736f6c63430006040033"

// DeployThreeArguments deploys a new Ethereum contract, binding an instance of ThreeArguments to it.
func DeployThreeArguments(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ThreeArguments, error) {
	parsed, err := abi.JSON(strings.NewReader(ThreeArgumentsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ThreeArgumentsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ThreeArguments{ThreeArgumentsCaller: ThreeArgumentsCaller{contract: contract}, ThreeArgumentsTransactor: ThreeArgumentsTransactor{contract: contract}, ThreeArgumentsFilterer: ThreeArgumentsFilterer{contract: contract}}, nil
}

// ThreeArguments is an auto generated Go binding around an Ethereum contract.
type ThreeArguments struct {
	ThreeArgumentsCaller     // Read-only binding to the contract
	ThreeArgumentsTransactor // Write-only binding to the contract
	ThreeArgumentsFilterer   // Log filterer for contract events
}

// ThreeArgumentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ThreeArgumentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreeArgumentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ThreeArgumentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreeArgumentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ThreeArgumentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreeArgumentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ThreeArgumentsSession struct {
	Contract     *ThreeArguments   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ThreeArgumentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ThreeArgumentsCallerSession struct {
	Contract *ThreeArgumentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ThreeArgumentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ThreeArgumentsTransactorSession struct {
	Contract     *ThreeArgumentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ThreeArgumentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ThreeArgumentsRaw struct {
	Contract *ThreeArguments // Generic contract binding to access the raw methods on
}

// ThreeArgumentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ThreeArgumentsCallerRaw struct {
	Contract *ThreeArgumentsCaller // Generic read-only contract binding to access the raw methods on
}

// ThreeArgumentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ThreeArgumentsTransactorRaw struct {
	Contract *ThreeArgumentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewThreeArguments creates a new instance of ThreeArguments, bound to a specific deployed contract.
func NewThreeArguments(address common.Address, backend bind.ContractBackend) (*ThreeArguments, error) {
	contract, err := bindThreeArguments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ThreeArguments{ThreeArgumentsCaller: ThreeArgumentsCaller{contract: contract}, ThreeArgumentsTransactor: ThreeArgumentsTransactor{contract: contract}, ThreeArgumentsFilterer: ThreeArgumentsFilterer{contract: contract}}, nil
}

// NewThreeArgumentsCaller creates a new read-only instance of ThreeArguments, bound to a specific deployed contract.
func NewThreeArgumentsCaller(address common.Address, caller bind.ContractCaller) (*ThreeArgumentsCaller, error) {
	contract, err := bindThreeArguments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ThreeArgumentsCaller{contract: contract}, nil
}

// NewThreeArgumentsTransactor creates a new write-only instance of ThreeArguments, bound to a specific deployed contract.
func NewThreeArgumentsTransactor(address common.Address, transactor bind.ContractTransactor) (*ThreeArgumentsTransactor, error) {
	contract, err := bindThreeArguments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ThreeArgumentsTransactor{contract: contract}, nil
}

// NewThreeArgumentsFilterer creates a new log filterer instance of ThreeArguments, bound to a specific deployed contract.
func NewThreeArgumentsFilterer(address common.Address, filterer bind.ContractFilterer) (*ThreeArgumentsFilterer, error) {
	contract, err := bindThreeArguments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ThreeArgumentsFilterer{contract: contract}, nil
}

// bindThreeArguments binds a generic wrapper to an already deployed contract.
func bindThreeArguments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ThreeArgumentsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThreeArguments *ThreeArgumentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ThreeArguments.Contract.ThreeArgumentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThreeArguments *ThreeArgumentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreeArguments.Contract.ThreeArgumentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThreeArguments *ThreeArgumentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThreeArguments.Contract.ThreeArgumentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThreeArguments *ThreeArgumentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ThreeArguments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThreeArguments *ThreeArgumentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreeArguments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThreeArguments *ThreeArgumentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThreeArguments.Contract.contract.Transact(opts, method, params...)
}

// ThreeArguments is a paid mutator transaction binding the contract method 0x9280b905.
//
// Solidity: function threeArguments(string argumentOne, int8 argumentTwo, bool argumentThree) returns()
func (_ThreeArguments *ThreeArgumentsTransactor) ThreeArguments(opts *bind.TransactOpts, argumentOne string, argumentTwo int8, argumentThree bool) (*types.Transaction, error) {
	return _ThreeArguments.contract.Transact(opts, "threeArguments", argumentOne, argumentTwo, argumentThree)
}

// ThreeArguments is a paid mutator transaction binding the contract method 0x9280b905.
//
// Solidity: function threeArguments(string argumentOne, int8 argumentTwo, bool argumentThree) returns()
func (_ThreeArguments *ThreeArgumentsSession) ThreeArguments(argumentOne string, argumentTwo int8, argumentThree bool) (*types.Transaction, error) {
	return _ThreeArguments.Contract.ThreeArguments(&_ThreeArguments.TransactOpts, argumentOne, argumentTwo, argumentThree)
}

// ThreeArguments is a paid mutator transaction binding the contract method 0x9280b905.
//
// Solidity: function threeArguments(string argumentOne, int8 argumentTwo, bool argumentThree) returns()
func (_ThreeArguments *ThreeArgumentsTransactorSession) ThreeArguments(argumentOne string, argumentTwo int8, argumentThree bool) (*types.Transaction, error) {
	return _ThreeArguments.Contract.ThreeArguments(&_ThreeArguments.TransactOpts, argumentOne, argumentTwo, argumentThree)
}

// ThreeArgumentsThreeArgumentsCalledIterator is returned from FilterThreeArgumentsCalled and is used to iterate over the raw logs and unpacked data for ThreeArgumentsCalled events raised by the ThreeArguments contract.
type ThreeArgumentsThreeArgumentsCalledIterator struct {
	Event *ThreeArgumentsThreeArgumentsCalled // Event containing the contract specifics and raw log

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
func (it *ThreeArgumentsThreeArgumentsCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreeArgumentsThreeArgumentsCalled)
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
		it.Event = new(ThreeArgumentsThreeArgumentsCalled)
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
func (it *ThreeArgumentsThreeArgumentsCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreeArgumentsThreeArgumentsCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreeArgumentsThreeArgumentsCalled represents a ThreeArgumentsCalled event raised by the ThreeArguments contract.
type ThreeArgumentsThreeArgumentsCalled struct {
	ArgumentOne   string
	ArgumentTwo   int8
	ArgumentThree bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterThreeArgumentsCalled is a free log retrieval operation binding the contract event 0xd589183661fa75f94e2db32f4eb7ebb50f4154c160e15eb43f772a46f360a3a8.
//
// Solidity: event ThreeArgumentsCalled(string argumentOne, int8 argumentTwo, bool argumentThree)
func (_ThreeArguments *ThreeArgumentsFilterer) FilterThreeArgumentsCalled(opts *bind.FilterOpts) (*ThreeArgumentsThreeArgumentsCalledIterator, error) {

	logs, sub, err := _ThreeArguments.contract.FilterLogs(opts, "ThreeArgumentsCalled")
	if err != nil {
		return nil, err
	}
	return &ThreeArgumentsThreeArgumentsCalledIterator{contract: _ThreeArguments.contract, event: "ThreeArgumentsCalled", logs: logs, sub: sub}, nil
}

// WatchThreeArgumentsCalled is a free log subscription operation binding the contract event 0xd589183661fa75f94e2db32f4eb7ebb50f4154c160e15eb43f772a46f360a3a8.
//
// Solidity: event ThreeArgumentsCalled(string argumentOne, int8 argumentTwo, bool argumentThree)
func (_ThreeArguments *ThreeArgumentsFilterer) WatchThreeArgumentsCalled(opts *bind.WatchOpts, sink chan<- *ThreeArgumentsThreeArgumentsCalled) (event.Subscription, error) {

	logs, sub, err := _ThreeArguments.contract.WatchLogs(opts, "ThreeArgumentsCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreeArgumentsThreeArgumentsCalled)
				if err := _ThreeArguments.contract.UnpackLog(event, "ThreeArgumentsCalled", log); err != nil {
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

// ParseThreeArgumentsCalled is a log parse operation binding the contract event 0xd589183661fa75f94e2db32f4eb7ebb50f4154c160e15eb43f772a46f360a3a8.
//
// Solidity: event ThreeArgumentsCalled(string argumentOne, int8 argumentTwo, bool argumentThree)
func (_ThreeArguments *ThreeArgumentsFilterer) ParseThreeArgumentsCalled(log types.Log) (*ThreeArgumentsThreeArgumentsCalled, error) {
	event := new(ThreeArgumentsThreeArgumentsCalled)
	if err := _ThreeArguments.contract.UnpackLog(event, "ThreeArgumentsCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
