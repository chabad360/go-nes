package cpu

type Instruction struct {
	Opcode  uint8
	Size    int
	Mode    AddressingMode
	Cycles  int
	Execute ExecutionFunc
}

type AddressingMode int

const (
	Immediate AddressingMode = iota
	ZeroPage
	ZeroPageX
	ZeroPageY
	Absolute
	AbsoluteX
	AbsoluteY
	Indirect
	IndirectX
	IndirectY
	Relative
	Accumulator
	Implicit
	Branch
	None
)

type ExecutionFunc func(*CPU, uint8, uint8)

var Instructions = map[uint8]Instruction{
	0x69: {Opcode: 0x69, Size: 2, Mode: Immediate, Cycles: 2, Execute: AddWithCarryIMM},

	0xEA: NOPInstruction,
}
var NOPInstruction = Instruction{Opcode: 0xEA, Size: 1, Mode: Implicit, Cycles: 2, Execute: NOP}

// NOP is a no-operation instruction
func NOP(cpu *CPU, operand1 uint8, operand2 uint8) {
	// No operation
}

// AddWithCarryIMM adds the immediate value to the accumulator with carry
func AddWithCarryIMM(cpu *CPU, operand1 uint8, operand2 uint8) {
	r := uint16(cpu.A)
	r += uint16(operand1)
	if cpu.C {
		r++
	}

	cpu.C = r > 0xFF
	cpu.Z = uint8(r) == 0
	cpu.N = r&0x80 != 0
	cpu.V = (r^uint16(cpu.A))&(r^uint16(operand1))&0x80 != 0
	cpu.A = uint8(r)

}
