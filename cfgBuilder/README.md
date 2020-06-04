# Eth Config Builder

This is a simple CLI tool to help automate new ETH deployments.

An input JSON config looks like this:

```json
{
  "relayerThreshold": "3",
  "relayers": [
    "0xff93B45308FD417dF303D6515aB04D9e89a750Ca",
    "0x8e0a907331554AF72563Bd8D43051C2E64Be5d35",
    "0x24962717f8fA5BA3b931bACaF9ac03924EB475a0"
  ],
  "chains": [
    {
      "name": "goerli",
      "chainId": "1",
      "endpoint": "http://localhost:8545",
      "bridge": "0x35542aC472082524e5D815763b2531dFf98Ac548",
      "erc20Handler": "0xF8eD8035856241900B23F230b5589f72678Aedfa",
      "erc721Handler": "0xAf65aEa42847bcb4897d3CF566Cd89248A196B17",
      "genericHandler": "0x30663188630403e7df0288B5Bd18c119A9Ef75ED",
      "gasLimit": "1000000",
      "gasPrice": "20000000",
      "startBlock": "0",
      "http": "false"
    },
    {
      "name": "kotti",
      "chainId": "2",
      "endpoint": "http://localhost:8546",
      "bridge": "0x35542aC472082524e5D815763b2531dFf98Ac548",
      "erc20Handler": "0xF8eD8035856241900B23F230b5589f72678Aedfa",
      "erc721Handler": "0xAf65aEa42847bcb4897d3CF566Cd89248A196B17",
      "genericHandler": "0x30663188630403e7df0288B5Bd18c119A9Ef75ED",
      "gasLimit": "1000000",
      "gasPrice": "20000000",
      "startBlock": "0",
      "http": "true"
    }
  ]
}
```
Compile an executable to run `cfgBuilder` with `make install-cfgBuilder ` or `make build-cfgBuilder` 
Running `./build/cfgBuilder <json-path> <output-path>` will produce config files for ChainBridge, and (optional) place them in the specified path. For each relayer in `relayers` a corresponding config will be generated.