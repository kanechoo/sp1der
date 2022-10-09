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
	wg           *sync.WaitGroup
	sourceChan   *chan string
}

func (e *HttpExecutor) Source(source *chan string) *HttpExecutor {
	e.sourceChan = source
	return e
}
func (e *HttpExecutor) HttpClient(httpClient *http.Client) *HttpExecutor {
	e.httpClient = httpClient
	return e
}
func (e *HttpExecutor) ParallelSize(size int) *HttpExecutor {
	e.parallelSize = &size
	return e
}
func (e *HttpExecutor) Sync(wg *sync.WaitGroup) *HttpExecutor {
	e.wg = wg
	return e
}
func (e *HttpExecutor) Target(targetChan *chan []byte) *HttpExecutor {
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
			time.Sleep(500 * time.Millisecond)
		case <-time.After(10 * time.Second):
			break loop
		}
	}
}
