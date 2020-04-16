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
const CentrifugeAssetHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResourceIDAndContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"originChainContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metaDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structCentrifugeAssetHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CentrifugeAssetHandlerBin is the compiled bytecode used for deploying new contracts.
var CentrifugeAssetHandlerBin = "0x60806040523480156200001157600080fd5b5060405162001a4738038062001a4783398181016040528101906200003791906200035a565b80518251146200007e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000759062000470565b60405180910390fd5b826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b8251811015620001165762000108838281518110620000df57fe5b6020026020010151838381518110620000f457fe5b60200260200101516200012060201b60201c565b8080600101915050620000c4565b5050505062000595565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600081519050620002238162000561565b92915050565b600082601f8301126200023b57600080fd5b8151620002526200024c82620004c0565b62000492565b915081818352602084019350602081019050838560208402820111156200027857600080fd5b60005b83811015620002ac578162000291888262000212565b8452602084019350602083019250506001810190506200027b565b5050505092915050565b600082601f830112620002c857600080fd5b8151620002df620002d982620004e9565b62000492565b915081818352602084019350602081019050838560208402820111156200030557600080fd5b60005b838110156200033957816200031e888262000343565b84526020840193506020830192505060018101905062000308565b5050505092915050565b60008151905062000354816200057b565b92915050565b6000806000606084860312156200037057600080fd5b6000620003808682870162000212565b935050602084015167ffffffffffffffff8111156200039e57600080fd5b620003ac86828701620002b6565b925050604084015167ffffffffffffffff811115620003ca57600080fd5b620003d88682870162000229565b9150509250925092565b6000620003f160478362000512565b91507f6d69736d61746368206c656e677468206265747765656e20696e697469616c5260008301527f65736f7572636549447320616e6420696e697469616c436f6e7472616374416460208301527f64726573736573000000000000000000000000000000000000000000000000006040830152606082019050919050565b600060208201905081810360008301526200048b81620003e2565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715620004b657600080fd5b8060405250919050565b600067ffffffffffffffff821115620004d857600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200050157600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620005308262000541565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200056c8162000523565b81146200057857600080fd5b50565b620005868162000537565b81146200059257600080fd5b50565b6114a280620005a56000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80637f79bea8116100665780637f79bea8146101325780638a025ce814610162578063c8ba6c871461017e578063db95e75c146101ae578063fc9539cd146101de57610093565b80630a6d55d814610098578063318c136e146100c85780633cf5040a146100e657806345a104db14610116575b600080fd5b6100b260048036038101906100ad9190610d13565b6101fa565b6040516100bf91906111e2565b60405180910390f35b6100d061022d565b6040516100dd91906111e2565b60405180910390f35b61010060048036038101906100fb9190610d13565b610252565b60405161010d91906111fd565b60405180910390f35b610130600480360381019061012b9190610de2565b61027c565b005b61014c60048036038101906101479190610cea565b6105a8565b60405161015991906111fd565b60405180910390f35b61017c60048036038101906101779190610d3c565b6105c8565b005b61019860048036038101906101939190610cea565b61074e565b6040516101a59190611218565b60405180910390f35b6101c860048036038101906101c39190610db9565b610766565b6040516101d59190611313565b60405180910390f35b6101f860048036038101906101f39190610d78565b6108c6565b005b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006004600083815260200190815260200160002060009054906101000a900460ff169050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461030b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030290611273565b60405180910390fd5b600080600060208401519250604084015191506060840151905060006001600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905061036681610a79565b6103a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039c90611293565b60405180910390fd5b600015156004600084815260200190815260200160002060009054906101000a900460ff1615151461040c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610403906112f3565b60405180910390fd5b6040518060c001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018960ff1681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff168152602001838152506003600089815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff1602179055506040820151816001015560608201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a082015181600401559050505050505050505050565b60056020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166001600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461066a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610661906112b3565b60405180910390fd5b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016106c191906111c7565b60405160208183030381529060405280519060200120826040516020016106e891906111c7565b604051602081830303815290604052805190602001201461073e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073590611253565b60405180910390fd5b6107488484610acf565b50505050565b60026020528060005260406000206000915090505481565b61076e610bc1565b600360008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820154815250509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610955576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094c90611273565b60405180910390fd5b60008060208301519150604083015190506109a26001600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610a79565b6109e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d8906112d3565b60405180910390fd5b600015156004600083815260200190815260200160002060009054906101000a900460ff16151514610a48576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3f90611233565b60405180910390fd5b60016004600083815260200190815260200160002060006101000a81548160ff021916908315150217905550505050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff16815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600080191681525090565b600081359050610c5181611410565b92915050565b600081359050610c6681611427565b92915050565b600082601f830112610c7d57600080fd5b8135610c90610c8b8261135b565b61132e565b91508082526020830160208301858383011115610cac57600080fd5b610cb78382846113f7565b50505092915050565b600081359050610ccf8161143e565b92915050565b600081359050610ce481611455565b92915050565b600060208284031215610cfc57600080fd5b6000610d0a84828501610c42565b91505092915050565b600060208284031215610d2557600080fd5b6000610d3384828501610c57565b91505092915050565b60008060408385031215610d4f57600080fd5b6000610d5d85828601610c57565b9250506020610d6e85828601610c42565b9150509250929050565b600060208284031215610d8a57600080fd5b600082013567ffffffffffffffff811115610da457600080fd5b610db084828501610c6c565b91505092915050565b600060208284031215610dcb57600080fd5b6000610dd984828501610cc0565b91505092915050565b60008060008060808587031215610df857600080fd5b6000610e0687828801610cd5565b9450506020610e1787828801610cc0565b9350506040610e2887828801610c42565b925050606085013567ffffffffffffffff811115610e4557600080fd5b610e5187828801610c6c565b91505092959194509250565b610e6681611398565b82525050565b610e7581611398565b82525050565b610e84816113aa565b82525050565b610e93816113b6565b82525050565b610ea2816113b6565b82525050565b610eb9610eb4826113b6565b611406565b82525050565b6000610ecc601983611387565b91507f617373657420686173206265656e206465706f736974656421000000000000006000830152602082019050919050565b6000610f0c603583611387565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000610f72601e83611387565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000610fb2603683611387565b91507f70726f7669646564206f726967696e436861696e436f6e74726163744164647260008301527f657373206973206e6f742077686974656c6973746564000000000000000000006020830152604082019050919050565b6000611018603783611387565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b600061107e602883611387565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b60006110e4603683611387565b91507f61737365742068617320616c7265616479206265656e206465706f736974656460008301527f20616e642063616e6e6f74206265206368616e676564000000000000000000006020830152604082019050919050565b60c0820160008201516111536000850182610e5d565b50602082015161116660208501826111b8565b5060408201516111796040850182610e8a565b50606082015161118c6060850182610e5d565b50608082015161119f6080850182610e5d565b5060a08201516111b260a0850182610e8a565b50505050565b6111c1816113ea565b82525050565b60006111d38284610ea8565b60208201915081905092915050565b60006020820190506111f76000830184610e6c565b92915050565b60006020820190506112126000830184610e7b565b92915050565b600060208201905061122d6000830184610e99565b92915050565b6000602082019050818103600083015261124c81610ebf565b9050919050565b6000602082019050818103600083015261126c81610eff565b9050919050565b6000602082019050818103600083015261128c81610f65565b9050919050565b600060208201905081810360008301526112ac81610fa5565b9050919050565b600060208201905081810360008301526112cc8161100b565b9050919050565b600060208201905081810360008301526112ec81611071565b9050919050565b6000602082019050818103600083015261130c816110d7565b9050919050565b600060c082019050611328600083018461113d565b92915050565b6000604051905081810181811067ffffffffffffffff8211171561135157600080fd5b8060405250919050565b600067ffffffffffffffff82111561137257600080fd5b601f19601f8301169050602081019050919050565b600082825260208201905092915050565b60006113a3826113c0565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b6000819050919050565b61141981611398565b811461142457600080fd5b50565b611430816113b6565b811461143b57600080fd5b50565b611447816113e0565b811461145257600080fd5b50565b61145e816113ea565b811461146957600080fd5b5056fea2646970667358221220e6e2e5bbee8eb21c7367b5bf89d9a088a51799ba586bc3b0410afb0f76e2fab764736f6c63430006040033"

