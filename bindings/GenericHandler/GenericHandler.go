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
	DestinationChainID          uint8
	ResourceID                  [32]byte
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaData                    []byte
}

// GenericHandlerABI is the input ABI used to generate the binding from.
const GenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structGenericHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GenericHandlerBin is the compiled bytecode used for deploying new contracts.
var GenericHandlerBin = "0x608060405234801561001057600080fd5b50604051620011be380380620011be83398181016040528101906100349190610090565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610102565b60008151905061008a816100eb565b92915050565b6000602082840312156100a257600080fd5b60006100b08482850161007b565b91505092915050565b60006100c4826100cb565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6100f4816100b9565b81146100ff57600080fd5b50565b6110ac80620001126000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806395601f091161005b57806395601f09146100ec578063db95e75c14610108578063e245386f14610138578063fc9539cd1461016c5761007d565b8063318c136e1461008257806345a104db146100a05780636ebcf607146100bc575b600080fd5b61008a610188565b6040516100979190610da8565b60405180910390f35b6100ba60048036038101906100b59190610b5c565b6101ae565b005b6100d660048036038101906100d19190610a51565b6103b8565b6040516100e39190610e5c565b60405180910390f35b61010660048036038101906101019190610a7a565b6103d0565b005b610122600480360381019061011d9190610b33565b6104fe565b60405161012f9190610e3a565b60405180910390f35b610152600480360381019061014d9190610b33565b6106a0565b604051610163959493929190610e77565b60405180910390f35b61018660048036038101906101819190610af2565b6107bb565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461023e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023590610dfa565b60405180910390fd5b600080606060208401519250604084015191506040519050836060015180820160600160405260e4360360e48337506040518060a001604052808860ff1681526020018381526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001828152506002600088815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040190805190602001906103ab9291906108a3565b5090505050505050505050565b60006020528060005260406000206000915090505481565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b815260040161041293929190610dc3565b602060405180830381600087803b15801561042c57600080fd5b505af1158015610440573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104649190610ac9565b506104b6826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461084e90919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b610506610923565b600260008381526020019081526020016000206040518060a00160405290816000820160009054906101000a900460ff1660ff1660ff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106905780601f1061066557610100808354040283529160200191610690565b820191906000526020600020905b81548152906001019060200180831161067357829003601f168201915b5050505050815250509050919050565b60026020528060005260406000206000915090508060000160009054906101000a900460ff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806004018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107b15780601f10610786576101008083540402835291602001916107b1565b820191906000526020600020905b81548152906001019060200180831161079457829003601f168201915b5050505050905085565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461084b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084290610dfa565b60405180910390fd5b50565b600080828401905083811015610899576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089090610e1a565b60405180910390fd5b8091505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106108e457805160ff1916838001178555610912565b82800160010185558215610912579182015b828111156109115782518255916020019190600101906108f6565b5b50905061091f9190610984565b5090565b6040518060a00160405280600060ff16815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b6109a691905b808211156109a257600081600090555060010161098a565b5090565b90565b6000813590506109b88161101a565b92915050565b6000815190506109cd81611031565b92915050565b600082601f8301126109e457600080fd5b81356109f76109f282610efe565b610ed1565b91508082526020830160208301858383011115610a1357600080fd5b610a1e838284610fc7565b50505092915050565b600081359050610a3681611048565b92915050565b600081359050610a4b8161105f565b92915050565b600060208284031215610a6357600080fd5b6000610a71848285016109a9565b91505092915050565b600080600060608486031215610a8f57600080fd5b6000610a9d868287016109a9565b9350506020610aae868287016109a9565b9250506040610abf86828701610a27565b9150509250925092565b600060208284031215610adb57600080fd5b6000610ae9848285016109be565b91505092915050565b600060208284031215610b0457600080fd5b600082013567ffffffffffffffff811115610b1e57600080fd5b610b2a848285016109d3565b91505092915050565b600060208284031215610b4557600080fd5b6000610b5384828501610a27565b91505092915050565b60008060008060808587031215610b7257600080fd5b6000610b8087828801610a3c565b9450506020610b9187828801610a27565b9350506040610ba2878288016109a9565b925050606085013567ffffffffffffffff811115610bbf57600080fd5b610bcb878288016109d3565b91505092959194509250565b610be081610f68565b82525050565b610bef81610f68565b82525050565b610bfe81610f86565b82525050565b610c0d81610f86565b82525050565b6000610c1e82610f2a565b610c288185610f35565b9350610c38818560208601610fd6565b610c4181611009565b840191505092915050565b6000610c5782610f2a565b610c618185610f46565b9350610c71818560208601610fd6565b610c7a81611009565b840191505092915050565b6000610c92601e83610f57565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000610cd2601b83610f57565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b600060a083016000830151610d1d6000860182610d8a565b506020830151610d306020860182610bf5565b506040830151610d436040860182610bd7565b506060830151610d566060860182610bd7565b5060808301518482036080860152610d6e8282610c13565b9150508091505092915050565b610d8481610fb0565b82525050565b610d9381610fba565b82525050565b610da281610fba565b82525050565b6000602082019050610dbd6000830184610be6565b92915050565b6000606082019050610dd86000830186610be6565b610de56020830185610be6565b610df26040830184610d7b565b949350505050565b60006020820190508181036000830152610e1381610c85565b9050919050565b60006020820190508181036000830152610e3381610cc5565b9050919050565b60006020820190508181036000830152610e548184610d05565b905092915050565b6000602082019050610e716000830184610d7b565b92915050565b600060a082019050610e8c6000830188610d99565b610e996020830187610c04565b610ea66040830186610be6565b610eb36060830185610be6565b8181036080830152610ec58184610c4c565b90509695505050505050565b6000604051905081810181811067ffffffffffffffff82111715610ef457600080fd5b8060405250919050565b600067ffffffffffffffff821115610f1557600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610f7382610f90565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015610ff4578082015181840152602081019050610fd9565b83811115611003576000848401525b50505050565b6000601f19601f8301169050919050565b61102381610f68565b811461102e57600080fd5b50565b61103a81610f7a565b811461104557600080fd5b50565b61105181610fb0565b811461105c57600080fd5b50565b61106881610fba565b811461107357600080fd5b5056fea2646970667358221220337ec2ffd59f4d16694bfdcffc2c8795f772f53f416153ab35da08b303bacab564736f6c63430006040033"

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
// Solidity: function _depositRecords(uint256 ) constant returns(uint8 _destinationChainID, bytes32 _resourceID, address _destinationRecipientAddress, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DestinationChainID          uint8
	ResourceID                  [32]byte
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaData                    []byte
}, error) {
	ret := new(struct {
		DestinationChainID          uint8
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
// Solidity: function _depositRecords(uint256 ) constant returns(uint8 _destinationChainID, bytes32 _resourceID, address _destinationRecipientAddress, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerSession) DepositRecords(arg0 *big.Int) (struct {
	DestinationChainID          uint8
	ResourceID                  [32]byte
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaData                    []byte
}, error) {
	return _GenericHandler.Contract.DepositRecords(&_GenericHandler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(uint8 _destinationChainID, bytes32 _resourceID, address _destinationRecipientAddress, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerCallerSession) DepositRecords(arg0 *big.Int) (struct {
	DestinationChainID          uint8
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

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_GenericHandler *GenericHandlerTransactor) Deposit(opts *bind.TransactOpts, destinationChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "deposit", destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_GenericHandler *GenericHandlerSession) Deposit(destinationChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Deposit(&_GenericHandler.TransactOpts, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_GenericHandler *GenericHandlerTransactorSession) Deposit(destinationChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
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

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_GenericHandler *GenericHandlerTransactor) FundERC20(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "fundERC20", tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_GenericHandler *GenericHandlerSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GenericHandler.Contract.FundERC20(&_GenericHandler.TransactOpts, tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_GenericHandler *GenericHandlerTransactorSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GenericHandler.Contract.FundERC20(&_GenericHandler.TransactOpts, tokenAddress, owner, amount)
}

var RuntimeBytecode = "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c806395601f091161005b57806395601f09146100ec578063db95e75c14610108578063e245386f14610138578063fc9539cd1461016c5761007d565b8063318c136e1461008257806345a104db146100a05780636ebcf607146100bc575b600080fd5b61008a610188565b6040516100979190610da8565b60405180910390f35b6100ba60048036038101906100b59190610b5c565b6101ae565b005b6100d660048036038101906100d19190610a51565b6103b8565b6040516100e39190610e5c565b60405180910390f35b61010660048036038101906101019190610a7a565b6103d0565b005b610122600480360381019061011d9190610b33565b6104fe565b60405161012f9190610e3a565b60405180910390f35b610152600480360381019061014d9190610b33565b6106a0565b604051610163959493929190610e77565b60405180910390f35b61018660048036038101906101819190610af2565b6107bb565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461023e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023590610dfa565b60405180910390fd5b600080606060208401519250604084015191506040519050836060015180820160600160405260e4360360e48337506040518060a001604052808860ff1681526020018381526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001828152506002600088815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040190805190602001906103ab9291906108a3565b5090505050505050505050565b60006020528060005260406000206000915090505481565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b815260040161041293929190610dc3565b602060405180830381600087803b15801561042c57600080fd5b505af1158015610440573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104649190610ac9565b506104b6826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461084e90919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b610506610923565b600260008381526020019081526020016000206040518060a00160405290816000820160009054906101000a900460ff1660ff1660ff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106905780601f1061066557610100808354040283529160200191610690565b820191906000526020600020905b81548152906001019060200180831161067357829003601f168201915b5050505050815250509050919050565b60026020528060005260406000206000915090508060000160009054906101000a900460ff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806004018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107b15780601f10610786576101008083540402835291602001916107b1565b820191906000526020600020905b81548152906001019060200180831161079457829003601f168201915b5050505050905085565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461084b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084290610dfa565b60405180910390fd5b50565b600080828401905083811015610899576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089090610e1a565b60405180910390fd5b8091505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106108e457805160ff1916838001178555610912565b82800160010185558215610912579182015b828111156109115782518255916020019190600101906108f6565b5b50905061091f9190610984565b5090565b6040518060a00160405280600060ff16815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b6109a691905b808211156109a257600081600090555060010161098a565b5090565b90565b6000813590506109b88161101a565b92915050565b6000815190506109cd81611031565b92915050565b600082601f8301126109e457600080fd5b81356109f76109f282610efe565b610ed1565b91508082526020830160208301858383011115610a1357600080fd5b610a1e838284610fc7565b50505092915050565b600081359050610a3681611048565b92915050565b600081359050610a4b8161105f565b92915050565b600060208284031215610a6357600080fd5b6000610a71848285016109a9565b91505092915050565b600080600060608486031215610a8f57600080fd5b6000610a9d868287016109a9565b9350506020610aae868287016109a9565b9250506040610abf86828701610a27565b9150509250925092565b600060208284031215610adb57600080fd5b6000610ae9848285016109be565b91505092915050565b600060208284031215610b0457600080fd5b600082013567ffffffffffffffff811115610b1e57600080fd5b610b2a848285016109d3565b91505092915050565b600060208284031215610b4557600080fd5b6000610b5384828501610a27565b91505092915050565b60008060008060808587031215610b7257600080fd5b6000610b8087828801610a3c565b9450506020610b9187828801610a27565b9350506040610ba2878288016109a9565b925050606085013567ffffffffffffffff811115610bbf57600080fd5b610bcb878288016109d3565b91505092959194509250565b610be081610f68565b82525050565b610bef81610f68565b82525050565b610bfe81610f86565b82525050565b610c0d81610f86565b82525050565b6000610c1e82610f2a565b610c288185610f35565b9350610c38818560208601610fd6565b610c4181611009565b840191505092915050565b6000610c5782610f2a565b610c618185610f46565b9350610c71818560208601610fd6565b610c7a81611009565b840191505092915050565b6000610c92601e83610f57565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000610cd2601b83610f57565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b600060a083016000830151610d1d6000860182610d8a565b506020830151610d306020860182610bf5565b506040830151610d436040860182610bd7565b506060830151610d566060860182610bd7565b5060808301518482036080860152610d6e8282610c13565b9150508091505092915050565b610d8481610fb0565b82525050565b610d9381610fba565b82525050565b610da281610fba565b82525050565b6000602082019050610dbd6000830184610be6565b92915050565b6000606082019050610dd86000830186610be6565b610de56020830185610be6565b610df26040830184610d7b565b949350505050565b60006020820190508181036000830152610e1381610c85565b9050919050565b60006020820190508181036000830152610e3381610cc5565b9050919050565b60006020820190508181036000830152610e548184610d05565b905092915050565b6000602082019050610e716000830184610d7b565b92915050565b600060a082019050610e8c6000830188610d99565b610e996020830187610c04565b610ea66040830186610be6565b610eb36060830185610be6565b8181036080830152610ec58184610c4c565b90509695505050505050565b6000604051905081810181811067ffffffffffffffff82111715610ef457600080fd5b8060405250919050565b600067ffffffffffffffff821115610f1557600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610f7382610f90565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015610ff4578082015181840152602081019050610fd9565b83811115611003576000848401525b50505050565b6000601f19601f8301169050919050565b61102381610f68565b811461102e57600080fd5b50565b61103a81610f7a565b811461104557600080fd5b50565b61105181610fb0565b811461105c57600080fd5b50565b61106881610fba565b811461107357600080fd5b5056fea2646970667358221220337ec2ffd59f4d16694bfdcffc2c8795f772f53f416153ab35da08b303bacab564736f6c63430006040033"
