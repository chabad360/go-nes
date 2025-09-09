package clock

import (
	"context"
	"time"
)

type Standard int

func (s Standard) Frequency() uint32 {
	return StandardFrequencies[s]
}

func (s Standard) CPUDivider() uint32 {
	return StandardCPUDividers[s]
}

const (
	NTSC Standard = iota
	PAL
	Dendy
)

var StandardFrequencies = map[Standard]uint32{
	NTSC:  236250000 / 11,
	PAL:   26601712,
	Dendy: 26601712,
}

var StandardCPUDividers = map[Standard]uint32{
	NTSC:  12,
	PAL:   16,
	Dendy: 15,
}

type Clock interface {
	// Wait waits until the specified number of ticks have passed and returns the current tick count
	Wait(uint32) uint32

	// Tick ticks the clock and returns the current tick count
	Tick() uint32

	// Standard returns the clock standard (NTSC, PAL, Dendy)
	Standard() Standard

	// Ticks returns the current tick count
	Ticks() uint32
}

type PPUClock interface {
	Clock
	PPUWait(uint32) uint32
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
