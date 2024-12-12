package tests

import (
	"context"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	
	"your-project/models"
	"your-project/services"
	"your-project/pkg/errors"
)

type MockAgent struct {
	mock.Mock
}

func (m *MockAgent) ExecuteExperiment(ctx context.Context, exp *models.Experiment) error {
	args := m.Called(ctx, exp)
	return args.Error(0)
}

func TestExperimentExecution(t *testing.T) {
	// 设置测试用例
	tests := []struct {
		name     string
		exp      *models.Experiment
		mockResp error
		wantErr  bool
	}{
		{
			name: "成功执行Java异常注入",
			exp: &models.Experiment{
				Scope: models.ScopeJavaApp,
				Action: "java-exception",
				Target: models.Target{
					JavaApp: &models.JavaAppTarget{
						ProcessName: "test-service",
					},
				},
			},
			mockResp: nil,
			wantErr:  false,
		},
		{
			name: "Agent执行失败",
			exp: &models.Experiment{
				Scope: models.ScopeJavaApp,
				Action: "java-exception",
			},
			mockResp: errors.New("agent error"),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建mock agent
			mockAgent := new(MockAgent)
			mockAgent.On("ExecuteExperiment", mock.Anything, tt.exp).Return(tt.mockResp)

			// 创建service
			svc := services.NewExperimentService(mockAgent)

			// 执行测试
			err := svc.ExecuteExperiment(context.Background(), tt.exp)

			// 验证结果
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// 验证mock调用
			mockAgent.AssertExpectations(t)
		})
	}
} 