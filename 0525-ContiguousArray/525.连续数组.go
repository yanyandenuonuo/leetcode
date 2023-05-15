/*
 * @lc app=leetcode.cn id=525 lang=golang
 *
 * [525] 连续数组
 */

// @lc code=start
func findMaxLength(nums []int) int {
	// solution1: 滑动窗口
	// solution2: 前缀和+hash，将0看作-1，则题目转化为和为0的子数组最大长度
	maxLength, preSum, sum := 0, make(map[int]int, 0), 0
	preSum[0] = -1
	for idx, num := range nums {
		if num == 0 {
			num = -1
		}

		sum += num

		if minIdx, isExist := preSum[sum]; isExist {
			if idx-minIdx > maxLength {
				maxLength = idx - minIdx
			}
		} else {
			preSum[sum] = idx
		}
	}

	return maxLength
}

// @lc code=end

