/*
 * @lc app=leetcode.cn id=908 lang=golang
 *
 * [908] 最小差值 I
 */

// @lc code=start
func smallestRangeI(nums []int, k int) int {
	// solution: 找最大值和最小值
	minVal, maxVal := 1<<31-1, -1<<31
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}

		if num > maxVal {
			maxVal = num
		}
	}

	if maxVal-minVal <= 2*k {
		return 0
	}

	return maxVal - minVal - 2*k
}

// @lc code=end

