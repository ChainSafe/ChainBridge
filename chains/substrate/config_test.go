// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"testing"

	"github.com/ChainSafe/chainbridge-utils/core"
)

func TestParseStartBlock(t *testing.T) {
	// Valid option included in config
	cfg := &core.ChainConfig{Opts: map[string]string{"startBlock": "1000"}}

	blk := parseStartBlock(cfg)

	if blk != 1000 {
		t.Fatalf("Got: %d Expected: %d", blk, 1000)
	}

	// Not included in config
	cfg = &core.ChainConfig{Opts: map[string]string{}}

	blk = parseStartBlock(cfg)

	if blk != 0 {
		t.Fatalf("Got: %d Expected: %d", blk, 0)
	}
}
