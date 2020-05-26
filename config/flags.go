// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package config

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

	BlockstorePathFlag = cli.StringFlag{
		Name:  "blockstore",
		Usage: "Specify path for blockstore",
		Value: "", // Empty will use home dir
	}

	FreshStartFlag = cli.BoolFlag{
		Name:  "fresh",
		Usage: "Disables loading from blockstore at start. Opts will still be used if specified.",
	}
)

// Generate subcommand flags
var (
	PasswordFlag = cli.StringFlag{
		Name:  "password",
		Usage: "Password used to encrypt the keystore. Used with --generate, --import, or --unlock",
	}
	Sr25519Flag = cli.BoolFlag{
		Name:  "sr25519",
		Usage: "Specify account/key type as sr25519.",
	}
	Secp256k1Flag = cli.BoolFlag{
		Name:  "secp256k1",
		Usage: "Specify account/key type as secp256k1.",
	}
)

var (
	EthereumImportFlag = cli.BoolFlag{
		Name:  "ethereum",
		Usage: "Import an existing ethereum keystore, such as from geth.",
	}
	PrivateKeyFlag = cli.StringFlag{
		Name:  "privateKey",
		Usage: "Import a hex representation of a private key into a keystore.",
	}
)

// Test Setting Flags
var (
	TestKeyFlag = cli.StringFlag{
		Name:  "testkey",
		Usage: "Applies a predetermined test keystore to the chains.",
	}
)
