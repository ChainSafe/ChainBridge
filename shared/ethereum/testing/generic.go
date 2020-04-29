// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethtest

import (
	"bytes"
	"testing"

	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func AssertGenericResourceAddress(t *testing.T, client *ethclient.Client, handler common.Address, rId msg.ResourceId, expected common.Address) {
	actual, err := utils.GetGenericResourceAddress(client, handler, rId)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(actual.Bytes(), expected.Bytes()) {
		t.Fatalf("Generic resoruce mismatch for ID %x. Expected address: %x Got: %x", rId, expected, actual)
	}
	log15.Info("Asserted generic resource ID", "handler", handler, "rId", rId.Hex(), "contract", actual)
}
