/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	// solution1: 模拟四向
	total := len(matrix) * len(matrix[0])
	res := make([]int, 0, total)
	minRow, maxRow, minColumn, maxColumn := 0, len(matrix)-1, 0, len(matrix[0])-1
	for total > 0 {
		// top
		for rowIdx, columnIdx := minRow, minColumn; columnIdx <= maxColumn && total > 0; columnIdx += 1 {
			res = append(res, matrix[rowIdx][columnIdx])
			total -= 1
		}
		minRow += 1

		// right
		for rowIdx, columnIdx := minRow, maxColumn; rowIdx <= maxRow && total > 0; rowIdx += 1 {
			res = append(res, matrix[rowIdx][columnIdx])
			total -= 1
		}
		maxColumn -= 1

		// bottom
		for rowIdx, columnIdx := maxRow, maxColumn; columnIdx >= minColumn && total > 0; columnIdx -= 1 {
			res = append(res, matrix[rowIdx][columnIdx])
			total -= 1
		}
		maxRow -= 1

		// left
		for rowIdx, columnIdx := maxRow, minColumn; rowIdx >= minRow && total > 0; rowIdx -= 1 {
			res = append(res, matrix[rowIdx][columnIdx])
			total -= 1
		}
		minColumn += 1
	}

	return res

	// solution2: 关注点的边界，存在bad case需要特殊处理，如只有一行时上下顶点重合，只有1列时左右顶点重合
	/**
	total := len(matrix) * len(matrix[0])
	res := make([]int, 0, total)
	minRow, maxRow, minColumn, maxColumn := 0, len(matrix)-1, 0, len(matrix[0])-1
	// rowOffset, columnOffset
	//  0,  1 从左到右
	//  1,  0 从上到下
	//  0, -1 从右到左
	// -1,  0 从下到上
	rowOffset, columnOffset := 0, 1 // 初始为从左到右
	for rowIdx, columnIdx := minRow, minColumn; total > 0; total -= 1 {
		res = append(res, matrix[rowIdx][columnIdx])
		rowIdx, columnIdx = rowIdx+rowOffset, columnIdx+columnOffset

		// bad case: 只存在1列的场景
		if columnIdx > maxColumn {
			minRow += 1
			rowIdx, columnIdx = minRow, minColumn
			continue
		}

		if rowIdx == minRow && columnIdx == minColumn {
			// 左上顶点
			rowOffset, columnOffset = 0, 1

			// 左上顶点时重置区域并重置rowIdx, columnIdx
			if minRow+1 <= maxRow-1 {
				minRow += 1
				maxRow -= 1
			} else if minRow+1 <= maxRow {
				minRow += 1
			}
			if minColumn+1 <= maxColumn-1 {
				minColumn += 1
				maxColumn -= 1
			} else if minColumn+1 <= maxColumn {
				minColumn += 1
			}

			rowIdx, columnIdx = minRow, minColumn
		} else if rowIdx == minRow && columnIdx == maxColumn {
			// 右上顶点
			rowOffset, columnOffset = 1, 0
		} else if rowIdx == maxRow && columnIdx == maxColumn {
			// 右下顶点
			rowOffset, columnOffset = 0, -1
		} else if rowIdx == maxRow && columnIdx == minColumn {
			// 左下顶点
			rowOffset, columnOffset = -1, 0
		}
	}

	return res
	*/
}

// @lc code=end

