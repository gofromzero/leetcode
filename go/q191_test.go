package program

import "testing"

func Test_hammingWeight2(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "hammingWeight2",
			args: args{
				num: 15,
			},
			want: 4,
		},
		{
			name: "hammingWeight2",
			args: args{
				num: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight2(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "hammingWeight",
			args: args{
				num: 15,
			},
			want: 4,
		},
		{
			name: "hammingWeight",
			args: args{
				num: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
