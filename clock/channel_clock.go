package clock

import (
	"sync"
)

type ChannelClock struct {
	// Current tick
	ticks uint

	standard Standard

	m *sync.Mutex

	Listeners []chan uint
}

func (c *ChannelClock) Ticks() uint {
	return c.ticks
}

func (c *ChannelClock) Standard() Standard {
	return c.standard
}

// NewChannelClock creates a new clock with the given frequency
func NewChannelClock(s Standard) *ChannelClock {
	return &ChannelClock{
		ticks:     0,
		standard:  s,
		Listeners: []chan uint{},
		m:         &sync.Mutex{},
	}
}

// AddListener adds a listener to the clock
func (c *ChannelClock) AddListener(listener chan uint) {
	c.m.Lock()
	defer c.m.Unlock()
	c.Listeners = append(c.Listeners, listener)
}

// Tick ticks the clock and notifies all listeners
func (c *ChannelClock) Tick() uint {
	c.m.Lock()
	defer c.m.Unlock()
	if len(c.Listeners) == 0 {
		return c.ticks
	}
	c.ticks++
	for _, listener := range c.Listeners {
		listener <- c.ticks
	}
	c.Listeners = c.Listeners[:0]
	return c.ticks
}

func (c *ChannelClock) wait() uint {
	listener := make(chan uint)
	c.AddListener(listener)
	return <-listener
}

func (c *ChannelClock) Wait(until uint) uint {
	for c.Ticks() < until {
		c.wait()
	}
	return c.Ticks()
}
