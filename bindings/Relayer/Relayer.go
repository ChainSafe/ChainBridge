// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Relayer

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

// RelayerABI is the input ABI used to generate the binding from.
const RelayerABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"initialRelayers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"initialRelayerThreshold\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"RelayerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIRelayer.RelayerActionType\",\"name\":\"relayerActionType\",\"type\":\"uint8\"}],\"name\":\"RelayerProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIRelayer.VoteStatus\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"RelayerProposalVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"RelayerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"RelayerThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposedValue\",\"type\":\"uint256\"}],\"name\":\"RelayerThresholdProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumIRelayer.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"RelayerThresholdProposalVote\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_relayerProposals\",\"outputs\":[{\"internalType\":\"enumIRelayer.RelayerActionType\",\"name\":\"_action\",\"type\":\"uint8\"},{\"internalType\":\"enumIRelayer.VoteStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_relayerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_relayers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalRelayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalRelayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRelayerThresholdProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"enumIRelayer.VoteStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"}],\"name\":\"getRelayerProposal\",\"outputs\":[{\"internalType\":\"enumIRelayer.RelayerActionType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"enumIRelayer.VoteStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"},{\"internalType\":\"enumIRelayer.RelayerActionType\",\"name\":\"action\",\"type\":\"uint8\"}],\"name\":\"voteRelayerProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposedValue\",\"type\":\"uint256\"}],\"name\":\"createRelayerThresholdProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIRelayer.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"voteRelayerThresholdProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RelayerBin is the compiled bytecode used for deploying new contracts.
var RelayerBin = "0x60806040523480156200001157600080fd5b506040516200202238038062002022833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660208202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620000c6578082015181840152602081019050620000a9565b505050509050016040526020018051906020019092919050505060005b8251811015620001205762000112838281518110620000fe57fe5b60200260200101516200013060201b60201c565b8080600101915050620000e3565b50806000819055505050620001e0565b6001600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600081548092919060010191905055508073ffffffffffffffffffffffffffffffffffffffff167f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c560405160405180910390a250565b611e3280620001f06000396000f3fe608060405234801561001057600080fd5b50600436106100a85760003560e01c80635f31b69c116100715780635f31b69c1461033c578063802aabe8146103ff578063933b46671461041d578063d7a9cd791461043b578063df26906014610459578063e9cdaead14610487576100a8565b806251f244146100ad57806330fddf01146101b85780634d4ebd7a14610233578063541d5548146102845780635d0db4db146102e0575b600080fd5b6100ef600480360360208110156100c357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104b8565b604051808560018111156100ff57fe5b60ff168152602001806020018060200184600181111561011b57fe5b60ff168152602001838103835286818151815260200191508051906020019060200280838360005b8381101561015e578082015181840152602081019050610143565b50505050905001838103825285818151815260200191508051906020019060200280838360005b838110156101a0578082015181840152602081019050610185565b50505050905001965050505050505060405180910390f35b6101fa600480360360208110156101ce57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106b3565b6040518083600181111561020a57fe5b60ff16815260200182600181111561021e57fe5b60ff1681526020019250505060405180910390f35b6102826004803603604081101561024957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff1690602001909291905050506106f1565b005b6102c66004803603602081101561029a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610eec565b604051808215151515815260200191505060405180910390f35b610322600480360360208110156102f657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f42565b604051808215151515815260200191505060405180910390f35b610344610f62565b60405180858152602001806020018060200184600181111561036257fe5b60ff168152602001838103835286818151815260200191508051906020019060200280838360005b838110156103a557808201518184015260208101905061038a565b50505050905001838103825285818151815260200191508051906020019060200280838360005b838110156103e75780820151818401526020810190506103cc565b50505050905001965050505050505060405180910390f35b6104076110a9565b6040518082815260200191505060405180910390f35b6104256110af565b6040518082815260200191505060405180910390f35b6104436110b9565b6040518082815260200191505060405180910390f35b6104856004803603602081101561046f57600080fd5b81019080803590602001909291905050506110bf565b005b6104b66004803603602081101561049d57600080fd5b81019080803560ff1690602001909291905050506114c4565b005b600060608060006104c7611c21565b600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060800160405290816000820160009054906101000a900460ff16600181111561052f57fe5b600181111561053a57fe5b8152602001600282018054806020026020016040519081016040528092919081815260200182805480156105c357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610579575b505050505081526020016003820180548060200260200160405190810160405280929190818152602001828054801561065157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610607575b505050505081526020016004820160009054906101000a900460ff16600181111561067857fe5b600181111561068357fe5b81525050905080600001518160200151826040015183606001518292508191509450945094509450509193509193565b60086020528060005260406000206000915090508060000160009054906101000a900460ff16908060040160009054906101000a900460ff16905082565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166107b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f73656e646572206973206e6f7420612072656c6179657200000000000000000081525060200191505060405180910390fd5b6000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600182600181111561080157fe5b60ff16111561085b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611db96023913960400191505060405180910390fd5b600015158160010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514610906576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180611d2d6025913960400191505060405180910390fd5b6000600181111561091357fe5b82600181111561091f57fe5b14156109cc57600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166109c7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180611ddc6021913960400191505060405180910390fd5b610a70565b600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610a6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180611d526025913960400191505060405180910390fd5b5b60006001811115610a7d57fe5b8160040160009054906101000a900460ff166001811115610a9a57fe5b1415610cb2576040518060800160405280836001811115610ab757fe5b81526020016001604051908082528060200260200182016040528015610aec5781602001602082028036833780820191505090505b5081526020016000604051908082528060200260200182016040528015610b225781602001602082028036833780820191505090505b508152602001600180811115610b3457fe5b815250600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690836001811115610b9a57fe5b02179055506020820151816002019080519060200190610bbb929190611c5f565b506040820151816003019080519060200190610bd8929190611c5f565b5060608201518160040160006101000a81548160ff02191690836001811115610bfd57fe5b02179055509050503381600201600081548110610c1657fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816001811115610c6a57fe5b8373ffffffffffffffffffffffffffffffffffffffff167fb12a1fbc5d87bccd40c7ac1f669d842fa9fdfe7f65d2988ceac40543b646807e60405160405180910390a3610d18565b80600201339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b60018160010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508060040160009054906101000a900460ff166001811115610d8f57fe5b8373ffffffffffffffffffffffffffffffffffffffff167fb6a4771df73f6cb79535dd475a69a813362d4b13633fdd88d406a3ed78a6033260405160405180910390a36001600054111580610ded5750600054816002018054905010155b15610ee75760008160040160006101000a81548160ff02191690836001811115610e1357fe5b021790555060006001811115610e2557fe5b8160000160009054906101000a900460ff166001811115610e4257fe5b1415610e9957610e51836119b6565b8273ffffffffffffffffffffffffffffffffffffffff167f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b60405160405180910390a2610ee6565b610ea283611a67565b8273ffffffffffffffffffffffffffffffffffffffff167f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c560405160405180910390a25b5b505050565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60076020528060005260406000206000915054906101000a900460ff1681565b60006060806000600260000154600280016002600301600260040160009054906101000a900460ff168280548060200260200160405190810160405280929190818152602001828054801561100c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610fc2575b505050505092508180548060200260200160405190810160405280929190818152602001828054801561109457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161104a575b50505050509150935093509350935090919293565b60015481565b6000600154905090565b60005481565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661117e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f73656e646572206973206e6f7420612072656c6179657200000000000000000081525060200191505060405180910390fd5b6000600181111561118b57fe5b600260040160009054906101000a900460ff1660018111156111a957fe5b1461121c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f612070726f706f73616c2069732063757272656e746c7920616374697665000081525060200191505060405180910390fd5b600154811115611277576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526042815260200180611d776042913960600191505060405180910390fd5b604051806080016040528082815260200160016040519080825280602002602001820160405280156112b85781602001602082028036833780820191505090505b50815260200160006040519080825280602002602001820160405280156112ee5781602001602082028036833780820191505090505b50815260200160018081111561130057fe5b815250600260008201518160000155602082015181600201908051906020019061132b929190611c5f565b506040820151816003019080519060200190611348929190611c5f565b5060608201518160040160006101000a81548160ff0219169083600181111561136d57fe5b02179055509050506001600054116113e0576002600001546000819055506000600260040160006101000a81548160ff021916908360018111156113ad57fe5b0217905550807fa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c860405160405180910390a25b33600280016000815481106113f157fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550807f79a7cdf066fccb1627ec87a3dd0bf8dcb3a186313f941c8bf80ed979aa62d38d60405160405180910390a250565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611583576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f73656e646572206973206e6f7420612072656c6179657200000000000000000081525060200191505060405180910390fd5b600181600181111561159157fe5b60ff161115611608576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f766f7465206f7574206f662074686520766f746520656e756d2072616e67650081525060200191505060405180910390fd5b60018081111561161457fe5b600260040160009054906101000a900460ff16600181111561163257fe5b146116a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f6e6f2070726f706f73616c2069732063757272656e746c79206163746976650081525060200191505060405180910390fd5b600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611768576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f72656c617965722068617320616c726561647920766f7465640000000000000081525060200191505060405180910390fd5b60018081111561177457fe5b81600181111561178057fe5b14156117f05760028001339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611857565b6002600301339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b6001600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508060018111156118be57fe5b7fe6124fe2b7a19e7cdd807fb16247ec258d8fb0bfde2949b023bcb0c1eea19cb160405160405180910390a2600054600280018054905010611964576002600001546000819055506000600260040160006101000a81548160ff0219169083600181111561192857fe5b02179055506002600001547fa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c860405160405180910390a26119b3565b600054611984600260030180549050600154611b1790919063ffffffff16565b10156119b2576000600260040160006101000a81548160ff021916908360018111156119ac57fe5b02179055505b5b50565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160008154809291906001900391905055508073ffffffffffffffffffffffffffffffffffffffff167f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b60405160405180910390a250565b6001600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600081548092919060010191905055508073ffffffffffffffffffffffffffffffffffffffff167f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c560405160405180910390a250565b6000611b5983836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611b61565b905092915050565b6000838311158290611c0e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611bd3578082015181840152602081019050611bb8565b50505050905090810190601f168015611c005780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b604051806080016040528060006001811115611c3957fe5b8152602001606081526020016060815260200160006001811115611c5957fe5b81525090565b828054828255906000526020600020908101928215611cd8579160200282015b82811115611cd75782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190611c7f565b5b509050611ce59190611ce9565b5090565b611d2991905b80821115611d2557600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101611cef565b5090565b9056fe72656c617965722068617320616c726561647920766f746564206f6e2070726f706f73616c70726f706f736564206164647265737320697320616c726561647920612072656c6179657270726f706f7365642076616c75652063616e6e6f742062652067726561746572207468616e2074686520746f74616c206e756d626572206f662072656c6179657273616374696f6e206f7574206f662074686520616374696f6e20656e756d2072616e676570726f706f7365642061646472657373206973206e6f7420612072656c61796572a2646970667358221220d871218f009df167ebaab13b975cb9f5f2b8355144c315b764ac2c2581a0e05f64736f6c63430006040033"

