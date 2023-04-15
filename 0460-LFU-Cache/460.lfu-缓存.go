/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 */

// @lc code=start
type LFUCache struct {
	capacity     int
	minFrequency int
	keyMap       map[int]*FrequencyLinkedNode
	frequencyMap map[int]*FrequencyLinkedNode
}

type FrequencyLinkedNode struct {
	keyLRUMap      map[int]*LinkedNode
	headKeyLRUList *LinkedNode
	tailKeyLRUList *LinkedNode
}

type LinkedNode struct {
	key       int
	val       int
	frequency int
	prev      *LinkedNode
	next      *LinkedNode
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:     capacity,
		minFrequency: 0,
		keyMap:       make(map[int]*FrequencyLinkedNode, capacity),
		frequencyMap: make(map[int]*FrequencyLinkedNode, capacity),
	}
}

func (this *LFUCache) Get(key int) int {
	if this.keyMap[key] == nil {
		return -1
	}
	// update frequency
	linkedNode := this.popLinkedNode(key)
	if this.frequencyMap[this.minFrequency] == nil {
		this.minFrequency = linkedNode.frequency + 1
	}
	linkedNode.frequency += 1
	this.updateFrequencyNode(linkedNode)

	return this.keyMap[key].keyLRUMap[key].val
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	// update
	if this.keyMap[key] != nil {
		linkedNode := this.popLinkedNode(key)
		if this.frequencyMap[this.minFrequency] == nil {
			this.minFrequency = linkedNode.frequency + 1
		}
		linkedNode.val = value
		linkedNode.frequency += 1
		this.updateFrequencyNode(linkedNode)
		return
	}

	// add
	if len(this.keyMap) == this.capacity {
		// delete lfu and lru key
		deleteKey := this.frequencyMap[this.minFrequency].headKeyLRUList.key
		this.popLinkedNode(deleteKey)
		delete(this.keyMap, deleteKey)
	}

	this.minFrequency = 1

	linkedNode := &LinkedNode{
		key:       key,
		val:       value,
		frequency: 1,
		prev:      nil,
		next:      nil,
	}
	this.updateFrequencyNode(linkedNode)
}

func (this *LFUCache) popLinkedNode(key int) *LinkedNode {
	// update frequency
	frequencyLinkedNode := this.keyMap[key]

	linkedNode := frequencyLinkedNode.keyLRUMap[key]
	frequency := linkedNode.frequency

	if linkedNode == frequencyLinkedNode.headKeyLRUList && linkedNode == frequencyLinkedNode.tailKeyLRUList {
		frequencyLinkedNode.headKeyLRUList = nil
		frequencyLinkedNode.tailKeyLRUList = nil
	} else if linkedNode == frequencyLinkedNode.headKeyLRUList && linkedNode != frequencyLinkedNode.tailKeyLRUList {
		frequencyLinkedNode.headKeyLRUList = frequencyLinkedNode.headKeyLRUList.next
		frequencyLinkedNode.headKeyLRUList.prev = nil
	} else if linkedNode != frequencyLinkedNode.headKeyLRUList && linkedNode == frequencyLinkedNode.tailKeyLRUList {
		frequencyLinkedNode.tailKeyLRUList = frequencyLinkedNode.tailKeyLRUList.prev
		frequencyLinkedNode.tailKeyLRUList.next = nil
	} else {
		linkedNode.prev.next = linkedNode.next
		linkedNode.next.prev = linkedNode.prev
	}

	delete(frequencyLinkedNode.keyLRUMap, key)

	// 是否清除frequencyLinkedNode节点
	if len(frequencyLinkedNode.keyLRUMap) == 0 {
		delete(this.frequencyMap, frequency)
	}
	return linkedNode
}

func (this *LFUCache) updateFrequencyNode(linkedNode *LinkedNode) {
	if this.minFrequency > linkedNode.frequency {
		this.minFrequency = linkedNode.frequency
	}

	var frequencyLinkedNode *FrequencyLinkedNode
	if this.frequencyMap[linkedNode.frequency] == nil {
		// add frequency
		frequencyLinkedNode = &FrequencyLinkedNode{
			keyLRUMap: map[int]*LinkedNode{
				linkedNode.key: linkedNode,
			},
			headKeyLRUList: linkedNode,
			tailKeyLRUList: linkedNode,
		}
	} else {
		// update frequency
		frequencyLinkedNode = this.frequencyMap[linkedNode.frequency]
		frequencyLinkedNode.keyLRUMap[linkedNode.key] = linkedNode

		frequencyLinkedNode.tailKeyLRUList.next = linkedNode
		linkedNode.prev = frequencyLinkedNode.tailKeyLRUList

		frequencyLinkedNode.tailKeyLRUList = linkedNode
	}
	this.keyMap[linkedNode.key] = frequencyLinkedNode
	this.frequencyMap[linkedNode.frequency] = frequencyLinkedNode
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

