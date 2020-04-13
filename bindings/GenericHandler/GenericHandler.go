// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GenericHandler

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

// GenericHandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type GenericHandlerDepositRecord struct {
	DestinationChainID          *big.Int
	ResourceID                  [32]byte
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaData                    []byte
}

// GenericHandlerABI is the input ABI used to generate the binding from.
const GenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structGenericHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GenericHandlerBin is the compiled bytecode used for deploying new contracts.
var GenericHandlerBin = "0x608060405234801561001057600080fd5b50604051610e35380380610e358339818101604052810190610032919061008e565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610100565b600081519050610088816100e9565b92915050565b6000602082840312156100a057600080fd5b60006100ae84828501610079565b91505092915050565b60006100c2826100c9565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6100f2816100b7565b81146100fd57600080fd5b50565b610d268061010f6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063318c136e146100675780636ebcf60714610085578063cb65d221146100b5578063db95e75c146100d1578063e245386f14610101578063fc9539cd14610135575b600080fd5b61006f610151565b60405161007c9190610ac0565b60405180910390f35b61009f600480360381019061009a9190610830565b610177565b6040516100ac9190610b1d565b60405180910390f35b6100cf60048036038101906100ca91906108c3565b61018f565b005b6100eb60048036038101906100e6919061089a565b61037f565b6040516100f89190610afb565b60405180910390f35b61011b6004803603810190610116919061089a565b61050e565b60405161012c959493929190610b38565b60405180910390f35b61014f600480360381019061014a9190610859565b61061c565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461021f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021690610adb565b60405180910390fd5b600080606060208401519250604084015191506040519050836060015180820160600160405260e4360360e48337506040518060a001604052808881526020018381526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018281525060026000888152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040190805190602001906103729291906106af565b5090505050505050505050565b61038761072f565b600260008381526020019081526020016000206040518060a001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104fe5780601f106104d3576101008083540402835291602001916104fe565b820191906000526020600020905b8154815290600101906020018083116104e157829003601f168201915b5050505050815250509050919050565b60026020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806004018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106125780601f106105e757610100808354040283529160200191610612565b820191906000526020600020905b8154815290600101906020018083116105f557829003601f168201915b5050505050905085565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a390610adb565b60405180910390fd5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106f057805160ff191683800117855561071e565b8280016001018555821561071e579182015b8281111561071d578251825591602001919060010190610702565b5b50905061072b919061078d565b5090565b6040518060a001604052806000815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b6107af91905b808211156107ab576000816000905550600101610793565b5090565b90565b6000813590506107c181610cc2565b92915050565b600082601f8301126107d857600080fd5b81356107eb6107e682610bbf565b610b92565b9150808252602083016020830185838301111561080757600080fd5b610812838284610c6f565b50505092915050565b60008135905061082a81610cd9565b92915050565b60006020828403121561084257600080fd5b6000610850848285016107b2565b91505092915050565b60006020828403121561086b57600080fd5b600082013567ffffffffffffffff81111561088557600080fd5b610891848285016107c7565b91505092915050565b6000602082840312156108ac57600080fd5b60006108ba8482850161081b565b91505092915050565b600080600080608085870312156108d957600080fd5b60006108e78782880161081b565b94505060206108f88782880161081b565b9350506040610909878288016107b2565b925050606085013567ffffffffffffffff81111561092657600080fd5b610932878288016107c7565b91505092959194509250565b61094781610c29565b82525050565b61095681610c29565b82525050565b61096581610c3b565b82525050565b61097481610c3b565b82525050565b600061098582610beb565b61098f8185610bf6565b935061099f818560208601610c7e565b6109a881610cb1565b840191505092915050565b60006109be82610beb565b6109c88185610c07565b93506109d8818560208601610c7e565b6109e181610cb1565b840191505092915050565b60006109f9601e83610c18565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b600060a083016000830151610a446000860182610aa2565b506020830151610a57602086018261095c565b506040830151610a6a604086018261093e565b506060830151610a7d606086018261093e565b5060808301518482036080860152610a95828261097a565b9150508091505092915050565b610aab81610c65565b82525050565b610aba81610c65565b82525050565b6000602082019050610ad5600083018461094d565b92915050565b60006020820190508181036000830152610af4816109ec565b9050919050565b60006020820190508181036000830152610b158184610a2c565b905092915050565b6000602082019050610b326000830184610ab1565b92915050565b600060a082019050610b4d6000830188610ab1565b610b5a602083018761096b565b610b67604083018661094d565b610b74606083018561094d565b8181036080830152610b8681846109b3565b90509695505050505050565b6000604051905081810181811067ffffffffffffffff82111715610bb557600080fd5b8060405250919050565b600067ffffffffffffffff821115610bd657600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610c3482610c45565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610c9c578082015181840152602081019050610c81565b83811115610cab576000848401525b50505050565b6000601f19601f8301169050919050565b610ccb81610c29565b8114610cd657600080fd5b50565b610ce281610c65565b8114610ced57600080fd5b5056fea264697066735822122002423c7cf5ddc77453344e2b3eff5d680bc25b30426fe657b23b4f5fb7c58cd764736f6c63430006040033"

