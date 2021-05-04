// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var BlockRetryInterval = time.Second * 5

type Connection struct {
	endpoint      string
	http          bool
	kp            *secp256k1.Keypair
	gasLimit      *big.Int
	maxGasPrice   *big.Int
	gasMultiplier *big.Float
	conn          *ethclient.Client
	// signer    ethtypes.Signer
	opts     *bind.TransactOpts
	callOpts *bind.CallOpts
	nonce    uint64
	optsLock sync.Mutex
	log      log15.Logger
	stop     chan int // All routines should exit when this channel is closed
}

// NewConnection returns an uninitialized connection, must call Connection.Connect() before using.
func NewConnection(endpoint string, http bool, kp *secp256k1.Keypair, log log15.Logger, gasLimit, gasPrice *big.Int, gasMultiplier *big.Float) *Connection {
	return &Connection{
		endpoint:      endpoint,
		http:          http,
		kp:            kp,
		gasLimit:      gasLimit,
		maxGasPrice:   gasPrice,
		gasMultiplier: gasMultiplier,
		log:           log,
		stop:          make(chan int),
	}
}

// Connect starts the ethereum WS connection
func (c *Connection) Connect() error {
	c.log.Info("Connecting to ethereum chain...", "url", c.endpoint)
	var rpcClient *rpc.Client
	var err error
	// Start http or ws client
	if c.http {
		rpcClient, err = rpc.DialHTTP(c.endpoint)
	} else {
		rpcClient, err = rpc.DialWebsocket(context.Background(), c.endpoint, "/ws")
	}
	if err != nil {
		return err
	}
	c.conn = ethclient.NewClient(rpcClient)

	// Construct tx opts, call opts, and nonce mechanism
	opts, _, err := c.newTransactOpts(big.NewInt(0), c.gasLimit, c.maxGasPrice)
	if err != nil {
		return err
	}
	c.opts = opts
	c.nonce = 0
	c.callOpts = &bind.CallOpts{From: c.kp.CommonAddress()}
	return nil
}

// newTransactOpts builds the TransactOpts for the connection's keypair.
func (c *Connection) newTransactOpts(value, gasLimit, gasPrice *big.Int) (*bind.TransactOpts, uint64, error) {
	privateKey := c.kp.PrivateKey()
	address := ethcrypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := c.conn.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, 0, err
	}

	id, err := c.conn.ChainID(context.Background())
	if err != nil {
		return nil, 0, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, id)
	if err != nil {
		return nil, 0, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	auth.GasLimit = uint64(gasLimit.Int64())
	auth.GasPrice = gasPrice
	auth.Context = context.Background()

	return auth, nonce, nil
}

func (c *Connection) Keypair() *secp256k1.Keypair {
	return c.kp
}

func (c *Connection) Client() *ethclient.Client {
	return c.conn
}

func (c *Connection) Opts() *bind.TransactOpts {
	return c.opts
}

func (c *Connection) CallOpts() *bind.CallOpts {
	return c.callOpts
}

func (c *Connection) SafeEstimateGas(ctx context.Context, apiKey, estimation string) (*big.Int, error) {
	// TODO: if the apikey is stored in a config file, we won't need to pass as an arg but can pull
	//       that info from the configuration file.
	var gsnRes *gsnResponse
	retries := 10
	gsnRes = callGSN(apiKey, retries)

	// TODO: will need to ask the user for what type of tx type. Options are:
	//       fast, fastest, safeLow, average, block_time, blockNum, speed, safeLowWait, avgWait, fastWait, fastestWait
	//       https://docs.ethgasstation.info/gas-price
	//       default to "fastest" if no estimation is given

	if gsnRes != nil {
		// fmt.Println("Response: ", gsnRes)
		switch estimation {
		case "fast":
			return big.NewInt(gsnRes.Fast), nil
		case "fastest":
			return big.NewInt(gsnRes.Fastest), nil
		case "safeLow":
			return big.NewInt(gsnRes.SafeLow), nil
		case "average":
			return big.NewInt(gsnRes.Average), nil
		default:
			return big.NewInt(gsnRes.Fast), nil
		}
	}

	suggestedGasPrice, err := c.conn.SuggestGasPrice(context.TODO())
	if err != nil {
		return nil, err
	}

	gasPrice := multiplyGasPrice(suggestedGasPrice, c.gasMultiplier)

	// Check we aren't exceeding our limit
	if gasPrice.Cmp(c.maxGasPrice) == 1 {
		return c.maxGasPrice, nil
	} else {
		return gasPrice, nil
	}
}

