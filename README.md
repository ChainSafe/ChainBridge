# Forte Bridge Relayer
The Forte Bridge Relayer is a Multi-Directional Blockchain Bridge that interacts with
multiple chains; Ethereum, Ethereum Classic, Substrate, Custodial Aleo based chains.

## Software Setup
* [GO](https://golang.org/doc/install)
* [Docker](https://docs.docker.com/get-docker/)

## Local Development

### Local Environment 
Make a copy of `.env.example` and name it `.env` within the `ChainBridge` directory.

```bash
cp .env.example .env
```

Environment variables that need to come from Secrets are marked # TODO: Move to Secrets Storage


### Local Build and Setup

To build and set up the local environment, run the following commands.

```bash
make get
make build
make install
```

Once you have run the above command, you can run your local build of chainbridge.
However, additional configuration is required to deploy contracts and generate a configuration file.
See [forte_bridge_devops](https://github.com/fortelabsinc/forte_bridge_devops) for instructions on deploying contracts
and creating local configuration file.

### Ethereum Options

Ethereum chains support the following additional options:

```
{
    "bridge": "0x12345...",          // Address of the bridge contract (required)
    "erc20Handler": "0x1234...",     // Address of erc20 handler (required)
    "erc721Handler": "0x1234...",    // Address of erc721 handler (required)
    "genericHandler": "0x1234...",   // Address of generic handler (required)
    "maxGasPrice": "0x1234",         // Gas price for transactions (default: 20000000000)
    "minGasPrice": "0x1234",         // Minimum gas price for transactions (default: 0)
    "gasLimit": "0x1234",            // Gas limit for transactions (default: 6721975)
    "gasMultiplier": "1.25",         // Multiplies the gas price by the supplied value (default: 1)
    "http": "true",                  // Whether the chain connection is ws or http (default: false)
    "startBlock": "1234",            // The block to start processing events from (default: 0)
    "blockConfirmations": "10"       // Number of blocks to wait before processing a block
    "useExtendedCall": "true"        // Extend extrinsic calls to substrate with ResourceID. Used for backward compatibility with example pallet. *Default: false*
    "egsApiKey": "xxx..."            // API key for Eth Gas Station (https://www.ethgasstation.info/)
    "egsSpeed": "fast"               // Desired speed for gas price selection, the options are: "average", "fast", "fastest"
}
```

### Substrate Options

Substrate supports the following additonal options:

```
{
    "startBlock": "1234" // The block to start processing events from (default: 0)
}
```

### ChainBrdige Commands

* Show a list of useful commands

  ```bash
  make help
  # Alternative
  make
  ```

* Generate Keystore files Using a Private Key

    *Generate needed keystore files for the relayer contracts Private Keys*
    ```bash
      ./build/chainbridge accounts import --privateKey "${Your Private Key Here}"
    ```

This command requires a properly configured json file in `./config/config.json` and a configured keystore in
`./keys`. Make sure you have moved the files you generated in [forte_bridge_devops](https://github.com/fortelabsinc/forte_bridge_devops) 
to the local directory for running the relayer here. See [forte_bridge_devops](https://github.com/fortelabsinc/forte_bridge_devops) 
for instructions on deploying contracts and creating the configuration file.

If using the Aleo Custodian. Make sure it is running on Localhost and the ports specified in the `config.json`
file.

* Run ChainBridge Relayer

  *Run ChainBridge Relayer*
  ```bash
  make start
  # Alternative
  ./build/chainbridge --config config/config.json --verbosity trace --latest
  ```

  

