package program

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
			name: "",
			args: args{
				s: []byte{'h', 'e', 'l', 'l', 'o'},
			},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)

			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("intersect2() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
