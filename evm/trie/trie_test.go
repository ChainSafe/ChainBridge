// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package trie

import (
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"github.com/ethereum/go-ethereum/rlp"
)

var (
	defaultTriesToStore = 3
)

func createNewTxTries(numHistoricalTries int) *TxTries {
	return NewTxTries(numHistoricalTries)
}

func createTempDB() *leveldb.Database {

	diskdb, err := leveldb.New("./temp-database", 256, 0, "")
	if err != nil {
		panic(fmt.Sprintf("unable to create testing database: %v", err))
	}
	return diskdb
}

func deleteTempDB() {
	os.RemoveAll("./temp-database")
}

func addTrie(txTries *TxTries, root common.Hash, transactions []common.Hash, db *leveldb.Database) error {

	if db == nil {
		db = createTempDB()
	}

	if transactions == nil {
		transactions = make([]common.Hash, 0)
	}

	_, err := txTries.AddTrie(root, db, transactions)

	if err != nil {
		return err
	}

	return nil

}

func TestEmptyTxTries(t *testing.T) {
	txTries := createNewTxTries(defaultTriesToStore)

	if txTries.triesToStore != defaultTriesToStore {
		t.Fatalf("tries to store not set properly, expected: %x, got: %x", defaultTriesToStore, txTries.triesToStore)
	}
}

func TestAddEmptyTrie(t *testing.T) {
	txTries := createNewTxTries(defaultTriesToStore)
	addTrie(txTries, emptyRoot, nil, nil)

	if txTries.txRoots[0] != emptyRoot {
		t.Fatalf("failed to set txRoot in txTries properly, expected: %x, got: %x", emptyRoot, txTries.txRoots[0])
	}

	if txTries.txTries[0].trie.Hash() != emptyRoot {
		t.Fatalf("trie does not have empty hash as root, expected: %x, got: %x", emptyRoot, txTries.txTries[0].trie.Hash())
	}

	deleteTempDB()
}

func TestAddEmptyTrieRetrieveProof_Fails(t *testing.T) {
	txTries := createNewTxTries(defaultTriesToStore)
	addTrie(txTries, emptyRoot, nil, nil)

	if txTries.txRoots[0] != emptyRoot {
		t.Fatalf("failed to set txRoot in txTries properly, expected: %x, got: %x", emptyRoot, txTries.txRoots[0])
	}

	if txTries.txTries[0].trie.Hash() != emptyRoot {
		t.Fatalf("trie does not have empty hash as root, expected: %x, got: %x", emptyRoot, txTries.txTries[0].trie.Hash())
	}

	deleteTempDB()
}

func TestAddSingleTrieUpdate(t *testing.T) {
	vals := []common.Hash{common.HexToHash("123"), common.HexToHash("456")}
	expectedRoot := common.HexToHash("8e4ba8a974b3dfa9dbfb76d61c633ba2a1250a4d79e3e912e841392fdb232c17")

	txTries := createNewTxTries(defaultTriesToStore)
	err := addTrie(txTries, expectedRoot, vals, nil)
	if err != nil {
		t.Fatal(err)
	}

	if txTries.txRoots[0] != expectedRoot {
		t.Fatalf("failed to set txRoot in txTries properly, expected: %x, got: %x", expectedRoot, txTries.txRoots[0])
	}

	if txTries.txTries[0].trie.Hash() != expectedRoot {
		t.Fatalf("trie does not have empty hash as root, expected: %x, got: %x", expectedRoot, txTries.txTries[0].trie.Hash())
	}

	deleteTempDB()

}

func TestAddSingleTrieRetrieveProof(t *testing.T) {
	vals := []common.Hash{common.HexToHash("123"), common.HexToHash("456")}
	expectedRoot := common.HexToHash("8e4ba8a974b3dfa9dbfb76d61c633ba2a1250a4d79e3e912e841392fdb232c17")

	txTries := createNewTxTries(defaultTriesToStore)
	err := addTrie(txTries, expectedRoot, vals, nil)
	if err != nil {
		t.Fatal(err)
	}

	if txTries.txRoots[0] != expectedRoot {
		t.Fatalf("failed to set txRoot in txTries properly, expected: %x, got: %x", expectedRoot, txTries.txRoots[0])
	}

	if txTries.txTries[0].trie.Hash() != expectedRoot {
		t.Fatalf("trie does not have empty hash as root, expected: %x, got: %x", expectedRoot, txTries.txTries[0].trie.Hash())
	}

	bytesIndex, err := intToBytes(0)
	if err != nil {
		t.Fatal(err)
	}
	keyToRetrieve, err := rlp.EncodeToBytes(bytesIndex)
	if err != nil {
		t.Fatal(err)
	}

	proofDb, err := txTries.RetrieveProof(expectedRoot, keyToRetrieve)
	if err != nil {
		t.Fatal(err)
	}

	exists, err := verifyProof(expectedRoot, keyToRetrieve, proofDb)
	if err != nil {
		t.Fatal(err)
	}

	if exists != true {
		t.Fatalf("not able to verify retrieved proof!")
	}

	deleteTempDB()

}

