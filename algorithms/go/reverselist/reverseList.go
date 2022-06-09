// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/9

package reverselist

func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	var root = head
	for root != nil {
		temp := root.Next
		root.Next = result
		result = root
		root = temp
	}
	return result
}
