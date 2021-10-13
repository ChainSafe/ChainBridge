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
	"github.com/ChainSafe/chainbridge-utils/msg"
)

func (l *listener) handleArc721DepositedEvent(destId msg.ChainId, nonce msg.Nonce) (msg.Message, error) {
	l.log.Info("Handling nonfungible deposit event")

	record, err := l.conn.Arc721DepositRecord(destId, nonce)

	if err != nil {
		l.log.Error("Error Unpacking ARC721 Deposit Record", "err", err)
		return msg.Message{}, err
	}

	l.log.Info("Record output ", "record", record)

	resourceID, err := HexStringToBytes32(record.ResourceID)
	if err != nil {
		l.log.Error("decode error: couldn't hex decode resource id bytes", "err", err)
		return msg.Message{}, err
	}

	tokenID, err := HexStringToBigInt(record.TokenID)
	if err != nil {
		l.log.Error("decode error: couldn't hex decode token id bytes", "err", err)
		return msg.Message{}, err
	}

	destinationRecipientAddress, err := HexStringToBytes(record.DestinationRecipientAddress)
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
