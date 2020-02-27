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
