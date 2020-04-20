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
	DestinationChainID             uint8
	ResourceID                     [32]byte
	LenDestinationRecipientAddress *big.Int
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	TokenID                        *big.Int
	MetaData                       []byte
}

// ERC721HandlerABI is the input ABI used to generate the binding from.
const ERC721HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"fundERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structERC721Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResourceIDAndContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721HandlerBin is the compiled bytecode used for deploying new contracts.
var ERC721HandlerBin = "0x60806040523480156200001157600080fd5b50604051620027973803806200279783398181016040528101906200003791906200035b565b80518251146200007e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000759062000471565b60405180910390fd5b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b8251811015620001175762000109838281518110620000e057fe5b6020026020010151838381518110620000f557fe5b60200260200101516200012160201b60201c565b8080600101915050620000c5565b5050505062000596565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600081519050620002248162000562565b92915050565b600082601f8301126200023c57600080fd5b8151620002536200024d82620004c1565b62000493565b915081818352602084019350602081019050838560208402820111156200027957600080fd5b60005b83811015620002ad578162000292888262000213565b8452602084019350602083019250506001810190506200027c565b5050505092915050565b600082601f830112620002c957600080fd5b8151620002e0620002da82620004ea565b62000493565b915081818352602084019350602081019050838560208402820111156200030657600080fd5b60005b838110156200033a57816200031f888262000344565b84526020840193506020830192505060018101905062000309565b5050505092915050565b60008151905062000355816200057c565b92915050565b6000806000606084860312156200037157600080fd5b6000620003818682870162000213565b935050602084015167ffffffffffffffff8111156200039f57600080fd5b620003ad86828701620002b7565b925050604084015167ffffffffffffffff811115620003cb57600080fd5b620003d9868287016200022a565b9150509250925092565b6000620003f260478362000513565b91507f6d69736d61746368206c656e677468206265747765656e20696e697469616c5260008301527f65736f7572636549447320616e6420696e697469616c436f6e7472616374416460208301527f64726573736573000000000000000000000000000000000000000000000000006040830152606082019050919050565b600060208201905081810360008301526200048c81620003e3565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715620004b757600080fd5b8060405250919050565b600067ffffffffffffffff821115620004d957600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200050257600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620005318262000542565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200056d8162000524565b81146200057957600080fd5b50565b620005878162000538565b81146200059357600080fd5b50565b6121f180620005a66000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638a025ce8116100715780638a025ce81461019f578063c8ba6c87146101bb578063d9caed12146101eb578063db95e75c14610207578063e245386f14610237578063fc9539cd1461026e576100b4565b80630a6d55d8146100b9578063318c136e146100e957806345a104db146101075780636ebcf6071461012357806373542980146101535780637f79bea81461016f575b600080fd5b6100d360048036038101906100ce919061174a565b61028a565b6040516100e09190611d38565b60405180910390f35b6100f16102bd565b6040516100fe9190611d38565b60405180910390f35b610121600480360381019061011c9190611842565b6102e3565b005b61013d600480360381019061013891906116a9565b6105eb565b60405161014a9190611f8e565b60405180910390f35b61016d600480360381019061016891906116d2565b610603565b005b610189600480360381019061018491906116a9565b610711565b6040516101969190611e54565b60405180910390f35b6101b960048036038101906101b49190611773565b610731565b005b6101d560048036038101906101d091906116a9565b6108b7565b6040516101e29190611e6f565b60405180910390f35b610205600480360381019061020091906116d2565b6108cf565b005b610221600480360381019061021c91906117f0565b610970565b60405161022e9190611f6c565b60405180910390f35b610251600480360381019061024c91906117f0565b610bc9565b604051610265989796959493929190611dc8565b60405180910390f35b610288600480360381019061028391906117af565b610d8e565b005b60036020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610373576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036a90611ecc565b60405180910390fd5b60008060006060806020860151945060408601519250604051915060608601519350836080018601518483016040016040528060200160e401360360e4843760405191508082016040016040528461010401803603818437505060006003600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905061040e8161105b565b61044d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044490611f4c565b60405180910390fd5b610459818930876110b1565b6040518061010001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018b60ff1681526020018781526020018681526020018481526020018973ffffffffffffffffffffffffffffffffffffffff16815260200185815260200183815250600260008b815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff1602179055506040820151816001015560608201518160020155608082015181600301908051906020019061056d9291906114bb565b5060a08201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816005015560e08201518160060190805190602001906105db9291906114bb565b5090505050505050505050505050565b60006020528060005260406000206000915090505481565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b815260040161064593929190611d91565b600060405180830381600087803b15801561065f57600080fd5b505af1158015610673573d6000803e3d6000fd5b505050506106c960016000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111c090919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60056020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166003600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146107d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ca90611f0c565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008060405160200161082a9190611d1d565b60405160208183030381529060405280519060200120826040516020016108519190611d1d565b60405160208183030381529060405280519060200120146108a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089e90611eac565b60405180910390fd5b6108b18484611215565b50505050565b60046020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461095f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095690611ecc565b60405180910390fd5b61096b83308484611307565b505050565b61097861153b565b60026000838152602001908152602001600020604051806101000160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ab75780601f10610a8c57610100808354040283529160200191610ab7565b820191906000526020600020905b815481529060010190602001808311610a9a57829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160058201548152602001600682018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bb95780601f10610b8e57610100808354040283529160200191610bb9565b820191906000526020600020905b815481529060010190602001808311610b9c57829003601f168201915b5050505050815250509050919050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff1690806001015490806002015490806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610cba5780601f10610c8f57610100808354040283529160200191610cba565b820191906000526020600020905b815481529060010190602001808311610c9d57829003601f168201915b5050505050908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806005015490806006018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d845780601f10610d5957610100808354040283529160200191610d84565b820191906000526020600020905b815481529060010190602001808311610d6757829003601f168201915b5050505050905088565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1590611ecc565b60405180910390fd5b600080606080602085015193506040850151925060405191506060850151806080018601518184016040016040528060200160840136036084853760405192508083016040016040528160a40180360381853750505060008060208401519150604b8701519050610e918160601c61105b565b610ed0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ec790611f2c565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663beab71316040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610f4157600080fd5b505af1158015610f55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f799190611819565b90508060ff1687601f60208110610f8c57fe5b1a60f81b60f81c60ff161415610fb357610fae8360601c308660601c8b611307565b611050565b60008360601c90508073ffffffffffffffffffffffffffffffffffffffff16638832e6e38660601c8b896040518463ffffffff1660e01b8152600401610ffb93929190611d53565b602060405180830381600087803b15801561101557600080fd5b505af1158015611029573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061104d9190611721565b50505b505050505050505050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b81526004016110f393929190611d91565b600060405180830381600087803b15801561110d57600080fd5b505af1158015611121573d6000803e3d6000fd5b5050505061117760016000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111c090919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b60008082840190508381101561120b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161120290611eec565b60405180910390fd5b8091505092915050565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b815260040161134993929190611d91565b600060405180830381600087803b15801561136357600080fd5b505af1158015611377573d6000803e3d6000fd5b505050506113cd60016000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461141690919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b600061145883836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611460565b905092915050565b60008383111582906114a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161149f9190611e8a565b60405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106114fc57805160ff191683800117855561152a565b8280016001018555821561152a579182015b8281111561152957825182559160200191906001019061150e565b5b50905061153791906115b2565b5090565b604051806101000160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600080191681526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001606081525090565b6115d491905b808211156115d05760008160009055506001016115b8565b5090565b90565b6000813590506115e681612148565b92915050565b6000815190506115fb8161215f565b92915050565b60008135905061161081612176565b92915050565b600082601f83011261162757600080fd5b813561163a61163582611fd6565b611fa9565b9150808252602083016020830185838301111561165657600080fd5b6116618382846120eb565b50505092915050565b6000813590506116798161218d565b92915050565b60008135905061168e816121a4565b92915050565b6000815190506116a3816121a4565b92915050565b6000602082840312156116bb57600080fd5b60006116c9848285016115d7565b91505092915050565b6000806000606084860312156116e757600080fd5b60006116f5868287016115d7565b9350506020611706868287016115d7565b92505060406117178682870161166a565b9150509250925092565b60006020828403121561173357600080fd5b6000611741848285016115ec565b91505092915050565b60006020828403121561175c57600080fd5b600061176a84828501611601565b91505092915050565b6000806040838503121561178657600080fd5b600061179485828601611601565b92505060206117a5858286016115d7565b9150509250929050565b6000602082840312156117c157600080fd5b600082013567ffffffffffffffff8111156117db57600080fd5b6117e784828501611616565b91505092915050565b60006020828403121561180257600080fd5b60006118108482850161166a565b91505092915050565b60006020828403121561182b57600080fd5b600061183984828501611694565b91505092915050565b6000806000806080858703121561185857600080fd5b60006118668782880161167f565b94505060206118778782880161166a565b9350506040611888878288016115d7565b925050606085013567ffffffffffffffff8111156118a557600080fd5b6118b187828801611616565b91505092959194509250565b6118c6816120b5565b82525050565b6118d581612056565b82525050565b6118e481612056565b82525050565b6118f381612068565b82525050565b61190281612074565b82525050565b61191181612074565b82525050565b61192861192382612074565b61212d565b82525050565b60006119398261200d565b6119438185612034565b93506119538185602086016120fa565b61195c81612137565b840191505092915050565b600061197282612002565b61197c8185612023565b935061198c8185602086016120fa565b61199581612137565b840191505092915050565b60006119ab82612002565b6119b58185612034565b93506119c58185602086016120fa565b6119ce81612137565b840191505092915050565b60006119e482612018565b6119ee8185612045565b93506119fe8185602086016120fa565b611a0781612137565b840191505092915050565b6000611a1f603583612045565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611a85601e83612045565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611ac5601b83612045565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000611b05603783612045565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b6000611b6b602883612045565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611bd1603383612045565b91507f70726f7669646564206f726967696e436861696e546f6b656e4164647265737360008301527f206973206e6f742077686974656c6973746564000000000000000000000000006020830152604082019050919050565b600061010083016000830151611c4360008601826118cc565b506020830151611c566020860182611cff565b506040830151611c6960408601826118f9565b506060830151611c7c6060860182611ce1565b5060808301518482036080860152611c948282611967565b91505060a0830151611ca960a08601826118cc565b5060c0830151611cbc60c0860182611ce1565b5060e083015184820360e0860152611cd48282611967565b9150508091505092915050565b611cea8161209e565b82525050565b611cf98161209e565b82525050565b611d08816120a8565b82525050565b611d17816120a8565b82525050565b6000611d298284611917565b60208201915081905092915050565b6000602082019050611d4d60008301846118db565b92915050565b6000606082019050611d6860008301866118bd565b611d756020830185611cf0565b8181036040830152611d87818461192e565b9050949350505050565b6000606082019050611da660008301866118db565b611db360208301856118db565b611dc06040830184611cf0565b949350505050565b600061010082019050611dde600083018b6118db565b611deb602083018a611d0e565b611df86040830189611908565b611e056060830188611cf0565b8181036080830152611e1781876119a0565b9050611e2660a08301866118db565b611e3360c0830185611cf0565b81810360e0830152611e4581846119a0565b90509998505050505050505050565b6000602082019050611e6960008301846118ea565b92915050565b6000602082019050611e846000830184611908565b92915050565b60006020820190508181036000830152611ea481846119d9565b905092915050565b60006020820190508181036000830152611ec581611a12565b9050919050565b60006020820190508181036000830152611ee581611a78565b9050919050565b60006020820190508181036000830152611f0581611ab8565b9050919050565b60006020820190508181036000830152611f2581611af8565b9050919050565b60006020820190508181036000830152611f4581611b5e565b9050919050565b60006020820190508181036000830152611f6581611bc4565b9050919050565b60006020820190508181036000830152611f868184611c2a565b905092915050565b6000602082019050611fa36000830184611cf0565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611fcc57600080fd5b8060405250919050565b600067ffffffffffffffff821115611fed57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006120618261207e565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60006120c0826120c7565b9050919050565b60006120d2826120d9565b9050919050565b60006120e48261207e565b9050919050565b82818337600083830152505050565b60005b838110156121185780820151818401526020810190506120fd565b83811115612127576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b61215181612056565b811461215c57600080fd5b50565b61216881612068565b811461217357600080fd5b50565b61217f81612074565b811461218a57600080fd5b50565b6121968161209e565b81146121a157600080fd5b50565b6121ad816120a8565b81146121b857600080fd5b5056fea2646970667358221220aa318d65d6c4f77d4fefa97a923d70b6d3116149ef30c2c5fd730238f491e5fc64736f6c63430006040033"

