// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

type MockRouter struct{}

func (m *MockRouter) Send(message msg.Message) error { return nil }

func setupListener(t *testing.T, config *Config) *Listener {
	conn := newLocalConnection(t, config)
	bridgeContract := createBridgeInstance(t, *conn, config.contract)

	listener := NewListener(conn, config)
	router := &MockRouter{}
	listener.SetRouter(router)
	listener.SetBridgeContract(bridgeContract)
	return listener
}

func TestListener_start_stop(t *testing.T) {
	cfg, _ := testDeployContracts(t, defaultDeployOpts)
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

func TestListener_Event(t *testing.T) {
	cfg, contracts := testDeployContracts(t, defaultDeployOpts)
	l := setupListener(t, cfg)

	opts, _, err := l.conn.newTransactOpts(big.NewInt(0), cfg.gasLimit, cfg.gasPrice)
	if err != nil {
		t.Fatal(err)
	}

	amount := big.NewInt(1)
	if err := createFreshErc20Deposit(
		l.bridgeContract,
		l.conn,
		opts,
		common.HexToAddress(AliceKp.Address()),
		contracts.ERC20HandlerAddress,
		// Values below are random and do not matter since we are not doing an e2e test
		contracts.ERC20HandlerAddress,
		common.HexToAddress(BobKp.Address()),
		big.NewInt(1),
		amount,
	); err != nil {
		t.Fatal(err)
	}
}

// TODO TESTS
// registerEventHandler
// watchEvent <- Must trigger an event
