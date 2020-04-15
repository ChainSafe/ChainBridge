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
const ERC721HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structERC721Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResourceIDAndContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721HandlerBin is the compiled bytecode used for deploying new contracts.
var ERC721HandlerBin = "0x60806040523480156200001157600080fd5b50604051620026623803806200266283398181016040528101906200003791906200035b565b80518251146200007e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000759062000471565b60405180910390fd5b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b8251811015620001175762000109838281518110620000e057fe5b6020026020010151838381518110620000f557fe5b60200260200101516200012160201b60201c565b8080600101915050620000c5565b5050505062000596565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600081519050620002248162000562565b92915050565b600082601f8301126200023c57600080fd5b8151620002536200024d82620004c1565b62000493565b915081818352602084019350602081019050838560208402820111156200027957600080fd5b60005b83811015620002ad578162000292888262000213565b8452602084019350602083019250506001810190506200027c565b5050505092915050565b600082601f830112620002c957600080fd5b8151620002e0620002da82620004ea565b62000493565b915081818352602084019350602081019050838560208402820111156200030657600080fd5b60005b838110156200033a57816200031f888262000344565b84526020840193506020830192505060018101905062000309565b5050505092915050565b60008151905062000355816200057c565b92915050565b6000806000606084860312156200037157600080fd5b6000620003818682870162000213565b935050602084015167ffffffffffffffff8111156200039f57600080fd5b620003ad86828701620002b7565b925050604084015167ffffffffffffffff811115620003cb57600080fd5b620003d9868287016200022a565b9150509250925092565b6000620003f260478362000513565b91507f6d69736d61746368206c656e677468206265747765656e20696e697469616c5260008301527f65736f7572636549447320616e6420696e697469616c436f6e7472616374416460208301527f64726573736573000000000000000000000000000000000000000000000000006040830152606082019050919050565b600060208201905081810360008301526200048c81620003e3565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715620004b757600080fd5b8060405250919050565b600067ffffffffffffffff821115620004d957600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200050257600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620005318262000542565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200056d8162000524565b81146200057957600080fd5b50565b620005878162000538565b81146200059357600080fd5b50565b6120bc80620005a66000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638a025ce8116100715780638a025ce814610178578063c8ba6c8714610194578063d9caed12146101c4578063db95e75c146101e0578063e245386f14610210578063fc9539cd14610247576100a9565b80630a6d55d8146100ae578063318c136e146100de57806345a104db146100fc5780636ebcf607146101185780637f79bea814610148575b600080fd5b6100c860048036038101906100c39190611615565b610263565b6040516100d59190611c03565b60405180910390f35b6100e6610296565b6040516100f39190611c03565b60405180910390f35b6101166004803603810190610111919061170d565b6102bc565b005b610132600480360381019061012d9190611574565b6105c4565b60405161013f9190611e59565b60405180910390f35b610162600480360381019061015d9190611574565b6105dc565b60405161016f9190611d1f565b60405180910390f35b610192600480360381019061018d919061163e565b6105fc565b005b6101ae60048036038101906101a99190611574565b610782565b6040516101bb9190611d3a565b60405180910390f35b6101de60048036038101906101d9919061159d565b61079a565b005b6101fa60048036038101906101f591906116bb565b61083b565b6040516102079190611e37565b60405180910390f35b61022a600480360381019061022591906116bb565b610a94565b60405161023e989796959493929190611c93565b60405180910390f35b610261600480360381019061025c919061167a565b610c59565b005b60036020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461034c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034390611d97565b60405180910390fd5b60008060006060806020860151945060408601519250604051915060608601519350836080018601518483016040016040528060200160e401360360e4843760405191508082016040016040528461010401803603818437505060006003600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506103e781610f26565b610426576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041d90611e17565b60405180910390fd5b61043281893087610f7c565b6040518061010001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018b60ff1681526020018781526020018681526020018481526020018973ffffffffffffffffffffffffffffffffffffffff16815260200185815260200183815250600260008b815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff16021790555060408201518160010155606082015181600201556080820151816003019080519060200190610546929190611386565b5060a08201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816005015560e08201518160060190805190602001906105b4929190611386565b5090505050505050505050505050565b60006020528060005260406000206000915090505481565b60056020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166003600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461069e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069590611dd7565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016106f59190611be8565b604051602081830303815290604052805190602001208260405160200161071c9190611be8565b6040516020818303038152906040528051906020012014610772576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076990611d77565b60405180910390fd5b61077c848461108b565b50505050565b60046020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461082a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082190611d97565b60405180910390fd5b6108368330848461117d565b505050565b610843611406565b60026000838152602001908152602001600020604051806101000160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109825780601f1061095757610100808354040283529160200191610982565b820191906000526020600020905b81548152906001019060200180831161096557829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160058201548152602001600682018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a845780601f10610a5957610100808354040283529160200191610a84565b820191906000526020600020905b815481529060010190602001808311610a6757829003601f168201915b5050505050815250509050919050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff1690806001015490806002015490806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b855780601f10610b5a57610100808354040283529160200191610b85565b820191906000526020600020905b815481529060010190602001808311610b6857829003601f168201915b5050505050908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806005015490806006018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c4f5780601f10610c2457610100808354040283529160200191610c4f565b820191906000526020600020905b815481529060010190602001808311610c3257829003601f168201915b5050505050905088565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ce9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ce090611d97565b60405180910390fd5b600080606080602085015193506040850151925060405191506060850151806080018601518184016040016040528060200160840136036084853760405192508083016040016040528160a40180360381853750505060008060208401519150604b8701519050610d5c8160601c610f26565b610d9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d9290611df7565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663beab71316040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610e0c57600080fd5b505af1158015610e20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e4491906116e4565b90508060ff1687601f60208110610e5757fe5b1a60f81b60f81c60ff161415610e7e57610e798360601c308660601c8b61117d565b610f1b565b60008360601c90508073ffffffffffffffffffffffffffffffffffffffff16638832e6e38660601c8b896040518463ffffffff1660e01b8152600401610ec693929190611c1e565b602060405180830381600087803b158015610ee057600080fd5b505af1158015610ef4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f1891906115ec565b50505b505050505050505050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610fbe93929190611c5c565b600060405180830381600087803b158015610fd857600080fd5b505af1158015610fec573d6000803e3d6000fd5b5050505061104260016000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461128c90919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b81526004016111bf93929190611c5c565b600060405180830381600087803b1580156111d957600080fd5b505af11580156111ed573d6000803e3d6000fd5b5050505061124360016000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546112e190919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b6000808284019050838110156112d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112ce90611db7565b60405180910390fd5b8091505092915050565b600061132383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061132b565b905092915050565b6000838311158290611373576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161136a9190611d55565b60405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106113c757805160ff19168380011785556113f5565b828001600101855582156113f5579182015b828111156113f45782518255916020019190600101906113d9565b5b509050611402919061147d565b5090565b604051806101000160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600080191681526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001606081525090565b61149f91905b8082111561149b576000816000905550600101611483565b5090565b90565b6000813590506114b181612013565b92915050565b6000815190506114c68161202a565b92915050565b6000813590506114db81612041565b92915050565b600082601f8301126114f257600080fd5b813561150561150082611ea1565b611e74565b9150808252602083016020830185838301111561152157600080fd5b61152c838284611fb6565b50505092915050565b60008135905061154481612058565b92915050565b6000813590506115598161206f565b92915050565b60008151905061156e8161206f565b92915050565b60006020828403121561158657600080fd5b6000611594848285016114a2565b91505092915050565b6000806000606084860312156115b257600080fd5b60006115c0868287016114a2565b93505060206115d1868287016114a2565b92505060406115e286828701611535565b9150509250925092565b6000602082840312156115fe57600080fd5b600061160c848285016114b7565b91505092915050565b60006020828403121561162757600080fd5b6000611635848285016114cc565b91505092915050565b6000806040838503121561165157600080fd5b600061165f858286016114cc565b9250506020611670858286016114a2565b9150509250929050565b60006020828403121561168c57600080fd5b600082013567ffffffffffffffff8111156116a657600080fd5b6116b2848285016114e1565b91505092915050565b6000602082840312156116cd57600080fd5b60006116db84828501611535565b91505092915050565b6000602082840312156116f657600080fd5b60006117048482850161155f565b91505092915050565b6000806000806080858703121561172357600080fd5b60006117318782880161154a565b945050602061174287828801611535565b9350506040611753878288016114a2565b925050606085013567ffffffffffffffff81111561177057600080fd5b61177c878288016114e1565b91505092959194509250565b61179181611f80565b82525050565b6117a081611f21565b82525050565b6117af81611f21565b82525050565b6117be81611f33565b82525050565b6117cd81611f3f565b82525050565b6117dc81611f3f565b82525050565b6117f36117ee82611f3f565b611ff8565b82525050565b600061180482611ed8565b61180e8185611eff565b935061181e818560208601611fc5565b61182781612002565b840191505092915050565b600061183d82611ecd565b6118478185611eee565b9350611857818560208601611fc5565b61186081612002565b840191505092915050565b600061187682611ecd565b6118808185611eff565b9350611890818560208601611fc5565b61189981612002565b840191505092915050565b60006118af82611ee3565b6118b98185611f10565b93506118c9818560208601611fc5565b6118d281612002565b840191505092915050565b60006118ea603583611f10565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611950601e83611f10565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611990601b83611f10565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b60006119d0603783611f10565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b6000611a36602883611f10565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611a9c603383611f10565b91507f70726f7669646564206f726967696e436861696e546f6b656e4164647265737360008301527f206973206e6f742077686974656c6973746564000000000000000000000000006020830152604082019050919050565b600061010083016000830151611b0e6000860182611797565b506020830151611b216020860182611bca565b506040830151611b3460408601826117c4565b506060830151611b476060860182611bac565b5060808301518482036080860152611b5f8282611832565b91505060a0830151611b7460a0860182611797565b5060c0830151611b8760c0860182611bac565b5060e083015184820360e0860152611b9f8282611832565b9150508091505092915050565b611bb581611f69565b82525050565b611bc481611f69565b82525050565b611bd381611f73565b82525050565b611be281611f73565b82525050565b6000611bf482846117e2565b60208201915081905092915050565b6000602082019050611c1860008301846117a6565b92915050565b6000606082019050611c336000830186611788565b611c406020830185611bbb565b8181036040830152611c5281846117f9565b9050949350505050565b6000606082019050611c7160008301866117a6565b611c7e60208301856117a6565b611c8b6040830184611bbb565b949350505050565b600061010082019050611ca9600083018b6117a6565b611cb6602083018a611bd9565b611cc360408301896117d3565b611cd06060830188611bbb565b8181036080830152611ce2818761186b565b9050611cf160a08301866117a6565b611cfe60c0830185611bbb565b81810360e0830152611d10818461186b565b90509998505050505050505050565b6000602082019050611d3460008301846117b5565b92915050565b6000602082019050611d4f60008301846117d3565b92915050565b60006020820190508181036000830152611d6f81846118a4565b905092915050565b60006020820190508181036000830152611d90816118dd565b9050919050565b60006020820190508181036000830152611db081611943565b9050919050565b60006020820190508181036000830152611dd081611983565b9050919050565b60006020820190508181036000830152611df0816119c3565b9050919050565b60006020820190508181036000830152611e1081611a29565b9050919050565b60006020820190508181036000830152611e3081611a8f565b9050919050565b60006020820190508181036000830152611e518184611af5565b905092915050565b6000602082019050611e6e6000830184611bbb565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611e9757600080fd5b8060405250919050565b600067ffffffffffffffff821115611eb857600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611f2c82611f49565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000611f8b82611f92565b9050919050565b6000611f9d82611fa4565b9050919050565b6000611faf82611f49565b9050919050565b82818337600083830152505050565b60005b83811015611fe3578082015181840152602081019050611fc8565b83811115611ff2576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b61201c81611f21565b811461202757600080fd5b50565b61203381611f33565b811461203e57600080fd5b50565b61204a81611f3f565b811461205557600080fd5b50565b61206181611f69565b811461206c57600080fd5b50565b61207881611f73565b811461208357600080fd5b5056fea26469706673582212208afce8a20d2206ae31a1a5c130b0291640b7de2008fada6f8246e98f309b15ec64736f6c63430006040033"

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

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) constant returns(bool)
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
// Solidity: function _contractWhitelist(address ) constant returns(bool)
func (_ERC721Handler *ERC721HandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.ContractWhitelist(&_ERC721Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) constant returns(bool)
func (_ERC721Handler *ERC721HandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.ContractWhitelist(&_ERC721Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) constant returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
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
// Solidity: function _depositRecords(uint256 ) constant returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
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
// Solidity: function _depositRecords(uint256 ) constant returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
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
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) constant returns(address)
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
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) constant returns(address)
func (_ERC721Handler *ERC721HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC721Handler.Contract.ResourceIDToTokenContractAddress(&_ERC721Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) constant returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC721Handler.Contract.ResourceIDToTokenContractAddress(&_ERC721Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) constant returns(bytes32)
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
// Solidity: function _tokenContractAddressToResourceID(address ) constant returns(bytes32)
func (_ERC721Handler *ERC721HandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC721Handler.Contract.TokenContractAddressToResourceID(&_ERC721Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) constant returns(bytes32)
func (_ERC721Handler *ERC721HandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC721Handler.Contract.TokenContractAddressToResourceID(&_ERC721Handler.CallOpts, arg0)
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

var RuntimeBytecode = "0x608060405234801561001057600080fd5b50600436106100a95760003560e01c80638a025ce8116100715780638a025ce814610178578063c8ba6c8714610194578063d9caed12146101c4578063db95e75c146101e0578063e245386f14610210578063fc9539cd14610247576100a9565b80630a6d55d8146100ae578063318c136e146100de57806345a104db146100fc5780636ebcf607146101185780637f79bea814610148575b600080fd5b6100c860048036038101906100c39190611615565b610263565b6040516100d59190611c03565b60405180910390f35b6100e6610296565b6040516100f39190611c03565b60405180910390f35b6101166004803603810190610111919061170d565b6102bc565b005b610132600480360381019061012d9190611574565b6105c4565b60405161013f9190611e59565b60405180910390f35b610162600480360381019061015d9190611574565b6105dc565b60405161016f9190611d1f565b60405180910390f35b610192600480360381019061018d919061163e565b6105fc565b005b6101ae60048036038101906101a99190611574565b610782565b6040516101bb9190611d3a565b60405180910390f35b6101de60048036038101906101d9919061159d565b61079a565b005b6101fa60048036038101906101f591906116bb565b61083b565b6040516102079190611e37565b60405180910390f35b61022a600480360381019061022591906116bb565b610a94565b60405161023e989796959493929190611c93565b60405180910390f35b610261600480360381019061025c919061167a565b610c59565b005b60036020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461034c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034390611d97565b60405180910390fd5b60008060006060806020860151945060408601519250604051915060608601519350836080018601518483016040016040528060200160e401360360e4843760405191508082016040016040528461010401803603818437505060006003600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506103e781610f26565b610426576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041d90611e17565b60405180910390fd5b61043281893087610f7c565b6040518061010001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018b60ff1681526020018781526020018681526020018481526020018973ffffffffffffffffffffffffffffffffffffffff16815260200185815260200183815250600260008b815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff16021790555060408201518160010155606082015181600201556080820151816003019080519060200190610546929190611386565b5060a08201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816005015560e08201518160060190805190602001906105b4929190611386565b5090505050505050505050505050565b60006020528060005260406000206000915090505481565b60056020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166003600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461069e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069590611dd7565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016106f59190611be8565b604051602081830303815290604052805190602001208260405160200161071c9190611be8565b6040516020818303038152906040528051906020012014610772576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076990611d77565b60405180910390fd5b61077c848461108b565b50505050565b60046020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461082a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082190611d97565b60405180910390fd5b6108368330848461117d565b505050565b610843611406565b60026000838152602001908152602001600020604051806101000160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109825780601f1061095757610100808354040283529160200191610982565b820191906000526020600020905b81548152906001019060200180831161096557829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160058201548152602001600682018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a845780601f10610a5957610100808354040283529160200191610a84565b820191906000526020600020905b815481529060010190602001808311610a6757829003601f168201915b5050505050815250509050919050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff1690806001015490806002015490806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b855780601f10610b5a57610100808354040283529160200191610b85565b820191906000526020600020905b815481529060010190602001808311610b6857829003601f168201915b5050505050908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806005015490806006018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c4f5780601f10610c2457610100808354040283529160200191610c4f565b820191906000526020600020905b815481529060010190602001808311610c3257829003601f168201915b5050505050905088565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ce9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ce090611d97565b60405180910390fd5b600080606080602085015193506040850151925060405191506060850151806080018601518184016040016040528060200160840136036084853760405192508083016040016040528160a40180360381853750505060008060208401519150604b8701519050610d5c8160601c610f26565b610d9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d9290611df7565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663beab71316040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610e0c57600080fd5b505af1158015610e20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e4491906116e4565b90508060ff1687601f60208110610e5757fe5b1a60f81b60f81c60ff161415610e7e57610e798360601c308660601c8b61117d565b610f1b565b60008360601c90508073ffffffffffffffffffffffffffffffffffffffff16638832e6e38660601c8b896040518463ffffffff1660e01b8152600401610ec693929190611c1e565b602060405180830381600087803b158015610ee057600080fd5b505af1158015610ef4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f1891906115ec565b50505b505050505050505050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610fbe93929190611c5c565b600060405180830381600087803b158015610fd857600080fd5b505af1158015610fec573d6000803e3d6000fd5b5050505061104260016000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461128c90919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b81526004016111bf93929190611c5c565b600060405180830381600087803b1580156111d957600080fd5b505af11580156111ed573d6000803e3d6000fd5b5050505061124360016000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546112e190919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b6000808284019050838110156112d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112ce90611db7565b60405180910390fd5b8091505092915050565b600061132383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061132b565b905092915050565b6000838311158290611373576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161136a9190611d55565b60405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106113c757805160ff19168380011785556113f5565b828001600101855582156113f5579182015b828111156113f45782518255916020019190600101906113d9565b5b509050611402919061147d565b5090565b604051806101000160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600080191681526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001606081525090565b61149f91905b8082111561149b576000816000905550600101611483565b5090565b90565b6000813590506114b181612013565b92915050565b6000815190506114c68161202a565b92915050565b6000813590506114db81612041565b92915050565b600082601f8301126114f257600080fd5b813561150561150082611ea1565b611e74565b9150808252602083016020830185838301111561152157600080fd5b61152c838284611fb6565b50505092915050565b60008135905061154481612058565b92915050565b6000813590506115598161206f565b92915050565b60008151905061156e8161206f565b92915050565b60006020828403121561158657600080fd5b6000611594848285016114a2565b91505092915050565b6000806000606084860312156115b257600080fd5b60006115c0868287016114a2565b93505060206115d1868287016114a2565b92505060406115e286828701611535565b9150509250925092565b6000602082840312156115fe57600080fd5b600061160c848285016114b7565b91505092915050565b60006020828403121561162757600080fd5b6000611635848285016114cc565b91505092915050565b6000806040838503121561165157600080fd5b600061165f858286016114cc565b9250506020611670858286016114a2565b9150509250929050565b60006020828403121561168c57600080fd5b600082013567ffffffffffffffff8111156116a657600080fd5b6116b2848285016114e1565b91505092915050565b6000602082840312156116cd57600080fd5b60006116db84828501611535565b91505092915050565b6000602082840312156116f657600080fd5b60006117048482850161155f565b91505092915050565b6000806000806080858703121561172357600080fd5b60006117318782880161154a565b945050602061174287828801611535565b9350506040611753878288016114a2565b925050606085013567ffffffffffffffff81111561177057600080fd5b61177c878288016114e1565b91505092959194509250565b61179181611f80565b82525050565b6117a081611f21565b82525050565b6117af81611f21565b82525050565b6117be81611f33565b82525050565b6117cd81611f3f565b82525050565b6117dc81611f3f565b82525050565b6117f36117ee82611f3f565b611ff8565b82525050565b600061180482611ed8565b61180e8185611eff565b935061181e818560208601611fc5565b61182781612002565b840191505092915050565b600061183d82611ecd565b6118478185611eee565b9350611857818560208601611fc5565b61186081612002565b840191505092915050565b600061187682611ecd565b6118808185611eff565b9350611890818560208601611fc5565b61189981612002565b840191505092915050565b60006118af82611ee3565b6118b98185611f10565b93506118c9818560208601611fc5565b6118d281612002565b840191505092915050565b60006118ea603583611f10565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611950601e83611f10565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611990601b83611f10565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b60006119d0603783611f10565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b6000611a36602883611f10565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611a9c603383611f10565b91507f70726f7669646564206f726967696e436861696e546f6b656e4164647265737360008301527f206973206e6f742077686974656c6973746564000000000000000000000000006020830152604082019050919050565b600061010083016000830151611b0e6000860182611797565b506020830151611b216020860182611bca565b506040830151611b3460408601826117c4565b506060830151611b476060860182611bac565b5060808301518482036080860152611b5f8282611832565b91505060a0830151611b7460a0860182611797565b5060c0830151611b8760c0860182611bac565b5060e083015184820360e0860152611b9f8282611832565b9150508091505092915050565b611bb581611f69565b82525050565b611bc481611f69565b82525050565b611bd381611f73565b82525050565b611be281611f73565b82525050565b6000611bf482846117e2565b60208201915081905092915050565b6000602082019050611c1860008301846117a6565b92915050565b6000606082019050611c336000830186611788565b611c406020830185611bbb565b8181036040830152611c5281846117f9565b9050949350505050565b6000606082019050611c7160008301866117a6565b611c7e60208301856117a6565b611c8b6040830184611bbb565b949350505050565b600061010082019050611ca9600083018b6117a6565b611cb6602083018a611bd9565b611cc360408301896117d3565b611cd06060830188611bbb565b8181036080830152611ce2818761186b565b9050611cf160a08301866117a6565b611cfe60c0830185611bbb565b81810360e0830152611d10818461186b565b90509998505050505050505050565b6000602082019050611d3460008301846117b5565b92915050565b6000602082019050611d4f60008301846117d3565b92915050565b60006020820190508181036000830152611d6f81846118a4565b905092915050565b60006020820190508181036000830152611d90816118dd565b9050919050565b60006020820190508181036000830152611db081611943565b9050919050565b60006020820190508181036000830152611dd081611983565b9050919050565b60006020820190508181036000830152611df0816119c3565b9050919050565b60006020820190508181036000830152611e1081611a29565b9050919050565b60006020820190508181036000830152611e3081611a8f565b9050919050565b60006020820190508181036000830152611e518184611af5565b905092915050565b6000602082019050611e6e6000830184611bbb565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611e9757600080fd5b8060405250919050565b600067ffffffffffffffff821115611eb857600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611f2c82611f49565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000611f8b82611f92565b9050919050565b6000611f9d82611fa4565b9050919050565b6000611faf82611f49565b9050919050565b82818337600083830152505050565b60005b83811015611fe3578082015181840152602081019050611fc8565b83811115611ff2576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b61201c81611f21565b811461202757600080fd5b50565b61203381611f33565b811461203e57600080fd5b50565b61204a81611f3f565b811461205557600080fd5b50565b61206181611f69565b811461206c57600080fd5b50565b61207881611f73565b811461208357600080fd5b5056fea26469706673582212208afce8a20d2206ae31a1a5c130b0291640b7de2008fada6f8246e98f309b15ec64736f6c63430006040033"
