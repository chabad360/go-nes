package cpu

// ArithmeticShiftLeft value = value << 1
func ArithmeticShiftLeft(cpu *CPU, value uint8, address uint16, accumulator bool) {
	result := value << 1

	cpu.C = value&0x80 != 0
	cpu.Z = result == 0
	cpu.N = result&0x80 != 0

	if accumulator {
		cpu.A = result
	} else {
		cpu.Mapper.Write(cpu, address, value)
		cpu.Wait()
		cpu.Mapper.Write(cpu, address, result)
		cpu.Wait()
	}
}

func ArithmeticShiftLeftA(cpu *CPU, operand1 uint8, operand2 uint8) {
	ArithmeticShiftLeft(cpu, cpu.A, 0, true)
}

func ArithmeticShiftLeftZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPage(cpu, operand1, operand2)
	ArithmeticShiftLeft(cpu, value(), address, false)
}

func ArithmeticShiftLeftZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPageX(cpu, operand1, operand2)
	ArithmeticShiftLeft(cpu, value(), address, false)
}

func ArithmeticShiftLeftABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := Absolute(cpu, operand1, operand2)
	ArithmeticShiftLeft(cpu, value(), address, false)
}

func ArithmeticShiftLeftABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteX(cpu, operand1, operand2)
	ArithmeticShiftLeft(cpu, value(), address, false)
}
