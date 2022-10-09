package process

import (
	"sync"
	"time"
)

type Processor struct {
	sourceChan   *chan []byte
	parallelSize *int
	wg           *sync.WaitGroup
}

func (p *Processor) Source(pullChan *chan []byte) *Processor {
	p.sourceChan = pullChan
	return p
}
func (p *Processor) Run(f func(s *[]byte)) {
	for i := 0; i < *p.parallelSize; i++ {
		p.wg.Add(1)
		go p.execute(f)
	}
}
func (p *Processor) ParallelSize(parallelSize int) *Processor {
	p.parallelSize = &parallelSize
	return p
}
func (p *Processor) Sync(wg *sync.WaitGroup) *Processor {
	p.wg = wg
	return p
}
func (p *Processor) execute(f func(s *[]byte)) {
	defer p.wg.Done()
loop:
	for {
		select {
		case data := <-*p.sourceChan:
			//execute next processor chain
			f(&data)
		case <-time.After(10 * time.Second):
			break loop
		}
	}
}
