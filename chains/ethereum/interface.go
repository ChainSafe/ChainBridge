package ethereum

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type ReceiverContract interface {
	Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error
	Transfer(opts *bind.TransactOpts) (*types.Transaction, error)
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)
}

type EmitterContract interface {
	FilterERCTransfer(opts *bind.FilterOpts) (EventIterator, error)
	FilterGenericTransfer(opts *bind.FilterOpts) (EventIterator, error)
}

type EventIterator interface {
	Next() bool
}
