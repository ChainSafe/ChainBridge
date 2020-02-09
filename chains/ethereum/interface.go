package ethereum

import (
	"math/big"

	emitter "github.com/ChainSafe/ChainBridgeV2/contracts/Emitter"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

// ReceiverContract provides the interface for Receiver.sol.
// Any bound contract should satisfy this interface.
type ReceiverContract interface {
	Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error
	Transfer(opts *bind.TransactOpts) (*types.Transaction, error)
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)
}

// EmitterContract provides the interface to filter events from Emitter.sol.
type EmitterContract interface {
	FilterERCTransfer(*bind.FilterOpts, []*big.Int, []*big.Int) (*emitter.EmitterERCTransferIterator, error)
	FilterGenericTransfer(*bind.FilterOpts, []*big.Int, []*big.Int) (*emitter.EmitterGenericTransferIterator, error)
	FilterNFTTransfer(opts *bind.FilterOpts, _destChain []*big.Int, _depositId []*big.Int) (*emitter.EmitterNFTTransferIterator, error)
}
