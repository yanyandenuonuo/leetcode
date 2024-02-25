/*
 * @lc app=leetcode.cn id=1137 lang=golang
 *
 * [1137] 第 N 个泰波那契数
 */

// @lc code=start
func tribonacci(n int) int {
	// solution1: 递归

	/**
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		return tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)
	}
	*/

	// solution2: 正向计算
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		pre3, pre2, pre1 := 0, 1, 1
		for idx := 3; idx < n; idx += 1 {
			pre3, pre2, pre1 = pre2, pre1, pre3+pre2+pre1
		}

		return pre3 + pre2 + pre1
	}

	// solution3: 矩阵快速幂，时间复杂度O(logn)，空间复杂度O(1)
}

// @lc code=end

