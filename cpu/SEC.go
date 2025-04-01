package cpu

// SetCarry C = 1
func SetCarry(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.C = true
}
