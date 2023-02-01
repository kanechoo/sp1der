package dist

import (
	"fmt"
	"sp1der/task"
	"sp1der/util"
	"sync"
	"time"
)

type Distributor struct {
	urlChan            *chan string
	httpClientPoolSize *int
	documentChan       *chan []byte
	waitGroup          *sync.WaitGroup
	requestSleepTime   *time.Duration
}

func (d *Distributor) UrlChan(urlChannel *chan string) *Distributor {
	d.urlChan = urlChannel
	return d
}
func (d *Distributor) HttpClientPoolSize(size int) *Distributor {
	d.httpClientPoolSize = &size
	return d
}
func (d *Distributor) SyncWaitGroup(waitGroup *sync.WaitGroup) *Distributor {
	d.waitGroup = waitGroup
	return d
}
func (d *Distributor) RequestSleepTime(duration time.Duration) *Distributor {
	d.requestSleepTime = &duration
	return d
}
func (d *Distributor) DocumentChan(documentChan *chan []byte) *Distributor {
	d.documentChan = documentChan
	return d
}
func (d *Distributor) Distribute() {
	//check param
	if nil == d.urlChan {
		panic(fmt.Sprintf("You must be define urlChan if you want to start distribute"))
	}
	if nil == d.waitGroup {
		panic(fmt.Sprintf("You must be define waitGrop if you want to start distribute"))
	}
	if nil == d.httpClientPoolSize {
		panic(fmt.Sprintf("You must be define httpClientPoolSize if you want to start distribute"))
	}
	if nil == d.requestSleepTime {
		panic(fmt.Sprintf("You must be define requestSleepTime if you want to start distribute"))
	}
	if nil == d.documentChan {
		panic(fmt.Sprintf("You must be define documentChan if you want to start distribute"))
	}
	executor := task.HttpExecutor{}
	executor.HttpClient(util.DefaultHttpClient()).
		ParallelExecutorSize(*d.httpClientPoolSize).
		SleepDuration(*d.requestSleepTime).
		SyncWaitGroup(d.waitGroup).UrlChannel(d.urlChan).
		ResultChannel(d.documentChan).Run()
}
func (d *Distributor) NextUrlFunc(f func() string) *Distributor {
	if nil == d.urlChan {
		panic("Please defined a url channel for the Distributor struct first then run this function")
	}
	go func() {
		for {
			s := f()
			if "" == s {
				break
			}
			select {
			case *d.urlChan <- s:
				fmt.Printf("Added request url to chan : %s\n", s)
			case <-time.After(30 * time.Second):
				fmt.Printf("Timeout\n")
			}
		}
	}()
	return d
}
