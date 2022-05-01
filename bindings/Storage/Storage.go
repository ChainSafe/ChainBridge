// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igloo

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"PutData\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"dataStorage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"putData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"65244594": "dataStorage(address,string)",
		"77236939": "putData(address,string,bytes)",
	},
	Bin: "0x60c06040526002608081905261060f60f31b60a09081526100239160009190610036565b5034801561003057600080fd5b5061010a565b828054610042906100cf565b90600052602060002090601f01602090048101928261006457600085556100aa565b82601f1061007d57805160ff19168380011785556100aa565b828001600101855582156100aa579182015b828111156100aa57825182559160200191906001019061008f565b506100b69291506100ba565b5090565b5b808211156100b657600081556001016100bb565b600181811c908216806100e357607f821691505b6020821081141561010457634e487b7160e01b600052602260045260246000fd5b50919050565b6105b9806101196000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063652445941461003b5780637723693914610064575b600080fd5b61004e6100493660046103bb565b610079565b60405161005b9190610465565b60405180910390f35b610077610072366004610478565b61012f565b005b60016020908152600092835260409092208151808301840180519281529084019290930191909120915280546100ae90610500565b80601f01602080910402602001604051908101604052809291908181526020018280546100da90610500565b80156101275780601f106100fc57610100808354040283529160200191610127565b820191906000526020600020905b81548152906001019060200180831161010a57829003601f168201915b505050505081565b6000805461013c90610500565b6001600160a01b0385166000908152600160205260409081902090519192509061016790859061053b565b9081526020016040518091039020805461018090610500565b9050146101c95760405162461bcd60e51b815260206004820152601360248201527263616e2774206f76657272696465206461746160681b604482015260640160405180910390fd5b6001600160a01b0383166000908152600160205260409081902090518291906101f390859061053b565b90815260200160405180910390209080519060200190610214929190610253565b507f6e4d9a64feabe0b2707f40883945812e71d177fd5b5c682b726ccba233bf575b8383604051610246929190610557565b60405180910390a1505050565b82805461025f90610500565b90600052602060002090601f01602090048101928261028157600085556102c7565b82601f1061029a57805160ff19168380011785556102c7565b828001600101855582156102c7579182015b828111156102c75782518255916020019190600101906102ac565b506102d39291506102d7565b5090565b5b808211156102d357600081556001016102d8565b80356001600160a01b038116811461030357600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff8084111561033957610339610308565b604051601f8501601f19908116603f0116810190828211818310171561036157610361610308565b8160405280935085815286868601111561037a57600080fd5b858560208301376000602087830101525050509392505050565b600082601f8301126103a557600080fd5b6103b48383356020850161031e565b9392505050565b600080604083850312156103ce57600080fd5b6103d7836102ec565b9150602083013567ffffffffffffffff8111156103f357600080fd5b6103ff85828601610394565b9150509250929050565b60005b8381101561042457818101518382015260200161040c565b83811115610433576000848401525b50505050565b60008151808452610451816020860160208601610409565b601f01601f19169290920160200192915050565b6020815260006103b46020830184610439565b60008060006060848603121561048d57600080fd5b610496846102ec565b9250602084013567ffffffffffffffff808211156104b357600080fd5b6104bf87838801610394565b935060408601359150808211156104d557600080fd5b508401601f810186136104e757600080fd5b6104f68682356020840161031e565b9150509250925092565b600181811c9082168061051457607f821691505b6020821081141561053557634e487b7160e01b600052602260045260246000fd5b50919050565b6000825161054d818460208701610409565b9190910192915050565b6001600160a01b038316815260406020820181905260009061057b90830184610439565b94935050505056fea2646970667358221220abda6fd148a8fb0cdd1ab6682517bb031c7c6b8af2f94e2342228ef5760cf59d64736f6c63430008090033",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// Deprecated: Use StorageMetaData.Sigs instead.
// StorageFuncSigs maps the 4-byte function signature to its string representation.
var StorageFuncSigs = StorageMetaData.Sigs

// StorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageMetaData.Bin instead.
var StorageBin = StorageMetaData.Bin

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// DataStorage is a free data retrieval call binding the contract method 0x65244594.
//
// Solidity: function dataStorage(address , string ) view returns(bytes)
func (_Storage *StorageCaller) DataStorage(opts *bind.CallOpts, arg0 common.Address, arg1 string) ([]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "dataStorage", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DataStorage is a free data retrieval call binding the contract method 0x65244594.
//
// Solidity: function dataStorage(address , string ) view returns(bytes)
func (_Storage *StorageSession) DataStorage(arg0 common.Address, arg1 string) ([]byte, error) {
	return _Storage.Contract.DataStorage(&_Storage.CallOpts, arg0, arg1)
}

// DataStorage is a free data retrieval call binding the contract method 0x65244594.
//
// Solidity: function dataStorage(address , string ) view returns(bytes)
func (_Storage *StorageCallerSession) DataStorage(arg0 common.Address, arg1 string) ([]byte, error) {
	return _Storage.Contract.DataStorage(&_Storage.CallOpts, arg0, arg1)
}

// PutData is a paid mutator transaction binding the contract method 0x77236939.
//
// Solidity: function putData(address userAddress, string key, bytes value) returns()
func (_Storage *StorageTransactor) PutData(opts *bind.TransactOpts, userAddress common.Address, key string, value []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "putData", userAddress, key, value)
}

// PutData is a paid mutator transaction binding the contract method 0x77236939.
//
// Solidity: function putData(address userAddress, string key, bytes value) returns()
func (_Storage *StorageSession) PutData(userAddress common.Address, key string, value []byte) (*types.Transaction, error) {
	return _Storage.Contract.PutData(&_Storage.TransactOpts, userAddress, key, value)
}

// PutData is a paid mutator transaction binding the contract method 0x77236939.
//
// Solidity: function putData(address userAddress, string key, bytes value) returns()
func (_Storage *StorageTransactorSession) PutData(userAddress common.Address, key string, value []byte) (*types.Transaction, error) {
	return _Storage.Contract.PutData(&_Storage.TransactOpts, userAddress, key, value)
}

// StoragePutDataIterator is returned from FilterPutData and is used to iterate over the raw logs and unpacked data for PutData events raised by the Storage contract.
type StoragePutDataIterator struct {
	Event *StoragePutData // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoragePutDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoragePutData)
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
		it.Event = new(StoragePutData)
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
func (it *StoragePutDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoragePutDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoragePutData represents a PutData event raised by the Storage contract.
type StoragePutData struct {
	UserAddress common.Address
	Key         string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPutData is a free log retrieval operation binding the contract event 0x6e4d9a64feabe0b2707f40883945812e71d177fd5b5c682b726ccba233bf575b.
//
// Solidity: event PutData(address userAddress, string key)
func (_Storage *StorageFilterer) FilterPutData(opts *bind.FilterOpts) (*StoragePutDataIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "PutData")
	if err != nil {
		return nil, err
	}
	return &StoragePutDataIterator{contract: _Storage.contract, event: "PutData", logs: logs, sub: sub}, nil
}

// WatchPutData is a free log subscription operation binding the contract event 0x6e4d9a64feabe0b2707f40883945812e71d177fd5b5c682b726ccba233bf575b.
//
// Solidity: event PutData(address userAddress, string key)
func (_Storage *StorageFilterer) WatchPutData(opts *bind.WatchOpts, sink chan<- *StoragePutData) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "PutData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoragePutData)
				if err := _Storage.contract.UnpackLog(event, "PutData", log); err != nil {
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

// ParsePutData is a log parse operation binding the contract event 0x6e4d9a64feabe0b2707f40883945812e71d177fd5b5c682b726ccba233bf575b.
//
// Solidity: event PutData(address userAddress, string key)
func (_Storage *StorageFilterer) ParsePutData(log types.Log) (*StoragePutData, error) {
	event := new(StoragePutData)
	if err := _Storage.contract.UnpackLog(event, "PutData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
