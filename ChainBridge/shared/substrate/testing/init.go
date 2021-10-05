// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package subtest

import (
	"fmt"
	"testing"

	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

// WARNING: THIS METHOD IS UNSAFE AND MAY PANIC
func EnsureInitializedChain(t *testing.T, client *utils.Client, relayers []types.AccountID, chains []msg.ChainId, resources map[msg.ResourceId]utils.Method, threshold uint32) {
	var count types.U32
	_, err := utils.QueryStorage(client, utils.BridgeStoragePrefix, "RelayerCount", nil, nil, &count)
	if err != nil {
		t.Fatal(err)
	}

	// Only perform setup if no relayers exist already (ie. state is uninitialized)
	if count == 0 {
		err = utils.InitializeChain(client, relayers, chains, resources, threshold)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println("======================== Substrate Initialized ========================")
		fmt.Printf("Relayers: 				%x\n", relayers)
		fmt.Printf("Whitelisted Chain IDs: 	%x\n", chains)
		fmt.Printf("Relayer Threshold: 		%x\n", threshold)
		for id, method := range resources {
			fmt.Printf("Resource - Id: 			%x, %s\n", id, method)
		}
		fmt.Println("========================================================================")

	} else {
		fmt.Println("=========================================================================")
		fmt.Println("! WARNING: Running tests against an initialized chain, results may vary !")
		fmt.Println("=========================================================================")
	}
}
