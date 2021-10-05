// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package subtest

import (
	"bytes"
	"math/big"
	"testing"

	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

func CreateClient(t *testing.T, key *signature.KeyringPair, endpoint string) *utils.Client {
	client, err := utils.CreateClient(key, endpoint)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func AddRelayer(t *testing.T, client *utils.Client, relayer types.AccountID) {
	err := client.AddRelayer(relayer)
	if err != nil {
		t.Fatal(err)
	}
}

func WhitelistChain(t *testing.T, client *utils.Client, id msg.ChainId) {
	err := client.WhitelistChain(id)
	if err != nil {
		t.Fatal(err)
	}
}

func InitiateNativeTransfer(t *testing.T, client *utils.Client, amount types.U128, recipient []byte, destId msg.ChainId) {
	err := client.InitiateNativeTransfer(amount, recipient, destId)
	if err != nil {
		t.Fatal(err)
	}
}

func InitiateNonFungibleTransfer(t *testing.T, client *utils.Client, tokenId types.U256, recipient []byte, destId msg.ChainId) {
	err := client.InitiateNonFungibleTransfer(tokenId, recipient, destId)
	if err != nil {
		t.Fatal(err)
	}
}

func InitiateHashTransfer(t *testing.T, client *utils.Client, hash types.Hash, destId msg.ChainId) {
	err := client.InitiateHashTransfer(hash, destId)
	if err != nil {
		t.Fatal(err)
	}
}

func RegisterResource(t *testing.T, client *utils.Client, id msg.ResourceId, method string) {
	err := client.RegisterResource(id, method)
	if err != nil {
		t.Fatal(err)
	}
}

func BalanceOf(t *testing.T, client *utils.Client, publicKey []byte) *big.Int {
	balance, err := utils.BalanceOf(client, publicKey)
	if err != nil {
		t.Fatal(err)
	}
	return balance
}

func AssertBalanceOf(t *testing.T, client *utils.Client, publicKey []byte, expected *big.Int) {
	current, err := utils.BalanceOf(client, publicKey)
	if err != nil {
		t.Fatal(err)
	}

	if expected.Cmp(current) != 0 {
		t.Fatalf("Balance does not match expected. Expected: %s Got: %s (change %s) \n", expected.String(), current.String(), big.NewInt(0).Sub(current, expected).String())
	}
}

func MintErc721(t *testing.T, client *utils.Client, tokenId *big.Int, metadata []byte, recipient *signature.KeyringPair) {
	err := client.MintErc721(tokenId, metadata, recipient)
	if err != nil {
		t.Fatal(err)
	}
}

func AssertOwnerOf(t *testing.T, client *utils.Client, tokenId *big.Int, expected types.AccountID) {
	owner, err := client.OwnerOf(tokenId)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(owner[:], expected[:]) {
		t.Fatalf("Owner does not match for token %s. Got: %x expected: %x", tokenId.String(), owner, expected)
	}
	log15.Info("Asserted ownership of erc721", "tokenId", tokenId.String(), "owner", owner)
}

func GetDepositNonce(t *testing.T, client *utils.Client, chain msg.ChainId) uint64 {
	count, err := client.GetDepositNonce(chain)
	if err != nil {
		t.Fatal(err)
	}
	return count
}

func NewNativeTransferCall(t *testing.T, client *utils.Client, amount types.U128, recipient []byte, destId msg.ChainId) types.Call {
	call, err := client.NewNativeTransferCall(amount, recipient, destId)
	if err != nil {
		t.Fatal(err)
	}
	return call
}
