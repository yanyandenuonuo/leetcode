/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	// solution1: use extra matrix store data
	//			  row0 => column(length-1)
	//			  row1 => column(length-2)
	//			  ...
	//			  row(length-1) => column(0)
	// solution2: use two times reverse replace rotate
	//			  reverse by horizontal
	//			  reverse by vertical
	// solution3: rotate relative four point every time
	//			  A  B	=> D  A
	//			  D  C	=> C  B
	//
	//			  1   2   3   4		13          1		13  9       1
	//			  5   6   7   8							            2
	//			  9	  10  11  12						15
	//			  13  14  15  16	16          4		16      8   4
	for rowIdx := 0; rowIdx < len(matrix)/2; rowIdx += 1 {
		for columnIdx := 0; columnIdx < len(matrix)/2+len(matrix)%2; columnIdx += 1 {
			matrix[rowIdx][columnIdx], matrix[columnIdx][len(matrix)-1-rowIdx], matrix[len(matrix)-1-rowIdx][len(matrix)-1-columnIdx], matrix[len(matrix)-1-columnIdx][rowIdx] =
				matrix[len(matrix)-1-columnIdx][rowIdx], matrix[rowIdx][columnIdx], matrix[columnIdx][len(matrix)-1-rowIdx], matrix[len(matrix)-1-rowIdx][len(matrix)-1-columnIdx]
		}
	}
}

// @lc code=end

