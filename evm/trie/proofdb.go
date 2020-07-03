// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package trie

import (
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type ProofDatabase struct {
	db   map[string][]byte
	lock sync.RWMutex
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
	_, val := db.db[string(key)]
	return val, nil
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
