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
const GenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"initialDepositFunctionSignatures\",\"type\":\"bytes4[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"initialExecuteFunctionSignatures\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToExecuteFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structGenericHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GenericHandlerBin is the compiled bytecode used for deploying new contracts.
var GenericHandlerBin = "0x60806040523480156200001157600080fd5b50604051620024a5380380620024a5833981810160405281019062000037919062000570565b82518451146200007e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000759062000844565b60405180910390fd5b8151835114620000c5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000bc9062000800565b60405180910390fd5b80518251146200010c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001039062000822565b60405180910390fd5b846000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b8451811015620001ce57620001c08582815181106200016d57fe5b60200260200101518583815181106200018257fe5b60200260200101518584815181106200019757fe5b6020026020010151858581518110620001ac57fe5b6020026020010151620001da60201b60201c565b808060010191505062000152565b505050505050620009d8565b826002600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c021790555080600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c02179055506001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b60008151905062000395816200098a565b92915050565b600082601f830112620003ad57600080fd5b8151620003c4620003be8262000894565b62000866565b91508181835260208401935060208101905083856020840282011115620003ea57600080fd5b60005b838110156200041e578162000403888262000384565b845260208401935060208301925050600181019050620003ed565b5050505092915050565b600082601f8301126200043a57600080fd5b8151620004516200044b82620008bd565b62000866565b915081818352602084019350602081019050838560208402820111156200047757600080fd5b60005b83811015620004ab578162000490888262000542565b8452602084019350602083019250506001810190506200047a565b5050505092915050565b600082601f830112620004c757600080fd5b8151620004de620004d882620008e6565b62000866565b915081818352602084019350602081019050838560208402820111156200050457600080fd5b60005b838110156200053857816200051d888262000559565b84526020840193506020830192505060018101905062000507565b5050505092915050565b6000815190506200055381620009a4565b92915050565b6000815190506200056a81620009be565b92915050565b600080600080600060a086880312156200058957600080fd5b6000620005998882890162000384565b955050602086015167ffffffffffffffff811115620005b757600080fd5b620005c58882890162000428565b945050604086015167ffffffffffffffff811115620005e357600080fd5b620005f1888289016200039b565b935050606086015167ffffffffffffffff8111156200060f57600080fd5b6200061d88828901620004b5565b925050608086015167ffffffffffffffff8111156200063b57600080fd5b6200064988828901620004b5565b9150509295509295909350565b600062000665604b836200090f565b91507f6d69736d61746368206c656e677468206265747765656e2070726f766964656460008301527f20636f6e74726163742061646472657373657320616e642066756e6374696f6e60208301527f207369676e6174757265730000000000000000000000000000000000000000006040830152606082019050919050565b6000620006f36048836200090f565b91507f6d69736d61746368206c656e677468206265747765656e2070726f766964656460008301527f206465706f73697420616e6420657865637574652066756e6374696f6e20736960208301527f676e6174757265730000000000000000000000000000000000000000000000006040830152606082019050919050565b6000620007816047836200090f565b91507f6d69736d61746368206c656e677468206265747765656e20696e697469616c5260008301527f65736f7572636549447320616e6420696e697469616c436f6e7472616374416460208301527f64726573736573000000000000000000000000000000000000000000000000006040830152606082019050919050565b600060208201905081810360008301526200081b8162000656565b9050919050565b600060208201905081810360008301526200083d81620006e4565b9050919050565b600060208201905081810360008301526200085f8162000772565b9050919050565b6000604051905081810181811067ffffffffffffffff821117156200088a57600080fd5b8060405250919050565b600067ffffffffffffffff821115620008ac57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115620008d557600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115620008fe57600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b60006200092d826200096a565b9050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620009958162000920565b8114620009a157600080fd5b50565b620009af8162000934565b8114620009bb57600080fd5b50565b620009c9816200093e565b8114620009d557600080fd5b50565b611abd80620009e86000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063bba8185a11610071578063bba8185a14610178578063c54c2a1114610194578063cb624463146101c4578063ce3d5784146101f4578063ec97d3b414610227578063fc9539cd14610257576100a9565b8063318c136e146100ae57806345a104db146100cc5780637f79bea8146100e8578063a5c3a98514610118578063b631865314610148575b600080fd5b6100b6610273565b6040516100c391906116fb565b60405180910390f35b6100e660048036038101906100e191906112a2565b610298565b005b61010260048036038101906100fd9190611134565b610658565b60405161010f9190611716565b60405180910390f35b610132600480360381019061012d9190611134565b610678565b60405161013f919061174c565b60405180910390f35b610162600480360381019061015d919061122a565b610698565b60405161016f9190611807565b60405180910390f35b610192600480360381019061018d9190611186565b6107fc565b005b6101ae60048036038101906101a9919061115d565b610986565b6040516101bb91906116fb565b60405180910390f35b6101de60048036038101906101d99190611134565b6109b9565b6040516101eb919061174c565b60405180910390f35b61020e60048036038101906102099190611266565b6109d9565b60405161021e9493929190611829565b60405180910390f35b610241600480360381019061023c9190611134565b610adb565b60405161024e9190611731565b60405180910390f35b610271600480360381019061026c91906111e9565b610af3565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610327576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031e90611787565b60405180910390fd5b60008060606020840151925060408401519150604051905081810160200160405260c4360360c4823760006002600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040b906117e7565b60405180910390fd5b600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146105585760008173ffffffffffffffffffffffffffffffffffffffff16836040516104d091906116e4565b6000604051808303816000865af19150503d806000811461050d576040519150601f19603f3d011682016040523d82523d6000602084013e610512565b606091505b5050905080610556576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054d906117a7565b60405180910390fd5b505b60405180608001604052808960ff1681526020018581526020018773ffffffffffffffffffffffffffffffffffffffff16815260200183815250600160008a60ff1660ff168152602001908152602001600020600089815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301908051906020019061064a929190610f8e565b509050505050505050505050565b60066020528060005260406000206000915054906101000a900460ff1681565b60056020528060005260406000206000915054906101000a900460e01b81565b6106a061100e565b600160008360ff1660ff16815260200190815260200160002060008481526020019081526020016000206040518060800160405290816000820160009054906101000a900460ff1660ff1660ff168152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107eb5780601f106107c0576101008083540402835291602001916107eb565b820191906000526020600020905b8154815290600101906020018083116107ce57829003601f168201915b505050505081525050905092915050565b600073ffffffffffffffffffffffffffffffffffffffff166002600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461089e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610895906117c7565b60405180910390fd5b6000600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016108f591906116a1565b604051602081830303815290604052805190602001208260405160200161091c91906116a1565b6040516020818303038152906040528051906020012014610972576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096990611767565b60405180910390fd5b61097e86868686610de4565b505050505050565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915054906101000a900460e01b81565b6001602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900460ff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ad15780601f10610aa657610100808354040283529160200191610ad1565b820191906000526020600020905b815481529060010190602001808311610ab457829003601f168201915b5050505050905084565b60036020528060005260406000206000915090505481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b82576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7990611787565b60405180910390fd5b60006060602083015191506040519050826040015180820160600160405260643603606483375060006002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610c6d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c64906117e7565b60405180910390fd5b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610ddd5760608184604051602001610d1c9291906116bc565b604051602081830303815290604052905060008373ffffffffffffffffffffffffffffffffffffffff1682604051610d5491906116e4565b6000604051808303816000865af19150503d8060008114610d91576040519150601f19603f3d011682016040523d82523d6000602084013e610d96565b606091505b5050905080610dda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dd1906117a7565b60405180910390fd5b50505b5050505050565b826002600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c021790555080600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c02179055506001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610fcf57805160ff1916838001178555610ffd565b82800160010185558215610ffd579182015b82811115610ffc578251825591602001919060010190610fe1565b5b50905061100a9190611052565b5090565b6040518060800160405280600060ff16815260200160008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b61107491905b80821115611070576000816000905550600101611058565b5090565b90565b60008135905061108681611a14565b92915050565b60008135905061109b81611a2b565b92915050565b6000813590506110b081611a42565b92915050565b600082601f8301126110c757600080fd5b81356110da6110d5826118a2565b611875565b915080825260208301602083018583830111156110f657600080fd5b6111018382846119ad565b50505092915050565b60008135905061111981611a59565b92915050565b60008135905061112e81611a70565b92915050565b60006020828403121561114657600080fd5b600061115484828501611077565b91505092915050565b60006020828403121561116f57600080fd5b600061117d8482850161108c565b91505092915050565b6000806000806080858703121561119c57600080fd5b60006111aa8782880161108c565b94505060206111bb87828801611077565b93505060406111cc878288016110a1565b92505060606111dd878288016110a1565b91505092959194509250565b6000602082840312156111fb57600080fd5b600082013567ffffffffffffffff81111561121557600080fd5b611221848285016110b6565b91505092915050565b6000806040838503121561123d57600080fd5b600061124b8582860161110a565b925050602061125c8582860161111f565b9150509250929050565b6000806040838503121561127957600080fd5b60006112878582860161111f565b92505060206112988582860161110a565b9150509250929050565b600080600080608085870312156112b857600080fd5b60006112c68782880161111f565b94505060206112d78782880161110a565b93505060406112e887828801611077565b925050606085013567ffffffffffffffff81111561130557600080fd5b611311878288016110b6565b91505092959194509250565b61132681611922565b82525050565b61133581611922565b82525050565b61134481611934565b82525050565b61135381611940565b82525050565b61136281611940565b82525050565b61137961137482611940565b6119ef565b82525050565b6113888161194a565b82525050565b61139f61139a8261194a565b6119f9565b82525050565b60006113b0826118d9565b6113ba8185611906565b93506113ca8185602086016119bc565b80840191505092915050565b60006113e1826118ce565b6113eb81856118e4565b93506113fb8185602086016119bc565b61140481611a03565b840191505092915050565b600061141a826118ce565b61142481856118f5565b93506114348185602086016119bc565b61143d81611a03565b840191505092915050565b6000611455603583611911565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b60006114bb601e83611911565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b60006114fb602683611911565b91507f64656c656761746563616c6c20746f20636f6e7472616374416464726573732060008301527f6661696c656400000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611561603783611911565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b60006115c7602b83611911565b91507f70726f766964656420636f6e747261637441646472657373206973206e6f742060008301527f77686974656c69737465640000000000000000000000000000000000000000006020830152604082019050919050565b60006080830160008301516116386000860182611683565b50602083015161164b602086018261134a565b50604083015161165e604086018261131d565b506060830151848203606086015261167682826113d6565b9150508091505092915050565b61168c816119a0565b82525050565b61169b816119a0565b82525050565b60006116ad8284611368565b60208201915081905092915050565b60006116c8828561138e565b6004820191506116d882846113a5565b91508190509392505050565b60006116f082846113a5565b915081905092915050565b6000602082019050611710600083018461132c565b92915050565b600060208201905061172b600083018461133b565b92915050565b60006020820190506117466000830184611359565b92915050565b6000602082019050611761600083018461137f565b92915050565b6000602082019050818103600083015261178081611448565b9050919050565b600060208201905081810360008301526117a0816114ae565b9050919050565b600060208201905081810360008301526117c0816114ee565b9050919050565b600060208201905081810360008301526117e081611554565b9050919050565b60006020820190508181036000830152611800816115ba565b9050919050565b600060208201905081810360008301526118218184611620565b905092915050565b600060808201905061183e6000830187611692565b61184b6020830186611359565b611858604083018561132c565b818103606083015261186a818461140f565b905095945050505050565b6000604051905081810181811067ffffffffffffffff8211171561189857600080fd5b8060405250919050565b600067ffffffffffffffff8211156118b957600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061192d82611976565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156119da5780820151818401526020810190506119bf565b838111156119e9576000848401525b50505050565b6000819050919050565b6000819050919050565b6000601f19601f8301169050919050565b611a1d81611922565b8114611a2857600080fd5b50565b611a3481611940565b8114611a3f57600080fd5b50565b611a4b8161194a565b8114611a5657600080fd5b50565b611a6281611996565b8114611a6d57600080fd5b50565b611a79816119a0565b8114611a8457600080fd5b5056fea2646970667358221220b4529f2956743fea7ef8e3e57801b6023f5ebae078bcc60e0f3767aab1c7218d64736f6c63430006040033"

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

