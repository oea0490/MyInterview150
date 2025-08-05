package BinaryTree

var maxSum int

func maxPathSum(root *TreeNode) int {
	maxSum = -1 << 63
	maxPathSumDfs(root)
	return maxSum
}

func maxPathSumDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftScore := maxPathSumDfs(root.Left)
	rightScore := maxPathSumDfs(root.Right)
	maxScore := max(leftScore, rightScore) + root.Val
	maxScore = max(maxScore, root.Val)
	maxSum = max(maxSum, maxScore)
	maxSum = max(maxSum, leftScore+rightScore+root.Val)
	return maxScore
}
