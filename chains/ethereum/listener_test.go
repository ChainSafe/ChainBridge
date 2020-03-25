// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var listenerTestConfig = &Config{
	id:       msg.EthereumId,
	endpoint: TestEndpoint,
	from:     keystore.AliceKey,
}

func listener_deployContracts(t *testing.T) {
	port := "8545"
	numRelayers := 2
	relayerThreshold := big.NewInt(1)
	pk := hexutil.Encode(AliceKp.Encode())[2:]
	DeployedContracts, err := DeployContracts(pk, port, numRelayers, relayerThreshold, uint8(0))
	if err != nil {
		t.Fatal(err)
	}

	listenerTestConfig.contract = DeployedContracts.BridgeAddress
}

func TestListener_start_stop(t *testing.T) {
	conn := newLocalConnection(t, listenerTestConfig)
	defer conn.Close()

	listener := NewListener(conn, listenerTestConfig)

	err := listener.Start()
	if err != nil {
		t.Fatal(err)
	}

	err = listener.Stop()
	if err != nil {
		t.Fatal(err)
	}
}

// TODO TESTS
// registerEventHandler
// watchEvent <- Must trigger an event
