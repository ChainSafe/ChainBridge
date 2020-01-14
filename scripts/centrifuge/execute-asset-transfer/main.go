package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ChainSafe/log15"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/config"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type SubstrateProof struct {
	LeafHash     [32]byte
	SortedHashes [][32]byte
}

func main() {
	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		panic(err)
	}

	bob, err := types.NewAddressFromHexAccountID("0x8eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a48")
	if err != nil {
		panic(err)
	}

	// produce a random asset hash
	rand.Seed(time.Now().UnixNano())
	hsh, err := types.GetHash(rand.Float32())
	if err != nil {
		panic(err)
	}

	var to [20]byte
	copy(to[:], bob.AsAccountID[:20])
	var staticProofs [3][32]byte
	c, err := types.NewCall(
		meta,
		"Nfts.validate_mint",
		hsh,
		to,
		[]SubstrateProof{
			{
				LeafHash:     hsh,
				SortedHashes: [][32]byte{hsh},
			},
		},
		staticProofs)

	if err != nil {
		log15.Error("Failed to create call", "err", err)
	}
	// Create the extrinsic
	ext := types.NewExtrinsic(c)
	if err != nil {
		panic(err)
	}

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		panic(err)
	}

	rv, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		panic(err)
	}

	key, err := types.CreateStorageKey(meta, "System", "AccountNonce", signature.TestKeyringPairAlice.PublicKey, nil)
	if err != nil {
		panic(err)
	}

	var nonce uint32
	err = api.RPC.State.GetStorageLatest(key, &nonce)
	if err != nil {
		panic(err)
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
		panic(err)
	}

	// Send the extrinsic
	_, err = api.RPC.Author.SubmitExtrinsic(ext)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Transfer sent with NFT hash %#x\n", hsh)
}
