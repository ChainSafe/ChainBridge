// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"fmt"

	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

func SubmitTx(client *Client, method Method, args ...interface{}) error {
	// Create call and extrinsic
	call, err := types.NewCall(
		client.Meta,
		string(method),
		args...,
	)
	if err != nil {
		return err
	}
	ext := types.NewExtrinsic(call)

	// Get latest runtime version
	rv, err := client.Api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return err
	}

	var acct AccountData
	_, err = QueryStorage(client, "System", "Account", client.Key.PublicKey, nil, &acct)
	if err != nil {
		return err
	}

	// Sign the extrinsic
	o := types.SignatureOptions{
		BlockHash:   client.Genesis,
		Era:         types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash: client.Genesis,
		Nonce:       types.UCompact(acct.Nonce),
		SpecVersion: rv.SpecVersion,
		Tip:         0,
	}
	err = ext.Sign(*client.Key, o)
	if err != nil {
		return err
	}

	// Submit and watch the extrinsic
	sub, err := client.Api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return err
	}

	for {
		status := <-sub.Chan()
		switch {
		case status.IsInBlock:
			log15.Info("Extrinsic in block", "block", status.AsInBlock.Hex())
			return nil
		case status.IsDropped:
			return fmt.Errorf("extrinsic dropped")
		case status.IsInvalid:
			return fmt.Errorf("extrinsic invalid")
		}
	}
}

func SubmitAndWait() {
	// TODO

}

func SubmitSudoTx(client *Client, method Method, args ...interface{}) error {
	call, err := types.NewCall(client.Meta, string(method), args...)
	if err != nil {
		return err
	}
	return SubmitTx(client, SudoMethod, call)
}
