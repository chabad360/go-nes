package clock

import (
	"sync"
)

type Channel2Clock struct {
	// Current tick
	ticks uint32

	standard Standard

	m *sync.Mutex

	cpuListener chan uint32
	ppuListener chan uint32
}

func (c *Channel2Clock) Ticks() uint32 {
	return c.ticks
}

func (c *Channel2Clock) Standard() Standard {
	return c.standard
}

// NewChannel2Clock creates a new clock with the given frequency
func NewChannel2Clock(s Standard) *Channel2Clock {
	return &Channel2Clock{
		ticks:    0,
		standard: s,

		cpuListener: make(chan uint32),
		ppuListener: make(chan uint32),
	}
}

// Tick ticks the clock and notifies all listeners
func (c *Channel2Clock) Tick() uint32 {
	c.ticks++
	c.cpuListener <- c.ticks
	c.ppuListener <- c.ticks
	return c.ticks
}

func (c *Channel2Clock) Wait(until uint32) uint32 {
	for c.Ticks() < until {
		<-c.cpuListener
	}
	return c.Ticks()
}

func (c *Channel2Clock) PPUWait(until uint32) uint32 {
	for c.Ticks() < until {
		<-c.ppuListener
	}
	return c.Ticks()
}
