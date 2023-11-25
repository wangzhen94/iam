package analytics

import (
	"github.com/vmihailenco/msgpack/v5"
	"github.com/wangzhen94/iam/pkg/log"
	"github.com/wangzhen94/iam/pkg/storage"
	"sync"
	"sync/atomic"
	"time"
)

const analyticsKeyName = "iam-system-analytics"

const (
	recordsBufferForcedFlushInterval = 1 * time.Second
)

type AnalyticsRecord struct {
	TimeStamp  int64     `json:"timestamp"`
	Username   string    `json:"username"`
	Effect     string    `json:"effect"`
	Conclusion string    `json:"conclusion"`
	Request    string    `json:"request"`
	Policies   string    `json:"policies"`
	Deciders   string    `json:"deciders"`
	ExpireAt   time.Time `json:"expireAt"   bson:"expireAt"`
}

var analytics *Analytics

type Analytics struct {
	store                      storage.AnalyticsHandler
	poolSize                   int
	recordsChan                chan *AnalyticsRecord
	workerBufferSize           uint64
	recordsBufferFlushInterval uint64
	shouldStop                 uint32
	poolWg                     sync.WaitGroup
}

func GetAnalytics() *Analytics {
	return analytics
}

func (a *Analytics) Start() {
	a.store.Connect()

	atomic.SwapUint32(&a.shouldStop, 0)
	for i := 0; i < a.poolSize; i++ {
		a.poolWg.Add(1)
		go a.recordWorker()
	}

}

func (r *Analytics) recordWorker() {
	defer r.poolWg.Done()

	recordsBuffer := make([][]byte, 0, r.workerBufferSize)
	lastSendTS := time.Now()
	for {
		var readyToSend bool
		select {
		case record, ok := <-r.recordsChan:
			if !ok {
				r.store.AppendToSetPipelined(analyticsKeyName, recordsBuffer)

				return
			}

			if encoded, err := msgpack.Marshal(record); err != nil {
				log.Errorf("Error encoding analytics data: %s", err.Error())
			} else {
				recordsBuffer = append(recordsBuffer, encoded)
			}

			readyToSend = uint64(len(recordsBuffer)) == r.workerBufferSize

		case <-time.After(time.Duration(r.recordsBufferFlushInterval) * time.Millisecond):
			readyToSend = true
		}

		if len(recordsBuffer) > 0 && (readyToSend || time.Since(lastSendTS) > recordsBufferForcedFlushInterval) {
			r.store.AppendToSetPipelined(analyticsKeyName, recordsBuffer)
			recordsBuffer = recordsBuffer[:0]
			lastSendTS = time.Now()
		}

	}

}

// RecordHit will store an AnalyticsRecord in Redis.
func (r *Analytics) RecordHit(record *AnalyticsRecord) error {
	// check if we should stop sending records 1st
	if atomic.LoadUint32(&r.shouldStop) > 0 {
		return nil
	}

	// just send record to channel consumed by pool of workers
	// leave all data crunching and Redis I/O work for pool workers
	r.recordsChan <- record

	return nil
}

//var analytics *

func NewAnalytics(options *AnalyticsOptions, store storage.AnalyticsHandler) *Analytics {
	ps := options.PoolSize
	recordsBufferSize := options.RecordsBufferSize
	workerBufferSize := recordsBufferSize / uint64(ps)
	log.Debug("Analytics pool worker buffer size", log.Uint64("workerBufferSize", workerBufferSize))

	recordsChan := make(chan *AnalyticsRecord, recordsBufferSize)

	analytics = &Analytics{
		store:                      store,
		poolSize:                   ps,
		recordsChan:                recordsChan,
		workerBufferSize:           workerBufferSize,
		recordsBufferFlushInterval: options.FlushInterval,
	}

	return analytics
}

// Stop stop the analytics service.
func (r *Analytics) Stop() {
	// flag to stop sending records into channel
	atomic.SwapUint32(&r.shouldStop, 1)

	// close channel to stop workers
	close(r.recordsChan)

	// wait for all workers to be done
	r.poolWg.Wait()
}
