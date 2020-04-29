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
const ERC20HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"burnableContractAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_useContractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"internalType\":\"structERC20Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20HandlerBin is the compiled bytecode used for deploying new contracts.
var ERC20HandlerBin = "0x60806040523480156200001157600080fd5b5060405162002c8338038062002c838339818101604052810190620000379190620004a4565b81518351146200007e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000759062000673565b60405180910390fd5b83600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b8351811015620001175762000109848281518110620000e057fe5b6020026020010151848381518110620000f557fe5b60200260200101516200016660201b60201c565b8080600101915050620000c5565b5060008090505b81518110156200015b576200014d8282815181106200013957fe5b60200260200101516200025860201b60201c565b80806001019150506200011e565b505050505062000798565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b62000269816200030660201b60201c565b620002ab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002a29062000651565b60405180910390fd5b6001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000815190506200036d8162000764565b92915050565b600082601f8301126200038557600080fd5b81516200039c6200039682620006c3565b62000695565b91508181835260208401935060208101905083856020840282011115620003c257600080fd5b60005b83811015620003f65781620003db88826200035c565b845260208401935060208301925050600181019050620003c5565b5050505092915050565b600082601f8301126200041257600080fd5b8151620004296200042382620006ec565b62000695565b915081818352602084019350602081019050838560208402820111156200044f57600080fd5b60005b838110156200048357816200046888826200048d565b84526020840193506020830192505060018101905062000452565b5050505092915050565b6000815190506200049e816200077e565b92915050565b60008060008060808587031215620004bb57600080fd5b6000620004cb878288016200035c565b945050602085015167ffffffffffffffff811115620004e957600080fd5b620004f78782880162000400565b935050604085015167ffffffffffffffff8111156200051557600080fd5b620005238782880162000373565b925050606085015167ffffffffffffffff8111156200054157600080fd5b6200054f8782880162000373565b91505092959194509250565b60006200056a60248362000715565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000620005d260478362000715565b91507f6d69736d61746368206c656e677468206265747765656e20696e697469616c5260008301527f65736f7572636549447320616e6420696e697469616c436f6e7472616374416460208301527f64726573736573000000000000000000000000000000000000000000000000006040830152606082019050919050565b600060208201905081810360008301526200066c816200055b565b9050919050565b600060208201905081810360008301526200068e81620005c3565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715620006b957600080fd5b8060405250919050565b600067ffffffffffffffff821115620006db57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200070457600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620007338262000744565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200076f8162000726565b81146200077b57600080fd5b50565b62000789816200073a565b81146200079557600080fd5b50565b6124db80620007a86000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80636ebcf607116100a2578063c8ba6c8711610071578063c8ba6c87146102dc578063d9caed121461030c578063db95e75c14610328578063e245386f14610358578063fc9539cd1461038e5761010b565b80636ebcf607146102445780637f79bea81461027457806395601f09146102a4578063b8fa3736146102c05761010b565b80633af32abf116100de5780633af32abf1461019857806345274738146101c857806345a104db146101f85780636a70d081146102145761010b565b806307b7ed99146101105780630a6d55d81461012c5780632521ab411461015c578063318c136e1461017a575b600080fd5b61012a60048036038101906101259190611a05565b6103aa565b005b61014660048036038101906101419190611aa6565b610446565b604051610153919061206e565b60405180910390f35b610164610479565b604051610171919061215f565b60405180910390f35b61018261048c565b60405161018f919061206e565b60405180910390f35b6101b260048036038101906101ad9190611a05565b6104b2565b6040516101bf919061215f565b60405180910390f35b6101e260048036038101906101dd9190611a05565b610508565b6040516101ef91906122b9565b60405180910390f35b610212600480360381019061020d9190611b75565b610520565b005b61022e60048036038101906102299190611a05565b61083e565b60405161023b919061215f565b60405180910390f35b61025e60048036038101906102599190611a05565b61085e565b60405161026b91906122b9565b60405180910390f35b61028e60048036038101906102899190611a05565b610876565b60405161029b919061215f565b60405180910390f35b6102be60048036038101906102b99190611a2e565b610896565b005b6102da60048036038101906102d59190611acf565b6109c4565b005b6102f660048036038101906102f19190611a05565b610bda565b604051610303919061217a565b60405180910390f35b61032660048036038101906103219190611a2e565b610bf2565b005b610342600480360381019061033d9190611b4c565b610c92565b60405161034f9190612297565b60405180910390f35b610372600480360381019061036d9190611b4c565b610e48565b60405161038597969594939291906120e9565b60405180910390f35b6103a860048036038101906103a39190611b0b565b610f6f565b005b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461043a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610431906121d7565b60405180910390fd5b61044381611132565b50565b60036020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260149054906101000a900460ff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60016020528060005260406000206000915090505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a7906121d7565b60405180910390fd5b60006060600080602085015193506040850151915060405192508460600151905080830160200160405260e4360360e4843760006003600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050610623816104b2565b610662576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065990612277565b60405180910390fd5b600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156106c4576106bf8188856111d5565b6106d1565b6106d0818830866112e2565b5b6040518060e001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018a60ff1681526020018681526020018381526020018581526020018873ffffffffffffffffffffffffffffffffffffffff16815260200184815250600760008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff160217905550604082015181600101556060820151816002015560808201518160030190805190602001906107de929190611834565b5060a08201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160050155905050505050505050505050565b60066020528060005260406000206000915054906101000a900460ff1681565b60006020528060005260406000206000915090505481565b60056020528060005260406000206000915054906101000a900460ff1681565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b81526004016108d893929190612089565b602060405180830381600087803b1580156108f257600080fd5b505af1158015610906573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061092a9190611a7d565b5061097c826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461141190919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4b906121d7565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166003600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610af6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aed90612237565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600080604051602001610b4d9190612053565b6040516020818303038152906040528051906020012082604051602001610b749190612053565b6040516020818303038152906040528051906020012014610bca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc1906121b7565b60405180910390fd5b610bd48484611466565b50505050565b60046020528060005260406000206000915090505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c82576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c79906121d7565b60405180910390fd5b610c8d838383611558565b505050565b610c9a6118b4565b600760008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610dd85780601f10610dad57610100808354040283529160200191610dd8565b820191906000526020600020905b815481529060010190602001808311610dbb57829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b60076020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff1690806001015490806002015490806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f395780601f10610f0e57610100808354040283529160200191610f39565b820191906000526020600020905b815481529060010190602001808311610f1c57829003601f168201915b5050505050908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154905087565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff6906121d7565b60405180910390fd5b60008060606020840151915060408401519250604051905083606001518082016020016040526084360360848337506000806003600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060208301519150611077816104b2565b6110b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110ad90612257565b60405180910390fd5b600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561111b57611116818360601c87611684565b61112a565b611129818360601c87611558565b5b505050505050565b61113b816104b2565b61117a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611171906121f7565b60405180910390fd5b6001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166379cc679084846040518363ffffffff1660e01b81526004016112159291906120c0565b600060405180830381600087803b15801561122f57600080fd5b505af1158015611243573d6000803e3d6000fd5b5050505061129982600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461141190919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b815260040161132493929190612089565b602060405180830381600087803b15801561133e57600080fd5b505af1158015611352573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113769190611a7d565b506113c8826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461141190919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b60008082840190508381101561145c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161145390612217565b60405180910390fd5b8091505092915050565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b81526004016115989291906120c0565b602060405180830381600087803b1580156115b257600080fd5b505af11580156115c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ea9190611a7d565b5061163c826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461178f90919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b81526004016116c49291906120c0565b600060405180830381600087803b1580156116de57600080fd5b505af11580156116f2573d6000803e3d6000fd5b50505050611747826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461141190919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60006117d183836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506117d9565b905092915050565b6000838311158290611821576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118189190612195565b60405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061187557805160ff19168380011785556118a3565b828001600101855582156118a3579182015b828111156118a2578251825591602001919060010190611887565b5b5090506118b09190611923565b5090565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600080191681526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b61194591905b80821115611941576000816000905550600101611929565b5090565b90565b60008135905061195781612432565b92915050565b60008151905061196c81612449565b92915050565b60008135905061198181612460565b92915050565b600082601f83011261199857600080fd5b81356119ab6119a682612301565b6122d4565b915080825260208301602083018583830111156119c757600080fd5b6119d28382846123d5565b50505092915050565b6000813590506119ea81612477565b92915050565b6000813590506119ff8161248e565b92915050565b600060208284031215611a1757600080fd5b6000611a2584828501611948565b91505092915050565b600080600060608486031215611a4357600080fd5b6000611a5186828701611948565b9350506020611a6286828701611948565b9250506040611a73868287016119db565b9150509250925092565b600060208284031215611a8f57600080fd5b6000611a9d8482850161195d565b91505092915050565b600060208284031215611ab857600080fd5b6000611ac684828501611972565b91505092915050565b60008060408385031215611ae257600080fd5b6000611af085828601611972565b9250506020611b0185828601611948565b9150509250929050565b600060208284031215611b1d57600080fd5b600082013567ffffffffffffffff811115611b3757600080fd5b611b4384828501611987565b91505092915050565b600060208284031215611b5e57600080fd5b6000611b6c848285016119db565b91505092915050565b60008060008060808587031215611b8b57600080fd5b6000611b99878288016119f0565b9450506020611baa878288016119db565b9350506040611bbb87828801611948565b925050606085013567ffffffffffffffff811115611bd857600080fd5b611be487828801611987565b91505092959194509250565b611bf981612376565b82525050565b611c0881612376565b82525050565b611c1781612388565b82525050565b611c2681612394565b82525050565b611c3581612394565b82525050565b611c4c611c4782612394565b612417565b82525050565b6000611c5d8261232d565b611c678185612343565b9350611c778185602086016123e4565b611c8081612421565b840191505092915050565b6000611c968261232d565b611ca08185612354565b9350611cb08185602086016123e4565b611cb981612421565b840191505092915050565b6000611ccf82612338565b611cd98185612365565b9350611ce98185602086016123e4565b611cf281612421565b840191505092915050565b6000611d0a603583612365565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611d70601e83612365565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611db0602483612365565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611e16601b83612365565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000611e56603783612365565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b6000611ebc602883612365565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611f22602e83612365565b91507f70726f7669646564206f726967696e546f6b656e41646472657373206973206e60008301527f6f742077686974656c69737465640000000000000000000000000000000000006020830152604082019050919050565b600060e083016000830151611f936000860182611bf0565b506020830151611fa66020860182612035565b506040830151611fb96040860182611c1d565b506060830151611fcc6060860182612017565b5060808301518482036080860152611fe48282611c52565b91505060a0830151611ff960a0860182611bf0565b5060c083015161200c60c0860182612017565b508091505092915050565b612020816123be565b82525050565b61202f816123be565b82525050565b61203e816123c8565b82525050565b61204d816123c8565b82525050565b600061205f8284611c3b565b60208201915081905092915050565b60006020820190506120836000830184611bff565b92915050565b600060608201905061209e6000830186611bff565b6120ab6020830185611bff565b6120b86040830184612026565b949350505050565b60006040820190506120d56000830185611bff565b6120e26020830184612026565b9392505050565b600060e0820190506120fe600083018a611bff565b61210b6020830189612044565b6121186040830188611c2c565b6121256060830187612026565b81810360808301526121378186611c8b565b905061214660a0830185611bff565b61215360c0830184612026565b98975050505050505050565b60006020820190506121746000830184611c0e565b92915050565b600060208201905061218f6000830184611c2c565b92915050565b600060208201905081810360008301526121af8184611cc4565b905092915050565b600060208201905081810360008301526121d081611cfd565b9050919050565b600060208201905081810360008301526121f081611d63565b9050919050565b6000602082019050818103600083015261221081611da3565b9050919050565b6000602082019050818103600083015261223081611e09565b9050919050565b6000602082019050818103600083015261225081611e49565b9050919050565b6000602082019050818103600083015261227081611eaf565b9050919050565b6000602082019050818103600083015261229081611f15565b9050919050565b600060208201905081810360008301526122b18184611f7b565b905092915050565b60006020820190506122ce6000830184612026565b92915050565b6000604051905081810181811067ffffffffffffffff821117156122f757600080fd5b8060405250919050565b600067ffffffffffffffff82111561231857600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006123818261239e565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156124025780820151818401526020810190506123e7565b83811115612411576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b61243b81612376565b811461244657600080fd5b50565b61245281612388565b811461245d57600080fd5b50565b61246981612394565b811461247457600080fd5b50565b612480816123be565b811461248b57600080fd5b50565b612497816123c8565b81146124a257600080fd5b5056fea2646970667358221220b5a11cac1d2d6340e5998d678f726226efeac6db1416e6d1ddc1ca0add87d96f64736f6c63430006040033"

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

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResource(&_ERC20Handler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResource(&_ERC20Handler.TransactOpts, resourceID, contractAddress)
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

var RuntimeBytecode = "0x608060405234801561001057600080fd5b506004361061010b5760003560e01c80636ebcf607116100a2578063c8ba6c8711610071578063c8ba6c87146102dc578063d9caed121461030c578063db95e75c14610328578063e245386f14610358578063fc9539cd1461038e5761010b565b80636ebcf607146102445780637f79bea81461027457806395601f09146102a4578063b8fa3736146102c05761010b565b80633af32abf116100de5780633af32abf1461019857806345274738146101c857806345a104db146101f85780636a70d081146102145761010b565b806307b7ed99146101105780630a6d55d81461012c5780632521ab411461015c578063318c136e1461017a575b600080fd5b61012a60048036038101906101259190611a05565b6103aa565b005b61014660048036038101906101419190611aa6565b610446565b604051610153919061206e565b60405180910390f35b610164610479565b604051610171919061215f565b60405180910390f35b61018261048c565b60405161018f919061206e565b60405180910390f35b6101b260048036038101906101ad9190611a05565b6104b2565b6040516101bf919061215f565b60405180910390f35b6101e260048036038101906101dd9190611a05565b610508565b6040516101ef91906122b9565b60405180910390f35b610212600480360381019061020d9190611b75565b610520565b005b61022e60048036038101906102299190611a05565b61083e565b60405161023b919061215f565b60405180910390f35b61025e60048036038101906102599190611a05565b61085e565b60405161026b91906122b9565b60405180910390f35b61028e60048036038101906102899190611a05565b610876565b60405161029b919061215f565b60405180910390f35b6102be60048036038101906102b99190611a2e565b610896565b005b6102da60048036038101906102d59190611acf565b6109c4565b005b6102f660048036038101906102f19190611a05565b610bda565b604051610303919061217a565b60405180910390f35b61032660048036038101906103219190611a2e565b610bf2565b005b610342600480360381019061033d9190611b4c565b610c92565b60405161034f9190612297565b60405180910390f35b610372600480360381019061036d9190611b4c565b610e48565b60405161038597969594939291906120e9565b60405180910390f35b6103a860048036038101906103a39190611b0b565b610f6f565b005b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461043a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610431906121d7565b60405180910390fd5b61044381611132565b50565b60036020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260149054906101000a900460ff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60016020528060005260406000206000915090505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a7906121d7565b60405180910390fd5b60006060600080602085015193506040850151915060405192508460600151905080830160200160405260e4360360e4843760006003600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050610623816104b2565b610662576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065990612277565b60405180910390fd5b600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156106c4576106bf8188856111d5565b6106d1565b6106d0818830866112e2565b5b6040518060e001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018a60ff1681526020018681526020018381526020018581526020018873ffffffffffffffffffffffffffffffffffffffff16815260200184815250600760008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff160217905550604082015181600101556060820151816002015560808201518160030190805190602001906107de929190611834565b5060a08201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160050155905050505050505050505050565b60066020528060005260406000206000915054906101000a900460ff1681565b60006020528060005260406000206000915090505481565b60056020528060005260406000206000915054906101000a900460ff1681565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b81526004016108d893929190612089565b602060405180830381600087803b1580156108f257600080fd5b505af1158015610906573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061092a9190611a7d565b5061097c826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461141190919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4b906121d7565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166003600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610af6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aed90612237565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600080604051602001610b4d9190612053565b6040516020818303038152906040528051906020012082604051602001610b749190612053565b6040516020818303038152906040528051906020012014610bca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc1906121b7565b60405180910390fd5b610bd48484611466565b50505050565b60046020528060005260406000206000915090505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c82576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c79906121d7565b60405180910390fd5b610c8d838383611558565b505050565b610c9a6118b4565b600760008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610dd85780601f10610dad57610100808354040283529160200191610dd8565b820191906000526020600020905b815481529060010190602001808311610dbb57829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b60076020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff1690806001015490806002015490806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f395780601f10610f0e57610100808354040283529160200191610f39565b820191906000526020600020905b815481529060010190602001808311610f1c57829003601f168201915b5050505050908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154905087565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff6906121d7565b60405180910390fd5b60008060606020840151915060408401519250604051905083606001518082016020016040526084360360848337506000806003600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060208301519150611077816104b2565b6110b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110ad90612257565b60405180910390fd5b600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561111b57611116818360601c87611684565b61112a565b611129818360601c87611558565b5b505050505050565b61113b816104b2565b61117a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611171906121f7565b60405180910390fd5b6001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166379cc679084846040518363ffffffff1660e01b81526004016112159291906120c0565b600060405180830381600087803b15801561122f57600080fd5b505af1158015611243573d6000803e3d6000fd5b5050505061129982600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461141190919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b815260040161132493929190612089565b602060405180830381600087803b15801561133e57600080fd5b505af1158015611352573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113769190611a7d565b506113c8826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461141190919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b60008082840190508381101561145c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161145390612217565b60405180910390fd5b8091505092915050565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b81526004016115989291906120c0565b602060405180830381600087803b1580156115b257600080fd5b505af11580156115c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ea9190611a7d565b5061163c826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461178f90919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b81526004016116c49291906120c0565b600060405180830381600087803b1580156116de57600080fd5b505af11580156116f2573d6000803e3d6000fd5b50505050611747826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461141190919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60006117d183836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506117d9565b905092915050565b6000838311158290611821576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118189190612195565b60405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061187557805160ff19168380011785556118a3565b828001600101855582156118a3579182015b828111156118a2578251825591602001919060010190611887565b5b5090506118b09190611923565b5090565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600080191681526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b61194591905b80821115611941576000816000905550600101611929565b5090565b90565b60008135905061195781612432565b92915050565b60008151905061196c81612449565b92915050565b60008135905061198181612460565b92915050565b600082601f83011261199857600080fd5b81356119ab6119a682612301565b6122d4565b915080825260208301602083018583830111156119c757600080fd5b6119d28382846123d5565b50505092915050565b6000813590506119ea81612477565b92915050565b6000813590506119ff8161248e565b92915050565b600060208284031215611a1757600080fd5b6000611a2584828501611948565b91505092915050565b600080600060608486031215611a4357600080fd5b6000611a5186828701611948565b9350506020611a6286828701611948565b9250506040611a73868287016119db565b9150509250925092565b600060208284031215611a8f57600080fd5b6000611a9d8482850161195d565b91505092915050565b600060208284031215611ab857600080fd5b6000611ac684828501611972565b91505092915050565b60008060408385031215611ae257600080fd5b6000611af085828601611972565b9250506020611b0185828601611948565b9150509250929050565b600060208284031215611b1d57600080fd5b600082013567ffffffffffffffff811115611b3757600080fd5b611b4384828501611987565b91505092915050565b600060208284031215611b5e57600080fd5b6000611b6c848285016119db565b91505092915050565b60008060008060808587031215611b8b57600080fd5b6000611b99878288016119f0565b9450506020611baa878288016119db565b9350506040611bbb87828801611948565b925050606085013567ffffffffffffffff811115611bd857600080fd5b611be487828801611987565b91505092959194509250565b611bf981612376565b82525050565b611c0881612376565b82525050565b611c1781612388565b82525050565b611c2681612394565b82525050565b611c3581612394565b82525050565b611c4c611c4782612394565b612417565b82525050565b6000611c5d8261232d565b611c678185612343565b9350611c778185602086016123e4565b611c8081612421565b840191505092915050565b6000611c968261232d565b611ca08185612354565b9350611cb08185602086016123e4565b611cb981612421565b840191505092915050565b6000611ccf82612338565b611cd98185612365565b9350611ce98185602086016123e4565b611cf281612421565b840191505092915050565b6000611d0a603583612365565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611d70601e83612365565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611db0602483612365565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611e16601b83612365565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000611e56603783612365565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b6000611ebc602883612365565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611f22602e83612365565b91507f70726f7669646564206f726967696e546f6b656e41646472657373206973206e60008301527f6f742077686974656c69737465640000000000000000000000000000000000006020830152604082019050919050565b600060e083016000830151611f936000860182611bf0565b506020830151611fa66020860182612035565b506040830151611fb96040860182611c1d565b506060830151611fcc6060860182612017565b5060808301518482036080860152611fe48282611c52565b91505060a0830151611ff960a0860182611bf0565b5060c083015161200c60c0860182612017565b508091505092915050565b612020816123be565b82525050565b61202f816123be565b82525050565b61203e816123c8565b82525050565b61204d816123c8565b82525050565b600061205f8284611c3b565b60208201915081905092915050565b60006020820190506120836000830184611bff565b92915050565b600060608201905061209e6000830186611bff565b6120ab6020830185611bff565b6120b86040830184612026565b949350505050565b60006040820190506120d56000830185611bff565b6120e26020830184612026565b9392505050565b600060e0820190506120fe600083018a611bff565b61210b6020830189612044565b6121186040830188611c2c565b6121256060830187612026565b81810360808301526121378186611c8b565b905061214660a0830185611bff565b61215360c0830184612026565b98975050505050505050565b60006020820190506121746000830184611c0e565b92915050565b600060208201905061218f6000830184611c2c565b92915050565b600060208201905081810360008301526121af8184611cc4565b905092915050565b600060208201905081810360008301526121d081611cfd565b9050919050565b600060208201905081810360008301526121f081611d63565b9050919050565b6000602082019050818103600083015261221081611da3565b9050919050565b6000602082019050818103600083015261223081611e09565b9050919050565b6000602082019050818103600083015261225081611e49565b9050919050565b6000602082019050818103600083015261227081611eaf565b9050919050565b6000602082019050818103600083015261229081611f15565b9050919050565b600060208201905081810360008301526122b18184611f7b565b905092915050565b60006020820190506122ce6000830184612026565b92915050565b6000604051905081810181811067ffffffffffffffff821117156122f757600080fd5b8060405250919050565b600067ffffffffffffffff82111561231857600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006123818261239e565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156124025780820151818401526020810190506123e7565b83811115612411576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b61243b81612376565b811461244657600080fd5b50565b61245281612388565b811461245d57600080fd5b50565b61246981612394565b811461247457600080fd5b50565b612480816123be565b811461248b57600080fd5b50565b612497816123c8565b81146124a257600080fd5b5056fea2646970667358221220b5a11cac1d2d6340e5998d678f726226efeac6db1416e6d1ddc1ca0add87d96f64736f6c63430006040033"