// DeployRelayer deploys a new Ethereum contract, binding an instance of Relayer to it.
func DeployRelayer(auth *bind.TransactOpts, backend bind.ContractBackend, initialRelayers []common.Address, initialRelayerThreshold *big.Int) (common.Address, *types.Transaction, *Relayer, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RelayerBin), backend, initialRelayers, initialRelayerThreshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Relayer{RelayerCaller: RelayerCaller{contract: contract}, RelayerTransactor: RelayerTransactor{contract: contract}, RelayerFilterer: RelayerFilterer{contract: contract}}, nil
}

// Relayer is an auto generated Go binding around an Ethereum contract.
type Relayer struct {
	RelayerCaller     // Read-only binding to the contract
	RelayerTransactor // Write-only binding to the contract
	RelayerFilterer   // Log filterer for contract events
}

// RelayerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayerSession struct {
	Contract     *Relayer          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayerCallerSession struct {
	Contract *RelayerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RelayerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayerTransactorSession struct {
	Contract     *RelayerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RelayerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayerRaw struct {
	Contract *Relayer // Generic contract binding to access the raw methods on
}

// RelayerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayerCallerRaw struct {
	Contract *RelayerCaller // Generic read-only contract binding to access the raw methods on
}

// RelayerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayerTransactorRaw struct {
	Contract *RelayerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayer creates a new instance of Relayer, bound to a specific deployed contract.
func NewRelayer(address common.Address, backend bind.ContractBackend) (*Relayer, error) {
	contract, err := bindRelayer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relayer{RelayerCaller: RelayerCaller{contract: contract}, RelayerTransactor: RelayerTransactor{contract: contract}, RelayerFilterer: RelayerFilterer{contract: contract}}, nil
}

// NewRelayerCaller creates a new read-only instance of Relayer, bound to a specific deployed contract.
func NewRelayerCaller(address common.Address, caller bind.ContractCaller) (*RelayerCaller, error) {
	contract, err := bindRelayer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerCaller{contract: contract}, nil
}

// NewRelayerTransactor creates a new write-only instance of Relayer, bound to a specific deployed contract.
func NewRelayerTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayerTransactor, error) {
	contract, err := bindRelayer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerTransactor{contract: contract}, nil
}

// NewRelayerFilterer creates a new log filterer instance of Relayer, bound to a specific deployed contract.
func NewRelayerFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayerFilterer, error) {
	contract, err := bindRelayer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayerFilterer{contract: contract}, nil
}

// bindRelayer binds a generic wrapper to an already deployed contract.
func bindRelayer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relayer *RelayerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relayer.Contract.RelayerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayer *RelayerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayer.Contract.RelayerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayer *RelayerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayer.Contract.RelayerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relayer *RelayerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relayer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayer *RelayerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayer *RelayerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayer.Contract.contract.Transact(opts, method, params...)
}

