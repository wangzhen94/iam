package cache

import (
	"github.com/dgraph-io/ristretto"
	pb "github.com/marmotedu/api/proto/apiserver/v1"
	"github.com/marmotedu/errors"
	"github.com/ory/ladon"
	"github.com/wangzhen94/iam/internal/authzserver/store"
	"sync"
)

type Cache struct {
	lock     *sync.RWMutex
	cli      store.Factory
	secrets  *ristretto.Cache
	policies *ristretto.Cache
}

var (
	// ErrSecretNotFound defines secret not found error.
	ErrSecretNotFound = errors.New("secret not found")
	// ErrPolicyNotFound defines policy not found error.
	ErrPolicyNotFound = errors.New("policy not found")
)

var (
	onceCache sync.Once
	cacheIns  *Cache
)

func (c *Cache) Reload() error {
	c.lock.Lock()
	defer c.lock.Unlock()

	secrets, err := c.cli.Secrets().List()
	if err != nil {
		return errors.Wrap(err, "grpc list secrets failed")
	}
	for key, val := range secrets {
		c.secrets.Set(key, val, 1)
	}

	policys, err := c.cli.Policies().List()
	if err != nil {
		return errors.Wrap(err, "grpc list policys failed")
	}

	for key, val := range policys {
		c.policies.Set(key, val, 1)
	}

	return nil
}

func GetCacheInsOr(cli store.Factory) (*Cache, error) {
	var err error
	if cli != nil {
		var (
			secretCache *ristretto.Cache
			policyCache *ristretto.Cache
		)
		onceCache.Do(func() {
			c := &ristretto.Config{
				NumCounters: 1e7,
				MaxCost:     1 << 30, // maximum cost of cache (1GB).
				BufferItems: 64,
				Cost:        nil,
			}

			secretCache, err = ristretto.NewCache(c)
			if err != nil {
				return
			}
			policyCache, err = ristretto.NewCache(c)
			if err != nil {
				return
			}

			cacheIns = &Cache{
				cli:      cli,
				lock:     new(sync.RWMutex),
				secrets:  secretCache,
				policies: policyCache,
			}
		})

	}
	return cacheIns, nil
}

func (c *Cache) GetSecret(key string) (*pb.SecretInfo, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	value, ok := c.secrets.Get(key)
	if !ok {
		return nil, ErrSecretNotFound
	}

	info := value.(*pb.SecretInfo)
	return info, nil
}

func (c *Cache) GetPolicy(key string) ([]*ladon.DefaultPolicy, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	ps, ok := c.policies.Get(key)
	if !ok {
		return nil, ErrPolicyNotFound
	}

	return ps.([]*ladon.DefaultPolicy), nil
}
