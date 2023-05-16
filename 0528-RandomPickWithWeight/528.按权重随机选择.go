/*
 * @lc app=leetcode.cn id=528 lang=golang
 *
 * [528] 按权重随机选择
 */

// @lc code=start
type Solution struct {
	// solution: 前缀和+二分查找
	preSumList []int
}

func Constructor(w []int) Solution {
	preSumList := make([]int, len(w))
	for idx, preSum := 0, 0; idx < len(w); idx += 1 {
		preSum += w[idx]
		preSumList[idx] = preSum
	}

	return Solution{
		preSumList: preSumList,
	}
}

func (this *Solution) PickIndex() int {
	randNum := rand.Intn(this.preSumList[len(this.preSumList)-1]) + 1
	leftIdx := 0
	for rightIdx := len(this.preSumList) - 1; leftIdx < rightIdx; {
		midIdx := leftIdx + (rightIdx-leftIdx)/2
		if this.preSumList[midIdx] > randNum {
			rightIdx = midIdx
		} else if this.preSumList[midIdx] < randNum {
			leftIdx = midIdx + 1
		} else {
			return midIdx
		}
	}

	return leftIdx
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
// @lc code=end

