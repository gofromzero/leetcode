package program

import "testing"

func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		H     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "minEatingSpeed",
			args: args{
				piles: []int{3, 6, 7, 11},
				H:     8,
			},
			want: 4,
		},
		{
			name: "minEatingSpeed",
			args: args{
				piles: []int{30, 11, 23, 4, 20},
				H:     5,
			},
			want: 30,
		},
		{
			name: "minEatingSpeed",
			args: args{
				piles: []int{30, 11, 23, 4, 20},
				H:     6,
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.args.piles, tt.args.H); got != tt.want {
				t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
