# ChaosBlade 混沌测试平台

基于 ChaosBlade 的分布式混沌工程测试平台，提供可视化界面来管理和执行混沌实验。

## 功能特性

- 实验管理：创建、执行、监控混沌实验
- 用例管理：复用实验配置，快速执行常用场景
- 监控分析：实时监控系统指标，分析实验影响
- 报告管理：自动生成实验报告，记录详细数据

## 快速开始

1. 环境要求
   - Docker 19.03+
   - Kubernetes 1.16+
   - Go 1.17+
   - Node.js 14+

2. 安装部署
```bash
# 克隆项目
git clone https://github.com/chaos-mesh/chaos-platform.git

# 安装依赖
make init

# 构建项目
make build

# 启动服务
docker-compose up -d
```

## 项目结构

```
chaos-platform/
├── api/                    # API定义
├── backend/               # 后端服务
├── agent/                 # Agent组件
├── frontend/             # 前端项目
├── deploy/               # 部署配置
└── docs/                 # 项目文档
```

## 开发指南

1. 后端开发
```bash
cd backend
go mod download
go run main.go
```

2. 前端开发
```bash
cd frontend
npm install
npm run serve
```

## 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交代码
4. 发起 Pull Request

## 开源协议

MIT License