// DeployCentrifugeAssetHandler deploys a new Ethereum contract, binding an instance of CentrifugeAssetHandler to it.
func DeployCentrifugeAssetHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address) (common.Address, *types.Transaction, *CentrifugeAssetHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(CentrifugeAssetHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CentrifugeAssetHandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses)
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

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) constant returns(bool)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CentrifugeAssetHandler.contract.Call(opts, out, "_contractWhitelist", arg0)
	return *ret0, err
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) constant returns(bool)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _CentrifugeAssetHandler.Contract.ContractWhitelist(&_CentrifugeAssetHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) constant returns(bool)
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _CentrifugeAssetHandler.Contract.ContractWhitelist(&_CentrifugeAssetHandler.CallOpts, arg0)
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

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerTransactor) SetResourceIDAndContractAddress(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.contract.Transact(opts, "setResourceIDAndContractAddress", resourceID, contractAddress)
}

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerSession) SetResourceIDAndContractAddress(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.Contract.SetResourceIDAndContractAddress(&_CentrifugeAssetHandler.TransactOpts, resourceID, contractAddress)
}

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_CentrifugeAssetHandler *CentrifugeAssetHandlerTransactorSession) SetResourceIDAndContractAddress(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _CentrifugeAssetHandler.Contract.SetResourceIDAndContractAddress(&_CentrifugeAssetHandler.TransactOpts, resourceID, contractAddress)
}

