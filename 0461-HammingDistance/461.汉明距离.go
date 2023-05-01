/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	counter := 0
	for num := x ^ y; num != 0; num &= (num - 1) {
		counter += 1
	}
	return counter
}

// @lc code=end

