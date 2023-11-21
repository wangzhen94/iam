package cache

import (
	"context"
	"fmt"
	pb "github.com/marmotedu/api/proto/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	"github.com/wangzhen94/iam/internal/apiserver/store"
	"github.com/wangzhen94/iam/internal/pkg/code"
	"github.com/wangzhen94/iam/pkg/log"
	"sync"
)

type Cache struct {
	store store.Factory
}

var (
	cacheServer *Cache
	once        sync.Once
)

func GetCacheInsOr(store store.Factory) (*Cache, error) {
	if store != nil {
		once.Do(func() {
			cacheServer = &Cache{store}
		})
	}

	if cacheServer == nil {
		return nil, fmt.Errorf("got nil cache server")
	}

	return cacheServer, nil
}

func (c *Cache) ListSecrets(ctx context.Context, r *pb.ListSecretsRequest) (*pb.ListSecretsResponse, error) {
	log.L(ctx).Info("cache list secrets function called.")
	opts := metav1.ListOptions{
		Offset: r.Offset,
		Limit:  r.Limit,
	}

	secrets, err := c.store.Secrets().List(ctx, "", opts)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	items := make([]*pb.SecretInfo, 0)
	for _, secret := range secrets.Items {
		items = append(items, &pb.SecretInfo{
			SecretId:    secret.SecretID,
			Username:    secret.Username,
			SecretKey:   secret.SecretKey,
			Expires:     secret.Expires,
			Description: secret.Description,
			CreatedAt:   secret.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   secret.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.ListSecretsResponse{
		Items:      items,
		TotalCount: secrets.TotalCount,
	}, nil
}

func (c *Cache) ListPolicies(ctx context.Context, r *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	log.L(ctx).Info("cache list policies function called.")
	opts := metav1.ListOptions{
		Offset: r.Offset,
		Limit:  r.Limit,
	}

	policies, err := c.store.Policies().List(ctx, "", opts)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	items := make([]*pb.PolicyInfo, 0)
	for _, pol := range policies.Items {
		items = append(items, &pb.PolicyInfo{
			Name:         pol.Name,
			Username:     pol.Username,
			PolicyShadow: pol.PolicyShadow,
			CreatedAt:    pol.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.ListPoliciesResponse{
		TotalCount: policies.TotalCount,
		Items:      items,
	}, nil
}
