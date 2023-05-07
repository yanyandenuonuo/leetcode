/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	// solution1: 记录原始0的位置，然后基于0的位置清空
	// solution2: 使用row[0]和column[0]标记该行列是否存在0，然后用两个额外的变量记录0行和0列是否存在0
	// solution3: 使用row[0]和column[0]标记该行列是否存在0，然后用一个额外的变量记录0行是否存在0，优化点在[0, 0]
	clearRow0 := false

	for rowIdx := 0; rowIdx < len(matrix); rowIdx += 1 {
		if matrix[rowIdx][0] == 0 {
			clearRow0 = true
		}

		for columnIdx := 0; columnIdx < len(matrix[rowIdx]); columnIdx += 1 {
			if matrix[rowIdx][columnIdx] == 0 {
				// 清零该行
				matrix[rowIdx][0] = 0
				if columnIdx != 0 {
					// 清零该列
					matrix[0][columnIdx] = 0
				}
			}
		}
	}

	for rowIdx := len(matrix) - 1; rowIdx >= 0; rowIdx -= 1 {
		for columnIdx := len(matrix[rowIdx]) - 1; columnIdx > 0; columnIdx -= 1 {
			if matrix[0][columnIdx] == 0 || matrix[rowIdx][0] == 0 {
				matrix[rowIdx][columnIdx] = 0
			}
		}

		if clearRow0 {
			matrix[rowIdx][0] = 0
		}
	}
}

// @lc code=end