// DeployGenericHandler deploys a new Ethereum contract, binding an instance of GenericHandler to it.
func DeployGenericHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *GenericHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GenericHandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GenericHandler{GenericHandlerCaller: GenericHandlerCaller{contract: contract}, GenericHandlerTransactor: GenericHandlerTransactor{contract: contract}, GenericHandlerFilterer: GenericHandlerFilterer{contract: contract}}, nil
}

// GenericHandler is an auto generated Go binding around an Ethereum contract.
type GenericHandler struct {
	GenericHandlerCaller     // Read-only binding to the contract
	GenericHandlerTransactor // Write-only binding to the contract
	GenericHandlerFilterer   // Log filterer for contract events
}

// GenericHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenericHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericHandlerSession struct {
	Contract     *GenericHandler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GenericHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericHandlerCallerSession struct {
	Contract *GenericHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GenericHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericHandlerTransactorSession struct {
	Contract     *GenericHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GenericHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenericHandlerRaw struct {
	Contract *GenericHandler // Generic contract binding to access the raw methods on
}

// GenericHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericHandlerCallerRaw struct {
	Contract *GenericHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// GenericHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericHandlerTransactorRaw struct {
	Contract *GenericHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericHandler creates a new instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandler(address common.Address, backend bind.ContractBackend) (*GenericHandler, error) {
	contract, err := bindGenericHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericHandler{GenericHandlerCaller: GenericHandlerCaller{contract: contract}, GenericHandlerTransactor: GenericHandlerTransactor{contract: contract}, GenericHandlerFilterer: GenericHandlerFilterer{contract: contract}}, nil
}

// NewGenericHandlerCaller creates a new read-only instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerCaller(address common.Address, caller bind.ContractCaller) (*GenericHandlerCaller, error) {
	contract, err := bindGenericHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerCaller{contract: contract}, nil
}

// NewGenericHandlerTransactor creates a new write-only instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*GenericHandlerTransactor, error) {
	contract, err := bindGenericHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerTransactor{contract: contract}, nil
}

// NewGenericHandlerFilterer creates a new log filterer instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*GenericHandlerFilterer, error) {
	contract, err := bindGenericHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerFilterer{contract: contract}, nil
}

