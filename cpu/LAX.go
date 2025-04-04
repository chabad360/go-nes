package cpu

// LoadAX A, X = memory
func LoadAX(cpu *CPU, memory uint8) {
	cpu.A = memory
	cpu.X = memory

	cpu.Z = cpu.A == 0
	cpu.N = cpu.A&0x80 != 0
}

func LoadAXIMM(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Immediate(cpu, operand1, operand2)
	LoadAX(cpu, memory())
}

func LoadAXZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPage(cpu, operand1, operand2)
	LoadAX(cpu, memory())
}

func LoadAXZPY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPageY(cpu, operand1, operand2)
	LoadAX(cpu, memory())
}

func LoadAXABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Absolute(cpu, operand1, operand2)
	LoadAX(cpu, memory())
}

func LoadAXABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := AbsoluteY(cpu, operand1, operand2)
	LoadAX(cpu, memory())
}

func LoadAXINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndexedIndirect(cpu, operand1, operand2)
	LoadAX(cpu, memory())
}

func LoadAXINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := IndirectIndexed(cpu, operand1, operand2)
	LoadAX(cpu, memory())
}
