package centrifuge

import (
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type EventAssetTransfer struct {
	Phase       types.Phase
	Destination types.Bytes
	DepositID   types.Bytes
	To          types.Bytes
	TokenID     types.Bytes
	Metadata    types.Bytes
	Topics      []types.Hash
}

type Events struct {
	types.EventRecords
	Bridge_AssetTransfer []EventAssetTransfer //nolint:stylecheck,golint
}

func assetTransferHandler(evtI interface{}) msg.Message {
	evt, ok := evtI.(EventAssetTransfer)
	if !ok {
		log15.Error("failed to cast EventAssetTransfer type")
	}

	log15.Info("Got asset transfer event!", "destination", evt.Destination)

	return msg.Message{
		Source:      msg.CentrifugeId,
		Destination: msg.EthereumId,
		Type:        msg.DepositAssetType,
		Data:        evt.Destination[:], // TODO: Pack data
	}
}
