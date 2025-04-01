package cpu

import (
	"fmt"
	"log/slog"
)

// CPU represents the state of the NES 6502 CPU.
type CPU struct {
	A, X, Y, SP         uint8
	PC                  uint16
	C, Z, I, D, B, V, N bool
	Memory              [0x07FF]uint8

	Mapper Mapper

	Clock  <-chan int
	Cycles int

	Log     bool
	Stubbed bool
}

// PowerOn initializes the CPU to a powered-on state.
func (cpu *CPU) PowerOn() {
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

	cpu.Memory = [0x07FF]uint8{}
}

// Reset initializes the CPU to the reset state.
func (cpu *CPU) Reset() {
	cpu.PC = 0xFFFC
	cpu.SP = cpu.SP - 3
	cpu.I = true
}

type Exec struct {
	Instruction Instruction
	Opcode      uint8
	Operand1    uint8
	Operand2    uint8
	PC          uint16
}

// Decode fetches the next opcode from memory and returns the corresponding instruction.
func (cpu *CPU) Decode() *Exec {
	opcode := cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC)
	cpu.Wait()
	instruction, exists := Instructions[opcode]
	if !exists {
		instruction = Instruction{
			Opcode:  opcode,
			Size:    0,
			Cycles:  2,
			Execute: UNKNOWN,
		}
	}

	exec := &Exec{
		Instruction: instruction,
		Opcode:      opcode,
		PC:          cpu.PC,
	}

	cpu.Wait() // Fake fetch (in theory, a cycle accurate system should also load and store the next byte here)
	cpu.PC++

	return exec
}

// Fetch retrieves the operands for the instruction based on its size.
func (cpu *CPU) Fetch(exec *Exec) {
	if exec.Instruction.Size >= 2 {
		exec.Operand1 = cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC)
		cpu.PC++
		// We don't need to wait, because the 6502 always prefetches the next byte
	}
	if exec.Instruction.Size == 3 {
		exec.Operand2 = cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC)
		cpu.PC++
		cpu.Wait()
	}
}

// Execute runs the instruction with the provided operands.
func (cpu *CPU) Execute(exec *Exec) {
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

	exec.Instruction.Execute(cpu, exec.Operand1, exec.Operand2)
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
	cpu.Mapper.Write(cpu, cpu.Clock, 0x0100+uint16(cpu.SP), value)
	cpu.SP--
}

// Pop pops a value from the stack.
func (cpu *CPU) Pop() uint8 {
	cpu.SP++
	value := cpu.Mapper.Read(cpu, cpu.Clock, 0x0100+uint16(cpu.SP))
	return value
}

// Wait waits for the next clock cycle.
func (cpu *CPU) Wait() {
	cpu.Cycles = <-cpu.Clock
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
