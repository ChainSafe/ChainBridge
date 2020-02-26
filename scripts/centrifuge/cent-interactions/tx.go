// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"fmt"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/rpc/author"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

func submitTx(api *gsrpc.SubstrateAPI, c types.Call, meta *types.Metadata) error {

	// Create the extrinsic
	ext := types.NewExtrinsic(c)

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return err
	}

	rv, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return err
	}

	key, err := types.CreateStorageKey(meta, "System", "AccountNonce", signature.TestKeyringPairAlice.PublicKey, nil)
	if err != nil {
		return err
	}

	var nonce uint32
	err = api.RPC.State.GetStorageLatest(key, &nonce)
	if err != nil {
		return err
	}

	o := types.SignatureOptions{
		BlockHash:   genesisHash,
		Era:         types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash: genesisHash,
		Nonce:       types.UCompact(nonce),
		SpecVersion: rv.SpecVersion,
		Tip:         0,
	}

	// Sign the transaction using Alice's default account
	err = ext.Sign(signature.TestKeyringPairAlice, o)
	if err != nil {
		return err
	}

	// Send the extrinsic
	sub, err := api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return err
	}

	return waitForExtrinsic(sub)
}

func waitForExtrinsic(sub *author.ExtrinsicStatusSubscription) error {
	for {
		select {
		case status := <-sub.Chan():
			if status.IsFinalized {
				fmt.Println("Extrinsic successful.")
				return nil
			}
		case err := <-sub.Err():
			return err
		}
	}

}
