package cpu

// ReturnFromInterrupt N | V | B | D | I | Z | C = stack[SP]; PC = stack[SP] + 1
func ReturnFromInterrupt(cpu *CPU, operand1 uint8, operand2 uint8) {
	flags := cpu.Pop()
	cpu.FromFlags(flags)

	low := cpu.Pop()
	high := cpu.Pop()
	cpu.PC = uint16(low) | (uint16(high) << 8)
	cpu.Wait()
}
