package cpu

// StoreX memory = X
func StoreX(cpu *CPU, memory uint16) {
	cpu.Mapper.Write(cpu, cpu.Clock, memory, cpu.X)
	cpu.Wait()
}

func StoreXZP(cpu *CPU, operand1, operand2 uint8) {
	_, address := ZeroPage(cpu, operand1, operand2)
	StoreX(cpu, address)
}

func StoreXZPY(cpu *CPU, operand1, operand2 uint8) {
	_, address := ZeroPageY(cpu, operand1, operand2)
	StoreX(cpu, address)
}

func StoreXABS(cpu *CPU, operand1, operand2 uint8) {
	_, address := Absolute(cpu, operand1, operand2)
	StoreX(cpu, address)
}
