package BinaryTree

var ancestor *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ancestor = nil
	ancestorDfs(root, p, q)
	return ancestor
}

func ancestorDfs(root, p, q *TreeNode) bool {
	if root == nil {
		return false
	}
	leftSon := ancestorDfs(root.Left, p, q)
	rightSon := ancestorDfs(root.Right, p, q)
	if (leftSon && rightSon) || ((root.Val == p.Val || root.Val == q.Val) && (leftSon || rightSon)) {
		ancestor = root
	}
	return leftSon || rightSon || (root.Val == p.Val || root.Val == q.Val)
}
