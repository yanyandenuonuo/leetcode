/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	// solution1: 依次除4
	/**
	for n > 0 && n%4 == 0 {
		n = n / 4
	}
	return n == 1
	*/

	// solution2: 二进制中1的位置，若是4的幂则一定是(2^2)^x，则一定只能有1个1且1在偶数位(从0位开始计数)
	// 构造1个mask，基数bit为1，偶数bit为0
	/**
	const mask = 0b10101010101010101010101010101010
	return n > 0 && n&(n-1) == 0 && n&mask == 0
	*/

	// solution3: 取模，n = 4^x = (3+1)^x
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}

// @lc code=end

