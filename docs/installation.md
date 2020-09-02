# Installation

## Docker

The official [ChainBridge Docker image](https://hub.docker.com/r/chainsafe/chainbridge)

To build the Docker image use 

```
docker build -t chainsafe/chainbridge .
```

To start ChainBridge

```
docker run -v ./config.toml:/config.toml chainsafe/chainbridge
```

## Dependencies

- [Subkey](https://github.com/paritytech/substrate): 
Used for substrate key management. Only required if connecting to a substrate chain.

```
make install-subkey
```


## Building from Source

To build `chainbridge` in `./build`.
```
make build
```

**or**

Use`go install` to add `chainbridge` to your GOBIN.

```
make install
```