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
const ERC20HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"burnableContractAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_useContractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"internalType\":\"structERC20Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResourceIDAndContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20HandlerBin is the compiled bytecode used for deploying new contracts.
var ERC20HandlerBin = "0x60806040523480156200001157600080fd5b5060405162002c3238038062002c328339818101604052810190620000379190620004a4565b81518351146200007e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000759062000673565b60405180910390fd5b83600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b8351811015620001175762000109848281518110620000e057fe5b6020026020010151848381518110620000f557fe5b60200260200101516200016660201b60201c565b8080600101915050620000c5565b5060008090505b81518110156200015b576200014d8282815181106200013957fe5b60200260200101516200025860201b60201c565b80806001019150506200011e565b505050505062000798565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b62000269816200030660201b60201c565b620002ab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002a29062000651565b60405180910390fd5b6001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000815190506200036d8162000764565b92915050565b600082601f8301126200038557600080fd5b81516200039c6200039682620006c3565b62000695565b91508181835260208401935060208101905083856020840282011115620003c257600080fd5b60005b83811015620003f65781620003db88826200035c565b845260208401935060208301925050600181019050620003c5565b5050505092915050565b600082601f8301126200041257600080fd5b8151620004296200042382620006ec565b62000695565b915081818352602084019350602081019050838560208402820111156200044f57600080fd5b60005b838110156200048357816200046888826200048d565b84526020840193506020830192505060018101905062000452565b5050505092915050565b6000815190506200049e816200077e565b92915050565b60008060008060808587031215620004bb57600080fd5b6000620004cb878288016200035c565b945050602085015167ffffffffffffffff811115620004e957600080fd5b620004f78782880162000400565b935050604085015167ffffffffffffffff8111156200051557600080fd5b620005238782880162000373565b925050606085015167ffffffffffffffff8111156200054157600080fd5b6200054f8782880162000373565b91505092959194509250565b60006200056a60248362000715565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000620005d260478362000715565b91507f6d69736d61746368206c656e677468206265747765656e20696e697469616c5260008301527f65736f7572636549447320616e6420696e697469616c436f6e7472616374416460208301527f64726573736573000000000000000000000000000000000000000000000000006040830152606082019050919050565b600060208201905081810360008301526200066c816200055b565b9050919050565b600060208201905081810360008301526200068e81620005c3565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715620006b957600080fd5b8060405250919050565b600067ffffffffffffffff821115620006db57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200070457600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620007338262000744565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200076f8162000726565b81146200077b57600080fd5b50565b62000789816200073a565b81146200079557600080fd5b50565b61248a80620007a86000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80636ebcf607116100a2578063c8ba6c8711610071578063c8ba6c87146102dc578063d9caed121461030c578063db95e75c14610328578063e245386f14610358578063fc9539cd1461038e5761010b565b80636ebcf607146102445780637f79bea8146102745780638a025ce8146102a457806395601f09146102c05761010b565b80633af32abf116100de5780633af32abf1461019857806345274738146101c857806345a104db146101f85780636a70d081146102145761010b565b806307b7ed99146101105780630a6d55d81461012c5780632521ab411461015c578063318c136e1461017a575b600080fd5b61012a6004803603810190610125919061198b565b6103aa565b005b61014660048036038101906101419190611a2c565b61044d565b604051610153919061201d565b60405180910390f35b610164610480565b604051610171919061210e565b60405180910390f35b610182610493565b60405161018f919061201d565b60405180910390f35b6101b260048036038101906101ad919061198b565b6104b9565b6040516101bf919061210e565b60405180910390f35b6101e260048036038101906101dd919061198b565b61050f565b6040516101ef9190612268565b60405180910390f35b610212600480360381019061020d9190611b24565b610527565b005b61022e6004803603810190610229919061198b565b610845565b60405161023b919061210e565b60405180910390f35b61025e6004803603810190610259919061198b565b610865565b60405161026b9190612268565b60405180910390f35b61028e6004803603810190610289919061198b565b61087d565b60405161029b919061210e565b60405180910390f35b6102be60048036038101906102b99190611a55565b61089d565b005b6102da60048036038101906102d591906119b4565b610a23565b005b6102f660048036038101906102f1919061198b565b610b51565b6040516103039190612129565b60405180910390f35b610326600480360381019061032191906119b4565b610b69565b005b610342600480360381019061033d9190611ad2565b610c09565b60405161034f9190612246565b60405180910390f35b610372600480360381019061036d9190611ad2565b610dbf565b6040516103859796959493929190612098565b60405180910390f35b6103a860048036038101906103a39190611a91565b610ee6565b005b6103b3816104b9565b6103f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e9906121a6565b60405180910390fd5b6001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60036020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260149054906101000a900460ff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60016020528060005260406000206000915090505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ae90612186565b60405180910390fd5b60006060600080602085015193506040850151915060405192508460600151905080830160200160405260e4360360e4843760006003600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905061062a816104b9565b610669576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066090612226565b60405180910390fd5b600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156106cb576106c6818885611125565b6106d8565b6106d781883086611232565b5b6040518060e001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018a60ff1681526020018681526020018381526020018581526020018873ffffffffffffffffffffffffffffffffffffffff16815260200184815250600760008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff160217905550604082015181600101556060820151816002015560808201518160030190805190602001906107e59291906117a5565b5060a08201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160050155905050505050505050505050565b60066020528060005260406000206000915054906101000a900460ff1681565b60006020528060005260406000206000915090505481565b60056020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166003600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461093f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610936906121e6565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016109969190612002565b60405160208183030381529060405280519060200120826040516020016109bd9190612002565b6040516020818303038152906040528051906020012014610a13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0a90612166565b60405180910390fd5b610a1d8484611361565b50505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b8152600401610a6593929190612038565b602060405180830381600087803b158015610a7f57600080fd5b505af1158015610a93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab79190611a03565b50610b09826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461145390919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60046020528060005260406000206000915090505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bf9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bf090612186565b60405180910390fd5b610c048383836114a8565b505050565b610c11611825565b600760008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d4f5780601f10610d2457610100808354040283529160200191610d4f565b820191906000526020600020905b815481529060010190602001808311610d3257829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b60076020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff1690806001015490806002015490806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610eb05780601f10610e8557610100808354040283529160200191610eb0565b820191906000526020600020905b815481529060010190602001808311610e9357829003601f168201915b5050505050908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154905087565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6d90612186565b60405180910390fd5b60008060606020840151925060408401519150604051905083606001518082016020016040526084360360848337506000806003600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060208301519150610fee816104b9565b61102d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102490612206565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663beab71316040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561109e57600080fd5b505af11580156110b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d69190611afb565b90508060ff1686601f602081106110e957fe5b1a60f81b60f81c60ff16141561110c57611107838560601c896114a8565b61111b565b61111a838560601c896115d4565b5b5050505050505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166379cc679084846040518363ffffffff1660e01b815260040161116592919061206f565b600060405180830381600087803b15801561117f57600080fd5b505af1158015611193573d6000803e3d6000fd5b505050506111e982600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461145390919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b815260040161127493929190612038565b602060405180830381600087803b15801561128e57600080fd5b505af11580156112a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112c69190611a03565b50611318826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461145390919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008082840190508381101561149e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611495906121c6565b60405180910390fd5b8091505092915050565b60008390508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b81526004016114e892919061206f565b602060405180830381600087803b15801561150257600080fd5b505af1158015611516573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061153a9190611a03565b5061158c826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461170090919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b815260040161161492919061206f565b602060405180830381600087803b15801561162e57600080fd5b505af1158015611642573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116669190611a03565b506116b8826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461145390919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b600061174283836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061174a565b905092915050565b6000838311158290611792576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117899190612144565b60405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106117e657805160ff1916838001178555611814565b82800160010185558215611814579182015b828111156118135782518255916020019190600101906117f8565b5b5090506118219190611894565b5090565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600080191681526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b6118b691905b808211156118b257600081600090555060010161189a565b5090565b90565b6000813590506118c8816123e1565b92915050565b6000815190506118dd816123f8565b92915050565b6000813590506118f28161240f565b92915050565b600082601f83011261190957600080fd5b813561191c611917826122b0565b612283565b9150808252602083016020830185838301111561193857600080fd5b611943838284612384565b50505092915050565b60008135905061195b81612426565b92915050565b6000813590506119708161243d565b92915050565b6000815190506119858161243d565b92915050565b60006020828403121561199d57600080fd5b60006119ab848285016118b9565b91505092915050565b6000806000606084860312156119c957600080fd5b60006119d7868287016118b9565b93505060206119e8868287016118b9565b92505060406119f98682870161194c565b9150509250925092565b600060208284031215611a1557600080fd5b6000611a23848285016118ce565b91505092915050565b600060208284031215611a3e57600080fd5b6000611a4c848285016118e3565b91505092915050565b60008060408385031215611a6857600080fd5b6000611a76858286016118e3565b9250506020611a87858286016118b9565b9150509250929050565b600060208284031215611aa357600080fd5b600082013567ffffffffffffffff811115611abd57600080fd5b611ac9848285016118f8565b91505092915050565b600060208284031215611ae457600080fd5b6000611af28482850161194c565b91505092915050565b600060208284031215611b0d57600080fd5b6000611b1b84828501611976565b91505092915050565b60008060008060808587031215611b3a57600080fd5b6000611b4887828801611961565b9450506020611b598782880161194c565b9350506040611b6a878288016118b9565b925050606085013567ffffffffffffffff811115611b8757600080fd5b611b93878288016118f8565b91505092959194509250565b611ba881612325565b82525050565b611bb781612325565b82525050565b611bc681612337565b82525050565b611bd581612343565b82525050565b611be481612343565b82525050565b611bfb611bf682612343565b6123c6565b82525050565b6000611c0c826122dc565b611c1681856122f2565b9350611c26818560208601612393565b611c2f816123d0565b840191505092915050565b6000611c45826122dc565b611c4f8185612303565b9350611c5f818560208601612393565b611c68816123d0565b840191505092915050565b6000611c7e826122e7565b611c888185612314565b9350611c98818560208601612393565b611ca1816123d0565b840191505092915050565b6000611cb9603583612314565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611d1f601e83612314565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611d5f602483612314565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611dc5601b83612314565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000611e05603783612314565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b6000611e6b602883612314565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611ed1603383612314565b91507f70726f7669646564206f726967696e436861696e546f6b656e4164647265737360008301527f206973206e6f742077686974656c6973746564000000000000000000000000006020830152604082019050919050565b600060e083016000830151611f426000860182611b9f565b506020830151611f556020860182611fe4565b506040830151611f686040860182611bcc565b506060830151611f7b6060860182611fc6565b5060808301518482036080860152611f938282611c01565b91505060a0830151611fa860a0860182611b9f565b5060c0830151611fbb60c0860182611fc6565b508091505092915050565b611fcf8161236d565b82525050565b611fde8161236d565b82525050565b611fed81612377565b82525050565b611ffc81612377565b82525050565b600061200e8284611bea565b60208201915081905092915050565b60006020820190506120326000830184611bae565b92915050565b600060608201905061204d6000830186611bae565b61205a6020830185611bae565b6120676040830184611fd5565b949350505050565b60006040820190506120846000830185611bae565b6120916020830184611fd5565b9392505050565b600060e0820190506120ad600083018a611bae565b6120ba6020830189611ff3565b6120c76040830188611bdb565b6120d46060830187611fd5565b81810360808301526120e68186611c3a565b90506120f560a0830185611bae565b61210260c0830184611fd5565b98975050505050505050565b60006020820190506121236000830184611bbd565b92915050565b600060208201905061213e6000830184611bdb565b92915050565b6000602082019050818103600083015261215e8184611c73565b905092915050565b6000602082019050818103600083015261217f81611cac565b9050919050565b6000602082019050818103600083015261219f81611d12565b9050919050565b600060208201905081810360008301526121bf81611d52565b9050919050565b600060208201905081810360008301526121df81611db8565b9050919050565b600060208201905081810360008301526121ff81611df8565b9050919050565b6000602082019050818103600083015261221f81611e5e565b9050919050565b6000602082019050818103600083015261223f81611ec4565b9050919050565b600060208201905081810360008301526122608184611f2a565b905092915050565b600060208201905061227d6000830184611fd5565b92915050565b6000604051905081810181811067ffffffffffffffff821117156122a657600080fd5b8060405250919050565b600067ffffffffffffffff8211156122c757600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006123308261234d565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156123b1578082015181840152602081019050612396565b838111156123c0576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b6123ea81612325565b81146123f557600080fd5b50565b61240181612337565b811461240c57600080fd5b50565b61241881612343565b811461242357600080fd5b50565b61242f8161236d565b811461243a57600080fd5b50565b61244681612377565b811461245157600080fd5b5056fea26469706673582212209533cdb3c25a1a93f594193c729ce41dc9106dc5a333282b99058bae0f22772f64736f6c63430006040033"