// DepositRecords is a free data retrieval call binding the contract method 0xce3d5784.
//
// Solidity: function _depositRecords(uint8 , uint256 ) view returns(uint8 _destinationChainID, bytes32 _resourceID, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 uint8, arg1 *big.Int) (struct {
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
	err := _GenericHandler.contract.Call(opts, out, "_depositRecords", arg0, arg1)
	return *ret, err
}

// DepositRecords is a free data retrieval call binding the contract method 0xce3d5784.
//
// Solidity: function _depositRecords(uint8 , uint256 ) view returns(uint8 _destinationChainID, bytes32 _resourceID, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerSession) DepositRecords(arg0 uint8, arg1 *big.Int) (struct {
	DestinationChainID uint8
	ResourceID         [32]byte
	Depositer          common.Address
	MetaData           []byte
}, error) {
	return _GenericHandler.Contract.DepositRecords(&_GenericHandler.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0xce3d5784.
//
// Solidity: function _depositRecords(uint8 , uint256 ) view returns(uint8 _destinationChainID, bytes32 _resourceID, address _depositer, bytes _metaData)
func (_GenericHandler *GenericHandlerCallerSession) DepositRecords(arg0 uint8, arg1 *big.Int) (struct {
	DestinationChainID uint8
	ResourceID         [32]byte
	Depositer          common.Address
	MetaData           []byte
}, error) {
	return _GenericHandler.Contract.DepositRecords(&_GenericHandler.CallOpts, arg0, arg1)
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

// GetDepositRecord is a free data retrieval call binding the contract method 0xb6318653.
//
// Solidity: function getDepositRecord(uint256 depositNonce, uint8 destId) view returns(GenericHandlerDepositRecord)
func (_GenericHandler *GenericHandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositNonce *big.Int, destId uint8) (GenericHandlerDepositRecord, error) {
	var (
		ret0 = new(GenericHandlerDepositRecord)
	)
	out := ret0
	err := _GenericHandler.contract.Call(opts, out, "getDepositRecord", depositNonce, destId)
	return *ret0, err
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xb6318653.
//
// Solidity: function getDepositRecord(uint256 depositNonce, uint8 destId) view returns(GenericHandlerDepositRecord)
func (_GenericHandler *GenericHandlerSession) GetDepositRecord(depositNonce *big.Int, destId uint8) (GenericHandlerDepositRecord, error) {
	return _GenericHandler.Contract.GetDepositRecord(&_GenericHandler.CallOpts, depositNonce, destId)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xb6318653.
//
// Solidity: function getDepositRecord(uint256 depositNonce, uint8 destId) view returns(GenericHandlerDepositRecord)
func (_GenericHandler *GenericHandlerCallerSession) GetDepositRecord(depositNonce *big.Int, destId uint8) (GenericHandlerDepositRecord, error) {
	return _GenericHandler.Contract.GetDepositRecord(&_GenericHandler.CallOpts, depositNonce, destId)
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
