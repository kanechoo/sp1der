package task

import (
	"net/http"
	"sp1der/util"
	"sync"
	"time"
)

type HttpExecutor struct {
	httpClient   *http.Client
	targetChan   *chan []byte
	parallelSize *int
	sleepTime    *time.Duration
	wg           *sync.WaitGroup
	sourceChan   *chan string
}

func (e *HttpExecutor) UrlChannel(source *chan string) *HttpExecutor {
	e.sourceChan = source
	return e
}
func (e *HttpExecutor) HttpClient(httpClient *http.Client) *HttpExecutor {
	e.httpClient = httpClient
	return e
}
func (e *HttpExecutor) ParallelExecutorSize(parallelSize int) *HttpExecutor {
	e.parallelSize = &parallelSize
	return e
}
func (e *HttpExecutor) Sync(wg *sync.WaitGroup) *HttpExecutor {
	e.wg = wg
	return e
}
func (e *HttpExecutor) SleepDuration(sleepTime time.Duration) *HttpExecutor {
	e.sleepTime = &sleepTime
	return e
}
func (e *HttpExecutor) ResultChannel(targetChan *chan []byte) *HttpExecutor {
	e.targetChan = targetChan
	return e
}
func (e *HttpExecutor) Run() {
	for i := 0; i < *e.parallelSize; i++ {
		e.wg.Add(1)
		go e.execute()
	}
}
func (e *HttpExecutor) execute() {
	defer e.wg.Done()
loop:
	for {
		select {
		case v := <-*e.sourceChan:
			res := util.HttpGet(e.httpClient, v)
			*e.targetChan <- *res
			time.Sleep(*e.sleepTime)
		case <-time.After(10 * time.Second):
			break loop
		}
	}
}
