package program

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if s.Val == t.Val {
		if same(s, t) {
			return true
		}
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func same(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val == t2.Val {
		return same(t1.Left, t2.Left) && same(t1.Right, t2.Right)
	}

	return false
}
