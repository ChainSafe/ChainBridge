package ethereum

import (
	"math/big"

	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// Queries the contract for the current deposit status
func (w *Writer) DepositStatus(depositId *big.Int, originChain *big.Int) (uint8, error) {
	ret := new(struct {
		Hash        [32]byte
		Id          *big.Int
		OriginChain *big.Int
		NumYes      *big.Int
		NumNo       *big.Int
		Status      uint8
	})
	out := ret
	err := w.receiverContract.Call(new(bind.CallOpts), out, "DepositProposals", depositId, originChain)
	if err != nil {
		log15.Error("failed to call DepositProposal", "error", err)
		return 0, err
	}

	return out.Status, err
}
