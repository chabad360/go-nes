package cpu

// PullProcessorStatus N | V | 0x20 | 0x10 | D | I | Z | C = stack[SP]
func PullProcessorStatus(cpu *CPU, operand1 uint8, operand2 uint8) {
	flags := cpu.Pop()
	cpu.Wait()

	cpu.N = flags&0x80 != 0
	cpu.V = flags&0x40 != 0
	cpu.D = flags&0x08 != 0
	cpu.I = flags&0x04 != 0
	cpu.Z = flags&0x02 != 0
	cpu.C = flags&0x01 != 0
	cpu.Wait()
}
