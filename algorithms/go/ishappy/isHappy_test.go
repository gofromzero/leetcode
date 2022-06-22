package ishappy

import "testing"

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isHappy",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "isHappy",
			args: args{
				n: 2,
			},
			want: false,
		},
		{
			name: "isHappy",
			args: args{
				n: 16,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
