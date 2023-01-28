package logger

import (
	"atom/container"
	"atom/providers/config"
	"log"
	"strings"

	"github.com/go-micro/plugins/v4/logger/zap"
	"go-micro.dev/v4/logger"
)

type Logger struct {
}

type LevelWriter struct {
	Logger logger.Logger
	Level  string
}

func (w LevelWriter) Write(p []byte) (n int, err error) {
	level, _ := logger.GetLevel(w.Level)
	w.Logger.Logf(level, strings.TrimSpace(string(p)))
	return len(p), nil
}

func init() {
	if err := container.Container.Provide(NewLogger); err != nil {
		log.Fatal(err)
	}
}

func NewLogger(conf *config.Config) (*Logger, error) {
	zapLogger, err := zap.NewLogger()
	if err != nil {
		return nil, err
	}
	logger.DefaultLogger = zapLogger

	return &Logger{}, nil
}

func (l *Logger) Init(options ...logger.Option) error                { panic("do not use") }
func (l *Logger) Options() logger.Options                            { panic("do not use") }
func (l *Logger) Fields(fields map[string]interface{}) logger.Logger { panic("do not use") }
func (l *Logger) String() string                                     { panic("do not use") }

func (l *Logger) Logf(level logger.Level, format string, v ...interface{}) {
	logger.Logf(level, format, v...)
}
func (l *Logger) Log(level logger.Level, v ...interface{}) {
	logger.Log(level, v...)
}

func (l *Logger) Info(args ...interface{}) {
	logger.Log(logger.InfoLevel, args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	logger.Logf(logger.InfoLevel, template, args...)
}

func (l *Logger) Trace(args ...interface{}) {
	logger.Log(logger.TraceLevel, args...)
}

func (l *Logger) Tracef(template string, args ...interface{}) {
	logger.Logf(logger.TraceLevel, template, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	logger.Log(logger.DebugLevel, args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	logger.Logf(logger.DebugLevel, template, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	logger.Log(logger.WarnLevel, args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	logger.Logf(logger.WarnLevel, template, args...)
}

func (l *Logger) Error(args ...interface{}) {
	logger.Log(logger.ErrorLevel, args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	logger.Logf(logger.ErrorLevel, template, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	logger.Log(logger.FatalLevel, args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	logger.Logf(logger.FatalLevel, template, args...)
}
