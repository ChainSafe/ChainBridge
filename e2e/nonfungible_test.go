// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package e2e

import (
	"fmt"
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
	tokens := ethtest.GenerateErc721Tokens(1, numberOfTxs)

	ethtest.Erc721MintMany(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Sub, tokens)
	ethtest.Erc721ApproveMany(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Sub, ctx.ethA.BaseContracts.ERC721HandlerAddress, tokens)
	nonce := ethtest.GetDepositNonce(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, SubChainId) + 1
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			tok := tokens[i-1]
			// Verify ownership and intact metadata
			ethtest.Erc721AssertOwner(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Sub, tok.Id, ctx.ethA.Client.Opts.From)
			ethtest.Erc721AssertMetadata(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Sub, tok.Id, string(tok.Metadata[:]))

			// Initiate transfer
			log.Info("Submitting transaction", "number", i, "from", ctx.ethA.Client.Opts.From, "handler", ctx.ethA.BaseContracts.ERC721HandlerAddress.String(), "tokenId", tok.Id.String())
			eth.CreateErc721Deposit(t, ctx.ethA.Client, SubChainId, subRecipient, tok.Id, ctx.ethA.BaseContracts, ctx.EthSubErc721ResourceId)

			// Wait for event
			sub.WaitForProposalSuccessOrFail(t, ctx.subClient, types.NewU64(nonce), types.U8(EthAChainId))
			nonce++

			// Verify substrate ownership and metadata
			subtest.AssertOwnerOf(t, ctx.subClient, tok.Id, types.NewAccountID(subRecipient))
			subtest.AssertErc721Metadata(t, ctx.subClient, tok.Id, tok.Metadata[:])

			// Verify token no longer exists on ethA
			ethtest.Erc721AssertNonExistence(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Sub, tok.Id)
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
			tok := tokens[i-1]

			// Execute transfer
			subtest.InitiateNonFungibleTransfer(t, ctx.subClient, types.NewU256(*tok.Id), ethRecipient.Bytes(), EthAChainId)

			// Wait for event
			eth.WaitForProposalActive(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			eth.WaitForProposalExecutedEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			nonce++

			// Verify ownership and intact metadata
			ethtest.Erc721AssertOwner(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Sub, tok.Id, ethRecipient)
			ethtest.Erc721AssertMetadata(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Sub, tok.Id, string(tok.Metadata[:]))

			// Verify token no longer exists on substrate
			subtest.AssertErc721NonExistence(t, ctx.subClient, tok.Id)
		})
		if !ok {
			return
		}
	}

}

func testErc721EthToEthRoundTrip(t *testing.T, ctx *testContext) {
	ethARecipient := eth.AliceKp.CommonAddress()
	ethBRecipient := eth.AliceKp.CommonAddress()
	numberOfTxs := 5
	tokens := ethtest.GenerateErc721Tokens(11, numberOfTxs)
	ethtest.Erc721MintMany(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Eth, tokens)
	ethtest.Erc721ApproveMany(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Eth, ctx.ethA.BaseContracts.ERC721HandlerAddress, tokens)

	nonce := ethtest.GetDepositNonce(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, EthBChainId) + 1
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer ethA to ethB %d", i), func(t *testing.T) {
			tok := tokens[i-1]

			// Verify ownership and intact metadata on ethA
			ethtest.Erc721AssertOwner(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Eth, tok.Id, ctx.ethA.Client.Opts.From)
			ethtest.Erc721AssertMetadata(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Eth, tok.Id, string(tok.Metadata[:]))

			// Verify non-existence on ethB
			ethtest.Erc721AssertNonExistence(t, ctx.ethB.Client, ctx.ethB.TestContracts.Erc721Eth, tok.Id)

			// Initiate transfer
			log.Info("Submitting transaction", "number", i, "recipient", ethBRecipient, "resourcId", ctx.EthEthErc721ResourceId.Hex(), "tokenId", tok.Id.String(), "from", ctx.ethA.Client.Opts.From, "handler", ctx.ethA.BaseContracts.ERC721HandlerAddress)
			eth.CreateErc721Deposit(t, ctx.ethA.Client, EthBChainId, ethBRecipient.Bytes(), tok.Id, ctx.ethA.BaseContracts, ctx.EthEthErc721ResourceId)

			// Wait for proposal events
			eth.WaitForProposalActive(t, ctx.ethB.Client, ctx.ethB.BaseContracts.BridgeAddress, nonce)
			eth.WaitForProposalExecutedEvent(t, ctx.ethB.Client, ctx.ethB.BaseContracts.BridgeAddress, nonce)
			nonce++

			// Verify ownership on ethB
			ethtest.Erc721AssertOwner(t, ctx.ethB.Client, ctx.ethB.TestContracts.Erc721Eth, tok.Id, ethBRecipient)
			ethtest.Erc721AssertMetadata(t, ctx.ethB.Client, ctx.ethB.TestContracts.Erc721Eth, tok.Id, string(tok.Metadata[:]))

			// Verify non-existence on ethA
			ethtest.Erc721AssertNonExistence(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Eth, tok.Id)
		})
		if !ok {
			return
		}
	}

	// Aprove handler to now move the tokens
	ethtest.Erc721ApproveMany(t, ctx.ethB.Client, ctx.ethB.TestContracts.Erc721Eth, ctx.ethB.BaseContracts.ERC721HandlerAddress, tokens)

	nonce = ethtest.GetDepositNonce(t, ctx.ethB.Client, ctx.ethB.BaseContracts.BridgeAddress, EthAChainId) + 1
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer ethB to ethA %d", i), func(t *testing.T) {
			tok := tokens[i-1]

			// Initiate transfer
			log.Info("Submitting transaction", "number", i, "recipient", ethBRecipient, "resourceId", ctx.EthEthErc721ResourceId.Hex(), "tokenId", tok.Id.String(), "from", ctx.ethB.Client.Opts.From, "handler", ctx.ethB.BaseContracts.ERC721HandlerAddress)
			eth.CreateErc721Deposit(t, ctx.ethB.Client, EthAChainId, ethARecipient.Bytes(), tok.Id, ctx.ethB.BaseContracts, ctx.EthEthErc721ResourceId)

			// Wait for proposal events
			eth.WaitForProposalActive(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			eth.WaitForProposalExecutedEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			nonce++

			// Verify ownership and metadata on ethA
			ethtest.Erc721AssertOwner(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Eth, tok.Id, ethARecipient)
			//ethtest.Erc721AssertMetadata(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc721Eth, tok.Id, string(tok.Metadata[:]))

			// Verfy non-existence on ethB
			ethtest.Erc721AssertNonExistence(t, ctx.ethB.Client, ctx.ethB.TestContracts.Erc721Eth, tok.Id)
		})
		if !ok {
			return
		}
	}
}
