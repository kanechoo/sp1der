package counter

import (
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	Value int
}

func (c *Counter) Plus(delta int) {
	c.mu.Lock()
	c.Value = c.Value + delta
	c.mu.Unlock()
}
func (c *Counter) Sub(delta int) {
	c.mu.Lock()
	c.Value = c.Value - delta
	c.mu.Unlock()
}
