package cpu

// BitwiseOr A = A | memory
func BitwiseOr(cpu *CPU, memory uint8) {
	cpu.A |= memory

	cpu.Z = cpu.A == 0
	cpu.N = cpu.A&0x80 != 0
}

func BitwiseOrIMM(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Immediate(cpu, operand1, operand2)
	BitwiseOr(cpu, memory)
}

func BitwiseOrZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPage(cpu, operand1, operand2)
	BitwiseOr(cpu, memory)
}

func BitwiseOrZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPageX(cpu, operand1, operand2)
	BitwiseOr(cpu, memory)
}

func BitwiseOrABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Absolute(cpu, operand1, operand2)
	BitwiseOr(cpu, memory)
}

func BitwiseOrABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteX(cpu, operand1, operand2)
	BitwiseOr(cpu, memory)
}

func BitwiseOrABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteY(cpu, operand1, operand2)
	BitwiseOr(cpu, memory)
}

func BitwiseOrINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndexedIndirect(cpu, operand1, operand2)
	BitwiseOr(cpu, memory)
}

func BitwiseOrINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndirectIndexed(cpu, operand1, operand2)
	BitwiseOr(cpu, memory)
}
