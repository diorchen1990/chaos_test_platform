package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsCollector struct {
	// 实验相关指标
	experimentCounter *prometheus.CounterVec
	experimentDuration *prometheus.HistogramVec
	
	// 系统资源指标
	cpuUsage    *prometheus.GaugeVec
	memoryUsage *prometheus.GaugeVec
	diskUsage   *prometheus.GaugeVec
	
	// 业务指标
	javaMetrics  *prometheus.GaugeVec
	dbMetrics    *prometheus.GaugeVec
}

func NewMetricsCollector(cfg *Config) (*MetricsCollector, error) {
	mc := &MetricsCollector{
		experimentCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "chaos_experiment_total",
				Help: "Total number of chaos experiments",
			},
			[]string{"type", "status"},
		),
		// ... 其他指标初始化
	}
	
	// 注册所有指标
	prometheus.MustRegister(
		mc.experimentCounter,
		mc.experimentDuration,
		mc.cpuUsage,
		mc.memoryUsage,
		mc.diskUsage,
		mc.javaMetrics,
		mc.dbMetrics,
	)
	
	return mc, nil
} 