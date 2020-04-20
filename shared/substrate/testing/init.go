// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package subtest

import (
	"fmt"
	"testing"

	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type Resource struct {
	Id     msg.ResourceId
	Method utils.Method
}

// WARNING: THIS METHOD IS UNSAFE AND MAY PANIC
func EnsureInitializedChain(t *testing.T, client *utils.Client, relayers []types.AccountID, chains []msg.ChainId, resources []Resource, threshold uint32) {
	var count types.U32
	_, err := utils.QueryStorage(client, "Bridge", "RelayerCount", nil, nil, &count)
	if err != nil {
		t.Fatal(err)
	}

	// Only perform setup if no relayers exist already (ie. state is uninitialized)
	if count == 0 {
		calls := []types.Call{}

		// Create AddRelayer calls
		for _, relayer := range relayers {
			call, err := client.NewAddRelayerCall(relayer)
			if err != nil {
				t.Fatal(err)
			}
			calls = append(calls, call)
		}
		// Create WhitelistChain calls
		for _, chain := range chains {
			call, err := client.NewWhitelistChainCall(chain)
			if err != nil {
				t.Fatal(err)
			}
			calls = append(calls, call)
		}

		// Create SetResource calls
		for _, r := range resources {
			call, err := client.NewRegisterResourceCall(r.Id, string(r.Method))
			if err != nil {
				t.Fatal(err)
			}
			calls = append(calls, call)
		}

		// Create a SetThreshold call
		call, err := client.NewSetRelayerThresholdCall(types.U32(threshold))
		if err != nil {
			t.Fatal(err)
		}
		calls = append(calls, call)

		err = utils.BatchSubmit(client, calls)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println("======================== Substrate Initialized ========================")
		fmt.Printf("Relayers: 				%x\n", relayers)
		fmt.Printf("Whitelisted Chain IDs: 	%x\n", chains)
		fmt.Printf("Relayer Threshold: 		%x\n", threshold)
		for _, r := range resources {
			fmt.Printf("Resource - Id: 			%x, %s\n", r.Id, r.Method)
		}
		fmt.Println("========================================================================")

	} else {
		fmt.Println("=========================================================================")
		fmt.Println("! WARNING: Running tests against an initialized chain, results may vary !")
		fmt.Println("=========================================================================")
	}
}
