package cpu

// BitwiseXOR A = A ^ memory
func BitwiseXOR(cpu *CPU, memory uint8) {
	cpu.A = cpu.A ^ memory

	cpu.Z = cpu.A == 0
	cpu.N = cpu.A&0x80 != 0
}

func BitwiseXORIMM(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Immediate(cpu, operand1, operand2)
	BitwiseXOR(cpu, memory())
}

func BitwiseXORZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPage(cpu, operand1, operand2)
	BitwiseXOR(cpu, memory())
}

func BitwiseXORZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPageX(cpu, operand1, operand2)
	BitwiseXOR(cpu, memory())
}

func BitwiseXORABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Absolute(cpu, operand1, operand2)
	BitwiseXOR(cpu, memory())
}

func BitwiseXORABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteX(cpu, operand1, operand2)
	BitwiseXOR(cpu, memory())
}

func BitwiseXORABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteY(cpu, operand1, operand2)
	BitwiseXOR(cpu, memory())
}

func BitwiseXORINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndexedIndirect(cpu, operand1, operand2)
	BitwiseXOR(cpu, memory())
}

func BitwiseXORINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndirectIndexed(cpu, operand1, operand2)
	BitwiseXOR(cpu, memory())
}
