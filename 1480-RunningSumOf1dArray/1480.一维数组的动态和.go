/*
 * @lc app=leetcode.cn id=1480 lang=golang
 *
 * [1480] 一维数组的动态和
 */

// @lc code=start
func runningSum(nums []int) []int {
	for idx := 1; idx < len(nums); idx += 1 {
		nums[idx] = nums[idx] + nums[idx-1]
	}
	return nums
}

// @lc code=end

