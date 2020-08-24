package program

import (
	"strings"
	"testing"
)

func Test_rotateString(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "rotateString case 1",
			args: args{
				A: "",
				B: "",
			},
			want: true,
		},
		{
			name: "rotateString case 2",
			args: args{
				A: "abcde",
				B: "cdeab",
			},
			want: true,
		},
		{
			name: "rotateString case 3",
			args: args{
				A: "abcde",
				B: "abcdd",
			},
			want: false,
		},
		{
			name: "rotateString case 4",
			args: args{
				A: "abcde",
				B: "abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateStringBetter(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "rotateStringBetter case 1",
			args: args{
				A: "",
				B: "",
			},
			want: true,
		},
		{
			name: "rotateStringBetter case 2",
			args: args{
				A: "abcde",
				B: "cdeab",
			},
			want: true,
		},
		{
			name: "rotateStringBetter case 3",
			args: args{
				A: "abcde",
				B: "abcdd",
			},
			want: false,
		},
		{
			name: "rotateStringBetter case 4",
			args: args{
				A: "abcde",
				B: "abc",
			},
			want: false,
		},
		{
			name: "rotateStringBetter case 5 len ",
			args: args{
				A: strings.Repeat("1", 101),
				B: strings.Repeat("1", 101),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateStringBetter(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("rotateStringBetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
