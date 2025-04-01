package cpu

// SetInterruptDisable I = 1
func SetInterruptDisable(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.I = true
}
