// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only
/*
Provides the command-line interface for the chainbridge application.

For configuration and CLI commands see the README: https://github.com/ChainSafe/ChainBridge.
*/
package health

// Start spins up the metrics server
func Start() {
	// TODO add to config
	httpPort := 8000
	httpEnabled := true

	if httpEnabled {
		opts := &httpMetricOptions{
			port: httpPort,
		}
		httpServer := newHTTPMetricServer(*opts)
		httpServer.Start()
	}

	// TODO add prometheus
}
