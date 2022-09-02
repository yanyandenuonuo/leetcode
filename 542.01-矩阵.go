/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
func updateMatrix(mat [][]int) [][]int {
	queueElement := make([][]int, 0, len(mat)*len(mat[0])-1)

	matrix := make([][]int, len(mat))

	for row := range mat {
		for column := range mat[row] {
			if matrix[row] == nil {
				matrix[row] = make([]int, len(mat[row]))
			}

			if mat[row][column] == 0 {
				matrix[row][column] = 0
				queueElement = append(queueElement, []int{row, column})
				continue
			}

			matrix[row][column] = -1
		}
	}

	for len(queueElement) > 0 {
		row := queueElement[0][0]
		column := queueElement[0][1]

		queueElement = queueElement[1:]

		// up
		if row-1 >= 0 && matrix[row-1][column] == -1 {
			matrix[row-1][column] = matrix[row][column] + 1
			queueElement = append(queueElement, []int{row - 1, column})
		}

		// down
		if row+1 < len(mat) && matrix[row+1][column] == -1 {
			matrix[row+1][column] = matrix[row][column] + 1
			queueElement = append(queueElement, []int{row + 1, column})
		}

		// left
		if column-1 >= 0 && matrix[row][column-1] == -1 {
			matrix[row][column-1] = matrix[row][column] + 1
			queueElement = append(queueElement, []int{row, column - 1})
		}

		// right
		if column+1 < len(mat[row]) && matrix[row][column+1] == -1 {
			matrix[row][column+1] = matrix[row][column] + 1
			queueElement = append(queueElement, []int{row, column + 1})
		}

		// fmt.Printf("row: %v, column: %v, matrix: %v \n", row, column, matrix)
	}

	return matrix
}

// @lc code=end

