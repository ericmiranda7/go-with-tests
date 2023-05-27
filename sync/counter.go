package sync

import "sync"

type Counter struct {
	v  int
	mu sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.v += 1
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	return c.v
}
