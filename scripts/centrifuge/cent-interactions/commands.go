package main

import (
	"encoding/hex"
	"fmt"

	"github.com/ChainSafe/log15"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/config"
	"github.com/centrifuge/go-substrate-rpc-client/types"
	"github.com/urfave/cli"
)

// parseHexArgument converts a hex string to an byte representation, ignoring 0x if present.
func parseHexArgument(arg string) []byte {
	// Ignore 0x prefix
	if len(arg) >= 2 && arg[0:2] == "0x" {
		arg = arg[2:]
	}

	// normalize number of characters
	if len(arg)%2 != 0 {
		arg = "0" + arg
	}

	if hx, err := hex.DecodeString(arg); err != nil {
		err = fmt.Errorf("failed to parse hex argument: %s", err.Error())
		panic(err)
	} else {
		return hx
	}
}

func setEmitterAddress(ctx *cli.Context) error {
	// Parse new address from CLI
	var addrStr string
	if addrStr = ctx.Args().Get(0); addrStr == "" {
		return fmt.Errorf("must provide a new address to use")
	}

	newAddr := parseHexArgument(addrStr)

	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		return err
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return err
	}

	c, err := types.NewCall(
		meta,
		"Bridge.set_address",
		newAddr,
	)

	if err != nil {
		log15.Error("Failed to create call", "err", err)
	}
	fmt.Printf("Submitting extrinsic to set emitter address to 0x%X\n", newAddr)
	return submitTx(api, c, meta)
}

func getEmitterAddress(ctx *cli.Context) error {
	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		panic(err)
	}

	key, err := types.CreateStorageKey(meta, "Bridge", "EmitterAddress", nil, nil)
	if err != nil {
		panic(err)
	}

	var addr []byte
	err = api.RPC.State.GetStorageLatest(key, &addr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Emitter Address: %X\n", addr)
	return nil
}

func whitelistChain(ctx *cli.Context) error {
	// Parse new address from CLI
	var idStr string
	if idStr = ctx.Args().Get(0); idStr == "" {
		return fmt.Errorf("must provide a chain id to use")
	}
	id := parseHexArgument(idStr)

	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		return err
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return err
	}

	c, err := types.NewCall(
		meta,
		"Bridge.whitelist_chain",
		id,
	)

	if err != nil {
		log15.Error("Failed to create call", "err", err)
	}
	fmt.Printf("Submitting extrinsic to whitelist chain 0x%X\n", id)
	return submitTx(api, c, meta)
}

func submitAssetTx(ctx *cli.Context) error {
	// Parse new CLi params
	var destIdStr, toStr, tokenIdStr, metadataStr string
	if destIdStr = ctx.Args().Get(0); destIdStr == "" {
		return fmt.Errorf("must provide a destination chain id to use")
	}
	if toStr = ctx.Args().Get(1); toStr == "" {
		return fmt.Errorf("must provide a to address to use")
	}
	if tokenIdStr = ctx.Args().Get(2); tokenIdStr == "" {
		return fmt.Errorf("must provide a token id id to use")
	}
	metadataStr = ctx.Args().Get(3)

	destId := parseHexArgument(destIdStr)

	to := parseHexArgument(toStr)

	tokenId := parseHexArgument(tokenIdStr)

	metadata := parseHexArgument(metadataStr)

	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		return err
	}

	// dest, to token_id, metadata
	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		panic(err)
	}

	c, err := types.NewCall(
		meta,
		"Bridge.transfer_asset",
		destId,
		to,
		tokenId,
		metadata,
	)

	if err != nil {
		return fmt.Errorf("failed to create call: %s", err)
	}

	return submitTx(api, c, meta)
}
