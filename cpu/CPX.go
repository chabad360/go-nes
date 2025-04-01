package cpu

// CompareX X - memory
func CompareX(cpu *CPU, memory uint8) {
	result := int(cpu.X) - int(memory)

	cpu.Z = result == 0
	cpu.C = result >= 0
	cpu.N = result&0x80 != 0
}

func CompareXIMM(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Immediate(cpu, operand1, operand2)
	CompareX(cpu, memory())
}

func CompareXZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPage(cpu, operand1, operand2)
	CompareX(cpu, memory())
}

func CompareXABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Absolute(cpu, operand1, operand2)
	CompareX(cpu, memory())
}
