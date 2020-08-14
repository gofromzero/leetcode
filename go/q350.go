package program

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}
	res := make([]int, 0, len(m))
	for _, v := range nums2 {
		if temp, ok := m[v]; ok {
			if temp > 0 {
				res = append(res, v)
				m[v]--
			}
		}
	}

	return res
}

func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := make([]int, 0)
	var i, j int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return res
}
