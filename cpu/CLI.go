package cpu

// ClearInterrupt I = false
func ClearInterrupt(cpu *CPU, operand1, operand2 uint8) {
	cpu.I = false
}
