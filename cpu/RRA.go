package cpu

// RotateRightAdd C = A[0]; A = A >> 1; A[7] = C; A = A + memory
func RotateRightAdd(cpu *CPU, memory uint8, address uint16) {
	result := RotateRight(cpu, memory, address, false)
	AddWithCarry(cpu, result)
}

func RotateRightAddZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPage(cpu, operand1, operand2, false)
	RotateRightAdd(cpu, value(), address)
}

func RotateRightAddZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPageX(cpu, operand1, operand2, false)
	RotateRightAdd(cpu, value(), address)
}

func RotateRightAddABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := Absolute(cpu, operand1, operand2, false)
	RotateRightAdd(cpu, value(), address)
}

func RotateRightAddABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteX(cpu, operand1, operand2, false)
	RotateRightAdd(cpu, value(), address)
}

func RotateRightAddABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteY(cpu, operand1, operand2, false)
	RotateRightAdd(cpu, value(), address)
}

func RotateRightAddINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := IndexedIndirect(cpu, operand1, operand2, false)
	RotateRightAdd(cpu, value(), address)
}

func RotateRightAddINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := IndirectIndexed(cpu, operand1, operand2, false)
	RotateRightAdd(cpu, value(), address)
}
