// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/1

package reversestring

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "reverseString",
			args: args{
				s: []byte{'h', 'e', 'l', 'l', 'o'},
			},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name: "reverseString",
			args: args{
				s: []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("reverseString() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
