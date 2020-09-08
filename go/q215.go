package program

import (
	"container/heap"
)

func findKthLargest(nums []int, k int) int {
	temp := IntHeap(nums[0:k])
	arr := &temp
	heap.Init(arr)

	for i := k; i < len(nums); i++ {
		heap.Push(arr, nums[i])
	}
	for i := 0; i < k-1; i++ {
		heap.Pop(arr)
	}

	return heap.Pop(arr).(int)
}
