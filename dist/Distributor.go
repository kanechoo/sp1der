package dist

import (
	"fmt"
	"sp1der/task"
	"sp1der/util"
	"sync"
	"time"
)

type HttpWalker struct {
	UrlChan            *chan string
	HttpClientPoolSize int
	DocumentChan       *chan []byte
	WaitGroup          *sync.WaitGroup
	SleepTime          time.Duration
}

func (d *HttpWalker) UrlStoreChan(urlChannel *chan string) *HttpWalker {
	d.UrlChan = urlChannel
	return d
}
func (d *HttpWalker) SetHttpClientPoolSize(size int) *HttpWalker {
	d.HttpClientPoolSize = size
	return d
}
func (d *HttpWalker) SetWaitGroup(waitGroup *sync.WaitGroup) *HttpWalker {
	d.WaitGroup = waitGroup
	return d
}
func (d *HttpWalker) SetSleepTime(duration time.Duration) *HttpWalker {
	d.SleepTime = duration
	return d
}
func (d *HttpWalker) SetDocumentChan(documentChan *chan []byte) *HttpWalker {
	d.DocumentChan = documentChan
	return d
}
func (d *HttpWalker) Walk() {
	//check param
	if nil == d.UrlChan {
		panic(fmt.Sprintf("You must be define UrlChan if you want to start distribute"))
	}
	if nil == d.WaitGroup {
		panic(fmt.Sprintf("You must be define waitGrop if you want to start distribute"))
	}
	if nil == &d.HttpClientPoolSize {
		panic(fmt.Sprintf("You must be define HttpClientPoolSize if you want to start distribute"))
	}
	if nil == &d.SleepTime {
		panic(fmt.Sprintf("You must be define SleepTime if you want to start distribute"))
	}
	if nil == d.DocumentChan {
		panic(fmt.Sprintf("You must be define DocumentChan if you want to start distribute"))
	}
	executor := task.HttpExecutor{}
	executor.HttpClient(util.DefaultHttpClient()).
		ParallelExecutorSize(d.HttpClientPoolSize).
		SleepDuration(d.SleepTime).
		SyncWaitGroup(d.WaitGroup).UrlChannel(d.UrlChan).
		ResultChannel(d.DocumentChan).Run()
}
func (d *HttpWalker) UrlGenerateFunc(f func() string) *HttpWalker {
	if nil == d.UrlChan {
		panic("Please defined a url channel for the HttpWalker struct first then run this function")
	}
	go func() {
		for {
			s := f()
			if "" == s {
				break
			}
			select {
			case *d.UrlChan <- s:
			case <-time.After(30 * time.Second):
				fmt.Printf("Timeout\n")
			}
		}
	}()
	return d
}
