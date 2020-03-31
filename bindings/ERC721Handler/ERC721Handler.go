// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC721Handler

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

// ERC721HandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type ERC721HandlerDepositRecord struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationChainTokenAddress   common.Address
	DestinationRecipientAddress    common.Address
	Depositer                      common.Address
	TokenID                        *big.Int
	MetaData                       []byte
}

// ERC721HandlerABI is the input ABI used to generate the binding from.
const ERC721HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destinationChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destinationChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structERC721Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721HandlerBin is the compiled bytecode used for deploying new contracts.
var ERC721HandlerBin = "0x60806040523480156200001157600080fd5b506040516200180638038062001806833981810160405281019062000037919062000096565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000110565b6000815190506200009081620000f6565b92915050565b600060208284031215620000a957600080fd5b6000620000b9848285016200007f565b91505092915050565b6000620000cf82620000d6565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200010181620000c2565b81146200010d57600080fd5b50565b6116e680620001206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063d9caed121161005b578063d9caed12146100ec578063db95e75c14610108578063e245386f14610138578063fc9539cd1461016f5761007d565b8063318c136e146100825780636ebcf607146100a0578063cb65d221146100d0575b600080fd5b61008a61018b565b604051610097919061136f565b60405180910390f35b6100ba60048036038101906100b59190610f99565b6101b1565b6040516100c79190611508565b60405180910390f35b6100ea60048036038101906100e591906110a4565b6101c9565b005b61010660048036038101906101019190610fc2565b61050e565b005b610122600480360381019061011d919061107b565b6105af565b60405161012f91906114e6565b60405180910390f35b610152600480360381019061014d919061107b565b610841565b6040516101669897969594939291906113c1565b60405180910390f35b6101896004803603810190610184919061103a565b6109c1565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610259576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610250906114a6565b60405180910390fd5b600080600080600060606020870151955060408701519450606087015193506080870151925060a0870151915060405190508660c0015180820160a00160405261014436036101448337506102b086893085610a96565b6040518061010001604052808773ffffffffffffffffffffffffffffffffffffffff1681526020018b81526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815250600260008b815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816006015560e08201518160070190805190602001906104fe929190610dae565b5090505050505050505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461059e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610595906114a6565b60405180910390fd5b6105aa83308484610ba5565b505050565b6105b7610e2e565b60026000838152602001908152602001600020604051806101000160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160068201548152602001600782018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108315780601f1061080657610100808354040283529160200191610831565b820191906000526020600020905b81548152906001019060200180831161081457829003601f168201915b5050505050815250509050919050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806006015490806007018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109b75780601f1061098c576101008083540402835291602001916109b7565b820191906000526020600020905b81548152906001019060200180831161099a57829003601f168201915b5050505050905088565b600080600060606020850151935060408501519250606085015191506040519050846080015180820160600160405260a4360360a483375060008490508073ffffffffffffffffffffffffffffffffffffffff16638832e6e38585856040518463ffffffff1660e01b8152600401610a3b93929190611446565b602060405180830381600087803b158015610a5557600080fd5b505af1158015610a69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a8d9190611011565b50505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610ad89392919061138a565b600060405180830381600087803b158015610af257600080fd5b505af1158015610b06573d6000803e3d6000fd5b50505050610b5c60016000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610cb490919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610be79392919061138a565b600060405180830381600087803b158015610c0157600080fd5b505af1158015610c15573d6000803e3d6000fd5b50505050610c6b60016000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610d0990919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b600080828401905083811015610cff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf6906114c6565b60405180910390fd5b8091505092915050565b6000610d4b83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610d53565b905092915050565b6000838311158290610d9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d929190611484565b60405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610def57805160ff1916838001178555610e1d565b82800160010185558215610e1d579182015b82811115610e1c578251825591602001919060010190610e01565b5b509050610e2a9190610ee1565b5090565b604051806101000160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001606081525090565b610f0391905b80821115610eff576000816000905550600101610ee7565b5090565b90565b600081359050610f158161166b565b92915050565b600081519050610f2a81611682565b92915050565b600082601f830112610f4157600080fd5b8135610f54610f4f82611550565b611523565b91508082526020830160208301858383011115610f7057600080fd5b610f7b838284611618565b50505092915050565b600081359050610f9381611699565b92915050565b600060208284031215610fab57600080fd5b6000610fb984828501610f06565b91505092915050565b600080600060608486031215610fd757600080fd5b6000610fe586828701610f06565b9350506020610ff686828701610f06565b925050604061100786828701610f84565b9150509250925092565b60006020828403121561102357600080fd5b600061103184828501610f1b565b91505092915050565b60006020828403121561104c57600080fd5b600082013567ffffffffffffffff81111561106657600080fd5b61107284828501610f30565b91505092915050565b60006020828403121561108d57600080fd5b600061109b84828501610f84565b91505092915050565b600080600080608085870312156110ba57600080fd5b60006110c887828801610f84565b94505060206110d987828801610f84565b93505060406110ea87828801610f06565b925050606085013567ffffffffffffffff81111561110757600080fd5b61111387828801610f30565b91505092959194509250565b611128816115d0565b82525050565b611137816115d0565b82525050565b600061114882611587565b61115281856115ae565b9350611162818560208601611627565b61116b8161165a565b840191505092915050565b60006111818261157c565b61118b818561159d565b935061119b818560208601611627565b6111a48161165a565b840191505092915050565b60006111ba8261157c565b6111c481856115ae565b93506111d4818560208601611627565b6111dd8161165a565b840191505092915050565b60006111f382611592565b6111fd81856115bf565b935061120d818560208601611627565b6112168161165a565b840191505092915050565b600061122e601e836115bf565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b600061126e601b836115bf565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000610100830160008301516112ba600086018261111f565b5060208301516112cd6020860182611351565b5060408301516112e0604086018261111f565b5060608301516112f3606086018261111f565b506080830151611306608086018261111f565b5060a083015161131960a086018261111f565b5060c083015161132c60c0860182611351565b5060e083015184820360e08601526113448282611176565b9150508091505092915050565b61135a8161160e565b82525050565b6113698161160e565b82525050565b6000602082019050611384600083018461112e565b92915050565b600060608201905061139f600083018661112e565b6113ac602083018561112e565b6113b96040830184611360565b949350505050565b6000610100820190506113d7600083018b61112e565b6113e4602083018a611360565b6113f1604083018961112e565b6113fe606083018861112e565b61140b608083018761112e565b61141860a083018661112e565b61142560c0830185611360565b81810360e083015261143781846111af565b90509998505050505050505050565b600060608201905061145b600083018661112e565b6114686020830185611360565b818103604083015261147a818461113d565b9050949350505050565b6000602082019050818103600083015261149e81846111e8565b905092915050565b600060208201905081810360008301526114bf81611221565b9050919050565b600060208201905081810360008301526114df81611261565b9050919050565b6000602082019050818103600083015261150081846112a1565b905092915050565b600060208201905061151d6000830184611360565b92915050565b6000604051905081810181811067ffffffffffffffff8211171561154657600080fd5b8060405250919050565b600067ffffffffffffffff82111561156757600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006115db826115ee565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561164557808201518184015260208101905061162a565b83811115611654576000848401525b50505050565b6000601f19601f8301169050919050565b611674816115d0565b811461167f57600080fd5b50565b61168b816115e2565b811461169657600080fd5b50565b6116a28161160e565b81146116ad57600080fd5b5056fea26469706673582212207a17b2d3d2bb40b2e95c9894e366807991e3384ff4e382756bb8912159a0b69b64736f6c63430006040033"

