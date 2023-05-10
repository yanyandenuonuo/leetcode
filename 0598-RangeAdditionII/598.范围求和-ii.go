/*
 * @lc app=leetcode.cn id=598 lang=golang
 *
 * [598] 范围求和 II
 */

// @lc code=start
func maxCount(m int, n int, ops [][]int) int {
	// solution: 找出最小的row和column
	minRow, minColumn := m-1, n-1
	for _, pos := range ops {
		if pos[0] > 0 && pos[0]-1 < minRow {
			minRow = pos[0] - 1
		}

		if pos[1] > 0 && pos[1]-1 < minColumn {
			minColumn = pos[1] - 1
		}
	}

	return (minRow + 1) * (minColumn + 1)
}

// @lc code=end

