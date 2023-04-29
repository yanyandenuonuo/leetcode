/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	// solution1: 模拟
	/**
	preRowList := make([]int, 0, rowIndex+1)
	for idx := 0; idx < rowIndex+1; idx += 1 {
		currentRowList := make([]int, idx+1)
		currentRowList[0], currentRowList[idx] = 1, 1
		for rowIdx := 1; rowIdx < idx; rowIdx += 1 {
			currentRowList[rowIdx] = preRowList[rowIdx-1] + preRowList[rowIdx]
		}
		preRowList = currentRowList
	}
	return preRowList
	*/

	// solution2: 利用组合数公式直接推导对应位置的值
	//	 m	   m-1
	// C n = c n 	* (n-m+1)/m
	rowList := make([]int, rowIndex+1)
	rowList[0], rowList[rowIndex] = 1, 1
	for idx := 1; idx < rowIndex; idx += 1 {
		rowList[idx] = rowList[idx-1] * (rowIndex - idx + 1) / idx
	}
	return rowList
}

// @lc code=end

