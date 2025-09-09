package cpu

import (
	"fmt"
	"go-nes/clock"
	"log/slog"
)

// CPU represents the state of the NES 6502 CPU.
type CPU struct {
	A, X, Y, SP         uint8
	PC                  uint16
	C, Z, I, D, B, V, N bool
	Memory              [0x0800]uint8

	Mapper Mapper

	Clock  clock.Clock
	Cycles uint32

	Log     bool
	Stubbed bool
}

// SetPowerOnState initializes the CPU to a powered-on state.
func (cpu *CPU) SetPowerOnState() {
	cpu.A = 0
	cpu.X = 0
	cpu.Y = 0
	cpu.SP = 0xFD
	cpu.PC = 0xFFFC
	cpu.C = false
	cpu.Z = false
	cpu.I = true
	cpu.D = false
	cpu.B = false
	cpu.V = false
	cpu.N = false

	cpu.Memory = [0x0800]uint8{}
}

func (cpu *CPU) PowerOn() {
	cpu.SetPowerOnState()
	cpu.Wait()
	cpu.Wait()
	cpu.Wait()
	cpu.Wait()
	cpu.Wait()
	cpu.Wait()
	cpu.Wait()
	JumpABS(cpu,
		cpu.Mapper.Read(cpu, cpu.PC),
		cpu.Mapper.Read(cpu, cpu.PC+1))
}

// Reset initializes the CPU to the reset state.
func (cpu *CPU) Reset() {
	cpu.PC = 0xFFFC
	cpu.SP = cpu.SP - 3
	cpu.I = true
}

// Decode fetches the next opcode from memory and returns the corresponding instruction.
func (cpu *CPU) Decode() (Instruction, uint8, uint16) {
	opcode := cpu.Mapper.Read(cpu, cpu.PC)
	cpu.Wait()
	pc := cpu.PC

	cpu.Wait() // Fake fetch (in theory, a cycle accurate system should also load and store the next byte here)
	cpu.PC++

	return Instructions[opcode], opcode, pc
}

// Fetch retrieves the operands for the instruction based on its size.
func (cpu *CPU) Fetch(instr Instruction) (operand1 uint8, operand2 uint8) {
	if instr.Size >= 2 {
		operand1 = cpu.Mapper.Read(cpu, cpu.PC)
		cpu.PC++
		// We don't need to wait, because the 6502 always prefetches the next byte
	}
	if instr.Size == 3 {
		operand2 = cpu.Mapper.Read(cpu, cpu.PC)
		cpu.PC++
		cpu.Wait()
	}
	return operand1, operand2
}

// Execute runs the instruction with the provided operands.
func (cpu *CPU) Execute(instr Instruction, operand1 uint8, operand2 uint8) {
	if cpu.Log {
		slog.Info("CPU State Pre-Operation",
			"A", fmt.Sprintf("0x%02X", cpu.A),
			"X", fmt.Sprintf("0x%02X", cpu.X),
			"Y", fmt.Sprintf("0x%02X", cpu.Y),
			"SP", fmt.Sprintf("0x%02X", cpu.SP),
			"PC", fmt.Sprintf("0x%04X", cpu.PC),
			"C", cpu.C,
			"Z", cpu.Z,
			"I", cpu.I,
			"D", cpu.D,
			"B", cpu.B,
			"V", cpu.V,
			"N", cpu.N,
			"Flags", fmt.Sprintf("0x%02X", cpu.Flags()),
		)
	}
	cpu.Stubbed = false

	instr.Execute(cpu, operand1, operand2)
	if cpu.Log {
		slog.Info("CPU State Post-Operation",
			"A", fmt.Sprintf("0x%02X", cpu.A),
			"X", fmt.Sprintf("0x%02X", cpu.X),
			"Y", fmt.Sprintf("0x%02X", cpu.Y),
			"SP", fmt.Sprintf("0x%02X", cpu.SP),
			"PC", fmt.Sprintf("0x%04X", cpu.PC),
			"C", cpu.C,
			"Z", cpu.Z,
			"I", cpu.I,
			"D", cpu.D,
			"B", cpu.B,
			"V", cpu.V,
			"N", cpu.N,
			"Flags", fmt.Sprintf("0x%02X", cpu.Flags()),
		)
	}
}

// Push pushes a value onto the stack.
func (cpu *CPU) Push(value uint8) {
	cpu.Mapper.Write(cpu, 0x0100+uint16(cpu.SP), value)
	cpu.SP--
	cpu.Wait()
}

// Pop pops a value from the stack.
func (cpu *CPU) Pop() uint8 {
	cpu.SP++
	value := cpu.Mapper.Read(cpu, 0x0100+uint16(cpu.SP))
	cpu.Wait()
	return value
}

// Wait waits for the next clock cycle.
func (cpu *CPU) Wait() {
	div := cpu.Clock.Standard().CPUDivider()
	cpu.Cycles = cpu.Clock.Wait((cpu.Cycles+1)*div) / div
}

func (c *CPU) Step() {
	instr, _, _ := c.Decode()
	op1, op2 := c.Fetch(instr)
	c.Execute(instr, op1, op2)
}

// Flags returns the current status flags as a byte.
func (cpu *CPU) Flags() uint8 {
	flags := uint8(0x30)
	if cpu.C {
		flags |= 0x01
	}
	if cpu.Z {
		flags |= 0x02
	}
	if cpu.I {
		flags |= 0x04
	}
	if cpu.D {
		flags |= 0x08
	}
	if cpu.B {
		flags |= 0x10
	}
	if cpu.V {
		flags |= 0x40
	}
	if cpu.N {
		flags |= 0x80
	}
	return flags
}

// FromFlags sets the CPU flags based on the provided byte.
func (cpu *CPU) FromFlags(flags uint8) {
	cpu.N = flags&0x80 != 0
	cpu.V = flags&0x40 != 0
	cpu.D = flags&0x08 != 0
	cpu.I = flags&0x04 != 0
	cpu.Z = flags&0x02 != 0
	cpu.C = flags&0x01 != 0
}
