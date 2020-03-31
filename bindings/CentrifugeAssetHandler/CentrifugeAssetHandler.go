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
var CentrifugeAssetHandlerBin = "0x608060405234801561001057600080fd5b50604051610d01380380610d018339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b600081519050610087816100e8565b92915050565b60006020828403121561009f57600080fd5b60006100ad84828501610078565b91505092915050565b60006100c1826100c8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6100f1816100b6565b81146100fc57600080fd5b50565b610bf38061010e6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063318c136e14610051578063cb65d2211461006f578063db95e75c1461008b578063fc9539cd146100bb575b600080fd5b6100596100d7565b6040516100669190610a3a565b60405180910390f35b610089600480360381019061008491906107fc565b6100fc565b005b6100a560048036038101906100a091906107d3565b610430565b6040516100b29190610ab5565b60405180910390f35b6100d560048036038101906100d09190610792565b6105c9565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461018b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018290610a55565b60405180910390fd5b60008060008060208501519350604085015192506060850151915060808501519050600060028111156101ba57fe5b6002600083815260200190815260200160002060009054906101000a900460ff1660028111156101e657fe5b14610226576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021d90610a95565b60405180910390fd5b60016002600083815260200190815260200160002060006101000a81548160ff0219169083600281111561025657fe5b02179055506040518060c001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff168152602001828152506001600089815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a082015181600501559050505050505050505050565b610438610683565b600160008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b600060208201519050600160028111156105df57fe5b6002600083815260200190815260200160002060009054906101000a900460ff16600281111561060b57fe5b1461064b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064290610a75565b60405180910390fd5b600280600083815260200190815260200160002060006101000a81548160ff0219169083600281111561067a57fe5b02179055505050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600080191681525090565b60008135905061072381610b8f565b92915050565b600082601f83011261073a57600080fd5b813561074d61074882610afd565b610ad0565b9150808252602083016020830185838301111561076957600080fd5b610774838284610b80565b50505092915050565b60008135905061078c81610ba6565b92915050565b6000602082840312156107a457600080fd5b600082013567ffffffffffffffff8111156107be57600080fd5b6107ca84828501610729565b91505092915050565b6000602082840312156107e557600080fd5b60006107f38482850161077d565b91505092915050565b6000806000806080858703121561081257600080fd5b60006108208782880161077d565b94505060206108318782880161077d565b935050604061084287828801610714565b925050606085013567ffffffffffffffff81111561085f57600080fd5b61086b87828801610729565b91505092959194509250565b61088081610b3a565b82525050565b61088f81610b3a565b82525050565b61089e81610b4c565b82525050565b60006108b1601e83610b29565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b60006108f1603983610b29565b91507f6173736574206861736e2774206265656e206465706f7369746564206f72206860008301527f617320616c7265616479206265656e2066696e616c697a6564000000000000006020830152604082019050919050565b6000610957603683610b29565b91507f61737365742068617320616c7265616479206265656e206465706f736974656460008301527f20616e642063616e6e6f74206265206368616e676564000000000000000000006020830152604082019050919050565b60c0820160008201516109c66000850182610877565b5060208201516109d96020850182610a2b565b5060408201516109ec6040850182610877565b5060608201516109ff6060850182610877565b506080820151610a126080850182610877565b5060a0820151610a2560a0850182610895565b50505050565b610a3481610b76565b82525050565b6000602082019050610a4f6000830184610886565b92915050565b60006020820190508181036000830152610a6e816108a4565b9050919050565b60006020820190508181036000830152610a8e816108e4565b9050919050565b60006020820190508181036000830152610aae8161094a565b9050919050565b600060c082019050610aca60008301846109b0565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610af357600080fd5b8060405250919050565b600067ffffffffffffffff821115610b1457600080fd5b601f19601f8301169050602081019050919050565b600082825260208201905092915050565b6000610b4582610b56565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b610b9881610b3a565b8114610ba357600080fd5b50565b610baf81610b76565b8114610bba57600080fd5b5056fea264697066735822122088302170e1e327773100bbf6d4ad35daa18b0a69941f25957ceda9058cbeb2e464736f6c63430006040033"

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
var RuntimeBytecode = "0x608060405234801561001057600080fd5b506004361061004c5760003560e01c8063318c136e14610051578063cb65d2211461006f578063db95e75c1461008b578063fc9539cd146100bb575b600080fd5b6100596100d7565b6040516100669190610a3a565b60405180910390f35b610089600480360381019061008491906107fc565b6100fc565b005b6100a560048036038101906100a091906107d3565b610430565b6040516100b29190610ab5565b60405180910390f35b6100d560048036038101906100d09190610792565b6105c9565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461018b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018290610a55565b60405180910390fd5b60008060008060208501519350604085015192506060850151915060808501519050600060028111156101ba57fe5b6002600083815260200190815260200160002060009054906101000a900460ff1660028111156101e657fe5b14610226576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021d90610a95565b60405180910390fd5b60016002600083815260200190815260200160002060006101000a81548160ff0219169083600281111561025657fe5b02179055506040518060c001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff168152602001828152506001600089815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a082015181600501559050505050505050505050565b610438610683565b600160008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b600060208201519050600160028111156105df57fe5b6002600083815260200190815260200160002060009054906101000a900460ff16600281111561060b57fe5b1461064b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064290610a75565b60405180910390fd5b600280600083815260200190815260200160002060006101000a81548160ff0219169083600281111561067a57fe5b02179055505050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600080191681525090565b60008135905061072381610b8f565b92915050565b600082601f83011261073a57600080fd5b813561074d61074882610afd565b610ad0565b9150808252602083016020830185838301111561076957600080fd5b610774838284610b80565b50505092915050565b60008135905061078c81610ba6565b92915050565b6000602082840312156107a457600080fd5b600082013567ffffffffffffffff8111156107be57600080fd5b6107ca84828501610729565b91505092915050565b6000602082840312156107e557600080fd5b60006107f38482850161077d565b91505092915050565b6000806000806080858703121561081257600080fd5b60006108208782880161077d565b94505060206108318782880161077d565b935050604061084287828801610714565b925050606085013567ffffffffffffffff81111561085f57600080fd5b61086b87828801610729565b91505092959194509250565b61088081610b3a565b82525050565b61088f81610b3a565b82525050565b61089e81610b4c565b82525050565b60006108b1601e83610b29565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b60006108f1603983610b29565b91507f6173736574206861736e2774206265656e206465706f7369746564206f72206860008301527f617320616c7265616479206265656e2066696e616c697a6564000000000000006020830152604082019050919050565b6000610957603683610b29565b91507f61737365742068617320616c7265616479206265656e206465706f736974656460008301527f20616e642063616e6e6f74206265206368616e676564000000000000000000006020830152604082019050919050565b60c0820160008201516109c66000850182610877565b5060208201516109d96020850182610a2b565b5060408201516109ec6040850182610877565b5060608201516109ff6060850182610877565b506080820151610a126080850182610877565b5060a0820151610a2560a0850182610895565b50505050565b610a3481610b76565b82525050565b6000602082019050610a4f6000830184610886565b92915050565b60006020820190508181036000830152610a6e816108a4565b9050919050565b60006020820190508181036000830152610a8e816108e4565b9050919050565b60006020820190508181036000830152610aae8161094a565b9050919050565b600060c082019050610aca60008301846109b0565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610af357600080fd5b8060405250919050565b600067ffffffffffffffff821115610b1457600080fd5b601f19601f8301169050602081019050919050565b600082825260208201905092915050565b6000610b4582610b56565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b610b9881610b3a565b8114610ba357600080fd5b50565b610baf81610b76565b8114610bba57600080fd5b5056fea264697066735822122088302170e1e327773100bbf6d4ad35daa18b0a69941f25957ceda9058cbeb2e464736f6c63430006040033"