// bindGenericHandler binds a generic wrapper to an already deployed contract.
func bindGenericHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericHandler *GenericHandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericHandler.Contract.GenericHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericHandler *GenericHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericHandler.Contract.GenericHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericHandler *GenericHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericHandler.Contract.GenericHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericHandler *GenericHandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenericHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericHandler *GenericHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericHandler *GenericHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericHandler.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_GenericHandler *GenericHandlerCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericHandler.contract.Call(opts, out, "_balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_GenericHandler *GenericHandlerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GenericHandler.Contract.Balances(&_GenericHandler.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_GenericHandler *GenericHandlerCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GenericHandler.Contract.Balances(&_GenericHandler.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() constant returns(address)
func (_GenericHandler *GenericHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GenericHandler.contract.Call(opts, out, "_bridgeAddress")
	return *ret0, err
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() constant returns(address)
func (_GenericHandler *GenericHandlerSession) BridgeAddress() (common.Address, error) {
	return _GenericHandler.Contract.BridgeAddress(&_GenericHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() constant returns(address)
func (_GenericHandler *GenericHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _GenericHandler.Contract.BridgeAddress(&_GenericHandler.CallOpts)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(uint256 _destinationChainID, bytes32 _resourceID, address _destinationRecipientAddress, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DestinationChainID          *big.Int
	ResourceID                  [32]byte
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaData                    []byte
}, error) {
	ret := new(struct {
		DestinationChainID          *big.Int
		ResourceID                  [32]byte
		DestinationRecipientAddress common.Address
		Depositer                   common.Address
		MetaData                    []byte
	})
	out := ret
	err := _GenericHandler.contract.Call(opts, out, "_depositRecords", arg0)
	return *ret, err
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(uint256 _destinationChainID, bytes32 _resourceID, address _destinationRecipientAddress, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerSession) DepositRecords(arg0 *big.Int) (struct {
	DestinationChainID          *big.Int
	ResourceID                  [32]byte
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaData                    []byte
}, error) {
	return _GenericHandler.Contract.DepositRecords(&_GenericHandler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(uint256 _destinationChainID, bytes32 _resourceID, address _destinationRecipientAddress, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerCallerSession) DepositRecords(arg0 *big.Int) (struct {
	DestinationChainID          *big.Int
	ResourceID                  [32]byte
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaData                    []byte
}, error) {
	return _GenericHandler.Contract.DepositRecords(&_GenericHandler.CallOpts, arg0)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositNonce) constant returns(GenericHandlerDepositRecord)
func (_GenericHandler *GenericHandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositNonce *big.Int) (GenericHandlerDepositRecord, error) {
	var (
		ret0 = new(GenericHandlerDepositRecord)
	)
	out := ret0
	err := _GenericHandler.contract.Call(opts, out, "getDepositRecord", depositNonce)
	return *ret0, err
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositNonce) constant returns(GenericHandlerDepositRecord)
func (_GenericHandler *GenericHandlerSession) GetDepositRecord(depositNonce *big.Int) (GenericHandlerDepositRecord, error) {
	return _GenericHandler.Contract.GetDepositRecord(&_GenericHandler.CallOpts, depositNonce)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositNonce) constant returns(GenericHandlerDepositRecord)
func (_GenericHandler *GenericHandlerCallerSession) GetDepositRecord(depositNonce *big.Int) (GenericHandlerDepositRecord, error) {
	return _GenericHandler.Contract.GetDepositRecord(&_GenericHandler.CallOpts, depositNonce)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_GenericHandler *GenericHandlerTransactor) Deposit(opts *bind.TransactOpts, destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "deposit", destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_GenericHandler *GenericHandlerSession) Deposit(destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Deposit(&_GenericHandler.TransactOpts, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_GenericHandler *GenericHandlerTransactorSession) Deposit(destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Deposit(&_GenericHandler.TransactOpts, destinationChainID, depositNonce, depositer, data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_GenericHandler *GenericHandlerTransactor) ExecuteDeposit(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "executeDeposit", data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_GenericHandler *GenericHandlerSession) ExecuteDeposit(data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.ExecuteDeposit(&_GenericHandler.TransactOpts, data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_GenericHandler *GenericHandlerTransactorSession) ExecuteDeposit(data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.ExecuteDeposit(&_GenericHandler.TransactOpts, data)
}

var RuntimeBytecode = "0x608060405234801561001057600080fd5b50600436106100625760003560e01c8063318c136e146100675780636ebcf60714610085578063cb65d221146100b5578063db95e75c146100d1578063e245386f14610101578063fc9539cd14610135575b600080fd5b61006f610151565b60405161007c9190610ac0565b60405180910390f35b61009f600480360381019061009a9190610830565b610177565b6040516100ac9190610b1d565b60405180910390f35b6100cf60048036038101906100ca91906108c3565b61018f565b005b6100eb60048036038101906100e6919061089a565b61037f565b6040516100f89190610afb565b60405180910390f35b61011b6004803603810190610116919061089a565b61050e565b60405161012c959493929190610b38565b60405180910390f35b61014f600480360381019061014a9190610859565b61061c565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461021f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021690610adb565b60405180910390fd5b600080606060208401519250604084015191506040519050836060015180820160600160405260e4360360e48337506040518060a001604052808881526020018381526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018281525060026000888152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040190805190602001906103729291906106af565b5090505050505050505050565b61038761072f565b600260008381526020019081526020016000206040518060a001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104fe5780601f106104d3576101008083540402835291602001916104fe565b820191906000526020600020905b8154815290600101906020018083116104e157829003601f168201915b5050505050815250509050919050565b60026020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806004018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106125780601f106105e757610100808354040283529160200191610612565b820191906000526020600020905b8154815290600101906020018083116105f557829003601f168201915b5050505050905085565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a390610adb565b60405180910390fd5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106f057805160ff191683800117855561071e565b8280016001018555821561071e579182015b8281111561071d578251825591602001919060010190610702565b5b50905061072b919061078d565b5090565b6040518060a001604052806000815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b6107af91905b808211156107ab576000816000905550600101610793565b5090565b90565b6000813590506107c181610cc2565b92915050565b600082601f8301126107d857600080fd5b81356107eb6107e682610bbf565b610b92565b9150808252602083016020830185838301111561080757600080fd5b610812838284610c6f565b50505092915050565b60008135905061082a81610cd9565b92915050565b60006020828403121561084257600080fd5b6000610850848285016107b2565b91505092915050565b60006020828403121561086b57600080fd5b600082013567ffffffffffffffff81111561088557600080fd5b610891848285016107c7565b91505092915050565b6000602082840312156108ac57600080fd5b60006108ba8482850161081b565b91505092915050565b600080600080608085870312156108d957600080fd5b60006108e78782880161081b565b94505060206108f88782880161081b565b9350506040610909878288016107b2565b925050606085013567ffffffffffffffff81111561092657600080fd5b610932878288016107c7565b91505092959194509250565b61094781610c29565b82525050565b61095681610c29565b82525050565b61096581610c3b565b82525050565b61097481610c3b565b82525050565b600061098582610beb565b61098f8185610bf6565b935061099f818560208601610c7e565b6109a881610cb1565b840191505092915050565b60006109be82610beb565b6109c88185610c07565b93506109d8818560208601610c7e565b6109e181610cb1565b840191505092915050565b60006109f9601e83610c18565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b600060a083016000830151610a446000860182610aa2565b506020830151610a57602086018261095c565b506040830151610a6a604086018261093e565b506060830151610a7d606086018261093e565b5060808301518482036080860152610a95828261097a565b9150508091505092915050565b610aab81610c65565b82525050565b610aba81610c65565b82525050565b6000602082019050610ad5600083018461094d565b92915050565b60006020820190508181036000830152610af4816109ec565b9050919050565b60006020820190508181036000830152610b158184610a2c565b905092915050565b6000602082019050610b326000830184610ab1565b92915050565b600060a082019050610b4d6000830188610ab1565b610b5a602083018761096b565b610b67604083018661094d565b610b74606083018561094d565b8181036080830152610b8681846109b3565b90509695505050505050565b6000604051905081810181811067ffffffffffffffff82111715610bb557600080fd5b8060405250919050565b600067ffffffffffffffff821115610bd657600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610c3482610c45565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610c9c578082015181840152602081019050610c81565b83811115610cab576000848401525b50505050565b6000601f19601f8301169050919050565b610ccb81610c29565b8114610cd657600080fd5b50565b610ce281610c65565b8114610ced57600080fd5b5056fea264697066735822122002423c7cf5ddc77453344e2b3eff5d680bc25b30426fe657b23b4f5fb7c58cd764736f6c63430006040033"
