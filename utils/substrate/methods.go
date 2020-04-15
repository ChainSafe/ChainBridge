package utils


// An available method on the substrate chain
type Method string

var AddRelayer Method = "Bridge.add_relayer"
var SetResource Method = "Bridge.set_resource"
var SetThreshold Method = "Bridge.set_threshold"
var WhitelistChain Method = "Bridge.whitelist_chain"
var ExampleTransferNative Method = "Example.transfer_native"
var ExampleTransferHash Method = "Example.transfer_hash"
var ExampleTransfer Method = "Example.transfer"
var ExampleRemark Method = "Example.remark"
var Sudo Method = "Sudo.sudo"