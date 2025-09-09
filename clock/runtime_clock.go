package clock

import (
	"runtime"
	"sync/atomic"
)

type RuntimeClock struct {
	s Standard

	ticks *uint32
	held  *uint32
}

func NewRuntimeClock(s Standard) *RuntimeClock {
	return &RuntimeClock{
		ticks: new(uint32),
		held:  new(uint32),
		s:     s,
	}
}

func (c *RuntimeClock) Wait(until uint32) uint32 {
	t := atomic.LoadUint32(c.ticks)
	for until > t {
		atomic.StoreUint32(c.held, 1)
		runtime.Gosched()
		t = atomic.LoadUint32(c.ticks)

	}
	return t
}

func (c *RuntimeClock) Tick() uint32 {
	if !atomic.CompareAndSwapUint32(c.held, 1, 0) {
		return atomic.LoadUint32(c.ticks)
	}

	return atomic.AddUint32(c.ticks, 1)
}

func (c *RuntimeClock) Standard() Standard {
	return c.s
}

func (c *RuntimeClock) Ticks() uint32 {
	return atomic.LoadUint32(c.ticks)
}
