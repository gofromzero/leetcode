package program

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := 0
	root := &ListNode{}
	result := root
	for l1 != nil || l2 != nil || temp != 0 {
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}

		root.Next = &ListNode{
			Val: temp % 10,
		}
		root = root.Next
		temp /= 10
	}

	return result.Next
}
