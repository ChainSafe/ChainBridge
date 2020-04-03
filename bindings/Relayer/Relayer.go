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
const RelayerABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"initialRelayers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"initialRelayerThreshold\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"RelayerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIRelayer.RelayerActionType\",\"name\":\"relayerActionType\",\"type\":\"uint8\"}],\"name\":\"RelayerProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIRelayer.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"RelayerProposalVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"RelayerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"RelayerThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposedValue\",\"type\":\"uint256\"}],\"name\":\"RelayerThresholdProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumIRelayer.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"RelayerThresholdProposalVote\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_relayerProposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_proposedAddress\",\"type\":\"address\"},{\"internalType\":\"enumIRelayer.RelayerActionType\",\"name\":\"_action\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_numYes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numNo\",\"type\":\"uint256\"},{\"internalType\":\"enumIRelayer.VoteStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_relayerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_relayers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalRelayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelayerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalRelayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRelayerThresholdProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"}],\"name\":\"getRelayerProposal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"},{\"internalType\":\"enumIRelayer.RelayerActionType\",\"name\":\"action\",\"type\":\"uint8\"}],\"name\":\"createRelayerProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"},{\"internalType\":\"enumIRelayer.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"voteRelayerProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposedValue\",\"type\":\"uint256\"}],\"name\":\"createRelayerThresholdProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIRelayer.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"voteRelayerThresholdProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RelayerBin is the compiled bytecode used for deploying new contracts.
var RelayerBin = "0x608060405260405180604001604052806040518060400160405280600681526020017f72656d6f7665000000000000000000000000000000000000000000000000000081525081526020016040518060400160405280600381526020017f616464000000000000000000000000000000000000000000000000000000000081525081525060079060026200009592919062000304565b5060405180604001604052806040518060400160405280600881526020017f696e61637469766500000000000000000000000000000000000000000000000081525081526020016040518060400160405280600681526020017f616374697665000000000000000000000000000000000000000000000000000081525081525060089060026200012792919062000304565b503480156200013557600080fd5b50604051620027a9380380620027a9833981810160405260408110156200015b57600080fd5b81019080805160405193929190846401000000008211156200017c57600080fd5b838201915060208201858111156200019357600080fd5b8251866020820283011164010000000082111715620001b157600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620001ea578082015181840152602081019050620001cd565b505050509050016040526020018051906020019092919050505060005b82518110156200024457620002368382815181106200022257fe5b60200260200101516200025460201b60201c565b808060010191505062000207565b5080600081905550505062000497565b6001600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600081548092919060010191905055508073ffffffffffffffffffffffffffffffffffffffff167f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c560405160405180910390a250565b82805482825590600052602060002090810192821562000358579160200282015b8281111562000357578251829080519060200190620003469291906200036b565b509160200191906001019062000325565b5b509050620003679190620003f2565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003ae57805160ff1916838001178555620003df565b82800160010185558215620003df579182015b82811115620003de578251825591602001919060010190620003c1565b5b509050620003ee919062000423565b5090565b6200042091905b808211156200041c57600081816200041291906200044b565b50600101620003f9565b5090565b90565b6200044891905b80821115620004445760008160009055506001016200042a565b5090565b90565b50805460018160011615610100020316600290046000825580601f1062000473575062000494565b601f01602090049060005260206000209081019062000493919062000423565b5b50565b61230280620004a76000396000f3fe608060405234801561001057600080fd5b50600436106100ce5760003560e01c8063677ce4e11161008c578063a795911911610066578063a7959119146104f4578063d7a9cd7914610545578063df26906014610563578063e9cdaead14610591576100ce565b8063677ce4e11461049a578063802aabe8146104b8578063933b4667146104d6576100ce565b806251f244146100d357806330fddf011461023d5780634d4ebd7a146102f9578063541d55481461034a5780635d0db4db146103a65780635f31b69c14610402575b600080fd5b610115600480360360208110156100e957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105c2565b604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018060200185815260200184815260200180602001838103835287818151815260200191508051906020019080838360005b8381101561019757808201518184015260208101905061017c565b50505050905090810190601f1680156101c45780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156101fd5780820151818401526020810190506101e2565b50505050905090810190601f16801561022a5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b61027f6004803603602081101561025357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061088a565b604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018560018111156102c157fe5b60ff1681526020018481526020018381526020018260018111156102e157fe5b60ff1681526020019550505050505060405180910390f35b6103486004803603604081101561030f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff1690602001909291905050506108fa565b005b61038c6004803603602081101561036057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fb7565b604051808215151515815260200191505060405180910390f35b6103e8600480360360208110156103bc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061100d565b604051808215151515815260200191505060405180910390f35b61040a61102d565b6040518085815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561045c578082015181840152602081019050610441565b50505050905090810190601f1680156104895780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b6104a2611123565b6040518082815260200191505060405180910390f35b6104c061112c565b6040518082815260200191505060405180910390f35b6104de611132565b6040518082815260200191505060405180910390f35b6105436004803603604081101561050a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff16906020019092919050505061113c565b005b61054d6117c1565b6040518082815260200191505060405180910390f35b61058f6004803603602081101561057957600080fd5b81019080803590602001909291905050506117c7565b005b6105c0600480360360208110156105a757600080fd5b81019080803560ff169060200190929190505050611aef565b005b6000606060008060606105d36121ae565b600a60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff16600181111561069157fe5b600181111561069c57fe5b815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff1660018111156106d257fe5b60018111156106dd57fe5b81525050905080600001516007826020015160018111156106fa57fe5b8154811061070457fe5b906000526020600020018260400151836060015160088560800151600181111561072a57fe5b8154811061073457fe5b90600052602060002001838054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107d35780601f106107a8576101008083540402835291602001916107d3565b820191906000526020600020905b8154815290600101906020018083116107b657829003601f168201915b50505050509350808054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561086f5780601f106108445761010080835404028352916020019161086f565b820191906000526020600020905b81548152906001019060200180831161085257829003601f168201915b50505050509050955095509550955095505091939590929450565b600a6020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16908060020154908060030154908060040160009054906101000a900460ff16905085565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166109b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f73656e646572206973206e6f7420612072656c6179657200000000000000000081525060200191505060405180910390fd5b6001808111156109c557fe5b600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160009054906101000a900460ff166001811115610a2057fe5b14610a76576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c81526020018061226d602c913960400191505060405180910390fd5b600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610b76576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f72656c617965722068617320616c726561647920766f7465640000000000000081525060200191505060405180910390fd5b6001816001811115610b8457fe5b1115610bf8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f766f7465206f7574206f662074686520766f746520656e756d2072616e67650081525060200191505060405180910390fd5b600180811115610c0457fe5b816001811115610c1057fe5b1415610c6d57600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160008154809291906001019190505550610cc0565b600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301600081548092919060010191905055505b6001600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167fb6a4771df73f6cb79535dd475a69a813362d4b13633fdd88d406a3ed78a603328260405180826001811115610da157fe5b60ff16815260200191505060405180910390a2600054600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015410610eed57600180811115610e0b57fe5b600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff166001811115610e6657fe5b1415610e7a57610e7582611f43565b610e84565b610e8382611ff3565b5b6000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160006101000a81548160ff02191690836001811115610ee357fe5b0217905550610fb3565b600054610f47600a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301546001546120a490919063ffffffff16565b1015610fb2576000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160006101000a81548160ff02191690836001811115610fac57fe5b02179055505b5b5050565b6000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60096020528060005260406000206000915054906101000a900460ff1681565b6000806000606060026000015460028001546002600301546008600260040160009054906101000a900460ff16600181111561106557fe5b8154811061106f57fe5b90600052602060002001808054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561110e5780601f106110e35761010080835404028352916020019161110e565b820191906000526020600020905b8154815290600101906020018083116110f157829003601f168201915b50505050509050935093509350935090919293565b60008054905090565b60015481565b6000600154905090565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166111fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f73656e646572206973206e6f7420612072656c6179657200000000000000000081525060200191505060405180910390fd5b600181600181111561120957fe5b1115611260576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018061224c6021913960400191505060405180910390fd5b6000600181111561126d57fe5b81600181111561127957fe5b1480156112d6575060011515600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515145b611348576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f61646472657373206973206e6f7420612072656c61796572000000000000000081525060200191505060405180910390fd5b60018081111561135457fe5b81600181111561136057fe5b1480156113bd575060001515600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515145b61142f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f616464726573732069732063757272656e746c7920612072656c61796572000081525060200191505060405180910390fd5b6000600181111561143c57fe5b600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160009054906101000a900460ff16600181111561149757fe5b146114ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806122996034913960400191505060405180910390fd5b6040518060a001604052808373ffffffffffffffffffffffffffffffffffffffff16815260200182600181111561152057fe5b8152602001600181526020016000815260200160018081111561153f57fe5b815250600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360018111156115ec57fe5b0217905550604082015181600201556060820151816003015560808201518160040160006101000a81548160ff0219169083600181111561162957fe5b02179055509050506001600054116116d6576000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160006101000a81548160ff0219169083600181111561169a57fe5b02179055506001808111156116ab57fe5b8160018111156116b757fe5b14156116cb576116c682611f43565b6116d5565b6116d482611ff3565b5b5b6001600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555080600181111561177a57fe5b8273ffffffffffffffffffffffffffffffffffffffff167fb12a1fbc5d87bccd40c7ac1f669d842fa9fdfe7f65d2988ceac40543b646807e60405160405180910390a35050565b60005481565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611886576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f73656e646572206973206e6f7420612072656c6179657200000000000000000081525060200191505060405180910390fd5b6000600181111561189357fe5b600260040160009054906101000a900460ff1660018111156118b157fe5b14611924576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f612070726f706f73616c2069732063757272656e746c7920616374697665000081525060200191505060405180910390fd5b60015481111561197f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604281526020018061220a6042913960600191505060405180910390fd5b604051806080016040528082815260200160018152602001600081526020016001808111156119aa57fe5b815250600260008201518160000155602082015181600201556040820151816003015560608201518160040160006101000a81548160ff021916908360018111156119f157fe5b0217905550905050600160005411611a64576002600001546000819055506000600260040160006101000a81548160ff02191690836001811115611a3157fe5b0217905550807fa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c860405160405180910390a25b6001600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550807f79a7cdf066fccb1627ec87a3dd0bf8dcb3a186313f941c8bf80ed979aa62d38d60405160405180910390a250565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611bae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f73656e646572206973206e6f7420612072656c6179657200000000000000000081525060200191505060405180910390fd5b600180811115611bba57fe5b600260040160009054906101000a900460ff166001811115611bd857fe5b14611c4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f6e6f2070726f706f73616c2069732063757272656e746c79206163746976650081525060200191505060405180910390fd5b600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611d0e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f72656c617965722068617320616c726561647920766f7465640000000000000081525060200191505060405180910390fd5b6001816001811115611d1c57fe5b1115611d90576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f766f7465206f7574206f662074686520766f746520656e756d2072616e67650081525060200191505060405180910390fd5b600180811115611d9c57fe5b816001811115611da857fe5b1415611dc7576002800160008154809291906001019190505550611ddd565b6002600301600081548092919060010191905055505b6001600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fe6124fe2b7a19e7cdd807fb16247ec258d8fb0bfde2949b023bcb0c1eea19cb18160405180826001811115611e6a57fe5b60ff16815260200191505060405180910390a1600054600280015410611ef4576002600001546000819055506000600260040160006101000a81548160ff02191690836001811115611eb857fe5b02179055506002600001547fa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c860405160405180910390a2611f40565b600054611f116002600301546001546120a490919063ffffffff16565b1015611f3f576000600260040160006101000a81548160ff02191690836001811115611f3957fe5b02179055505b5b50565b6001600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600081548092919060010191905055508073ffffffffffffffffffffffffffffffffffffffff167f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c560405160405180910390a250565b6000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160008154809291906001900391905055508073ffffffffffffffffffffffffffffffffffffffff167f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b60405160405180910390a250565b60006120e683836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506120ee565b905092915050565b600083831115829061219b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612160578082015181840152602081019050612145565b50505050905090810190601f16801561218d5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060018111156121e357fe5b815260200160008152602001600081526020016000600181111561220357fe5b8152509056fe70726f706f7365642076616c75652063616e6e6f742062652067726561746572207468616e2074686520746f74616c206e756d626572206f662072656c6179657273616374696f6e206f7574206f662074686520766f746520656e756d2072616e67657468657265206973206e6f206163746976652070726f706f73616c20666f7220746869732061646472657373746865726520697320616c726561647920616e206163746976652070726f706f73616c20666f7220746869732061646472657373a2646970667358221220a3b64ba2f9b5befb7c7cdcf6545e2983b6117664460fc85812ea636512dc407f64736f6c63430006040033"

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
// Solidity: function _relayerProposals(address ) constant returns(address _proposedAddress, uint8 _action, uint256 _numYes, uint256 _numNo, uint8 _status)
func (_Relayer *RelayerCaller) RelayerProposals(opts *bind.CallOpts, arg0 common.Address) (struct {
	ProposedAddress common.Address
	Action          uint8
	NumYes          *big.Int
	NumNo           *big.Int
	Status          uint8
}, error) {
	ret := new(struct {
		ProposedAddress common.Address
		Action          uint8
		NumYes          *big.Int
		NumNo           *big.Int
		Status          uint8
	})
	out := ret
	err := _Relayer.contract.Call(opts, out, "_relayerProposals", arg0)
	return *ret, err
}

// RelayerProposals is a free data retrieval call binding the contract method 0x30fddf01.
//
// Solidity: function _relayerProposals(address ) constant returns(address _proposedAddress, uint8 _action, uint256 _numYes, uint256 _numNo, uint8 _status)
func (_Relayer *RelayerSession) RelayerProposals(arg0 common.Address) (struct {
	ProposedAddress common.Address
	Action          uint8
	NumYes          *big.Int
	NumNo           *big.Int
	Status          uint8
}, error) {
	return _Relayer.Contract.RelayerProposals(&_Relayer.CallOpts, arg0)
}

// RelayerProposals is a free data retrieval call binding the contract method 0x30fddf01.
//
// Solidity: function _relayerProposals(address ) constant returns(address _proposedAddress, uint8 _action, uint256 _numYes, uint256 _numNo, uint8 _status)
func (_Relayer *RelayerCallerSession) RelayerProposals(arg0 common.Address) (struct {
	ProposedAddress common.Address
	Action          uint8
	NumYes          *big.Int
	NumNo           *big.Int
	Status          uint8
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
// Solidity: function getCurrentRelayerThresholdProposal() constant returns(uint256, uint256, uint256, string)
func (_Relayer *RelayerCaller) GetCurrentRelayerThresholdProposal(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(string)
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
// Solidity: function getCurrentRelayerThresholdProposal() constant returns(uint256, uint256, uint256, string)
func (_Relayer *RelayerSession) GetCurrentRelayerThresholdProposal() (*big.Int, *big.Int, *big.Int, string, error) {
	return _Relayer.Contract.GetCurrentRelayerThresholdProposal(&_Relayer.CallOpts)
}

// GetCurrentRelayerThresholdProposal is a free data retrieval call binding the contract method 0x5f31b69c.
//
// Solidity: function getCurrentRelayerThresholdProposal() constant returns(uint256, uint256, uint256, string)
func (_Relayer *RelayerCallerSession) GetCurrentRelayerThresholdProposal() (*big.Int, *big.Int, *big.Int, string, error) {
	return _Relayer.Contract.GetCurrentRelayerThresholdProposal(&_Relayer.CallOpts)
}

// GetRelayerProposal is a free data retrieval call binding the contract method 0x0051f244.
//
// Solidity: function getRelayerProposal(address proposedAddress) constant returns(address, string, uint256, uint256, string)
func (_Relayer *RelayerCaller) GetRelayerProposal(opts *bind.CallOpts, proposedAddress common.Address) (common.Address, string, *big.Int, *big.Int, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Relayer.contract.Call(opts, out, "getRelayerProposal", proposedAddress)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetRelayerProposal is a free data retrieval call binding the contract method 0x0051f244.
//
// Solidity: function getRelayerProposal(address proposedAddress) constant returns(address, string, uint256, uint256, string)
func (_Relayer *RelayerSession) GetRelayerProposal(proposedAddress common.Address) (common.Address, string, *big.Int, *big.Int, string, error) {
	return _Relayer.Contract.GetRelayerProposal(&_Relayer.CallOpts, proposedAddress)
}

// GetRelayerProposal is a free data retrieval call binding the contract method 0x0051f244.
//
// Solidity: function getRelayerProposal(address proposedAddress) constant returns(address, string, uint256, uint256, string)
func (_Relayer *RelayerCallerSession) GetRelayerProposal(proposedAddress common.Address) (common.Address, string, *big.Int, *big.Int, string, error) {
	return _Relayer.Contract.GetRelayerProposal(&_Relayer.CallOpts, proposedAddress)
}

// GetRelayerThreshold is a free data retrieval call binding the contract method 0x677ce4e1.
//
// Solidity: function getRelayerThreshold() constant returns(uint256)
func (_Relayer *RelayerCaller) GetRelayerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayer.contract.Call(opts, out, "getRelayerThreshold")
	return *ret0, err
}

// GetRelayerThreshold is a free data retrieval call binding the contract method 0x677ce4e1.
//
// Solidity: function getRelayerThreshold() constant returns(uint256)
func (_Relayer *RelayerSession) GetRelayerThreshold() (*big.Int, error) {
	return _Relayer.Contract.GetRelayerThreshold(&_Relayer.CallOpts)
}

// GetRelayerThreshold is a free data retrieval call binding the contract method 0x677ce4e1.
//
// Solidity: function getRelayerThreshold() constant returns(uint256)
func (_Relayer *RelayerCallerSession) GetRelayerThreshold() (*big.Int, error) {
	return _Relayer.Contract.GetRelayerThreshold(&_Relayer.CallOpts)
}

// CreateRelayerProposal is a paid mutator transaction binding the contract method 0xa7959119.
//
// Solidity: function createRelayerProposal(address proposedAddress, uint8 action) returns()
func (_Relayer *RelayerTransactor) CreateRelayerProposal(opts *bind.TransactOpts, proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _Relayer.contract.Transact(opts, "createRelayerProposal", proposedAddress, action)
}

// CreateRelayerProposal is a paid mutator transaction binding the contract method 0xa7959119.
//
// Solidity: function createRelayerProposal(address proposedAddress, uint8 action) returns()
func (_Relayer *RelayerSession) CreateRelayerProposal(proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _Relayer.Contract.CreateRelayerProposal(&_Relayer.TransactOpts, proposedAddress, action)
}

// CreateRelayerProposal is a paid mutator transaction binding the contract method 0xa7959119.
//
// Solidity: function createRelayerProposal(address proposedAddress, uint8 action) returns()
func (_Relayer *RelayerTransactorSession) CreateRelayerProposal(proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _Relayer.Contract.CreateRelayerProposal(&_Relayer.TransactOpts, proposedAddress, action)
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
// Solidity: function voteRelayerProposal(address proposedAddress, uint8 vote) returns()
func (_Relayer *RelayerTransactor) VoteRelayerProposal(opts *bind.TransactOpts, proposedAddress common.Address, vote uint8) (*types.Transaction, error) {
	return _Relayer.contract.Transact(opts, "voteRelayerProposal", proposedAddress, vote)
}

// VoteRelayerProposal is a paid mutator transaction binding the contract method 0x4d4ebd7a.
//
// Solidity: function voteRelayerProposal(address proposedAddress, uint8 vote) returns()
func (_Relayer *RelayerSession) VoteRelayerProposal(proposedAddress common.Address, vote uint8) (*types.Transaction, error) {
	return _Relayer.Contract.VoteRelayerProposal(&_Relayer.TransactOpts, proposedAddress, vote)
}

// VoteRelayerProposal is a paid mutator transaction binding the contract method 0x4d4ebd7a.
//
// Solidity: function voteRelayerProposal(address proposedAddress, uint8 vote) returns()
func (_Relayer *RelayerTransactorSession) VoteRelayerProposal(proposedAddress common.Address, vote uint8) (*types.Transaction, error) {
	return _Relayer.Contract.VoteRelayerProposal(&_Relayer.TransactOpts, proposedAddress, vote)
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
// Solidity: event RelayerProposalVote(address indexed proposedAddress, uint8 vote)
func (_Relayer *RelayerFilterer) FilterRelayerProposalVote(opts *bind.FilterOpts, proposedAddress []common.Address) (*RelayerRelayerProposalVoteIterator, error) {

	var proposedAddressRule []interface{}
	for _, proposedAddressItem := range proposedAddress {
		proposedAddressRule = append(proposedAddressRule, proposedAddressItem)
	}

	logs, sub, err := _Relayer.contract.FilterLogs(opts, "RelayerProposalVote", proposedAddressRule)
	if err != nil {
		return nil, err
	}
	return &RelayerRelayerProposalVoteIterator{contract: _Relayer.contract, event: "RelayerProposalVote", logs: logs, sub: sub}, nil
}

// WatchRelayerProposalVote is a free log subscription operation binding the contract event 0xb6a4771df73f6cb79535dd475a69a813362d4b13633fdd88d406a3ed78a60332.
//
// Solidity: event RelayerProposalVote(address indexed proposedAddress, uint8 vote)
func (_Relayer *RelayerFilterer) WatchRelayerProposalVote(opts *bind.WatchOpts, sink chan<- *RelayerRelayerProposalVote, proposedAddress []common.Address) (event.Subscription, error) {

	var proposedAddressRule []interface{}
	for _, proposedAddressItem := range proposedAddress {
		proposedAddressRule = append(proposedAddressRule, proposedAddressItem)
	}

	logs, sub, err := _Relayer.contract.WatchLogs(opts, "RelayerProposalVote", proposedAddressRule)
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
// Solidity: event RelayerProposalVote(address indexed proposedAddress, uint8 vote)
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
// Solidity: event RelayerThresholdProposalVote(uint8 vote)
func (_Relayer *RelayerFilterer) FilterRelayerThresholdProposalVote(opts *bind.FilterOpts) (*RelayerRelayerThresholdProposalVoteIterator, error) {

	logs, sub, err := _Relayer.contract.FilterLogs(opts, "RelayerThresholdProposalVote")
	if err != nil {
		return nil, err
	}
	return &RelayerRelayerThresholdProposalVoteIterator{contract: _Relayer.contract, event: "RelayerThresholdProposalVote", logs: logs, sub: sub}, nil
}

// WatchRelayerThresholdProposalVote is a free log subscription operation binding the contract event 0xe6124fe2b7a19e7cdd807fb16247ec258d8fb0bfde2949b023bcb0c1eea19cb1.
//
// Solidity: event RelayerThresholdProposalVote(uint8 vote)
func (_Relayer *RelayerFilterer) WatchRelayerThresholdProposalVote(opts *bind.WatchOpts, sink chan<- *RelayerRelayerThresholdProposalVote) (event.Subscription, error) {

	logs, sub, err := _Relayer.contract.WatchLogs(opts, "RelayerThresholdProposalVote")
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
// Solidity: event RelayerThresholdProposalVote(uint8 vote)
func (_Relayer *RelayerFilterer) ParseRelayerThresholdProposalVote(log types.Log) (*RelayerRelayerThresholdProposalVote, error) {
	event := new(RelayerRelayerThresholdProposalVote)
	if err := _Relayer.contract.UnpackLog(event, "RelayerThresholdProposalVote", log); err != nil {
		return nil, err
	}
	return event, nil
}

var RuntimeBytecode = "0x608060405234801561001057600080fd5b50600436106100ce5760003560e01c8063677ce4e11161008c578063a795911911610066578063a7959119146104f4578063d7a9cd7914610545578063df26906014610563578063e9cdaead14610591576100ce565b8063677ce4e11461049a578063802aabe8146104b8578063933b4667146104d6576100ce565b806251f244146100d357806330fddf011461023d5780634d4ebd7a146102f9578063541d55481461034a5780635d0db4db146103a65780635f31b69c14610402575b600080fd5b610115600480360360208110156100e957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105c2565b604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018060200185815260200184815260200180602001838103835287818151815260200191508051906020019080838360005b8381101561019757808201518184015260208101905061017c565b50505050905090810190601f1680156101c45780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156101fd5780820151818401526020810190506101e2565b50505050905090810190601f16801561022a5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b61027f6004803603602081101561025357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061088a565b604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018560018111156102c157fe5b60ff1681526020018481526020018381526020018260018111156102e157fe5b60ff1681526020019550505050505060405180910390f35b6103486004803603604081101561030f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff1690602001909291905050506108fa565b005b61038c6004803603602081101561036057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fb7565b604051808215151515815260200191505060405180910390f35b6103e8600480360360208110156103bc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061100d565b604051808215151515815260200191505060405180910390f35b61040a61102d565b6040518085815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561045c578082015181840152602081019050610441565b50505050905090810190601f1680156104895780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b6104a2611123565b6040518082815260200191505060405180910390f35b6104c061112c565b6040518082815260200191505060405180910390f35b6104de611132565b6040518082815260200191505060405180910390f35b6105436004803603604081101561050a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff16906020019092919050505061113c565b005b61054d6117c1565b6040518082815260200191505060405180910390f35b61058f6004803603602081101561057957600080fd5b81019080803590602001909291905050506117c7565b005b6105c0600480360360208110156105a757600080fd5b81019080803560ff169060200190929190505050611aef565b005b6000606060008060606105d36121ae565b600a60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff16600181111561069157fe5b600181111561069c57fe5b815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff1660018111156106d257fe5b60018111156106dd57fe5b81525050905080600001516007826020015160018111156106fa57fe5b8154811061070457fe5b906000526020600020018260400151836060015160088560800151600181111561072a57fe5b8154811061073457fe5b90600052602060002001838054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107d35780601f106107a8576101008083540402835291602001916107d3565b820191906000526020600020905b8154815290600101906020018083116107b657829003601f168201915b50505050509350808054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561086f5780601f106108445761010080835404028352916020019161086f565b820191906000526020600020905b81548152906001019060200180831161085257829003601f168201915b50505050509050955095509550955095505091939590929450565b600a6020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16908060020154908060030154908060040160009054906101000a900460ff16905085565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166109b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f73656e646572206973206e6f7420612072656c6179657200000000000000000081525060200191505060405180910390fd5b6001808111156109c557fe5b600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160009054906101000a900460ff166001811115610a2057fe5b14610a76576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c81526020018061226d602c913960400191505060405180910390fd5b600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610b76576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f72656c617965722068617320616c726561647920766f7465640000000000000081525060200191505060405180910390fd5b6001816001811115610b8457fe5b1115610bf8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f766f7465206f7574206f662074686520766f746520656e756d2072616e67650081525060200191505060405180910390fd5b600180811115610c0457fe5b816001811115610c1057fe5b1415610c6d57600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160008154809291906001019190505550610cc0565b600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301600081548092919060010191905055505b6001600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167fb6a4771df73f6cb79535dd475a69a813362d4b13633fdd88d406a3ed78a603328260405180826001811115610da157fe5b60ff16815260200191505060405180910390a2600054600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015410610eed57600180811115610e0b57fe5b600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff166001811115610e6657fe5b1415610e7a57610e7582611f43565b610e84565b610e8382611ff3565b5b6000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160006101000a81548160ff02191690836001811115610ee357fe5b0217905550610fb3565b600054610f47600a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301546001546120a490919063ffffffff16565b1015610fb2576000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160006101000a81548160ff02191690836001811115610fac57fe5b02179055505b5b5050565b6000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60096020528060005260406000206000915054906101000a900460ff1681565b6000806000606060026000015460028001546002600301546008600260040160009054906101000a900460ff16600181111561106557fe5b8154811061106f57fe5b90600052602060002001808054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561110e5780601f106110e35761010080835404028352916020019161110e565b820191906000526020600020905b8154815290600101906020018083116110f157829003601f168201915b50505050509050935093509350935090919293565b60008054905090565b60015481565b6000600154905090565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166111fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f73656e646572206973206e6f7420612072656c6179657200000000000000000081525060200191505060405180910390fd5b600181600181111561120957fe5b1115611260576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018061224c6021913960400191505060405180910390fd5b6000600181111561126d57fe5b81600181111561127957fe5b1480156112d6575060011515600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515145b611348576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f61646472657373206973206e6f7420612072656c61796572000000000000000081525060200191505060405180910390fd5b60018081111561135457fe5b81600181111561136057fe5b1480156113bd575060001515600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515145b61142f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f616464726573732069732063757272656e746c7920612072656c61796572000081525060200191505060405180910390fd5b6000600181111561143c57fe5b600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160009054906101000a900460ff16600181111561149757fe5b146114ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806122996034913960400191505060405180910390fd5b6040518060a001604052808373ffffffffffffffffffffffffffffffffffffffff16815260200182600181111561152057fe5b8152602001600181526020016000815260200160018081111561153f57fe5b815250600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360018111156115ec57fe5b0217905550604082015181600201556060820151816003015560808201518160040160006101000a81548160ff0219169083600181111561162957fe5b02179055509050506001600054116116d6576000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160006101000a81548160ff0219169083600181111561169a57fe5b02179055506001808111156116ab57fe5b8160018111156116b757fe5b14156116cb576116c682611f43565b6116d5565b6116d482611ff3565b5b5b6001600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555080600181111561177a57fe5b8273ffffffffffffffffffffffffffffffffffffffff167fb12a1fbc5d87bccd40c7ac1f669d842fa9fdfe7f65d2988ceac40543b646807e60405160405180910390a35050565b60005481565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611886576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f73656e646572206973206e6f7420612072656c6179657200000000000000000081525060200191505060405180910390fd5b6000600181111561189357fe5b600260040160009054906101000a900460ff1660018111156118b157fe5b14611924576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f612070726f706f73616c2069732063757272656e746c7920616374697665000081525060200191505060405180910390fd5b60015481111561197f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604281526020018061220a6042913960600191505060405180910390fd5b604051806080016040528082815260200160018152602001600081526020016001808111156119aa57fe5b815250600260008201518160000155602082015181600201556040820151816003015560608201518160040160006101000a81548160ff021916908360018111156119f157fe5b0217905550905050600160005411611a64576002600001546000819055506000600260040160006101000a81548160ff02191690836001811115611a3157fe5b0217905550807fa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c860405160405180910390a25b6001600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550807f79a7cdf066fccb1627ec87a3dd0bf8dcb3a186313f941c8bf80ed979aa62d38d60405160405180910390a250565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611bae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f73656e646572206973206e6f7420612072656c6179657200000000000000000081525060200191505060405180910390fd5b600180811115611bba57fe5b600260040160009054906101000a900460ff166001811115611bd857fe5b14611c4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f6e6f2070726f706f73616c2069732063757272656e746c79206163746976650081525060200191505060405180910390fd5b600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611d0e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f72656c617965722068617320616c726561647920766f7465640000000000000081525060200191505060405180910390fd5b6001816001811115611d1c57fe5b1115611d90576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f766f7465206f7574206f662074686520766f746520656e756d2072616e67650081525060200191505060405180910390fd5b600180811115611d9c57fe5b816001811115611da857fe5b1415611dc7576002800160008154809291906001019190505550611ddd565b6002600301600081548092919060010191905055505b6001600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fe6124fe2b7a19e7cdd807fb16247ec258d8fb0bfde2949b023bcb0c1eea19cb18160405180826001811115611e6a57fe5b60ff16815260200191505060405180910390a1600054600280015410611ef4576002600001546000819055506000600260040160006101000a81548160ff02191690836001811115611eb857fe5b02179055506002600001547fa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c860405160405180910390a2611f40565b600054611f116002600301546001546120a490919063ffffffff16565b1015611f3f576000600260040160006101000a81548160ff02191690836001811115611f3957fe5b02179055505b5b50565b6001600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600081548092919060010191905055508073ffffffffffffffffffffffffffffffffffffffff167f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c560405160405180910390a250565b6000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160008154809291906001900391905055508073ffffffffffffffffffffffffffffffffffffffff167f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b60405160405180910390a250565b60006120e683836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506120ee565b905092915050565b600083831115829061219b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612160578082015181840152602081019050612145565b50505050905090810190601f16801561218d5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060018111156121e357fe5b815260200160008152602001600081526020016000600181111561220357fe5b8152509056fe70726f706f7365642076616c75652063616e6e6f742062652067726561746572207468616e2074686520746f74616c206e756d626572206f662072656c6179657273616374696f6e206f7574206f662074686520766f746520656e756d2072616e67657468657265206973206e6f206163746976652070726f706f73616c20666f7220746869732061646472657373746865726520697320616c726561647920616e206163746976652070726f706f73616c20666f7220746869732061646472657373a2646970667358221220a3b64ba2f9b5befb7c7cdcf6545e2983b6117664460fc85812ea636512dc407f64736f6c63430006040033"
