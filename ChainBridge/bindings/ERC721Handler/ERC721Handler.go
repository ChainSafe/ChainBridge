// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC721Handler

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC721HandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type ERC721HandlerDepositRecord struct {
	TokenAddress                common.Address
	DestinationChainID          uint8
	ResourceID                  [32]byte
	DestinationRecipientAddress []byte
	Depositer                   common.Address
	TokenID                     *big.Int
	MetaData                    []byte
}

// ERC721HandlerMetaData contains all meta data concerning the ERC721Handler contract.
var ERC721HandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"burnableContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"lockMintUnlockableContractAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_lockMintUnlockList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"fundERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setLockMintUnlockable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structERC721Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001e6138038062001e61833981016040819052620000349162000371565b8251845114620000615760405162461bcd60e51b8152600401620000589062000597565b60405180910390fd5b600080546001600160a01b0319166001600160a01b0387161781555b8451811015620000ca57620000c18582815181106200009857fe5b6020026020010151858381518110620000ad57fe5b60200260200101516200014e60201b60201c565b6001016200007d565b5060005b82518110156200010657620000fd838281518110620000e957fe5b60200260200101516200019d60201b60201c565b600101620000ce565b5060005b81518110156200014257620001398282815181106200012557fe5b60200260200101516200023860201b60201c565b6001016200010a565b5050505050506200063b565b600082815260016020818152604080842080546001600160a01b039096166001600160a01b0319909616861790559383526002815283832094909455600390935220805460ff19169091179055565b6001600160a01b03811660009081526003602052604090205460ff16620001d85760405162461bcd60e51b8152600401620000589062000553565b6001600160a01b03811660009081526005602052604090205460ff1615620002145760405162461bcd60e51b8152600401620000589062000499565b6001600160a01b03166000908152600460205260409020805460ff19166001179055565b6001600160a01b03811660009081526003602052604090205460ff16620002735760405162461bcd60e51b8152600401620000589062000553565b6001600160a01b03811660009081526004602052604090205460ff1615620002af5760405162461bcd60e51b81526004016200005890620004f6565b6001600160a01b03166000908152600560205260409020805460ff19166001179055565b80516001600160a01b0381168114620002eb57600080fd5b92915050565b600082601f83011262000302578081fd5b81516200031962000313826200061b565b620005f4565b8181529150602080830190848101818402860182018710156200033b57600080fd5b60005b848110156200036657620003538883620002d3565b845292820192908201906001016200033e565b505050505092915050565b600080600080600060a0868803121562000389578081fd5b620003958787620002d3565b602087810151919650906001600160401b0380821115620003b4578384fd5b818901915089601f830112620003c8578384fd5b8151620003d962000313826200061b565b81815284810190848601868402860187018e1015620003f6578788fd5b8795505b838610156200041a578051835260019590950194918601918601620003fa565b5060408c0151909950945050508083111562000434578384fd5b620004428a848b01620002f1565b9550606089015192508083111562000458578384fd5b620004668a848b01620002f1565b945060808901519250808311156200047c578384fd5b50506200048c88828901620002f1565b9150509295509295909350565b6020808252603b908201527f70726f766964656420636f6e74726163742063616e6e6f74206265206c6f636b60408201527f4d696e74556e6c6f636b61626c6520616e64206275726e61626c650000000000606082015260800190565b6020808252603b908201527f70726f766964656420636f6e74726163742063616e6e6f74206265206275726e60408201527f61626c6520616e64206c6f636b4d696e74556e6c6f636b61626c650000000000606082015260800190565b60208082526024908201527f70726f766964656420636f6e7472616374206973206e6f742077686974656c696040820152631cdd195960e21b606082015260800190565b6020808252603c908201527f696e697469616c5265736f7572636549447320616e6420696e697469616c436f60408201527f6e7472616374416464726573736573206c656e206d69736d6174636800000000606082015260800190565b6040518181016001600160401b03811182821017156200061357600080fd5b604052919050565b60006001600160401b0382111562000631578081fd5b5060209081020190565b611816806200064b6000396000f3fe608060405234801561001057600080fd5b50600436106100f45760003560e01c80637354298011610097578063ba484c0911610066578063ba484c09146101f7578063c8ba6c8714610217578063d9caed1214610237578063e248cff21461024a576100f4565b806373542980146101ab5780637f79bea8146101be578063b1f06463146101d1578063b8fa3736146101e4576100f4565b8063318c136e116100d3578063318c136e1461014a57806338995da9146101525780634402027f146101655780636a70d0811461018b576100f4565b8062d8e7ee146100f957806307b7ed991461010e5780630a6d55d814610121575b600080fd5b61010c610107366004611169565b61025d565b005b61010c61011c366004611169565b610271565b61013461012f3660046111e4565b610282565b6040516101419190611483565b60405180910390f35b61013461029d565b61010c610160366004611275565b6102ac565b6101786101733660046113fd565b610598565b60405161014197969594939291906114e2565b61019e610199366004611169565b610708565b6040516101419190611542565b61010c6101b9366004611184565b61071d565b61019e6101cc366004611169565b610787565b61019e6101df366004611169565b61079c565b61010c6101f23660046111fc565b6107b1565b61020a6102053660046113c9565b6107c7565b60405161014191906116e8565b61022a610225366004611169565b61097c565b604051610141919061154d565b61010c610245366004611184565b61098e565b61010c61025836600461122b565b6109a7565b610265610be9565b61026e81610c15565b50565b610279610be9565b61026e81610caa565b6001602052600090815260409020546001600160a01b031681565b6000546001600160a01b031681565b6102b4610be9565b6000806060806102c6858701876113a8565b90945092506102db6040808501908789611767565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052508e8152600160209081526040808320546001600160a01b03168084526003909252909120549496509360ff1692506103619150505760405162461bcd60e51b8152600401610358906116a0565b60405180910390fd5b61037b6001600160a01b038216635b5e139f60e01b610d3f565b156104065760405163c87b56dd60e01b815281906001600160a01b0382169063c87b56dd906103ae90899060040161154d565b60006040518083038186803b1580156103c657600080fd5b505afa1580156103da573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261040291908101906112f5565b9250505b6001600160a01b03811660009081526004602052604090205460ff1615610436576104318186610d62565b610442565b61044281893088610dc7565b6040518060e00160405280826001600160a01b031681526020018b60ff1681526020018c8152602001848152602001896001600160a01b0316815260200186815260200183815250600660008c60ff1660ff16815260200190815260200160002060008b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160000160146101000a81548160ff021916908360ff16021790555060408201518160010155606082015181600201908051906020019061053c929190611001565b5060808201516003820180546001600160a01b0319166001600160a01b0390921691909117905560a0820151600482015560c08201518051610588916005840191602090910190611001565b5050505050505050505050505050565b600660209081526000928352604080842082529183529181902080546001808301546002808501805487516101009582161595909502600019011691909104601f81018890048802840188019096528583526001600160a01b03841696600160a01b90940460ff16959194939091908301828280156106585780601f1061062d57610100808354040283529160200191610658565b820191906000526020600020905b81548152906001019060200180831161063b57829003601f168201915b505050506003830154600484015460058501805460408051602060026101006001861615026000190190941693909304601f810184900484028201840190925281815296976001600160a01b0390951696939550908301828280156106fe5780601f106106d3576101008083540402835291602001916106fe565b820191906000526020600020905b8154815290600101906020018083116106e157829003601f168201915b5050505050905087565b60046020526000908152604090205460ff1681565b6040516323b872dd60e01b815283906001600160a01b038216906323b872dd9061074f90869030908790600401611497565b600060405180830381600087803b15801561076957600080fd5b505af115801561077d573d6000803e3d6000fd5b5050505050505050565b60036020526000908152604090205460ff1681565b60056020526000908152604090205460ff1681565b6107b9610be9565b6107c38282610e32565b5050565b6107cf61107f565b60ff828116600090815260066020908152604080832067ffffffffffffffff88168452825291829020825160e08101845281546001600160a01b0381168252600160a01b900490941684830152600180820154858501526002808301805486516101009482161594909402600019011691909104601f8101859004850283018501909552848252919360608601939192918301828280156108b15780601f10610886576101008083540402835291602001916108b1565b820191906000526020600020905b81548152906001019060200180831161089457829003601f168201915b505050918352505060038201546001600160a01b03166020808301919091526004830154604080840191909152600584018054825160026101006001841615026000190190921691909104601f81018590048502820185019093528281526060909401939290919083018282801561096a5780601f1061093f5761010080835404028352916020019161096a565b820191906000526020600020905b81548152906001019060200180831161094d57829003601f168201915b50505050508152505090505b92915050565b60026020526000908152604090205481565b610996610be9565b6109a283308484610dc7565b505050565b6109af610be9565b60008060608180826109c3878901896113a8565b9096509450604080860193506109dc908490898b611767565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929650610a229250899150859050818b611767565b810190610a2f91906111e4565b9150610a4560208484018101908501898b611767565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052506020898101518f8352600182526040808420546001600160a01b0316808552600390935290922054959650909490935060ff169150610ac890505760405162461bcd60e51b8152600401610358906116a0565b6001600160a01b03811660009081526004602052604090205460ff1615610afd57610af8818360601c8a86610e81565b610bdc565b6001600160a01b03811660009081526005602052604090205460ff1615610bcd576040516370a0823160e01b815281906000906001600160a01b038316906370a0823190610b4f903090600401611483565b60206040518083038186803b158015610b6757600080fd5b505afa158015610b7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9f9190611390565b11610bb857610bb3828460601c8b87610e81565b610bc7565b610bc782308560601c8c610dc7565b50610bdc565b610bdc81308460601c8b610dc7565b5050505050505050505050565b6000546001600160a01b03163314610c135760405162461bcd60e51b8152600401610358906115c8565b565b6001600160a01b03811660009081526003602052604090205460ff16610c4d5760405162461bcd60e51b81526004016103589061165c565b6001600160a01b03811660009081526004602052604090205460ff1615610c865760405162461bcd60e51b8152600401610358906115ff565b6001600160a01b03166000908152600560205260409020805460ff19166001179055565b6001600160a01b03811660009081526003602052604090205460ff16610ce25760405162461bcd60e51b81526004016103589061165c565b6001600160a01b03811660009081526005602052604090205460ff1615610d1b5760405162461bcd60e51b81526004016103589061156b565b6001600160a01b03166000908152600460205260409020805460ff19166001179055565b6000610d4a83610eb3565b8015610d5b5750610d5b8383610ee6565b9392505050565b604051630852cd8d60e31b815282906001600160a01b038216906342966c6890610d9090859060040161154d565b600060405180830381600087803b158015610daa57600080fd5b505af1158015610dbe573d6000803e3d6000fd5b50505050505050565b6040516323b872dd60e01b815284906001600160a01b038216906323b872dd90610df990879087908790600401611497565b600060405180830381600087803b158015610e1357600080fd5b505af1158015610e27573d6000803e3d6000fd5b505050505050505050565b600082815260016020818152604080842080546001600160a01b039096166001600160a01b0319909616861790559383526002815283832094909455600390935220805460ff19169091179055565b6040516334ff261960e21b815284906001600160a01b0382169063d3fc986490610df9908790879087906004016114bb565b6000610ec6826301ffc9a760e01b610ee6565b80156109765750610edf826001600160e01b0319610ee6565b1592915050565b6000806000610ef58585610f0c565b91509150818015610f035750805b95945050505050565b60008060606301ffc9a760e01b84604051602401610f2a9190611556565b604051602081830303815290604052906001600160e01b0319166020820180516001600160e01b038381831617835250505050905060006060866001600160a01b031661753084604051610f7e9190611467565b6000604051808303818686fa925050503d8060008114610fba576040519150601f19603f3d011682016040523d82523d6000602084013e610fbf565b606091505b5091509150602081511015610fdd5760008094509450505050610ffa565b8181806020019051810190610ff291906111c4565b945094505050505b9250929050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061104257805160ff191683800117855561106f565b8280016001018555821561106f579182015b8281111561106f578251825591602001919060010190611054565b5061107b9291506110d4565b5090565b6040518060e0016040528060006001600160a01b03168152602001600060ff168152602001600080191681526020016060815260200160006001600160a01b0316815260200160008152602001606081525090565b5b8082111561107b57600081556001016110d5565b80356001600160a01b038116811461097657600080fd5b60008083601f840112611111578182fd5b50813567ffffffffffffffff811115611128578182fd5b602083019150836020828501011115610ffa57600080fd5b803567ffffffffffffffff8116811461097657600080fd5b803560ff8116811461097657600080fd5b60006020828403121561117a578081fd5b610d5b83836110e9565b600080600060608486031215611198578182fd5b83356111a3816117cb565b925060208401356111b3816117cb565b929592945050506040919091013590565b6000602082840312156111d5578081fd5b81518015158114610d5b578182fd5b6000602082840312156111f5578081fd5b5035919050565b6000806040838503121561120e578182fd5b823591506020830135611220816117cb565b809150509250929050565b60008060006040848603121561123f578283fd5b83359250602084013567ffffffffffffffff81111561125c578283fd5b61126886828701611100565b9497909650939450505050565b60008060008060008060a0878903121561128d578182fd5b8635955061129e8860208901611158565b94506112ad8860408901611140565b93506112bc88606089016110e9565b9250608087013567ffffffffffffffff8111156112d7578283fd5b6112e389828a01611100565b979a9699509497509295939492505050565b600060208284031215611306578081fd5b815167ffffffffffffffff8082111561131d578283fd5b818401915084601f830112611330578283fd5b81518181111561133e578384fd5b604051601f8201601f19168101602001838111828210171561135e578586fd5b604052818152838201602001871015611375578485fd5b61138682602083016020870161179b565b9695505050505050565b6000602082840312156113a1578081fd5b5051919050565b600080604083850312156113ba578182fd5b50508035926020909101359150565b600080604083850312156113db578182fd5b6113e58484611140565b91506113f48460208501611158565b90509250929050565b6000806040838503121561140f578182fd5b823560ff8116811461141f578283fd5b9150602083013567ffffffffffffffff81168114611220578182fd5b6000815180845261145381602086016020860161179b565b601f01601f19169290920160200192915050565b6000825161147981846020870161179b565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b600060018060a01b038516825283602083015260606040830152610f03606083018461143b565b600060018060a01b03808a16835260ff8916602084015287604084015260e0606084015261151360e084018861143b565b81871660808501528560a085015283810360c0850152611533818661143b565b9b9a5050505050505050505050565b901515815260200190565b90815260200190565b6001600160e01b031991909116815260200190565b6020808252603b908201527f70726f766964656420636f6e74726163742063616e6e6f74206265206c6f636b60408201527f4d696e74556e6c6f636b61626c6520616e64206275726e61626c650000000000606082015260800190565b6020808252601e908201527f73656e646572206d7573742062652062726964676520636f6e74726163740000604082015260600190565b6020808252603b908201527f70726f766964656420636f6e74726163742063616e6e6f74206265206275726e60408201527f61626c6520616e64206c6f636b4d696e74556e6c6f636b61626c650000000000606082015260800190565b60208082526024908201527f70726f766964656420636f6e7472616374206973206e6f742077686974656c696040820152631cdd195960e21b606082015260800190565b60208082526028908201527f70726f766964656420746f6b656e41646472657373206973206e6f74207768696040820152671d195b1a5cdd195960c21b606082015260800190565b6000602082526116f8835161178f565b602083015260ff602084015116604083015260408301516060830152606083015160e0608084015261172e61010084018261143b565b905060018060a01b0360808501511660a084015260a084015160c084015260c0840151601f198483030160e0850152610f03828261143b565b60008085851115611776578182fd5b83861115611782578182fd5b5050820193919092039150565b6001600160a01b031690565b60005b838110156117b657818101518382015260200161179e565b838111156117c5576000848401525b50505050565b6001600160a01b038116811461026e57600080fdfea2646970667358221220ff8f4d34265b65cb35592592871139fe7b9489888d5643523f6426be2776884e64736f6c634300060c0033",
}

