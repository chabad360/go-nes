package cpu

// JumpToSubroutine push PC + 2 to stack, PC = memory
func JumpToSubroutine(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := Absolute(cpu, operand1, operand2)
	returnAddress := cpu.PC - 1
	cpu.Push(uint8(returnAddress >> 8))
	cpu.Wait()
	cpu.Push(uint8(returnAddress & 0x00FF))
	cpu.Wait()
	cpu.PC = address
	cpu.Wait()
}
