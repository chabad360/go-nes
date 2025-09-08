package cpu

// StoreY memory = Y
func StoreY(cpu *CPU, address uint16) {
	cpu.Mapper.Write(cpu, address, cpu.Y)
	cpu.Wait()
}

func StoreYZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := ZeroPage(cpu, operand1, operand2, true)
	StoreY(cpu, address)
}

func StoreYZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := ZeroPageX(cpu, operand1, operand2, true)
	StoreY(cpu, address)
}

func StoreYABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := Absolute(cpu, operand1, operand2, true)
	StoreY(cpu, address)
}
