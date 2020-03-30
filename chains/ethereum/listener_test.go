// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	msg "github.com/ChainSafe/ChainBridge/message"
)

const ListenerTimeout = time.Second * 10

type MockRouter struct {
	msgs chan msg.Message
}

func (r *MockRouter) Send(message msg.Message) error {
	r.msgs <- message
	return nil
}

func setupListener(t *testing.T, config *Config) (*Listener, *MockRouter) {
	conn := newLocalConnection(t, config)
	bridgeContract := createBridgeInstance(t, conn, config.contract)

	router := &MockRouter{msgs: make(chan msg.Message)}
	listener := NewListener(conn, config)
	listener.SetBridgeContract(bridgeContract)
	listener.SetRouter(router)
	// Start the listener
	err := listener.Start()
	if err != nil {
		t.Fatal(err)
	}
	return listener, router
}

func TestListener_start_stop(t *testing.T) {
	cfg, _ := deployContracts(t, defaultDeployOpts)
	conn := newLocalConnection(t, cfg)
	defer conn.Close()

	listener := NewListener(conn, cfg)

	err := listener.Start()
	if err != nil {
		t.Fatal(err)
	}

	err = listener.Stop()
	if err != nil {
		t.Fatal(err)
	}
}

func TestListener_depositEvent(t *testing.T) {
	cfg, contracts := deployContracts(t, defaultDeployOpts)
	l, router := setupListener(t, cfg)

	// Get transaction ready
	opts, _, err := l.conn.newTransactOpts(big.NewInt(0), cfg.gasLimit, cfg.gasPrice)
	if err != nil {
		t.Fatal(err)
	}

	// Set some constants
	amount := big.NewInt(1)
	destId := big.NewInt(1)

	expectedMessage := msg.Message{
		Source:       msg.ChainId(0),
		Destination:  msg.ChainId(destId.Uint64()),
		Type:         msg.CreateDepositProposalType,
		DepositNonce: uint32(1),
		To:           contracts.ERC20HandlerAddress.Bytes(),
		Metadata:     amount.Bytes(),
	}

	// Create an ERC20 Deposit
	if err := createErc20Deposit(
		l.bridgeContract,
		l.conn,
		opts,
		common.HexToAddress(AliceKp.Address()),
		contracts.ERC20HandlerAddress,
		// Values below are random and do not matter since we are not doing an e2e test
		contracts.ERC20HandlerAddress,
		common.HexToAddress(BobKp.Address()),
		destId,
		amount,
	); err != nil {
		t.Fatal(err)
	}

	// Verify message
	select {
	case m := <-router.msgs:
		if !reflect.DeepEqual(expectedMessage, m) {
			t.Fatalf("Unexpected message.\n\tExpected: %#v\n\tGot: %#v\n", expectedMessage, m)
		}
	case <-time.After(ListenerTimeout):
		t.Fatalf("test timed out")
	}
}

func TestListener_createProposalEvent(t *testing.T) {
	cfg, _ := deployContracts(t, defaultDeployOpts)
	l, router := setupListener(t, cfg)

	// Get transaction ready
	opts, _, err := l.conn.newTransactOpts(big.NewInt(0), cfg.gasLimit, cfg.gasPrice)
	if err != nil {
		t.Fatal(err)
	}

	metadata := hash([]byte{})

	expectedMessage := msg.Message{
		Source:       msg.ChainId(0),
		Destination:  msg.ChainId(0),
		Type:         msg.VoteDepositProposalType,
		DepositNonce: uint32(1),
		Metadata:     metadata[:],
	}

	// Create an random deposit proposal
	if err := createDepositProposal(
		l.bridgeContract,
		l.conn,
		opts,
		big.NewInt(0),
		big.NewInt(1),
		metadata,
	); err != nil {
		t.Fatal(err)
	}

	// Verify message
	select {
	case m := <-router.msgs:
		if !reflect.DeepEqual(expectedMessage, m) {
			t.Fatalf("Unexpected message.\n\tExpected: %#v\n\tGot: %#v\n", expectedMessage, m)
		}
	case <-time.After(ListenerTimeout):
		t.Fatalf("test timed out")
	}
}
