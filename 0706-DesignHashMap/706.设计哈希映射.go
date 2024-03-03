/*
 * @lc app=leetcode.cn id=706 lang=golang
 *
 * [706] 设计哈希映射
 */

// @lc code=start

type HashValue struct {
	key   int
	value int
}

type MyHashMap struct {
	hashBase int
	hashList [][]HashValue
}

func Constructor() MyHashMap {
	return MyHashMap{
		hashBase: 997,
		hashList: make([][]HashValue, 997),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	hashIdx := key % this.hashBase

	for idx, hashVal := range this.hashList[hashIdx] {
		if hashVal.key == key {
			this.hashList[hashIdx][idx].value = value
			return
		}
	}

	this.hashList[hashIdx] = append(this.hashList[hashIdx], HashValue{
		key:   key,
		value: value,
	})
}

func (this *MyHashMap) Get(key int) int {
	hashIdx := key % this.hashBase

	for _, hashVal := range this.hashList[hashIdx] {
		if hashVal.key == key {
			return hashVal.value
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	hashIdx := key % this.hashBase

	hashList := make([]HashValue, 0, len(this.hashList[hashIdx]))

	for _, hashVal := range this.hashList[hashIdx] {
		if hashVal.key == key {
			continue
		}

		hashList = append(hashList, HashValue{
			key:   hashVal.key,
			value: hashVal.value,
		})
	}

	this.hashList[hashIdx] = hashList
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end

