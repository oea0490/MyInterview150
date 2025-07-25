package LinkedList

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		nextStep := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextStep
	}
	return pre
}
