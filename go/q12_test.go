package program

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "intToRoman",
			args: args{
				num: 3,
			},
			want: "III",
		},
		{
			name: "intToRoman",
			args: args{
				num: 4,
			},
			want: "IV",
		},
		{
			name: "intToRoman",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "intToRoman",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
		{
			name: "intToRoman",
			args: args{
				num: 449,
			},
			want: "CDXLIX",
		},
		{
			name: "intToRoman",
			args: args{
				num: 600,
			},
			want: "DC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
