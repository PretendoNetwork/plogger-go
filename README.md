# Plogger
***P***retendo Network ***logger*** library

## Logs
Plogger supports 5 log levels:

- Critical - Used to indicate something has gone horribly wrong
- Error - A general error has occured
- Warning - Something is likely not right, but not overly important
- Success - Success
- Info - A general log message

All log messages are formatted using the following format: `[time] [level] spacing [function_name package_name/file_name:file_line] : message`. The `spacing` is used to align sections

For example:

```
[2025-02-16T17:22:50] [WARNING]  [func 0() main.init/init.go:34] : Error loading .env file: open .env: no such file or directory
[2025-02-16T17:22:50] [ERROR]    [func 0() main.init/init.go:52] : PN_FRIENDS_CONFIG_DATABASE_URI environment variable not set
```

## API

### Creating a logger
To create a logger, call `plogger.NewLogger()` with an optional base path. If no path is given, `.` is used. When the logger writes to a file, the file path `BASE/log/TYPE.log` is written to

```go
var logger = plogger.NewLogger() // Writes log files to ./log/TYPE.log
```

```go
var logger = plogger.NewLogger("./some/other/base") // Writes log files to ./some/other/base/log/TYPE.log
```

### Logging a message
```go
logger.Critical("Log message") // Creates a "Critical" level log message
logger.Error("Log message") // Creates a "Error" level log message
logger.Warning("Log message") // Creates a "Warning" level log message
logger.Success("Log message") // Creates a "Success" level log message
logger.Info("Log message") // Creates a "Info" level log message

logger.Criticalf("Log %s", "message") // Creates a "Critical" level log message, with additional formatting
logger.Errorf("Log %s", "message") // Creates a "Error" level log message, with additional formatting
logger.Warningf("Log %s", "message") // Creates a "Warning" level log message, with additional formatting
logger.Successf("Log %s", "message") // Creates a "Success" level log message, with additional formatting
logger.Infof("Log %s", "message") // Creates a "Info" level log message, with additional formatting
```

### Opt-out of logging
By default all loggers will write logs to both the console, to the log levels log file, and the `all.log` file. To opt-out of a type of logging, configure the logger after creation:

```go
logger.SetLogToStdOut(false) // Disables console logging for this specific logger
logger.SetLogToFile(false) // Disables file logging for this specific logger

plogger.SetGlobalLogToStdOut(false) // Disables console logging for all loggers
plogger.SetGlobalLogToFile(false) // Disables file logging for all loggers
```

The global settings may also be set using environment variables:

```env
PLOGGER_DISABLE_CONSOLE_LOGGING_GLOBAL=true # Disables console logging for all loggers
PLOGGER_DISABLE_FILE_LOGGING_GLOBAL=true # Disables file logging for all loggers
```