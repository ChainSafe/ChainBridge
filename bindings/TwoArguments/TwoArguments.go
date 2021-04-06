// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TwoArguments

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

// TwoArgumentsABI is the input ABI used to generate the binding from.
const TwoArgumentsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"argumentOne\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"argumentTwo\",\"type\":\"bytes4\"}],\"name\":\"TwoArgumentsCalled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"argumentOne\",\"type\":\"address[]\"},{\"internalType\":\"bytes4\",\"name\":\"argumentTwo\",\"type\":\"bytes4\"}],\"name\":\"twoArguments\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TwoArgumentsBin is the compiled bytecode used for deploying new contracts.
var TwoArgumentsBin = "0x608060405234801561001057600080fd5b506101b9806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806372e0745c14610030575b600080fd5b6100d06004803603604081101561004657600080fd5b810190808035906020019064010000000081111561006357600080fd5b82018360208201111561007557600080fd5b8035906020019184602083028401116401000000008311171561009757600080fd5b909192939192939080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690602001909291905050506100d2565b005b7fc983106aca50fad459fb18ede1630e8ff8147ff28ad451a856427931fd7f15e38383836040518080602001837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020018281038252858582818152602001925060200280828437600081840152601f19601f82011690508083019250505094505050505060405180910390a150505056fea2646970667358221220a31b9056a86d9052aeda1eb343a1499abd98c78fd9482413cdc186678529d7f864736f6c63430006040033"

