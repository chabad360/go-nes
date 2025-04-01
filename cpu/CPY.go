package cpu

// CompareY Y - memory
func CompareY(cpu *CPU, memory uint8) {
	result := int(cpu.Y) - int(memory)

	cpu.Z = result == 0
	cpu.C = result >= 0
	cpu.N = result&0x80 != 0
}

func CompareYIMM(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Immediate(cpu, operand1, operand2)
	CompareY(cpu, memory())
}

func CompareYZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPage(cpu, operand1, operand2)
	CompareY(cpu, memory())
}

func CompareYABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Absolute(cpu, operand1, operand2)
	CompareY(cpu, memory())
}
