package collector

import (
	"github.com/istio/pkg/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/zhanglianx111/harbor-exporter/pkgs/config"
	"github.com/zhanglianx111/harbor-exporter/pkgs/harbor/client"
	"sync"
)

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
			// gauge
			"harbor_gauge_metric": newGlobalMetric(namespace, "status",  "Current the status of harbor", []string{"component","host"}),
			"registry_gauge_metric": newGlobalMetric(namespace, "registry_status",  "Current the status of registry in harbor", []string{"component", "host"}),
			"redis_gauge_metric": newGlobalMetric(namespace, "redis_status",  "Current the status of redis in harbor", []string{"component", "host"}),
			"core_gauge_metric": newGlobalMetric(namespace, "core_status","Current the status of core in harbor", []string{"component", "host"}),
			"portal_gauge_metric": newGlobalMetric(namespace, "portal_status","Current the status of portal in harbor", []string{"component", "host"}),
			"jobservice_gauge_metric": newGlobalMetric(namespace, "jobservice_status","Current the status of jobservice in harbor", []string{"component", "host"}),
			"database_gauge_metric": newGlobalMetric(namespace, "database_status","Current the status of database in harbor", []string{"component", "host"}),
			"registryctl_gauge_metric": newGlobalMetric(namespace, "registryctl_status","Current the status of registryctl in harbor", []string{"component", "host"}),
			"harbor_gauge_volume": newGlobalMetric(namespace, "volume", "Current the size of volume", []string{"size", "host"}),
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

	/* TODO 加入counter类型的metrics */
	/*
	for host, currentValue := range mockCounterMetricData {
		ch <-prometheus.MustNewConstMetric(c.metrics["my_counter_metric"], prometheus.CounterValue, float64(currentValue), host)
	}
	*/


	volumeGaugeMetricsData := c.GetVolumeInfo()
	for name, value := range volumeGaugeMetricsData {
		ch <-prometheus.MustNewConstMetric(c.metrics["harbor_gauge_volume"], prometheus.GaugeValue, float64(value), name, config.Config.Harbor)
	}

	harborGaugeMetricsData := c.GetHarborStatus()
		for comp, currentValue := range harborGaugeMetricsData {
			switch comp {
			case "harbor":
				ch <-prometheus.MustNewConstMetric(c.metrics["harbor_gauge_metric"], prometheus.GaugeValue, float64(currentValue), comp, config.Config.Harbor)
			case "core":
				ch <-prometheus.MustNewConstMetric(c.metrics["core_gauge_metric"], prometheus.GaugeValue, float64(currentValue), comp, config.Config.Harbor)
			case "registry":
				ch <-prometheus.MustNewConstMetric(c.metrics["registry_gauge_metric"], prometheus.GaugeValue, float64(currentValue), comp, config.Config.Harbor)
			case "registryctl":
				ch <-prometheus.MustNewConstMetric(c.metrics["registryctl_gauge_metric"], prometheus.GaugeValue, float64(currentValue), comp, config.Config.Harbor)
			case "database":
				ch <-prometheus.MustNewConstMetric(c.metrics["database_gauge_metric"], prometheus.GaugeValue, float64(currentValue), comp, config.Config.Harbor)
			case "jobservice":
				ch <-prometheus.MustNewConstMetric(c.metrics["jobservice_gauge_metric"], prometheus.GaugeValue, float64(currentValue), comp, config.Config.Harbor)
			case "redis":
				ch <-prometheus.MustNewConstMetric(c.metrics["redis_gauge_metric"], prometheus.GaugeValue, float64(currentValue), comp, config.Config.Harbor)
			}
		}

}


/*
	harbor健康状态metrics
*/
func (c *Metrics) GetHarborStatus() (harborGaugeMetricsData map[string]int8) {
	//harborGaugeMetricsData = new(map[string]int8)
	status := harbor.GetHealthStatus()
	harborGaugeMetricsData = map[string]int8 {}
	for comp, value := range status {
		log.Debugf("comp:%s, valule:%d\n", comp, value)
		harborGaugeMetricsData[comp] = value
	}
	return
}

/*
	harbor使用磁盘情况
*/
func (c *Metrics) GetVolumeInfo() (volumeGaugeMetricsData map[string]uint64) {
	volume := harbor.GetVolumeInfo()
	volumeGaugeMetricsData = map[string]uint64{}
	for name, value := range volume {
		log.Debugf("%s: %d\n", name, value)
		volumeGaugeMetricsData[name] = value
	}
	return
}



