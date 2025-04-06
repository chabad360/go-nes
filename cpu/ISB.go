package cpu

// IncrementSubtract memory++; A = A - memory
func IncrementSubtract(cpu *CPU, memory uint8, address uint16) {
	result := Increment(cpu, memory, address)
	SubtractWithCarry(cpu, result)
}

func IncrementSubtractZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPage(cpu, operand1, operand2)
	IncrementSubtract(cpu, value(), address)
}

func IncrementSubtractZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPageX(cpu, operand1, operand2)
	IncrementSubtract(cpu, value(), address)
}

func IncrementSubtractABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := Absolute(cpu, operand1, operand2)
	IncrementSubtract(cpu, value(), address)
}

func IncrementSubtractABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteX(cpu, operand1, operand2)
	IncrementSubtract(cpu, value(), address)
}

func IncrementSubtractABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteY(cpu, operand1, operand2)
	IncrementSubtract(cpu, value(), address)
}

func IncrementSubtractINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := IndexedIndirect(cpu, operand1, operand2)
	IncrementSubtract(cpu, value(), address)
}

func IncrementSubtractINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := IndirectIndexed(cpu, operand1, operand2)
	IncrementSubtract(cpu, value(), address)
}
