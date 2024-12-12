package scheduler

import (
	"context"
	"sync"
	"time"
)

type Scheduler struct {
	experiments sync.Map
	agents      sync.Map
	executor    ExperimentExecutor
}

func (s *Scheduler) ScheduleExperiment(exp *models.Experiment) error {
	// 1. 查找可用 Agent
	agents := s.findAvailableAgents(exp.Target)
	if len(agents) == 0 {
		return ErrNoAvailableAgent
	}

	// 2. 分发实验任务
	for _, agent := range agents {
		if err := s.dispatchToAgent(agent, exp); err != nil {
			continue
		}
	}

	// 3. 更新实验状态
	return s.updateExperimentStatus(exp.ID, "running")
}

func (s *Scheduler) dispatchToAgent(agent *Agent, exp *models.Experiment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return agent.ExecuteExperiment(ctx, exp)
} 