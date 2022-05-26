// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/5/26

package addTwoNumbers

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var left int
	var root ListNode
	var l = &root

	for l1 != nil || l2 != nil || left != 0 {
		if l1 != nil {
			left += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			left += l2.Val
			l2 = l2.Next
		}

		l.Next = &ListNode{Val: left % 10}
		l = l.Next
		left /= 10
	}

	return root.Next
}
