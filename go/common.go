package program

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}
	return fmt.Sprintf("%d %v", l.Val, l.Next)
}

func GenerateListNode(list []int) *ListNode {
	if list == nil {
		return nil
	}
	head := &ListNode{}
	root := head

	for _, v := range list {
		root.Next = &ListNode{Val: v}
		root = root.Next
	}
	return head.Next
}

func CheckListEqual(left, right *ListNode) bool {
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}

	return true
}
