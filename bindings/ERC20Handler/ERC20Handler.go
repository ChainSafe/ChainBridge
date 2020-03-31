// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20Handler

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

// ERC20HandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type ERC20HandlerDepositRecord struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationChainTokenAddress   common.Address
	DestinationRecipientAddress    common.Address
	Depositer                      common.Address
	Amount                         *big.Int
}

// ERC20HandlerABI is the input ABI used to generate the binding from.
const ERC20HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destinationChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destinationChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"internalType\":\"structERC20Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20HandlerBin is the compiled bytecode used for deploying new contracts.
var ERC20HandlerBin = "0x60806040523480156200001157600080fd5b506040516200154a3803806200154a833981810160405281019062000037919062000096565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000110565b6000815190506200009081620000f6565b92915050565b600060208284031215620000a957600080fd5b6000620000b9848285016200007f565b91505092915050565b6000620000cf82620000d6565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200010181620000c2565b81146200010d57600080fd5b50565b61142a80620001206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063d9caed121161005b578063d9caed12146100ec578063db95e75c14610108578063e245386f14610138578063fc9539cd1461016e5761007d565b8063318c136e146100825780636ebcf607146100a0578063cb65d221146100d0575b600080fd5b61008a61018a565b604051610097919061111d565b60405180910390f35b6100ba60048036038101906100b59190610e14565b6101b0565b6040516100c79190611284565b60405180910390f35b6100ea60048036038101906100e59190610f1f565b6101c8565b005b61010660048036038101906101019190610e3d565b6104c8565b005b610122600480360381019061011d9190610ef6565b610569565b60405161012f9190611269565b60405180910390f35b610152600480360381019061014d9190610ef6565b610758565b6040516101659796959493929190611198565b60405180910390f35b61018860048036038101906101839190610eb5565b61083a565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610258576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024f90611229565b60405180910390fd5b60008060008060006020860151945060408601519350606086015192506080860151915060a0860151905061028f8588308461097e565b6040518060e001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018a81526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200182815250600260008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160060155905050505050505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610558576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054f90611229565b60405180910390fd5b61056483308484610aad565b505050565b610571610cd6565b600260008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820154815250509050919050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060060154905087565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c190611229565b60405180910390fd5b600080600060208401519250604084015191506060840151905060008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b815260040161092492919061116f565b602060405180830381600087803b15801561093e57600080fd5b505af1158015610952573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109769190610e8c565b505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b81526004016109c093929190611138565b602060405180830381600087803b1580156109da57600080fd5b505af11580156109ee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a129190610e8c565b50610a64826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610bdc90919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610aef93929190611138565b602060405180830381600087803b158015610b0957600080fd5b505af1158015610b1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b419190610e8c565b50610b93826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c3190919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b600080828401905083811015610c27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1e90611249565b60405180910390fd5b8091505092915050565b6000610c7383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610c7b565b905092915050565b6000838311158290610cc3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cba9190611207565b60405180910390fd5b5060008385039050809150509392505050565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600081359050610d90816113af565b92915050565b600081519050610da5816113c6565b92915050565b600082601f830112610dbc57600080fd5b8135610dcf610dca826112cc565b61129f565b91508082526020830160208301858383011115610deb57600080fd5b610df683828461135c565b50505092915050565b600081359050610e0e816113dd565b92915050565b600060208284031215610e2657600080fd5b6000610e3484828501610d81565b91505092915050565b600080600060608486031215610e5257600080fd5b6000610e6086828701610d81565b9350506020610e7186828701610d81565b9250506040610e8286828701610dff565b9150509250925092565b600060208284031215610e9e57600080fd5b6000610eac84828501610d96565b91505092915050565b600060208284031215610ec757600080fd5b600082013567ffffffffffffffff811115610ee157600080fd5b610eed84828501610dab565b91505092915050565b600060208284031215610f0857600080fd5b6000610f1684828501610dff565b91505092915050565b60008060008060808587031215610f3557600080fd5b6000610f4387828801610dff565b9450506020610f5487828801610dff565b9350506040610f6587828801610d81565b925050606085013567ffffffffffffffff811115610f8257600080fd5b610f8e87828801610dab565b91505092959194509250565b610fa381611314565b82525050565b610fb281611314565b82525050565b6000610fc3826112f8565b610fcd8185611303565b9350610fdd81856020860161136b565b610fe68161139e565b840191505092915050565b6000610ffe601e83611303565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b600061103e601b83611303565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b60e0820160008201516110876000850182610f9a565b50602082015161109a60208501826110ff565b5060408201516110ad6040850182610f9a565b5060608201516110c06060850182610f9a565b5060808201516110d36080850182610f9a565b5060a08201516110e660a0850182610f9a565b5060c08201516110f960c08501826110ff565b50505050565b61110881611352565b82525050565b61111781611352565b82525050565b60006020820190506111326000830184610fa9565b92915050565b600060608201905061114d6000830186610fa9565b61115a6020830185610fa9565b611167604083018461110e565b949350505050565b60006040820190506111846000830185610fa9565b611191602083018461110e565b9392505050565b600060e0820190506111ad600083018a610fa9565b6111ba602083018961110e565b6111c76040830188610fa9565b6111d46060830187610fa9565b6111e16080830186610fa9565b6111ee60a0830185610fa9565b6111fb60c083018461110e565b98975050505050505050565b600060208201905081810360008301526112218184610fb8565b905092915050565b6000602082019050818103600083015261124281610ff1565b9050919050565b6000602082019050818103600083015261126281611031565b9050919050565b600060e08201905061127e6000830184611071565b92915050565b6000602082019050611299600083018461110e565b92915050565b6000604051905081810181811067ffffffffffffffff821117156112c257600080fd5b8060405250919050565b600067ffffffffffffffff8211156112e357600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600061131f82611332565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561138957808201518184015260208101905061136e565b83811115611398576000848401525b50505050565b6000601f19601f8301169050919050565b6113b881611314565b81146113c357600080fd5b50565b6113cf81611326565b81146113da57600080fd5b50565b6113e681611352565b81146113f157600080fd5b5056fea2646970667358221220df2b14d6858f94a03b55525a9b1a557cc3d60d848809ff87a2bc0aeda1c861ba64736f6c63430006040033"

