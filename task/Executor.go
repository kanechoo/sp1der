package task

import (
	"fmt"
	"net/http"
	"sp1der/util/httpv2"
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
func (e *HttpExecutor) SyncWaitGroup(wg *sync.WaitGroup) *HttpExecutor {
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
			fmt.Printf("Request : %s\n", v)
			oneRequest := func() (*[]byte, error) {
				return httpv2.GetRequest(e.httpClient, v)
			}
			res, err := tryToRequest(3, oneRequest)
			if nil != err {
				fmt.Printf("Can't get any response from : %s\n", v)
			} else {
				*e.targetChan <- *res
			}
			time.Sleep(*e.sleepTime)
		case <-time.After(10 * time.Second):
			break loop
		}
	}
}

func tryToRequest(tryTimes int, fun func() (*[]byte, error)) (*[]byte, error) {
	b := make([]byte, 0)
	var newErr error = nil
	for i := 1; i < tryTimes; i++ {
		result, err := fun()
		if nil != err {
			newErr = err
			time.Sleep(200 * time.Millisecond)
			continue
		} else {
			return result, nil
		}
	}
	return &b, newErr
}
