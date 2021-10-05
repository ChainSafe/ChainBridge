// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"
	"os"
	"testing"
	"time"

	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

const TestSubEndpoint = "ws://localhost:9944"

var TestTimeout = time.Second * 30

var log = log15.New("e2e", "substrate")

var AliceKp = keystore.TestKeyRing.SubstrateKeys[keystore.AliceKey]
var BobKp = keystore.TestKeyRing.SubstrateKeys[keystore.BobKey]
var CharlieKp = keystore.TestKeyRing.SubstrateKeys[keystore.CharlieKey]
var DaveKp = keystore.TestKeyRing.SubstrateKeys[keystore.DaveKey]
var EveKp = keystore.TestKeyRing.SubstrateKeys[keystore.EveKey]

var RelayerSet = []types.AccountID{
	types.NewAccountID(BobKp.AsKeyringPair().PublicKey),
	types.NewAccountID(CharlieKp.AsKeyringPair().PublicKey),
	types.NewAccountID(DaveKp.AsKeyringPair().PublicKey),
}

func CreateConfig(key string, chain msg.ChainId) *core.ChainConfig {
	return &core.ChainConfig{
		Name:           fmt.Sprintf("substrate(%s)", key),
		Id:             chain,
		Endpoint:       TestSubEndpoint,
		From:           "",
		KeystorePath:   key,
		Insecure:       true,
		FreshStart:     true,
		BlockstorePath: os.TempDir(),
		Opts:           map[string]string{"useExtendedCall": "true"},
	}
}

func WaitForProposalSuccessOrFail(t *testing.T, client *utils.Client, nonce types.U64, chain types.U8) {
	key, err := types.CreateStorageKey(client.Meta, "System", "Events", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	sub, err := client.Api.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
	if err != nil {
		t.Fatal(err)
	}
	defer sub.Unsubscribe()

	timeout := time.After(TestTimeout)
	for {
		select {
		case <-timeout:
			t.Fatalf("Timed out waiting for proposal success/fail event")
		case set := <-sub.Chan():
			for _, chng := range set.Changes {
				if !types.Eq(chng.StorageKey, key) || !chng.HasStorageData {
					// skip, we are only interested in events with content
					continue
				}

				// Decode the event records
				events := utils.Events{}
				err = types.EventRecordsRaw(chng.StorageData).DecodeEventRecords(client.Meta, &events)
				if err != nil {
					t.Fatal(err)
				}

				for _, evt := range events.ChainBridge_ProposalSucceeded {
					if evt.DepositNonce == nonce && evt.SourceId == chain {
						log.Info("Proposal succeeded", "depositNonce", evt.DepositNonce, "source", evt.SourceId)
						return
					} else {
						log.Info("Found mismatched success event", "depositNonce", evt.DepositNonce, "source", evt.SourceId)
					}
				}

				for _, evt := range events.ChainBridge_ProposalFailed {
					if evt.DepositNonce == nonce && evt.SourceId == chain {
						log.Info("Proposal failed", "depositNonce", evt.DepositNonce, "source", evt.SourceId)
						t.Fatalf("Proposal failed. Nonce: %d Source: %d", evt.DepositNonce, evt.SourceId)
					} else {
						log.Info("Found mismatched fail event", "depositNonce", evt.DepositNonce, "source", evt.SourceId)
					}
				}
			}
		}

	}
}

func HashInt(i int) types.Hash {
	hash, err := types.GetHash(types.NewI64(int64(i)))
	if err != nil {
		panic(err)
	}
	return hash
}
