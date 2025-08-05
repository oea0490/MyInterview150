package BinaryTree

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var level []int
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, level)
	}
	return ans
}
