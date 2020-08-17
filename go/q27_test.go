package program

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "removeElement case 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "removeElement case 2",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			if !reflect.DeepEqual(tt.args.nums[:got], tt.args.nums[:tt.want]) {
				t.Errorf("removeElement() = %v, want %v", tt.args.nums[:got], tt.args.nums[:tt.want])
			}
		})
	}
}
