package cpu

// PullProcessorStatus N | V | 0x20 | 0x10 | D | I | Z | C = stack[SP]
func PullProcessorStatus(cpu *CPU, operand1 uint8, operand2 uint8) {
	flags := cpu.Pop()

	cpu.FromFlags(flags)
	cpu.Wait()
}
