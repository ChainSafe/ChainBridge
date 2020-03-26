package ethereum

import (
	"math/big"

	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type DepositProposal struct {
	DestinationChainID *big.Int
	DepositNonce       *big.Int
	DataHash           [32]byte
	NumYes             *big.Int
	NumNo              *big.Int
	Status             string
}

// Queries the contract for the current deposit status
func (w *Writer) GetDepositStatus(originChain *big.Int, depositNonce *big.Int) (uint8, error) {
	out := new(DepositProposal)

	err := w.bridgeContract.Call(new(bind.CallOpts), out, "getDepositProposal", originChain, depositNonce)
	if err != nil {
		log15.Error("Failed to call getDepositProposal", "error", out)
		return 0, err
	}

	return out.Status, err
}
