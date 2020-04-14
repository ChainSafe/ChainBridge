// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"
)

// Name of pallet defined within substrate
const PalletName = "Bridge"

// An available method on the substrate chain
type Method string

func createMethod(method string) Method {
	return Method(fmt.Sprintf("%s.%s", PalletName, method))
}

func (m Method) String() string {
	return string(m)
}

var AddRelayer = createMethod("add_relayer")
var SetResource = createMethod("set_resource")
var SetThreshold = createMethod("set_threshold")
var WhitelistChain = createMethod("whitelist_chain")
var AcknowledgeProposal = createMethod("acknowledge_proposal")
var ExampleTransferNative Method = "Example.transfer_native"
var ExampleTransferHash Method = "Example.transfer_hash"
var ExampleTransfer Method = "Example.transfer"
var ExampleRemark Method = "Example.remark"
var Sudo Method = "Sudo.sudo"
