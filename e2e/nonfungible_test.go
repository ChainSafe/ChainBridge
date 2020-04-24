// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package e2e

import (
	"fmt"
	"math/big"
	"testing"

	eth "github.com/ChainSafe/ChainBridge/e2e/ethereum"
	sub "github.com/ChainSafe/ChainBridge/e2e/substrate"
	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
	subtest "github.com/ChainSafe/ChainBridge/shared/substrate/testing"
	log "github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

func testErc721ToSubstrateRoundTrip(t *testing.T, ctx *testContext) {
	numberOfTxs := 5
	subRecipient := sub.AliceKp.AsKeyringPair().PublicKey // use alice so we can send back
	ethRecipient := eth.AliceKp.CommonAddress()
	tokenIds := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(5)}
	metadata := []byte("hello, world")

	nonce := ethtest.GetDepositNonce(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, SubChainId) + 1
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			id := tokenIds[i-1]
			log.Info("Submitting transaction", "number", i, "from", ctx.ethA.Opts.From, "handler", ctx.ethA.BaseContracts.ERC721HandlerAddress.String(), "tokenId", id.String())
			// Initiate transfer
			eth.CreateErc721Deposit(t, ctx.ethA.Client, ctx.ethA.Opts, SubChainId, subRecipient, id, metadata, ctx.ethA.BaseContracts, ctx.EthSubErc721ResourceId)

			// Wait for event
			sub.WaitForProposalSuccessOrFail(t, ctx.subClient, types.NewU64(nonce), types.U8(EthAChainId))
			nonce++

			// Verify ownership
			subtest.AssertOwnerOf(t, ctx.subClient, id, types.NewAccountID(subRecipient))
		})
		if !ok {
			return
		}
	}

	nonce = subtest.GetDepositNonce(t, ctx.subClient, EthAChainId) + 1
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			// Get latest eth block
			id := types.NewU256(*tokenIds[i-1])

			// Execute transfer
			subtest.InitiateNonFungibleTransfer(t, ctx.subClient, id, ethRecipient.Bytes(), EthAChainId)

			// Wait for event
			eth.WaitForDepositCreatedEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			eth.WaitForDepositExecutedEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			nonce++

			// Verify ownership
			ethtest.Erc721AssertOwner(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Sub, id.Int, ethRecipient)
		})
		if !ok {
			return
		}
	}

}

func testErc721EthToEthRoundTrip(t *testing.T, ctx *testContext) {
	ethARecipient := eth.AliceKp.CommonAddress()
	ethBRecipient := eth.AliceKp.CommonAddress()
	tokenIds := []*big.Int{big.NewInt(11), big.NewInt(12), big.NewInt(13), big.NewInt(14), big.NewInt(15)}

	numberOfTxs := 5
	nonce := ethtest.GetDepositNonce(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, EthBChainId) + 1
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer ethA to ethB %d", i), func(t *testing.T) {
			id := tokenIds[i-1]
			log.Info("Submitting transaction", "number", i, "recipient", ethBRecipient, "resourcId", ctx.EthEthErc721ResourceId.Hex(), "tokenId", id.String(), "from", ctx.ethA.Opts.From, "handler", ctx.ethA.BaseContracts.ERC721HandlerAddress)
			// Initiate transfer
			eth.CreateErc721Deposit(t, ctx.ethA.Client, ctx.ethA.Opts, EthBChainId, ethBRecipient.Bytes(), id, []byte{}, ctx.ethA.BaseContracts, ctx.EthEthErc721ResourceId)

			eth.WaitForDepositCreatedEvent(t, ctx.ethB.Client, ctx.ethB.BaseContracts.BridgeAddress, nonce)
			eth.WaitForDepositExecutedEvent(t, ctx.ethB.Client, ctx.ethB.BaseContracts.BridgeAddress, nonce)
			nonce++

			// Verify ownership
			ethtest.Erc721AssertOwner(t, ctx.ethB.Client, ctx.ethB.TestContracts.Erc721Eth, id, ethBRecipient)
		})
		if !ok {
			return
		}
	}

	// Aprove handler to now move the tokens
	ethtest.Erc721ApproveMany(t, ctx.ethB.Client, ctx.ethB.Opts, ctx.ethB.TestContracts.Erc721Eth, ctx.ethB.BaseContracts.ERC721HandlerAddress, 11, 5)
	nonce = ethtest.GetDepositNonce(t, ctx.ethB.Client, ctx.ethB.BaseContracts.BridgeAddress, EthAChainId) + 1
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer ethB to ethA %d", i), func(t *testing.T) {
			id := tokenIds[i-1]
			// TODO: Remove
			ethtest.Erc721AssertOwner(t, ctx.ethB.Client, ctx.ethB.TestContracts.Erc721Eth, id, ctx.ethB.Opts.From)

			log.Info("Submitting transaction", "number", i, "recipient", ethBRecipient, "resourceId", ctx.EthEthErc721ResourceId.Hex(), "tokenId", id.String(), "from", ctx.ethB.Opts.From, "handler", ctx.ethB.BaseContracts.ERC721HandlerAddress)
			// Initiate transfer
			eth.CreateErc721Deposit(t, ctx.ethB.Client, ctx.ethB.Opts, EthAChainId, ethARecipient.Bytes(), id, []byte{}, ctx.ethB.BaseContracts, ctx.EthEthErc721ResourceId)
			eth.WaitForDepositCreatedEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			eth.WaitForDepositExecutedEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			nonce++

			// Verify ownership
			ethtest.Erc721AssertOwner(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Eth, id, ethARecipient)
		})
		if !ok {
			return
		}
	}
}
