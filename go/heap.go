package program

// IntHeap heap []int
type IntHeap []int

// Len Len
func (h IntHeap) Len() int { return len(h) }

// Less Less
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }

// Swap Swap
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push Push
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop Pop
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
