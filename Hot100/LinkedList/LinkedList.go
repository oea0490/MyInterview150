package LinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitListNodeWithValue(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func InitListNode() *ListNode {
	return &ListNode{
		Val:  0,
		Next: nil,
	}
}
