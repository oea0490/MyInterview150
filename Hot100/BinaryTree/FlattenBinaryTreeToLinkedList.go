package BinaryTree

var preNode *TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	root.Left = nil
	root.Right = preNode
	preNode = root
}
