package log

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

type LevelWriter struct {
	Logger *Logger
	Level  zapcore.Level
}

func (w LevelWriter) Write(p []byte) (n int, err error) {
	str := strings.TrimSpace(string(p))
	switch w.Level {
	case zapcore.InfoLevel:
		Info(str)
	case zapcore.ErrorLevel:
		Error(str)
	}
	return len(p), nil
}
