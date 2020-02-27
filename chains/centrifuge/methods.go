package substrate

import (
	"fmt"
)

// Name of pallet defined within substrate
const PalletName = "Bridge"

// An available method on the substrate chain
type Method string

func (m Method) String() string {
	return string(m)
}

var SetAddressMethod = Method(fmt.Sprintf("%s.%s", PalletName, "set_address"))

var WhitelistChainMethod = Method(fmt.Sprintf("%s.%s", PalletName, "whitelist_chain"))

var AssetTxMethod = Method(fmt.Sprintf("%s.%s", PalletName, "transfer_asset"))
