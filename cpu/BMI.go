package cpu

// BranchIfMinus PC = PC + offset if N == true
func BranchIfMinus(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := Relative(cpu, operand1, operand2)
	if cpu.N {
		cpu.Wait()
		cpu.PC = address
	}
}