// DeployERC721Handler deploys a new Ethereum contract, binding an instance of ERC721Handler to it.
func DeployERC721Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *ERC721Handler, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721HandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721HandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Handler{ERC721HandlerCaller: ERC721HandlerCaller{contract: contract}, ERC721HandlerTransactor: ERC721HandlerTransactor{contract: contract}, ERC721HandlerFilterer: ERC721HandlerFilterer{contract: contract}}, nil
}

// ERC721Handler is an auto generated Go binding around an Ethereum contract.
type ERC721Handler struct {
	ERC721HandlerCaller     // Read-only binding to the contract
	ERC721HandlerTransactor // Write-only binding to the contract
	ERC721HandlerFilterer   // Log filterer for contract events
}

// ERC721HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721HandlerSession struct {
	Contract     *ERC721Handler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721HandlerCallerSession struct {
	Contract *ERC721HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC721HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721HandlerTransactorSession struct {
	Contract     *ERC721HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC721HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721HandlerRaw struct {
	Contract *ERC721Handler // Generic contract binding to access the raw methods on
}

// ERC721HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721HandlerCallerRaw struct {
	Contract *ERC721HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721HandlerTransactorRaw struct {
	Contract *ERC721HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Handler creates a new instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721Handler(address common.Address, backend bind.ContractBackend) (*ERC721Handler, error) {
	contract, err := bindERC721Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Handler{ERC721HandlerCaller: ERC721HandlerCaller{contract: contract}, ERC721HandlerTransactor: ERC721HandlerTransactor{contract: contract}, ERC721HandlerFilterer: ERC721HandlerFilterer{contract: contract}}, nil
}

// NewERC721HandlerCaller creates a new read-only instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC721HandlerCaller, error) {
	contract, err := bindERC721Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerCaller{contract: contract}, nil
}

// NewERC721HandlerTransactor creates a new write-only instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721HandlerTransactor, error) {
	contract, err := bindERC721Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerTransactor{contract: contract}, nil
}

// NewERC721HandlerFilterer creates a new log filterer instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721HandlerFilterer, error) {
	contract, err := bindERC721Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerFilterer{contract: contract}, nil
}

