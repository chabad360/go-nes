package clock

type Standard int

func (s Standard) Frequency() int {
	return StandardFrequencies[s]
}

func (s Standard) CPUDivider() int {
	return StandardCPUDividers[s]
}

const (
	NTSC Standard = iota
	PAL
	Dendy
)

var StandardFrequencies = map[Standard]int{
	NTSC:  236250000 / 11,
	PAL:   26601712,
	Dendy: 26601712,
}

var StandardCPUDividers = map[Standard]int{
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
