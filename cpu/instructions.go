package cpu

import (
	"fmt"
	"log/slog"
)

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

	// BCC: Branch if Carry Clear
	0x90: {Opcode: 0x90, Size: 2, Mode: Relative, Cycles: 2, Execute: BranchIfCarryClear},

	// BCS: Branch if Carry Set
	0xB0: {Opcode: 0xB0, Size: 2, Mode: Relative, Cycles: 2, Execute: BranchIfCarrySet},

	// BEQ: Branch if Equal
	0xF0: {Opcode: 0xF0, Size: 2, Mode: Relative, Cycles: 2, Execute: BranchIfEqual},

	// BNE: Branch if Not Equal
	0xD0: {Opcode: 0xD0, Size: 2, Mode: Relative, Cycles: 2, Execute: BranchIfNotEqual},

	// CLC: Clear Carry Flag
	0x18: {Opcode: 0x18, Size: 1, Mode: Implicit, Cycles: 2, Execute: ClearCarry},

	// JMP: Jump
	0x4C: {Opcode: 0x4C, Size: 3, Mode: Absolute, Cycles: 3, Execute: JumpABS},
	0x6C: {Opcode: 0x6C, Size: 3, Mode: Indirect, Cycles: 5, Execute: JumpIND},

	// JSR: Jump to Subroutine
	0x20: {Opcode: 0x20, Size: 3, Mode: Absolute, Cycles: 6, Execute: JumpToSubroutineABS},

	// LDA: Load Accumulator
	0xA9: {Opcode: 0xA9, Size: 2, Mode: Immediate, Cycles: 2, Execute: LoadAIMM},
	0xA5: {Opcode: 0xA5, Size: 2, Mode: ZeroPage, Cycles: 3, Execute: LoadAZP},
	0xB5: {Opcode: 0xB5, Size: 2, Mode: ZeroPageX, Cycles: 4, Execute: LoadAZPX},
	0xAD: {Opcode: 0xAD, Size: 3, Mode: Absolute, Cycles: 4, Execute: LoadAABS},
	0xBD: {Opcode: 0xBD, Size: 3, Mode: AbsoluteX, Cycles: 4, Execute: LoadAABSX},
	0xB9: {Opcode: 0xB9, Size: 3, Mode: AbsoluteY, Cycles: 4, Execute: LoadAABSY},
	0xA1: {Opcode: 0xA1, Size: 2, Mode: IndexedIndirect, Cycles: 6, Execute: LoadAINDX},
	0xB1: {Opcode: 0xB1, Size: 2, Mode: IndirectIndexed, Cycles: 5, Execute: LoadAINDY},

	// LDX: Load X Register
	0xA2: {Opcode: 0xA2, Size: 2, Mode: Immediate, Cycles: 2, Execute: LoadXIMM},
	0xA6: {Opcode: 0xA6, Size: 2, Mode: ZeroPage, Cycles: 3, Execute: LoadXZP},
	0xB6: {Opcode: 0xB6, Size: 2, Mode: ZeroPageY, Cycles: 4, Execute: LoadXZPY},
	0xAE: {Opcode: 0xAE, Size: 3, Mode: Absolute, Cycles: 4, Execute: LoadXABS},
	0xBE: {Opcode: 0xBE, Size: 3, Mode: AbsoluteY, Cycles: 4, Execute: LoadXABSY},

	// SEC: Set Carry Flag
	0x38: {Opcode: 0x38, Size: 1, Mode: Implicit, Cycles: 2, Execute: SetCarry},

	// STX: Store X Register
	0x86: {Opcode: 0x86, Size: 2, Mode: ZeroPage, Cycles: 3, Execute: StoreXZP},
	0x96: {Opcode: 0x96, Size: 2, Mode: ZeroPageY, Cycles: 4, Execute: StoreXZPY},
	0x8E: {Opcode: 0x8E, Size: 3, Mode: Absolute, Cycles: 4, Execute: StoreXABS},

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

// StubImmediate is a placeholder for immediate addressing mode
func StubImmediate(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Stubbed = true
	slog.Warn("Immediate Mode Instruction Stubbed",
		"PC", cpu.PC,
		"Operand1", operand1,
		"Operand2", "None",
		"Opcode", cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC),
		"Address", "None",
		"Value", operand1,
	)
}

// StubImplicit is a placeholder for implicit addressing mode
func StubImplicit(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Stubbed = true
	slog.Warn("Implicit Mode Instruction Stubbed",
		"PC", cpu.PC,
		"Operand1", "None",
		"Operand2", "None",
		"Opcode", cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC),
		"Address", "None",
		"Value", "None",
	)
}

