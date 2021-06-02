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

const ForwarderAbi = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"struct MinimalForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"struct MinimalForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

type ForwarderClient struct {
	client             *ethclient.Client
	forwarderAddress   common.Address
	forwarderAbi       abi.ABI
	forwarderNonceLock sync.Mutex
	forwarderNonce     *big.Int
	fromAddress        common.Address
}

func NewForwarderClient(
	client *ethclient.Client,
	forwarderAddress common.Address,
	fromAddress common.Address,
) (*ForwarderClient, error) {
	forwarderAbi, err := abi.JSON(strings.NewReader(ForwarderAbi))
	if err != nil {
		return nil, err
	}

	return &ForwarderClient{
		client:           client,
		forwarderAddress: forwarderAddress,
		fromAddress:      fromAddress,
		forwarderAbi:     forwarderAbi,
	}, nil
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

func (c *ForwarderClient) LockAndNextNonce() (*big.Int, error) {
	c.forwarderNonceLock.Lock()

	if c.forwarderNonce == nil {
		return c.GetOnChainNonce()
	} else {
		return c.forwarderNonce.Add(c.forwarderNonce, big.NewInt(1)), nil
	}
}

func (c *ForwarderClient) UnlockAndSetNonce(nonce *big.Int) {
	// CHRIS:TODO check all the comments in this file, some of them are todos
	// CHRIS:TODO docs about the nil

	if nonce != nil {
		c.forwarderNonce = nonce
	}
	c.forwarderNonceLock.Unlock()
}

func (c *ForwarderClient) TypedHash(
	from, to string,
	data []byte,
	value, gas *math.HexOrDecimal256,
	nonce *big.Int,
	chainId *math.HexOrDecimal256,
	verifyingContract string) ([]byte, error) {

	typedData := signer.TypedData{
		Types: signer.Types{
			// EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
			"EIP712Domain": []signer.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			// ForwardRequest(address from,address to,uint256 value,uint256 gas,uint256 nonce,bytes data)
			"ForwardRequest": []signer.Type{
				{Name: "from", Type: "address"},
				{Name: "to", Type: "address"},
				{Name: "value", Type: "uint256"},
				{Name: "gas", Type: "uint256"},
				{Name: "nonce", Type: "uint256"},
				{Name: "data", Type: "bytes"},
			},
		},
		PrimaryType: "ForwardRequest",
		Domain: signer.TypedDataDomain{
			Name:              "MinimalForwarder",
			ChainId:           chainId,
			Version:           "0.0.1",
			VerifyingContract: verifyingContract,
		},
		Message: signer.TypedDataMessage{
			"from":  from,
			"to":    to,
			"value": value,
			"gas":   gas,
			"data":  data,
			"nonce": math.NewHexOrDecimal256(nonce.Int64()),
		},
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		// CHRIS: add the proper errors everywhere
		return nil, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, err
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	return crypto.Keccak256(rawData), nil
}

func (c *ForwarderClient) PackAndSignForwarderArg(
	from, to common.Address,
	data []byte,
	nonce, value, gas *big.Int,
	chainId uint,
	kp secp256k1.Keypair) ([]byte, error) {
	// CHRIS:TODO: need to revert this nonce if we fail to send

	forwarderHash, err := c.TypedHash(
		from.String(),
		to.String(),
		data,
		math.NewHexOrDecimal256(int64(value.Uint64())),
		math.NewHexOrDecimal256(int64(gas.Uint64())),
		nonce,
		math.NewHexOrDecimal256(int64(chainId)),
		c.forwarderAddress.String())

	if err != nil {
		return nil, err
	}

	// now sign that hash
	sig, err := crypto.Sign(forwarderHash, kp.PrivateKey())
	sig[64] += 27 // Transform V from 0/1 to 27/28 according to the yellow paper

	if err != nil {
		return nil, err
	}

	// struct ForwardRequest {
	//     address from;
	//     address to;
	//     uint256 value;
	//     uint256 gas;
	//     uint256 nonce;
	//     bytes data;
	// }
	forwardReq := struct {
		From  common.Address
		To    common.Address
		Value *big.Int
		Gas   *big.Int
		Nonce *big.Int
		Data  []byte
	}{
		From:  from,
		To:    to,
		Value: value,
		Gas:   gas,
		Nonce: nonce,
		Data:  data,
	}

	packed, err := c.forwarderAbi.Pack("execute", forwardReq, sig)
	if err != nil {
		return nil, err
	}

	return packed, nil
}
