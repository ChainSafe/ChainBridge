package prometheus

import (
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rcrowley/go-metrics"

	"github.com/ChainSafe/ChainBridge/core"
	msg "github.com/ChainSafe/ChainBridge/message"
	log "github.com/ChainSafe/log15"
)

var (
	erc20deposits          = metrics.NewCounter()
	erc721deposits         = metrics.NewCounter()
	totalDepositsProcessed = metrics.NewCounter()
	timeRunning            = metrics.GetOrRegisterTimer("time-running", nil)
)

func recordDepositMetrics() {
	metrics.Register("erc20", erc20deposits)

	go func() {
		for {
			// increments the counter by 1 every 1 sec
			// add dynamic counter from ethereum/chain events
			erc20deposits.Inc(10)
			erc721deposits.Inc(11)
			totalDepositsProcessed.Inc(11)

			timeRunning.Time(func() {

			})
			timeRunning.Update(11)

			time.Sleep(1 * time.Second)
		}
	}()
}

type HttpMetricServer struct {
	port         int
	core         *core.Core
	blockheights map[msg.ChainId]BlockHeightInfo
}

type BlockHeightInfo struct {
	height      *big.Int
	lastUpdated time.Time
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
		port:         opts.port,
		core:         opts.core,
		blockheights: make(map[msg.ChainId]BlockHeightInfo),
	}
}

// Start starts the http metrics server
func (s HttpMetricServer) Start() {
	log.Info("Prometheus metrics server started", "port", s.port)

	recordDepositMetrics()

	//NOTE: Be sure to unregister short-lived meters and timers otherwise they will leak memory:
	// Will call Stop() on the Meter to allow for garbage collection
	//metrics.Unregister("time-running")

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":"+strconv.Itoa(s.port), nil)
}
