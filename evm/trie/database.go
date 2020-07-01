// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package trie

import (
	"errors"
	"sync"

	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type Database struct {
	path string
	db   *leveldb.DB
	lock sync.RWMutex
	log  log15.Logger
}

type ProofDatabase struct {
	db 		map[[]byte][]byte
	lock 	sync.RWMutex
}

// NewProofDatabase returns a wrapped map
func NewProofDatabase(path string, log log15.Logger) (*ProofDatabase) {
	db := &ProofDatabase{
		db: make(map[[]byte][]byte),
	}

	return db
}

func (db *ProofDatabase) Close() error {
	db.lock.Lock()
	defer db.lock.Unlock()

	db.db = nil
	return nil
}

func (db *ProofDatabase) Has(key []byte) (bool, error) {
	if db.db == nil {
		return false, errors.New("database does not exist")
	}
	_, val := db.db[key]
	return val, nil
}

func (db *ProofDatabase) Get(key []byte) ([]byte, error) {
	if db.db == nil {
		return nil, errors.New("database does not exist")
	}
	if val, exists := db.db[key]; exists {
		return common.CopyBytes(val), nil
	}

	return nil, errors.New("key not in database")
} 

func (db *ProofDatabase) Put(key []byte, value []byte) error {

	if db.db == nil {
		return errors.New("database does not exist")
	}
	db.db[key] = common.CopyBytes(value)
	return nil
}

func (db *ProofDatabase) Delete(key []byte) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	if db.db == nil {
		return errors.New("database does not exist")
	}
	delete(db.db, key)
	return nil
}

// NewDatabase returns a wrapped LevelDB obejct
func NewDatabase(path string, log log15.Logger) (*Database, error) {
	//TODO: Look into appropriate opts here
	db, err := leveldb.OpenFile(path, nil)

	if err != nil {
		return nil, err
	}

	database := &Database{
		path: path,
		db:   db,
		log:  log,
	}

	return database, nil
}

func (db *Database) Close() error {
	db.lock.Lock()
	defer db.lock.Unlock()

	return db.db.Close()
}

func (db *Database) Compact(start []byte, limit []byte) error {
	return db.db.CompactRange(util.Range{Start: start, Limit: limit})
}

func (db *Database) Has(key []byte) (bool, error) {
	return db.db.Has(key, nil)
}

func (db *Database) Get(key []byte) ([]byte, error) {
	dat, err := db.db.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func (db *Database) Put(key []byte, value []byte) error {
	return db.db.Put(key, value, nil)
}

func (db *Database) Delete(key []byte) error {
	return db.db.Delete(key, nil)
}
