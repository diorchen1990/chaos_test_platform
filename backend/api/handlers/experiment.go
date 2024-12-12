package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExperimentHandler struct {
	scheduler *scheduler.Scheduler
	monitor   *monitor.MetricsCollector
}

func (h *ExperimentHandler) CreateExperiment(c *gin.Context) {
	var exp models.Experiment
	if err := c.ShouldBindJSON(&exp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. 验证实验参数
	if err := h.validateExperiment(&exp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. 调度实验
	if err := h.scheduler.ScheduleExperiment(&exp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3. 记录监控指标
	h.monitor.RecordExperiment(&exp)

	c.JSON(http.StatusOK, exp)
} 