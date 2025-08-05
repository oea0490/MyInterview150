package BinaryTree

var inOrderMap map[int]int

func buildTree(preorder []int, inorder []int) *TreeNode {
	inOrderMap = make(map[int]int)
	for i, v := range inorder {
		inOrderMap[v] = i
	}
	return buildTree2(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func buildTree2(preorder []int, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	rootVal := preorder[preStart]
	rootIndex := inOrderMap[rootVal]
	leftLen := rootIndex - inStart
	root := &TreeNode{Val: rootVal}
	root.Left = buildTree2(preorder, inorder, preStart+1, preStart+leftLen, inStart, rootIndex-1)
	root.Right = buildTree2(preorder, inorder, preStart+leftLen+1, preEnd, rootIndex+1, inEnd)
	return root
}
