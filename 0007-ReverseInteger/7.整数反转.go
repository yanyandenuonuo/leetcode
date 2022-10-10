/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	res := 0
	for x != 0 {
		// math.MinInt == -1<<(intSize-1)
		// math.MaxInt == 1<<(intSize-1) - 1
		if res > (1<<31-1)/10 || res < (-1<<31)/10 {
			return 0
		}
		res = res*10 + x%10
		x /= 10
	}

	return res
}

// @lc code=end

