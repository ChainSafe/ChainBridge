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
	OriginChainContractAddress  common.Address
	DestinationChainID          uint8
	ResourceID                  [32]byte
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaDataHash                [32]byte
}

// CentrifugeAssetHandlerABI is the input ABI used to generate the binding from.
const CentrifugeAssetHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"originChainContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metaDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structCentrifugeAssetHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CentrifugeAssetHandlerBin is the compiled bytecode used for deploying new contracts.
var CentrifugeAssetHandlerBin = "0x60806040523480156200001157600080fd5b506040516200130138038062001301833981810160405281019062000037919062000095565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200010f565b6000815190506200008f81620000f5565b92915050565b600060208284031215620000a857600080fd5b6000620000b8848285016200007e565b91505092915050565b6000620000ce82620000d5565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200010081620000c1565b81146200010c57600080fd5b50565b6111e2806200011f6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806345a104db1161005b57806345a104db14610100578063c8ba6c871461011c578063db95e75c1461014c578063fc9539cd1461017c5761007d565b80630a6d55d814610082578063318c136e146100b25780633cf5040a146100d0575b600080fd5b61009c60048036038101906100979190610b19565b610198565b6040516100a99190610ed0565b60405180910390f35b6100ba6101cb565b6040516100c79190610ed0565b60405180910390f35b6100ea60048036038101906100e59190610b19565b6101f0565b6040516100f79190610eeb565b60405180910390f35b61011a60048036038101906101159190610bd5565b61021a565b005b61013660048036038101906101319190610af0565b6106b0565b6040516101439190610f06565b60405180910390f35b61016660048036038101906101619190610b83565b6106c8565b6040516101739190610fa3565b60405180910390f35b61019660048036038101906101919190610b42565b610828565b005b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006004600083815260200190815260200160002060009054906101000a900460ff169050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a090610f63565b60405180910390fd5b60008060006020840151925060408401519150606084015190506000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008060405160200161031a9190610eb5565b60405160208183030381529060405280519060200120826040516020016103419190610eb5565b6040516020818303038152906040528051906020012014156104ac5760008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663beab71316040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156103cd57600080fd5b505af11580156103e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104059190610bac565b90506104118782610957565b935083600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550866001600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505b600015156004600085815260200190815260200160002060009054906101000a900460ff16151514610513576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050a90610f83565b60405180910390fd5b6040518060c001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018a60ff1681526020018381526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200184815250600360008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff1602179055506040820151816001015560608201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160040155905050505050505050505050565b60026020528060005260406000206000915090505481565b6106d06109b2565b600360008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820154815250509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ae90610f63565b60405180910390fd5b600060208201519050600015156004600083815260200190815260200160002060009054906101000a900460ff16151514610927576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091e90610f43565b60405180910390fd5b60016004600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60006060838360405160200161096e929190610e89565b60405160208183030381529060405260405160200161098d9190610f21565b6040516020818303038152906040529050600060208201519050809250505092915050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff16815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600080191681525090565b600081359050610a4281611150565b92915050565b600081359050610a5781611167565b92915050565b600082601f830112610a6e57600080fd5b8135610a81610a7c82610feb565b610fbe565b91508082526020830160208301858383011115610a9d57600080fd5b610aa88382846110a3565b50505092915050565b600081359050610ac08161117e565b92915050565b600081359050610ad581611195565b92915050565b600081519050610aea81611195565b92915050565b600060208284031215610b0257600080fd5b6000610b1084828501610a33565b91505092915050565b600060208284031215610b2b57600080fd5b6000610b3984828501610a48565b91505092915050565b600060208284031215610b5457600080fd5b600082013567ffffffffffffffff811115610b6e57600080fd5b610b7a84828501610a5d565b91505092915050565b600060208284031215610b9557600080fd5b6000610ba384828501610ab1565b91505092915050565b600060208284031215610bbe57600080fd5b6000610bcc84828501610adb565b91505092915050565b60008060008060808587031215610beb57600080fd5b6000610bf987828801610ac6565b9450506020610c0a87828801610ab1565b9350506040610c1b87828801610a33565b925050606085013567ffffffffffffffff811115610c3857600080fd5b610c4487828801610a5d565b91505092959194509250565b610c5981611044565b82525050565b610c6881611044565b82525050565b610c7f610c7a82611044565b6110e5565b82525050565b610c8e81611056565b82525050565b610c9d81611062565b82525050565b610cac81611062565b82525050565b610cc3610cbe82611062565b6110f7565b82525050565b6000610cd482611017565b610cde8185611022565b9350610cee8185602086016110b2565b610cf781611125565b840191505092915050565b6000610d0f601983611033565b91507f617373657420686173206265656e206465706f736974656421000000000000006000830152602082019050919050565b6000610d4f601e83611033565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000610d8f603683611033565b91507f61737365742068617320616c7265616479206265656e206465706f736974656460008301527f20616e642063616e6e6f74206265206368616e676564000000000000000000006020830152604082019050919050565b60c082016000820151610dfe6000850182610c50565b506020820151610e116020850182610e63565b506040820151610e246040850182610c94565b506060820151610e376060850182610c50565b506080820151610e4a6080850182610c50565b5060a0820151610e5d60a0850182610c94565b50505050565b610e6c81611096565b82525050565b610e83610e7e82611096565b611113565b82525050565b6000610e958285610c6e565b601482019150610ea58284610e72565b6001820191508190509392505050565b6000610ec18284610cb2565b60208201915081905092915050565b6000602082019050610ee56000830184610c5f565b92915050565b6000602082019050610f006000830184610c85565b92915050565b6000602082019050610f1b6000830184610ca3565b92915050565b60006020820190508181036000830152610f3b8184610cc9565b905092915050565b60006020820190508181036000830152610f5c81610d02565b9050919050565b60006020820190508181036000830152610f7c81610d42565b9050919050565b60006020820190508181036000830152610f9c81610d82565b9050919050565b600060c082019050610fb86000830184610de8565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610fe157600080fd5b8060405250919050565b600067ffffffffffffffff82111561100257600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600061104f8261106c565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156110d05780820151818401526020810190506110b5565b838111156110df576000848401525b50505050565b60006110f082611101565b9050919050565b6000819050919050565b600061110c82611143565b9050919050565b600061111e82611136565b9050919050565b6000601f19601f8301169050919050565b60008160f81b9050919050565b60008160601b9050919050565b61115981611044565b811461116457600080fd5b50565b61117081611062565b811461117b57600080fd5b50565b6111878161108c565b811461119257600080fd5b50565b61119e81611096565b81146111a957600080fd5b5056fea2646970667358221220134782c61740f625396ca0f5ddddc06d47cea65e02547b1088ca319ebfa8534364736f6c63430006040033"

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

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) constant returns(address)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CentrifugeAssetHandler.contract.Call(opts, out, "_resourceIDToTokenContractAddress", arg0)
	return *ret0, err
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) constant returns(address)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _CentrifugeAssetHandler.Contract.ResourceIDToTokenContractAddress(&_CentrifugeAssetHandler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) constant returns(address)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _CentrifugeAssetHandler.Contract.ResourceIDToTokenContractAddress(&_CentrifugeAssetHandler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) constant returns(bytes32)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _CentrifugeAssetHandler.contract.Call(opts, out, "_tokenContractAddressToResourceID", arg0)
	return *ret0, err
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) constant returns(bytes32)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _CentrifugeAssetHandler.Contract.TokenContractAddressToResourceID(&_CentrifugeAssetHandler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) constant returns(bytes32)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _CentrifugeAssetHandler.Contract.TokenContractAddressToResourceID(&_CentrifugeAssetHandler.CallOpts, arg0)
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

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 originChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerTransactor) Deposit(opts *bind.TransactOpts, originChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.contract.Transact(opts, "deposit", originChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 originChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerSession) Deposit(originChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.Contract.Deposit(&_CentrifugeAssetHandler.TransactOpts, originChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 originChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerTransactorSession) Deposit(originChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
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

var RuntimeBytecode = "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c806345a104db1161005b57806345a104db14610100578063c8ba6c871461011c578063db95e75c1461014c578063fc9539cd1461017c5761007d565b80630a6d55d814610082578063318c136e146100b25780633cf5040a146100d0575b600080fd5b61009c60048036038101906100979190610b19565b610198565b6040516100a99190610ed0565b60405180910390f35b6100ba6101cb565b6040516100c79190610ed0565b60405180910390f35b6100ea60048036038101906100e59190610b19565b6101f0565b6040516100f79190610eeb565b60405180910390f35b61011a60048036038101906101159190610bd5565b61021a565b005b61013660048036038101906101319190610af0565b6106b0565b6040516101439190610f06565b60405180910390f35b61016660048036038101906101619190610b83565b6106c8565b6040516101739190610fa3565b60405180910390f35b61019660048036038101906101919190610b42565b610828565b005b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006004600083815260200190815260200160002060009054906101000a900460ff169050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a090610f63565b60405180910390fd5b60008060006020840151925060408401519150606084015190506000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008060405160200161031a9190610eb5565b60405160208183030381529060405280519060200120826040516020016103419190610eb5565b6040516020818303038152906040528051906020012014156104ac5760008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663beab71316040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156103cd57600080fd5b505af11580156103e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104059190610bac565b90506104118782610957565b935083600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550866001600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505b600015156004600085815260200190815260200160002060009054906101000a900460ff16151514610513576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050a90610f83565b60405180910390fd5b6040518060c001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018a60ff1681526020018381526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200184815250600360008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff1602179055506040820151816001015560608201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160040155905050505050505050505050565b60026020528060005260406000206000915090505481565b6106d06109b2565b600360008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820154815250509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ae90610f63565b60405180910390fd5b600060208201519050600015156004600083815260200190815260200160002060009054906101000a900460ff16151514610927576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091e90610f43565b60405180910390fd5b60016004600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60006060838360405160200161096e929190610e89565b60405160208183030381529060405260405160200161098d9190610f21565b6040516020818303038152906040529050600060208201519050809250505092915050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff16815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600080191681525090565b600081359050610a4281611150565b92915050565b600081359050610a5781611167565b92915050565b600082601f830112610a6e57600080fd5b8135610a81610a7c82610feb565b610fbe565b91508082526020830160208301858383011115610a9d57600080fd5b610aa88382846110a3565b50505092915050565b600081359050610ac08161117e565b92915050565b600081359050610ad581611195565b92915050565b600081519050610aea81611195565b92915050565b600060208284031215610b0257600080fd5b6000610b1084828501610a33565b91505092915050565b600060208284031215610b2b57600080fd5b6000610b3984828501610a48565b91505092915050565b600060208284031215610b5457600080fd5b600082013567ffffffffffffffff811115610b6e57600080fd5b610b7a84828501610a5d565b91505092915050565b600060208284031215610b9557600080fd5b6000610ba384828501610ab1565b91505092915050565b600060208284031215610bbe57600080fd5b6000610bcc84828501610adb565b91505092915050565b60008060008060808587031215610beb57600080fd5b6000610bf987828801610ac6565b9450506020610c0a87828801610ab1565b9350506040610c1b87828801610a33565b925050606085013567ffffffffffffffff811115610c3857600080fd5b610c4487828801610a5d565b91505092959194509250565b610c5981611044565b82525050565b610c6881611044565b82525050565b610c7f610c7a82611044565b6110e5565b82525050565b610c8e81611056565b82525050565b610c9d81611062565b82525050565b610cac81611062565b82525050565b610cc3610cbe82611062565b6110f7565b82525050565b6000610cd482611017565b610cde8185611022565b9350610cee8185602086016110b2565b610cf781611125565b840191505092915050565b6000610d0f601983611033565b91507f617373657420686173206265656e206465706f736974656421000000000000006000830152602082019050919050565b6000610d4f601e83611033565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000610d8f603683611033565b91507f61737365742068617320616c7265616479206265656e206465706f736974656460008301527f20616e642063616e6e6f74206265206368616e676564000000000000000000006020830152604082019050919050565b60c082016000820151610dfe6000850182610c50565b506020820151610e116020850182610e63565b506040820151610e246040850182610c94565b506060820151610e376060850182610c50565b506080820151610e4a6080850182610c50565b5060a0820151610e5d60a0850182610c94565b50505050565b610e6c81611096565b82525050565b610e83610e7e82611096565b611113565b82525050565b6000610e958285610c6e565b601482019150610ea58284610e72565b6001820191508190509392505050565b6000610ec18284610cb2565b60208201915081905092915050565b6000602082019050610ee56000830184610c5f565b92915050565b6000602082019050610f006000830184610c85565b92915050565b6000602082019050610f1b6000830184610ca3565b92915050565b60006020820190508181036000830152610f3b8184610cc9565b905092915050565b60006020820190508181036000830152610f5c81610d02565b9050919050565b60006020820190508181036000830152610f7c81610d42565b9050919050565b60006020820190508181036000830152610f9c81610d82565b9050919050565b600060c082019050610fb86000830184610de8565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610fe157600080fd5b8060405250919050565b600067ffffffffffffffff82111561100257600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600061104f8261106c565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156110d05780820151818401526020810190506110b5565b838111156110df576000848401525b50505050565b60006110f082611101565b9050919050565b6000819050919050565b600061110c82611143565b9050919050565b600061111e82611136565b9050919050565b6000601f19601f8301169050919050565b60008160f81b9050919050565b60008160601b9050919050565b61115981611044565b811461116457600080fd5b50565b61117081611062565b811461117b57600080fd5b50565b6111878161108c565b811461119257600080fd5b50565b61119e81611096565b81146111a957600080fd5b5056fea2646970667358221220134782c61740f625396ca0f5ddddc06d47cea65e02547b1088ca319ebfa8534364736f6c63430006040033"
