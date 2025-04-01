package cpu

import (
	"fmt"
	"log/slog"
)

type AddressingMode func(cpu *CPU, operand1, operand2 uint8) (value uint8, address uint16)

func logInstruction(cpu *CPU, mode string, operand1, operand2 uint8, address uint16, value uint8) {
	if cpu.Log {
		slog.Info(mode+" Mode Instruction",
			"PC", fmt.Sprintf("0x%04X", cpu.PC),
			"Opcode", fmt.Sprintf("0x%02X", cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC)),
			"Operand1", fmt.Sprintf("0x%02X", operand1),
			"Operand2", fmt.Sprintf("0x%02X", operand2),
			"Address", fmt.Sprintf("0x%04X", address),
			"Value", fmt.Sprintf("0x%02X", value),
		)
	}
}

// Immediate returns the immediate operand without any modification
func Immediate(cpu *CPU, operand1, operand2 uint8) (uint8, uint16) {
	logInstruction(cpu, "Immediate", operand1, operand2, 0, operand1)
	return operand1, 0
}

// Implicit is used for instructions that do not require an operand
func Implicit(cpu *CPU, operand1, operand2 uint8) (uint8, uint16) {
	// No operation, just return 0
	return 0, 0
}

// Accumulator is used for instructions that operate directly on the accumulator
func Accumulator(cpu *CPU, operand1, operand2 uint8) (uint8, uint16) {
	return cpu.A, 0 // No address, just return the accumulator value
}

// Relative calculates the address for branch instructions based on the operand
func Relative(cpu *CPU, operand1, operand2 uint8) (uint8, uint16) {
	address := uint16(int(cpu.PC) + int(int8(operand1)))
	logInstruction(cpu, "Relative", operand1, operand2, address, operand1)
	return operand1, address
}

// ZeroPage reads a byte from the zero page using the operand as the address
func ZeroPage(cpu *CPU, operand1, operand2 uint8) (value uint8, address uint16) {
	value = cpu.Mapper.Read(cpu, cpu.Clock, uint16(operand1))
	address = uint16(operand1)
	logInstruction(cpu, "Zero Page", operand1, operand2, address, value)
	return
}

// ZeroPageX reads a byte from the zero page using the operand plus the X register as the address
func ZeroPageX(cpu *CPU, operand1, operand2 uint8) (value uint8, address uint16) {
	address = uint16(operand1+cpu.X) % 256
	value = cpu.Mapper.Read(cpu, cpu.Clock, address)
	logInstruction(cpu, "Zero Page X", operand1, operand2, address, value)
	return
}

// ZeroPageY reads a byte from the zero page using the operand plus the Y register as the address
func ZeroPageY(cpu *CPU, operand1, operand2 uint8) (value uint8, address uint16) {
	address = uint16(operand1+cpu.Y) % 256
	value = cpu.Mapper.Read(cpu, cpu.Clock, address)
	logInstruction(cpu, "Zero Page Y", operand1, operand2, address, value)
	return
}

// Absolute reads a byte from the absolute 16-bit address specified by the operand
func Absolute(cpu *CPU, operand1, operand2 uint8) (value uint8, address uint16) {
	address = uint16(operand2)<<8 | uint16(operand1)
	value = cpu.Mapper.Read(cpu, cpu.Clock, address)
	logInstruction(cpu, "Absolute", operand1, operand2, address, value)
	return
}

// AbsoluteX reads a byte from the absolute address specified by the operand plus the X register
func AbsoluteX(cpu *CPU, operand1, operand2 uint8) (value uint8, address uint16) {
	address = uint16(operand2)<<8 | uint16(operand1)
	address += uint16(cpu.X)
	value = cpu.Mapper.Read(cpu, cpu.Clock, address)
	logInstruction(cpu, "Absolute X", operand1, operand2, address, value)
	return
}

// AbsoluteY reads a byte from the absolute address specified by the operand plus the Y register
func AbsoluteY(cpu *CPU, operand1, operand2 uint8) (value uint8, address uint16) {
	address = uint16(operand2)<<8 | uint16(operand1)
	address += uint16(cpu.Y)
	value = cpu.Mapper.Read(cpu, cpu.Clock, address)
	logInstruction(cpu, "Absolute Y Mode", operand1, operand2, address, value)
	return
}

// Indirect reads a byte from the address specified by the operand, which is treated as a pointer (doesn't return the address)
func Indirect(cpu *CPU, operand1, operand2 uint8) (uint8, uint16) {
	address := uint16(operand2)<<8 | uint16(operand1)

	// Read the low byte and high byte from the address
	low := cpu.Mapper.Read(cpu, cpu.Clock, address)

	hAddr := address + 1
	if address&0x00FF == 0x00FF {
		hAddr = address & 0xFF00 // wrap around for page boundary bug
	}
	high := cpu.Mapper.Read(cpu, cpu.Clock, hAddr)

	// Combine the high and low bytes to get the actual address
	address = uint16(high)<<8 | uint16(low)

	low = cpu.Mapper.Read(cpu, cpu.Clock, address)

	hAddr = address + 1
	high = cpu.Mapper.Read(cpu, cpu.Clock, hAddr)

	logInstruction(cpu, "Indirect", operand1, operand2, uint16(operand2)<<8|uint16(operand1), low)

	// Return the value at the indirect address
	return low, uint16(high)<<8 | uint16(low)
}

// IndexedIndirect reads a byte from an indirect address specified by the operand plus the X register
func IndexedIndirect(cpu *CPU, operand1, operand2 uint8) (value uint8, address uint16) {
	address = uint16(operand1 + cpu.X)
	low := cpu.Mapper.Read(cpu, cpu.Clock, address%256)

	hAddr := (address + 1) % 256
	high := cpu.Mapper.Read(cpu, cpu.Clock, hAddr)

	// Combine the high and low bytes to get the actual address
	address = uint16(high)<<8 | uint16(low)

	value = cpu.Mapper.Read(cpu, cpu.Clock, address)

	logInstruction(cpu, "Indexed Indirect", operand1, operand2, address, value)
	return
}

// IndirectIndexed reads a byte from an indirect address specified by the operand plus the Y register
func IndirectIndexed(cpu *CPU, operand1, operand2 uint8) (value uint8, address uint16) {
	address = uint16(operand1)
	low := cpu.Mapper.Read(cpu, cpu.Clock, address)

	hAddr := address + 1
	high := cpu.Mapper.Read(cpu, cpu.Clock, hAddr)

	// Combine the high and low bytes to get the actual address
	address = uint16(high)<<8 | uint16(low)

	// Add the Y register to the address
	address += uint16(cpu.Y)

	value = cpu.Mapper.Read(cpu, cpu.Clock, address)

	logInstruction(cpu, "Indirect Indexed", operand1, operand2, address, value)
	return
}