// DeployERC721Handler deploys a new Ethereum contract, binding an instance of ERC721Handler to it.
func DeployERC721Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address) (common.Address, *types.Transaction, *ERC721Handler, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721HandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721HandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses)
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
// Solidity: function _balances(address ) view returns(uint256)
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
// Solidity: function _balances(address ) view returns(uint256)
func (_ERC721Handler *ERC721HandlerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC721Handler.Contract.Balances(&_ERC721Handler.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) view returns(uint256)
func (_ERC721Handler *ERC721HandlerCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC721Handler.Contract.Balances(&_ERC721Handler.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
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
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC721Handler.Contract.BridgeAddress(&_ERC721Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC721Handler.Contract.BridgeAddress(&_ERC721Handler.CallOpts)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Handler.contract.Call(opts, out, "_contractWhitelist", arg0)
	return *ret0, err
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.ContractWhitelist(&_ERC721Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.ContractWhitelist(&_ERC721Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) view returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
func (_ERC721Handler *ERC721HandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             uint8
	ResourceID                     [32]byte
	LenDestinationRecipientAddress *big.Int
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	TokenID                        *big.Int
	MetaData                       []byte
}, error) {
	ret := new(struct {
		OriginChainTokenAddress        common.Address
		DestinationChainID             uint8
		ResourceID                     [32]byte
		LenDestinationRecipientAddress *big.Int
		DestinationRecipientAddress    []byte
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
// Solidity: function _depositRecords(uint256 ) view returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
func (_ERC721Handler *ERC721HandlerSession) DepositRecords(arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             uint8
	ResourceID                     [32]byte
	LenDestinationRecipientAddress *big.Int
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	TokenID                        *big.Int
	MetaData                       []byte
}, error) {
	return _ERC721Handler.Contract.DepositRecords(&_ERC721Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) view returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
func (_ERC721Handler *ERC721HandlerCallerSession) DepositRecords(arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             uint8
	ResourceID                     [32]byte
	LenDestinationRecipientAddress *big.Int
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	TokenID                        *big.Int
	MetaData                       []byte
}, error) {
	return _ERC721Handler.Contract.DepositRecords(&_ERC721Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Handler.contract.Call(opts, out, "_resourceIDToTokenContractAddress", arg0)
	return *ret0, err
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC721Handler.Contract.ResourceIDToTokenContractAddress(&_ERC721Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC721Handler.Contract.ResourceIDToTokenContractAddress(&_ERC721Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ERC721Handler.contract.Call(opts, out, "_tokenContractAddressToResourceID", arg0)
	return *ret0, err
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC721Handler.Contract.TokenContractAddressToResourceID(&_ERC721Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC721Handler.Contract.TokenContractAddressToResourceID(&_ERC721Handler.CallOpts, arg0)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) view returns(ERC721HandlerDepositRecord)
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
// Solidity: function getDepositRecord(uint256 depositID) view returns(ERC721HandlerDepositRecord)
func (_ERC721Handler *ERC721HandlerSession) GetDepositRecord(depositID *big.Int) (ERC721HandlerDepositRecord, error) {
	return _ERC721Handler.Contract.GetDepositRecord(&_ERC721Handler.CallOpts, depositID)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) view returns(ERC721HandlerDepositRecord)
func (_ERC721Handler *ERC721HandlerCallerSession) GetDepositRecord(depositID *big.Int) (ERC721HandlerDepositRecord, error) {
	return _ERC721Handler.Contract.GetDepositRecord(&_ERC721Handler.CallOpts, depositID)
}

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactor) Deposit(opts *bind.TransactOpts, destinationChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "deposit", destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerSession) Deposit(destinationChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) Deposit(destinationChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
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

// FundERC721 is a paid mutator transaction binding the contract method 0x73542980.
//
// Solidity: function fundERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerTransactor) FundERC721(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "fundERC721", tokenAddress, owner, tokenID)
}

// FundERC721 is a paid mutator transaction binding the contract method 0x73542980.
//
// Solidity: function fundERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerSession) FundERC721(tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.Contract.FundERC721(&_ERC721Handler.TransactOpts, tokenAddress, owner, tokenID)
}

// FundERC721 is a paid mutator transaction binding the contract method 0x73542980.
//
// Solidity: function fundERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) FundERC721(tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.Contract.FundERC721(&_ERC721Handler.TransactOpts, tokenAddress, owner, tokenID)
}

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetResourceIDAndContractAddress(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setResourceIDAndContractAddress", resourceID, contractAddress)
}

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerSession) SetResourceIDAndContractAddress(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetResourceIDAndContractAddress(&_ERC721Handler.TransactOpts, resourceID, contractAddress)
}

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetResourceIDAndContractAddress(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetResourceIDAndContractAddress(&_ERC721Handler.TransactOpts, resourceID, contractAddress)
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

var RuntimeBytecode = "0x608060405234801561001057600080fd5b50600436106100b45760003560e01c80638a025ce8116100715780638a025ce81461019f578063c8ba6c87146101bb578063d9caed12146101eb578063db95e75c14610207578063e245386f14610237578063fc9539cd1461026e576100b4565b80630a6d55d8146100b9578063318c136e146100e957806345a104db146101075780636ebcf6071461012357806373542980146101535780637f79bea81461016f575b600080fd5b6100d360048036038101906100ce919061174a565b61028a565b6040516100e09190611d38565b60405180910390f35b6100f16102bd565b6040516100fe9190611d38565b60405180910390f35b610121600480360381019061011c9190611842565b6102e3565b005b61013d600480360381019061013891906116a9565b6105eb565b60405161014a9190611f8e565b60405180910390f35b61016d600480360381019061016891906116d2565b610603565b005b610189600480360381019061018491906116a9565b610711565b6040516101969190611e54565b60405180910390f35b6101b960048036038101906101b49190611773565b610731565b005b6101d560048036038101906101d091906116a9565b6108b7565b6040516101e29190611e6f565b60405180910390f35b610205600480360381019061020091906116d2565b6108cf565b005b610221600480360381019061021c91906117f0565b610970565b60405161022e9190611f6c565b60405180910390f35b610251600480360381019061024c91906117f0565b610bc9565b604051610265989796959493929190611dc8565b60405180910390f35b610288600480360381019061028391906117af565b610d8e565b005b60036020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610373576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036a90611ecc565b60405180910390fd5b60008060006060806020860151945060408601519250604051915060608601519350836080018601518483016040016040528060200160e401360360e4843760405191508082016040016040528461010401803603818437505060006003600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905061040e8161105b565b61044d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044490611f4c565b60405180910390fd5b610459818930876110b1565b6040518061010001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018b60ff1681526020018781526020018681526020018481526020018973ffffffffffffffffffffffffffffffffffffffff16815260200185815260200183815250600260008b815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff1602179055506040820151816001015560608201518160020155608082015181600301908051906020019061056d9291906114bb565b5060a08201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816005015560e08201518160060190805190602001906105db9291906114bb565b5090505050505050505050505050565b60006020528060005260406000206000915090505481565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b815260040161064593929190611d91565b600060405180830381600087803b15801561065f57600080fd5b505af1158015610673573d6000803e3d6000fd5b505050506106c960016000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111c090919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60056020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166003600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146107d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ca90611f0c565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008060405160200161082a9190611d1d565b60405160208183030381529060405280519060200120826040516020016108519190611d1d565b60405160208183030381529060405280519060200120146108a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089e90611eac565b60405180910390fd5b6108b18484611215565b50505050565b60046020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461095f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095690611ecc565b60405180910390fd5b61096b83308484611307565b505050565b61097861153b565b60026000838152602001908152602001600020604051806101000160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ab75780601f10610a8c57610100808354040283529160200191610ab7565b820191906000526020600020905b815481529060010190602001808311610a9a57829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160058201548152602001600682018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bb95780601f10610b8e57610100808354040283529160200191610bb9565b820191906000526020600020905b815481529060010190602001808311610b9c57829003601f168201915b5050505050815250509050919050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff1690806001015490806002015490806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610cba5780601f10610c8f57610100808354040283529160200191610cba565b820191906000526020600020905b815481529060010190602001808311610c9d57829003601f168201915b5050505050908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806005015490806006018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d845780601f10610d5957610100808354040283529160200191610d84565b820191906000526020600020905b815481529060010190602001808311610d6757829003601f168201915b5050505050905088565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1590611ecc565b60405180910390fd5b600080606080602085015193506040850151925060405191506060850151806080018601518184016040016040528060200160840136036084853760405192508083016040016040528160a40180360381853750505060008060208401519150604b8701519050610e918160601c61105b565b610ed0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ec790611f2c565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663beab71316040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610f4157600080fd5b505af1158015610f55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f799190611819565b90508060ff1687601f60208110610f8c57fe5b1a60f81b60f81c60ff161415610fb357610fae8360601c308660601c8b611307565b611050565b60008360601c90508073ffffffffffffffffffffffffffffffffffffffff16638832e6e38660601c8b896040518463ffffffff1660e01b8152600401610ffb93929190611d53565b602060405180830381600087803b15801561101557600080fd5b505af1158015611029573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061104d9190611721565b50505b505050505050505050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b81526004016110f393929190611d91565b600060405180830381600087803b15801561110d57600080fd5b505af1158015611121573d6000803e3d6000fd5b5050505061117760016000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111c090919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b60008082840190508381101561120b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161120290611eec565b60405180910390fd5b8091505092915050565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b815260040161134993929190611d91565b600060405180830381600087803b15801561136357600080fd5b505af1158015611377573d6000803e3d6000fd5b505050506113cd60016000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461141690919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b600061145883836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611460565b905092915050565b60008383111582906114a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161149f9190611e8a565b60405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106114fc57805160ff191683800117855561152a565b8280016001018555821561152a579182015b8281111561152957825182559160200191906001019061150e565b5b50905061153791906115b2565b5090565b604051806101000160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600080191681526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001606081525090565b6115d491905b808211156115d05760008160009055506001016115b8565b5090565b90565b6000813590506115e681612148565b92915050565b6000815190506115fb8161215f565b92915050565b60008135905061161081612176565b92915050565b600082601f83011261162757600080fd5b813561163a61163582611fd6565b611fa9565b9150808252602083016020830185838301111561165657600080fd5b6116618382846120eb565b50505092915050565b6000813590506116798161218d565b92915050565b60008135905061168e816121a4565b92915050565b6000815190506116a3816121a4565b92915050565b6000602082840312156116bb57600080fd5b60006116c9848285016115d7565b91505092915050565b6000806000606084860312156116e757600080fd5b60006116f5868287016115d7565b9350506020611706868287016115d7565b92505060406117178682870161166a565b9150509250925092565b60006020828403121561173357600080fd5b6000611741848285016115ec565b91505092915050565b60006020828403121561175c57600080fd5b600061176a84828501611601565b91505092915050565b6000806040838503121561178657600080fd5b600061179485828601611601565b92505060206117a5858286016115d7565b9150509250929050565b6000602082840312156117c157600080fd5b600082013567ffffffffffffffff8111156117db57600080fd5b6117e784828501611616565b91505092915050565b60006020828403121561180257600080fd5b60006118108482850161166a565b91505092915050565b60006020828403121561182b57600080fd5b600061183984828501611694565b91505092915050565b6000806000806080858703121561185857600080fd5b60006118668782880161167f565b94505060206118778782880161166a565b9350506040611888878288016115d7565b925050606085013567ffffffffffffffff8111156118a557600080fd5b6118b187828801611616565b91505092959194509250565b6118c6816120b5565b82525050565b6118d581612056565b82525050565b6118e481612056565b82525050565b6118f381612068565b82525050565b61190281612074565b82525050565b61191181612074565b82525050565b61192861192382612074565b61212d565b82525050565b60006119398261200d565b6119438185612034565b93506119538185602086016120fa565b61195c81612137565b840191505092915050565b600061197282612002565b61197c8185612023565b935061198c8185602086016120fa565b61199581612137565b840191505092915050565b60006119ab82612002565b6119b58185612034565b93506119c58185602086016120fa565b6119ce81612137565b840191505092915050565b60006119e482612018565b6119ee8185612045565b93506119fe8185602086016120fa565b611a0781612137565b840191505092915050565b6000611a1f603583612045565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611a85601e83612045565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611ac5601b83612045565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000611b05603783612045565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b6000611b6b602883612045565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611bd1603383612045565b91507f70726f7669646564206f726967696e436861696e546f6b656e4164647265737360008301527f206973206e6f742077686974656c6973746564000000000000000000000000006020830152604082019050919050565b600061010083016000830151611c4360008601826118cc565b506020830151611c566020860182611cff565b506040830151611c6960408601826118f9565b506060830151611c7c6060860182611ce1565b5060808301518482036080860152611c948282611967565b91505060a0830151611ca960a08601826118cc565b5060c0830151611cbc60c0860182611ce1565b5060e083015184820360e0860152611cd48282611967565b9150508091505092915050565b611cea8161209e565b82525050565b611cf98161209e565b82525050565b611d08816120a8565b82525050565b611d17816120a8565b82525050565b6000611d298284611917565b60208201915081905092915050565b6000602082019050611d4d60008301846118db565b92915050565b6000606082019050611d6860008301866118bd565b611d756020830185611cf0565b8181036040830152611d87818461192e565b9050949350505050565b6000606082019050611da660008301866118db565b611db360208301856118db565b611dc06040830184611cf0565b949350505050565b600061010082019050611dde600083018b6118db565b611deb602083018a611d0e565b611df86040830189611908565b611e056060830188611cf0565b8181036080830152611e1781876119a0565b9050611e2660a08301866118db565b611e3360c0830185611cf0565b81810360e0830152611e4581846119a0565b90509998505050505050505050565b6000602082019050611e6960008301846118ea565b92915050565b6000602082019050611e846000830184611908565b92915050565b60006020820190508181036000830152611ea481846119d9565b905092915050565b60006020820190508181036000830152611ec581611a12565b9050919050565b60006020820190508181036000830152611ee581611a78565b9050919050565b60006020820190508181036000830152611f0581611ab8565b9050919050565b60006020820190508181036000830152611f2581611af8565b9050919050565b60006020820190508181036000830152611f4581611b5e565b9050919050565b60006020820190508181036000830152611f6581611bc4565b9050919050565b60006020820190508181036000830152611f868184611c2a565b905092915050565b6000602082019050611fa36000830184611cf0565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611fcc57600080fd5b8060405250919050565b600067ffffffffffffffff821115611fed57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006120618261207e565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60006120c0826120c7565b9050919050565b60006120d2826120d9565b9050919050565b60006120e48261207e565b9050919050565b82818337600083830152505050565b60005b838110156121185780820151818401526020810190506120fd565b83811115612127576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b61215181612056565b811461215c57600080fd5b50565b61216881612068565b811461217357600080fd5b50565b61217f81612074565b811461218a57600080fd5b50565b6121968161209e565b81146121a157600080fd5b50565b6121ad816120a8565b81146121b857600080fd5b5056fea2646970667358221220aa318d65d6c4f77d4fefa97a923d70b6d3116149ef30c2c5fd730238f491e5fc64736f6c63430006040033"
