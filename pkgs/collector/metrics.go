package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/zhanglianx111/harbor-exporter/pkgs/harbor/client"
	"math/rand"
	"sync"
)
/*
var HarborStatus = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "harbor_status",
		Help: "Current the status of harbor.",
	}
	)

var RedisStatus = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "redis_status",
		Help: "Current the status of redis in harbor",
}

	)

/*
hdFailures := prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hd_errors_total",
		Help: "Number of hard-disk errors.",
	},
	[]string{"device"},
)
*/

type Metrics struct {
	metrics map[string]*prometheus.Desc
	mutex	sync.Mutex
}

func newGlobalMetric(namespace , metricName , docString string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(namespace+"_"+metricName, docString, labels, nil)
}

/* TODO 支持更多的label */
func NewMetrics(namespace string) *Metrics {
	return &Metrics{
		metrics: map[string]*prometheus.Desc{
			"harbor_gauge_metric": newGlobalMetric(namespace, "status",  "Current the status of harbor", []string{"component"}),
			"registry_gauge_metric": newGlobalMetric(namespace, "registry_status",  "Current the status of registry in harbor", []string{"component"}),
			"redis_gauge_metric": newGlobalMetric(namespace, "redis_status",  "Current the status of redis in harbor", []string{"component"}),
			"core_gauge_metric": newGlobalMetric(namespace, "core_status","Current the status of core in harbor", []string{"component"}),
			"portal_gauge_metric": newGlobalMetric(namespace, "portal_status","Current the status of portal in harbor", []string{"component"}),
			"jobservice_gauge_metric": newGlobalMetric(namespace, "jobservice_status","Current the status of jobservice in harbor", []string{"component"}),
			"database_gauge_metric": newGlobalMetric(namespace, "database_status","Current the status of database in harbor", []string{"component"}),
			"registryctl_gauge_metric": newGlobalMetric(namespace, "registryctl_status","Current the status of registryctl in harbor", []string{"component"}),
		},
	}
}

/**
 * 接口：Describe
 * 功能：传递结构体中的指标描述符到channel
 */
func (c *Metrics) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		ch <- m
	}
}

/**
 * 接口：Collect
 * 功能：抓取最新的数据，传递给channel
 */
func (c *Metrics) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()  // 加锁
	defer c.mutex.Unlock()

	//_, mockGaugeMetricData := c.GenerateMockData()
	/* TODO 加入counter类型的metrics */
	/*
	for host, currentValue := range mockCounterMetricData {
		ch <-prometheus.MustNewConstMetric(c.metrics["my_counter_metric"], prometheus.CounterValue, float64(currentValue), host)
	}

	for host, currentValue := range mockGaugeMetricData {
		ch <-prometheus.MustNewConstMetric(c.metrics["harbor_gauge_metric"], prometheus.GaugeValue, float64(currentValue), host)
	}
	*/
	harborGaugeMetricsData := c.GetHarborStatus()
		for comp, currentValue := range harborGaugeMetricsData {
			switch comp {
			case "harbor":
				ch <-prometheus.MustNewConstMetric(c.metrics["harbor_gauge_metric"], prometheus.GaugeValue, float64(currentValue), comp)
			case "core":
				ch <-prometheus.MustNewConstMetric(c.metrics["core_gauge_metric"], prometheus.GaugeValue, float64(currentValue), comp)
			case "registry":
				ch <-prometheus.MustNewConstMetric(c.metrics["registry_gauge_metric"], prometheus.GaugeValue, float64(currentValue), comp)
			case "registryctl":
				ch <-prometheus.MustNewConstMetric(c.metrics["registryctl_gauge_metric"], prometheus.GaugeValue, float64(currentValue), comp)
			case "database":
				ch <-prometheus.MustNewConstMetric(c.metrics["database_gauge_metric"], prometheus.GaugeValue, float64(currentValue), comp)
			case "jobservice":
				ch <-prometheus.MustNewConstMetric(c.metrics["jobservice_gauge_metric"], prometheus.GaugeValue, float64(currentValue), comp)
			case "redis":
				ch <-prometheus.MustNewConstMetric(c.metrics["redis_gauge_metric"], prometheus.GaugeValue, float64(currentValue), comp)
			}
		}

}


/**
 * 函数：GenerateMockData
 * 功能：生成模拟数据
 */
func (c *Metrics) GenerateMockData() (mockCounterMetricData map[string]int, mockGaugeMetricData map[string]int) {
	mockCounterMetricData = map[string]int{
		"yahoo.com": int(rand.Int31n(1000)),
		"google.com": int(rand.Int31n(1000)),
	}
	mockGaugeMetricData = map[string]int{
		"yahoo.com": int(rand.Int31n(10)),
		"google.com": int(rand.Int31n(10)),
	}
	return
}

/*
	harbor健康状态metrics
*/
func (c *Metrics) GetHarborStatus() (harborGaugeMetricsData map[string]int8) {
	//harborGaugeMetricsData = new(map[string]int8)
	status := harbor.GetHealthStatus()
	fmt.Println(status)
	harborGaugeMetricsData = map[string]int8 {}
	for comp, value := range status {
		fmt.Printf("comp:%s, valule:%d\n", comp, value)
		harborGaugeMetricsData[comp] = value
	}
	return
}




