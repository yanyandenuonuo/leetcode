/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	// solution1: 模拟四向
	if n == 0 {
		return [][]int{}
	}
	matrix := make([][]int, n)
	for idx := 0; idx < n; idx += 1 {
		matrix[idx] = make([]int, n)
	}

	minRow, maxRow, minColumn, maxColumn := 0, n-1, 0, n-1
	for total := 1; total <= n*n; {
		// top
		for rowIdx, columnIdx := minRow, minColumn; columnIdx <= maxColumn && total <= n*n; columnIdx += 1 {
			matrix[rowIdx][columnIdx] = total
			total += 1
		}
		minRow += 1

		// right
		for rowIdx, columnIdx := minRow, maxColumn; rowIdx <= maxRow && total <= n*n; rowIdx += 1 {
			matrix[rowIdx][columnIdx] = total
			total += 1
		}
		maxColumn -= 1

		// bottom
		for rowIdx, columnIdx := maxRow, maxColumn; columnIdx >= minColumn && total <= n*n; columnIdx -= 1 {
			matrix[rowIdx][columnIdx] = total
			total += 1
		}
		maxRow -= 1

		// left
		for rowIdx, columnIdx := maxRow, minColumn; rowIdx >= minRow && total <= n*n; rowIdx -= 1 {
			matrix[rowIdx][columnIdx] = total
			total += 1
		}
		minColumn += 1
	}

	return matrix
}

// @lc code=end

