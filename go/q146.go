package program

// LRUCache LRUCache
type LRUCache struct {
	m          map[int]*LinkNode
	cap        int
	head, tail *LinkNode
}

// Constructor Constructor
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

// RemoveNode RemoveNode
func (c *LRUCache) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

// AddNode AddNode
func (c *LRUCache) AddNode(node *LinkNode) {
	node.pre = c.head
	node.next = c.head.next
	c.head.next.pre = node
	c.head.next = node
}

// Get Get
func (c *LRUCache) Get(key int) int {
	if v, ok := c.m[key]; ok {
		c.RemoveNode(v)
		c.AddNode(v)
		return v.val
	}
	return -1
}

// MoveToHead MoveToHead
func (c *LRUCache) MoveToHead(node *LinkNode) {
	c.RemoveNode(node)
	c.AddNode(node)
}

// Put Put
func (c *LRUCache) Put(key int, value int) {
	if v, exist := c.m[key]; exist {
		v.val = value
		c.MoveToHead(v)
		return
	}
	node := &LinkNode{
		key: key,
		val: value,
	}
	if len(c.m) == c.cap {
		delete(c.m, c.tail.pre.key)
		c.RemoveNode(c.tail.pre)
	}

	c.m[key] = node
	c.AddNode(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
