// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CentrifugeAssetHandler

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

// CentrifugeAssetHandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type CentrifugeAssetHandlerDepositRecord struct {
	OriginChainContractAddress     common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationRecipientAddress    common.Address
	Depositer                      common.Address
	MetaDataHash                   [32]byte
}

// CentrifugeAssetHandlerABI is the input ABI used to generate the binding from.
const CentrifugeAssetHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"originChainContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metaDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structCentrifugeAssetHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CentrifugeAssetHandlerBin is the compiled bytecode used for deploying new contracts.
var CentrifugeAssetHandlerBin = "0x608060405234801561001057600080fd5b50604051610e1f380380610e1f8339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b600081519050610087816100e8565b92915050565b60006020828403121561009f57600080fd5b60006100ad84828501610078565b91505092915050565b60006100c1826100c8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6100f1816100b6565b81146100fc57600080fd5b50565b610d118061010e6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063318c136e1461005c5780633cf5040a1461007a578063cb65d221146100aa578063db95e75c146100c6578063fc9539cd146100f6575b600080fd5b610064610112565b6040516100719190610b1a565b60405180910390f35b610094600480360381019061008f919061083a565b610137565b6040516100a19190610b35565b60405180910390f35b6100c460048036038101906100bf91906108cd565b610161565b005b6100e060048036038101906100db91906108a4565b61044e565b6040516100ed9190610bb0565b60405180910390f35b610110600480360381019061010b9190610863565b6105e7565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006002600083815260200190815260200160002060009054906101000a900460ff169050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e790610b50565b60405180910390fd5b60008060008060208501519350604085015192506060850151915060808501519050600015156002600083815260200190815260200160002060009054906101000a900460ff16151514610279576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027090610b90565b60405180910390fd5b6040518060c001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff168152602001828152506001600089815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a082015181600501559050505050505050505050565b610456610716565b600160008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610676576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066d90610b50565b60405180910390fd5b600060208201519050600015156002600083815260200190815260200160002060009054906101000a900460ff161515146106e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106dd90610b70565b60405180910390fd5b60016002600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600080191681525090565b6000813590506107b681610c96565b92915050565b6000813590506107cb81610cad565b92915050565b600082601f8301126107e257600080fd5b81356107f56107f082610bf8565b610bcb565b9150808252602083016020830185838301111561081157600080fd5b61081c838284610c87565b50505092915050565b60008135905061083481610cc4565b92915050565b60006020828403121561084c57600080fd5b600061085a848285016107bc565b91505092915050565b60006020828403121561087557600080fd5b600082013567ffffffffffffffff81111561088f57600080fd5b61089b848285016107d1565b91505092915050565b6000602082840312156108b657600080fd5b60006108c484828501610825565b91505092915050565b600080600080608085870312156108e357600080fd5b60006108f187828801610825565b945050602061090287828801610825565b9350506040610913878288016107a7565b925050606085013567ffffffffffffffff81111561093057600080fd5b61093c878288016107d1565b91505092959194509250565b61095181610c35565b82525050565b61096081610c35565b82525050565b61096f81610c47565b82525050565b61097e81610c53565b82525050565b6000610991601e83610c24565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b60006109d1603983610c24565b91507f6173736574206861736e2774206265656e206465706f7369746564206f72206860008301527f617320616c7265616479206265656e2066696e616c697a6564000000000000006020830152604082019050919050565b6000610a37603683610c24565b91507f61737365742068617320616c7265616479206265656e206465706f736974656460008301527f20616e642063616e6e6f74206265206368616e676564000000000000000000006020830152604082019050919050565b60c082016000820151610aa66000850182610948565b506020820151610ab96020850182610b0b565b506040820151610acc6040850182610948565b506060820151610adf6060850182610948565b506080820151610af26080850182610948565b5060a0820151610b0560a0850182610975565b50505050565b610b1481610c7d565b82525050565b6000602082019050610b2f6000830184610957565b92915050565b6000602082019050610b4a6000830184610966565b92915050565b60006020820190508181036000830152610b6981610984565b9050919050565b60006020820190508181036000830152610b89816109c4565b9050919050565b60006020820190508181036000830152610ba981610a2a565b9050919050565b600060c082019050610bc56000830184610a90565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610bee57600080fd5b8060405250919050565b600067ffffffffffffffff821115610c0f57600080fd5b601f19601f8301169050602081019050919050565b600082825260208201905092915050565b6000610c4082610c5d565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b610c9f81610c35565b8114610caa57600080fd5b50565b610cb681610c53565b8114610cc157600080fd5b50565b610ccd81610c7d565b8114610cd857600080fd5b5056fea264697066735822122072faff63a2b6d58e875d908788a41b670cfb71325b82f5f8c0fb8102590daf0c64736f6c63430006040033"

