package clock

import (
	"sync"
)

type Channel2Clock struct {
	// Current tick
	ticks uint

	standard Standard

	m *sync.Mutex

	cpuListener chan uint
	ppuListener chan uint
}

func (c *Channel2Clock) Ticks() uint {
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

		cpuListener: make(chan uint),
		ppuListener: make(chan uint),
	}
}

// Tick ticks the clock and notifies all listeners
func (c *Channel2Clock) Tick() uint {
	c.ticks++
	c.cpuListener <- c.ticks
	c.ppuListener <- c.ticks
	return c.ticks
}

func (c *Channel2Clock) Wait(until uint) uint {
	for c.Ticks() < until {
		<-c.cpuListener
	}
	return c.Ticks()
}

func (c *Channel2Clock) PPUWait(until uint) uint {
	for c.Ticks() < until {
		<-c.ppuListener
	}
	return c.Ticks()
}
