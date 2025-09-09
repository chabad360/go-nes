package main

import (
	"context"
	"go-nes/cart"
	"go-nes/clock"
	"go-nes/cpu"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	i := make(chan os.Signal, 1)
	signal.Notify(i, os.Interrupt, os.Kill)

	// Load the Official instructions test cartridge
	cartridge, err := cart.NewCartidge("cpu/testdata/official_only.nes")
	if err != nil {
		slog.Error("Failed to load Official cartridge", "error", err)
		return
	}

	c := clock.NewRuntimeClock(clock.NTSC)

	CPU := &cpu.CPU{
		Mapper: &cpu.NROM{
			PRGROM: cartridge.PRGROM,
			CHRROM: cartridge.CHRROM,
			RAM:    [0x2000]uint8{},
		},
		Clock: c,

		//Log: true,
	}

	if len(cartridge.PRGROM) == 0 {
		slog.Error("No PRG ROM found in cartridge")
		return
	}

	slog.Info("Starting clock")
	done := clock.StartClock(context.Background(), c)

	slog.Info("Powering on CPU")
	CPU.PowerOn()

	slog.Info("Starting CPU loop")
	go func() {
		for {
			CPU.Step()
		}
	}()

	<-i
	slog.Info("Stopping clock")
	done()
	slog.Info("Exiting")
}
