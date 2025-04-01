package cpu

// Compare A A - memory
func Compare(cpu *CPU, memory uint8) {
	result := int(cpu.A) - int(memory)

	cpu.C = result >= 0
	cpu.Z = uint8(result) == 0
	cpu.N = result&0x80 != 0
}

func CompareIMM(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Immediate(cpu, operand1, operand2)
	Compare(cpu, memory())
}

func CompareZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPage(cpu, operand1, operand2)
	Compare(cpu, memory())
}

func CompareZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPageX(cpu, operand1, operand2)
	Compare(cpu, memory())
}

func CompareABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Absolute(cpu, operand1, operand2)
	Compare(cpu, memory())
}

func CompareABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteX(cpu, operand1, operand2)
	Compare(cpu, memory())
}

func CompareABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteY(cpu, operand1, operand2)
	Compare(cpu, memory())
}

func CompareINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndexedIndirect(cpu, operand1, operand2)
	Compare(cpu, memory())
}

func CompareINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndirectIndexed(cpu, operand1, operand2)
	Compare(cpu, memory())
}
