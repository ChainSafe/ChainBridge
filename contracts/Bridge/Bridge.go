// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Bridge

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

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialValidatorThreshold\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"DepositProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"DepositProposalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumBridge.Vote\",\"name\":\"vote\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumBridge.DepositProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"DepositProposalVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"ERC20Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"ERC721Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"GenericDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"ValidatorThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposedValue\",\"type\":\"uint256\"}],\"name\":\"ValidatorThresholdProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumBridge.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"ValidatorThresholdProposalVote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"_currentValidatorThresholdProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposedValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numYes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numNo\",\"type\":\"uint256\"},{\"internalType\":\"enumBridge.ValidatorThresholdProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositProposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_originChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_depositID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_numYes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numNo\",\"type\":\"uint256\"},{\"internalType\":\"enumBridge.DepositProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_erc20DepositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_originChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destinationChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_erc721DepositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_originChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destinationChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_genericDepositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_originChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destinationChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_validatorContract\",\"outputs\":[{\"internalType\":\"contractIValidator\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_validatorThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"}],\"name\":\"getDepositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getGenericDepositRecord\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getERC20DepositRecord\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getERC721DepositRecord\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentValidatorThresholdProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"originChainContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"depositGeneric\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"depositGeneric\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"depositERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"createDepositProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"},{\"internalType\":\"enumBridge.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"voteDepositProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationChainHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDepositProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposedValue\",\"type\":\"uint256\"}],\"name\":\"createValidatorThresholdProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"enumBridge.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"voteValidatorThresholdProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x608060405260405180604001604052806040518060400160405280600881526020017f696e61637469766500000000000000000000000000000000000000000000000081525081526020016040518060400160405280600681526020017f616374697665000000000000000000000000000000000000000000000000000081525081525060079060026200009592919062000276565b506040518060a001604052806040518060400160405280600881526020017f696e61637469766500000000000000000000000000000000000000000000000081525081526020016040518060400160405280600681526020017f616374697665000000000000000000000000000000000000000000000000000081525081526020016040518060400160405280600681526020017f64656e696564000000000000000000000000000000000000000000000000000081525081526020016040518060400160405280600681526020017f706173736564000000000000000000000000000000000000000000000000000081525081526020016040518060400160405280600b81526020017f7472616e736665727265640000000000000000000000000000000000000000008152508152506008906005620001d8929190620002dd565b50348015620001e657600080fd5b5060405162004c1938038062004c19833981810160405260408110156200020c57600080fd5b810190808051906020019092919080519060200190929190505050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600181905550505062000470565b828054828255906000526020600020908101928215620002ca579160200282015b82811115620002c9578251829080519060200190620002b892919062000344565b509160200191906001019062000297565b5b509050620002d99190620003cb565b5090565b82805482825590600052602060002090810192821562000331579160200282015b82811115620003305782518290805190602001906200031f92919062000344565b5091602001919060010190620002fe565b5b509050620003409190620003cb565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200038757805160ff1916838001178555620003b8565b82800160010185558215620003b8579182015b82811115620003b75782518255916020019190600101906200039a565b5b509050620003c79190620003fc565b5090565b620003f991905b80821115620003f55760008181620003eb919062000424565b50600101620003d2565b5090565b90565b6200042191905b808211156200041d57600081600090555060010162000403565b5090565b90565b50805460018160011615610100020316600290046000825580601f106200044c57506200046d565b601f0160209004906000526020600020908101906200046c9190620003fc565b5b50565b61479980620004806000396000f3fe608060405234801561001057600080fd5b50600436106101735760003560e01c80639bd61e6a116100de578063cb4815e011610097578063e7d3696c11610071578063e7d3696c14611041578063ef35ac5714611086578063f95644b9146111d5578063fe434e0c1461136057610173565b8063cb4815e014610ee2578063d852416e14610fb6578063dd8ecfe61461100057610173565b80639bd61e6a14610bac578063a157727614610bdd578063b06e1cc314610cfc578063b613a6ac14610d79578063b8260def14610dbb578063ba39157314610ea057610173565b806378bb28b91161013057806378bb28b9146106e05780637fa9aaaf1461077857806386ae761d146107a65780638a7c016a146108c55780638cd1a91c146108e35780639829ecca14610a2857610173565b8063030e7ddd146101785780630d39a1d31461026757806326397c9d146102d75780632d994e1f1461045b578063356e124f1461049d578063683fd38814610628575b600080fd5b6102656004803603608081101561018e57600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156101df57600080fd5b8201836020820111156101f157600080fd5b8035906020019184600183028401116401000000008311171561021357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061137e565b005b6102bd6004803603606081101561027d57600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061162a565b604051808215151515815260200191505060405180910390f35b61030d600480360360408110156102ed57600080fd5b8101908080359060200190929190803590602001909291905050506116a7565b604051808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561041b578082015181840152602081019050610400565b50505050905090810190601f1680156104485780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b61049b6004803603606081101561047157600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050611926565b005b6104d3600480360360408110156104b357600080fd5b810190808035906020019092919080359060200190929190505050611d42565b604051808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156105e75780820151818401526020810190506105cc565b50505050905090810190601f1680156106145780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6106de600480360360c081101561063e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611ea9565b005b6106e86121b8565b6040518085815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561073a57808201518184015260208101905061071f565b50505050905090810190601f1680156107675780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b6107a46004803603602081101561078e57600080fd5b81019080803590602001909291905050506122ae565b005b6107dc600480360360408110156107bc57600080fd5b8101908080359060200190929190803590602001909291905050506126a1565b604051808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001965050505050505060405180910390f35b6108cd612884565b6040518082815260200191505060405180910390f35b610a26600480360360c08110156108f957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156109a057600080fd5b8201836020820111156109b257600080fd5b803590602001918460018302840111640100000000831117156109d457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061288a565b005b610a5e60048036036040811015610a3e57600080fd5b810190808035906020019092919080359060200190929190505050612ad3565b604051808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610b6c578082015181840152602081019050610b51565b50505050905090810190601f168015610b995780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b610bdb60048036036020811015610bc257600080fd5b81019080803560ff169060200190929190505050612c34565b005b610c1360048036036040811015610bf357600080fd5b810190808035906020019092919080359060200190929190505050613154565b604051808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001965050505050505060405180910390f35b610d3260048036036040811015610d1257600080fd5b81019080803590602001909291908035906020019092919050505061321d565b60405180878152602001868152602001858152602001848152602001838152602001826004811115610d6057fe5b60ff168152602001965050505050505060405180910390f35b610da560048036036020811015610d8f57600080fd5b8101908080359060200190929190505050613273565b6040518082815260200191505060405180910390f35b610e9e60048036036060811015610dd157600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190640100000000811115610e1857600080fd5b820183602082011115610e2a57600080fd5b80359060200191846001830284011164010000000083111715610e4c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050613290565b005b610ecc60048036036020811015610eb657600080fd5b81019080803590602001909291905050506134d9565b6040518082815260200191505060405180910390f35b610f1860048036036040811015610ef857600080fd5b8101908080359060200190929190803590602001909291905050506134f1565b6040518087815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610f76578082015181840152602081019050610f5b565b50505050905090810190601f168015610fa35780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b610fbe613685565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6110086136aa565b6040518085815260200184815260200183815260200182600181111561102a57fe5b60ff16815260200194505050505060405180910390f35b6110846004803603606081101561105757600080fd5b810190808035906020019092919080359060200190929190803560ff1690602001909291905050506136d5565b005b6111d3600480360360e081101561109c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561114d57600080fd5b82018360208201111561115f57600080fd5b8035906020019184600183028401116401000000008311171561118157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050613cf5565b005b61120b600480360360408110156111eb57600080fd5b810190808035906020019092919080359060200190929190505050614028565b604051808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561131f578082015181840152602081019050611304565b50505050905090810190601f16801561134c5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6113686142bb565b6040518082815260200191505060405180910390f35b6000600d600086815260200190815260200160002060008581526020019081526020016000209050600060048111156113b357fe5b8160060160009054906101000a900460ff1660048111156113d057fe5b1415611444576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f70726f706f73616c206973206e6f74206163746976650000000000000000000081525060200191505060405180910390fd5b6003600481111561145157fe5b8160060160009054906101000a900460ff16600481111561146e57fe5b146114c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260378152602001806146ea6037913960400191505060405180910390fd5b8060020154828051906020012014611527576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260318152602001806146b96031913960400191505060405180910390fd5b60008390508073ffffffffffffffffffffffffffffffffffffffff1663fc9539cd846040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561159857808201518184015260208101905061157d565b50505050905090810190601f1680156115c55780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b1580156115e457600080fd5b505af11580156115f8573d6000803e3d6000fd5b5050505060048260060160006101000a81548160ff0219169083600481111561161d57fe5b0217905550505050505050565b6000600d6000858152602001908152602001600020600084815260200190815260200160002060030160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690509392505050565b600080600080600060606116b96143cf565b600a60008a815260200190815260200160002060008981526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600582018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156118e35780601f106118b8576101008083540402835291602001916118e3565b820191906000526020600020905b8154815290600101906020018083116118c657829003601f168201915b5050505050815250509050806000015181602001518260400151836060015184608001518560a00151809050965096509650965096509650509295509295509295565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663facd743b336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156119cb57600080fd5b505af11580156119df573d6000803e3d6000fd5b505050506040513d60208110156119f557600080fd5b8101908080519060200190929190505050611a0f57600080fd5b60006004811115611a1c57fe5b600d6000868152602001908152602001600020600085815260200190815260200160002060060160009054906101000a900460ff166004811115611a5c57fe5b1480611ab2575060026004811115611a7057fe5b600d6000868152602001908152602001600020600085815260200190815260200160002060060160009054906101000a900460ff166004811115611ab057fe5b145b611b07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604f81526020018061466a604f913960600191505060405180910390fd5b6001805411611bd2576040518060c00160405280858152602001848152602001838152602001600181526020016000815260200160036004811115611b4857fe5b815250600d60008681526020019081526020016000206000858152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600401556080820151816005015560a08201518160060160006101000a81548160ff02191690836004811115611bc557fe5b0217905550905050611c90565b6040518060c00160405280858152602001848152602001838152602001600181526020016000815260200160016004811115611c0a57fe5b815250600d60008681526020019081526020016000206000858152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600401556080820151816005015560a08201518160060160006101000a81548160ff02191690836004811115611c8757fe5b02179055509050505b6001600d6000868152602001908152602001600020600085815260200190815260200160002060030160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508183857fd79daef3492e6769de74ee113f9383423c3e43830a24b1be81e49dceecc639fd60405160405180910390a450505050565b600c602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806005015490806006018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611e9f5780601f10611e7457610100808354040283529160200191611e9f565b820191906000526020600020905b815481529060010190602001808311611e8257829003601f168201915b5050505050905087565b60008590508073ffffffffffffffffffffffffffffffffffffffff16631cad5a408833856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015611f6957600080fd5b505af1158015611f7d573d6000803e3d6000fd5b505050506000600960008781526020019081526020016000206000815460010191905081905590506040518060c001604052808973ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018781526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff16815260200184815250600b6000888152602001908152602001600020600083815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050155905050807f2199cfceaff95f329a0db94bd11b97a8f85f9e983679d755198045fd67851e0460405160405180910390a25050505050505050565b6000806000606060026000015460028001546002600301546007600260040160009054906101000a900460ff1660018111156121f057fe5b815481106121fa57fe5b90600052602060002001808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156122995780601f1061226e57610100808354040283529160200191612299565b820191906000526020600020905b81548152906001019060200180831161227c57829003601f168201915b50505050509050935093509350935090919293565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663facd743b336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561235357600080fd5b505af1158015612367573d6000803e3d6000fd5b505050506040513d602081101561237d57600080fd5b810190808051906020019092919050505061239757600080fd5b600060018111156123a457fe5b600260040160009054906101000a900460ff1660018111156123c257fe5b14612435576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f612070726f706f73616c2069732063757272656e746c7920616374697665000081525060200191505060405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663295912df6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561249e57600080fd5b505af11580156124b2573d6000803e3d6000fd5b505050506040513d60208110156124c857600080fd5b8101908080519060200190929190505050821115612531576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260448152602001806147216044913960600191505060405180910390fd5b6040518060800160405280838152602001600181526020016000815260200160018081111561255c57fe5b815250600260008201518160000155602082015181600201556040820151816003015560608201518160040160006101000a81548160ff021916908360018111156125a357fe5b02179055509050506001805411612615576002600001546001819055506000600260040160006101000a81548160ff021916908360018111156125e257fe5b0217905550817f95c586411a87fb2ca28c34b78bce3bf57d0c29769a83ae46b484bd7fc049e8ee60405160405180910390a25b6001600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550817fc7876b43f0eccd5339779e2a0ddc65954cea640a0bc330145ead56fd205f107660405160405180910390a25050565b6000806000806000806126b261445d565b600b60008a815260200190815260200160002060008981526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050806000015181602001518260400151836060015184608001518560a00151965096509650965096509650509295509295509295565b60015481565b6000600960008681526020019081526020016000206000815460010191905081905590506040518060c001604052808873ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183815250600a6000878152602001908152602001600020600083815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a0820151816005019080519060200190612a999291906144eb565b50905050807f5f7c049578fd3ad850c1b539b94320934c5afc6dc53f4eab709dc45ee4cabe8160405160405180910390a250505050505050565b600a602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806005018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612c2a5780601f10612bff57610100808354040283529160200191612c2a565b820191906000526020600020905b815481529060010190602001808311612c0d57829003601f168201915b5050505050905086565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663facd743b336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015612cd957600080fd5b505af1158015612ced573d6000803e3d6000fd5b505050506040513d6020811015612d0357600080fd5b8101908080519060200190929190505050612d1d57600080fd5b600180811115612d2957fe5b600260040160009054906101000a900460ff166001811115612d4757fe5b14612dba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f6e6f2070726f706f73616c2069732063757272656e746c79206163746976650081525060200191505060405180910390fd5b600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615612e7d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f76616c696461746f722068617320616c726561647920766f746564000000000081525060200191505060405180910390fd5b6001826001811115612e8b57fe5b1115612eff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f766f7465206f7574206f662074686520766f746520656e756d2072616e67650081525060200191505060405180910390fd5b600180811115612f0b57fe5b826001811115612f1757fe5b1415612f36576002800160008154809291906001019190505550612f4c565b6002600301600081548092919060010191905055505b6001600260010160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f1c5e1f97eaeb3bcba56e89ed4597ed8074e3bf825c6a65528f5e27a621139abe8260405180826001811115612fd957fe5b60ff16815260200191505060405180910390a1600154600280015410613063576002600001546001819055506000600260040160006101000a81548160ff0219169083600181111561302757fe5b02179055506002600001547f95c586411a87fb2ca28c34b78bce3bf57d0c29769a83ae46b484bd7fc049e8ee60405160405180910390a2613150565b6001546131216002600301546000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663295912df6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156130d857600080fd5b505af11580156130ec573d6000803e3d6000fd5b505050506040513d602081101561310257600080fd5b81019080805190602001909291905050506142c590919063ffffffff16565b101561314f576000600260040160006101000a81548160ff0219169083600181111561314957fe5b02179055505b5b5050565b600b602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154905086565b600d602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060020154908060040154908060050154908060060160009054906101000a900460ff16905086565b600060096000838152602001908152602001600020549050919050565b6000600960008581526020019081526020016000206000815460010191905081905590506040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001858152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183815250600a6000868152602001908152602001600020600083815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050190805190602001906134a29291906144eb565b50905050807f5f7c049578fd3ad850c1b539b94320934c5afc6dc53f4eab709dc45ee4cabe8160405160405180910390a250505050565b60096020528060005260406000206000915090505481565b6000806000806000606061350361456b565b600d60008a815260200190815260200160002060008981526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff16600481111561358257fe5b600481111561358d57fe5b8152505090508060000151816020015182604001518360600151846080015160088660a0015160048111156135be57fe5b815481106135c857fe5b90600052602060002001808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156136675780601f1061363c57610100808354040283529160200191613667565b820191906000526020600020905b81548152906001019060200180831161364a57829003601f168201915b50505050509050965096509650965096509650509295509295509295565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60028060000154908060020154908060030154908060040160009054906101000a900460ff16905084565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663facd743b336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561377a57600080fd5b505af115801561378e573d6000803e3d6000fd5b505050506040513d60208110156137a457600080fd5b81019080805190602001909291905050506137be57600080fd5b6000600d600086815260200190815260200160002060008581526020019081526020016000209050600060048111156137f357fe5b8160060160009054906101000a900460ff16600481111561381057fe5b1415613884576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f70726f706f73616c206973206e6f74206163746976650000000000000000000081525060200191505060405180910390fd5b6001600481111561389157fe5b8160060160009054906101000a900460ff1660048111156138ae57fe5b14613921576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f70726f706f73616c20686173206265656e2066696e616c697a6564000000000081525060200191505060405180910390fd5b8060030160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156139e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f76616c696461746f722068617320616c726561647920766f746564000000000081525060200191505060405180910390fd5b60018360018111156139f157fe5b1115613a65576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f696e76616c696420766f7465000000000000000000000000000000000000000081525060200191505060405180910390fd5b600180811115613a7157fe5b836001811115613a7d57fe5b1415613a9c578060040160008154809291906001019190505550613ab1565b80600501600081548092919060010191905055505b60018160030160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600154816004015410613b715760038160060160006101000a81548160ff02191690836004811115613b3957fe5b021790555083857fbc724fbe62906cf888d9fef0688a894cdea9b5c36290addb9d815699d343a99d60405160405180910390a3613c8a565b600154613c2e82600501546000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663295912df6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015613be557600080fd5b505af1158015613bf9573d6000803e3d6000fd5b505050506040513d6020811015613c0f57600080fd5b81019080805190602001909291905050506142c590919063ffffffff16565b1015613c895760028160060160006101000a81548160ff02191690836004811115613c5557fe5b021790555083857fbc724fbe62906cf888d9fef0688a894cdea9b5c36290addb9d815699d343a99d60405160405180910390a35b5b826001811115613c9657fe5b84867f40fc0ff8dcb86fbd4d70d3e1237dbd6dab58a03c8a1a3af7b27a4f70b540a66b8460060160009054906101000a900460ff1660405180826004811115613cdb57fe5b60ff16815260200191505060405180910390a45050505050565b60008690508073ffffffffffffffffffffffffffffffffffffffff1663331ded1a8933866040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015613db557600080fd5b505af1158015613dc9573d6000803e3d6000fd5b505050506000600960008881526020019081526020016000206000815460010191905081905590506040518060e001604052808a73ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff1681526020018881526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200184815250600c6000898152602001908152602001600020600083815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a0820151816005015560c0820151816006019080519060200190613fec9291906144eb565b50905050807f14f91f5d57b330e0bf94ee405aac47227ebc44b43634f4f0bf64052913dcfed260405160405180910390a2505050505050505050565b600080600080600080606061403b6145af565b600c60008b815260200190815260200160002060008a81526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160058201548152602001600682018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561426f5780601f106142445761010080835404028352916020019161426f565b820191906000526020600020905b81548152906001019060200180831161425257829003601f168201915b5050505050815250509050806000015181602001518260400151836060015184608001518560a001518660c0015180905097509750975097509750975097505092959891949750929550565b6000600154905090565b600061430783836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061430f565b905092915050565b60008383111582906143bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015614381578082015181840152602081019050614366565b50505050905090810190601f1680156143ae5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061452c57805160ff191683800117855561455a565b8280016001018555821561455a579182015b8281111561455957825182559160200191906001019061453e565b5b5090506145679190614644565b5090565b6040518060c001604052806000815260200160008152602001600080191681526020016000815260200160008152602001600060048111156145a957fe5b81525090565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001606081525090565b61466691905b8082111561466257600081600090555060010161464a565b5090565b9056fe746869732070726f706f73616c206973206569746865722063757272656e746c7920616374697665206f722068617320616c7265616479206265656e207061737365642f7472616e7366657272656470726f7669646564206461746120646f6573206e6f74206d617463682070726f706f73616c27732064617461206861736870726f706f73616c20776173206e6f7420706173736564206f722068617320616c7265616479206265656e207472616e7366657272656470726f706f7365642076616c75652063616e6e6f742062652067726561746572207468616e2074686520746f74616c206e756d626572206f662076616c696461746f7273a265627a7a723158200531a67937738838a5a1bb6696a7cdb4c0b6144557cf4d24b3020855afef84e064736f6c63430005100032"

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, validatorContract common.Address, initialValidatorThreshold *big.Int) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend, validatorContract, initialValidatorThreshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// CurrentValidatorThresholdProposal is a free data retrieval call binding the contract method 0xdd8ecfe6.
//
// Solidity: function _currentValidatorThresholdProposal() constant returns(uint256 _proposedValue, uint256 _numYes, uint256 _numNo, uint8 _status)
func (_Bridge *BridgeCaller) CurrentValidatorThresholdProposal(opts *bind.CallOpts) (struct {
	ProposedValue *big.Int
	NumYes        *big.Int
	NumNo         *big.Int
	Status        uint8
}, error) {
	ret := new(struct {
		ProposedValue *big.Int
		NumYes        *big.Int
		NumNo         *big.Int
		Status        uint8
	})
	out := ret
	err := _Bridge.contract.Call(opts, out, "_currentValidatorThresholdProposal")
	return *ret, err
}

