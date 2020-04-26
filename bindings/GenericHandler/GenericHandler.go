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
	DestinationChainID uint8
	ResourceID         [32]byte
	Depositer          common.Address
	MetaData           []byte
}

// GenericHandlerABI is the input ABI used to generate the binding from.
const GenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"initialDepositFunctionSignatures\",\"type\":\"bytes4[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"initialExecuteFunctionSignatures\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToExecuteFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structGenericHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GenericHandlerBin is the compiled bytecode used for deploying new contracts.
var GenericHandlerBin = "0x60806040523480156200001157600080fd5b506040516200241a3803806200241a833981810160405281019062000037919062000570565b82518451146200007e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000759062000844565b60405180910390fd5b8151835114620000c5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000bc9062000800565b60405180910390fd5b80518251146200010c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001039062000822565b60405180910390fd5b846000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b8451811015620001ce57620001c08582815181106200016d57fe5b60200260200101518583815181106200018257fe5b60200260200101518584815181106200019757fe5b6020026020010151858581518110620001ac57fe5b6020026020010151620001da60201b60201c565b808060010191505062000152565b505050505050620009d8565b826002600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c021790555080600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c02179055506001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b60008151905062000395816200098a565b92915050565b600082601f830112620003ad57600080fd5b8151620003c4620003be8262000894565b62000866565b91508181835260208401935060208101905083856020840282011115620003ea57600080fd5b60005b838110156200041e578162000403888262000384565b845260208401935060208301925050600181019050620003ed565b5050505092915050565b600082601f8301126200043a57600080fd5b8151620004516200044b82620008bd565b62000866565b915081818352602084019350602081019050838560208402820111156200047757600080fd5b60005b83811015620004ab578162000490888262000542565b8452602084019350602083019250506001810190506200047a565b5050505092915050565b600082601f830112620004c757600080fd5b8151620004de620004d882620008e6565b62000866565b915081818352602084019350602081019050838560208402820111156200050457600080fd5b60005b838110156200053857816200051d888262000559565b84526020840193506020830192505060018101905062000507565b5050505092915050565b6000815190506200055381620009a4565b92915050565b6000815190506200056a81620009be565b92915050565b600080600080600060a086880312156200058957600080fd5b6000620005998882890162000384565b955050602086015167ffffffffffffffff811115620005b757600080fd5b620005c58882890162000428565b945050604086015167ffffffffffffffff811115620005e357600080fd5b620005f1888289016200039b565b935050606086015167ffffffffffffffff8111156200060f57600080fd5b6200061d88828901620004b5565b925050608086015167ffffffffffffffff8111156200063b57600080fd5b6200064988828901620004b5565b9150509295509295909350565b600062000665604b836200090f565b91507f6d69736d61746368206c656e677468206265747765656e2070726f766964656460008301527f20636f6e74726163742061646472657373657320616e642066756e6374696f6e60208301527f207369676e6174757265730000000000000000000000000000000000000000006040830152606082019050919050565b6000620006f36048836200090f565b91507f6d69736d61746368206c656e677468206265747765656e2070726f766964656460008301527f206465706f73697420616e6420657865637574652066756e6374696f6e20736960208301527f676e6174757265730000000000000000000000000000000000000000000000006040830152606082019050919050565b6000620007816047836200090f565b91507f6d69736d61746368206c656e677468206265747765656e20696e697469616c5260008301527f65736f7572636549447320616e6420696e697469616c436f6e7472616374416460208301527f64726573736573000000000000000000000000000000000000000000000000006040830152606082019050919050565b600060208201905081810360008301526200081b8162000656565b9050919050565b600060208201905081810360008301526200083d81620006e4565b9050919050565b600060208201905081810360008301526200085f8162000772565b9050919050565b6000604051905081810181811067ffffffffffffffff821117156200088a57600080fd5b8060405250919050565b600067ffffffffffffffff821115620008ac57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115620008d557600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115620008fe57600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b60006200092d826200096a565b9050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620009958162000920565b8114620009a157600080fd5b50565b620009af8162000934565b8114620009bb57600080fd5b50565b620009c9816200093e565b8114620009d557600080fd5b50565b611a3280620009e86000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063c54c2a1111610071578063c54c2a1114610164578063cb62446314610194578063db95e75c146101c4578063e245386f146101f4578063ec97d3b414610227578063fc9539cd14610257576100a9565b8063318c136e146100ae57806345a104db146100cc5780637f79bea8146100e8578063a5c3a98514610118578063bba8185a14610148575b600080fd5b6100b6610273565b6040516100c39190611670565b60405180910390f35b6100e660048036038101906100e19190611217565b610298565b005b61010260048036038101906100fd91906110f8565b610641565b60405161010f919061168b565b60405180910390f35b610132600480360381019061012d91906110f8565b610661565b60405161013f91906116c1565b60405180910390f35b610162600480360381019061015d919061114a565b610681565b005b61017e60048036038101906101799190611121565b61080b565b60405161018b9190611670565b60405180910390f35b6101ae60048036038101906101a991906110f8565b61083e565b6040516101bb91906116c1565b60405180910390f35b6101de60048036038101906101d991906111ee565b61085e565b6040516101eb919061177c565b60405180910390f35b61020e600480360381019061020991906111ee565b6109aa565b60405161021e949392919061179e565b60405180910390f35b610241600480360381019061023c91906110f8565b610a9f565b60405161024e91906116a6565b60405180910390f35b610271600480360381019061026c91906111ad565b610ab7565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610327576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031e906116fc565b60405180910390fd5b60008060606020840151925060408401519150604051905081810160200160405260c4360360c4823760006002600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040b9061175c565b60405180910390fd5b600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146105585760008173ffffffffffffffffffffffffffffffffffffffff16836040516104d09190611659565b6000604051808303816000865af19150503d806000811461050d576040519150601f19603f3d011682016040523d82523d6000602084013e610512565b606091505b5050905080610556576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054d9061171c565b60405180910390fd5b505b60405180608001604052808960ff1681526020018581526020018773ffffffffffffffffffffffffffffffffffffffff168152602001838152506001600089815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003019080519060200190610633929190610f52565b509050505050505050505050565b60066020528060005260406000206000915054906101000a900460ff1681565b60056020528060005260406000206000915054906101000a900460e01b81565b600073ffffffffffffffffffffffffffffffffffffffff166002600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610723576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071a9061173c565b60405180910390fd5b6000600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008060405160200161077a9190611616565b60405160208183030381529060405280519060200120826040516020016107a19190611616565b60405160208183030381529060405280519060200120146107f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ee906116dc565b60405180910390fd5b61080386868686610da8565b505050505050565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915054906101000a900460e01b81565b610866610fd2565b600160008381526020019081526020016000206040518060800160405290816000820160009054906101000a900460ff1660ff1660ff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561099a5780601f1061096f5761010080835404028352916020019161099a565b820191906000526020600020905b81548152906001019060200180831161097d57829003601f168201915b5050505050815250509050919050565b60016020528060005260406000206000915090508060000160009054906101000a900460ff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a955780601f10610a6a57610100808354040283529160200191610a95565b820191906000526020600020905b815481529060010190602001808311610a7857829003601f168201915b5050505050905084565b60036020528060005260406000206000915090505481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b46576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3d906116fc565b60405180910390fd5b60006060602083015191506040519050826040015180820160600160405260643603606483375060006002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610c31576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c289061175c565b60405180910390fd5b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610da15760608184604051602001610ce0929190611631565b604051602081830303815290604052905060008373ffffffffffffffffffffffffffffffffffffffff1682604051610d189190611659565b6000604051808303816000865af19150503d8060008114610d55576040519150601f19603f3d011682016040523d82523d6000602084013e610d5a565b606091505b5050905080610d9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d959061171c565b60405180910390fd5b50505b5050505050565b826002600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c021790555080600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c02179055506001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f9357805160ff1916838001178555610fc1565b82800160010185558215610fc1579182015b82811115610fc0578251825591602001919060010190610fa5565b5b509050610fce9190611016565b5090565b6040518060800160405280600060ff16815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b61103891905b8082111561103457600081600090555060010161101c565b5090565b90565b60008135905061104a81611989565b92915050565b60008135905061105f816119a0565b92915050565b600081359050611074816119b7565b92915050565b600082601f83011261108b57600080fd5b813561109e61109982611817565b6117ea565b915080825260208301602083018583830111156110ba57600080fd5b6110c5838284611922565b50505092915050565b6000813590506110dd816119ce565b92915050565b6000813590506110f2816119e5565b92915050565b60006020828403121561110a57600080fd5b60006111188482850161103b565b91505092915050565b60006020828403121561113357600080fd5b600061114184828501611050565b91505092915050565b6000806000806080858703121561116057600080fd5b600061116e87828801611050565b945050602061117f8782880161103b565b935050604061119087828801611065565b92505060606111a187828801611065565b91505092959194509250565b6000602082840312156111bf57600080fd5b600082013567ffffffffffffffff8111156111d957600080fd5b6111e58482850161107a565b91505092915050565b60006020828403121561120057600080fd5b600061120e848285016110ce565b91505092915050565b6000806000806080858703121561122d57600080fd5b600061123b878288016110e3565b945050602061124c878288016110ce565b935050604061125d8782880161103b565b925050606085013567ffffffffffffffff81111561127a57600080fd5b6112868782880161107a565b91505092959194509250565b61129b81611897565b82525050565b6112aa81611897565b82525050565b6112b9816118a9565b82525050565b6112c8816118b5565b82525050565b6112d7816118b5565b82525050565b6112ee6112e9826118b5565b611964565b82525050565b6112fd816118bf565b82525050565b61131461130f826118bf565b61196e565b82525050565b60006113258261184e565b61132f818561187b565b935061133f818560208601611931565b80840191505092915050565b600061135682611843565b6113608185611859565b9350611370818560208601611931565b61137981611978565b840191505092915050565b600061138f82611843565b611399818561186a565b93506113a9818560208601611931565b6113b281611978565b840191505092915050565b60006113ca603583611886565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611430601e83611886565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611470602683611886565b91507f64656c656761746563616c6c20746f20636f6e7472616374416464726573732060008301527f6661696c656400000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006114d6603783611886565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b600061153c602b83611886565b91507f70726f766964656420636f6e747261637441646472657373206973206e6f742060008301527f77686974656c69737465640000000000000000000000000000000000000000006020830152604082019050919050565b60006080830160008301516115ad60008601826115f8565b5060208301516115c060208601826112bf565b5060408301516115d36040860182611292565b50606083015184820360608601526115eb828261134b565b9150508091505092915050565b61160181611915565b82525050565b61161081611915565b82525050565b600061162282846112dd565b60208201915081905092915050565b600061163d8285611303565b60048201915061164d828461131a565b91508190509392505050565b6000611665828461131a565b915081905092915050565b600060208201905061168560008301846112a1565b92915050565b60006020820190506116a060008301846112b0565b92915050565b60006020820190506116bb60008301846112ce565b92915050565b60006020820190506116d660008301846112f4565b92915050565b600060208201905081810360008301526116f5816113bd565b9050919050565b6000602082019050818103600083015261171581611423565b9050919050565b6000602082019050818103600083015261173581611463565b9050919050565b60006020820190508181036000830152611755816114c9565b9050919050565b600060208201905081810360008301526117758161152f565b9050919050565b600060208201905081810360008301526117968184611595565b905092915050565b60006080820190506117b36000830187611607565b6117c060208301866112ce565b6117cd60408301856112a1565b81810360608301526117df8184611384565b905095945050505050565b6000604051905081810181811067ffffffffffffffff8211171561180d57600080fd5b8060405250919050565b600067ffffffffffffffff82111561182e57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006118a2826118eb565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b8381101561194f578082015181840152602081019050611934565b8381111561195e576000848401525b50505050565b6000819050919050565b6000819050919050565b6000601f19601f8301169050919050565b61199281611897565b811461199d57600080fd5b50565b6119a9816118b5565b81146119b457600080fd5b50565b6119c0816118bf565b81146119cb57600080fd5b50565b6119d78161190b565b81146119e257600080fd5b50565b6119ee81611915565b81146119f957600080fd5b5056fea2646970667358221220f26b846a5b46561d6488834e79e5de41b6c4fff221bfc8176199a5359e06967964736f6c63430006040033"

