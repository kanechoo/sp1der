package unit

import (
	"sp1der/docs"
	"sync"
	"time"
)

type Processor struct {
	documentChan *chan []byte
	parallelSize *int
	wg           *sync.WaitGroup
	file         *string
}

func (p *Processor) DocumentChan(documentChan *chan []byte) *Processor {
	p.documentChan = documentChan
	return p
}
func (p *Processor) CallBack(f func(s *[]byte)) {
	for i := 0; i < *p.parallelSize; i++ {
		p.wg.Add(1)
		go p.execute(f)
	}
}
func (p *Processor) Parallel(parallelSize int) *Processor {
	p.parallelSize = &parallelSize
	return p
}
func (p *Processor) SelectorYamlFile(file string) *Processor {
	p.file = &file
	return p
}
func (p *Processor) ExecuteSelectorQuery() *[]map[string]string {
	items := make([]map[string]string, 0)
	for i := 0; i < *p.parallelSize; i++ {
		p.wg.Add(1)
		go p.execute(func(s *[]byte) {
			doc := docs.DocumentParser{}
			item := doc.ToDoc(s).SelectorYamlFile(*p.file).ExecuteSelectorQuery()
			items = append(items, *item...)
		})
	}
	p.wg.Wait()
	return &items
}
func (p *Processor) SyncWaitGroup(wg *sync.WaitGroup) *Processor {
	p.wg = wg
	return p
}
func (p *Processor) execute(f func(s *[]byte)) {
	defer p.wg.Done()
loop:
	for {
		select {
		case data := <-*p.documentChan:
			//execute next unit chain
			f(&data)
		case <-time.After(10 * time.Second):
			break loop
		}
	}
}
