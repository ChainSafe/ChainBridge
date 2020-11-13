// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package trie

import (
	"errors"
	"sync"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type ProofDatabase struct {
	db   			map[string][]byte
	lock 			sync.RWMutex
}

// NewProofDatabase returns a wrapped map
func NewProofDatabase() *ProofDatabase {
	db := &ProofDatabase{
		db: make(map[string][]byte),
	}

	return db
}

// Close sets ProofDatabase db's db to nil
func (db *ProofDatabase) Close() error {
	db.lock.Lock()
	defer db.lock.Unlock()

	db.db = nil
	return nil
}

// Has checks if key exists in ProofDatabase db
func (db *ProofDatabase) Has(key []byte) (bool, error) {
	if db.db == nil {
		return false, errors.New("database does not exist")
	}
	_, exists := db.db[string(key)]
	return exists, nil
}

// Get retrieves value associated with key in ProofDatabase db, and checks for existance
func (db *ProofDatabase) Get(key []byte) ([]byte, error) {
	if db.db == nil {
		return nil, errors.New("database does not exist")
	}
	if val, exists := db.db[string(key)]; exists {
		return common.CopyBytes(val), nil
	}

	return nil, errors.New("key not in database")
}

// Put insets and associates value with key in ProofDatabase db
func (db *ProofDatabase) Put(key []byte, value []byte) error {

	if db.db == nil {
		return errors.New("database does not exist")
	}
	db.db[string(key)] = common.CopyBytes(value)
	return nil
}

// Delete removes the key and associated value from the database
func (db *ProofDatabase) Delete(key []byte) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	if db.db == nil {
		return errors.New("database does not exist")
	}
	delete(db.db, string(key))
	return nil
}

// Encodes a proof Database to a format parsable by the on chain contract
func encodeProofDB(rootHash common.Hash, key []byte, proofDb ProofDatabase) (encodedProof [][]byte, err error) {
	key = keybytesToHex(key)
	wantHash := rootHash

	// we want to repeat until we have reached the desired value node
	for i := 0; ; i++ {
		// retrieve target node from proofDB
		buf, _ := proofDb.Get(wantHash[:])
		if buf == nil {
			return nil, fmt.Errorf("proof node %d (hash %064x) missing", i, wantHash)
		}

		n, err := decodeNode(wantHash[:], buf)
		if err != nil {
			return nil, fmt.Errorf("bad proof node %d: %v", i, err)
		}

		// we know that the RLP encoded node of the root element should be the next element of our encoded proof
		// there are two cases: Either this node is a full node or a short node.

		// if the node is a full node it is a 17 element array of format [v0 ... v15, value] 
		// where v0 ... v15 represent each possible value by next nibble in the key
		// where value holds the element we are constructing the proof for if we have reached the end of the path.

		// if the node is a short node it is a 2 element array of format [encodedPath, key]
		// where encodedPath contains a shortcut to skip ahead in the key
		// where key is the hash of the next node we want to retrieve.
		encodedProof = append(encodedProof, n.toArray())

		// we want to retrieve the next node on the key path
		keyrest, cld := get(n, key, true)
		switch cld := cld.(type) {
		case nil:
			// The trie doesn't contain the key
			return nil, nil
		case hashNode:
			// There are more elements in the proof to retrieve
			key = keyrest
			copy(wantHash[:], cld)
		case valueNode:
			// We have reached the desired value node
			return encodedProof, nil
		}
	}
}