// DeployGenericHandler deploys a new Ethereum contract, binding an instance of GenericHandler to it.
func DeployGenericHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, initialDepositFunctionSignatures [][4]byte, initialExecuteFunctionSignatures [][4]byte) (common.Address, *types.Transaction, *GenericHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GenericHandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses, initialDepositFunctionSignatures, initialExecuteFunctionSignatures)
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

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToDepositFunctionSignature(opts *bind.CallOpts, arg0 common.Address) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _GenericHandler.contract.Call(opts, out, "_contractAddressToDepositFunctionSignature", arg0)
	return *ret0, err
}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerSession) ContractAddressToDepositFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToDepositFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToDepositFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToDepositFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToExecuteFunctionSignature(opts *bind.CallOpts, arg0 common.Address) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _GenericHandler.contract.Call(opts, out, "_contractAddressToExecuteFunctionSignature", arg0)
	return *ret0, err
}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerSession) ContractAddressToExecuteFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToExecuteFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToExecuteFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToExecuteFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _GenericHandler.contract.Call(opts, out, "_contractAddressToResourceID", arg0)
	return *ret0, err
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.ContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.ContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
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
// Solidity: function _depositRecords(uint256 ) view returns(uint8 _destinationChainID, bytes32 _resourceID, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DestinationChainID uint8
	ResourceID         [32]byte
	Depositer          common.Address
	MetaData           []byte
}, error) {
	ret := new(struct {
		DestinationChainID uint8
		ResourceID         [32]byte
		Depositer          common.Address
		MetaData           []byte
	})
	out := ret
	err := _GenericHandler.contract.Call(opts, out, "_depositRecords", arg0)
	return *ret, err
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) view returns(uint8 _destinationChainID, bytes32 _resourceID, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerSession) DepositRecords(arg0 *big.Int) (struct {
	DestinationChainID uint8
	ResourceID         [32]byte
	Depositer          common.Address
	MetaData           []byte
}, error) {
	return _GenericHandler.Contract.DepositRecords(&_GenericHandler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) view returns(uint8 _destinationChainID, bytes32 _resourceID, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerCallerSession) DepositRecords(arg0 *big.Int) (struct {
	DestinationChainID uint8
	ResourceID         [32]byte
	Depositer          common.Address
	MetaData           []byte
}, error) {
	return _GenericHandler.Contract.DepositRecords(&_GenericHandler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerCaller) ResourceIDToContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GenericHandler.contract.Call(opts, out, "_resourceIDToContractAddress", arg0)
	return *ret0, err
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToContractAddress(&_GenericHandler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerCallerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToContractAddress(&_GenericHandler.CallOpts, arg0)
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

// SetResource is a paid mutator transaction binding the contract method 0xbba8185a.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "setResource", resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// SetResource is a paid mutator transaction binding the contract method 0xbba8185a.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetResource(&_GenericHandler.TransactOpts, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// SetResource is a paid mutator transaction binding the contract method 0xbba8185a.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetResource(&_GenericHandler.TransactOpts, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

var RuntimeBytecode = "0x608060405234801561001057600080fd5b50600436106100a95760003560e01c8063c54c2a1111610071578063c54c2a1114610164578063cb62446314610194578063db95e75c146101c4578063e245386f146101f4578063ec97d3b414610227578063fc9539cd14610257576100a9565b8063318c136e146100ae57806345a104db146100cc5780637f79bea8146100e8578063a5c3a98514610118578063bba8185a14610148575b600080fd5b6100b6610273565b6040516100c39190611670565b60405180910390f35b6100e660048036038101906100e19190611217565b610298565b005b61010260048036038101906100fd91906110f8565b610641565b60405161010f919061168b565b60405180910390f35b610132600480360381019061012d91906110f8565b610661565b60405161013f91906116c1565b60405180910390f35b610162600480360381019061015d919061114a565b610681565b005b61017e60048036038101906101799190611121565b61080b565b60405161018b9190611670565b60405180910390f35b6101ae60048036038101906101a991906110f8565b61083e565b6040516101bb91906116c1565b60405180910390f35b6101de60048036038101906101d991906111ee565b61085e565b6040516101eb919061177c565b60405180910390f35b61020e600480360381019061020991906111ee565b6109aa565b60405161021e949392919061179e565b60405180910390f35b610241600480360381019061023c91906110f8565b610a9f565b60405161024e91906116a6565b60405180910390f35b610271600480360381019061026c91906111ad565b610ab7565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610327576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031e906116fc565b60405180910390fd5b60008060606020840151925060408401519150604051905081810160200160405260c4360360c4823760006002600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040b9061175c565b60405180910390fd5b600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146105585760008173ffffffffffffffffffffffffffffffffffffffff16836040516104d09190611659565b6000604051808303816000865af19150503d806000811461050d576040519150601f19603f3d011682016040523d82523d6000602084013e610512565b606091505b5050905080610556576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054d9061171c565b60405180910390fd5b505b60405180608001604052808960ff1681526020018581526020018773ffffffffffffffffffffffffffffffffffffffff168152602001838152506001600089815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003019080519060200190610633929190610f52565b509050505050505050505050565b60066020528060005260406000206000915054906101000a900460ff1681565b60056020528060005260406000206000915054906101000a900460e01b81565b600073ffffffffffffffffffffffffffffffffffffffff166002600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610723576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071a9061173c565b60405180910390fd5b6000600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008060405160200161077a9190611616565b60405160208183030381529060405280519060200120826040516020016107a19190611616565b60405160208183030381529060405280519060200120146107f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ee906116dc565b60405180910390fd5b61080386868686610da8565b505050505050565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915054906101000a900460e01b81565b610866610fd2565b600160008381526020019081526020016000206040518060800160405290816000820160009054906101000a900460ff1660ff1660ff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561099a5780601f1061096f5761010080835404028352916020019161099a565b820191906000526020600020905b81548152906001019060200180831161097d57829003601f168201915b5050505050815250509050919050565b60016020528060005260406000206000915090508060000160009054906101000a900460ff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a955780601f10610a6a57610100808354040283529160200191610a95565b820191906000526020600020905b815481529060010190602001808311610a7857829003601f168201915b5050505050905084565b60036020528060005260406000206000915090505481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b46576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3d906116fc565b60405180910390fd5b60006060602083015191506040519050826040015180820160600160405260643603606483375060006002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610c31576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c289061175c565b60405180910390fd5b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610da15760608184604051602001610ce0929190611631565b604051602081830303815290604052905060008373ffffffffffffffffffffffffffffffffffffffff1682604051610d189190611659565b6000604051808303816000865af19150503d8060008114610d55576040519150601f19603f3d011682016040523d82523d6000602084013e610d5a565b606091505b5050905080610d9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d959061171c565b60405180910390fd5b50505b5050505050565b826002600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c021790555080600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c02179055506001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f9357805160ff1916838001178555610fc1565b82800160010185558215610fc1579182015b82811115610fc0578251825591602001919060010190610fa5565b5b509050610fce9190611016565b5090565b6040518060800160405280600060ff16815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b61103891905b8082111561103457600081600090555060010161101c565b5090565b90565b60008135905061104a81611989565b92915050565b60008135905061105f816119a0565b92915050565b600081359050611074816119b7565b92915050565b600082601f83011261108b57600080fd5b813561109e61109982611817565b6117ea565b915080825260208301602083018583830111156110ba57600080fd5b6110c5838284611922565b50505092915050565b6000813590506110dd816119ce565b92915050565b6000813590506110f2816119e5565b92915050565b60006020828403121561110a57600080fd5b60006111188482850161103b565b91505092915050565b60006020828403121561113357600080fd5b600061114184828501611050565b91505092915050565b6000806000806080858703121561116057600080fd5b600061116e87828801611050565b945050602061117f8782880161103b565b935050604061119087828801611065565b92505060606111a187828801611065565b91505092959194509250565b6000602082840312156111bf57600080fd5b600082013567ffffffffffffffff8111156111d957600080fd5b6111e58482850161107a565b91505092915050565b60006020828403121561120057600080fd5b600061120e848285016110ce565b91505092915050565b6000806000806080858703121561122d57600080fd5b600061123b878288016110e3565b945050602061124c878288016110ce565b935050604061125d8782880161103b565b925050606085013567ffffffffffffffff81111561127a57600080fd5b6112868782880161107a565b91505092959194509250565b61129b81611897565b82525050565b6112aa81611897565b82525050565b6112b9816118a9565b82525050565b6112c8816118b5565b82525050565b6112d7816118b5565b82525050565b6112ee6112e9826118b5565b611964565b82525050565b6112fd816118bf565b82525050565b61131461130f826118bf565b61196e565b82525050565b60006113258261184e565b61132f818561187b565b935061133f818560208601611931565b80840191505092915050565b600061135682611843565b6113608185611859565b9350611370818560208601611931565b61137981611978565b840191505092915050565b600061138f82611843565b611399818561186a565b93506113a9818560208601611931565b6113b281611978565b840191505092915050565b60006113ca603583611886565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611430601e83611886565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611470602683611886565b91507f64656c656761746563616c6c20746f20636f6e7472616374416464726573732060008301527f6661696c656400000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006114d6603783611886565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b600061153c602b83611886565b91507f70726f766964656420636f6e747261637441646472657373206973206e6f742060008301527f77686974656c69737465640000000000000000000000000000000000000000006020830152604082019050919050565b60006080830160008301516115ad60008601826115f8565b5060208301516115c060208601826112bf565b5060408301516115d36040860182611292565b50606083015184820360608601526115eb828261134b565b9150508091505092915050565b61160181611915565b82525050565b61161081611915565b82525050565b600061162282846112dd565b60208201915081905092915050565b600061163d8285611303565b60048201915061164d828461131a565b91508190509392505050565b6000611665828461131a565b915081905092915050565b600060208201905061168560008301846112a1565b92915050565b60006020820190506116a060008301846112b0565b92915050565b60006020820190506116bb60008301846112ce565b92915050565b60006020820190506116d660008301846112f4565b92915050565b600060208201905081810360008301526116f5816113bd565b9050919050565b6000602082019050818103600083015261171581611423565b9050919050565b6000602082019050818103600083015261173581611463565b9050919050565b60006020820190508181036000830152611755816114c9565b9050919050565b600060208201905081810360008301526117758161152f565b9050919050565b600060208201905081810360008301526117968184611595565b905092915050565b60006080820190506117b36000830187611607565b6117c060208301866112ce565b6117cd60408301856112a1565b81810360608301526117df8184611384565b905095945050505050565b6000604051905081810181811067ffffffffffffffff8211171561180d57600080fd5b8060405250919050565b600067ffffffffffffffff82111561182e57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006118a2826118eb565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b8381101561194f578082015181840152602081019050611934565b8381111561195e576000848401525b50505050565b6000819050919050565b6000819050919050565b6000601f19601f8301169050919050565b61199281611897565b811461199d57600080fd5b50565b6119a9816118b5565b81146119b457600080fd5b50565b6119c0816118bf565b81146119cb57600080fd5b50565b6119d78161190b565b81146119e257600080fd5b50565b6119ee81611915565b81146119f957600080fd5b5056fea2646970667358221220f26b846a5b46561d6488834e79e5de41b6c4fff221bfc8176199a5359e06967964736f6c63430006040033"
