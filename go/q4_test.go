package program

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "findMedianSortedArrays",
			args: args{
				nums1: []int{1, 3, 5, 7, 9, 11, 12, 13, 14},
				nums2: []int{2, 4, 6, 8, 10},
			},
			want: 7.5,
		},
		{
			name: "findMedianSortedArrays",
			args: args{
				nums1: []int{1, 3, 5, 7, 9, 11, 12, 13, 14},
				nums2: []int{2},
			},
			want: 8,
		},
		{
			name: "findMedianSortedArrays",
			args: args{
				nums1: []int{2},
				nums2: []int{1, 3, 5, 7, 9, 11, 12, 13, 14},
			},
			want: 8,
		},
		{
			name: "findMedianSortedArrays",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
