.PHONY: init build test clean

# 版本信息
VERSION := $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')

# 构建标记
LDFLAGS := -X main.Version=$(VERSION) \
           -X main.GitCommit=$(COMMIT) \
           -X main.BuildTime=$(BUILD_TIME)

# 初始化
init:
	go mod download
	cd frontend && npm install

# 构建
build: build-backend build-agent build-frontend

build-backend:
	cd backend && go build -ldflags "$(LDFLAGS)" -o ../bin/chaos-controller

build-agent:
	cd agent && go build -ldflags "$(LDFLAGS)" -o ../bin/chaos-agent

build-frontend:
	cd frontend && npm run build

# 测试
test: test-backend test-frontend

test-backend:
	cd backend && go test -v -race ./...

test-frontend:
	cd frontend && npm run test

# 清理
clean:
	rm -rf bin/
	rm -rf frontend/dist/ 