// bindERC721Handler binds a generic wrapper to an already deployed contract.
func bindERC721Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Handler *ERC721HandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Handler.Contract.ERC721HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Handler *ERC721HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ERC721HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Handler *ERC721HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ERC721HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Handler *ERC721HandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Handler *ERC721HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Handler *ERC721HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Handler.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC721Handler *ERC721HandlerCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Handler.contract.Call(opts, out, "_balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC721Handler *ERC721HandlerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC721Handler.Contract.Balances(&_ERC721Handler.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC721Handler *ERC721HandlerCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC721Handler.Contract.Balances(&_ERC721Handler.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() constant returns(address)
func (_ERC721Handler *ERC721HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Handler.contract.Call(opts, out, "_bridgeAddress")
	return *ret0, err
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() constant returns(address)
func (_ERC721Handler *ERC721HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC721Handler.Contract.BridgeAddress(&_ERC721Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() constant returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC721Handler.Contract.BridgeAddress(&_ERC721Handler.CallOpts)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(address _originChainTokenAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationChainTokenAddress, address _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
func (_ERC721Handler *ERC721HandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationChainTokenAddress   common.Address
	DestinationRecipientAddress    common.Address
	Depositer                      common.Address
	TokenID                        *big.Int
	MetaData                       []byte
}, error) {
	ret := new(struct {
		OriginChainTokenAddress        common.Address
		DestinationChainID             *big.Int
		DestinationChainHandlerAddress common.Address
		DestinationChainTokenAddress   common.Address
		DestinationRecipientAddress    common.Address
		Depositer                      common.Address
		TokenID                        *big.Int
		MetaData                       []byte
	})
	out := ret
	err := _ERC721Handler.contract.Call(opts, out, "_depositRecords", arg0)
	return *ret, err
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(address _originChainTokenAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationChainTokenAddress, address _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
func (_ERC721Handler *ERC721HandlerSession) DepositRecords(arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationChainTokenAddress   common.Address
	DestinationRecipientAddress    common.Address
	Depositer                      common.Address
	TokenID                        *big.Int
	MetaData                       []byte
}, error) {
	return _ERC721Handler.Contract.DepositRecords(&_ERC721Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(address _originChainTokenAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationChainTokenAddress, address _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
func (_ERC721Handler *ERC721HandlerCallerSession) DepositRecords(arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationChainTokenAddress   common.Address
	DestinationRecipientAddress    common.Address
	Depositer                      common.Address
	TokenID                        *big.Int
	MetaData                       []byte
}, error) {
	return _ERC721Handler.Contract.DepositRecords(&_ERC721Handler.CallOpts, arg0)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) constant returns(ERC721HandlerDepositRecord)
func (_ERC721Handler *ERC721HandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositID *big.Int) (ERC721HandlerDepositRecord, error) {
	var (
		ret0 = new(ERC721HandlerDepositRecord)
	)
	out := ret0
	err := _ERC721Handler.contract.Call(opts, out, "getDepositRecord", depositID)
	return *ret0, err
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) constant returns(ERC721HandlerDepositRecord)
func (_ERC721Handler *ERC721HandlerSession) GetDepositRecord(depositID *big.Int) (ERC721HandlerDepositRecord, error) {
	return _ERC721Handler.Contract.GetDepositRecord(&_ERC721Handler.CallOpts, depositID)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) constant returns(ERC721HandlerDepositRecord)
func (_ERC721Handler *ERC721HandlerCallerSession) GetDepositRecord(depositID *big.Int) (ERC721HandlerDepositRecord, error) {
	return _ERC721Handler.Contract.GetDepositRecord(&_ERC721Handler.CallOpts, depositID)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactor) Deposit(opts *bind.TransactOpts, destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "deposit", destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerSession) Deposit(destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) Deposit(destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, destinationChainID, depositNonce, depositer, data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactor) ExecuteDeposit(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "executeDeposit", data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_ERC721Handler *ERC721HandlerSession) ExecuteDeposit(data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ExecuteDeposit(&_ERC721Handler.TransactOpts, data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) ExecuteDeposit(data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ExecuteDeposit(&_ERC721Handler.TransactOpts, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "withdraw", tokenAddress, recipient, tokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerSession) Withdraw(tokenAddress common.Address, recipient common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Withdraw(&_ERC721Handler.TransactOpts, tokenAddress, recipient, tokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) Withdraw(tokenAddress common.Address, recipient common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Withdraw(&_ERC721Handler.TransactOpts, tokenAddress, recipient, tokenID)
}
var RuntimeBytecode = "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c8063d9caed121161005b578063d9caed12146100ec578063db95e75c14610108578063e245386f14610138578063fc9539cd1461016f5761007d565b8063318c136e146100825780636ebcf607146100a0578063cb65d221146100d0575b600080fd5b61008a61018b565b604051610097919061136f565b60405180910390f35b6100ba60048036038101906100b59190610f99565b6101b1565b6040516100c79190611508565b60405180910390f35b6100ea60048036038101906100e591906110a4565b6101c9565b005b61010660048036038101906101019190610fc2565b61050e565b005b610122600480360381019061011d919061107b565b6105af565b60405161012f91906114e6565b60405180910390f35b610152600480360381019061014d919061107b565b610841565b6040516101669897969594939291906113c1565b60405180910390f35b6101896004803603810190610184919061103a565b6109c1565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610259576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610250906114a6565b60405180910390fd5b600080600080600060606020870151955060408701519450606087015193506080870151925060a0870151915060405190508660c0015180820160a00160405261014436036101448337506102b086893085610a96565b6040518061010001604052808773ffffffffffffffffffffffffffffffffffffffff1681526020018b81526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815250600260008b815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816006015560e08201518160070190805190602001906104fe929190610dae565b5090505050505050505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461059e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610595906114a6565b60405180910390fd5b6105aa83308484610ba5565b505050565b6105b7610e2e565b60026000838152602001908152602001600020604051806101000160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160068201548152602001600782018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108315780601f1061080657610100808354040283529160200191610831565b820191906000526020600020905b81548152906001019060200180831161081457829003601f168201915b5050505050815250509050919050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806006015490806007018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109b75780601f1061098c576101008083540402835291602001916109b7565b820191906000526020600020905b81548152906001019060200180831161099a57829003601f168201915b5050505050905088565b600080600060606020850151935060408501519250606085015191506040519050846080015180820160600160405260a4360360a483375060008490508073ffffffffffffffffffffffffffffffffffffffff16638832e6e38585856040518463ffffffff1660e01b8152600401610a3b93929190611446565b602060405180830381600087803b158015610a5557600080fd5b505af1158015610a69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a8d9190611011565b50505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610ad89392919061138a565b600060405180830381600087803b158015610af257600080fd5b505af1158015610b06573d6000803e3d6000fd5b50505050610b5c60016000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610cb490919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610be79392919061138a565b600060405180830381600087803b158015610c0157600080fd5b505af1158015610c15573d6000803e3d6000fd5b50505050610c6b60016000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610d0990919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b600080828401905083811015610cff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf6906114c6565b60405180910390fd5b8091505092915050565b6000610d4b83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610d53565b905092915050565b6000838311158290610d9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d929190611484565b60405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610def57805160ff1916838001178555610e1d565b82800160010185558215610e1d579182015b82811115610e1c578251825591602001919060010190610e01565b5b509050610e2a9190610ee1565b5090565b604051806101000160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001606081525090565b610f0391905b80821115610eff576000816000905550600101610ee7565b5090565b90565b600081359050610f158161166b565b92915050565b600081519050610f2a81611682565b92915050565b600082601f830112610f4157600080fd5b8135610f54610f4f82611550565b611523565b91508082526020830160208301858383011115610f7057600080fd5b610f7b838284611618565b50505092915050565b600081359050610f9381611699565b92915050565b600060208284031215610fab57600080fd5b6000610fb984828501610f06565b91505092915050565b600080600060608486031215610fd757600080fd5b6000610fe586828701610f06565b9350506020610ff686828701610f06565b925050604061100786828701610f84565b9150509250925092565b60006020828403121561102357600080fd5b600061103184828501610f1b565b91505092915050565b60006020828403121561104c57600080fd5b600082013567ffffffffffffffff81111561106657600080fd5b61107284828501610f30565b91505092915050565b60006020828403121561108d57600080fd5b600061109b84828501610f84565b91505092915050565b600080600080608085870312156110ba57600080fd5b60006110c887828801610f84565b94505060206110d987828801610f84565b93505060406110ea87828801610f06565b925050606085013567ffffffffffffffff81111561110757600080fd5b61111387828801610f30565b91505092959194509250565b611128816115d0565b82525050565b611137816115d0565b82525050565b600061114882611587565b61115281856115ae565b9350611162818560208601611627565b61116b8161165a565b840191505092915050565b60006111818261157c565b61118b818561159d565b935061119b818560208601611627565b6111a48161165a565b840191505092915050565b60006111ba8261157c565b6111c481856115ae565b93506111d4818560208601611627565b6111dd8161165a565b840191505092915050565b60006111f382611592565b6111fd81856115bf565b935061120d818560208601611627565b6112168161165a565b840191505092915050565b600061122e601e836115bf565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b600061126e601b836115bf565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000610100830160008301516112ba600086018261111f565b5060208301516112cd6020860182611351565b5060408301516112e0604086018261111f565b5060608301516112f3606086018261111f565b506080830151611306608086018261111f565b5060a083015161131960a086018261111f565b5060c083015161132c60c0860182611351565b5060e083015184820360e08601526113448282611176565b9150508091505092915050565b61135a8161160e565b82525050565b6113698161160e565b82525050565b6000602082019050611384600083018461112e565b92915050565b600060608201905061139f600083018661112e565b6113ac602083018561112e565b6113b96040830184611360565b949350505050565b6000610100820190506113d7600083018b61112e565b6113e4602083018a611360565b6113f1604083018961112e565b6113fe606083018861112e565b61140b608083018761112e565b61141860a083018661112e565b61142560c0830185611360565b81810360e083015261143781846111af565b90509998505050505050505050565b600060608201905061145b600083018661112e565b6114686020830185611360565b818103604083015261147a818461113d565b9050949350505050565b6000602082019050818103600083015261149e81846111e8565b905092915050565b600060208201905081810360008301526114bf81611221565b9050919050565b600060208201905081810360008301526114df81611261565b9050919050565b6000602082019050818103600083015261150081846112a1565b905092915050565b600060208201905061151d6000830184611360565b92915050565b6000604051905081810181811067ffffffffffffffff8211171561154657600080fd5b8060405250919050565b600067ffffffffffffffff82111561156757600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006115db826115ee565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561164557808201518184015260208101905061162a565b83811115611654576000848401525b50505050565b6000601f19601f8301169050919050565b611674816115d0565b811461167f57600080fd5b50565b61168b816115e2565b811461169657600080fd5b50565b6116a28161160e565b81146116ad57600080fd5b5056fea26469706673582212207a17b2d3d2bb40b2e95c9894e366807991e3384ff4e382756bb8912159a0b69b64736f6c63430006040033"
