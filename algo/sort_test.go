package algo

import (
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "bubbleSort",
			args: args{
				arr: []int{3, 4, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubbleSort(tt.args.arr)
			if !checkSort(tt.args.arr) {
				t.Errorf("bubbleSort() = %v, want sort", tt.args.arr)
			}

		})
	}
}

func checkSort(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func Test_radixSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "radixSort",
			args: args{
				arr: []int{3, 4, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			radixSort(tt.args.arr)
			if !checkSort(tt.args.arr) {
				t.Errorf("radixSort() = %v, want sort", tt.args.arr)
			}
		})
	}
}

func Test_selectSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "selectSort",
			args: args{
				arr: []int{3, 4, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selectSort(tt.args.arr)
			if !checkSort(tt.args.arr) {
				t.Errorf("selectSort() = %v, want sort", tt.args.arr)
			}
		})
	}
}

func Test_insertSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "insertSort",
			args: args{
				arr: []int{3, 4, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertSort(tt.args.arr)
			if !checkSort(tt.args.arr) {
				t.Errorf("insertSort() = %v, want sort", tt.args.arr)
			}
		})
	}
}

func Test_shellSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "shellSort",
			args: args{
				arr: []int{3, 4, 1},
			},
		},
		{
			name: "shellSort",
			args: args{
				arr: []int{3, 4, 3, 411, 331, 3131313, 1331, 122, 12},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shellSort(tt.args.arr)
			if !checkSort(tt.args.arr) {
				t.Errorf("shellSort() = %v, want sort", tt.args.arr)
			}
		})
	}
}

func Test_mergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "mergeSort",
			args: args{
				arr: []int{3, 4, 1},
			},
		},
		{
			name: "mergeSort",
			args: args{
				arr: []int{3, 4, 3, 411, 331, 3131313, 1331, 122, 12},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.arr); !checkSort(got) {
				t.Errorf("mergeSort() = %v", got)
			}
		})
	}
}
