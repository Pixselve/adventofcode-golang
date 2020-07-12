package main

import (
	"reflect"
	"testing"
)

func TestOpcodeManager(t *testing.T) {
	type args struct {
		position uint
		integers []uint
	}
	tests := []struct {
		name    string
		args    args
		want    []uint
		wantErr bool
	}{
		{name: "1,9,10,3,2,3,11,0,99,30,40,50", args: args{
			position: 0,
			integers: []uint{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
		}, want: []uint{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, wantErr: false},

		{name: "1,0,0,0,99", args: args{
			position: 0,
			integers: []uint{1, 0, 0, 0, 99},
		}, want: []uint{2, 0, 0, 0, 99}, wantErr: false},

		{name: "2,3,0,3,99", args: args{
			position: 0,
			integers: []uint{2, 3, 0, 3, 99},
		}, want: []uint{2, 3, 0, 6, 99}, wantErr: false},

		{name: "2,4,4,5,99,0", args: args{
			position: 0,
			integers: []uint{2, 4, 4, 5, 99, 0},
		}, want: []uint{2, 4, 4, 5, 99, 9801}, wantErr: false},

		{name: "1,0,0,0,99", args: args{
			position: 0,
			integers: []uint{1, 1, 1, 4, 99, 5, 6, 0, 99},
		}, want: []uint{30, 1, 1, 4, 2, 5, 6, 0, 99}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OpcodeManager(tt.args.position, tt.args.integers)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpcodeManager() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpcodeManager() got = %v, want %v", got, tt.want)
			}
		})
	}
}
