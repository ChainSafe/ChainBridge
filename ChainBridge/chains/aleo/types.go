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
	"encoding/hex"
	"fmt"
	"math/big"
	"regexp"
)

type DepositLog struct {
	DestinationChainID uint8  `json:"destination_chain_id"`
	Nonce              uint64 `json:"nonce"`
	Handler            string `json:"handler"`
}

type ARC721DepositRecord struct {
	DestinationChainID          uint8  `json:"destination_chain_id"`
	ResourceID                  string `json:"resource_id"`
	Depositer                   string `json:"depositer"`
	DestinationRecipientAddress string `json:"destination_recipient_address"`
	TokenID                     string `json:"token_id"`
	TokenURI                    string `json:"token_uri"`
}

// HexStringFormat Removes leading 0x from a hex string
func HexStringFormat(hexString string) string {
	re := regexp.MustCompile(`(?i)0x`)
	return re.ReplaceAllString(hexString, "")
}

// HexStringToBytes Converts a hex string to a byte array
func HexStringToBytes(hexString string) ([]byte, error) {
	return hex.DecodeString(HexStringFormat(hexString))
}

// HexStringToBytes32 Converts a hex string to a [32]byte array representing the resourceID
func HexStringToBytes32(hexString string) ([32]byte, error) {
	var byte32 [32]byte
	resultBytes, err := HexStringToBytes(hexString)
	if len(resultBytes) > 32 {
		err = fmt.Errorf("invalid byte length: length must be 32 bytes")
	}
	copy(byte32[:], resultBytes)
	return byte32, err
}

// HexStringToBigInt Converts a hex string to a big Int
func HexStringToBigInt(hexString string) (*big.Int, error) {
	bigInt := new(big.Int)
	resultBytes, err := HexStringToBytes(hexString)
	bigInt.SetBytes(resultBytes)
	return bigInt, err
}
