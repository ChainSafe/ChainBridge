// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC721Mintable

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

// ERC721MintableABI is the input ABI used to generate the binding from.
const ERC721MintableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721MintableBin is the compiled bytecode used for deploying new contracts.
var ERC721MintableBin = "0x60806040526200001c6301ffc9a760e01b6200005a60201b60201c565b620000346380ac58cd60e01b6200005a60201b60201c565b62000054620000486200016360201b60201c565b6200016b60201b60201c565b62000390565b63ffffffff60e01b817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161415620000f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4552433136353a20696e76616c696420696e746572666163652069640000000081525060200191505060405180910390fd5b6001600080837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600033905090565b62000186816005620001cc60201b620020671790919060201c565b8073ffffffffffffffffffffffffffffffffffffffff167f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f660405160405180910390a250565b620001de8282620002b060201b60201c565b1562000252576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000339576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180620029a66022913960400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b61260680620003a06000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806370a08231116100a2578063a144819411610071578063a1448194146105b2578063a22cb46514610618578063aa271e1a14610668578063b88d4fde146106c4578063e985e9c5146107c95761010b565b806370a082311461040f5780638832e6e314610467578063983b2d561461056457806398650275146105a85761010b565b806340c10f19116100de57806340c10f191461029f57806342842e0e1461030557806342966c68146103735780636352211e146103a15761010b565b806301ffc9a714610110578063081812fc14610175578063095ea7b3146101e357806323b872dd14610231575b600080fd5b61015b6004803603602081101561012657600080fd5b8101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190505050610845565b604051808215151515815260200191505060405180910390f35b6101a16004803603602081101561018b57600080fd5b81019080803590602001909291905050506108ac565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61022f600480360360408110156101f957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610947565b005b61029d6004803603606081101561024757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a8b565b005b6102eb600480360360408110156102b557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b01565b604051808215151515815260200191505060405180910390f35b6103716004803603606081101561031b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b7c565b005b61039f6004803603602081101561038957600080fd5b8101908080359060200190929190505050610b9c565b005b6103cd600480360360208110156103b757600080fd5b8101908080359060200190929190505050610c07565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104516004803603602081101561042557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ccf565b6040518082815260200191505060405180910390f35b61054a6004803603606081101561047d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156104c457600080fd5b8201836020820111156104d657600080fd5b803590602001918460018302840111640100000000831117156104f857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610da4565b604051808215151515815260200191505060405180910390f35b6105a66004803603602081101561057a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e21565b005b6105b0610e92565b005b6105fe600480360360408110156105c857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ea4565b604051808215151515815260200191505060405180910390f35b6106666004803603604081101561062e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803515159060200190929190505050610f1f565b005b6106aa6004803603602081101561067e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110d7565b604051808215151515815260200191505060405180910390f35b6107c7600480360360808110156106da57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561074157600080fd5b82018360208201111561075357600080fd5b8035906020019184600183028401116401000000008311171561077557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506110f4565b005b61082b600480360360408110156107df57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061116c565b604051808215151515815260200191505060405180910390f35b6000806000837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060009054906101000a900460ff169050919050565b60006108b782611200565b61090c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c8152602001806124b3602c913960400191505060405180910390fd5b6002600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600061095282610c07565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156109d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018061252a6021913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166109f8611272565b73ffffffffffffffffffffffffffffffffffffffff161480610a275750610a2681610a21611272565b61116c565b5b610a7c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260388152602001806123d76038913960400191505060405180910390fd5b610a86838361127a565b505050565b610a9c610a96611272565b82611333565b610af1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603181526020018061254b6031913960400191505060405180910390fd5b610afc838383611427565b505050565b6000610b13610b0e611272565b6110d7565b610b68576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806124626030913960400191505060405180910390fd5b610b72838361168f565b6001905092915050565b610b97838383604051806020016040528060008152506110f4565b505050565b610ba63382611333565b610bfb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806125a16030913960400191505060405180910390fd5b610c04816118b3565b50565b6000806001600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610cc6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806124396029913960400191505060405180910390fd5b80915050919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610d56576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a81526020018061240f602a913960400191505060405180910390fd5b610d9d600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206118c8565b9050919050565b6000610db6610db1611272565b6110d7565b610e0b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806124626030913960400191505060405180910390fd5b610e168484846118d6565b600190509392505050565b610e31610e2c611272565b6110d7565b610e86576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806124626030913960400191505060405180910390fd5b610e8f81611947565b50565b610ea2610e9d611272565b6119a1565b565b6000610eb6610eb1611272565b6110d7565b610f0b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806124626030913960400191505060405180910390fd5b610f1583836119fb565b6001905092915050565b610f27611272565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610fc8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4552433732313a20617070726f766520746f2063616c6c65720000000000000081525060200191505060405180910390fd5b8060046000610fd5611272565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff16611082611272565b73ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051808215151515815260200191505060405180910390a35050565b60006110ed826005611a1990919063ffffffff16565b9050919050565b6111056110ff611272565b83611333565b61115a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603181526020018061254b6031913960400191505060405180910390fd5b61116684848484611af7565b50505050565b6000600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000806001600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415915050919050565b600033905090565b816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff166112ed83610c07565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600061133e82611200565b611393576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c8152602001806123ab602c913960400191505060405180910390fd5b600061139e83610c07565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16148061140d57508373ffffffffffffffffffffffffffffffffffffffff166113f5846108ac565b73ffffffffffffffffffffffffffffffffffffffff16145b8061141e575061141d818561116c565b5b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff1661144782610c07565b73ffffffffffffffffffffffffffffffffffffffff16146114b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806125016029913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611539576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806123876024913960400191505060405180910390fd5b611544838383611b69565b61154f60008261127a565b611596600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020611b6e565b6115dd600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020611b91565b816001600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611732576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4552433732313a206d696e7420746f20746865207a65726f206164647265737381525060200191505060405180910390fd5b61173b81611200565b156117ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000081525060200191505060405180910390fd5b6117ba60008383611b69565b816001600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611853600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020611b91565b808273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a45050565b6118c56118bf82610c07565b82611ba7565b50565b600081600001549050919050565b6118e0838361168f565b6118ed6000848484611d44565b611942576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001806123556032913960400191505060405180910390fd5b505050565b61195b81600561206790919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f660405160405180910390a250565b6119b581600561214290919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669260405160405180910390a250565b611a158282604051806020016040528060008152506118d6565b5050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611aa0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806124df6022913960400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b611b02848484611427565b611b0e84848484611d44565b611b63576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001806123556032913960400191505060405180910390fd5b50505050565b505050565b611b86600182600001546121ff90919063ffffffff16565b816000018190555050565b6001816000016000828254019250508190555050565b8173ffffffffffffffffffffffffffffffffffffffff16611bc782610c07565b73ffffffffffffffffffffffffffffffffffffffff1614611c33576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061257c6025913960400191505060405180910390fd5b611c3f82600083611b69565b611c4a60008261127a565b611c91600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020611b6e565b60006001600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a45050565b6000611d658473ffffffffffffffffffffffffffffffffffffffff16612249565b611d72576001905061205f565b600060608573ffffffffffffffffffffffffffffffffffffffff1663150b7a0260e01b611d9d611272565b898888604051602401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611e4d578082015181840152602081019050611e32565b50505050905090810190601f168015611e7a5780820380516001836020036101000a031916815260200191505b5095505050505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040518082805190602001908083835b60208310611f125780518252602082019150602081019050602083039250611eef565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611f74576040519150601f19603f3d011682016040523d82523d6000602084013e611f79565b606091505b509150915081611fe757600081511115611f965780518082602001fd5b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001806123556032913960400191505060405180910390fd5b6000818060200190516020811015611ffe57600080fd5b8101908080519060200190929190505050905063150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161493505050505b949350505050565b6120718282611a19565b156120e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b61214c8282611a19565b6121a1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806124926021913960400191505060405180910390fd5b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600061224183836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250612294565b905092915050565b60008060007fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060001b9050833f915080821415801561228b57506000801b8214155b92505050919050565b6000838311158290612341576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156123065780820151818401526020810190506122eb565b50505050905090810190601f1680156123335780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838503905080915050939250505056fe4552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f20616464726573734552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c654552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e526f6c65733a206163636f756e7420697320746865207a65726f20616464726573734552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65724552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644552433732313a206275726e206f6620746f6b656e2074686174206973206e6f74206f776e4552433732314275726e61626c653a2063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564a2646970667358221220b4bae99e380e9266072820bb45ba1f6cd90a7eba70bea96528e6c8b4d8a6abcc64736f6c63430006040033526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373"

