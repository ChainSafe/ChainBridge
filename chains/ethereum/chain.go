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
	"fmt"
	"math/big"

	"github.com/ChainSafe/ChainBridge/chains/ethereum/client"
	"github.com/ChainSafe/ChainBridge/chains/ethereum/config"

	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	erc20Handler "github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	erc721Handler "github.com/ChainSafe/ChainBridge/bindings/ERC721Handler"
	"github.com/ChainSafe/ChainBridge/bindings/GenericHandler"
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

type Chain struct {
	cfg      *config.Config // The config of the chain
	conn     *client.Client // THe chains connection
	listener *listener      // The listener of this chain
	writer   *writer        // The writer of the chain
	stop     chan<- int
}

// checkBlockstore queries the blockstore for the latest known block. If the latest block is
// greater than cfg.startBlock, then cfg.startBlock is replaced with the latest known block.
func setupBlockstore(cfg *config.Config, kp *secp256k1.Keypair) (*blockstore.Blockstore, error) {
	bs, err := blockstore.NewBlockstore(cfg.BlockstorePath, cfg.Id, kp.Address())
	if err != nil {
		return nil, err
	}

	if !cfg.FreshStart {
		latestBlock, err := bs.TryLoadLatestBlock()
		if err != nil {
			return nil, err
		}

		if latestBlock.Cmp(cfg.StartBlock) == 1 {
			cfg.StartBlock = latestBlock
		}
	}

	return bs, nil
}

func InitializeChain(cfg *config.Config, logger log15.Logger, sysErr chan<- error, m *metrics.ChainMetrics) (*Chain, error) {
	kpI, err := keystore.KeypairFromAddress(cfg.From, keystore.EthChain, cfg.KeystorePath, cfg.Insecure)
	if err != nil {
		return nil, err
	}
	kp, _ := kpI.(*secp256k1.Keypair)

	bs, err := setupBlockstore(cfg, kp)
	if err != nil {
		return nil, err
	}

	stop := make(chan int)
	conn := client.NewClient(cfg.Endpoint, cfg.Http, kp, logger, cfg.GasLimit, cfg.MaxGasPrice, cfg.GasMultiplier)
	err = conn.Connect()
	if err != nil {
		return nil, err
	}

	bridgeContract, err := bridge.NewBridge(cfg.BridgeContract, conn)
	if err != nil {
		return nil, err
	}

	chainId, err := bridgeContract.ChainID(conn.CallOpts())
	if err != nil {
		return nil, err
	}

	if chainId != uint8(cfg.Id) {
		return nil, fmt.Errorf("chainId (%d) and configuration chainId (%d) do not match", chainId, cfg.Id)
	}

	erc20HandlerContract, err := erc20Handler.NewERC20Handler(cfg.ERC20HandlerContract, conn)
	if err != nil {
		return nil, err
	}

	erc721HandlerContract, err := erc721Handler.NewERC721Handler(cfg.ERC721HandlerContract, conn)
	if err != nil {
		return nil, err
	}

	genericHandlerContract, err := GenericHandler.NewGenericHandler(cfg.GenericHandlerContract, conn)
	if err != nil {
		return nil, err
	}

	if cfg.LatestBlock {
		curr, err := conn.LatestBlock()
		if err != nil {
			return nil, err
		}
		cfg.StartBlock = curr
	}

	listener := NewListener(cfg, conn, logger, bs, stop, sysErr, m)
	listener.SetContracts(bridgeContract, erc20HandlerContract, erc721HandlerContract, genericHandlerContract)

	writer := NewWriter(cfg, conn, logger, stop, sysErr, m)
	writer.setContract(bridgeContract)

	return &Chain{
		cfg:      cfg,
		conn:     conn,
		writer:   writer,
		listener: listener,
		stop:     stop,
	}, nil
}

func (c *Chain) SetRouter(r *core.Router) {
	r.Listen(c.cfg.Id, c.writer)
	c.listener.SetRouter(r)
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
