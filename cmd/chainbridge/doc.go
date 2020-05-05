/*
Package main (under cmd/chainbridge) provide command-line interface functionality for the chainbridge application.

Calling the Bridge

Running the Bridge can be called with the `chainbridge command.`
For more information on configuring the chain, refer to: https://pkg.go.dev/mod/github.com/ChainSafe/ChainBridge#configuration
For more information on running the chain, refer to [TO BE WRITTEN]

Managing Accounts

In the context of the bridge, managing accounts means generating and importing keystore files for the bridge.

To generate new keystores, call `chainbridge accounts generate`. The supported types are secp256k1 and sr25519, and can be specfied with flags (--secp256k1, etc).
To pass in a password, the `password` field should be called. If unsupplied, the generate function will ask for a password via the CLI.

To list the currenty keystore, call `chainbridge accounts list`. This will give a list of all the keystore files, listed by filename.

Importing keys

To import keys, the base command `chainbridge accounts import`

If importing a chainbridge keystore from another directory, a path to the keystore must be provided.

If importing a private key, the `--privateKey` flag must be provided. The private key must be a hex representation of the key, but the `0x` prefix is optional.
To pass in a password, the `password` field should be called. If unsupplied, the generate function will ask for a password via the CLI.
Keytype flags can be passed in to the import, and the supported types are secp256k1 and sr25519. If no key type is provided, the key is assumed to be secp256k1.

If importing an ethereum key (such as a geth key), the --ethereum flag must be used, and then the path to the key.
To pass in a password, the `password` field should be called. If unsupplied, the generate function will ask for a password via the CLI.
A new password will be needed for the new chainbridge keystore, and that must be provided via the CLI.
Only Ethereum format keys are supported at this time.

Config

The configs section of main support converting the ChainConfig structs to TOML and vice versa. For more information on configuration, refer to https://pkg.go.dev/mod/github.com/ChainSafe/ChainBridge#configuration.

*/
package main