// CurrentValidatorThresholdProposal is a free data retrieval call binding the contract method 0xdd8ecfe6.
//
// Solidity: function _currentValidatorThresholdProposal() constant returns(uint256 _proposedValue, uint256 _numYes, uint256 _numNo, uint8 _status)
func (_Bridge *BridgeSession) CurrentValidatorThresholdProposal() (struct {
	ProposedValue *big.Int
	NumYes        *big.Int
	NumNo         *big.Int
	Status        uint8
}, error) {
	return _Bridge.Contract.CurrentValidatorThresholdProposal(&_Bridge.CallOpts)
}

// CurrentValidatorThresholdProposal is a free data retrieval call binding the contract method 0xdd8ecfe6.
//
// Solidity: function _currentValidatorThresholdProposal() constant returns(uint256 _proposedValue, uint256 _numYes, uint256 _numNo, uint8 _status)
func (_Bridge *BridgeCallerSession) CurrentValidatorThresholdProposal() (struct {
	ProposedValue *big.Int
	NumYes        *big.Int
	NumNo         *big.Int
	Status        uint8
}, error) {
	return _Bridge.Contract.CurrentValidatorThresholdProposal(&_Bridge.CallOpts)
}

// DepositCounts is a free data retrieval call binding the contract method 0xba391573.
//
// Solidity: function _depositCounts(uint256 ) constant returns(uint256)
func (_Bridge *BridgeCaller) DepositCounts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "_depositCounts", arg0)
	return *ret0, err
}

