package clock

import (
	"sync"
	"time"
)

type TimeClock struct {
	// Frequency of the clock in Hz
	Frequency int
	// Current tick
	ticks uint

	standard Standard

	m *sync.Mutex

	Listeners []chan uint
}

func (c *TimeClock) Ticks() uint {
	return c.ticks
}

func (c *TimeClock) Standard() Standard {
	return c.standard
}

// New creates a new clock with the given frequency
func New(s Standard) *TimeClock {
	return &TimeClock{
		Frequency: StandardFrequencies[s],
		ticks:     0,
		standard:  s,
		Listeners: []chan uint{},
		m:         &sync.Mutex{},
	}
}

// AddListener adds a listener to the clock
func (c *TimeClock) AddListener(listener chan uint) {
	c.m.Lock()
	defer c.m.Unlock()
	c.Listeners = append(c.Listeners, listener)
}

// Tick ticks the clock and notifies all listeners
func (c *TimeClock) Tick() uint {
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

func (c *TimeClock) wait() uint {
	listener := make(chan uint)
	c.AddListener(listener)
	return <-listener
}

func (c *TimeClock) Wait(until uint) uint {
	for c.Ticks() < until {
		c.wait()
	}
	return c.Ticks()
}

// Start starts the clock
func (c *TimeClock) Start() {
	ticker := time.NewTicker(time.Second / time.Duration(c.Frequency))

	for {
		select {
		case <-ticker.C:
			c.Tick()
		}
	}
}
