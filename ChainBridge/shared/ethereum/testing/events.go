// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethtest

import (
	"context"
	"fmt"
	"math/big"

	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func WatchEvent(client *utils.Client, bridge common.Address, subStr utils.EventSig) {
	fmt.Printf("Watching for event: %s\n", subStr)
	query := eth.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{bridge},
		Topics: [][]common.Hash{
			{subStr.GetTopic()},
		},
	}

	ch := make(chan ethtypes.Log)
	sub, err := client.Client.SubscribeFilterLogs(context.Background(), query, ch)
	if err != nil {
		log15.Error("Failed to subscribe to event", "event", subStr)
		return
	}
	defer sub.Unsubscribe()

	for {
		select {
		case evt := <-ch:
			fmt.Printf("%s (block: %d): %#v\n", subStr, evt.BlockNumber, evt.Topics)

		case err := <-sub.Err():
			if err != nil {
				log15.Error("Subscription error", "event", subStr, "err", err)
				return
			}
		}
	}
}
