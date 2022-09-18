/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心下标
 */

// @lc code=start
func pivotIndex(nums []int) int {
	totalSum, leftSum := 0, 0
	for _, val := range nums {
		totalSum += val
	}
	for idx, val := range nums {
		if 2*leftSum+val == totalSum {
			return idx
		}
		leftSum += val
	}
	return -1
}

// @lc code=end
