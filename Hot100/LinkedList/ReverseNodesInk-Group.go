package LinkedList

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil {
		start := pre.Next
		end := start
		for i := 1; i < k; i++ {
			end = end.Next
			if end == nil {
				return dummy.Next
			}
		}
		next := end.Next
		end.Next = nil
		// prev -> end -> start -> next
		pre.Next = reverseList(start)
		start.Next = next
		pre = start
	}
	return dummy.Next
}
