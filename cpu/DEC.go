package cpu

// Decrement memory--
func Decrement(cpu *CPU, value uint8, address uint16) {
	result := value - 1

	cpu.Z = result == 0
	cpu.N = result&0x80 != 0

	cpu.Mapper.Write(cpu, address, value)
	cpu.Wait()
	cpu.Mapper.Write(cpu, address, result)
	cpu.Wait()
}

func DecrementZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPage(cpu, operand1, operand2)
	Decrement(cpu, value(), address)
}

func DecrementZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPageX(cpu, operand1, operand2)
	Decrement(cpu, value(), address)
}

func DecrementABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := Absolute(cpu, operand1, operand2)
	Decrement(cpu, value(), address)
}

func DecrementABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteX(cpu, operand1, operand2)
	cpu.Wait()
	Decrement(cpu, value(), address)
}
