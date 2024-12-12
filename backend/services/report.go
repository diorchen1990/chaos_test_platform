package services

import (
	"context"
	"time"
	"your-project/models"
	"your-project/pkg/utils"
)

type ReportService struct {
	metricsCollector *monitor.MetricsCollector
	experimentSvc    *ExperimentService
}

func (s *ReportService) GenerateReport(ctx context.Context, experimentID string) (*models.Report, error) {
	// 获取实验信息
	exp, err := s.experimentSvc.GetExperiment(experimentID)
	if err != nil {
		return nil, err
	}

	// 收集指标数据
	metrics, err := s.metricsCollector.GetExperimentMetrics(experimentID)
	if err != nil {
		return nil, err
	}

	// 生成报告
	report := &models.Report{
		ID:           utils.GenerateID(),
		ExperimentID: experimentID,
		Name:         exp.Name + "_Report",
		StartTime:    exp.CreateTime,
		EndTime:      time.Now(),
		Status:       exp.Status,
		Summary: models.ReportSummary{
			TotalSteps:    len(exp.Steps),
			Duration:      time.Since(exp.CreateTime).Seconds(),
			AffectedNodes: exp.AgentIDs,
		},
		Metrics: metrics,
	}

	// 生成步骤信息
	for _, step := range exp.Steps {
		reportStep := models.ReportStep{
			Name:      step.Name,
			Status:    step.Status,
			StartTime: step.StartTime,
			EndTime:   step.EndTime,
			Action:    step.Action,
			Target:    step.Target,
			Error:     step.Error,
		}
		report.Steps = append(report.Steps, reportStep)
	}

	return report, nil
}

func (s *ReportService) ExportReportPDF(report *models.Report) ([]byte, error) {
	// 使用模板生成PDF
	template := template.New("report")
	template.Parse(reportTemplate)
	
	var buf bytes.Buffer
	if err := template.Execute(&buf, report); err != nil {
		return nil, err
	}

	return utils.GeneratePDF(buf.Bytes())
} 