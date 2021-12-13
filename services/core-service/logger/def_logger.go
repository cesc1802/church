package logger

func Fatal(format string, args ...interface{}) {
	DefaultLogger.Fatal(format, args...)
}

// Error prints a error level log message to the stderr. Arguments are handled in the manner of fmt.Printf.
func Error(format string, args ...interface{}) {
	DefaultLogger.Error(format, args...)
}

// Warning prints a warning level log message to the stderr. Arguments are handled in the manner of fmt.Printf.
func Warning(format string, args ...interface{}) {
	DefaultLogger.Warn(format, args...)
}

// Info prints a info level log message to the stderr. Arguments are handled in the manner of fmt.Printf.
func Info(format string, args ...interface{}) {
	DefaultLogger.Info(format, args...)
}

// Debug prints a debug level log message to the stderr. Arguments are handled in the manner of fmt.Printf.
func Debug(format string, args ...interface{}) {
	DefaultLogger.Debug(format, args...)
}
