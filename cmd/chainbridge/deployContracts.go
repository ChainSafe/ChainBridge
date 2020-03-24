// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"math/big"

	eth "github.com/ChainSafe/ChainBridgeV2/chains/ethereum"
	log "github.com/ChainSafe/log15"
	"github.com/urfave/cli"
)

var deployContractsCommand = cli.Command{
	Action:      parseCommands,
	Name:        "deploycontracts",
	Usage:       "deploys contracts",
	Category:    "tests",
	Flags:       deployContractsFlags,
	Description: "\tthe deploycontracts command is used to deploy contracts on a local network for testing purposes\n",
}

func parseCommands(ctx *cli.Context) error {
	log.Info("Deploying Contracts")

	port := ctx.String(PortFlag.Name)
	relayers := ctx.Int(NumRelayersFlag.Name)
	relayerThreshold := ctx.Int(RelayerThresholdFlag.Name)
	minCount := ctx.Int(MinCountFlag.Name)
	deployPK := ctx.String(PKFlag.Name)

	deployedContracts, err := eth.DeployContracts(deployPK, port, relayers, big.NewInt(int64(relayerThreshold)), uint8(minCount))
	if err != nil {
		return err
	}

	log.Info("Bridge Contract Deployed at: " + deployedContracts.BridgeAddress.Hex())
	log.Info("Relayer Contract Deployed at: " + deployedContracts.RelayerAddress.Hex())
	log.Info("ERC20 Handler Contract Deployed at: " + deployedContracts.ERC20HandlerAddress.Hex())
	log.Info("ERC721 Handler Contract Deployed at: " + deployedContracts.ERC721HandlerAddress.Hex())
	log.Info("Centrifuge Asset Handler Contract Deployed at: " + deployedContracts.CentrifugeHandlerAddress.Hex())

	return nil
}
