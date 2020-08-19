package program

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	temp := &ListNode{Next: head}
	root := head
	old := temp
	for root.Next != nil {
		if n > 0 {
			n--
		}
		if n == 0 {
			old = old.Next
		}

		root = root.Next
	}
	if n <= 1 {
		old.Next = old.Next.Next
	}
	return temp.Next
}
