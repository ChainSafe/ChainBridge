// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"
	"sync"

	"github.com/ChainSafe/log15"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/rpc/author"
	"github.com/centrifuge/go-substrate-rpc-client/rpc/state"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type Connection struct {
	api         *gsrpc.SubstrateAPI
	url         string
	name        string
	meta        types.Metadata
	genesisHash types.Hash
	key         *signature.KeyringPair
	metaLock    sync.RWMutex
}

func NewConnection(url string, name string, key *signature.KeyringPair) *Connection {
	return &Connection{url: url, name: name, key: key}
}

func (c *Connection) getMetadata() (meta types.Metadata) {
	c.metaLock.RLock()
	meta = c.meta
	c.metaLock.RUnlock()
	return
}

func (c *Connection) Connect() error {
	log15.Info("Connecting to substrate chain...", "chain", c.name)
	api, err := gsrpc.NewSubstrateAPI(c.url)
	if err != nil {
		return err
	}
	c.api = api

	// Fetch metadata
	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return err
	}
	c.meta = *meta
	log15.Debug("Fetched substrate metadata")

	// Fetch genesis hash
	genesisHash, err := c.api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return err
	}
	c.genesisHash = genesisHash
	log15.Debug("Fetched substrate genesis hash", "hash", genesisHash.Hex())
	return nil
}

// SubmitTx constructs and submits an extrinsic to call the method with the given arguments.
// All args are passed directly into GSRPC. GSRPC types are recommended to avoid serialization inconsistencies.
func (c *Connection) SubmitTx(method Method, args ...interface{}) error {
	log15.Debug("Submitting substrate call...", "method", method, "sender", c.key.Address)

	meta := c.getMetadata()

	// Create call and extrinsic
	call, err := types.NewCall(
		&meta,
		method.String(),
		args...,
	)
	if err != nil {
		return fmt.Errorf("failed to construct call: %s", err.Error())
	}
	ext := types.NewExtrinsic(call)

	// Get latest runtime version
	rv, err := c.api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return err
	}

	// TODO: Does this actually return more than nonce?
	var nonce uint32
	_, err = c.queryStorage("System", "Account", c.key.PublicKey, nil, &nonce)
	if err != nil {
		return err
	}

	// Sign the extrinsic
	o := types.SignatureOptions{
		BlockHash:   c.genesisHash,
		Era:         types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash: c.genesisHash,
		Nonce:       types.UCompact(nonce),
		SpecVersion: rv.SpecVersion,
		Tip:         0,
	}
	err = ext.Sign(*c.key, o)
	if err != nil {
		return err
	}

	// Submit and watch the extrinsic
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return fmt.Errorf("submission of extrinsic failed: %s", err.Error())
	}
	log15.Debug("Extrinsic submission succeeded")
	defer sub.Unsubscribe()
	return watchSubmission(sub)
}

func watchSubmission(sub *author.ExtrinsicStatusSubscription) error {
	for {
		select {
		case status := <-sub.Chan():
			switch {
			case status.IsInBlock:
				log15.Trace("Successful extrinsic finalized", "block", status.AsInBlock.Hex())
				return nil
			case status.IsRetracted:
				return fmt.Errorf("extrinsic retracted: %s", status.AsRetracted.Hex())
			case status.IsDropped:
				return fmt.Errorf("extrinsic dropped from network")
			case status.IsInvalid:
				return fmt.Errorf("extrinsic invalid")
			default:
				log15.Trace("Other status", "status", fmt.Sprintf("%+v", status))
			}
		case err := <-sub.Err():
			log15.Trace("Extrinsic subscription error", "err", err)
			return err
		}
	}
}

// Subscribe creates a subscription to all events.
func (c *Connection) Subscribe() (*state.StorageSubscription, error) {
	meta := c.getMetadata()
	key, err := types.CreateStorageKey(&meta, "System", "Events", nil, nil)
	if err != nil {
		return nil, err
	}

	sub, err := c.api.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
	if err != nil {
		return nil, err
	}
	return sub, nil
}

// queryStorage performs a storage lookup. Arguments may be nil, result must be a pointer.
func (c *Connection) queryStorage(prefix, method string, arg1, arg2 []byte, result interface{}) (bool, error) {
	// Fetch account nonce
	data := c.getMetadata()
	key, err := types.CreateStorageKey(&data, prefix, method, arg1, arg2)
	if err != nil {
		return false, err
	}
	return c.api.RPC.State.GetStorageLatest(key, result)
}

func (c *Connection) Close() {
	// TODO: Anything required to shutdown GRPC?
}
