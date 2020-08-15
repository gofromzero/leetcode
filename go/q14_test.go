package program

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{
		//	name: "case 1",
		//	args: args{
		//		strs: []string{"flower","flow","flight"},
		//	},
		//	want: "fl",
		//},
		//{
		//	name: "case 2",
		//	args: args{
		//		strs: []string{"dog","racecar","car"},
		//	},
		//	want: "",
		//},
		//{
		//	name: "case 3",
		//	args: args{
		//		strs: []string{},
		//	},
		//	want: "",
		//},
		{
			name: "case 3",
			args: args{
				strs: []string{"", "1"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
