package cpu

// LoadA A = memory
func LoadA(cpu *CPU, memory uint8) {
	cpu.A = memory

	cpu.Z = cpu.A == 0
	cpu.N = cpu.A&0x80 != 0
}

func LoadAIMM(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Immediate(cpu, operand1, operand2)
	LoadA(cpu, memory)
}

func LoadAZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPage(cpu, operand1, operand2)
	LoadA(cpu, memory)
}

func LoadAZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPageX(cpu, operand1, operand2)
	LoadA(cpu, memory)
}

func LoadAABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Absolute(cpu, operand1, operand2)
	LoadA(cpu, memory)
}

func LoadAABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteX(cpu, operand1, operand2)
	LoadA(cpu, memory)
}

func LoadAABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteY(cpu, operand1, operand2)
	LoadA(cpu, memory)
}

func LoadAINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndexedIndirect(cpu, operand1, operand2)
	LoadA(cpu, memory)
}

func LoadAINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndirectIndexed(cpu, operand1, operand2)
	LoadA(cpu, memory)
}
