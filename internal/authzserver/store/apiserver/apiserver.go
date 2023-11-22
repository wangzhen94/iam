package apiserver

import (
	pb "github.com/marmotedu/api/proto/apiserver/v1"
	"github.com/wangzhen94/iam/internal/authzserver/store"
	"github.com/wangzhen94/iam/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"sync"
)

type datastore struct {
	cli pb.CacheClient
}

func (d *datastore) Secrets() store.SecretStore {
	return newSecrets(d)
}

func (d *datastore) Policies() store.PolicyStore {
	return nil
}

var (
	apiServerFactory store.Factory
	once             sync.Once
)

func GetAPIServerFactoryOrDie(address string, clientCA string) store.Factory {
	once.Do(func() {
		var (
			err   error
			conn  *grpc.ClientConn
			creds credentials.TransportCredentials
		)

		creds, err = credentials.NewClientTLSFromFile(clientCA, "")
		if err != nil {
			log.Panicf("credentials.NewClientTLSFromFile err: %v", err)
		}

		conn, err = grpc.Dial(address, grpc.WithBlock(), grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Panicf("Connect to grpc server failed, error: %s", err.Error())
		}

		apiServerFactory = &datastore{pb.NewCacheClient(conn)}
		log.Infof("Connected to grpc server, address: %s", address)
	})

	if apiServerFactory == nil {
		log.Panicf("failed to get apiserver store fatory")
	}
	return apiServerFactory
}
