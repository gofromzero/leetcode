// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/5/26

package reverselist

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func GenList(arr ...int) *ListNode {
	var root = &ListNode{}
	var l = root
	for _, val := range arr {
		l.Next = &ListNode{Val: val}
		l = l.Next

	}
	return root.Next
}
