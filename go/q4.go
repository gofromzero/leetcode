package program

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	total := len1 + len2
	left := (total + 1) / 2
	right := (total + 2) / 2
	return (float64(findK(nums1, 0, nums2, 0, left)) + float64(findK(nums1, 0, nums2, 0, right))) / 2
}

func findK(nums1 []int, i int, nums2 []int, j int, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		return nums1[i+k-1]
	}
	if k == 1 {
		return min(nums1[i], nums2[j])
	}

	var mid1, mid2 int
	if i+k/2-1 < len(nums1) {
		mid1 = nums1[i+k/2-1]
	} else {
		mid1 = math.MaxInt64
	}
	if j+k/2-1 < len(nums2) {
		mid2 = nums2[j+k/2-1]
	} else {
		mid2 = math.MaxInt64
	}
	if mid1 < mid2 {
		return findK(nums1, i+k/2, nums2, j, k-k/2)
	}
	return findK(nums1, i, nums2, j+k/2, k-k/2)
}
