package clock

import (
	"runtime"
	"sync/atomic"
)

type RuntimeClock struct {
	s     Standard
	ticks *atomic.Uint32

	held *atomic.Bool
}

func NewRuntimeClock(s Standard) *RuntimeClock {
	return &RuntimeClock{
		ticks: &atomic.Uint32{},
		held:  &atomic.Bool{},
		s:     s,
	}
}

func (c *RuntimeClock) Wait(until uint) uint {
	t := c.ticks.Load()
	for uint32(until) > t {
		c.held.Store(true)
		runtime.Gosched()
		t = c.ticks.Load()
	}
	return uint(t)
}

func (c *RuntimeClock) Tick() uint {
	if !c.held.CompareAndSwap(true, false) {
		return uint(c.ticks.Load())
	}

	t := c.ticks.Add(1)
	return uint(t)
}

func (c *RuntimeClock) Standard() Standard {
	return c.s
}

func (c *RuntimeClock) Ticks() uint {
	return uint(c.ticks.Load())
}
