// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package health

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/ChainSafe/ChainBridge/core"
	msg "github.com/ChainSafe/ChainBridge/message"
	log "github.com/ChainSafe/log15"
)

type httpMetricServer struct {
	port         int
	core         *core.Core
	blockheights map[msg.ChainId]blockHeightInfo
}

type blockHeightInfo struct {
	height      *big.Int
	lastUpdated time.Time
}

type httpMetricOptions struct {
	port int
	core *core.Core
}

func newhttpMetricServer(opts httpMetricOptions) *httpMetricServer {
	return &httpMetricServer{
		port:         opts.port,
		core:         opts.core,
		blockheights: make(map[msg.ChainId]blockHeightInfo),
	}
}

// Start starts the http metrics server
func (s httpMetricServer) Start() {
	log.Info("Metrics server started", "port", s.port)

	// Setup routes
	http.HandleFunc("/health", s.healthStatus)

	// Start http server
	err := http.ListenAndServe(":"+strconv.Itoa(s.port), nil)

	if err == http.ErrServerClosed {
		log.Info("Server is shutting down", err)
	} else {
		log.Error("Shutting down, server error: ", err)
	}
}

// healthStatus is a catch-all update that grabs the latest updates on the running chains
// It assumes that the configuration was set correctly, therefore the relevant chains are
// only those that are in the core.Core registry.
func (s httpMetricServer) healthStatus(w http.ResponseWriter, r *http.Request) {

	// Grab all chains
	chains := s.core.GetChains()
	requestTime := time.Now()

	// Iterate through their block heads and update the cache accordingly
	for _, chain := range chains {
		latestHeight, err := chain.GetLatestBlock()
		if err != nil {
			// TODO better error messaging
			errorMsg := fmt.Sprintf("%s%d%s%s", "Failed to receive latest head for: ", chain.Id(), "Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(errorMsg))
			if err != nil {
				log.Error("Failed to write, failed to get latest height and writing error", err)
			}
		}

		// Get old blockheight
		if prevHeight, ok := s.blockheights[chain.Id()]; ok {
			// Note The listener performing the polling should already be accounting for the block delay
			// TODO Account for timestamps
			timeDiff := requestTime.Sub(prevHeight.lastUpdated)
			if latestHeight.Cmp(prevHeight.height) >= 0 && timeDiff < 120 {
				s.blockheights[chain.Id()] = blockHeightInfo{
					height:      latestHeight,
					lastUpdated: requestTime,
				}
			} else {
				if timeDiff.Seconds() > 120 {
					// Error if we exceeded the time limit
					errorMsg := fmt.Sprintf("%s%s%s%s", "Chain height hasn't changed in: ", timeDiff, "Current height", latestHeight)
					w.WriteHeader(http.StatusInternalServerError)
					_, err = w.Write([]byte(errorMsg))
					if err != nil {
						log.Error("Failed to write, chain height hasn't changed in: %d seconds, Current height: %d", timeDiff.Seconds(), latestHeight, err)
					}
				} else {
					// Error for having a smaller blockheight than previous
					errorMsg := fmt.Sprintf("%s%s%s%s%s", "latestHeight is <= previousHeight", "previousHeight", prevHeight.height, "latestHeight", latestHeight)
					w.WriteHeader(http.StatusInternalServerError)
					_, err := w.Write([]byte(errorMsg))
					if err != nil {
						log.Error("Failed to write, latest height less than previous height, latest height: %d, previous height: %d", latestHeight, prevHeight.height, err)
					}
				}
			}
		} else {
			// Note: Could be edge case where chain never started, perhaps push initialization to different step
			// First time we've received a block for this chain
			s.blockheights[chain.Id()] = blockHeightInfo{
				height:      latestHeight,
				lastUpdated: requestTime,
			}
		}
	}
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("200 - Operational"))
	if err != nil {
		log.Error("Failed to write, 200 - Operational")
	}
}
