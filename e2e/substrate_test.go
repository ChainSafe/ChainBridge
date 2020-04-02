package e2e

import (
	"testing"

	"github.com/ChainSafe/ChainBridge/chains/substrate"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type subClient struct {
	api *gsrpc.SubstrateAPI
	meta *types.Metadata
	genesis types.Hash
}

func createSubClient(t *testing.T) *subClient {
	c := &subClient{}
	api, err := gsrpc.NewSubstrateAPI(TestSubEndpoint)
	if err != nil {
		t.Fatal(err)
	}
	c.api = api

	// Fetch metadata
	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		t.Fatal(err)
	}
	c.meta = meta

	// Fetch genesis hash
	genesisHash, err := c.api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		t.Fatal(err)
	}
	c.genesis = genesisHash

	return c
}

func watchForProposalSucceeded(t *testing.T, client *subClient, success chan bool) {
	key, err := types.CreateStorageKey(client.meta, "System", "Events", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	sub, err := client.api.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
	if err != nil {
		t.Fatal(err)
	}

	for {
		set := <- sub.Chan()
		for _, chng := range set.Changes {
			if !types.Eq(chng.StorageKey, key) || !chng.HasStorageData {
				// skip, we are only interested in events with countent
				continue
			}

			// Decode the event records
			events := substrate.Events{}
			err = types.EventRecordsRaw(chng.StorageData).DecodeEventRecords(client.meta, &events)
			if err != nil {
				t.Fatal(err)
			}

			// Show what we are busy with
			for _, _ = range events.Bridge_ProposalSucceeded {
				success <- true
			}
		}
	}
}

func watchForProposalFailed(t *testing.T, client *subClient, fail chan bool) {
	key, err := types.CreateStorageKey(client.meta, "System", "Events", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	sub, err := client.api.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
	if err != nil {
		t.Fatal(err)
	}

	for {
		set := <- sub.Chan()
		for _, chng := range set.Changes {
			if !types.Eq(chng.StorageKey, key) || !chng.HasStorageData {
				// skip, we are only interested in events with countent
				continue
			}

			// Decode the event records
			events := substrate.Events{}
			err = types.EventRecordsRaw(chng.StorageData).DecodeEventRecords(client.meta, &events)
			if err != nil {
				t.Fatal(err)
			}

			// Show what we are busy with
			for _, _ = range events.Bridge_ProposalFailed {
				fail <- true
			}
		}
	}
}