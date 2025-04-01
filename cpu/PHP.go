package cpu

// PushProcessorStatus stack[SP] = N | V | 0x20 | 0x10 | D | I | Z | C
func PushProcessorStatus(cpu *CPU, operand1 uint8, operand2 uint8) {
	flags := cpu.Flags()
	cpu.Push(flags)
}
