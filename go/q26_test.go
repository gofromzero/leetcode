package program

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "removeDuplicates case 1",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
		{
			name: "removeDuplicates case 2",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "removeDuplicates case 3",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "removeDuplicates case 4",
			args: args{
				nums: []int{1, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums[:got], tt.args.nums[:tt.want]) {
				t.Errorf("removeElement() = %v, want %v", tt.args.nums[:got], tt.args.nums[:tt.want])
			}
		})
	}
}
