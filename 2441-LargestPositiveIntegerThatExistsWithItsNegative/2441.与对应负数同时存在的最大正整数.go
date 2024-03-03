/*
 * @lc app=leetcode.cn id=2441 lang=golang
 *
 * [2441] 与对应负数同时存在的最大正整数
 */

// @lc code=start
func findMaxK(nums []int) int {
	// solution1: 暴力枚举
	// solution2: hash表
	numMap := make(map[int]int, len(nums))
	maxVal := -1
	for _, val := range nums {
		if _, isExist := numMap[-val]; !isExist {
			numMap[val] = 1
			continue
		}
		if val < 0 {
			val = -val
		}
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal

	// solution3: 排序+双指针
}

// @lc code=end

