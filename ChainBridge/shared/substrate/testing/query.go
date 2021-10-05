// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package subtest

import (
	"bytes"
	"math/big"
	"testing"

	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

func QueryStorage(t *testing.T, client *utils.Client, prefix, method string, arg1, arg2 []byte, result interface{}) bool {
	exists, err := utils.QueryStorage(client, prefix, method, arg1, arg2, result)
	if err != nil {
		t.Fatal(err)
	}
	return exists
}

func QueryConst(t *testing.T, client *utils.Client, prefix, name string, res interface{}) {
	err := utils.QueryConst(client, prefix, name, res)
	if err != nil {
		t.Fatal(err)
	}
}

func AssertErc721Metadata(t *testing.T, client *utils.Client, id *big.Int, expected types.Bytes) {
	token, err := utils.GetErc721Token(client, types.NewU256(*id))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(token.Metadata, expected) {
		t.Fatalf("Erc721 metadata mismatch for token %s. Got: %x Expected: %x", id.String(), token.Metadata, expected)
	}
	log15.Info("Asserted erc721 metadata", "tokenId", id, "metadata", token.Metadata)
}

func AssertErc721NonExistence(t *testing.T, client *utils.Client, id *big.Int) {
	var res types.Bytes // No expecting to decode is
	exists := QueryStorage(t, client, "TokenStorage", "Tokens", id.Bytes(), nil, &res)
	if exists {
		t.Fatalf("erc721 token %s exists but should have been burned", id)
	}
}