// DeployERC20Handler deploys a new Ethereum contract, binding an instance of ERC20Handler to it.
func DeployERC20Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *ERC20Handler, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20HandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// ERC20Handler is an auto generated Go binding around an Ethereum contract.
type ERC20Handler struct {
	ERC20HandlerCaller     // Read-only binding to the contract
	ERC20HandlerTransactor // Write-only binding to the contract
	ERC20HandlerFilterer   // Log filterer for contract events
}

// ERC20HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20HandlerSession struct {
	Contract     *ERC20Handler     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20HandlerCallerSession struct {
	Contract *ERC20HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC20HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20HandlerTransactorSession struct {
	Contract     *ERC20HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC20HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20HandlerRaw struct {
	Contract *ERC20Handler // Generic contract binding to access the raw methods on
}

// ERC20HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20HandlerCallerRaw struct {
	Contract *ERC20HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactorRaw struct {
	Contract *ERC20HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Handler creates a new instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20Handler(address common.Address, backend bind.ContractBackend) (*ERC20Handler, error) {
	contract, err := bindERC20Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// NewERC20HandlerCaller creates a new read-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC20HandlerCaller, error) {
	contract, err := bindERC20Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerCaller{contract: contract}, nil
}

// NewERC20HandlerTransactor creates a new write-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20HandlerTransactor, error) {
	contract, err := bindERC20Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerTransactor{contract: contract}, nil
}

// NewERC20HandlerFilterer creates a new log filterer instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20HandlerFilterer, error) {
	contract, err := bindERC20Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerFilterer{contract: contract}, nil
}

// bindERC20Handler binds a generic wrapper to an already deployed contract.
func bindERC20Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.ERC20HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC20Handler *ERC20HandlerCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC20Handler *ERC20HandlerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC20Handler.Contract.Balances(&_ERC20Handler.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) constant returns(uint256)
func (_ERC20Handler *ERC20HandlerCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC20Handler.Contract.Balances(&_ERC20Handler.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() constant returns(address)
func (_ERC20Handler *ERC20HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_bridgeAddress")
	return *ret0, err
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() constant returns(address)
func (_ERC20Handler *ERC20HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() constant returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(address _originChainTokenAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationChainTokenAddress, address _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationChainTokenAddress   common.Address
	DestinationRecipientAddress    common.Address
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	ret := new(struct {
		OriginChainTokenAddress        common.Address
		DestinationChainID             *big.Int
		DestinationChainHandlerAddress common.Address
		DestinationChainTokenAddress   common.Address
		DestinationRecipientAddress    common.Address
		Depositer                      common.Address
		Amount                         *big.Int
	})
	out := ret
	err := _ERC20Handler.contract.Call(opts, out, "_depositRecords", arg0)
	return *ret, err
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(address _originChainTokenAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationChainTokenAddress, address _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerSession) DepositRecords(arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationChainTokenAddress   common.Address
	DestinationRecipientAddress    common.Address
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	return _ERC20Handler.Contract.DepositRecords(&_ERC20Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(address _originChainTokenAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationChainTokenAddress, address _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerCallerSession) DepositRecords(arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationChainTokenAddress   common.Address
	DestinationRecipientAddress    common.Address
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	return _ERC20Handler.Contract.DepositRecords(&_ERC20Handler.CallOpts, arg0)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) constant returns(ERC20HandlerDepositRecord)
func (_ERC20Handler *ERC20HandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositID *big.Int) (ERC20HandlerDepositRecord, error) {
	var (
		ret0 = new(ERC20HandlerDepositRecord)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "getDepositRecord", depositID)
	return *ret0, err
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) constant returns(ERC20HandlerDepositRecord)
func (_ERC20Handler *ERC20HandlerSession) GetDepositRecord(depositID *big.Int) (ERC20HandlerDepositRecord, error) {
	return _ERC20Handler.Contract.GetDepositRecord(&_ERC20Handler.CallOpts, depositID)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) constant returns(ERC20HandlerDepositRecord)
func (_ERC20Handler *ERC20HandlerCallerSession) GetDepositRecord(depositID *big.Int) (ERC20HandlerDepositRecord, error) {
	return _ERC20Handler.Contract.GetDepositRecord(&_ERC20Handler.CallOpts, depositID)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactor) Deposit(opts *bind.TransactOpts, destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "deposit", destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerSession) Deposit(destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) Deposit(destinationChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, destinationChainID, depositNonce, depositer, data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactor) ExecuteDeposit(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "executeDeposit", data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_ERC20Handler *ERC20HandlerSession) ExecuteDeposit(data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteDeposit(&_ERC20Handler.TransactOpts, data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) ExecuteDeposit(data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteDeposit(&_ERC20Handler.TransactOpts, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "withdraw", tokenAddress, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerSession) Withdraw(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Withdraw(&_ERC20Handler.TransactOpts, tokenAddress, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) Withdraw(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Withdraw(&_ERC20Handler.TransactOpts, tokenAddress, recipient, amount)
}
var RuntimeBytecode = "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c8063d9caed121161005b578063d9caed12146100ec578063db95e75c14610108578063e245386f14610138578063fc9539cd1461016e5761007d565b8063318c136e146100825780636ebcf607146100a0578063cb65d221146100d0575b600080fd5b61008a61018a565b604051610097919061111d565b60405180910390f35b6100ba60048036038101906100b59190610e14565b6101b0565b6040516100c79190611284565b60405180910390f35b6100ea60048036038101906100e59190610f1f565b6101c8565b005b61010660048036038101906101019190610e3d565b6104c8565b005b610122600480360381019061011d9190610ef6565b610569565b60405161012f9190611269565b60405180910390f35b610152600480360381019061014d9190610ef6565b610758565b6040516101659796959493929190611198565b60405180910390f35b61018860048036038101906101839190610eb5565b61083a565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610258576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024f90611229565b60405180910390fd5b60008060008060006020860151945060408601519350606086015192506080860151915060a0860151905061028f8588308461097e565b6040518060e001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018a81526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200182815250600260008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160060155905050505050505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610558576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054f90611229565b60405180910390fd5b61056483308484610aad565b505050565b610571610cd6565b600260008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820154815250509050919050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060060154905087565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c190611229565b60405180910390fd5b600080600060208401519250604084015191506060840151905060008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b815260040161092492919061116f565b602060405180830381600087803b15801561093e57600080fd5b505af1158015610952573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109769190610e8c565b505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b81526004016109c093929190611138565b602060405180830381600087803b1580156109da57600080fd5b505af11580156109ee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a129190610e8c565b50610a64826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610bdc90919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610aef93929190611138565b602060405180830381600087803b158015610b0957600080fd5b505af1158015610b1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b419190610e8c565b50610b93826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c3190919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b600080828401905083811015610c27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1e90611249565b60405180910390fd5b8091505092915050565b6000610c7383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610c7b565b905092915050565b6000838311158290610cc3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cba9190611207565b60405180910390fd5b5060008385039050809150509392505050565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600081359050610d90816113af565b92915050565b600081519050610da5816113c6565b92915050565b600082601f830112610dbc57600080fd5b8135610dcf610dca826112cc565b61129f565b91508082526020830160208301858383011115610deb57600080fd5b610df683828461135c565b50505092915050565b600081359050610e0e816113dd565b92915050565b600060208284031215610e2657600080fd5b6000610e3484828501610d81565b91505092915050565b600080600060608486031215610e5257600080fd5b6000610e6086828701610d81565b9350506020610e7186828701610d81565b9250506040610e8286828701610dff565b9150509250925092565b600060208284031215610e9e57600080fd5b6000610eac84828501610d96565b91505092915050565b600060208284031215610ec757600080fd5b600082013567ffffffffffffffff811115610ee157600080fd5b610eed84828501610dab565b91505092915050565b600060208284031215610f0857600080fd5b6000610f1684828501610dff565b91505092915050565b60008060008060808587031215610f3557600080fd5b6000610f4387828801610dff565b9450506020610f5487828801610dff565b9350506040610f6587828801610d81565b925050606085013567ffffffffffffffff811115610f8257600080fd5b610f8e87828801610dab565b91505092959194509250565b610fa381611314565b82525050565b610fb281611314565b82525050565b6000610fc3826112f8565b610fcd8185611303565b9350610fdd81856020860161136b565b610fe68161139e565b840191505092915050565b6000610ffe601e83611303565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b600061103e601b83611303565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b60e0820160008201516110876000850182610f9a565b50602082015161109a60208501826110ff565b5060408201516110ad6040850182610f9a565b5060608201516110c06060850182610f9a565b5060808201516110d36080850182610f9a565b5060a08201516110e660a0850182610f9a565b5060c08201516110f960c08501826110ff565b50505050565b61110881611352565b82525050565b61111781611352565b82525050565b60006020820190506111326000830184610fa9565b92915050565b600060608201905061114d6000830186610fa9565b61115a6020830185610fa9565b611167604083018461110e565b949350505050565b60006040820190506111846000830185610fa9565b611191602083018461110e565b9392505050565b600060e0820190506111ad600083018a610fa9565b6111ba602083018961110e565b6111c76040830188610fa9565b6111d46060830187610fa9565b6111e16080830186610fa9565b6111ee60a0830185610fa9565b6111fb60c083018461110e565b98975050505050505050565b600060208201905081810360008301526112218184610fb8565b905092915050565b6000602082019050818103600083015261124281610ff1565b9050919050565b6000602082019050818103600083015261126281611031565b9050919050565b600060e08201905061127e6000830184611071565b92915050565b6000602082019050611299600083018461110e565b92915050565b6000604051905081810181811067ffffffffffffffff821117156112c257600080fd5b8060405250919050565b600067ffffffffffffffff8211156112e357600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600061131f82611332565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561138957808201518184015260208101905061136e565b83811115611398576000848401525b50505050565b6000601f19601f8301169050919050565b6113b881611314565b81146113c357600080fd5b50565b6113cf81611326565b81146113da57600080fd5b50565b6113e681611352565b81146113f157600080fd5b5056fea2646970667358221220df2b14d6858f94a03b55525a9b1a557cc3d60d848809ff87a2bc0aeda1c861ba64736f6c63430006040033"
