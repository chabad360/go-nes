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

func (c *RuntimeClock) Wait(until uint) uint {
	t := atomic.LoadUint32(c.ticks)
	for uint32(until) > t {
		atomic.StoreUint32(c.held, 1)
		runtime.Gosched()
		t = atomic.LoadUint32(c.ticks)

	}
	return uint(t)
}

func (c *RuntimeClock) Tick() uint {
	if !atomic.CompareAndSwapUint32(c.held, 1, 0) {
		return uint(atomic.LoadUint32(c.ticks))
	}

	return uint(atomic.AddUint32(c.ticks, 1))
}

func (c *RuntimeClock) Standard() Standard {
	return c.s
}

func (c *RuntimeClock) Ticks() uint {
	return uint(atomic.LoadUint32(c.ticks))
}
