package logger

import (
	"atom/container"
	"atom/providers/config"
	"log"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	if err := container.Container.Provide(NewLogger); err != nil {
		log.Fatal(err)
	}
}

type Logger struct {
	logger *zap.SugaredLogger
}

type LevelWriter struct {
	Logger *Logger
	Level  zapcore.Level
}

func (w LevelWriter) Write(p []byte) (n int, err error) {
	str := strings.TrimSpace(string(p))
	switch w.Level {
	case zapcore.InfoLevel:
		w.Logger.Info(str)
	case zapcore.ErrorLevel:
		w.Logger.Info(str)
	}
	return len(p), nil
}

func NewLogger(conf *config.Config) (*Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return &Logger{logger: logger.Sugar()}, nil
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (l *Logger) DPanic(args ...interface{}) {
	l.logger.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *Logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *Logger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (l *Logger) DPanicf(template string, args ...interface{}) {
	l.logger.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) DPanicw(msg string, keysAndValues ...interface{}) {
	l.logger.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.logger.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

// Sync flushes any buffered log entries.
func (l *Logger) Sync() error {
	return l.logger.Sync()
}
