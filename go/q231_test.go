package program

import "testing"

func Test_isPowerOfTwoBetter(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isPowerOfTwoBetter",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "isPowerOfTwoBetter",
			args: args{
				n: 16,
			},
			want: true,
		},
		{
			name: "isPowerOfTwoBetter",
			args: args{
				n: 218,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwoBetter(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwoBetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isPowerOfTwoBetter",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "isPowerOfTwoBetter",
			args: args{
				n: 16,
			},
			want: true,
		},
		{
			name: "isPowerOfTwoBetter",
			args: args{
				n: 218,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
