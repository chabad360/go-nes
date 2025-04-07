package cpu

// Break: stack[SP] = PC + 2; SP--; PC = $FFFE
func Break(cpu *CPU, operand1 uint8, operand2 uint8) {

	returnAddress := cpu.PC
	cpu.Push(uint8(returnAddress >> 8))
	cpu.Push(uint8(returnAddress & 0x00FF))

	cpu.Push(cpu.Flags())

	cpu.PC = 0xFFFE
	cpu.Wait()
	cpu.Wait()

	cpu.B = true
}
