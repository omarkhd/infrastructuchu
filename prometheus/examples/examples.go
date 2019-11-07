package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	pa "github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	e := newExamples()
	e.serveMetrics()
	e.emitCounters()
	select {}
}

type examples struct {
	metrics struct {
		requestsTotal *prometheus.CounterVec
	}
	metricsPort int
}

func newExamples() *examples {
	e := &examples{
		metricsPort: 5100,
	}
	e.metrics.requestsTotal = pa.NewCounterVec(
		prometheus.CounterOpts{
			Name: "examples_requests_total",
			Help: "Total number of requests",
		},
		[]string{"endpoint", "host", "zone"},
	)
	return e
}

func (e *examples) serveMetrics() {
	go func() {
		log.Printf("Starting metrics HTTP server on port %d", e.metricsPort)
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(fmt.Sprintf(":%d", e.metricsPort), nil)
	}()
}
