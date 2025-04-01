package cpu

// BitTest A & memory
func BitTest(cpu *CPU, memory uint8) {
	result := cpu.A & memory
	cpu.Wait()

	cpu.Z = result == 0
	cpu.V = memory&0x40 != 0
	cpu.N = memory&0x80 != 0
}

func BitTestZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := ZeroPage(cpu, operand1, operand2)
	BitTest(cpu, memory)
}

func BitTestABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	memory, _ := Absolute(cpu, operand1, operand2)
	BitTest(cpu, memory)
}
