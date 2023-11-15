package apiserver

import (
	"context"
	"github.com/wangzhen94/iam/internal/apiserver/config"
	genericoptions "github.com/wangzhen94/iam/internal/pkg/options"
	"github.com/wangzhen94/iam/pkg/storage"
)

import (
	genericapiserver "github.com/wangzhen94/iam/internal/pkg/server"
)

type apiServer struct {
	//gs           *shutdown.GracefulShutdown
	redisOptions *genericoptions.RedisOptions
	//gRPCAPIServer    *grpcAPIServer
	genericAPIServer *genericapiserver.GenericAPIServer
}

func createAPIServer(cfg *config.Config) (*apiServer, error) {
	//gs := shutdown.New()
	//gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}

	//extraConfig, err := buildExtraConfig(cfg)
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}
	//extraServer, err := extraConfig.complete().New()
	if err != nil {
		return nil, err
	}

	server := &apiServer{
		//gs:               gs,
		redisOptions:     cfg.RedisOptions,
		genericAPIServer: genericServer,
		//gRPCAPIServer:    extraServer,
	}

	return server, nil
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()
	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}

type preparedAPIServer struct {
	*apiServer
}

func (s *apiServer) PrepareRun() preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)

	s.initRedisStore()

	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run() error {
	//go s.gRPCAPIServer.Run()

	// start shutdown managers
	//if err := s.gs.Start(); err != nil {
	//	log.Fatalf("start shutdown manager failed: %s", err.Error())
	//}

	return s.genericAPIServer.Run()
}

func (s *apiServer) initRedisStore() {
	ctx, _ := context.WithCancel(context.Background())
	//s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
	//	cancel()
	//
	//	return nil
	//}))

	config := &storage.Config{
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

	// try to connect to redis
	go storage.ConnectToRedis(ctx, config)
}
