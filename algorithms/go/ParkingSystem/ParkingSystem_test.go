package ParkingSystem

import "testing"

func TestParkingSystem_AddCar(t *testing.T) {
	type fields struct {
		Arr []int
	}

	tests := []struct {
		name   string
		fields fields
		args   []int
		want   []bool
	}{
		{
			name: "",
			fields: fields{
				Arr: []int{1, 1, 0},
			},
			args: []int{1, 2, 3, 1},
			want: []bool{true, true, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.fields.Arr[0], tt.fields.Arr[1], tt.fields.Arr[2])
			for i, val := range tt.args {
				if got := this.AddCar(val); got != tt.want[i] {
					t.Errorf("this.AddCar() = %v, want %v", got, tt.want[i])
				}
			}

		})
	}
}
