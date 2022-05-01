// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only
/*
The ethereum package contains the logic for interacting with ethereum chains.

There are 3 major components: the connection, the listener, and the writer.
The currently supported transfer types are Fungible (ERC20), Non-Fungible (ERC721), and generic.

Connection

The connection contains the ethereum RPC client and can be accessed by both the writer and listener.

Listener

The listener polls for each new block and looks for deposit events in the bridge contract. If a deposit occurs, the listener will fetch additional information from the handler before constructing a message and forwarding it to the router.

Writer

The writer recieves the message and creates a proposals on-chain. Once a proposal is made, the writer then watches for a finalization event and will attempt to execute the proposal if a matching event occurs. The writer skips over any proposals it has already seen.
*/
package ethereum

import (
	"math/big"

	erc721Handler "github.com/ChainSafe/ChainBridge/bindings/ERC721Handler"
	bridge "github.com/ChainSafe/ChainBridge/bindings/SnowBridge"
	connection "github.com/ChainSafe/ChainBridge/connections/ethereum"
	"github.com/ChainSafe/chainbridge-utils/blockstore"
	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	metrics "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var _ core.Chain = &Chain{}

var _ Connection = &connection.Connection{}

type Connection interface {
	Connect() error
	Keypair() *secp256k1.Keypair
	Opts() *bind.TransactOpts
	CallOpts() *bind.CallOpts
	LockAndUpdateOpts() error
	UnlockOpts()
	Client() *ethclient.Client
	EnsureHasBytecode(address common.Address) error
	LatestBlock() (*big.Int, error)
	WaitForBlock(block *big.Int, delay *big.Int) error
	Close()
}

// Message is used as a generic format to communicate between chains
type Message struct {
	Source       ChainId      // Source where message was initiated
	Destination  ChainId      // Destination chain of message
	Type         TransferType // type of bridge transfer
	DepositNonce Nonce        // Nonce for the deposit
	ResourceId   ResourceId
	Payload      []interface{} // data associated with event sequence
}

type Chain struct {
	cfg      *core.ChainConfig // The config of the chain
	conn     Connection        // THe chains connection
	listener *listener         // The listener of this chain
	writer   *writer           // The writer of the chain
	stop     chan<- int
}

// checkBlockstore queries the blockstore for the latest known block. If the latest block is
// greater than cfg.startBlock, then cfg.startBlock is replaced with the latest known block.
func setupBlockstore(cfg *Config, kp *secp256k1.Keypair) (*blockstore.Blockstore, error) {
	bs, err := blockstore.NewBlockstore(cfg.blockstorePath, cfg.id, kp.Address())
	if err != nil {
		return nil, err
	}

	if !cfg.freshStart {
		latestBlock, err := bs.TryLoadLatestBlock()
		if err != nil {
			return nil, err
		}

		if latestBlock.Cmp(cfg.startBlock) == 1 {
			cfg.startBlock = latestBlock
		}
	}

	return bs, nil
}

func InitializeChain(chainCfg *core.ChainConfig, logger log15.Logger, sysErr chan<- error, m *metrics.ChainMetrics) (*Chain, error) {
	cfg, err := parseChainConfig(chainCfg)
	if err != nil {
		return nil, err
	}

	kpI, err := keystore.KeypairFromAddress(cfg.from, keystore.EthChain, cfg.keystorePath, chainCfg.Insecure)
	if err != nil {
		return nil, err
	}
	kp, _ := kpI.(*secp256k1.Keypair)

	bs, err := setupBlockstore(cfg, kp)
	if err != nil {
		return nil, err
	}

	stop := make(chan int)
	conn := connection.NewConnection(cfg.endpoint, cfg.http, kp, logger, cfg.gasLimit, cfg.maxGasPrice, cfg.gasMultiplier)
	err = conn.Connect()
	if err != nil {
		return nil, err
	}
	err = conn.EnsureHasBytecode(cfg.bridgeContract)
	if err != nil {
		return nil, err
	}

	// err = conn.EnsureHasBytecode(cfg.erc721HandlerContract)
	// if err != nil {
	// 	return nil, err
	// }

	bridgeContract, err := bridge.NewSnowBridge(cfg.bridgeContract, conn.Client()) // initialize our snowbridge contract
	if err != nil {
		return nil, err
	}

	// chainId, err := bridgeContract.DomainID(conn.CallOpts()) // get ChainID
	// if err != nil {
	// 	return nil, err
	// }

	// if chainId != uint8(chainCfg.Id) {
	// 	return nil, fmt.Errorf("chainId (%d) and configuration chainId (%d) do not match", chainId, chainCfg.Id)
	// }

	erc721HandlerContract, err := erc721Handler.NewERC721Handler(cfg.erc721HandlerContract, conn.Client()) // initialize our ERC721 handler
	if err != nil {
		return nil, err
	}

	if chainCfg.LatestBlock {
		curr, err := conn.LatestBlock()
		if err != nil {
			return nil, err
		}
		cfg.startBlock = curr
	}

	listener := NewListener(conn, cfg, logger, bs, stop, sysErr, m)
	listener.setContracts(bridgeContract, erc721HandlerContract)

	writer := NewWriter(conn, cfg, logger, stop, sysErr, m)
	writer.setContract(bridgeContract)

	return &Chain{
		cfg:      chainCfg,
		conn:     conn,
		writer:   writer,
		listener: listener,
		stop:     stop,
	}, nil
}

func (c *Chain) SetRouter(r *core.Router) {
	r.Listen(c.cfg.Id, c.writer)
	c.listener.setRouter(r)
}

func (c *Chain) Start() error {
	err := c.listener.start()
	if err != nil {
		return err
	}

	err = c.writer.start()
	if err != nil {
		return err
	}

	c.writer.log.Debug("Successfully started chain")
	return nil
}

func (c *Chain) Id() msg.ChainId {
	return c.cfg.Id
}

func (c *Chain) Name() string {
	return c.cfg.Name
}

func (c *Chain) LatestBlock() metrics.LatestBlock {
	return c.listener.latestBlock
}

// Stop signals to any running routines to exit
func (c *Chain) Stop() {
	close(c.stop)
	if c.conn != nil {
		c.conn.Close()
	}
}
