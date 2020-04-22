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
const GenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResourceIDAndContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structGenericHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GenericHandlerBin is the compiled bytecode used for deploying new contracts.
var GenericHandlerBin = "0x60806040523480156200001157600080fd5b5060405162001c7e38038062001c7e83398181016040528101906200003791906200035b565b80518251146200007e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000759062000471565b60405180910390fd5b82600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b8251811015620001175762000109838281518110620000e057fe5b6020026020010151838381518110620000f557fe5b60200260200101516200012160201b60201c565b8080600101915050620000c5565b5050505062000596565b806004600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600081519050620002248162000562565b92915050565b600082601f8301126200023c57600080fd5b8151620002536200024d82620004c1565b62000493565b915081818352602084019350602081019050838560208402820111156200027957600080fd5b60005b83811015620002ad578162000292888262000213565b8452602084019350602083019250506001810190506200027c565b5050505092915050565b600082601f830112620002c957600080fd5b8151620002e0620002da82620004ea565b62000493565b915081818352602084019350602081019050838560208402820111156200030657600080fd5b60005b838110156200033a57816200031f888262000344565b84526020840193506020830192505060018101905062000309565b5050505092915050565b60008151905062000355816200057c565b92915050565b6000806000606084860312156200037157600080fd5b6000620003818682870162000213565b935050602084015167ffffffffffffffff8111156200039f57600080fd5b620003ad86828701620002b7565b925050604084015167ffffffffffffffff811115620003cb57600080fd5b620003d9868287016200022a565b9150509250925092565b6000620003f260478362000513565b91507f6d69736d61746368206c656e677468206265747765656e20696e697469616c5260008301527f65736f7572636549447320616e6420696e697469616c436f6e7472616374416460208301527f64726573736573000000000000000000000000000000000000000000000000006040830152606082019050919050565b600060208201905081810360008301526200048c81620003e3565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715620004b757600080fd5b8060405250919050565b600067ffffffffffffffff821115620004d957600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200050257600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620005318262000542565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200056d8162000524565b81146200057957600080fd5b50565b620005878162000538565b81146200059357600080fd5b50565b6116d880620005a66000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638a025ce8116100715780638a025ce8146101b357806395601f09146101cf578063c8ba6c87146101eb578063db95e75c1461021b578063e245386f1461024b578063fc9539cd1461027f576100b4565b80630a6d55d8146100b9578063318c136e146100e9578063452747381461010757806345a104db146101375780636ebcf607146101535780637f79bea814610183575b600080fd5b6100d360048036038101906100ce9190610f15565b61029b565b6040516100e0919061133d565b60405180910390f35b6100f16102ce565b6040516100fe919061133d565b60405180910390f35b610121600480360381019061011c9190610e74565b6102f4565b60405161012e9190611467565b60405180910390f35b610151600480360381019061014c9190610fe4565b61030c565b005b61016d60048036038101906101689190610e74565b610516565b60405161017a9190611467565b60405180910390f35b61019d60048036038101906101989190610e74565b61052e565b6040516101aa919061138f565b60405180910390f35b6101cd60048036038101906101c89190610f3e565b61054e565b005b6101e960048036038101906101e49190610e9d565b6106d4565b005b61020560048036038101906102009190610e74565b610802565b60405161021291906113aa565b60405180910390f35b61023560048036038101906102309190610fbb565b61081a565b6040516102429190611445565b60405180910390f35b61026560048036038101906102609190610fbb565b6109bc565b604051610276959493929190611482565b60405180910390f35b61029960048036038101906102949190610f7a565b610ad7565b005b60046020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461039c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610393906113e5565b60405180910390fd5b600080606060208401519250604084015191506040519050836060015180820160600160405260e4360360e48337506040518060a001604052808860ff1681526020018381526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001828152506003600088815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004019080519060200190610509929190610cb1565b5090505050505050505050565b60006020528060005260406000206000915090505481565b60076020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166004600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146105f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e790611425565b60405180910390fd5b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016106479190611322565b604051602081830303815290604052805190602001208260405160200161066e9190611322565b60405160208183030381529060405280519060200120146106c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106bb906113c5565b60405180910390fd5b6106ce8484610b6a565b50505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b815260040161071693929190611358565b602060405180830381600087803b15801561073057600080fd5b505af1158015610744573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107689190610eec565b506107ba826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c5c90919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60056020528060005260406000206000915090505481565b610822610d31565b600360008381526020019081526020016000206040518060a00160405290816000820160009054906101000a900460ff1660ff1660ff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109ac5780601f10610981576101008083540402835291602001916109ac565b820191906000526020600020905b81548152906001019060200180831161098f57829003601f168201915b5050505050815250509050919050565b60036020528060005260406000206000915090508060000160009054906101000a900460ff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806004018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610acd5780601f10610aa257610100808354040283529160200191610acd565b820191906000526020600020905b815481529060010190602001808311610ab057829003601f168201915b5050505050905085565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b67576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5e906113e5565b60405180910390fd5b50565b806004600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600080828401905083811015610ca7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9e90611405565b60405180910390fd5b8091505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610cf257805160ff1916838001178555610d20565b82800160010185558215610d20579182015b82811115610d1f578251825591602001919060010190610d04565b5b509050610d2d9190610d92565b5090565b6040518060a00160405280600060ff16815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b610db491905b80821115610db0576000816000905550600101610d98565b5090565b90565b600081359050610dc68161162f565b92915050565b600081519050610ddb81611646565b92915050565b600081359050610df08161165d565b92915050565b600082601f830112610e0757600080fd5b8135610e1a610e1582611509565b6114dc565b91508082526020830160208301858383011115610e3657600080fd5b610e418382846115d2565b50505092915050565b600081359050610e5981611674565b92915050565b600081359050610e6e8161168b565b92915050565b600060208284031215610e8657600080fd5b6000610e9484828501610db7565b91505092915050565b600080600060608486031215610eb257600080fd5b6000610ec086828701610db7565b9350506020610ed186828701610db7565b9250506040610ee286828701610e4a565b9150509250925092565b600060208284031215610efe57600080fd5b6000610f0c84828501610dcc565b91505092915050565b600060208284031215610f2757600080fd5b6000610f3584828501610de1565b91505092915050565b60008060408385031215610f5157600080fd5b6000610f5f85828601610de1565b9250506020610f7085828601610db7565b9150509250929050565b600060208284031215610f8c57600080fd5b600082013567ffffffffffffffff811115610fa657600080fd5b610fb284828501610df6565b91505092915050565b600060208284031215610fcd57600080fd5b6000610fdb84828501610e4a565b91505092915050565b60008060008060808587031215610ffa57600080fd5b600061100887828801610e5f565b945050602061101987828801610e4a565b935050604061102a87828801610db7565b925050606085013567ffffffffffffffff81111561104757600080fd5b61105387828801610df6565b91505092959194509250565b61106881611573565b82525050565b61107781611573565b82525050565b61108681611585565b82525050565b61109581611591565b82525050565b6110a481611591565b82525050565b6110bb6110b682611591565b611614565b82525050565b60006110cc82611535565b6110d68185611540565b93506110e68185602086016115e1565b6110ef8161161e565b840191505092915050565b600061110582611535565b61110f8185611551565b935061111f8185602086016115e1565b6111288161161e565b840191505092915050565b6000611140603583611562565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b60006111a6601e83611562565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b60006111e6601b83611562565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000611226603783611562565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b600060a0830160008301516112976000860182611304565b5060208301516112aa602086018261108c565b5060408301516112bd604086018261105f565b5060608301516112d0606086018261105f565b50608083015184820360808601526112e882826110c1565b9150508091505092915050565b6112fe816115bb565b82525050565b61130d816115c5565b82525050565b61131c816115c5565b82525050565b600061132e82846110aa565b60208201915081905092915050565b6000602082019050611352600083018461106e565b92915050565b600060608201905061136d600083018661106e565b61137a602083018561106e565b61138760408301846112f5565b949350505050565b60006020820190506113a4600083018461107d565b92915050565b60006020820190506113bf600083018461109b565b92915050565b600060208201905081810360008301526113de81611133565b9050919050565b600060208201905081810360008301526113fe81611199565b9050919050565b6000602082019050818103600083015261141e816111d9565b9050919050565b6000602082019050818103600083015261143e81611219565b9050919050565b6000602082019050818103600083015261145f818461127f565b905092915050565b600060208201905061147c60008301846112f5565b92915050565b600060a0820190506114976000830188611313565b6114a4602083018761109b565b6114b1604083018661106e565b6114be606083018561106e565b81810360808301526114d081846110fa565b90509695505050505050565b6000604051905081810181811067ffffffffffffffff821117156114ff57600080fd5b8060405250919050565b600067ffffffffffffffff82111561152057600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061157e8261159b565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156115ff5780820151818401526020810190506115e4565b8381111561160e576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b61163881611573565b811461164357600080fd5b50565b61164f81611585565b811461165a57600080fd5b50565b61166681611591565b811461167157600080fd5b50565b61167d816115bb565b811461168857600080fd5b50565b611694816115c5565b811461169f57600080fd5b5056fea26469706673582212204232fe4a8cab608274b46f34042bd5ae7f0a90e9115e9441e9ff7edb02af4e1764736f6c63430006040033"

