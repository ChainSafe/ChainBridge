// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package health

// Start spins up the metrics server
func Start() {
	// TODO add to config file
	httpPort := 8000
	httpEnabled := true

	if httpEnabled {
		opts := &httpMetricOptions{
			port: httpPort,
		}
		httpServer := newhttpMetricServer(*opts)
		httpServer.Start()
	}
}
