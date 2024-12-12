package models

import (
    "time"
    "encoding/json"
)

type TestCase struct {
    ID          string          `json:"id" gorm:"primaryKey"`
    Name        string          `json:"name"`
    Category    string          `json:"category"`
    Description string          `json:"description"`
    Tags        []string        `json:"tags" gorm:"type:json"`
    Experiment  Experiment      `json:"experiment" gorm:"type:json"`
    Creator     string          `json:"creator"`
    IsFavorite  bool           `json:"isFavorite"`
    UsageCount  int            `json:"usageCount"`
    CreateTime  time.Time       `json:"createTime"`
    UpdateTime  time.Time       `json:"updateTime"`
}

type TestCaseCategory struct {
    ID       string    `json:"id" gorm:"primaryKey"`
    Name     string    `json:"name"`
    ParentID string    `json:"parentId"`
    Level    int       `json:"level"`
    Sort     int       `json:"sort"`
} 