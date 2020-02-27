package ethereum

import (
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/contracts/Receiver"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// Queries the contract for the current deposit status
func (w *Writer) DepositStatus(depositId *big.Int, originChain *big.Int) (uint8, error) {
	caller, err := Receiver.NewReceiverCaller(w.cfg.receiver, w.conn.conn)
	if err != nil {
		log15.Error("failed to init caller instance", "error", err)
		return 0, err
	}

	ret, err := caller.DepositProposals(new(bind.CallOpts), depositId, originChain)
	if err != nil {
		log15.Error("failed to call DepositProposal", "error", err)
		return 0, err
	}

	return ret.Status, nil
}
