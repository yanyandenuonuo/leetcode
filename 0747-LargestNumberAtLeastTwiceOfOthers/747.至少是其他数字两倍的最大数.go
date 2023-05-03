/*
 * @lc app=leetcode.cn id=747 lang=golang
 *
 * [747] 至少是其他数字两倍的最大数
 */

// @lc code=start
func dominantIndex(nums []int) int {
	// solution1: 找出最大值和次大值
	maxIdx, preMaxIdx := -1, -1
	for idx, num := range nums {
		if maxIdx == -1 || num >= nums[maxIdx] {
			maxIdx, preMaxIdx = idx, maxIdx
			continue
		}

		if preMaxIdx == -1 || num > nums[preMaxIdx] {
			preMaxIdx = idx
		}
	}

	if nums[maxIdx] >= 2*nums[preMaxIdx] {
		return maxIdx
	}

	return -1
}

// @lc code=end

