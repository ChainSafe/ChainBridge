// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package ERC20Handler

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ERC20HandlerABI is the input ABI used to generate the binding from.
const ERC20HandlerABI = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"constant":true,"inputs":[],"name":"isOwner","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"renounceOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"address","name":"","type":"address"}],"name":"synthetics","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"address","name":"","type":"address"}],"name":"whitelist","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"uint256","name":"_originChain","type":"uint256"},{"internalType":"address","name":"_originTokenAddress","type":"address"},{"internalType":"address","name":"_to","type":"address"},{"internalType":"uint256","name":"_value","type":"uint256"}],"name":"executeTransfer","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`

// ERC20HandlerBin is the compiled bytecode used for deploying new contracts.
const ERC20HandlerBin = `6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a361239e806100cf6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063834cdf351161005b578063834cdf35146101a85780638da5cb5b146102205780638f32d59b1461026a578063f2fde38b1461028c5761007d565b80634b25bfce14610082578063593fd4bf14610110578063715018a61461019e575b600080fd5b6100ce6004803603604081101561009857600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102d0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61015c6004803603604081101561012657600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610312565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101a6610354565b005b61021e600480360360808110156101be57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610424565b005b61022861098d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102726109b6565b604051808215151515815260200191505060405180910390f35b6102ce600480360360208110156102a257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a0d565b005b60016020528160005260406000206020528060005260406000206000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60026020528160005260406000206020528060005260406000206000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61035c6109b6565b61036557600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b61042c6109b6565b61043557600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166001600086815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461061d5760006001600086815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008190508073ffffffffffffffffffffffffffffffffffffffff166340c10f1985856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156105da57600080fd5b505af11580156105ee573d6000803e3d6000fd5b505050506040513d602081101561060457600080fd5b8101908080519060200190929190505050505050610987565b600073ffffffffffffffffffffffffffffffffffffffff166002600086815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146108055760006002600086815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008190508073ffffffffffffffffffffffffffffffffffffffff166340c10f1985856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156107c257600080fd5b505af11580156107d6573d6000803e3d6000fd5b505050506040513d60208110156107ec57600080fd5b8101908080519060200190929190505050505050610986565b600060405161081390610b22565b604051809103906000f08015801561082f573d6000803e3d6000fd5b509050806002600087815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561094857600080fd5b505af115801561095c573d6000803e3d6000fd5b505050506040513d602081101561097257600080fd5b810190808051906020019092919050505050505b5b50505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b610a156109b6565b610a1e57600080fd5b610a2781610a2a565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a6457600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61183a80610b308339019056fe608060405262000024620000186200002a60201b60201c565b6200003260201b60201c565b62000257565b600033905090565b6200004d8160036200009360201b620012701790919060201c565b8073ffffffffffffffffffffffffffffffffffffffff167f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f660405160405180910390a250565b620000a582826200017760201b60201c565b1562000119576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000200576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180620018186022913960400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6115b180620002676000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063983b2d5611610071578063983b2d56146102e7578063986502751461032b578063a457c2d714610335578063a9059cbb1461039b578063aa271e1a14610401578063dd62ed3e1461045d576100b4565b8063095ea7b3146100b957806318160ddd1461011f57806323b872dd1461013d57806339509351146101c357806340c10f191461022957806370a082311461028f575b600080fd5b610105600480360360408110156100cf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506104d5565b604051808215151515815260200191505060405180910390f35b6101276104ec565b6040518082815260200191505060405180910390f35b6101a96004803603606081101561015357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506104f6565b604051808215151515815260200191505060405180910390f35b61020f600480360360408110156101d957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506105c1565b604051808215151515815260200191505060405180910390f35b6102756004803603604081101561023f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610666565b604051808215151515815260200191505060405180910390f35b6102d1600480360360208110156102a557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106e1565b6040518082815260200191505060405180910390f35b610329600480360360208110156102fd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610729565b005b61033361079a565b005b6103816004803603604081101561034b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506107ac565b604051808215151515815260200191505060405180910390f35b6103e7600480360360408110156103b157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061086b565b604051808215151515815260200191505060405180910390f35b6104436004803603602081101561041757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610882565b604051808215151515815260200191505060405180910390f35b6104bf6004803603604081101561047357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061089f565b6040518082815260200191505060405180910390f35b60006104e2338484610926565b6001905092915050565b6000600254905090565b6000610503848484610b1d565b6105b684336105b1856040518060600160405280602881526020016114c560289139600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610dd39092919063ffffffff16565b610926565b600190509392505050565b600061065c338461065785600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e9390919063ffffffff16565b610926565b6001905092915050565b6000610678610673610f1b565b610882565b6106cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806114746030913960400191505060405180910390fd5b6106d78383610f23565b6001905092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610739610734610f1b565b610882565b61078e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806114746030913960400191505060405180910390fd5b610797816110de565b50565b6107aa6107a5610f1b565b611138565b565b6000610861338461085c8560405180606001604052806025815260200161155860259139600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610dd39092919063ffffffff16565b610926565b6001905092915050565b6000610878338484610b1d565b6001905092915050565b600061089882600361119290919063ffffffff16565b9050919050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156109ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806115346024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610a32576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018061142c6022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610ba3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061150f6025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610c29576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806114096023913960400191505060405180910390fd5b610c948160405180606001604052806026815260200161144e602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610dd39092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610d27816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e9390919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000838311158290610e80576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610e45578082015181840152602081019050610e2a565b50505050905090810190601f168015610e725780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b600080828401905083811015610f11576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610fc6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b610fdb81600254610e9390919063ffffffff16565b600281905550611032816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e9390919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b6110f281600361127090919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f660405160405180910390a250565b61114c81600361134b90919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669260405160405180910390a250565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611219576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806114ed6022913960400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b61127a8282611192565b156112ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6113558282611192565b6113aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806114a46021913960400191505060405180910390fd5b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e63654d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c6545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e6365526f6c65733a206163636f756e7420697320746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a7231582019576d72d96e6bfcccfae245889f59766d3c143a569e094f31da6ff72a26962564736f6c634300050c0032526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373a265627a7a72315820d5efe6612cca94d60b96400132ca95dc92fa995d7a13e61f6fab97154236a0ec64736f6c634300050c0032`

