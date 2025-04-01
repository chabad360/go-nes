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

	// AND: Logical AND
	0x29: {Opcode: 0x29, Size: 2, Mode: Immediate, Cycles: 2, Execute: AndIMM},
	0x25: {Opcode: 0x25, Size: 2, Mode: ZeroPage, Cycles: 3, Execute: AndZP},
	0x35: {Opcode: 0x35, Size: 2, Mode: ZeroPageX, Cycles: 4, Execute: AndZPX},
	0x2D: {Opcode: 0x2D, Size: 3, Mode: Absolute, Cycles: 4, Execute: AndABS},
	0x3D: {Opcode: 0x3D, Size: 3, Mode: AbsoluteX, Cycles: 4, Execute: AndABSX},
	0x39: {Opcode: 0x39, Size: 3, Mode: AbsoluteY, Cycles: 4, Execute: AndABSY},
	0x21: {Opcode: 0x21, Size: 2, Mode: IndexedIndirect, Cycles: 6, Execute: AndINDX},
	0x31: {Opcode: 0x31, Size: 2, Mode: IndirectIndexed, Cycles: 5, Execute: AndINDY},

	// BCC: Branch if Carry Clear
	0x90: {Opcode: 0x90, Size: 2, Mode: Relative, Cycles: 2, Execute: BranchIfCarryClear},

	// BCS: Branch if Carry Set
	0xB0: {Opcode: 0xB0, Size: 2, Mode: Relative, Cycles: 2, Execute: BranchIfCarrySet},

	// BEQ: Branch if Equal
	0xF0: {Opcode: 0xF0, Size: 2, Mode: Relative, Cycles: 2, Execute: BranchIfEqual},

	// BIT: Bit Test
	0x24: {Opcode: 0x24, Size: 2, Mode: ZeroPage, Cycles: 3, Execute: BitTestZP},
	0x2C: {Opcode: 0x2C, Size: 3, Mode: Absolute, Cycles: 4, Execute: BitTestABS},

	// BMI: Branch if Minus
	0x30: {Opcode: 0x30, Size: 2, Mode: Relative, Cycles: 2, Execute: BranchIfMinus},

	// BNE: Branch if Not Equal
	0xD0: {Opcode: 0xD0, Size: 2, Mode: Relative, Cycles: 2, Execute: BranchIfNotEqual},

	// BPL: Branch if Positive
	0x10: {Opcode: 0x10, Size: 2, Mode: Relative, Cycles: 2, Execute: BranchIfPlus},

	// BVC: Branch if Overflow Clear
	0x50: {Opcode: 0x50, Size: 2, Mode: Relative, Cycles: 2, Execute: BranchIfOverflowClear},

	// BVS: Branch if Overflow Set
	0x70: {Opcode: 0x70, Size: 2, Mode: Relative, Cycles: 2, Execute: BranchIfOverflowSet},

	// CLC: Clear Carry Flag
	0x18: {Opcode: 0x18, Size: 1, Mode: Implicit, Cycles: 2, Execute: ClearCarry},

	// CLD: Clear Decimal Flag
	0xD8: {Opcode: 0xD8, Size: 1, Mode: Implicit, Cycles: 2, Execute: ClearDecimal},

	// CLV: Clear Overflow Flag
	0xB8: {Opcode: 0xB8, Size: 1, Mode: Implicit, Cycles: 2, Execute: ClearOverflow},

	// CMP: Compare
	0xC9: {Opcode: 0xC9, Size: 2, Mode: Immediate, Cycles: 2, Execute: CompareIMM},
	0xC5: {Opcode: 0xC5, Size: 2, Mode: ZeroPage, Cycles: 3, Execute: CompareZP},
	0xD5: {Opcode: 0xD5, Size: 2, Mode: ZeroPageX, Cycles: 4, Execute: CompareZPX},
	0xCD: {Opcode: 0xCD, Size: 3, Mode: Absolute, Cycles: 4, Execute: CompareABS},
	0xDD: {Opcode: 0xDD, Size: 3, Mode: AbsoluteX, Cycles: 4, Execute: CompareABSX},
	0xD9: {Opcode: 0xD9, Size: 3, Mode: AbsoluteY, Cycles: 4, Execute: CompareABSY},
	0xC1: {Opcode: 0xC1, Size: 2, Mode: IndexedIndirect, Cycles: 6, Execute: CompareINDX},
	0xD1: {Opcode: 0xD1, Size: 2, Mode: IndirectIndexed, Cycles: 5, Execute: CompareINDY},

	// CPX: Compare X Register
	0xE0: {Opcode: 0xE0, Size: 2, Mode: Immediate, Cycles: 2, Execute: CompareXIMM},
	0xE4: {Opcode: 0xE4, Size: 2, Mode: ZeroPage, Cycles: 3, Execute: CompareXZP},
	0xEC: {Opcode: 0xEC, Size: 3, Mode: Absolute, Cycles: 4, Execute: CompareXABS},

	// CPY: Compare Y Register
	0xC0: {Opcode: 0xC0, Size: 2, Mode: Immediate, Cycles: 2, Execute: CompareYIMM},
	0xC4: {Opcode: 0xC4, Size: 2, Mode: ZeroPage, Cycles: 3, Execute: CompareYZP},
	0xCC: {Opcode: 0xCC, Size: 3, Mode: Absolute, Cycles: 4, Execute: CompareYABS},

	// EOR: Exclusive OR
	0x49: {Opcode: 0x49, Size: 2, Mode: Immediate, Cycles: 2, Execute: BitwiseXORIMM},
	0x45: {Opcode: 0x45, Size: 2, Mode: ZeroPage, Cycles: 3, Execute: BitwiseXORZP},
	0x55: {Opcode: 0x55, Size: 2, Mode: ZeroPageX, Cycles: 4, Execute: BitwiseXORZPX},
	0x4D: {Opcode: 0x4D, Size: 3, Mode: Absolute, Cycles: 4, Execute: BitwiseXORABS},
	0x5D: {Opcode: 0x5D, Size: 3, Mode: AbsoluteX, Cycles: 4, Execute: BitwiseXORABSX},
	0x59: {Opcode: 0x59, Size: 3, Mode: AbsoluteY, Cycles: 4, Execute: BitwiseXORABSY},
	0x41: {Opcode: 0x41, Size: 2, Mode: IndexedIndirect, Cycles: 6, Execute: BitwiseXORINDX},
	0x51: {Opcode: 0x51, Size: 2, Mode: IndirectIndexed, Cycles: 5, Execute: BitwiseXORINDY},

	// JMP: Jump
	0x4C: {Opcode: 0x4C, Size: 3, Mode: Absolute, Cycles: 3, Execute: JumpABS},
	0x6C: {Opcode: 0x6C, Size: 3, Mode: Indirect, Cycles: 5, Execute: JumpIND},

	// JSR: Jump to Subroutine
	0x20: {Opcode: 0x20, Size: 3, Mode: Absolute, Cycles: 6, Execute: JumpToSubroutine},

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

	// LDY: Load Y Register
	0xA0: {Opcode: 0xA0, Size: 2, Mode: Immediate, Cycles: 2, Execute: LoadYIMM},
	0xA4: {Opcode: 0xA4, Size: 2, Mode: ZeroPage, Cycles: 3, Execute: LoadYZP},
	0xB4: {Opcode: 0xB4, Size: 2, Mode: ZeroPageX, Cycles: 4, Execute: LoadYZPX},
	0xAC: {Opcode: 0xAC, Size: 3, Mode: Absolute, Cycles: 4, Execute: LoadYABS},
	0xBC: {Opcode: 0xBC, Size: 3, Mode: AbsoluteX, Cycles: 4, Execute: LoadYABSX},

	// ORA: Logical OR
	0x09: {Opcode: 0x09, Size: 2, Mode: Immediate, Cycles: 2, Execute: BitwiseOrIMM},
	0x05: {Opcode: 0x05, Size: 2, Mode: ZeroPage, Cycles: 3, Execute: BitwiseOrZP},
	0x15: {Opcode: 0x15, Size: 2, Mode: ZeroPageX, Cycles: 4, Execute: BitwiseOrZPX},
	0x0D: {Opcode: 0x0D, Size: 3, Mode: Absolute, Cycles: 4, Execute: BitwiseOrABS},
	0x1D: {Opcode: 0x1D, Size: 3, Mode: AbsoluteX, Cycles: 4, Execute: BitwiseOrABSX},
	0x19: {Opcode: 0x19, Size: 3, Mode: AbsoluteY, Cycles: 4, Execute: BitwiseOrABSY},
	0x01: {Opcode: 0x01, Size: 2, Mode: IndexedIndirect, Cycles: 6, Execute: BitwiseOrINDX},
	0x11: {Opcode: 0x11, Size: 2, Mode: IndirectIndexed, Cycles: 5, Execute: BitwiseOrINDY},

	// PHA: Push Accumulator
	0x48: {Opcode: 0x48, Size: 1, Mode: Implicit, Cycles: 3, Execute: PushA},

	// PHP: Push Processor Status
	0x08: {Opcode: 0x08, Size: 1, Mode: Implicit, Cycles: 3, Execute: PushProcessorStatus},

	// PLA: Pull Accumulator
	0x68: {Opcode: 0x68, Size: 1, Mode: Implicit, Cycles: 4, Execute: PullA},

	// PLP: Pull Processor Status
	0x28: {Opcode: 0x28, Size: 1, Mode: Implicit, Cycles: 4, Execute: PullProcessorStatus},

	// RTS: Return from Subroutine
	0x60: {Opcode: 0x60, Size: 1, Mode: Implicit, Cycles: 6, Execute: ReturnFromSubroutine},

	// SEC: Set Carry Flag
	0x38: {Opcode: 0x38, Size: 1, Mode: Implicit, Cycles: 2, Execute: SetCarry},

	// SED: Set Decimal Flag
	0xF8: {Opcode: 0xF8, Size: 1, Mode: Implicit, Cycles: 2, Execute: SetDecimal},

	// SEI: Set Interrupt Disable
	0x78: {Opcode: 0x78, Size: 1, Mode: Implicit, Cycles: 2, Execute: SetInterruptDisable},

	// STA: Store Accumulator
	0x85: {Opcode: 0x85, Size: 2, Mode: ZeroPage, Cycles: 3, Execute: StoreAZP},
	0x95: {Opcode: 0x95, Size: 2, Mode: ZeroPageX, Cycles: 4, Execute: StoreAZPX},
	0x8D: {Opcode: 0x8D, Size: 3, Mode: Absolute, Cycles: 4, Execute: StoreAABS},
	0x9D: {Opcode: 0x9D, Size: 3, Mode: AbsoluteX, Cycles: 5, Execute: StoreAABSX},
	0x99: {Opcode: 0x99, Size: 3, Mode: AbsoluteY, Cycles: 5, Execute: StoreAABSY},
	0x81: {Opcode: 0x81, Size: 2, Mode: IndexedIndirect, Cycles: 6, Execute: StoreAINDX},
	0x91: {Opcode: 0x91, Size: 2, Mode: IndirectIndexed, Cycles: 6, Execute: StoreAINDY},

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