// ERC721HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721HandlerMetaData.ABI instead.
var ERC721HandlerABI = ERC721HandlerMetaData.ABI

// ERC721HandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC721HandlerMetaData.Bin instead.
var ERC721HandlerBin = ERC721HandlerMetaData.Bin

// DeployERC721Handler deploys a new Ethereum contract, binding an instance of ERC721Handler to it.
func DeployERC721Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, burnableContractAddresses []common.Address, lockMintUnlockableContractAddresses []common.Address) (common.Address, *types.Transaction, *ERC721Handler, error) {
	parsed, err := ERC721HandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC721HandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses, burnableContractAddresses, lockMintUnlockableContractAddresses)
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
func (_ERC721Handler *ERC721HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_ERC721Handler *ERC721HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.BurnList(&_ERC721Handler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.BurnList(&_ERC721Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
func (_ERC721Handler *ERC721HandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 uint8, arg1 uint64) (struct {
	TokenAddress                common.Address
	DestinationChainID          uint8
	ResourceID                  [32]byte
	DestinationRecipientAddress []byte
	Depositer                   common.Address
	TokenID                     *big.Int
	MetaData                    []byte
}, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	outstruct := new(struct {
		TokenAddress                common.Address
		DestinationChainID          uint8
		ResourceID                  [32]byte
		DestinationRecipientAddress []byte
		Depositer                   common.Address
		TokenID                     *big.Int
		MetaData                    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.DestinationChainID = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.ResourceID = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.DestinationRecipientAddress = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.Depositer = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.TokenID = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.MetaData = *abi.ConvertType(out[6], new([]byte)).(*[]byte)

	return *outstruct, err

}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
func (_ERC721Handler *ERC721HandlerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	TokenAddress                common.Address
	DestinationChainID          uint8
	ResourceID                  [32]byte
	DestinationRecipientAddress []byte
	Depositer                   common.Address
	TokenID                     *big.Int
	MetaData                    []byte
}, error) {
	return _ERC721Handler.Contract.DepositRecords(&_ERC721Handler.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
func (_ERC721Handler *ERC721HandlerCallerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	TokenAddress                common.Address
	DestinationChainID          uint8
	ResourceID                  [32]byte
	DestinationRecipientAddress []byte
	Depositer                   common.Address
	TokenID                     *big.Int
	MetaData                    []byte
}, error) {
	return _ERC721Handler.Contract.DepositRecords(&_ERC721Handler.CallOpts, arg0, arg1)
}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCaller) LockMintUnlockList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_lockMintUnlockList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerSession) LockMintUnlockList(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.LockMintUnlockList(&_ERC721Handler.CallOpts, arg0)
}

// LockMintUnlockList is a free data retrieval call binding the contract method 0xb1f06463.
//
// Solidity: function _lockMintUnlockList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCallerSession) LockMintUnlockList(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.LockMintUnlockList(&_ERC721Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

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

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,bytes32,bytes,address,uint256,bytes))
func (_ERC721Handler *ERC721HandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositNonce uint64, destId uint8) (ERC721HandlerDepositRecord, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "getDepositRecord", depositNonce, destId)

	if err != nil {
		return *new(ERC721HandlerDepositRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC721HandlerDepositRecord)).(*ERC721HandlerDepositRecord)

	return out0, err

}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,bytes32,bytes,address,uint256,bytes))
func (_ERC721Handler *ERC721HandlerSession) GetDepositRecord(depositNonce uint64, destId uint8) (ERC721HandlerDepositRecord, error) {
	return _ERC721Handler.Contract.GetDepositRecord(&_ERC721Handler.CallOpts, depositNonce, destId)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,bytes32,bytes,address,uint256,bytes))
func (_ERC721Handler *ERC721HandlerCallerSession) GetDepositRecord(depositNonce uint64, destId uint8) (ERC721HandlerDepositRecord, error) {
	return _ERC721Handler.Contract.GetDepositRecord(&_ERC721Handler.CallOpts, depositNonce, destId)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "deposit", resourceID, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ExecuteProposal(&_ERC721Handler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ExecuteProposal(&_ERC721Handler.TransactOpts, resourceID, data)
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

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetBurnable(&_ERC721Handler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetBurnable(&_ERC721Handler.TransactOpts, contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetLockMintUnlockable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setLockMintUnlockable", contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerSession) SetLockMintUnlockable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetLockMintUnlockable(&_ERC721Handler.TransactOpts, contractAddress)
}

// SetLockMintUnlockable is a paid mutator transaction binding the contract method 0x00d8e7ee.
//
// Solidity: function setLockMintUnlockable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetLockMintUnlockable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetLockMintUnlockable(&_ERC721Handler.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetResource(&_ERC721Handler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetResource(&_ERC721Handler.TransactOpts, resourceID, contractAddress)
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
