// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/8

package ispoweroftwo

import "testing"

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
			name: "",
			args: args{
				n: 1 << 3,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				n: 1<<3 + 3,
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
