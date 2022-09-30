package process

import (
	"fmt"
	"sync"
	"time"
)

type Processor[S any] struct {
	sourceChan     *chan S
	processor      func(b *S, t *Transport) *Transport
	processorCount int
	processorChain []func(x *Transport) *Transport
	tr             *Transport
	parallelSize   *int
	wg             *sync.WaitGroup
}
type ProcessorFunc[T any] func(b *T, t *Transport) *Transport
type Transport struct {
	Data interface{}
}

func (p *Processor[T]) Source(pullChan *chan T) *Processor[T] {
	p.sourceChan = pullChan
	return p
}
func (p *Processor[S]) Processor(processor ProcessorFunc[S]) *Processor[S] {
	if p.processorCount == 0 {
		p.processor = processor
		p.processorCount = 1
	} else {
		panic(fmt.Sprintf("only support one processor func define"))
	}
	return p
}

// NextProcessor first processor next type is your channel type
func (p *Processor[S]) NextProcessor(processor func(x *Transport) *Transport) *Processor[S] {
	//check processor exists
	if p.processorCount != 1 {
		panic(fmt.Sprintf("must be define processor func first"))
	}
	if nil != processor {
		p.processorChain = append(p.processorChain, processor)
	}
	return p
}
func (p *Processor[S]) ParallelSize(parallelSize int) *Processor[S] {
	p.parallelSize = &parallelSize
	return p
}
func (p *Processor[S]) Sync(wg *sync.WaitGroup) *Processor[S] {
	p.wg = wg
	return p
}
func (p *Processor[S]) Run() {
	for i := 0; i < *p.parallelSize; i++ {
		p.wg.Add(1)
		go p.execute()
	}
}
func (p *Processor[S]) execute() {
	defer p.wg.Done()
loop:
	for {
		select {
		case data := <-*p.sourceChan:
			//user processor chain
			p.tr = &Transport{
				Data: data,
			}
			//execute first process
			p.tr = p.processor(&data, &Transport{})
			//execute next processor chain
			for i := range p.processorChain {
				if 0 == i {
					p.tr = p.processorChain[i](p.tr)
				} else {
					p.tr = (p.processorChain[i])(p.tr)
				}
			}
		case <-time.After(10 * time.Second):
			break loop
		}
	}
}
