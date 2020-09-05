package program

type LRUCache struct {
	m          map[int]*LinkNode
	cap        int
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{key: 0, val: 0, pre: nil, next: nil}
	tail := &LinkNode{key: 0, val: 0, pre: nil, next: nil}
	head.next = tail
	tail.pre = head

	return LRUCache{
		m:    make(map[int]*LinkNode, capacity),
		cap:  capacity,
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) AddNode(node *LinkNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.RemoveNode(v)
		this.AddNode(v)
		return v.val
	}
	return -1
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

func (this *LRUCache) Put(key int, value int) {
	if v, exist := this.m[key]; exist {
		v.val = value
		this.MoveToHead(v)
		return
	}
	node := &LinkNode{
		key: key,
		val: value,
	}
	if len(this.m) == this.cap {
		delete(this.m, this.tail.pre.key)
		this.RemoveNode(this.tail.pre)
	}

	this.m[key] = node
	this.AddNode(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
