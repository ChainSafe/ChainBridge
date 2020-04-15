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
	DestinationChainID             uint8
	ResourceID                     [32]byte
	LenDestinationRecipientAddress *big.Int
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}

// ERC20HandlerABI is the input ABI used to generate the binding from.
const ERC20HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_useContractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"internalType\":\"structERC20Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResourceIDAndContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20HandlerBin is the compiled bytecode used for deploying new contracts.
var ERC20HandlerBin = "0x60806040523480156200001157600080fd5b50604051620026c2380380620026c283398181016040528101906200003791906200035b565b80518251146200007e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000759062000471565b60405180910390fd5b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b8251811015620001175762000109838281518110620000e057fe5b6020026020010151838381518110620000f557fe5b60200260200101516200012160201b60201c565b8080600101915050620000c5565b5050505062000596565b806002600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600081519050620002248162000562565b92915050565b600082601f8301126200023c57600080fd5b8151620002536200024d82620004c1565b62000493565b915081818352602084019350602081019050838560208402820111156200027957600080fd5b60005b83811015620002ad578162000292888262000213565b8452602084019350602083019250506001810190506200027c565b5050505092915050565b600082601f830112620002c957600080fd5b8151620002e0620002da82620004ea565b62000493565b915081818352602084019350602081019050838560208402820111156200030657600080fd5b60005b838110156200033a57816200031f888262000344565b84526020840193506020830192505060018101905062000309565b5050505092915050565b60008151905062000355816200057c565b92915050565b6000806000606084860312156200037157600080fd5b6000620003818682870162000213565b935050602084015167ffffffffffffffff8111156200039f57600080fd5b620003ad86828701620002b7565b925050604084015167ffffffffffffffff811115620003cb57600080fd5b620003d9868287016200022a565b9150509250925092565b6000620003f260478362000513565b91507f6d69736d61746368206c656e677468206265747765656e20696e697469616c5260008301527f65736f7572636549447320616e6420696e697469616c436f6e7472616374416460208301527f64726573736573000000000000000000000000000000000000000000000000006040830152606082019050919050565b600060208201905081810360008301526200048c81620003e3565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715620004b757600080fd5b8060405250919050565b600067ffffffffffffffff821115620004d957600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200050257600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620005318262000542565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200056d8162000524565b81146200057957600080fd5b50565b620005878162000538565b81146200059357600080fd5b50565b61211c80620005a66000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638a025ce81161008c578063d9caed1211610066578063d9caed121461026f578063db95e75c1461028b578063e245386f146102bb578063fc9539cd146102f1576100ea565b80638a025ce81461020757806395601f0914610223578063c8ba6c871461023f576100ea565b80633af32abf116100c85780633af32abf1461015b57806345a104db1461018b5780636ebcf607146101a75780637f79bea8146101d7576100ea565b80630a6d55d8146100ef5780632521ab411461011f578063318c136e1461013d575b600080fd5b61010960048036038101906101049190611744565b61030d565b6040516101169190611ccf565b60405180910390f35b610127610340565b6040516101349190611dc0565b60405180910390f35b610145610353565b6040516101529190611ccf565b60405180910390f35b610175600480360381019061017091906116a3565b610379565b6040516101829190611dc0565b60405180910390f35b6101a560048036038101906101a0919061183c565b6103cf565b005b6101c160048036038101906101bc91906116a3565b61068a565b6040516101ce9190611efa565b60405180910390f35b6101f160048036038101906101ec91906116a3565b6106a2565b6040516101fe9190611dc0565b60405180910390f35b610221600480360381019061021c919061176d565b6106c2565b005b61023d600480360381019061023891906116cc565b610848565b005b610259600480360381019061025491906116a3565b610976565b6040516102669190611ddb565b60405180910390f35b610289600480360381019061028491906116cc565b61098e565b005b6102a560048036038101906102a091906117ea565b610a2e565b6040516102b29190611ed8565b60405180910390f35b6102d560048036038101906102d091906117ea565b610be4565b6040516102e89796959493929190611d4a565b60405180910390f35b61030b600480360381019061030691906117a9565b610d0b565b005b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160149054906101000a900460ff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461045f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045690611e38565b60405180910390fd5b60006060600080602085015193506040850151915060405192508460600151905080830160200160405260e4360360e4843760006002600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506104d281610379565b610511576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050890611eb8565b60405180910390fd5b61051d81883086610f4a565b6040518060e001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018a60ff1681526020018681526020018381526020018581526020018873ffffffffffffffffffffffffffffffffffffffff16815260200184815250600560008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff1602179055506040820151816001015560608201518160020155608082015181600301908051906020019061062a9291906114bd565b5060a08201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160050155905050505050505050505050565b60006020528060005260406000206000915090505481565b60046020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610764576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075b90611e78565b60405180910390fd5b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016107bb9190611cb4565b60405160208183030381529060405280519060200120826040516020016107e29190611cb4565b6040516020818303038152906040528051906020012014610838576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082f90611e18565b60405180910390fd5b6108428484611079565b50505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b815260040161088a93929190611cea565b602060405180830381600087803b1580156108a457600080fd5b505af11580156108b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108dc919061171b565b5061092e826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461116b90919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60036020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1590611e38565b60405180910390fd5b610a298383836111c0565b505050565b610a3661153d565b600560008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b745780601f10610b4957610100808354040283529160200191610b74565b820191906000526020600020905b815481529060010190602001808311610b5757829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b60056020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff1690806001015490806002015490806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610cd55780601f10610caa57610100808354040283529160200191610cd5565b820191906000526020600020905b815481529060010190602001808311610cb857829003601f168201915b5050505050908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154905087565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d9290611e38565b60405180910390fd5b60008060606020840151925060408401519150604051905083606001518082016020016040526084360360848337506000806002600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060208301519150610e1381610379565b610e52576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4990611e98565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663beab71316040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610ec357600080fd5b505af1158015610ed7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610efb9190611813565b90508060ff1686601f60208110610f0e57fe5b1a60f81b60f81c60ff161415610f3157610f2c838560601c896111c0565b610f40565b610f3f838560601c896112ec565b5b5050505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610f8c93929190611cea565b602060405180830381600087803b158015610fa657600080fd5b505af1158015610fba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fde919061171b565b50611030826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461116b90919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b806002600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6000808284019050838110156111b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111ad90611e58565b60405180910390fd5b8091505092915050565b60008390508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b8152600401611200929190611d21565b602060405180830381600087803b15801561121a57600080fd5b505af115801561122e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611252919061171b565b506112a4826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461141890919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b815260040161132c929190611d21565b602060405180830381600087803b15801561134657600080fd5b505af115801561135a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061137e919061171b565b506113d0826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461116b90919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b600061145a83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611462565b905092915050565b60008383111582906114aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114a19190611df6565b60405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106114fe57805160ff191683800117855561152c565b8280016001018555821561152c579182015b8281111561152b578251825591602001919060010190611510565b5b50905061153991906115ac565b5090565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600080191681526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b6115ce91905b808211156115ca5760008160009055506001016115b2565b5090565b90565b6000813590506115e081612073565b92915050565b6000815190506115f58161208a565b92915050565b60008135905061160a816120a1565b92915050565b600082601f83011261162157600080fd5b813561163461162f82611f42565b611f15565b9150808252602083016020830185838301111561165057600080fd5b61165b838284612016565b50505092915050565b600081359050611673816120b8565b92915050565b600081359050611688816120cf565b92915050565b60008151905061169d816120cf565b92915050565b6000602082840312156116b557600080fd5b60006116c3848285016115d1565b91505092915050565b6000806000606084860312156116e157600080fd5b60006116ef868287016115d1565b9350506020611700868287016115d1565b925050604061171186828701611664565b9150509250925092565b60006020828403121561172d57600080fd5b600061173b848285016115e6565b91505092915050565b60006020828403121561175657600080fd5b6000611764848285016115fb565b91505092915050565b6000806040838503121561178057600080fd5b600061178e858286016115fb565b925050602061179f858286016115d1565b9150509250929050565b6000602082840312156117bb57600080fd5b600082013567ffffffffffffffff8111156117d557600080fd5b6117e184828501611610565b91505092915050565b6000602082840312156117fc57600080fd5b600061180a84828501611664565b91505092915050565b60006020828403121561182557600080fd5b60006118338482850161168e565b91505092915050565b6000806000806080858703121561185257600080fd5b600061186087828801611679565b945050602061187187828801611664565b9350506040611882878288016115d1565b925050606085013567ffffffffffffffff81111561189f57600080fd5b6118ab87828801611610565b91505092959194509250565b6118c081611fb7565b82525050565b6118cf81611fb7565b82525050565b6118de81611fc9565b82525050565b6118ed81611fd5565b82525050565b6118fc81611fd5565b82525050565b61191361190e82611fd5565b612058565b82525050565b600061192482611f6e565b61192e8185611f84565b935061193e818560208601612025565b61194781612062565b840191505092915050565b600061195d82611f6e565b6119678185611f95565b9350611977818560208601612025565b61198081612062565b840191505092915050565b600061199682611f79565b6119a08185611fa6565b93506119b0818560208601612025565b6119b981612062565b840191505092915050565b60006119d1603583611fa6565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611a37601e83611fa6565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611a77601b83611fa6565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000611ab7603783611fa6565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b6000611b1d602883611fa6565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611b83603383611fa6565b91507f70726f7669646564206f726967696e436861696e546f6b656e4164647265737360008301527f206973206e6f742077686974656c6973746564000000000000000000000000006020830152604082019050919050565b600060e083016000830151611bf460008601826118b7565b506020830151611c076020860182611c96565b506040830151611c1a60408601826118e4565b506060830151611c2d6060860182611c78565b5060808301518482036080860152611c458282611919565b91505060a0830151611c5a60a08601826118b7565b5060c0830151611c6d60c0860182611c78565b508091505092915050565b611c8181611fff565b82525050565b611c9081611fff565b82525050565b611c9f81612009565b82525050565b611cae81612009565b82525050565b6000611cc08284611902565b60208201915081905092915050565b6000602082019050611ce460008301846118c6565b92915050565b6000606082019050611cff60008301866118c6565b611d0c60208301856118c6565b611d196040830184611c87565b949350505050565b6000604082019050611d3660008301856118c6565b611d436020830184611c87565b9392505050565b600060e082019050611d5f600083018a6118c6565b611d6c6020830189611ca5565b611d7960408301886118f3565b611d866060830187611c87565b8181036080830152611d988186611952565b9050611da760a08301856118c6565b611db460c0830184611c87565b98975050505050505050565b6000602082019050611dd560008301846118d5565b92915050565b6000602082019050611df060008301846118f3565b92915050565b60006020820190508181036000830152611e10818461198b565b905092915050565b60006020820190508181036000830152611e31816119c4565b9050919050565b60006020820190508181036000830152611e5181611a2a565b9050919050565b60006020820190508181036000830152611e7181611a6a565b9050919050565b60006020820190508181036000830152611e9181611aaa565b9050919050565b60006020820190508181036000830152611eb181611b10565b9050919050565b60006020820190508181036000830152611ed181611b76565b9050919050565b60006020820190508181036000830152611ef28184611bdc565b905092915050565b6000602082019050611f0f6000830184611c87565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611f3857600080fd5b8060405250919050565b600067ffffffffffffffff821115611f5957600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611fc282611fdf565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015612043578082015181840152602081019050612028565b83811115612052576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b61207c81611fb7565b811461208757600080fd5b50565b61209381611fc9565b811461209e57600080fd5b50565b6120aa81611fd5565b81146120b557600080fd5b50565b6120c181611fff565b81146120cc57600080fd5b50565b6120d881612009565b81146120e357600080fd5b5056fea26469706673582212208aec434279802894910beb7c3092c07935d166df77b9b6808e39e5fbffab763964736f6c63430006040033"

