package BinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initTreeNodeWithValue(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func initTreeNode() *TreeNode {
	return &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
}
