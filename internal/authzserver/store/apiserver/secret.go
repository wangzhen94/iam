package apiserver

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/avast/retry-go"
	pb "github.com/marmotedu/api/proto/apiserver/v1"
	"github.com/marmotedu/errors"
	"github.com/wangzhen94/iam/pkg/log"
)

type secrets struct {
	cli pb.CacheClient
}

func (s *secrets) List() (map[string]*pb.SecretInfo, error) {
	res := make(map[string]*pb.SecretInfo)

	log.Info("loading secrets")

	req := &pb.ListSecretsRequest{
		Offset: pointer.ToInt64(0),
		Limit:  pointer.ToInt64(-1),
	}

	var resp *pb.ListSecretsResponse

	err := retry.Do(
		func() error {
			var listErr error
			resp, listErr = s.cli.ListSecrets(context.Background(), req)
			if listErr != nil {
				return listErr
			}

			return nil
		}, retry.Attempts(3),
	)
	if err != nil {
		return nil, errors.Wrap(err, "list res failed.")
	}

	log.Infof("Secrets found (%d total):", len(resp.Items))

	for _, v := range resp.Items {
		log.Infof(" - %s:%s", v.Username, v.SecretId)
		res[v.SecretId] = v
	}
	return res, nil
}

func newSecrets(ds *datastore) *secrets {
	return &secrets{ds.cli}
}
