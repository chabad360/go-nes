package cpu

import (
	"fmt"
	"log/slog"
)

type AddressingMode func(cpu *CPU, operand1, operand2 uint8) (value ValueGuard, address uint16)

// ValueGuard is a closure used to wrap the value of memory, to only cost an instruction when the value is used
type ValueGuard func() uint8

func GuardValue(cpu *CPU, value uint8) ValueGuard {
	return func() uint8 {
		cpu.Wait()
		return value
	}
}

func UnguardedValue(value uint8) ValueGuard {
	return func() uint8 {
		return value
	}
}

func logInstruction(cpu *CPU, mode string, operand1, operand2 uint8, address uint16, value uint8, size int) {
	if cpu.Log {
		slog.Info(mode+" Mode Instruction",
			"PC", fmt.Sprintf("0x%04X", cpu.PC-uint16(size)),
			"Opcode", fmt.Sprintf("0x%02X", cpu.Mapper.Read(cpu, cpu.PC-uint16(size))),
			"Operand1", fmt.Sprintf("0x%02X", operand1),
			"Operand2", fmt.Sprintf("0x%02X", operand2),
			"Address", fmt.Sprintf("0x%04X", address),
			"Value", fmt.Sprintf("0x%02X", value),
		)
	}
}

// Immediate returns the immediate operand without any modification
func Immediate(cpu *CPU, operand1, operand2 uint8) (ValueGuard, uint16) {
	logInstruction(cpu, "Immediate", operand1, operand2, 0, operand1, 2)
	return UnguardedValue(operand1), 0
}

// Implicit is used for instructions that do not require an operand
func Implicit(cpu *CPU, operand1, operand2 uint8) (ValueGuard, uint16) {
	// No operation, just return 0
	return UnguardedValue(0), 0
}

// Accumulator is used for instructions that operate directly on the accumulator
func Accumulator(cpu *CPU, operand1, operand2 uint8) (ValueGuard, uint16) {
	return UnguardedValue(cpu.A), 0 // No address, just return the accumulator value
}

// Relative calculates the address for branch instructions based on the operand
func Relative(cpu *CPU, operand1, operand2 uint8) (ValueGuard, uint16) {
	address := uint16(int(cpu.PC) + int(int8(operand1)))
	logInstruction(cpu, "Relative", operand1, operand2, address, operand1, 3)
	return UnguardedValue(operand1), address
}

// ZeroPage reads a byte from the zero page using the operand as the address
func ZeroPage(cpu *CPU, operand1, operand2 uint8) (value ValueGuard, address uint16) {
	val := cpu.Mapper.Read(cpu, uint16(operand1))
	address = uint16(operand1)
	logInstruction(cpu, "Zero Page", operand1, operand2, address, val, 2)
	value = GuardValue(cpu, val)
	return
}

// ZeroPageX reads a byte from the zero page using the operand plus the X register as the address
func ZeroPageX(cpu *CPU, operand1, operand2 uint8) (value ValueGuard, address uint16) {
	address = uint16(operand1+cpu.X) % 256
	val := cpu.Mapper.Read(cpu, address)
	logInstruction(cpu, "Zero Page X", operand1, operand2, address, val, 2)
	value = GuardValue(cpu, val)
	return
}

// ZeroPageY reads a byte from the zero page using the operand plus the Y register as the address
func ZeroPageY(cpu *CPU, operand1, operand2 uint8) (value ValueGuard, address uint16) {
	address = uint16(operand1+cpu.Y) % 256
	val := cpu.Mapper.Read(cpu, address)
	logInstruction(cpu, "Zero Page Y", operand1, operand2, address, val, 2)
	value = GuardValue(cpu, val)
	return
}

// Absolute reads a byte from the absolute 16-bit address specified by the operand
func Absolute(cpu *CPU, operand1, operand2 uint8) (value ValueGuard, address uint16) {
	address = uint16(operand2)<<8 | uint16(operand1)
	val := cpu.Mapper.Read(cpu, address)
	logInstruction(cpu, "Absolute", operand1, operand2, address, val, 3)
	value = GuardValue(cpu, val)
	return
}

// AbsoluteX reads a byte from the absolute address specified by the operand plus the X register
func AbsoluteX(cpu *CPU, operand1, operand2 uint8) (value ValueGuard, address uint16) {
	address = uint16(operand2)<<8 | uint16(operand1)
	address += uint16(cpu.X)
	val := cpu.Mapper.Read(cpu, address)
	logInstruction(cpu, "Absolute X", operand1, operand2, address, val, 3)
	value = GuardValue(cpu, val)
	return
}

// AbsoluteY reads a byte from the absolute address specified by the operand plus the Y register
func AbsoluteY(cpu *CPU, operand1, operand2 uint8) (value ValueGuard, address uint16) {
	address = uint16(operand2)<<8 | uint16(operand1)
	address += uint16(cpu.Y)
	val := cpu.Mapper.Read(cpu, address)
	logInstruction(cpu, "Absolute Y Mode", operand1, operand2, address, val, 3)
	value = GuardValue(cpu, val)
	return
}

// Indirect reads a byte from the address specified by the operand, which is treated as a pointer (doesn't return the address)
func Indirect(cpu *CPU, operand1, operand2 uint8) (ValueGuard, uint16) {
	address := uint16(operand2)<<8 | uint16(operand1)

	// Read the low byte and high byte from the address
	low := cpu.Mapper.Read(cpu, address)

	hAddr := address + 1
	if address&0x00FF == 0x00FF {
		hAddr = address & 0xFF00 // wrap around for page boundary bug
	}
	high := cpu.Mapper.Read(cpu, hAddr)

	// Combine the high and low bytes to get the actual address
	address = uint16(high)<<8 | uint16(low)

	low = cpu.Mapper.Read(cpu, address)

	hAddr = address + 1
	high = cpu.Mapper.Read(cpu, hAddr)

	logInstruction(cpu, "Indirect", operand1, operand2, uint16(operand2)<<8|uint16(operand1), low, 3)

	// Return the value at the indirect address
	return UnguardedValue(low), uint16(high)<<8 | uint16(low)
}

// IndexedIndirect reads a byte from an indirect address specified by the operand plus the X register
func IndexedIndirect(cpu *CPU, operand1, operand2 uint8) (value ValueGuard, address uint16) {
	address = uint16(operand1 + cpu.X)
	low := cpu.Mapper.Read(cpu, address%256)

	hAddr := (address + 1) % 256
	high := cpu.Mapper.Read(cpu, hAddr)

	// Combine the high and low bytes to get the actual address
	address = uint16(high)<<8 | uint16(low)

	val := cpu.Mapper.Read(cpu, address)

	logInstruction(cpu, "Indexed Indirect", operand1, operand2, address, val, 3)
	value = GuardValue(cpu, val)
	return
}

// IndirectIndexed reads a byte from an indirect address specified by the operand plus the Y register
func IndirectIndexed(cpu *CPU, operand1, operand2 uint8) (value ValueGuard, address uint16) {
	address = uint16(operand1)
	low := cpu.Mapper.Read(cpu, address)

	hAddr := address + 1
	high := cpu.Mapper.Read(cpu, hAddr)

	// Combine the high and low bytes to get the actual address
	address = uint16(high)<<8 | uint16(low)

	// Add the Y register to the address
	address += uint16(cpu.Y)

	val := cpu.Mapper.Read(cpu, address)

	logInstruction(cpu, "Indirect Indexed", operand1, operand2, address, val, 3)
	value = GuardValue(cpu, val)
	return
}
