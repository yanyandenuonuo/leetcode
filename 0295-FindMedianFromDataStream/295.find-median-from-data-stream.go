/*
 * @lc app=leetcode id=295 lang=golang
 *
 * [295] Find Median from Data Stream
 *
 * https://leetcode.com/problems/find-median-from-data-stream/description/
 *
 * algorithms
 * Hard (45.26%)
 * Likes:    3150
 * Dislikes: 59
 * Total Accepted:    234.1K
 * Total Submissions: 517.1K
 * Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n' +
  '[[],[1],[2],[],[3],[]]'
 *
 * Median is the middle value in an ordered integer list. If the size of the
 * list is even, there is no middle value. So the median is the mean of the two
 * middle value.
 * For example,
 *
 * [2,3,4], the median is 3
 *
 * [2,3], the median is (2 + 3) / 2 = 2.5
 *
 * Design a data structure that supports the following two operations:
 *
 *
 * void addNum(int num) - Add a integer number from the data stream to the data
 * structure.
 * double findMedian() - Return the median of all elements so far.
 *
 *
 *
 *
 * Example:
 *
 *
 * addNum(1)
 * addNum(2)
 * findMedian() -> 1.5
 * addNum(3)
 * findMedian() -> 2
 *
 *
 *
 *
 * Follow up:
 *
 *
 * If all integer numbers from the stream are between 0 and 100, how would you
 * optimize it?
 * If 99% of all integer numbers from the stream are between 0 and 100, how
 * would you optimize it?
 *
 *
*/

// @lc code=start
type hp struct {
	sort.IntSlice
}

func (h *hp) Push(val interface{}) {
	h.IntSlice = append(h.IntSlice, val.(int))
}

func (h *hp) Pop() interface{} {
	val := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return val
}

type MedianFinder struct {
	// solution1: 直接使用数组，每次读取时进行排序
	//			  每次插入时间复杂度为O(1)，读取为O(nlog(n))，空间复杂度为O(n)
	// solution2: 利用有序集合（实质为红黑树）+双指针
	//			  每次插入时间复杂度为O(log(n))，读取为O(1)，空间复杂度为O(n)
	// solution3: 利用优先队列（实质为小顶堆），Pop弹出的为最小元素
	// 			  minQueue采用负数保存，以实现大顶堆的效果
	//			  优先把数据放入minQueue，两者一直保持为长度相等或len(minQueue)-len(maxQueue)=1
	//			  每次插入时间复杂度为O(log(n))，读取为O(1)，空间复杂度为O(n)
	maxQueue hp
	minQueue hp
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minQueue.Len() == 0 || num <= -this.minQueue.IntSlice[0] {
		heap.Push(&this.minQueue, -num)
		if this.maxQueue.Len()+1 < this.minQueue.Len() {
			heap.Push(&this.maxQueue, -heap.Pop(&this.minQueue).(int))
		}
	} else {
		heap.Push(&this.maxQueue, num)
		if this.maxQueue.Len() > this.minQueue.Len() {
			heap.Push(&this.minQueue, -heap.Pop(&this.maxQueue).(int))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minQueue.Len() > this.maxQueue.Len() {
		return float64(-this.minQueue.IntSlice[0])
	}
	return float64(this.maxQueue.IntSlice[0]-this.minQueue.IntSlice[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

