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
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaData                    []byte
}

// GenericHandlerABI is the input ABI used to generate the binding from.
const GenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structGenericHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GenericHandlerBin is the compiled bytecode used for deploying new contracts.
var GenericHandlerBin = "0x608060405234801561001057600080fd5b50604051610db8380380610db88339818101604052810190610032919061008e565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610100565b600081519050610088816100e9565b92915050565b6000602082840312156100a057600080fd5b60006100ae84828501610079565b91505092915050565b60006100c2826100c9565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6100f2816100b7565b81146100fd57600080fd5b50565b610ca98061010f6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063318c136e146100675780636ebcf60714610085578063cb65d221146100b5578063db95e75c146100d1578063e245386f14610101578063fc9539cd14610134575b600080fd5b61006f610150565b60405161007c9190610a5b565b60405180910390f35b61009f600480360381019061009a91906107fc565b610176565b6040516100ac9190610ab8565b60405180910390f35b6100cf60048036038101906100ca919061088f565b61018e565b005b6100eb60048036038101906100e69190610866565b610365565b6040516100f89190610a96565b60405180910390f35b61011b60048036038101906101169190610866565b6104ea565b60405161012b9493929190610ad3565b60405180910390f35b61014e60048036038101906101499190610825565b6105f2565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461021e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021590610a76565b60405180910390fd5b60006060602083015191506040519050826040015180820160400160405260c4360360c483375060405180608001604052808781526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff16815260200182815250600260008781526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003019080519060200190610359929190610685565b50905050505050505050565b61036d610705565b60026000838152602001908152602001600020604051806080016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104da5780601f106104af576101008083540402835291602001916104da565b820191906000526020600020905b8154815290600101906020018083116104bd57829003601f168201915b5050505050815250509050919050565b60026020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105e85780601f106105bd576101008083540402835291602001916105e8565b820191906000526020600020905b8154815290600101906020018083116105cb57829003601f168201915b5050505050905084565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610682576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067990610a76565b60405180910390fd5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106c657805160ff19168380011785556106f4565b828001600101855582156106f4579182015b828111156106f35782518255916020019190600101906106d8565b5b5090506107019190610759565b5090565b604051806080016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b61077b91905b8082111561077757600081600090555060010161075f565b5090565b90565b60008135905061078d81610c45565b92915050565b600082601f8301126107a457600080fd5b81356107b76107b282610b4c565b610b1f565b915080825260208301602083018583830111156107d357600080fd5b6107de838284610bf2565b50505092915050565b6000813590506107f681610c5c565b92915050565b60006020828403121561080e57600080fd5b600061081c8482850161077e565b91505092915050565b60006020828403121561083757600080fd5b600082013567ffffffffffffffff81111561085157600080fd5b61085d84828501610793565b91505092915050565b60006020828403121561087857600080fd5b6000610886848285016107e7565b91505092915050565b600080600080608085870312156108a557600080fd5b60006108b3878288016107e7565b94505060206108c4878288016107e7565b93505060406108d58782880161077e565b925050606085013567ffffffffffffffff8111156108f257600080fd5b6108fe87828801610793565b91505092959194509250565b61091381610bb6565b82525050565b61092281610bb6565b82525050565b600061093382610b78565b61093d8185610b83565b935061094d818560208601610c01565b61095681610c34565b840191505092915050565b600061096c82610b78565b6109768185610b94565b9350610986818560208601610c01565b61098f81610c34565b840191505092915050565b60006109a7601e83610ba5565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b60006080830160008301516109f26000860182610a3d565b506020830151610a05602086018261090a565b506040830151610a18604086018261090a565b5060608301518482036060860152610a308282610928565b9150508091505092915050565b610a4681610be8565b82525050565b610a5581610be8565b82525050565b6000602082019050610a706000830184610919565b92915050565b60006020820190508181036000830152610a8f8161099a565b9050919050565b60006020820190508181036000830152610ab081846109da565b905092915050565b6000602082019050610acd6000830184610a4c565b92915050565b6000608082019050610ae86000830187610a4c565b610af56020830186610919565b610b026040830185610919565b8181036060830152610b148184610961565b905095945050505050565b6000604051905081810181811067ffffffffffffffff82111715610b4257600080fd5b8060405250919050565b600067ffffffffffffffff821115610b6357600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610bc182610bc8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610c1f578082015181840152602081019050610c04565b83811115610c2e576000848401525b50505050565b6000601f19601f8301169050919050565b610c4e81610bb6565b8114610c5957600080fd5b50565b610c6581610be8565b8114610c7057600080fd5b5056fea26469706673582212209f90fc2aef59ce8c28fe2b0b71e077860860a81a4b1e1c19a2f437977818ed7564736f6c63430006040033"

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
// Solidity: function _depositRecords(uint256 ) constant returns(uint256 _destinationChainID, address _destinationRecipientAddress, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DestinationChainID          *big.Int
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaData                    []byte
}, error) {
	ret := new(struct {
		DestinationChainID          *big.Int
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
// Solidity: function _depositRecords(uint256 ) constant returns(uint256 _destinationChainID, address _destinationRecipientAddress, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerSession) DepositRecords(arg0 *big.Int) (struct {
	DestinationChainID          *big.Int
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaData                    []byte
}, error) {
	return _GenericHandler.Contract.DepositRecords(&_GenericHandler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(uint256 _destinationChainID, address _destinationRecipientAddress, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerCallerSession) DepositRecords(arg0 *big.Int) (struct {
	DestinationChainID          *big.Int
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaData                    []byte
}, error) {
	return _GenericHandler.Contract.DepositRecords(&_GenericHandler.CallOpts, arg0)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) constant returns(GenericHandlerDepositRecord)
func (_GenericHandler *GenericHandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositID *big.Int) (GenericHandlerDepositRecord, error) {
	var (
		ret0 = new(GenericHandlerDepositRecord)
	)
	out := ret0
	err := _GenericHandler.contract.Call(opts, out, "getDepositRecord", depositID)
	return *ret0, err
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) constant returns(GenericHandlerDepositRecord)
func (_GenericHandler *GenericHandlerSession) GetDepositRecord(depositID *big.Int) (GenericHandlerDepositRecord, error) {
	return _GenericHandler.Contract.GetDepositRecord(&_GenericHandler.CallOpts, depositID)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) constant returns(GenericHandlerDepositRecord)
func (_GenericHandler *GenericHandlerCallerSession) GetDepositRecord(depositID *big.Int) (GenericHandlerDepositRecord, error) {
	return _GenericHandler.Contract.GetDepositRecord(&_GenericHandler.CallOpts, depositID)
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
var RuntimeBytecode = "0x608060405234801561001057600080fd5b50600436106100625760003560e01c8063318c136e146100675780636ebcf60714610085578063cb65d221146100b5578063db95e75c146100d1578063e245386f14610101578063fc9539cd14610134575b600080fd5b61006f610150565b60405161007c9190610a5b565b60405180910390f35b61009f600480360381019061009a91906107fc565b610176565b6040516100ac9190610ab8565b60405180910390f35b6100cf60048036038101906100ca919061088f565b61018e565b005b6100eb60048036038101906100e69190610866565b610365565b6040516100f89190610a96565b60405180910390f35b61011b60048036038101906101169190610866565b6104ea565b60405161012b9493929190610ad3565b60405180910390f35b61014e60048036038101906101499190610825565b6105f2565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461021e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021590610a76565b60405180910390fd5b60006060602083015191506040519050826040015180820160400160405260c4360360c483375060405180608001604052808781526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff16815260200182815250600260008781526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003019080519060200190610359929190610685565b50905050505050505050565b61036d610705565b60026000838152602001908152602001600020604051806080016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104da5780601f106104af576101008083540402835291602001916104da565b820191906000526020600020905b8154815290600101906020018083116104bd57829003601f168201915b5050505050815250509050919050565b60026020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105e85780601f106105bd576101008083540402835291602001916105e8565b820191906000526020600020905b8154815290600101906020018083116105cb57829003601f168201915b5050505050905084565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610682576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067990610a76565b60405180910390fd5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106c657805160ff19168380011785556106f4565b828001600101855582156106f4579182015b828111156106f35782518255916020019190600101906106d8565b5b5090506107019190610759565b5090565b604051806080016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b61077b91905b8082111561077757600081600090555060010161075f565b5090565b90565b60008135905061078d81610c45565b92915050565b600082601f8301126107a457600080fd5b81356107b76107b282610b4c565b610b1f565b915080825260208301602083018583830111156107d357600080fd5b6107de838284610bf2565b50505092915050565b6000813590506107f681610c5c565b92915050565b60006020828403121561080e57600080fd5b600061081c8482850161077e565b91505092915050565b60006020828403121561083757600080fd5b600082013567ffffffffffffffff81111561085157600080fd5b61085d84828501610793565b91505092915050565b60006020828403121561087857600080fd5b6000610886848285016107e7565b91505092915050565b600080600080608085870312156108a557600080fd5b60006108b3878288016107e7565b94505060206108c4878288016107e7565b93505060406108d58782880161077e565b925050606085013567ffffffffffffffff8111156108f257600080fd5b6108fe87828801610793565b91505092959194509250565b61091381610bb6565b82525050565b61092281610bb6565b82525050565b600061093382610b78565b61093d8185610b83565b935061094d818560208601610c01565b61095681610c34565b840191505092915050565b600061096c82610b78565b6109768185610b94565b9350610986818560208601610c01565b61098f81610c34565b840191505092915050565b60006109a7601e83610ba5565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b60006080830160008301516109f26000860182610a3d565b506020830151610a05602086018261090a565b506040830151610a18604086018261090a565b5060608301518482036060860152610a308282610928565b9150508091505092915050565b610a4681610be8565b82525050565b610a5581610be8565b82525050565b6000602082019050610a706000830184610919565b92915050565b60006020820190508181036000830152610a8f8161099a565b9050919050565b60006020820190508181036000830152610ab081846109da565b905092915050565b6000602082019050610acd6000830184610a4c565b92915050565b6000608082019050610ae86000830187610a4c565b610af56020830186610919565b610b026040830185610919565b8181036060830152610b148184610961565b905095945050505050565b6000604051905081810181811067ffffffffffffffff82111715610b4257600080fd5b8060405250919050565b600067ffffffffffffffff821115610b6357600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610bc182610bc8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610c1f578082015181840152602081019050610c04565b83811115610c2e576000848401525b50505050565b6000601f19601f8301169050919050565b610c4e81610bb6565b8114610c5957600080fd5b50565b610c6581610be8565b8114610c7057600080fd5b5056fea26469706673582212209f90fc2aef59ce8c28fe2b0b71e077860860a81a4b1e1c19a2f437977818ed7564736f6c63430006040033"
