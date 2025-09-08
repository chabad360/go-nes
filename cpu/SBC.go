package cpu

// SubtractWithCarry A - memory - ~ C
func SubtractWithCarry(cpu *CPU, memory uint8) {
	result := uint16(cpu.A)
	result -= uint16(memory)
	if !cpu.C {
		result--
	}

	cpu.C = result&0x100 == 0
	cpu.Z = uint8(result) == 0
	cpu.N = result&0x80 != 0
	cpu.V = ((cpu.A^memory)&0x80 != 0) && ((cpu.A^uint8(result))&0x80 != 0)
	cpu.A = uint8(result)
}

func SubtractWithCarryIMM(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Immediate(cpu, operand1, operand2, false)
	SubtractWithCarry(cpu, memory())
}

func SubtractWithCarryZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPage(cpu, operand1, operand2, false)
	SubtractWithCarry(cpu, memory())
}

func SubtractWithCarryZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPageX(cpu, operand1, operand2, false)
	SubtractWithCarry(cpu, memory())
}

func SubtractWithCarryABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Absolute(cpu, operand1, operand2, false)
	SubtractWithCarry(cpu, memory())
}

func SubtractWithCarryABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteX(cpu, operand1, operand2, false)
	SubtractWithCarry(cpu, memory())
}

func SubtractWithCarryABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteY(cpu, operand1, operand2, false)
	SubtractWithCarry(cpu, memory())
}

func SubtractWithCarryINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndexedIndirect(cpu, operand1, operand2, false)
	SubtractWithCarry(cpu, memory())
}

func SubtractWithCarryINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndirectIndexed(cpu, operand1, operand2, false)
	SubtractWithCarry(cpu, memory())
}
