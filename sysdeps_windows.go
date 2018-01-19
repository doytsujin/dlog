package dlog

import "golang.org/x/sys/windows/svc/eventlog"

type systemLogger struct {
	inner *eventlog.Log
}

func NewSystemLogger(appName string, facility string) (*systemLogger, error) {
	err := eventlog.InstallAsEventCreate(appName, eventlog.Error|eventlog.Warning|eventlog.Info)
	if err != nil {
		return nil, err
	}
	var eventLogger *eventlog.Log
	if eventLogger, err = eventlog.Open(appName); err != nil {
		return nil, err
	}
	return &systemLogger{inner: eventLogger}, nil
}

func (systemLogger *systemLogger) WriteString(severity Severity, message string) {
	switch severity {
	case SeverityError:
	case SeverityCritical:
	case SeverityFatal:
		systemLogger.inner.Error(uint32(severity), message)
	case SeverityWarning:
		systemLogger.inner.Warning(uint32(severity), message)
	default:
		systemLogger.inner.Info(uint32(severity), message)
	}
}