// DepositCounts is a free data retrieval call binding the contract method 0xba391573.
//
// Solidity: function _depositCounts(uint256 ) constant returns(uint256)
func (_Bridge *BridgeSession) DepositCounts(arg0 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.DepositCounts(&_Bridge.CallOpts, arg0)
}

// DepositCounts is a free data retrieval call binding the contract method 0xba391573.
//
// Solidity: function _depositCounts(uint256 ) constant returns(uint256)
func (_Bridge *BridgeCallerSession) DepositCounts(arg0 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.DepositCounts(&_Bridge.CallOpts, arg0)
}

// DepositProposals is a free data retrieval call binding the contract method 0xb06e1cc3.
//
// Solidity: function _depositProposals(uint256 , uint256 ) constant returns(uint256 _originChainID, uint256 _depositID, bytes32 _dataHash, uint256 _numYes, uint256 _numNo, uint8 _status)
func (_Bridge *BridgeCaller) DepositProposals(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	OriginChainID *big.Int
	DepositID     *big.Int
	DataHash      [32]byte
	NumYes        *big.Int
	NumNo         *big.Int
	Status        uint8
}, error) {
	ret := new(struct {
		OriginChainID *big.Int
		DepositID     *big.Int
		DataHash      [32]byte
		NumYes        *big.Int
		NumNo         *big.Int
		Status        uint8
	})
	out := ret
	err := _Bridge.contract.Call(opts, out, "_depositProposals", arg0, arg1)
	return *ret, err
}

// DepositProposals is a free data retrieval call binding the contract method 0xb06e1cc3.
//
// Solidity: function _depositProposals(uint256 , uint256 ) constant returns(uint256 _originChainID, uint256 _depositID, bytes32 _dataHash, uint256 _numYes, uint256 _numNo, uint8 _status)
func (_Bridge *BridgeSession) DepositProposals(arg0 *big.Int, arg1 *big.Int) (struct {
	OriginChainID *big.Int
	DepositID     *big.Int
	DataHash      [32]byte
	NumYes        *big.Int
	NumNo         *big.Int
	Status        uint8
}, error) {
	return _Bridge.Contract.DepositProposals(&_Bridge.CallOpts, arg0, arg1)
}

// DepositProposals is a free data retrieval call binding the contract method 0xb06e1cc3.
//
// Solidity: function _depositProposals(uint256 , uint256 ) constant returns(uint256 _originChainID, uint256 _depositID, bytes32 _dataHash, uint256 _numYes, uint256 _numNo, uint8 _status)
func (_Bridge *BridgeCallerSession) DepositProposals(arg0 *big.Int, arg1 *big.Int) (struct {
	OriginChainID *big.Int
	DepositID     *big.Int
	DataHash      [32]byte
	NumYes        *big.Int
	NumNo         *big.Int
	Status        uint8
}, error) {
	return _Bridge.Contract.DepositProposals(&_Bridge.CallOpts, arg0, arg1)
}

// Erc20DepositRecords is a free data retrieval call binding the contract method 0xa1577276.
//
// Solidity: function _erc20DepositRecords(uint256 , uint256 ) constant returns(address _originChainTokenAddress, address _originChainHandlerAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationRecipientAddress, uint256 _amount)
func (_Bridge *BridgeCaller) Erc20DepositRecords(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	OriginChainHandlerAddress      common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationRecipientAddress    common.Address
	Amount                         *big.Int
}, error) {
	ret := new(struct {
		OriginChainTokenAddress        common.Address
		OriginChainHandlerAddress      common.Address
		DestinationChainID             *big.Int
		DestinationChainHandlerAddress common.Address
		DestinationRecipientAddress    common.Address
		Amount                         *big.Int
	})
	out := ret
	err := _Bridge.contract.Call(opts, out, "_erc20DepositRecords", arg0, arg1)
	return *ret, err
}

// Erc20DepositRecords is a free data retrieval call binding the contract method 0xa1577276.
//
// Solidity: function _erc20DepositRecords(uint256 , uint256 ) constant returns(address _originChainTokenAddress, address _originChainHandlerAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationRecipientAddress, uint256 _amount)
func (_Bridge *BridgeSession) Erc20DepositRecords(arg0 *big.Int, arg1 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	OriginChainHandlerAddress      common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationRecipientAddress    common.Address
	Amount                         *big.Int
}, error) {
	return _Bridge.Contract.Erc20DepositRecords(&_Bridge.CallOpts, arg0, arg1)
}

// Erc20DepositRecords is a free data retrieval call binding the contract method 0xa1577276.
//
// Solidity: function _erc20DepositRecords(uint256 , uint256 ) constant returns(address _originChainTokenAddress, address _originChainHandlerAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationRecipientAddress, uint256 _amount)
func (_Bridge *BridgeCallerSession) Erc20DepositRecords(arg0 *big.Int, arg1 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	OriginChainHandlerAddress      common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationRecipientAddress    common.Address
	Amount                         *big.Int
}, error) {
	return _Bridge.Contract.Erc20DepositRecords(&_Bridge.CallOpts, arg0, arg1)
}

// Erc721DepositRecords is a free data retrieval call binding the contract method 0x356e124f.
//
// Solidity: function _erc721DepositRecords(uint256 , uint256 ) constant returns(address _originChainTokenAddress, address _originChainHandlerAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationRecipientAddress, uint256 _tokenID, bytes _data)
func (_Bridge *BridgeCaller) Erc721DepositRecords(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	OriginChainHandlerAddress      common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationRecipientAddress    common.Address
	TokenID                        *big.Int
	Data                           []byte
}, error) {
	ret := new(struct {
		OriginChainTokenAddress        common.Address
		OriginChainHandlerAddress      common.Address
		DestinationChainID             *big.Int
		DestinationChainHandlerAddress common.Address
		DestinationRecipientAddress    common.Address
		TokenID                        *big.Int
		Data                           []byte
	})
	out := ret
	err := _Bridge.contract.Call(opts, out, "_erc721DepositRecords", arg0, arg1)
	return *ret, err
}

// Erc721DepositRecords is a free data retrieval call binding the contract method 0x356e124f.
//
// Solidity: function _erc721DepositRecords(uint256 , uint256 ) constant returns(address _originChainTokenAddress, address _originChainHandlerAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationRecipientAddress, uint256 _tokenID, bytes _data)
func (_Bridge *BridgeSession) Erc721DepositRecords(arg0 *big.Int, arg1 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	OriginChainHandlerAddress      common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationRecipientAddress    common.Address
	TokenID                        *big.Int
	Data                           []byte
}, error) {
	return _Bridge.Contract.Erc721DepositRecords(&_Bridge.CallOpts, arg0, arg1)
}

// Erc721DepositRecords is a free data retrieval call binding the contract method 0x356e124f.
//
// Solidity: function _erc721DepositRecords(uint256 , uint256 ) constant returns(address _originChainTokenAddress, address _originChainHandlerAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationRecipientAddress, uint256 _tokenID, bytes _data)
func (_Bridge *BridgeCallerSession) Erc721DepositRecords(arg0 *big.Int, arg1 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	OriginChainHandlerAddress      common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationRecipientAddress    common.Address
	TokenID                        *big.Int
	Data                           []byte
}, error) {
	return _Bridge.Contract.Erc721DepositRecords(&_Bridge.CallOpts, arg0, arg1)
}

// GenericDepositRecords is a free data retrieval call binding the contract method 0x9829ecca.
//
// Solidity: function _genericDepositRecords(uint256 , uint256 ) constant returns(address _originChainTokenAddress, address _originChainHandlerAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationRecipientAddress, bytes _data)
func (_Bridge *BridgeCaller) GenericDepositRecords(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	OriginChainHandlerAddress      common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationRecipientAddress    common.Address
	Data                           []byte
}, error) {
	ret := new(struct {
		OriginChainTokenAddress        common.Address
		OriginChainHandlerAddress      common.Address
		DestinationChainID             *big.Int
		DestinationChainHandlerAddress common.Address
		DestinationRecipientAddress    common.Address
		Data                           []byte
	})
	out := ret
	err := _Bridge.contract.Call(opts, out, "_genericDepositRecords", arg0, arg1)
	return *ret, err
}

// GenericDepositRecords is a free data retrieval call binding the contract method 0x9829ecca.
//
// Solidity: function _genericDepositRecords(uint256 , uint256 ) constant returns(address _originChainTokenAddress, address _originChainHandlerAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationRecipientAddress, bytes _data)
func (_Bridge *BridgeSession) GenericDepositRecords(arg0 *big.Int, arg1 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	OriginChainHandlerAddress      common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationRecipientAddress    common.Address
	Data                           []byte
}, error) {
	return _Bridge.Contract.GenericDepositRecords(&_Bridge.CallOpts, arg0, arg1)
}

