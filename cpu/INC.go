package cpu

// Increment memory++
func Increment(cpu *CPU, value uint8, address uint16) (result uint8) {
	result = value + 1

	cpu.Z = result == 0
	cpu.N = result&0x80 != 0

	cpu.Mapper.Write(cpu, address, value)
	cpu.Wait()
	cpu.Mapper.Write(cpu, address, result)
	cpu.Wait()

	return result
}

func IncrementZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPage(cpu, operand1, operand2)
	Increment(cpu, value(), address)
}

func IncrementZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPageX(cpu, operand1, operand2)
	Increment(cpu, value(), address)
}

func IncrementABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := Absolute(cpu, operand1, operand2)
	Increment(cpu, value(), address)
}

func IncrementABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteX(cpu, operand1, operand2)
	cpu.Wait()
	Increment(cpu, value(), address)
}
