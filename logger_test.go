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

func TestCritialfLog(t *testing.T) {
	logger.Criticalf("Critical errorf: %s, %t, %d", "I'm a string", false, 42)
}

func TestErrorfLog(t *testing.T) {
	logger.Errorf("Errorf: %s, %t, %d", "I'm a string", false, 42)
}

func TestWarningfLog(t *testing.T) {
	logger.Warningf("Warningf: %s, %t, %d", "I'm a string", false, 42)
}

func TestSuccessfLog(t *testing.T) {
	logger.Successf("Successf: %s, %t, %d", "I'm a string", false, 42)
}

func TestInfofLog(t *testing.T) {
	logger.Infof("Infof: %s, %t, %d", "I'm a string", false, 42)
}
