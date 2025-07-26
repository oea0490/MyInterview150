package BinaryTree

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	inorder(root, &ans)
	return ans
}

func inorder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, ans)
	*ans = append(*ans, root.Val)
	inorder(root.Right, ans)
}
