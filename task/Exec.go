package task

import (
	"net/http"
	"sp1der/util"
	"sync"
	"time"
)

type Executor[S any, T any] struct {
	httpClient   *http.Client
	targetChan   *chan T
	parallelSize *int
	wg           *sync.WaitGroup
	sourceChan   *chan S
}
type Builder[S any, T any] struct {
	sourceChan   *chan S
	targetChan   *chan T
	wg           *sync.WaitGroup
	parallelSize *int
	httpClient   *http.Client
	s            func(s *S) *string
	t            func(b *[]byte) *T
}

func (e *Executor[S, T]) Source(source *chan S) *Executor[S, T] {
	e.sourceChan = source
	return e
}
func (e *Executor[S, T]) HttpClient(httpClient *http.Client) *Executor[S, T] {
	e.httpClient = httpClient
	return e
}
func (e *Executor[S, T]) ParallelSize(size int) *Executor[S, T] {
	e.parallelSize = &size
	return e
}
func (e *Executor[S, T]) Sync(wg *sync.WaitGroup) *Executor[S, T] {
	e.wg = wg
	return e
}
func (e *Executor[S, T]) Target(targetChan *chan T) *Executor[S, T] {
	e.targetChan = targetChan
	return e
}
func (e *Executor[S, T]) Build() *Builder[S, T] {
	return &Builder[S, T]{sourceChan: e.sourceChan, targetChan: e.targetChan, wg: e.wg, parallelSize: e.parallelSize, httpClient: e.httpClient}
}
func (b *Builder[S, T]) S(s func(s *S) *string) *Builder[S, T] {
	b.s = s
	return b
}
func (b *Builder[S, T]) T(t func(b *[]byte) *T) *Builder[S, T] {
	b.t = t
	return b
}
func (b *Builder[S, T]) Run() {
	for i := 0; i < *b.parallelSize; i++ {
		b.wg.Add(1)
		go b.execute()
	}
}
func (b *Builder[S, T]) execute() {
	defer b.wg.Done()
loop:
	for {
		select {
		case v := <-*b.sourceChan:
			res := util.DefaultHttpGet(b.httpClient, *b.s(&v))
			*b.targetChan <- *b.t(res)
			time.Sleep(500 * time.Millisecond)
		case <-time.After(10 * time.Second):
			break loop
		}
	}
}
