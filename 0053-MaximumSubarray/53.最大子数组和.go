/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	maxSum := math.MinInt32
	sum := 0
	for _, val := range nums {
		sum += val
		if maxSum < sum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}

// @lc code=end

