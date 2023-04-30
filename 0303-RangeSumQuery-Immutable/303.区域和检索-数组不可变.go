/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	for idx := 1; idx < len(nums); idx += 1 {
		preSum[idx] = preSum[idx-1] + nums[idx]
	}
	return NumArray{
		preSum: preSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > 0 {
		return this.preSum[right] - this.preSum[left-1]

	}
	return this.preSum[right]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

