/*
 * @lc app=leetcode.cn id=561 lang=golang
 *
 * [561] 数组拆分
 */

// @lc code=start
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	maxSum := 0
	for idx := 0; idx < len(nums); idx += 2 {
		maxSum += nums[idx]
	}
	return maxSum
}

// @lc code=end

