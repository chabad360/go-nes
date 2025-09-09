package clock

import (
	"sync"
)

type CondClock struct {
	m *sync.RWMutex
	C *sync.Cond

	s     Standard
	ticks uint32

	held bool
}

func NewCondClock(s Standard) *CondClock {
	m := &sync.RWMutex{}
	return &CondClock{
		m: m,
		C: sync.NewCond(m.RLocker()),
		s: s,
	}
}

func (c *CondClock) Wait(until uint32) uint32 {
	c.C.L.Lock()
	defer c.C.L.Unlock()

	for until > c.ticks {
		c.held = true
		c.C.Wait()
	}
	return c.ticks
}

func (c *CondClock) Tick() uint32 {
	c.m.Lock()

	if !c.held {
		c.m.Unlock()
		return c.ticks
	}

	c.ticks++
	c.held = false
	t := c.ticks
	c.m.Unlock()

	c.C.Broadcast()
	return t
}

func (c *CondClock) Standard() Standard {
	return c.s
}

func (c *CondClock) Ticks() uint32 {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.ticks
}
