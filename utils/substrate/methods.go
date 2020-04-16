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
var ExampleTransferHashMethod Method = "Example.transfer_hash"
var ExampleTransferMethod Method = "Example.transfer"
var ExampleRemarkMethod Method = "Example.remark"
var SudoMethod Method = "Sudo.sudo"
