// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package HandlerHelpers

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

// HandlerHelpersABI is the input ABI used to generate the binding from.
const HandlerHelpersABI = "[{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HandlerHelpersBin is the compiled bytecode used for deploying new contracts.
var HandlerHelpersBin = "0x608060405234801561001057600080fd5b50610983806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80636a70d0811161005b5780636a70d0811461017e5780637f79bea8146101da578063b8fa373614610236578063c8ba6c87146102845761007d565b806307b7ed99146100825780630a6d55d8146100c6578063318c136e14610134575b600080fd5b6100c46004803603602081101561009857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102dc565b005b6100f2600480360360208110156100dc57600080fd5b81019080803590602001909291905050506103aa565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61013c6103dd565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101c06004803603602081101561019457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610402565b604051808215151515815260200191505060405180910390f35b61021c600480360360208110156101f057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610422565b604051808215151515815260200191505060405180910390f35b6102826004803603604081101561024c57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610442565b005b6102c66004803603602081101561029a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106b6565b6040518082815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461039e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f73656e646572206d7573742062652062726964676520636f6e7472616374000081525060200191505060405180910390fd5b6103a7816106ce565b50565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915054906101000a900460ff1681565b60036020528060005260406000206000915054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610504576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f73656e646572206d7573742062652062726964676520636f6e7472616374000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166001600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146105bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260378152602001806109176037913960400191505060405180910390fd5b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020018082815260200191505060405160208183030381529060405280519060200120826040516020018082815260200191505060405160208183030381529060405280519060200120146106a6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260358152602001806108be6035913960400191505060405180910390fd5b6106b084846107cb565b50505050565b60026020528060005260406000206000915090505481565b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610770576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806108f36024913960400191505060405180910390fd5b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550505056fe636f6e7472616374206164647265737320616c72656164792068617320636f72726573706f6e64696e67207265736f75726365494470726f766964656420636f6e7472616374206973206e6f742077686974656c69737465647265736f75726365494420616c726561647920686173206120636f72726573706f6e64696e6720636f6e74726163742061646472657373a26469706673582212204d8a08ad921ab5015de5cc7daa9f0ecbc00e27031f714ea4e220c498160d4bd264736f6c63430006040033"

