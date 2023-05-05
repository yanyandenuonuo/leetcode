/*
 * @lc app=leetcode.cn id=566 lang=golang
 *
 * [566] 重塑矩阵
 */

// @lc code=start
func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat) == 0 || len(mat)*len(mat[0]) != r*c {
		return mat
	}
	newMat := make([][]int, r)

	for rowIdx := 0; rowIdx < len(mat); rowIdx += 1 {
		for columnIdx := 0; columnIdx < len(mat[rowIdx]); columnIdx += 1 {
			idx := rowIdx*len(mat[rowIdx]) + columnIdx
			if idx%c == 0 {
				newMat[idx/c] = make([]int, c)
			}
			newMat[idx/c][idx%c] = mat[rowIdx][columnIdx]
		}
	}

	return newMat
}

// @lc code=end

