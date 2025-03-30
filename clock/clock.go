package clock

import "time"

type Clock struct {
	// Frequency of the clock in Hz
	Frequency int
	// Current tick
	Ticks int

	Listeners []chan int
}

// New creates a new clock with the given frequency
func New(frequency int) *Clock {
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

// RemoveListener removes a listener from the clock
func (c *Clock) RemoveListener(listener chan int) {
	for i, l := range c.Listeners {
		if l == listener {
			c.Listeners = append(c.Listeners[:i], c.Listeners[i+1:]...)
			break
		}
	}
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
	for {
		c.Tick()
		// Sleep for the duration of one clock tick
		time.Sleep(time.Second / time.Duration(c.Frequency))
	}
}
