package clock

import (
	"runtime"
)

type AtomiclessClock struct {
	s Standard

	ticks *uint32
	held  *bool
}

func NewAtomiclessClock(s Standard) *AtomiclessClock {
	return &AtomiclessClock{
		ticks: new(uint32),
		held:  new(bool),
		s:     s,
	}
}

func (c *AtomiclessClock) Wait(until uint32) uint32 {
	t := *c.ticks
	for until > t {
		*c.held = true
		runtime.Gosched()
		t = *c.ticks
	}
	return t
}

func (c *AtomiclessClock) Tick() uint32 {
	if !*c.held {
		return *c.ticks
	}

	*c.held = false
	*c.ticks++
	return *c.ticks
}

func (c *AtomiclessClock) Standard() Standard {
	return c.s
}

func (c *AtomiclessClock) Ticks() uint32 {
	return *c.ticks
}
