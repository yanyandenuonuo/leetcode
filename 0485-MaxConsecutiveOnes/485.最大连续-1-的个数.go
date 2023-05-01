/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续 1 的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	for idx, counter := 0, 0; idx < len(nums); idx += 1 {
		if nums[idx] == 0 {
			counter = 0
			continue
		}

		counter += 1

		if counter > max {
			max = counter
		}
	}
	return max
}

// @lc code=end

