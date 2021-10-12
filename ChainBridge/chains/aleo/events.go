/* Copyright (c) 2021 Forte Labs
 * All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and remains the property of
 * Forte Labs and their suppliers, if any.  The intellectual and technical
 * concepts contained herein are proprietary to Forte Labs and their suppliers
 * and may be covered by U.S. and Foreign Patents, patents in process, and are
 * protected by trade secret or copyright law. Dissemination of this information
 * or reproduction of this material is strictly forbidden unless prior written
 * permission is obtained from Forte Labs.
 */

package aleo

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ChainSafe/chainbridge-utils/msg"
)

type ARC721DepositRecord struct {
	DestinationChainID          uint8  `json:"destination_chain_id"`
	ResourceID                  string `json:"resource_id"`
	Depositer                   string `json:"depositer"`
	DestinationRecipientAddress string `json:"destination_recipient_address"`
	TokenID                     string `json:"token_id"`
	TokenURI                    string `json:"token_uri"`
}

func hexStringFormat(hexString string) string {
	repl := strings.Replace(hexString, "0x", "", -1)
	repl = strings.Replace(repl, "0X", "", -1)
	return repl
}

func (l *listener) handleArc721DepositedEvent(destId msg.ChainId, nonce msg.Nonce) (msg.Message, error) {
	l.log.Info("Handling nonfungible deposit event")
	arg := map[string]interface{}{
		"destination_chain_id": uint64(destId),
		"nonce":                uint64(nonce),
	}
	var record ARC721DepositRecord

	if err := l.conn.client.CallContext(context.Background(), &record, "deposit_record", arg); err != nil {
		l.log.Error("Error Unpacking ARC721 Deposit Record", "err", err)
		return msg.Message{}, err
	}
	l.log.Info("Record output ", "record", record)

	var resourceID [32]byte
	formattedResourceID := hexStringFormat(record.ResourceID)
	resourceBytes, err := hex.DecodeString(formattedResourceID)
	if err != nil {
		l.log.Error("decode error: couldn't hex decode resource id bytes", "err", err)
		return msg.Message{}, err
	}
	if len(resourceBytes) > 32 {
		l.log.Error("Resource ID must be 32 bytes")
		return msg.Message{}, fmt.Errorf("resource ID must be a 32 bytes")
	}
	copy(resourceID[:], resourceBytes)

	tokenID := new(big.Int)
	formattedTokenID := hexStringFormat(record.TokenID)
	tokenIDBytes, err := hex.DecodeString(formattedTokenID)
	if err != nil {
		l.log.Error("decode error: couldn't hex decode token id bytes", "err", err)
		return msg.Message{}, err
	}
	tokenID.SetBytes(tokenIDBytes)
	l.log.Info("Token id ", "token_id", tokenID)

	formattedDestinationRecipientAddress := hexStringFormat(record.DestinationRecipientAddress)
	destinationRecipientAddress, err := hex.DecodeString(formattedDestinationRecipientAddress)
	if err != nil {
		l.log.Error("decode error: couldn't hex decode DestinationRecipientAddress bytes", "err", err)
		return msg.Message{}, err
	}

	MetaData := []byte(record.TokenURI)

	return msg.NewNonFungibleTransfer(
		l.cfg.id,
		destId,
		nonce,
		resourceID,
		tokenID,
		destinationRecipientAddress,
		MetaData,
	), nil

}
