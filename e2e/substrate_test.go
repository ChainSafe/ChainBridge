// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package e2e

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ChainSafe/ChainBridge/chains/substrate"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type subClient struct {
	api     *gsrpc.SubstrateAPI
	meta    *types.Metadata
	genesis types.Hash
	key     *signature.KeyringPair
}

func createSubClient(t *testing.T, key *signature.KeyringPair) *subClient {
	c := &subClient{key: key}
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

func watchForProposalSuccessOrFail(t *testing.T, client *subClient, success chan bool, fail chan error) {
	key, err := types.CreateStorageKey(client.meta, "System", "Events", nil, nil)
	if err != nil {
		fail <- err
		close(success)
		close(fail)
		return
	}

	sub, err := client.api.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
	if err != nil {
		fail <- err
		close(success)
		close(fail)
		return
	}

	for {
		set := <-sub.Chan()
		for _, chng := range set.Changes {
			if !types.Eq(chng.StorageKey, key) || !chng.HasStorageData {
				// skip, we are only interested in events with countent
				continue
			}

			// Decode the event records
			events := substrate.Events{}
			err = types.EventRecordsRaw(chng.StorageData).DecodeEventRecords(client.meta, &events)
			if err != nil {
				fail <- err
				close(success)
				close(fail)
				return
			}

			for range events.Bridge_ProposalSucceeded {
				success <- true
			}

			for range events.Bridge_ProposalFailed {
				fail <- errors.New("Proposal failed")
			}
		}
	}
}

func submitTx(t *testing.T, client *subClient, method substrate.Method, args ...interface{}) {
	log15.Info("Submiting test tx", "method", method)
	// Create call and extrinsic
	call, err := types.NewCall(
		client.meta,
		method.String(),
		args...,
	)
	if err != nil {
		t.Fatal(err)
	}
	ext := types.NewExtrinsic(call)

	// Get latest runtime version
	rv, err := client.api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		t.Fatal(err)
	}

	var acct substrate.AccountData
	_, err = queryStorage(client, "System", "Account", client.key.PublicKey, nil, &acct)
	if err != nil {
		t.Fatal(err)
	}

	// Sign the extrinsic
	o := types.SignatureOptions{
		BlockHash:   client.genesis,
		Era:         types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash: client.genesis,
		Nonce:       types.UCompact(acct.Nonce),
		SpecVersion: rv.SpecVersion,
		Tip:         0,
	}
	err = ext.Sign(*client.key, o)
	if err != nil {
		t.Fatal(err)
	}

	// Submit and watch the extrinsic
	sub, err := client.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		t.Fatal(err)
	}

	for {
		status := <-sub.Chan()
		switch {
		case status.IsInBlock:
			fmt.Printf("Completed at block hash: %#x\n", status.AsInBlock)
			return
		case status.IsDropped:
			t.Fatal("Extrinsic dropped")
		case status.IsInvalid:
			t.Fatal("Extrinsic invalid")
		}
	}
}

func submitSudoTx(t *testing.T, client *subClient, method substrate.Method, args ...interface{}) {
	call, err := types.NewCall(client.meta, method.String(), args...)
	if err != nil {
		t.Fatal(err)
	}
	submitTx(t, client, substrate.Sudo, call)
}

// queryStorage performs a storage lookup. Arguments may be nil, result must be a pointer.
func queryStorage(client *subClient, prefix, method string, arg1, arg2 []byte, result interface{}) (bool, error) {
	key, err := types.CreateStorageKey(client.meta, prefix, method, arg1, arg2)
	if err != nil {
		return false, err
	}
	return client.api.RPC.State.GetStorageLatest(key, result)
}

func whitelistChain(t *testing.T, client *subClient, id msg.ChainId) {
	submitSudoTx(t, client, substrate.WhitelistChain, types.U32(id))
}

func initiateHashTransfer(t *testing.T, client *subClient, hash types.Hash, destId msg.ChainId) {
	recipient := types.Bytes{}
	submitTx(t, client, substrate.ExampleTransferHash, hash, recipient, types.U32(destId))
}
