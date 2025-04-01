package cpu

// TransferYToA A = Y
func TransferYToA(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.A = cpu.Y
	cpu.Z = cpu.A == 0
	cpu.N = cpu.A&0x80 != 0
}
