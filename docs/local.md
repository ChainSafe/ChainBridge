# Running Locally
## Prerequisites

- Docker, docker-compose
- chainbridge `v1.0.0` binary (see [README](https://github.com/chainsafe/chainbridge/#building))
- Polkadot JS Portal ([https://portal.chain.centrifuge.io](https://portal.chain.centrifuge.io/))
    - Connect to your local node (below) by clicking in the top-left corner and using `ws://localhost:9944`
    - You will need to setup the type definitions for the chain by selecting `Settings` -> `Developer`
    - Type definitions can be found here: [PolkadotJS Apps](https://github.com/ChainSafe/chainbridge-substrate-chain#polkadot-js-apps)
- cb-sol-cli (see [README](https://github.com/ChainSafe/chainbridge-deploy/tree/master/cb-sol-cli#cb-sol-cli-documentation))

## Starting Local Chains

The easiest way to get started is to use the docker-compose file found here: [https://gist.github.com/ansermino/f1571bb354d2007b26dce53d52dbca75](https://gist.github.com/ansermino/f1571bb354d2007b26dce53d52dbca75). This will start a geth instance and an instance of chainbridge-substrate-chain:

```bash
docker-compose -f ./docker-compose-geth-substrate.yml up -V
```

(Use `-V` to always start with new chains. These instructions depend on deterministic ethereum addresses, which are used as defaults implicitly by some of these commands. Avoid re-deploying the contracts without restarting both chains, or ensure to specify all the required parameters.)

## On-Chain Setup (Ethereum)

### Deploy Contracts

To deploy the contracts on to the Ethereum chain, run the following:

```bash
cb-sol-cli deploy --all --relayerThreshold 1
```

After running, the expected output looks like this:

```bash
================================================================
Url:        http://localhost:8545
Deployer:   0xff93B45308FD417dF303D6515aB04D9e89a750Ca
Gas Limit:   8000000
Gas Price:   20000000
Deploy Cost: 0.0

Options
=======
Chain Id:    0
Threshold:   2
Relayers:    0xff93B45308FD417dF303D6515aB04D9e89a750Ca,0x8e0a907331554AF72563Bd8D43051C2E64Be5d35,0x24962717f8fA5BA3b931bACaF9ac03924EB475a0,0x148FfB2074A9e59eD58142822b3eB3fcBffb0cd7,0x4CEEf6139f00F9F4535Ad19640Ff7A0137708485
Bridge Fee:  0
Expiry:      100

Contract Addresses
================================================================
Bridge:             0x62877dDCd49aD22f5eDfc6ac108e9a4b5D2bD88B
----------------------------------------------------------------
Erc20 Handler:      0x3167776db165D8eA0f51790CA2bbf44Db5105ADF
----------------------------------------------------------------
Erc721 Handler:     0x3f709398808af36ADBA86ACC617FeB7F5B7B193E
----------------------------------------------------------------
Generic Handler:    0x2B6Ab4b880A45a07d83Cf4d664Df4Ab85705Bc07
----------------------------------------------------------------
Erc20:              0x21605f71845f372A9ed84253d2D024B7B10999f4
----------------------------------------------------------------
Erc721:             0xd7E33e1bbf65dC001A0Eb1552613106CD7e40C31
----------------------------------------------------------------
Centrifuge Asset:   0xc279648CE5cAa25B9bA753dAb0Dfef44A069BaF4
================================================================
```

### Register Resources

```bash
# Register fungible resource ID with erc20 contract
cb-sol-cli bridge register-resource --resourceId "0x000000000000000000000000000000c76ebe4a02bbc34786d860b355f5a5ce00" --targetContract "0x21605f71845f372A9ed84253d2D024B7B10999f4"

# Register non-fungible resource ID with erc721 contract
cb-sol-cli bridge register-resource --resourceId "0x000000000000000000000000000000e389d61c11e5fe32ec1735b3cd38c69501" --targetContract "0xd7E33e1bbf65dC001A0Eb1552613106CD7e40C31" --handler "0x3f709398808af36ADBA86ACC617FeB7F5B7B193E"

# Register generic resource ID
cb-sol-cli bridge register-generic-resource --resourceId "0x000000000000000000000000000000f44be64d2de895454c3467021928e55e01" --targetContract "0xc279648CE5cAa25B9bA753dAb0Dfef44A069BaF4" --handler "0x2B6Ab4b880A45a07d83Cf4d664Df4Ab85705Bc07" --hash --deposit "" --execute "store(bytes32)"
```

### Specify Token Semantics

To allow for a variety of use cases, the Ethereum contracts support both the `transfer` and the `mint/burn` ERC methods.

For simplicity's sake the following examples only make use of the  `mint/burn` method:

```bash
# Register the erc20 contract as mintable/burnable
cb-sol-cli bridge set-burn --tokenContract "0x21605f71845f372A9ed84253d2D024B7B10999f4"

# Register the associated handler as a minter
cb-sol-cli erc20 add-minter --minter "0x3167776db165D8eA0f51790CA2bbf44Db5105ADF"

# Register the erc721 contract as mintable/burnable
cb-sol-cli bridge set-burn --tokenContract "0xd7E33e1bbf65dC001A0Eb1552613106CD7e40C31" --handler "0x3f709398808af36ADBA86ACC617FeB7F5B7B193E"

# Add the handler as a minter
cb-sol-cli erc721 add-minter --minter "0x3f709398808af36ADBA86ACC617FeB7F5B7B193E"
```

## On-Chain Setup (Substrate)

### Registering Relayers

First we need to register the account of the relayer on substrate (cb-sol-cli deploys contracts with the 5 test keys preloaded). 

Select the `Sudo` tab in the PolkadotJS UI. Choose the `addRelayer` method of `chainBridge`, and select Alice as the relayer.

### Register Resources

Select the `Sudo` tab and call `chainBridge.setResourceId` for each of the transfer types you wish to use:

**Fungible (Native asset):**

Id: `0x000000000000000000000000000000c76ebe4a02bbc34786d860b355f5a5ce00`

Method: `0x4578616d706c652e7472616e73666572` (utf-8 encoding of "Example.transfer")

**NonFungible(ERC721):**

Id: `0x000000000000000000000000000000e389d61c11e5fe32ec1735b3cd38c69501`

Method: `0x4578616d706c652e6d696e745f657263373231` (utf-8 encoding of "Example.mint_erc721")

**Generic (Hash Transfer):**

Id: `0x000000000000000000000000000000f44be64d2de895454c3467021928e55e01`

Method:  `0x4578616d706c652e72656d61726b` (utf-8 encoding of "Example.remark")

### Whitelist Chains

Using the `Sudo` tab, call `chainBridge.whitelistChain`, specifying `0` for out ethereum chain ID.

## Running A Relayer

Here is an example config file for a single relayer ("Alice") using the contracts we've deployed.

```json
{
  "chains": [
    {
      "name": "eth",
      "type": "ethereum",
      "id": "0",
      "endpoint": "ws://localhost:8545",
      "from": "0xff93B45308FD417dF303D6515aB04D9e89a750Ca",
      "opts": {
        "bridge": "0x62877dDCd49aD22f5eDfc6ac108e9a4b5D2bD88B",
        "erc20Handler": "0x3167776db165D8eA0f51790CA2bbf44Db5105ADF",
        "erc721Handler": "0x3f709398808af36ADBA86ACC617FeB7F5B7B193E",
        "genericHandler": "0x2B6Ab4b880A45a07d83Cf4d664Df4Ab85705Bc07",
        "gasLimit": "1000000",
        "maxGasPrice": "20000000"
      }
    },
    {
      "name": "sub",
      "type": "substrate",
      "id": "1",
      "endpoint": "ws://localhost:9944",
      "from": "5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY",
      "opts": {}
    }
  ]
}
```

You can then start a relayer using the default "Alice" key:

```bash
chainbridge --config config.json --testkey alice --latest
```

## Fungible Transfers

### Substrate Native Token ⇒ ERC 20

In the substrate UI select the `Extrinsics` tab, and call `example.transferNative` with these parameters:

- Amount: `1000` **(select `Pico` for units)**
- Recipient: `0xff93B45308FD417dF303D6515aB04D9e89a750Ca`
- Dest Id: `0`

You can query the recipients balance on ethereum with this:

```bash
cb-sol-cli erc20 balance --address "0xff93B45308FD417dF303D6515aB04D9e89a750Ca"
```

### ERC20 ⇒ Substrate Native Token

If necessary, you can mint some tokens:

```bash
cb-sol-cli erc20 mint --amount 1000
```

Before initiating the transfer we have to approve the bridge to take ownership of the tokens:

```bash
cb-sol-cli erc20 approve --amount 1000 --recipient "0x3167776db165D8eA0f51790CA2bbf44Db5105ADF"
```

To initiate a transfer on the ethereum chain use this command (Note: there will be a 10 block delay before the relayer will process the transfer):

```bash
cb-sol-cli erc20 deposit --amount 1000 --dest 1 --recipient "0xd43593c715fdd31c61141abd04a99fd6822c8558854ccde39a5684e7a56da27d" --resourceId "0x000000000000000000000000000000c76ebe4a02bbc34786d860b355f5a5ce00"
```

## Non-Fungible Transfers

### Substrate NFT ⇒ ERC721

First, you'll need to mint a token. Select the `Sudo` tab and call `erc721.mint` with parameters such as these:

- Owner: `Alice`
- TokenId: `1`
- Metadata: `""`

Now the owner of the token can initiate a transfer by calling `example.transferErc721`:

- Recipient: `0xff93B45308FD417dF303D6515aB04D9e89a750Ca`
- TokenId: `1`
- DestId: `0`

You can query ownership of tokens on ethereum with this:

```bash
cb-sol-cli erc721 owner --id 0x1
```

### ERC721 ⇒ Substrate NFT

If necessary, you can mint an erc721 token like this:

```bash
cb-sol-cli erc721 mint --id 0x99
```

Before initiating the transfer, we must approve the bridge to take ownership of the tokens:

```bash
cb-sol-cli erc721 approve --id 0x99 --recipient "0x3f709398808af36ADBA86ACC617FeB7F5B7B193E"
```

Now we can initiate the transfer:

```bash
cb-sol-cli erc721 deposit --id 0x99 --dest 1 --resourceId "0x000000000000000000000000000000e389d61c11e5fe32ec1735b3cd38c69501" --recipient "0xd43593c715fdd31c61141abd04a99fd6822c8558854ccde39a5684e7a56da27d"
```

### Generic Data Transfer

To demonstrate a possible use of the generic data transfer, we have a hash registry on ethereum. We also have a method on the example substrate chain to emit a hash inside an event, which we can trigger from ethereum. 

### Generic Data Substrate ⇒ Eth

For this example we will transfer a 32 byte hash to a registry on ethereum. Using the `Extrinsics` tab, call `example.transferHash`:

- Hash: `0x699c776c7e6ce8e6d96d979b60e41135a13a2303ae1610c8d546f31f0c6dc730`
- Dest ID: `0`

You can verify the transfer with this command:

```bash
cb-sol-cli cent getHash --hash 0x699c776c7e6ce8e6d96d979b60e41135a13a2303ae1610c8d546f31f0c6dc730
```
