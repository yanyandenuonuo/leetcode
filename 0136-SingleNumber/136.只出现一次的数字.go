/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
	ans := nums[0]
	for idx := 1; idx < len(nums); idx += 1 {
		ans ^= nums[idx]
	}
	return ans
}

// @lc code=end

