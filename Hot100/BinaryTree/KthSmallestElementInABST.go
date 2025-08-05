package BinaryTree

var (
	index int
	ans   int
)

func kthSmallest(root *TreeNode, k int) int {
	index = 0
	ans = 0
	kthSmallestInorder(root, k)
	return ans
}

func kthSmallestInorder(root *TreeNode, k int) {
	if root == nil {
		return
	}
	kthSmallestInorder(root.Left, k)
	index++
	if index == k {
		ans = root.Val
		return
	}
	kthSmallestInorder(root.Right, k)
}
