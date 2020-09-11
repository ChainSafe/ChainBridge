// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package types

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type ChainMetrics struct {
	BlocksProcessed prometheus.Counter
	VotesSubmitted  prometheus.Counter
}

func NewChainMetrics(chain string) *ChainMetrics {
	metrics := &ChainMetrics{
		BlocksProcessed: prometheus.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_latest_block", chain),
			Help: "Number of blocks processed by the chain's listener",
		}),
		VotesSubmitted: prometheus.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_votes_submitted", chain),
			Help: "Number of votes submitted to chain",
		}),
	}

	prometheus.MustRegister(metrics.BlocksProcessed)
	prometheus.MustRegister(metrics.VotesSubmitted)

	return metrics
}
