// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package celo

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	connection "github.com/ChainSafe/ChainBridge/connections/ethereum"
	"github.com/ChainSafe/ChainBridge/keystore"
	ethutils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var TestEndpoint = "ws://localhost:8545"
var AliceKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var GasLimit = big.NewInt(ethutils.DefaultGasLimit)
var GasPrice = big.NewInt(ethutils.DefaultGasPrice)
var ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

func createTestListener(t *testing.T) *listener {
	conn := connection.NewConnection(TestEndpoint, false, AliceKp, log15.Root(), GasLimit, GasPrice)
	l := NewListener(conn)

	return l
}

// creating and sending a new transaction
func newTransaction(t *testing.T, l *listener) common.Hash {

	// Creating a new transaction
	nonce := l.conn.Opts().Nonce
	tx := types.NewTransaction(nonce.Uint64(), ZeroAddress, big.NewInt(0), ethutils.DefaultGasLimit, GasPrice, nil)

	chainId, err := l.conn.Client().NetworkID(context.Background())
	signer := types.NewEIP155Signer(chainId)
	if err != nil {
		t.Fatal(err)
	}

	// Sign the transaction
	signedTx, err := types.SignTx(tx, signer, AliceKp.PrivateKey()) //SignTx(tx *Transaction, s Signer, prv *ecdsa.PrivateKey) (*Transaction, error)
	if err != nil {
		t.Fatal(err)
	}

	// Send the transaction for execution
	err = l.conn.Client().SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatal(err)
	}
	return signedTx.Hash()
}

// WaitForTx will query the chain at ExpectedBlockTime intervals, until a receipt is returned.
// Returns an error if the tx failed.
func waitForTx(client *ethclient.Client, hash common.Hash) (*types.Receipt, error) {
	retry := 10
	for retry > 0 {
		receipt, err := client.TransactionReceipt(context.Background(), hash)
		if err != nil {
			retry--
			time.Sleep(ExpectedBlockTime)
			continue
		}

		if receipt.Status != 1 {
			return nil, fmt.Errorf("transaction failed on chain")
		}
		return receipt, nil
	}
	return nil, fmt.Errorf("transaction after retries failed")
}

func TestListener_start_stop(t *testing.T) {
	l := createTestListener(t)

	err := l.start()
	if err != nil {
		t.Fatal(err)
	}

	// Initiate shutdown
	l.close()
}

// Testing transaction Block hash
func TestListener_TransactionBlockHash(t *testing.T) {

	l := createTestListener(t)
	err := l.start()
	if err != nil {
		t.Fatal(err)
	}

	// Create and submit a new transaction and return the signed transaction hash
	hash := newTransaction(t, l)

	receipt, err := waitForTx(l.conn.Client(), hash)
	if err != nil {
		t.Fatal(err)
	}

	blockHash, err := l.getTransactionBlockHash(hash)
	if err != nil {
		t.Fatal(err)
	}

	// Get the block
	retrievedBlock, err := l.conn.Client().BlockByHash(context.Background(), blockHash)
	if err != nil {
		t.Fatal(err)
	}

	// Confirm that the receipt blockhash and the block's blockhash are the same
	if blockHash != receipt.BlockHash {
		t.Fatalf("block hash are not equal, expected: %x, %x", retrievedBlock.Hash(), receipt.TxHash)
	}
}

func TestListener_BlockTransactionsByHash(t *testing.T) {
	l := createTestListener(t)
	err := l.start()
	if err != nil {
		t.Fatal(err)
	}

	// Create and submit a new transaction
	hash := newTransaction(t, l)

	// Get receipt from hash
	receipt, err := waitForTx(l.conn.Client(), hash)
	if err != nil {
		t.Fatal(err)
	}

	// Get txHashes and txroot from blockHash
	txs, _, err := l.getBlockTransactionsByHash(receipt.BlockHash)
	if err != nil {
		t.Fatal(err)
	}

	if len(txs) != 1 {
		t.Fatalf("transaction hashes should have a length of one, %x", txs)
	}

	if txs[0] != hash {
		t.Fatalf("hash and transactions should be the same: hash, %x txs, %x", hash, txs)
	}
}
