package BinaryTree

var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	diameter = 0
	leftDepth := diameterDepth(root.Left)
	rightDepth := diameterDepth(root.Right)
	diameter = max(diameter, leftDepth+rightDepth)
	return diameter
}

func diameterDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := diameterDepth(root.Left)
	rightDepth := diameterDepth(root.Right)
	diameter = max(diameter, leftDepth+rightDepth)
	return max(leftDepth, rightDepth) + 1
}
