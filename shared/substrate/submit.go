// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"fmt"
	"math/big"
	"sync"

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

	var acct types.AccountInfo
	_, err = QueryStorage(client, "System", "Account", client.Key.PublicKey, nil, &acct)
	if err != nil {
		return err
	}

	// Sign the extrinsic
	o := types.SignatureOptions{
		BlockHash:          client.Genesis,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        client.Genesis,
		Nonce:              types.NewUCompactFromUInt(uint64(acct.Nonce)),
		SpecVersion:        rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: rv.TransactionVersion,
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

func SubmitSudoTx(client *Client, method Method, args ...interface{}) error {
	call, err := types.NewCall(client.Meta, string(method), args...)
	if err != nil {
		return err
	}
	return SubmitTx(client, SudoMethod, call)
}

// Batch submit take multiple calls and attempts to sumbit all of them, then waits until they complete.
// This should allow for the calls to be processed in a single block (to some limit).
// WARNING: Failed calls are not reported
func BatchSubmit(client *Client, calls []types.Call) error {
	// Get latest runtime version
	rv, err := client.Api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return err
	}

	var acct types.AccountInfo
	_, err = QueryStorage(client, "System", "Account", client.Key.PublicKey, nil, &acct)
	if err != nil {
		return err
	}

	// Sign the extrinsic
	o := types.SignatureOptions{
		BlockHash:          client.Genesis,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        client.Genesis,
		Nonce:              types.NewUCompactFromUInt(uint64(acct.Nonce)),
		SpecVersion:        rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: rv.TransactionVersion,
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(calls))

	for _, call := range calls {
		ext := types.NewExtrinsic(call)

		err = ext.Sign(*client.Key, o)
		if err != nil {
			return err
		}

		// Submit and watch the extrinsic
		sub, err := client.Api.RPC.Author.SubmitAndWatchExtrinsic(ext)
		if err != nil {
			return err
		}

		go func() {
			for {
				status := <-sub.Chan()
				switch {
				case status.IsInBlock:
					wg.Done()
					return
				case status.IsDropped:
					wg.Done()
					return
				case status.IsInvalid:
					wg.Done()
					return
				}
			}
		}()

		bigNonce := big.Int(o.Nonce)
		o.Nonce = types.NewUCompactFromUInt(bigNonce.Uint64() + 1)
	}

	wg.Wait()
	return nil
}
