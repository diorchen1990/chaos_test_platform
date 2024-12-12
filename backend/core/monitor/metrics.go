package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsCollector struct {
	// 实验相关指标
	experimentCounter   *prometheus.CounterVec   // 实验总数
	experimentDuration  *prometheus.HistogramVec // 实验执行时长
	experimentSuccess   *prometheus.CounterVec   // 成功实验数
	experimentFailure   *prometheus.CounterVec   // 失败实验数
	
	// 系统资源指标  
	resourceUsage      *prometheus.GaugeVec    // CPU/内存/磁盘使用率
	
	// 应用指标
	applicationMetrics *prometheus.GaugeVec    // JVM/DB等应用指标
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
		mc.resourceUsage,
		mc.applicationMetrics,
	)
	
	return mc, nil
} 