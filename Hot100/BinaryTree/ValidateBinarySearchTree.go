package BinaryTree

func isValidBST(root *TreeNode) bool {
	min := -1 << 63
	max := 1<<63 - 1
	return isValidBST2(root, min, max)
}

func isValidBST2(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return isValidBST2(root.Left, min, root.Val) && isValidBST2(root.Right, root.Val, max)
}
