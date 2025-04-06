package cpu

// BranchIfPlus PC = PC + offset if N == false
func BranchIfPlus(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := Relative(cpu, operand1, operand2)
	if !cpu.N {
		cpu.Wait()
		if address&0xFF00 != cpu.PC&0xFF00 {
			cpu.Wait()
		}
		cpu.PC = address
	}
}
