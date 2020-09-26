package program

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return string(leafSimilarCheck(root1)) == string(leafSimilarCheck(root2))
}

func leafSimilarCheck(root *TreeNode) []byte {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []byte{byte(root.Val + 48)}
	}

	result := leafSimilarCheck(root.Left)
	result = append(result, leafSimilarCheck(root.Right)...)

	return result
}
