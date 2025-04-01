package cpu

// LoadX X = memory
func LoadX(cpu *CPU, memory uint8) {
	cpu.X = memory
	cpu.Z = cpu.X == 0
	cpu.N = cpu.X&0x80 != 0
}

func LoadXIMM(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Immediate(cpu, operand1, operand2)
	LoadX(cpu, memory)
}

func LoadXZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPage(cpu, operand1, operand2)
	LoadX(cpu, memory)
}

func LoadXZPY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPageY(cpu, operand1, operand2)
	LoadX(cpu, memory)
}

func LoadXABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Absolute(cpu, operand1, operand2)
	LoadX(cpu, memory)
}

func LoadXABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteY(cpu, operand1, operand2)
	LoadX(cpu, memory)
}
