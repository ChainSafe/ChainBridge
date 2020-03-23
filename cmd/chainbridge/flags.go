// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	log "github.com/ChainSafe/log15"
	"github.com/urfave/cli"
)

var (
	ConfigFileFlag = cli.StringFlag{
		Name:  "config",
		Usage: "TOML configuration file",
	}

	VerbosityFlag = cli.StringFlag{
		Name:  "verbosity",
		Usage: "Supports levels crit (silent) to trce (trace)",
		Value: log.LvlInfo.String(),
	}

	KeystorePathFlag = cli.StringFlag{
		Name:  "keystore",
		Usage: "Path to keystore directory",
		Value: DefaultKeystorePath,
	}
)

// Generate subcommand flags
var (
	PrivateKeyFlag = cli.StringFlag{
		Name:  "privateKey",
		Usage: "Hex string private key used to generate a keypair.",
	}
	PasswordFlag = cli.StringFlag{
		Name:  "password",
		Usage: "Password used to encrypt the keystore. Used with --generate or --unlock",
	}
	Sr25519Flag = cli.BoolFlag{
		Name:  "sr25519",
		Usage: "Specify account type as sr25519",
	}
	Secp256k1Flag = cli.BoolFlag{
		Name:  "secp256k1",
		Usage: "Specify account type as secp256k1",
	}
)

// Deploy Contracts Flags
var (
	PortFlag = cli.StringFlag{
		Name:  "port",
		Usage: "The port at which your local chain instance is running",
		Value: "8545",
	}
	NumRelayersFlag = cli.IntFlag{
		Name:  "validators",
		Usage: "Specify total number of validators",
		Value: 2,
	}
	RelayerThresholdFlag = cli.IntFlag{
		Name:  "validatorthreshold",
		Usage: "Specify validator threshold",
		Value: 2,
	}
	PKFlag = cli.StringFlag{
		Name:  "pk",
		Usage: "Specify private key of account you wish to deploy contracts from",
		Value: "000000000000000000000000000000000000000000000000000000416c696365",
	}
	MinCountFlag = cli.IntFlag{
		Name:  "mincount",
		Usage: "Specify min count for bridge asset",
		Value: 10,
	}
)

// Test Setting Flags
var (
	TestKeyFlag = cli.StringFlag{
		Name:  "testkey",
		Usage: "Applies a predetermined test keystore to the chains",
	}
)
