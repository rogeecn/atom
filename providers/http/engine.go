package http

import (
	"atom/container"
	"atom/providers/config"
	"atom/providers/logger"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := container.Container.Provide(NewService); err != nil {
		log.Fatal(err)
	}
}

type Service struct {
	Engine *gin.Engine
	conf   *config.Config
}

func (e *Service) Serve() error {
	if e.conf.Http.Https {
		return e.Engine.RunTLS(e.conf.Http.PortString(), e.conf.Http.HttpsCert, e.conf.Http.HttpKey)
	}
	return e.Engine.Run(e.conf.Http.PortString())
}

func NewService(log *logger.Logger, cfg *config.Config) *Service {
	log.Info("init http service with gin...")

	gin.DefaultWriter = &logger.LevelWriter{Logger: log, Level: "info"}
	gin.DefaultErrorWriter = &logger.LevelWriter{Logger: log, Level: "error"}

	if cfg.App.IsDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	engine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	engine.Use(gin.Recovery())

	return &Service{Engine: engine, conf: cfg}
}
