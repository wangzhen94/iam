package authzserver

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/wangzhen94/iam/internal/authzserver/analytics"
	"github.com/wangzhen94/iam/internal/authzserver/config"
	"github.com/wangzhen94/iam/internal/authzserver/controller/v1/authorize"
	"github.com/wangzhen94/iam/internal/authzserver/load"
	"github.com/wangzhen94/iam/internal/authzserver/load/cache"
	"github.com/wangzhen94/iam/internal/authzserver/store/apiserver"
	"github.com/wangzhen94/iam/internal/pkg/code"
	genericoptions "github.com/wangzhen94/iam/internal/pkg/options"
	genericapiserver "github.com/wangzhen94/iam/internal/pkg/server"
	"github.com/wangzhen94/iam/pkg/log"
	"github.com/wangzhen94/iam/pkg/storage"
	"golang.org/x/net/context"
)

const RedisKeyPrefix = "analytics-"

type authzServer struct {
	rpcServer        string
	clientCA         string
	redisOptions     *genericoptions.RedisOptions
	genericAPIServer *genericapiserver.GenericAPIServer
	analyticsOptions *analytics.AnalyticsOptions
	redisCancelFunc  context.CancelFunc
}

func (s *authzServer) PrepareRun() prepareAuthzServer {
	_ = s.initialize()

	initRouter(s.genericAPIServer.Engine)

	return prepareAuthzServer{s}
}

func initRouter(g *gin.Engine) {
	installController(g)
}

func installController(g *gin.Engine) {
	auth := newCacheAuth()
	g.NoRoute(auth.AuthFunc(), func(c *gin.Context) {
		core.WriteResponse(c, errors.WithCode(code.ErrPageNotFound, "page not found."), nil)
	})

	cacheIns, _ := cache.GetCacheInsOr(nil)
	if cacheIns == nil {
		log.Panic("get nil cache instance")
	}
	v1 := g.Group("/v1", auth.AuthFunc())
	{
		authzController := authorize.NewAuthzController(cacheIns)

		v1.POST("/authz", authzController.Authorize)
	}
}

func (s *authzServer) initialize() error {
	ctx, cancel := context.WithCancel(context.Background())
	s.redisCancelFunc = cancel

	go storage.ConnectToRedis(ctx, s.buildStorageConfig())

	cacheIns, err := cache.GetCacheInsOr(apiserver.GetAPIServerFactoryOrDie(s.rpcServer, s.clientCA))
	if err != nil {
		return errors.Wrap(err, "get cache instance failed.")
	}

	load.NewLoader(ctx, cacheIns).Start()

	if s.analyticsOptions.Enable {
		analyticsStore := storage.RedisCluster{KeyPrefix: RedisKeyPrefix}
		analyticsIns := analytics.NewAnalytics(s.analyticsOptions, &analyticsStore)
		analyticsIns.Start()
	}

	return nil
}

type prepareAuthzServer struct {
	*authzServer
}

func (s prepareAuthzServer) Run() error {
	stopCh := make(chan struct{})

	go s.genericAPIServer.Run()

	<-stopCh
	// todo debug this
	return nil
}

func createAuthzServer(cfg *config.Config) (*authzServer, error) {

	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}

	genericAPIServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	server := &authzServer{
		rpcServer:        cfg.RPCServer,
		clientCA:         cfg.ClientCA,
		redisOptions:     cfg.RedisOptions,
		genericAPIServer: genericAPIServer,
		analyticsOptions: cfg.AnalyticsOptions,
	}

	return server, nil
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()
	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	if lastErr = cfg.SecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	if lastErr = cfg.InsecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}

func PrepareRun() {

}

func (s *authzServer) buildStorageConfig() *storage.Config {
	return &storage.Config{
		Host:                  s.redisOptions.Host,
		Port:                  s.redisOptions.Port,
		Addrs:                 s.redisOptions.Addrs,
		MasterName:            s.redisOptions.MasterName,
		Username:              s.redisOptions.Username,
		Password:              s.redisOptions.Password,
		Database:              s.redisOptions.Database,
		MaxIdle:               s.redisOptions.MaxIdle,
		MaxActive:             s.redisOptions.MaxActive,
		Timeout:               s.redisOptions.Timeout,
		EnableCluster:         s.redisOptions.EnableCluster,
		UseSSL:                s.redisOptions.UseSSL,
		SSLInsecureSkipVerify: s.redisOptions.SSLInsecureSkipVerify,
	}
}
