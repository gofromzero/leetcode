// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/9

package mergetwolists

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result = &ListNode{}
	root := result
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			root.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			root.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		root = root.Next
	}

	if list1 == nil {
		root.Next = list2
	} else {
		root.Next = list1
	}

	return result.Next
}
