// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"testing"
	"time"

	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func TestChain_ListenerShutdownOnFailure(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, msg.ChainId(1))
	cfg := &core.ChainConfig{
		Id:             msg.ChainId(1),
		Name:           "alice",
		Endpoint:       TestEndpoint,
		From:           keystore.AliceKey,
		Insecure:       true,
		KeystorePath:   keystore.AliceKey,
		BlockstorePath: "",
		FreshStart:     true,
		Opts: map[string]string{
			"bridge":         contracts.BridgeAddress.Hex(),
			"erc20Handler":   contracts.ERC20HandlerAddress.Hex(),
			"erc721Handler":  contracts.ERC721HandlerAddress.Hex(),
			"genericHandler": contracts.GenericHandlerAddress.Hex(),
			"gasLimit":       big.NewInt(DefaultGasLimit).String(),
			"maxGasPrice":    big.NewInt(DefaultGasPrice).String(),
		},
	}
	sysErr := make(chan error)
	chain, err := InitializeChain(cfg, TestLogger, sysErr, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = chain.Start()
	if err != nil {
		t.Fatal(err)
	}

	// Cause critical failure to polling mechanism
	chain.conn.Client().Close()

	// Pull expected error
	select {
	case err := <-sysErr:
		if err.Error() != ErrFatalPolling.Error() {
			t.Fatalf("Unexpected error: %s", err)
		}
	case <-time.After(time.Second * 30):
		t.Fatal("Test timed out")
	}

	// Tell everyone to shutdown
	chain.Stop()
}

func TestChain_WriterShutdownOnFailure(t *testing.T) {
	// Setup contracts and params for erc20 transfer
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, msg.ChainId(1))
	erc20Contract := ethtest.DeployMintApproveErc20(t, client, contracts.ERC20HandlerAddress, big.NewInt(100))
	src := msg.ChainId(5) // Not yet used, nonce should be 0
	dst := msg.ChainId(1)
	amount := big.NewInt(10)
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc20Contract.Bytes(), 31), uint8(src)))
	recipient := ethcrypto.PubkeyToAddress(BobKp.PrivateKey().PublicKey)
	ethtest.RegisterResource(t, client, contracts.BridgeAddress, contracts.ERC20HandlerAddress, resourceId, erc20Contract)

	// Start a chain
	cfg := &core.ChainConfig{
		Id:             dst,
		Name:           "alice",
		Endpoint:       TestEndpoint,
		From:           keystore.AliceKey,
		Insecure:       true,
		KeystorePath:   keystore.AliceKey,
		BlockstorePath: "",
		FreshStart:     true,
		Opts: map[string]string{
			"bridge":         contracts.BridgeAddress.Hex(),
			"erc20Handler":   contracts.ERC20HandlerAddress.Hex(),
			"erc721Handler":  contracts.ERC721HandlerAddress.Hex(),
			"genericHandler": contracts.GenericHandlerAddress.Hex(),
			"gasLimit":       big.NewInt(DefaultGasLimit).String(),
			"maxGasPrice":    big.NewInt(DefaultGasPrice).String(),
		},
	}
	sysErr := make(chan error)
	chain, err := InitializeChain(cfg, TestLogger, sysErr, nil)
	if err != nil {
		t.Fatal(err)
	}

	r := core.NewRouter(TestLogger)
	chain.SetRouter(r)

	err = chain.Start()
	if err != nil {
		t.Fatal(err)
	}

	// Submit some messages
	message := msg.NewFungibleTransfer(src, dst, 1, amount, resourceId, recipient.Bytes())

	for i := 0; i < 5; i++ {
		err = chain.listener.router.Send(message)
		if err != nil {
			t.Fatal(err)
		}

		message.DepositNonce++
	}

	time.Sleep(time.Second)
	// Cause critical failure for submitting txs
	chain.conn.Client().Close()

	// Pull expected error
	select {
	case err := <-sysErr:
		if err.Error() != ErrFatalPolling.Error() &&
			err.Error() != ErrFatalTx.Error() &&
			err.Error() != ErrFatalQuery.Error() {
			t.Fatalf("Unexpected error: %s", err)
		}
	case <-time.After(time.Second * 30):
		t.Fatal("Test timed out")
	}

	// Tell everyone to shutdown
	chain.Stop()
}
