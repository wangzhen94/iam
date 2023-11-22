package apiserver

import pb "github.com/marmotedu/api/proto/apiserver/v1"

type secrets struct {
	cli pb.CacheClient
}

func (s *secrets) List() (map[string]*pb.SecretInfo, error) {
	//TODO implement me
	panic("implement me")
}

func newSecrets(ds *datastore) *secrets {
	return &secrets{ds.cli}
}
