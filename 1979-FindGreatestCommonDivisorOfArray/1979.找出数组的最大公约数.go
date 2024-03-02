/*
 * @lc app=leetcode.cn id=1979 lang=golang
 *
 * [1979] 找出数组的最大公约数
 */

// @lc code=start
func findGCD(nums []int) int {
	minVal, maxVal := nums[0], nums[0]

	for _, val := range nums {
		if val > maxVal {
			maxVal = val
		} else if val < minVal {
			minVal = val
		}
	}

	for minVal != 0 {
		minVal, maxVal = maxVal%minVal, minVal
	}

	return maxVal
}

// @lc code=end

