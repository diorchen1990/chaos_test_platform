# ChaosBlade 混沌测试平台

基于 ChaosBlade 的分布式混沌工程测试平台，提供可视化界面来管理和执行混沌实验。

## 目录
- [功能特性](#功能特性)
- [快速开始](#快速开始)
- [系统架构](#系统架构)
- [安装部署](#安装部署)
  - [环境要求](#环境要求)
  - [快速部署](#快速部署)
  - [手动部署](#手动部署)
  - [验证部署](#验证部署)
- [配置说明](#配置说明)
  - [核心配置](#核心配置)
  - [环境变量](#环境变量)
  - [权限要求](#权限要求)
- [探针管理](#探针管理)
  - [安装方式](#安装方式)
  - [配置说明](#配置说明)
  - [验证安装](#验证安装)
- [ChaosBlade](#chaosblade)
  - [安装配置](#安装配置)
  - [使用说明](#使用说明)
  - [注意事项](#注意事项)
- [常见问题](#常见问题)
- [开发指南](#开发指南)

> **重要提示：**
> 1. 在生产环境使用前，请确保完整阅读配置说明和注意事项
> 2. 建议先在测试环境进行充分验证
> 3. 请妥善保管各类密钥和密码信息

## 功能特性

### 核心功能
- 🌐 分布式系统故障注入
- 📊 可视化实验管理
- 🔍 全方位监控能力
- 📝 用例管理与复用
- 🚀 自动化部署支持

### 故障类型支持
- **Java应用故障**
  - 方法延迟注入
  - 异常注入
  - 线程池故障
  - GC压力模拟
  - 内存泄漏模拟

- **数据库故障**
  - 连接超时
  - 慢查询模拟
  - 连接池满载
  - SQL执行异常

- **分布式故障**
  - 服务调用延迟
  - 网络分区模拟
  - 消息延迟投递
  - RPC调用异常

## 快速开始

### 环境要求
```bash
# 必需组件
- Docker 20.10+
- Docker Compose 2.0+
- MySQL 8.0+
- Prometheus 2.30+
- Grafana 8.0+

# 系统要求
- CPU: 4核+
- 内存: 8GB+
- 磁盘: 50GB+
- 操作系统: Linux (CentOS 7+/Ubuntu 18.04+)
```
### 安装前准备
```bash
# 1. 系统环境检查
df -h
free -h
nproc

# 2. 安装依赖
# CentOS
yum install -y wget curl git docker

# Ubuntu
apt-get update
apt-get install -y wget curl git docker.io

# 3. 配置Docker
systemctl enable docker
systemctl start docker
```

### 快速部署
```bash
# 1. 克隆项目
git clone https://github.com/diorchen1990/chaos_test_platform
cd chaos-platform

# 2. 配置环境变量
cp .env.example .env
vim .env

# 3. 启动服务
docker-compose up -d

# 4. 验证部署
curl http://localhost:8080/api/health
```

## 系统架构

### 组件说明
1. **控制平台**
   - Web界面：实验管理、监控展示
   - API服务：提供RESTful接口
   - 调度系统：实验编排和执行
   - 监控系统：指标采集和展示

2. **Agent组件**
   - 故障注入执行器
   - 指标采集器
   - 状态管理器

3. **存储系统**
   - MySQL：元数据存储
   - Prometheus：监控数据
   - Elasticsearch：日志存储

## 配置说明

### 核心配置文件
```yaml
# backend/config/config.yaml
server:
  port: 8080
  host: 0.0.0.0
  debug: false

database:
  driver: mysql
  host: localhost
  port: 3306
  username: your_username
  password: your_password

chaosblade:
  path: /usr/local/bin/blade
  prepare:
    jvm: true
    docker: true
    kubernetes: false
```

### 环境变量
```bash
# 必需的环境变量
MYSQL_ROOT_PASSWORD=xxx     # MySQL root密码
MYSQL_DATABASE=chaos        # 数据库名
DB_USER=xxx                # 数据库用户
DB_PASSWORD=xxx            # 数据库密码

# 可选的环境变量
LOG_LEVEL=info             # 日志级别
PROMETHEUS_URL=xxx         # Prometheus地址
```
## 自定义配置说明

### 1. 项目配置替换

#### 1.1 包导入路径
在所有 Go 文件中需要替换包导入路径：

```go
// 替换前
import (
    "your-project/models"
    "your-project/services"
    "your-project/pkg/errors"
)

// 替换后
import (
    "github.com/chaos-mesh/chaos-platform/models"
    "github.com/chaos-mesh/chaos-platform/services"
    "github.com/chaos-mesh/chaos-platform/pkg/errors"
)
```

#### 1.2 Docker镜像
在 docker-compose.yml 中替换镜像名称：

```yaml
services:
  backend:
    image: <your-registry>/chaos-platform-backend:latest  # 替换为实际的镜像仓库地址
  frontend:
    image: <your-registry>/chaos-platform-frontend:latest # 替换为实际的镜像仓库地址
```

#### 1.3 服务地址
在配置文件中替换服务地址：

```yaml
# backend/config/config.yaml
server:
  host: <your-host>          # 替换为实际的服务器地址
  port: 8080

monitor:
  prometheus:
    url: <prometheus-url>    # 替换为实际的Prometheus地址
  grafana:
    url: <grafana-url>       # 替换为实际的Grafana地址

alert:
  dingtalk:
    webhook: <webhook-url>   # 替换为实际的钉钉webhook地址
```

### 2. 自定义开发

#### 2.1 故障类型
在 `backend/models/experiment.go` 中添加新的故障类型：

```go
// 添加自定义故障类型
const (
    // Java应用故障
    ActionJavaException   = "java-exception"
    ActionJavaLatency    = "java-latency"
    ActionJavaOOM        = "java-oom"
    
    // 数据库故障
    ActionDBTimeout      = "db-timeout"
    ActionDBError       = "db-error"
    
    // 自定义故障类型
    ActionCustomFault    = "custom-fault"
)

// 实现故障处理器
func (h *ExperimentHandler) handleCustomFault(ctx context.Context, exp *models.Experiment) error {
    // 实现自定义故障注入逻辑
    return nil
}
```

#### 2.2 监控指标
在 `backend/core/monitor/metrics.go` 中添加自定义指标：

```go
var (
    // 实验相关指标
    experimentExecutions = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "chaos_experiment_executions_total",
            Help: "Total number of chaos experiments executed",
        },
        []string{"type", "status"},
    )

    // 自定义业务指标
    customMetrics = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "chaos_custom_metric",
            Help: "Custom business metrics",
        },
        []string{"service", "method"},
    )
)

func init() {
    // 注册指标
    prometheus.MustRegister(experimentExecutions)
    prometheus.MustRegister(customMetrics)
}
```

#### 2.3 探针实现
在 `agent/probe/custom_probe.go` 中实现自定义探针：

```go
type CustomProbe struct {
    config ProbeConfig
    logger *zap.Logger
}

func (p *CustomProbe) Init(config ProbeConfig) error {
    p.config = config
    return nil
}

func (p *CustomProbe) InjectFault(fault Fault) error {
    // 实现故障注入逻辑
    switch fault.Type {
    case "custom-fault":
        return p.injectCustomFault(fault)
    default:
        return errors.New("unsupported fault type")
    }
}

func (p *CustomProbe) injectCustomFault(fault Fault) error {
    // 实现自定义故障注入
    return nil
}
```

### 3. 配置项说明

#### 3.1 实验配置
```yaml
experiment:
  timeout: 300              # 实验超时时间(秒)
  concurrent: 10            # 最大并发数
  retries: 3               # 重试次数
  workDir: /opt/chaos/experiments  # 工作目录

recovery:
  auto: true               # 是否自动恢复
  timeout: 60              # 恢复超时时间(秒)
  retries: 3               # 恢复重试次数
```

#### 3.2 监控配置
```yaml
monitor:
  prometheus:
    scrapeInterval: 15s    # 采集间隔
    retentionTime: 15d     # 数据保留时间
  grafana:
    dashboardDir: /opt/chaos/dashboards  # 面板目录
```

#### 3.3 告警配置
```yaml
alert:
  enable: true
  providers:
    - type: email
      host: smtp.example.com
      port: 587
      username: your-email
      password: your-password
    - type: dingtalk
      webhook: https://oapi.dingtalk.com/robot/send?access_token=xxx
```

### 4. 注意事项

1. 权限控制
   - 生产环境必须启用认证
   - 实验执行需要审批
   - 关键操作需要二次确认

2. 数据安全
   - 定期备份实验数据
   - 加密敏感配置信息
   - 保留操作审计日志

3. 监控告警
   - 配置合理的告警阈值
   - 设置告警升级策略
   - 确保告警及时送达

4. 故障恢复
   - 验证恢复机制可用性
   - 准备手动恢复方案
   - 记录恢复操作步骤
## 故障排查指南

### 1. 服务启动问题
```bash
# 检查服务状态
docker-compose ps

# 查看服务日志
docker-compose logs -f backend
docker-compose logs -f frontend

# 检查配置
cat config/config.yaml
```

### 2. Agent连接问题
```bash
# 检查Agent状态
curl http://localhost:8080/api/agents/status

# 查看Agent日志
tail -f /var/log/chaos-agent.log
```

### 3. 实验执行问题
```bash
# 检查实验状态
curl http://localhost:8080/api/experiments/{id}

# 查看ChaosBlade日志
blade status
```

## 安全建议

### 1. 访问控制
- 使用强密码和密钥
- 启用HTTPS
- 配置防火墙规则
- 实施IP白名单

### 2. 数据安全
- 定期备份数据
- 加密敏感信息
- 实施审计日志
- 监控异常访问

## 维护指南

### 1. 日常维护
```bash
# 清理日志
find /var/log/chaos-platform -name "*.log" -mtime +30 -delete

# 备份数据
mysqldump -u root -p chaos_platform > backup.sql

# 更新系统
docker-compose pull
docker-compose up -d
```

### 2. 监控检查
- 定期检查系统资源使用情况
- 监控服务运行状态
- 检查日志是否正常轮转
- 验证备份是否可用

## 开发指南

1. 后端开发
```bash
# 安装依赖
make setup

# 生成代码
make generate

# 运行测试
make test
```

### 2. 本地开发
```bash
# 启动后端服务
make run

# 启动前端服务
cd web && npm run serve
```

### 3. 构建部署
```bash
# 构建二进制
make build

# 构建Docker镜像
make docker

# 部署服务
docker-compose up -d
```

### 4. 代码规范
- 遵循 [Go 代码规范](https://golang.org/doc/effective_go)
- 使用 `gofmt` 格式化代码
- 运行 `golangci-lint` 检查代码质量
```

## 维护者

- 维护者：[姓名]
- 邮箱：[邮箱]

## 开源协议

MIT License
