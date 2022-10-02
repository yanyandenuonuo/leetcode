/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start
type LRUCache struct {
	size     int
	capacity int
	keyMap   map[int]*DLinkedNode
	head     *DLinkedNode
	tail     *DLinkedNode
}

type DLinkedNode struct {
	key   int
	value int
	prev  *DLinkedNode
	next  *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		size:     0,
		capacity: capacity,
		keyMap:   map[int]*DLinkedNode{},
		head: &DLinkedNode{
			key:   0,
			value: 0,
			prev:  nil,
			next:  nil,
		},
		tail: &DLinkedNode{
			key:   0,
			value: 0,
			prev:  nil,
			next:  nil,
		},
	}
	lruCache.head.next = lruCache.tail
	lruCache.tail.prev = lruCache.head
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if keyNode, isExist := this.keyMap[key]; isExist {
		// move to first
		keyNode.prev.next = keyNode.next
		keyNode.next.prev = keyNode.prev

		this.head.next.prev = keyNode
		keyNode.next = this.head.next
		keyNode.prev = this.head
		this.head.next = keyNode

		return keyNode.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if keyNode, isExist := this.keyMap[key]; isExist {
		// update value
		keyNode.value = value

		// move to first
		keyNode.prev.next = keyNode.next
		keyNode.next.prev = keyNode.prev

		this.head.next.prev = keyNode
		keyNode.next = this.head.next
		keyNode.prev = this.head
		this.head.next = keyNode

		return
	}

	if this.size == this.capacity {
		// delete one
		delete(this.keyMap, this.tail.prev.key)

		this.tail.prev = this.tail.prev.prev
		this.tail.prev.next = this.tail

		this.size -= 1
	}
	keyNode := &DLinkedNode{
		key:   key,
		value: value,
		prev:  this.head,
		next:  this.head.next,
	}
	keyNode.next.prev = keyNode
	this.head.next = keyNode

	this.keyMap[key] = keyNode
	this.size += 1
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

