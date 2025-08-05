package BinaryTree

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ans := pathSumDfs(root, targetSum)
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)
	return ans
}

func pathSumDfs(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == targetSum {
		res++
	}
	res += pathSumDfs(root.Left, targetSum-root.Val)
	res += pathSumDfs(root.Right, targetSum-root.Val)
	return res
}