// DeployCentrifugeAssetHandler deploys a new Ethereum contract, binding an instance of CentrifugeAssetHandler to it.
func DeployCentrifugeAssetHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *CentrifugeAssetHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(CentrifugeAssetHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CentrifugeAssetHandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CentrifugeAssetHandler{CentrifugeAssetHandlerCaller: CentrifugeAssetHandlerCaller{contract: contract}, CentrifugeAssetHandlerTransactor: CentrifugeAssetHandlerTransactor{contract: contract}, CentrifugeAssetHandlerFilterer: CentrifugeAssetHandlerFilterer{contract: contract}}, nil
}

// CentrifugeAssetHandler is an auto generated Go binding around an Ethereum contract.
type CentrifugeAssetHandler struct {
	CentrifugeAssetHandlerCaller     // Read-only binding to the contract
	CentrifugeAssetHandlerTransactor // Write-only binding to the contract
	CentrifugeAssetHandlerFilterer   // Log filterer for contract events
}

// CentrifugeAssetHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CentrifugeAssetHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CentrifugeAssetHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CentrifugeAssetHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CentrifugeAssetHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CentrifugeAssetHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CentrifugeAssetHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CentrifugeAssetHandlerSession struct {
	Contract     *CentrifugeAssetHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CentrifugeAssetHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CentrifugeAssetHandlerCallerSession struct {
	Contract *CentrifugeAssetHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// CentrifugeAssetHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CentrifugeAssetHandlerTransactorSession struct {
	Contract     *CentrifugeAssetHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// CentrifugeAssetHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CentrifugeAssetHandlerRaw struct {
	Contract *CentrifugeAssetHandler // Generic contract binding to access the raw methods on
}

// CentrifugeAssetHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CentrifugeAssetHandlerCallerRaw struct {
	Contract *CentrifugeAssetHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// CentrifugeAssetHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CentrifugeAssetHandlerTransactorRaw struct {
	Contract *CentrifugeAssetHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCentrifugeAssetHandler creates a new instance of CentrifugeAssetHandler, bound to a specific deployed contract.
func NewCentrifugeAssetHandler(address common.Address, backend bind.ContractBackend) (*CentrifugeAssetHandler, error) {
	contract, err := bindCentrifugeAssetHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CentrifugeAssetHandler{CentrifugeAssetHandlerCaller: CentrifugeAssetHandlerCaller{contract: contract}, CentrifugeAssetHandlerTransactor: CentrifugeAssetHandlerTransactor{contract: contract}, CentrifugeAssetHandlerFilterer: CentrifugeAssetHandlerFilterer{contract: contract}}, nil
}

// NewCentrifugeAssetHandlerCaller creates a new read-only instance of CentrifugeAssetHandler, bound to a specific deployed contract.
func NewCentrifugeAssetHandlerCaller(address common.Address, caller bind.ContractCaller) (*CentrifugeAssetHandlerCaller, error) {
	contract, err := bindCentrifugeAssetHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CentrifugeAssetHandlerCaller{contract: contract}, nil
}

// NewCentrifugeAssetHandlerTransactor creates a new write-only instance of CentrifugeAssetHandler, bound to a specific deployed contract.
func NewCentrifugeAssetHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*CentrifugeAssetHandlerTransactor, error) {
	contract, err := bindCentrifugeAssetHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CentrifugeAssetHandlerTransactor{contract: contract}, nil
}

