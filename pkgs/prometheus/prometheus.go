package prometheus

import "github.com/prometheus/client_golang/prometheus"

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
}