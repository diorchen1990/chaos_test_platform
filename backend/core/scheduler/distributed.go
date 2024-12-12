package scheduler

import (
    "context"
    "sync"
)

type DistributedScheduler struct {
    *Scheduler
    serviceRegistry ServiceRegistry
}

func (s *DistributedScheduler) ScheduleDistributedExperiment(exp *models.Experiment) error {
    // 1. 获取服务依赖关系
    deps := s.serviceRegistry.GetDependencies(exp.Target.Distributed.ServiceName)
    
    // 2. 创建分布式实验计划
    plan := s.createExperimentPlan(exp, deps)
    
    // 3. 按顺序执行实验
    return s.executeDistributedPlan(plan)
}

func (s *DistributedScheduler) executeDistributedPlan(plan *ExperimentPlan) error {
    var wg sync.WaitGroup
    errCh := make(chan error, len(plan.Steps))
    
    for _, step := range plan.Steps {
        wg.Add(1)
        go func(step *ExperimentStep) {
            defer wg.Done()
            if err := s.executeStep(step); err != nil {
                errCh <- err
            }
        }(step)
    }
    
    wg.Wait()
    close(errCh)
    
    // 检查错误
    for err := range errCh {
        if err != nil {
            return err
        }
    }
    
    return nil
} 