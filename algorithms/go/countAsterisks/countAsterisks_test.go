package countAsterisks

import "testing"

func Test_countAsterisks(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s: "l|*e*et|c**o|*de|",
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				s: "iamprogrammer|",
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				s: "yo|uar|e**|b|e***au|tifu|l",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAsterisks(tt.args.s); got != tt.want {
				t.Errorf("countAsterisks() = %v, want %v", got, tt.want)
			}
		})
	}
}
