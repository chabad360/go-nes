package cpu

// ClearDecimal D = 0
func ClearDecimal(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.D = false
}
