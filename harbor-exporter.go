package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zhanglianx111/harbor-exporter/pkgs/collector"
	"log"
	"net/http"
)

var (
	listenAddr       = flag.String("listen-port", "9001", "An port for getting metrics")
	metricsPath      = flag.String("metrics-path", "/metrics", "expose metrics url path")
	metricsNamespace = flag.String("metrics-namespace", "harbor", "Prometheus metrics namespace, as the prefix of metrics name")
)

func main() {
	/*
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":8080", nil))
	*/

	flag.Parse()

	metrics := collector.NewMetrics(*metricsNamespace)
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)

	http.Handle(*metricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	log.Printf("Starting Server at http://localhost:%s%s", *listenAddr, *metricsPath)
	log.Fatal(http.ListenAndServe(":"+*listenAddr, nil))
}
