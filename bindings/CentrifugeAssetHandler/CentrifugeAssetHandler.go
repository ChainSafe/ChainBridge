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
const CentrifugeAssetHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"originChainContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metaDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structCentrifugeAssetHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CentrifugeAssetHandlerBin is the compiled bytecode used for deploying new contracts.
var CentrifugeAssetHandlerBin = "0x608060405234801561001057600080fd5b50604051610d90380380610d908339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b600081519050610087816100e8565b92915050565b60006020828403121561009f57600080fd5b60006100ad84828501610078565b91505092915050565b60006100c1826100c8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6100f1816100b6565b81146100fc57600080fd5b50565b610c828061010e6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063318c136e14610051578063cb65d2211461006f578063db95e75c1461008b578063fc9539cd146100bb575b600080fd5b6100596100d7565b6040516100669190610ac9565b60405180910390f35b6100896004803603810190610084919061088b565b6100fc565b005b6100a560048036038101906100a09190610862565b610430565b6040516100b29190610b44565b60405180910390f35b6100d560048036038101906100d09190610821565b6105c9565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461018b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018290610ae4565b60405180910390fd5b60008060008060208501519350604085015192506060850151915060808501519050600060028111156101ba57fe5b6002600083815260200190815260200160002060009054906101000a900460ff1660028111156101e657fe5b14610226576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021d90610b24565b60405180910390fd5b60016002600083815260200190815260200160002060006101000a81548160ff0219169083600281111561025657fe5b02179055506040518060c001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff168152602001828152506001600089815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a082015181600501559050505050505050505050565b610438610712565b600160008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610658576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064f90610ae4565b60405180910390fd5b6000602082015190506001600281111561066e57fe5b6002600083815260200190815260200160002060009054906101000a900460ff16600281111561069a57fe5b146106da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d190610b04565b60405180910390fd5b600280600083815260200190815260200160002060006101000a81548160ff0219169083600281111561070957fe5b02179055505050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600080191681525090565b6000813590506107b281610c1e565b92915050565b600082601f8301126107c957600080fd5b81356107dc6107d782610b8c565b610b5f565b915080825260208301602083018583830111156107f857600080fd5b610803838284610c0f565b50505092915050565b60008135905061081b81610c35565b92915050565b60006020828403121561083357600080fd5b600082013567ffffffffffffffff81111561084d57600080fd5b610859848285016107b8565b91505092915050565b60006020828403121561087457600080fd5b60006108828482850161080c565b91505092915050565b600080600080608085870312156108a157600080fd5b60006108af8782880161080c565b94505060206108c08782880161080c565b93505060406108d1878288016107a3565b925050606085013567ffffffffffffffff8111156108ee57600080fd5b6108fa878288016107b8565b91505092959194509250565b61090f81610bc9565b82525050565b61091e81610bc9565b82525050565b61092d81610bdb565b82525050565b6000610940601e83610bb8565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000610980603983610bb8565b91507f6173736574206861736e2774206265656e206465706f7369746564206f72206860008301527f617320616c7265616479206265656e2066696e616c697a6564000000000000006020830152604082019050919050565b60006109e6603683610bb8565b91507f61737365742068617320616c7265616479206265656e206465706f736974656460008301527f20616e642063616e6e6f74206265206368616e676564000000000000000000006020830152604082019050919050565b60c082016000820151610a556000850182610906565b506020820151610a686020850182610aba565b506040820151610a7b6040850182610906565b506060820151610a8e6060850182610906565b506080820151610aa16080850182610906565b5060a0820151610ab460a0850182610924565b50505050565b610ac381610c05565b82525050565b6000602082019050610ade6000830184610915565b92915050565b60006020820190508181036000830152610afd81610933565b9050919050565b60006020820190508181036000830152610b1d81610973565b9050919050565b60006020820190508181036000830152610b3d816109d9565b9050919050565b600060c082019050610b596000830184610a3f565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610b8257600080fd5b8060405250919050565b600067ffffffffffffffff821115610ba357600080fd5b601f19601f8301169050602081019050919050565b600082825260208201905092915050565b6000610bd482610be5565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b610c2781610bc9565b8114610c3257600080fd5b50565b610c3e81610c05565b8114610c4957600080fd5b5056fea264697066735822122089b303eef9e421921de1e56c88aa13a1861ef5e8781dde98996301ae775b536f64736f6c63430006040033"

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

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerTransactor) Deposit(opts *bind.TransactOpts, destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.contract.Transact(opts, "deposit", destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerSession) Deposit(destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.Contract.Deposit(&_CentrifugeAssetHandler.TransactOpts, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerTransactorSession) Deposit(destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.Contract.Deposit(&_CentrifugeAssetHandler.TransactOpts, destinationChainID, depositNonce, depositer, data)
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
var RuntimeBytecode = "0x608060405234801561001057600080fd5b506004361061004c5760003560e01c8063318c136e14610051578063cb65d2211461006f578063db95e75c1461008b578063fc9539cd146100bb575b600080fd5b6100596100d7565b6040516100669190610ac9565b60405180910390f35b6100896004803603810190610084919061088b565b6100fc565b005b6100a560048036038101906100a09190610862565b610430565b6040516100b29190610b44565b60405180910390f35b6100d560048036038101906100d09190610821565b6105c9565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461018b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018290610ae4565b60405180910390fd5b60008060008060208501519350604085015192506060850151915060808501519050600060028111156101ba57fe5b6002600083815260200190815260200160002060009054906101000a900460ff1660028111156101e657fe5b14610226576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021d90610b24565b60405180910390fd5b60016002600083815260200190815260200160002060006101000a81548160ff0219169083600281111561025657fe5b02179055506040518060c001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff168152602001828152506001600089815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a082015181600501559050505050505050505050565b610438610712565b600160008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610658576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064f90610ae4565b60405180910390fd5b6000602082015190506001600281111561066e57fe5b6002600083815260200190815260200160002060009054906101000a900460ff16600281111561069a57fe5b146106da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d190610b04565b60405180910390fd5b600280600083815260200190815260200160002060006101000a81548160ff0219169083600281111561070957fe5b02179055505050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600080191681525090565b6000813590506107b281610c1e565b92915050565b600082601f8301126107c957600080fd5b81356107dc6107d782610b8c565b610b5f565b915080825260208301602083018583830111156107f857600080fd5b610803838284610c0f565b50505092915050565b60008135905061081b81610c35565b92915050565b60006020828403121561083357600080fd5b600082013567ffffffffffffffff81111561084d57600080fd5b610859848285016107b8565b91505092915050565b60006020828403121561087457600080fd5b60006108828482850161080c565b91505092915050565b600080600080608085870312156108a157600080fd5b60006108af8782880161080c565b94505060206108c08782880161080c565b93505060406108d1878288016107a3565b925050606085013567ffffffffffffffff8111156108ee57600080fd5b6108fa878288016107b8565b91505092959194509250565b61090f81610bc9565b82525050565b61091e81610bc9565b82525050565b61092d81610bdb565b82525050565b6000610940601e83610bb8565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000610980603983610bb8565b91507f6173736574206861736e2774206265656e206465706f7369746564206f72206860008301527f617320616c7265616479206265656e2066696e616c697a6564000000000000006020830152604082019050919050565b60006109e6603683610bb8565b91507f61737365742068617320616c7265616479206265656e206465706f736974656460008301527f20616e642063616e6e6f74206265206368616e676564000000000000000000006020830152604082019050919050565b60c082016000820151610a556000850182610906565b506020820151610a686020850182610aba565b506040820151610a7b6040850182610906565b506060820151610a8e6060850182610906565b506080820151610aa16080850182610906565b5060a0820151610ab460a0850182610924565b50505050565b610ac381610c05565b82525050565b6000602082019050610ade6000830184610915565b92915050565b60006020820190508181036000830152610afd81610933565b9050919050565b60006020820190508181036000830152610b1d81610973565b9050919050565b60006020820190508181036000830152610b3d816109d9565b9050919050565b600060c082019050610b596000830184610a3f565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610b8257600080fd5b8060405250919050565b600067ffffffffffffffff821115610ba357600080fd5b601f19601f8301169050602081019050919050565b600082825260208201905092915050565b6000610bd482610be5565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b610c2781610bc9565b8114610c3257600080fd5b50565b610c3e81610c05565b8114610c4957600080fd5b5056fea264697066735822122089b303eef9e421921de1e56c88aa13a1861ef5e8781dde98996301ae775b536f64736f6c63430006040033"