var RuntimeBytecode = "0x608060405234801561001057600080fd5b50600436106100935760003560e01c80637f79bea8116100665780637f79bea8146101325780638a025ce814610162578063c8ba6c871461017e578063db95e75c146101ae578063fc9539cd146101de57610093565b80630a6d55d814610098578063318c136e146100c85780633cf5040a146100e657806345a104db14610116575b600080fd5b6100b260048036038101906100ad9190610d13565b6101fa565b6040516100bf91906111e2565b60405180910390f35b6100d061022d565b6040516100dd91906111e2565b60405180910390f35b61010060048036038101906100fb9190610d13565b610252565b60405161010d91906111fd565b60405180910390f35b610130600480360381019061012b9190610de2565b61027c565b005b61014c60048036038101906101479190610cea565b6105a8565b60405161015991906111fd565b60405180910390f35b61017c60048036038101906101779190610d3c565b6105c8565b005b61019860048036038101906101939190610cea565b61074e565b6040516101a59190611218565b60405180910390f35b6101c860048036038101906101c39190610db9565b610766565b6040516101d59190611313565b60405180910390f35b6101f860048036038101906101f39190610d78565b6108c6565b005b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006004600083815260200190815260200160002060009054906101000a900460ff169050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461030b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030290611273565b60405180910390fd5b600080600060208401519250604084015191506060840151905060006001600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905061036681610a79565b6103a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039c90611293565b60405180910390fd5b600015156004600084815260200190815260200160002060009054906101000a900460ff1615151461040c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610403906112f3565b60405180910390fd5b6040518060c001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018960ff1681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff168152602001838152506003600089815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff1602179055506040820151816001015560608201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a082015181600401559050505050505050505050565b60056020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166001600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461066a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610661906112b3565b60405180910390fd5b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016106c191906111c7565b60405160208183030381529060405280519060200120826040516020016106e891906111c7565b604051602081830303815290604052805190602001201461073e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073590611253565b60405180910390fd5b6107488484610acf565b50505050565b60026020528060005260406000206000915090505481565b61076e610bc1565b600360008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820154815250509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610955576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094c90611273565b60405180910390fd5b60008060208301519150604083015190506109a26001600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610a79565b6109e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d8906112d3565b60405180910390fd5b600015156004600083815260200190815260200160002060009054906101000a900460ff16151514610a48576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3f90611233565b60405180910390fd5b60016004600083815260200190815260200160002060006101000a81548160ff021916908315150217905550505050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff16815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600080191681525090565b600081359050610c5181611410565b92915050565b600081359050610c6681611427565b92915050565b600082601f830112610c7d57600080fd5b8135610c90610c8b8261135b565b61132e565b91508082526020830160208301858383011115610cac57600080fd5b610cb78382846113f7565b50505092915050565b600081359050610ccf8161143e565b92915050565b600081359050610ce481611455565b92915050565b600060208284031215610cfc57600080fd5b6000610d0a84828501610c42565b91505092915050565b600060208284031215610d2557600080fd5b6000610d3384828501610c57565b91505092915050565b60008060408385031215610d4f57600080fd5b6000610d5d85828601610c57565b9250506020610d6e85828601610c42565b9150509250929050565b600060208284031215610d8a57600080fd5b600082013567ffffffffffffffff811115610da457600080fd5b610db084828501610c6c565b91505092915050565b600060208284031215610dcb57600080fd5b6000610dd984828501610cc0565b91505092915050565b60008060008060808587031215610df857600080fd5b6000610e0687828801610cd5565b9450506020610e1787828801610cc0565b9350506040610e2887828801610c42565b925050606085013567ffffffffffffffff811115610e4557600080fd5b610e5187828801610c6c565b91505092959194509250565b610e6681611398565b82525050565b610e7581611398565b82525050565b610e84816113aa565b82525050565b610e93816113b6565b82525050565b610ea2816113b6565b82525050565b610eb9610eb4826113b6565b611406565b82525050565b6000610ecc601983611387565b91507f617373657420686173206265656e206465706f736974656421000000000000006000830152602082019050919050565b6000610f0c603583611387565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000610f72601e83611387565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000610fb2603683611387565b91507f70726f7669646564206f726967696e436861696e436f6e74726163744164647260008301527f657373206973206e6f742077686974656c6973746564000000000000000000006020830152604082019050919050565b6000611018603783611387565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b600061107e602883611387565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b60006110e4603683611387565b91507f61737365742068617320616c7265616479206265656e206465706f736974656460008301527f20616e642063616e6e6f74206265206368616e676564000000000000000000006020830152604082019050919050565b60c0820160008201516111536000850182610e5d565b50602082015161116660208501826111b8565b5060408201516111796040850182610e8a565b50606082015161118c6060850182610e5d565b50608082015161119f6080850182610e5d565b5060a08201516111b260a0850182610e8a565b50505050565b6111c1816113ea565b82525050565b60006111d38284610ea8565b60208201915081905092915050565b60006020820190506111f76000830184610e6c565b92915050565b60006020820190506112126000830184610e7b565b92915050565b600060208201905061122d6000830184610e99565b92915050565b6000602082019050818103600083015261124c81610ebf565b9050919050565b6000602082019050818103600083015261126c81610eff565b9050919050565b6000602082019050818103600083015261128c81610f65565b9050919050565b600060208201905081810360008301526112ac81610fa5565b9050919050565b600060208201905081810360008301526112cc8161100b565b9050919050565b600060208201905081810360008301526112ec81611071565b9050919050565b6000602082019050818103600083015261130c816110d7565b9050919050565b600060c082019050611328600083018461113d565b92915050565b6000604051905081810181811067ffffffffffffffff8211171561135157600080fd5b8060405250919050565b600067ffffffffffffffff82111561137257600080fd5b601f19601f8301169050602081019050919050565b600082825260208201905092915050565b60006113a3826113c0565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b6000819050919050565b61141981611398565b811461142457600080fd5b50565b611430816113b6565b811461143b57600080fd5b50565b611447816113e0565b811461145257600080fd5b50565b61145e816113ea565b811461146957600080fd5b5056fea2646970667358221220e6e2e5bbee8eb21c7367b5bf89d9a088a51799ba586bc3b0410afb0f76e2fab764736f6c63430006040033"
