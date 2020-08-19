package program

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

type listNode struct {
	Val  int
	Next *listNode
}

func generateListNode(list []int) *listNode {
	root := &listNode{}
	head := root
	for _, v := range list {
		root.Val = v
		root.Next = &listNode{}
		root = root.Next
	}
	return head
}

func checkListEqual(left, right *listNode) bool {
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
