//go:build !consulent
// +build !consulent

package consul

type EntDeps struct{}

func (rm *ReportingManager) initReporting(deps Deps) error {
	return nil
}

func (rm *ReportingManager) startReportingAgent() error {
	// no op
	return nil
}

func (rm *ReportingManager) stopReportingAgent() error {
	// no op
	return nil
}