func TestAddMultipleTries(t *testing.T) {
	txTries := createNewTxTries(defaultTriesToStore)
	db := createTempDB()

	vals1 := []common.Hash{common.HexToHash("123"), common.HexToHash("456")}
	expectedRoot1 := common.HexToHash("8e4ba8a974b3dfa9dbfb76d61c633ba2a1250a4d79e3e912e841392fdb232c17")

	err := addTrie(txTries, expectedRoot1, vals1, db)
	if err != nil {
		t.Fatal(err)
	}

	if txTries.txRoots[0] != expectedRoot1 {
		t.Fatalf("failed to set txRoot in txTries properly, expected: %x, got: %x", expectedRoot1, txTries.txRoots[0])
	}

	if txTries.txTries[0].trie.Hash() != expectedRoot1 {
		t.Fatalf("trie does not have empty hash as root, expected: %x, got: %x", expectedRoot1, txTries.txTries[0].trie.Hash())
	}

	vals2 := []common.Hash{common.HexToHash("abc"), common.HexToHash("def")}
	expectedRoot2 := common.HexToHash("f004b10977aed387a552429aac0228822c5502d6cf2d7c27ab41729bded6de88")

	err = addTrie(txTries, expectedRoot2, vals2, db)
	if err != nil {
		t.Fatal(err)
	}

	if txTries.txRoots[1] != expectedRoot2 {
		t.Fatalf("failed to set txRoot in txTries properly, expected: %x, got: %x", expectedRoot2, txTries.txRoots[1])
	}

	if txTries.txTries[1].trie.Hash() != expectedRoot2 {
		t.Fatalf("trie does not have empty hash as root, expected: %x, got: %x", expectedRoot2, txTries.txTries[1].trie.Hash())
	}

	vals3 := []common.Hash{common.HexToHash("qwe"), common.HexToHash("rty")}
	expectedRoot3 := common.HexToHash("76cdfcecc7b62db5a03b511f2b50bc56c69e33a3a46bc42bcdb1a46d77245739")

	err = addTrie(txTries, expectedRoot3, vals3, db)
	if err != nil {
		t.Fatal(err)
	}

	if txTries.txRoots[2] != expectedRoot3 {
		t.Fatalf("failed to set txRoot in txTries properly, expected: %x, got: %x", expectedRoot3, txTries.txRoots[2])
	}

	if txTries.txTries[2].trie.Hash() != expectedRoot3 {
		t.Fatalf("trie does not have empty hash as root, expected: %x, got: %x", expectedRoot3, txTries.txTries[2].trie.Hash())
	}

	err = addTrie(txTries, expectedRoot1, vals1, db)
	if err != nil {
		t.Fatal(err)
	}

	if txTries.txRoots[2] != expectedRoot1 {
		t.Fatalf("failed to set txRoot in txTries properly, expected: %x, got: %x", expectedRoot1, txTries.txRoots[2])
	}

	if txTries.txTries[2].trie.Hash() != expectedRoot1 {
		t.Fatalf("trie does not have empty hash as root, expected: %x, got: %x", expectedRoot1, txTries.txTries[2].trie.Hash())
	}

	if txTries.txRoots[0] != expectedRoot2 {
		t.Fatalf("failed to set txRoot in txTries properly, expected: %x, got: %x", expectedRoot2, txTries.txRoots[0])
	}

	if txTries.txTries[0].trie.Hash() != expectedRoot2 {
		t.Fatalf("trie does not have empty hash as root, expected: %x, got: %x", expectedRoot2, txTries.txTries[0].trie.Hash())
	}

	deleteTempDB()

}

