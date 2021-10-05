// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package subtest

import (
	"testing"
	"time"

	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

var TestTimeout = time.Second * 15

func WaitForRemarkEvent(t *testing.T, client *utils.Client, hash types.Hash) {
	key, err := types.CreateStorageKey(client.Meta, "System", "Events", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	sub, err := client.Api.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
	if err != nil {
		t.Fatal(err)
	}

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

				for _, evt := range events.Example_Remark {
					if evt.Hash == hash {
						log15.Info("Found matching Remark event", "hash", evt.Hash)
						return
					} else {
						log15.Info("Found mismatched Remark event", "hash", evt.Hash)
					}
				}
			}
		}

	}
}
