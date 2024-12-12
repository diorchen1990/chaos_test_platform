package models

import (
    "time"
    "encoding/json"
)

type Report struct {
    ID            string          `json:"id" gorm:"primaryKey"`
    Name          string          `json:"name"`
    ExperimentID  string          `json:"experimentId"`
    TestCaseID    string          `json:"testCaseId"`
    StartTime     time.Time       `json:"startTime"`
    EndTime       time.Time       `json:"endTime"`
    Status        string          `json:"status"`
    Summary       ReportSummary   `json:"summary" gorm:"type:json"`
    Steps         []ReportStep    `json:"steps" gorm:"type:json"`
    Metrics       []MetricData    `json:"metrics" gorm:"type:json"`
    Logs          []LogEntry      `json:"logs" gorm:"type:json"`
    Creator       string          `json:"creator"`
}

type ReportSummary struct {
    TotalSteps    int     `json:"totalSteps"`
    SuccessSteps  int     `json:"successSteps"`
    FailedSteps   int     `json:"failedSteps"`
    Duration      float64 `json:"duration"`
    ErrorCount    int     `json:"errorCount"`
    AffectedNodes []string `json:"affectedNodes"`
}

type ReportStep struct {
    Name          string    `json:"name"`
    Status        string    `json:"status"`
    StartTime     time.Time `json:"startTime"`
    EndTime       time.Time `json:"endTime"`
    Action        string    `json:"action"`
    Target        string    `json:"target"`
    Error         string    `json:"error,omitempty"`
    Screenshots   []string  `json:"screenshots,omitempty"`
}

type MetricData struct {
    Name      string    `json:"name"`
    Labels    map[string]string `json:"labels"`
    Values    []MetricPoint `json:"values"`
}

type MetricPoint struct {
    Timestamp int64   `json:"timestamp"`
    Value     float64 `json:"value"`
}

type LogEntry struct {
    Timestamp time.Time `json:"timestamp"`
    Level     string    `json:"level"`
    Message   string    `json:"message"`
    Source    string    `json:"source"`
} 