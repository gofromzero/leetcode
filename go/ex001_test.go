package program

import (
	"reflect"
	"testing"
)

func Test_printNumbers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "printNumbers",
			args: args{
				n: 1,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printNumbers(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printNumbers2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "printNumbers",
			args: args{
				n: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printNumbers2(2)
		})
	}
}
