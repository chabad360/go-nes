package cpu

import "fmt"

type Instruction struct {
	Opcode  uint8
	Size    int
	Mode    AddressingMode
	Cycles  int
	Execute ExecutionFunc
}

type ExecutionFunc func(*CPU, uint8, uint8)

var Instructions = map[uint8]Instruction{
	// ADC: Add with Carry
	0x69: {Opcode: 0x69, Size: 2, Mode: Immediate, Cycles: 2, Execute: AddWithCarryIMM},
	0x65: {Opcode: 0x65, Size: 2, Mode: ZeroPage, Cycles: 3, Execute: AddWithCarryZP},
	0x75: {Opcode: 0x75, Size: 2, Mode: ZeroPageX, Cycles: 4, Execute: AddWithCarryZPX},
	0x6D: {Opcode: 0x6D, Size: 3, Mode: Absolute, Cycles: 4, Execute: AddWithCarryABS},
	0x7D: {Opcode: 0x7D, Size: 3, Mode: AbsoluteX, Cycles: 4, Execute: AddWithCarryABSX},
	0x79: {Opcode: 0x79, Size: 3, Mode: AbsoluteY, Cycles: 4, Execute: AddWithCarryABSY},
	0x61: {Opcode: 0x61, Size: 2, Mode: IndexedIndirect, Cycles: 6, Execute: AddWithCarryINDX},
	0x71: {Opcode: 0x71, Size: 2, Mode: IndirectIndexed, Cycles: 5, Execute: AddWithCarryINDY},

	0xEA: {Opcode: 0xEA, Size: 1, Mode: Implicit, Cycles: 2, Execute: NOP},
}

// NOP is a no-operation instruction
func NOP(cpu *CPU, operand1 uint8, operand2 uint8) {
	// No operation
}

func UNKNOWN(cpu *CPU, operand1 uint8, operand2 uint8) {
	// Handle unknown opcode
	// This could be a good place to log an error or set a flag
	// For now, we'll just do nothing
	value := cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC)

	panic(fmt.Sprintf("Unknown opcode: %x at PC: %x", value, cpu.PC))
}
