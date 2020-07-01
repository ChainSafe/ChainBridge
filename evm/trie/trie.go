// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package trie

import (

	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

type Trie struct {
	t         *trie.Trie
	trieRoots []common.Hash
}

var (
	emptyRoot = common.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")
)

func NewTrie(root common.Hash, db *Database) (*Trie, error) {

	// TODO: look into cache values
	// this creates a new trie database with our KVDB as the diskDB for node storage
	newTrie, err := trie.New(root, trie.NewDatabaseWithCache(db, 0))

	if err != nil {
		return nil, err
	}

	trie := &Trie{
		t: newTrie,
	}

	if root != (common.Hash{}) && root != emptyRoot {
		trie.trieRoots = []common.Hash{root}
	}

	return trie, nil
}

func (t *Trie) UpdateTrie(transactions []common.Hash, transactionRoot common.Hash) error {
	for i, tx := range transactions {
		key, err := rlp.EncodeToBytes(i)
		if err != nil {
			return err
		}

		t.t.Update(key, tx)
	}

	if *Trie.t.root != transactionRoot{
		return errors.New("transaction roots don't match")
	}

	*Trie.trieRoots = append(Trie.trieRoots, transactionRoot)
	return nil
}

// retrieves a proof from a trie object, given the root of the trie it is contained in and the key
func (t *Trie) RetrieveProof(root common.Hash, key []byte) *ProofDatabase {
	var proof = *ProofDatabase
	err := t.t.Prove(key, 0, &proof)
	if err != nil {
		return proof
	}
}
