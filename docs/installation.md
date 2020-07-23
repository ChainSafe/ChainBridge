# Installation

## Dependencies

- [Subkey](https://github.com/paritytech/substrate): 
Used for substrate key management. Only required if connecting to a substrate chain.

  `make install-subkey`


## Building

`make build`: Builds `chainbridge` in `./build`.

**or**

`make install`: Uses `go install` to add `chainbridge` to your GOBIN.