// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package subtest

import (
	"fmt"
	"testing"

	"github.com/ChainSafe/ChainBridge/e2e/substrate"
	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

func TestChain_Events(t *testing.T) {
	targetURL := substrate.TestSubEndpoint // Replace with desired endpoint
	api, err := gsrpc.NewSubstrateAPI(targetURL)
	if err != nil {
		panic(err)
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		panic(err)
	}

	key, err := types.CreateStorageKey(meta, "System", "Events", nil, nil)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%x\n", key)

	//latest, err := api.RPC.Chain.GetBlockLatest()
	//if err != nil {
	//	panic(err)
	//}
	latestNumber := uint32(1) // Set to uint32(latest.Block.Header.Number)

	batchSize := uint32(1) // Set to higher value accordingly, like 1000
	current := uint64(0)   // Start block
	numBatches := latestNumber - uint32(current)/batchSize

	// Not smart enough to adjust batch size to last batch
	// Manually trigger the last one with minimum batch size
	for i := uint32(0); i < numBatches; i++ {
		lower, err := api.RPC.Chain.GetBlockHash(current)
		if err != nil {
			panic(err)
		}

		upperBlock := current + uint64(batchSize)
		upper, err := api.RPC.Chain.GetBlockHash(upperBlock)
		if err != nil {
			panic(err)
		}

		raws, err := api.RPC.State.QueryStorage([]types.StorageKey{key}, lower, upper)
		if err != nil {
			panic(err)
		}

		for j := 0; j < len(raws); j++ {
			events := utils.Events{}
			for k := 0; k < len(raws[j].Changes); k++ {
				raw := raws[j].Changes[k]
				fmt.Printf("Processing events for block %s with data: %x\n", raws[j].Block.Hex(), raw.StorageData)
				err = types.EventRecordsRaw(raw.StorageData).DecodeEventRecords(meta, &events)
				if err != nil {
					panic(err)
				}
			}
		}

		fmt.Println("Events batch successfully processed: ", i, "until block", upperBlock)
		current += uint64(batchSize)
	}

}
