package logger

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

// Logger -.
type Logger interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

// Logger -.
type logger struct {
	logger *zerolog.Logger
}

var _ Logger = (*logger)(nil)

type (
	// Level represent severity of logs
	Level int
)

var (
	// DefaultLogger holds default logger
	DefaultLogger Logger = NewDefaultLogger()

	// DefaultLevel holds default value for loggers
	DefaultLevel Level = INFO

	// StdoutHandler holds a handler with outputting to stdout
	StdoutHandler = os.Stdout

	// StderrHandler holds a handler with outputting to stderr
	StderrHandler = os.Stderr
)

// Logging levels.
const (
	CRITICAL Level = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

// LevelNames provides mapping for log levels
var LevelNames = map[Level]string{
	ERROR:   "ERROR",
	WARNING: "WARNING",
	NOTICE:  "NOTICE",
	INFO:    "INFO",
	DEBUG:   "DEBUG",
}

var ZeroLogLevelToAppLogLevel = map[Level]zerolog.Level{
	ERROR:   zerolog.ErrorLevel,
	WARNING: zerolog.WarnLevel,
	INFO:    zerolog.InfoLevel,
	DEBUG:   zerolog.DebugLevel,
}

func NewLogger(level Level) *logger {
	l := ZeroLogLevelToAppLogLevel[level]
	zerolog.SetGlobalLevel(l)
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &logger{
		logger: &log,
	}
}

func (l *logger) WithHook(hookHdl zerolog.Hook) {
	if hookHdl == nil {
		return
	}
	log := l.logger.Hook(hookHdl)
	l.logger = &log
}

func NewDefaultLogger() *logger {
	l := ZeroLogLevelToAppLogLevel[DefaultLevel]
	zerolog.SetGlobalLevel(l)
	log := zerolog.New(StdoutHandler).With().Timestamp().Logger()

	return &logger{
		logger: &log,
	}
}

// Debug -.
func (l *logger) Debug(message interface{}, args ...interface{}) {
	l.msg("DEBUG", message, args...)
}

// Info -.
func (l *logger) Info(message string, args ...interface{}) {
	l.log(message, args...)
}

// Warn -.
func (l *logger) Warn(message string, args ...interface{}) {
	l.log(message, args...)
}

// Error -.
func (l *logger) Error(message interface{}, args ...interface{}) {
	if l.logger.GetLevel() == zerolog.DebugLevel {
		l.Debug(message, args...)
	}

	l.msg("ERROR", message, args...)
}

// Fatal -.
func (l *logger) Fatal(message interface{}, args ...interface{}) {
	l.msg("FATAL", message, args...)

	os.Exit(1)
}

func (l *logger) log(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info().Msg(message)
	} else {
		l.logger.Info().Msgf(message, args...)
	}
}

func (l *logger) msg(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(msg.Error(), args...)
	case string:
		l.log(msg, args...)
	default:
		l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}

type loggerKey struct{}

func WithLogger(ctx context.Context, log *logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, log)
}

func FromContext(ctx context.Context) *logger {
	if ctx == nil {
		return NewDefaultLogger()
	}
	if l, ok := ctx.Value(loggerKey{}).(*logger); ok {
		return l
	}

	return NewDefaultLogger()
}
