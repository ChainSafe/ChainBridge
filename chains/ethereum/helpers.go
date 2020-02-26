// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

const (
	DepositAsset = "DepositAsset"
	NftTransfer  = "NftTransfer"
	ErcTransfer  = "ErcTransfer"

	// For internal usage only
	DepositAssetSignature = "DepositAsset(address,bytes32)"
	NftTransferSignature  = "NFTTransfer(uint256,uint256,address,address,uint256,bytes)"
	ErcTransferSignature  = "ERCTransfer(uint256,uint256,address,uint256,address)"
)

func BuildEventSubscriptions(events []string) []string {
	var arr = []string{}
	for _, event := range events {
		if event == DepositAsset {
			arr = append(arr, DepositAssetSignature)
		} else if event == NftTransfer {
			arr = append(arr, NftTransferSignature)
		} else if event == ErcTransfer {
			arr = append(arr, ErcTransferSignature)
		}
	}
	return arr
}
