/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	maxLCIS := 1
	preLCIS := 1
	for idx := 1; idx < len(nums); idx += 1 {
		if nums[idx] > nums[idx-1] {
			preLCIS += 1
			if preLCIS > maxLCIS {
				maxLCIS = preLCIS
			}
		} else {
			preLCIS = 1
		}
	}
	return maxLCIS
}

// @lc code=end

