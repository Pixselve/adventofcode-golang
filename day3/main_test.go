package main

import (
	"reflect"
	"testing"
)

func TestStringToInstruction(t *testing.T) {
	type args struct {
		instruction string
	}
	tests := []struct {
		name    string
		args    args
		want    Instruction
		wantErr bool
	}{
		{name: "Test UP Direction with Amount of 1", args: args{instruction: "U1"}, want: Instruction{
			Direction: "U",
			Amount:    1,
		}, wantErr: false},
		{name: "Test LEFT Direction with Amount of 1", args: args{instruction: "L1"}, want: Instruction{
			Direction: "L",
			Amount:    1,
		}, wantErr: false},
		{name: "Test RIGHT Direction with Amount of 1", args: args{instruction: "R1"}, want: Instruction{
			Direction: "R",
			Amount:    1,
		}, wantErr: false},
		{name: "Test DOWN Direction with Amount of 1", args: args{instruction: "D1"}, want: Instruction{
			Direction: "D",
			Amount:    1,
		}, wantErr: false},

		{name: "Test UP Direction with Amount of 2", args: args{instruction: "U2"}, want: Instruction{
			Direction: "U",
			Amount:    2,
		}, wantErr: false},
		{name: "Test LEFT Direction with Amount of 2", args: args{instruction: "L2"}, want: Instruction{
			Direction: "L",
			Amount:    2,
		}, wantErr: false},
		{name: "Test RIGHT Direction with Amount of 2", args: args{instruction: "R2"}, want: Instruction{
			Direction: "R",
			Amount:    2,
		}, wantErr: false},
		{name: "Test DOWN Direction with Amount of 2", args: args{instruction: "D2"}, want: Instruction{
			Direction: "D",
			Amount:    2,
		}, wantErr: false},

		{name: "Test UP Direction with Amount of 3", args: args{instruction: "U3"}, want: Instruction{
			Direction: "U",
			Amount:    3,
		}, wantErr: false},
		{name: "Test LEFT Direction with Amount of 3", args: args{instruction: "L3"}, want: Instruction{
			Direction: "L",
			Amount:    3,
		}, wantErr: false},
		{name: "Test RIGHT Direction with Amount of 3", args: args{instruction: "R3"}, want: Instruction{
			Direction: "R",
			Amount:    3,
		}, wantErr: false},
		{name: "Test DOWN Direction with Amount of 3", args: args{instruction: "D3"}, want: Instruction{
			Direction: "D",
			Amount:    3,
		}, wantErr: false},

		{name: "Test UP Direction with INVALID Amount", args: args{instruction: "U-1"}, want: Instruction{}, wantErr: true},
		{name: "Test LEFT Direction with INVALID Amount", args: args{instruction: "L-1"}, want: Instruction{}, wantErr: true},
		{name: "Test RIGHT Direction with INVALID Amount", args: args{instruction: "R-1"}, want: Instruction{}, wantErr: true},
		{name: "Test DOWN Direction with INVALID Amount", args: args{instruction: "D-1"}, want: Instruction{}, wantErr: true},

		{name: "Test UP Direction with no Amount", args: args{instruction: "U"}, want: Instruction{}, wantErr: true},
		{name: "Test LEFT Direction with no Amount", args: args{instruction: "L"}, want: Instruction{}, wantErr: true},
		{name: "Test RIGHT Direction with no Amount", args: args{instruction: "R"}, want: Instruction{}, wantErr: true},
		{name: "Test DOWN Direction with no Amount", args: args{instruction: "D"}, want: Instruction{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToInstruction(tt.args.instruction)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToInstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToInstruction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawStringInstructions(t *testing.T) {
	type args struct {
		instructionsRaw string
	}
	tests := []struct {
		name string
		args args
		want []Instruction
	}{
		{
			name: "Empty string",
			args: args{
				instructionsRaw: "",
			},
			want: nil,
		},
		{
			name: "1 instruction",
			args: args{
				instructionsRaw: "U1",
			},
			want: []Instruction{{
				Direction: "U",
				Amount:    1,
			}},
		},
		{
			name: "2 instruction",
			args: args{
				instructionsRaw: "U7,D5",
			},
			want: []Instruction{{
				Direction: "U",
				Amount:    7,
			}, {
				Direction: "D",
				Amount:    5,
			}},
		},
		{
			name: "3 instruction",
			args: args{
				instructionsRaw: "U8,D2,L3",
			},
			want: []Instruction{{
				Direction: "U",
				Amount:    8,
			}, {
				Direction: "D",
				Amount:    2,
			}, {
				Direction: "L",
				Amount:    3,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RawStringInstructions(tt.args.instructionsRaw); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RawStringInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstructionsToCoordinatesMap(t *testing.T) {
	type args struct {
		instructions []Instruction
	}
	tests := []struct {
		name string
		args args
		want map[Position]uint
	}{
		{
			name: "1 Instruction",
			args: args{
				instructions: []Instruction{{Direction: "U", Amount: 1}},
			},
			want: map[Position]uint{Position{
				X: 0,
				Y: 1,
			}: 1},
		},
		{
			name: "2 Instruction",
			args: args{
				instructions: []Instruction{{Direction: "U", Amount: 1}, {
					Direction: "R",
					Amount:    1,
				}},
			},
			want: map[Position]uint{Position{
				X: 0,
				Y: 1,
			}: 1, Position{
				X: 1,
				Y: 1,
			}: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InstructionsToCoordinatesMap(tt.args.instructions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstructionsToCoordinatesMap() = %v, want %v", got, tt.want)
			}
		})
	}
}