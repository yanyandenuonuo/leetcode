/*
 * @lc app=leetcode.cn id=766 lang=golang
 *
 * [766] 托普利茨矩阵
 */

// @lc code=start
func isToeplitzMatrix(matrix [][]int) bool {
	// solution1: 比较对角线元素
	// solution2: 比较相邻元素，选择与右下元素或左上元素对比，如果选择右下则每次需要判断边界，可以选择左上，起点设置为[1, 1]即可
	for rowIdx := 1; rowIdx < len(matrix); rowIdx += 1 {
		for columnIdx := 1; columnIdx < len(matrix[rowIdx]); columnIdx += 1 {
			if matrix[rowIdx][columnIdx] != matrix[rowIdx-1][columnIdx-1] {
				return false
			}
		}
	}

	return true
}

// @lc code=end

