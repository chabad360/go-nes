package clock

import (
	"context"
	"time"
)

type Standard int

func (s Standard) Frequency() uint {
	return StandardFrequencies[s]
}

func (s Standard) CPUDivider() uint {
	return StandardCPUDividers[s]
}

const (
	NTSC Standard = iota
	PAL
	Dendy
)

var StandardFrequencies = map[Standard]uint{
	NTSC:  236250000 / 11,
	PAL:   26601712,
	Dendy: 26601712,
}

var StandardCPUDividers = map[Standard]uint{
	NTSC:  12,
	PAL:   16,
	Dendy: 15,
}

type Clock interface {
	// Wait waits until the specified number of ticks have passed and returns the current tick count
	Wait(uint) uint

	// Tick ticks the clock and returns the current tick count
	Tick() uint

	// Standard returns the clock standard (NTSC, PAL, Dendy)
	Standard() Standard

	// Ticks returns the current tick count
	Ticks() uint
}

type PPUClock interface {
	Clock
	PPUWait(uint) uint
}

func StartClock(ctx context.Context, c Clock) context.CancelFunc {
	ticker := time.NewTicker(time.Second / time.Duration(c.Standard().Frequency()))
	_, cancel := context.WithCancel(ctx)

	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				c.Tick()
			}
		}
	}()

	return cancel
}
