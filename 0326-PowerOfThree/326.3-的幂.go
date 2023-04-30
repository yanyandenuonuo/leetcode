/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3 的幂
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	// solution1: 依次除3
	/**
	for n > 0 && n%3 == 0 {
		n = n / 3
	}
	return n == 1
	*/

	// solution2: 找出有效范围内的最大的3的幂数，查看n是否为其因数即可
	//			  3^19 = 1162261467
	//			  3^20 > 2^31 -1
	return n > 0 && 1162261467%n == 0
}

// @lc code=end

