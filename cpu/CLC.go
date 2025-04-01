package cpu

// ClearCarry C = 0
func ClearCarry(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.C = false
}
