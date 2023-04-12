package consul

import (
	"github.com/hashicorp/go-hclog"
)

type ReportingManager struct {
	logger hclog.Logger
	EntDeps
}

func NewReportingManager(logger hclog.Logger) *ReportingManager {
	return &ReportingManager{
		logger: logger.Named("reporting"),
	}
}
