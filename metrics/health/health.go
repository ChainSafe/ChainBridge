// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package health

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ChainSafe/chainbridge-utils/msg"
	log "github.com/ChainSafe/log15"
)

// After this duration with no changes a chain will return an error
const BlockTimeout = 20

type httpMetricServer struct {
	port      int
	timeDelay int
	chains    []core.Chain
	stats     []ChainInfo
}

type httpMetricOptions struct {
	port      int
	timeDelay int
	chains    []core.Chain
}

type httpResponse struct {
	Chains []ChainInfo `json:"chains,omitempty"`
	Error  string      `json:"error,omitempty"`
}

type ChainInfo struct {
	ChainId     msg.ChainId `json:"chainId"`
	Height      *big.Int    `json:"height"`
	LastUpdated time.Time   `json:"lastUpdated"`
}

// Start spins up the metrics server
func Start(port int, chains []core.Chain) {
	opts := httpMetricOptions{
		port:      port,
		timeDelay: BlockTimeout,
		chains:    chains,
	}
	httpServer := newhttpMetricServer(opts)
	httpServer.Start()
}

func newhttpMetricServer(opts httpMetricOptions) *httpMetricServer {
	return &httpMetricServer{
		port:      opts.port,
		chains:    opts.chains,
		timeDelay: opts.timeDelay,
		stats:     make([]ChainInfo, len(opts.chains)),
	}
}

// Start starts the http metrics server
func (s httpMetricServer) Start() {
	log.Info("Health status server started listening on", "port", s.port)

	// Setup routes
	http.HandleFunc("/health", s.healthStatus)

	// Start http server
	err := http.ListenAndServe(":"+strconv.Itoa(s.port), nil)

	if err == http.ErrServerClosed {
		log.Info("Health status server is shutting down", err)
	} else {
		log.Error("Health status server shutting down, server error: ", err)
	}
}

// healthStatus is a catch-all update that grabs the latest updates on the running chains
// It assumes that the configuration was set correctly, therefore the relevant chains are
// only those that are in the core.Core registry.
func (s httpMetricServer) healthStatus(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Iterate through their block heads and update the cache accordingly
	for i, chain := range s.chains {
		current := chain.LatestBlock()
		prev := s.stats[i]
		if s.stats[i].LastUpdated.IsZero() {
			// First time we've received a block for this chain
			s.stats[chain.Id()] = ChainInfo{
				ChainId:     chain.Id(),
				Height:      current.Height,
				LastUpdated: current.LastUpdated,
			}
		} else {
			now := time.Now()
			timeDiff := now.Sub(prev.LastUpdated)
			// If block has changed, update it
			if current.Height.Cmp(prev.Height) == 1 {
				s.stats[chain.Id()].LastUpdated = current.LastUpdated
				s.stats[chain.Id()].Height = current.Height
			} else if int(timeDiff.Seconds()) >= s.timeDelay { // Error if we exceeded the time limit
				response := &httpResponse{
					Chains: []ChainInfo{},
					Error:  fmt.Sprintf("chain %d height hasn't changed for %f seconds. Current Height: %s", prev.ChainId, timeDiff.Seconds(), current.Height),
				}
				w.WriteHeader(http.StatusInternalServerError)
				err := json.NewEncoder(w).Encode(response)
				if err != nil {
					log.Error("Failed to write metrics", "err", err)
				}
				return
			} else if current.Height != nil && prev.Height != nil && current.Height.Cmp(prev.Height) == -1 { // Error for having a smaller blockheight than previous
				response := &httpResponse{
					Chains: []ChainInfo{},
					Error:  fmt.Sprintf("unexpected block height. previous = %s current = %s", prev.Height, current.Height),
				}
				w.WriteHeader(http.StatusInternalServerError)
				err := json.NewEncoder(w).Encode(response)
				if err != nil {
					log.Error("Failed to write metrics", "err", err)
				}
				return
			}
		}
	}

	response := &httpResponse{
		Chains: s.stats,
		Error:  "",
	}
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error("Failed to serve metrics")
	}
}
