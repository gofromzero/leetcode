package program

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "addTwoNumbers case 1",
			args: args{
				l1: []int{2, 4, 3},
				l2: []int{5, 6, 4},
			},
			want: []int{7, 0, 8},
		},
		{
			name: "addTwoNumbers case 2",
			args: args{
				l1: []int{2, 4, 6},
				l2: []int{5, 6, 4},
			},
			want: []int{7, 0, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := GenerateListNode(tt.args.l1)
			l2 := GenerateListNode(tt.args.l2)
			want := GenerateListNode(tt.want)

			if got := addTwoNumbers(l1, l2); !reflect.DeepEqual(got, want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
