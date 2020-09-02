# Testing

Unit tests require an ethereum node running on `localhost:8545` and a substrate node running on `localhost:9944`. E2E tests require an additional ethereum node on `localhost:8546`. 

A docker-compose file is provided to run two Geth nodes and a chainbridge-substrate-chain node in isolated environments:
```
$ docker-compose -f ./docker-compose-e2e.yml up
```

See [chainbridge-solidity](https://github.com/chainsafe/chainbridge-solidity) and [chainbridge-substrate-chain](https://github.com/ChainSafe/chainbridge-substrate-chain) for more information on testing facilities.

All Go tests can be run with:
```
$ make test
```
Go tests specifically for ethereum, substrate and E2E can be run with
```
$ make test-eth
$ make test-sub
$ make test-e2e
```

The bindings for the solidity contracts live in `bindings/`. To update the bindings modify `scripts/setup-contracts.sh` and then run `make clean && make setup-contracts`
