// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

// An available method on the substrate chain
type Method string

var AddRelayerMethod Method = "Bridge.add_relayer"
var SetResourceMethod Method = "Bridge.set_resource"
var SetThresholdMethod Method = "Bridge.set_threshold"
var WhitelistChainMethod Method = "Bridge.whitelist_chain"
var ExampleTransferNativeMethod Method = "Example.transfer_native"
var ExampleTransferErc721Method Method = "Example.transfer_erc721"
var ExampleTransferHashMethod Method = "Example.transfer_hash"
var ExampleMintErc721Method Method = "Example.mint_erc721"
var ExampleTransferMethod Method = "Example.transfer"
var ExampleRemarkMethod Method = "Example.remark"
var Erc721MintMethod Method = "Erc721.mint"
var SudoMethod Method = "Sudo.sudo"
