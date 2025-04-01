package cpu

type WriteFunc func(cpu *CPU, clock <-chan int, address uint16, value uint8)
type ReadFunc func(cpu *CPU, clock <-chan int, address uint16) uint8

type Mapper interface {
	Write(cpu *CPU, address uint16, value uint8)
	Read(cpu *CPU, address uint16) uint8
}

type GenericMapper struct{}

func (g *GenericMapper) Write(cpu *CPU, address uint16, value uint8) {
	switch {
	case address < 0x2000:
		cpu.Memory[address%0x0800] = value
	case address < 0x4000:
		// PPU registers
	case address < 0x4018:
		// APU/IO registers
	case address < 0x4020:
		// APU/IO registers
	}
	return // ignore writes to other addresses
}

func (g *GenericMapper) Read(cpu *CPU, address uint16) uint8 {
	switch {
	case address < 0x2000:
		return cpu.Memory[address%0x0800]
	case address < 0x4000:
		// PPU registers
	case address < 0x4018:
		// APU/IO registers
	case address < 0x4020:
		// APU/IO registers
	}
	return 0 // ignore reads from other addresses
}
