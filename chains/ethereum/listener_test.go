// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridge/blockstore"
	msg "github.com/ChainSafe/ChainBridge/message"
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

func createTestListener(t *testing.T, config *Config, contracts *DeployedContracts) (*Listener, *MockRouter) {
	// Create copy and add deployed contract addresses
	newConfig := *config
	newConfig.bridgeContract = contracts.BridgeAddress
	newConfig.erc20HandlerContract = contracts.ERC20HandlerAddress
	newConfig.startBlock = big.NewInt(0)

	conn := newLocalConnection(t, &newConfig)
	bridgeContract, err := createBridgeContract(newConfig.bridgeContract, conn)
	if err != nil {
		t.Fatal(err)
	}
	erc20HandlerContract, err := createErc20HandlerContract(newConfig.erc20HandlerContract, conn)
	if err != nil {
		t.Fatal(err)
	}

	router := &MockRouter{msgs: make(chan msg.Message)}
	listener := NewListener(conn, &newConfig, TestLogger, &blockstore.EmptyStore{})
	listener.SetContracts(bridgeContract, erc20HandlerContract)
	listener.SetRouter(router)
	// Start the listener
	err = listener.Start()
	if err != nil {
		t.Fatal(err)
	}

	return listener, router
}

func TestListener_start_stop(t *testing.T) {
	contracts := deployTestContracts(t, aliceTestConfig.id)
	l, _ := createTestListener(t, aliceTestConfig, contracts)

	err := l.Start()
	if err != nil {
		t.Fatal(err)
	}

	err = l.Stop()
	if err != nil {
		t.Fatal(err)
	}
}

func TestListener_depositEvent(t *testing.T) {
	contracts := deployTestContracts(t, aliceTestConfig.id)
	l, router := createTestListener(t, aliceTestConfig, contracts)

	// For debugging
	go watchEvent(l.conn, Deposit)

	// Get transaction ready
	opts, nonce, err := l.conn.newTransactOpts(big.NewInt(0), big.NewInt(DefaultGasLimit), big.NewInt(DefaultGasPrice))
	nonce.lock.Unlock() // We manual increment nonce in tests
	if err != nil {
		t.Fatal(err)
	}

	erc20Contract := deployMintApproveErc20(t, l.conn, opts)

	amount := big.NewInt(10)
	src := msg.ChainId(0)
	dst := msg.ChainId(1)
	resourceId := append(common.LeftPadBytes(erc20Contract.Bytes(), 31), uint8(src))
	recipient := ethcrypto.PubkeyToAddress(BobKp.PrivateKey().PublicKey)

	whitelistResourceId(t, l.conn.conn, opts, contracts.ERC20HandlerAddress, msg.ResourceIdFromSlice(resourceId), erc20Contract)

	expectedMessage := msg.NewFungibleTransfer(
		src,
		dst,
		1,
		amount,
		msg.ResourceIdFromSlice(resourceId),
		common.HexToAddress(BobKp.Address()).Bytes(),
	)
	// Create an ERC20 Deposit
	err = createErc20Deposit(
		l.bridgeContract,
		opts,
		erc20Contract,
		l.cfg.erc20HandlerContract,

		recipient,
		dst,
		amount,
	)
	if err != nil {
		t.Fatal(err)
	}

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
		msg.ResourceIdFromSlice(resourceId),
		common.HexToAddress(BobKp.Address()).Bytes(),
	)
	err = createErc20Deposit(
		l.bridgeContract,
		opts,
		erc20Contract,
		l.cfg.erc20HandlerContract,

		recipient,
		dst,
		amount,
	)
	if err != nil {
		t.Fatal(err)
	}

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
