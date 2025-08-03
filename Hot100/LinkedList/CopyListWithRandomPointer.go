package LinkedList

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	dummy := &Node{
		Val: -1,
	}
	// 记录原链表节点和索引
	NtoImap := make(map[*Node]int)
	// 记录新链表索引和节点
	ImapToN := make(map[int]*Node)
	cur := head
	curCopy := dummy
	index := 0
	// 一次遍历创建新链表
	for cur != nil {
		newCopyNode := &Node{
			Val: cur.Val,
		}
		curCopy.Next = newCopyNode
		NtoImap[cur] = index
		ImapToN[index] = newCopyNode
		cur = cur.Next
		curCopy = curCopy.Next
		index++
	}
	// 二次遍历更新随机指针
	cur = head
	curCopy = dummy.Next
	for cur != nil {
		if cur.Random != nil {
			randomIndex := NtoImap[cur.Random]
			curCopy.Random = ImapToN[randomIndex]
		}
		cur = cur.Next
		curCopy = curCopy.Next
	}
	return dummy.Next
}
