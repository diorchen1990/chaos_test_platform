type MetricsCollector struct {
    // 实验相关指标
    experimentGauge   *prometheus.GaugeVec
    experimentCounter *prometheus.CounterVec
    
    // 系统资源指标
    cpuUsage    *prometheus.GaugeVec
    memoryUsage *prometheus.GaugeVec
    diskIO      *prometheus.GaugeVec
    
    // 应用指标
    jvmMetrics  *prometheus.GaugeVec
    dbMetrics   *prometheus.GaugeVec
    rpcMetrics  *prometheus.HistogramVec
}

func NewMetricsCollector() *MetricsCollector {
    return &MetricsCollector{
        experimentGauge: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Name: "chaos_experiment_status",
                Help: "Current status of chaos experiments",
            },
            []string{"experiment_id", "type"},
        ),
        // ... 其他指标初始化
    }
} 