package ethereum

import (
	"math/big"

	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	_, _, _, _, _, Status, err := w.bridgeContract.GetDepositProposal(new(bind.CallOpts), originChainId, depositNonce)
	if err != nil {
		log15.Error("Failed to call getDepositProposal", "error")
		return "inactive", err
	}

	threshold := big.NewInt(0)
	err2 := w.bridgeContract.Call(new(bind.CallOpts), &threshold, "getRelayerThreshold")
	if err2 != nil {
		log15.Error("Failed to fetch threshold", "err", err2)
	} else {
		log15.Trace("Got threshold", "threshold", threshold.String())
	}

	return DepositProposalStatus(Status), err
}

func (w *Writer) GetHasVotedOnDepositProposal(account common.Address, destinationChainId *big.Int, depositNonce *big.Int) (bool, error) {
	_, _, _, yesVotes, noVotes, _, err := w.bridgeContract.GetDepositProposal(new(bind.CallOpts), destinationChainId, depositNonce)
	if err != nil {
		log15.Debug("Errored")
		log15.Error("Failed to call getDepositProposal", "error")
		return false, err
	}

	log15.Error("Yes votes", "error", len(yesVotes))
	log15.Error("No votes", "error", len(noVotes))

	// TODO: Check if in yes & no via race condition or async
	for _, n := range yesVotes {
		if account == n {
			log15.Debug("found yes")
			return true, nil
		}
	}
	for _, n := range noVotes {
		if account == n {
			log15.Debug("found no")
			return true, nil
		}
	}
	log15.Debug("Not found")
	return false, nil
}
