/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2 的幂
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	// solution1: 判断1的数量
	// return n > 0 && n&(n-1) == 0x00

	// solution2: 判断是否为2^31的约数
	return n > 0 && (1<<30)%n == 0
}

// @lc code=end