// StubAccumulator is a placeholder for accumulator addressing mode
func StubAccumulator(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Stubbed = true
	slog.Warn("Accumulator Mode Instruction Stubbed",
		"PC", cpu.PC,
		"Operand1", "None",
		"Operand2", "None",
		"Opcode", cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC),
		"Address", "None",
		"Value", cpu.A,
	)
}

// StubRelative is a placeholder for relative addressing mode
func StubRelative(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Stubbed = true
	value, address := Relative(cpu, operand1, operand2)
	slog.Warn("Relative Mode Instruction Stubbed",
		"PC", cpu.PC,
		"Operand1", operand1,
		"Operand2", "None",
		"Opcode", cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC),
		"Address", address,
		"Value", value,
	)
}

// StubZeroPage is a placeholder for zero page addressing mode operation
func StubZeroPage(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Stubbed = true
	value, address := ZeroPage(cpu, operand1, operand2)
	slog.Warn("Zero Page Mode Instruction Stubbed",
		"PC", cpu.PC,
		"Operand1", operand1,
		"Operand2", "None",
		"Opcode", cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC),
		"Address", address,
		"Value", value,
	)
}

// StubZeroPageX is a placeholder for zero page X addressing mode
func StubZeroPageX(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Stubbed = true
	value, address := ZeroPageX(cpu, operand1, operand2)
	slog.Warn("Zero Page X Mode Instruction Stubbed",
		"PC", cpu.PC,
		"Operand1", operand1,
		"Operand2", "None",
		"Opcode", cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC),
		"Address", address,
		"Value", value,
	)
}

// StubZeroPageY is a placeholder for zero page Y addressing mode operation
func StubZeroPageY(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Stubbed = true
	value, address := ZeroPageY(cpu, operand1, operand2)
	slog.Warn("Zero Page Y Mode Instruction Stubbed",
		"PC", cpu.PC,
		"Operand1", operand1,
		"Operand2", "None",
		"Opcode", cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC),
		"Address", address,
		"Value", value,
	)
}

// StubAbsolute is a placeholder for absolute addressing mode operation
func StubAbsolute(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Stubbed = true
	value, address := Absolute(cpu, operand1, operand2)
	slog.Warn("Absolute Mode Instruction Stubbed",
		"PC", cpu.PC,
		"Operand1", operand1,
		"Operand2", operand2,
		"Opcode", cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC),
		"Address", address,
		"Value", value,
	)
}

// StubAbsoluteX is a placeholder for absolute X addressing mode operation
func StubAbsoluteX(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Stubbed = true
	value, address := AbsoluteX(cpu, operand1, operand2)
	slog.Warn("Absolute X Mode Instruction Stubbed",
		"PC", cpu.PC,
		"Operand1", operand1,
		"Operand2", operand2,
		"Opcode", cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC),
		"Address", address,
		"Value", value,
	)
}

// StubAbsoluteY is a placeholder for absolute Y addressing mode operation
func StubAbsoluteY(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Stubbed = true
	value, address := AbsoluteY(cpu, operand1, operand2)
	slog.Warn("Absolute Y Mode Instruction Stubbed",
		"PC", cpu.PC,
		"Operand1", operand1,
		"Operand2", operand2,
		"Opcode", cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC),
		"Address", address,
		"Value", value,
	)
}

// StubIndirect is a placeholder for indirect addressing mode operation
func StubIndirect(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Stubbed = true
	_, value2 := Indirect(cpu, operand1, operand2)
	slog.Warn("Indirect Mode Instruction Stubbed",
		"PC", cpu.PC,
		"Operand1", operand1,
		"Operand2", operand2,
		"Opcode", cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC),
		"Address", uint16(operand2)<<8|uint16(operand1),
		"Value", value2,
	)
}

// StubIndexedIndirect is a placeholder for indexed indirect addressing mode
func StubIndexedIndirect(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Stubbed = true
	value, address := IndexedIndirect(cpu, operand1, operand2)
	slog.Warn("Indexed Indirect Mode Instruction Stubbed",
		"PC", cpu.PC,
		"Operand1", operand1,
		"Operand2", "None",
		"Opcode", cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC),
		"Address", address,
		"Value", value,
	)
}

// StubIndirectIndexed is a placeholder for an indirect indexed addressing mode operation
func StubIndirectIndexed(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Stubbed = true
	value, address := IndirectIndexed(cpu, operand1, operand2)
	slog.Warn("Indirect Indexed Mode Instruction Stubbed",
		"PC", cpu.PC,
		"Operand1", operand1,
		"Operand2", "None",
		"Opcode", cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC),
		"Address", address,
		"Value", value,
	)
}
