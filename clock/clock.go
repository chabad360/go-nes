package clock

import "time"

type Standard int

const (
	NTSC Standard = iota
	PAL
	Dendy
)

type Clock struct {
	// Frequency of the clock in Hz
	Frequency int
	// Current tick
	Ticks int

	Listeners []chan int
}

// New creates a new clock with the given frequency
func New(s Standard) *Clock {
	var frequency int
	switch s {
	case NTSC:
		frequency = 236250000 / 11
	case PAL:
		frequency = 26601712
	}
	return &Clock{
		Frequency: frequency,
		Ticks:     0,
		Listeners: []chan int{},
	}
}

// AddListener adds a listener to the clock
func (c *Clock) AddListener(listener chan int) {
	c.Listeners = append(c.Listeners, listener)
}

// Tick ticks the clock and notifies all listeners
func (c *Clock) Tick() {
	c.Ticks++
	for _, listener := range c.Listeners {
		listener <- c.Ticks
	}
}

// Start starts the clock
func (c *Clock) Start() {
	ticker := time.NewTicker(time.Second / time.Duration(c.Frequency))

	for {
		select {
		case <-ticker.C:
			c.Tick()
		}
	}
}
