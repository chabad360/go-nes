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

	Clock <-chan int

	Log bool
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
func (cpu *CPU) Decode() Exec {
	opcode := cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC)
	instruction, exists := Instructions[opcode]
	if !exists {
		instruction = Instruction{
			Opcode:  opcode,
			Size:    0,
			Cycles:  2,
			Execute: UNKNOWN,
		}
	}
	cpu.PC++

	return Exec{
		Instruction: instruction,
		Opcode:      opcode,
		PC:          cpu.PC - 1,
	}
}

// Fetch retrieves the operands for the instruction based on its size.
func (cpu *CPU) Fetch(exec Exec) {
	// Fetch the operands based on the instruction size
	var operand1, operand2 uint8
	if exec.Instruction.Size == 2 {
		operand1 = cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC)
		cpu.PC++
	}
	if exec.Instruction.Size == 3 {
		operand2 = cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC)
		cpu.PC++
	}

	exec.Operand1 = operand1
	exec.Operand2 = operand2
}

// Execute runs the instruction with the provided operands.
func (cpu *CPU) Execute(exec Exec) {
	if cpu.Log {
		slog.Debug(fmt.Sprintf("%x %x %x", exec.Instruction.Opcode, exec.Operand1, exec.Operand2))
	}

	exec.Instruction.Execute(cpu, exec.Operand1, exec.Operand2)
	for i := 0; i < exec.Instruction.Cycles; i++ {
		<-cpu.Clock
	}
}
