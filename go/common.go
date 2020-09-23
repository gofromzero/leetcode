package program

import "fmt"

// TreeNode TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkNode LinkNode
type LinkNode struct {
	key, val  int
	pre, next *LinkNode
}

type Node struct {
	Val      int
	Children []*Node
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}
	return fmt.Sprintf("%d %v", l.Val, l.Next)
}

// GenerateListNode arr to list
func GenerateListNode(list []int) *ListNode {
	if len(list) == 0 {
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

// CheckListEqual check list equal
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
