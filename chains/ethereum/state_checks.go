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

type DepositProposalStatus string

const (
	Inactive    DepositProposalStatus = "inactive"
	Active      DepositProposalStatus = "active"
	Denied      DepositProposalStatus = "denied"
	Passed      DepositProposalStatus = "passed"
	Transferred DepositProposalStatus = "transferred"
)

// Queries the contract for the current deposit status
func (w *Writer) GetDepositStatus(originChainId *big.Int, depositNonce *big.Int) (DepositProposalStatus, error) {
	out := new(DepositProposal)

	_, _, _, _, _, Status, err := w.bridgeContract.GetDepositProposal(new(bind.CallOpts), originChainId, depositNonce)
	if err != nil {
		log15.Error("Failed to call getDepositProposal", "error", out)
		return "inactive", err
	}

	return DepositProposalStatus(Status), err
}
