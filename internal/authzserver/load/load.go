package load

import (
	"context"
	"github.com/marmotedu/errors"
	"github.com/wangzhen94/iam/pkg/log"
	"github.com/wangzhen94/iam/pkg/storage"
	"sync"
	"time"
)

type Loader interface {
	Reload() error
}

type Load struct {
	ctx    context.Context
	lock   *sync.RWMutex
	loader Loader
}

func (l *Load) Start() {
	go startPubSubLoop()
	go l.reloadQueueLoop()
	go l.reloadLoop()
	l.DoReload()
}

func (l *Load) reloadQueueLoop(cb ...func()) {
	for {
		select {
		case <-l.ctx.Done():
			return
		case fn := <-reloadQueue:
			requeueLock.Lock()
			requeue = append(requeue, fn)
			requeueLock.Unlock()
			log.Info("Reload queued")
			if len(cb) != 0 {
				cb[0]()
			}
		}
	}
}

func (l *Load) reloadLoop(complete ...func()) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-l.ctx.Done():
			return
		case <-ticker.C:
			cb, ok := shouldReload()
			if !ok {
				continue
			}
			start := time.Now()
			l.DoReload()
			for _, c := range cb {
				if c != nil {
					c()
				}
			}
			if len(complete) != 0 {
				complete[0]()
			}

			log.Infof("reload: cycle completed in %v", time.Since(start))
		}
	}
}

func shouldReload() ([]func(), bool) {
	requeueLock.Lock()
	defer requeueLock.Lock()

	if len(requeue) == 0 {
		return nil, false
	}

	n := requeue
	requeue = []func(){}

	return n, true
}

func startPubSubLoop() {
	cacheStore := storage.RedisCluster{}
	cacheStore.Connect()

	for {
		err := cacheStore.StartPubSubHandler(RedisPubSubChannel, func(v interface{}) {
			handleRedisEvent(v, nil, nil)
		})

		if err != nil {
			if !errors.Is(err, storage.ErrRedisIsDown) {
				log.Errorf("Connection to Redis failed, reconnect in 10s: %s", err.Error())
			}

			time.Sleep(10 * time.Second)
			log.Warnf("Reconnecting: %s", err.Error())
		}
	}
}

func (l *Load) DoReload() {
	l.lock.Lock()
	defer l.lock.Unlock()

	if err := l.loader.Reload(); err != nil {
		log.Errorf("faild to refresh target storage: %s", err.Error())
	}

	log.Debugf("refresh target storage succ")
}

func NewLoader(ctx context.Context, loader Loader) *Load {
	return &Load{
		ctx:    ctx,
		lock:   new(sync.RWMutex),
		loader: loader,
	}
}
