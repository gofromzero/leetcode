package program

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
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
			name: "mergeTwoLists case 1",
			args: args{
				l1: []int{1, 2, 3, 4, 5},
				l2: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		},
		{
			name: "mergeTwoLists case 2",
			args: args{
				l1: []int{-10, -9, -6, -4, 1, 9, 9},
				l2: []int{-5, -3, 0, 7, 8, 8},
			},
			want: []int{-10, -9, -6, -5, -4, -3, 0, 1, 7, 8, 8, 9, 9},
		},
		{
			name: "mergeTwoLists case 3",
			args: args{
				l1: []int{-10, -9, -6, -4, 1, 9, 9},
				l2: []int{},
			},
			want: []int{-10, -9, -6, -4, 1, 9, 9},
		},
		{
			name: "mergeTwoLists case 4",
			args: args{
				l1: []int{},
				l2: []int{-5, -3, 0, 7, 8, 8},
			},
			want: []int{-5, -3, 0, 7, 8, 8},
		},
		{
			name: "mergeTwoLists case 5",
			args: args{
				l1: []int{-5, -3, 0, 7, 8, 8},
				l2: []int{-10, -9, -6, -4, 1, 9, 9},
			},
			want: []int{-10, -9, -6, -5, -4, -3, 0, 1, 7, 8, 8, 9, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := GenerateListNode(tt.args.l1)
			l2 := GenerateListNode(tt.args.l2)
			want := GenerateListNode(tt.want)
			if got := mergeTwoLists(l1, l2); !reflect.DeepEqual(got, want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, want)
			}
		})
	}
}
