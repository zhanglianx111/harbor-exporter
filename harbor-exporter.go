package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zhanglianx111/harbor-exporter/pkgs/collector"
	"github.com/zhanglianx111/harbor-exporter/pkgs/harbor/client"
	_ "github.com/zhanglianx111/harbor-exporter/pkgs/config"
	"net/http"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	listenAddr       = flag.String("listen-port", "0.0.0.0:9001", "An port for getting metrics")
	metricsPath      = flag.String("metrics-path", "/metrics", "expose metrics url path")
	metricsNamespace = flag.String("metrics-namespace", "harbor", "Prometheus metrics namespace, as the prefix of metrics name")
//	harborUsername	 = flag.String("username", "superuser", "the supersuser in harbor")
//	harborPassword	 = flag.String("password", "password", "the password for superuser in harbor")
	harborAddr		 = flag.String("harbor-address", "http[s]://example.aibee.cn", "harbor address")

	log 			 = logrus.New()

)

func init() {
	/* TODO 单独出个log包 */
	log.Level = logrus.DebugLevel
	log.Out = os.Stdout

	harbor.Login()
}

func main() {
//	log.Out = os.Stdout
	flag.Parse()

	metrics := collector.NewMetrics(*metricsNamespace)
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)

	http.Handle(*metricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	log.Infof("Starting Server at http://%s%s", *listenAddr, *metricsPath)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
