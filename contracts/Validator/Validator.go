// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Validator

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

// ValidatorABI is the input ABI used to generate the binding from.
const ValidatorABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"initialValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"initialValidatorThreshold\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIValidator.ValidatorActionType\",\"name\":\"validatorActionType\",\"type\":\"uint8\"}],\"name\":\"ValidatorProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIValidator.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"ValidatorProposalVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"ValidatorThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposedValue\",\"type\":\"uint256\"}],\"name\":\"ValidatorThresholdProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumIValidator.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"ValidatorThresholdProposalVote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"_totalValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_validatorProposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_proposedAddress\",\"type\":\"address\"},{\"internalType\":\"enumIValidator.ValidatorActionType\",\"name\":\"_action\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_numYes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numNo\",\"type\":\"uint256\"},{\"internalType\":\"enumIValidator.VoteStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_validatorThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_validators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getTotalValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentValidatorThresholdProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"}],\"name\":\"getValidatorProposal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"},{\"internalType\":\"enumIValidator.ValidatorActionType\",\"name\":\"action\",\"type\":\"uint8\"}],\"name\":\"createValidatorProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"},{\"internalType\":\"enumIValidator.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"voteValidatorProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposedValue\",\"type\":\"uint256\"}],\"name\":\"createValidatorThresholdProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"enumIValidator.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"voteValidatorThresholdProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ValidatorBin is the compiled bytecode used for deploying new contracts.
var ValidatorBin = "0x608060405260405180604001604052806040518060400160405280600681526020017f72656d6f7665000000000000000000000000000000000000000000000000000081525081526020016040518060400160405280600381526020017f616464000000000000000000000000000000000000000000000000000000000081525081525060079060026200009592919062000304565b5060405180604001604052806040518060400160405280600881526020017f696e61637469766500000000000000000000000000000000000000000000000081525081526020016040518060400160405280600681526020017f616374697665000000000000000000000000000000000000000000000000000081525081525060089060026200012792919062000304565b503480156200013557600080fd5b50604051620027ab380380620027ab833981810160405260408110156200015b57600080fd5b81019080805160405193929190846401000000008211156200017c57600080fd5b838201915060208201858111156200019357600080fd5b8251866020820283011164010000000082111715620001b157600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620001ea578082015181840152602081019050620001cd565b505050509050016040526020018051906020019092919050505060005b82518110156200024457620002368382815181106200022257fe5b60200260200101516200025460201b60201c565b808060010191505062000207565b5080600081905550505062000497565b6001600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600081548092919060010191905055508073ffffffffffffffffffffffffffffffffffffffff167fe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec388498760405160405180910390a250565b82805482825590600052602060002090810192821562000358579160200282015b8281111562000357578251829080519060200190620003469291906200036b565b509160200191906001019062000325565b5b509050620003679190620003f2565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003ae57805160ff1916838001178555620003df565b82800160010185558215620003df579182015b82811115620003de578251825591602001919060010190620003c1565b5b509050620003ee919062000423565b5090565b6200042091905b808211156200041c57600081816200041291906200044b565b50600101620003f9565b5090565b90565b6200044891905b80821115620004445760008160009055506001016200042a565b5090565b90565b50805460018160011615610100020316600290046000825580601f1062000473575062000494565b601f01602090049060005260206000209081019062000493919062000423565b5b50565b61230480620004a76000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80638a7c016a1161008c578063d6932fec11610066578063d6932fec14610383578063d7cb54a2146104ed578063facd743b14610549578063fe434e0c146105a5576100cf565b80638a7c016a146103165780639bd61e6a14610334578063b63e24ec14610365576100cf565b8063295912df146100d4578063331ba913146100f257806359e44de71461014357806362e73263146101ff57806378bb28b9146102505780637fa9aaaf146102e8575b600080fd5b6100dc6105c3565b6040518082815260200191505060405180910390f35b6101416004803603604081101561010857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff1690602001909291905050506105cd565b005b6101856004803603602081101561015957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c52565b604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018560018111156101c757fe5b60ff1681526020018481526020018381526020018260018111156101e757fe5b60ff1681526020019550505050505060405180910390f35b61024e6004803603604081101561021557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff169060200190929190505050610cc2565b005b61025861137f565b6040518085815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156102aa57808201518184015260208101905061028f565b50505050905090810190601f1680156102d75780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b610314600480360360208110156102fe57600080fd5b8101908080359060200190929190505050611475565b005b61031e61179d565b6040518082815260200191505060405180910390f35b6103636004803603602081101561034a57600080fd5b81019080803560ff1690602001909291905050506117a3565b005b61036d611bf7565b6040518082815260200191505060405180910390f35b6103c56004803603602081101561039957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611bfd565b604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018060200185815260200184815260200180602001838103835287818151815260200191508051906020019080838360005b8381101561044757808201518184015260208101905061042c565b50505050905090810190601f1680156104745780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156104ad578082015181840152602081019050610492565b50505050905090810190601f1680156104da5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b61052f6004803603602081101561050357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611ec5565b604051808215151515815260200191505060405180910390f35b61058b6004803603602081101561055f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611ee5565b604051808215151515815260200191505060405180910390f35b6105ad611f3b565b6040518082815260200191505060405180910390f35b6000600154905090565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661068c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f73656e646572206973206e6f7420612076616c696461746f720000000000000081525060200191505060405180910390fd5b600181600181111561069a57fe5b11156106f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018061220b6021913960400191505060405180910390fd5b600060018111156106fe57fe5b81600181111561070a57fe5b148015610767575060011515600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515145b6107d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f61646472657373206973206e6f7420612076616c696461746f7200000000000081525060200191505060405180910390fd5b6001808111156107e557fe5b8160018111156107f157fe5b14801561084e575060001515600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515145b6108c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f616464726573732069732063757272656e746c7920612076616c696461746f7281525060200191505060405180910390fd5b600060018111156108cd57fe5b600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160009054906101000a900460ff16600181111561092857fe5b1461097e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603481526020018061229c6034913960400191505060405180910390fd5b6040518060a001604052808373ffffffffffffffffffffffffffffffffffffffff1681526020018260018111156109b157fe5b815260200160018152602001600081526020016001808111156109d057fe5b815250600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff02191690836001811115610a7d57fe5b0217905550604082015181600201556060820151816003015560808201518160040160006101000a81548160ff02191690836001811115610aba57fe5b0217905550905050600160005411610b67576000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160006101000a81548160ff02191690836001811115610b2b57fe5b0217905550600180811115610b3c57fe5b816001811115610b4857fe5b1415610b5c57610b5782611f44565b610b66565b610b6582611ff4565b5b5b6001600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550806001811115610c0b57fe5b8273ffffffffffffffffffffffffffffffffffffffff167fed03ea03758fc396e99de60b8dc7b559d5969f3665a5088be2b2cdd8fbebbd9260405160405180910390a35050565b600a6020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16908060020154908060030154908060040160009054906101000a900460ff16905085565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610d81576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f73656e646572206973206e6f7420612076616c696461746f720000000000000081525060200191505060405180910390fd5b600180811115610d8d57fe5b600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160009054906101000a900460ff166001811115610de857fe5b14610e3e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c81526020018061222c602c913960400191505060405180910390fd5b600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610f3e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f76616c696461746f722068617320616c726561647920766f746564000000000081525060200191505060405180910390fd5b6001816001811115610f4c57fe5b1115610fc0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f766f7465206f7574206f662074686520766f746520656e756d2072616e67650081525060200191505060405180910390fd5b600180811115610fcc57fe5b816001811115610fd857fe5b141561103557600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160008154809291906001019190505550611088565b600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301600081548092919060010191905055505b6001600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f673869db4419bd0f5d03308d0606a2dcdd44d44205032c739f5cc03d37bab176826040518082600181111561116957fe5b60ff16815260200191505060405180910390a2600054600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020154106112b5576001808111156111d357fe5b600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff16600181111561122e57fe5b14156112425761123d82611f44565b61124c565b61124b82611ff4565b5b6000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160006101000a81548160ff021916908360018111156112ab57fe5b021790555061137b565b60005461130f600a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301546001546120a590919063ffffffff16565b101561137a576000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160006101000a81548160ff0219169083600181111561137457fe5b02179055505b5b5050565b6000806000606060026000015460028001546002600301546008600260040160009054906101000a900460ff1660018111156113b757fe5b815481106113c157fe5b90600052602060002001808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156114605780601f1061143557610100808354040283529160200191611460565b820191906000526020600020905b81548152906001019060200180831161144357829003601f168201915b50505050509050935093509350935090919293565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611534576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f73656e646572206973206e6f7420612076616c696461746f720000000000000081525060200191505060405180910390fd5b6000600181111561154157fe5b600260040160009054906101000a900460ff16600181111561155f57fe5b146115d2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f612070726f706f73616c2069732063757272656e746c7920616374697665000081525060200191505060405180910390fd5b60015481111561162d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260448152602001806122586044913960600191505060405180910390fd5b6040518060800160405280828152602001600181526020016000815260200160018081111561165857fe5b815250600260008201518160000155602082015181600201556040820151816003015560608201518160040160006101000a81548160ff0219169083600181111561169f57fe5b0217905550905050600160005411611712576002600001546000819055506000600260040160006101000a81548160ff021916908360018111156116df57fe5b0217905550807f95c586411a87fb2ca28c34b78bce3bf57d0c29769a83ae46b484bd7fc049e8ee60405160405180910390a25b6001600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550807fc7876b43f0eccd5339779e2a0ddc65954cea640a0bc330145ead56fd205f107660405160405180910390a250565b60005481565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611862576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f73656e646572206973206e6f7420612076616c696461746f720000000000000081525060200191505060405180910390fd5b60018081111561186e57fe5b600260040160009054906101000a900460ff16600181111561188c57fe5b146118ff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f6e6f2070726f706f73616c2069732063757272656e746c79206163746976650081525060200191505060405180910390fd5b600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156119c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f76616c696461746f722068617320616c726561647920766f746564000000000081525060200191505060405180910390fd5b60018160018111156119d057fe5b1115611a44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f766f7465206f7574206f662074686520766f746520656e756d2072616e67650081525060200191505060405180910390fd5b600180811115611a5057fe5b816001811115611a5c57fe5b1415611a7b576002800160008154809291906001019190505550611a91565b6002600301600081548092919060010191905055505b6001600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f1c5e1f97eaeb3bcba56e89ed4597ed8074e3bf825c6a65528f5e27a621139abe8160405180826001811115611b1e57fe5b60ff16815260200191505060405180910390a1600054600280015410611ba8576002600001546000819055506000600260040160006101000a81548160ff02191690836001811115611b6c57fe5b02179055506002600001547f95c586411a87fb2ca28c34b78bce3bf57d0c29769a83ae46b484bd7fc049e8ee60405160405180910390a2611bf4565b600054611bc56002600301546001546120a590919063ffffffff16565b1015611bf3576000600260040160006101000a81548160ff02191690836001811115611bed57fe5b02179055505b5b50565b60015481565b600060606000806060611c0e6121af565b600a60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff166001811115611ccc57fe5b6001811115611cd757fe5b815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff166001811115611d0d57fe5b6001811115611d1857fe5b8152505090508060000151600782602001516001811115611d3557fe5b81548110611d3f57fe5b9060005260206000200182604001518360600151600885608001516001811115611d6557fe5b81548110611d6f57fe5b90600052602060002001838054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611e0e5780601f10611de357610100808354040283529160200191611e0e565b820191906000526020600020905b815481529060010190602001808311611df157829003601f168201915b50505050509350808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611eaa5780601f10611e7f57610100808354040283529160200191611eaa565b820191906000526020600020905b815481529060010190602001808311611e8d57829003601f168201915b50505050509050955095509550955095505091939590929450565b60096020528060005260406000206000915054906101000a900460ff1681565b6000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60008054905090565b6001600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600081548092919060010191905055508073ffffffffffffffffffffffffffffffffffffffff167fe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec388498760405160405180910390a250565b6000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160008154809291906001900391905055508073ffffffffffffffffffffffffffffffffffffffff167fe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f160405160405180910390a250565b60006120e783836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506120ef565b905092915050565b600083831115829061219c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612161578082015181840152602081019050612146565b50505050905090810190601f16801561218e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060018111156121e457fe5b815260200160008152602001600081526020016000600181111561220457fe5b8152509056fe616374696f6e206f7574206f662074686520766f746520656e756d2072616e67657468657265206973206e6f206163746976652070726f706f73616c20666f722074686973206164647265737370726f706f7365642076616c75652063616e6e6f742062652067726561746572207468616e2074686520746f74616c206e756d626572206f662076616c696461746f7273746865726520697320616c726561647920616e206163746976652070726f706f73616c20666f7220746869732061646472657373a265627a7a723158202121eebb43af4f381ac5977dde164da073fa81f6b5f313b98a612a61d6c452ab64736f6c63430005100032"

// DeployValidator deploys a new Ethereum contract, binding an instance of Validator to it.
func DeployValidator(auth *bind.TransactOpts, backend bind.ContractBackend, initialValidators []common.Address, initialValidatorThreshold *big.Int) (common.Address, *types.Transaction, *Validator, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorBin), backend, initialValidators, initialValidatorThreshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// Validator is an auto generated Go binding around an Ethereum contract.
type Validator struct {
	ValidatorCaller     // Read-only binding to the contract
	ValidatorTransactor // Write-only binding to the contract
	ValidatorFilterer   // Log filterer for contract events
}

// ValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSession struct {
	Contract     *Validator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorCallerSession struct {
	Contract *ValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorTransactorSession struct {
	Contract     *ValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRaw struct {
	Contract *Validator // Generic contract binding to access the raw methods on
}

// ValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorCallerRaw struct {
	Contract *ValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorTransactorRaw struct {
	Contract *ValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidator creates a new instance of Validator, bound to a specific deployed contract.
func NewValidator(address common.Address, backend bind.ContractBackend) (*Validator, error) {
	contract, err := bindValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// NewValidatorCaller creates a new read-only instance of Validator, bound to a specific deployed contract.
func NewValidatorCaller(address common.Address, caller bind.ContractCaller) (*ValidatorCaller, error) {
	contract, err := bindValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorCaller{contract: contract}, nil
}

// NewValidatorTransactor creates a new write-only instance of Validator, bound to a specific deployed contract.
func NewValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorTransactor, error) {
	contract, err := bindValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorTransactor{contract: contract}, nil
}

// NewValidatorFilterer creates a new log filterer instance of Validator, bound to a specific deployed contract.
func NewValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorFilterer, error) {
	contract, err := bindValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorFilterer{contract: contract}, nil
}

// bindValidator binds a generic wrapper to an already deployed contract.
func bindValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.ValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transact(opts, method, params...)
}

// TotalValidators is a free data retrieval call binding the contract method 0xb63e24ec.
//
// Solidity: function _totalValidators() constant returns(uint256)
func (_Validator *ValidatorCaller) TotalValidators(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "_totalValidators")
	return *ret0, err
}

// TotalValidators is a free data retrieval call binding the contract method 0xb63e24ec.
//
// Solidity: function _totalValidators() constant returns(uint256)
func (_Validator *ValidatorSession) TotalValidators() (*big.Int, error) {
	return _Validator.Contract.TotalValidators(&_Validator.CallOpts)
}

// TotalValidators is a free data retrieval call binding the contract method 0xb63e24ec.
//
// Solidity: function _totalValidators() constant returns(uint256)
func (_Validator *ValidatorCallerSession) TotalValidators() (*big.Int, error) {
	return _Validator.Contract.TotalValidators(&_Validator.CallOpts)
}

// ValidatorProposals is a free data retrieval call binding the contract method 0x59e44de7.
//
// Solidity: function _validatorProposals(address ) constant returns(address _proposedAddress, uint8 _action, uint256 _numYes, uint256 _numNo, uint8 _status)
func (_Validator *ValidatorCaller) ValidatorProposals(opts *bind.CallOpts, arg0 common.Address) (struct {
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
	err := _Validator.contract.Call(opts, out, "_validatorProposals", arg0)
	return *ret, err
}

// ValidatorProposals is a free data retrieval call binding the contract method 0x59e44de7.
//
// Solidity: function _validatorProposals(address ) constant returns(address _proposedAddress, uint8 _action, uint256 _numYes, uint256 _numNo, uint8 _status)
func (_Validator *ValidatorSession) ValidatorProposals(arg0 common.Address) (struct {
	ProposedAddress common.Address
	Action          uint8
	NumYes          *big.Int
	NumNo           *big.Int
	Status          uint8
}, error) {
	return _Validator.Contract.ValidatorProposals(&_Validator.CallOpts, arg0)
}

// ValidatorProposals is a free data retrieval call binding the contract method 0x59e44de7.
//
// Solidity: function _validatorProposals(address ) constant returns(address _proposedAddress, uint8 _action, uint256 _numYes, uint256 _numNo, uint8 _status)
func (_Validator *ValidatorCallerSession) ValidatorProposals(arg0 common.Address) (struct {
	ProposedAddress common.Address
	Action          uint8
	NumYes          *big.Int
	NumNo           *big.Int
	Status          uint8
}, error) {
	return _Validator.Contract.ValidatorProposals(&_Validator.CallOpts, arg0)
}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x8a7c016a.
//
// Solidity: function _validatorThreshold() constant returns(uint256)
func (_Validator *ValidatorCaller) ValidatorThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "_validatorThreshold")
	return *ret0, err
}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x8a7c016a.
//
// Solidity: function _validatorThreshold() constant returns(uint256)
func (_Validator *ValidatorSession) ValidatorThreshold() (*big.Int, error) {
	return _Validator.Contract.ValidatorThreshold(&_Validator.CallOpts)
}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x8a7c016a.
//
// Solidity: function _validatorThreshold() constant returns(uint256)
func (_Validator *ValidatorCallerSession) ValidatorThreshold() (*big.Int, error) {
	return _Validator.Contract.ValidatorThreshold(&_Validator.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0xd7cb54a2.
//
// Solidity: function _validators(address ) constant returns(bool)
func (_Validator *ValidatorCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "_validators", arg0)
	return *ret0, err
}

// Validators is a free data retrieval call binding the contract method 0xd7cb54a2.
//
// Solidity: function _validators(address ) constant returns(bool)
func (_Validator *ValidatorSession) Validators(arg0 common.Address) (bool, error) {
	return _Validator.Contract.Validators(&_Validator.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xd7cb54a2.
//
// Solidity: function _validators(address ) constant returns(bool)
func (_Validator *ValidatorCallerSession) Validators(arg0 common.Address) (bool, error) {
	return _Validator.Contract.Validators(&_Validator.CallOpts, arg0)
}

// GetCurrentValidatorThresholdProposal is a free data retrieval call binding the contract method 0x78bb28b9.
//
// Solidity: function getCurrentValidatorThresholdProposal() constant returns(uint256, uint256, uint256, string)
func (_Validator *ValidatorCaller) GetCurrentValidatorThresholdProposal(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, string, error) {
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
	err := _Validator.contract.Call(opts, out, "getCurrentValidatorThresholdProposal")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetCurrentValidatorThresholdProposal is a free data retrieval call binding the contract method 0x78bb28b9.
//
// Solidity: function getCurrentValidatorThresholdProposal() constant returns(uint256, uint256, uint256, string)
func (_Validator *ValidatorSession) GetCurrentValidatorThresholdProposal() (*big.Int, *big.Int, *big.Int, string, error) {
	return _Validator.Contract.GetCurrentValidatorThresholdProposal(&_Validator.CallOpts)
}

// GetCurrentValidatorThresholdProposal is a free data retrieval call binding the contract method 0x78bb28b9.
//
// Solidity: function getCurrentValidatorThresholdProposal() constant returns(uint256, uint256, uint256, string)
func (_Validator *ValidatorCallerSession) GetCurrentValidatorThresholdProposal() (*big.Int, *big.Int, *big.Int, string, error) {
	return _Validator.Contract.GetCurrentValidatorThresholdProposal(&_Validator.CallOpts)
}

// GetValidatorProposal is a free data retrieval call binding the contract method 0xd6932fec.
//
// Solidity: function getValidatorProposal(address proposedAddress) constant returns(address, string, uint256, uint256, string)
func (_Validator *ValidatorCaller) GetValidatorProposal(opts *bind.CallOpts, proposedAddress common.Address) (common.Address, string, *big.Int, *big.Int, string, error) {
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
	err := _Validator.contract.Call(opts, out, "getValidatorProposal", proposedAddress)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetValidatorProposal is a free data retrieval call binding the contract method 0xd6932fec.
//
// Solidity: function getValidatorProposal(address proposedAddress) constant returns(address, string, uint256, uint256, string)
func (_Validator *ValidatorSession) GetValidatorProposal(proposedAddress common.Address) (common.Address, string, *big.Int, *big.Int, string, error) {
	return _Validator.Contract.GetValidatorProposal(&_Validator.CallOpts, proposedAddress)
}

// GetValidatorProposal is a free data retrieval call binding the contract method 0xd6932fec.
//
// Solidity: function getValidatorProposal(address proposedAddress) constant returns(address, string, uint256, uint256, string)
func (_Validator *ValidatorCallerSession) GetValidatorProposal(proposedAddress common.Address) (common.Address, string, *big.Int, *big.Int, string, error) {
	return _Validator.Contract.GetValidatorProposal(&_Validator.CallOpts, proposedAddress)
}

// GetValidatorThreshold is a free data retrieval call binding the contract method 0xfe434e0c.
//
// Solidity: function getValidatorThreshold() constant returns(uint256)
func (_Validator *ValidatorCaller) GetValidatorThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "getValidatorThreshold")
	return *ret0, err
}

// GetValidatorThreshold is a free data retrieval call binding the contract method 0xfe434e0c.
//
// Solidity: function getValidatorThreshold() constant returns(uint256)
func (_Validator *ValidatorSession) GetValidatorThreshold() (*big.Int, error) {
	return _Validator.Contract.GetValidatorThreshold(&_Validator.CallOpts)
}

// GetValidatorThreshold is a free data retrieval call binding the contract method 0xfe434e0c.
//
// Solidity: function getValidatorThreshold() constant returns(uint256)
func (_Validator *ValidatorCallerSession) GetValidatorThreshold() (*big.Int, error) {
	return _Validator.Contract.GetValidatorThreshold(&_Validator.CallOpts)
}

// CreateValidatorProposal is a paid mutator transaction binding the contract method 0x331ba913.
//
// Solidity: function createValidatorProposal(address proposedAddress, uint8 action) returns()
func (_Validator *ValidatorTransactor) CreateValidatorProposal(opts *bind.TransactOpts, proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "createValidatorProposal", proposedAddress, action)
}

// CreateValidatorProposal is a paid mutator transaction binding the contract method 0x331ba913.
//
// Solidity: function createValidatorProposal(address proposedAddress, uint8 action) returns()
func (_Validator *ValidatorSession) CreateValidatorProposal(proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _Validator.Contract.CreateValidatorProposal(&_Validator.TransactOpts, proposedAddress, action)
}

// CreateValidatorProposal is a paid mutator transaction binding the contract method 0x331ba913.
//
// Solidity: function createValidatorProposal(address proposedAddress, uint8 action) returns()
func (_Validator *ValidatorTransactorSession) CreateValidatorProposal(proposedAddress common.Address, action uint8) (*types.Transaction, error) {
	return _Validator.Contract.CreateValidatorProposal(&_Validator.TransactOpts, proposedAddress, action)
}

// CreateValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x7fa9aaaf.
//
// Solidity: function createValidatorThresholdProposal(uint256 proposedValue) returns()
func (_Validator *ValidatorTransactor) CreateValidatorThresholdProposal(opts *bind.TransactOpts, proposedValue *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "createValidatorThresholdProposal", proposedValue)
}

// CreateValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x7fa9aaaf.
//
// Solidity: function createValidatorThresholdProposal(uint256 proposedValue) returns()
func (_Validator *ValidatorSession) CreateValidatorThresholdProposal(proposedValue *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.CreateValidatorThresholdProposal(&_Validator.TransactOpts, proposedValue)
}

// CreateValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x7fa9aaaf.
//
// Solidity: function createValidatorThresholdProposal(uint256 proposedValue) returns()
func (_Validator *ValidatorTransactorSession) CreateValidatorThresholdProposal(proposedValue *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.CreateValidatorThresholdProposal(&_Validator.TransactOpts, proposedValue)
}

// GetTotalValidators is a paid mutator transaction binding the contract method 0x295912df.
//
// Solidity: function getTotalValidators() returns(uint256)
func (_Validator *ValidatorTransactor) GetTotalValidators(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "getTotalValidators")
}

