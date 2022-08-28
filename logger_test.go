package plogger

import "testing"

var logger = NewLogger()

func TestCritialLog(t *testing.T) {
	logger.Critical("Critical error")
}

func TestErrorLog(t *testing.T) {
	logger.Error("Error")
}

func TestWarningLog(t *testing.T) {
	logger.Warning("Warning")
}

func TestSuccessLog(t *testing.T) {
	logger.Success("Success")
}

func TestInfoLog(t *testing.T) {
	logger.Info("Info")
}