// DeployERC20Handler deploys a new Ethereum contract, binding an instance of ERC20Handler to it.
func DeployERC20Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address) (common.Address, *types.Transaction, *ERC20Handler, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20HandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses)
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

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) constant returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_contractWhitelist", arg0)
	return *ret0, err
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) constant returns(bool)
func (_ERC20Handler *ERC20HandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.ContractWhitelist(&_ERC20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) constant returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.ContractWhitelist(&_ERC20Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             uint8
	ResourceID                     [32]byte
	LenDestinationRecipientAddress *big.Int
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	ret := new(struct {
		OriginChainTokenAddress        common.Address
		DestinationChainID             uint8
		ResourceID                     [32]byte
		LenDestinationRecipientAddress *big.Int
		DestinationRecipientAddress    []byte
		Depositer                      common.Address
		Amount                         *big.Int
	})
	out := ret
	err := _ERC20Handler.contract.Call(opts, out, "_depositRecords", arg0)
	return *ret, err
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerSession) DepositRecords(arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             uint8
	ResourceID                     [32]byte
	LenDestinationRecipientAddress *big.Int
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	return _ERC20Handler.Contract.DepositRecords(&_ERC20Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerCallerSession) DepositRecords(arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             uint8
	ResourceID                     [32]byte
	LenDestinationRecipientAddress *big.Int
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	return _ERC20Handler.Contract.DepositRecords(&_ERC20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) constant returns(address)
func (_ERC20Handler *ERC20HandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_resourceIDToTokenContractAddress", arg0)
	return *ret0, err
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) constant returns(address)
func (_ERC20Handler *ERC20HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToTokenContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) constant returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToTokenContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) constant returns(bytes32)
func (_ERC20Handler *ERC20HandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_tokenContractAddressToResourceID", arg0)
	return *ret0, err
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) constant returns(bytes32)
func (_ERC20Handler *ERC20HandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.TokenContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) constant returns(bytes32)
func (_ERC20Handler *ERC20HandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.TokenContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// UseContractWhitelist is a free data retrieval call binding the contract method 0x2521ab41.
//
// Solidity: function _useContractWhitelist() constant returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) UseContractWhitelist(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_useContractWhitelist")
	return *ret0, err
}

// UseContractWhitelist is a free data retrieval call binding the contract method 0x2521ab41.
//
// Solidity: function _useContractWhitelist() constant returns(bool)
func (_ERC20Handler *ERC20HandlerSession) UseContractWhitelist() (bool, error) {
	return _ERC20Handler.Contract.UseContractWhitelist(&_ERC20Handler.CallOpts)
}

// UseContractWhitelist is a free data retrieval call binding the contract method 0x2521ab41.
//
// Solidity: function _useContractWhitelist() constant returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) UseContractWhitelist() (bool, error) {
	return _ERC20Handler.Contract.UseContractWhitelist(&_ERC20Handler.CallOpts)
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

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address contractAddress) constant returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) IsWhitelisted(opts *bind.CallOpts, contractAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "isWhitelisted", contractAddress)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address contractAddress) constant returns(bool)
func (_ERC20Handler *ERC20HandlerSession) IsWhitelisted(contractAddress common.Address) (bool, error) {
	return _ERC20Handler.Contract.IsWhitelisted(&_ERC20Handler.CallOpts, contractAddress)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address contractAddress) constant returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) IsWhitelisted(contractAddress common.Address) (bool, error) {
	return _ERC20Handler.Contract.IsWhitelisted(&_ERC20Handler.CallOpts, contractAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactor) Deposit(opts *bind.TransactOpts, destinationChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "deposit", destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerSession) Deposit(destinationChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) Deposit(destinationChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
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

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactor) FundERC20(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "fundERC20", tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.FundERC20(&_ERC20Handler.TransactOpts, tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.FundERC20(&_ERC20Handler.TransactOpts, tokenAddress, owner, amount)
}

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetResourceIDAndContractAddress(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setResourceIDAndContractAddress", resourceID, contractAddress)
}

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetResourceIDAndContractAddress(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResourceIDAndContractAddress(&_ERC20Handler.TransactOpts, resourceID, contractAddress)
}

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetResourceIDAndContractAddress(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResourceIDAndContractAddress(&_ERC20Handler.TransactOpts, resourceID, contractAddress)
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

var RuntimeBytecode = "0x608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638a025ce81161008c578063d9caed1211610066578063d9caed121461026f578063db95e75c1461028b578063e245386f146102bb578063fc9539cd146102f1576100ea565b80638a025ce81461020757806395601f0914610223578063c8ba6c871461023f576100ea565b80633af32abf116100c85780633af32abf1461015b57806345a104db1461018b5780636ebcf607146101a75780637f79bea8146101d7576100ea565b80630a6d55d8146100ef5780632521ab411461011f578063318c136e1461013d575b600080fd5b61010960048036038101906101049190611744565b61030d565b6040516101169190611ccf565b60405180910390f35b610127610340565b6040516101349190611dc0565b60405180910390f35b610145610353565b6040516101529190611ccf565b60405180910390f35b610175600480360381019061017091906116a3565b610379565b6040516101829190611dc0565b60405180910390f35b6101a560048036038101906101a0919061183c565b6103cf565b005b6101c160048036038101906101bc91906116a3565b61068a565b6040516101ce9190611efa565b60405180910390f35b6101f160048036038101906101ec91906116a3565b6106a2565b6040516101fe9190611dc0565b60405180910390f35b610221600480360381019061021c919061176d565b6106c2565b005b61023d600480360381019061023891906116cc565b610848565b005b610259600480360381019061025491906116a3565b610976565b6040516102669190611ddb565b60405180910390f35b610289600480360381019061028491906116cc565b61098e565b005b6102a560048036038101906102a091906117ea565b610a2e565b6040516102b29190611ed8565b60405180910390f35b6102d560048036038101906102d091906117ea565b610be4565b6040516102e89796959493929190611d4a565b60405180910390f35b61030b600480360381019061030691906117a9565b610d0b565b005b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160149054906101000a900460ff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461045f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045690611e38565b60405180910390fd5b60006060600080602085015193506040850151915060405192508460600151905080830160200160405260e4360360e4843760006002600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506104d281610379565b610511576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050890611eb8565b60405180910390fd5b61051d81883086610f4a565b6040518060e001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018a60ff1681526020018681526020018381526020018581526020018873ffffffffffffffffffffffffffffffffffffffff16815260200184815250600560008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff1602179055506040820151816001015560608201518160020155608082015181600301908051906020019061062a9291906114bd565b5060a08201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160050155905050505050505050505050565b60006020528060005260406000206000915090505481565b60046020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610764576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075b90611e78565b60405180910390fd5b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016107bb9190611cb4565b60405160208183030381529060405280519060200120826040516020016107e29190611cb4565b6040516020818303038152906040528051906020012014610838576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082f90611e18565b60405180910390fd5b6108428484611079565b50505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b815260040161088a93929190611cea565b602060405180830381600087803b1580156108a457600080fd5b505af11580156108b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108dc919061171b565b5061092e826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461116b90919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60036020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1590611e38565b60405180910390fd5b610a298383836111c0565b505050565b610a3661153d565b600560008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b745780601f10610b4957610100808354040283529160200191610b74565b820191906000526020600020905b815481529060010190602001808311610b5757829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b60056020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff1690806001015490806002015490806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610cd55780601f10610caa57610100808354040283529160200191610cd5565b820191906000526020600020905b815481529060010190602001808311610cb857829003601f168201915b5050505050908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154905087565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d9290611e38565b60405180910390fd5b60008060606020840151925060408401519150604051905083606001518082016020016040526084360360848337506000806002600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060208301519150610e1381610379565b610e52576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4990611e98565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663beab71316040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610ec357600080fd5b505af1158015610ed7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610efb9190611813565b90508060ff1686601f60208110610f0e57fe5b1a60f81b60f81c60ff161415610f3157610f2c838560601c896111c0565b610f40565b610f3f838560601c896112ec565b5b5050505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610f8c93929190611cea565b602060405180830381600087803b158015610fa657600080fd5b505af1158015610fba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fde919061171b565b50611030826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461116b90919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b806002600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6000808284019050838110156111b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111ad90611e58565b60405180910390fd5b8091505092915050565b60008390508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b8152600401611200929190611d21565b602060405180830381600087803b15801561121a57600080fd5b505af115801561122e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611252919061171b565b506112a4826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461141890919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b815260040161132c929190611d21565b602060405180830381600087803b15801561134657600080fd5b505af115801561135a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061137e919061171b565b506113d0826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461116b90919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b600061145a83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611462565b905092915050565b60008383111582906114aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114a19190611df6565b60405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106114fe57805160ff191683800117855561152c565b8280016001018555821561152c579182015b8281111561152b578251825591602001919060010190611510565b5b50905061153991906115ac565b5090565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600080191681526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b6115ce91905b808211156115ca5760008160009055506001016115b2565b5090565b90565b6000813590506115e081612073565b92915050565b6000815190506115f58161208a565b92915050565b60008135905061160a816120a1565b92915050565b600082601f83011261162157600080fd5b813561163461162f82611f42565b611f15565b9150808252602083016020830185838301111561165057600080fd5b61165b838284612016565b50505092915050565b600081359050611673816120b8565b92915050565b600081359050611688816120cf565b92915050565b60008151905061169d816120cf565b92915050565b6000602082840312156116b557600080fd5b60006116c3848285016115d1565b91505092915050565b6000806000606084860312156116e157600080fd5b60006116ef868287016115d1565b9350506020611700868287016115d1565b925050604061171186828701611664565b9150509250925092565b60006020828403121561172d57600080fd5b600061173b848285016115e6565b91505092915050565b60006020828403121561175657600080fd5b6000611764848285016115fb565b91505092915050565b6000806040838503121561178057600080fd5b600061178e858286016115fb565b925050602061179f858286016115d1565b9150509250929050565b6000602082840312156117bb57600080fd5b600082013567ffffffffffffffff8111156117d557600080fd5b6117e184828501611610565b91505092915050565b6000602082840312156117fc57600080fd5b600061180a84828501611664565b91505092915050565b60006020828403121561182557600080fd5b60006118338482850161168e565b91505092915050565b6000806000806080858703121561185257600080fd5b600061186087828801611679565b945050602061187187828801611664565b9350506040611882878288016115d1565b925050606085013567ffffffffffffffff81111561189f57600080fd5b6118ab87828801611610565b91505092959194509250565b6118c081611fb7565b82525050565b6118cf81611fb7565b82525050565b6118de81611fc9565b82525050565b6118ed81611fd5565b82525050565b6118fc81611fd5565b82525050565b61191361190e82611fd5565b612058565b82525050565b600061192482611f6e565b61192e8185611f84565b935061193e818560208601612025565b61194781612062565b840191505092915050565b600061195d82611f6e565b6119678185611f95565b9350611977818560208601612025565b61198081612062565b840191505092915050565b600061199682611f79565b6119a08185611fa6565b93506119b0818560208601612025565b6119b981612062565b840191505092915050565b60006119d1603583611fa6565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611a37601e83611fa6565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611a77601b83611fa6565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000611ab7603783611fa6565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b6000611b1d602883611fa6565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611b83603383611fa6565b91507f70726f7669646564206f726967696e436861696e546f6b656e4164647265737360008301527f206973206e6f742077686974656c6973746564000000000000000000000000006020830152604082019050919050565b600060e083016000830151611bf460008601826118b7565b506020830151611c076020860182611c96565b506040830151611c1a60408601826118e4565b506060830151611c2d6060860182611c78565b5060808301518482036080860152611c458282611919565b91505060a0830151611c5a60a08601826118b7565b5060c0830151611c6d60c0860182611c78565b508091505092915050565b611c8181611fff565b82525050565b611c9081611fff565b82525050565b611c9f81612009565b82525050565b611cae81612009565b82525050565b6000611cc08284611902565b60208201915081905092915050565b6000602082019050611ce460008301846118c6565b92915050565b6000606082019050611cff60008301866118c6565b611d0c60208301856118c6565b611d196040830184611c87565b949350505050565b6000604082019050611d3660008301856118c6565b611d436020830184611c87565b9392505050565b600060e082019050611d5f600083018a6118c6565b611d6c6020830189611ca5565b611d7960408301886118f3565b611d866060830187611c87565b8181036080830152611d988186611952565b9050611da760a08301856118c6565b611db460c0830184611c87565b98975050505050505050565b6000602082019050611dd560008301846118d5565b92915050565b6000602082019050611df060008301846118f3565b92915050565b60006020820190508181036000830152611e10818461198b565b905092915050565b60006020820190508181036000830152611e31816119c4565b9050919050565b60006020820190508181036000830152611e5181611a2a565b9050919050565b60006020820190508181036000830152611e7181611a6a565b9050919050565b60006020820190508181036000830152611e9181611aaa565b9050919050565b60006020820190508181036000830152611eb181611b10565b9050919050565b60006020820190508181036000830152611ed181611b76565b9050919050565b60006020820190508181036000830152611ef28184611bdc565b905092915050565b6000602082019050611f0f6000830184611c87565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611f3857600080fd5b8060405250919050565b600067ffffffffffffffff821115611f5957600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611fc282611fdf565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015612043578082015181840152602081019050612028565b83811115612052576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b61207c81611fb7565b811461208757600080fd5b50565b61209381611fc9565b811461209e57600080fd5b50565b6120aa81611fd5565b81146120b557600080fd5b50565b6120c181611fff565b81146120cc57600080fd5b50565b6120d881612009565b81146120e357600080fd5b5056fea26469706673582212208aec434279802894910beb7c3092c07935d166df77b9b6808e39e5fbffab763964736f6c63430006040033"
