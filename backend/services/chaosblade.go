package services

import (
    "os/exec"
    "path/filepath"
)

type ChaosBladeService struct {
    bladePath string
}

func NewChaosBladeService(bladePath string) *ChaosBladeService {
    return &ChaosBladeService{
        bladePath: bladePath,
    }
}

func (s *ChaosBladeService) ExecuteCommand(args ...string) error {
    cmd := exec.Command(s.bladePath, args...)
    return cmd.Run()
}

func (s *ChaosBladeService) PrepareJVM() error {
    return s.ExecuteCommand("prepare", "jvm")
}

func (s *ChaosBladeService) CreateExperiment(exp *models.Experiment) error {
    args := []string{"create"}
    
    // 根据实验类型构建参数
    switch exp.Scope {
    case models.ScopeJavaApp:
        args = append(args, "jvm")
        // 添加Java相关参数
    case models.ScopeContainer:
        args = append(args, "docker")
        // 添加容器相关参数
    }
    
    return s.ExecuteCommand(args...)
}

func (s *ChaosBladeService) DestroyExperiment(expID string) error {
    cmd := exec.Command("blade", "destroy", expID)
    _, err := cmd.CombinedOutput()
    return err
}

func (s *ChaosBladeService) ListExperiments() ([]models.Experiment, error) {
    cmd := exec.Command("blade", "status")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return nil, err
    }

    var experiments []models.Experiment
    // 解析输出到实验列表
    // TODO: 实现解析逻辑
    
    return experiments, nil
}

func (s *ChaosBladeService) GetExperimentStatus(expID string) (*models.ExperimentStatus, error) {
    cmd := exec.Command("blade", "status", expID)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return nil, err
    }

    status := &models.ExperimentStatus{}
    // TODO: 解析状态
    return status, nil
} 