package program

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "firstUniqChar case 1",
			args: args{
				s: "leetcode",
			},
			want: 0,
		},
		{
			name: "firstUniqChar case 2",
			args: args{
				s: "loveleetcode",
			},
			want: 2,
		},
		{
			name: "firstUniqChar case 3",
			args: args{
				s: "",
			},
			want: -1,
		},
		{
			name: "firstUniqChar case 4",
			args: args{
				s: "aaabbb",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstUniqCharBetter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "firstUniqChar case 1",
			args: args{
				s: "leetcode",
			},
			want: 0,
		},
		{
			name: "firstUniqChar case 2",
			args: args{
				s: "loveleetcode",
			},
			want: 2,
		},
		{
			name: "firstUniqChar case 3",
			args: args{
				s: "",
			},
			want: -1,
		},
		{
			name: "firstUniqChar case 4",
			args: args{
				s: "aaabbb",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqCharBetter(tt.args.s); got != tt.want {
				t.Errorf("firstUniqCharBetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
