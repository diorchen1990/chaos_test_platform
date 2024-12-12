package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type ReportHandler struct {
    reportService *services.ReportService
}

func (h *ReportHandler) GenerateReport(c *gin.Context) {
    experimentID := c.Param("id")
    
    report, err := h.reportService.GenerateReport(c, experimentID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) DownloadReport(c *gin.Context) {
    reportID := c.Param("id")
    format := c.Query("format") // pdf or html
    
    var data []byte
    var err error
    
    switch format {
    case "pdf":
        data, err = h.reportService.ExportReportPDF(reportID)
    case "html":
        data, err = h.reportService.ExportReportHTML(reportID)
    default:
        c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported format"})
        return
    }

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    filename := fmt.Sprintf("chaos_report_%s.%s", reportID, format)
    c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
    c.Data(http.StatusOK, "application/octet-stream", data)
} 