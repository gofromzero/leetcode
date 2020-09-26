package program

var node *TreeNode

func _inorder(root *TreeNode) {
	if root == nil {
		return
	}
	_inorder(root.Left)
	node.Left = nil
	node.Right = root
	node = root
	_inorder(root.Right)

}

func increasingBST(root *TreeNode) *TreeNode {
	node = &TreeNode{}
	head := node
	_inorder(root)
	return head.Right
}

//
//func increasingBST(root *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	nodes := inOrder(root)
//	var head = &TreeNode{}
//	var result = head
//	for _, v := range nodes {
//		head.Right = &TreeNode{
//			Val: v,
//		}
//		head = head.Right
//	}
//
//	return result.Right
//}
//
//func inOrder(root *TreeNode) []int  {
//	if root == nil {
//		return nil
//	}
//	result := inOrder(root.Left)
//	result = append(result, root.Val)
//	result = append(result, inOrder(root.Right)...)
//	return  result
//}
