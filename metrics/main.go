package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
	Name: "chainbridge_total_process_ops",
	Help: "The total number of ops process events",
})

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(1 * time.Second)
		}
	}()
}
func main()  {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":1111", nil)
}