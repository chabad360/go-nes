package cpu

// Jump PC = memory
func Jump(cpu *CPU, memory uint16) {
	cpu.PC = memory
}

func JumpABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := Absolute(cpu, operand1, operand2, true)
	Jump(cpu, address)
}

func JumpIND(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := Indirect(cpu, operand1, operand2, true)
	Jump(cpu, address)
}