// NewCentrifugeAssetHandlerFilterer creates a new log filterer instance of CentrifugeAssetHandler, bound to a specific deployed contract.
func NewCentrifugeAssetHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*CentrifugeAssetHandlerFilterer, error) {
	contract, err := bindCentrifugeAssetHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CentrifugeAssetHandlerFilterer{contract: contract}, nil
}

// bindCentrifugeAssetHandler binds a generic wrapper to an already deployed contract.
func bindCentrifugeAssetHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CentrifugeAssetHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CentrifugeAssetHandler.Contract.CentrifugeAssetHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.Contract.CentrifugeAssetHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.Contract.CentrifugeAssetHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CentrifugeAssetHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() constant returns(address)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CentrifugeAssetHandler.contract.Call(opts, out, "_bridgeAddress")
	return *ret0, err
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() constant returns(address)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerSession) BridgeAddress() (common.Address, error) {
	return _CentrifugeAssetHandler.Contract.BridgeAddress(&_CentrifugeAssetHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() constant returns(address)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _CentrifugeAssetHandler.Contract.BridgeAddress(&_CentrifugeAssetHandler.CallOpts)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) constant returns(CentrifugeAssetHandlerDepositRecord)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositID *big.Int) (CentrifugeAssetHandlerDepositRecord, error) {
	var (
		ret0 = new(CentrifugeAssetHandlerDepositRecord)
	)
	out := ret0
	err := _CentrifugeAssetHandler.contract.Call(opts, out, "getDepositRecord", depositID)
	return *ret0, err
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) constant returns(CentrifugeAssetHandlerDepositRecord)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerSession) GetDepositRecord(depositID *big.Int) (CentrifugeAssetHandlerDepositRecord, error) {
	return _CentrifugeAssetHandler.Contract.GetDepositRecord(&_CentrifugeAssetHandler.CallOpts, depositID)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) constant returns(CentrifugeAssetHandlerDepositRecord)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerCallerSession) GetDepositRecord(depositID *big.Int) (CentrifugeAssetHandlerDepositRecord, error) {
	return _CentrifugeAssetHandler.Contract.GetDepositRecord(&_CentrifugeAssetHandler.CallOpts, depositID)
}

// GetHash is a free data retrieval call binding the contract method 0x3cf5040a.
//
// Solidity: function getHash(bytes32 hash) constant returns(bool)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerCaller) GetHash(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CentrifugeAssetHandler.contract.Call(opts, out, "getHash", hash)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0x3cf5040a.
//
// Solidity: function getHash(bytes32 hash) constant returns(bool)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerSession) GetHash(hash [32]byte) (bool, error) {
	return _CentrifugeAssetHandler.Contract.GetHash(&_CentrifugeAssetHandler.CallOpts, hash)
}

