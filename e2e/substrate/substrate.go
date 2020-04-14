// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ChainSafe/ChainBridge/chains/substrate"
	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

const TestSubEndpoint = "ws://localhost:9944"

var log = log15.New("e2e", "substrate")

var AliceKp = keystore.TestKeyRing.SubstrateKeys[keystore.AliceKey]
var BobKp = keystore.TestKeyRing.SubstrateKeys[keystore.BobKey]

func CreateConfig(key string, chain msg.ChainId) *core.ChainConfig {
	return &core.ChainConfig{
		Name:         fmt.Sprintf("substrate(%s)", key),
		Id:           chain,
		Endpoint:     TestSubEndpoint,
		From:         "",
		KeystorePath: key,
		Insecure:     true,
		Opts:         map[string]string{},
	}
}

type subClient struct {
	api     *gsrpc.SubstrateAPI
	meta    *types.Metadata
	genesis types.Hash
	key     *signature.KeyringPair
}

func CreateSubClient(t *testing.T, key *signature.KeyringPair) *subClient {
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

func WatchForProposalSuccessOrFail(client *subClient, expectedNonce types.U64, success chan bool, fail chan error) {
	key, err := types.CreateStorageKey(client.meta, "System", "Events", nil, nil)
	if err != nil {
		fail <- err
		return
	}

	sub, err := client.api.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
	if err != nil {
		fail <- err
		return
	}

	for {
		set := <-sub.Chan()
		for _, chng := range set.Changes {
			if !types.Eq(chng.StorageKey, key) || !chng.HasStorageData {
				// skip, we are only interested in events with content
				continue
			}

			// Decode the event records
			events := substrate.Events{}
			err = types.EventRecordsRaw(chng.StorageData).DecodeEventRecords(client.meta, &events)
			if err != nil {
				fail <- err
				return
			}

			for _, evt := range events.Bridge_ProposalSucceeded {
				if evt.DepositNonce == expectedNonce {
					success <- true
					return
				} else {
					log15.Info("Found mismatched event", "depositNonce", evt.DepositNonce)
				}
			}

			for _, evt := range events.Bridge_ProposalFailed {
				if evt.DepositNonce == expectedNonce {
					fail <- errors.New("proposal failed")
					return
				}
			}
		}
	}
}

func SubmitTx(t *testing.T, client *subClient, method substrate.Method, args ...interface{}) {
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
			log.Info("Extrinsic in block", "block", status.AsInBlock.Hex())
			return
		case status.IsDropped:
			t.Fatal("Extrinsic dropped")
		case status.IsInvalid:
			t.Fatal("Extrinsic invalid")
		}
	}
}

func SubmitSudoTx(t *testing.T, client *subClient, method substrate.Method, args ...interface{}) {
	call, err := types.NewCall(client.meta, method.String(), args...)
	if err != nil {
		t.Fatal(err)
	}
	SubmitTx(t, client, substrate.Sudo, call)
}

// queryStorage performs a storage lookup. Arguments may be nil, result must be a pointer.
func queryStorage(client *subClient, prefix, method string, arg1, arg2 []byte, result interface{}) (bool, error) {
	key, err := types.CreateStorageKey(client.meta, prefix, method, arg1, arg2)
	if err != nil {
		return false, err
	}
	return client.api.RPC.State.GetStorageLatest(key, result)
}

func AddRelayer(t *testing.T, client *subClient, relayer types.AccountID) {
	log.Info("Adding relayer", "accountId", relayer)
	SubmitSudoTx(t, client, substrate.AddRelayer, relayer)
}

func WhitelistChain(t *testing.T, client *subClient, id msg.ChainId) {
	log.Info("Whitelisting chain", "chainId", id)
	SubmitSudoTx(t, client, substrate.WhitelistChain, types.U8(id))
}

func InitiateHashTransfer(t *testing.T, client *subClient, hash types.Hash, destId msg.ChainId) {
	log.Info("Initiating hash transfer", "hash", hash.Hex())
	SubmitTx(t, client, substrate.ExampleTransferHash, hash, types.U8(destId))
}

func InitiateSubstrateNativeTransfer(t *testing.T, client *subClient, amount int, recipient []byte, destId msg.ChainId) { //nolint:unused,deadcode
	log.Info("Initiating Substrate native transfer", "amount", amount, "recipient", recipient, "destId", destId)
	SubmitTx(t, client, substrate.ExampleTransfer, amount, recipient, types.U8(destId))
}

func HashInt(i int) types.Hash {
	hash, err := types.GetHash(types.NewI64(int64(i)))
	if err != nil {
		panic(err)
	}
	return hash
}

func RegisterResource(t *testing.T, client *subClient, id msg.ResourceId, method string) {
	log.Info("Registering resource", "rId", id, "method", []byte(method))
	SubmitSudoTx(t, client, substrate.SetResource, types.NewBytes32(id), []byte(method))
}
