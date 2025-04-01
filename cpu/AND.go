package cpu

// BitwiseAnd A = A & memory
func BitwiseAnd(cpu *CPU, memory uint8) {
	cpu.A = cpu.A & memory

	cpu.Z = cpu.A == 0
	cpu.N = cpu.A&0x80 != 0
}

func AndIMM(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Immediate(cpu, operand1, operand2)
	BitwiseAnd(cpu, memory)
}

func AndZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPage(cpu, operand1, operand2)
	BitwiseAnd(cpu, memory)
}

func AndZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPageX(cpu, operand1, operand2)
	BitwiseAnd(cpu, memory)
}

func AndABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Absolute(cpu, operand1, operand2)
	BitwiseAnd(cpu, memory)
}

func AndABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteX(cpu, operand1, operand2)
	BitwiseAnd(cpu, memory)
}

func AndABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteY(cpu, operand1, operand2)
	BitwiseAnd(cpu, memory)
}

func AndINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndexedIndirect(cpu, operand1, operand2)
	BitwiseAnd(cpu, memory)
}

func AndINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndirectIndexed(cpu, operand1, operand2)
	BitwiseAnd(cpu, memory)
}
