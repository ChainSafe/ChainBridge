package centrifuge

import (
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

// Centrifuge event
type EventNFTDeposited struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

type Events struct {
	types.EventRecords
	Nfts_DepositAsset []EventNFTDeposited //nolint:stylecheck,golint
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
