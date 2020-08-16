package program

func rotate(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
