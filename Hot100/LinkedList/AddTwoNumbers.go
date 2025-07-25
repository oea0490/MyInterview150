package LinkedList

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := InitListNode()
	carry := 0
	cur := dummy
	for l1 != nil || l2 != nil || carry != 0 {
		val := carry
		if l1 != nil {
			val += l1.Val
		}
		if l2 != nil {
			val += l2.Val
		}
		cur.Next = InitListNodeWithVal(val % 10)
		carry = val / 10
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return dummy.Next
}
