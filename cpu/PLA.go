package cpu

// PullA A = stack[SP]
func PullA(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.A = cpu.Pop()
	cpu.Z = cpu.A == 0
	cpu.N = cpu.A&0x80 != 0
	cpu.Wait()
}
