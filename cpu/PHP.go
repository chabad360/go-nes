package cpu

// PushProcessorStatus stack[SP] = N | V | 0x20 | 0x10 | D | I | Z | C
func PushProcessorStatus(cpu *CPU, operand1 uint8, operand2 uint8) {
	flags := uint8(0x30)
	if cpu.N {
		flags |= 0x80
	}
	if cpu.V {
		flags |= 0x40
	}
	if cpu.D {
		flags |= 0x08
	}
	if cpu.I {
		flags |= 0x04
	}
	if cpu.Z {
		flags |= 0x02
	}
	if cpu.C {
		flags |= 0x01
	}
	cpu.Push(flags)
	cpu.Wait()
}