// DeployGenericHandler deploys a new Ethereum contract, binding an instance of GenericHandler to it.
func DeployGenericHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address) (common.Address, *types.Transaction, *GenericHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GenericHandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses)
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
// Solidity: function _balances(address ) view returns(uint256)
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
// Solidity: function _balances(address ) view returns(uint256)
func (_GenericHandler *GenericHandlerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GenericHandler.Contract.Balances(&_GenericHandler.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) view returns(uint256)
func (_GenericHandler *GenericHandlerCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GenericHandler.Contract.Balances(&_GenericHandler.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
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
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerSession) BridgeAddress() (common.Address, error) {
	return _GenericHandler.Contract.BridgeAddress(&_GenericHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _GenericHandler.Contract.BridgeAddress(&_GenericHandler.CallOpts)
}

// BurnedTokens is a free data retrieval call binding the contract method 0x45274738.
//
// Solidity: function _burnedTokens(address ) view returns(uint256)
func (_GenericHandler *GenericHandlerCaller) BurnedTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenericHandler.contract.Call(opts, out, "_burnedTokens", arg0)
	return *ret0, err
}

// BurnedTokens is a free data retrieval call binding the contract method 0x45274738.
//
// Solidity: function _burnedTokens(address ) view returns(uint256)
func (_GenericHandler *GenericHandlerSession) BurnedTokens(arg0 common.Address) (*big.Int, error) {
	return _GenericHandler.Contract.BurnedTokens(&_GenericHandler.CallOpts, arg0)
}

// BurnedTokens is a free data retrieval call binding the contract method 0x45274738.
//
// Solidity: function _burnedTokens(address ) view returns(uint256)
func (_GenericHandler *GenericHandlerCallerSession) BurnedTokens(arg0 common.Address) (*big.Int, error) {
	return _GenericHandler.Contract.BurnedTokens(&_GenericHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GenericHandler.contract.Call(opts, out, "_contractWhitelist", arg0)
	return *ret0, err
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _GenericHandler.Contract.ContractWhitelist(&_GenericHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _GenericHandler.Contract.ContractWhitelist(&_GenericHandler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) view returns(uint8 _destinationChainID, bytes32 _resourceID, address _destinationRecipientAddress, address _depositer, bytes _metaData)
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
// Solidity: function _depositRecords(uint256 ) view returns(uint8 _destinationChainID, bytes32 _resourceID, address _destinationRecipientAddress, address _depositer, bytes _metaData)
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
// Solidity: function _depositRecords(uint256 ) view returns(uint8 _destinationChainID, bytes32 _resourceID, address _destinationRecipientAddress, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerCallerSession) DepositRecords(arg0 *big.Int) (struct {
	DestinationChainID          uint8
	ResourceID                  [32]byte
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaData                    []byte
}, error) {
	return _GenericHandler.Contract.DepositRecords(&_GenericHandler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GenericHandler.contract.Call(opts, out, "_resourceIDToTokenContractAddress", arg0)
	return *ret0, err
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToTokenContractAddress(&_GenericHandler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToTokenContractAddress(&_GenericHandler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _GenericHandler.contract.Call(opts, out, "_tokenContractAddressToResourceID", arg0)
	return *ret0, err
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.TokenContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.TokenContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositNonce) view returns(GenericHandlerDepositRecord)
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
// Solidity: function getDepositRecord(uint256 depositNonce) view returns(GenericHandlerDepositRecord)
func (_GenericHandler *GenericHandlerSession) GetDepositRecord(depositNonce *big.Int) (GenericHandlerDepositRecord, error) {
	return _GenericHandler.Contract.GetDepositRecord(&_GenericHandler.CallOpts, depositNonce)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositNonce) view returns(GenericHandlerDepositRecord)
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

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_GenericHandler *GenericHandlerTransactor) SetResourceIDAndContractAddress(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "setResourceIDAndContractAddress", resourceID, contractAddress)
}

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_GenericHandler *GenericHandlerSession) SetResourceIDAndContractAddress(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetResourceIDAndContractAddress(&_GenericHandler.TransactOpts, resourceID, contractAddress)
}

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_GenericHandler *GenericHandlerTransactorSession) SetResourceIDAndContractAddress(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetResourceIDAndContractAddress(&_GenericHandler.TransactOpts, resourceID, contractAddress)
}

var RuntimeBytecode = "0x608060405234801561001057600080fd5b50600436106100b45760003560e01c80638a025ce8116100715780638a025ce8146101b357806395601f09146101cf578063c8ba6c87146101eb578063db95e75c1461021b578063e245386f1461024b578063fc9539cd1461027f576100b4565b80630a6d55d8146100b9578063318c136e146100e9578063452747381461010757806345a104db146101375780636ebcf607146101535780637f79bea814610183575b600080fd5b6100d360048036038101906100ce9190610f15565b61029b565b6040516100e0919061133d565b60405180910390f35b6100f16102ce565b6040516100fe919061133d565b60405180910390f35b610121600480360381019061011c9190610e74565b6102f4565b60405161012e9190611467565b60405180910390f35b610151600480360381019061014c9190610fe4565b61030c565b005b61016d60048036038101906101689190610e74565b610516565b60405161017a9190611467565b60405180910390f35b61019d60048036038101906101989190610e74565b61052e565b6040516101aa919061138f565b60405180910390f35b6101cd60048036038101906101c89190610f3e565b61054e565b005b6101e960048036038101906101e49190610e9d565b6106d4565b005b61020560048036038101906102009190610e74565b610802565b60405161021291906113aa565b60405180910390f35b61023560048036038101906102309190610fbb565b61081a565b6040516102429190611445565b60405180910390f35b61026560048036038101906102609190610fbb565b6109bc565b604051610276959493929190611482565b60405180910390f35b61029960048036038101906102949190610f7a565b610ad7565b005b60046020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461039c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610393906113e5565b60405180910390fd5b600080606060208401519250604084015191506040519050836060015180820160600160405260e4360360e48337506040518060a001604052808860ff1681526020018381526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001828152506003600088815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004019080519060200190610509929190610cb1565b5090505050505050505050565b60006020528060005260406000206000915090505481565b60076020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166004600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146105f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e790611425565b60405180910390fd5b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016106479190611322565b604051602081830303815290604052805190602001208260405160200161066e9190611322565b60405160208183030381529060405280519060200120146106c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106bb906113c5565b60405180910390fd5b6106ce8484610b6a565b50505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b815260040161071693929190611358565b602060405180830381600087803b15801561073057600080fd5b505af1158015610744573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107689190610eec565b506107ba826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c5c90919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60056020528060005260406000206000915090505481565b610822610d31565b600360008381526020019081526020016000206040518060a00160405290816000820160009054906101000a900460ff1660ff1660ff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109ac5780601f10610981576101008083540402835291602001916109ac565b820191906000526020600020905b81548152906001019060200180831161098f57829003601f168201915b5050505050815250509050919050565b60036020528060005260406000206000915090508060000160009054906101000a900460ff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806004018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610acd5780601f10610aa257610100808354040283529160200191610acd565b820191906000526020600020905b815481529060010190602001808311610ab057829003601f168201915b5050505050905085565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b67576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5e906113e5565b60405180910390fd5b50565b806004600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600080828401905083811015610ca7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9e90611405565b60405180910390fd5b8091505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610cf257805160ff1916838001178555610d20565b82800160010185558215610d20579182015b82811115610d1f578251825591602001919060010190610d04565b5b509050610d2d9190610d92565b5090565b6040518060a00160405280600060ff16815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b610db491905b80821115610db0576000816000905550600101610d98565b5090565b90565b600081359050610dc68161162f565b92915050565b600081519050610ddb81611646565b92915050565b600081359050610df08161165d565b92915050565b600082601f830112610e0757600080fd5b8135610e1a610e1582611509565b6114dc565b91508082526020830160208301858383011115610e3657600080fd5b610e418382846115d2565b50505092915050565b600081359050610e5981611674565b92915050565b600081359050610e6e8161168b565b92915050565b600060208284031215610e8657600080fd5b6000610e9484828501610db7565b91505092915050565b600080600060608486031215610eb257600080fd5b6000610ec086828701610db7565b9350506020610ed186828701610db7565b9250506040610ee286828701610e4a565b9150509250925092565b600060208284031215610efe57600080fd5b6000610f0c84828501610dcc565b91505092915050565b600060208284031215610f2757600080fd5b6000610f3584828501610de1565b91505092915050565b60008060408385031215610f5157600080fd5b6000610f5f85828601610de1565b9250506020610f7085828601610db7565b9150509250929050565b600060208284031215610f8c57600080fd5b600082013567ffffffffffffffff811115610fa657600080fd5b610fb284828501610df6565b91505092915050565b600060208284031215610fcd57600080fd5b6000610fdb84828501610e4a565b91505092915050565b60008060008060808587031215610ffa57600080fd5b600061100887828801610e5f565b945050602061101987828801610e4a565b935050604061102a87828801610db7565b925050606085013567ffffffffffffffff81111561104757600080fd5b61105387828801610df6565b91505092959194509250565b61106881611573565b82525050565b61107781611573565b82525050565b61108681611585565b82525050565b61109581611591565b82525050565b6110a481611591565b82525050565b6110bb6110b682611591565b611614565b82525050565b60006110cc82611535565b6110d68185611540565b93506110e68185602086016115e1565b6110ef8161161e565b840191505092915050565b600061110582611535565b61110f8185611551565b935061111f8185602086016115e1565b6111288161161e565b840191505092915050565b6000611140603583611562565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b60006111a6601e83611562565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b60006111e6601b83611562565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000611226603783611562565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b600060a0830160008301516112976000860182611304565b5060208301516112aa602086018261108c565b5060408301516112bd604086018261105f565b5060608301516112d0606086018261105f565b50608083015184820360808601526112e882826110c1565b9150508091505092915050565b6112fe816115bb565b82525050565b61130d816115c5565b82525050565b61131c816115c5565b82525050565b600061132e82846110aa565b60208201915081905092915050565b6000602082019050611352600083018461106e565b92915050565b600060608201905061136d600083018661106e565b61137a602083018561106e565b61138760408301846112f5565b949350505050565b60006020820190506113a4600083018461107d565b92915050565b60006020820190506113bf600083018461109b565b92915050565b600060208201905081810360008301526113de81611133565b9050919050565b600060208201905081810360008301526113fe81611199565b9050919050565b6000602082019050818103600083015261141e816111d9565b9050919050565b6000602082019050818103600083015261143e81611219565b9050919050565b6000602082019050818103600083015261145f818461127f565b905092915050565b600060208201905061147c60008301846112f5565b92915050565b600060a0820190506114976000830188611313565b6114a4602083018761109b565b6114b1604083018661106e565b6114be606083018561106e565b81810360808301526114d081846110fa565b90509695505050505050565b6000604051905081810181811067ffffffffffffffff821117156114ff57600080fd5b8060405250919050565b600067ffffffffffffffff82111561152057600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061157e8261159b565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156115ff5780820151818401526020810190506115e4565b8381111561160e576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b61163881611573565b811461164357600080fd5b50565b61164f81611585565b811461165a57600080fd5b50565b61166681611591565b811461167157600080fd5b50565b61167d816115bb565b811461168857600080fd5b50565b611694816115c5565b811461169f57600080fd5b5056fea26469706673582212204232fe4a8cab608274b46f34042bd5ae7f0a90e9115e9441e9ff7edb02af4e1764736f6c63430006040033"
