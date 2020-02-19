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
const EmitterABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_destChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"ERCTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_destChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"GenericTransfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_destChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_destChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"depositGenericErc\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"lockErc\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"releaseErc\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// Lock is a paid mutator transaction binding the contract method 0x3fea56b8.
//
// Solidity: function lock(_tokenAddress address, _value uint256, _to address, _from address) returns()
func (_Emitter *EmitterTransactor) Lock(opts *bind.TransactOpts, _tokenAddress common.Address, _value *big.Int, _to common.Address, _from common.Address) (*types.Transaction, error) {
	return _Emitter.contract.Transact(opts, "lock", _tokenAddress, _value, _to, _from)
}

// Lock is a paid mutator transaction binding the contract method 0x3fea56b8.
//
// Solidity: function lock(_tokenAddress address, _value uint256, _to address, _from address) returns()
func (_Emitter *EmitterSession) Lock(_tokenAddress common.Address, _value *big.Int, _to common.Address, _from common.Address) (*types.Transaction, error) {
	return _Emitter.Contract.Lock(&_Emitter.TransactOpts, _tokenAddress, _value, _to, _from)
}

// Lock is a paid mutator transaction binding the contract method 0x3fea56b8.
//
// Solidity: function lock(_tokenAddress address, _value uint256, _to address, _from address) returns()
func (_Emitter *EmitterTransactorSession) Lock(_tokenAddress common.Address, _value *big.Int, _to common.Address, _from common.Address) (*types.Transaction, error) {
	return _Emitter.Contract.Lock(&_Emitter.TransactOpts, _tokenAddress, _value, _to, _from)
}

// LockErc is a paid mutator transaction binding the contract method 0x6d3cdd97.
//
// Solidity: function lockErc(_tokenAddress address, _value uint256, _to address, _from address) returns()
func (_Emitter *EmitterTransactor) LockErc(opts *bind.TransactOpts, _tokenAddress common.Address, _value *big.Int, _to common.Address, _from common.Address) (*types.Transaction, error) {
	return _Emitter.contract.Transact(opts, "lockErc", _tokenAddress, _value, _to, _from)
}

// LockErc is a paid mutator transaction binding the contract method 0x6d3cdd97.
//
// Solidity: function lockErc(_tokenAddress address, _value uint256, _to address, _from address) returns()
func (_Emitter *EmitterSession) LockErc(_tokenAddress common.Address, _value *big.Int, _to common.Address, _from common.Address) (*types.Transaction, error) {
	return _Emitter.Contract.LockErc(&_Emitter.TransactOpts, _tokenAddress, _value, _to, _from)
}

// LockErc is a paid mutator transaction binding the contract method 0x6d3cdd97.
//
// Solidity: function lockErc(_tokenAddress address, _value uint256, _to address, _from address) returns()
func (_Emitter *EmitterTransactorSession) LockErc(_tokenAddress common.Address, _value *big.Int, _to common.Address, _from common.Address) (*types.Transaction, error) {
	return _Emitter.Contract.LockErc(&_Emitter.TransactOpts, _tokenAddress, _value, _to, _from)
}

// Release is a paid mutator transaction binding the contract method 0x15f41501.
//
// Solidity: function release(_tokenAddress address, _value uint256, _to address) returns()
func (_Emitter *EmitterTransactor) Release(opts *bind.TransactOpts, _tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Emitter.contract.Transact(opts, "release", _tokenAddress, _value, _to)
}

// Release is a paid mutator transaction binding the contract method 0x15f41501.
//
// Solidity: function release(_tokenAddress address, _value uint256, _to address) returns()
func (_Emitter *EmitterSession) Release(_tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Emitter.Contract.Release(&_Emitter.TransactOpts, _tokenAddress, _value, _to)
}

// Release is a paid mutator transaction binding the contract method 0x15f41501.
//
// Solidity: function release(_tokenAddress address, _value uint256, _to address) returns()
func (_Emitter *EmitterTransactorSession) Release(_tokenAddress common.Address, _value *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Emitter.Contract.Release(&_Emitter.TransactOpts, _tokenAddress, _value, _to)
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
// Solidity: e ERCTransfer(_destChain uint256, _depositId uint256, _to address, _amount uint256, _tokenAddress address)
func (_Emitter *EmitterFilterer) FilterERCTransfer(opts *bind.FilterOpts) (*EmitterERCTransferIterator, error) {

	logs, sub, err := _Emitter.contract.FilterLogs(opts, "ERCTransfer")
	if err != nil {
		return nil, err
	}
	return &EmitterERCTransferIterator{contract: _Emitter.contract, event: "ERCTransfer", logs: logs, sub: sub}, nil
}

// WatchERCTransfer is a free log subscription operation binding the contract event 0x12a20368a5b6a9eb35068693542bef5613d39588e50c62fca6312f1b7d9242fe.
//
// Solidity: e ERCTransfer(_destChain uint256, _depositId uint256, _to address, _amount uint256, _tokenAddress address)
func (_Emitter *EmitterFilterer) WatchERCTransfer(opts *bind.WatchOpts, sink chan<- *EmitterERCTransfer) (event.Subscription, error) {

	logs, sub, err := _Emitter.contract.WatchLogs(opts, "ERCTransfer")
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
// Solidity: e GenericTransfer(_destChain uint256, _depositId uint256, _destAddress address, _data bytes)
func (_Emitter *EmitterFilterer) FilterGenericTransfer(opts *bind.FilterOpts) (*EmitterGenericTransferIterator, error) {

	logs, sub, err := _Emitter.contract.FilterLogs(opts, "GenericTransfer")
	if err != nil {
		return nil, err
	}
	return &EmitterGenericTransferIterator{contract: _Emitter.contract, event: "GenericTransfer", logs: logs, sub: sub}, nil
}

// WatchGenericTransfer is a free log subscription operation binding the contract event 0x6cc01912ea1b73308485cfdf6b52b3a8c1e8f33825352b5d5fc0b5ddec3a2a2b.
//
// Solidity: e GenericTransfer(_destChain uint256, _depositId uint256, _destAddress address, _data bytes)
func (_Emitter *EmitterFilterer) WatchGenericTransfer(opts *bind.WatchOpts, sink chan<- *EmitterGenericTransfer) (event.Subscription, error) {

	logs, sub, err := _Emitter.contract.WatchLogs(opts, "GenericTransfer")
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
