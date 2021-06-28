// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
)

type RelayTx struct {
	to       common.Address
	data     []byte
	gasLimit uint
	schedule string
}

// ABI encode the vote proposal function args
func packVoteProposalData(chainID uint8, depositNonce uint64, resourceID [32]byte, dataHash [32]byte) ([]byte, error) {
	parsed, err := abi.JSON(strings.NewReader(Bridge.BridgeABI))
	if err != nil {
		return nil, err
	}
	packed, err := parsed.Pack("voteProposal", chainID, depositNonce, resourceID, dataHash)
	if err != nil {
		return nil, err
	}
	return packed, nil
}

// ABI encode the execute proposal function args
func packExecuteProposalData(chainID uint8, depositNonce uint64, data []byte, resourceID [32]byte) ([]byte, error) {
	parsed, err := abi.JSON(strings.NewReader(Bridge.BridgeABI))
	if err != nil {
		return nil, err
	}
	packed, err := parsed.Pack("executeProposal", chainID, depositNonce, data, resourceID)
	if err != nil {
		return nil, err
	}
	return packed, nil
}

type SignedRelayTx struct {
	tx   RelayTx
	sig  []byte
	txId string
}

// Form and sign a relay transaction
func toSignedRelayTx(to string, data []byte, gasLimit uint, chainId uint, keyPair *secp256k1.Keypair) (*SignedRelayTx, error) {
	// see https://infura.io/docs/transactions#section/Developing-with-ITX/Signing-a-relay-request
	toAddress := common.HexToAddress(to)
	gasInt := big.NewInt(int64(gasLimit))
	chainIdBigInt := big.NewInt(int64(chainId))
	schedule := "fast"

	uint256Ty, err := abi.NewType("uint256", "uint256", nil)
	if err != nil {
		return nil, err
	}

	addressTy, err := abi.NewType("address", "address", nil)
	if err != nil {
		return nil, err
	}

	bytesTy, err := abi.NewType("bytes", "bytes", nil)
	if err != nil {
		return nil, err
	}

	stringTy, err := abi.NewType("string", "string", nil)
	if err != nil {
		return nil, err
	}

	arguments := abi.Arguments{
		{
			Type: addressTy,
		},
		{
			Type: bytesTy,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: stringTy,
		},
	}

	packed, err := arguments.Pack(toAddress, data, gasInt, chainIdBigInt, schedule)
	if err != nil {
		return nil, err
	}

	relayTxId := crypto.Keccak256Hash(packed)
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(relayTxId), string(relayTxId.Bytes()))
	msgHash := crypto.Keccak256Hash([]byte(msg))

	sig, err := crypto.Sign(msgHash.Bytes(), keyPair.PrivateKey())
	if err != nil {
		return nil, err
	}

	return &SignedRelayTx{
		tx: RelayTx{
			to:       toAddress,
			data:     data,
			gasLimit: gasLimit,
			schedule: schedule,
		},
		sig:  sig,
		txId: relayTxId.String(),
	}, nil
}

// Send a relay transaction to the provided rpc client
func sendRelayTransaction(rpc *rpc.Client, ctx context.Context, relayTx RelayTx, sig []byte) (*string, error) {
	var hex hexutil.Bytes
	sigHex := "0x" + common.Bytes2Hex(sig)
	txArg := map[string]interface{}{
		"to":       relayTx.to.String(),
		"data":     "0x" + common.Bytes2Hex(relayTx.data),
		"gasLimit": relayTx.gasLimit,
		"schedule": relayTx.schedule,
	}
	err := rpc.CallContext(ctx, &hex, "relay_sendTransaction", txArg, sigHex)

	if err != nil {
		return nil, err
	} else {
		txId := hex.String()
		return &txId, nil
	}
}
