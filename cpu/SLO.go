package cpu

// ShiftLeftAndOr memory = C < memory << 1; A = A | memory
func ShiftLeftAndOr(cpu *CPU, memory uint8, address uint16) {
	result := ArithmeticShiftLeft(cpu, memory, address, false)
	BitwiseOr(cpu, result)
}

func ShiftLeftAndOrZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPage(cpu, operand1, operand2)
	ShiftLeftAndOr(cpu, value(), address)
}

func ShiftLeftAndOrZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPageX(cpu, operand1, operand2)
	ShiftLeftAndOr(cpu, value(), address)
}

func ShiftLeftAndOrABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := Absolute(cpu, operand1, operand2)
	ShiftLeftAndOr(cpu, value(), address)
}

func ShiftLeftAndOrABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteX(cpu, operand1, operand2)
	ShiftLeftAndOr(cpu, value(), address)
}

func ShiftLeftAndOrABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteY(cpu, operand1, operand2)
	ShiftLeftAndOr(cpu, value(), address)
}

func ShiftLeftAndOrINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := IndexedIndirect(cpu, operand1, operand2)
	ShiftLeftAndOr(cpu, value(), address)
}

func ShiftLeftAndOrINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := IndirectIndexed(cpu, operand1, operand2)
	ShiftLeftAndOr(cpu, value(), address)
}
