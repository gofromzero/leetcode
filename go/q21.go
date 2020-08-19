package program

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	root := &ListNode{}
	head := root
	for l1 != nil || l2 != nil {
		if l2 == nil {
			*head = *l1
			break
		}
		if l1 == nil {
			*head = *l2
			break
		}
		if l1.Val < l2.Val {
			head.Val = l1.Val
			l1 = l1.Next
		} else {
			head.Val = l2.Val
			l2 = l2.Next
		}

		head.Next = &ListNode{}
		head = head.Next
	}
	return root
}
