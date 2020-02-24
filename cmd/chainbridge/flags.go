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

// Account management flags
var (
	ImportFlag = cli.StringFlag{
		Name:  "import",
		Usage: "Import encrypted keystore file generated with bridge cli",
	}
	ListFlag = cli.BoolFlag{
		Name:  "list",
		Usage: "List node keys",
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
	Ed25519Flag = cli.BoolFlag{
		Name:  "ed25519",
		Usage: "Specify account type as ed25519",
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

// Test Setting Flags
var (
	TestKeyFlag = cli.StringFlag{
		Name:  "testkey",
		Usage: "Applies a predetermined test keystore to the chains",
	}
)
