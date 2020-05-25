// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethtest

import (
	"errors"
	"math/big"
	"testing"

	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
)

// Erc721 is used for convenience to represent an Erc721
type Erc721 struct {
	Id       *big.Int
	Metadata [32]byte
}

func GenerateErc721Tokens(start int, numOfTokens int) []Erc721 {
	var res []Erc721
	for i := start; i < start+numOfTokens; i++ {
		token := Erc721{
			Id:       big.NewInt(int64(i)),
			Metadata: utils.Hash([]byte{byte(i)}),
		}
		res = append(res, token)
	}
	return res
}

func Erc721Deploy(t *testing.T, client *utils.Client) common.Address {
	addr, err := utils.DeployErc721(client)
	if err != nil {
		t.Fatal(err)
	}
	return addr
}

func Erc721Mint(t *testing.T, client *utils.Client, erc721Contract common.Address, id *big.Int, metadata []byte) {
	log15.Info("Minting erc721 token", "contract", erc721Contract.Hex(), "Id", id)
	err := utils.Erc721Mint(client, erc721Contract, id, metadata)
	if err != nil {
		t.Fatal(err)
	}
}

func Erc721MintMany(t *testing.T, client *utils.Client, erc721Contract common.Address, tokens []Erc721) {
	for _, tok := range tokens {
		Erc721Mint(t, client, erc721Contract, tok.Id, tok.Metadata[:])
	}
}

func Erc721Approve(t *testing.T, client *utils.Client, erc721Contract, recipient common.Address, tokenId *big.Int) {
	log15.Info("Approving erc721 token for transfer", "contract", erc721Contract.Hex(), "Id", tokenId, "recipient", recipient)
	err := utils.ApproveErc721(client, erc721Contract, recipient, tokenId)
	if err != nil {
		t.Fatal(err)
	}
}

func Erc721ApproveMany(t *testing.T, client *utils.Client, erc721Contract, recipient common.Address, tokens []Erc721) {
	for _, tok := range tokens {
		Erc721Approve(t, client, erc721Contract, recipient, tok.Id)
	}
}

func Erc721AssertOwner(t *testing.T, client *utils.Client, erc721Contract common.Address, tokenId *big.Int, expected common.Address) {
	addr, err := utils.OwnerOf(client, erc721Contract, tokenId)
	if err != nil {
		t.Fatal(err)
	}

	if addr != expected {
		t.Fatalf("address %s does not own %x, %s does", expected.Hex(), tokenId.Bytes(), addr.Hex())
	}
	log15.Info("Asserted ownership of erc721", "tokenId", tokenId, "owner", addr)
}

func Erc721FundHandler(t *testing.T, client *utils.Client, handler, erc721Contract common.Address, tokenId *big.Int) {
	err := utils.FundErc721Handler(client, handler, erc721Contract, tokenId)
	if err != nil {
		t.Fatal(err)
	}
	log15.Info("Funded handler with erc721 token", "tokenId", tokenId, "handler", handler)
}

func Erc721FundHandlerMany(t *testing.T, client *utils.Client, handler, erc721Contract common.Address, start, numberOfTokens int) {
	for i := start; i < start+numberOfTokens; i++ {
		Erc721FundHandler(t, client, handler, erc721Contract, big.NewInt(int64(i)))
	}
}

func Erc721AddMinter(t *testing.T, client *utils.Client, erc721Contract common.Address, minter common.Address) {
	err := utils.Erc721AddMinter(client, erc721Contract, minter)
	if err != nil {
		t.Fatal(err)
	}
}

func Erc721AssertMetadata(t *testing.T, client *utils.Client, erc721Contract common.Address, tokenId *big.Int, expected string) {
	actual, err := utils.Erc721GetTokenURI(client, erc721Contract, tokenId)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("erc721 metadata mismatch for token %s. Expected: %s Got: %s", tokenId.String(), expected, actual)
	}
}

var NonExistentTokenError = errors.New("VM Exception while processing transaction: revert ERC721: owner query for nonexistent token")

func Erc721AssertNonExistence(t *testing.T, client *utils.Client, erc721Contract common.Address, id *big.Int) {
	_, err := utils.OwnerOf(client, erc721Contract, id)
	// TODO: Assert actual revert, not possible with geth currently
	if err != nil {
		//if err.Error() != NonExistentTokenError.Error() {
		//	t.Fatal(err)
		//}
		log15.Info("Asserted non-existence of erc721", "tokenId", id, "result", err)
	}

}
