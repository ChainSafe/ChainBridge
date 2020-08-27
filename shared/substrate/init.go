// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

func InitializeChain(client *Client, relayers []types.AccountID, chains []msg.ChainId, resources map[msg.ResourceId]Method, threshold uint32) error {
	calls := []types.Call{}

	// Create AddRelayer calls
	for _, relayer := range relayers {
		call, err := client.NewAddRelayerCall(relayer)
		if err != nil {
			return err
		}
		calls = append(calls, call)
	}
	// Create WhitelistChain calls
	for _, chain := range chains {
		call, err := client.NewWhitelistChainCall(chain)
		if err != nil {
			return err
		}
		calls = append(calls, call)
	}

	// Create SetResource calls
	for id, method := range resources {
		call, err := client.NewRegisterResourceCall(id, string(method))
		if err != nil {
			return err
		}
		calls = append(calls, call)
	}

	// Create a SetThreshold call
	call, err := client.NewSetRelayerThresholdCall(types.U32(threshold))
	if err != nil {
		return err
	}
	calls = append(calls, call)

	return BatchSubmit(client, calls)
}
