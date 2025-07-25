package LinkedList

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	slowPrev, slow, fast := &ListNode{}, dummy, head
	for i := 1; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slowPrev = slow
		slow = slow.Next
	}
	slowPrev.Next = slow.Next
	return dummy.Next
}
