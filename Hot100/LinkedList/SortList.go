package LinkedList

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 断开链表
	right := slow.Next
	slow.Next = nil
	left := head
	left = sortList(left)
	right = sortList(right)
	return mergeTwoLists(left, right)
}
