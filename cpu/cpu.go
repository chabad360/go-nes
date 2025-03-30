package cpu

// CPU represents the state of the NES 6502 CPU.
type CPU struct {
	A, X, Y, SP         uint8
	PC                  uint16
	C, Z, I, D, B, V, N bool
	Memory              [0x07FF]uint8

	Mapper Mapper

	Clock <-chan int
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
	PC          uint16
}

func (cpu *CPU) Decode() Exec {
	opcode := cpu.Mapper.Read(cpu, cpu.Clock, cpu.PC)
	instruction, exists := Instructions[opcode]
	if !exists {
		instruction = NOPInstruction
	}
	cpu.PC += uint16(instruction.Size)

	return Exec{
		Instruction: instruction,
		PC:          cpu.PC - uint16(instruction.Size),
	}
}

func (cpu *CPU) Execute(exec Exec) {
	var operand1, operand2 uint8
	if exec.Instruction.Size == 2 {
		operand1 = cpu.Mapper.Read(cpu, cpu.Clock, exec.PC+1)
	}
	if exec.Instruction.Size == 3 {
		operand2 = cpu.Mapper.Read(cpu, cpu.Clock, exec.PC+2)
	}
	exec.Instruction.Execute(cpu, operand1, operand2)
	for i := 0; i < exec.Instruction.Cycles; i++ {
		<-cpu.Clock
	}
}
