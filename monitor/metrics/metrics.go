package metrics

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rcrowley/go-metrics"
	"net/http"
	"time"
)

var (
	erc20deposits = metrics.NewCounter()
	erc721deposits = metrics.NewCounter()
	totalDepositsProcessed = metrics.NewCounter()
	timeRunning = metrics.GetOrRegisterTimer("time-running", nil)
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


func main() {
	recordDepositMetrics()

	//NOTE: Be sure to unregister short-lived meters and timers otherwise they will leak memory:
	// Will call Stop() on the Meter to allow for garbage collection
	//metrics.Unregister("time-running")


	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":1111", nil)
}