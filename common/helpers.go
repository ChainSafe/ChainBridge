package common

import (
	"encoding/hex"
	"fmt"
	"syscall"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	// In your configuration file, please register as many of the following events
	DepositAsset = "DepositAsset"
	NftTransfer  = "NftTransfer"
	ErcTransfer  = "ErcTransfer"

	// For internal usage only
	DepositAssetSignature = "DepositAsset(address,bytes32)"
	NftTransferSignature  = "NFTTransfer(uint256,uint256,address,address,uint256,bytes)"
	ErcTransferSignature  = "ERCTransfer(uint256,uint256,address,uint256,address)"
)

// GetPassword displays a message and prompts user to enter a password
func GetPassword(msg string) []byte {
	for {
		fmt.Println(msg)
		fmt.Print("> ")
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Printf("invalid input: %s\n", err)
		} else {
			fmt.Printf("\n")
			return password
		}
	}
}

// StringToAddress turns an Ethereum address in string form into a byte array
func StringToAddress(addr string) [20]byte {
	addrBytes, err := hex.DecodeString(addr)
	if err != nil {
		return [20]byte{}
	}
	addrArr := [20]byte{}
	copy(addrArr[:], addrBytes)
	return addrArr
}

// FunctionId returns a 4-byte function ID given the signature
func FunctionId(sig string) []byte {
	bytes := []byte(sig)
	hash := ethcrypto.Keccak256(bytes)
	return hash[0:4]
}

func BuildEventSubscriptions(events []string) []string {
	var arr = []string{}
	for _, event := range events {
		if event == DepositAsset {
			arr = append(arr, DepositAssetSignature)
		} else if event == NftTransfer {
			arr = append(arr, NftTransferSignature)
		} else if event == ErcTransfer {
			arr = append(arr, ErcTransferSignature)
		}
	}
	return arr
}
