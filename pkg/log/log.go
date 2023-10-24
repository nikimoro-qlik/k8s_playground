package log

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger appliaction logger in json format
type Logger struct {
	sugar       *zap.SugaredLogger
	zapLevel    zap.AtomicLevel
	writeSyncer zapcore.WriteSyncer
	sync.Once
}

// singleton logger
var logger = &Logger{writeSyncer: os.Stdout}

// GetLogger returns a global singleton logger for generic use
func GetLogger() *Logger {
	logger.Once.Do(func() {
		logger.zapLevel = zap.NewAtomicLevel()

		encoderCfg := zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			NameKey:        "logger",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		}
		core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), logger.writeSyncer, logger.zapLevel)
		customLogger := zap.New(core)
		defer customLogger.Sync()
		logger.sugar = customLogger.Sugar()

	})
	return logger
}

// GetSugar returns the underlying sugar logger
func (s *Logger) GetSugar() *zap.SugaredLogger {
	return s.sugar
}

// Info info level log using fmt.Sprint to construct and log a message.
func (s *Logger) Info(args ...interface{}) {
	s.sugar.Info(args...)
}

// Infow info level logs of msg  with some additional context. The variadic key-value
// pairs where the first element of the pair is used as the field key
// and the second as the field value.
func (s *Logger) Infow(msg string, keysAndValues ...interface{}) {
	s.sugar.Infow(msg, keysAndValues...)
}

// Infof info level logs uses fmt.Sprintf to log a templated message.
func (s *Logger) Infof(template string, args ...interface{}) {
	s.sugar.Infof(template, args...)
}

// Warn warning level log using fmt.Sprint to construct and log a message.
func (s *Logger) Warn(args ...interface{}) {
	s.sugar.Warn(args...)
}

// Warnw warning level logs of msg  with some additional context. The variadic key-value
// pairs where the first element of the pair is used as the field key
// and the second as the field value.
func (s *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	s.sugar.Warnw(msg, keysAndValues...)
}

// Warnf warning level logs uses fmt.Sprintf to log a templated message.
func (s *Logger) Warnf(template string, args ...interface{}) {
	s.sugar.Warnf(template, args...)
}

// Error error level log using fmt.Sprint to construct and log a message.
func (s *Logger) Error(args ...interface{}) {
	s.sugar.Error(args...)
}

// Errorw error level logs of msg  with some additional context. The variadic key-value
// pairs where the first element of the pair is used as the field key
// and the second as the field value.
func (s *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	s.sugar.Errorw(msg, keysAndValues...)
}

// Errorf error level logs uses fmt.Sprintf to log a templated message.
func (s *Logger) Errorf(template string, args ...interface{}) {
	s.sugar.Errorf(template, args...)
}

// Debug debug level log using fmt.Sprint to construct and log a message.
func (s *Logger) Debug(args ...interface{}) {
	s.sugar.Debug(args...)
}

// Debugw debug level logs of msg  with some additional context. The variadic key-value
// pairs where the first element of the pair is used as the field key
// and the second as the field value.
func (s *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	s.sugar.Debugw(msg, keysAndValues...)
}

// Debugf debug level logs uses fmt.Sprintf to log a templated message.
func (s *Logger) Debugf(template string, args ...interface{}) {
	s.sugar.Debugf(template, args...)
}

// Log log interface with error
func (s *Logger) Log(args ...interface{}) error {
	s.Error(args...)
	return nil
}

// SetLevel set level use debug, info, error, warning
// unexpected string set level to info
func (s *Logger) SetLevel(level string) {
	s.zapLevel.SetLevel(getLevelFromString(level))
}

// GetLevelFromString takes the log level as string and converts it to zap's type
func getLevelFromString(level string) zapcore.Level {
	switch level {
	case "info", "INFO", "":
		return zapcore.InfoLevel
	case "warn", "WARN", "warning", "WARNING":
		return zapcore.WarnLevel
	case "error", "ERROR":
		return zapcore.ErrorLevel
	case "debug", "DEBUG":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