func multiplyGasPrice(gasEstimate *big.Int, gasMultiplier *big.Float) *big.Int {

	gasEstimateFloat := new(big.Float).SetInt(gasEstimate)

	result := gasEstimateFloat.Mul(gasEstimateFloat, gasMultiplier)

	gasPrice := new(big.Int)

	result.Int(gasPrice)

	return gasPrice
}

// LockAndUpdateOpts acquires a lock on the opts before updating the nonce
// and gas price.
func (c *Connection) LockAndUpdateOpts() error {
	c.optsLock.Lock()

	gasPrice, err := c.SafeEstimateGas(context.TODO())
	if err != nil {
		return err
	}
	c.opts.GasPrice = gasPrice

	nonce, err := c.conn.PendingNonceAt(context.Background(), c.opts.From)
	if err != nil {
		c.optsLock.Unlock()
		return err
	}
	c.opts.Nonce.SetUint64(nonce)
	return nil
}

func (c *Connection) UnlockOpts() {
	c.optsLock.Unlock()
}

// LatestBlock returns the latest block from the current chain
func (c *Connection) LatestBlock() (*big.Int, error) {
	header, err := c.conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

// EnsureHasBytecode asserts if contract code exists at the specified address
func (c *Connection) EnsureHasBytecode(addr ethcommon.Address) error {
	code, err := c.conn.CodeAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}

	if len(code) == 0 {
		return fmt.Errorf("no bytecode found at %s", addr.Hex())
	}
	return nil
}

// WaitForBlock will poll for the block number until the current block is equal or greater.
// If delay is provided it will wait until currBlock - delay = targetBlock
func (c *Connection) WaitForBlock(targetBlock *big.Int, delay *big.Int) error {
	for {
		select {
		case <-c.stop:
			return errors.New("connection terminated")
		default:
			currBlock, err := c.LatestBlock()
			if err != nil {
				return err
			}

			if delay != nil {
				currBlock.Sub(currBlock, delay)
			}

			// Equal or greater than target
			if currBlock.Cmp(targetBlock) >= 0 {
				return nil
			}
			c.log.Trace("Block not ready, waiting", "target", targetBlock, "current", currBlock, "delay", delay)
			time.Sleep(BlockRetryInterval)
			continue
		}
	}
}

// Close terminates the client connection and stops any running routines
func (c *Connection) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
	close(c.stop)
}

// Utility methods for connections

type RPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Response struct {
	Error  *RPCError       `json:"error"`
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
}

type gsnResponse struct {
	Fast          int64       `json:"fast"`
	Fastest       int64       `json:"fastest"`
	SafeLow       int64       `json:"safeLow"`
	Average       int64       `json:"average"`
	Block_time    float32     `json:"block_time"`
	BlockNum      int64       `json:"blockNum"`
	Speed         float32     `json:"speed"`
	SafeLowWait   float32     `json:"safeLowWait"`
	AvgWait       float32     `json:"avgWait"`
	FastWait      float32     `json:"fastWait"`
	FastestWait   float32     `json:"fastestWait"`
	GasPriceRange interface{} `json:"gasPriceRange"`
}

// callGSN will call the Gas Station Network and request the gas prices on the Ethereum network
// TODO: if the apikey is stored in a config file, we won't need to pass as an arg but can pull
//       that info from the configuration file.
func callGSN(apiKey string, retries int) *gsnResponse {
	time.Sleep(1 * time.Second)

	gsnURL := "https://ethgasstation.info/api/ethgasAPI.json"
	if len(apiKey) != 0 {
		gsnURL = gsnURL + "?api-key=" + apiKey
	}

	var gsnRes gsnResponse
	var body []byte

	for i := 0; i < retries; i++ {
		res, err := http.Get(gsnURL)
		if err != nil {
			return nil
		}

		body, err = ioutil.ReadAll(res.Body)
		err = json.Unmarshal(body, &gsnRes)
		if err != nil {
			return nil
		}

		if body != nil && res.StatusCode == http.StatusOK {
			break
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Println(big.NewInt(gsnRes.Fast))
	fmt.Println(big.NewInt(gsnRes.Fastest))
	fmt.Println(big.NewInt(gsnRes.SafeLow))
	fmt.Println(big.NewInt(gsnRes.Average))

	return &gsnRes
}
