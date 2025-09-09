package clock

import (
	"sync"
)

type ChannelClock struct {
	// Current tick
	ticks uint32

	standard Standard

	m *sync.Mutex

	Listeners []chan uint32
}

func (c *ChannelClock) Ticks() uint32 {
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
		Listeners: []chan uint32{},
		m:         &sync.Mutex{},
	}
}

// AddListener adds a listener to the clock
func (c *ChannelClock) AddListener(listener chan uint32) {
	c.m.Lock()
	defer c.m.Unlock()
	c.Listeners = append(c.Listeners, listener)
}

// Tick ticks the clock and notifies all listeners
func (c *ChannelClock) Tick() uint32 {
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

func (c *ChannelClock) wait() uint32 {
	listener := make(chan uint32)
	c.AddListener(listener)
	return <-listener
}

func (c *ChannelClock) Wait(until uint32) uint32 {
	for c.Ticks() < until {
		c.wait()
	}
	return c.Ticks()
}
