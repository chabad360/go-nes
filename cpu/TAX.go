package cpu

// TransferAToX X = A
func TransferAToX(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.X = cpu.A
	cpu.Z = cpu.X == 0
	cpu.N = cpu.X&0x80 != 0
}
