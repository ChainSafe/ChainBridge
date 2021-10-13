/* Copyright (c) 2021 Forte Labs
 * All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and remains the property of
 * Forte Labs and their suppliers, if any.  The intellectual and technical
 * concepts contained herein are proprietary to Forte Labs and their suppliers
 * and may be covered by U.S. and Foreign Patents, patents in process, and are
 * protected by trade secret or copyright law. Dissemination of this information
 * or reproduction of this material is strictly forbidden unless prior written
 * permission is obtained from Forte Labs.
 */

package aleo

import (
	"context"
	"math/big"

	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/rpc"
)

type Connection struct {
	endpoint string
	http     bool
	log      log15.Logger
	client   *rpc.Client
	stop     chan int
}

// NewConnection returns an uninitialized connection, must call Connection.Connect() before using.
func NewConnection(endpoint string, http bool, log log15.Logger) *Connection {
	return &Connection{
		endpoint: endpoint,
		http:     http,
		log:      log,
		stop:     make(chan int),
	}
}

// Connect starts the http or WS connection to the Aleo Custodian, this is established via RPC but does not make the same calls
func (c *Connection) Connect() error {
	c.log.Info("Connecting to the aleo custodian...", "url", c.endpoint)
	var rpcClient *rpc.Client
	var err error

	if c.http {
		rpcClient, err = rpc.DialHTTP(c.endpoint)
	} else {
		rpcClient, err = rpc.DialContext(context.Background(), c.endpoint)
	}

	if err != nil {
		return err
	}
	c.client = rpcClient

	return nil
}

// LatestBlock queries the Aleo Custodian for it's advertised "latest block", from our custodian's RPC replication endpoint
func (c *Connection) LatestBlock() (*big.Int, error) {
	var latestBlock *big.Int
	err := c.client.CallContext(context.Background(), &latestBlock, "latest_block")
	return latestBlock, err
}

// DepositEvents queries the Aleo Custodian for the deposit events at the latest block
func (c *Connection) DepositEvents(latestBlock *big.Int) ([]DepositLog, error) {
	var results []DepositLog
	arg := map[string]interface{}{
		"latest_block": latestBlock,
	}
	err := c.client.CallContext(context.Background(), &results, "deposit_events", arg)
	return results, err
}

// Arc721DepositRecord gets a requested ARC721 Deposit record from the custodian
func (c *Connection) Arc721DepositRecord(destId msg.ChainId, nonce msg.Nonce) (ARC721DepositRecord, error) {
	var record ARC721DepositRecord
	arg := map[string]interface{}{
		"destination_chain_id": uint64(destId),
		"nonce":                uint64(nonce),
	}
	err := c.client.CallContext(context.Background(), &record, "deposit_record", arg)
	return record, err
}

// Close terminates the client connection and stops any running routines
func (c *Connection) Close() {
	if c.client != nil {
		c.client.Close()
	}
	close(c.stop)
}
