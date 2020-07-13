package main

import (
	"reflect"
	"testing"
)

func TestNewInstruction(t *testing.T) {
	type args struct {
		instruction int
	}
	tests := []struct {
		name string
		args args
		want Instruction
	}{
		{name: "", args: args{
			instruction: 1002,
		}, want: Instruction{
			opcode:     02,
			mode1param: 0,
			mode2param: 1,
			mode3param: 0,
		}},
		{name: "", args: args{
			instruction: 99,
		}, want: Instruction{
			opcode:     99,
			mode1param: 0,
			mode2param: 0,
			mode3param: 0,
		}},
		{name: "", args: args{
			instruction: 12345,
		}, want: Instruction{
			opcode:     45,
			mode1param: 3,
			mode2param: 2,
			mode3param: 1,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInstruction(tt.args.instruction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