// RelayerProposals is a free data retrieval call binding the contract method 0x30fddf01.
//
// Solidity: function _relayerProposals(address ) constant returns(uint8 _action, uint8 _status)
func (_Relayer *RelayerCaller) RelayerProposals(opts *bind.CallOpts, arg0 common.Address) (struct {
	Action uint8
	Status uint8
}, error) {
	ret := new(struct {
		Action uint8
		Status uint8
	})
	out := ret
	err := _Relayer.contract.Call(opts, out, "_relayerProposals", arg0)
	return *ret, err
}

// RelayerProposals is a free data retrieval call binding the contract method 0x30fddf01.
//
// Solidity: function _relayerProposals(address ) constant returns(uint8 _action, uint8 _status)
func (_Relayer *RelayerSession) RelayerProposals(arg0 common.Address) (struct {
	Action uint8
	Status uint8
}, error) {
	return _Relayer.Contract.RelayerProposals(&_Relayer.CallOpts, arg0)
}

// RelayerProposals is a free data retrieval call binding the contract method 0x30fddf01.
//
// Solidity: function _relayerProposals(address ) constant returns(uint8 _action, uint8 _status)
func (_Relayer *RelayerCallerSession) RelayerProposals(arg0 common.Address) (struct {
	Action uint8
	Status uint8
}, error) {
	return _Relayer.Contract.RelayerProposals(&_Relayer.CallOpts, arg0)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() constant returns(uint256)
func (_Relayer *RelayerCaller) RelayerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayer.contract.Call(opts, out, "_relayerThreshold")
	return *ret0, err
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() constant returns(uint256)
func (_Relayer *RelayerSession) RelayerThreshold() (*big.Int, error) {
	return _Relayer.Contract.RelayerThreshold(&_Relayer.CallOpts)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() constant returns(uint256)
func (_Relayer *RelayerCallerSession) RelayerThreshold() (*big.Int, error) {
	return _Relayer.Contract.RelayerThreshold(&_Relayer.CallOpts)
}

// Relayers is a free data retrieval call binding the contract method 0x5d0db4db.
//
// Solidity: function _relayers(address ) constant returns(bool)
func (_Relayer *RelayerCaller) Relayers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relayer.contract.Call(opts, out, "_relayers", arg0)
	return *ret0, err
}

// Relayers is a free data retrieval call binding the contract method 0x5d0db4db.
//
// Solidity: function _relayers(address ) constant returns(bool)
func (_Relayer *RelayerSession) Relayers(arg0 common.Address) (bool, error) {
	return _Relayer.Contract.Relayers(&_Relayer.CallOpts, arg0)
}

// Relayers is a free data retrieval call binding the contract method 0x5d0db4db.
//
// Solidity: function _relayers(address ) constant returns(bool)
func (_Relayer *RelayerCallerSession) Relayers(arg0 common.Address) (bool, error) {
	return _Relayer.Contract.Relayers(&_Relayer.CallOpts, arg0)
}

// TotalRelayers is a free data retrieval call binding the contract method 0x802aabe8.
//
// Solidity: function _totalRelayers() constant returns(uint256)
func (_Relayer *RelayerCaller) TotalRelayers(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayer.contract.Call(opts, out, "_totalRelayers")
	return *ret0, err
}

// TotalRelayers is a free data retrieval call binding the contract method 0x802aabe8.
//
// Solidity: function _totalRelayers() constant returns(uint256)
func (_Relayer *RelayerSession) TotalRelayers() (*big.Int, error) {
	return _Relayer.Contract.TotalRelayers(&_Relayer.CallOpts)
}

// TotalRelayers is a free data retrieval call binding the contract method 0x802aabe8.
//
// Solidity: function _totalRelayers() constant returns(uint256)
func (_Relayer *RelayerCallerSession) TotalRelayers() (*big.Int, error) {
	return _Relayer.Contract.TotalRelayers(&_Relayer.CallOpts)
}

// GetCurrentRelayerThresholdProposal is a free data retrieval call binding the contract method 0x5f31b69c.
//
// Solidity: function getCurrentRelayerThresholdProposal() constant returns(uint256, address[], address[], uint8)
func (_Relayer *RelayerCaller) GetCurrentRelayerThresholdProposal(opts *bind.CallOpts) (*big.Int, []common.Address, []common.Address, uint8, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([]common.Address)
		ret2 = new([]common.Address)
		ret3 = new(uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Relayer.contract.Call(opts, out, "getCurrentRelayerThresholdProposal")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetCurrentRelayerThresholdProposal is a free data retrieval call binding the contract method 0x5f31b69c.
//
// Solidity: function getCurrentRelayerThresholdProposal() constant returns(uint256, address[], address[], uint8)
func (_Relayer *RelayerSession) GetCurrentRelayerThresholdProposal() (*big.Int, []common.Address, []common.Address, uint8, error) {
	return _Relayer.Contract.GetCurrentRelayerThresholdProposal(&_Relayer.CallOpts)
}

// GetCurrentRelayerThresholdProposal is a free data retrieval call binding the contract method 0x5f31b69c.
//
// Solidity: function getCurrentRelayerThresholdProposal() constant returns(uint256, address[], address[], uint8)
func (_Relayer *RelayerCallerSession) GetCurrentRelayerThresholdProposal() (*big.Int, []common.Address, []common.Address, uint8, error) {
	return _Relayer.Contract.GetCurrentRelayerThresholdProposal(&_Relayer.CallOpts)
}

// GetRelayerProposal is a free data retrieval call binding the contract method 0x0051f244.
//
// Solidity: function getRelayerProposal(address proposedAddress) constant returns(uint8, address[], address[], uint8)
func (_Relayer *RelayerCaller) GetRelayerProposal(opts *bind.CallOpts, proposedAddress common.Address) (uint8, []common.Address, []common.Address, uint8, error) {
	var (
		ret0 = new(uint8)
		ret1 = new([]common.Address)
		ret2 = new([]common.Address)
		ret3 = new(uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Relayer.contract.Call(opts, out, "getRelayerProposal", proposedAddress)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetRelayerProposal is a free data retrieval call binding the contract method 0x0051f244.
//
// Solidity: function getRelayerProposal(address proposedAddress) constant returns(uint8, address[], address[], uint8)
func (_Relayer *RelayerSession) GetRelayerProposal(proposedAddress common.Address) (uint8, []common.Address, []common.Address, uint8, error) {
	return _Relayer.Contract.GetRelayerProposal(&_Relayer.CallOpts, proposedAddress)
}

// GetRelayerProposal is a free data retrieval call binding the contract method 0x0051f244.
//
// Solidity: function getRelayerProposal(address proposedAddress) constant returns(uint8, address[], address[], uint8)
func (_Relayer *RelayerCallerSession) GetRelayerProposal(proposedAddress common.Address) (uint8, []common.Address, []common.Address, uint8, error) {
	return _Relayer.Contract.GetRelayerProposal(&_Relayer.CallOpts, proposedAddress)
}

// CreateRelayerThresholdProposal is a paid mutator transaction binding the contract method 0xdf269060.
//
// Solidity: function createRelayerThresholdProposal(uint256 proposedValue) returns()
func (_Relayer *RelayerTransactor) CreateRelayerThresholdProposal(opts *bind.TransactOpts, proposedValue *big.Int) (*types.Transaction, error) {
	return _Relayer.contract.Transact(opts, "createRelayerThresholdProposal", proposedValue)
}

// CreateRelayerThresholdProposal is a paid mutator transaction binding the contract method 0xdf269060.
//
// Solidity: function createRelayerThresholdProposal(uint256 proposedValue) returns()
func (_Relayer *RelayerSession) CreateRelayerThresholdProposal(proposedValue *big.Int) (*types.Transaction, error) {
	return _Relayer.Contract.CreateRelayerThresholdProposal(&_Relayer.TransactOpts, proposedValue)
}

// CreateRelayerThresholdProposal is a paid mutator transaction binding the contract method 0xdf269060.
//
// Solidity: function createRelayerThresholdProposal(uint256 proposedValue) returns()
func (_Relayer *RelayerTransactorSession) CreateRelayerThresholdProposal(proposedValue *big.Int) (*types.Transaction, error) {
	return _Relayer.Contract.CreateRelayerThresholdProposal(&_Relayer.TransactOpts, proposedValue)
}

// GetTotalRelayers is a paid mutator transaction binding the contract method 0x933b4667.
//
// Solidity: function getTotalRelayers() returns(uint256)
func (_Relayer *RelayerTransactor) GetTotalRelayers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayer.contract.Transact(opts, "getTotalRelayers")
}

// GetTotalRelayers is a paid mutator transaction binding the contract method 0x933b4667.
//
// Solidity: function getTotalRelayers() returns(uint256)
func (_Relayer *RelayerSession) GetTotalRelayers() (*types.Transaction, error) {
	return _Relayer.Contract.GetTotalRelayers(&_Relayer.TransactOpts)
}

// GetTotalRelayers is a paid mutator transaction binding the contract method 0x933b4667.
//
// Solidity: function getTotalRelayers() returns(uint256)
func (_Relayer *RelayerTransactorSession) GetTotalRelayers() (*types.Transaction, error) {
	return _Relayer.Contract.GetTotalRelayers(&_Relayer.TransactOpts)
}

// IsRelayer is a paid mutator transaction binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayerAddress) returns(bool)
func (_Relayer *RelayerTransactor) IsRelayer(opts *bind.TransactOpts, relayerAddress common.Address) (*types.Transaction, error) {
	return _Relayer.contract.Transact(opts, "isRelayer", relayerAddress)
}

// IsRelayer is a paid mutator transaction binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayerAddress) returns(bool)
func (_Relayer *RelayerSession) IsRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _Relayer.Contract.IsRelayer(&_Relayer.TransactOpts, relayerAddress)
}

// IsRelayer is a paid mutator transaction binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayerAddress) returns(bool)
func (_Relayer *RelayerTransactorSession) IsRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _Relayer.Contract.IsRelayer(&_Relayer.TransactOpts, relayerAddress)
}

// VoteRelayerProposal is a paid mutator transaction binding the contract method 0x4d4ebd7a.
//
// Solidity: function voteRelayerProposal(address proposedAddress, uint8 action) returns()
func (_Relayer *RelayerTransactor) VoteRelayerProposal(opts *bind.TransactOpts, proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _Relayer.contract.Transact(opts, "voteRelayerProposal", proposedAddress, action)
}

// VoteRelayerProposal is a paid mutator transaction binding the contract method 0x4d4ebd7a.
//
// Solidity: function voteRelayerProposal(address proposedAddress, uint8 action) returns()
func (_Relayer *RelayerSession) VoteRelayerProposal(proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _Relayer.Contract.VoteRelayerProposal(&_Relayer.TransactOpts, proposedAddress, action)
}

// VoteRelayerProposal is a paid mutator transaction binding the contract method 0x4d4ebd7a.
//
// Solidity: function voteRelayerProposal(address proposedAddress, uint8 action) returns()
func (_Relayer *RelayerTransactorSession) VoteRelayerProposal(proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _Relayer.Contract.VoteRelayerProposal(&_Relayer.TransactOpts, proposedAddress, action)
}

// VoteRelayerThresholdProposal is a paid mutator transaction binding the contract method 0xe9cdaead.
//
// Solidity: function voteRelayerThresholdProposal(uint8 vote) returns()
func (_Relayer *RelayerTransactor) VoteRelayerThresholdProposal(opts *bind.TransactOpts, vote uint8) (*types.Transaction, error) {
	return _Relayer.contract.Transact(opts, "voteRelayerThresholdProposal", vote)
}

// VoteRelayerThresholdProposal is a paid mutator transaction binding the contract method 0xe9cdaead.
//
// Solidity: function voteRelayerThresholdProposal(uint8 vote) returns()
func (_Relayer *RelayerSession) VoteRelayerThresholdProposal(vote uint8) (*types.Transaction, error) {
	return _Relayer.Contract.VoteRelayerThresholdProposal(&_Relayer.TransactOpts, vote)
}

// VoteRelayerThresholdProposal is a paid mutator transaction binding the contract method 0xe9cdaead.
//
// Solidity: function voteRelayerThresholdProposal(uint8 vote) returns()
func (_Relayer *RelayerTransactorSession) VoteRelayerThresholdProposal(vote uint8) (*types.Transaction, error) {
	return _Relayer.Contract.VoteRelayerThresholdProposal(&_Relayer.TransactOpts, vote)
}

// RelayerRelayerAddedIterator is returned from FilterRelayerAdded and is used to iterate over the raw logs and unpacked data for RelayerAdded events raised by the Relayer contract.
type RelayerRelayerAddedIterator struct {
	Event *RelayerRelayerAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RelayerRelayerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerRelayerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RelayerRelayerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RelayerRelayerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerRelayerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerRelayerAdded represents a RelayerAdded event raised by the Relayer contract.
type RelayerRelayerAdded struct {
	RelayerAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRelayerAdded is a free log retrieval operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayerAddress)
func (_Relayer *RelayerFilterer) FilterRelayerAdded(opts *bind.FilterOpts, relayerAddress []common.Address) (*RelayerRelayerAddedIterator, error) {

	var relayerAddressRule []interface{}
	for _, relayerAddressItem := range relayerAddress {
		relayerAddressRule = append(relayerAddressRule, relayerAddressItem)
	}

	logs, sub, err := _Relayer.contract.FilterLogs(opts, "RelayerAdded", relayerAddressRule)
	if err != nil {
		return nil, err
	}
	return &RelayerRelayerAddedIterator{contract: _Relayer.contract, event: "RelayerAdded", logs: logs, sub: sub}, nil
}

// WatchRelayerAdded is a free log subscription operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayerAddress)
func (_Relayer *RelayerFilterer) WatchRelayerAdded(opts *bind.WatchOpts, sink chan<- *RelayerRelayerAdded, relayerAddress []common.Address) (event.Subscription, error) {

	var relayerAddressRule []interface{}
	for _, relayerAddressItem := range relayerAddress {
		relayerAddressRule = append(relayerAddressRule, relayerAddressItem)
	}

	logs, sub, err := _Relayer.contract.WatchLogs(opts, "RelayerAdded", relayerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerRelayerAdded)
				if err := _Relayer.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerAdded is a log parse operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayerAddress)
func (_Relayer *RelayerFilterer) ParseRelayerAdded(log types.Log) (*RelayerRelayerAdded, error) {
	event := new(RelayerRelayerAdded)
	if err := _Relayer.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerRelayerProposalCreatedIterator is returned from FilterRelayerProposalCreated and is used to iterate over the raw logs and unpacked data for RelayerProposalCreated events raised by the Relayer contract.
type RelayerRelayerProposalCreatedIterator struct {
	Event *RelayerRelayerProposalCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RelayerRelayerProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerRelayerProposalCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RelayerRelayerProposalCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RelayerRelayerProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerRelayerProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerRelayerProposalCreated represents a RelayerProposalCreated event raised by the Relayer contract.
type RelayerRelayerProposalCreated struct {
	ProposedAddress   common.Address
	RelayerActionType uint8
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRelayerProposalCreated is a free log retrieval operation binding the contract event 0xb12a1fbc5d87bccd40c7ac1f669d842fa9fdfe7f65d2988ceac40543b646807e.
//
// Solidity: event RelayerProposalCreated(address indexed proposedAddress, uint8 indexed relayerActionType)
func (_Relayer *RelayerFilterer) FilterRelayerProposalCreated(opts *bind.FilterOpts, proposedAddress []common.Address, relayerActionType []uint8) (*RelayerRelayerProposalCreatedIterator, error) {

	var proposedAddressRule []interface{}
	for _, proposedAddressItem := range proposedAddress {
		proposedAddressRule = append(proposedAddressRule, proposedAddressItem)
	}
	var relayerActionTypeRule []interface{}
	for _, relayerActionTypeItem := range relayerActionType {
		relayerActionTypeRule = append(relayerActionTypeRule, relayerActionTypeItem)
	}

	logs, sub, err := _Relayer.contract.FilterLogs(opts, "RelayerProposalCreated", proposedAddressRule, relayerActionTypeRule)
	if err != nil {
		return nil, err
	}
	return &RelayerRelayerProposalCreatedIterator{contract: _Relayer.contract, event: "RelayerProposalCreated", logs: logs, sub: sub}, nil
}

// WatchRelayerProposalCreated is a free log subscription operation binding the contract event 0xb12a1fbc5d87bccd40c7ac1f669d842fa9fdfe7f65d2988ceac40543b646807e.
//
// Solidity: event RelayerProposalCreated(address indexed proposedAddress, uint8 indexed relayerActionType)
func (_Relayer *RelayerFilterer) WatchRelayerProposalCreated(opts *bind.WatchOpts, sink chan<- *RelayerRelayerProposalCreated, proposedAddress []common.Address, relayerActionType []uint8) (event.Subscription, error) {

	var proposedAddressRule []interface{}
	for _, proposedAddressItem := range proposedAddress {
		proposedAddressRule = append(proposedAddressRule, proposedAddressItem)
	}
	var relayerActionTypeRule []interface{}
	for _, relayerActionTypeItem := range relayerActionType {
		relayerActionTypeRule = append(relayerActionTypeRule, relayerActionTypeItem)
	}

	logs, sub, err := _Relayer.contract.WatchLogs(opts, "RelayerProposalCreated", proposedAddressRule, relayerActionTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerRelayerProposalCreated)
				if err := _Relayer.contract.UnpackLog(event, "RelayerProposalCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerProposalCreated is a log parse operation binding the contract event 0xb12a1fbc5d87bccd40c7ac1f669d842fa9fdfe7f65d2988ceac40543b646807e.
//
// Solidity: event RelayerProposalCreated(address indexed proposedAddress, uint8 indexed relayerActionType)
func (_Relayer *RelayerFilterer) ParseRelayerProposalCreated(log types.Log) (*RelayerRelayerProposalCreated, error) {
	event := new(RelayerRelayerProposalCreated)
	if err := _Relayer.contract.UnpackLog(event, "RelayerProposalCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerRelayerProposalVoteIterator is returned from FilterRelayerProposalVote and is used to iterate over the raw logs and unpacked data for RelayerProposalVote events raised by the Relayer contract.
type RelayerRelayerProposalVoteIterator struct {
	Event *RelayerRelayerProposalVote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RelayerRelayerProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerRelayerProposalVote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RelayerRelayerProposalVote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RelayerRelayerProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerRelayerProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerRelayerProposalVote represents a RelayerProposalVote event raised by the Relayer contract.
type RelayerRelayerProposalVote struct {
	ProposedAddress common.Address
	Vote            uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRelayerProposalVote is a free log retrieval operation binding the contract event 0xb6a4771df73f6cb79535dd475a69a813362d4b13633fdd88d406a3ed78a60332.
//
// Solidity: event RelayerProposalVote(address indexed proposedAddress, uint8 indexed vote)
func (_Relayer *RelayerFilterer) FilterRelayerProposalVote(opts *bind.FilterOpts, proposedAddress []common.Address, vote []uint8) (*RelayerRelayerProposalVoteIterator, error) {

	var proposedAddressRule []interface{}
	for _, proposedAddressItem := range proposedAddress {
		proposedAddressRule = append(proposedAddressRule, proposedAddressItem)
	}
	var voteRule []interface{}
	for _, voteItem := range vote {
		voteRule = append(voteRule, voteItem)
	}

	logs, sub, err := _Relayer.contract.FilterLogs(opts, "RelayerProposalVote", proposedAddressRule, voteRule)
	if err != nil {
		return nil, err
	}
	return &RelayerRelayerProposalVoteIterator{contract: _Relayer.contract, event: "RelayerProposalVote", logs: logs, sub: sub}, nil
}

// WatchRelayerProposalVote is a free log subscription operation binding the contract event 0xb6a4771df73f6cb79535dd475a69a813362d4b13633fdd88d406a3ed78a60332.
//
// Solidity: event RelayerProposalVote(address indexed proposedAddress, uint8 indexed vote)
func (_Relayer *RelayerFilterer) WatchRelayerProposalVote(opts *bind.WatchOpts, sink chan<- *RelayerRelayerProposalVote, proposedAddress []common.Address, vote []uint8) (event.Subscription, error) {

	var proposedAddressRule []interface{}
	for _, proposedAddressItem := range proposedAddress {
		proposedAddressRule = append(proposedAddressRule, proposedAddressItem)
	}
	var voteRule []interface{}
	for _, voteItem := range vote {
		voteRule = append(voteRule, voteItem)
	}

	logs, sub, err := _Relayer.contract.WatchLogs(opts, "RelayerProposalVote", proposedAddressRule, voteRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerRelayerProposalVote)
				if err := _Relayer.contract.UnpackLog(event, "RelayerProposalVote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerProposalVote is a log parse operation binding the contract event 0xb6a4771df73f6cb79535dd475a69a813362d4b13633fdd88d406a3ed78a60332.
//
// Solidity: event RelayerProposalVote(address indexed proposedAddress, uint8 indexed vote)
func (_Relayer *RelayerFilterer) ParseRelayerProposalVote(log types.Log) (*RelayerRelayerProposalVote, error) {
	event := new(RelayerRelayerProposalVote)
	if err := _Relayer.contract.UnpackLog(event, "RelayerProposalVote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerRelayerRemovedIterator is returned from FilterRelayerRemoved and is used to iterate over the raw logs and unpacked data for RelayerRemoved events raised by the Relayer contract.
type RelayerRelayerRemovedIterator struct {
	Event *RelayerRelayerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RelayerRelayerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerRelayerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RelayerRelayerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RelayerRelayerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerRelayerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerRelayerRemoved represents a RelayerRemoved event raised by the Relayer contract.
type RelayerRelayerRemoved struct {
	RelayerAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRelayerRemoved is a free log retrieval operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayerAddress)
func (_Relayer *RelayerFilterer) FilterRelayerRemoved(opts *bind.FilterOpts, relayerAddress []common.Address) (*RelayerRelayerRemovedIterator, error) {

	var relayerAddressRule []interface{}
	for _, relayerAddressItem := range relayerAddress {
		relayerAddressRule = append(relayerAddressRule, relayerAddressItem)
	}

	logs, sub, err := _Relayer.contract.FilterLogs(opts, "RelayerRemoved", relayerAddressRule)
	if err != nil {
		return nil, err
	}
	return &RelayerRelayerRemovedIterator{contract: _Relayer.contract, event: "RelayerRemoved", logs: logs, sub: sub}, nil
}

// WatchRelayerRemoved is a free log subscription operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayerAddress)
func (_Relayer *RelayerFilterer) WatchRelayerRemoved(opts *bind.WatchOpts, sink chan<- *RelayerRelayerRemoved, relayerAddress []common.Address) (event.Subscription, error) {

	var relayerAddressRule []interface{}
	for _, relayerAddressItem := range relayerAddress {
		relayerAddressRule = append(relayerAddressRule, relayerAddressItem)
	}

	logs, sub, err := _Relayer.contract.WatchLogs(opts, "RelayerRemoved", relayerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerRelayerRemoved)
				if err := _Relayer.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerRemoved is a log parse operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayerAddress)
func (_Relayer *RelayerFilterer) ParseRelayerRemoved(log types.Log) (*RelayerRelayerRemoved, error) {
	event := new(RelayerRelayerRemoved)
	if err := _Relayer.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerRelayerThresholdChangedIterator is returned from FilterRelayerThresholdChanged and is used to iterate over the raw logs and unpacked data for RelayerThresholdChanged events raised by the Relayer contract.
type RelayerRelayerThresholdChangedIterator struct {
	Event *RelayerRelayerThresholdChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RelayerRelayerThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerRelayerThresholdChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RelayerRelayerThresholdChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RelayerRelayerThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerRelayerThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerRelayerThresholdChanged represents a RelayerThresholdChanged event raised by the Relayer contract.
type RelayerRelayerThresholdChanged struct {
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRelayerThresholdChanged is a free log retrieval operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 indexed newThreshold)
func (_Relayer *RelayerFilterer) FilterRelayerThresholdChanged(opts *bind.FilterOpts, newThreshold []*big.Int) (*RelayerRelayerThresholdChangedIterator, error) {

	var newThresholdRule []interface{}
	for _, newThresholdItem := range newThreshold {
		newThresholdRule = append(newThresholdRule, newThresholdItem)
	}

	logs, sub, err := _Relayer.contract.FilterLogs(opts, "RelayerThresholdChanged", newThresholdRule)
	if err != nil {
		return nil, err
	}
	return &RelayerRelayerThresholdChangedIterator{contract: _Relayer.contract, event: "RelayerThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchRelayerThresholdChanged is a free log subscription operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 indexed newThreshold)
func (_Relayer *RelayerFilterer) WatchRelayerThresholdChanged(opts *bind.WatchOpts, sink chan<- *RelayerRelayerThresholdChanged, newThreshold []*big.Int) (event.Subscription, error) {

	var newThresholdRule []interface{}
	for _, newThresholdItem := range newThreshold {
		newThresholdRule = append(newThresholdRule, newThresholdItem)
	}

	logs, sub, err := _Relayer.contract.WatchLogs(opts, "RelayerThresholdChanged", newThresholdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerRelayerThresholdChanged)
				if err := _Relayer.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerThresholdChanged is a log parse operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 indexed newThreshold)
func (_Relayer *RelayerFilterer) ParseRelayerThresholdChanged(log types.Log) (*RelayerRelayerThresholdChanged, error) {
	event := new(RelayerRelayerThresholdChanged)
	if err := _Relayer.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerRelayerThresholdProposalCreatedIterator is returned from FilterRelayerThresholdProposalCreated and is used to iterate over the raw logs and unpacked data for RelayerThresholdProposalCreated events raised by the Relayer contract.
type RelayerRelayerThresholdProposalCreatedIterator struct {
	Event *RelayerRelayerThresholdProposalCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RelayerRelayerThresholdProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerRelayerThresholdProposalCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RelayerRelayerThresholdProposalCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RelayerRelayerThresholdProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerRelayerThresholdProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerRelayerThresholdProposalCreated represents a RelayerThresholdProposalCreated event raised by the Relayer contract.
type RelayerRelayerThresholdProposalCreated struct {
	ProposedValue *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRelayerThresholdProposalCreated is a free log retrieval operation binding the contract event 0x79a7cdf066fccb1627ec87a3dd0bf8dcb3a186313f941c8bf80ed979aa62d38d.
//
// Solidity: event RelayerThresholdProposalCreated(uint256 indexed proposedValue)
func (_Relayer *RelayerFilterer) FilterRelayerThresholdProposalCreated(opts *bind.FilterOpts, proposedValue []*big.Int) (*RelayerRelayerThresholdProposalCreatedIterator, error) {

	var proposedValueRule []interface{}
	for _, proposedValueItem := range proposedValue {
		proposedValueRule = append(proposedValueRule, proposedValueItem)
	}

	logs, sub, err := _Relayer.contract.FilterLogs(opts, "RelayerThresholdProposalCreated", proposedValueRule)
	if err != nil {
		return nil, err
	}
	return &RelayerRelayerThresholdProposalCreatedIterator{contract: _Relayer.contract, event: "RelayerThresholdProposalCreated", logs: logs, sub: sub}, nil
}

// WatchRelayerThresholdProposalCreated is a free log subscription operation binding the contract event 0x79a7cdf066fccb1627ec87a3dd0bf8dcb3a186313f941c8bf80ed979aa62d38d.
//
// Solidity: event RelayerThresholdProposalCreated(uint256 indexed proposedValue)
func (_Relayer *RelayerFilterer) WatchRelayerThresholdProposalCreated(opts *bind.WatchOpts, sink chan<- *RelayerRelayerThresholdProposalCreated, proposedValue []*big.Int) (event.Subscription, error) {

	var proposedValueRule []interface{}
	for _, proposedValueItem := range proposedValue {
		proposedValueRule = append(proposedValueRule, proposedValueItem)
	}

	logs, sub, err := _Relayer.contract.WatchLogs(opts, "RelayerThresholdProposalCreated", proposedValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerRelayerThresholdProposalCreated)
				if err := _Relayer.contract.UnpackLog(event, "RelayerThresholdProposalCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerThresholdProposalCreated is a log parse operation binding the contract event 0x79a7cdf066fccb1627ec87a3dd0bf8dcb3a186313f941c8bf80ed979aa62d38d.
//
// Solidity: event RelayerThresholdProposalCreated(uint256 indexed proposedValue)
func (_Relayer *RelayerFilterer) ParseRelayerThresholdProposalCreated(log types.Log) (*RelayerRelayerThresholdProposalCreated, error) {
	event := new(RelayerRelayerThresholdProposalCreated)
	if err := _Relayer.contract.UnpackLog(event, "RelayerThresholdProposalCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerRelayerThresholdProposalVoteIterator is returned from FilterRelayerThresholdProposalVote and is used to iterate over the raw logs and unpacked data for RelayerThresholdProposalVote events raised by the Relayer contract.
type RelayerRelayerThresholdProposalVoteIterator struct {
	Event *RelayerRelayerThresholdProposalVote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RelayerRelayerThresholdProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerRelayerThresholdProposalVote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RelayerRelayerThresholdProposalVote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RelayerRelayerThresholdProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerRelayerThresholdProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerRelayerThresholdProposalVote represents a RelayerThresholdProposalVote event raised by the Relayer contract.
type RelayerRelayerThresholdProposalVote struct {
	Vote uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRelayerThresholdProposalVote is a free log retrieval operation binding the contract event 0xe6124fe2b7a19e7cdd807fb16247ec258d8fb0bfde2949b023bcb0c1eea19cb1.
//
// Solidity: event RelayerThresholdProposalVote(uint8 indexed vote)
func (_Relayer *RelayerFilterer) FilterRelayerThresholdProposalVote(opts *bind.FilterOpts, vote []uint8) (*RelayerRelayerThresholdProposalVoteIterator, error) {

	var voteRule []interface{}
	for _, voteItem := range vote {
		voteRule = append(voteRule, voteItem)
	}

	logs, sub, err := _Relayer.contract.FilterLogs(opts, "RelayerThresholdProposalVote", voteRule)
	if err != nil {
		return nil, err
	}
	return &RelayerRelayerThresholdProposalVoteIterator{contract: _Relayer.contract, event: "RelayerThresholdProposalVote", logs: logs, sub: sub}, nil
}

// WatchRelayerThresholdProposalVote is a free log subscription operation binding the contract event 0xe6124fe2b7a19e7cdd807fb16247ec258d8fb0bfde2949b023bcb0c1eea19cb1.
//
// Solidity: event RelayerThresholdProposalVote(uint8 indexed vote)
func (_Relayer *RelayerFilterer) WatchRelayerThresholdProposalVote(opts *bind.WatchOpts, sink chan<- *RelayerRelayerThresholdProposalVote, vote []uint8) (event.Subscription, error) {

	var voteRule []interface{}
	for _, voteItem := range vote {
		voteRule = append(voteRule, voteItem)
	}

	logs, sub, err := _Relayer.contract.WatchLogs(opts, "RelayerThresholdProposalVote", voteRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerRelayerThresholdProposalVote)
				if err := _Relayer.contract.UnpackLog(event, "RelayerThresholdProposalVote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerThresholdProposalVote is a log parse operation binding the contract event 0xe6124fe2b7a19e7cdd807fb16247ec258d8fb0bfde2949b023bcb0c1eea19cb1.
//
// Solidity: event RelayerThresholdProposalVote(uint8 indexed vote)
func (_Relayer *RelayerFilterer) ParseRelayerThresholdProposalVote(log types.Log) (*RelayerRelayerThresholdProposalVote, error) {
	event := new(RelayerRelayerThresholdProposalVote)
	if err := _Relayer.contract.UnpackLog(event, "RelayerThresholdProposalVote", log); err != nil {
		return nil, err
	}
	return event, nil
}

var RuntimeBytecode = "0x608060405234801561001057600080fd5b50600436106100a85760003560e01c80635f31b69c116100715780635f31b69c1461033c578063802aabe8146103ff578063933b46671461041d578063d7a9cd791461043b578063df26906014610459578063e9cdaead14610487576100a8565b806251f244146100ad57806330fddf01146101b85780634d4ebd7a14610233578063541d5548146102845780635d0db4db146102e0575b600080fd5b6100ef600480360360208110156100c357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104b8565b604051808560018111156100ff57fe5b60ff168152602001806020018060200184600181111561011b57fe5b60ff168152602001838103835286818151815260200191508051906020019060200280838360005b8381101561015e578082015181840152602081019050610143565b50505050905001838103825285818151815260200191508051906020019060200280838360005b838110156101a0578082015181840152602081019050610185565b50505050905001965050505050505060405180910390f35b6101fa600480360360208110156101ce57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106b3565b6040518083600181111561020a57fe5b60ff16815260200182600181111561021e57fe5b60ff1681526020019250505060405180910390f35b6102826004803603604081101561024957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff1690602001909291905050506106f1565b005b6102c66004803603602081101561029a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610eec565b604051808215151515815260200191505060405180910390f35b610322600480360360208110156102f657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f42565b604051808215151515815260200191505060405180910390f35b610344610f62565b60405180858152602001806020018060200184600181111561036257fe5b60ff168152602001838103835286818151815260200191508051906020019060200280838360005b838110156103a557808201518184015260208101905061038a565b50505050905001838103825285818151815260200191508051906020019060200280838360005b838110156103e75780820151818401526020810190506103cc565b50505050905001965050505050505060405180910390f35b6104076110a9565b6040518082815260200191505060405180910390f35b6104256110af565b6040518082815260200191505060405180910390f35b6104436110b9565b6040518082815260200191505060405180910390f35b6104856004803603602081101561046f57600080fd5b81019080803590602001909291905050506110bf565b005b6104b66004803603602081101561049d57600080fd5b81019080803560ff1690602001909291905050506114c4565b005b600060608060006104c7611c21565b600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060800160405290816000820160009054906101000a900460ff16600181111561052f57fe5b600181111561053a57fe5b8152602001600282018054806020026020016040519081016040528092919081815260200182805480156105c357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610579575b505050505081526020016003820180548060200260200160405190810160405280929190818152602001828054801561065157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610607575b505050505081526020016004820160009054906101000a900460ff16600181111561067857fe5b600181111561068357fe5b81525050905080600001518160200151826040015183606001518292508191509450945094509450509193509193565b60086020528060005260406000206000915090508060000160009054906101000a900460ff16908060040160009054906101000a900460ff16905082565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166107b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f73656e646572206973206e6f7420612072656c6179657200000000000000000081525060200191505060405180910390fd5b6000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600182600181111561080157fe5b60ff16111561085b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611db96023913960400191505060405180910390fd5b600015158160010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514610906576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180611d2d6025913960400191505060405180910390fd5b6000600181111561091357fe5b82600181111561091f57fe5b14156109cc57600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166109c7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180611ddc6021913960400191505060405180910390fd5b610a70565b600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610a6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180611d526025913960400191505060405180910390fd5b5b60006001811115610a7d57fe5b8160040160009054906101000a900460ff166001811115610a9a57fe5b1415610cb2576040518060800160405280836001811115610ab757fe5b81526020016001604051908082528060200260200182016040528015610aec5781602001602082028036833780820191505090505b5081526020016000604051908082528060200260200182016040528015610b225781602001602082028036833780820191505090505b508152602001600180811115610b3457fe5b815250600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690836001811115610b9a57fe5b02179055506020820151816002019080519060200190610bbb929190611c5f565b506040820151816003019080519060200190610bd8929190611c5f565b5060608201518160040160006101000a81548160ff02191690836001811115610bfd57fe5b02179055509050503381600201600081548110610c1657fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816001811115610c6a57fe5b8373ffffffffffffffffffffffffffffffffffffffff167fb12a1fbc5d87bccd40c7ac1f669d842fa9fdfe7f65d2988ceac40543b646807e60405160405180910390a3610d18565b80600201339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b60018160010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508060040160009054906101000a900460ff166001811115610d8f57fe5b8373ffffffffffffffffffffffffffffffffffffffff167fb6a4771df73f6cb79535dd475a69a813362d4b13633fdd88d406a3ed78a6033260405160405180910390a36001600054111580610ded5750600054816002018054905010155b15610ee75760008160040160006101000a81548160ff02191690836001811115610e1357fe5b021790555060006001811115610e2557fe5b8160000160009054906101000a900460ff166001811115610e4257fe5b1415610e9957610e51836119b6565b8273ffffffffffffffffffffffffffffffffffffffff167f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b60405160405180910390a2610ee6565b610ea283611a67565b8273ffffffffffffffffffffffffffffffffffffffff167f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c560405160405180910390a25b5b505050565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60076020528060005260406000206000915054906101000a900460ff1681565b60006060806000600260000154600280016002600301600260040160009054906101000a900460ff168280548060200260200160405190810160405280929190818152602001828054801561100c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610fc2575b505050505092508180548060200260200160405190810160405280929190818152602001828054801561109457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161104a575b50505050509150935093509350935090919293565b60015481565b6000600154905090565b60005481565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661117e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f73656e646572206973206e6f7420612072656c6179657200000000000000000081525060200191505060405180910390fd5b6000600181111561118b57fe5b600260040160009054906101000a900460ff1660018111156111a957fe5b1461121c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f612070726f706f73616c2069732063757272656e746c7920616374697665000081525060200191505060405180910390fd5b600154811115611277576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526042815260200180611d776042913960600191505060405180910390fd5b604051806080016040528082815260200160016040519080825280602002602001820160405280156112b85781602001602082028036833780820191505090505b50815260200160006040519080825280602002602001820160405280156112ee5781602001602082028036833780820191505090505b50815260200160018081111561130057fe5b815250600260008201518160000155602082015181600201908051906020019061132b929190611c5f565b506040820151816003019080519060200190611348929190611c5f565b5060608201518160040160006101000a81548160ff0219169083600181111561136d57fe5b02179055509050506001600054116113e0576002600001546000819055506000600260040160006101000a81548160ff021916908360018111156113ad57fe5b0217905550807fa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c860405160405180910390a25b33600280016000815481106113f157fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550807f79a7cdf066fccb1627ec87a3dd0bf8dcb3a186313f941c8bf80ed979aa62d38d60405160405180910390a250565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611583576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f73656e646572206973206e6f7420612072656c6179657200000000000000000081525060200191505060405180910390fd5b600181600181111561159157fe5b60ff161115611608576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f766f7465206f7574206f662074686520766f746520656e756d2072616e67650081525060200191505060405180910390fd5b60018081111561161457fe5b600260040160009054906101000a900460ff16600181111561163257fe5b146116a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f6e6f2070726f706f73616c2069732063757272656e746c79206163746976650081525060200191505060405180910390fd5b600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611768576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f72656c617965722068617320616c726561647920766f7465640000000000000081525060200191505060405180910390fd5b60018081111561177457fe5b81600181111561178057fe5b14156117f05760028001339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611857565b6002600301339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b6001600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508060018111156118be57fe5b7fe6124fe2b7a19e7cdd807fb16247ec258d8fb0bfde2949b023bcb0c1eea19cb160405160405180910390a2600054600280018054905010611964576002600001546000819055506000600260040160006101000a81548160ff0219169083600181111561192857fe5b02179055506002600001547fa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c860405160405180910390a26119b3565b600054611984600260030180549050600154611b1790919063ffffffff16565b10156119b2576000600260040160006101000a81548160ff021916908360018111156119ac57fe5b02179055505b5b50565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160008154809291906001900391905055508073ffffffffffffffffffffffffffffffffffffffff167f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b60405160405180910390a250565b6001600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600081548092919060010191905055508073ffffffffffffffffffffffffffffffffffffffff167f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c560405160405180910390a250565b6000611b5983836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611b61565b905092915050565b6000838311158290611c0e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611bd3578082015181840152602081019050611bb8565b50505050905090810190601f168015611c005780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b604051806080016040528060006001811115611c3957fe5b8152602001606081526020016060815260200160006001811115611c5957fe5b81525090565b828054828255906000526020600020908101928215611cd8579160200282015b82811115611cd75782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190611c7f565b5b509050611ce59190611ce9565b5090565b611d2991905b80821115611d2557600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101611cef565b5090565b9056fe72656c617965722068617320616c726561647920766f746564206f6e2070726f706f73616c70726f706f736564206164647265737320697320616c726561647920612072656c6179657270726f706f7365642076616c75652063616e6e6f742062652067726561746572207468616e2074686520746f74616c206e756d626572206f662072656c6179657273616374696f6e206f7574206f662074686520616374696f6e20656e756d2072616e676570726f706f7365642061646472657373206973206e6f7420612072656c61796572a2646970667358221220d871218f009df167ebaab13b975cb9f5f2b8355144c315b764ac2c2581a0e05f64736f6c63430006040033"
