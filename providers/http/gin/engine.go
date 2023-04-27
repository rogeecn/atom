package gin

import (
	"fmt"
	"time"

	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers"
	"github.com/rogeecn/atom/providers/http"
	"github.com/rogeecn/atom/providers/log"

	"github.com/gin-gonic/gin"
)

type Service struct {
	conf   *http.Config
	Engine *gin.Engine
}

func (e *Service) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return e.Engine.Use(middleware...)
}

func (e *Service) GetEngine() interface{} {
	return e.Engine
}

func (e *Service) Serve() error {
	if e.conf.Tls != nil {
		return e.Engine.RunTLS(e.conf.PortString(), e.conf.Tls.Cert, e.conf.Tls.Key)
	}
	return e.Engine.Run(e.conf.PortString())
}

func Provide(o *providers.Options) error {
	var config http.Config
	if err := o.UnmarshalConfig(&config); err != nil {
		log.Fatal(err)
	}

	return container.Container.Provide(func() (http.Service, error) {
		gin.DefaultWriter = log.LevelWriter{Level: log.InfoLevel}
		gin.DefaultErrorWriter = log.LevelWriter{Level: log.ErrorLevel}

		engine := gin.New()
		engine.Use(gin.Recovery())
		engine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			return fmt.Sprintf(`%s - [%s] "%s %s %s %d %s '%q' %s"\n`,
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

		return &Service{Engine: engine, conf: &config}, nil
	}, o.DiOptions()...)
}
