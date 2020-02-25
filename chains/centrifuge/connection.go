package substrate

import (
	"fmt"
	"time"

	"github.com/ChainSafe/log15"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/rpc/author"
	"github.com/centrifuge/go-substrate-rpc-client/rpc/state"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

//var _ chains.Connection = &Connection{}

type Connection struct {
	api *gsrpc.SubstrateAPI
	url string
	meta *types.Metadata
	genesisHash types.Hash
	key signature.KeyringPair
}

func NewConnection(url string) *Connection {
	return &Connection{url: url}
}

func (c *Connection) Connect() error {
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
	c.meta = meta
	log15.Debug("Fetched substrate metadata")

	// Fetch genesis hash
	genesisHash, err := c.api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return err
	}
	c.genesisHash = genesisHash
	log15.Debug("Fetched substrate genesis hash", "hash", genesisHash)
	return nil
}

func (c *Connection) SubmitTx(method Method, args ...interface{}) error {
	log15.Debug("Submitting substrate call...", "method", method)
	// Create call
	call, err := types.NewCall(c.meta, method.String(), args)
	if err != nil {
		return err
	}
	// Create the extrinsic
	ext := types.NewExtrinsic(call)

	rv, err := c.api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return err
	}

	key, err := types.CreateStorageKey(c.meta, "System", "AccountNonce", c.key.PublicKey, nil)
	if err != nil {
		return err
	}

	var nonce uint32
	err = c.api.RPC.State.GetStorageLatest(key, &nonce)
	if err != nil {
		return err
	}

	o := types.SignatureOptions{
		BlockHash:   c.genesisHash,
		Era:         types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash: c.genesisHash,
		Nonce:       types.UCompact(nonce),
		SpecVersion: rv.SpecVersion,
		Tip:         0,
	}

	// Sign the transaction using Alice's default account
	err = ext.Sign(c.key, o)
	log15.Debug("Ext complete", "signed", ext.IsSigned(), "method", ext.Method)
	if err != nil {
		return err
	}

	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	log15.Debug("Extrinsic submission succeeded")
	time.Sleep(time.Second * 1)
	return watchSubmission(sub)
}

func watchSubmission(sub *author.ExtrinsicStatusSubscription) error {
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


// Subscribe creates a subscription to all events
func (c *Connection) Subscribe() (*state.StorageSubscription, error) {
	meta, err := c.api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, err
	}
	key, err := types.CreateStorageKey(meta, "System", "Events", nil, nil)
	if err != nil {
		return nil, err
	}

	sub, err := c.api.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (c *Connection) Close() {
	// Do nothing for now
}
