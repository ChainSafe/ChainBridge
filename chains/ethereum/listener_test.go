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
	"github.com/ChainSafe/ChainBridge/blockstore"
	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
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

func createTestListener(t *testing.T, config *Config, contracts *utils.DeployedContracts) (*listener, *MockRouter) {
	// Create copy and add deployed contract addresses
	newConfig := *config
	newConfig.bridgeContract = contracts.BridgeAddress
	newConfig.erc20HandlerContract = contracts.ERC20HandlerAddress
	newConfig.erc721HandlerContract = contracts.ERC721HandlerAddress
	newConfig.genericHandlerContract = contracts.GenericHandlerAddress

	conn := newLocalConnection(t, &newConfig)
	latestBlock, err := conn.latestBlock()
	if err != nil {
		t.Fatal(err)
	}
	newConfig.startBlock = latestBlock

	bridgeContract, err := Bridge.NewBridge(newConfig.bridgeContract, conn.conn)
	if err != nil {
		t.Fatal(err)
	}
	erc20HandlerContract, err := ERC20Handler.NewERC20Handler(newConfig.erc20HandlerContract, conn.conn)
	if err != nil {
		t.Fatal(err)
	}
	erc721HandlerContract, err := ERC721Handler.NewERC721Handler(newConfig.erc721HandlerContract, conn.conn)
	if err != nil {
		t.Fatal(err)
	}
	genericHandlerContract, err := GenericHandler.NewGenericHandler(newConfig.genericHandlerContract, conn.conn)
	if err != nil {
		t.Fatal(err)
	}

	router := &MockRouter{msgs: make(chan msg.Message)}
	listener := NewListener(conn, &newConfig, TestLogger, &blockstore.EmptyStore{})
	listener.setContracts(bridgeContract, erc20HandlerContract, erc721HandlerContract, genericHandlerContract)
	listener.setRouter(router)
	// Start the listener
	err = listener.start()
	if err != nil {
		t.Fatal(err)
	}

	return listener, router
}

func TestListener_start_stop(t *testing.T) {
	contracts := deployTestContracts(t, aliceTestConfig.id, AliceKp)
	l, _ := createTestListener(t, aliceTestConfig, contracts)

	err := l.start()
	if err != nil {
		t.Fatal(err)
	}

	err = l.stop()
	if err != nil {
		t.Fatal(err)
	}
}

func TestListener_Erc20DepositedEvent(t *testing.T) {
	contracts := deployTestContracts(t, aliceTestConfig.id, AliceKp)
	l, router := createTestListener(t, aliceTestConfig, contracts)

	// For debugging
	go watchEvent(l.conn, utils.Deposit)

	// Get transaction ready
	opts, _, err := l.conn.newTransactOpts(big.NewInt(0), big.NewInt(DefaultGasLimit), big.NewInt(DefaultGasPrice))
	if err != nil {
		t.Fatal(err)
	}

	erc20Contract := ethtest.DeployMintApproveErc20(t, l.conn.conn, opts, contracts.ERC20HandlerAddress, big.NewInt(100))

	amount := big.NewInt(10)
	src := msg.ChainId(0)
	dst := msg.ChainId(1)
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc20Contract.Bytes(), 31), uint8(src)))
	recipient := ethcrypto.PubkeyToAddress(BobKp.PrivateKey().PublicKey)

	ethtest.RegisterResource(t, l.conn.conn, opts, contracts.BridgeAddress, contracts.ERC20HandlerAddress, resourceId, erc20Contract)

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
		opts,
		resourceId,
		l.cfg.erc20HandlerContract,

		recipient,
		dst,
		amount,
	)

	// Verify message
	select {
	case m := <-router.msgs:
		err = compareMessage(expectedMessage, m)
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(TestTimeout):
		t.Fatalf("test timed out")
	}

	// Create second deposit, verify nonce change
	expectedMessage2 := msg.NewFungibleTransfer(
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
		opts,
		resourceId,
		l.cfg.erc20HandlerContract,

		recipient,
		dst,
		amount,
	)

	// Verify message
	select {
	case m := <-router.msgs:
		err = compareMessage(expectedMessage2, m)
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(TestTimeout):
		t.Fatalf("test timed out")
	}
}

func TestListener_Erc721DepositedEvent(t *testing.T) {
	contracts := deployTestContracts(t, aliceTestConfig.id, AliceKp)
	l, router := createTestListener(t, aliceTestConfig, contracts)

	// For debugging
	go watchEvent(l.conn, utils.Deposit)

	// Get transaction ready
	opts, _, err := l.conn.newTransactOpts(big.NewInt(0), big.NewInt(DefaultGasLimit), big.NewInt(DefaultGasPrice))
	if err != nil {
		t.Fatal(err)
	}

	tokenId := big.NewInt(99)
	erc721Contract := ethtest.Erc721Deploy(t, l.conn.conn, opts)
	ethtest.Erc721Mint(t, l.conn.conn, opts, erc721Contract, tokenId, []byte{})
	ethtest.Erc721Approve(t, l.conn.conn, opts, erc721Contract, contracts.ERC721HandlerAddress, tokenId)
	log15.Info("Deployed erc721, minted and approved handler", "handler", contracts.ERC721HandlerAddress, "contract", erc721Contract, "tokenId", tokenId.Bytes())
	ethtest.Erc721AssertOwner(t, l.conn.conn, opts, erc721Contract, tokenId, opts.From)
	src := msg.ChainId(0)
	dst := msg.ChainId(1)
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc721Contract.Bytes(), 31), uint8(src)))
	recipient := BobKp.CommonAddress()

	ethtest.RegisterResource(t, l.conn.conn, opts, contracts.BridgeAddress, contracts.ERC721HandlerAddress, resourceId, erc721Contract)

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
		opts,
		resourceId,
		l.cfg.erc721HandlerContract,

		recipient,
		dst,
		tokenId,
		[]byte{}, // Empty metadata
	)

	// Verify message
	select {
	case m := <-router.msgs:
		err = compareMessage(expectedMessage, m)
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(TestTimeout):
		t.Fatalf("test timed out")
	}
}

func TestListener_GenericDepositedEvent(t *testing.T) {
	contracts := deployTestContracts(t, aliceTestConfig.id, AliceKp)
	l, router := createTestListener(t, aliceTestConfig, contracts)

	// For debugging
	go watchEvent(l.conn, utils.Deposit)

	// Get transaction ready
	opts, _, err := l.conn.newTransactOpts(big.NewInt(0), big.NewInt(DefaultGasLimit), big.NewInt(DefaultGasPrice))
	if err != nil {
		t.Fatal(err)
	}

	src := msg.ChainId(0)
	dst := msg.ChainId(1)
	hash := utils.Hash(common.LeftPadBytes([]byte{1}, 32))
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes([]byte{1}, 31), uint8(src)))
	depositSig := utils.CreateFunctionSignature("")
	executeSig := utils.CreateFunctionSignature("store()")
	ethtest.RegisterGenericResource(t, l.conn.conn, opts, contracts.BridgeAddress, contracts.GenericHandlerAddress, resourceId, utils.ZeroAddress, depositSig, executeSig)

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
		opts,
		resourceId,
		l.cfg.genericHandlerContract,

		dst,
		hash[:],
	)

	// Verify message
	select {
	case m := <-router.msgs:
		err = compareMessage(expectedMessage, m)
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(TestTimeout):
		t.Fatalf("test timed out")
	}
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
