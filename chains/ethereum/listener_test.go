// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	"github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	"github.com/ChainSafe/ChainBridge/bindings/ERC721Handler"
	"github.com/ChainSafe/ChainBridge/bindings/GenericHandler"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
	"github.com/ChainSafe/chainbridge-utils/blockstore"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

type MockRouter struct {
	msgs chan msg.Message
}

func (r *MockRouter) Send(message msg.Message) error {
	r.msgs <- message
	return nil
}

func createTestListener(t *testing.T, config *Config, contracts *utils.DeployedContracts, stop <-chan int, sysErr chan<- error) (*listener, *MockRouter) {
	// Create copy and add deployed contract addresses
	newConfig := *config
	newConfig.bridgeContract = contracts.BridgeAddress
	newConfig.erc20HandlerContract = contracts.ERC20HandlerAddress
	newConfig.erc721HandlerContract = contracts.ERC721HandlerAddress
	newConfig.genericHandlerContract = contracts.GenericHandlerAddress

	conn := newLocalConnection(t, &newConfig)
	latestBlock, err := conn.LatestBlock()
	if err != nil {
		t.Fatal(err)
	}
	newConfig.startBlock = latestBlock

	bridgeContract, err := Bridge.NewBridge(newConfig.bridgeContract, conn.Client())
	if err != nil {
		t.Fatal(err)
	}
	erc20HandlerContract, err := ERC20Handler.NewERC20Handler(newConfig.erc20HandlerContract, conn.Client())
	if err != nil {
		t.Fatal(err)
	}
	erc721HandlerContract, err := ERC721Handler.NewERC721Handler(newConfig.erc721HandlerContract, conn.Client())
	if err != nil {
		t.Fatal(err)
	}
	genericHandlerContract, err := GenericHandler.NewGenericHandler(newConfig.genericHandlerContract, conn.Client())
	if err != nil {
		t.Fatal(err)
	}

	router := &MockRouter{msgs: make(chan msg.Message)}
	listener := NewListener(conn, &newConfig, TestLogger, &blockstore.EmptyStore{}, stop, sysErr, nil)
	listener.setContracts(bridgeContract, erc20HandlerContract, erc721HandlerContract, genericHandlerContract)
	listener.setRouter(router)
	// Start the listener
	err = listener.start()
	if err != nil {
		t.Fatal(err)
	}

	return listener, router
}

func verifyMessage(t *testing.T, r *MockRouter, expected msg.Message, errs chan error) {
	// Verify message
	select {
	case m := <-r.msgs:
		err := compareMessage(expected, m)
		if err != nil {
			t.Fatal(err)
		}
	case err := <-errs:
		t.Fatalf("Fatal error: %s", err)
	case <-time.After(TestTimeout):
		t.Fatalf("test timed out")
	}
}

func TestListener_start_stop(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, aliceTestConfig.id)
	stop := make(chan int)
	l, _ := createTestListener(t, aliceTestConfig, contracts, stop, nil)

	err := l.start()
	if err != nil {
		t.Fatal(err)
	}

	// Initiate shutdown
	close(stop)
}

func TestListener_Erc20DepositedEvent(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, aliceTestConfig.id)
	errs := make(chan error)
	l, router := createTestListener(t, aliceTestConfig, contracts, make(chan int), errs)

	// For debugging
	go ethtest.WatchEvent(client, contracts.BridgeAddress, utils.Deposit)

	erc20Contract := ethtest.DeployMintApproveErc20(t, client, contracts.ERC20HandlerAddress, big.NewInt(100))

	amount := big.NewInt(10)
	src := msg.ChainId(0)
	dst := msg.ChainId(1)
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc20Contract.Bytes(), 31), uint8(src)))
	recipient := ethcrypto.PubkeyToAddress(BobKp.PrivateKey().PublicKey)

	ethtest.RegisterResource(t, client, contracts.BridgeAddress, contracts.ERC20HandlerAddress, resourceId, erc20Contract)

	expectedMessage := msg.NewFungibleTransfer(
		src,
		dst,
		1,
		amount,
		resourceId,
		common.HexToAddress(BobKp.Address()).Bytes(),
	)
	// Create an ERC20 Deposit
	createErc20Deposit(
		t,
		l.bridgeContract,
		client,
		resourceId,

		recipient,
		dst,
		amount,
	)

	verifyMessage(t, router, expectedMessage, errs)

	// Create second deposit, verify nonce change
	expectedMessage = msg.NewFungibleTransfer(
		src,
		dst,
		2,
		amount,
		resourceId,
		common.HexToAddress(BobKp.Address()).Bytes(),
	)
	createErc20Deposit(
		t,
		l.bridgeContract,
		client,
		resourceId,

		recipient,
		dst,
		amount,
	)

	verifyMessage(t, router, expectedMessage, errs)
}

