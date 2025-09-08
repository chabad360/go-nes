package cpu

// RotateLeftAndAnd C = A[7]; A = A << 1; A[0] = C; A = A & M
func RotateLeftAndAnd(cpu *CPU, value uint8, address uint16) {
	result := RotateLeft(cpu, value, address, false)
	BitwiseAnd(cpu, result)
}

func RotateLeftAndAndZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPage(cpu, operand1, operand2, false)
	RotateLeftAndAnd(cpu, value(), address)
}

func RotateLeftAndAndZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPageX(cpu, operand1, operand2, false)
	RotateLeftAndAnd(cpu, value(), address)
}

func RotateLeftAndAndABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := Absolute(cpu, operand1, operand2, false)
	RotateLeftAndAnd(cpu, value(), address)
}

func RotateLeftAndAndABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteX(cpu, operand1, operand2, false)
	RotateLeftAndAnd(cpu, value(), address)
}

func RotateLeftAndAndABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteY(cpu, operand1, operand2, false)
	RotateLeftAndAnd(cpu, value(), address)
}

func RotateLeftAndAndINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := IndexedIndirect(cpu, operand1, operand2, false)
	RotateLeftAndAnd(cpu, value(), address)
}

func RotateLeftAndAndINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := IndirectIndexed(cpu, operand1, operand2, false)
	RotateLeftAndAnd(cpu, value(), address)
}
