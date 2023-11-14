package server

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/component-base/pkg/version"
	"github.com/marmotedu/log"
	"github.com/wangzhen94/iam/internal/pkg/middleware"
	"net/http"
	"time"
)

type GenericAPIServer struct {
	middlewares []string
	// SecureServingInfo holds configuration of the TLS server.
	//SecureServingInfo *SecureServingInfo

	// InsecureServingInfo holds configuration of the insecure HTTP server.
	//InsecureServingInfo *InsecureServingInfo

	ShutdownTimeout time.Duration

	*gin.Engine
	healthz         bool
	enableMetrics   bool
	enableProfiling bool

	insecureServer, secureServer *http.Server
}

func initGenericAPIServer(s *GenericAPIServer) {
	// do some setup
	// s.GET(path, ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.Setup()
	s.InstallMiddlewares()
	s.InstallAPIs()
}

// InstallAPIs install generic apis.
func (s *GenericAPIServer) InstallAPIs() {
	// install healthz handler
	if s.healthz {
		s.GET("/healthz", func(c *gin.Context) {
			core.WriteResponse(c, nil, map[string]string{"status": "ok"})
		})
	}

	//// install metric handler
	//if s.enableMetrics {
	//	prometheus := ginprometheus.NewPrometheus("gin")
	//	prometheus.Use(s.Engine)
	//}

	// install pprof handler
	if s.enableProfiling {
		pprof.Register(s.Engine)
	}

	s.GET("/version", func(c *gin.Context) {
		core.WriteResponse(c, nil, version.Get())
	})
}

// Setup do some setup work for gin engine.
func (s *GenericAPIServer) Setup() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

// InstallMiddlewares install generic middlewares.
func (s *GenericAPIServer) InstallMiddlewares() {
	// necessary middlewares
	s.Use(middleware.RequestID())
	s.Use(middleware.Context())

	// install custom middlewares
	for _, m := range s.middlewares {
		mw, ok := middleware.Middlewares[m]
		if !ok {
			log.Warnf("can not find middleware: %s", m)

			continue
		}

		log.Infof("install middleware: %s", m)
		s.Use(mw)
	}
}