// DeployERC20Handler deploys a new Ethereum contract, binding an instance of ERC20Handler to it.
func DeployERC20Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, burnableContractAddresses []common.Address) (common.Address, *types.Transaction, *ERC20Handler, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20HandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses, burnableContractAddresses)
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
// Solidity: function _balances(address ) view returns(uint256)
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
// Solidity: function _balances(address ) view returns(uint256)
func (_ERC20Handler *ERC20HandlerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC20Handler.Contract.Balances(&_ERC20Handler.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) view returns(uint256)
func (_ERC20Handler *ERC20HandlerCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC20Handler.Contract.Balances(&_ERC20Handler.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
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
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_burnList", arg0)
	return *ret0, err
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.BurnList(&_ERC20Handler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.BurnList(&_ERC20Handler.CallOpts, arg0)
}

// BurnedTokens is a free data retrieval call binding the contract method 0x45274738.
//
// Solidity: function _burnedTokens(address ) view returns(uint256)
func (_ERC20Handler *ERC20HandlerCaller) BurnedTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_burnedTokens", arg0)
	return *ret0, err
}

// BurnedTokens is a free data retrieval call binding the contract method 0x45274738.
//
// Solidity: function _burnedTokens(address ) view returns(uint256)
func (_ERC20Handler *ERC20HandlerSession) BurnedTokens(arg0 common.Address) (*big.Int, error) {
	return _ERC20Handler.Contract.BurnedTokens(&_ERC20Handler.CallOpts, arg0)
}

// BurnedTokens is a free data retrieval call binding the contract method 0x45274738.
//
// Solidity: function _burnedTokens(address ) view returns(uint256)
func (_ERC20Handler *ERC20HandlerCallerSession) BurnedTokens(arg0 common.Address) (*big.Int, error) {
	return _ERC20Handler.Contract.BurnedTokens(&_ERC20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
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
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.ContractWhitelist(&_ERC20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.ContractWhitelist(&_ERC20Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) view returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
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
// Solidity: function _depositRecords(uint256 ) view returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
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
// Solidity: function _depositRecords(uint256 ) view returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
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
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
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
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToTokenContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToTokenContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
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
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.TokenContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.TokenContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// UseContractWhitelist is a free data retrieval call binding the contract method 0x2521ab41.
//
// Solidity: function _useContractWhitelist() view returns(bool)
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
// Solidity: function _useContractWhitelist() view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) UseContractWhitelist() (bool, error) {
	return _ERC20Handler.Contract.UseContractWhitelist(&_ERC20Handler.CallOpts)
}

// UseContractWhitelist is a free data retrieval call binding the contract method 0x2521ab41.
//
// Solidity: function _useContractWhitelist() view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) UseContractWhitelist() (bool, error) {
	return _ERC20Handler.Contract.UseContractWhitelist(&_ERC20Handler.CallOpts)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) view returns(ERC20HandlerDepositRecord)
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
// Solidity: function getDepositRecord(uint256 depositID) view returns(ERC20HandlerDepositRecord)
func (_ERC20Handler *ERC20HandlerSession) GetDepositRecord(depositID *big.Int) (ERC20HandlerDepositRecord, error) {
	return _ERC20Handler.Contract.GetDepositRecord(&_ERC20Handler.CallOpts, depositID)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) view returns(ERC20HandlerDepositRecord)
func (_ERC20Handler *ERC20HandlerCallerSession) GetDepositRecord(depositID *big.Int) (ERC20HandlerDepositRecord, error) {
	return _ERC20Handler.Contract.GetDepositRecord(&_ERC20Handler.CallOpts, depositID)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address contractAddress) view returns(bool)
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
// Solidity: function isWhitelisted(address contractAddress) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) IsWhitelisted(contractAddress common.Address) (bool, error) {
	return _ERC20Handler.Contract.IsWhitelisted(&_ERC20Handler.CallOpts, contractAddress)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address contractAddress) view returns(bool)
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

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetBurnable(&_ERC20Handler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetBurnable(&_ERC20Handler.TransactOpts, contractAddress)
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

var RuntimeBytecode = "0x608060405234801561001057600080fd5b506004361061010b5760003560e01c80636ebcf607116100a2578063c8ba6c8711610071578063c8ba6c87146102dc578063d9caed121461030c578063db95e75c14610328578063e245386f14610358578063fc9539cd1461038e5761010b565b80636ebcf607146102445780637f79bea8146102745780638a025ce8146102a457806395601f09146102c05761010b565b80633af32abf116100de5780633af32abf1461019857806345274738146101c857806345a104db146101f85780636a70d081146102145761010b565b806307b7ed99146101105780630a6d55d81461012c5780632521ab411461015c578063318c136e1461017a575b600080fd5b61012a6004803603810190610125919061198b565b6103aa565b005b61014660048036038101906101419190611a2c565b61044d565b604051610153919061201d565b60405180910390f35b610164610480565b604051610171919061210e565b60405180910390f35b610182610493565b60405161018f919061201d565b60405180910390f35b6101b260048036038101906101ad919061198b565b6104b9565b6040516101bf919061210e565b60405180910390f35b6101e260048036038101906101dd919061198b565b61050f565b6040516101ef9190612268565b60405180910390f35b610212600480360381019061020d9190611b24565b610527565b005b61022e6004803603810190610229919061198b565b610845565b60405161023b919061210e565b60405180910390f35b61025e6004803603810190610259919061198b565b610865565b60405161026b9190612268565b60405180910390f35b61028e6004803603810190610289919061198b565b61087d565b60405161029b919061210e565b60405180910390f35b6102be60048036038101906102b99190611a55565b61089d565b005b6102da60048036038101906102d591906119b4565b610a23565b005b6102f660048036038101906102f1919061198b565b610b51565b6040516103039190612129565b60405180910390f35b610326600480360381019061032191906119b4565b610b69565b005b610342600480360381019061033d9190611ad2565b610c09565b60405161034f9190612246565b60405180910390f35b610372600480360381019061036d9190611ad2565b610dbf565b6040516103859796959493929190612098565b60405180910390f35b6103a860048036038101906103a39190611a91565b610ee6565b005b6103b3816104b9565b6103f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e9906121a6565b60405180910390fd5b6001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60036020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260149054906101000a900460ff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60016020528060005260406000206000915090505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ae90612186565b60405180910390fd5b60006060600080602085015193506040850151915060405192508460600151905080830160200160405260e4360360e4843760006003600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905061062a816104b9565b610669576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066090612226565b60405180910390fd5b600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156106cb576106c6818885611125565b6106d8565b6106d781883086611232565b5b6040518060e001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018a60ff1681526020018681526020018381526020018581526020018873ffffffffffffffffffffffffffffffffffffffff16815260200184815250600760008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff160217905550604082015181600101556060820151816002015560808201518160030190805190602001906107e59291906117a5565b5060a08201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160050155905050505050505050505050565b60066020528060005260406000206000915054906101000a900460ff1681565b60006020528060005260406000206000915090505481565b60056020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166003600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461093f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610936906121e6565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016109969190612002565b60405160208183030381529060405280519060200120826040516020016109bd9190612002565b6040516020818303038152906040528051906020012014610a13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0a90612166565b60405180910390fd5b610a1d8484611361565b50505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b8152600401610a6593929190612038565b602060405180830381600087803b158015610a7f57600080fd5b505af1158015610a93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab79190611a03565b50610b09826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461145390919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60046020528060005260406000206000915090505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bf9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bf090612186565b60405180910390fd5b610c048383836114a8565b505050565b610c11611825565b600760008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d4f5780601f10610d2457610100808354040283529160200191610d4f565b820191906000526020600020905b815481529060010190602001808311610d3257829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b60076020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff1690806001015490806002015490806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610eb05780601f10610e8557610100808354040283529160200191610eb0565b820191906000526020600020905b815481529060010190602001808311610e9357829003601f168201915b5050505050908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154905087565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6d90612186565b60405180910390fd5b60008060606020840151925060408401519150604051905083606001518082016020016040526084360360848337506000806003600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060208301519150610fee816104b9565b61102d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102490612206565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663beab71316040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561109e57600080fd5b505af11580156110b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d69190611afb565b90508060ff1686601f602081106110e957fe5b1a60f81b60f81c60ff16141561110c57611107838560601c896114a8565b61111b565b61111a838560601c896115d4565b5b5050505050505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166379cc679084846040518363ffffffff1660e01b815260040161116592919061206f565b600060405180830381600087803b15801561117f57600080fd5b505af1158015611193573d6000803e3d6000fd5b505050506111e982600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461145390919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b815260040161127493929190612038565b602060405180830381600087803b15801561128e57600080fd5b505af11580156112a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112c69190611a03565b50611318826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461145390919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008082840190508381101561149e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611495906121c6565b60405180910390fd5b8091505092915050565b60008390508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b81526004016114e892919061206f565b602060405180830381600087803b15801561150257600080fd5b505af1158015611516573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061153a9190611a03565b5061158c826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461170090919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b815260040161161492919061206f565b602060405180830381600087803b15801561162e57600080fd5b505af1158015611642573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116669190611a03565b506116b8826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461145390919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b600061174283836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061174a565b905092915050565b6000838311158290611792576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117899190612144565b60405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106117e657805160ff1916838001178555611814565b82800160010185558215611814579182015b828111156118135782518255916020019190600101906117f8565b5b5090506118219190611894565b5090565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600080191681526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b6118b691905b808211156118b257600081600090555060010161189a565b5090565b90565b6000813590506118c8816123e1565b92915050565b6000815190506118dd816123f8565b92915050565b6000813590506118f28161240f565b92915050565b600082601f83011261190957600080fd5b813561191c611917826122b0565b612283565b9150808252602083016020830185838301111561193857600080fd5b611943838284612384565b50505092915050565b60008135905061195b81612426565b92915050565b6000813590506119708161243d565b92915050565b6000815190506119858161243d565b92915050565b60006020828403121561199d57600080fd5b60006119ab848285016118b9565b91505092915050565b6000806000606084860312156119c957600080fd5b60006119d7868287016118b9565b93505060206119e8868287016118b9565b92505060406119f98682870161194c565b9150509250925092565b600060208284031215611a1557600080fd5b6000611a23848285016118ce565b91505092915050565b600060208284031215611a3e57600080fd5b6000611a4c848285016118e3565b91505092915050565b60008060408385031215611a6857600080fd5b6000611a76858286016118e3565b9250506020611a87858286016118b9565b9150509250929050565b600060208284031215611aa357600080fd5b600082013567ffffffffffffffff811115611abd57600080fd5b611ac9848285016118f8565b91505092915050565b600060208284031215611ae457600080fd5b6000611af28482850161194c565b91505092915050565b600060208284031215611b0d57600080fd5b6000611b1b84828501611976565b91505092915050565b60008060008060808587031215611b3a57600080fd5b6000611b4887828801611961565b9450506020611b598782880161194c565b9350506040611b6a878288016118b9565b925050606085013567ffffffffffffffff811115611b8757600080fd5b611b93878288016118f8565b91505092959194509250565b611ba881612325565b82525050565b611bb781612325565b82525050565b611bc681612337565b82525050565b611bd581612343565b82525050565b611be481612343565b82525050565b611bfb611bf682612343565b6123c6565b82525050565b6000611c0c826122dc565b611c1681856122f2565b9350611c26818560208601612393565b611c2f816123d0565b840191505092915050565b6000611c45826122dc565b611c4f8185612303565b9350611c5f818560208601612393565b611c68816123d0565b840191505092915050565b6000611c7e826122e7565b611c888185612314565b9350611c98818560208601612393565b611ca1816123d0565b840191505092915050565b6000611cb9603583612314565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611d1f601e83612314565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611d5f602483612314565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611dc5601b83612314565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000611e05603783612314565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b6000611e6b602883612314565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611ed1603383612314565b91507f70726f7669646564206f726967696e436861696e546f6b656e4164647265737360008301527f206973206e6f742077686974656c6973746564000000000000000000000000006020830152604082019050919050565b600060e083016000830151611f426000860182611b9f565b506020830151611f556020860182611fe4565b506040830151611f686040860182611bcc565b506060830151611f7b6060860182611fc6565b5060808301518482036080860152611f938282611c01565b91505060a0830151611fa860a0860182611b9f565b5060c0830151611fbb60c0860182611fc6565b508091505092915050565b611fcf8161236d565b82525050565b611fde8161236d565b82525050565b611fed81612377565b82525050565b611ffc81612377565b82525050565b600061200e8284611bea565b60208201915081905092915050565b60006020820190506120326000830184611bae565b92915050565b600060608201905061204d6000830186611bae565b61205a6020830185611bae565b6120676040830184611fd5565b949350505050565b60006040820190506120846000830185611bae565b6120916020830184611fd5565b9392505050565b600060e0820190506120ad600083018a611bae565b6120ba6020830189611ff3565b6120c76040830188611bdb565b6120d46060830187611fd5565b81810360808301526120e68186611c3a565b90506120f560a0830185611bae565b61210260c0830184611fd5565b98975050505050505050565b60006020820190506121236000830184611bbd565b92915050565b600060208201905061213e6000830184611bdb565b92915050565b6000602082019050818103600083015261215e8184611c73565b905092915050565b6000602082019050818103600083015261217f81611cac565b9050919050565b6000602082019050818103600083015261219f81611d12565b9050919050565b600060208201905081810360008301526121bf81611d52565b9050919050565b600060208201905081810360008301526121df81611db8565b9050919050565b600060208201905081810360008301526121ff81611df8565b9050919050565b6000602082019050818103600083015261221f81611e5e565b9050919050565b6000602082019050818103600083015261223f81611ec4565b9050919050565b600060208201905081810360008301526122608184611f2a565b905092915050565b600060208201905061227d6000830184611fd5565b92915050565b6000604051905081810181811067ffffffffffffffff821117156122a657600080fd5b8060405250919050565b600067ffffffffffffffff8211156122c757600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006123308261234d565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156123b1578082015181840152602081019050612396565b838111156123c0576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b6123ea81612325565b81146123f557600080fd5b50565b61240181612337565b811461240c57600080fd5b50565b61241881612343565b811461242357600080fd5b50565b61242f8161236d565b811461243a57600080fd5b50565b61244681612377565b811461245157600080fd5b5056fea26469706673582212209533cdb3c25a1a93f594193c729ce41dc9106dc5a333282b99058bae0f22772f64736f6c63430006040033"
