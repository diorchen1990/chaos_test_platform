package executor

import (
	"context"
	"github.com/your-project/models"
	"github.com/your-project/probe"
)

type Executor interface {
	Execute(ctx context.Context, exp *models.Experiment) error
	Stop(expID string) error
}

type ChaosExecutor struct {
	probes map[string]probe.Probe
}

func (e *ChaosExecutor) Execute(ctx context.Context, exp *models.Experiment) error {
	probe, ok := e.probes[exp.Scope]
	if !ok {
		return ErrProbeNotFound
	}

	return probe.InjectFault(exp)
} 