func TestListener_Erc721DepositedEvent(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, aliceTestConfig.id)
	errs := make(chan error)
	l, router := createTestListener(t, aliceTestConfig, contracts, make(chan int), errs)

	// For debugging
	go ethtest.WatchEvent(client, contracts.BridgeAddress, utils.Deposit)

	tokenId := big.NewInt(99)

	erc721Contract := ethtest.Erc721Deploy(t, client)
	ethtest.Erc721Mint(t, client, erc721Contract, tokenId, []byte{})
	ethtest.Erc721Approve(t, client, erc721Contract, contracts.ERC721HandlerAddress, tokenId)
	log15.Info("Deployed erc721, minted and approved handler", "handler", contracts.ERC721HandlerAddress, "contract", erc721Contract, "tokenId", tokenId.Bytes())
	ethtest.Erc721AssertOwner(t, client, erc721Contract, tokenId, client.Opts.From)
	src := msg.ChainId(0)
	dst := msg.ChainId(1)
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc721Contract.Bytes(), 31), uint8(src)))
	recipient := BobKp.CommonAddress()

	ethtest.RegisterResource(t, client, contracts.BridgeAddress, contracts.ERC721HandlerAddress, resourceId, erc721Contract)

	expectedMessage := msg.NewNonFungibleTransfer(
		src,
		dst,
		1,
		resourceId,
		tokenId,
		recipient.Bytes(),
		[]byte{},
	)

	// Create an ERC20 Deposit
	createErc721Deposit(
		t,
		l.bridgeContract,
		client,
		resourceId,

		recipient,
		dst,
		tokenId,
	)

	verifyMessage(t, router, expectedMessage, errs)
}

func TestListener_GenericDepositedEvent(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, aliceTestConfig.id)
	errs := make(chan error)
	l, router := createTestListener(t, aliceTestConfig, contracts, make(chan int), errs)

	// For debugging
	go ethtest.WatchEvent(client, contracts.BridgeAddress, utils.Deposit)

	src := msg.ChainId(0)
	dst := msg.ChainId(1)
	hash := utils.Hash(common.LeftPadBytes([]byte{1}, 32))
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes([]byte{1}, 31), uint8(src)))
	depositSig := utils.CreateFunctionSignature("")
	executeSig := utils.CreateFunctionSignature("store()")
	ethtest.RegisterGenericResource(t, client, contracts.BridgeAddress, contracts.GenericHandlerAddress, resourceId, utils.ZeroAddress, depositSig, executeSig)

	expectedMessage := msg.NewGenericTransfer(
		src,
		dst,
		1,
		resourceId,
		hash[:],
	)

	// Create an ERC20 Deposit
	createGenericDeposit(
		t,
		l.bridgeContract,
		client,
		resourceId,

		dst,
		hash[:],
	)

	verifyMessage(t, router, expectedMessage, errs)
}

func compareMessage(expected, actual msg.Message) error {
	if !reflect.DeepEqual(expected, actual) {
		if !reflect.DeepEqual(expected.Source, actual.Source) {
			return fmt.Errorf("Source doesn't match. \n\tExpected: %#v\n\tGot: %#v\n", expected.Source, actual.Source)
		} else if !reflect.DeepEqual(expected.Destination, actual.Destination) {
			return fmt.Errorf("Destination doesn't match. \n\tExpected: %#v\n\tGot: %#v\n", expected.Destination, actual.Destination)
		} else if !reflect.DeepEqual(expected.DepositNonce, actual.DepositNonce) {
			return fmt.Errorf("Deposit nonce doesn't match. \n\tExpected: %#v\n\tGot: %#v\n", expected.DepositNonce, actual.DepositNonce)
		} else if !reflect.DeepEqual(expected.Payload, actual.Payload) {
			return fmt.Errorf("Payload doesn't match. \n\tExpected: %#v\n\tGot: %#v\n", expected.Payload, actual.Payload)
		}
	}
	return nil
}
