/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	memory := make([][]bool, len(board))
	for rowIdx := range memory {
		memory[rowIdx] = make([]bool, len(board[rowIdx]))
	}

	searchDirections := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	var findWord func(rowIdx, columnIdx, wordIdx int) bool
	findWord = func(rowIdx, columnIdx, wordIdx int) bool {
		if board[rowIdx][columnIdx] != word[wordIdx] {
			return false
		}

		if wordIdx == len(word)-1 {
			return true
		}

		if memory[rowIdx][columnIdx] {
			return false
		}

		memory[rowIdx][columnIdx] = true
		defer func() {
			memory[rowIdx][columnIdx] = false
		}()

		for _, direction := range searchDirections {
			nextRowIdx := rowIdx + direction[0]
			nextColumnIdx := columnIdx + direction[1]

			if nextRowIdx < 0 || nextRowIdx >= len(board) ||
				nextColumnIdx < 0 || nextColumnIdx >= len(board[0]) || memory[nextRowIdx][nextColumnIdx] {
				continue
			}

			if findWord(nextRowIdx, nextColumnIdx, wordIdx+1) {
				return true
			}
		}
		return false
	}

	for rowIdx := range board {
		for columnIdx := range board[rowIdx] {
			if findWord(rowIdx, columnIdx, 0) {
				return true
			}
		}
	}

	return false
}

// @lc code=end

