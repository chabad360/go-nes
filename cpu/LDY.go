package cpu

// LoadY Y = memory
func LoadY(cpu *CPU, memory uint8) {
	cpu.Y = memory

	cpu.Z = cpu.Y == 0
	cpu.N = cpu.Y&0x80 != 0
}

func LoadYIMM(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Immediate(cpu, operand1, operand2, false)
	LoadY(cpu, memory())
}

func LoadYZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPage(cpu, operand1, operand2, false)
	LoadY(cpu, memory())
}

func LoadYZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPageX(cpu, operand1, operand2, false)
	LoadY(cpu, memory())
}

func LoadYABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Absolute(cpu, operand1, operand2, false)
	LoadY(cpu, memory())
}

func LoadYABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteX(cpu, operand1, operand2, false)
	LoadY(cpu, memory())
}
