package centrifuge

import (
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type EventNFTDeposited struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

type EventAssetTransfer struct {
	Phase       types.Phase
	Destination types.Bytes
	DepositID   types.U32
	To          types.Bytes
	TokenID     types.Bytes
	Metadata    types.Bytes
	Topics      []types.Hash
}

type Events struct {
	types.EventRecords
	Nfts_DepositAsset    []EventNFTDeposited  //nolint:stylecheck,golint
	Bridge_AssetTransfer []EventAssetTransfer //nolint:stylecheck,golint
}

func nftHandler(evtI interface{}) msg.Message {
	evt, ok := evtI.(EventNFTDeposited)
	if !ok {
		log15.Error("failed to cast EventNFTDeposited type")
	}

	log15.Info("Got nfts event!", "evt", evt.Asset.Hex())

	return msg.Message{
		Source:      msg.CentrifugeId,
		Destination: msg.EthereumId,
		Type:        msg.DepositAssetType,
		Data:        evt.Asset[:],
	}
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