// DeployERC20Handler deploys a new Ethereum contract, binding an instance of ERC20Handler to it.
func DeployERC20Handler(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Handler, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20HandlerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}}, nil
}

// ERC20Handler is an auto generated Go binding around an Ethereum contract.
type ERC20Handler struct {
	ERC20HandlerCaller     // Read-only binding to the contract
	ERC20HandlerTransactor // Write-only binding to the contract
}

// ERC20HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactor struct {
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
	contract, err := bindERC20Handler(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}}, nil
}

// NewERC20HandlerCaller creates a new read-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC20HandlerCaller, error) {
	contract, err := bindERC20Handler(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerCaller{contract: contract}, nil
}

// NewERC20HandlerTransactor creates a new write-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20HandlerTransactor, error) {
	contract, err := bindERC20Handler(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerTransactor{contract: contract}, nil
}

// bindERC20Handler binds a generic wrapper to an already deployed contract.
func bindERC20Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ERC20Handler *ERC20HandlerSession) IsOwner() (bool, error) {
	return _ERC20Handler.Contract.IsOwner(&_ERC20Handler.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) IsOwner() (bool, error) {
	return _ERC20Handler.Contract.IsOwner(&_ERC20Handler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ERC20Handler *ERC20HandlerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ERC20Handler *ERC20HandlerSession) Owner() (common.Address, error) {
	return _ERC20Handler.Contract.Owner(&_ERC20Handler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) Owner() (common.Address, error) {
	return _ERC20Handler.Contract.Owner(&_ERC20Handler.CallOpts)
}

// Synthetics is a free data retrieval call binding the contract method 0x593fd4bf.
//
// Solidity: function synthetics( uint256,  address) constant returns(address)
func (_ERC20Handler *ERC20HandlerCaller) Synthetics(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "synthetics", arg0, arg1)
	return *ret0, err
}

// Synthetics is a free data retrieval call binding the contract method 0x593fd4bf.
//
// Solidity: function synthetics( uint256,  address) constant returns(address)
func (_ERC20Handler *ERC20HandlerSession) Synthetics(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC20Handler.Contract.Synthetics(&_ERC20Handler.CallOpts, arg0, arg1)
}

// Synthetics is a free data retrieval call binding the contract method 0x593fd4bf.
//
// Solidity: function synthetics( uint256,  address) constant returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) Synthetics(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC20Handler.Contract.Synthetics(&_ERC20Handler.CallOpts, arg0, arg1)
}

// Whitelist is a free data retrieval call binding the contract method 0x4b25bfce.
//
// Solidity: function whitelist( uint256,  address) constant returns(address)
func (_ERC20Handler *ERC20HandlerCaller) Whitelist(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "whitelist", arg0, arg1)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x4b25bfce.
//
// Solidity: function whitelist( uint256,  address) constant returns(address)
func (_ERC20Handler *ERC20HandlerSession) Whitelist(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC20Handler.Contract.Whitelist(&_ERC20Handler.CallOpts, arg0, arg1)
}

// Whitelist is a free data retrieval call binding the contract method 0x4b25bfce.
//
// Solidity: function whitelist( uint256,  address) constant returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) Whitelist(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC20Handler.Contract.Whitelist(&_ERC20Handler.CallOpts, arg0, arg1)
}

// ExecuteTransfer is a paid mutator transaction binding the contract method 0x834cdf35.
//
// Solidity: function executeTransfer(_originChain uint256, _originTokenAddress address, _to address, _value uint256) returns()
func (_ERC20Handler *ERC20HandlerTransactor) ExecuteTransfer(opts *bind.TransactOpts, _originChain *big.Int, _originTokenAddress common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "executeTransfer", _originChain, _originTokenAddress, _to, _value)
}

// ExecuteTransfer is a paid mutator transaction binding the contract method 0x834cdf35.
//
// Solidity: function executeTransfer(_originChain uint256, _originTokenAddress address, _to address, _value uint256) returns()
func (_ERC20Handler *ERC20HandlerSession) ExecuteTransfer(_originChain *big.Int, _originTokenAddress common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteTransfer(&_ERC20Handler.TransactOpts, _originChain, _originTokenAddress, _to, _value)
}

// ExecuteTransfer is a paid mutator transaction binding the contract method 0x834cdf35.
//
// Solidity: function executeTransfer(_originChain uint256, _originTokenAddress address, _to address, _value uint256) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) ExecuteTransfer(_originChain *big.Int, _originTokenAddress common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteTransfer(&_ERC20Handler.TransactOpts, _originChain, _originTokenAddress, _to, _value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Handler *ERC20HandlerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Handler *ERC20HandlerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20Handler.Contract.RenounceOwnership(&_ERC20Handler.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20Handler.Contract.RenounceOwnership(&_ERC20Handler.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ERC20Handler *ERC20HandlerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ERC20Handler *ERC20HandlerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.TransferOwnership(&_ERC20Handler.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.TransferOwnership(&_ERC20Handler.TransactOpts, newOwner)
}
