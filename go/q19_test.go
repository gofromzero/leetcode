package program

import (
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "removeNthFromEnd case 1",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				n:    1,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "removeNthFromEnd case 2",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				n:    2,
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "removeNthFromEnd case 3",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				n:    3,
			},
			want: []int{1, 2, 4, 5},
		},
		{
			name: "removeNthFromEnd case 4",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				n:    5,
			},
			want: []int{2, 3, 4, 5},
		},
		{
			name: "removeNthFromEnd case 5",
			args: args{
				head: nil,
				n:    5,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := GenerateListNode(tt.args.head)
			want := GenerateListNode(tt.want)

			if got := removeNthFromEnd(args, tt.args.n); !CheckListEqual(got, want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, want)
			}
		})
	}
}
