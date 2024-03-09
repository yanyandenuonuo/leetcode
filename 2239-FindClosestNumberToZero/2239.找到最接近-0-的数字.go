/*
 * @lc app=leetcode.cn id=2239 lang=golang
 *
 * [2239] 找到最接近 0 的数字
 */

// @lc code=start
func findClosestNumber(nums []int) int {
	minDiff, maxNum := 1<<31-1, -1<<31
	for _, num := range nums {
		positiveNum := num
		if positiveNum < 0 {
			positiveNum = -positiveNum
		}

		if positiveNum < minDiff {
			minDiff = positiveNum
			maxNum = num
		} else if positiveNum == minDiff && num > maxNum {
			maxNum = num
		}
	}

	return maxNum
}

// @lc code=end

