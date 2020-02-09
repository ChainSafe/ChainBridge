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
const BridgeAssetABI = `[{"inputs":[{"internalType":"uint8","name":"mc","type":"uint8"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"asset","type":"bytes32"}],"name":"AssetStored","type":"event"},{"constant":true,"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"assets","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"bytes32","name":"asset","type":"bytes32"}],"name":"isAssetValid","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"bytes32","name":"asset","type":"bytes32"}],"name":"store","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`

// BridgeAssetBin is the compiled bytecode used for deploying new contracts.
const BridgeAssetBin = `60806040523480156100115760006000fd5b506040516105303803806105308339818101604052610033919081019061006f565b5b80600060006101000a81548160ff021916908360ff1602179055505b506100c1566100c0565b600081519050610069816100a6565b92915050565b6000602082840312156100825760006000fd5b60006100908482850161005a565b91505092915050565b600060ff82169050919050565b6100af81610099565b811415156100bd5760006000fd5b50565b5b610460806100d06000396000f3fe60806040523480156100115760006000fd5b50600436106100465760003560e01c8063339f16a91461004c578063654cf88c1461007c5780639fda5b661461009857610046565b60006000fd5b610066600480360361006191908101906102ca565b6100c8565b6040516100739190610378565b60405180910390f35b610096600480360361009191908101906102ca565b610118565b005b6100b260048036036100ad91908101906102ca565b61028c565b6040516100bf91906103b3565b60405180910390f35b6000600160016000506000846000191660001916815260200190815260200160002060009054906101000a900460ff1660ff16141561010a5760019050610113565b60009050610113565b919050565b600160016000506000836000191660001916815260200190815260200160002060009054906101000a900460ff1660ff161415151561018c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018390610393565b60405180910390fd5b600a60016000506000836000191660001916815260200190815260200160002060008282829054906101000a900460ff160192506101000a81548160ff021916908360ff160217905550600060009054906101000a900460ff1660ff1660016000506000836000191660001916815260200190815260200160002060009054906101000a900460ff1660ff16141561028857600160016000506000836000191660001916815260200190815260200160002060006101000a81548160ff021916908360ff16021790555080600019167f08ae553713effae7116be03743b167b8b803449ee8fb912c2ec43dc2c824f53560405160405180910390a25b5b50565b600160005060205280600052604060002060009150909054906101000a900460ff16815661041c565b6000813590506102c481610402565b92915050565b6000602082840312156102dd5760006000fd5b60006102eb848285016102b5565b91505092915050565b6102fd816103df565b82525050565b60006103106026836103ce565b91507f41737365742063616e6e6f74206265206368616e676564206f6e636520636f6e60008301527f6669726d656400000000000000000000000000000000000000000000000000006020830152604082019050919050565b610372816103f5565b82525050565b600060208201905061038d60008301846102f4565b92915050565b600060208201905081810360008301526103ac81610303565b9050919050565b60006020820190506103c86000830184610369565b92915050565b600082825260208201905092915050565b60008115159050919050565b6000819050919050565b600060ff82169050919050565b61040b816103eb565b811415156104195760006000fd5b50565bfea365627a7a72315820f86d46b1645a10f8d7ff7cdb35ef4fcefe8a632a10cd965df055acfd71e1486c6c6578706572696d656e74616cf564736f6c634300050c0040`

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
