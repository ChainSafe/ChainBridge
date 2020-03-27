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

var SetAddress = createMethod("set_address")

var WhitelistChain = createMethod("whitelist_chain")

var AssetTx = createMethod("transfer_asset")

var CreateProposal = createMethod("create_proposal")

var Approve = createMethod("approve")

var ExampleTransfer Method = "Example.transfer"

var ExampleTransferHash = Method("Example.transfer_hash")

// TODO: May not be needed outside of testing
var AddValidator = createMethod("add_validator")
var RemoveValidator = createMethod("remove_validator")
var Sudo Method = "Sudo.sudo"
