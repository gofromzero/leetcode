package program

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "strStr case 1",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "strStr case 2",
			args: args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			want: -1,
		},
		{
			name: "strStr case 3",
			args: args{
				haystack: "",
				needle:   "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strStrInternal(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "strStr case 1",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "strStr case 2",
			args: args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			want: -1,
		},
		{
			name: "strStr case 3",
			args: args{
				haystack: "mississippi",
				needle:   "sippia",
			},
			want: -1,
		},
		{
			name: "strStr case 4",
			args: args{
				haystack: "",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "strStr case 1",
			args: args{
				haystack: "hexllo",
				needle:   "ll",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStrInternal(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStrInternal() = %v, want %v", got, tt.want)
			}
		})
	}
}
