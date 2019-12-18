package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EventSig string

func (en *EventSig) GetTopic() common.Hash {
	return crypto.Keccak256Hash([]byte(*en))
}
