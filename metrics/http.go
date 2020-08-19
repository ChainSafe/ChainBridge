package metrics

import (
	"net/http"
	"strconv"

	"github.com/ChainSafe/ChainBridge/core"
	log "github.com/ChainSafe/log15"
)

type HttpMetricServer struct {
	port int
	core *core.Core
}

type httpMetricOptions struct {
	port int
	core *core.Core
}

type chain struct {
	id int
}

func newHTTPMetricServer(opts httpMetricOptions) *HttpMetricServer {
	return &HttpMetricServer{
		port: opts.port,
		core: opts.core,
	}
}

// Start starts the http metrics server
func (s HttpMetricServer) Start() {
	log.Info("Metrics server started", "port", s.port)

	// Setup routes
	http.HandleFunc("/health", s.healthStatus)

	// Start http server
	// TODO Push strconv to cli parser
	http.ListenAndServe(":"+strconv.Itoa(s.port), nil)
}

func (s HttpMetricServer) healthStatus(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Internal Service Error"))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - Operational"))
}