// DeployHandlerHelpers deploys a new Ethereum contract, binding an instance of HandlerHelpers to it.
func DeployHandlerHelpers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HandlerHelpers, error) {
	parsed, err := abi.JSON(strings.NewReader(HandlerHelpersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HandlerHelpersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HandlerHelpers{HandlerHelpersCaller: HandlerHelpersCaller{contract: contract}, HandlerHelpersTransactor: HandlerHelpersTransactor{contract: contract}, HandlerHelpersFilterer: HandlerHelpersFilterer{contract: contract}}, nil
}

// HandlerHelpers is an auto generated Go binding around an Ethereum contract.
type HandlerHelpers struct {
	HandlerHelpersCaller     // Read-only binding to the contract
	HandlerHelpersTransactor // Write-only binding to the contract
	HandlerHelpersFilterer   // Log filterer for contract events
}

// HandlerHelpersCaller is an auto generated read-only Go binding around an Ethereum contract.
type HandlerHelpersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerHelpersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HandlerHelpersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerHelpersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HandlerHelpersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerHelpersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HandlerHelpersSession struct {
	Contract     *HandlerHelpers   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HandlerHelpersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HandlerHelpersCallerSession struct {
	Contract *HandlerHelpersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HandlerHelpersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HandlerHelpersTransactorSession struct {
	Contract     *HandlerHelpersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HandlerHelpersRaw is an auto generated low-level Go binding around an Ethereum contract.
type HandlerHelpersRaw struct {
	Contract *HandlerHelpers // Generic contract binding to access the raw methods on
}

// HandlerHelpersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HandlerHelpersCallerRaw struct {
	Contract *HandlerHelpersCaller // Generic read-only contract binding to access the raw methods on
}

// HandlerHelpersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HandlerHelpersTransactorRaw struct {
	Contract *HandlerHelpersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHandlerHelpers creates a new instance of HandlerHelpers, bound to a specific deployed contract.
func NewHandlerHelpers(address common.Address, backend bind.ContractBackend) (*HandlerHelpers, error) {
	contract, err := bindHandlerHelpers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HandlerHelpers{HandlerHelpersCaller: HandlerHelpersCaller{contract: contract}, HandlerHelpersTransactor: HandlerHelpersTransactor{contract: contract}, HandlerHelpersFilterer: HandlerHelpersFilterer{contract: contract}}, nil
}

// NewHandlerHelpersCaller creates a new read-only instance of HandlerHelpers, bound to a specific deployed contract.
func NewHandlerHelpersCaller(address common.Address, caller bind.ContractCaller) (*HandlerHelpersCaller, error) {
	contract, err := bindHandlerHelpers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HandlerHelpersCaller{contract: contract}, nil
}

// NewHandlerHelpersTransactor creates a new write-only instance of HandlerHelpers, bound to a specific deployed contract.
func NewHandlerHelpersTransactor(address common.Address, transactor bind.ContractTransactor) (*HandlerHelpersTransactor, error) {
	contract, err := bindHandlerHelpers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HandlerHelpersTransactor{contract: contract}, nil
}

// NewHandlerHelpersFilterer creates a new log filterer instance of HandlerHelpers, bound to a specific deployed contract.
func NewHandlerHelpersFilterer(address common.Address, filterer bind.ContractFilterer) (*HandlerHelpersFilterer, error) {
	contract, err := bindHandlerHelpers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HandlerHelpersFilterer{contract: contract}, nil
}

// bindHandlerHelpers binds a generic wrapper to an already deployed contract.
func bindHandlerHelpers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HandlerHelpersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HandlerHelpers *HandlerHelpersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HandlerHelpers.Contract.HandlerHelpersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HandlerHelpers *HandlerHelpersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.HandlerHelpersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HandlerHelpers *HandlerHelpersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.HandlerHelpersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HandlerHelpers *HandlerHelpersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HandlerHelpers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HandlerHelpers *HandlerHelpersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HandlerHelpers *HandlerHelpersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HandlerHelpers *HandlerHelpersCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HandlerHelpers.contract.Call(opts, out, "_bridgeAddress")
	return *ret0, err
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HandlerHelpers *HandlerHelpersSession) BridgeAddress() (common.Address, error) {
	return _HandlerHelpers.Contract.BridgeAddress(&_HandlerHelpers.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HandlerHelpers *HandlerHelpersCallerSession) BridgeAddress() (common.Address, error) {
	return _HandlerHelpers.Contract.BridgeAddress(&_HandlerHelpers.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HandlerHelpers.contract.Call(opts, out, "_burnList", arg0)
	return *ret0, err
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersSession) BurnList(arg0 common.Address) (bool, error) {
	return _HandlerHelpers.Contract.BurnList(&_HandlerHelpers.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _HandlerHelpers.Contract.BurnList(&_HandlerHelpers.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HandlerHelpers.contract.Call(opts, out, "_contractWhitelist", arg0)
	return *ret0, err
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _HandlerHelpers.Contract.ContractWhitelist(&_HandlerHelpers.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _HandlerHelpers.Contract.ContractWhitelist(&_HandlerHelpers.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_HandlerHelpers *HandlerHelpersCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HandlerHelpers.contract.Call(opts, out, "_resourceIDToTokenContractAddress", arg0)
	return *ret0, err
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_HandlerHelpers *HandlerHelpersSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _HandlerHelpers.Contract.ResourceIDToTokenContractAddress(&_HandlerHelpers.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_HandlerHelpers *HandlerHelpersCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _HandlerHelpers.Contract.ResourceIDToTokenContractAddress(&_HandlerHelpers.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_HandlerHelpers *HandlerHelpersCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _HandlerHelpers.contract.Call(opts, out, "_tokenContractAddressToResourceID", arg0)
	return *ret0, err
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_HandlerHelpers *HandlerHelpersSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _HandlerHelpers.Contract.TokenContractAddressToResourceID(&_HandlerHelpers.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_HandlerHelpers *HandlerHelpersCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _HandlerHelpers.Contract.TokenContractAddressToResourceID(&_HandlerHelpers.CallOpts, arg0)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetBurnable(&_HandlerHelpers.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetBurnable(&_HandlerHelpers.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetResource(&_HandlerHelpers.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetResource(&_HandlerHelpers.TransactOpts, resourceID, contractAddress)
}