// GenericDepositRecords is a free data retrieval call binding the contract method 0x9829ecca.
//
// Solidity: function _genericDepositRecords(uint256 , uint256 ) constant returns(address _originChainTokenAddress, address _originChainHandlerAddress, uint256 _destinationChainID, address _destinationChainHandlerAddress, address _destinationRecipientAddress, bytes _data)
func (_Bridge *BridgeCallerSession) GenericDepositRecords(arg0 *big.Int, arg1 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	OriginChainHandlerAddress      common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationRecipientAddress    common.Address
	Data                           []byte
}, error) {
	return _Bridge.Contract.GenericDepositRecords(&_Bridge.CallOpts, arg0, arg1)
}

// ValidatorContract is a free data retrieval call binding the contract method 0xd852416e.
//
// Solidity: function _validatorContract() constant returns(address)
func (_Bridge *BridgeCaller) ValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "_validatorContract")
	return *ret0, err
}

// ValidatorContract is a free data retrieval call binding the contract method 0xd852416e.
//
// Solidity: function _validatorContract() constant returns(address)
func (_Bridge *BridgeSession) ValidatorContract() (common.Address, error) {
	return _Bridge.Contract.ValidatorContract(&_Bridge.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0xd852416e.
//
// Solidity: function _validatorContract() constant returns(address)
func (_Bridge *BridgeCallerSession) ValidatorContract() (common.Address, error) {
	return _Bridge.Contract.ValidatorContract(&_Bridge.CallOpts)
}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x8a7c016a.
//
// Solidity: function _validatorThreshold() constant returns(uint256)
func (_Bridge *BridgeCaller) ValidatorThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "_validatorThreshold")
	return *ret0, err
}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x8a7c016a.
//
// Solidity: function _validatorThreshold() constant returns(uint256)
func (_Bridge *BridgeSession) ValidatorThreshold() (*big.Int, error) {
	return _Bridge.Contract.ValidatorThreshold(&_Bridge.CallOpts)
}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x8a7c016a.
//
// Solidity: function _validatorThreshold() constant returns(uint256)
func (_Bridge *BridgeCallerSession) ValidatorThreshold() (*big.Int, error) {
	return _Bridge.Contract.ValidatorThreshold(&_Bridge.CallOpts)
}

// GetCurrentValidatorThresholdProposal is a free data retrieval call binding the contract method 0x78bb28b9.
//
// Solidity: function getCurrentValidatorThresholdProposal() constant returns(uint256, uint256, uint256, string)
func (_Bridge *BridgeCaller) GetCurrentValidatorThresholdProposal(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, string, error) {
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
	err := _Bridge.contract.Call(opts, out, "getCurrentValidatorThresholdProposal")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetCurrentValidatorThresholdProposal is a free data retrieval call binding the contract method 0x78bb28b9.
//
// Solidity: function getCurrentValidatorThresholdProposal() constant returns(uint256, uint256, uint256, string)
func (_Bridge *BridgeSession) GetCurrentValidatorThresholdProposal() (*big.Int, *big.Int, *big.Int, string, error) {
	return _Bridge.Contract.GetCurrentValidatorThresholdProposal(&_Bridge.CallOpts)
}

// GetCurrentValidatorThresholdProposal is a free data retrieval call binding the contract method 0x78bb28b9.
//
// Solidity: function getCurrentValidatorThresholdProposal() constant returns(uint256, uint256, uint256, string)
func (_Bridge *BridgeCallerSession) GetCurrentValidatorThresholdProposal() (*big.Int, *big.Int, *big.Int, string, error) {
	return _Bridge.Contract.GetCurrentValidatorThresholdProposal(&_Bridge.CallOpts)
}

// GetDepositCount is a free data retrieval call binding the contract method 0xb613a6ac.
//
// Solidity: function getDepositCount(uint256 originChainID) constant returns(uint256)
func (_Bridge *BridgeCaller) GetDepositCount(opts *bind.CallOpts, originChainID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "getDepositCount", originChainID)
	return *ret0, err
}

// GetDepositCount is a free data retrieval call binding the contract method 0xb613a6ac.
//
// Solidity: function getDepositCount(uint256 originChainID) constant returns(uint256)
func (_Bridge *BridgeSession) GetDepositCount(originChainID *big.Int) (*big.Int, error) {
	return _Bridge.Contract.GetDepositCount(&_Bridge.CallOpts, originChainID)
}

// GetDepositCount is a free data retrieval call binding the contract method 0xb613a6ac.
//
// Solidity: function getDepositCount(uint256 originChainID) constant returns(uint256)
func (_Bridge *BridgeCallerSession) GetDepositCount(originChainID *big.Int) (*big.Int, error) {
	return _Bridge.Contract.GetDepositCount(&_Bridge.CallOpts, originChainID)
}

// GetDepositProposal is a free data retrieval call binding the contract method 0xcb4815e0.
//
// Solidity: function getDepositProposal(uint256 originChainID, uint256 depositID) constant returns(uint256, uint256, bytes32, uint256, uint256, string)
func (_Bridge *BridgeCaller) GetDepositProposal(opts *bind.CallOpts, originChainID *big.Int, depositID *big.Int) (*big.Int, *big.Int, [32]byte, *big.Int, *big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new([32]byte)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _Bridge.contract.Call(opts, out, "getDepositProposal", originChainID, depositID)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetDepositProposal is a free data retrieval call binding the contract method 0xcb4815e0.
//
// Solidity: function getDepositProposal(uint256 originChainID, uint256 depositID) constant returns(uint256, uint256, bytes32, uint256, uint256, string)
func (_Bridge *BridgeSession) GetDepositProposal(originChainID *big.Int, depositID *big.Int) (*big.Int, *big.Int, [32]byte, *big.Int, *big.Int, string, error) {
	return _Bridge.Contract.GetDepositProposal(&_Bridge.CallOpts, originChainID, depositID)
}

// GetDepositProposal is a free data retrieval call binding the contract method 0xcb4815e0.
//
// Solidity: function getDepositProposal(uint256 originChainID, uint256 depositID) constant returns(uint256, uint256, bytes32, uint256, uint256, string)
func (_Bridge *BridgeCallerSession) GetDepositProposal(originChainID *big.Int, depositID *big.Int) (*big.Int, *big.Int, [32]byte, *big.Int, *big.Int, string, error) {
	return _Bridge.Contract.GetDepositProposal(&_Bridge.CallOpts, originChainID, depositID)
}

// GetERC20DepositRecord is a free data retrieval call binding the contract method 0x86ae761d.
//
// Solidity: function getERC20DepositRecord(uint256 originChainID, uint256 depositID) constant returns(address, address, uint256, address, address, uint256)
func (_Bridge *BridgeCaller) GetERC20DepositRecord(opts *bind.CallOpts, originChainID *big.Int, depositID *big.Int) (common.Address, common.Address, *big.Int, common.Address, common.Address, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(common.Address)
		ret4 = new(common.Address)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _Bridge.contract.Call(opts, out, "getERC20DepositRecord", originChainID, depositID)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetERC20DepositRecord is a free data retrieval call binding the contract method 0x86ae761d.
//
// Solidity: function getERC20DepositRecord(uint256 originChainID, uint256 depositID) constant returns(address, address, uint256, address, address, uint256)
func (_Bridge *BridgeSession) GetERC20DepositRecord(originChainID *big.Int, depositID *big.Int) (common.Address, common.Address, *big.Int, common.Address, common.Address, *big.Int, error) {
	return _Bridge.Contract.GetERC20DepositRecord(&_Bridge.CallOpts, originChainID, depositID)
}

// GetERC20DepositRecord is a free data retrieval call binding the contract method 0x86ae761d.
//
// Solidity: function getERC20DepositRecord(uint256 originChainID, uint256 depositID) constant returns(address, address, uint256, address, address, uint256)
func (_Bridge *BridgeCallerSession) GetERC20DepositRecord(originChainID *big.Int, depositID *big.Int) (common.Address, common.Address, *big.Int, common.Address, common.Address, *big.Int, error) {
	return _Bridge.Contract.GetERC20DepositRecord(&_Bridge.CallOpts, originChainID, depositID)
}

// GetERC721DepositRecord is a free data retrieval call binding the contract method 0xf95644b9.
//
// Solidity: function getERC721DepositRecord(uint256 originChainID, uint256 depositID) constant returns(address, address, uint256, address, address, uint256, bytes)
func (_Bridge *BridgeCaller) GetERC721DepositRecord(opts *bind.CallOpts, originChainID *big.Int, depositID *big.Int) (common.Address, common.Address, *big.Int, common.Address, common.Address, *big.Int, []byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(common.Address)
		ret4 = new(common.Address)
		ret5 = new(*big.Int)
		ret6 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _Bridge.contract.Call(opts, out, "getERC721DepositRecord", originChainID, depositID)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetERC721DepositRecord is a free data retrieval call binding the contract method 0xf95644b9.
//
// Solidity: function getERC721DepositRecord(uint256 originChainID, uint256 depositID) constant returns(address, address, uint256, address, address, uint256, bytes)
func (_Bridge *BridgeSession) GetERC721DepositRecord(originChainID *big.Int, depositID *big.Int) (common.Address, common.Address, *big.Int, common.Address, common.Address, *big.Int, []byte, error) {
	return _Bridge.Contract.GetERC721DepositRecord(&_Bridge.CallOpts, originChainID, depositID)
}

// GetERC721DepositRecord is a free data retrieval call binding the contract method 0xf95644b9.
//
// Solidity: function getERC721DepositRecord(uint256 originChainID, uint256 depositID) constant returns(address, address, uint256, address, address, uint256, bytes)
func (_Bridge *BridgeCallerSession) GetERC721DepositRecord(originChainID *big.Int, depositID *big.Int) (common.Address, common.Address, *big.Int, common.Address, common.Address, *big.Int, []byte, error) {
	return _Bridge.Contract.GetERC721DepositRecord(&_Bridge.CallOpts, originChainID, depositID)
}

// GetGenericDepositRecord is a free data retrieval call binding the contract method 0x26397c9d.
//
// Solidity: function getGenericDepositRecord(uint256 originChainID, uint256 depositID) constant returns(address, address, uint256, address, address, bytes)
func (_Bridge *BridgeCaller) GetGenericDepositRecord(opts *bind.CallOpts, originChainID *big.Int, depositID *big.Int) (common.Address, common.Address, *big.Int, common.Address, common.Address, []byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(common.Address)
		ret4 = new(common.Address)
		ret5 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _Bridge.contract.Call(opts, out, "getGenericDepositRecord", originChainID, depositID)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetGenericDepositRecord is a free data retrieval call binding the contract method 0x26397c9d.
//
// Solidity: function getGenericDepositRecord(uint256 originChainID, uint256 depositID) constant returns(address, address, uint256, address, address, bytes)
func (_Bridge *BridgeSession) GetGenericDepositRecord(originChainID *big.Int, depositID *big.Int) (common.Address, common.Address, *big.Int, common.Address, common.Address, []byte, error) {
	return _Bridge.Contract.GetGenericDepositRecord(&_Bridge.CallOpts, originChainID, depositID)
}

// GetGenericDepositRecord is a free data retrieval call binding the contract method 0x26397c9d.
//
// Solidity: function getGenericDepositRecord(uint256 originChainID, uint256 depositID) constant returns(address, address, uint256, address, address, bytes)
func (_Bridge *BridgeCallerSession) GetGenericDepositRecord(originChainID *big.Int, depositID *big.Int) (common.Address, common.Address, *big.Int, common.Address, common.Address, []byte, error) {
	return _Bridge.Contract.GetGenericDepositRecord(&_Bridge.CallOpts, originChainID, depositID)
}

// GetValidatorThreshold is a free data retrieval call binding the contract method 0xfe434e0c.
//
// Solidity: function getValidatorThreshold() constant returns(uint256)
func (_Bridge *BridgeCaller) GetValidatorThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "getValidatorThreshold")
	return *ret0, err
}

// GetValidatorThreshold is a free data retrieval call binding the contract method 0xfe434e0c.
//
// Solidity: function getValidatorThreshold() constant returns(uint256)
func (_Bridge *BridgeSession) GetValidatorThreshold() (*big.Int, error) {
	return _Bridge.Contract.GetValidatorThreshold(&_Bridge.CallOpts)
}

// GetValidatorThreshold is a free data retrieval call binding the contract method 0xfe434e0c.
//
// Solidity: function getValidatorThreshold() constant returns(uint256)
func (_Bridge *BridgeCallerSession) GetValidatorThreshold() (*big.Int, error) {
	return _Bridge.Contract.GetValidatorThreshold(&_Bridge.CallOpts)
}

// HasVoted is a free data retrieval call binding the contract method 0x0d39a1d3.
//
// Solidity: function hasVoted(uint256 originChainID, uint256 depositID, address validatorAddress) constant returns(bool)
func (_Bridge *BridgeCaller) HasVoted(opts *bind.CallOpts, originChainID *big.Int, depositID *big.Int, validatorAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "hasVoted", originChainID, depositID, validatorAddress)
	return *ret0, err
}

// HasVoted is a free data retrieval call binding the contract method 0x0d39a1d3.
//
// Solidity: function hasVoted(uint256 originChainID, uint256 depositID, address validatorAddress) constant returns(bool)
func (_Bridge *BridgeSession) HasVoted(originChainID *big.Int, depositID *big.Int, validatorAddress common.Address) (bool, error) {
	return _Bridge.Contract.HasVoted(&_Bridge.CallOpts, originChainID, depositID, validatorAddress)
}

// HasVoted is a free data retrieval call binding the contract method 0x0d39a1d3.
//
// Solidity: function hasVoted(uint256 originChainID, uint256 depositID, address validatorAddress) constant returns(bool)
func (_Bridge *BridgeCallerSession) HasVoted(originChainID *big.Int, depositID *big.Int, validatorAddress common.Address) (bool, error) {
	return _Bridge.Contract.HasVoted(&_Bridge.CallOpts, originChainID, depositID, validatorAddress)
}

// CreateDepositProposal is a paid mutator transaction binding the contract method 0x2d994e1f.
//
// Solidity: function createDepositProposal(uint256 originChainID, uint256 depositID, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactor) CreateDepositProposal(opts *bind.TransactOpts, originChainID *big.Int, depositID *big.Int, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "createDepositProposal", originChainID, depositID, dataHash)
}

