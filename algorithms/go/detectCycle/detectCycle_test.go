package detectCycle

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {

	arr1 := GenList([]int{3, 2, 0, -4})
	head1 := arr1.Next
	var last1 = head1
	for last1.Next != nil {
		last1 = last1.Next
	}
	last1.Next = head1

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: arr1,
			},
			want: head1,
		},
		{
			name: "",
			args: args{
				head: GenList([]int{1}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
