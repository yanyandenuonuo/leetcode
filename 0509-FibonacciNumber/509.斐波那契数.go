/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n > 1 {
		preTwo := 0
		preOne := 1
		for idx := 2; idx < n; idx += 1 {
			preTwo, preOne = preOne, preOne+preTwo
		}
		return preOne + preTwo
	} else {
		// n < 0, wrong intput
		return -1
	}
}

// @lc code=end

