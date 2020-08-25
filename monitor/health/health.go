// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package health

import "github.com/ChainSafe/ChainBridge/core"

// Start spins up the metrics server
func Start(core *core.Core) {
	// TODO add to config file
	httpPort := 8000
	blockTimeDelay := 120
	httpEnabled := true

	if httpEnabled {
		opts := httpMetricOptions{
			port:      httpPort,
			timeDelay: blockTimeDelay,
			core:      core,
		}
		httpServer := newhttpMetricServer(opts)
		httpServer.Start()
	}
}
