// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package subtest

import (
	"testing"

	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
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
