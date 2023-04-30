/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */

// @lc code=start
func getSum(a int, b int) int {
	// solution: 无进位场景下 a + b = a ^ b, 需要进位的值为(a & b)<<1
	for b != 0 {
		bitCarry := uint(a&b) << 1
		a ^= b
		b = int(bitCarry)
	}

	return a
}

// @lc code=end

