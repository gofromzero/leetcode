// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/8

package removenthfromend

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	root := &ListNode{
		Val:  0,
		Next: head,
	}
	l, r := root, root
	for r.Next != nil {
		if n == 0 {
			l = l.Next
		} else {
			n--
		}

		r = r.Next
	}

	if n == 0 {
		l.Next = l.Next.Next
	}

	return root.Next
}
