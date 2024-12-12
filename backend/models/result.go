package models

type ExperimentResult struct {
    ExperimentID string    `json:"experimentId"`
    StartTime    int64     `json:"startTime"`
    EndTime      int64     `json:"endTime"`
    Success      bool      `json:"success"`
    Metrics      map[string]interface{} `json:"metrics"`
    Logs         []string  `json:"logs"`
} 