package types

import (
	"math/big"
	"time"
)

type LatestBlock struct {
	Height      *big.Int
	LastUpdated time.Time
}
