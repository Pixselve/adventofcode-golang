package main

import "testing"

func TestAscendingOrder(t *testing.T) {
	type args struct {
		input uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid repeat",
			args: args{
				input: 111111,
			},
			want: true,
		},
		{
			name: "Valid no repeat",
			args: args{
				input: 123456,
			},
			want: true,
		},
		{
			name: "Not valid no repeat",
			args: args{
				input: 123450,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AscendingOrder(tt.args.input); got != tt.want {
				t.Errorf("AscendingOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtLeastARepeat(t *testing.T) {
	type args struct {
		input uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid",
			args: args{
				input: 111111,
			},
			want: true,
		},
		{
			name: "Valid",
			args: args{
				input: 122345,
			},
			want: true,
		},
		{
			name: "Valid",
			args: args{
				input: 123789,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AtLeastARepeat(tt.args.input); got != tt.want {
				t.Errorf("AtLeastARepeat() = %v, want %v", got, tt.want)
			}
		})
	}
}