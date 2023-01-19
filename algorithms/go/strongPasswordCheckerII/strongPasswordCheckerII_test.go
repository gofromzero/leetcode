package strongPasswordCheckerII

import "testing"

func Test_strongPasswordCheckerII(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				password: "IloveLe3tcode!",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				password: "Ilove",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				password: "IIloveLe3tcode",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strongPasswordCheckerII(tt.args.password); got != tt.want {
				t.Errorf("strongPasswordCheckerII() = %v, want %v", got, tt.want)
			}
		})
	}
}
