package ethereum

import (
	"math/big"

	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type DepositProposal struct {
	Hash        [32]byte
	Id          *big.Int
	OriginChain *big.Int
	NumYes      *big.Int
	NumNo       *big.Int
	Status      uint8
}

// Queries the contract for the current deposit status
func (w *Writer) GetDepositStatus(originChain *big.Int, depositNonce *big.Int) (uint8, error) {
	out := new(DepositProposal)

	err := w.bridgeContract.Call(new(bind.CallOpts), out, "getDepositProposal", originChain, depositNonce)
	if err != nil {
		log15.Error("Failed to call getDepositProposal", "error", err)
		return 0, err
	}

	return out.Status, err
}
