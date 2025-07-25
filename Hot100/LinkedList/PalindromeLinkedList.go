package LinkedList

import "slices"

func isPalindrome(head *ListNode) bool {
	list := []byte{}
	cur := head
	for cur != nil {
		list = append(list, byte(cur.Val))
		cur = cur.Next
	}
	originList := slices.Clone(list)
	slices.Reverse(list)
	return string(list) == string(originList)
}