// DeployTwoArguments deploys a new Ethereum contract, binding an instance of TwoArguments to it.
func DeployTwoArguments(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TwoArguments, error) {
	parsed, err := abi.JSON(strings.NewReader(TwoArgumentsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TwoArgumentsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TwoArguments{TwoArgumentsCaller: TwoArgumentsCaller{contract: contract}, TwoArgumentsTransactor: TwoArgumentsTransactor{contract: contract}, TwoArgumentsFilterer: TwoArgumentsFilterer{contract: contract}}, nil
}

// TwoArguments is an auto generated Go binding around an Ethereum contract.
type TwoArguments struct {
	TwoArgumentsCaller     // Read-only binding to the contract
	TwoArgumentsTransactor // Write-only binding to the contract
	TwoArgumentsFilterer   // Log filterer for contract events
}

// TwoArgumentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TwoArgumentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoArgumentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TwoArgumentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoArgumentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TwoArgumentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoArgumentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TwoArgumentsSession struct {
	Contract     *TwoArguments     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TwoArgumentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TwoArgumentsCallerSession struct {
	Contract *TwoArgumentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TwoArgumentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TwoArgumentsTransactorSession struct {
	Contract     *TwoArgumentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TwoArgumentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TwoArgumentsRaw struct {
	Contract *TwoArguments // Generic contract binding to access the raw methods on
}

// TwoArgumentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TwoArgumentsCallerRaw struct {
	Contract *TwoArgumentsCaller // Generic read-only contract binding to access the raw methods on
}

// TwoArgumentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TwoArgumentsTransactorRaw struct {
	Contract *TwoArgumentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTwoArguments creates a new instance of TwoArguments, bound to a specific deployed contract.
func NewTwoArguments(address common.Address, backend bind.ContractBackend) (*TwoArguments, error) {
	contract, err := bindTwoArguments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TwoArguments{TwoArgumentsCaller: TwoArgumentsCaller{contract: contract}, TwoArgumentsTransactor: TwoArgumentsTransactor{contract: contract}, TwoArgumentsFilterer: TwoArgumentsFilterer{contract: contract}}, nil
}

// NewTwoArgumentsCaller creates a new read-only instance of TwoArguments, bound to a specific deployed contract.
func NewTwoArgumentsCaller(address common.Address, caller bind.ContractCaller) (*TwoArgumentsCaller, error) {
	contract, err := bindTwoArguments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TwoArgumentsCaller{contract: contract}, nil
}

// NewTwoArgumentsTransactor creates a new write-only instance of TwoArguments, bound to a specific deployed contract.
func NewTwoArgumentsTransactor(address common.Address, transactor bind.ContractTransactor) (*TwoArgumentsTransactor, error) {
	contract, err := bindTwoArguments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TwoArgumentsTransactor{contract: contract}, nil
}

// NewTwoArgumentsFilterer creates a new log filterer instance of TwoArguments, bound to a specific deployed contract.
func NewTwoArgumentsFilterer(address common.Address, filterer bind.ContractFilterer) (*TwoArgumentsFilterer, error) {
	contract, err := bindTwoArguments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TwoArgumentsFilterer{contract: contract}, nil
}

// bindTwoArguments binds a generic wrapper to an already deployed contract.
func bindTwoArguments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TwoArgumentsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwoArguments *TwoArgumentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwoArguments.Contract.TwoArgumentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwoArguments *TwoArgumentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwoArguments.Contract.TwoArgumentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwoArguments *TwoArgumentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwoArguments.Contract.TwoArgumentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwoArguments *TwoArgumentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwoArguments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwoArguments *TwoArgumentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwoArguments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwoArguments *TwoArgumentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwoArguments.Contract.contract.Transact(opts, method, params...)
}

// TwoArguments is a paid mutator transaction binding the contract method 0x72e0745c.
//
// Solidity: function twoArguments(address[] argumentOne, bytes4 argumentTwo) returns()
func (_TwoArguments *TwoArgumentsTransactor) TwoArguments(opts *bind.TransactOpts, argumentOne []common.Address, argumentTwo [4]byte) (*types.Transaction, error) {
	return _TwoArguments.contract.Transact(opts, "twoArguments", argumentOne, argumentTwo)
}

// TwoArguments is a paid mutator transaction binding the contract method 0x72e0745c.
//
// Solidity: function twoArguments(address[] argumentOne, bytes4 argumentTwo) returns()
func (_TwoArguments *TwoArgumentsSession) TwoArguments(argumentOne []common.Address, argumentTwo [4]byte) (*types.Transaction, error) {
	return _TwoArguments.Contract.TwoArguments(&_TwoArguments.TransactOpts, argumentOne, argumentTwo)
}

// TwoArguments is a paid mutator transaction binding the contract method 0x72e0745c.
//
// Solidity: function twoArguments(address[] argumentOne, bytes4 argumentTwo) returns()
func (_TwoArguments *TwoArgumentsTransactorSession) TwoArguments(argumentOne []common.Address, argumentTwo [4]byte) (*types.Transaction, error) {
	return _TwoArguments.Contract.TwoArguments(&_TwoArguments.TransactOpts, argumentOne, argumentTwo)
}

// TwoArgumentsTwoArgumentsCalledIterator is returned from FilterTwoArgumentsCalled and is used to iterate over the raw logs and unpacked data for TwoArgumentsCalled events raised by the TwoArguments contract.
type TwoArgumentsTwoArgumentsCalledIterator struct {
	Event *TwoArgumentsTwoArgumentsCalled // Event containing the contract specifics and raw log

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
func (it *TwoArgumentsTwoArgumentsCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoArgumentsTwoArgumentsCalled)
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
		it.Event = new(TwoArgumentsTwoArgumentsCalled)
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
func (it *TwoArgumentsTwoArgumentsCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoArgumentsTwoArgumentsCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoArgumentsTwoArgumentsCalled represents a TwoArgumentsCalled event raised by the TwoArguments contract.
type TwoArgumentsTwoArgumentsCalled struct {
	ArgumentOne []common.Address
	ArgumentTwo [4]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTwoArgumentsCalled is a free log retrieval operation binding the contract event 0xc983106aca50fad459fb18ede1630e8ff8147ff28ad451a856427931fd7f15e3.
//
// Solidity: event TwoArgumentsCalled(address[] argumentOne, bytes4 argumentTwo)
func (_TwoArguments *TwoArgumentsFilterer) FilterTwoArgumentsCalled(opts *bind.FilterOpts) (*TwoArgumentsTwoArgumentsCalledIterator, error) {

	logs, sub, err := _TwoArguments.contract.FilterLogs(opts, "TwoArgumentsCalled")
	if err != nil {
		return nil, err
	}
	return &TwoArgumentsTwoArgumentsCalledIterator{contract: _TwoArguments.contract, event: "TwoArgumentsCalled", logs: logs, sub: sub}, nil
}

// WatchTwoArgumentsCalled is a free log subscription operation binding the contract event 0xc983106aca50fad459fb18ede1630e8ff8147ff28ad451a856427931fd7f15e3.
//
// Solidity: event TwoArgumentsCalled(address[] argumentOne, bytes4 argumentTwo)
func (_TwoArguments *TwoArgumentsFilterer) WatchTwoArgumentsCalled(opts *bind.WatchOpts, sink chan<- *TwoArgumentsTwoArgumentsCalled) (event.Subscription, error) {

	logs, sub, err := _TwoArguments.contract.WatchLogs(opts, "TwoArgumentsCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoArgumentsTwoArgumentsCalled)
				if err := _TwoArguments.contract.UnpackLog(event, "TwoArgumentsCalled", log); err != nil {
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

// ParseTwoArgumentsCalled is a log parse operation binding the contract event 0xc983106aca50fad459fb18ede1630e8ff8147ff28ad451a856427931fd7f15e3.
//
// Solidity: event TwoArgumentsCalled(address[] argumentOne, bytes4 argumentTwo)
func (_TwoArguments *TwoArgumentsFilterer) ParseTwoArgumentsCalled(log types.Log) (*TwoArgumentsTwoArgumentsCalled, error) {
	event := new(TwoArgumentsTwoArgumentsCalled)
	if err := _TwoArguments.contract.UnpackLog(event, "TwoArgumentsCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
