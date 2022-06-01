// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/1

package middlenode

func middleNode(head *ListNode) *ListNode {
	l, r := head, head
	for r.Next != nil && r.Next.Next != nil {
		l = l.Next
		r = r.Next.Next
	}
	if r.Next != nil {
		l = l.Next
	}
	return l
}
