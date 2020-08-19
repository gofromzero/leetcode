package program

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head []int
		pos  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "hasCycle case 1",
			args: args{
				head: []int{3, 2, 0, -4},
				pos:  0,
			},
			want: true,
		},
		{
			name: "hasCycle case 2",
			args: args{
				head: []int{3, 2, 0, -4},
				pos:  -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := GenerateListNode(tt.args.head)
			tail := head
			for tail.Next != nil {
				tail = tail.Next
			}

			root := head
			for i := 0; i < tt.args.pos; i++ {
				root = root.Next
			}
			if tt.args.pos >= 0 {
				tail.Next = root
			}

			if got := hasCycle(head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
