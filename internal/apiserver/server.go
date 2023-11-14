package apiserver

import (
	"github.com/wangzhen94/iam/internal/apiserver/config"
)

import (
	genericapiserver "github.com/wangzhen94/iam/internal/pkg/server"
)

type apiServer struct {
	//gs               *shutdown.GracefulShutdown
	//redisOptions     *genericoptions.RedisOptions
	gRPCAPIServer    *grpcAPIServer
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
		//redisOptions:     cfg.RedisOptions,
		genericAPIServer: genericServer,
		//gRPCAPIServer:    extraServer,
	}

	return server, nil
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	//genericConfig = genericapiserver.NewConfig()
	//if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
	//	return
	//}
	//
	//if lastErr = cfg.FeatureOptions.ApplyTo(genericConfig); lastErr != nil {
	//	return
	//}
	//
	//if lastErr = cfg.SecureServing.ApplyTo(genericConfig); lastErr != nil {
	//	return
	//}
	//
	//if lastErr = cfg.InsecureServing.ApplyTo(genericConfig); lastErr != nil {
	//	return
	//}

	return
}

type preparedAPIServer struct {
	*apiServer
}

func (s *apiServer) PrepareRun() preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)

	//s.initRedisStore()

	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run() error {
	go s.gRPCAPIServer.Run()

	// start shutdown managers

	return s.genericAPIServer.Run()
}
