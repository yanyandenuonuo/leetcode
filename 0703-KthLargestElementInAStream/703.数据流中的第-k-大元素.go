/*
 * @lc app=leetcode.cn id=703 lang=golang
 *
 * [703] 数据流中的第 K 大元素
 */

// @lc code=start
type KthLargest struct {
	// solution: 优先队列
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := KthLargest{k: k}
	for _, num := range nums {
		kthLargest.Add(num)
	}

	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}

	return this.IntSlice[0]
}

func (this *KthLargest) Push(v interface{}) {
	this.IntSlice = append(this.IntSlice, v.(int))
}

func (this *KthLargest) Pop() interface{} {
	val := this.IntSlice[this.Len()-1]
	this.IntSlice = this.IntSlice[:this.Len()-1]
	return val
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end

