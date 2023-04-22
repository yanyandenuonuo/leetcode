/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	rowNumList, columnNumList, sudokuNumList := [9][10]bool{}, [9][10]bool{}, [3][3][10]bool{}
	for rowIdx := range board {
		for columnIdx := range board[rowIdx] {
			if board[rowIdx][columnIdx] < '0' || board[rowIdx][columnIdx] > '9' {
				continue
			}
			num := board[rowIdx][columnIdx] - '0'
			if rowNumList[rowIdx][num] {
				return false
			}
			rowNumList[rowIdx][num] = true

			if columnNumList[columnIdx][num] {
				return false
			}
			columnNumList[columnIdx][num] = true

			if sudokuNumList[rowIdx/3][columnIdx/3][num] {
				return false
			}
			sudokuNumList[rowIdx/3][columnIdx/3][num] = true
		}
	}
	return true
}

// @lc code=end

