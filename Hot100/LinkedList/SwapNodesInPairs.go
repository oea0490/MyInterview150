package LinkedList

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		first := pre.Next
		second := pre.Next.Next
		// prev -> second -> first -> second.next
		pre.Next = second
		first.Next = second.Next
		second.Next = first
		pre = first
	}
	return dummy.Next
}
