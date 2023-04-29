/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)
	for idx := 0; idx < numRows; idx += 1 {
		rowList := make([]int, idx+1)
		rowList[0], rowList[idx] = 1, 1
		for rowIdx := 1; rowIdx < idx; rowIdx += 1 {
			rowList[rowIdx] = res[idx-1][rowIdx-1] + res[idx-1][rowIdx]
		}
		res = append(res, rowList)
	}
	return res
}

// @lc code=end

