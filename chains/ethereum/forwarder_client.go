// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	signer "github.com/ethereum/go-ethereum/signer/core"
)

// Keeps track of forwarder info include the state of the current
// nonces
type ForwarderClient struct {
	client             *ethclient.Client
	forwarderAddress   common.Address
	forwarderAbi       abi.ABI
	forwarderNonceLock sync.Mutex
	forwarderNonce     *big.Int
	fromAddress        common.Address
	chainId            *big.Int
}

// https://github.com/opengsn/gsn/blob/bdce42a5fbd37d1abc7bd32bdbe10fc8c71dc602/packages/contracts/src/forwarder/Forwarder.sol
const GsnForwarderAbi = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"domainValue\",\"type\":\"bytes\"}],\"name\":\"DomainRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"typeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeStr\",\"type\":\"string\"}],\"name\":\"RequestTypeRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_TYPE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENERIC_PARAMS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"}],\"name\":\"_getEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"domains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"registerDomainSeparator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"typeSuffix\",\"type\":\"string\"}],\"name\":\"registerRequestType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"typeHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

var gsnForwarderAbi abi.ABI

func init() {
	var err error
	gsnForwarderAbi, err = abi.JSON(strings.NewReader(GsnForwarderAbi))
	if err != nil {
		panic(err)
	}
}

func NewForwarderClient(
	client *ethclient.Client,
	forwarderAddress common.Address,
	fromAddress common.Address,
	chainId *big.Int,
) *ForwarderClient {
	return &ForwarderClient{
		client:           client,
		forwarderAddress: forwarderAddress,
		fromAddress:      fromAddress,
		forwarderAbi:     gsnForwarderAbi,
		chainId:          chainId,
	}
}

func (c *ForwarderClient) GetOnChainNonce() (*big.Int, error) {
	packed, err := c.forwarderAbi.Pack("getNonce", c.fromAddress)
	if err != nil {
		return nil, err
	}

	data := eth.CallMsg{
		To:    &c.forwarderAddress,
		Data:  packed,
		From:  c.fromAddress,
		Gas:   200000,
		Value: big.NewInt(0),
	}

	callOpts := &bind.CallOpts{From: c.fromAddress, Context: context.Background()}
	res, err := c.client.CallContract(callOpts.Context, data, nil)
	if err != nil {
		return nil, err
	}

	p, err := c.forwarderAbi.Unpack("getNonce", res)
	if err != nil {
		return nil, err
	}

	nonce := p[0].(*big.Int)

	return nonce, nil
}

// Locks nonce access and returns the next nonce
func (c *ForwarderClient) LockAndNextNonce() (*big.Int, error) {
	c.forwarderNonceLock.Lock()

	if c.forwarderNonce == nil {
		nonce, err := c.GetOnChainNonce()
		if err != nil {
			c.forwarderNonceLock.Unlock()
			return nil, err
		}
		return nonce, nil
	} else {
		return c.forwarderNonce.Add(c.forwarderNonce, big.NewInt(1)), nil
	}
}

// Unlocks nonce access and sets the provided nonce
// If transaction usage of a nonce failes the nonce shouldnt be
// updated upon unlock, instead nil should be supplied
func (c *ForwarderClient) UnlockAndSetNonce(nonce *big.Int) {
	if nonce != nil {
		c.forwarderNonce = nonce
	}
	c.forwarderNonceLock.Unlock()
}

// Generate the 712 type hash for signing
func (c *ForwarderClient) TypedHash(
	from, to string,
	data []byte,
	value, gas *math.HexOrDecimal256,
	nonce *big.Int,
	verifyingContract string) ([]byte, *[32]byte, *[32]byte, error) {

	chainId := math.NewHexOrDecimal256(int64(c.chainId.Uint64()))

	typedData := signer.TypedData{
		Types: signer.Types{
			// EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
			"EIP712Domain": []signer.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			// ForwardRequest(address from,address to,uint256 value,uint256 gas,uint256 nonce,bytes data,uint256 validUntil)
			"ForwardRequest": []signer.Type{
				{Name: "from", Type: "address"},
				{Name: "to", Type: "address"},
				{Name: "value", Type: "uint256"},
				{Name: "gas", Type: "uint256"},
				{Name: "nonce", Type: "uint256"},
				{Name: "data", Type: "bytes"},
				{Name: "validUntil", Type: "uint256"},
			},
		},
		// domain hash - 0xAA937C9BB1C0E90D11F3420A75F7D8EF42DFF261ADAE5B58B57AB0F692C4F7B7
		// name "GSN Relayed Transaction"- "2"
		PrimaryType: "ForwardRequest",
		Domain: signer.TypedDataDomain{
			Name:              "GSN Relayed Transaction",
			ChainId:           chainId,
			Version:           "2",
			VerifyingContract: verifyingContract,
		},
		Message: signer.TypedDataMessage{
			"from":       from,
			"to":         to,
			"value":      value,
			"gas":        gas,
			"data":       data,
			"nonce":      math.NewHexOrDecimal256(nonce.Int64()),
			"validUntil": math.NewHexOrDecimal256(0),
		},
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())

	if err != nil {
		return nil, nil, nil, err
	}
	var fixedSizeDomainSeperator [32]byte
	copy(fixedSizeDomainSeperator[:], domainSeparator)

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, nil, nil, err
	}
	var fixedSizeTypeHash [32]byte
	copy(fixedSizeTypeHash[:], typedData.TypeHash(typedData.PrimaryType))

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	return crypto.Keccak256(rawData), &fixedSizeDomainSeperator, &fixedSizeTypeHash, nil
}

// Pack and sign forwarder contract 'execute' arguments
func (c *ForwarderClient) PackAndSignForwarderArg(
	from, to common.Address,
	data []byte,
	nonce, value, gas *big.Int,
	kp secp256k1.Keypair) ([]byte, error) {

	forwarderHash, domainSeperator, typeHash, err := c.TypedHash(
		from.String(),
		to.String(),
		data,
		math.NewHexOrDecimal256(int64(value.Uint64())),
		math.NewHexOrDecimal256(int64(gas.Uint64())),
		nonce,
		c.forwarderAddress.String(),
	)

	if err != nil {
		return nil, err
	}

	// now sign that hash
	sig, err := crypto.Sign(forwarderHash, kp.PrivateKey())
	sig[64] += 27 // Transform V from 0/1 to 27/28

	if err != nil {
		return nil, err
	}

	var suffixData = common.Hex2Bytes("0x")

	type ForwardRequest struct {
		From       common.Address
		To         common.Address
		Value      *big.Int
		Gas        *big.Int
		Nonce      *big.Int
		Data       []byte
		ValidUntil *big.Int
	}

	forwardReq := ForwardRequest{
		From:       from,
		To:         to,
		Value:      value,
		Gas:        gas,
		Nonce:      nonce,
		Data:       data,
		ValidUntil: big.NewInt(0),
	}

	packed, err := c.forwarderAbi.Pack("execute", forwardReq, domainSeperator, typeHash, suffixData, sig)
	if err != nil {
		return nil, err
	}

	return packed, nil
}
