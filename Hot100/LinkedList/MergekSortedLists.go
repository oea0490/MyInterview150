package LinkedList

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	leftNode := merge(lists, left, mid)
	rightNode := merge(lists, mid+1, right)
	return mergeTwoLists(leftNode, rightNode)
}