// CreateDepositProposal is a paid mutator transaction binding the contract method 0x2d994e1f.
//
// Solidity: function createDepositProposal(uint256 originChainID, uint256 depositID, bytes32 dataHash) returns()
func (_Bridge *BridgeSession) CreateDepositProposal(originChainID *big.Int, depositID *big.Int, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.CreateDepositProposal(&_Bridge.TransactOpts, originChainID, depositID, dataHash)
}

// CreateDepositProposal is a paid mutator transaction binding the contract method 0x2d994e1f.
//
// Solidity: function createDepositProposal(uint256 originChainID, uint256 depositID, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactorSession) CreateDepositProposal(originChainID *big.Int, depositID *big.Int, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.CreateDepositProposal(&_Bridge.TransactOpts, originChainID, depositID, dataHash)
}

// CreateValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x7fa9aaaf.
//
// Solidity: function createValidatorThresholdProposal(uint256 proposedValue) returns()
func (_Bridge *BridgeTransactor) CreateValidatorThresholdProposal(opts *bind.TransactOpts, proposedValue *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "createValidatorThresholdProposal", proposedValue)
}

// CreateValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x7fa9aaaf.
//
// Solidity: function createValidatorThresholdProposal(uint256 proposedValue) returns()
func (_Bridge *BridgeSession) CreateValidatorThresholdProposal(proposedValue *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.CreateValidatorThresholdProposal(&_Bridge.TransactOpts, proposedValue)
}

// CreateValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x7fa9aaaf.
//
// Solidity: function createValidatorThresholdProposal(uint256 proposedValue) returns()
func (_Bridge *BridgeTransactorSession) CreateValidatorThresholdProposal(proposedValue *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.CreateValidatorThresholdProposal(&_Bridge.TransactOpts, proposedValue)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x683fd388.
//
// Solidity: function depositERC20(address originChainTokenAddress, address originChainHandlerAddress, uint256 destinationChainID, address destinationChainHandlerAddress, address destinationRecipientAddress, uint256 amount) returns()
func (_Bridge *BridgeTransactor) DepositERC20(opts *bind.TransactOpts, originChainTokenAddress common.Address, originChainHandlerAddress common.Address, destinationChainID *big.Int, destinationChainHandlerAddress common.Address, destinationRecipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositERC20", originChainTokenAddress, originChainHandlerAddress, destinationChainID, destinationChainHandlerAddress, destinationRecipientAddress, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x683fd388.
//
// Solidity: function depositERC20(address originChainTokenAddress, address originChainHandlerAddress, uint256 destinationChainID, address destinationChainHandlerAddress, address destinationRecipientAddress, uint256 amount) returns()
func (_Bridge *BridgeSession) DepositERC20(originChainTokenAddress common.Address, originChainHandlerAddress common.Address, destinationChainID *big.Int, destinationChainHandlerAddress common.Address, destinationRecipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC20(&_Bridge.TransactOpts, originChainTokenAddress, originChainHandlerAddress, destinationChainID, destinationChainHandlerAddress, destinationRecipientAddress, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x683fd388.
//
// Solidity: function depositERC20(address originChainTokenAddress, address originChainHandlerAddress, uint256 destinationChainID, address destinationChainHandlerAddress, address destinationRecipientAddress, uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) DepositERC20(originChainTokenAddress common.Address, originChainHandlerAddress common.Address, destinationChainID *big.Int, destinationChainHandlerAddress common.Address, destinationRecipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC20(&_Bridge.TransactOpts, originChainTokenAddress, originChainHandlerAddress, destinationChainID, destinationChainHandlerAddress, destinationRecipientAddress, amount)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xef35ac57.
//
// Solidity: function depositERC721(address originChainTokenAddress, address originChainHandlerAddress, uint256 destinationChainID, address destinationChainHandlerAddress, address destinationRecipientAddress, uint256 tokenID, bytes data) returns()
func (_Bridge *BridgeTransactor) DepositERC721(opts *bind.TransactOpts, originChainTokenAddress common.Address, originChainHandlerAddress common.Address, destinationChainID *big.Int, destinationChainHandlerAddress common.Address, destinationRecipientAddress common.Address, tokenID *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositERC721", originChainTokenAddress, originChainHandlerAddress, destinationChainID, destinationChainHandlerAddress, destinationRecipientAddress, tokenID, data)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xef35ac57.
//
// Solidity: function depositERC721(address originChainTokenAddress, address originChainHandlerAddress, uint256 destinationChainID, address destinationChainHandlerAddress, address destinationRecipientAddress, uint256 tokenID, bytes data) returns()
func (_Bridge *BridgeSession) DepositERC721(originChainTokenAddress common.Address, originChainHandlerAddress common.Address, destinationChainID *big.Int, destinationChainHandlerAddress common.Address, destinationRecipientAddress common.Address, tokenID *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC721(&_Bridge.TransactOpts, originChainTokenAddress, originChainHandlerAddress, destinationChainID, destinationChainHandlerAddress, destinationRecipientAddress, tokenID, data)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xef35ac57.
//
// Solidity: function depositERC721(address originChainTokenAddress, address originChainHandlerAddress, uint256 destinationChainID, address destinationChainHandlerAddress, address destinationRecipientAddress, uint256 tokenID, bytes data) returns()
func (_Bridge *BridgeTransactorSession) DepositERC721(originChainTokenAddress common.Address, originChainHandlerAddress common.Address, destinationChainID *big.Int, destinationChainHandlerAddress common.Address, destinationRecipientAddress common.Address, tokenID *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC721(&_Bridge.TransactOpts, originChainTokenAddress, originChainHandlerAddress, destinationChainID, destinationChainHandlerAddress, destinationRecipientAddress, tokenID, data)
}

// DepositGeneric is a paid mutator transaction binding the contract method 0x8cd1a91c.
//
// Solidity: function depositGeneric(address originChainContractAddress, address originChainHandlerAddress, uint256 destinationChainID, address destinationChainHandlerAddress, address destinationRecipientAddress, bytes data) returns()
func (_Bridge *BridgeTransactor) DepositGeneric(opts *bind.TransactOpts, originChainContractAddress common.Address, originChainHandlerAddress common.Address, destinationChainID *big.Int, destinationChainHandlerAddress common.Address, destinationRecipientAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositGeneric", originChainContractAddress, originChainHandlerAddress, destinationChainID, destinationChainHandlerAddress, destinationRecipientAddress, data)
}

// DepositGeneric is a paid mutator transaction binding the contract method 0x8cd1a91c.
//
// Solidity: function depositGeneric(address originChainContractAddress, address originChainHandlerAddress, uint256 destinationChainID, address destinationChainHandlerAddress, address destinationRecipientAddress, bytes data) returns()
func (_Bridge *BridgeSession) DepositGeneric(originChainContractAddress common.Address, originChainHandlerAddress common.Address, destinationChainID *big.Int, destinationChainHandlerAddress common.Address, destinationRecipientAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.DepositGeneric(&_Bridge.TransactOpts, originChainContractAddress, originChainHandlerAddress, destinationChainID, destinationChainHandlerAddress, destinationRecipientAddress, data)
}

// DepositGeneric is a paid mutator transaction binding the contract method 0x8cd1a91c.
//
// Solidity: function depositGeneric(address originChainContractAddress, address originChainHandlerAddress, uint256 destinationChainID, address destinationChainHandlerAddress, address destinationRecipientAddress, bytes data) returns()
func (_Bridge *BridgeTransactorSession) DepositGeneric(originChainContractAddress common.Address, originChainHandlerAddress common.Address, destinationChainID *big.Int, destinationChainHandlerAddress common.Address, destinationRecipientAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.DepositGeneric(&_Bridge.TransactOpts, originChainContractAddress, originChainHandlerAddress, destinationChainID, destinationChainHandlerAddress, destinationRecipientAddress, data)
}

// DepositGeneric0 is a paid mutator transaction binding the contract method 0xb8260def.
//
// Solidity: function depositGeneric(uint256 destinationChainID, address destinationRecipientAddress, bytes data) returns()
func (_Bridge *BridgeTransactor) DepositGeneric0(opts *bind.TransactOpts, destinationChainID *big.Int, destinationRecipientAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositGeneric0", destinationChainID, destinationRecipientAddress, data)
}

// DepositGeneric0 is a paid mutator transaction binding the contract method 0xb8260def.
//
// Solidity: function depositGeneric(uint256 destinationChainID, address destinationRecipientAddress, bytes data) returns()
func (_Bridge *BridgeSession) DepositGeneric0(destinationChainID *big.Int, destinationRecipientAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.DepositGeneric0(&_Bridge.TransactOpts, destinationChainID, destinationRecipientAddress, data)
}

// DepositGeneric0 is a paid mutator transaction binding the contract method 0xb8260def.
//
// Solidity: function depositGeneric(uint256 destinationChainID, address destinationRecipientAddress, bytes data) returns()
func (_Bridge *BridgeTransactorSession) DepositGeneric0(destinationChainID *big.Int, destinationRecipientAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.DepositGeneric0(&_Bridge.TransactOpts, destinationChainID, destinationRecipientAddress, data)
}

// ExecuteDepositProposal is a paid mutator transaction binding the contract method 0x030e7ddd.
//
// Solidity: function executeDepositProposal(uint256 originChainID, uint256 depositID, address destinationChainHandlerAddress, bytes data) returns()
func (_Bridge *BridgeTransactor) ExecuteDepositProposal(opts *bind.TransactOpts, originChainID *big.Int, depositID *big.Int, destinationChainHandlerAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "executeDepositProposal", originChainID, depositID, destinationChainHandlerAddress, data)
}

// ExecuteDepositProposal is a paid mutator transaction binding the contract method 0x030e7ddd.
//
// Solidity: function executeDepositProposal(uint256 originChainID, uint256 depositID, address destinationChainHandlerAddress, bytes data) returns()
func (_Bridge *BridgeSession) ExecuteDepositProposal(originChainID *big.Int, depositID *big.Int, destinationChainHandlerAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteDepositProposal(&_Bridge.TransactOpts, originChainID, depositID, destinationChainHandlerAddress, data)
}

// ExecuteDepositProposal is a paid mutator transaction binding the contract method 0x030e7ddd.
//
// Solidity: function executeDepositProposal(uint256 originChainID, uint256 depositID, address destinationChainHandlerAddress, bytes data) returns()
func (_Bridge *BridgeTransactorSession) ExecuteDepositProposal(originChainID *big.Int, depositID *big.Int, destinationChainHandlerAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteDepositProposal(&_Bridge.TransactOpts, originChainID, depositID, destinationChainHandlerAddress, data)
}

// VoteDepositProposal is a paid mutator transaction binding the contract method 0xe7d3696c.
//
// Solidity: function voteDepositProposal(uint256 originChainID, uint256 depositID, uint8 vote) returns()
func (_Bridge *BridgeTransactor) VoteDepositProposal(opts *bind.TransactOpts, originChainID *big.Int, depositID *big.Int, vote uint8) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "voteDepositProposal", originChainID, depositID, vote)
}

// VoteDepositProposal is a paid mutator transaction binding the contract method 0xe7d3696c.
//
// Solidity: function voteDepositProposal(uint256 originChainID, uint256 depositID, uint8 vote) returns()
func (_Bridge *BridgeSession) VoteDepositProposal(originChainID *big.Int, depositID *big.Int, vote uint8) (*types.Transaction, error) {
	return _Bridge.Contract.VoteDepositProposal(&_Bridge.TransactOpts, originChainID, depositID, vote)
}

// VoteDepositProposal is a paid mutator transaction binding the contract method 0xe7d3696c.
//
// Solidity: function voteDepositProposal(uint256 originChainID, uint256 depositID, uint8 vote) returns()
func (_Bridge *BridgeTransactorSession) VoteDepositProposal(originChainID *big.Int, depositID *big.Int, vote uint8) (*types.Transaction, error) {
	return _Bridge.Contract.VoteDepositProposal(&_Bridge.TransactOpts, originChainID, depositID, vote)
}

// VoteValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x9bd61e6a.
//
// Solidity: function voteValidatorThresholdProposal(uint8 vote) returns()
func (_Bridge *BridgeTransactor) VoteValidatorThresholdProposal(opts *bind.TransactOpts, vote uint8) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "voteValidatorThresholdProposal", vote)
}

// VoteValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x9bd61e6a.
//
// Solidity: function voteValidatorThresholdProposal(uint8 vote) returns()
func (_Bridge *BridgeSession) VoteValidatorThresholdProposal(vote uint8) (*types.Transaction, error) {
	return _Bridge.Contract.VoteValidatorThresholdProposal(&_Bridge.TransactOpts, vote)
}

// VoteValidatorThresholdProposal is a paid mutator transaction binding the contract method 0x9bd61e6a.
//
// Solidity: function voteValidatorThresholdProposal(uint8 vote) returns()
func (_Bridge *BridgeTransactorSession) VoteValidatorThresholdProposal(vote uint8) (*types.Transaction, error) {
	return _Bridge.Contract.VoteValidatorThresholdProposal(&_Bridge.TransactOpts, vote)
}

// BridgeDepositProposalCreatedIterator is returned from FilterDepositProposalCreated and is used to iterate over the raw logs and unpacked data for DepositProposalCreated events raised by the Bridge contract.
type BridgeDepositProposalCreatedIterator struct {
	Event *BridgeDepositProposalCreated // Event containing the contract specifics and raw log

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
func (it *BridgeDepositProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositProposalCreated)
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
		it.Event = new(BridgeDepositProposalCreated)
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
func (it *BridgeDepositProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositProposalCreated represents a DepositProposalCreated event raised by the Bridge contract.
type BridgeDepositProposalCreated struct {
	OriginChainID *big.Int
	DepositID     *big.Int
	DataHash      [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDepositProposalCreated is a free log retrieval operation binding the contract event 0xd79daef3492e6769de74ee113f9383423c3e43830a24b1be81e49dceecc639fd.
//
// Solidity: event DepositProposalCreated(uint256 indexed originChainID, uint256 indexed depositID, bytes32 indexed dataHash)
func (_Bridge *BridgeFilterer) FilterDepositProposalCreated(opts *bind.FilterOpts, originChainID []*big.Int, depositID []*big.Int, dataHash [][32]byte) (*BridgeDepositProposalCreatedIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositIDRule []interface{}
	for _, depositIDItem := range depositID {
		depositIDRule = append(depositIDRule, depositIDItem)
	}
	var dataHashRule []interface{}
	for _, dataHashItem := range dataHash {
		dataHashRule = append(dataHashRule, dataHashItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositProposalCreated", originChainIDRule, depositIDRule, dataHashRule)
	if err != nil {
		return nil, err
	}
	return &BridgeDepositProposalCreatedIterator{contract: _Bridge.contract, event: "DepositProposalCreated", logs: logs, sub: sub}, nil
}

// WatchDepositProposalCreated is a free log subscription operation binding the contract event 0xd79daef3492e6769de74ee113f9383423c3e43830a24b1be81e49dceecc639fd.
//
// Solidity: event DepositProposalCreated(uint256 indexed originChainID, uint256 indexed depositID, bytes32 indexed dataHash)
func (_Bridge *BridgeFilterer) WatchDepositProposalCreated(opts *bind.WatchOpts, sink chan<- *BridgeDepositProposalCreated, originChainID []*big.Int, depositID []*big.Int, dataHash [][32]byte) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositIDRule []interface{}
	for _, depositIDItem := range depositID {
		depositIDRule = append(depositIDRule, depositIDItem)
	}
	var dataHashRule []interface{}
	for _, dataHashItem := range dataHash {
		dataHashRule = append(dataHashRule, dataHashItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositProposalCreated", originChainIDRule, depositIDRule, dataHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositProposalCreated)
				if err := _Bridge.contract.UnpackLog(event, "DepositProposalCreated", log); err != nil {
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

// ParseDepositProposalCreated is a log parse operation binding the contract event 0xd79daef3492e6769de74ee113f9383423c3e43830a24b1be81e49dceecc639fd.
//
// Solidity: event DepositProposalCreated(uint256 indexed originChainID, uint256 indexed depositID, bytes32 indexed dataHash)
func (_Bridge *BridgeFilterer) ParseDepositProposalCreated(log types.Log) (*BridgeDepositProposalCreated, error) {
	event := new(BridgeDepositProposalCreated)
	if err := _Bridge.contract.UnpackLog(event, "DepositProposalCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeDepositProposalFinalizedIterator is returned from FilterDepositProposalFinalized and is used to iterate over the raw logs and unpacked data for DepositProposalFinalized events raised by the Bridge contract.
type BridgeDepositProposalFinalizedIterator struct {
	Event *BridgeDepositProposalFinalized // Event containing the contract specifics and raw log

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
func (it *BridgeDepositProposalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositProposalFinalized)
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
		it.Event = new(BridgeDepositProposalFinalized)
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
func (it *BridgeDepositProposalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositProposalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositProposalFinalized represents a DepositProposalFinalized event raised by the Bridge contract.
type BridgeDepositProposalFinalized struct {
	OriginChainID *big.Int
	DepositID     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDepositProposalFinalized is a free log retrieval operation binding the contract event 0xbc724fbe62906cf888d9fef0688a894cdea9b5c36290addb9d815699d343a99d.
//
// Solidity: event DepositProposalFinalized(uint256 indexed originChainID, uint256 indexed depositID)
func (_Bridge *BridgeFilterer) FilterDepositProposalFinalized(opts *bind.FilterOpts, originChainID []*big.Int, depositID []*big.Int) (*BridgeDepositProposalFinalizedIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositIDRule []interface{}
	for _, depositIDItem := range depositID {
		depositIDRule = append(depositIDRule, depositIDItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositProposalFinalized", originChainIDRule, depositIDRule)
	if err != nil {
		return nil, err
	}
	return &BridgeDepositProposalFinalizedIterator{contract: _Bridge.contract, event: "DepositProposalFinalized", logs: logs, sub: sub}, nil
}

// WatchDepositProposalFinalized is a free log subscription operation binding the contract event 0xbc724fbe62906cf888d9fef0688a894cdea9b5c36290addb9d815699d343a99d.
//
// Solidity: event DepositProposalFinalized(uint256 indexed originChainID, uint256 indexed depositID)
func (_Bridge *BridgeFilterer) WatchDepositProposalFinalized(opts *bind.WatchOpts, sink chan<- *BridgeDepositProposalFinalized, originChainID []*big.Int, depositID []*big.Int) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositIDRule []interface{}
	for _, depositIDItem := range depositID {
		depositIDRule = append(depositIDRule, depositIDItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositProposalFinalized", originChainIDRule, depositIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositProposalFinalized)
				if err := _Bridge.contract.UnpackLog(event, "DepositProposalFinalized", log); err != nil {
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

// ParseDepositProposalFinalized is a log parse operation binding the contract event 0xbc724fbe62906cf888d9fef0688a894cdea9b5c36290addb9d815699d343a99d.
//
// Solidity: event DepositProposalFinalized(uint256 indexed originChainID, uint256 indexed depositID)
func (_Bridge *BridgeFilterer) ParseDepositProposalFinalized(log types.Log) (*BridgeDepositProposalFinalized, error) {
	event := new(BridgeDepositProposalFinalized)
	if err := _Bridge.contract.UnpackLog(event, "DepositProposalFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeDepositProposalVoteIterator is returned from FilterDepositProposalVote and is used to iterate over the raw logs and unpacked data for DepositProposalVote events raised by the Bridge contract.
type BridgeDepositProposalVoteIterator struct {
	Event *BridgeDepositProposalVote // Event containing the contract specifics and raw log

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
func (it *BridgeDepositProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositProposalVote)
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
		it.Event = new(BridgeDepositProposalVote)
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
func (it *BridgeDepositProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositProposalVote represents a DepositProposalVote event raised by the Bridge contract.
type BridgeDepositProposalVote struct {
	OriginChainID *big.Int
	DepositID     *big.Int
	Vote          uint8
	Status        uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDepositProposalVote is a free log retrieval operation binding the contract event 0x40fc0ff8dcb86fbd4d70d3e1237dbd6dab58a03c8a1a3af7b27a4f70b540a66b.
//
// Solidity: event DepositProposalVote(uint256 indexed originChainID, uint256 indexed depositID, uint8 indexed vote, uint8 status)
func (_Bridge *BridgeFilterer) FilterDepositProposalVote(opts *bind.FilterOpts, originChainID []*big.Int, depositID []*big.Int, vote []uint8) (*BridgeDepositProposalVoteIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositIDRule []interface{}
	for _, depositIDItem := range depositID {
		depositIDRule = append(depositIDRule, depositIDItem)
	}
	var voteRule []interface{}
	for _, voteItem := range vote {
		voteRule = append(voteRule, voteItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositProposalVote", originChainIDRule, depositIDRule, voteRule)
	if err != nil {
		return nil, err
	}
	return &BridgeDepositProposalVoteIterator{contract: _Bridge.contract, event: "DepositProposalVote", logs: logs, sub: sub}, nil
}

// WatchDepositProposalVote is a free log subscription operation binding the contract event 0x40fc0ff8dcb86fbd4d70d3e1237dbd6dab58a03c8a1a3af7b27a4f70b540a66b.
//
// Solidity: event DepositProposalVote(uint256 indexed originChainID, uint256 indexed depositID, uint8 indexed vote, uint8 status)
func (_Bridge *BridgeFilterer) WatchDepositProposalVote(opts *bind.WatchOpts, sink chan<- *BridgeDepositProposalVote, originChainID []*big.Int, depositID []*big.Int, vote []uint8) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositIDRule []interface{}
	for _, depositIDItem := range depositID {
		depositIDRule = append(depositIDRule, depositIDItem)
	}
	var voteRule []interface{}
	for _, voteItem := range vote {
		voteRule = append(voteRule, voteItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositProposalVote", originChainIDRule, depositIDRule, voteRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositProposalVote)
				if err := _Bridge.contract.UnpackLog(event, "DepositProposalVote", log); err != nil {
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

// ParseDepositProposalVote is a log parse operation binding the contract event 0x40fc0ff8dcb86fbd4d70d3e1237dbd6dab58a03c8a1a3af7b27a4f70b540a66b.
//
// Solidity: event DepositProposalVote(uint256 indexed originChainID, uint256 indexed depositID, uint8 indexed vote, uint8 status)
func (_Bridge *BridgeFilterer) ParseDepositProposalVote(log types.Log) (*BridgeDepositProposalVote, error) {
	event := new(BridgeDepositProposalVote)
	if err := _Bridge.contract.UnpackLog(event, "DepositProposalVote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeERC20DepositedIterator is returned from FilterERC20Deposited and is used to iterate over the raw logs and unpacked data for ERC20Deposited events raised by the Bridge contract.
type BridgeERC20DepositedIterator struct {
	Event *BridgeERC20Deposited // Event containing the contract specifics and raw log

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
func (it *BridgeERC20DepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20Deposited)
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
		it.Event = new(BridgeERC20Deposited)
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
func (it *BridgeERC20DepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20DepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20Deposited represents a ERC20Deposited event raised by the Bridge contract.
type BridgeERC20Deposited struct {
	DepositID *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20Deposited is a free log retrieval operation binding the contract event 0x2199cfceaff95f329a0db94bd11b97a8f85f9e983679d755198045fd67851e04.
//
// Solidity: event ERC20Deposited(uint256 indexed depositID)
func (_Bridge *BridgeFilterer) FilterERC20Deposited(opts *bind.FilterOpts, depositID []*big.Int) (*BridgeERC20DepositedIterator, error) {

	var depositIDRule []interface{}
	for _, depositIDItem := range depositID {
		depositIDRule = append(depositIDRule, depositIDItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ERC20Deposited", depositIDRule)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20DepositedIterator{contract: _Bridge.contract, event: "ERC20Deposited", logs: logs, sub: sub}, nil
}

// WatchERC20Deposited is a free log subscription operation binding the contract event 0x2199cfceaff95f329a0db94bd11b97a8f85f9e983679d755198045fd67851e04.
//
// Solidity: event ERC20Deposited(uint256 indexed depositID)
func (_Bridge *BridgeFilterer) WatchERC20Deposited(opts *bind.WatchOpts, sink chan<- *BridgeERC20Deposited, depositID []*big.Int) (event.Subscription, error) {

	var depositIDRule []interface{}
	for _, depositIDItem := range depositID {
		depositIDRule = append(depositIDRule, depositIDItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ERC20Deposited", depositIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20Deposited)
				if err := _Bridge.contract.UnpackLog(event, "ERC20Deposited", log); err != nil {
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

// ParseERC20Deposited is a log parse operation binding the contract event 0x2199cfceaff95f329a0db94bd11b97a8f85f9e983679d755198045fd67851e04.
//
// Solidity: event ERC20Deposited(uint256 indexed depositID)
func (_Bridge *BridgeFilterer) ParseERC20Deposited(log types.Log) (*BridgeERC20Deposited, error) {
	event := new(BridgeERC20Deposited)
	if err := _Bridge.contract.UnpackLog(event, "ERC20Deposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeERC721DepositedIterator is returned from FilterERC721Deposited and is used to iterate over the raw logs and unpacked data for ERC721Deposited events raised by the Bridge contract.
type BridgeERC721DepositedIterator struct {
	Event *BridgeERC721Deposited // Event containing the contract specifics and raw log

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
func (it *BridgeERC721DepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC721Deposited)
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
		it.Event = new(BridgeERC721Deposited)
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
func (it *BridgeERC721DepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC721DepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC721Deposited represents a ERC721Deposited event raised by the Bridge contract.
type BridgeERC721Deposited struct {
	DepositID *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC721Deposited is a free log retrieval operation binding the contract event 0x14f91f5d57b330e0bf94ee405aac47227ebc44b43634f4f0bf64052913dcfed2.
//
// Solidity: event ERC721Deposited(uint256 indexed depositID)
func (_Bridge *BridgeFilterer) FilterERC721Deposited(opts *bind.FilterOpts, depositID []*big.Int) (*BridgeERC721DepositedIterator, error) {

	var depositIDRule []interface{}
	for _, depositIDItem := range depositID {
		depositIDRule = append(depositIDRule, depositIDItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ERC721Deposited", depositIDRule)
	if err != nil {
		return nil, err
	}
	return &BridgeERC721DepositedIterator{contract: _Bridge.contract, event: "ERC721Deposited", logs: logs, sub: sub}, nil
}

// WatchERC721Deposited is a free log subscription operation binding the contract event 0x14f91f5d57b330e0bf94ee405aac47227ebc44b43634f4f0bf64052913dcfed2.
//
// Solidity: event ERC721Deposited(uint256 indexed depositID)
func (_Bridge *BridgeFilterer) WatchERC721Deposited(opts *bind.WatchOpts, sink chan<- *BridgeERC721Deposited, depositID []*big.Int) (event.Subscription, error) {

	var depositIDRule []interface{}
	for _, depositIDItem := range depositID {
		depositIDRule = append(depositIDRule, depositIDItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ERC721Deposited", depositIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC721Deposited)
				if err := _Bridge.contract.UnpackLog(event, "ERC721Deposited", log); err != nil {
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

// ParseERC721Deposited is a log parse operation binding the contract event 0x14f91f5d57b330e0bf94ee405aac47227ebc44b43634f4f0bf64052913dcfed2.
//
// Solidity: event ERC721Deposited(uint256 indexed depositID)
func (_Bridge *BridgeFilterer) ParseERC721Deposited(log types.Log) (*BridgeERC721Deposited, error) {
	event := new(BridgeERC721Deposited)
	if err := _Bridge.contract.UnpackLog(event, "ERC721Deposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeGenericDepositedIterator is returned from FilterGenericDeposited and is used to iterate over the raw logs and unpacked data for GenericDeposited events raised by the Bridge contract.
type BridgeGenericDepositedIterator struct {
	Event *BridgeGenericDeposited // Event containing the contract specifics and raw log

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
func (it *BridgeGenericDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeGenericDeposited)
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
		it.Event = new(BridgeGenericDeposited)
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
func (it *BridgeGenericDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeGenericDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeGenericDeposited represents a GenericDeposited event raised by the Bridge contract.
type BridgeGenericDeposited struct {
	DepositID *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGenericDeposited is a free log retrieval operation binding the contract event 0x5f7c049578fd3ad850c1b539b94320934c5afc6dc53f4eab709dc45ee4cabe81.
//
// Solidity: event GenericDeposited(uint256 indexed depositID)
func (_Bridge *BridgeFilterer) FilterGenericDeposited(opts *bind.FilterOpts, depositID []*big.Int) (*BridgeGenericDepositedIterator, error) {

	var depositIDRule []interface{}
	for _, depositIDItem := range depositID {
		depositIDRule = append(depositIDRule, depositIDItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "GenericDeposited", depositIDRule)
	if err != nil {
		return nil, err
	}
	return &BridgeGenericDepositedIterator{contract: _Bridge.contract, event: "GenericDeposited", logs: logs, sub: sub}, nil
}

// WatchGenericDeposited is a free log subscription operation binding the contract event 0x5f7c049578fd3ad850c1b539b94320934c5afc6dc53f4eab709dc45ee4cabe81.
//
// Solidity: event GenericDeposited(uint256 indexed depositID)
func (_Bridge *BridgeFilterer) WatchGenericDeposited(opts *bind.WatchOpts, sink chan<- *BridgeGenericDeposited, depositID []*big.Int) (event.Subscription, error) {

	var depositIDRule []interface{}
	for _, depositIDItem := range depositID {
		depositIDRule = append(depositIDRule, depositIDItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "GenericDeposited", depositIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeGenericDeposited)
				if err := _Bridge.contract.UnpackLog(event, "GenericDeposited", log); err != nil {
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

// ParseGenericDeposited is a log parse operation binding the contract event 0x5f7c049578fd3ad850c1b539b94320934c5afc6dc53f4eab709dc45ee4cabe81.
//
// Solidity: event GenericDeposited(uint256 indexed depositID)
func (_Bridge *BridgeFilterer) ParseGenericDeposited(log types.Log) (*BridgeGenericDeposited, error) {
	event := new(BridgeGenericDeposited)
	if err := _Bridge.contract.UnpackLog(event, "GenericDeposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeValidatorThresholdChangedIterator is returned from FilterValidatorThresholdChanged and is used to iterate over the raw logs and unpacked data for ValidatorThresholdChanged events raised by the Bridge contract.
type BridgeValidatorThresholdChangedIterator struct {
	Event *BridgeValidatorThresholdChanged // Event containing the contract specifics and raw log

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
func (it *BridgeValidatorThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeValidatorThresholdChanged)
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
		it.Event = new(BridgeValidatorThresholdChanged)
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
func (it *BridgeValidatorThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeValidatorThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeValidatorThresholdChanged represents a ValidatorThresholdChanged event raised by the Bridge contract.
type BridgeValidatorThresholdChanged struct {
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorThresholdChanged is a free log retrieval operation binding the contract event 0x95c586411a87fb2ca28c34b78bce3bf57d0c29769a83ae46b484bd7fc049e8ee.
//
// Solidity: event ValidatorThresholdChanged(uint256 indexed newThreshold)
func (_Bridge *BridgeFilterer) FilterValidatorThresholdChanged(opts *bind.FilterOpts, newThreshold []*big.Int) (*BridgeValidatorThresholdChangedIterator, error) {

	var newThresholdRule []interface{}
	for _, newThresholdItem := range newThreshold {
		newThresholdRule = append(newThresholdRule, newThresholdItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ValidatorThresholdChanged", newThresholdRule)
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorThresholdChangedIterator{contract: _Bridge.contract, event: "ValidatorThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorThresholdChanged is a free log subscription operation binding the contract event 0x95c586411a87fb2ca28c34b78bce3bf57d0c29769a83ae46b484bd7fc049e8ee.
//
// Solidity: event ValidatorThresholdChanged(uint256 indexed newThreshold)
func (_Bridge *BridgeFilterer) WatchValidatorThresholdChanged(opts *bind.WatchOpts, sink chan<- *BridgeValidatorThresholdChanged, newThreshold []*big.Int) (event.Subscription, error) {

	var newThresholdRule []interface{}
	for _, newThresholdItem := range newThreshold {
		newThresholdRule = append(newThresholdRule, newThresholdItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ValidatorThresholdChanged", newThresholdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeValidatorThresholdChanged)
				if err := _Bridge.contract.UnpackLog(event, "ValidatorThresholdChanged", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseValidatorThresholdChanged(log types.Log) (*BridgeValidatorThresholdChanged, error) {
	event := new(BridgeValidatorThresholdChanged)
	if err := _Bridge.contract.UnpackLog(event, "ValidatorThresholdChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeValidatorThresholdProposalCreatedIterator is returned from FilterValidatorThresholdProposalCreated and is used to iterate over the raw logs and unpacked data for ValidatorThresholdProposalCreated events raised by the Bridge contract.
type BridgeValidatorThresholdProposalCreatedIterator struct {
	Event *BridgeValidatorThresholdProposalCreated // Event containing the contract specifics and raw log

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
func (it *BridgeValidatorThresholdProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeValidatorThresholdProposalCreated)
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
		it.Event = new(BridgeValidatorThresholdProposalCreated)
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
func (it *BridgeValidatorThresholdProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeValidatorThresholdProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeValidatorThresholdProposalCreated represents a ValidatorThresholdProposalCreated event raised by the Bridge contract.
type BridgeValidatorThresholdProposalCreated struct {
	ProposedValue *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorThresholdProposalCreated is a free log retrieval operation binding the contract event 0xc7876b43f0eccd5339779e2a0ddc65954cea640a0bc330145ead56fd205f1076.
//
// Solidity: event ValidatorThresholdProposalCreated(uint256 indexed proposedValue)
func (_Bridge *BridgeFilterer) FilterValidatorThresholdProposalCreated(opts *bind.FilterOpts, proposedValue []*big.Int) (*BridgeValidatorThresholdProposalCreatedIterator, error) {

	var proposedValueRule []interface{}
	for _, proposedValueItem := range proposedValue {
		proposedValueRule = append(proposedValueRule, proposedValueItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ValidatorThresholdProposalCreated", proposedValueRule)
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorThresholdProposalCreatedIterator{contract: _Bridge.contract, event: "ValidatorThresholdProposalCreated", logs: logs, sub: sub}, nil
}

// WatchValidatorThresholdProposalCreated is a free log subscription operation binding the contract event 0xc7876b43f0eccd5339779e2a0ddc65954cea640a0bc330145ead56fd205f1076.
//
// Solidity: event ValidatorThresholdProposalCreated(uint256 indexed proposedValue)
func (_Bridge *BridgeFilterer) WatchValidatorThresholdProposalCreated(opts *bind.WatchOpts, sink chan<- *BridgeValidatorThresholdProposalCreated, proposedValue []*big.Int) (event.Subscription, error) {

	var proposedValueRule []interface{}
	for _, proposedValueItem := range proposedValue {
		proposedValueRule = append(proposedValueRule, proposedValueItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ValidatorThresholdProposalCreated", proposedValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeValidatorThresholdProposalCreated)
				if err := _Bridge.contract.UnpackLog(event, "ValidatorThresholdProposalCreated", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseValidatorThresholdProposalCreated(log types.Log) (*BridgeValidatorThresholdProposalCreated, error) {
	event := new(BridgeValidatorThresholdProposalCreated)
	if err := _Bridge.contract.UnpackLog(event, "ValidatorThresholdProposalCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeValidatorThresholdProposalVoteIterator is returned from FilterValidatorThresholdProposalVote and is used to iterate over the raw logs and unpacked data for ValidatorThresholdProposalVote events raised by the Bridge contract.
type BridgeValidatorThresholdProposalVoteIterator struct {
	Event *BridgeValidatorThresholdProposalVote // Event containing the contract specifics and raw log

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
func (it *BridgeValidatorThresholdProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeValidatorThresholdProposalVote)
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
		it.Event = new(BridgeValidatorThresholdProposalVote)
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
func (it *BridgeValidatorThresholdProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeValidatorThresholdProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeValidatorThresholdProposalVote represents a ValidatorThresholdProposalVote event raised by the Bridge contract.
type BridgeValidatorThresholdProposalVote struct {
	Vote uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorThresholdProposalVote is a free log retrieval operation binding the contract event 0x1c5e1f97eaeb3bcba56e89ed4597ed8074e3bf825c6a65528f5e27a621139abe.
//
// Solidity: event ValidatorThresholdProposalVote(uint8 vote)
func (_Bridge *BridgeFilterer) FilterValidatorThresholdProposalVote(opts *bind.FilterOpts) (*BridgeValidatorThresholdProposalVoteIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ValidatorThresholdProposalVote")
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorThresholdProposalVoteIterator{contract: _Bridge.contract, event: "ValidatorThresholdProposalVote", logs: logs, sub: sub}, nil
}

// WatchValidatorThresholdProposalVote is a free log subscription operation binding the contract event 0x1c5e1f97eaeb3bcba56e89ed4597ed8074e3bf825c6a65528f5e27a621139abe.
//
// Solidity: event ValidatorThresholdProposalVote(uint8 vote)
func (_Bridge *BridgeFilterer) WatchValidatorThresholdProposalVote(opts *bind.WatchOpts, sink chan<- *BridgeValidatorThresholdProposalVote) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ValidatorThresholdProposalVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeValidatorThresholdProposalVote)
				if err := _Bridge.contract.UnpackLog(event, "ValidatorThresholdProposalVote", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseValidatorThresholdProposalVote(log types.Log) (*BridgeValidatorThresholdProposalVote, error) {
	event := new(BridgeValidatorThresholdProposalVote)
	if err := _Bridge.contract.UnpackLog(event, "ValidatorThresholdProposalVote", log); err != nil {
		return nil, err
	}
	return event, nil
}
