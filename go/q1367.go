package program

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSubPaths(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func isSubPaths(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if root.Val != head.Val {
		return false
	}

	return isSubPaths(head.Next, root.Left) || isSubPaths(head.Next, root.Right)
}
