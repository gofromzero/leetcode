// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/1

package middlenode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func GenList(arr []int) *ListNode {
	var root = &ListNode{}
	var l = root
	for _, val := range arr {
		l.Next = &ListNode{Val: val}
		l = l.Next

	}
	return root.Next
}

func EqualListNode(l *ListNode, r *ListNode) bool {
	for l != nil && r != nil {
		if l.Val != r.Val {
			return false
		}
		l = l.Next
		r = r.Next
	}
	if l == nil && r == l {
		return true
	}

	return false
}