// GetTotalValidators is a paid mutator transaction binding the contract method 0x295912df.
//
// Solidity: function getTotalValidators() returns(uint256)
func (_Validator *ValidatorSession) GetTotalValidators() (*types.Transaction, error) {
	return _Validator.Contract.GetTotalValidators(&_Validator.TransactOpts)
}

// GetTotalValidators is a paid mutator transaction binding the contract method 0x295912df.
//
// Solidity: function getTotalValidators() returns(uint256)
func (_Validator *ValidatorTransactorSession) GetTotalValidators() (*types.Transaction, error) {
	return _Validator.Contract.GetTotalValidators(&_Validator.TransactOpts)
}

// IsValidator is a paid mutator transaction binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validatorAddress) returns(bool)
func (_Validator *ValidatorTransactor) IsValidator(opts *bind.TransactOpts, validatorAddress common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "isValidator", validatorAddress)
}

// IsValidator is a paid mutator transaction binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validatorAddress) returns(bool)
func (_Validator *ValidatorSession) IsValidator(validatorAddress common.Address) (*types.Transaction, error) {
	return _Validator.Contract.IsValidator(&_Validator.TransactOpts, validatorAddress)
}

// IsValidator is a paid mutator transaction binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validatorAddress) returns(bool)
func (_Validator *ValidatorTransactorSession) IsValidator(validatorAddress common.Address) (*types.Transaction, error) {
	return _Validator.Contract.IsValidator(&_Validator.TransactOpts, validatorAddress)
}

