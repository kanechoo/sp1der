package unit

import (
	"sp1der/docs"
	"sync"
	"time"
)

type Processor struct {
	DocumentChan     *chan []byte
	ParallelSize     int
	WaitGroup        *sync.WaitGroup
	SelectorYamlFile string
}

func (p *Processor) SetDocumentChan(documentChan *chan []byte) *Processor {
	p.DocumentChan = documentChan
	return p
}
func (p *Processor) CallBack(f func(s *[]byte)) {
	for i := 0; i < p.ParallelSize; i++ {
		p.WaitGroup.Add(1)
		go p.execute(f)
	}
}
func (p *Processor) SetParallel(parallelSize int) *Processor {
	p.ParallelSize = parallelSize
	return p
}
func (p *Processor) SetSelectorYamlFile(file string) *Processor {
	p.SelectorYamlFile = file
	return p
}
func (p *Processor) ExecuteSelectorQuery() *[]map[string]string {
	items := make([]map[string]string, 0)
	for i := 0; i < p.ParallelSize; i++ {
		p.WaitGroup.Add(1)
		go p.execute(func(s *[]byte) {
			doc := docs.DocumentParser{}
			item := doc.ToDoc(s).SelectorYamlFile(p.SelectorYamlFile).ExecuteSelectorQuery()
			items = append(items, *item...)
		})
	}
	p.WaitGroup.Wait()
	return &items
}
func (p *Processor) SetWaitGroup(wg *sync.WaitGroup) *Processor {
	p.WaitGroup = wg
	return p
}
func (p *Processor) execute(f func(s *[]byte)) {
	defer p.WaitGroup.Done()
loop:
	for {
		select {
		case data := <-*p.DocumentChan:
			//execute next unit chain
			f(&data)
		case <-time.After(10 * time.Second):
			break loop
		}
	}
}
