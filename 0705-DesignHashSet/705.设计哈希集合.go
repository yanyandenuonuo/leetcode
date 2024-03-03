/*
 * @lc app=leetcode.cn id=705 lang=golang
 *
 * [705] 设计哈希集合
 */

// @lc code=start
type MyHashSet struct {
	hashBase int
	hashSet  [][]int
}

func Constructor() MyHashSet {
	return MyHashSet{
		hashBase: 997,
		hashSet:  make([][]int, 997),
	}
}

func (this *MyHashSet) Add(key int) {
	hashIdx := key % this.hashBase

	for _, val := range this.hashSet[hashIdx] {
		if val == key {
			return
		}
	}

	this.hashSet[hashIdx] = append(this.hashSet[hashIdx], key)
}

func (this *MyHashSet) Remove(key int) {
	hashIdx := key % this.hashBase

	hashList := make([]int, 0, len(this.hashSet[hashIdx]))
	for _, val := range this.hashSet[hashIdx] {
		if val == key {
			continue
		}

		hashList = append(hashList, val)
	}

	this.hashSet[hashIdx] = hashList
}

func (this *MyHashSet) Contains(key int) bool {
	hashIdx := key % this.hashBase

	for _, val := range this.hashSet[hashIdx] {
		if val == key {
			return true
		}
	}

	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end

