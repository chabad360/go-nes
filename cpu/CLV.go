package cpu

// ClearOverflow V = 0
func ClearOverflow(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.V = false
}
