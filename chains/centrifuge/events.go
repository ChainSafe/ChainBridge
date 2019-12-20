package centrifuge

import (
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

// Centrofuge event
type EventNFTDeposited struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

type Events struct {
	types.EventRecords
	Nfts_DepositAsset []EventNFTDeposited //nolint:stylecheck,golint
}
