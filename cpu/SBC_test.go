package cpu

import "testing"

func TestSubtractWithCarry(t *testing.T) {
	tests := []struct {
		name   string
		cpu    *CPU
		memory uint8
		result ResultState
	}{
		{
			name:   "No Carry",
			cpu:    &CPU{A: 0x01, C: true},
			memory: 0x01,
			result: ResultState{A: 0x00, C: true, Z: true, N: false},
		},
		{
			name:   "With Carry",
			cpu:    &CPU{A: 0x01, C: false},
			memory: 0x01,
			result: ResultState{A: 0xff, C: false, Z: false, N: true},
		},
		{
			name:   "Carried",
			cpu:    &CPU{A: 0xFE, C: true},
			memory: 0xFF,
			result: ResultState{A: 0xFF, C: false, Z: false, N: true},
		},
		{
			name:   "Signed Overflow",
			cpu:    &CPU{A: 0x01, C: false},
			memory: 0x80,
			result: ResultState{A: 0x80, C: false, Z: false, N: true, V: true},
		},
		{
			name:   "Signed Overflow 2",
			cpu:    &CPU{A: 0x80, C: false},
			memory: 0x00,
			result: ResultState{A: 0x7f, C: true, Z: false, N: false, V: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SubtractWithCarry(tt.cpu, tt.memory)
			if !CompareResultState(tt.cpu, tt.result) {
				t.Errorf("SubtractWithCarry() = %#v, want %#v", CPUToResultState(tt.cpu), tt.result)
			}
		})
	}
}
