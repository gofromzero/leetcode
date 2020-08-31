package program

import "testing"

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n          int
		badVersion int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "firstBadVersion",
			args: args{
				n:          5,
				badVersion: 4,
			},
			want: 4,
		},

		{
			name: "firstBadVersion",
			args: args{
				n:          10,
				badVersion: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badVersion = tt.args.badVersion
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
