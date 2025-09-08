package cpu

// AddWithCarry A = A + memory + C
func AddWithCarry(cpu *CPU, memory uint8) {
	result := uint16(cpu.A)
	result += uint16(memory)
	if cpu.C {
		result++
	}

	cpu.C = result > 0xFF
	cpu.Z = uint8(result) == 0
	cpu.N = result&0x80 != 0
	cpu.V = (result^uint16(cpu.A))&(result^uint16(memory))&0x80 != 0
	cpu.A = uint8(result)
}

func AddWithCarryIMM(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Immediate(cpu, operand1, operand2, false)
	AddWithCarry(cpu, memory())
}
func AddWithCarryZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPage(cpu, operand1, operand2, false)
	AddWithCarry(cpu, memory())
}
func AddWithCarryZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPageX(cpu, operand1, operand2, false)
	AddWithCarry(cpu, memory())
}

func AddWithCarryABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Absolute(cpu, operand1, operand2, false)
	AddWithCarry(cpu, memory())
}

func AddWithCarryABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteX(cpu, operand1, operand2, false)
	AddWithCarry(cpu, memory())
}

func AddWithCarryABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteY(cpu, operand1, operand2, false)
	AddWithCarry(cpu, memory())
}

func AddWithCarryINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndexedIndirect(cpu, operand1, operand2, false)
	AddWithCarry(cpu, memory())
}

func AddWithCarryINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndirectIndexed(cpu, operand1, operand2, false)
	AddWithCarry(cpu, memory())
}
