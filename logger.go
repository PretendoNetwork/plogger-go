package plogger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/jwalton/go-supportscolor"
)

// tried to use const here but no worky
var (
	flags int = os.O_APPEND | os.O_CREATE | os.O_WRONLY

	logTemplate string = "[%s] [%s]%s[%s %s/%s:%s] : %s"

	criticalPrefix        string = "CRITICAL"
	criticalPrefixColored string = color.New(color.FgWhite, color.BgHiRed, color.Bold).Sprint(criticalPrefix)

	errorPrefix        string = "ERROR"
	errorPrefixColored string = color.New(color.FgRed, color.Bold).Sprint(errorPrefix)

	warningPrefix        string = "WARNING"
	warningPrefixColored string = color.New(color.FgYellow, color.Bold).Sprint(warningPrefix)

	successPrefix        string = "SUCCESS"
	successPrefixColored string = color.New(color.FgGreen, color.Bold).Sprint(successPrefix)

	infoPrefix        string = "INFO"
	infoPrefixColored string = color.New(color.FgCyan, color.Bold).Sprint(infoPrefix)

	bold     func(a ...interface{}) string = color.New(color.Bold).SprintFunc()
	grey     func(a ...interface{}) string = color.New(color.FgHiBlack, color.Bold).SprintFunc()
	magenta  func(a ...interface{}) string = color.New(color.FgMagenta, color.Bold).SprintFunc()
	darkCyan func(a ...interface{}) string = color.New(color.FgHiCyan, color.Bold).SprintFunc()
	green    func(a ...interface{}) string = color.New(color.FgGreen, color.Bold).SprintFunc()
	yellow   func(a ...interface{}) string = color.New(color.FgHiYellow, color.Bold).SprintFunc() // this ends up red-ish for some reason?

	colorSupported bool = supportscolor.Stdout().SupportsColor

	maxPrefixLength = len(criticalPrefix) // size of the longest prefix, for spacing later on
)

type Logger struct {
	allLogsFile     *os.File
	criticalLogFile *os.File
	errorLogFile    *os.File
	warningLogFile  *os.File
	successLogFile  *os.File
	infoLogFile     *os.File
}

func (logger *Logger) logLine(message, prefix, prefixColored string, logFile *os.File) {
	file, line, function, packageName := getCallerInfo()
	date := time.Now().UTC().Format("2006-01-02T15:04:05")

	prefixLength := len(prefix)
	spacing := strings.Repeat(" ", (maxPrefixLength-prefixLength)+1)

	logPlain := fmt.Sprintf(logTemplate+"\n", date, prefix, " ", "func "+function, packageName, file, line, message)
	logPlainSpaced := fmt.Sprintf(logTemplate+"\n", date, prefix, spacing, "func "+function, packageName, file, line, message)

	if colorSupported {
		fmt.Printf(logTemplate+"\n", grey(date), prefixColored, spacing, magenta("func ")+darkCyan(function), green(packageName), green(file), yellow(line), bold(message))
	} else {
		fmt.Println(logPlainSpaced)
	}

	if _, err := logFile.WriteString(logPlain); err != nil {
		log.Println(err)
	}

	if _, err := logger.allLogsFile.WriteString(logPlainSpaced); err != nil {
		log.Println(err)
	}
}

func (logger *Logger) Critical(message string) {
	logger.logLine(message, criticalPrefix, criticalPrefixColored, logger.criticalLogFile)
}

func (logger *Logger) Error(message string) {
	logger.logLine(message, errorPrefix, errorPrefixColored, logger.errorLogFile)
}

func (logger *Logger) Warning(message string) {
	logger.logLine(message, warningPrefix, warningPrefixColored, logger.warningLogFile)
}

func (logger *Logger) Success(message string) {
	logger.logLine(message, successPrefix, successPrefixColored, logger.successLogFile)
}

func (logger *Logger) Info(message string) {
	logger.logLine(message, infoPrefix, infoPrefixColored, logger.infoLogFile)
}

func NewLogger(args ...string) *Logger {
	var logFolderRoot string

	if len(args) == 0 {
		logFolderRoot = "."
	} else {
		logFolderRoot = args[0]
	}

	logFolderPath := filepath.Join(logFolderRoot, "log")

	err := os.MkdirAll(logFolderPath, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	logger := &Logger{
		allLogsFile:     createFileHandle(filepath.Join(logFolderPath, "all.log")),
		criticalLogFile: createFileHandle(filepath.Join(logFolderPath, "critical.log")),
		errorLogFile:    createFileHandle(filepath.Join(logFolderPath, "error.log")),
		warningLogFile:  createFileHandle(filepath.Join(logFolderPath, "warning.log")),
		successLogFile:  createFileHandle(filepath.Join(logFolderPath, "success.log")),
		infoLogFile:     createFileHandle(filepath.Join(logFolderPath, "info.log")),
	}

	return logger
}

func createFileHandle(path string) *os.File {
	f, err := os.OpenFile(path, flags, 0644)
	if err != nil {
		log.Println(err)
	}

	return f
}

func getCallerInfo() (string, string, string, string) {
	pc, file, line, ok := runtime.Caller(3) // step up 3 calls in the goroutine stack to find the function which called the log
	if ok {
		// https://stackoverflow.com/a/56960913
		parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
		partsLength := len(parts)
		packageName := ""
		function := parts[partsLength-1]
		if parts[partsLength-2][0] == '(' {
			function = parts[partsLength-2] + "." + function
			packageName = strings.Join(parts[0:partsLength-2], ".")
		} else {
			packageName = strings.Join(parts[0:partsLength-1], ".")
		}

		fileSplit := strings.Split(file, "/")
		file = fileSplit[len(fileSplit)-1]

		return file, strconv.Itoa(line), function + "()", packageName
	}

	return "Unknown source file", "0", "Unknown function", "Unknown package"
}