// GetHash is a free data retrieval call binding the contract method 0x3cf5040a.
//
// Solidity: function getHash(bytes32 hash) constant returns(bool)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerCallerSession) GetHash(hash [32]byte) (bool, error) {
	return _CentrifugeAssetHandler.Contract.GetHash(&_CentrifugeAssetHandler.CallOpts, hash)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 originChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerTransactor) Deposit(opts *bind.TransactOpts, originChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.contract.Transact(opts, "deposit", originChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 originChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerSession) Deposit(originChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.Contract.Deposit(&_CentrifugeAssetHandler.TransactOpts, originChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 originChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerTransactorSession) Deposit(originChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.Contract.Deposit(&_CentrifugeAssetHandler.TransactOpts, originChainID, depositNonce, depositer, data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerTransactor) ExecuteDeposit(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.contract.Transact(opts, "executeDeposit", data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerSession) ExecuteDeposit(data []byte) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.Contract.ExecuteDeposit(&_CentrifugeAssetHandler.TransactOpts, data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerTransactorSession) ExecuteDeposit(data []byte) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.Contract.ExecuteDeposit(&_CentrifugeAssetHandler.TransactOpts, data)
}
var RuntimeBytecode = "0x608060405234801561001057600080fd5b50600436106100575760003560e01c8063318c136e1461005c5780633cf5040a1461007a578063cb65d221146100aa578063db95e75c146100c6578063fc9539cd146100f6575b600080fd5b610064610112565b6040516100719190610b1a565b60405180910390f35b610094600480360381019061008f919061083a565b610137565b6040516100a19190610b35565b60405180910390f35b6100c460048036038101906100bf91906108cd565b610161565b005b6100e060048036038101906100db91906108a4565b61044e565b6040516100ed9190610bb0565b60405180910390f35b610110600480360381019061010b9190610863565b6105e7565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006002600083815260200190815260200160002060009054906101000a900460ff169050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e790610b50565b60405180910390fd5b60008060008060208501519350604085015192506060850151915060808501519050600015156002600083815260200190815260200160002060009054906101000a900460ff16151514610279576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027090610b90565b60405180910390fd5b6040518060c001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff168152602001828152506001600089815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a082015181600501559050505050505050505050565b610456610716565b600160008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610676576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066d90610b50565b60405180910390fd5b600060208201519050600015156002600083815260200190815260200160002060009054906101000a900460ff161515146106e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106dd90610b70565b60405180910390fd5b60016002600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600080191681525090565b6000813590506107b681610c96565b92915050565b6000813590506107cb81610cad565b92915050565b600082601f8301126107e257600080fd5b81356107f56107f082610bf8565b610bcb565b9150808252602083016020830185838301111561081157600080fd5b61081c838284610c87565b50505092915050565b60008135905061083481610cc4565b92915050565b60006020828403121561084c57600080fd5b600061085a848285016107bc565b91505092915050565b60006020828403121561087557600080fd5b600082013567ffffffffffffffff81111561088f57600080fd5b61089b848285016107d1565b91505092915050565b6000602082840312156108b657600080fd5b60006108c484828501610825565b91505092915050565b600080600080608085870312156108e357600080fd5b60006108f187828801610825565b945050602061090287828801610825565b9350506040610913878288016107a7565b925050606085013567ffffffffffffffff81111561093057600080fd5b61093c878288016107d1565b91505092959194509250565b61095181610c35565b82525050565b61096081610c35565b82525050565b61096f81610c47565b82525050565b61097e81610c53565b82525050565b6000610991601e83610c24565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b60006109d1603983610c24565b91507f6173736574206861736e2774206265656e206465706f7369746564206f72206860008301527f617320616c7265616479206265656e2066696e616c697a6564000000000000006020830152604082019050919050565b6000610a37603683610c24565b91507f61737365742068617320616c7265616479206265656e206465706f736974656460008301527f20616e642063616e6e6f74206265206368616e676564000000000000000000006020830152604082019050919050565b60c082016000820151610aa66000850182610948565b506020820151610ab96020850182610b0b565b506040820151610acc6040850182610948565b506060820151610adf6060850182610948565b506080820151610af26080850182610948565b5060a0820151610b0560a0850182610975565b50505050565b610b1481610c7d565b82525050565b6000602082019050610b2f6000830184610957565b92915050565b6000602082019050610b4a6000830184610966565b92915050565b60006020820190508181036000830152610b6981610984565b9050919050565b60006020820190508181036000830152610b89816109c4565b9050919050565b60006020820190508181036000830152610ba981610a2a565b9050919050565b600060c082019050610bc56000830184610a90565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610bee57600080fd5b8060405250919050565b600067ffffffffffffffff821115610c0f57600080fd5b601f19601f8301169050602081019050919050565b600082825260208201905092915050565b6000610c4082610c5d565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b610c9f81610c35565b8114610caa57600080fd5b50565b610cb681610c53565b8114610cc157600080fd5b50565b610ccd81610c7d565b8114610cd857600080fd5b5056fea264697066735822122072faff63a2b6d58e875d908788a41b670cfb71325b82f5f8c0fb8102590daf0c64736f6c63430006040033"
