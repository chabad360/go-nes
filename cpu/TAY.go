package cpu

// TransferAToY Y = A
func TransferAToY(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Y = cpu.A
	cpu.Z = cpu.Y == 0
	cpu.N = cpu.Y&0x80 != 0
}
