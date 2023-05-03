/*
 * @lc app=leetcode.cn id=867 lang=golang
 *
 * [867] 转置矩阵
 */

// @lc code=start
func transpose(matrix [][]int) [][]int {
	// solution: 模拟，因为row和column不一定相等，所以无法原地转
	newMatrix := make([][]int, len(matrix[0]))
	for rowIdx := 0; rowIdx < len(matrix); rowIdx += 1 {
		for columnIdx := 0; columnIdx < len(matrix[rowIdx]); columnIdx += 1 {
			if rowIdx == 0 {
				newMatrix[columnIdx] = make([]int, len(matrix))
			}
			newMatrix[columnIdx][rowIdx] = matrix[rowIdx][columnIdx]
		}
	}

	return newMatrix
}

// @lc code=end

