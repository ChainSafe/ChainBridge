package ethereum

import (
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/contracts/Receiver"
	"github.com/ChainSafe/log15"
)

func (w *Writer) DepositComplete(depositId big.Int, originChain big.Int) (bool, error) {
	caller, err := Receiver.NewReceiverCaller(w.cfg.receiver, w.conn.conn)
	if err != nil {
		log15.Error("failed to init caller")
		return false, err
	}

	caller.DepositProposals()
	return true, nil
}