// VoteValidatorProposal is a paid mutator transaction binding the contract method 0x62e73263.
//
// Solidity: function voteValidatorProposal(address proposedAddress, uint8 vote) returns()
func (_Validator *ValidatorTransactor) VoteValidatorProposal(opts *bind.TransactOpts, proposedAddress common.Address, vote uint8) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "voteValidatorProposal", proposedAddress, vote)
}

// VoteValidatorProposal is a paid mutator transaction binding the contract method 0x62e73263.
//
// Solidity: function voteValidatorProposal(address proposedAddress, uint8 vote) returns()
func (_Validator *ValidatorSession) VoteValidatorProposal(proposedAddress common.Address, vote uint8) (*types.Transaction, error) {
	return _Validator.Contract.VoteValidatorProposal(&_Validator.TransactOpts, proposedAddress, vote)
}

// VoteValidatorProposal is a paid mutator transaction binding the contract method 0x62e73263.
//
// Solidity: function voteValidatorProposal(address proposedAddress, uint8 vote) returns()
func (_Validator *ValidatorTransactorSession) VoteValidatorProposal(proposedAddress common.Address, vote uint8) (*types.Transaction, error) {
	return _Validator.Contract.VoteValidatorProposal(&_Validator.TransactOpts, proposedAddress, vote)
}

// VoteValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x9bd61e6a.
//
// Solidity: function voteValidatorThresholdProposal(uint8 vote) returns()
func (_Validator *ValidatorTransactor) VoteValidatorThresholdProposal(opts *bind.TransactOpts, vote uint8) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "voteValidatorThresholdProposal", vote)
}

// VoteValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x9bd61e6a.
//
// Solidity: function voteValidatorThresholdProposal(uint8 vote) returns()
func (_Validator *ValidatorSession) VoteValidatorThresholdProposal(vote uint8) (*types.Transaction, error) {
	return _Validator.Contract.VoteValidatorThresholdProposal(&_Validator.TransactOpts, vote)
}

// VoteValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x9bd61e6a.
//
// Solidity: function voteValidatorThresholdProposal(uint8 vote) returns()
func (_Validator *ValidatorTransactorSession) VoteValidatorThresholdProposal(vote uint8) (*types.Transaction, error) {
	return _Validator.Contract.VoteValidatorThresholdProposal(&_Validator.TransactOpts, vote)
}

// ValidatorValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the Validator contract.
type ValidatorValidatorAddedIterator struct {
	Event *ValidatorValidatorAdded // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorAdded)
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
		it.Event = new(ValidatorValidatorAdded)
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
func (it *ValidatorValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorAdded represents a ValidatorAdded event raised by the Validator contract.
type ValidatorValidatorAdded struct {
	ValidatorAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validatorAddress)
func (_Validator *ValidatorFilterer) FilterValidatorAdded(opts *bind.FilterOpts, validatorAddress []common.Address) (*ValidatorValidatorAddedIterator, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorAdded", validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorAddedIterator{contract: _Validator.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validatorAddress)
func (_Validator *ValidatorFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorAdded, validatorAddress []common.Address) (event.Subscription, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorAdded", validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorAdded)
				if err := _Validator.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ParseValidatorAdded is a log parse operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validatorAddress)
func (_Validator *ValidatorFilterer) ParseValidatorAdded(log types.Log) (*ValidatorValidatorAdded, error) {
	event := new(ValidatorValidatorAdded)
	if err := _Validator.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorValidatorProposalCreatedIterator is returned from FilterValidatorProposalCreated and is used to iterate over the raw logs and unpacked data for ValidatorProposalCreated events raised by the Validator contract.
type ValidatorValidatorProposalCreatedIterator struct {
	Event *ValidatorValidatorProposalCreated // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorProposalCreated)
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
		it.Event = new(ValidatorValidatorProposalCreated)
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
func (it *ValidatorValidatorProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorProposalCreated represents a ValidatorProposalCreated event raised by the Validator contract.
type ValidatorValidatorProposalCreated struct {
	ProposedAddress     common.Address
	ValidatorActionType uint8
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorProposalCreated is a free log retrieval operation binding the contract event 0xed03ea03758fc396e99de60b8dc7b559d5969f3665a5088be2b2cdd8fbebbd92.
//
// Solidity: event ValidatorProposalCreated(address indexed proposedAddress, uint8 indexed validatorActionType)
func (_Validator *ValidatorFilterer) FilterValidatorProposalCreated(opts *bind.FilterOpts, proposedAddress []common.Address, validatorActionType []uint8) (*ValidatorValidatorProposalCreatedIterator, error) {

	var proposedAddressRule []interface{}
	for _, proposedAddressItem := range proposedAddress {
		proposedAddressRule = append(proposedAddressRule, proposedAddressItem)
	}
	var validatorActionTypeRule []interface{}
	for _, validatorActionTypeItem := range validatorActionType {
		validatorActionTypeRule = append(validatorActionTypeRule, validatorActionTypeItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorProposalCreated", proposedAddressRule, validatorActionTypeRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorProposalCreatedIterator{contract: _Validator.contract, event: "ValidatorProposalCreated", logs: logs, sub: sub}, nil
}

// WatchValidatorProposalCreated is a free log subscription operation binding the contract event 0xed03ea03758fc396e99de60b8dc7b559d5969f3665a5088be2b2cdd8fbebbd92.
//
// Solidity: event ValidatorProposalCreated(address indexed proposedAddress, uint8 indexed validatorActionType)
func (_Validator *ValidatorFilterer) WatchValidatorProposalCreated(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorProposalCreated, proposedAddress []common.Address, validatorActionType []uint8) (event.Subscription, error) {

	var proposedAddressRule []interface{}
	for _, proposedAddressItem := range proposedAddress {
		proposedAddressRule = append(proposedAddressRule, proposedAddressItem)
	}
	var validatorActionTypeRule []interface{}
	for _, validatorActionTypeItem := range validatorActionType {
		validatorActionTypeRule = append(validatorActionTypeRule, validatorActionTypeItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorProposalCreated", proposedAddressRule, validatorActionTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorProposalCreated)
				if err := _Validator.contract.UnpackLog(event, "ValidatorProposalCreated", log); err != nil {
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

// ParseValidatorProposalCreated is a log parse operation binding the contract event 0xed03ea03758fc396e99de60b8dc7b559d5969f3665a5088be2b2cdd8fbebbd92.
//
// Solidity: event ValidatorProposalCreated(address indexed proposedAddress, uint8 indexed validatorActionType)
func (_Validator *ValidatorFilterer) ParseValidatorProposalCreated(log types.Log) (*ValidatorValidatorProposalCreated, error) {
	event := new(ValidatorValidatorProposalCreated)
	if err := _Validator.contract.UnpackLog(event, "ValidatorProposalCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorValidatorProposalVoteIterator is returned from FilterValidatorProposalVote and is used to iterate over the raw logs and unpacked data for ValidatorProposalVote events raised by the Validator contract.
type ValidatorValidatorProposalVoteIterator struct {
	Event *ValidatorValidatorProposalVote // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorProposalVote)
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
		it.Event = new(ValidatorValidatorProposalVote)
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
func (it *ValidatorValidatorProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorProposalVote represents a ValidatorProposalVote event raised by the Validator contract.
type ValidatorValidatorProposalVote struct {
	ProposedAddress common.Address
	Vote            uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorProposalVote is a free log retrieval operation binding the contract event 0x673869db4419bd0f5d03308d0606a2dcdd44d44205032c739f5cc03d37bab176.
//
// Solidity: event ValidatorProposalVote(address indexed proposedAddress, uint8 vote)
func (_Validator *ValidatorFilterer) FilterValidatorProposalVote(opts *bind.FilterOpts, proposedAddress []common.Address) (*ValidatorValidatorProposalVoteIterator, error) {

	var proposedAddressRule []interface{}
	for _, proposedAddressItem := range proposedAddress {
		proposedAddressRule = append(proposedAddressRule, proposedAddressItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorProposalVote", proposedAddressRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorProposalVoteIterator{contract: _Validator.contract, event: "ValidatorProposalVote", logs: logs, sub: sub}, nil
}

// WatchValidatorProposalVote is a free log subscription operation binding the contract event 0x673869db4419bd0f5d03308d0606a2dcdd44d44205032c739f5cc03d37bab176.
//
// Solidity: event ValidatorProposalVote(address indexed proposedAddress, uint8 vote)
func (_Validator *ValidatorFilterer) WatchValidatorProposalVote(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorProposalVote, proposedAddress []common.Address) (event.Subscription, error) {

	var proposedAddressRule []interface{}
	for _, proposedAddressItem := range proposedAddress {
		proposedAddressRule = append(proposedAddressRule, proposedAddressItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorProposalVote", proposedAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorProposalVote)
				if err := _Validator.contract.UnpackLog(event, "ValidatorProposalVote", log); err != nil {
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

// ParseValidatorProposalVote is a log parse operation binding the contract event 0x673869db4419bd0f5d03308d0606a2dcdd44d44205032c739f5cc03d37bab176.
//
// Solidity: event ValidatorProposalVote(address indexed proposedAddress, uint8 vote)
func (_Validator *ValidatorFilterer) ParseValidatorProposalVote(log types.Log) (*ValidatorValidatorProposalVote, error) {
	event := new(ValidatorValidatorProposalVote)
	if err := _Validator.contract.UnpackLog(event, "ValidatorProposalVote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the Validator contract.
type ValidatorValidatorRemovedIterator struct {
	Event *ValidatorValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorRemoved)
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
		it.Event = new(ValidatorValidatorRemoved)
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
func (it *ValidatorValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorRemoved represents a ValidatorRemoved event raised by the Validator contract.
type ValidatorValidatorRemoved struct {
	ValidatorAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validatorAddress)
func (_Validator *ValidatorFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, validatorAddress []common.Address) (*ValidatorValidatorRemovedIterator, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorRemoved", validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorRemovedIterator{contract: _Validator.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validatorAddress)
func (_Validator *ValidatorFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorRemoved, validatorAddress []common.Address) (event.Subscription, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorRemoved", validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorRemoved)
				if err := _Validator.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validatorAddress)
func (_Validator *ValidatorFilterer) ParseValidatorRemoved(log types.Log) (*ValidatorValidatorRemoved, error) {
	event := new(ValidatorValidatorRemoved)
	if err := _Validator.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorValidatorThresholdChangedIterator is returned from FilterValidatorThresholdChanged and is used to iterate over the raw logs and unpacked data for ValidatorThresholdChanged events raised by the Validator contract.
type ValidatorValidatorThresholdChangedIterator struct {
	Event *ValidatorValidatorThresholdChanged // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorThresholdChanged)
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
		it.Event = new(ValidatorValidatorThresholdChanged)
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
func (it *ValidatorValidatorThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorThresholdChanged represents a ValidatorThresholdChanged event raised by the Validator contract.
type ValidatorValidatorThresholdChanged struct {
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorThresholdChanged is a free log retrieval operation binding the contract event 0x95c586411a87fb2ca28c34b78bce3bf57d0c29769a83ae46b484bd7fc049e8ee.
//
// Solidity: event ValidatorThresholdChanged(uint256 indexed newThreshold)
func (_Validator *ValidatorFilterer) FilterValidatorThresholdChanged(opts *bind.FilterOpts, newThreshold []*big.Int) (*ValidatorValidatorThresholdChangedIterator, error) {

	var newThresholdRule []interface{}
	for _, newThresholdItem := range newThreshold {
		newThresholdRule = append(newThresholdRule, newThresholdItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorThresholdChanged", newThresholdRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorThresholdChangedIterator{contract: _Validator.contract, event: "ValidatorThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorThresholdChanged is a free log subscription operation binding the contract event 0x95c586411a87fb2ca28c34b78bce3bf57d0c29769a83ae46b484bd7fc049e8ee.
//
// Solidity: event ValidatorThresholdChanged(uint256 indexed newThreshold)
func (_Validator *ValidatorFilterer) WatchValidatorThresholdChanged(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorThresholdChanged, newThreshold []*big.Int) (event.Subscription, error) {

	var newThresholdRule []interface{}
	for _, newThresholdItem := range newThreshold {
		newThresholdRule = append(newThresholdRule, newThresholdItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorThresholdChanged", newThresholdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorThresholdChanged)
				if err := _Validator.contract.UnpackLog(event, "ValidatorThresholdChanged", log); err != nil {
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

// ParseValidatorThresholdChanged is a log parse operation binding the contract event 0x95c586411a87fb2ca28c34b78bce3bf57d0c29769a83ae46b484bd7fc049e8ee.
//
// Solidity: event ValidatorThresholdChanged(uint256 indexed newThreshold)
func (_Validator *ValidatorFilterer) ParseValidatorThresholdChanged(log types.Log) (*ValidatorValidatorThresholdChanged, error) {
	event := new(ValidatorValidatorThresholdChanged)
	if err := _Validator.contract.UnpackLog(event, "ValidatorThresholdChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorValidatorThresholdProposalCreatedIterator is returned from FilterValidatorThresholdProposalCreated and is used to iterate over the raw logs and unpacked data for ValidatorThresholdProposalCreated events raised by the Validator contract.
type ValidatorValidatorThresholdProposalCreatedIterator struct {
	Event *ValidatorValidatorThresholdProposalCreated // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorThresholdProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorThresholdProposalCreated)
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
		it.Event = new(ValidatorValidatorThresholdProposalCreated)
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
func (it *ValidatorValidatorThresholdProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorThresholdProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorThresholdProposalCreated represents a ValidatorThresholdProposalCreated event raised by the Validator contract.
type ValidatorValidatorThresholdProposalCreated struct {
	ProposedValue *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorThresholdProposalCreated is a free log retrieval operation binding the contract event 0xc7876b43f0eccd5339779e2a0ddc65954cea640a0bc330145ead56fd205f1076.
//
// Solidity: event ValidatorThresholdProposalCreated(uint256 indexed proposedValue)
func (_Validator *ValidatorFilterer) FilterValidatorThresholdProposalCreated(opts *bind.FilterOpts, proposedValue []*big.Int) (*ValidatorValidatorThresholdProposalCreatedIterator, error) {

	var proposedValueRule []interface{}
	for _, proposedValueItem := range proposedValue {
		proposedValueRule = append(proposedValueRule, proposedValueItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorThresholdProposalCreated", proposedValueRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorThresholdProposalCreatedIterator{contract: _Validator.contract, event: "ValidatorThresholdProposalCreated", logs: logs, sub: sub}, nil
}

// WatchValidatorThresholdProposalCreated is a free log subscription operation binding the contract event 0xc7876b43f0eccd5339779e2a0ddc65954cea640a0bc330145ead56fd205f1076.
//
// Solidity: event ValidatorThresholdProposalCreated(uint256 indexed proposedValue)
func (_Validator *ValidatorFilterer) WatchValidatorThresholdProposalCreated(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorThresholdProposalCreated, proposedValue []*big.Int) (event.Subscription, error) {

	var proposedValueRule []interface{}
	for _, proposedValueItem := range proposedValue {
		proposedValueRule = append(proposedValueRule, proposedValueItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorThresholdProposalCreated", proposedValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorThresholdProposalCreated)
				if err := _Validator.contract.UnpackLog(event, "ValidatorThresholdProposalCreated", log); err != nil {
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

// ParseValidatorThresholdProposalCreated is a log parse operation binding the contract event 0xc7876b43f0eccd5339779e2a0ddc65954cea640a0bc330145ead56fd205f1076.
//
// Solidity: event ValidatorThresholdProposalCreated(uint256 indexed proposedValue)
func (_Validator *ValidatorFilterer) ParseValidatorThresholdProposalCreated(log types.Log) (*ValidatorValidatorThresholdProposalCreated, error) {
	event := new(ValidatorValidatorThresholdProposalCreated)
	if err := _Validator.contract.UnpackLog(event, "ValidatorThresholdProposalCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorValidatorThresholdProposalVoteIterator is returned from FilterValidatorThresholdProposalVote and is used to iterate over the raw logs and unpacked data for ValidatorThresholdProposalVote events raised by the Validator contract.
type ValidatorValidatorThresholdProposalVoteIterator struct {
	Event *ValidatorValidatorThresholdProposalVote // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorThresholdProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorThresholdProposalVote)
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
		it.Event = new(ValidatorValidatorThresholdProposalVote)
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
func (it *ValidatorValidatorThresholdProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorThresholdProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorThresholdProposalVote represents a ValidatorThresholdProposalVote event raised by the Validator contract.
type ValidatorValidatorThresholdProposalVote struct {
	Vote uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorThresholdProposalVote is a free log retrieval operation binding the contract event 0x1c5e1f97eaeb3bcba56e89ed4597ed8074e3bf825c6a65528f5e27a621139abe.
//
// Solidity: event ValidatorThresholdProposalVote(uint8 vote)
func (_Validator *ValidatorFilterer) FilterValidatorThresholdProposalVote(opts *bind.FilterOpts) (*ValidatorValidatorThresholdProposalVoteIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorThresholdProposalVote")
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorThresholdProposalVoteIterator{contract: _Validator.contract, event: "ValidatorThresholdProposalVote", logs: logs, sub: sub}, nil
}

// WatchValidatorThresholdProposalVote is a free log subscription operation binding the contract event 0x1c5e1f97eaeb3bcba56e89ed4597ed8074e3bf825c6a65528f5e27a621139abe.
//
// Solidity: event ValidatorThresholdProposalVote(uint8 vote)
func (_Validator *ValidatorFilterer) WatchValidatorThresholdProposalVote(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorThresholdProposalVote) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorThresholdProposalVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorThresholdProposalVote)
				if err := _Validator.contract.UnpackLog(event, "ValidatorThresholdProposalVote", log); err != nil {
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

// ParseValidatorThresholdProposalVote is a log parse operation binding the contract event 0x1c5e1f97eaeb3bcba56e89ed4597ed8074e3bf825c6a65528f5e27a621139abe.
//
// Solidity: event ValidatorThresholdProposalVote(uint8 vote)
func (_Validator *ValidatorFilterer) ParseValidatorThresholdProposalVote(log types.Log) (*ValidatorValidatorThresholdProposalVote, error) {
	event := new(ValidatorValidatorThresholdProposalVote)
	if err := _Validator.contract.UnpackLog(event, "ValidatorThresholdProposalVote", log); err != nil {
		return nil, err
	}
	return event, nil
}
