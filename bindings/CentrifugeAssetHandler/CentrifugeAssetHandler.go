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
	DestinationChainID          *big.Int
	ResourceID                  [32]byte
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaDataHash                [32]byte
}

// CentrifugeAssetHandlerABI is the input ABI used to generate the binding from.
const CentrifugeAssetHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"originChainContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metaDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structCentrifugeAssetHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CentrifugeAssetHandlerBin is the compiled bytecode used for deploying new contracts.
var CentrifugeAssetHandlerBin = "0x60806040523480156200001157600080fd5b50604051620012bc380380620012bc833981810160405281019062000037919062000095565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200010f565b6000815190506200008f81620000f5565b92915050565b600060208284031215620000a857600080fd5b6000620000b8848285016200007e565b91505092915050565b6000620000ce82620000d5565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200010081620000c1565b81146200010c57600080fd5b50565b61119d806200011f6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063c8ba6c871161005b578063c8ba6c8714610100578063cb65d22114610130578063db95e75c1461014c578063fc9539cd1461017c5761007d565b80630a6d55d814610082578063318c136e146100b25780633cf5040a146100d0575b600080fd5b61009c60048036038101906100979190610ad4565b610198565b6040516100a99190610e8b565b60405180910390f35b6100ba6101cb565b6040516100c79190610e8b565b60405180910390f35b6100ea60048036038101906100e59190610ad4565b6101f0565b6040516100f79190610ea6565b60405180910390f35b61011a60048036038101906101159190610aab565b61021a565b6040516101279190610ec1565b60405180910390f35b61014a60048036038101906101459190610b67565b610232565b005b61016660048036038101906101619190610b3e565b6106ae565b6040516101739190610f5e565b60405180910390f35b61019660048036038101906101919190610afd565b6107fb565b005b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006004600083815260200190815260200160002060009054906101000a900460ff169050919050565b60026020528060005260406000206000915090505481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b890610f1e565b60405180910390fd5b60008060006020840151925060408401519150606084015190506000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016103329190610e70565b60405160208183030381529060405280519060200120826040516020016103599190610e70565b6040516020818303038152906040528051906020012014156104c45760008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663beab71316040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156103e557600080fd5b505af11580156103f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061041d9190610be2565b9050610429878261092a565b935083600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550866001600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505b600015156004600085815260200190815260200160002060009054906101000a900460ff1615151461052b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052290610f3e565b60405180910390fd5b6040518060c001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018a81526020018381526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200184815250600360008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050155905050505050505050505050565b6106b6610985565b600360008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461088a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088190610f1e565b60405180910390fd5b600060208201519050600015156004600083815260200190815260200160002060009054906101000a900460ff161515146108fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f190610efe565b60405180910390fd5b60016004600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600060608383604051602001610941929190610e44565b6040516020818303038152906040526040516020016109609190610edc565b6040516020818303038152906040529050600060208201519050809250505092915050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600080191681525090565b600081359050610a128161110b565b92915050565b600081359050610a2781611122565b92915050565b600082601f830112610a3e57600080fd5b8135610a51610a4c82610fa6565b610f79565b91508082526020830160208301858383011115610a6d57600080fd5b610a7883828461105e565b50505092915050565b600081359050610a9081611139565b92915050565b600081519050610aa581611150565b92915050565b600060208284031215610abd57600080fd5b6000610acb84828501610a03565b91505092915050565b600060208284031215610ae657600080fd5b6000610af484828501610a18565b91505092915050565b600060208284031215610b0f57600080fd5b600082013567ffffffffffffffff811115610b2957600080fd5b610b3584828501610a2d565b91505092915050565b600060208284031215610b5057600080fd5b6000610b5e84828501610a81565b91505092915050565b60008060008060808587031215610b7d57600080fd5b6000610b8b87828801610a81565b9450506020610b9c87828801610a81565b9350506040610bad87828801610a03565b925050606085013567ffffffffffffffff811115610bca57600080fd5b610bd687828801610a2d565b91505092959194509250565b600060208284031215610bf457600080fd5b6000610c0284828501610a96565b91505092915050565b610c1481610fff565b82525050565b610c2381610fff565b82525050565b610c3a610c3582610fff565b6110a0565b82525050565b610c4981611011565b82525050565b610c588161101d565b82525050565b610c678161101d565b82525050565b610c7e610c798261101d565b6110b2565b82525050565b6000610c8f82610fd2565b610c998185610fdd565b9350610ca981856020860161106d565b610cb2816110e0565b840191505092915050565b6000610cca601983610fee565b91507f617373657420686173206265656e206465706f736974656421000000000000006000830152602082019050919050565b6000610d0a601e83610fee565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000610d4a603683610fee565b91507f61737365742068617320616c7265616479206265656e206465706f736974656460008301527f20616e642063616e6e6f74206265206368616e676564000000000000000000006020830152604082019050919050565b60c082016000820151610db96000850182610c0b565b506020820151610dcc6020850182610e1e565b506040820151610ddf6040850182610c4f565b506060820151610df26060850182610c0b565b506080820151610e056080850182610c0b565b5060a0820151610e1860a0850182610c4f565b50505050565b610e2781611047565b82525050565b610e3e610e3982611051565b6110ce565b82525050565b6000610e508285610c29565b601482019150610e608284610e2d565b6001820191508190509392505050565b6000610e7c8284610c6d565b60208201915081905092915050565b6000602082019050610ea06000830184610c1a565b92915050565b6000602082019050610ebb6000830184610c40565b92915050565b6000602082019050610ed66000830184610c5e565b92915050565b60006020820190508181036000830152610ef68184610c84565b905092915050565b60006020820190508181036000830152610f1781610cbd565b9050919050565b60006020820190508181036000830152610f3781610cfd565b9050919050565b60006020820190508181036000830152610f5781610d3d565b9050919050565b600060c082019050610f736000830184610da3565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610f9c57600080fd5b8060405250919050565b600067ffffffffffffffff821115610fbd57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600061100a82611027565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b8381101561108b578082015181840152602081019050611070565b8381111561109a576000848401525b50505050565b60006110ab826110bc565b9050919050565b6000819050919050565b60006110c7826110fe565b9050919050565b60006110d9826110f1565b9050919050565b6000601f19601f8301169050919050565b60008160f81b9050919050565b60008160601b9050919050565b61111481610fff565b811461111f57600080fd5b50565b61112b8161101d565b811461113657600080fd5b50565b61114281611047565b811461114d57600080fd5b50565b61115981611051565b811461116457600080fd5b5056fea2646970667358221220454b9a3618fc05e05777dfa216b17850cc42c21135ee9bbcf85b78ac7ad5cefa64736f6c63430006040033"

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

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 originChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerTransactor) Deposit(opts *bind.TransactOpts, originChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.contract.Transact(opts, "deposit", originChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 originChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerSession) Deposit(originChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.Contract.Deposit(&_CentrifugeAssetHandler.TransactOpts, originChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xcb65d221.
//
// Solidity: function deposit(uint256 originChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerTransactorSession) Deposit(originChainID *big.Int, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
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

var RuntimeBytecode = "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c8063c8ba6c871161005b578063c8ba6c8714610100578063cb65d22114610130578063db95e75c1461014c578063fc9539cd1461017c5761007d565b80630a6d55d814610082578063318c136e146100b25780633cf5040a146100d0575b600080fd5b61009c60048036038101906100979190610ad4565b610198565b6040516100a99190610e8b565b60405180910390f35b6100ba6101cb565b6040516100c79190610e8b565b60405180910390f35b6100ea60048036038101906100e59190610ad4565b6101f0565b6040516100f79190610ea6565b60405180910390f35b61011a60048036038101906101159190610aab565b61021a565b6040516101279190610ec1565b60405180910390f35b61014a60048036038101906101459190610b67565b610232565b005b61016660048036038101906101619190610b3e565b6106ae565b6040516101739190610f5e565b60405180910390f35b61019660048036038101906101919190610afd565b6107fb565b005b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006004600083815260200190815260200160002060009054906101000a900460ff169050919050565b60026020528060005260406000206000915090505481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b890610f1e565b60405180910390fd5b60008060006020840151925060408401519150606084015190506000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016103329190610e70565b60405160208183030381529060405280519060200120826040516020016103599190610e70565b6040516020818303038152906040528051906020012014156104c45760008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663beab71316040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156103e557600080fd5b505af11580156103f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061041d9190610be2565b9050610429878261092a565b935083600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550866001600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505b600015156004600085815260200190815260200160002060009054906101000a900460ff1615151461052b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052290610f3e565b60405180910390fd5b6040518060c001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018a81526020018381526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200184815250600360008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050155905050505050505050505050565b6106b6610985565b600360008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461088a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088190610f1e565b60405180910390fd5b600060208201519050600015156004600083815260200190815260200160002060009054906101000a900460ff161515146108fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f190610efe565b60405180910390fd5b60016004600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600060608383604051602001610941929190610e44565b6040516020818303038152906040526040516020016109609190610edc565b6040516020818303038152906040529050600060208201519050809250505092915050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600080191681525090565b600081359050610a128161110b565b92915050565b600081359050610a2781611122565b92915050565b600082601f830112610a3e57600080fd5b8135610a51610a4c82610fa6565b610f79565b91508082526020830160208301858383011115610a6d57600080fd5b610a7883828461105e565b50505092915050565b600081359050610a9081611139565b92915050565b600081519050610aa581611150565b92915050565b600060208284031215610abd57600080fd5b6000610acb84828501610a03565b91505092915050565b600060208284031215610ae657600080fd5b6000610af484828501610a18565b91505092915050565b600060208284031215610b0f57600080fd5b600082013567ffffffffffffffff811115610b2957600080fd5b610b3584828501610a2d565b91505092915050565b600060208284031215610b5057600080fd5b6000610b5e84828501610a81565b91505092915050565b60008060008060808587031215610b7d57600080fd5b6000610b8b87828801610a81565b9450506020610b9c87828801610a81565b9350506040610bad87828801610a03565b925050606085013567ffffffffffffffff811115610bca57600080fd5b610bd687828801610a2d565b91505092959194509250565b600060208284031215610bf457600080fd5b6000610c0284828501610a96565b91505092915050565b610c1481610fff565b82525050565b610c2381610fff565b82525050565b610c3a610c3582610fff565b6110a0565b82525050565b610c4981611011565b82525050565b610c588161101d565b82525050565b610c678161101d565b82525050565b610c7e610c798261101d565b6110b2565b82525050565b6000610c8f82610fd2565b610c998185610fdd565b9350610ca981856020860161106d565b610cb2816110e0565b840191505092915050565b6000610cca601983610fee565b91507f617373657420686173206265656e206465706f736974656421000000000000006000830152602082019050919050565b6000610d0a601e83610fee565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000610d4a603683610fee565b91507f61737365742068617320616c7265616479206265656e206465706f736974656460008301527f20616e642063616e6e6f74206265206368616e676564000000000000000000006020830152604082019050919050565b60c082016000820151610db96000850182610c0b565b506020820151610dcc6020850182610e1e565b506040820151610ddf6040850182610c4f565b506060820151610df26060850182610c0b565b506080820151610e056080850182610c0b565b5060a0820151610e1860a0850182610c4f565b50505050565b610e2781611047565b82525050565b610e3e610e3982611051565b6110ce565b82525050565b6000610e508285610c29565b601482019150610e608284610e2d565b6001820191508190509392505050565b6000610e7c8284610c6d565b60208201915081905092915050565b6000602082019050610ea06000830184610c1a565b92915050565b6000602082019050610ebb6000830184610c40565b92915050565b6000602082019050610ed66000830184610c5e565b92915050565b60006020820190508181036000830152610ef68184610c84565b905092915050565b60006020820190508181036000830152610f1781610cbd565b9050919050565b60006020820190508181036000830152610f3781610cfd565b9050919050565b60006020820190508181036000830152610f5781610d3d565b9050919050565b600060c082019050610f736000830184610da3565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610f9c57600080fd5b8060405250919050565b600067ffffffffffffffff821115610fbd57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600061100a82611027565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b8381101561108b578082015181840152602081019050611070565b8381111561109a576000848401525b50505050565b60006110ab826110bc565b9050919050565b6000819050919050565b60006110c7826110fe565b9050919050565b60006110d9826110f1565b9050919050565b6000601f19601f8301169050919050565b60008160f81b9050919050565b60008160601b9050919050565b61111481610fff565b811461111f57600080fd5b50565b61112b8161101d565b811461113657600080fd5b50565b61114281611047565b811461114d57600080fd5b50565b61115981611051565b811461116457600080fd5b5056fea2646970667358221220454b9a3618fc05e05777dfa216b17850cc42c21135ee9bbcf85b78ac7ad5cefa64736f6c63430006040033"