func TestAddMultipleTriesRetrieveProof(t *testing.T) {
	txTries := createNewTxTries(defaultTriesToStore)
	db := createTempDB()

	vals1 := []common.Hash{common.HexToHash("123"), common.HexToHash("456")}
	expectedRoot1 := common.HexToHash("8e4ba8a974b3dfa9dbfb76d61c633ba2a1250a4d79e3e912e841392fdb232c17")

	err := addTrie(txTries, expectedRoot1, vals1, db)
	if err != nil {
		t.Fatal(err)
	}

	if txTries.txRoots[0] != expectedRoot1 {
		t.Fatalf("failed to set txRoot in txTries properly, expected: %x, got: %x", expectedRoot1, txTries.txRoots[0])
	}

	if txTries.txTries[0].trie.Hash() != expectedRoot1 {
		t.Fatalf("trie does not have empty hash as root, expected: %x, got: %x", expectedRoot1, txTries.txTries[0].trie.Hash())
	}

	vals2 := []common.Hash{common.HexToHash("abc"), common.HexToHash("def")}
	expectedRoot2 := common.HexToHash("f004b10977aed387a552429aac0228822c5502d6cf2d7c27ab41729bded6de88")

	err = addTrie(txTries, expectedRoot2, vals2, db)
	if err != nil {
		t.Fatal(err)
	}

	if txTries.txRoots[1] != expectedRoot2 {
		t.Fatalf("failed to set txRoot in txTries properly, expected: %x, got: %x", expectedRoot2, txTries.txRoots[1])
	}

	if txTries.txTries[1].trie.Hash() != expectedRoot2 {
		t.Fatalf("trie does not have empty hash as root, expected: %x, got: %x", expectedRoot2, txTries.txTries[1].trie.Hash())
	}

	vals3 := []common.Hash{common.HexToHash("qwe"), common.HexToHash("rty")}
	expectedRoot3 := common.HexToHash("76cdfcecc7b62db5a03b511f2b50bc56c69e33a3a46bc42bcdb1a46d77245739")

	err = addTrie(txTries, expectedRoot3, vals3, db)
	if err != nil {
		t.Fatal(err)
	}

	if txTries.txRoots[2] != expectedRoot3 {
		t.Fatalf("failed to set txRoot in txTries properly, expected: %x, got: %x", expectedRoot3, txTries.txRoots[2])
	}

	if txTries.txTries[2].trie.Hash() != expectedRoot3 {
		t.Fatalf("trie does not have empty hash as root, expected: %x, got: %x", expectedRoot3, txTries.txTries[2].trie.Hash())
	}

	err = addTrie(txTries, expectedRoot1, vals1, db)
	if err != nil {
		t.Fatal(err)
	}

	if txTries.txRoots[2] != expectedRoot1 {
		t.Fatalf("failed to set txRoot in txTries properly, expected: %x, got: %x", expectedRoot1, txTries.txRoots[2])
	}

	if txTries.txTries[2].trie.Hash() != expectedRoot1 {
		t.Fatalf("trie does not have empty hash as root, expected: %x, got: %x", expectedRoot1, txTries.txTries[2].trie.Hash())
	}

	if txTries.txRoots[0] != expectedRoot2 {
		t.Fatalf("failed to set txRoot in txTries properly, expected: %x, got: %x", expectedRoot2, txTries.txRoots[0])
	}

	if txTries.txTries[0].trie.Hash() != expectedRoot2 {
		t.Fatalf("trie does not have empty hash as root, expected: %x, got: %x", expectedRoot2, txTries.txTries[0].trie.Hash())
	}

	bytesIndex, err := intToBytes(0)
	if err != nil {
		t.Fatal(err)
	}
	keyToRetrieve, err := rlp.EncodeToBytes(bytesIndex)
	if err != nil {
		t.Fatal(err)
	}

	proofDb1, err := txTries.RetrieveProof(expectedRoot1, keyToRetrieve)
	if err != nil {
		t.Fatal(err)
	}

	exists1, err := verifyProof(expectedRoot1, keyToRetrieve, proofDb1)
	if err != nil {
		t.Fatal(err)
	}

	if exists1 != true {
		t.Fatalf("not able to verify retrieved proof!")
	}

	proofDb2, err := txTries.RetrieveProof(expectedRoot2, keyToRetrieve)
	if err != nil {
		t.Fatal(err)
	}

	exists2, err := verifyProof(expectedRoot2, keyToRetrieve, proofDb2)
	if err != nil {
		t.Fatal(err)
	}

	if exists2 != true {
		t.Fatalf("not able to verify retrieved proof!")
	}

	proofDb3, err := txTries.RetrieveProof(expectedRoot3, keyToRetrieve)
	if err != nil {
		t.Fatal(err)
	}

	exists3, err := verifyProof(expectedRoot3, keyToRetrieve, proofDb3)
	if err != nil {
		t.Fatal(err)
	}

	if exists3 != true {
		t.Fatalf("not able to verify retrieved proof!")
	}

	deleteTempDB()

}
