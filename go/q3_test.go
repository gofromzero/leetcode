package program

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	name: "lengthOfLongestSubstring",
		//	args: args{
		//		s: "abcabcbb",
		//	},
		//	want: 3,
		//},
		//{
		//	name: "lengthOfLongestSubstring",
		//	args: args{
		//		s: "bbbbb",
		//	},
		//	want: 1,
		//},
		{
			name: "lengthOfLongestSubstring",
			args: args{
				s: "abba",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
