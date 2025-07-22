package LinkedList

type LRUCache struct {
	cache    map[int]*DLinkedNode
	head     *DLinkedNode
	tail     *DLinkedNode
	size     int
	capacity int
}

type DLinkedNode struct {
	key  int
	val  int
	prev *DLinkedNode
	next *DLinkedNode
}

func initDLinkedNode(key, val int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		val: val,
	}
}

func Constructor(capacity int) LRUCache {
	LRUCache := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		size:     0,
		capacity: capacity,
	}
	LRUCache.head.next = LRUCache.tail
	LRUCache.tail.prev = LRUCache.head
	return LRUCache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if !ok {
		newNode := initDLinkedNode(key, value)
		this.cache[key] = newNode
		this.addToHead(newNode)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node.val = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