// DeployERC721Mintable deploys a new Ethereum contract, binding an instance of ERC721Mintable to it.
func DeployERC721Mintable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Mintable, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721MintableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721MintableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Mintable{ERC721MintableCaller: ERC721MintableCaller{contract: contract}, ERC721MintableTransactor: ERC721MintableTransactor{contract: contract}, ERC721MintableFilterer: ERC721MintableFilterer{contract: contract}}, nil
}

// ERC721Mintable is an auto generated Go binding around an Ethereum contract.
type ERC721Mintable struct {
	ERC721MintableCaller     // Read-only binding to the contract
	ERC721MintableTransactor // Write-only binding to the contract
	ERC721MintableFilterer   // Log filterer for contract events
}

// ERC721MintableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721MintableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MintableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721MintableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MintableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721MintableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MintableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721MintableSession struct {
	Contract     *ERC721Mintable   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721MintableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721MintableCallerSession struct {
	Contract *ERC721MintableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC721MintableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721MintableTransactorSession struct {
	Contract     *ERC721MintableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC721MintableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721MintableRaw struct {
	Contract *ERC721Mintable // Generic contract binding to access the raw methods on
}

// ERC721MintableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721MintableCallerRaw struct {
	Contract *ERC721MintableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721MintableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721MintableTransactorRaw struct {
	Contract *ERC721MintableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Mintable creates a new instance of ERC721Mintable, bound to a specific deployed contract.
func NewERC721Mintable(address common.Address, backend bind.ContractBackend) (*ERC721Mintable, error) {
	contract, err := bindERC721Mintable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Mintable{ERC721MintableCaller: ERC721MintableCaller{contract: contract}, ERC721MintableTransactor: ERC721MintableTransactor{contract: contract}, ERC721MintableFilterer: ERC721MintableFilterer{contract: contract}}, nil
}

// NewERC721MintableCaller creates a new read-only instance of ERC721Mintable, bound to a specific deployed contract.
func NewERC721MintableCaller(address common.Address, caller bind.ContractCaller) (*ERC721MintableCaller, error) {
	contract, err := bindERC721Mintable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MintableCaller{contract: contract}, nil
}

// NewERC721MintableTransactor creates a new write-only instance of ERC721Mintable, bound to a specific deployed contract.
func NewERC721MintableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721MintableTransactor, error) {
	contract, err := bindERC721Mintable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MintableTransactor{contract: contract}, nil
}

// NewERC721MintableFilterer creates a new log filterer instance of ERC721Mintable, bound to a specific deployed contract.
func NewERC721MintableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721MintableFilterer, error) {
	contract, err := bindERC721Mintable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721MintableFilterer{contract: contract}, nil
}

// bindERC721Mintable binds a generic wrapper to an already deployed contract.
func bindERC721Mintable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721MintableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Mintable *ERC721MintableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Mintable.Contract.ERC721MintableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Mintable *ERC721MintableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.ERC721MintableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Mintable *ERC721MintableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.ERC721MintableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Mintable *ERC721MintableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Mintable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Mintable *ERC721MintableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Mintable *ERC721MintableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Mintable *ERC721MintableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Mintable.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Mintable *ERC721MintableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Mintable.Contract.BalanceOf(&_ERC721Mintable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Mintable *ERC721MintableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Mintable.Contract.BalanceOf(&_ERC721Mintable.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Mintable *ERC721MintableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Mintable.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Mintable *ERC721MintableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Mintable.Contract.GetApproved(&_ERC721Mintable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Mintable *ERC721MintableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Mintable.Contract.GetApproved(&_ERC721Mintable.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Mintable *ERC721MintableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Mintable.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Mintable *ERC721MintableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Mintable.Contract.IsApprovedForAll(&_ERC721Mintable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Mintable *ERC721MintableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Mintable.Contract.IsApprovedForAll(&_ERC721Mintable.CallOpts, owner, operator)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ERC721Mintable *ERC721MintableCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Mintable.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ERC721Mintable *ERC721MintableSession) IsMinter(account common.Address) (bool, error) {
	return _ERC721Mintable.Contract.IsMinter(&_ERC721Mintable.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ERC721Mintable *ERC721MintableCallerSession) IsMinter(account common.Address) (bool, error) {
	return _ERC721Mintable.Contract.IsMinter(&_ERC721Mintable.CallOpts, account)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Mintable *ERC721MintableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Mintable.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Mintable *ERC721MintableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Mintable.Contract.OwnerOf(&_ERC721Mintable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Mintable *ERC721MintableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Mintable.Contract.OwnerOf(&_ERC721Mintable.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Mintable *ERC721MintableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Mintable.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Mintable *ERC721MintableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Mintable.Contract.SupportsInterface(&_ERC721Mintable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Mintable *ERC721MintableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Mintable.Contract.SupportsInterface(&_ERC721Mintable.CallOpts, interfaceId)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ERC721Mintable *ERC721MintableTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ERC721Mintable.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ERC721Mintable *ERC721MintableSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.AddMinter(&_ERC721Mintable.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ERC721Mintable *ERC721MintableTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.AddMinter(&_ERC721Mintable.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Mintable *ERC721MintableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Mintable *ERC721MintableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.Approve(&_ERC721Mintable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Mintable *ERC721MintableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.Approve(&_ERC721Mintable.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ERC721Mintable *ERC721MintableTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ERC721Mintable *ERC721MintableSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.Burn(&_ERC721Mintable.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ERC721Mintable *ERC721MintableTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.Burn(&_ERC721Mintable.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns(bool)
func (_ERC721Mintable *ERC721MintableTransactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.contract.Transact(opts, "mint", to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns(bool)
func (_ERC721Mintable *ERC721MintableSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.Mint(&_ERC721Mintable.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns(bool)
func (_ERC721Mintable *ERC721MintableTransactorSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.Mint(&_ERC721Mintable.TransactOpts, to, tokenId)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ERC721Mintable *ERC721MintableTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Mintable.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ERC721Mintable *ERC721MintableSession) RenounceMinter() (*types.Transaction, error) {
	return _ERC721Mintable.Contract.RenounceMinter(&_ERC721Mintable.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ERC721Mintable *ERC721MintableTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _ERC721Mintable.Contract.RenounceMinter(&_ERC721Mintable.TransactOpts)
}

// SafeMint is a paid mutator transaction binding the contract method 0x8832e6e3.
//
// Solidity: function safeMint(address to, uint256 tokenId, bytes _data) returns(bool)
func (_ERC721Mintable *ERC721MintableTransactor) SafeMint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Mintable.contract.Transact(opts, "safeMint", to, tokenId, _data)
}

// SafeMint is a paid mutator transaction binding the contract method 0x8832e6e3.
//
// Solidity: function safeMint(address to, uint256 tokenId, bytes _data) returns(bool)
func (_ERC721Mintable *ERC721MintableSession) SafeMint(to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.SafeMint(&_ERC721Mintable.TransactOpts, to, tokenId, _data)
}

// SafeMint is a paid mutator transaction binding the contract method 0x8832e6e3.
//
// Solidity: function safeMint(address to, uint256 tokenId, bytes _data) returns(bool)
func (_ERC721Mintable *ERC721MintableTransactorSession) SafeMint(to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.SafeMint(&_ERC721Mintable.TransactOpts, to, tokenId, _data)
}

// SafeMint0 is a paid mutator transaction binding the contract method 0xa1448194.
//
// Solidity: function safeMint(address to, uint256 tokenId) returns(bool)
func (_ERC721Mintable *ERC721MintableTransactor) SafeMint0(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.contract.Transact(opts, "safeMint0", to, tokenId)
}

// SafeMint0 is a paid mutator transaction binding the contract method 0xa1448194.
//
// Solidity: function safeMint(address to, uint256 tokenId) returns(bool)
func (_ERC721Mintable *ERC721MintableSession) SafeMint0(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.SafeMint0(&_ERC721Mintable.TransactOpts, to, tokenId)
}

// SafeMint0 is a paid mutator transaction binding the contract method 0xa1448194.
//
// Solidity: function safeMint(address to, uint256 tokenId) returns(bool)
func (_ERC721Mintable *ERC721MintableTransactorSession) SafeMint0(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.SafeMint0(&_ERC721Mintable.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Mintable *ERC721MintableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Mintable *ERC721MintableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.SafeTransferFrom(&_ERC721Mintable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Mintable *ERC721MintableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.SafeTransferFrom(&_ERC721Mintable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Mintable *ERC721MintableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Mintable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Mintable *ERC721MintableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.SafeTransferFrom0(&_ERC721Mintable.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Mintable *ERC721MintableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.SafeTransferFrom0(&_ERC721Mintable.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Mintable *ERC721MintableTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Mintable.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Mintable *ERC721MintableSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.SetApprovalForAll(&_ERC721Mintable.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Mintable *ERC721MintableTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.SetApprovalForAll(&_ERC721Mintable.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Mintable *ERC721MintableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Mintable *ERC721MintableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.TransferFrom(&_ERC721Mintable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Mintable *ERC721MintableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mintable.Contract.TransferFrom(&_ERC721Mintable.TransactOpts, from, to, tokenId)
}

// ERC721MintableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Mintable contract.
type ERC721MintableApprovalIterator struct {
	Event *ERC721MintableApproval // Event containing the contract specifics and raw log

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
func (it *ERC721MintableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MintableApproval)
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
		it.Event = new(ERC721MintableApproval)
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
func (it *ERC721MintableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MintableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MintableApproval represents a Approval event raised by the ERC721Mintable contract.
type ERC721MintableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Mintable *ERC721MintableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721MintableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Mintable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MintableApprovalIterator{contract: _ERC721Mintable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Mintable *ERC721MintableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721MintableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Mintable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MintableApproval)
				if err := _ERC721Mintable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Mintable *ERC721MintableFilterer) ParseApproval(log types.Log) (*ERC721MintableApproval, error) {
	event := new(ERC721MintableApproval)
	if err := _ERC721Mintable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721MintableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Mintable contract.
type ERC721MintableApprovalForAllIterator struct {
	Event *ERC721MintableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721MintableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MintableApprovalForAll)
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
		it.Event = new(ERC721MintableApprovalForAll)
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
func (it *ERC721MintableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MintableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MintableApprovalForAll represents a ApprovalForAll event raised by the ERC721Mintable contract.
type ERC721MintableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Mintable *ERC721MintableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721MintableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Mintable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MintableApprovalForAllIterator{contract: _ERC721Mintable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Mintable *ERC721MintableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721MintableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Mintable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MintableApprovalForAll)
				if err := _ERC721Mintable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Mintable *ERC721MintableFilterer) ParseApprovalForAll(log types.Log) (*ERC721MintableApprovalForAll, error) {
	event := new(ERC721MintableApprovalForAll)
	if err := _ERC721Mintable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721MintableMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the ERC721Mintable contract.
type ERC721MintableMinterAddedIterator struct {
	Event *ERC721MintableMinterAdded // Event containing the contract specifics and raw log

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
func (it *ERC721MintableMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MintableMinterAdded)
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
		it.Event = new(ERC721MintableMinterAdded)
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
func (it *ERC721MintableMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MintableMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MintableMinterAdded represents a MinterAdded event raised by the ERC721Mintable contract.
type ERC721MintableMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ERC721Mintable *ERC721MintableFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*ERC721MintableMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC721Mintable.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MintableMinterAddedIterator{contract: _ERC721Mintable.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ERC721Mintable *ERC721MintableFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *ERC721MintableMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC721Mintable.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MintableMinterAdded)
				if err := _ERC721Mintable.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ERC721Mintable *ERC721MintableFilterer) ParseMinterAdded(log types.Log) (*ERC721MintableMinterAdded, error) {
	event := new(ERC721MintableMinterAdded)
	if err := _ERC721Mintable.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721MintableMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the ERC721Mintable contract.
type ERC721MintableMinterRemovedIterator struct {
	Event *ERC721MintableMinterRemoved // Event containing the contract specifics and raw log

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
func (it *ERC721MintableMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MintableMinterRemoved)
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
		it.Event = new(ERC721MintableMinterRemoved)
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
func (it *ERC721MintableMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MintableMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MintableMinterRemoved represents a MinterRemoved event raised by the ERC721Mintable contract.
type ERC721MintableMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ERC721Mintable *ERC721MintableFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*ERC721MintableMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC721Mintable.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MintableMinterRemovedIterator{contract: _ERC721Mintable.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ERC721Mintable *ERC721MintableFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *ERC721MintableMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC721Mintable.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MintableMinterRemoved)
				if err := _ERC721Mintable.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ERC721Mintable *ERC721MintableFilterer) ParseMinterRemoved(log types.Log) (*ERC721MintableMinterRemoved, error) {
	event := new(ERC721MintableMinterRemoved)
	if err := _ERC721Mintable.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721MintableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Mintable contract.
type ERC721MintableTransferIterator struct {
	Event *ERC721MintableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721MintableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MintableTransfer)
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
		it.Event = new(ERC721MintableTransfer)
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
func (it *ERC721MintableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MintableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MintableTransfer represents a Transfer event raised by the ERC721Mintable contract.
type ERC721MintableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Mintable *ERC721MintableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721MintableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Mintable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MintableTransferIterator{contract: _ERC721Mintable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Mintable *ERC721MintableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721MintableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Mintable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MintableTransfer)
				if err := _ERC721Mintable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Mintable *ERC721MintableFilterer) ParseTransfer(log types.Log) (*ERC721MintableTransfer, error) {
	event := new(ERC721MintableTransfer)
	if err := _ERC721Mintable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

var RuntimeBytecode = "0x608060405234801561001057600080fd5b506004361061010b5760003560e01c806370a08231116100a2578063a144819411610071578063a1448194146105b2578063a22cb46514610618578063aa271e1a14610668578063b88d4fde146106c4578063e985e9c5146107c95761010b565b806370a082311461040f5780638832e6e314610467578063983b2d561461056457806398650275146105a85761010b565b806340c10f19116100de57806340c10f191461029f57806342842e0e1461030557806342966c68146103735780636352211e146103a15761010b565b806301ffc9a714610110578063081812fc14610175578063095ea7b3146101e357806323b872dd14610231575b600080fd5b61015b6004803603602081101561012657600080fd5b8101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190505050610845565b604051808215151515815260200191505060405180910390f35b6101a16004803603602081101561018b57600080fd5b81019080803590602001909291905050506108ac565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61022f600480360360408110156101f957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610947565b005b61029d6004803603606081101561024757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a8b565b005b6102eb600480360360408110156102b557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b01565b604051808215151515815260200191505060405180910390f35b6103716004803603606081101561031b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b7c565b005b61039f6004803603602081101561038957600080fd5b8101908080359060200190929190505050610b9c565b005b6103cd600480360360208110156103b757600080fd5b8101908080359060200190929190505050610c07565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104516004803603602081101561042557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ccf565b6040518082815260200191505060405180910390f35b61054a6004803603606081101561047d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156104c457600080fd5b8201836020820111156104d657600080fd5b803590602001918460018302840111640100000000831117156104f857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610da4565b604051808215151515815260200191505060405180910390f35b6105a66004803603602081101561057a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e21565b005b6105b0610e92565b005b6105fe600480360360408110156105c857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ea4565b604051808215151515815260200191505060405180910390f35b6106666004803603604081101561062e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803515159060200190929190505050610f1f565b005b6106aa6004803603602081101561067e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110d7565b604051808215151515815260200191505060405180910390f35b6107c7600480360360808110156106da57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561074157600080fd5b82018360208201111561075357600080fd5b8035906020019184600183028401116401000000008311171561077557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506110f4565b005b61082b600480360360408110156107df57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061116c565b604051808215151515815260200191505060405180910390f35b6000806000837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060009054906101000a900460ff169050919050565b60006108b782611200565b61090c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c8152602001806124b3602c913960400191505060405180910390fd5b6002600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600061095282610c07565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156109d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018061252a6021913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166109f8611272565b73ffffffffffffffffffffffffffffffffffffffff161480610a275750610a2681610a21611272565b61116c565b5b610a7c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260388152602001806123d76038913960400191505060405180910390fd5b610a86838361127a565b505050565b610a9c610a96611272565b82611333565b610af1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603181526020018061254b6031913960400191505060405180910390fd5b610afc838383611427565b505050565b6000610b13610b0e611272565b6110d7565b610b68576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806124626030913960400191505060405180910390fd5b610b72838361168f565b6001905092915050565b610b97838383604051806020016040528060008152506110f4565b505050565b610ba63382611333565b610bfb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806125a16030913960400191505060405180910390fd5b610c04816118b3565b50565b6000806001600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610cc6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806124396029913960400191505060405180910390fd5b80915050919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610d56576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a81526020018061240f602a913960400191505060405180910390fd5b610d9d600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206118c8565b9050919050565b6000610db6610db1611272565b6110d7565b610e0b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806124626030913960400191505060405180910390fd5b610e168484846118d6565b600190509392505050565b610e31610e2c611272565b6110d7565b610e86576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806124626030913960400191505060405180910390fd5b610e8f81611947565b50565b610ea2610e9d611272565b6119a1565b565b6000610eb6610eb1611272565b6110d7565b610f0b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806124626030913960400191505060405180910390fd5b610f1583836119fb565b6001905092915050565b610f27611272565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610fc8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4552433732313a20617070726f766520746f2063616c6c65720000000000000081525060200191505060405180910390fd5b8060046000610fd5611272565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff16611082611272565b73ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051808215151515815260200191505060405180910390a35050565b60006110ed826005611a1990919063ffffffff16565b9050919050565b6111056110ff611272565b83611333565b61115a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603181526020018061254b6031913960400191505060405180910390fd5b61116684848484611af7565b50505050565b6000600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000806001600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415915050919050565b600033905090565b816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff166112ed83610c07565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600061133e82611200565b611393576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c8152602001806123ab602c913960400191505060405180910390fd5b600061139e83610c07565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16148061140d57508373ffffffffffffffffffffffffffffffffffffffff166113f5846108ac565b73ffffffffffffffffffffffffffffffffffffffff16145b8061141e575061141d818561116c565b5b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff1661144782610c07565b73ffffffffffffffffffffffffffffffffffffffff16146114b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806125016029913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611539576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806123876024913960400191505060405180910390fd5b611544838383611b69565b61154f60008261127a565b611596600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020611b6e565b6115dd600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020611b91565b816001600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611732576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4552433732313a206d696e7420746f20746865207a65726f206164647265737381525060200191505060405180910390fd5b61173b81611200565b156117ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000081525060200191505060405180910390fd5b6117ba60008383611b69565b816001600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611853600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020611b91565b808273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a45050565b6118c56118bf82610c07565b82611ba7565b50565b600081600001549050919050565b6118e0838361168f565b6118ed6000848484611d44565b611942576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001806123556032913960400191505060405180910390fd5b505050565b61195b81600561206790919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f660405160405180910390a250565b6119b581600561214290919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669260405160405180910390a250565b611a158282604051806020016040528060008152506118d6565b5050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611aa0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806124df6022913960400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b611b02848484611427565b611b0e84848484611d44565b611b63576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001806123556032913960400191505060405180910390fd5b50505050565b505050565b611b86600182600001546121ff90919063ffffffff16565b816000018190555050565b6001816000016000828254019250508190555050565b8173ffffffffffffffffffffffffffffffffffffffff16611bc782610c07565b73ffffffffffffffffffffffffffffffffffffffff1614611c33576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061257c6025913960400191505060405180910390fd5b611c3f82600083611b69565b611c4a60008261127a565b611c91600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020611b6e565b60006001600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a45050565b6000611d658473ffffffffffffffffffffffffffffffffffffffff16612249565b611d72576001905061205f565b600060608573ffffffffffffffffffffffffffffffffffffffff1663150b7a0260e01b611d9d611272565b898888604051602401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611e4d578082015181840152602081019050611e32565b50505050905090810190601f168015611e7a5780820380516001836020036101000a031916815260200191505b5095505050505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040518082805190602001908083835b60208310611f125780518252602082019150602081019050602083039250611eef565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611f74576040519150601f19603f3d011682016040523d82523d6000602084013e611f79565b606091505b509150915081611fe757600081511115611f965780518082602001fd5b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001806123556032913960400191505060405180910390fd5b6000818060200190516020811015611ffe57600080fd5b8101908080519060200190929190505050905063150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161493505050505b949350505050565b6120718282611a19565b156120e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b61214c8282611a19565b6121a1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806124926021913960400191505060405180910390fd5b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600061224183836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250612294565b905092915050565b60008060007fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060001b9050833f915080821415801561228b57506000801b8214155b92505050919050565b6000838311158290612341576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156123065780820151818401526020810190506122eb565b50505050905090810190601f1680156123335780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838503905080915050939250505056fe4552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f20616464726573734552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c654552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e526f6c65733a206163636f756e7420697320746865207a65726f20616464726573734552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65724552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f7665644552433732313a206275726e206f6620746f6b656e2074686174206973206e6f74206f776e4552433732314275726e61626c653a2063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564a2646970667358221220b4bae99e380e9266072820bb45ba1f6cd90a7eba70bea96528e6c8b4d8a6abcc64736f6c63430006040033"
