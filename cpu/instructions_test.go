package cpu

type ResultState struct {
	A, X, Y, SP         uint8
	C, Z, I, D, B, V, N bool

	PC uint16
}

func CPUToResultState(cpu *CPU) ResultState {
	return ResultState{
		A: cpu.A, X: cpu.X, Y: cpu.Y, SP: cpu.SP,
		C: cpu.C, Z: cpu.Z, I: cpu.I, D: cpu.D,
		B: cpu.B, V: cpu.V, N: cpu.N,
		PC: cpu.PC,
	}
}

func CompareResultState(cpu *CPU, r ResultState) bool {
	if cpu.A != r.A || cpu.X != r.X || cpu.Y != r.Y || cpu.SP != r.SP ||
		cpu.C != r.C || cpu.Z != r.Z || cpu.I != r.I || cpu.D != r.D ||
		cpu.B != r.B || cpu.V != r.V || cpu.N != r.N {
		return false
	}
	if cpu.PC != r.PC {
		return false
	}
	//for i := 0; i < len(cpu.Memory); i++ {
	//	if cpu.Memory[i] != r.Memory[i] {
	//		return false
	//	}
	//}
	return true
}
