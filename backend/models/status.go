package models

type ExperimentStatus struct {
    ID          string    `json:"id"`
    State       string    `json:"state"`
    Error       string    `json:"error,omitempty"`
    StartTime   int64     `json:"startTime"`
    EndTime     int64     `json:"endTime,omitempty"`
    Metrics     map[string]interface{} `json:"metrics,omitempty"`
} 