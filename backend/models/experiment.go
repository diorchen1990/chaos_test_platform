package models

import (
    "time"
    "encoding/json"
)

type ExperimentScope string

const (
    // 系统层面故障
    ScopeHost       ExperimentScope = "host"
    ScopeContainer  ExperimentScope = "container"
    ScopeK8s       ExperimentScope = "kubernetes"
    
    // 应用层面故障
    ScopeJavaApp    ExperimentScope = "java-app"
    ScopeDatabase   ExperimentScope = "database"
    ScopeMQ         ExperimentScope = "message-queue"
    ScopeRPC        ExperimentScope = "rpc"
)

// 故障类型定义
const (
    // Java应用故障类型
    ActionJavaException    = "java-exception"
    ActionJavaGC          = "java-gc"
    ActionJavaThreadPool  = "java-threadpool"
    ActionJavaCodeCoverage = "java-codecoverage"
    
    // 数据库故障类型
    ActionDBTimeout       = "db-timeout"
    ActionDBError        = "db-error"
    ActionDBSlow         = "db-slow"
    ActionDBConnectionLeak = "db-connection-leak"
    
    // 分布式故障类型
    ActionRPCTimeout     = "rpc-timeout"
    ActionRPCError      = "rpc-error"
    ActionMQDelay       = "mq-delay"
    ActionNetworkPartition = "network-partition"
)

type Experiment struct {
    ID          string          `json:"id" gorm:"primaryKey"`
    Name        string          `json:"name"`
    Scope       ExperimentScope `json:"scope"`
    Target      Target          `json:"target" gorm:"type:json"`
    Action      string          `json:"action"`
    Parameters  json.RawMessage `json:"parameters" gorm:"type:json"`
    Status      string          `json:"status"`
    CreateTime  time.Time       `json:"createTime"`
    UpdateTime  time.Time       `json:"updateTime"`
    Creator     string          `json:"creator"`
    AgentIDs    []string        `json:"agentIds" gorm:"type:json"`
    Results     []ExperimentResult `json:"results" gorm:"type:json"`
}

type Target struct {
    Type        string            `json:"type"`
    Endpoint    string            `json:"endpoint"`
    Labels      map[string]string `json:"labels"`
    // 数据库特定字段
    Database    *DatabaseTarget   `json:"database,omitempty"`
    // Java应用特定字段
    JavaApp     *JavaAppTarget    `json:"javaApp,omitempty"`
    // 分布式系统特定字段
    Distributed *DistributedTarget `json:"distributed,omitempty"`
}

type DatabaseTarget struct {
    Type        string `json:"type"` // mysql, postgresql, mongodb等
    Host        string `json:"host"`
    Port        int    `json:"port"`
    Database    string `json:"database"`
    User        string `json:"user"`
    SQLPattern  string `json:"sqlPattern,omitempty"`
}

type JavaAppTarget struct {
    ProcessName string   `json:"processName"`
    Port       int      `json:"port"`
    Classes    []string `json:"classes,omitempty"`
    Methods    []string `json:"methods,omitempty"`
}

type DistributedTarget struct {
    ServiceName    string   `json:"serviceName"`
    Instances     []string `json:"instances"`
    APIPath       string   `json:"apiPath,omitempty"`
    Dependencies  []string `json:"dependencies,omitempty"`
}

type ExperimentResult struct {
    AgentID     string          `json:"agentId"`
    Status      string          `json:"status"`
    StartTime   time.Time       `json:"startTime"`
    EndTime     time.Time       `json:"endTime"`
    Metrics     json.RawMessage `json:"metrics"`
    Error       string          `json:"error,omitempty"`
} 