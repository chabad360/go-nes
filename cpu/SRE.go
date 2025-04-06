package cpu

// ShiftRightAndXOR memory >>= 1; A = A ^ memory
func ShiftRightAndXOR(cpu *CPU, memory uint8, address uint16) {
	result := LogicalShiftRight(cpu, memory, address, false)
	BitwiseXOR(cpu, result)
}

func ShiftRightAndXORZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPage(cpu, operand1, operand2)
	ShiftRightAndXOR(cpu, value(), address)
}

func ShiftRightAndXORZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPageX(cpu, operand1, operand2)
	ShiftRightAndXOR(cpu, value(), address)
}

func ShiftRightAndXORABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := Absolute(cpu, operand1, operand2)
	ShiftRightAndXOR(cpu, value(), address)
}

func ShiftRightAndXORABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteX(cpu, operand1, operand2)
	ShiftRightAndXOR(cpu, value(), address)
}

func ShiftRightAndXORABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteY(cpu, operand1, operand2)
	ShiftRightAndXOR(cpu, value(), address)
}

func ShiftRightAndXORINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := IndexedIndirect(cpu, operand1, operand2)
	ShiftRightAndXOR(cpu, value(), address)
}

func ShiftRightAndXORINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := IndirectIndexed(cpu, operand1, operand2)
	ShiftRightAndXOR(cpu, value(), address)
}
