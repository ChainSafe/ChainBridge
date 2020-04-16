// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package subtest

import (
	"fmt"

	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

// WARNING: THIS METHOD IS UNSAFE AND MAY PANIC
func EnsureInitializedChain(client *utils.Client, relayers []types.AccountID, chains []msg.ChainId, threshold uint32) {
	var count types.U32
	_, err := utils.QueryStorage(client, "Bridge", "RelayerCount", nil, nil, &count)
	if err != nil {
		panic(err)
	}

	// Only perform setup if no relayers exist already (ie. state is uninitialized)
	if count == 0 {
		for _, relayer := range relayers {
			err = client.AddRelayer(relayer)
			if err != nil {
				panic(err)
			}
		}

		for _, chain := range chains {
			err = client.WhitelistChain(chain)
			if err != nil {
				panic(err)
			}
		}

		err = client.SetRelayerThreshold(types.U32(threshold))
		if err != nil {
			panic(err)
		}

	} else {
		fmt.Println("=========================================================================")
		fmt.Println("! WARNING: Running tests against an initialized chain, results may vary !")
		fmt.Println("=========================================================================")
	}
}
