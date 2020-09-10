package algo

import "math"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func radixSort(arr []int) {
	digit := getMaxDigit(arr)
	radixSortIn(arr, digit)
}

func radixSortIn(arr []int, digit int) {
	mod, dev := 10, 1
	for i := 0; i < digit; i, dev, mod = i+1, dev*10, mod*10 {
		counter := make([][]int, mod*2)
		for j := 0; j < len(arr); j++ {
			bucket := (arr[j]%mod)/dev + mod
			counter[bucket] = append(counter[bucket], arr[j])
		}

		pos := 0
		for _, bucket := range counter {
			for _, val := range bucket {
				arr[pos] = val
				pos++
			}
		}
	}
}

func getMaxDigit(arr []int) int {
	maxValue := getMaxValue(arr)
	return getNumLength(maxValue)
}

func getMaxValue(arr []int) int {
	var max = math.MinInt64
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func getNumLength(value int) int {
	var count int
	for value > 0 {
		value >>= 1
		count++
	}
	return count
}

func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		var index = i
		for j := i; j < len(arr); j++ {
			if arr[index] > arr[j] {
				index = j
			}
		}

		arr[i], arr[index] = arr[index], arr[i]
	}
}

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		var temp = arr[i]
		var j = i
		for ; j > 0; j-- {
			if temp < arr[j-1] {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		arr[j] = temp
	}
}

func shellSort(arr []int) {
	length := len(arr)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1
	}

	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}
	return result
}

func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}
func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}

func heapSort(arr []int) []int {
	arrLen := len(arr)
	buildMaxHeap(arr, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		arrLen--
		heapify(arr, 0, arrLen)
	}
	return arr
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, largest, arrLen)
	}
}
