package metrics

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
