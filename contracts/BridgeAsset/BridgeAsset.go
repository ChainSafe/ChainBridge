// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package BridgeAsset

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// BridgeAssetABI is the input ABI used to generate the binding from.
const BridgeAssetABI = `[{"inputs":[{"internalType":"uint8","name":"mc","type":"uint8"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"asset","type":"bytes32"}],"name":"AssetStored","type":"event"},{"constant":true,"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"assets","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"bytes32","name":"asset","type":"bytes32"}],"name":"store","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"internalType":"bytes32","name":"asset","type":"bytes32"}],"name":"isAssetValid","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"}]`

// BridgeAssetBin is the compiled bytecode used for deploying new contracts.
const BridgeAssetBin = `608060405234801561001057600080fd5b506040516104c93803806104c983398181016040526100329190810190610067565b806000806101000a81548160ff021916908360ff160217905550506100b4565b6000815190506100618161009d565b92915050565b60006020828403121561007957600080fd5b600061008784828501610052565b91505092915050565b600060ff82169050919050565b6100a681610090565b81146100b157600080fd5b50565b610406806100c36000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063339f16a914610046578063654cf88c146100765780639fda5b6614610092575b600080fd5b610060600480360361005b9190810190610275565b6100c2565b60405161006d9190610322565b60405180910390f35b610090600480360361008b9190810190610275565b610102565b005b6100ac60048036036100a79190810190610275565b610240565b6040516100b9919061035d565b60405180910390f35b6000600180600084815260200190815260200160002060009054906101000a900460ff1660ff1614156100f857600190506100fd565b600090505b919050565b600180600083815260200190815260200160002060009054906101000a900460ff1660ff161415610168576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015f9061033d565b60405180910390fd5b600a6001600083815260200190815260200160002060008282829054906101000a900460ff160192506101000a81548160ff021916908360ff1602179055506000809054906101000a900460ff1660ff166001600083815260200190815260200160002060009054906101000a900460ff1660ff16141561023d57600180600083815260200190815260200160002060006101000a81548160ff021916908360ff160217905550807f08ae553713effae7116be03743b167b8b803449ee8fb912c2ec43dc2c824f53560405160405180910390a25b50565b60016020528060005260406000206000915054906101000a900460ff1681565b60008135905061026f816103ac565b92915050565b60006020828403121561028757600080fd5b600061029584828501610260565b91505092915050565b6102a781610389565b82525050565b60006102ba602683610378565b91507f41737365742063616e6e6f74206265206368616e676564206f6e636520636f6e60008301527f6669726d656400000000000000000000000000000000000000000000000000006020830152604082019050919050565b61031c8161039f565b82525050565b6000602082019050610337600083018461029e565b92915050565b60006020820190508181036000830152610356816102ad565b9050919050565b60006020820190506103726000830184610313565b92915050565b600082825260208201905092915050565b60008115159050919050565b6000819050919050565b600060ff82169050919050565b6103b581610395565b81146103c057600080fd5b5056fea365627a7a723158204d400c339e522dca1a2e44bdc1bc8982a4509a7ab9bfefe346f267c7dac7a9666c6578706572696d656e74616cf564736f6c634300050c0040`

// DeployBridgeAsset deploys a new Ethereum contract, binding an instance of BridgeAsset to it.
func DeployBridgeAsset(auth *bind.TransactOpts, backend bind.ContractBackend, mc uint8) (common.Address, *types.Transaction, *BridgeAsset, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeAssetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeAssetBin), backend, mc)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeAsset{BridgeAssetCaller: BridgeAssetCaller{contract: contract}, BridgeAssetTransactor: BridgeAssetTransactor{contract: contract}}, nil
}

// BridgeAsset is an auto generated Go binding around an Ethereum contract.
type BridgeAsset struct {
	BridgeAssetCaller     // Read-only binding to the contract
	BridgeAssetTransactor // Write-only binding to the contract
}

// BridgeAssetCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeAssetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeAssetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeAssetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeAssetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeAssetSession struct {
	Contract     *BridgeAsset      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeAssetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeAssetCallerSession struct {
	Contract *BridgeAssetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BridgeAssetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeAssetTransactorSession struct {
	Contract     *BridgeAssetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BridgeAssetRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeAssetRaw struct {
	Contract *BridgeAsset // Generic contract binding to access the raw methods on
}

// BridgeAssetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeAssetCallerRaw struct {
	Contract *BridgeAssetCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeAssetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeAssetTransactorRaw struct {
	Contract *BridgeAssetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeAsset creates a new instance of BridgeAsset, bound to a specific deployed contract.
func NewBridgeAsset(address common.Address, backend bind.ContractBackend) (*BridgeAsset, error) {
	contract, err := bindBridgeAsset(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &BridgeAsset{BridgeAssetCaller: BridgeAssetCaller{contract: contract}, BridgeAssetTransactor: BridgeAssetTransactor{contract: contract}}, nil
}

// NewBridgeAssetCaller creates a new read-only instance of BridgeAsset, bound to a specific deployed contract.
func NewBridgeAssetCaller(address common.Address, caller bind.ContractCaller) (*BridgeAssetCaller, error) {
	contract, err := bindBridgeAsset(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeAssetCaller{contract: contract}, nil
}

// NewBridgeAssetTransactor creates a new write-only instance of BridgeAsset, bound to a specific deployed contract.
func NewBridgeAssetTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeAssetTransactor, error) {
	contract, err := bindBridgeAsset(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BridgeAssetTransactor{contract: contract}, nil
}

// bindBridgeAsset binds a generic wrapper to an already deployed contract.
func bindBridgeAsset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeAssetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeAsset *BridgeAssetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeAsset.Contract.BridgeAssetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeAsset *BridgeAssetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeAsset.Contract.BridgeAssetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeAsset *BridgeAssetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeAsset.Contract.BridgeAssetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeAsset *BridgeAssetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeAsset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeAsset *BridgeAssetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeAsset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeAsset *BridgeAssetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeAsset.Contract.contract.Transact(opts, method, params...)
}

// Assets is a free data retrieval call binding the contract method 0x9fda5b66.
//
// Solidity: function assets( bytes32) constant returns(uint8)
func (_BridgeAsset *BridgeAssetCaller) Assets(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BridgeAsset.contract.Call(opts, out, "assets", arg0)
	return *ret0, err
}

// Assets is a free data retrieval call binding the contract method 0x9fda5b66.
//
// Solidity: function assets( bytes32) constant returns(uint8)
func (_BridgeAsset *BridgeAssetSession) Assets(arg0 [32]byte) (uint8, error) {
	return _BridgeAsset.Contract.Assets(&_BridgeAsset.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0x9fda5b66.
//
// Solidity: function assets( bytes32) constant returns(uint8)
func (_BridgeAsset *BridgeAssetCallerSession) Assets(arg0 [32]byte) (uint8, error) {
	return _BridgeAsset.Contract.Assets(&_BridgeAsset.CallOpts, arg0)
}

// IsAssetValid is a free data retrieval call binding the contract method 0x339f16a9.
//
// Solidity: function isAssetValid(asset bytes32) constant returns(bool)
func (_BridgeAsset *BridgeAssetCaller) IsAssetValid(opts *bind.CallOpts, asset [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeAsset.contract.Call(opts, out, "isAssetValid", asset)
	return *ret0, err
}

// IsAssetValid is a free data retrieval call binding the contract method 0x339f16a9.
//
// Solidity: function isAssetValid(asset bytes32) constant returns(bool)
func (_BridgeAsset *BridgeAssetSession) IsAssetValid(asset [32]byte) (bool, error) {
	return _BridgeAsset.Contract.IsAssetValid(&_BridgeAsset.CallOpts, asset)
}

// IsAssetValid is a free data retrieval call binding the contract method 0x339f16a9.
//
// Solidity: function isAssetValid(asset bytes32) constant returns(bool)
func (_BridgeAsset *BridgeAssetCallerSession) IsAssetValid(asset [32]byte) (bool, error) {
	return _BridgeAsset.Contract.IsAssetValid(&_BridgeAsset.CallOpts, asset)
}

// Store is a paid mutator transaction binding the contract method 0x654cf88c.
//
// Solidity: function store(asset bytes32) returns()
func (_BridgeAsset *BridgeAssetTransactor) Store(opts *bind.TransactOpts, asset [32]byte) (*types.Transaction, error) {
	return _BridgeAsset.contract.Transact(opts, "store", asset)
}

// Store is a paid mutator transaction binding the contract method 0x654cf88c.
//
// Solidity: function store(asset bytes32) returns()
func (_BridgeAsset *BridgeAssetSession) Store(asset [32]byte) (*types.Transaction, error) {
	return _BridgeAsset.Contract.Store(&_BridgeAsset.TransactOpts, asset)
}

// Store is a paid mutator transaction binding the contract method 0x654cf88c.
//
// Solidity: function store(asset bytes32) returns()
func (_BridgeAsset *BridgeAssetTransactorSession) Store(asset [32]byte) (*types.Transaction, error) {
	return _BridgeAsset.Contract.Store(&_BridgeAsset.TransactOpts, asset)
}
