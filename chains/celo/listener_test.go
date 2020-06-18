// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package celo

import (
	"context"
	connection "github.com/ChainSafe/ChainBridge/connections/ethereum"
	"github.com/ChainSafe/ChainBridge/keystore"
	ethutils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"testing"
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
	// get the block hash of the receipt
	blockHash := l.getTransactionBlockHash(hash)
	// using the receipt blockhash get the block
	block, err := l.conn.Client().BlockByHash(context.Background(), blockHash)
	if err != nil {
		t.Fatal(err)
	}
	// Confirm that the receipt blockhash and the block's blockhash are the same
	if block.Hash() != blockHash {
		t.Fatalf("block hash are not equal, expected: %x, %x", hash, block.TxHash())
	}
	log15.Info("Transaction worked")

}

func TestListener_BlockTransactionsByHash(t *testing.T) {
	l := createTestListener(t)
	err := l.start()
	if err != nil {
		t.Fatal(err)
	}

	// Create and submit a new transaction
	hash := newTransaction(t, l)
	// Get blockhash from receipt
	blockHash := l.getTransactionBlockHash(hash)

	// Get txHashes and txroot from blockHash
	l.getBlockTransactionsByHash(blockHash)

}
