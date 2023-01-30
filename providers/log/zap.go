package log

import (
	"atom/providers/config"
	"log"

	"go.uber.org/zap"
)

func NewZapLogger(conf *config.Config) (*Logger, error) {
	log.Print("init logger")
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	DefaultLogger = &Logger{logger: logger.Sugar()}
	return DefaultLogger, nil
}
