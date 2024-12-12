package main

import (
    "context"
    "github.com/chaos-mesh/chaos-platform/pkg/config"
    "github.com/chaos-mesh/chaos-platform/pkg/logger"
    "github.com/chaos-mesh/chaos-platform/api"
    "github.com/chaos-mesh/chaos-platform/core/monitor"
)

func main() {
    // 1. 加载配置
    cfg, err := config.Load()
    if err != nil {
        panic(err)
    }

    // 2. 初始化日志
    logger := logger.NewLogger(cfg.Log)
    defer logger.Sync()

    // 3. 初始化监控
    metrics, err := monitor.NewMetricsCollector(cfg.Monitor)
    if err != nil {
        logger.Fatal("Failed to initialize metrics collector", err)
    }

    // 4. 启动服务
    server := api.NewServer(cfg, logger, metrics)
    if err := server.Start(); err != nil {
        logger.Fatal("Failed to start server", err)
    }
} 