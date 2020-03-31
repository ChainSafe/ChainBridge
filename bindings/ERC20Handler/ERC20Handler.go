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
var ERC20HandlerBin = "0x60806040523480156200001157600080fd5b50604051620014ba380380620014ba833981810160405281019062000037919062000096565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000110565b6000815190506200009081620000f6565b92915050565b600060208284031215620000a957600080fd5b6000620000b9848285016200007f565b91505092915050565b6000620000cf82620000d6565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200010181620000c2565b81146200010d57600080fd5b50565b61139a80620001206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063d9caed121161005b578063d9caed12146100ec578063db95e75c14610108578063e245386f14610138578063fc9539cd1461016e5761007d565b8063318c136e146100825780636ebcf607146100a0578063cb65d221146100d0575b600080fd5b61008a61018a565b604051610097919061108d565b60405180910390f35b6100ba60048036038101906100b59190610d84565b6101b0565b6040516100c791906111f4565b60405180910390f35b6100ea60048036038101906100e59190610e8f565b6101c8565b005b61010660048036038101906101019190610dad565b6104c8565b005b610122600480360381019061011d9190610e66565b610569565b60405161012f91906111d9565b60405180910390f35b610152600480360381019061014d9190610e66565b610758565b6040516101659796959493929190611108565b60405180910390f35b61018860048036038101906101839190610e25565b61083a565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610258576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024f90611199565b60405180910390fd5b60008060008060006020860151945060408601519350606086015192506080860151915060a0860151905061028f858830846108ee565b6040518060e001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018a81526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200182815250600260008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160060155905050505050505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610558576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054f90611199565b60405180910390fd5b61056483308484610a1d565b505050565b610571610c46565b600260008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820154815250509050919050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060060154905087565b600080600060208401519250604084015191506060840151905060008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b81526004016108949291906110df565b602060405180830381600087803b1580156108ae57600080fd5b505af11580156108c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e69190610dfc565b505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610930939291906110a8565b602060405180830381600087803b15801561094a57600080fd5b505af115801561095e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109829190610dfc565b506109d4826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610b4c90919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610a5f939291906110a8565b602060405180830381600087803b158015610a7957600080fd5b505af1158015610a8d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab19190610dfc565b50610b03826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610ba190919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b600080828401905083811015610b97576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8e906111b9565b60405180910390fd5b8091505092915050565b6000610be383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610beb565b905092915050565b6000838311158290610c33576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2a9190611177565b60405180910390fd5b5060008385039050809150509392505050565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600081359050610d008161131f565b92915050565b600081519050610d1581611336565b92915050565b600082601f830112610d2c57600080fd5b8135610d3f610d3a8261123c565b61120f565b91508082526020830160208301858383011115610d5b57600080fd5b610d668382846112cc565b50505092915050565b600081359050610d7e8161134d565b92915050565b600060208284031215610d9657600080fd5b6000610da484828501610cf1565b91505092915050565b600080600060608486031215610dc257600080fd5b6000610dd086828701610cf1565b9350506020610de186828701610cf1565b9250506040610df286828701610d6f565b9150509250925092565b600060208284031215610e0e57600080fd5b6000610e1c84828501610d06565b91505092915050565b600060208284031215610e3757600080fd5b600082013567ffffffffffffffff811115610e5157600080fd5b610e5d84828501610d1b565b91505092915050565b600060208284031215610e7857600080fd5b6000610e8684828501610d6f565b91505092915050565b60008060008060808587031215610ea557600080fd5b6000610eb387828801610d6f565b9450506020610ec487828801610d6f565b9350506040610ed587828801610cf1565b925050606085013567ffffffffffffffff811115610ef257600080fd5b610efe87828801610d1b565b91505092959194509250565b610f1381611284565b82525050565b610f2281611284565b82525050565b6000610f3382611268565b610f3d8185611273565b9350610f4d8185602086016112db565b610f568161130e565b840191505092915050565b6000610f6e601e83611273565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000610fae601b83611273565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b60e082016000820151610ff76000850182610f0a565b50602082015161100a602085018261106f565b50604082015161101d6040850182610f0a565b5060608201516110306060850182610f0a565b5060808201516110436080850182610f0a565b5060a082015161105660a0850182610f0a565b5060c082015161106960c085018261106f565b50505050565b611078816112c2565b82525050565b611087816112c2565b82525050565b60006020820190506110a26000830184610f19565b92915050565b60006060820190506110bd6000830186610f19565b6110ca6020830185610f19565b6110d7604083018461107e565b949350505050565b60006040820190506110f46000830185610f19565b611101602083018461107e565b9392505050565b600060e08201905061111d600083018a610f19565b61112a602083018961107e565b6111376040830188610f19565b6111446060830187610f19565b6111516080830186610f19565b61115e60a0830185610f19565b61116b60c083018461107e565b98975050505050505050565b600060208201905081810360008301526111918184610f28565b905092915050565b600060208201905081810360008301526111b281610f61565b9050919050565b600060208201905081810360008301526111d281610fa1565b9050919050565b600060e0820190506111ee6000830184610fe1565b92915050565b6000602082019050611209600083018461107e565b92915050565b6000604051905081810181811067ffffffffffffffff8211171561123257600080fd5b8060405250919050565b600067ffffffffffffffff82111561125357600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600061128f826112a2565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156112f95780820151818401526020810190506112de565b83811115611308576000848401525b50505050565b6000601f19601f8301169050919050565b61132881611284565b811461133357600080fd5b50565b61133f81611296565b811461134a57600080fd5b50565b611356816112c2565b811461136157600080fd5b5056fea26469706673582212206067d1975411b41a9a303c1e9239cce5f81ed3089ec9545b702cb20dd9a0a20164736f6c63430006040033"

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
var RuntimeBytecode = "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c8063d9caed121161005b578063d9caed12146100ec578063db95e75c14610108578063e245386f14610138578063fc9539cd1461016e5761007d565b8063318c136e146100825780636ebcf607146100a0578063cb65d221146100d0575b600080fd5b61008a61018a565b604051610097919061108d565b60405180910390f35b6100ba60048036038101906100b59190610d84565b6101b0565b6040516100c791906111f4565b60405180910390f35b6100ea60048036038101906100e59190610e8f565b6101c8565b005b61010660048036038101906101019190610dad565b6104c8565b005b610122600480360381019061011d9190610e66565b610569565b60405161012f91906111d9565b60405180910390f35b610152600480360381019061014d9190610e66565b610758565b6040516101659796959493929190611108565b60405180910390f35b61018860048036038101906101839190610e25565b61083a565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610258576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024f90611199565b60405180910390fd5b60008060008060006020860151945060408601519350606086015192506080860151915060a0860151905061028f858830846108ee565b6040518060e001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018a81526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200182815250600260008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160060155905050505050505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610558576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054f90611199565b60405180910390fd5b61056483308484610a1d565b505050565b610571610c46565b600260008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820154815250509050919050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060060154905087565b600080600060208401519250604084015191506060840151905060008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b81526004016108949291906110df565b602060405180830381600087803b1580156108ae57600080fd5b505af11580156108c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e69190610dfc565b505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610930939291906110a8565b602060405180830381600087803b15801561094a57600080fd5b505af115801561095e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109829190610dfc565b506109d4826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610b4c90919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610a5f939291906110a8565b602060405180830381600087803b158015610a7957600080fd5b505af1158015610a8d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab19190610dfc565b50610b03826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610ba190919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b600080828401905083811015610b97576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8e906111b9565b60405180910390fd5b8091505092915050565b6000610be383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610beb565b905092915050565b6000838311158290610c33576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2a9190611177565b60405180910390fd5b5060008385039050809150509392505050565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600081359050610d008161131f565b92915050565b600081519050610d1581611336565b92915050565b600082601f830112610d2c57600080fd5b8135610d3f610d3a8261123c565b61120f565b91508082526020830160208301858383011115610d5b57600080fd5b610d668382846112cc565b50505092915050565b600081359050610d7e8161134d565b92915050565b600060208284031215610d9657600080fd5b6000610da484828501610cf1565b91505092915050565b600080600060608486031215610dc257600080fd5b6000610dd086828701610cf1565b9350506020610de186828701610cf1565b9250506040610df286828701610d6f565b9150509250925092565b600060208284031215610e0e57600080fd5b6000610e1c84828501610d06565b91505092915050565b600060208284031215610e3757600080fd5b600082013567ffffffffffffffff811115610e5157600080fd5b610e5d84828501610d1b565b91505092915050565b600060208284031215610e7857600080fd5b6000610e8684828501610d6f565b91505092915050565b60008060008060808587031215610ea557600080fd5b6000610eb387828801610d6f565b9450506020610ec487828801610d6f565b9350506040610ed587828801610cf1565b925050606085013567ffffffffffffffff811115610ef257600080fd5b610efe87828801610d1b565b91505092959194509250565b610f1381611284565b82525050565b610f2281611284565b82525050565b6000610f3382611268565b610f3d8185611273565b9350610f4d8185602086016112db565b610f568161130e565b840191505092915050565b6000610f6e601e83611273565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000610fae601b83611273565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b60e082016000820151610ff76000850182610f0a565b50602082015161100a602085018261106f565b50604082015161101d6040850182610f0a565b5060608201516110306060850182610f0a565b5060808201516110436080850182610f0a565b5060a082015161105660a0850182610f0a565b5060c082015161106960c085018261106f565b50505050565b611078816112c2565b82525050565b611087816112c2565b82525050565b60006020820190506110a26000830184610f19565b92915050565b60006060820190506110bd6000830186610f19565b6110ca6020830185610f19565b6110d7604083018461107e565b949350505050565b60006040820190506110f46000830185610f19565b611101602083018461107e565b9392505050565b600060e08201905061111d600083018a610f19565b61112a602083018961107e565b6111376040830188610f19565b6111446060830187610f19565b6111516080830186610f19565b61115e60a0830185610f19565b61116b60c083018461107e565b98975050505050505050565b600060208201905081810360008301526111918184610f28565b905092915050565b600060208201905081810360008301526111b281610f61565b9050919050565b600060208201905081810360008301526111d281610fa1565b9050919050565b600060e0820190506111ee6000830184610fe1565b92915050565b6000602082019050611209600083018461107e565b92915050565b6000604051905081810181811067ffffffffffffffff8211171561123257600080fd5b8060405250919050565b600067ffffffffffffffff82111561125357600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600061128f826112a2565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156112f95780820151818401526020810190506112de565b83811115611308576000848401525b50505050565b6000601f19601f8301169050919050565b61132881611284565b811461133357600080fd5b50565b61133f81611296565b811461134a57600080fd5b50565b611356816112c2565b811461136157600080fd5b5056fea26469706673582212206067d1975411b41a9a303c1e9239cce5f81ed3089ec9545b702cb20dd9a0a20164736f6c63430006040033"
