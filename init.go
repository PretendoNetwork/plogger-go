package plogger

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func init() {
	disableGlobalConsoleLogging := os.Getenv("PLOGGER_DISABLE_CONSOLE_LOGGING_GLOBAL")
	disableGlobalFileLogging := os.Getenv("PLOGGER_DISABLE_FILE_LOGGING_GLOBAL")

	if disableGlobalConsoleLogging != "" {
		globalLogToStdOut = false
	}

	if disableGlobalFileLogging != "" {
		globalLogToFile = false
	}
}
