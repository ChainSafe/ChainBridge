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
const GenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResourceIDAndContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structGenericHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GenericHandlerBin is the compiled bytecode used for deploying new contracts.
var GenericHandlerBin = "0x60806040523480156200001157600080fd5b5060405162001c2b38038062001c2b83398181016040528101906200003791906200035b565b80518251146200007e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000759062000471565b60405180910390fd5b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b8251811015620001175762000109838281518110620000e057fe5b6020026020010151838381518110620000f557fe5b60200260200101516200012160201b60201c565b8080600101915050620000c5565b5050505062000596565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600081519050620002248162000562565b92915050565b600082601f8301126200023c57600080fd5b8151620002536200024d82620004c1565b62000493565b915081818352602084019350602081019050838560208402820111156200027957600080fd5b60005b83811015620002ad578162000292888262000213565b8452602084019350602083019250506001810190506200027c565b5050505092915050565b600082601f830112620002c957600080fd5b8151620002e0620002da82620004ea565b62000493565b915081818352602084019350602081019050838560208402820111156200030657600080fd5b60005b838110156200033a57816200031f888262000344565b84526020840193506020830192505060018101905062000309565b5050505092915050565b60008151905062000355816200057c565b92915050565b6000806000606084860312156200037157600080fd5b6000620003818682870162000213565b935050602084015167ffffffffffffffff8111156200039f57600080fd5b620003ad86828701620002b7565b925050604084015167ffffffffffffffff811115620003cb57600080fd5b620003d9868287016200022a565b9150509250925092565b6000620003f260478362000513565b91507f6d69736d61746368206c656e677468206265747765656e20696e697469616c5260008301527f65736f7572636549447320616e6420696e697469616c436f6e7472616374416460208301527f64726573736573000000000000000000000000000000000000000000000000006040830152606082019050919050565b600060208201905081810360008301526200048c81620003e3565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715620004b757600080fd5b8060405250919050565b600067ffffffffffffffff821115620004d957600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200050257600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620005318262000542565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200056d8162000524565b81146200057957600080fd5b50565b620005878162000538565b81146200059357600080fd5b50565b61168580620005a66000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638a025ce8116100715780638a025ce81461017857806395601f0914610194578063c8ba6c87146101b0578063db95e75c146101e0578063e245386f14610210578063fc9539cd14610244576100a9565b80630a6d55d8146100ae578063318c136e146100de57806345a104db146100fc5780636ebcf607146101185780637f79bea814610148575b600080fd5b6100c860048036038101906100c39190610ec2565b610260565b6040516100d591906112ea565b60405180910390f35b6100e6610293565b6040516100f391906112ea565b60405180910390f35b61011660048036038101906101119190610f91565b6102b9565b005b610132600480360381019061012d9190610e21565b6104c3565b60405161013f9190611414565b60405180910390f35b610162600480360381019061015d9190610e21565b6104db565b60405161016f919061133c565b60405180910390f35b610192600480360381019061018d9190610eeb565b6104fb565b005b6101ae60048036038101906101a99190610e4a565b610681565b005b6101ca60048036038101906101c59190610e21565b6107af565b6040516101d79190611357565b60405180910390f35b6101fa60048036038101906101f59190610f68565b6107c7565b60405161020791906113f2565b60405180910390f35b61022a60048036038101906102259190610f68565b610969565b60405161023b95949392919061142f565b60405180910390f35b61025e60048036038101906102599190610f27565b610a84565b005b60036020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610349576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034090611392565b60405180910390fd5b600080606060208401519250604084015191506040519050836060015180820160600160405260e4360360e48337506040518060a001604052808860ff1681526020018381526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001828152506002600088815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040190805190602001906104b6929190610c5e565b5090505050505050505050565b60006020528060005260406000206000915090505481565b60066020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166003600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461059d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610594906113d2565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016105f491906112cf565b604051602081830303815290604052805190602001208260405160200161061b91906112cf565b6040516020818303038152906040528051906020012014610671576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066890611372565b60405180910390fd5b61067b8484610b17565b50505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b81526004016106c393929190611305565b602060405180830381600087803b1580156106dd57600080fd5b505af11580156106f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107159190610e99565b50610767826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c0990919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60046020528060005260406000206000915090505481565b6107cf610cde565b600260008381526020019081526020016000206040518060a00160405290816000820160009054906101000a900460ff1660ff1660ff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109595780601f1061092e57610100808354040283529160200191610959565b820191906000526020600020905b81548152906001019060200180831161093c57829003601f168201915b5050505050815250509050919050565b60026020528060005260406000206000915090508060000160009054906101000a900460ff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806004018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a7a5780601f10610a4f57610100808354040283529160200191610a7a565b820191906000526020600020905b815481529060010190602001808311610a5d57829003601f168201915b5050505050905085565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b14576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0b90611392565b60405180910390fd5b50565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600080828401905083811015610c54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4b906113b2565b60405180910390fd5b8091505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c9f57805160ff1916838001178555610ccd565b82800160010185558215610ccd579182015b82811115610ccc578251825591602001919060010190610cb1565b5b509050610cda9190610d3f565b5090565b6040518060a00160405280600060ff16815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b610d6191905b80821115610d5d576000816000905550600101610d45565b5090565b90565b600081359050610d73816115dc565b92915050565b600081519050610d88816115f3565b92915050565b600081359050610d9d8161160a565b92915050565b600082601f830112610db457600080fd5b8135610dc7610dc2826114b6565b611489565b91508082526020830160208301858383011115610de357600080fd5b610dee83828461157f565b50505092915050565b600081359050610e0681611621565b92915050565b600081359050610e1b81611638565b92915050565b600060208284031215610e3357600080fd5b6000610e4184828501610d64565b91505092915050565b600080600060608486031215610e5f57600080fd5b6000610e6d86828701610d64565b9350506020610e7e86828701610d64565b9250506040610e8f86828701610df7565b9150509250925092565b600060208284031215610eab57600080fd5b6000610eb984828501610d79565b91505092915050565b600060208284031215610ed457600080fd5b6000610ee284828501610d8e565b91505092915050565b60008060408385031215610efe57600080fd5b6000610f0c85828601610d8e565b9250506020610f1d85828601610d64565b9150509250929050565b600060208284031215610f3957600080fd5b600082013567ffffffffffffffff811115610f5357600080fd5b610f5f84828501610da3565b91505092915050565b600060208284031215610f7a57600080fd5b6000610f8884828501610df7565b91505092915050565b60008060008060808587031215610fa757600080fd5b6000610fb587828801610e0c565b9450506020610fc687828801610df7565b9350506040610fd787828801610d64565b925050606085013567ffffffffffffffff811115610ff457600080fd5b61100087828801610da3565b91505092959194509250565b61101581611520565b82525050565b61102481611520565b82525050565b61103381611532565b82525050565b6110428161153e565b82525050565b6110518161153e565b82525050565b6110686110638261153e565b6115c1565b82525050565b6000611079826114e2565b61108381856114ed565b935061109381856020860161158e565b61109c816115cb565b840191505092915050565b60006110b2826114e2565b6110bc81856114fe565b93506110cc81856020860161158e565b6110d5816115cb565b840191505092915050565b60006110ed60358361150f565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611153601e8361150f565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611193601b8361150f565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b60006111d360378361150f565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b600060a08301600083015161124460008601826112b1565b5060208301516112576020860182611039565b50604083015161126a604086018261100c565b50606083015161127d606086018261100c565b5060808301518482036080860152611295828261106e565b9150508091505092915050565b6112ab81611568565b82525050565b6112ba81611572565b82525050565b6112c981611572565b82525050565b60006112db8284611057565b60208201915081905092915050565b60006020820190506112ff600083018461101b565b92915050565b600060608201905061131a600083018661101b565b611327602083018561101b565b61133460408301846112a2565b949350505050565b6000602082019050611351600083018461102a565b92915050565b600060208201905061136c6000830184611048565b92915050565b6000602082019050818103600083015261138b816110e0565b9050919050565b600060208201905081810360008301526113ab81611146565b9050919050565b600060208201905081810360008301526113cb81611186565b9050919050565b600060208201905081810360008301526113eb816111c6565b9050919050565b6000602082019050818103600083015261140c818461122c565b905092915050565b600060208201905061142960008301846112a2565b92915050565b600060a08201905061144460008301886112c0565b6114516020830187611048565b61145e604083018661101b565b61146b606083018561101b565b818103608083015261147d81846110a7565b90509695505050505050565b6000604051905081810181811067ffffffffffffffff821117156114ac57600080fd5b8060405250919050565b600067ffffffffffffffff8211156114cd57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061152b82611548565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156115ac578082015181840152602081019050611591565b838111156115bb576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b6115e581611520565b81146115f057600080fd5b50565b6115fc81611532565b811461160757600080fd5b50565b6116138161153e565b811461161e57600080fd5b50565b61162a81611568565b811461163557600080fd5b50565b61164181611572565b811461164c57600080fd5b5056fea2646970667358221220595a279a00a6789fdf4f5e18eaa410749a9af1216da21fd6b3fcaed28179e14864736f6c63430006040033"

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

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) constant returns(bool)
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
// Solidity: function _contractWhitelist(address ) constant returns(bool)
func (_GenericHandler *GenericHandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _GenericHandler.Contract.ContractWhitelist(&_GenericHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) constant returns(bool)
func (_GenericHandler *GenericHandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _GenericHandler.Contract.ContractWhitelist(&_GenericHandler.CallOpts, arg0)
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

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) constant returns(address)
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
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) constant returns(address)
func (_GenericHandler *GenericHandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToTokenContractAddress(&_GenericHandler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) constant returns(address)
func (_GenericHandler *GenericHandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToTokenContractAddress(&_GenericHandler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) constant returns(bytes32)
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
// Solidity: function _tokenContractAddressToResourceID(address ) constant returns(bytes32)
func (_GenericHandler *GenericHandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.TokenContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) constant returns(bytes32)
func (_GenericHandler *GenericHandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.TokenContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
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

var RuntimeBytecode = "0x608060405234801561001057600080fd5b50600436106100a95760003560e01c80638a025ce8116100715780638a025ce81461017857806395601f0914610194578063c8ba6c87146101b0578063db95e75c146101e0578063e245386f14610210578063fc9539cd14610244576100a9565b80630a6d55d8146100ae578063318c136e146100de57806345a104db146100fc5780636ebcf607146101185780637f79bea814610148575b600080fd5b6100c860048036038101906100c39190610ec2565b610260565b6040516100d591906112ea565b60405180910390f35b6100e6610293565b6040516100f391906112ea565b60405180910390f35b61011660048036038101906101119190610f91565b6102b9565b005b610132600480360381019061012d9190610e21565b6104c3565b60405161013f9190611414565b60405180910390f35b610162600480360381019061015d9190610e21565b6104db565b60405161016f919061133c565b60405180910390f35b610192600480360381019061018d9190610eeb565b6104fb565b005b6101ae60048036038101906101a99190610e4a565b610681565b005b6101ca60048036038101906101c59190610e21565b6107af565b6040516101d79190611357565b60405180910390f35b6101fa60048036038101906101f59190610f68565b6107c7565b60405161020791906113f2565b60405180910390f35b61022a60048036038101906102259190610f68565b610969565b60405161023b95949392919061142f565b60405180910390f35b61025e60048036038101906102599190610f27565b610a84565b005b60036020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610349576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034090611392565b60405180910390fd5b600080606060208401519250604084015191506040519050836060015180820160600160405260e4360360e48337506040518060a001604052808860ff1681526020018381526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001828152506002600088815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040190805190602001906104b6929190610c5e565b5090505050505050505050565b60006020528060005260406000206000915090505481565b60066020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166003600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461059d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610594906113d2565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016105f491906112cf565b604051602081830303815290604052805190602001208260405160200161061b91906112cf565b6040516020818303038152906040528051906020012014610671576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066890611372565b60405180910390fd5b61067b8484610b17565b50505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b81526004016106c393929190611305565b602060405180830381600087803b1580156106dd57600080fd5b505af11580156106f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107159190610e99565b50610767826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c0990919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60046020528060005260406000206000915090505481565b6107cf610cde565b600260008381526020019081526020016000206040518060a00160405290816000820160009054906101000a900460ff1660ff1660ff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109595780601f1061092e57610100808354040283529160200191610959565b820191906000526020600020905b81548152906001019060200180831161093c57829003601f168201915b5050505050815250509050919050565b60026020528060005260406000206000915090508060000160009054906101000a900460ff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806004018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a7a5780601f10610a4f57610100808354040283529160200191610a7a565b820191906000526020600020905b815481529060010190602001808311610a5d57829003601f168201915b5050505050905085565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b14576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0b90611392565b60405180910390fd5b50565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600080828401905083811015610c54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4b906113b2565b60405180910390fd5b8091505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c9f57805160ff1916838001178555610ccd565b82800160010185558215610ccd579182015b82811115610ccc578251825591602001919060010190610cb1565b5b509050610cda9190610d3f565b5090565b6040518060a00160405280600060ff16815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b610d6191905b80821115610d5d576000816000905550600101610d45565b5090565b90565b600081359050610d73816115dc565b92915050565b600081519050610d88816115f3565b92915050565b600081359050610d9d8161160a565b92915050565b600082601f830112610db457600080fd5b8135610dc7610dc2826114b6565b611489565b91508082526020830160208301858383011115610de357600080fd5b610dee83828461157f565b50505092915050565b600081359050610e0681611621565b92915050565b600081359050610e1b81611638565b92915050565b600060208284031215610e3357600080fd5b6000610e4184828501610d64565b91505092915050565b600080600060608486031215610e5f57600080fd5b6000610e6d86828701610d64565b9350506020610e7e86828701610d64565b9250506040610e8f86828701610df7565b9150509250925092565b600060208284031215610eab57600080fd5b6000610eb984828501610d79565b91505092915050565b600060208284031215610ed457600080fd5b6000610ee284828501610d8e565b91505092915050565b60008060408385031215610efe57600080fd5b6000610f0c85828601610d8e565b9250506020610f1d85828601610d64565b9150509250929050565b600060208284031215610f3957600080fd5b600082013567ffffffffffffffff811115610f5357600080fd5b610f5f84828501610da3565b91505092915050565b600060208284031215610f7a57600080fd5b6000610f8884828501610df7565b91505092915050565b60008060008060808587031215610fa757600080fd5b6000610fb587828801610e0c565b9450506020610fc687828801610df7565b9350506040610fd787828801610d64565b925050606085013567ffffffffffffffff811115610ff457600080fd5b61100087828801610da3565b91505092959194509250565b61101581611520565b82525050565b61102481611520565b82525050565b61103381611532565b82525050565b6110428161153e565b82525050565b6110518161153e565b82525050565b6110686110638261153e565b6115c1565b82525050565b6000611079826114e2565b61108381856114ed565b935061109381856020860161158e565b61109c816115cb565b840191505092915050565b60006110b2826114e2565b6110bc81856114fe565b93506110cc81856020860161158e565b6110d5816115cb565b840191505092915050565b60006110ed60358361150f565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611153601e8361150f565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611193601b8361150f565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b60006111d360378361150f565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b600060a08301600083015161124460008601826112b1565b5060208301516112576020860182611039565b50604083015161126a604086018261100c565b50606083015161127d606086018261100c565b5060808301518482036080860152611295828261106e565b9150508091505092915050565b6112ab81611568565b82525050565b6112ba81611572565b82525050565b6112c981611572565b82525050565b60006112db8284611057565b60208201915081905092915050565b60006020820190506112ff600083018461101b565b92915050565b600060608201905061131a600083018661101b565b611327602083018561101b565b61133460408301846112a2565b949350505050565b6000602082019050611351600083018461102a565b92915050565b600060208201905061136c6000830184611048565b92915050565b6000602082019050818103600083015261138b816110e0565b9050919050565b600060208201905081810360008301526113ab81611146565b9050919050565b600060208201905081810360008301526113cb81611186565b9050919050565b600060208201905081810360008301526113eb816111c6565b9050919050565b6000602082019050818103600083015261140c818461122c565b905092915050565b600060208201905061142960008301846112a2565b92915050565b600060a08201905061144460008301886112c0565b6114516020830187611048565b61145e604083018661101b565b61146b606083018561101b565b818103608083015261147d81846110a7565b90509695505050505050565b6000604051905081810181811067ffffffffffffffff821117156114ac57600080fd5b8060405250919050565b600067ffffffffffffffff8211156114cd57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061152b82611548565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156115ac578082015181840152602081019050611591565b838111156115bb576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b6115e581611520565b81146115f057600080fd5b50565b6115fc81611532565b811461160757600080fd5b50565b6116138161153e565b811461161e57600080fd5b50565b61162a81611568565b811461163557600080fd5b50565b61164181611572565b811461164c57600080fd5b5056fea2646970667358221220595a279a00a6789fdf4f5e18eaa410749a9af1216da21fd6b3fcaed28179e14864736f6c63430006040033"
