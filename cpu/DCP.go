package cpu

// DecrementCompare memory--; A - memory
func DecrementCompare(cpu *CPU, memory uint8, address uint16) {
	result := Decrement(cpu, memory, address)
	Compare(cpu, result)
}

func DecrementCompareZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPage(cpu, operand1, operand2)
	DecrementCompare(cpu, value(), address)
}

func DecrementCompareZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPageX(cpu, operand1, operand2)
	DecrementCompare(cpu, value(), address)
}

func DecrementCompareABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := Absolute(cpu, operand1, operand2)
	DecrementCompare(cpu, value(), address)
}

func DecrementCompareABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteX(cpu, operand1, operand2)
	DecrementCompare(cpu, value(), address)
}

func DecrementCompareABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteY(cpu, operand1, operand2)
	DecrementCompare(cpu, value(), address)
}

func DecrementCompareINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := IndexedIndirect(cpu, operand1, operand2)
	DecrementCompare(cpu, value(), address)
}

func DecrementCompareINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := IndirectIndexed(cpu, operand1, operand2)
	DecrementCompare(cpu, value(), address)
}
