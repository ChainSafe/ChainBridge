// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"
	"sync"

	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/rpc/author"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type Connection struct {
	api         *gsrpc.SubstrateAPI
	log         log15.Logger
	url         string                 // API endpoint
	name        string                 // Chain name
	meta        types.Metadata         // Latest chain metadata
	metaLock    sync.RWMutex           // Lock metadata for updates, allows concurrent reads
	genesisHash types.Hash             // Chain genesis hash
	key         *signature.KeyringPair // Keyring used for signing
	nonce       types.U32              // Latest account nonce
	nonceLock   sync.Mutex             // Locks nonce for updates
	stop        <-chan int             // Signals system shutdown, should be observed in all selects and loops
	sysErr      chan<- error           // Propagates fatal errors to core
}

func NewConnection(url string, name string, key *signature.KeyringPair, log log15.Logger, stop <-chan int, sysErr chan<- error) *Connection {
	return &Connection{url: url, name: name, key: key, log: log, stop: stop, sysErr: sysErr}
}

func (c *Connection) getMetadata() (meta types.Metadata) {
	c.metaLock.RLock()
	meta = c.meta
	c.metaLock.RUnlock()
	return meta
}

func (c *Connection) updateMetatdata() error {
	c.metaLock.Lock()
	meta, err := c.api.RPC.State.GetMetadataLatest()
	if err != nil {
		c.metaLock.Unlock()
		return err
	}
	c.meta = *meta
	c.metaLock.Unlock()
	return nil
}

func (c *Connection) Connect() error {
	c.log.Info("Connecting to substrate chain...", "url", c.url)
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
	c.log.Debug("Fetched substrate metadata")

	// Fetch genesis hash
	genesisHash, err := c.api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return err
	}
	c.genesisHash = genesisHash
	c.log.Debug("Fetched substrate genesis hash", "hash", genesisHash.Hex())
	return nil
}

// SubmitTx constructs and submits an extrinsic to call the method with the given arguments.
// All args are passed directly into GSRPC. GSRPC types are recommended to avoid serialization inconsistencies.
func (c *Connection) SubmitTx(method utils.Method, args ...interface{}) error {
	c.log.Debug("Submitting substrate call...", "method", method, "sender", c.key.Address)

	meta := c.getMetadata()

	// Create call and extrinsic
	call, err := types.NewCall(
		&meta,
		string(method),
		args...,
	)
	if err != nil {
		return fmt.Errorf("failed to construct call: %w", err)
	}
	ext := types.NewExtrinsic(call)

	// Get latest runtime version
	rv, err := c.api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return err
	}

	c.nonceLock.Lock()
	latestNonce, err := c.getLatestNonce()
	if err != nil {
		c.nonceLock.Unlock()
		return err
	}
	if latestNonce > c.nonce {
		c.nonce = latestNonce
	}

	// Sign the extrinsic
	o := types.SignatureOptions{
		BlockHash:          c.genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        c.genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(c.nonce)),
		SpecVersion:        rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: rv.TransactionVersion,
	}

	err = ext.Sign(*c.key, o)
	if err != nil {
		c.nonceLock.Unlock()
		return err
	}

	// Submit and watch the extrinsic
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	c.nonce++
	c.nonceLock.Unlock()
	if err != nil {
		return fmt.Errorf("submission of extrinsic failed: %w", err)
	}
	c.log.Trace("Extrinsic submission succeeded")
	defer sub.Unsubscribe()

	return c.watchSubmission(sub)
}

func (c *Connection) watchSubmission(sub *author.ExtrinsicStatusSubscription) error {
	for {
		select {
		case <-c.stop:
			return TerminatedError
		case status := <-sub.Chan():
			switch {
			case status.IsInBlock:
				c.log.Trace("Extrinsic included in block", "block", status.AsInBlock.Hex())
				return nil
			case status.IsRetracted:
				return fmt.Errorf("extrinsic retracted: %s", status.AsRetracted.Hex())
			case status.IsDropped:
				return fmt.Errorf("extrinsic dropped from network")
			case status.IsInvalid:
				return fmt.Errorf("extrinsic invalid")
			}
		case err := <-sub.Err():
			c.log.Trace("Extrinsic subscription error", "err", err)
			return err
		}
	}
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

// TODO: Add this to GSRPC
func getConst(meta *types.Metadata, prefix, name string, res interface{}) error {
	for _, mod := range meta.AsMetadataV12.Modules {
		if string(mod.Name) == prefix {
			for _, cons := range mod.Constants {
				if string(cons.Name) == name {
					return types.DecodeFromBytes(cons.Value, res)
				}
			}
		}
	}
	return fmt.Errorf("could not find constant %s.%s", prefix, name)
}

func (c *Connection) getConst(prefix, name string, res interface{}) error {
	meta := c.getMetadata()
	return getConst(&meta, prefix, name, res)
}

func (c *Connection) checkChainId(expected msg.ChainId) error {
	var actual msg.ChainId
	err := c.getConst(utils.BridgePalletName, "ChainIdentity", &actual)
	if err != nil {
		return err
	}

	if actual != expected {
		return fmt.Errorf("ChainID is incorrect, Expected chainId: %d, got chainId: %d", expected, actual)
	}

	return nil
}

func (c *Connection) getLatestNonce() (types.U32, error) {
	var acct types.AccountInfo
	exists, err := c.queryStorage("System", "Account", c.key.PublicKey, nil, &acct)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, nil
	}

	return acct.Nonce, nil
}
func (c *Connection) Close() {
	// TODO: Anything required to shutdown GRPC?
}
