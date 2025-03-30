package cpu

type NROM struct {
	GenericMapper
	PRGROM []uint8

	CHRROM []uint8
}

func (n *NROM) Write(cpu *CPU, clock <-chan int, address uint16, value uint8) {
	n.GenericMapper.Write(cpu, clock, address, value) // NROM has no special write behavior
}

func (n *NROM) Read(cpu *CPU, clock <-chan int, address uint16) uint8 {
	if address >= 0x8000 {
		if address < 0xBFFF {
			return n.PRGROM[address-0x8000]
		} else if address < 0xFFFF {
			address = address - 0x8000
			if len(n.PRGROM) <= 0x4000 {
				address = address % 0x4000 // Loop back if PRGROM is less than 16KB
			}
			return n.PRGROM[address]
		}
	}
	return n.GenericMapper.Read(cpu, clock, address)
}
