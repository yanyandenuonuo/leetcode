/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	// solution: 利用val作为idx占位
	for _, val := range nums {
		if val < 0 {
			val = -val
		}
		if nums[val-1] < 0 {
			return val
		}
		nums[val-1] = -nums[val-1]
	}
	return -1
}

// @lc code=end

