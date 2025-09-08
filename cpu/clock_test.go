package cpu

import (
	"go-nes/cart"
	"go-nes/clock"
	"testing"
)

func BenchmarkClock_NESTest(b *testing.B) {
	clocks := []struct {
		name string
		new  func() clock.Clock
	}{
		{"ChannelClock", func() clock.Clock { return clock.New(clock.NTSC) }},
		{"CondClock", func() clock.Clock { return clock.NewCondClock(clock.NTSC) }},
		{"RuntimeClock", func() clock.Clock { return clock.NewRuntimeClock(clock.NTSC) }},
	}
	for _, clk := range clocks {
		b.Run(clk.name, func(b *testing.B) {
			b.ReportAllocs()
			// Load the NESTest cartridge
			cartidge, err := cart.NewCartidge("testdata/nestest.nes")
			if err != nil {
				b.Fatalf("Failed to load NESTest cartidge: %v", err)
			}

			c := clk.new()

			cpu := &CPU{
				Mapper: &NROM{
					PRGROM: cartidge.PRGROM,
					CHRROM: cartidge.CHRROM,
				},
				Clock: c,
				//Log:   true,
			}

			go func() {
				//ti := time.NewTicker(time.Second / (236250000 / 11))
				for {
					select {
					case <-b.Context().Done():
						return
					default:
						c.Tick()
					}
				}
			}()

			cpu.PowerOn()
			cpu.PC = 0xC000 // Start at the beginning of nestest
			b.ResetTimer()
			for b.Loop() {
				exec := cpu.Decode()
				cpu.Fetch(exec)
				cpu.Execute(exec)
			}
			b.StopTimer()

			result1 := cpu.Memory[0x0002]
			result2 := cpu.Memory[0x0003]

			if result1 != 0x00 {
				b.Errorf("Expected 0x00 at address 0x0002, got 0x%02X", result1)
			}
			if result2 != 0x00 {
				b.Errorf("Expected 0x00 at address 0x0003, got 0x%02X", result2)
			}

			//b.Logf("Cycles: %d", cpu.Cycles)
		})
	}
}
