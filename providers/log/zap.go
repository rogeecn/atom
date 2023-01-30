package log

import (
	"atom/providers/config"

	"go.uber.org/zap"
)

func NewZapLogger(conf *config.Config) (*Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	DefaultLogger = &Logger{logger: logger.Sugar()}
	return DefaultLogger, nil
}
